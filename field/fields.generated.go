package field

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/tag"
	"github.com/shopspring/decimal"
	"time"
)

//AccountField is a STRING field
type AccountField struct{ quickfix.FIXString }

//Tag returns tag.Account (1)
func (f AccountField) Tag() quickfix.Tag { return tag.Account }

//NewAccount returns a new AccountField initialized with val
func NewAccount(val string) AccountField {
	return AccountField{quickfix.FIXString(val)}
}

func (f AccountField) Value() string { return f.String() }

//AccountTypeField is a enum.AccountType field
type AccountTypeField struct{ quickfix.FIXString }

//Tag returns tag.AccountType (581)
func (f AccountTypeField) Tag() quickfix.Tag { return tag.AccountType }

func NewAccountType(val enum.AccountType) AccountTypeField {
	return AccountTypeField{quickfix.FIXString(val)}
}

func (f AccountTypeField) Value() enum.AccountType { return enum.AccountType(f.String()) }

//AccruedInterestAmtField is a AMT field
type AccruedInterestAmtField struct{ quickfix.FIXDecimal }

//Tag returns tag.AccruedInterestAmt (159)
func (f AccruedInterestAmtField) Tag() quickfix.Tag { return tag.AccruedInterestAmt }

//NewAccruedInterestAmt returns a new AccruedInterestAmtField initialized with val and scale
func NewAccruedInterestAmt(val decimal.Decimal, scale int32) AccruedInterestAmtField {
	return AccruedInterestAmtField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AccruedInterestAmtField) Value() (val decimal.Decimal) { return f.Decimal }

//AccruedInterestRateField is a PERCENTAGE field
type AccruedInterestRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.AccruedInterestRate (158)
func (f AccruedInterestRateField) Tag() quickfix.Tag { return tag.AccruedInterestRate }

//NewAccruedInterestRate returns a new AccruedInterestRateField initialized with val and scale
func NewAccruedInterestRate(val decimal.Decimal, scale int32) AccruedInterestRateField {
	return AccruedInterestRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AccruedInterestRateField) Value() (val decimal.Decimal) { return f.Decimal }

//AcctIDSourceField is a enum.AcctIDSource field
type AcctIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.AcctIDSource (660)
func (f AcctIDSourceField) Tag() quickfix.Tag { return tag.AcctIDSource }

func NewAcctIDSource(val enum.AcctIDSource) AcctIDSourceField {
	return AcctIDSourceField{quickfix.FIXString(val)}
}

func (f AcctIDSourceField) Value() enum.AcctIDSource { return enum.AcctIDSource(f.String()) }

//AdjustmentField is a enum.Adjustment field
type AdjustmentField struct{ quickfix.FIXString }

//Tag returns tag.Adjustment (334)
func (f AdjustmentField) Tag() quickfix.Tag { return tag.Adjustment }

func NewAdjustment(val enum.Adjustment) AdjustmentField {
	return AdjustmentField{quickfix.FIXString(val)}
}

func (f AdjustmentField) Value() enum.Adjustment { return enum.Adjustment(f.String()) }

//AdjustmentTypeField is a enum.AdjustmentType field
type AdjustmentTypeField struct{ quickfix.FIXString }

//Tag returns tag.AdjustmentType (718)
func (f AdjustmentTypeField) Tag() quickfix.Tag { return tag.AdjustmentType }

func NewAdjustmentType(val enum.AdjustmentType) AdjustmentTypeField {
	return AdjustmentTypeField{quickfix.FIXString(val)}
}

func (f AdjustmentTypeField) Value() enum.AdjustmentType { return enum.AdjustmentType(f.String()) }

//AdvIdField is a STRING field
type AdvIdField struct{ quickfix.FIXString }

//Tag returns tag.AdvId (2)
func (f AdvIdField) Tag() quickfix.Tag { return tag.AdvId }

//NewAdvId returns a new AdvIdField initialized with val
func NewAdvId(val string) AdvIdField {
	return AdvIdField{quickfix.FIXString(val)}
}

func (f AdvIdField) Value() string { return f.String() }

//AdvRefIDField is a STRING field
type AdvRefIDField struct{ quickfix.FIXString }

//Tag returns tag.AdvRefID (3)
func (f AdvRefIDField) Tag() quickfix.Tag { return tag.AdvRefID }

//NewAdvRefID returns a new AdvRefIDField initialized with val
func NewAdvRefID(val string) AdvRefIDField {
	return AdvRefIDField{quickfix.FIXString(val)}
}

func (f AdvRefIDField) Value() string { return f.String() }

//AdvSideField is a enum.AdvSide field
type AdvSideField struct{ quickfix.FIXString }

//Tag returns tag.AdvSide (4)
func (f AdvSideField) Tag() quickfix.Tag { return tag.AdvSide }

func NewAdvSide(val enum.AdvSide) AdvSideField {
	return AdvSideField{quickfix.FIXString(val)}
}

func (f AdvSideField) Value() enum.AdvSide { return enum.AdvSide(f.String()) }

//AdvTransTypeField is a enum.AdvTransType field
type AdvTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.AdvTransType (5)
func (f AdvTransTypeField) Tag() quickfix.Tag { return tag.AdvTransType }

func NewAdvTransType(val enum.AdvTransType) AdvTransTypeField {
	return AdvTransTypeField{quickfix.FIXString(val)}
}

func (f AdvTransTypeField) Value() enum.AdvTransType { return enum.AdvTransType(f.String()) }

//AffectedOrderIDField is a STRING field
type AffectedOrderIDField struct{ quickfix.FIXString }

//Tag returns tag.AffectedOrderID (535)
func (f AffectedOrderIDField) Tag() quickfix.Tag { return tag.AffectedOrderID }

//NewAffectedOrderID returns a new AffectedOrderIDField initialized with val
func NewAffectedOrderID(val string) AffectedOrderIDField {
	return AffectedOrderIDField{quickfix.FIXString(val)}
}

func (f AffectedOrderIDField) Value() string { return f.String() }

//AffectedSecondaryOrderIDField is a STRING field
type AffectedSecondaryOrderIDField struct{ quickfix.FIXString }

//Tag returns tag.AffectedSecondaryOrderID (536)
func (f AffectedSecondaryOrderIDField) Tag() quickfix.Tag { return tag.AffectedSecondaryOrderID }

//NewAffectedSecondaryOrderID returns a new AffectedSecondaryOrderIDField initialized with val
func NewAffectedSecondaryOrderID(val string) AffectedSecondaryOrderIDField {
	return AffectedSecondaryOrderIDField{quickfix.FIXString(val)}
}

func (f AffectedSecondaryOrderIDField) Value() string { return f.String() }

//AffirmStatusField is a enum.AffirmStatus field
type AffirmStatusField struct{ quickfix.FIXString }

//Tag returns tag.AffirmStatus (940)
func (f AffirmStatusField) Tag() quickfix.Tag { return tag.AffirmStatus }

func NewAffirmStatus(val enum.AffirmStatus) AffirmStatusField {
	return AffirmStatusField{quickfix.FIXString(val)}
}

func (f AffirmStatusField) Value() enum.AffirmStatus { return enum.AffirmStatus(f.String()) }

//AggregatedBookField is a BOOLEAN field
type AggregatedBookField struct{ quickfix.FIXBoolean }

//Tag returns tag.AggregatedBook (266)
func (f AggregatedBookField) Tag() quickfix.Tag { return tag.AggregatedBook }

//NewAggregatedBook returns a new AggregatedBookField initialized with val
func NewAggregatedBook(val bool) AggregatedBookField {
	return AggregatedBookField{quickfix.FIXBoolean(val)}
}

func (f AggregatedBookField) Value() bool { return f.Bool() }

//AggressorIndicatorField is a BOOLEAN field
type AggressorIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.AggressorIndicator (1057)
func (f AggressorIndicatorField) Tag() quickfix.Tag { return tag.AggressorIndicator }

//NewAggressorIndicator returns a new AggressorIndicatorField initialized with val
func NewAggressorIndicator(val bool) AggressorIndicatorField {
	return AggressorIndicatorField{quickfix.FIXBoolean(val)}
}

func (f AggressorIndicatorField) Value() bool { return f.Bool() }

//AgreementCurrencyField is a CURRENCY field
type AgreementCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.AgreementCurrency (918)
func (f AgreementCurrencyField) Tag() quickfix.Tag { return tag.AgreementCurrency }

//NewAgreementCurrency returns a new AgreementCurrencyField initialized with val
func NewAgreementCurrency(val string) AgreementCurrencyField {
	return AgreementCurrencyField{quickfix.FIXString(val)}
}

func (f AgreementCurrencyField) Value() string { return f.String() }

//AgreementDateField is a LOCALMKTDATE field
type AgreementDateField struct{ quickfix.FIXString }

//Tag returns tag.AgreementDate (915)
func (f AgreementDateField) Tag() quickfix.Tag { return tag.AgreementDate }

//NewAgreementDate returns a new AgreementDateField initialized with val
func NewAgreementDate(val string) AgreementDateField {
	return AgreementDateField{quickfix.FIXString(val)}
}

func (f AgreementDateField) Value() string { return f.String() }

//AgreementDescField is a STRING field
type AgreementDescField struct{ quickfix.FIXString }

//Tag returns tag.AgreementDesc (913)
func (f AgreementDescField) Tag() quickfix.Tag { return tag.AgreementDesc }

//NewAgreementDesc returns a new AgreementDescField initialized with val
func NewAgreementDesc(val string) AgreementDescField {
	return AgreementDescField{quickfix.FIXString(val)}
}

func (f AgreementDescField) Value() string { return f.String() }

//AgreementIDField is a STRING field
type AgreementIDField struct{ quickfix.FIXString }

//Tag returns tag.AgreementID (914)
func (f AgreementIDField) Tag() quickfix.Tag { return tag.AgreementID }

//NewAgreementID returns a new AgreementIDField initialized with val
func NewAgreementID(val string) AgreementIDField {
	return AgreementIDField{quickfix.FIXString(val)}
}

func (f AgreementIDField) Value() string { return f.String() }

//AllocAccountField is a STRING field
type AllocAccountField struct{ quickfix.FIXString }

//Tag returns tag.AllocAccount (79)
func (f AllocAccountField) Tag() quickfix.Tag { return tag.AllocAccount }

//NewAllocAccount returns a new AllocAccountField initialized with val
func NewAllocAccount(val string) AllocAccountField {
	return AllocAccountField{quickfix.FIXString(val)}
}

func (f AllocAccountField) Value() string { return f.String() }

//AllocAccountTypeField is a enum.AllocAccountType field
type AllocAccountTypeField struct{ quickfix.FIXString }

//Tag returns tag.AllocAccountType (798)
func (f AllocAccountTypeField) Tag() quickfix.Tag { return tag.AllocAccountType }

func NewAllocAccountType(val enum.AllocAccountType) AllocAccountTypeField {
	return AllocAccountTypeField{quickfix.FIXString(val)}
}

func (f AllocAccountTypeField) Value() enum.AllocAccountType { return enum.AllocAccountType(f.String()) }

//AllocAccruedInterestAmtField is a AMT field
type AllocAccruedInterestAmtField struct{ quickfix.FIXDecimal }

//Tag returns tag.AllocAccruedInterestAmt (742)
func (f AllocAccruedInterestAmtField) Tag() quickfix.Tag { return tag.AllocAccruedInterestAmt }

//NewAllocAccruedInterestAmt returns a new AllocAccruedInterestAmtField initialized with val and scale
func NewAllocAccruedInterestAmt(val decimal.Decimal, scale int32) AllocAccruedInterestAmtField {
	return AllocAccruedInterestAmtField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AllocAccruedInterestAmtField) Value() (val decimal.Decimal) { return f.Decimal }

//AllocAcctIDSourceField is a INT field
type AllocAcctIDSourceField struct{ quickfix.FIXInt }

//Tag returns tag.AllocAcctIDSource (661)
func (f AllocAcctIDSourceField) Tag() quickfix.Tag { return tag.AllocAcctIDSource }

//NewAllocAcctIDSource returns a new AllocAcctIDSourceField initialized with val
func NewAllocAcctIDSource(val int) AllocAcctIDSourceField {
	return AllocAcctIDSourceField{quickfix.FIXInt(val)}
}

func (f AllocAcctIDSourceField) Value() int { return f.Int() }

//AllocAvgPxField is a PRICE field
type AllocAvgPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.AllocAvgPx (153)
func (f AllocAvgPxField) Tag() quickfix.Tag { return tag.AllocAvgPx }

//NewAllocAvgPx returns a new AllocAvgPxField initialized with val and scale
func NewAllocAvgPx(val decimal.Decimal, scale int32) AllocAvgPxField {
	return AllocAvgPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AllocAvgPxField) Value() (val decimal.Decimal) { return f.Decimal }

//AllocCancReplaceReasonField is a enum.AllocCancReplaceReason field
type AllocCancReplaceReasonField struct{ quickfix.FIXString }

//Tag returns tag.AllocCancReplaceReason (796)
func (f AllocCancReplaceReasonField) Tag() quickfix.Tag { return tag.AllocCancReplaceReason }

func NewAllocCancReplaceReason(val enum.AllocCancReplaceReason) AllocCancReplaceReasonField {
	return AllocCancReplaceReasonField{quickfix.FIXString(val)}
}

func (f AllocCancReplaceReasonField) Value() enum.AllocCancReplaceReason {
	return enum.AllocCancReplaceReason(f.String())
}

//AllocClearingFeeIndicatorField is a STRING field
type AllocClearingFeeIndicatorField struct{ quickfix.FIXString }

//Tag returns tag.AllocClearingFeeIndicator (1136)
func (f AllocClearingFeeIndicatorField) Tag() quickfix.Tag { return tag.AllocClearingFeeIndicator }

//NewAllocClearingFeeIndicator returns a new AllocClearingFeeIndicatorField initialized with val
func NewAllocClearingFeeIndicator(val string) AllocClearingFeeIndicatorField {
	return AllocClearingFeeIndicatorField{quickfix.FIXString(val)}
}

func (f AllocClearingFeeIndicatorField) Value() string { return f.String() }

//AllocCustomerCapacityField is a STRING field
type AllocCustomerCapacityField struct{ quickfix.FIXString }

//Tag returns tag.AllocCustomerCapacity (993)
func (f AllocCustomerCapacityField) Tag() quickfix.Tag { return tag.AllocCustomerCapacity }

//NewAllocCustomerCapacity returns a new AllocCustomerCapacityField initialized with val
func NewAllocCustomerCapacity(val string) AllocCustomerCapacityField {
	return AllocCustomerCapacityField{quickfix.FIXString(val)}
}

func (f AllocCustomerCapacityField) Value() string { return f.String() }

//AllocHandlInstField is a enum.AllocHandlInst field
type AllocHandlInstField struct{ quickfix.FIXString }

//Tag returns tag.AllocHandlInst (209)
func (f AllocHandlInstField) Tag() quickfix.Tag { return tag.AllocHandlInst }

func NewAllocHandlInst(val enum.AllocHandlInst) AllocHandlInstField {
	return AllocHandlInstField{quickfix.FIXString(val)}
}

func (f AllocHandlInstField) Value() enum.AllocHandlInst { return enum.AllocHandlInst(f.String()) }

//AllocIDField is a STRING field
type AllocIDField struct{ quickfix.FIXString }

//Tag returns tag.AllocID (70)
func (f AllocIDField) Tag() quickfix.Tag { return tag.AllocID }

//NewAllocID returns a new AllocIDField initialized with val
func NewAllocID(val string) AllocIDField {
	return AllocIDField{quickfix.FIXString(val)}
}

func (f AllocIDField) Value() string { return f.String() }

//AllocInterestAtMaturityField is a AMT field
type AllocInterestAtMaturityField struct{ quickfix.FIXDecimal }

//Tag returns tag.AllocInterestAtMaturity (741)
func (f AllocInterestAtMaturityField) Tag() quickfix.Tag { return tag.AllocInterestAtMaturity }

//NewAllocInterestAtMaturity returns a new AllocInterestAtMaturityField initialized with val and scale
func NewAllocInterestAtMaturity(val decimal.Decimal, scale int32) AllocInterestAtMaturityField {
	return AllocInterestAtMaturityField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AllocInterestAtMaturityField) Value() (val decimal.Decimal) { return f.Decimal }

//AllocIntermedReqTypeField is a enum.AllocIntermedReqType field
type AllocIntermedReqTypeField struct{ quickfix.FIXString }

//Tag returns tag.AllocIntermedReqType (808)
func (f AllocIntermedReqTypeField) Tag() quickfix.Tag { return tag.AllocIntermedReqType }

func NewAllocIntermedReqType(val enum.AllocIntermedReqType) AllocIntermedReqTypeField {
	return AllocIntermedReqTypeField{quickfix.FIXString(val)}
}

func (f AllocIntermedReqTypeField) Value() enum.AllocIntermedReqType {
	return enum.AllocIntermedReqType(f.String())
}

//AllocLinkIDField is a STRING field
type AllocLinkIDField struct{ quickfix.FIXString }

//Tag returns tag.AllocLinkID (196)
func (f AllocLinkIDField) Tag() quickfix.Tag { return tag.AllocLinkID }

//NewAllocLinkID returns a new AllocLinkIDField initialized with val
func NewAllocLinkID(val string) AllocLinkIDField {
	return AllocLinkIDField{quickfix.FIXString(val)}
}

func (f AllocLinkIDField) Value() string { return f.String() }

//AllocLinkTypeField is a enum.AllocLinkType field
type AllocLinkTypeField struct{ quickfix.FIXString }

//Tag returns tag.AllocLinkType (197)
func (f AllocLinkTypeField) Tag() quickfix.Tag { return tag.AllocLinkType }

func NewAllocLinkType(val enum.AllocLinkType) AllocLinkTypeField {
	return AllocLinkTypeField{quickfix.FIXString(val)}
}

func (f AllocLinkTypeField) Value() enum.AllocLinkType { return enum.AllocLinkType(f.String()) }

//AllocMethodField is a enum.AllocMethod field
type AllocMethodField struct{ quickfix.FIXString }

//Tag returns tag.AllocMethod (1002)
func (f AllocMethodField) Tag() quickfix.Tag { return tag.AllocMethod }

func NewAllocMethod(val enum.AllocMethod) AllocMethodField {
	return AllocMethodField{quickfix.FIXString(val)}
}

func (f AllocMethodField) Value() enum.AllocMethod { return enum.AllocMethod(f.String()) }

//AllocNetMoneyField is a AMT field
type AllocNetMoneyField struct{ quickfix.FIXDecimal }

//Tag returns tag.AllocNetMoney (154)
func (f AllocNetMoneyField) Tag() quickfix.Tag { return tag.AllocNetMoney }

//NewAllocNetMoney returns a new AllocNetMoneyField initialized with val and scale
func NewAllocNetMoney(val decimal.Decimal, scale int32) AllocNetMoneyField {
	return AllocNetMoneyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AllocNetMoneyField) Value() (val decimal.Decimal) { return f.Decimal }

//AllocNoOrdersTypeField is a enum.AllocNoOrdersType field
type AllocNoOrdersTypeField struct{ quickfix.FIXString }

//Tag returns tag.AllocNoOrdersType (857)
func (f AllocNoOrdersTypeField) Tag() quickfix.Tag { return tag.AllocNoOrdersType }

func NewAllocNoOrdersType(val enum.AllocNoOrdersType) AllocNoOrdersTypeField {
	return AllocNoOrdersTypeField{quickfix.FIXString(val)}
}

func (f AllocNoOrdersTypeField) Value() enum.AllocNoOrdersType {
	return enum.AllocNoOrdersType(f.String())
}

//AllocPositionEffectField is a enum.AllocPositionEffect field
type AllocPositionEffectField struct{ quickfix.FIXString }

//Tag returns tag.AllocPositionEffect (1047)
func (f AllocPositionEffectField) Tag() quickfix.Tag { return tag.AllocPositionEffect }

func NewAllocPositionEffect(val enum.AllocPositionEffect) AllocPositionEffectField {
	return AllocPositionEffectField{quickfix.FIXString(val)}
}

func (f AllocPositionEffectField) Value() enum.AllocPositionEffect {
	return enum.AllocPositionEffect(f.String())
}

//AllocPriceField is a PRICE field
type AllocPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.AllocPrice (366)
func (f AllocPriceField) Tag() quickfix.Tag { return tag.AllocPrice }

//NewAllocPrice returns a new AllocPriceField initialized with val and scale
func NewAllocPrice(val decimal.Decimal, scale int32) AllocPriceField {
	return AllocPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AllocPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//AllocQtyField is a QTY field
type AllocQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.AllocQty (80)
func (f AllocQtyField) Tag() quickfix.Tag { return tag.AllocQty }

//NewAllocQty returns a new AllocQtyField initialized with val and scale
func NewAllocQty(val decimal.Decimal, scale int32) AllocQtyField {
	return AllocQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AllocQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//AllocRejCodeField is a enum.AllocRejCode field
type AllocRejCodeField struct{ quickfix.FIXString }

//Tag returns tag.AllocRejCode (88)
func (f AllocRejCodeField) Tag() quickfix.Tag { return tag.AllocRejCode }

func NewAllocRejCode(val enum.AllocRejCode) AllocRejCodeField {
	return AllocRejCodeField{quickfix.FIXString(val)}
}

func (f AllocRejCodeField) Value() enum.AllocRejCode { return enum.AllocRejCode(f.String()) }

//AllocReportIDField is a STRING field
type AllocReportIDField struct{ quickfix.FIXString }

//Tag returns tag.AllocReportID (755)
func (f AllocReportIDField) Tag() quickfix.Tag { return tag.AllocReportID }

//NewAllocReportID returns a new AllocReportIDField initialized with val
func NewAllocReportID(val string) AllocReportIDField {
	return AllocReportIDField{quickfix.FIXString(val)}
}

func (f AllocReportIDField) Value() string { return f.String() }

//AllocReportRefIDField is a STRING field
type AllocReportRefIDField struct{ quickfix.FIXString }

//Tag returns tag.AllocReportRefID (795)
func (f AllocReportRefIDField) Tag() quickfix.Tag { return tag.AllocReportRefID }

//NewAllocReportRefID returns a new AllocReportRefIDField initialized with val
func NewAllocReportRefID(val string) AllocReportRefIDField {
	return AllocReportRefIDField{quickfix.FIXString(val)}
}

func (f AllocReportRefIDField) Value() string { return f.String() }

//AllocReportTypeField is a enum.AllocReportType field
type AllocReportTypeField struct{ quickfix.FIXString }

//Tag returns tag.AllocReportType (794)
func (f AllocReportTypeField) Tag() quickfix.Tag { return tag.AllocReportType }

func NewAllocReportType(val enum.AllocReportType) AllocReportTypeField {
	return AllocReportTypeField{quickfix.FIXString(val)}
}

func (f AllocReportTypeField) Value() enum.AllocReportType { return enum.AllocReportType(f.String()) }

//AllocSettlCurrAmtField is a AMT field
type AllocSettlCurrAmtField struct{ quickfix.FIXDecimal }

//Tag returns tag.AllocSettlCurrAmt (737)
func (f AllocSettlCurrAmtField) Tag() quickfix.Tag { return tag.AllocSettlCurrAmt }

//NewAllocSettlCurrAmt returns a new AllocSettlCurrAmtField initialized with val and scale
func NewAllocSettlCurrAmt(val decimal.Decimal, scale int32) AllocSettlCurrAmtField {
	return AllocSettlCurrAmtField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AllocSettlCurrAmtField) Value() (val decimal.Decimal) { return f.Decimal }

//AllocSettlCurrencyField is a CURRENCY field
type AllocSettlCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.AllocSettlCurrency (736)
func (f AllocSettlCurrencyField) Tag() quickfix.Tag { return tag.AllocSettlCurrency }

//NewAllocSettlCurrency returns a new AllocSettlCurrencyField initialized with val
func NewAllocSettlCurrency(val string) AllocSettlCurrencyField {
	return AllocSettlCurrencyField{quickfix.FIXString(val)}
}

func (f AllocSettlCurrencyField) Value() string { return f.String() }

//AllocSettlInstTypeField is a enum.AllocSettlInstType field
type AllocSettlInstTypeField struct{ quickfix.FIXString }

//Tag returns tag.AllocSettlInstType (780)
func (f AllocSettlInstTypeField) Tag() quickfix.Tag { return tag.AllocSettlInstType }

func NewAllocSettlInstType(val enum.AllocSettlInstType) AllocSettlInstTypeField {
	return AllocSettlInstTypeField{quickfix.FIXString(val)}
}

func (f AllocSettlInstTypeField) Value() enum.AllocSettlInstType {
	return enum.AllocSettlInstType(f.String())
}

//AllocSharesField is a QTY field
type AllocSharesField struct{ quickfix.FIXDecimal }

//Tag returns tag.AllocShares (80)
func (f AllocSharesField) Tag() quickfix.Tag { return tag.AllocShares }

//NewAllocShares returns a new AllocSharesField initialized with val and scale
func NewAllocShares(val decimal.Decimal, scale int32) AllocSharesField {
	return AllocSharesField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AllocSharesField) Value() (val decimal.Decimal) { return f.Decimal }

//AllocStatusField is a enum.AllocStatus field
type AllocStatusField struct{ quickfix.FIXString }

//Tag returns tag.AllocStatus (87)
func (f AllocStatusField) Tag() quickfix.Tag { return tag.AllocStatus }

func NewAllocStatus(val enum.AllocStatus) AllocStatusField {
	return AllocStatusField{quickfix.FIXString(val)}
}

func (f AllocStatusField) Value() enum.AllocStatus { return enum.AllocStatus(f.String()) }

//AllocTextField is a STRING field
type AllocTextField struct{ quickfix.FIXString }

//Tag returns tag.AllocText (161)
func (f AllocTextField) Tag() quickfix.Tag { return tag.AllocText }

//NewAllocText returns a new AllocTextField initialized with val
func NewAllocText(val string) AllocTextField {
	return AllocTextField{quickfix.FIXString(val)}
}

func (f AllocTextField) Value() string { return f.String() }

//AllocTransTypeField is a enum.AllocTransType field
type AllocTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.AllocTransType (71)
func (f AllocTransTypeField) Tag() quickfix.Tag { return tag.AllocTransType }

func NewAllocTransType(val enum.AllocTransType) AllocTransTypeField {
	return AllocTransTypeField{quickfix.FIXString(val)}
}

func (f AllocTransTypeField) Value() enum.AllocTransType { return enum.AllocTransType(f.String()) }

//AllocTypeField is a enum.AllocType field
type AllocTypeField struct{ quickfix.FIXString }

//Tag returns tag.AllocType (626)
func (f AllocTypeField) Tag() quickfix.Tag { return tag.AllocType }

func NewAllocType(val enum.AllocType) AllocTypeField {
	return AllocTypeField{quickfix.FIXString(val)}
}

func (f AllocTypeField) Value() enum.AllocType { return enum.AllocType(f.String()) }

//AllowableOneSidednessCurrField is a CURRENCY field
type AllowableOneSidednessCurrField struct{ quickfix.FIXString }

//Tag returns tag.AllowableOneSidednessCurr (767)
func (f AllowableOneSidednessCurrField) Tag() quickfix.Tag { return tag.AllowableOneSidednessCurr }

//NewAllowableOneSidednessCurr returns a new AllowableOneSidednessCurrField initialized with val
func NewAllowableOneSidednessCurr(val string) AllowableOneSidednessCurrField {
	return AllowableOneSidednessCurrField{quickfix.FIXString(val)}
}

func (f AllowableOneSidednessCurrField) Value() string { return f.String() }

//AllowableOneSidednessPctField is a PERCENTAGE field
type AllowableOneSidednessPctField struct{ quickfix.FIXDecimal }

//Tag returns tag.AllowableOneSidednessPct (765)
func (f AllowableOneSidednessPctField) Tag() quickfix.Tag { return tag.AllowableOneSidednessPct }

//NewAllowableOneSidednessPct returns a new AllowableOneSidednessPctField initialized with val and scale
func NewAllowableOneSidednessPct(val decimal.Decimal, scale int32) AllowableOneSidednessPctField {
	return AllowableOneSidednessPctField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AllowableOneSidednessPctField) Value() (val decimal.Decimal) { return f.Decimal }

//AllowableOneSidednessValueField is a AMT field
type AllowableOneSidednessValueField struct{ quickfix.FIXDecimal }

//Tag returns tag.AllowableOneSidednessValue (766)
func (f AllowableOneSidednessValueField) Tag() quickfix.Tag { return tag.AllowableOneSidednessValue }

//NewAllowableOneSidednessValue returns a new AllowableOneSidednessValueField initialized with val and scale
func NewAllowableOneSidednessValue(val decimal.Decimal, scale int32) AllowableOneSidednessValueField {
	return AllowableOneSidednessValueField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AllowableOneSidednessValueField) Value() (val decimal.Decimal) { return f.Decimal }

//AltMDSourceIDField is a STRING field
type AltMDSourceIDField struct{ quickfix.FIXString }

//Tag returns tag.AltMDSourceID (817)
func (f AltMDSourceIDField) Tag() quickfix.Tag { return tag.AltMDSourceID }

//NewAltMDSourceID returns a new AltMDSourceIDField initialized with val
func NewAltMDSourceID(val string) AltMDSourceIDField {
	return AltMDSourceIDField{quickfix.FIXString(val)}
}

func (f AltMDSourceIDField) Value() string { return f.String() }

//ApplBegSeqNumField is a SEQNUM field
type ApplBegSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.ApplBegSeqNum (1182)
func (f ApplBegSeqNumField) Tag() quickfix.Tag { return tag.ApplBegSeqNum }

//NewApplBegSeqNum returns a new ApplBegSeqNumField initialized with val
func NewApplBegSeqNum(val int) ApplBegSeqNumField {
	return ApplBegSeqNumField{quickfix.FIXInt(val)}
}

func (f ApplBegSeqNumField) Value() int { return f.Int() }

//ApplEndSeqNumField is a SEQNUM field
type ApplEndSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.ApplEndSeqNum (1183)
func (f ApplEndSeqNumField) Tag() quickfix.Tag { return tag.ApplEndSeqNum }

//NewApplEndSeqNum returns a new ApplEndSeqNumField initialized with val
func NewApplEndSeqNum(val int) ApplEndSeqNumField {
	return ApplEndSeqNumField{quickfix.FIXInt(val)}
}

func (f ApplEndSeqNumField) Value() int { return f.Int() }

//ApplExtIDField is a INT field
type ApplExtIDField struct{ quickfix.FIXInt }

//Tag returns tag.ApplExtID (1156)
func (f ApplExtIDField) Tag() quickfix.Tag { return tag.ApplExtID }

//NewApplExtID returns a new ApplExtIDField initialized with val
func NewApplExtID(val int) ApplExtIDField {
	return ApplExtIDField{quickfix.FIXInt(val)}
}

func (f ApplExtIDField) Value() int { return f.Int() }

//ApplIDField is a STRING field
type ApplIDField struct{ quickfix.FIXString }

//Tag returns tag.ApplID (1180)
func (f ApplIDField) Tag() quickfix.Tag { return tag.ApplID }

//NewApplID returns a new ApplIDField initialized with val
func NewApplID(val string) ApplIDField {
	return ApplIDField{quickfix.FIXString(val)}
}

func (f ApplIDField) Value() string { return f.String() }

//ApplLastSeqNumField is a SEQNUM field
type ApplLastSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.ApplLastSeqNum (1350)
func (f ApplLastSeqNumField) Tag() quickfix.Tag { return tag.ApplLastSeqNum }

//NewApplLastSeqNum returns a new ApplLastSeqNumField initialized with val
func NewApplLastSeqNum(val int) ApplLastSeqNumField {
	return ApplLastSeqNumField{quickfix.FIXInt(val)}
}

func (f ApplLastSeqNumField) Value() int { return f.Int() }

//ApplNewSeqNumField is a SEQNUM field
type ApplNewSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.ApplNewSeqNum (1399)
func (f ApplNewSeqNumField) Tag() quickfix.Tag { return tag.ApplNewSeqNum }

//NewApplNewSeqNum returns a new ApplNewSeqNumField initialized with val
func NewApplNewSeqNum(val int) ApplNewSeqNumField {
	return ApplNewSeqNumField{quickfix.FIXInt(val)}
}

func (f ApplNewSeqNumField) Value() int { return f.Int() }

//ApplQueueActionField is a enum.ApplQueueAction field
type ApplQueueActionField struct{ quickfix.FIXString }

//Tag returns tag.ApplQueueAction (815)
func (f ApplQueueActionField) Tag() quickfix.Tag { return tag.ApplQueueAction }

func NewApplQueueAction(val enum.ApplQueueAction) ApplQueueActionField {
	return ApplQueueActionField{quickfix.FIXString(val)}
}

func (f ApplQueueActionField) Value() enum.ApplQueueAction { return enum.ApplQueueAction(f.String()) }

//ApplQueueDepthField is a INT field
type ApplQueueDepthField struct{ quickfix.FIXInt }

//Tag returns tag.ApplQueueDepth (813)
func (f ApplQueueDepthField) Tag() quickfix.Tag { return tag.ApplQueueDepth }

//NewApplQueueDepth returns a new ApplQueueDepthField initialized with val
func NewApplQueueDepth(val int) ApplQueueDepthField {
	return ApplQueueDepthField{quickfix.FIXInt(val)}
}

func (f ApplQueueDepthField) Value() int { return f.Int() }

//ApplQueueMaxField is a INT field
type ApplQueueMaxField struct{ quickfix.FIXInt }

//Tag returns tag.ApplQueueMax (812)
func (f ApplQueueMaxField) Tag() quickfix.Tag { return tag.ApplQueueMax }

//NewApplQueueMax returns a new ApplQueueMaxField initialized with val
func NewApplQueueMax(val int) ApplQueueMaxField {
	return ApplQueueMaxField{quickfix.FIXInt(val)}
}

func (f ApplQueueMaxField) Value() int { return f.Int() }

//ApplQueueResolutionField is a enum.ApplQueueResolution field
type ApplQueueResolutionField struct{ quickfix.FIXString }

//Tag returns tag.ApplQueueResolution (814)
func (f ApplQueueResolutionField) Tag() quickfix.Tag { return tag.ApplQueueResolution }

func NewApplQueueResolution(val enum.ApplQueueResolution) ApplQueueResolutionField {
	return ApplQueueResolutionField{quickfix.FIXString(val)}
}

func (f ApplQueueResolutionField) Value() enum.ApplQueueResolution {
	return enum.ApplQueueResolution(f.String())
}

//ApplReportIDField is a STRING field
type ApplReportIDField struct{ quickfix.FIXString }

//Tag returns tag.ApplReportID (1356)
func (f ApplReportIDField) Tag() quickfix.Tag { return tag.ApplReportID }

//NewApplReportID returns a new ApplReportIDField initialized with val
func NewApplReportID(val string) ApplReportIDField {
	return ApplReportIDField{quickfix.FIXString(val)}
}

func (f ApplReportIDField) Value() string { return f.String() }

//ApplReportTypeField is a enum.ApplReportType field
type ApplReportTypeField struct{ quickfix.FIXString }

//Tag returns tag.ApplReportType (1426)
func (f ApplReportTypeField) Tag() quickfix.Tag { return tag.ApplReportType }

func NewApplReportType(val enum.ApplReportType) ApplReportTypeField {
	return ApplReportTypeField{quickfix.FIXString(val)}
}

func (f ApplReportTypeField) Value() enum.ApplReportType { return enum.ApplReportType(f.String()) }

//ApplReqIDField is a STRING field
type ApplReqIDField struct{ quickfix.FIXString }

//Tag returns tag.ApplReqID (1346)
func (f ApplReqIDField) Tag() quickfix.Tag { return tag.ApplReqID }

//NewApplReqID returns a new ApplReqIDField initialized with val
func NewApplReqID(val string) ApplReqIDField {
	return ApplReqIDField{quickfix.FIXString(val)}
}

func (f ApplReqIDField) Value() string { return f.String() }

//ApplReqTypeField is a enum.ApplReqType field
type ApplReqTypeField struct{ quickfix.FIXString }

//Tag returns tag.ApplReqType (1347)
func (f ApplReqTypeField) Tag() quickfix.Tag { return tag.ApplReqType }

func NewApplReqType(val enum.ApplReqType) ApplReqTypeField {
	return ApplReqTypeField{quickfix.FIXString(val)}
}

func (f ApplReqTypeField) Value() enum.ApplReqType { return enum.ApplReqType(f.String()) }

//ApplResendFlagField is a BOOLEAN field
type ApplResendFlagField struct{ quickfix.FIXBoolean }

//Tag returns tag.ApplResendFlag (1352)
func (f ApplResendFlagField) Tag() quickfix.Tag { return tag.ApplResendFlag }

//NewApplResendFlag returns a new ApplResendFlagField initialized with val
func NewApplResendFlag(val bool) ApplResendFlagField {
	return ApplResendFlagField{quickfix.FIXBoolean(val)}
}

func (f ApplResendFlagField) Value() bool { return f.Bool() }

//ApplResponseErrorField is a enum.ApplResponseError field
type ApplResponseErrorField struct{ quickfix.FIXString }

//Tag returns tag.ApplResponseError (1354)
func (f ApplResponseErrorField) Tag() quickfix.Tag { return tag.ApplResponseError }

func NewApplResponseError(val enum.ApplResponseError) ApplResponseErrorField {
	return ApplResponseErrorField{quickfix.FIXString(val)}
}

func (f ApplResponseErrorField) Value() enum.ApplResponseError {
	return enum.ApplResponseError(f.String())
}

//ApplResponseIDField is a STRING field
type ApplResponseIDField struct{ quickfix.FIXString }

//Tag returns tag.ApplResponseID (1353)
func (f ApplResponseIDField) Tag() quickfix.Tag { return tag.ApplResponseID }

//NewApplResponseID returns a new ApplResponseIDField initialized with val
func NewApplResponseID(val string) ApplResponseIDField {
	return ApplResponseIDField{quickfix.FIXString(val)}
}

func (f ApplResponseIDField) Value() string { return f.String() }

//ApplResponseTypeField is a enum.ApplResponseType field
type ApplResponseTypeField struct{ quickfix.FIXString }

//Tag returns tag.ApplResponseType (1348)
func (f ApplResponseTypeField) Tag() quickfix.Tag { return tag.ApplResponseType }

func NewApplResponseType(val enum.ApplResponseType) ApplResponseTypeField {
	return ApplResponseTypeField{quickfix.FIXString(val)}
}

func (f ApplResponseTypeField) Value() enum.ApplResponseType { return enum.ApplResponseType(f.String()) }

//ApplSeqNumField is a SEQNUM field
type ApplSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.ApplSeqNum (1181)
func (f ApplSeqNumField) Tag() quickfix.Tag { return tag.ApplSeqNum }

//NewApplSeqNum returns a new ApplSeqNumField initialized with val
func NewApplSeqNum(val int) ApplSeqNumField {
	return ApplSeqNumField{quickfix.FIXInt(val)}
}

func (f ApplSeqNumField) Value() int { return f.Int() }

//ApplTotalMessageCountField is a INT field
type ApplTotalMessageCountField struct{ quickfix.FIXInt }

//Tag returns tag.ApplTotalMessageCount (1349)
func (f ApplTotalMessageCountField) Tag() quickfix.Tag { return tag.ApplTotalMessageCount }

//NewApplTotalMessageCount returns a new ApplTotalMessageCountField initialized with val
func NewApplTotalMessageCount(val int) ApplTotalMessageCountField {
	return ApplTotalMessageCountField{quickfix.FIXInt(val)}
}

func (f ApplTotalMessageCountField) Value() int { return f.Int() }

//ApplVerIDField is a enum.ApplVerID field
type ApplVerIDField struct{ quickfix.FIXString }

//Tag returns tag.ApplVerID (1128)
func (f ApplVerIDField) Tag() quickfix.Tag { return tag.ApplVerID }

func NewApplVerID(val enum.ApplVerID) ApplVerIDField {
	return ApplVerIDField{quickfix.FIXString(val)}
}

func (f ApplVerIDField) Value() enum.ApplVerID { return enum.ApplVerID(f.String()) }

//AsOfIndicatorField is a enum.AsOfIndicator field
type AsOfIndicatorField struct{ quickfix.FIXString }

//Tag returns tag.AsOfIndicator (1015)
func (f AsOfIndicatorField) Tag() quickfix.Tag { return tag.AsOfIndicator }

func NewAsOfIndicator(val enum.AsOfIndicator) AsOfIndicatorField {
	return AsOfIndicatorField{quickfix.FIXString(val)}
}

func (f AsOfIndicatorField) Value() enum.AsOfIndicator { return enum.AsOfIndicator(f.String()) }

//AsgnReqIDField is a STRING field
type AsgnReqIDField struct{ quickfix.FIXString }

//Tag returns tag.AsgnReqID (831)
func (f AsgnReqIDField) Tag() quickfix.Tag { return tag.AsgnReqID }

//NewAsgnReqID returns a new AsgnReqIDField initialized with val
func NewAsgnReqID(val string) AsgnReqIDField {
	return AsgnReqIDField{quickfix.FIXString(val)}
}

func (f AsgnReqIDField) Value() string { return f.String() }

//AsgnRptIDField is a STRING field
type AsgnRptIDField struct{ quickfix.FIXString }

//Tag returns tag.AsgnRptID (833)
func (f AsgnRptIDField) Tag() quickfix.Tag { return tag.AsgnRptID }

//NewAsgnRptID returns a new AsgnRptIDField initialized with val
func NewAsgnRptID(val string) AsgnRptIDField {
	return AsgnRptIDField{quickfix.FIXString(val)}
}

func (f AsgnRptIDField) Value() string { return f.String() }

//AssignmentMethodField is a enum.AssignmentMethod field
type AssignmentMethodField struct{ quickfix.FIXString }

//Tag returns tag.AssignmentMethod (744)
func (f AssignmentMethodField) Tag() quickfix.Tag { return tag.AssignmentMethod }

func NewAssignmentMethod(val enum.AssignmentMethod) AssignmentMethodField {
	return AssignmentMethodField{quickfix.FIXString(val)}
}

func (f AssignmentMethodField) Value() enum.AssignmentMethod { return enum.AssignmentMethod(f.String()) }

//AssignmentUnitField is a QTY field
type AssignmentUnitField struct{ quickfix.FIXDecimal }

//Tag returns tag.AssignmentUnit (745)
func (f AssignmentUnitField) Tag() quickfix.Tag { return tag.AssignmentUnit }

//NewAssignmentUnit returns a new AssignmentUnitField initialized with val and scale
func NewAssignmentUnit(val decimal.Decimal, scale int32) AssignmentUnitField {
	return AssignmentUnitField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AssignmentUnitField) Value() (val decimal.Decimal) { return f.Decimal }

//AttachmentPointField is a PERCENTAGE field
type AttachmentPointField struct{ quickfix.FIXDecimal }

//Tag returns tag.AttachmentPoint (1457)
func (f AttachmentPointField) Tag() quickfix.Tag { return tag.AttachmentPoint }

//NewAttachmentPoint returns a new AttachmentPointField initialized with val and scale
func NewAttachmentPoint(val decimal.Decimal, scale int32) AttachmentPointField {
	return AttachmentPointField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AttachmentPointField) Value() (val decimal.Decimal) { return f.Decimal }

//AutoAcceptIndicatorField is a BOOLEAN field
type AutoAcceptIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.AutoAcceptIndicator (754)
func (f AutoAcceptIndicatorField) Tag() quickfix.Tag { return tag.AutoAcceptIndicator }

//NewAutoAcceptIndicator returns a new AutoAcceptIndicatorField initialized with val
func NewAutoAcceptIndicator(val bool) AutoAcceptIndicatorField {
	return AutoAcceptIndicatorField{quickfix.FIXBoolean(val)}
}

func (f AutoAcceptIndicatorField) Value() bool { return f.Bool() }

//AvgParPxField is a PRICE field
type AvgParPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.AvgParPx (860)
func (f AvgParPxField) Tag() quickfix.Tag { return tag.AvgParPx }

//NewAvgParPx returns a new AvgParPxField initialized with val and scale
func NewAvgParPx(val decimal.Decimal, scale int32) AvgParPxField {
	return AvgParPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AvgParPxField) Value() (val decimal.Decimal) { return f.Decimal }

//AvgPrxPrecisionField is a INT field
type AvgPrxPrecisionField struct{ quickfix.FIXInt }

//Tag returns tag.AvgPrxPrecision (74)
func (f AvgPrxPrecisionField) Tag() quickfix.Tag { return tag.AvgPrxPrecision }

//NewAvgPrxPrecision returns a new AvgPrxPrecisionField initialized with val
func NewAvgPrxPrecision(val int) AvgPrxPrecisionField {
	return AvgPrxPrecisionField{quickfix.FIXInt(val)}
}

func (f AvgPrxPrecisionField) Value() int { return f.Int() }

//AvgPxField is a PRICE field
type AvgPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.AvgPx (6)
func (f AvgPxField) Tag() quickfix.Tag { return tag.AvgPx }

//NewAvgPx returns a new AvgPxField initialized with val and scale
func NewAvgPx(val decimal.Decimal, scale int32) AvgPxField {
	return AvgPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AvgPxField) Value() (val decimal.Decimal) { return f.Decimal }

//AvgPxIndicatorField is a enum.AvgPxIndicator field
type AvgPxIndicatorField struct{ quickfix.FIXString }

//Tag returns tag.AvgPxIndicator (819)
func (f AvgPxIndicatorField) Tag() quickfix.Tag { return tag.AvgPxIndicator }

func NewAvgPxIndicator(val enum.AvgPxIndicator) AvgPxIndicatorField {
	return AvgPxIndicatorField{quickfix.FIXString(val)}
}

func (f AvgPxIndicatorField) Value() enum.AvgPxIndicator { return enum.AvgPxIndicator(f.String()) }

//AvgPxPrecisionField is a INT field
type AvgPxPrecisionField struct{ quickfix.FIXInt }

//Tag returns tag.AvgPxPrecision (74)
func (f AvgPxPrecisionField) Tag() quickfix.Tag { return tag.AvgPxPrecision }

//NewAvgPxPrecision returns a new AvgPxPrecisionField initialized with val
func NewAvgPxPrecision(val int) AvgPxPrecisionField {
	return AvgPxPrecisionField{quickfix.FIXInt(val)}
}

func (f AvgPxPrecisionField) Value() int { return f.Int() }

//BasisFeatureDateField is a LOCALMKTDATE field
type BasisFeatureDateField struct{ quickfix.FIXString }

//Tag returns tag.BasisFeatureDate (259)
func (f BasisFeatureDateField) Tag() quickfix.Tag { return tag.BasisFeatureDate }

//NewBasisFeatureDate returns a new BasisFeatureDateField initialized with val
func NewBasisFeatureDate(val string) BasisFeatureDateField {
	return BasisFeatureDateField{quickfix.FIXString(val)}
}

func (f BasisFeatureDateField) Value() string { return f.String() }

//BasisFeaturePriceField is a PRICE field
type BasisFeaturePriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.BasisFeaturePrice (260)
func (f BasisFeaturePriceField) Tag() quickfix.Tag { return tag.BasisFeaturePrice }

//NewBasisFeaturePrice returns a new BasisFeaturePriceField initialized with val and scale
func NewBasisFeaturePrice(val decimal.Decimal, scale int32) BasisFeaturePriceField {
	return BasisFeaturePriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f BasisFeaturePriceField) Value() (val decimal.Decimal) { return f.Decimal }

//BasisPxTypeField is a enum.BasisPxType field
type BasisPxTypeField struct{ quickfix.FIXString }

//Tag returns tag.BasisPxType (419)
func (f BasisPxTypeField) Tag() quickfix.Tag { return tag.BasisPxType }

func NewBasisPxType(val enum.BasisPxType) BasisPxTypeField {
	return BasisPxTypeField{quickfix.FIXString(val)}
}

func (f BasisPxTypeField) Value() enum.BasisPxType { return enum.BasisPxType(f.String()) }

//BeginSeqNoField is a SEQNUM field
type BeginSeqNoField struct{ quickfix.FIXInt }

//Tag returns tag.BeginSeqNo (7)
func (f BeginSeqNoField) Tag() quickfix.Tag { return tag.BeginSeqNo }

//NewBeginSeqNo returns a new BeginSeqNoField initialized with val
func NewBeginSeqNo(val int) BeginSeqNoField {
	return BeginSeqNoField{quickfix.FIXInt(val)}
}

func (f BeginSeqNoField) Value() int { return f.Int() }

//BeginStringField is a STRING field
type BeginStringField struct{ quickfix.FIXString }

//Tag returns tag.BeginString (8)
func (f BeginStringField) Tag() quickfix.Tag { return tag.BeginString }

//NewBeginString returns a new BeginStringField initialized with val
func NewBeginString(val string) BeginStringField {
	return BeginStringField{quickfix.FIXString(val)}
}

func (f BeginStringField) Value() string { return f.String() }

//BenchmarkField is a enum.Benchmark field
type BenchmarkField struct{ quickfix.FIXString }

//Tag returns tag.Benchmark (219)
func (f BenchmarkField) Tag() quickfix.Tag { return tag.Benchmark }

func NewBenchmark(val enum.Benchmark) BenchmarkField {
	return BenchmarkField{quickfix.FIXString(val)}
}

func (f BenchmarkField) Value() enum.Benchmark { return enum.Benchmark(f.String()) }

//BenchmarkCurveCurrencyField is a CURRENCY field
type BenchmarkCurveCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.BenchmarkCurveCurrency (220)
func (f BenchmarkCurveCurrencyField) Tag() quickfix.Tag { return tag.BenchmarkCurveCurrency }

//NewBenchmarkCurveCurrency returns a new BenchmarkCurveCurrencyField initialized with val
func NewBenchmarkCurveCurrency(val string) BenchmarkCurveCurrencyField {
	return BenchmarkCurveCurrencyField{quickfix.FIXString(val)}
}

func (f BenchmarkCurveCurrencyField) Value() string { return f.String() }

//BenchmarkCurveNameField is a enum.BenchmarkCurveName field
type BenchmarkCurveNameField struct{ quickfix.FIXString }

//Tag returns tag.BenchmarkCurveName (221)
func (f BenchmarkCurveNameField) Tag() quickfix.Tag { return tag.BenchmarkCurveName }

func NewBenchmarkCurveName(val enum.BenchmarkCurveName) BenchmarkCurveNameField {
	return BenchmarkCurveNameField{quickfix.FIXString(val)}
}

func (f BenchmarkCurveNameField) Value() enum.BenchmarkCurveName {
	return enum.BenchmarkCurveName(f.String())
}

//BenchmarkCurvePointField is a STRING field
type BenchmarkCurvePointField struct{ quickfix.FIXString }

//Tag returns tag.BenchmarkCurvePoint (222)
func (f BenchmarkCurvePointField) Tag() quickfix.Tag { return tag.BenchmarkCurvePoint }

//NewBenchmarkCurvePoint returns a new BenchmarkCurvePointField initialized with val
func NewBenchmarkCurvePoint(val string) BenchmarkCurvePointField {
	return BenchmarkCurvePointField{quickfix.FIXString(val)}
}

func (f BenchmarkCurvePointField) Value() string { return f.String() }

//BenchmarkPriceField is a PRICE field
type BenchmarkPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.BenchmarkPrice (662)
func (f BenchmarkPriceField) Tag() quickfix.Tag { return tag.BenchmarkPrice }

//NewBenchmarkPrice returns a new BenchmarkPriceField initialized with val and scale
func NewBenchmarkPrice(val decimal.Decimal, scale int32) BenchmarkPriceField {
	return BenchmarkPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f BenchmarkPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//BenchmarkPriceTypeField is a INT field
type BenchmarkPriceTypeField struct{ quickfix.FIXInt }

//Tag returns tag.BenchmarkPriceType (663)
func (f BenchmarkPriceTypeField) Tag() quickfix.Tag { return tag.BenchmarkPriceType }

//NewBenchmarkPriceType returns a new BenchmarkPriceTypeField initialized with val
func NewBenchmarkPriceType(val int) BenchmarkPriceTypeField {
	return BenchmarkPriceTypeField{quickfix.FIXInt(val)}
}

func (f BenchmarkPriceTypeField) Value() int { return f.Int() }

//BenchmarkSecurityIDField is a STRING field
type BenchmarkSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.BenchmarkSecurityID (699)
func (f BenchmarkSecurityIDField) Tag() quickfix.Tag { return tag.BenchmarkSecurityID }

//NewBenchmarkSecurityID returns a new BenchmarkSecurityIDField initialized with val
func NewBenchmarkSecurityID(val string) BenchmarkSecurityIDField {
	return BenchmarkSecurityIDField{quickfix.FIXString(val)}
}

func (f BenchmarkSecurityIDField) Value() string { return f.String() }

//BenchmarkSecurityIDSourceField is a STRING field
type BenchmarkSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.BenchmarkSecurityIDSource (761)
func (f BenchmarkSecurityIDSourceField) Tag() quickfix.Tag { return tag.BenchmarkSecurityIDSource }

//NewBenchmarkSecurityIDSource returns a new BenchmarkSecurityIDSourceField initialized with val
func NewBenchmarkSecurityIDSource(val string) BenchmarkSecurityIDSourceField {
	return BenchmarkSecurityIDSourceField{quickfix.FIXString(val)}
}

func (f BenchmarkSecurityIDSourceField) Value() string { return f.String() }

//BidDescriptorField is a STRING field
type BidDescriptorField struct{ quickfix.FIXString }

//Tag returns tag.BidDescriptor (400)
func (f BidDescriptorField) Tag() quickfix.Tag { return tag.BidDescriptor }

//NewBidDescriptor returns a new BidDescriptorField initialized with val
func NewBidDescriptor(val string) BidDescriptorField {
	return BidDescriptorField{quickfix.FIXString(val)}
}

func (f BidDescriptorField) Value() string { return f.String() }

//BidDescriptorTypeField is a enum.BidDescriptorType field
type BidDescriptorTypeField struct{ quickfix.FIXString }

//Tag returns tag.BidDescriptorType (399)
func (f BidDescriptorTypeField) Tag() quickfix.Tag { return tag.BidDescriptorType }

func NewBidDescriptorType(val enum.BidDescriptorType) BidDescriptorTypeField {
	return BidDescriptorTypeField{quickfix.FIXString(val)}
}

func (f BidDescriptorTypeField) Value() enum.BidDescriptorType {
	return enum.BidDescriptorType(f.String())
}

//BidForwardPointsField is a PRICEOFFSET field
type BidForwardPointsField struct{ quickfix.FIXDecimal }

//Tag returns tag.BidForwardPoints (189)
func (f BidForwardPointsField) Tag() quickfix.Tag { return tag.BidForwardPoints }

//NewBidForwardPoints returns a new BidForwardPointsField initialized with val and scale
func NewBidForwardPoints(val decimal.Decimal, scale int32) BidForwardPointsField {
	return BidForwardPointsField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f BidForwardPointsField) Value() (val decimal.Decimal) { return f.Decimal }

//BidForwardPoints2Field is a PRICEOFFSET field
type BidForwardPoints2Field struct{ quickfix.FIXDecimal }

//Tag returns tag.BidForwardPoints2 (642)
func (f BidForwardPoints2Field) Tag() quickfix.Tag { return tag.BidForwardPoints2 }

//NewBidForwardPoints2 returns a new BidForwardPoints2Field initialized with val and scale
func NewBidForwardPoints2(val decimal.Decimal, scale int32) BidForwardPoints2Field {
	return BidForwardPoints2Field{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f BidForwardPoints2Field) Value() (val decimal.Decimal) { return f.Decimal }

//BidIDField is a STRING field
type BidIDField struct{ quickfix.FIXString }

//Tag returns tag.BidID (390)
func (f BidIDField) Tag() quickfix.Tag { return tag.BidID }

//NewBidID returns a new BidIDField initialized with val
func NewBidID(val string) BidIDField {
	return BidIDField{quickfix.FIXString(val)}
}

func (f BidIDField) Value() string { return f.String() }

//BidPxField is a PRICE field
type BidPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.BidPx (132)
func (f BidPxField) Tag() quickfix.Tag { return tag.BidPx }

//NewBidPx returns a new BidPxField initialized with val and scale
func NewBidPx(val decimal.Decimal, scale int32) BidPxField {
	return BidPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f BidPxField) Value() (val decimal.Decimal) { return f.Decimal }

//BidRequestTransTypeField is a enum.BidRequestTransType field
type BidRequestTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.BidRequestTransType (374)
func (f BidRequestTransTypeField) Tag() quickfix.Tag { return tag.BidRequestTransType }

func NewBidRequestTransType(val enum.BidRequestTransType) BidRequestTransTypeField {
	return BidRequestTransTypeField{quickfix.FIXString(val)}
}

func (f BidRequestTransTypeField) Value() enum.BidRequestTransType {
	return enum.BidRequestTransType(f.String())
}

//BidSizeField is a QTY field
type BidSizeField struct{ quickfix.FIXDecimal }

//Tag returns tag.BidSize (134)
func (f BidSizeField) Tag() quickfix.Tag { return tag.BidSize }

//NewBidSize returns a new BidSizeField initialized with val and scale
func NewBidSize(val decimal.Decimal, scale int32) BidSizeField {
	return BidSizeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f BidSizeField) Value() (val decimal.Decimal) { return f.Decimal }

//BidSpotRateField is a PRICE field
type BidSpotRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.BidSpotRate (188)
func (f BidSpotRateField) Tag() quickfix.Tag { return tag.BidSpotRate }

//NewBidSpotRate returns a new BidSpotRateField initialized with val and scale
func NewBidSpotRate(val decimal.Decimal, scale int32) BidSpotRateField {
	return BidSpotRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f BidSpotRateField) Value() (val decimal.Decimal) { return f.Decimal }

//BidSwapPointsField is a PRICEOFFSET field
type BidSwapPointsField struct{ quickfix.FIXDecimal }

//Tag returns tag.BidSwapPoints (1065)
func (f BidSwapPointsField) Tag() quickfix.Tag { return tag.BidSwapPoints }

//NewBidSwapPoints returns a new BidSwapPointsField initialized with val and scale
func NewBidSwapPoints(val decimal.Decimal, scale int32) BidSwapPointsField {
	return BidSwapPointsField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f BidSwapPointsField) Value() (val decimal.Decimal) { return f.Decimal }

//BidTradeTypeField is a enum.BidTradeType field
type BidTradeTypeField struct{ quickfix.FIXString }

//Tag returns tag.BidTradeType (418)
func (f BidTradeTypeField) Tag() quickfix.Tag { return tag.BidTradeType }

func NewBidTradeType(val enum.BidTradeType) BidTradeTypeField {
	return BidTradeTypeField{quickfix.FIXString(val)}
}

func (f BidTradeTypeField) Value() enum.BidTradeType { return enum.BidTradeType(f.String()) }

//BidTypeField is a enum.BidType field
type BidTypeField struct{ quickfix.FIXString }

//Tag returns tag.BidType (394)
func (f BidTypeField) Tag() quickfix.Tag { return tag.BidType }

func NewBidType(val enum.BidType) BidTypeField {
	return BidTypeField{quickfix.FIXString(val)}
}

func (f BidTypeField) Value() enum.BidType { return enum.BidType(f.String()) }

//BidYieldField is a PERCENTAGE field
type BidYieldField struct{ quickfix.FIXDecimal }

//Tag returns tag.BidYield (632)
func (f BidYieldField) Tag() quickfix.Tag { return tag.BidYield }

//NewBidYield returns a new BidYieldField initialized with val and scale
func NewBidYield(val decimal.Decimal, scale int32) BidYieldField {
	return BidYieldField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f BidYieldField) Value() (val decimal.Decimal) { return f.Decimal }

//BodyLengthField is a LENGTH field
type BodyLengthField struct{ quickfix.FIXInt }

//Tag returns tag.BodyLength (9)
func (f BodyLengthField) Tag() quickfix.Tag { return tag.BodyLength }

//NewBodyLength returns a new BodyLengthField initialized with val
func NewBodyLength(val int) BodyLengthField {
	return BodyLengthField{quickfix.FIXInt(val)}
}

func (f BodyLengthField) Value() int { return f.Int() }

//BookingRefIDField is a STRING field
type BookingRefIDField struct{ quickfix.FIXString }

//Tag returns tag.BookingRefID (466)
func (f BookingRefIDField) Tag() quickfix.Tag { return tag.BookingRefID }

//NewBookingRefID returns a new BookingRefIDField initialized with val
func NewBookingRefID(val string) BookingRefIDField {
	return BookingRefIDField{quickfix.FIXString(val)}
}

func (f BookingRefIDField) Value() string { return f.String() }

//BookingTypeField is a enum.BookingType field
type BookingTypeField struct{ quickfix.FIXString }

//Tag returns tag.BookingType (775)
func (f BookingTypeField) Tag() quickfix.Tag { return tag.BookingType }

func NewBookingType(val enum.BookingType) BookingTypeField {
	return BookingTypeField{quickfix.FIXString(val)}
}

func (f BookingTypeField) Value() enum.BookingType { return enum.BookingType(f.String()) }

//BookingUnitField is a enum.BookingUnit field
type BookingUnitField struct{ quickfix.FIXString }

//Tag returns tag.BookingUnit (590)
func (f BookingUnitField) Tag() quickfix.Tag { return tag.BookingUnit }

func NewBookingUnit(val enum.BookingUnit) BookingUnitField {
	return BookingUnitField{quickfix.FIXString(val)}
}

func (f BookingUnitField) Value() enum.BookingUnit { return enum.BookingUnit(f.String()) }

//BrokerOfCreditField is a STRING field
type BrokerOfCreditField struct{ quickfix.FIXString }

//Tag returns tag.BrokerOfCredit (92)
func (f BrokerOfCreditField) Tag() quickfix.Tag { return tag.BrokerOfCredit }

//NewBrokerOfCredit returns a new BrokerOfCreditField initialized with val
func NewBrokerOfCredit(val string) BrokerOfCreditField {
	return BrokerOfCreditField{quickfix.FIXString(val)}
}

func (f BrokerOfCreditField) Value() string { return f.String() }

//BusinessRejectReasonField is a enum.BusinessRejectReason field
type BusinessRejectReasonField struct{ quickfix.FIXString }

//Tag returns tag.BusinessRejectReason (380)
func (f BusinessRejectReasonField) Tag() quickfix.Tag { return tag.BusinessRejectReason }

func NewBusinessRejectReason(val enum.BusinessRejectReason) BusinessRejectReasonField {
	return BusinessRejectReasonField{quickfix.FIXString(val)}
}

func (f BusinessRejectReasonField) Value() enum.BusinessRejectReason {
	return enum.BusinessRejectReason(f.String())
}

//BusinessRejectRefIDField is a STRING field
type BusinessRejectRefIDField struct{ quickfix.FIXString }

//Tag returns tag.BusinessRejectRefID (379)
func (f BusinessRejectRefIDField) Tag() quickfix.Tag { return tag.BusinessRejectRefID }

//NewBusinessRejectRefID returns a new BusinessRejectRefIDField initialized with val
func NewBusinessRejectRefID(val string) BusinessRejectRefIDField {
	return BusinessRejectRefIDField{quickfix.FIXString(val)}
}

func (f BusinessRejectRefIDField) Value() string { return f.String() }

//BuyVolumeField is a QTY field
type BuyVolumeField struct{ quickfix.FIXDecimal }

//Tag returns tag.BuyVolume (330)
func (f BuyVolumeField) Tag() quickfix.Tag { return tag.BuyVolume }

//NewBuyVolume returns a new BuyVolumeField initialized with val and scale
func NewBuyVolume(val decimal.Decimal, scale int32) BuyVolumeField {
	return BuyVolumeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f BuyVolumeField) Value() (val decimal.Decimal) { return f.Decimal }

//CFICodeField is a STRING field
type CFICodeField struct{ quickfix.FIXString }

//Tag returns tag.CFICode (461)
func (f CFICodeField) Tag() quickfix.Tag { return tag.CFICode }

//NewCFICode returns a new CFICodeField initialized with val
func NewCFICode(val string) CFICodeField {
	return CFICodeField{quickfix.FIXString(val)}
}

func (f CFICodeField) Value() string { return f.String() }

//CPProgramField is a enum.CPProgram field
type CPProgramField struct{ quickfix.FIXString }

//Tag returns tag.CPProgram (875)
func (f CPProgramField) Tag() quickfix.Tag { return tag.CPProgram }

func NewCPProgram(val enum.CPProgram) CPProgramField {
	return CPProgramField{quickfix.FIXString(val)}
}

func (f CPProgramField) Value() enum.CPProgram { return enum.CPProgram(f.String()) }

//CPRegTypeField is a STRING field
type CPRegTypeField struct{ quickfix.FIXString }

//Tag returns tag.CPRegType (876)
func (f CPRegTypeField) Tag() quickfix.Tag { return tag.CPRegType }

//NewCPRegType returns a new CPRegTypeField initialized with val
func NewCPRegType(val string) CPRegTypeField {
	return CPRegTypeField{quickfix.FIXString(val)}
}

func (f CPRegTypeField) Value() string { return f.String() }

//CalculatedCcyLastQtyField is a QTY field
type CalculatedCcyLastQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.CalculatedCcyLastQty (1056)
func (f CalculatedCcyLastQtyField) Tag() quickfix.Tag { return tag.CalculatedCcyLastQty }

//NewCalculatedCcyLastQty returns a new CalculatedCcyLastQtyField initialized with val and scale
func NewCalculatedCcyLastQty(val decimal.Decimal, scale int32) CalculatedCcyLastQtyField {
	return CalculatedCcyLastQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CalculatedCcyLastQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//CancellationRightsField is a enum.CancellationRights field
type CancellationRightsField struct{ quickfix.FIXString }

//Tag returns tag.CancellationRights (480)
func (f CancellationRightsField) Tag() quickfix.Tag { return tag.CancellationRights }

func NewCancellationRights(val enum.CancellationRights) CancellationRightsField {
	return CancellationRightsField{quickfix.FIXString(val)}
}

func (f CancellationRightsField) Value() enum.CancellationRights {
	return enum.CancellationRights(f.String())
}

//CapPriceField is a PRICE field
type CapPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.CapPrice (1199)
func (f CapPriceField) Tag() quickfix.Tag { return tag.CapPrice }

//NewCapPrice returns a new CapPriceField initialized with val and scale
func NewCapPrice(val decimal.Decimal, scale int32) CapPriceField {
	return CapPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CapPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//CardExpDateField is a LOCALMKTDATE field
type CardExpDateField struct{ quickfix.FIXString }

//Tag returns tag.CardExpDate (490)
func (f CardExpDateField) Tag() quickfix.Tag { return tag.CardExpDate }

//NewCardExpDate returns a new CardExpDateField initialized with val
func NewCardExpDate(val string) CardExpDateField {
	return CardExpDateField{quickfix.FIXString(val)}
}

func (f CardExpDateField) Value() string { return f.String() }

//CardHolderNameField is a STRING field
type CardHolderNameField struct{ quickfix.FIXString }

//Tag returns tag.CardHolderName (488)
func (f CardHolderNameField) Tag() quickfix.Tag { return tag.CardHolderName }

//NewCardHolderName returns a new CardHolderNameField initialized with val
func NewCardHolderName(val string) CardHolderNameField {
	return CardHolderNameField{quickfix.FIXString(val)}
}

func (f CardHolderNameField) Value() string { return f.String() }

//CardIssNoField is a STRING field
type CardIssNoField struct{ quickfix.FIXString }

//Tag returns tag.CardIssNo (491)
func (f CardIssNoField) Tag() quickfix.Tag { return tag.CardIssNo }

//NewCardIssNo returns a new CardIssNoField initialized with val
func NewCardIssNo(val string) CardIssNoField {
	return CardIssNoField{quickfix.FIXString(val)}
}

func (f CardIssNoField) Value() string { return f.String() }

//CardIssNumField is a STRING field
type CardIssNumField struct{ quickfix.FIXString }

//Tag returns tag.CardIssNum (491)
func (f CardIssNumField) Tag() quickfix.Tag { return tag.CardIssNum }

//NewCardIssNum returns a new CardIssNumField initialized with val
func NewCardIssNum(val string) CardIssNumField {
	return CardIssNumField{quickfix.FIXString(val)}
}

func (f CardIssNumField) Value() string { return f.String() }

//CardNumberField is a STRING field
type CardNumberField struct{ quickfix.FIXString }

//Tag returns tag.CardNumber (489)
func (f CardNumberField) Tag() quickfix.Tag { return tag.CardNumber }

//NewCardNumber returns a new CardNumberField initialized with val
func NewCardNumber(val string) CardNumberField {
	return CardNumberField{quickfix.FIXString(val)}
}

func (f CardNumberField) Value() string { return f.String() }

//CardStartDateField is a LOCALMKTDATE field
type CardStartDateField struct{ quickfix.FIXString }

//Tag returns tag.CardStartDate (503)
func (f CardStartDateField) Tag() quickfix.Tag { return tag.CardStartDate }

//NewCardStartDate returns a new CardStartDateField initialized with val
func NewCardStartDate(val string) CardStartDateField {
	return CardStartDateField{quickfix.FIXString(val)}
}

func (f CardStartDateField) Value() string { return f.String() }

//CashDistribAgentAcctNameField is a STRING field
type CashDistribAgentAcctNameField struct{ quickfix.FIXString }

//Tag returns tag.CashDistribAgentAcctName (502)
func (f CashDistribAgentAcctNameField) Tag() quickfix.Tag { return tag.CashDistribAgentAcctName }

//NewCashDistribAgentAcctName returns a new CashDistribAgentAcctNameField initialized with val
func NewCashDistribAgentAcctName(val string) CashDistribAgentAcctNameField {
	return CashDistribAgentAcctNameField{quickfix.FIXString(val)}
}

func (f CashDistribAgentAcctNameField) Value() string { return f.String() }

//CashDistribAgentAcctNumberField is a STRING field
type CashDistribAgentAcctNumberField struct{ quickfix.FIXString }

//Tag returns tag.CashDistribAgentAcctNumber (500)
func (f CashDistribAgentAcctNumberField) Tag() quickfix.Tag { return tag.CashDistribAgentAcctNumber }

//NewCashDistribAgentAcctNumber returns a new CashDistribAgentAcctNumberField initialized with val
func NewCashDistribAgentAcctNumber(val string) CashDistribAgentAcctNumberField {
	return CashDistribAgentAcctNumberField{quickfix.FIXString(val)}
}

func (f CashDistribAgentAcctNumberField) Value() string { return f.String() }

//CashDistribAgentCodeField is a STRING field
type CashDistribAgentCodeField struct{ quickfix.FIXString }

//Tag returns tag.CashDistribAgentCode (499)
func (f CashDistribAgentCodeField) Tag() quickfix.Tag { return tag.CashDistribAgentCode }

//NewCashDistribAgentCode returns a new CashDistribAgentCodeField initialized with val
func NewCashDistribAgentCode(val string) CashDistribAgentCodeField {
	return CashDistribAgentCodeField{quickfix.FIXString(val)}
}

func (f CashDistribAgentCodeField) Value() string { return f.String() }

//CashDistribAgentNameField is a STRING field
type CashDistribAgentNameField struct{ quickfix.FIXString }

//Tag returns tag.CashDistribAgentName (498)
func (f CashDistribAgentNameField) Tag() quickfix.Tag { return tag.CashDistribAgentName }

//NewCashDistribAgentName returns a new CashDistribAgentNameField initialized with val
func NewCashDistribAgentName(val string) CashDistribAgentNameField {
	return CashDistribAgentNameField{quickfix.FIXString(val)}
}

func (f CashDistribAgentNameField) Value() string { return f.String() }

//CashDistribCurrField is a CURRENCY field
type CashDistribCurrField struct{ quickfix.FIXString }

//Tag returns tag.CashDistribCurr (478)
func (f CashDistribCurrField) Tag() quickfix.Tag { return tag.CashDistribCurr }

//NewCashDistribCurr returns a new CashDistribCurrField initialized with val
func NewCashDistribCurr(val string) CashDistribCurrField {
	return CashDistribCurrField{quickfix.FIXString(val)}
}

func (f CashDistribCurrField) Value() string { return f.String() }

//CashDistribPayRefField is a STRING field
type CashDistribPayRefField struct{ quickfix.FIXString }

//Tag returns tag.CashDistribPayRef (501)
func (f CashDistribPayRefField) Tag() quickfix.Tag { return tag.CashDistribPayRef }

//NewCashDistribPayRef returns a new CashDistribPayRefField initialized with val
func NewCashDistribPayRef(val string) CashDistribPayRefField {
	return CashDistribPayRefField{quickfix.FIXString(val)}
}

func (f CashDistribPayRefField) Value() string { return f.String() }

//CashMarginField is a enum.CashMargin field
type CashMarginField struct{ quickfix.FIXString }

//Tag returns tag.CashMargin (544)
func (f CashMarginField) Tag() quickfix.Tag { return tag.CashMargin }

func NewCashMargin(val enum.CashMargin) CashMarginField {
	return CashMarginField{quickfix.FIXString(val)}
}

func (f CashMarginField) Value() enum.CashMargin { return enum.CashMargin(f.String()) }

//CashOrderQtyField is a QTY field
type CashOrderQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.CashOrderQty (152)
func (f CashOrderQtyField) Tag() quickfix.Tag { return tag.CashOrderQty }

//NewCashOrderQty returns a new CashOrderQtyField initialized with val and scale
func NewCashOrderQty(val decimal.Decimal, scale int32) CashOrderQtyField {
	return CashOrderQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CashOrderQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//CashOutstandingField is a AMT field
type CashOutstandingField struct{ quickfix.FIXDecimal }

//Tag returns tag.CashOutstanding (901)
func (f CashOutstandingField) Tag() quickfix.Tag { return tag.CashOutstanding }

//NewCashOutstanding returns a new CashOutstandingField initialized with val and scale
func NewCashOutstanding(val decimal.Decimal, scale int32) CashOutstandingField {
	return CashOutstandingField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CashOutstandingField) Value() (val decimal.Decimal) { return f.Decimal }

//CashSettlAgentAcctNameField is a STRING field
type CashSettlAgentAcctNameField struct{ quickfix.FIXString }

//Tag returns tag.CashSettlAgentAcctName (185)
func (f CashSettlAgentAcctNameField) Tag() quickfix.Tag { return tag.CashSettlAgentAcctName }

//NewCashSettlAgentAcctName returns a new CashSettlAgentAcctNameField initialized with val
func NewCashSettlAgentAcctName(val string) CashSettlAgentAcctNameField {
	return CashSettlAgentAcctNameField{quickfix.FIXString(val)}
}

func (f CashSettlAgentAcctNameField) Value() string { return f.String() }

//CashSettlAgentAcctNumField is a STRING field
type CashSettlAgentAcctNumField struct{ quickfix.FIXString }

//Tag returns tag.CashSettlAgentAcctNum (184)
func (f CashSettlAgentAcctNumField) Tag() quickfix.Tag { return tag.CashSettlAgentAcctNum }

//NewCashSettlAgentAcctNum returns a new CashSettlAgentAcctNumField initialized with val
func NewCashSettlAgentAcctNum(val string) CashSettlAgentAcctNumField {
	return CashSettlAgentAcctNumField{quickfix.FIXString(val)}
}

func (f CashSettlAgentAcctNumField) Value() string { return f.String() }

//CashSettlAgentCodeField is a STRING field
type CashSettlAgentCodeField struct{ quickfix.FIXString }

//Tag returns tag.CashSettlAgentCode (183)
func (f CashSettlAgentCodeField) Tag() quickfix.Tag { return tag.CashSettlAgentCode }

//NewCashSettlAgentCode returns a new CashSettlAgentCodeField initialized with val
func NewCashSettlAgentCode(val string) CashSettlAgentCodeField {
	return CashSettlAgentCodeField{quickfix.FIXString(val)}
}

func (f CashSettlAgentCodeField) Value() string { return f.String() }

//CashSettlAgentContactNameField is a STRING field
type CashSettlAgentContactNameField struct{ quickfix.FIXString }

//Tag returns tag.CashSettlAgentContactName (186)
func (f CashSettlAgentContactNameField) Tag() quickfix.Tag { return tag.CashSettlAgentContactName }

//NewCashSettlAgentContactName returns a new CashSettlAgentContactNameField initialized with val
func NewCashSettlAgentContactName(val string) CashSettlAgentContactNameField {
	return CashSettlAgentContactNameField{quickfix.FIXString(val)}
}

func (f CashSettlAgentContactNameField) Value() string { return f.String() }

//CashSettlAgentContactPhoneField is a STRING field
type CashSettlAgentContactPhoneField struct{ quickfix.FIXString }

//Tag returns tag.CashSettlAgentContactPhone (187)
func (f CashSettlAgentContactPhoneField) Tag() quickfix.Tag { return tag.CashSettlAgentContactPhone }

//NewCashSettlAgentContactPhone returns a new CashSettlAgentContactPhoneField initialized with val
func NewCashSettlAgentContactPhone(val string) CashSettlAgentContactPhoneField {
	return CashSettlAgentContactPhoneField{quickfix.FIXString(val)}
}

func (f CashSettlAgentContactPhoneField) Value() string { return f.String() }

//CashSettlAgentNameField is a STRING field
type CashSettlAgentNameField struct{ quickfix.FIXString }

//Tag returns tag.CashSettlAgentName (182)
func (f CashSettlAgentNameField) Tag() quickfix.Tag { return tag.CashSettlAgentName }

//NewCashSettlAgentName returns a new CashSettlAgentNameField initialized with val
func NewCashSettlAgentName(val string) CashSettlAgentNameField {
	return CashSettlAgentNameField{quickfix.FIXString(val)}
}

func (f CashSettlAgentNameField) Value() string { return f.String() }

//CcyAmtField is a AMT field
type CcyAmtField struct{ quickfix.FIXDecimal }

//Tag returns tag.CcyAmt (1157)
func (f CcyAmtField) Tag() quickfix.Tag { return tag.CcyAmt }

//NewCcyAmt returns a new CcyAmtField initialized with val and scale
func NewCcyAmt(val decimal.Decimal, scale int32) CcyAmtField {
	return CcyAmtField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CcyAmtField) Value() (val decimal.Decimal) { return f.Decimal }

//CheckSumField is a STRING field
type CheckSumField struct{ quickfix.FIXString }

//Tag returns tag.CheckSum (10)
func (f CheckSumField) Tag() quickfix.Tag { return tag.CheckSum }

//NewCheckSum returns a new CheckSumField initialized with val
func NewCheckSum(val string) CheckSumField {
	return CheckSumField{quickfix.FIXString(val)}
}

func (f CheckSumField) Value() string { return f.String() }

//ClOrdIDField is a STRING field
type ClOrdIDField struct{ quickfix.FIXString }

//Tag returns tag.ClOrdID (11)
func (f ClOrdIDField) Tag() quickfix.Tag { return tag.ClOrdID }

//NewClOrdID returns a new ClOrdIDField initialized with val
func NewClOrdID(val string) ClOrdIDField {
	return ClOrdIDField{quickfix.FIXString(val)}
}

func (f ClOrdIDField) Value() string { return f.String() }

//ClOrdLinkIDField is a STRING field
type ClOrdLinkIDField struct{ quickfix.FIXString }

//Tag returns tag.ClOrdLinkID (583)
func (f ClOrdLinkIDField) Tag() quickfix.Tag { return tag.ClOrdLinkID }

//NewClOrdLinkID returns a new ClOrdLinkIDField initialized with val
func NewClOrdLinkID(val string) ClOrdLinkIDField {
	return ClOrdLinkIDField{quickfix.FIXString(val)}
}

func (f ClOrdLinkIDField) Value() string { return f.String() }

//ClearingAccountField is a STRING field
type ClearingAccountField struct{ quickfix.FIXString }

//Tag returns tag.ClearingAccount (440)
func (f ClearingAccountField) Tag() quickfix.Tag { return tag.ClearingAccount }

//NewClearingAccount returns a new ClearingAccountField initialized with val
func NewClearingAccount(val string) ClearingAccountField {
	return ClearingAccountField{quickfix.FIXString(val)}
}

func (f ClearingAccountField) Value() string { return f.String() }

//ClearingBusinessDateField is a LOCALMKTDATE field
type ClearingBusinessDateField struct{ quickfix.FIXString }

//Tag returns tag.ClearingBusinessDate (715)
func (f ClearingBusinessDateField) Tag() quickfix.Tag { return tag.ClearingBusinessDate }

//NewClearingBusinessDate returns a new ClearingBusinessDateField initialized with val
func NewClearingBusinessDate(val string) ClearingBusinessDateField {
	return ClearingBusinessDateField{quickfix.FIXString(val)}
}

func (f ClearingBusinessDateField) Value() string { return f.String() }

//ClearingFeeIndicatorField is a enum.ClearingFeeIndicator field
type ClearingFeeIndicatorField struct{ quickfix.FIXString }

//Tag returns tag.ClearingFeeIndicator (635)
func (f ClearingFeeIndicatorField) Tag() quickfix.Tag { return tag.ClearingFeeIndicator }

func NewClearingFeeIndicator(val enum.ClearingFeeIndicator) ClearingFeeIndicatorField {
	return ClearingFeeIndicatorField{quickfix.FIXString(val)}
}

func (f ClearingFeeIndicatorField) Value() enum.ClearingFeeIndicator {
	return enum.ClearingFeeIndicator(f.String())
}

//ClearingFirmField is a STRING field
type ClearingFirmField struct{ quickfix.FIXString }

//Tag returns tag.ClearingFirm (439)
func (f ClearingFirmField) Tag() quickfix.Tag { return tag.ClearingFirm }

//NewClearingFirm returns a new ClearingFirmField initialized with val
func NewClearingFirm(val string) ClearingFirmField {
	return ClearingFirmField{quickfix.FIXString(val)}
}

func (f ClearingFirmField) Value() string { return f.String() }

//ClearingInstructionField is a enum.ClearingInstruction field
type ClearingInstructionField struct{ quickfix.FIXString }

//Tag returns tag.ClearingInstruction (577)
func (f ClearingInstructionField) Tag() quickfix.Tag { return tag.ClearingInstruction }

func NewClearingInstruction(val enum.ClearingInstruction) ClearingInstructionField {
	return ClearingInstructionField{quickfix.FIXString(val)}
}

func (f ClearingInstructionField) Value() enum.ClearingInstruction {
	return enum.ClearingInstruction(f.String())
}

//ClientBidIDField is a STRING field
type ClientBidIDField struct{ quickfix.FIXString }

//Tag returns tag.ClientBidID (391)
func (f ClientBidIDField) Tag() quickfix.Tag { return tag.ClientBidID }

//NewClientBidID returns a new ClientBidIDField initialized with val
func NewClientBidID(val string) ClientBidIDField {
	return ClientBidIDField{quickfix.FIXString(val)}
}

func (f ClientBidIDField) Value() string { return f.String() }

//ClientIDField is a STRING field
type ClientIDField struct{ quickfix.FIXString }

//Tag returns tag.ClientID (109)
func (f ClientIDField) Tag() quickfix.Tag { return tag.ClientID }

//NewClientID returns a new ClientIDField initialized with val
func NewClientID(val string) ClientIDField {
	return ClientIDField{quickfix.FIXString(val)}
}

func (f ClientIDField) Value() string { return f.String() }

//CollActionField is a enum.CollAction field
type CollActionField struct{ quickfix.FIXString }

//Tag returns tag.CollAction (944)
func (f CollActionField) Tag() quickfix.Tag { return tag.CollAction }

func NewCollAction(val enum.CollAction) CollActionField {
	return CollActionField{quickfix.FIXString(val)}
}

func (f CollActionField) Value() enum.CollAction { return enum.CollAction(f.String()) }

//CollApplTypeField is a enum.CollApplType field
type CollApplTypeField struct{ quickfix.FIXString }

//Tag returns tag.CollApplType (1043)
func (f CollApplTypeField) Tag() quickfix.Tag { return tag.CollApplType }

func NewCollApplType(val enum.CollApplType) CollApplTypeField {
	return CollApplTypeField{quickfix.FIXString(val)}
}

func (f CollApplTypeField) Value() enum.CollApplType { return enum.CollApplType(f.String()) }

//CollAsgnIDField is a STRING field
type CollAsgnIDField struct{ quickfix.FIXString }

//Tag returns tag.CollAsgnID (902)
func (f CollAsgnIDField) Tag() quickfix.Tag { return tag.CollAsgnID }

//NewCollAsgnID returns a new CollAsgnIDField initialized with val
func NewCollAsgnID(val string) CollAsgnIDField {
	return CollAsgnIDField{quickfix.FIXString(val)}
}

func (f CollAsgnIDField) Value() string { return f.String() }

//CollAsgnReasonField is a enum.CollAsgnReason field
type CollAsgnReasonField struct{ quickfix.FIXString }

//Tag returns tag.CollAsgnReason (895)
func (f CollAsgnReasonField) Tag() quickfix.Tag { return tag.CollAsgnReason }

func NewCollAsgnReason(val enum.CollAsgnReason) CollAsgnReasonField {
	return CollAsgnReasonField{quickfix.FIXString(val)}
}

func (f CollAsgnReasonField) Value() enum.CollAsgnReason { return enum.CollAsgnReason(f.String()) }

//CollAsgnRefIDField is a STRING field
type CollAsgnRefIDField struct{ quickfix.FIXString }

//Tag returns tag.CollAsgnRefID (907)
func (f CollAsgnRefIDField) Tag() quickfix.Tag { return tag.CollAsgnRefID }

//NewCollAsgnRefID returns a new CollAsgnRefIDField initialized with val
func NewCollAsgnRefID(val string) CollAsgnRefIDField {
	return CollAsgnRefIDField{quickfix.FIXString(val)}
}

func (f CollAsgnRefIDField) Value() string { return f.String() }

//CollAsgnRejectReasonField is a enum.CollAsgnRejectReason field
type CollAsgnRejectReasonField struct{ quickfix.FIXString }

//Tag returns tag.CollAsgnRejectReason (906)
func (f CollAsgnRejectReasonField) Tag() quickfix.Tag { return tag.CollAsgnRejectReason }

func NewCollAsgnRejectReason(val enum.CollAsgnRejectReason) CollAsgnRejectReasonField {
	return CollAsgnRejectReasonField{quickfix.FIXString(val)}
}

func (f CollAsgnRejectReasonField) Value() enum.CollAsgnRejectReason {
	return enum.CollAsgnRejectReason(f.String())
}

//CollAsgnRespTypeField is a enum.CollAsgnRespType field
type CollAsgnRespTypeField struct{ quickfix.FIXString }

//Tag returns tag.CollAsgnRespType (905)
func (f CollAsgnRespTypeField) Tag() quickfix.Tag { return tag.CollAsgnRespType }

func NewCollAsgnRespType(val enum.CollAsgnRespType) CollAsgnRespTypeField {
	return CollAsgnRespTypeField{quickfix.FIXString(val)}
}

func (f CollAsgnRespTypeField) Value() enum.CollAsgnRespType { return enum.CollAsgnRespType(f.String()) }

//CollAsgnTransTypeField is a enum.CollAsgnTransType field
type CollAsgnTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.CollAsgnTransType (903)
func (f CollAsgnTransTypeField) Tag() quickfix.Tag { return tag.CollAsgnTransType }

func NewCollAsgnTransType(val enum.CollAsgnTransType) CollAsgnTransTypeField {
	return CollAsgnTransTypeField{quickfix.FIXString(val)}
}

func (f CollAsgnTransTypeField) Value() enum.CollAsgnTransType {
	return enum.CollAsgnTransType(f.String())
}

//CollInquiryIDField is a STRING field
type CollInquiryIDField struct{ quickfix.FIXString }

//Tag returns tag.CollInquiryID (909)
func (f CollInquiryIDField) Tag() quickfix.Tag { return tag.CollInquiryID }

//NewCollInquiryID returns a new CollInquiryIDField initialized with val
func NewCollInquiryID(val string) CollInquiryIDField {
	return CollInquiryIDField{quickfix.FIXString(val)}
}

func (f CollInquiryIDField) Value() string { return f.String() }

//CollInquiryQualifierField is a enum.CollInquiryQualifier field
type CollInquiryQualifierField struct{ quickfix.FIXString }

//Tag returns tag.CollInquiryQualifier (896)
func (f CollInquiryQualifierField) Tag() quickfix.Tag { return tag.CollInquiryQualifier }

func NewCollInquiryQualifier(val enum.CollInquiryQualifier) CollInquiryQualifierField {
	return CollInquiryQualifierField{quickfix.FIXString(val)}
}

func (f CollInquiryQualifierField) Value() enum.CollInquiryQualifier {
	return enum.CollInquiryQualifier(f.String())
}

//CollInquiryResultField is a enum.CollInquiryResult field
type CollInquiryResultField struct{ quickfix.FIXString }

//Tag returns tag.CollInquiryResult (946)
func (f CollInquiryResultField) Tag() quickfix.Tag { return tag.CollInquiryResult }

func NewCollInquiryResult(val enum.CollInquiryResult) CollInquiryResultField {
	return CollInquiryResultField{quickfix.FIXString(val)}
}

func (f CollInquiryResultField) Value() enum.CollInquiryResult {
	return enum.CollInquiryResult(f.String())
}

//CollInquiryStatusField is a enum.CollInquiryStatus field
type CollInquiryStatusField struct{ quickfix.FIXString }

//Tag returns tag.CollInquiryStatus (945)
func (f CollInquiryStatusField) Tag() quickfix.Tag { return tag.CollInquiryStatus }

func NewCollInquiryStatus(val enum.CollInquiryStatus) CollInquiryStatusField {
	return CollInquiryStatusField{quickfix.FIXString(val)}
}

func (f CollInquiryStatusField) Value() enum.CollInquiryStatus {
	return enum.CollInquiryStatus(f.String())
}

//CollReqIDField is a STRING field
type CollReqIDField struct{ quickfix.FIXString }

//Tag returns tag.CollReqID (894)
func (f CollReqIDField) Tag() quickfix.Tag { return tag.CollReqID }

//NewCollReqID returns a new CollReqIDField initialized with val
func NewCollReqID(val string) CollReqIDField {
	return CollReqIDField{quickfix.FIXString(val)}
}

func (f CollReqIDField) Value() string { return f.String() }

//CollRespIDField is a STRING field
type CollRespIDField struct{ quickfix.FIXString }

//Tag returns tag.CollRespID (904)
func (f CollRespIDField) Tag() quickfix.Tag { return tag.CollRespID }

//NewCollRespID returns a new CollRespIDField initialized with val
func NewCollRespID(val string) CollRespIDField {
	return CollRespIDField{quickfix.FIXString(val)}
}

func (f CollRespIDField) Value() string { return f.String() }

//CollRptIDField is a STRING field
type CollRptIDField struct{ quickfix.FIXString }

//Tag returns tag.CollRptID (908)
func (f CollRptIDField) Tag() quickfix.Tag { return tag.CollRptID }

//NewCollRptID returns a new CollRptIDField initialized with val
func NewCollRptID(val string) CollRptIDField {
	return CollRptIDField{quickfix.FIXString(val)}
}

func (f CollRptIDField) Value() string { return f.String() }

//CollStatusField is a enum.CollStatus field
type CollStatusField struct{ quickfix.FIXString }

//Tag returns tag.CollStatus (910)
func (f CollStatusField) Tag() quickfix.Tag { return tag.CollStatus }

func NewCollStatus(val enum.CollStatus) CollStatusField {
	return CollStatusField{quickfix.FIXString(val)}
}

func (f CollStatusField) Value() enum.CollStatus { return enum.CollStatus(f.String()) }

//CommCurrencyField is a CURRENCY field
type CommCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.CommCurrency (479)
func (f CommCurrencyField) Tag() quickfix.Tag { return tag.CommCurrency }

//NewCommCurrency returns a new CommCurrencyField initialized with val
func NewCommCurrency(val string) CommCurrencyField {
	return CommCurrencyField{quickfix.FIXString(val)}
}

func (f CommCurrencyField) Value() string { return f.String() }

//CommTypeField is a enum.CommType field
type CommTypeField struct{ quickfix.FIXString }

//Tag returns tag.CommType (13)
func (f CommTypeField) Tag() quickfix.Tag { return tag.CommType }

func NewCommType(val enum.CommType) CommTypeField {
	return CommTypeField{quickfix.FIXString(val)}
}

func (f CommTypeField) Value() enum.CommType { return enum.CommType(f.String()) }

//CommissionField is a AMT field
type CommissionField struct{ quickfix.FIXDecimal }

//Tag returns tag.Commission (12)
func (f CommissionField) Tag() quickfix.Tag { return tag.Commission }

//NewCommission returns a new CommissionField initialized with val and scale
func NewCommission(val decimal.Decimal, scale int32) CommissionField {
	return CommissionField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CommissionField) Value() (val decimal.Decimal) { return f.Decimal }

//ComplexEventConditionField is a enum.ComplexEventCondition field
type ComplexEventConditionField struct{ quickfix.FIXString }

//Tag returns tag.ComplexEventCondition (1490)
func (f ComplexEventConditionField) Tag() quickfix.Tag { return tag.ComplexEventCondition }

func NewComplexEventCondition(val enum.ComplexEventCondition) ComplexEventConditionField {
	return ComplexEventConditionField{quickfix.FIXString(val)}
}

func (f ComplexEventConditionField) Value() enum.ComplexEventCondition {
	return enum.ComplexEventCondition(f.String())
}

//ComplexEventEndDateField is a UTCTIMESTAMP field
type ComplexEventEndDateField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.ComplexEventEndDate (1493)
func (f ComplexEventEndDateField) Tag() quickfix.Tag { return tag.ComplexEventEndDate }

//NewComplexEventEndDate returns a new ComplexEventEndDateField initialized with val
func NewComplexEventEndDate(val time.Time) ComplexEventEndDateField {
	return ComplexEventEndDateField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewComplexEventEndDateNoMillis returns a new ComplexEventEndDateField initialized with val without millisecs
func NewComplexEventEndDateNoMillis(val time.Time) ComplexEventEndDateField {
	return ComplexEventEndDateField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f ComplexEventEndDateField) Value() time.Time { return f.Time }

//ComplexEventEndTimeField is a UTCTIMEONLY field
type ComplexEventEndTimeField struct{ quickfix.FIXString }

//Tag returns tag.ComplexEventEndTime (1496)
func (f ComplexEventEndTimeField) Tag() quickfix.Tag { return tag.ComplexEventEndTime }

//NewComplexEventEndTime returns a new ComplexEventEndTimeField initialized with val
func NewComplexEventEndTime(val string) ComplexEventEndTimeField {
	return ComplexEventEndTimeField{quickfix.FIXString(val)}
}

func (f ComplexEventEndTimeField) Value() string { return f.String() }

//ComplexEventPriceField is a PRICE field
type ComplexEventPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.ComplexEventPrice (1486)
func (f ComplexEventPriceField) Tag() quickfix.Tag { return tag.ComplexEventPrice }

//NewComplexEventPrice returns a new ComplexEventPriceField initialized with val and scale
func NewComplexEventPrice(val decimal.Decimal, scale int32) ComplexEventPriceField {
	return ComplexEventPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f ComplexEventPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//ComplexEventPriceBoundaryMethodField is a enum.ComplexEventPriceBoundaryMethod field
type ComplexEventPriceBoundaryMethodField struct{ quickfix.FIXString }

//Tag returns tag.ComplexEventPriceBoundaryMethod (1487)
func (f ComplexEventPriceBoundaryMethodField) Tag() quickfix.Tag {
	return tag.ComplexEventPriceBoundaryMethod
}

func NewComplexEventPriceBoundaryMethod(val enum.ComplexEventPriceBoundaryMethod) ComplexEventPriceBoundaryMethodField {
	return ComplexEventPriceBoundaryMethodField{quickfix.FIXString(val)}
}

func (f ComplexEventPriceBoundaryMethodField) Value() enum.ComplexEventPriceBoundaryMethod {
	return enum.ComplexEventPriceBoundaryMethod(f.String())
}

//ComplexEventPriceBoundaryPrecisionField is a PERCENTAGE field
type ComplexEventPriceBoundaryPrecisionField struct{ quickfix.FIXDecimal }

//Tag returns tag.ComplexEventPriceBoundaryPrecision (1488)
func (f ComplexEventPriceBoundaryPrecisionField) Tag() quickfix.Tag {
	return tag.ComplexEventPriceBoundaryPrecision
}

//NewComplexEventPriceBoundaryPrecision returns a new ComplexEventPriceBoundaryPrecisionField initialized with val and scale
func NewComplexEventPriceBoundaryPrecision(val decimal.Decimal, scale int32) ComplexEventPriceBoundaryPrecisionField {
	return ComplexEventPriceBoundaryPrecisionField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f ComplexEventPriceBoundaryPrecisionField) Value() (val decimal.Decimal) { return f.Decimal }

//ComplexEventPriceTimeTypeField is a enum.ComplexEventPriceTimeType field
type ComplexEventPriceTimeTypeField struct{ quickfix.FIXString }

//Tag returns tag.ComplexEventPriceTimeType (1489)
func (f ComplexEventPriceTimeTypeField) Tag() quickfix.Tag { return tag.ComplexEventPriceTimeType }

func NewComplexEventPriceTimeType(val enum.ComplexEventPriceTimeType) ComplexEventPriceTimeTypeField {
	return ComplexEventPriceTimeTypeField{quickfix.FIXString(val)}
}

func (f ComplexEventPriceTimeTypeField) Value() enum.ComplexEventPriceTimeType {
	return enum.ComplexEventPriceTimeType(f.String())
}

//ComplexEventStartDateField is a UTCTIMESTAMP field
type ComplexEventStartDateField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.ComplexEventStartDate (1492)
func (f ComplexEventStartDateField) Tag() quickfix.Tag { return tag.ComplexEventStartDate }

//NewComplexEventStartDate returns a new ComplexEventStartDateField initialized with val
func NewComplexEventStartDate(val time.Time) ComplexEventStartDateField {
	return ComplexEventStartDateField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewComplexEventStartDateNoMillis returns a new ComplexEventStartDateField initialized with val without millisecs
func NewComplexEventStartDateNoMillis(val time.Time) ComplexEventStartDateField {
	return ComplexEventStartDateField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f ComplexEventStartDateField) Value() time.Time { return f.Time }

//ComplexEventStartTimeField is a UTCTIMEONLY field
type ComplexEventStartTimeField struct{ quickfix.FIXString }

//Tag returns tag.ComplexEventStartTime (1495)
func (f ComplexEventStartTimeField) Tag() quickfix.Tag { return tag.ComplexEventStartTime }

//NewComplexEventStartTime returns a new ComplexEventStartTimeField initialized with val
func NewComplexEventStartTime(val string) ComplexEventStartTimeField {
	return ComplexEventStartTimeField{quickfix.FIXString(val)}
}

func (f ComplexEventStartTimeField) Value() string { return f.String() }

//ComplexEventTypeField is a enum.ComplexEventType field
type ComplexEventTypeField struct{ quickfix.FIXString }

//Tag returns tag.ComplexEventType (1484)
func (f ComplexEventTypeField) Tag() quickfix.Tag { return tag.ComplexEventType }

func NewComplexEventType(val enum.ComplexEventType) ComplexEventTypeField {
	return ComplexEventTypeField{quickfix.FIXString(val)}
}

func (f ComplexEventTypeField) Value() enum.ComplexEventType { return enum.ComplexEventType(f.String()) }

//ComplexOptPayoutAmountField is a AMT field
type ComplexOptPayoutAmountField struct{ quickfix.FIXDecimal }

//Tag returns tag.ComplexOptPayoutAmount (1485)
func (f ComplexOptPayoutAmountField) Tag() quickfix.Tag { return tag.ComplexOptPayoutAmount }

//NewComplexOptPayoutAmount returns a new ComplexOptPayoutAmountField initialized with val and scale
func NewComplexOptPayoutAmount(val decimal.Decimal, scale int32) ComplexOptPayoutAmountField {
	return ComplexOptPayoutAmountField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f ComplexOptPayoutAmountField) Value() (val decimal.Decimal) { return f.Decimal }

//ComplianceIDField is a STRING field
type ComplianceIDField struct{ quickfix.FIXString }

//Tag returns tag.ComplianceID (376)
func (f ComplianceIDField) Tag() quickfix.Tag { return tag.ComplianceID }

//NewComplianceID returns a new ComplianceIDField initialized with val
func NewComplianceID(val string) ComplianceIDField {
	return ComplianceIDField{quickfix.FIXString(val)}
}

func (f ComplianceIDField) Value() string { return f.String() }

//ConcessionField is a AMT field
type ConcessionField struct{ quickfix.FIXDecimal }

//Tag returns tag.Concession (238)
func (f ConcessionField) Tag() quickfix.Tag { return tag.Concession }

//NewConcession returns a new ConcessionField initialized with val and scale
func NewConcession(val decimal.Decimal, scale int32) ConcessionField {
	return ConcessionField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f ConcessionField) Value() (val decimal.Decimal) { return f.Decimal }

//ConfirmIDField is a STRING field
type ConfirmIDField struct{ quickfix.FIXString }

//Tag returns tag.ConfirmID (664)
func (f ConfirmIDField) Tag() quickfix.Tag { return tag.ConfirmID }

//NewConfirmID returns a new ConfirmIDField initialized with val
func NewConfirmID(val string) ConfirmIDField {
	return ConfirmIDField{quickfix.FIXString(val)}
}

func (f ConfirmIDField) Value() string { return f.String() }

//ConfirmRefIDField is a STRING field
type ConfirmRefIDField struct{ quickfix.FIXString }

//Tag returns tag.ConfirmRefID (772)
func (f ConfirmRefIDField) Tag() quickfix.Tag { return tag.ConfirmRefID }

//NewConfirmRefID returns a new ConfirmRefIDField initialized with val
func NewConfirmRefID(val string) ConfirmRefIDField {
	return ConfirmRefIDField{quickfix.FIXString(val)}
}

func (f ConfirmRefIDField) Value() string { return f.String() }

//ConfirmRejReasonField is a enum.ConfirmRejReason field
type ConfirmRejReasonField struct{ quickfix.FIXString }

//Tag returns tag.ConfirmRejReason (774)
func (f ConfirmRejReasonField) Tag() quickfix.Tag { return tag.ConfirmRejReason }

func NewConfirmRejReason(val enum.ConfirmRejReason) ConfirmRejReasonField {
	return ConfirmRejReasonField{quickfix.FIXString(val)}
}

func (f ConfirmRejReasonField) Value() enum.ConfirmRejReason { return enum.ConfirmRejReason(f.String()) }

//ConfirmReqIDField is a STRING field
type ConfirmReqIDField struct{ quickfix.FIXString }

//Tag returns tag.ConfirmReqID (859)
func (f ConfirmReqIDField) Tag() quickfix.Tag { return tag.ConfirmReqID }

//NewConfirmReqID returns a new ConfirmReqIDField initialized with val
func NewConfirmReqID(val string) ConfirmReqIDField {
	return ConfirmReqIDField{quickfix.FIXString(val)}
}

func (f ConfirmReqIDField) Value() string { return f.String() }

//ConfirmStatusField is a enum.ConfirmStatus field
type ConfirmStatusField struct{ quickfix.FIXString }

//Tag returns tag.ConfirmStatus (665)
func (f ConfirmStatusField) Tag() quickfix.Tag { return tag.ConfirmStatus }

func NewConfirmStatus(val enum.ConfirmStatus) ConfirmStatusField {
	return ConfirmStatusField{quickfix.FIXString(val)}
}

func (f ConfirmStatusField) Value() enum.ConfirmStatus { return enum.ConfirmStatus(f.String()) }

//ConfirmTransTypeField is a enum.ConfirmTransType field
type ConfirmTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.ConfirmTransType (666)
func (f ConfirmTransTypeField) Tag() quickfix.Tag { return tag.ConfirmTransType }

func NewConfirmTransType(val enum.ConfirmTransType) ConfirmTransTypeField {
	return ConfirmTransTypeField{quickfix.FIXString(val)}
}

func (f ConfirmTransTypeField) Value() enum.ConfirmTransType { return enum.ConfirmTransType(f.String()) }

//ConfirmTypeField is a enum.ConfirmType field
type ConfirmTypeField struct{ quickfix.FIXString }

//Tag returns tag.ConfirmType (773)
func (f ConfirmTypeField) Tag() quickfix.Tag { return tag.ConfirmType }

func NewConfirmType(val enum.ConfirmType) ConfirmTypeField {
	return ConfirmTypeField{quickfix.FIXString(val)}
}

func (f ConfirmTypeField) Value() enum.ConfirmType { return enum.ConfirmType(f.String()) }

//ContAmtCurrField is a CURRENCY field
type ContAmtCurrField struct{ quickfix.FIXString }

//Tag returns tag.ContAmtCurr (521)
func (f ContAmtCurrField) Tag() quickfix.Tag { return tag.ContAmtCurr }

//NewContAmtCurr returns a new ContAmtCurrField initialized with val
func NewContAmtCurr(val string) ContAmtCurrField {
	return ContAmtCurrField{quickfix.FIXString(val)}
}

func (f ContAmtCurrField) Value() string { return f.String() }

//ContAmtTypeField is a enum.ContAmtType field
type ContAmtTypeField struct{ quickfix.FIXString }

//Tag returns tag.ContAmtType (519)
func (f ContAmtTypeField) Tag() quickfix.Tag { return tag.ContAmtType }

func NewContAmtType(val enum.ContAmtType) ContAmtTypeField {
	return ContAmtTypeField{quickfix.FIXString(val)}
}

func (f ContAmtTypeField) Value() enum.ContAmtType { return enum.ContAmtType(f.String()) }

//ContAmtValueField is a FLOAT field
type ContAmtValueField struct{ quickfix.FIXDecimal }

//Tag returns tag.ContAmtValue (520)
func (f ContAmtValueField) Tag() quickfix.Tag { return tag.ContAmtValue }

//NewContAmtValue returns a new ContAmtValueField initialized with val and scale
func NewContAmtValue(val decimal.Decimal, scale int32) ContAmtValueField {
	return ContAmtValueField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f ContAmtValueField) Value() (val decimal.Decimal) { return f.Decimal }

//ContIntRptIDField is a STRING field
type ContIntRptIDField struct{ quickfix.FIXString }

//Tag returns tag.ContIntRptID (977)
func (f ContIntRptIDField) Tag() quickfix.Tag { return tag.ContIntRptID }

//NewContIntRptID returns a new ContIntRptIDField initialized with val
func NewContIntRptID(val string) ContIntRptIDField {
	return ContIntRptIDField{quickfix.FIXString(val)}
}

func (f ContIntRptIDField) Value() string { return f.String() }

//ContextPartyIDField is a STRING field
type ContextPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.ContextPartyID (1523)
func (f ContextPartyIDField) Tag() quickfix.Tag { return tag.ContextPartyID }

//NewContextPartyID returns a new ContextPartyIDField initialized with val
func NewContextPartyID(val string) ContextPartyIDField {
	return ContextPartyIDField{quickfix.FIXString(val)}
}

func (f ContextPartyIDField) Value() string { return f.String() }

//ContextPartyIDSourceField is a CHAR field
type ContextPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.ContextPartyIDSource (1524)
func (f ContextPartyIDSourceField) Tag() quickfix.Tag { return tag.ContextPartyIDSource }

//NewContextPartyIDSource returns a new ContextPartyIDSourceField initialized with val
func NewContextPartyIDSource(val string) ContextPartyIDSourceField {
	return ContextPartyIDSourceField{quickfix.FIXString(val)}
}

func (f ContextPartyIDSourceField) Value() string { return f.String() }

//ContextPartyRoleField is a INT field
type ContextPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.ContextPartyRole (1525)
func (f ContextPartyRoleField) Tag() quickfix.Tag { return tag.ContextPartyRole }

//NewContextPartyRole returns a new ContextPartyRoleField initialized with val
func NewContextPartyRole(val int) ContextPartyRoleField {
	return ContextPartyRoleField{quickfix.FIXInt(val)}
}

func (f ContextPartyRoleField) Value() int { return f.Int() }

//ContextPartySubIDField is a STRING field
type ContextPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.ContextPartySubID (1527)
func (f ContextPartySubIDField) Tag() quickfix.Tag { return tag.ContextPartySubID }

//NewContextPartySubID returns a new ContextPartySubIDField initialized with val
func NewContextPartySubID(val string) ContextPartySubIDField {
	return ContextPartySubIDField{quickfix.FIXString(val)}
}

func (f ContextPartySubIDField) Value() string { return f.String() }

//ContextPartySubIDTypeField is a INT field
type ContextPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.ContextPartySubIDType (1528)
func (f ContextPartySubIDTypeField) Tag() quickfix.Tag { return tag.ContextPartySubIDType }

//NewContextPartySubIDType returns a new ContextPartySubIDTypeField initialized with val
func NewContextPartySubIDType(val int) ContextPartySubIDTypeField {
	return ContextPartySubIDTypeField{quickfix.FIXInt(val)}
}

func (f ContextPartySubIDTypeField) Value() int { return f.Int() }

//ContingencyTypeField is a enum.ContingencyType field
type ContingencyTypeField struct{ quickfix.FIXString }

//Tag returns tag.ContingencyType (1385)
func (f ContingencyTypeField) Tag() quickfix.Tag { return tag.ContingencyType }

func NewContingencyType(val enum.ContingencyType) ContingencyTypeField {
	return ContingencyTypeField{quickfix.FIXString(val)}
}

func (f ContingencyTypeField) Value() enum.ContingencyType { return enum.ContingencyType(f.String()) }

//ContraBrokerField is a STRING field
type ContraBrokerField struct{ quickfix.FIXString }

//Tag returns tag.ContraBroker (375)
func (f ContraBrokerField) Tag() quickfix.Tag { return tag.ContraBroker }

//NewContraBroker returns a new ContraBrokerField initialized with val
func NewContraBroker(val string) ContraBrokerField {
	return ContraBrokerField{quickfix.FIXString(val)}
}

func (f ContraBrokerField) Value() string { return f.String() }

//ContraLegRefIDField is a STRING field
type ContraLegRefIDField struct{ quickfix.FIXString }

//Tag returns tag.ContraLegRefID (655)
func (f ContraLegRefIDField) Tag() quickfix.Tag { return tag.ContraLegRefID }

//NewContraLegRefID returns a new ContraLegRefIDField initialized with val
func NewContraLegRefID(val string) ContraLegRefIDField {
	return ContraLegRefIDField{quickfix.FIXString(val)}
}

func (f ContraLegRefIDField) Value() string { return f.String() }

//ContraTradeQtyField is a QTY field
type ContraTradeQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.ContraTradeQty (437)
func (f ContraTradeQtyField) Tag() quickfix.Tag { return tag.ContraTradeQty }

//NewContraTradeQty returns a new ContraTradeQtyField initialized with val and scale
func NewContraTradeQty(val decimal.Decimal, scale int32) ContraTradeQtyField {
	return ContraTradeQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f ContraTradeQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//ContraTradeTimeField is a UTCTIMESTAMP field
type ContraTradeTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.ContraTradeTime (438)
func (f ContraTradeTimeField) Tag() quickfix.Tag { return tag.ContraTradeTime }

//NewContraTradeTime returns a new ContraTradeTimeField initialized with val
func NewContraTradeTime(val time.Time) ContraTradeTimeField {
	return ContraTradeTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewContraTradeTimeNoMillis returns a new ContraTradeTimeField initialized with val without millisecs
func NewContraTradeTimeNoMillis(val time.Time) ContraTradeTimeField {
	return ContraTradeTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f ContraTradeTimeField) Value() time.Time { return f.Time }

//ContraTraderField is a STRING field
type ContraTraderField struct{ quickfix.FIXString }

//Tag returns tag.ContraTrader (337)
func (f ContraTraderField) Tag() quickfix.Tag { return tag.ContraTrader }

//NewContraTrader returns a new ContraTraderField initialized with val
func NewContraTrader(val string) ContraTraderField {
	return ContraTraderField{quickfix.FIXString(val)}
}

func (f ContraTraderField) Value() string { return f.String() }

//ContractMultiplierField is a FLOAT field
type ContractMultiplierField struct{ quickfix.FIXDecimal }

//Tag returns tag.ContractMultiplier (231)
func (f ContractMultiplierField) Tag() quickfix.Tag { return tag.ContractMultiplier }

//NewContractMultiplier returns a new ContractMultiplierField initialized with val and scale
func NewContractMultiplier(val decimal.Decimal, scale int32) ContractMultiplierField {
	return ContractMultiplierField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f ContractMultiplierField) Value() (val decimal.Decimal) { return f.Decimal }

//ContractMultiplierUnitField is a enum.ContractMultiplierUnit field
type ContractMultiplierUnitField struct{ quickfix.FIXString }

//Tag returns tag.ContractMultiplierUnit (1435)
func (f ContractMultiplierUnitField) Tag() quickfix.Tag { return tag.ContractMultiplierUnit }

func NewContractMultiplierUnit(val enum.ContractMultiplierUnit) ContractMultiplierUnitField {
	return ContractMultiplierUnitField{quickfix.FIXString(val)}
}

func (f ContractMultiplierUnitField) Value() enum.ContractMultiplierUnit {
	return enum.ContractMultiplierUnit(f.String())
}

//ContractSettlMonthField is a MONTHYEAR field
type ContractSettlMonthField struct{ quickfix.FIXString }

//Tag returns tag.ContractSettlMonth (667)
func (f ContractSettlMonthField) Tag() quickfix.Tag { return tag.ContractSettlMonth }

//NewContractSettlMonth returns a new ContractSettlMonthField initialized with val
func NewContractSettlMonth(val string) ContractSettlMonthField {
	return ContractSettlMonthField{quickfix.FIXString(val)}
}

func (f ContractSettlMonthField) Value() string { return f.String() }

//ContraryInstructionIndicatorField is a BOOLEAN field
type ContraryInstructionIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.ContraryInstructionIndicator (719)
func (f ContraryInstructionIndicatorField) Tag() quickfix.Tag { return tag.ContraryInstructionIndicator }

//NewContraryInstructionIndicator returns a new ContraryInstructionIndicatorField initialized with val
func NewContraryInstructionIndicator(val bool) ContraryInstructionIndicatorField {
	return ContraryInstructionIndicatorField{quickfix.FIXBoolean(val)}
}

func (f ContraryInstructionIndicatorField) Value() bool { return f.Bool() }

//CopyMsgIndicatorField is a BOOLEAN field
type CopyMsgIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.CopyMsgIndicator (797)
func (f CopyMsgIndicatorField) Tag() quickfix.Tag { return tag.CopyMsgIndicator }

//NewCopyMsgIndicator returns a new CopyMsgIndicatorField initialized with val
func NewCopyMsgIndicator(val bool) CopyMsgIndicatorField {
	return CopyMsgIndicatorField{quickfix.FIXBoolean(val)}
}

func (f CopyMsgIndicatorField) Value() bool { return f.Bool() }

//CorporateActionField is a enum.CorporateAction field
type CorporateActionField struct{ quickfix.FIXString }

//Tag returns tag.CorporateAction (292)
func (f CorporateActionField) Tag() quickfix.Tag { return tag.CorporateAction }

func NewCorporateAction(val enum.CorporateAction) CorporateActionField {
	return CorporateActionField{quickfix.FIXString(val)}
}

func (f CorporateActionField) Value() enum.CorporateAction { return enum.CorporateAction(f.String()) }

//CountryField is a COUNTRY field
type CountryField struct{ quickfix.FIXString }

//Tag returns tag.Country (421)
func (f CountryField) Tag() quickfix.Tag { return tag.Country }

//NewCountry returns a new CountryField initialized with val
func NewCountry(val string) CountryField {
	return CountryField{quickfix.FIXString(val)}
}

func (f CountryField) Value() string { return f.String() }

//CountryOfIssueField is a COUNTRY field
type CountryOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.CountryOfIssue (470)
func (f CountryOfIssueField) Tag() quickfix.Tag { return tag.CountryOfIssue }

//NewCountryOfIssue returns a new CountryOfIssueField initialized with val
func NewCountryOfIssue(val string) CountryOfIssueField {
	return CountryOfIssueField{quickfix.FIXString(val)}
}

func (f CountryOfIssueField) Value() string { return f.String() }

//CouponPaymentDateField is a LOCALMKTDATE field
type CouponPaymentDateField struct{ quickfix.FIXString }

//Tag returns tag.CouponPaymentDate (224)
func (f CouponPaymentDateField) Tag() quickfix.Tag { return tag.CouponPaymentDate }

//NewCouponPaymentDate returns a new CouponPaymentDateField initialized with val
func NewCouponPaymentDate(val string) CouponPaymentDateField {
	return CouponPaymentDateField{quickfix.FIXString(val)}
}

func (f CouponPaymentDateField) Value() string { return f.String() }

//CouponRateField is a PERCENTAGE field
type CouponRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.CouponRate (223)
func (f CouponRateField) Tag() quickfix.Tag { return tag.CouponRate }

//NewCouponRate returns a new CouponRateField initialized with val and scale
func NewCouponRate(val decimal.Decimal, scale int32) CouponRateField {
	return CouponRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CouponRateField) Value() (val decimal.Decimal) { return f.Decimal }

//CoveredOrUncoveredField is a enum.CoveredOrUncovered field
type CoveredOrUncoveredField struct{ quickfix.FIXString }

//Tag returns tag.CoveredOrUncovered (203)
func (f CoveredOrUncoveredField) Tag() quickfix.Tag { return tag.CoveredOrUncovered }

func NewCoveredOrUncovered(val enum.CoveredOrUncovered) CoveredOrUncoveredField {
	return CoveredOrUncoveredField{quickfix.FIXString(val)}
}

func (f CoveredOrUncoveredField) Value() enum.CoveredOrUncovered {
	return enum.CoveredOrUncovered(f.String())
}

//CreditRatingField is a STRING field
type CreditRatingField struct{ quickfix.FIXString }

//Tag returns tag.CreditRating (255)
func (f CreditRatingField) Tag() quickfix.Tag { return tag.CreditRating }

//NewCreditRating returns a new CreditRatingField initialized with val
func NewCreditRating(val string) CreditRatingField {
	return CreditRatingField{quickfix.FIXString(val)}
}

func (f CreditRatingField) Value() string { return f.String() }

//CrossIDField is a STRING field
type CrossIDField struct{ quickfix.FIXString }

//Tag returns tag.CrossID (548)
func (f CrossIDField) Tag() quickfix.Tag { return tag.CrossID }

//NewCrossID returns a new CrossIDField initialized with val
func NewCrossID(val string) CrossIDField {
	return CrossIDField{quickfix.FIXString(val)}
}

func (f CrossIDField) Value() string { return f.String() }

//CrossPercentField is a PERCENTAGE field
type CrossPercentField struct{ quickfix.FIXDecimal }

//Tag returns tag.CrossPercent (413)
func (f CrossPercentField) Tag() quickfix.Tag { return tag.CrossPercent }

//NewCrossPercent returns a new CrossPercentField initialized with val and scale
func NewCrossPercent(val decimal.Decimal, scale int32) CrossPercentField {
	return CrossPercentField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CrossPercentField) Value() (val decimal.Decimal) { return f.Decimal }

//CrossPrioritizationField is a enum.CrossPrioritization field
type CrossPrioritizationField struct{ quickfix.FIXString }

//Tag returns tag.CrossPrioritization (550)
func (f CrossPrioritizationField) Tag() quickfix.Tag { return tag.CrossPrioritization }

func NewCrossPrioritization(val enum.CrossPrioritization) CrossPrioritizationField {
	return CrossPrioritizationField{quickfix.FIXString(val)}
}

func (f CrossPrioritizationField) Value() enum.CrossPrioritization {
	return enum.CrossPrioritization(f.String())
}

//CrossTypeField is a enum.CrossType field
type CrossTypeField struct{ quickfix.FIXString }

//Tag returns tag.CrossType (549)
func (f CrossTypeField) Tag() quickfix.Tag { return tag.CrossType }

func NewCrossType(val enum.CrossType) CrossTypeField {
	return CrossTypeField{quickfix.FIXString(val)}
}

func (f CrossTypeField) Value() enum.CrossType { return enum.CrossType(f.String()) }

//CstmApplVerIDField is a STRING field
type CstmApplVerIDField struct{ quickfix.FIXString }

//Tag returns tag.CstmApplVerID (1129)
func (f CstmApplVerIDField) Tag() quickfix.Tag { return tag.CstmApplVerID }

//NewCstmApplVerID returns a new CstmApplVerIDField initialized with val
func NewCstmApplVerID(val string) CstmApplVerIDField {
	return CstmApplVerIDField{quickfix.FIXString(val)}
}

func (f CstmApplVerIDField) Value() string { return f.String() }

//CumQtyField is a QTY field
type CumQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.CumQty (14)
func (f CumQtyField) Tag() quickfix.Tag { return tag.CumQty }

//NewCumQty returns a new CumQtyField initialized with val and scale
func NewCumQty(val decimal.Decimal, scale int32) CumQtyField {
	return CumQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CumQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//CurrencyField is a CURRENCY field
type CurrencyField struct{ quickfix.FIXString }

//Tag returns tag.Currency (15)
func (f CurrencyField) Tag() quickfix.Tag { return tag.Currency }

//NewCurrency returns a new CurrencyField initialized with val
func NewCurrency(val string) CurrencyField {
	return CurrencyField{quickfix.FIXString(val)}
}

func (f CurrencyField) Value() string { return f.String() }

//CurrencyRatioField is a FLOAT field
type CurrencyRatioField struct{ quickfix.FIXDecimal }

//Tag returns tag.CurrencyRatio (1382)
func (f CurrencyRatioField) Tag() quickfix.Tag { return tag.CurrencyRatio }

//NewCurrencyRatio returns a new CurrencyRatioField initialized with val and scale
func NewCurrencyRatio(val decimal.Decimal, scale int32) CurrencyRatioField {
	return CurrencyRatioField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CurrencyRatioField) Value() (val decimal.Decimal) { return f.Decimal }

//CustDirectedOrderField is a BOOLEAN field
type CustDirectedOrderField struct{ quickfix.FIXBoolean }

//Tag returns tag.CustDirectedOrder (1029)
func (f CustDirectedOrderField) Tag() quickfix.Tag { return tag.CustDirectedOrder }

//NewCustDirectedOrder returns a new CustDirectedOrderField initialized with val
func NewCustDirectedOrder(val bool) CustDirectedOrderField {
	return CustDirectedOrderField{quickfix.FIXBoolean(val)}
}

func (f CustDirectedOrderField) Value() bool { return f.Bool() }

//CustOrderCapacityField is a enum.CustOrderCapacity field
type CustOrderCapacityField struct{ quickfix.FIXString }

//Tag returns tag.CustOrderCapacity (582)
func (f CustOrderCapacityField) Tag() quickfix.Tag { return tag.CustOrderCapacity }

func NewCustOrderCapacity(val enum.CustOrderCapacity) CustOrderCapacityField {
	return CustOrderCapacityField{quickfix.FIXString(val)}
}

func (f CustOrderCapacityField) Value() enum.CustOrderCapacity {
	return enum.CustOrderCapacity(f.String())
}

//CustOrderHandlingInstField is a enum.CustOrderHandlingInst field
type CustOrderHandlingInstField struct{ quickfix.FIXString }

//Tag returns tag.CustOrderHandlingInst (1031)
func (f CustOrderHandlingInstField) Tag() quickfix.Tag { return tag.CustOrderHandlingInst }

func NewCustOrderHandlingInst(val enum.CustOrderHandlingInst) CustOrderHandlingInstField {
	return CustOrderHandlingInstField{quickfix.FIXString(val)}
}

func (f CustOrderHandlingInstField) Value() enum.CustOrderHandlingInst {
	return enum.CustOrderHandlingInst(f.String())
}

//CustomerOrFirmField is a enum.CustomerOrFirm field
type CustomerOrFirmField struct{ quickfix.FIXString }

//Tag returns tag.CustomerOrFirm (204)
func (f CustomerOrFirmField) Tag() quickfix.Tag { return tag.CustomerOrFirm }

func NewCustomerOrFirm(val enum.CustomerOrFirm) CustomerOrFirmField {
	return CustomerOrFirmField{quickfix.FIXString(val)}
}

func (f CustomerOrFirmField) Value() enum.CustomerOrFirm { return enum.CustomerOrFirm(f.String()) }

//CxlQtyField is a QTY field
type CxlQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.CxlQty (84)
func (f CxlQtyField) Tag() quickfix.Tag { return tag.CxlQty }

//NewCxlQty returns a new CxlQtyField initialized with val and scale
func NewCxlQty(val decimal.Decimal, scale int32) CxlQtyField {
	return CxlQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CxlQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//CxlRejReasonField is a enum.CxlRejReason field
type CxlRejReasonField struct{ quickfix.FIXString }

//Tag returns tag.CxlRejReason (102)
func (f CxlRejReasonField) Tag() quickfix.Tag { return tag.CxlRejReason }

func NewCxlRejReason(val enum.CxlRejReason) CxlRejReasonField {
	return CxlRejReasonField{quickfix.FIXString(val)}
}

func (f CxlRejReasonField) Value() enum.CxlRejReason { return enum.CxlRejReason(f.String()) }

//CxlRejResponseToField is a enum.CxlRejResponseTo field
type CxlRejResponseToField struct{ quickfix.FIXString }

//Tag returns tag.CxlRejResponseTo (434)
func (f CxlRejResponseToField) Tag() quickfix.Tag { return tag.CxlRejResponseTo }

func NewCxlRejResponseTo(val enum.CxlRejResponseTo) CxlRejResponseToField {
	return CxlRejResponseToField{quickfix.FIXString(val)}
}

func (f CxlRejResponseToField) Value() enum.CxlRejResponseTo { return enum.CxlRejResponseTo(f.String()) }

//CxlTypeField is a enum.CxlType field
type CxlTypeField struct{ quickfix.FIXString }

//Tag returns tag.CxlType (125)
func (f CxlTypeField) Tag() quickfix.Tag { return tag.CxlType }

func NewCxlType(val enum.CxlType) CxlTypeField {
	return CxlTypeField{quickfix.FIXString(val)}
}

func (f CxlTypeField) Value() enum.CxlType { return enum.CxlType(f.String()) }

//DKReasonField is a enum.DKReason field
type DKReasonField struct{ quickfix.FIXString }

//Tag returns tag.DKReason (127)
func (f DKReasonField) Tag() quickfix.Tag { return tag.DKReason }

func NewDKReason(val enum.DKReason) DKReasonField {
	return DKReasonField{quickfix.FIXString(val)}
}

func (f DKReasonField) Value() enum.DKReason { return enum.DKReason(f.String()) }

//DateOfBirthField is a LOCALMKTDATE field
type DateOfBirthField struct{ quickfix.FIXString }

//Tag returns tag.DateOfBirth (486)
func (f DateOfBirthField) Tag() quickfix.Tag { return tag.DateOfBirth }

//NewDateOfBirth returns a new DateOfBirthField initialized with val
func NewDateOfBirth(val string) DateOfBirthField {
	return DateOfBirthField{quickfix.FIXString(val)}
}

func (f DateOfBirthField) Value() string { return f.String() }

//DatedDateField is a LOCALMKTDATE field
type DatedDateField struct{ quickfix.FIXString }

//Tag returns tag.DatedDate (873)
func (f DatedDateField) Tag() quickfix.Tag { return tag.DatedDate }

//NewDatedDate returns a new DatedDateField initialized with val
func NewDatedDate(val string) DatedDateField {
	return DatedDateField{quickfix.FIXString(val)}
}

func (f DatedDateField) Value() string { return f.String() }

//DayAvgPxField is a PRICE field
type DayAvgPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.DayAvgPx (426)
func (f DayAvgPxField) Tag() quickfix.Tag { return tag.DayAvgPx }

//NewDayAvgPx returns a new DayAvgPxField initialized with val and scale
func NewDayAvgPx(val decimal.Decimal, scale int32) DayAvgPxField {
	return DayAvgPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DayAvgPxField) Value() (val decimal.Decimal) { return f.Decimal }

//DayBookingInstField is a enum.DayBookingInst field
type DayBookingInstField struct{ quickfix.FIXString }

//Tag returns tag.DayBookingInst (589)
func (f DayBookingInstField) Tag() quickfix.Tag { return tag.DayBookingInst }

func NewDayBookingInst(val enum.DayBookingInst) DayBookingInstField {
	return DayBookingInstField{quickfix.FIXString(val)}
}

func (f DayBookingInstField) Value() enum.DayBookingInst { return enum.DayBookingInst(f.String()) }

//DayCumQtyField is a QTY field
type DayCumQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.DayCumQty (425)
func (f DayCumQtyField) Tag() quickfix.Tag { return tag.DayCumQty }

//NewDayCumQty returns a new DayCumQtyField initialized with val and scale
func NewDayCumQty(val decimal.Decimal, scale int32) DayCumQtyField {
	return DayCumQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DayCumQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//DayOrderQtyField is a QTY field
type DayOrderQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.DayOrderQty (424)
func (f DayOrderQtyField) Tag() quickfix.Tag { return tag.DayOrderQty }

//NewDayOrderQty returns a new DayOrderQtyField initialized with val and scale
func NewDayOrderQty(val decimal.Decimal, scale int32) DayOrderQtyField {
	return DayOrderQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DayOrderQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//DealingCapacityField is a enum.DealingCapacity field
type DealingCapacityField struct{ quickfix.FIXString }

//Tag returns tag.DealingCapacity (1048)
func (f DealingCapacityField) Tag() quickfix.Tag { return tag.DealingCapacity }

func NewDealingCapacity(val enum.DealingCapacity) DealingCapacityField {
	return DealingCapacityField{quickfix.FIXString(val)}
}

func (f DealingCapacityField) Value() enum.DealingCapacity { return enum.DealingCapacity(f.String()) }

//DefBidSizeField is a QTY field
type DefBidSizeField struct{ quickfix.FIXDecimal }

//Tag returns tag.DefBidSize (293)
func (f DefBidSizeField) Tag() quickfix.Tag { return tag.DefBidSize }

//NewDefBidSize returns a new DefBidSizeField initialized with val and scale
func NewDefBidSize(val decimal.Decimal, scale int32) DefBidSizeField {
	return DefBidSizeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DefBidSizeField) Value() (val decimal.Decimal) { return f.Decimal }

//DefOfferSizeField is a QTY field
type DefOfferSizeField struct{ quickfix.FIXDecimal }

//Tag returns tag.DefOfferSize (294)
func (f DefOfferSizeField) Tag() quickfix.Tag { return tag.DefOfferSize }

//NewDefOfferSize returns a new DefOfferSizeField initialized with val and scale
func NewDefOfferSize(val decimal.Decimal, scale int32) DefOfferSizeField {
	return DefOfferSizeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DefOfferSizeField) Value() (val decimal.Decimal) { return f.Decimal }

//DefaultApplExtIDField is a INT field
type DefaultApplExtIDField struct{ quickfix.FIXInt }

//Tag returns tag.DefaultApplExtID (1407)
func (f DefaultApplExtIDField) Tag() quickfix.Tag { return tag.DefaultApplExtID }

//NewDefaultApplExtID returns a new DefaultApplExtIDField initialized with val
func NewDefaultApplExtID(val int) DefaultApplExtIDField {
	return DefaultApplExtIDField{quickfix.FIXInt(val)}
}

func (f DefaultApplExtIDField) Value() int { return f.Int() }

//DefaultApplVerIDField is a STRING field
type DefaultApplVerIDField struct{ quickfix.FIXString }

//Tag returns tag.DefaultApplVerID (1137)
func (f DefaultApplVerIDField) Tag() quickfix.Tag { return tag.DefaultApplVerID }

//NewDefaultApplVerID returns a new DefaultApplVerIDField initialized with val
func NewDefaultApplVerID(val string) DefaultApplVerIDField {
	return DefaultApplVerIDField{quickfix.FIXString(val)}
}

func (f DefaultApplVerIDField) Value() string { return f.String() }

//DefaultCstmApplVerIDField is a STRING field
type DefaultCstmApplVerIDField struct{ quickfix.FIXString }

//Tag returns tag.DefaultCstmApplVerID (1408)
func (f DefaultCstmApplVerIDField) Tag() quickfix.Tag { return tag.DefaultCstmApplVerID }

//NewDefaultCstmApplVerID returns a new DefaultCstmApplVerIDField initialized with val
func NewDefaultCstmApplVerID(val string) DefaultCstmApplVerIDField {
	return DefaultCstmApplVerIDField{quickfix.FIXString(val)}
}

func (f DefaultCstmApplVerIDField) Value() string { return f.String() }

//DefaultVerIndicatorField is a BOOLEAN field
type DefaultVerIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.DefaultVerIndicator (1410)
func (f DefaultVerIndicatorField) Tag() quickfix.Tag { return tag.DefaultVerIndicator }

//NewDefaultVerIndicator returns a new DefaultVerIndicatorField initialized with val
func NewDefaultVerIndicator(val bool) DefaultVerIndicatorField {
	return DefaultVerIndicatorField{quickfix.FIXBoolean(val)}
}

func (f DefaultVerIndicatorField) Value() bool { return f.Bool() }

//DeleteReasonField is a enum.DeleteReason field
type DeleteReasonField struct{ quickfix.FIXString }

//Tag returns tag.DeleteReason (285)
func (f DeleteReasonField) Tag() quickfix.Tag { return tag.DeleteReason }

func NewDeleteReason(val enum.DeleteReason) DeleteReasonField {
	return DeleteReasonField{quickfix.FIXString(val)}
}

func (f DeleteReasonField) Value() enum.DeleteReason { return enum.DeleteReason(f.String()) }

//DeliverToCompIDField is a STRING field
type DeliverToCompIDField struct{ quickfix.FIXString }

//Tag returns tag.DeliverToCompID (128)
func (f DeliverToCompIDField) Tag() quickfix.Tag { return tag.DeliverToCompID }

//NewDeliverToCompID returns a new DeliverToCompIDField initialized with val
func NewDeliverToCompID(val string) DeliverToCompIDField {
	return DeliverToCompIDField{quickfix.FIXString(val)}
}

func (f DeliverToCompIDField) Value() string { return f.String() }

//DeliverToLocationIDField is a STRING field
type DeliverToLocationIDField struct{ quickfix.FIXString }

//Tag returns tag.DeliverToLocationID (145)
func (f DeliverToLocationIDField) Tag() quickfix.Tag { return tag.DeliverToLocationID }

//NewDeliverToLocationID returns a new DeliverToLocationIDField initialized with val
func NewDeliverToLocationID(val string) DeliverToLocationIDField {
	return DeliverToLocationIDField{quickfix.FIXString(val)}
}

func (f DeliverToLocationIDField) Value() string { return f.String() }

//DeliverToSubIDField is a STRING field
type DeliverToSubIDField struct{ quickfix.FIXString }

//Tag returns tag.DeliverToSubID (129)
func (f DeliverToSubIDField) Tag() quickfix.Tag { return tag.DeliverToSubID }

//NewDeliverToSubID returns a new DeliverToSubIDField initialized with val
func NewDeliverToSubID(val string) DeliverToSubIDField {
	return DeliverToSubIDField{quickfix.FIXString(val)}
}

func (f DeliverToSubIDField) Value() string { return f.String() }

//DeliveryDateField is a LOCALMKTDATE field
type DeliveryDateField struct{ quickfix.FIXString }

//Tag returns tag.DeliveryDate (743)
func (f DeliveryDateField) Tag() quickfix.Tag { return tag.DeliveryDate }

//NewDeliveryDate returns a new DeliveryDateField initialized with val
func NewDeliveryDate(val string) DeliveryDateField {
	return DeliveryDateField{quickfix.FIXString(val)}
}

func (f DeliveryDateField) Value() string { return f.String() }

//DeliveryFormField is a enum.DeliveryForm field
type DeliveryFormField struct{ quickfix.FIXString }

//Tag returns tag.DeliveryForm (668)
func (f DeliveryFormField) Tag() quickfix.Tag { return tag.DeliveryForm }

func NewDeliveryForm(val enum.DeliveryForm) DeliveryFormField {
	return DeliveryFormField{quickfix.FIXString(val)}
}

func (f DeliveryFormField) Value() enum.DeliveryForm { return enum.DeliveryForm(f.String()) }

//DeliveryTypeField is a enum.DeliveryType field
type DeliveryTypeField struct{ quickfix.FIXString }

//Tag returns tag.DeliveryType (919)
func (f DeliveryTypeField) Tag() quickfix.Tag { return tag.DeliveryType }

func NewDeliveryType(val enum.DeliveryType) DeliveryTypeField {
	return DeliveryTypeField{quickfix.FIXString(val)}
}

func (f DeliveryTypeField) Value() enum.DeliveryType { return enum.DeliveryType(f.String()) }

//DerivFlexProductEligibilityIndicatorField is a BOOLEAN field
type DerivFlexProductEligibilityIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.DerivFlexProductEligibilityIndicator (1243)
func (f DerivFlexProductEligibilityIndicatorField) Tag() quickfix.Tag {
	return tag.DerivFlexProductEligibilityIndicator
}

//NewDerivFlexProductEligibilityIndicator returns a new DerivFlexProductEligibilityIndicatorField initialized with val
func NewDerivFlexProductEligibilityIndicator(val bool) DerivFlexProductEligibilityIndicatorField {
	return DerivFlexProductEligibilityIndicatorField{quickfix.FIXBoolean(val)}
}

func (f DerivFlexProductEligibilityIndicatorField) Value() bool { return f.Bool() }

//DerivativeCFICodeField is a STRING field
type DerivativeCFICodeField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeCFICode (1248)
func (f DerivativeCFICodeField) Tag() quickfix.Tag { return tag.DerivativeCFICode }

//NewDerivativeCFICode returns a new DerivativeCFICodeField initialized with val
func NewDerivativeCFICode(val string) DerivativeCFICodeField {
	return DerivativeCFICodeField{quickfix.FIXString(val)}
}

func (f DerivativeCFICodeField) Value() string { return f.String() }

//DerivativeCapPriceField is a PRICE field
type DerivativeCapPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.DerivativeCapPrice (1321)
func (f DerivativeCapPriceField) Tag() quickfix.Tag { return tag.DerivativeCapPrice }

//NewDerivativeCapPrice returns a new DerivativeCapPriceField initialized with val and scale
func NewDerivativeCapPrice(val decimal.Decimal, scale int32) DerivativeCapPriceField {
	return DerivativeCapPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DerivativeCapPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//DerivativeContractMultiplierField is a FLOAT field
type DerivativeContractMultiplierField struct{ quickfix.FIXDecimal }

//Tag returns tag.DerivativeContractMultiplier (1266)
func (f DerivativeContractMultiplierField) Tag() quickfix.Tag { return tag.DerivativeContractMultiplier }

//NewDerivativeContractMultiplier returns a new DerivativeContractMultiplierField initialized with val and scale
func NewDerivativeContractMultiplier(val decimal.Decimal, scale int32) DerivativeContractMultiplierField {
	return DerivativeContractMultiplierField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DerivativeContractMultiplierField) Value() (val decimal.Decimal) { return f.Decimal }

//DerivativeContractMultiplierUnitField is a INT field
type DerivativeContractMultiplierUnitField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeContractMultiplierUnit (1438)
func (f DerivativeContractMultiplierUnitField) Tag() quickfix.Tag {
	return tag.DerivativeContractMultiplierUnit
}

//NewDerivativeContractMultiplierUnit returns a new DerivativeContractMultiplierUnitField initialized with val
func NewDerivativeContractMultiplierUnit(val int) DerivativeContractMultiplierUnitField {
	return DerivativeContractMultiplierUnitField{quickfix.FIXInt(val)}
}

func (f DerivativeContractMultiplierUnitField) Value() int { return f.Int() }

//DerivativeContractSettlMonthField is a MONTHYEAR field
type DerivativeContractSettlMonthField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeContractSettlMonth (1285)
func (f DerivativeContractSettlMonthField) Tag() quickfix.Tag { return tag.DerivativeContractSettlMonth }

//NewDerivativeContractSettlMonth returns a new DerivativeContractSettlMonthField initialized with val
func NewDerivativeContractSettlMonth(val string) DerivativeContractSettlMonthField {
	return DerivativeContractSettlMonthField{quickfix.FIXString(val)}
}

func (f DerivativeContractSettlMonthField) Value() string { return f.String() }

//DerivativeCountryOfIssueField is a COUNTRY field
type DerivativeCountryOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeCountryOfIssue (1258)
func (f DerivativeCountryOfIssueField) Tag() quickfix.Tag { return tag.DerivativeCountryOfIssue }

//NewDerivativeCountryOfIssue returns a new DerivativeCountryOfIssueField initialized with val
func NewDerivativeCountryOfIssue(val string) DerivativeCountryOfIssueField {
	return DerivativeCountryOfIssueField{quickfix.FIXString(val)}
}

func (f DerivativeCountryOfIssueField) Value() string { return f.String() }

//DerivativeEncodedIssuerField is a DATA field
type DerivativeEncodedIssuerField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeEncodedIssuer (1278)
func (f DerivativeEncodedIssuerField) Tag() quickfix.Tag { return tag.DerivativeEncodedIssuer }

//NewDerivativeEncodedIssuer returns a new DerivativeEncodedIssuerField initialized with val
func NewDerivativeEncodedIssuer(val string) DerivativeEncodedIssuerField {
	return DerivativeEncodedIssuerField{quickfix.FIXString(val)}
}

func (f DerivativeEncodedIssuerField) Value() string { return f.String() }

//DerivativeEncodedIssuerLenField is a LENGTH field
type DerivativeEncodedIssuerLenField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeEncodedIssuerLen (1277)
func (f DerivativeEncodedIssuerLenField) Tag() quickfix.Tag { return tag.DerivativeEncodedIssuerLen }

//NewDerivativeEncodedIssuerLen returns a new DerivativeEncodedIssuerLenField initialized with val
func NewDerivativeEncodedIssuerLen(val int) DerivativeEncodedIssuerLenField {
	return DerivativeEncodedIssuerLenField{quickfix.FIXInt(val)}
}

func (f DerivativeEncodedIssuerLenField) Value() int { return f.Int() }

//DerivativeEncodedSecurityDescField is a DATA field
type DerivativeEncodedSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeEncodedSecurityDesc (1281)
func (f DerivativeEncodedSecurityDescField) Tag() quickfix.Tag {
	return tag.DerivativeEncodedSecurityDesc
}

//NewDerivativeEncodedSecurityDesc returns a new DerivativeEncodedSecurityDescField initialized with val
func NewDerivativeEncodedSecurityDesc(val string) DerivativeEncodedSecurityDescField {
	return DerivativeEncodedSecurityDescField{quickfix.FIXString(val)}
}

func (f DerivativeEncodedSecurityDescField) Value() string { return f.String() }

//DerivativeEncodedSecurityDescLenField is a LENGTH field
type DerivativeEncodedSecurityDescLenField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeEncodedSecurityDescLen (1280)
func (f DerivativeEncodedSecurityDescLenField) Tag() quickfix.Tag {
	return tag.DerivativeEncodedSecurityDescLen
}

//NewDerivativeEncodedSecurityDescLen returns a new DerivativeEncodedSecurityDescLenField initialized with val
func NewDerivativeEncodedSecurityDescLen(val int) DerivativeEncodedSecurityDescLenField {
	return DerivativeEncodedSecurityDescLenField{quickfix.FIXInt(val)}
}

func (f DerivativeEncodedSecurityDescLenField) Value() int { return f.Int() }

//DerivativeEventDateField is a LOCALMKTDATE field
type DerivativeEventDateField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeEventDate (1288)
func (f DerivativeEventDateField) Tag() quickfix.Tag { return tag.DerivativeEventDate }

//NewDerivativeEventDate returns a new DerivativeEventDateField initialized with val
func NewDerivativeEventDate(val string) DerivativeEventDateField {
	return DerivativeEventDateField{quickfix.FIXString(val)}
}

func (f DerivativeEventDateField) Value() string { return f.String() }

//DerivativeEventPxField is a PRICE field
type DerivativeEventPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.DerivativeEventPx (1290)
func (f DerivativeEventPxField) Tag() quickfix.Tag { return tag.DerivativeEventPx }

//NewDerivativeEventPx returns a new DerivativeEventPxField initialized with val and scale
func NewDerivativeEventPx(val decimal.Decimal, scale int32) DerivativeEventPxField {
	return DerivativeEventPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DerivativeEventPxField) Value() (val decimal.Decimal) { return f.Decimal }

//DerivativeEventTextField is a STRING field
type DerivativeEventTextField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeEventText (1291)
func (f DerivativeEventTextField) Tag() quickfix.Tag { return tag.DerivativeEventText }

//NewDerivativeEventText returns a new DerivativeEventTextField initialized with val
func NewDerivativeEventText(val string) DerivativeEventTextField {
	return DerivativeEventTextField{quickfix.FIXString(val)}
}

func (f DerivativeEventTextField) Value() string { return f.String() }

//DerivativeEventTimeField is a UTCTIMESTAMP field
type DerivativeEventTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.DerivativeEventTime (1289)
func (f DerivativeEventTimeField) Tag() quickfix.Tag { return tag.DerivativeEventTime }

//NewDerivativeEventTime returns a new DerivativeEventTimeField initialized with val
func NewDerivativeEventTime(val time.Time) DerivativeEventTimeField {
	return DerivativeEventTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewDerivativeEventTimeNoMillis returns a new DerivativeEventTimeField initialized with val without millisecs
func NewDerivativeEventTimeNoMillis(val time.Time) DerivativeEventTimeField {
	return DerivativeEventTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f DerivativeEventTimeField) Value() time.Time { return f.Time }

//DerivativeEventTypeField is a INT field
type DerivativeEventTypeField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeEventType (1287)
func (f DerivativeEventTypeField) Tag() quickfix.Tag { return tag.DerivativeEventType }

//NewDerivativeEventType returns a new DerivativeEventTypeField initialized with val
func NewDerivativeEventType(val int) DerivativeEventTypeField {
	return DerivativeEventTypeField{quickfix.FIXInt(val)}
}

func (f DerivativeEventTypeField) Value() int { return f.Int() }

//DerivativeExerciseStyleField is a CHAR field
type DerivativeExerciseStyleField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeExerciseStyle (1299)
func (f DerivativeExerciseStyleField) Tag() quickfix.Tag { return tag.DerivativeExerciseStyle }

//NewDerivativeExerciseStyle returns a new DerivativeExerciseStyleField initialized with val
func NewDerivativeExerciseStyle(val string) DerivativeExerciseStyleField {
	return DerivativeExerciseStyleField{quickfix.FIXString(val)}
}

func (f DerivativeExerciseStyleField) Value() string { return f.String() }

//DerivativeFloorPriceField is a PRICE field
type DerivativeFloorPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.DerivativeFloorPrice (1322)
func (f DerivativeFloorPriceField) Tag() quickfix.Tag { return tag.DerivativeFloorPrice }

//NewDerivativeFloorPrice returns a new DerivativeFloorPriceField initialized with val and scale
func NewDerivativeFloorPrice(val decimal.Decimal, scale int32) DerivativeFloorPriceField {
	return DerivativeFloorPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DerivativeFloorPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//DerivativeFlowScheduleTypeField is a INT field
type DerivativeFlowScheduleTypeField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeFlowScheduleType (1442)
func (f DerivativeFlowScheduleTypeField) Tag() quickfix.Tag { return tag.DerivativeFlowScheduleType }

//NewDerivativeFlowScheduleType returns a new DerivativeFlowScheduleTypeField initialized with val
func NewDerivativeFlowScheduleType(val int) DerivativeFlowScheduleTypeField {
	return DerivativeFlowScheduleTypeField{quickfix.FIXInt(val)}
}

func (f DerivativeFlowScheduleTypeField) Value() int { return f.Int() }

//DerivativeFuturesValuationMethodField is a STRING field
type DerivativeFuturesValuationMethodField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeFuturesValuationMethod (1319)
func (f DerivativeFuturesValuationMethodField) Tag() quickfix.Tag {
	return tag.DerivativeFuturesValuationMethod
}

//NewDerivativeFuturesValuationMethod returns a new DerivativeFuturesValuationMethodField initialized with val
func NewDerivativeFuturesValuationMethod(val string) DerivativeFuturesValuationMethodField {
	return DerivativeFuturesValuationMethodField{quickfix.FIXString(val)}
}

func (f DerivativeFuturesValuationMethodField) Value() string { return f.String() }

//DerivativeInstrAttribTypeField is a INT field
type DerivativeInstrAttribTypeField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeInstrAttribType (1313)
func (f DerivativeInstrAttribTypeField) Tag() quickfix.Tag { return tag.DerivativeInstrAttribType }

//NewDerivativeInstrAttribType returns a new DerivativeInstrAttribTypeField initialized with val
func NewDerivativeInstrAttribType(val int) DerivativeInstrAttribTypeField {
	return DerivativeInstrAttribTypeField{quickfix.FIXInt(val)}
}

func (f DerivativeInstrAttribTypeField) Value() int { return f.Int() }

//DerivativeInstrAttribValueField is a STRING field
type DerivativeInstrAttribValueField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeInstrAttribValue (1314)
func (f DerivativeInstrAttribValueField) Tag() quickfix.Tag { return tag.DerivativeInstrAttribValue }

//NewDerivativeInstrAttribValue returns a new DerivativeInstrAttribValueField initialized with val
func NewDerivativeInstrAttribValue(val string) DerivativeInstrAttribValueField {
	return DerivativeInstrAttribValueField{quickfix.FIXString(val)}
}

func (f DerivativeInstrAttribValueField) Value() string { return f.String() }

//DerivativeInstrRegistryField is a STRING field
type DerivativeInstrRegistryField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeInstrRegistry (1257)
func (f DerivativeInstrRegistryField) Tag() quickfix.Tag { return tag.DerivativeInstrRegistry }

//NewDerivativeInstrRegistry returns a new DerivativeInstrRegistryField initialized with val
func NewDerivativeInstrRegistry(val string) DerivativeInstrRegistryField {
	return DerivativeInstrRegistryField{quickfix.FIXString(val)}
}

func (f DerivativeInstrRegistryField) Value() string { return f.String() }

//DerivativeInstrmtAssignmentMethodField is a CHAR field
type DerivativeInstrmtAssignmentMethodField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeInstrmtAssignmentMethod (1255)
func (f DerivativeInstrmtAssignmentMethodField) Tag() quickfix.Tag {
	return tag.DerivativeInstrmtAssignmentMethod
}

//NewDerivativeInstrmtAssignmentMethod returns a new DerivativeInstrmtAssignmentMethodField initialized with val
func NewDerivativeInstrmtAssignmentMethod(val string) DerivativeInstrmtAssignmentMethodField {
	return DerivativeInstrmtAssignmentMethodField{quickfix.FIXString(val)}
}

func (f DerivativeInstrmtAssignmentMethodField) Value() string { return f.String() }

//DerivativeInstrumentPartyIDField is a STRING field
type DerivativeInstrumentPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeInstrumentPartyID (1293)
func (f DerivativeInstrumentPartyIDField) Tag() quickfix.Tag { return tag.DerivativeInstrumentPartyID }

//NewDerivativeInstrumentPartyID returns a new DerivativeInstrumentPartyIDField initialized with val
func NewDerivativeInstrumentPartyID(val string) DerivativeInstrumentPartyIDField {
	return DerivativeInstrumentPartyIDField{quickfix.FIXString(val)}
}

func (f DerivativeInstrumentPartyIDField) Value() string { return f.String() }

//DerivativeInstrumentPartyIDSourceField is a STRING field
type DerivativeInstrumentPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeInstrumentPartyIDSource (1294)
func (f DerivativeInstrumentPartyIDSourceField) Tag() quickfix.Tag {
	return tag.DerivativeInstrumentPartyIDSource
}

//NewDerivativeInstrumentPartyIDSource returns a new DerivativeInstrumentPartyIDSourceField initialized with val
func NewDerivativeInstrumentPartyIDSource(val string) DerivativeInstrumentPartyIDSourceField {
	return DerivativeInstrumentPartyIDSourceField{quickfix.FIXString(val)}
}

func (f DerivativeInstrumentPartyIDSourceField) Value() string { return f.String() }

//DerivativeInstrumentPartyRoleField is a INT field
type DerivativeInstrumentPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeInstrumentPartyRole (1295)
func (f DerivativeInstrumentPartyRoleField) Tag() quickfix.Tag {
	return tag.DerivativeInstrumentPartyRole
}

//NewDerivativeInstrumentPartyRole returns a new DerivativeInstrumentPartyRoleField initialized with val
func NewDerivativeInstrumentPartyRole(val int) DerivativeInstrumentPartyRoleField {
	return DerivativeInstrumentPartyRoleField{quickfix.FIXInt(val)}
}

func (f DerivativeInstrumentPartyRoleField) Value() int { return f.Int() }

//DerivativeInstrumentPartySubIDField is a STRING field
type DerivativeInstrumentPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeInstrumentPartySubID (1297)
func (f DerivativeInstrumentPartySubIDField) Tag() quickfix.Tag {
	return tag.DerivativeInstrumentPartySubID
}

//NewDerivativeInstrumentPartySubID returns a new DerivativeInstrumentPartySubIDField initialized with val
func NewDerivativeInstrumentPartySubID(val string) DerivativeInstrumentPartySubIDField {
	return DerivativeInstrumentPartySubIDField{quickfix.FIXString(val)}
}

func (f DerivativeInstrumentPartySubIDField) Value() string { return f.String() }

//DerivativeInstrumentPartySubIDTypeField is a INT field
type DerivativeInstrumentPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeInstrumentPartySubIDType (1298)
func (f DerivativeInstrumentPartySubIDTypeField) Tag() quickfix.Tag {
	return tag.DerivativeInstrumentPartySubIDType
}

//NewDerivativeInstrumentPartySubIDType returns a new DerivativeInstrumentPartySubIDTypeField initialized with val
func NewDerivativeInstrumentPartySubIDType(val int) DerivativeInstrumentPartySubIDTypeField {
	return DerivativeInstrumentPartySubIDTypeField{quickfix.FIXInt(val)}
}

func (f DerivativeInstrumentPartySubIDTypeField) Value() int { return f.Int() }

//DerivativeIssueDateField is a LOCALMKTDATE field
type DerivativeIssueDateField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeIssueDate (1276)
func (f DerivativeIssueDateField) Tag() quickfix.Tag { return tag.DerivativeIssueDate }

//NewDerivativeIssueDate returns a new DerivativeIssueDateField initialized with val
func NewDerivativeIssueDate(val string) DerivativeIssueDateField {
	return DerivativeIssueDateField{quickfix.FIXString(val)}
}

func (f DerivativeIssueDateField) Value() string { return f.String() }

//DerivativeIssuerField is a STRING field
type DerivativeIssuerField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeIssuer (1275)
func (f DerivativeIssuerField) Tag() quickfix.Tag { return tag.DerivativeIssuer }

//NewDerivativeIssuer returns a new DerivativeIssuerField initialized with val
func NewDerivativeIssuer(val string) DerivativeIssuerField {
	return DerivativeIssuerField{quickfix.FIXString(val)}
}

func (f DerivativeIssuerField) Value() string { return f.String() }

//DerivativeListMethodField is a INT field
type DerivativeListMethodField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeListMethod (1320)
func (f DerivativeListMethodField) Tag() quickfix.Tag { return tag.DerivativeListMethod }

//NewDerivativeListMethod returns a new DerivativeListMethodField initialized with val
func NewDerivativeListMethod(val int) DerivativeListMethodField {
	return DerivativeListMethodField{quickfix.FIXInt(val)}
}

func (f DerivativeListMethodField) Value() int { return f.Int() }

//DerivativeLocaleOfIssueField is a STRING field
type DerivativeLocaleOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeLocaleOfIssue (1260)
func (f DerivativeLocaleOfIssueField) Tag() quickfix.Tag { return tag.DerivativeLocaleOfIssue }

//NewDerivativeLocaleOfIssue returns a new DerivativeLocaleOfIssueField initialized with val
func NewDerivativeLocaleOfIssue(val string) DerivativeLocaleOfIssueField {
	return DerivativeLocaleOfIssueField{quickfix.FIXString(val)}
}

func (f DerivativeLocaleOfIssueField) Value() string { return f.String() }

//DerivativeMaturityDateField is a LOCALMKTDATE field
type DerivativeMaturityDateField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeMaturityDate (1252)
func (f DerivativeMaturityDateField) Tag() quickfix.Tag { return tag.DerivativeMaturityDate }

//NewDerivativeMaturityDate returns a new DerivativeMaturityDateField initialized with val
func NewDerivativeMaturityDate(val string) DerivativeMaturityDateField {
	return DerivativeMaturityDateField{quickfix.FIXString(val)}
}

func (f DerivativeMaturityDateField) Value() string { return f.String() }

//DerivativeMaturityMonthYearField is a MONTHYEAR field
type DerivativeMaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeMaturityMonthYear (1251)
func (f DerivativeMaturityMonthYearField) Tag() quickfix.Tag { return tag.DerivativeMaturityMonthYear }

//NewDerivativeMaturityMonthYear returns a new DerivativeMaturityMonthYearField initialized with val
func NewDerivativeMaturityMonthYear(val string) DerivativeMaturityMonthYearField {
	return DerivativeMaturityMonthYearField{quickfix.FIXString(val)}
}

func (f DerivativeMaturityMonthYearField) Value() string { return f.String() }

//DerivativeMaturityTimeField is a TZTIMEONLY field
type DerivativeMaturityTimeField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeMaturityTime (1253)
func (f DerivativeMaturityTimeField) Tag() quickfix.Tag { return tag.DerivativeMaturityTime }

//NewDerivativeMaturityTime returns a new DerivativeMaturityTimeField initialized with val
func NewDerivativeMaturityTime(val string) DerivativeMaturityTimeField {
	return DerivativeMaturityTimeField{quickfix.FIXString(val)}
}

func (f DerivativeMaturityTimeField) Value() string { return f.String() }

//DerivativeMinPriceIncrementField is a FLOAT field
type DerivativeMinPriceIncrementField struct{ quickfix.FIXDecimal }

//Tag returns tag.DerivativeMinPriceIncrement (1267)
func (f DerivativeMinPriceIncrementField) Tag() quickfix.Tag { return tag.DerivativeMinPriceIncrement }

//NewDerivativeMinPriceIncrement returns a new DerivativeMinPriceIncrementField initialized with val and scale
func NewDerivativeMinPriceIncrement(val decimal.Decimal, scale int32) DerivativeMinPriceIncrementField {
	return DerivativeMinPriceIncrementField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DerivativeMinPriceIncrementField) Value() (val decimal.Decimal) { return f.Decimal }

//DerivativeMinPriceIncrementAmountField is a AMT field
type DerivativeMinPriceIncrementAmountField struct{ quickfix.FIXDecimal }

//Tag returns tag.DerivativeMinPriceIncrementAmount (1268)
func (f DerivativeMinPriceIncrementAmountField) Tag() quickfix.Tag {
	return tag.DerivativeMinPriceIncrementAmount
}

//NewDerivativeMinPriceIncrementAmount returns a new DerivativeMinPriceIncrementAmountField initialized with val and scale
func NewDerivativeMinPriceIncrementAmount(val decimal.Decimal, scale int32) DerivativeMinPriceIncrementAmountField {
	return DerivativeMinPriceIncrementAmountField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DerivativeMinPriceIncrementAmountField) Value() (val decimal.Decimal) { return f.Decimal }

//DerivativeNTPositionLimitField is a INT field
type DerivativeNTPositionLimitField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeNTPositionLimit (1274)
func (f DerivativeNTPositionLimitField) Tag() quickfix.Tag { return tag.DerivativeNTPositionLimit }

//NewDerivativeNTPositionLimit returns a new DerivativeNTPositionLimitField initialized with val
func NewDerivativeNTPositionLimit(val int) DerivativeNTPositionLimitField {
	return DerivativeNTPositionLimitField{quickfix.FIXInt(val)}
}

func (f DerivativeNTPositionLimitField) Value() int { return f.Int() }

//DerivativeOptAttributeField is a CHAR field
type DerivativeOptAttributeField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeOptAttribute (1265)
func (f DerivativeOptAttributeField) Tag() quickfix.Tag { return tag.DerivativeOptAttribute }

//NewDerivativeOptAttribute returns a new DerivativeOptAttributeField initialized with val
func NewDerivativeOptAttribute(val string) DerivativeOptAttributeField {
	return DerivativeOptAttributeField{quickfix.FIXString(val)}
}

func (f DerivativeOptAttributeField) Value() string { return f.String() }

//DerivativeOptPayAmountField is a AMT field
type DerivativeOptPayAmountField struct{ quickfix.FIXDecimal }

//Tag returns tag.DerivativeOptPayAmount (1225)
func (f DerivativeOptPayAmountField) Tag() quickfix.Tag { return tag.DerivativeOptPayAmount }

//NewDerivativeOptPayAmount returns a new DerivativeOptPayAmountField initialized with val and scale
func NewDerivativeOptPayAmount(val decimal.Decimal, scale int32) DerivativeOptPayAmountField {
	return DerivativeOptPayAmountField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DerivativeOptPayAmountField) Value() (val decimal.Decimal) { return f.Decimal }

//DerivativePositionLimitField is a INT field
type DerivativePositionLimitField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativePositionLimit (1273)
func (f DerivativePositionLimitField) Tag() quickfix.Tag { return tag.DerivativePositionLimit }

//NewDerivativePositionLimit returns a new DerivativePositionLimitField initialized with val
func NewDerivativePositionLimit(val int) DerivativePositionLimitField {
	return DerivativePositionLimitField{quickfix.FIXInt(val)}
}

func (f DerivativePositionLimitField) Value() int { return f.Int() }

//DerivativePriceQuoteMethodField is a STRING field
type DerivativePriceQuoteMethodField struct{ quickfix.FIXString }

//Tag returns tag.DerivativePriceQuoteMethod (1318)
func (f DerivativePriceQuoteMethodField) Tag() quickfix.Tag { return tag.DerivativePriceQuoteMethod }

//NewDerivativePriceQuoteMethod returns a new DerivativePriceQuoteMethodField initialized with val
func NewDerivativePriceQuoteMethod(val string) DerivativePriceQuoteMethodField {
	return DerivativePriceQuoteMethodField{quickfix.FIXString(val)}
}

func (f DerivativePriceQuoteMethodField) Value() string { return f.String() }

//DerivativePriceUnitOfMeasureField is a STRING field
type DerivativePriceUnitOfMeasureField struct{ quickfix.FIXString }

//Tag returns tag.DerivativePriceUnitOfMeasure (1315)
func (f DerivativePriceUnitOfMeasureField) Tag() quickfix.Tag { return tag.DerivativePriceUnitOfMeasure }

//NewDerivativePriceUnitOfMeasure returns a new DerivativePriceUnitOfMeasureField initialized with val
func NewDerivativePriceUnitOfMeasure(val string) DerivativePriceUnitOfMeasureField {
	return DerivativePriceUnitOfMeasureField{quickfix.FIXString(val)}
}

func (f DerivativePriceUnitOfMeasureField) Value() string { return f.String() }

//DerivativePriceUnitOfMeasureQtyField is a QTY field
type DerivativePriceUnitOfMeasureQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.DerivativePriceUnitOfMeasureQty (1316)
func (f DerivativePriceUnitOfMeasureQtyField) Tag() quickfix.Tag {
	return tag.DerivativePriceUnitOfMeasureQty
}

//NewDerivativePriceUnitOfMeasureQty returns a new DerivativePriceUnitOfMeasureQtyField initialized with val and scale
func NewDerivativePriceUnitOfMeasureQty(val decimal.Decimal, scale int32) DerivativePriceUnitOfMeasureQtyField {
	return DerivativePriceUnitOfMeasureQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DerivativePriceUnitOfMeasureQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//DerivativeProductField is a INT field
type DerivativeProductField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeProduct (1246)
func (f DerivativeProductField) Tag() quickfix.Tag { return tag.DerivativeProduct }

//NewDerivativeProduct returns a new DerivativeProductField initialized with val
func NewDerivativeProduct(val int) DerivativeProductField {
	return DerivativeProductField{quickfix.FIXInt(val)}
}

func (f DerivativeProductField) Value() int { return f.Int() }

//DerivativeProductComplexField is a STRING field
type DerivativeProductComplexField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeProductComplex (1228)
func (f DerivativeProductComplexField) Tag() quickfix.Tag { return tag.DerivativeProductComplex }

//NewDerivativeProductComplex returns a new DerivativeProductComplexField initialized with val
func NewDerivativeProductComplex(val string) DerivativeProductComplexField {
	return DerivativeProductComplexField{quickfix.FIXString(val)}
}

func (f DerivativeProductComplexField) Value() string { return f.String() }

//DerivativePutOrCallField is a INT field
type DerivativePutOrCallField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativePutOrCall (1323)
func (f DerivativePutOrCallField) Tag() quickfix.Tag { return tag.DerivativePutOrCall }

//NewDerivativePutOrCall returns a new DerivativePutOrCallField initialized with val
func NewDerivativePutOrCall(val int) DerivativePutOrCallField {
	return DerivativePutOrCallField{quickfix.FIXInt(val)}
}

func (f DerivativePutOrCallField) Value() int { return f.Int() }

//DerivativeSecurityAltIDField is a STRING field
type DerivativeSecurityAltIDField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityAltID (1219)
func (f DerivativeSecurityAltIDField) Tag() quickfix.Tag { return tag.DerivativeSecurityAltID }

//NewDerivativeSecurityAltID returns a new DerivativeSecurityAltIDField initialized with val
func NewDerivativeSecurityAltID(val string) DerivativeSecurityAltIDField {
	return DerivativeSecurityAltIDField{quickfix.FIXString(val)}
}

func (f DerivativeSecurityAltIDField) Value() string { return f.String() }

//DerivativeSecurityAltIDSourceField is a STRING field
type DerivativeSecurityAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityAltIDSource (1220)
func (f DerivativeSecurityAltIDSourceField) Tag() quickfix.Tag {
	return tag.DerivativeSecurityAltIDSource
}

//NewDerivativeSecurityAltIDSource returns a new DerivativeSecurityAltIDSourceField initialized with val
func NewDerivativeSecurityAltIDSource(val string) DerivativeSecurityAltIDSourceField {
	return DerivativeSecurityAltIDSourceField{quickfix.FIXString(val)}
}

func (f DerivativeSecurityAltIDSourceField) Value() string { return f.String() }

//DerivativeSecurityDescField is a STRING field
type DerivativeSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityDesc (1279)
func (f DerivativeSecurityDescField) Tag() quickfix.Tag { return tag.DerivativeSecurityDesc }

//NewDerivativeSecurityDesc returns a new DerivativeSecurityDescField initialized with val
func NewDerivativeSecurityDesc(val string) DerivativeSecurityDescField {
	return DerivativeSecurityDescField{quickfix.FIXString(val)}
}

func (f DerivativeSecurityDescField) Value() string { return f.String() }

//DerivativeSecurityExchangeField is a EXCHANGE field
type DerivativeSecurityExchangeField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityExchange (1272)
func (f DerivativeSecurityExchangeField) Tag() quickfix.Tag { return tag.DerivativeSecurityExchange }

//NewDerivativeSecurityExchange returns a new DerivativeSecurityExchangeField initialized with val
func NewDerivativeSecurityExchange(val string) DerivativeSecurityExchangeField {
	return DerivativeSecurityExchangeField{quickfix.FIXString(val)}
}

func (f DerivativeSecurityExchangeField) Value() string { return f.String() }

//DerivativeSecurityGroupField is a STRING field
type DerivativeSecurityGroupField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityGroup (1247)
func (f DerivativeSecurityGroupField) Tag() quickfix.Tag { return tag.DerivativeSecurityGroup }

//NewDerivativeSecurityGroup returns a new DerivativeSecurityGroupField initialized with val
func NewDerivativeSecurityGroup(val string) DerivativeSecurityGroupField {
	return DerivativeSecurityGroupField{quickfix.FIXString(val)}
}

func (f DerivativeSecurityGroupField) Value() string { return f.String() }

//DerivativeSecurityIDField is a STRING field
type DerivativeSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityID (1216)
func (f DerivativeSecurityIDField) Tag() quickfix.Tag { return tag.DerivativeSecurityID }

//NewDerivativeSecurityID returns a new DerivativeSecurityIDField initialized with val
func NewDerivativeSecurityID(val string) DerivativeSecurityIDField {
	return DerivativeSecurityIDField{quickfix.FIXString(val)}
}

func (f DerivativeSecurityIDField) Value() string { return f.String() }

//DerivativeSecurityIDSourceField is a STRING field
type DerivativeSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityIDSource (1217)
func (f DerivativeSecurityIDSourceField) Tag() quickfix.Tag { return tag.DerivativeSecurityIDSource }

//NewDerivativeSecurityIDSource returns a new DerivativeSecurityIDSourceField initialized with val
func NewDerivativeSecurityIDSource(val string) DerivativeSecurityIDSourceField {
	return DerivativeSecurityIDSourceField{quickfix.FIXString(val)}
}

func (f DerivativeSecurityIDSourceField) Value() string { return f.String() }

//DerivativeSecurityListRequestTypeField is a enum.DerivativeSecurityListRequestType field
type DerivativeSecurityListRequestTypeField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityListRequestType (1307)
func (f DerivativeSecurityListRequestTypeField) Tag() quickfix.Tag {
	return tag.DerivativeSecurityListRequestType
}

func NewDerivativeSecurityListRequestType(val enum.DerivativeSecurityListRequestType) DerivativeSecurityListRequestTypeField {
	return DerivativeSecurityListRequestTypeField{quickfix.FIXString(val)}
}

func (f DerivativeSecurityListRequestTypeField) Value() enum.DerivativeSecurityListRequestType {
	return enum.DerivativeSecurityListRequestType(f.String())
}

//DerivativeSecurityStatusField is a STRING field
type DerivativeSecurityStatusField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityStatus (1256)
func (f DerivativeSecurityStatusField) Tag() quickfix.Tag { return tag.DerivativeSecurityStatus }

//NewDerivativeSecurityStatus returns a new DerivativeSecurityStatusField initialized with val
func NewDerivativeSecurityStatus(val string) DerivativeSecurityStatusField {
	return DerivativeSecurityStatusField{quickfix.FIXString(val)}
}

func (f DerivativeSecurityStatusField) Value() string { return f.String() }

//DerivativeSecuritySubTypeField is a STRING field
type DerivativeSecuritySubTypeField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecuritySubType (1250)
func (f DerivativeSecuritySubTypeField) Tag() quickfix.Tag { return tag.DerivativeSecuritySubType }

//NewDerivativeSecuritySubType returns a new DerivativeSecuritySubTypeField initialized with val
func NewDerivativeSecuritySubType(val string) DerivativeSecuritySubTypeField {
	return DerivativeSecuritySubTypeField{quickfix.FIXString(val)}
}

func (f DerivativeSecuritySubTypeField) Value() string { return f.String() }

//DerivativeSecurityTypeField is a STRING field
type DerivativeSecurityTypeField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityType (1249)
func (f DerivativeSecurityTypeField) Tag() quickfix.Tag { return tag.DerivativeSecurityType }

//NewDerivativeSecurityType returns a new DerivativeSecurityTypeField initialized with val
func NewDerivativeSecurityType(val string) DerivativeSecurityTypeField {
	return DerivativeSecurityTypeField{quickfix.FIXString(val)}
}

func (f DerivativeSecurityTypeField) Value() string { return f.String() }

//DerivativeSecurityXMLField is a DATA field
type DerivativeSecurityXMLField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityXML (1283)
func (f DerivativeSecurityXMLField) Tag() quickfix.Tag { return tag.DerivativeSecurityXML }

//NewDerivativeSecurityXML returns a new DerivativeSecurityXMLField initialized with val
func NewDerivativeSecurityXML(val string) DerivativeSecurityXMLField {
	return DerivativeSecurityXMLField{quickfix.FIXString(val)}
}

func (f DerivativeSecurityXMLField) Value() string { return f.String() }

//DerivativeSecurityXMLLenField is a LENGTH field
type DerivativeSecurityXMLLenField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeSecurityXMLLen (1282)
func (f DerivativeSecurityXMLLenField) Tag() quickfix.Tag { return tag.DerivativeSecurityXMLLen }

//NewDerivativeSecurityXMLLen returns a new DerivativeSecurityXMLLenField initialized with val
func NewDerivativeSecurityXMLLen(val int) DerivativeSecurityXMLLenField {
	return DerivativeSecurityXMLLenField{quickfix.FIXInt(val)}
}

func (f DerivativeSecurityXMLLenField) Value() int { return f.Int() }

//DerivativeSecurityXMLSchemaField is a STRING field
type DerivativeSecurityXMLSchemaField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityXMLSchema (1284)
func (f DerivativeSecurityXMLSchemaField) Tag() quickfix.Tag { return tag.DerivativeSecurityXMLSchema }

//NewDerivativeSecurityXMLSchema returns a new DerivativeSecurityXMLSchemaField initialized with val
func NewDerivativeSecurityXMLSchema(val string) DerivativeSecurityXMLSchemaField {
	return DerivativeSecurityXMLSchemaField{quickfix.FIXString(val)}
}

func (f DerivativeSecurityXMLSchemaField) Value() string { return f.String() }

//DerivativeSettlMethodField is a CHAR field
type DerivativeSettlMethodField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSettlMethod (1317)
func (f DerivativeSettlMethodField) Tag() quickfix.Tag { return tag.DerivativeSettlMethod }

//NewDerivativeSettlMethod returns a new DerivativeSettlMethodField initialized with val
func NewDerivativeSettlMethod(val string) DerivativeSettlMethodField {
	return DerivativeSettlMethodField{quickfix.FIXString(val)}
}

func (f DerivativeSettlMethodField) Value() string { return f.String() }

//DerivativeSettleOnOpenFlagField is a STRING field
type DerivativeSettleOnOpenFlagField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSettleOnOpenFlag (1254)
func (f DerivativeSettleOnOpenFlagField) Tag() quickfix.Tag { return tag.DerivativeSettleOnOpenFlag }

//NewDerivativeSettleOnOpenFlag returns a new DerivativeSettleOnOpenFlagField initialized with val
func NewDerivativeSettleOnOpenFlag(val string) DerivativeSettleOnOpenFlagField {
	return DerivativeSettleOnOpenFlagField{quickfix.FIXString(val)}
}

func (f DerivativeSettleOnOpenFlagField) Value() string { return f.String() }

//DerivativeStateOrProvinceOfIssueField is a STRING field
type DerivativeStateOrProvinceOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeStateOrProvinceOfIssue (1259)
func (f DerivativeStateOrProvinceOfIssueField) Tag() quickfix.Tag {
	return tag.DerivativeStateOrProvinceOfIssue
}

//NewDerivativeStateOrProvinceOfIssue returns a new DerivativeStateOrProvinceOfIssueField initialized with val
func NewDerivativeStateOrProvinceOfIssue(val string) DerivativeStateOrProvinceOfIssueField {
	return DerivativeStateOrProvinceOfIssueField{quickfix.FIXString(val)}
}

func (f DerivativeStateOrProvinceOfIssueField) Value() string { return f.String() }

//DerivativeStrikeCurrencyField is a CURRENCY field
type DerivativeStrikeCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeStrikeCurrency (1262)
func (f DerivativeStrikeCurrencyField) Tag() quickfix.Tag { return tag.DerivativeStrikeCurrency }

//NewDerivativeStrikeCurrency returns a new DerivativeStrikeCurrencyField initialized with val
func NewDerivativeStrikeCurrency(val string) DerivativeStrikeCurrencyField {
	return DerivativeStrikeCurrencyField{quickfix.FIXString(val)}
}

func (f DerivativeStrikeCurrencyField) Value() string { return f.String() }

//DerivativeStrikeMultiplierField is a FLOAT field
type DerivativeStrikeMultiplierField struct{ quickfix.FIXDecimal }

//Tag returns tag.DerivativeStrikeMultiplier (1263)
func (f DerivativeStrikeMultiplierField) Tag() quickfix.Tag { return tag.DerivativeStrikeMultiplier }

//NewDerivativeStrikeMultiplier returns a new DerivativeStrikeMultiplierField initialized with val and scale
func NewDerivativeStrikeMultiplier(val decimal.Decimal, scale int32) DerivativeStrikeMultiplierField {
	return DerivativeStrikeMultiplierField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DerivativeStrikeMultiplierField) Value() (val decimal.Decimal) { return f.Decimal }

//DerivativeStrikePriceField is a PRICE field
type DerivativeStrikePriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.DerivativeStrikePrice (1261)
func (f DerivativeStrikePriceField) Tag() quickfix.Tag { return tag.DerivativeStrikePrice }

//NewDerivativeStrikePrice returns a new DerivativeStrikePriceField initialized with val and scale
func NewDerivativeStrikePrice(val decimal.Decimal, scale int32) DerivativeStrikePriceField {
	return DerivativeStrikePriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DerivativeStrikePriceField) Value() (val decimal.Decimal) { return f.Decimal }

//DerivativeStrikeValueField is a FLOAT field
type DerivativeStrikeValueField struct{ quickfix.FIXDecimal }

//Tag returns tag.DerivativeStrikeValue (1264)
func (f DerivativeStrikeValueField) Tag() quickfix.Tag { return tag.DerivativeStrikeValue }

//NewDerivativeStrikeValue returns a new DerivativeStrikeValueField initialized with val and scale
func NewDerivativeStrikeValue(val decimal.Decimal, scale int32) DerivativeStrikeValueField {
	return DerivativeStrikeValueField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DerivativeStrikeValueField) Value() (val decimal.Decimal) { return f.Decimal }

//DerivativeSymbolField is a STRING field
type DerivativeSymbolField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSymbol (1214)
func (f DerivativeSymbolField) Tag() quickfix.Tag { return tag.DerivativeSymbol }

//NewDerivativeSymbol returns a new DerivativeSymbolField initialized with val
func NewDerivativeSymbol(val string) DerivativeSymbolField {
	return DerivativeSymbolField{quickfix.FIXString(val)}
}

func (f DerivativeSymbolField) Value() string { return f.String() }

//DerivativeSymbolSfxField is a STRING field
type DerivativeSymbolSfxField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSymbolSfx (1215)
func (f DerivativeSymbolSfxField) Tag() quickfix.Tag { return tag.DerivativeSymbolSfx }

//NewDerivativeSymbolSfx returns a new DerivativeSymbolSfxField initialized with val
func NewDerivativeSymbolSfx(val string) DerivativeSymbolSfxField {
	return DerivativeSymbolSfxField{quickfix.FIXString(val)}
}

func (f DerivativeSymbolSfxField) Value() string { return f.String() }

//DerivativeTimeUnitField is a STRING field
type DerivativeTimeUnitField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeTimeUnit (1271)
func (f DerivativeTimeUnitField) Tag() quickfix.Tag { return tag.DerivativeTimeUnit }

//NewDerivativeTimeUnit returns a new DerivativeTimeUnitField initialized with val
func NewDerivativeTimeUnit(val string) DerivativeTimeUnitField {
	return DerivativeTimeUnitField{quickfix.FIXString(val)}
}

func (f DerivativeTimeUnitField) Value() string { return f.String() }

//DerivativeUnitOfMeasureField is a STRING field
type DerivativeUnitOfMeasureField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeUnitOfMeasure (1269)
func (f DerivativeUnitOfMeasureField) Tag() quickfix.Tag { return tag.DerivativeUnitOfMeasure }

//NewDerivativeUnitOfMeasure returns a new DerivativeUnitOfMeasureField initialized with val
func NewDerivativeUnitOfMeasure(val string) DerivativeUnitOfMeasureField {
	return DerivativeUnitOfMeasureField{quickfix.FIXString(val)}
}

func (f DerivativeUnitOfMeasureField) Value() string { return f.String() }

//DerivativeUnitOfMeasureQtyField is a QTY field
type DerivativeUnitOfMeasureQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.DerivativeUnitOfMeasureQty (1270)
func (f DerivativeUnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.DerivativeUnitOfMeasureQty }

//NewDerivativeUnitOfMeasureQty returns a new DerivativeUnitOfMeasureQtyField initialized with val and scale
func NewDerivativeUnitOfMeasureQty(val decimal.Decimal, scale int32) DerivativeUnitOfMeasureQtyField {
	return DerivativeUnitOfMeasureQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DerivativeUnitOfMeasureQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//DerivativeValuationMethodField is a STRING field
type DerivativeValuationMethodField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeValuationMethod (1319)
func (f DerivativeValuationMethodField) Tag() quickfix.Tag { return tag.DerivativeValuationMethod }

//NewDerivativeValuationMethod returns a new DerivativeValuationMethodField initialized with val
func NewDerivativeValuationMethod(val string) DerivativeValuationMethodField {
	return DerivativeValuationMethodField{quickfix.FIXString(val)}
}

func (f DerivativeValuationMethodField) Value() string { return f.String() }

//DesignationField is a STRING field
type DesignationField struct{ quickfix.FIXString }

//Tag returns tag.Designation (494)
func (f DesignationField) Tag() quickfix.Tag { return tag.Designation }

//NewDesignation returns a new DesignationField initialized with val
func NewDesignation(val string) DesignationField {
	return DesignationField{quickfix.FIXString(val)}
}

func (f DesignationField) Value() string { return f.String() }

//DeskIDField is a STRING field
type DeskIDField struct{ quickfix.FIXString }

//Tag returns tag.DeskID (284)
func (f DeskIDField) Tag() quickfix.Tag { return tag.DeskID }

//NewDeskID returns a new DeskIDField initialized with val
func NewDeskID(val string) DeskIDField {
	return DeskIDField{quickfix.FIXString(val)}
}

func (f DeskIDField) Value() string { return f.String() }

//DeskOrderHandlingInstField is a enum.DeskOrderHandlingInst field
type DeskOrderHandlingInstField struct{ quickfix.FIXString }

//Tag returns tag.DeskOrderHandlingInst (1035)
func (f DeskOrderHandlingInstField) Tag() quickfix.Tag { return tag.DeskOrderHandlingInst }

func NewDeskOrderHandlingInst(val enum.DeskOrderHandlingInst) DeskOrderHandlingInstField {
	return DeskOrderHandlingInstField{quickfix.FIXString(val)}
}

func (f DeskOrderHandlingInstField) Value() enum.DeskOrderHandlingInst {
	return enum.DeskOrderHandlingInst(f.String())
}

//DeskTypeField is a enum.DeskType field
type DeskTypeField struct{ quickfix.FIXString }

//Tag returns tag.DeskType (1033)
func (f DeskTypeField) Tag() quickfix.Tag { return tag.DeskType }

func NewDeskType(val enum.DeskType) DeskTypeField {
	return DeskTypeField{quickfix.FIXString(val)}
}

func (f DeskTypeField) Value() enum.DeskType { return enum.DeskType(f.String()) }

//DeskTypeSourceField is a enum.DeskTypeSource field
type DeskTypeSourceField struct{ quickfix.FIXString }

//Tag returns tag.DeskTypeSource (1034)
func (f DeskTypeSourceField) Tag() quickfix.Tag { return tag.DeskTypeSource }

func NewDeskTypeSource(val enum.DeskTypeSource) DeskTypeSourceField {
	return DeskTypeSourceField{quickfix.FIXString(val)}
}

func (f DeskTypeSourceField) Value() enum.DeskTypeSource { return enum.DeskTypeSource(f.String()) }

//DetachmentPointField is a PERCENTAGE field
type DetachmentPointField struct{ quickfix.FIXDecimal }

//Tag returns tag.DetachmentPoint (1458)
func (f DetachmentPointField) Tag() quickfix.Tag { return tag.DetachmentPoint }

//NewDetachmentPoint returns a new DetachmentPointField initialized with val and scale
func NewDetachmentPoint(val decimal.Decimal, scale int32) DetachmentPointField {
	return DetachmentPointField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DetachmentPointField) Value() (val decimal.Decimal) { return f.Decimal }

//DiscretionInstField is a enum.DiscretionInst field
type DiscretionInstField struct{ quickfix.FIXString }

//Tag returns tag.DiscretionInst (388)
func (f DiscretionInstField) Tag() quickfix.Tag { return tag.DiscretionInst }

func NewDiscretionInst(val enum.DiscretionInst) DiscretionInstField {
	return DiscretionInstField{quickfix.FIXString(val)}
}

func (f DiscretionInstField) Value() enum.DiscretionInst { return enum.DiscretionInst(f.String()) }

//DiscretionLimitTypeField is a enum.DiscretionLimitType field
type DiscretionLimitTypeField struct{ quickfix.FIXString }

//Tag returns tag.DiscretionLimitType (843)
func (f DiscretionLimitTypeField) Tag() quickfix.Tag { return tag.DiscretionLimitType }

func NewDiscretionLimitType(val enum.DiscretionLimitType) DiscretionLimitTypeField {
	return DiscretionLimitTypeField{quickfix.FIXString(val)}
}

func (f DiscretionLimitTypeField) Value() enum.DiscretionLimitType {
	return enum.DiscretionLimitType(f.String())
}

//DiscretionMoveTypeField is a enum.DiscretionMoveType field
type DiscretionMoveTypeField struct{ quickfix.FIXString }

//Tag returns tag.DiscretionMoveType (841)
func (f DiscretionMoveTypeField) Tag() quickfix.Tag { return tag.DiscretionMoveType }

func NewDiscretionMoveType(val enum.DiscretionMoveType) DiscretionMoveTypeField {
	return DiscretionMoveTypeField{quickfix.FIXString(val)}
}

func (f DiscretionMoveTypeField) Value() enum.DiscretionMoveType {
	return enum.DiscretionMoveType(f.String())
}

//DiscretionOffsetField is a PRICEOFFSET field
type DiscretionOffsetField struct{ quickfix.FIXDecimal }

//Tag returns tag.DiscretionOffset (389)
func (f DiscretionOffsetField) Tag() quickfix.Tag { return tag.DiscretionOffset }

//NewDiscretionOffset returns a new DiscretionOffsetField initialized with val and scale
func NewDiscretionOffset(val decimal.Decimal, scale int32) DiscretionOffsetField {
	return DiscretionOffsetField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DiscretionOffsetField) Value() (val decimal.Decimal) { return f.Decimal }

//DiscretionOffsetTypeField is a enum.DiscretionOffsetType field
type DiscretionOffsetTypeField struct{ quickfix.FIXString }

//Tag returns tag.DiscretionOffsetType (842)
func (f DiscretionOffsetTypeField) Tag() quickfix.Tag { return tag.DiscretionOffsetType }

func NewDiscretionOffsetType(val enum.DiscretionOffsetType) DiscretionOffsetTypeField {
	return DiscretionOffsetTypeField{quickfix.FIXString(val)}
}

func (f DiscretionOffsetTypeField) Value() enum.DiscretionOffsetType {
	return enum.DiscretionOffsetType(f.String())
}

//DiscretionOffsetValueField is a FLOAT field
type DiscretionOffsetValueField struct{ quickfix.FIXDecimal }

//Tag returns tag.DiscretionOffsetValue (389)
func (f DiscretionOffsetValueField) Tag() quickfix.Tag { return tag.DiscretionOffsetValue }

//NewDiscretionOffsetValue returns a new DiscretionOffsetValueField initialized with val and scale
func NewDiscretionOffsetValue(val decimal.Decimal, scale int32) DiscretionOffsetValueField {
	return DiscretionOffsetValueField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DiscretionOffsetValueField) Value() (val decimal.Decimal) { return f.Decimal }

//DiscretionPriceField is a PRICE field
type DiscretionPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.DiscretionPrice (845)
func (f DiscretionPriceField) Tag() quickfix.Tag { return tag.DiscretionPrice }

//NewDiscretionPrice returns a new DiscretionPriceField initialized with val and scale
func NewDiscretionPrice(val decimal.Decimal, scale int32) DiscretionPriceField {
	return DiscretionPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DiscretionPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//DiscretionRoundDirectionField is a enum.DiscretionRoundDirection field
type DiscretionRoundDirectionField struct{ quickfix.FIXString }

//Tag returns tag.DiscretionRoundDirection (844)
func (f DiscretionRoundDirectionField) Tag() quickfix.Tag { return tag.DiscretionRoundDirection }

func NewDiscretionRoundDirection(val enum.DiscretionRoundDirection) DiscretionRoundDirectionField {
	return DiscretionRoundDirectionField{quickfix.FIXString(val)}
}

func (f DiscretionRoundDirectionField) Value() enum.DiscretionRoundDirection {
	return enum.DiscretionRoundDirection(f.String())
}

//DiscretionScopeField is a enum.DiscretionScope field
type DiscretionScopeField struct{ quickfix.FIXString }

//Tag returns tag.DiscretionScope (846)
func (f DiscretionScopeField) Tag() quickfix.Tag { return tag.DiscretionScope }

func NewDiscretionScope(val enum.DiscretionScope) DiscretionScopeField {
	return DiscretionScopeField{quickfix.FIXString(val)}
}

func (f DiscretionScopeField) Value() enum.DiscretionScope { return enum.DiscretionScope(f.String()) }

//DisplayHighQtyField is a QTY field
type DisplayHighQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.DisplayHighQty (1086)
func (f DisplayHighQtyField) Tag() quickfix.Tag { return tag.DisplayHighQty }

//NewDisplayHighQty returns a new DisplayHighQtyField initialized with val and scale
func NewDisplayHighQty(val decimal.Decimal, scale int32) DisplayHighQtyField {
	return DisplayHighQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DisplayHighQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//DisplayLowQtyField is a QTY field
type DisplayLowQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.DisplayLowQty (1085)
func (f DisplayLowQtyField) Tag() quickfix.Tag { return tag.DisplayLowQty }

//NewDisplayLowQty returns a new DisplayLowQtyField initialized with val and scale
func NewDisplayLowQty(val decimal.Decimal, scale int32) DisplayLowQtyField {
	return DisplayLowQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DisplayLowQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//DisplayMethodField is a enum.DisplayMethod field
type DisplayMethodField struct{ quickfix.FIXString }

//Tag returns tag.DisplayMethod (1084)
func (f DisplayMethodField) Tag() quickfix.Tag { return tag.DisplayMethod }

func NewDisplayMethod(val enum.DisplayMethod) DisplayMethodField {
	return DisplayMethodField{quickfix.FIXString(val)}
}

func (f DisplayMethodField) Value() enum.DisplayMethod { return enum.DisplayMethod(f.String()) }

//DisplayMinIncrField is a QTY field
type DisplayMinIncrField struct{ quickfix.FIXDecimal }

//Tag returns tag.DisplayMinIncr (1087)
func (f DisplayMinIncrField) Tag() quickfix.Tag { return tag.DisplayMinIncr }

//NewDisplayMinIncr returns a new DisplayMinIncrField initialized with val and scale
func NewDisplayMinIncr(val decimal.Decimal, scale int32) DisplayMinIncrField {
	return DisplayMinIncrField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DisplayMinIncrField) Value() (val decimal.Decimal) { return f.Decimal }

//DisplayQtyField is a QTY field
type DisplayQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.DisplayQty (1138)
func (f DisplayQtyField) Tag() quickfix.Tag { return tag.DisplayQty }

//NewDisplayQty returns a new DisplayQtyField initialized with val and scale
func NewDisplayQty(val decimal.Decimal, scale int32) DisplayQtyField {
	return DisplayQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DisplayQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//DisplayWhenField is a enum.DisplayWhen field
type DisplayWhenField struct{ quickfix.FIXString }

//Tag returns tag.DisplayWhen (1083)
func (f DisplayWhenField) Tag() quickfix.Tag { return tag.DisplayWhen }

func NewDisplayWhen(val enum.DisplayWhen) DisplayWhenField {
	return DisplayWhenField{quickfix.FIXString(val)}
}

func (f DisplayWhenField) Value() enum.DisplayWhen { return enum.DisplayWhen(f.String()) }

//DistribPaymentMethodField is a enum.DistribPaymentMethod field
type DistribPaymentMethodField struct{ quickfix.FIXString }

//Tag returns tag.DistribPaymentMethod (477)
func (f DistribPaymentMethodField) Tag() quickfix.Tag { return tag.DistribPaymentMethod }

func NewDistribPaymentMethod(val enum.DistribPaymentMethod) DistribPaymentMethodField {
	return DistribPaymentMethodField{quickfix.FIXString(val)}
}

func (f DistribPaymentMethodField) Value() enum.DistribPaymentMethod {
	return enum.DistribPaymentMethod(f.String())
}

//DistribPercentageField is a PERCENTAGE field
type DistribPercentageField struct{ quickfix.FIXDecimal }

//Tag returns tag.DistribPercentage (512)
func (f DistribPercentageField) Tag() quickfix.Tag { return tag.DistribPercentage }

//NewDistribPercentage returns a new DistribPercentageField initialized with val and scale
func NewDistribPercentage(val decimal.Decimal, scale int32) DistribPercentageField {
	return DistribPercentageField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DistribPercentageField) Value() (val decimal.Decimal) { return f.Decimal }

//DividendYieldField is a PERCENTAGE field
type DividendYieldField struct{ quickfix.FIXDecimal }

//Tag returns tag.DividendYield (1380)
func (f DividendYieldField) Tag() quickfix.Tag { return tag.DividendYield }

//NewDividendYield returns a new DividendYieldField initialized with val and scale
func NewDividendYield(val decimal.Decimal, scale int32) DividendYieldField {
	return DividendYieldField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f DividendYieldField) Value() (val decimal.Decimal) { return f.Decimal }

//DlvyInstField is a STRING field
type DlvyInstField struct{ quickfix.FIXString }

//Tag returns tag.DlvyInst (86)
func (f DlvyInstField) Tag() quickfix.Tag { return tag.DlvyInst }

//NewDlvyInst returns a new DlvyInstField initialized with val
func NewDlvyInst(val string) DlvyInstField {
	return DlvyInstField{quickfix.FIXString(val)}
}

func (f DlvyInstField) Value() string { return f.String() }

//DlvyInstTypeField is a enum.DlvyInstType field
type DlvyInstTypeField struct{ quickfix.FIXString }

//Tag returns tag.DlvyInstType (787)
func (f DlvyInstTypeField) Tag() quickfix.Tag { return tag.DlvyInstType }

func NewDlvyInstType(val enum.DlvyInstType) DlvyInstTypeField {
	return DlvyInstTypeField{quickfix.FIXString(val)}
}

func (f DlvyInstTypeField) Value() enum.DlvyInstType { return enum.DlvyInstType(f.String()) }

//DueToRelatedField is a BOOLEAN field
type DueToRelatedField struct{ quickfix.FIXBoolean }

//Tag returns tag.DueToRelated (329)
func (f DueToRelatedField) Tag() quickfix.Tag { return tag.DueToRelated }

//NewDueToRelated returns a new DueToRelatedField initialized with val
func NewDueToRelated(val bool) DueToRelatedField {
	return DueToRelatedField{quickfix.FIXBoolean(val)}
}

func (f DueToRelatedField) Value() bool { return f.Bool() }

//EFPTrackingErrorField is a PERCENTAGE field
type EFPTrackingErrorField struct{ quickfix.FIXDecimal }

//Tag returns tag.EFPTrackingError (405)
func (f EFPTrackingErrorField) Tag() quickfix.Tag { return tag.EFPTrackingError }

//NewEFPTrackingError returns a new EFPTrackingErrorField initialized with val and scale
func NewEFPTrackingError(val decimal.Decimal, scale int32) EFPTrackingErrorField {
	return EFPTrackingErrorField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f EFPTrackingErrorField) Value() (val decimal.Decimal) { return f.Decimal }

//EffectiveTimeField is a UTCTIMESTAMP field
type EffectiveTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.EffectiveTime (168)
func (f EffectiveTimeField) Tag() quickfix.Tag { return tag.EffectiveTime }

//NewEffectiveTime returns a new EffectiveTimeField initialized with val
func NewEffectiveTime(val time.Time) EffectiveTimeField {
	return EffectiveTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewEffectiveTimeNoMillis returns a new EffectiveTimeField initialized with val without millisecs
func NewEffectiveTimeNoMillis(val time.Time) EffectiveTimeField {
	return EffectiveTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f EffectiveTimeField) Value() time.Time { return f.Time }

//EmailThreadIDField is a STRING field
type EmailThreadIDField struct{ quickfix.FIXString }

//Tag returns tag.EmailThreadID (164)
func (f EmailThreadIDField) Tag() quickfix.Tag { return tag.EmailThreadID }

//NewEmailThreadID returns a new EmailThreadIDField initialized with val
func NewEmailThreadID(val string) EmailThreadIDField {
	return EmailThreadIDField{quickfix.FIXString(val)}
}

func (f EmailThreadIDField) Value() string { return f.String() }

//EmailTypeField is a enum.EmailType field
type EmailTypeField struct{ quickfix.FIXString }

//Tag returns tag.EmailType (94)
func (f EmailTypeField) Tag() quickfix.Tag { return tag.EmailType }

func NewEmailType(val enum.EmailType) EmailTypeField {
	return EmailTypeField{quickfix.FIXString(val)}
}

func (f EmailTypeField) Value() enum.EmailType { return enum.EmailType(f.String()) }

//EncodedAllocTextField is a DATA field
type EncodedAllocTextField struct{ quickfix.FIXString }

//Tag returns tag.EncodedAllocText (361)
func (f EncodedAllocTextField) Tag() quickfix.Tag { return tag.EncodedAllocText }

//NewEncodedAllocText returns a new EncodedAllocTextField initialized with val
func NewEncodedAllocText(val string) EncodedAllocTextField {
	return EncodedAllocTextField{quickfix.FIXString(val)}
}

func (f EncodedAllocTextField) Value() string { return f.String() }

//EncodedAllocTextLenField is a LENGTH field
type EncodedAllocTextLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedAllocTextLen (360)
func (f EncodedAllocTextLenField) Tag() quickfix.Tag { return tag.EncodedAllocTextLen }

//NewEncodedAllocTextLen returns a new EncodedAllocTextLenField initialized with val
func NewEncodedAllocTextLen(val int) EncodedAllocTextLenField {
	return EncodedAllocTextLenField{quickfix.FIXInt(val)}
}

func (f EncodedAllocTextLenField) Value() int { return f.Int() }

//EncodedHeadlineField is a DATA field
type EncodedHeadlineField struct{ quickfix.FIXString }

//Tag returns tag.EncodedHeadline (359)
func (f EncodedHeadlineField) Tag() quickfix.Tag { return tag.EncodedHeadline }

//NewEncodedHeadline returns a new EncodedHeadlineField initialized with val
func NewEncodedHeadline(val string) EncodedHeadlineField {
	return EncodedHeadlineField{quickfix.FIXString(val)}
}

func (f EncodedHeadlineField) Value() string { return f.String() }

//EncodedHeadlineLenField is a LENGTH field
type EncodedHeadlineLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedHeadlineLen (358)
func (f EncodedHeadlineLenField) Tag() quickfix.Tag { return tag.EncodedHeadlineLen }

//NewEncodedHeadlineLen returns a new EncodedHeadlineLenField initialized with val
func NewEncodedHeadlineLen(val int) EncodedHeadlineLenField {
	return EncodedHeadlineLenField{quickfix.FIXInt(val)}
}

func (f EncodedHeadlineLenField) Value() int { return f.Int() }

//EncodedIssuerField is a DATA field
type EncodedIssuerField struct{ quickfix.FIXString }

//Tag returns tag.EncodedIssuer (349)
func (f EncodedIssuerField) Tag() quickfix.Tag { return tag.EncodedIssuer }

//NewEncodedIssuer returns a new EncodedIssuerField initialized with val
func NewEncodedIssuer(val string) EncodedIssuerField {
	return EncodedIssuerField{quickfix.FIXString(val)}
}

func (f EncodedIssuerField) Value() string { return f.String() }

//EncodedIssuerLenField is a LENGTH field
type EncodedIssuerLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedIssuerLen (348)
func (f EncodedIssuerLenField) Tag() quickfix.Tag { return tag.EncodedIssuerLen }

//NewEncodedIssuerLen returns a new EncodedIssuerLenField initialized with val
func NewEncodedIssuerLen(val int) EncodedIssuerLenField {
	return EncodedIssuerLenField{quickfix.FIXInt(val)}
}

func (f EncodedIssuerLenField) Value() int { return f.Int() }

//EncodedLegIssuerField is a DATA field
type EncodedLegIssuerField struct{ quickfix.FIXString }

//Tag returns tag.EncodedLegIssuer (619)
func (f EncodedLegIssuerField) Tag() quickfix.Tag { return tag.EncodedLegIssuer }

//NewEncodedLegIssuer returns a new EncodedLegIssuerField initialized with val
func NewEncodedLegIssuer(val string) EncodedLegIssuerField {
	return EncodedLegIssuerField{quickfix.FIXString(val)}
}

func (f EncodedLegIssuerField) Value() string { return f.String() }

//EncodedLegIssuerLenField is a LENGTH field
type EncodedLegIssuerLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedLegIssuerLen (618)
func (f EncodedLegIssuerLenField) Tag() quickfix.Tag { return tag.EncodedLegIssuerLen }

//NewEncodedLegIssuerLen returns a new EncodedLegIssuerLenField initialized with val
func NewEncodedLegIssuerLen(val int) EncodedLegIssuerLenField {
	return EncodedLegIssuerLenField{quickfix.FIXInt(val)}
}

func (f EncodedLegIssuerLenField) Value() int { return f.Int() }

//EncodedLegSecurityDescField is a DATA field
type EncodedLegSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.EncodedLegSecurityDesc (622)
func (f EncodedLegSecurityDescField) Tag() quickfix.Tag { return tag.EncodedLegSecurityDesc }

//NewEncodedLegSecurityDesc returns a new EncodedLegSecurityDescField initialized with val
func NewEncodedLegSecurityDesc(val string) EncodedLegSecurityDescField {
	return EncodedLegSecurityDescField{quickfix.FIXString(val)}
}

func (f EncodedLegSecurityDescField) Value() string { return f.String() }

//EncodedLegSecurityDescLenField is a LENGTH field
type EncodedLegSecurityDescLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedLegSecurityDescLen (621)
func (f EncodedLegSecurityDescLenField) Tag() quickfix.Tag { return tag.EncodedLegSecurityDescLen }

//NewEncodedLegSecurityDescLen returns a new EncodedLegSecurityDescLenField initialized with val
func NewEncodedLegSecurityDescLen(val int) EncodedLegSecurityDescLenField {
	return EncodedLegSecurityDescLenField{quickfix.FIXInt(val)}
}

func (f EncodedLegSecurityDescLenField) Value() int { return f.Int() }

//EncodedListExecInstField is a DATA field
type EncodedListExecInstField struct{ quickfix.FIXString }

//Tag returns tag.EncodedListExecInst (353)
func (f EncodedListExecInstField) Tag() quickfix.Tag { return tag.EncodedListExecInst }

//NewEncodedListExecInst returns a new EncodedListExecInstField initialized with val
func NewEncodedListExecInst(val string) EncodedListExecInstField {
	return EncodedListExecInstField{quickfix.FIXString(val)}
}

func (f EncodedListExecInstField) Value() string { return f.String() }

//EncodedListExecInstLenField is a LENGTH field
type EncodedListExecInstLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedListExecInstLen (352)
func (f EncodedListExecInstLenField) Tag() quickfix.Tag { return tag.EncodedListExecInstLen }

//NewEncodedListExecInstLen returns a new EncodedListExecInstLenField initialized with val
func NewEncodedListExecInstLen(val int) EncodedListExecInstLenField {
	return EncodedListExecInstLenField{quickfix.FIXInt(val)}
}

func (f EncodedListExecInstLenField) Value() int { return f.Int() }

//EncodedListStatusTextField is a DATA field
type EncodedListStatusTextField struct{ quickfix.FIXString }

//Tag returns tag.EncodedListStatusText (446)
func (f EncodedListStatusTextField) Tag() quickfix.Tag { return tag.EncodedListStatusText }

//NewEncodedListStatusText returns a new EncodedListStatusTextField initialized with val
func NewEncodedListStatusText(val string) EncodedListStatusTextField {
	return EncodedListStatusTextField{quickfix.FIXString(val)}
}

func (f EncodedListStatusTextField) Value() string { return f.String() }

//EncodedListStatusTextLenField is a LENGTH field
type EncodedListStatusTextLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedListStatusTextLen (445)
func (f EncodedListStatusTextLenField) Tag() quickfix.Tag { return tag.EncodedListStatusTextLen }

//NewEncodedListStatusTextLen returns a new EncodedListStatusTextLenField initialized with val
func NewEncodedListStatusTextLen(val int) EncodedListStatusTextLenField {
	return EncodedListStatusTextLenField{quickfix.FIXInt(val)}
}

func (f EncodedListStatusTextLenField) Value() int { return f.Int() }

//EncodedMktSegmDescField is a DATA field
type EncodedMktSegmDescField struct{ quickfix.FIXString }

//Tag returns tag.EncodedMktSegmDesc (1398)
func (f EncodedMktSegmDescField) Tag() quickfix.Tag { return tag.EncodedMktSegmDesc }

//NewEncodedMktSegmDesc returns a new EncodedMktSegmDescField initialized with val
func NewEncodedMktSegmDesc(val string) EncodedMktSegmDescField {
	return EncodedMktSegmDescField{quickfix.FIXString(val)}
}

func (f EncodedMktSegmDescField) Value() string { return f.String() }

//EncodedMktSegmDescLenField is a LENGTH field
type EncodedMktSegmDescLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedMktSegmDescLen (1397)
func (f EncodedMktSegmDescLenField) Tag() quickfix.Tag { return tag.EncodedMktSegmDescLen }

//NewEncodedMktSegmDescLen returns a new EncodedMktSegmDescLenField initialized with val
func NewEncodedMktSegmDescLen(val int) EncodedMktSegmDescLenField {
	return EncodedMktSegmDescLenField{quickfix.FIXInt(val)}
}

func (f EncodedMktSegmDescLenField) Value() int { return f.Int() }

//EncodedSecurityDescField is a DATA field
type EncodedSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.EncodedSecurityDesc (351)
func (f EncodedSecurityDescField) Tag() quickfix.Tag { return tag.EncodedSecurityDesc }

//NewEncodedSecurityDesc returns a new EncodedSecurityDescField initialized with val
func NewEncodedSecurityDesc(val string) EncodedSecurityDescField {
	return EncodedSecurityDescField{quickfix.FIXString(val)}
}

func (f EncodedSecurityDescField) Value() string { return f.String() }

//EncodedSecurityDescLenField is a LENGTH field
type EncodedSecurityDescLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedSecurityDescLen (350)
func (f EncodedSecurityDescLenField) Tag() quickfix.Tag { return tag.EncodedSecurityDescLen }

//NewEncodedSecurityDescLen returns a new EncodedSecurityDescLenField initialized with val
func NewEncodedSecurityDescLen(val int) EncodedSecurityDescLenField {
	return EncodedSecurityDescLenField{quickfix.FIXInt(val)}
}

func (f EncodedSecurityDescLenField) Value() int { return f.Int() }

//EncodedSecurityListDescField is a DATA field
type EncodedSecurityListDescField struct{ quickfix.FIXString }

//Tag returns tag.EncodedSecurityListDesc (1469)
func (f EncodedSecurityListDescField) Tag() quickfix.Tag { return tag.EncodedSecurityListDesc }

//NewEncodedSecurityListDesc returns a new EncodedSecurityListDescField initialized with val
func NewEncodedSecurityListDesc(val string) EncodedSecurityListDescField {
	return EncodedSecurityListDescField{quickfix.FIXString(val)}
}

func (f EncodedSecurityListDescField) Value() string { return f.String() }

//EncodedSecurityListDescLenField is a LENGTH field
type EncodedSecurityListDescLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedSecurityListDescLen (1468)
func (f EncodedSecurityListDescLenField) Tag() quickfix.Tag { return tag.EncodedSecurityListDescLen }

//NewEncodedSecurityListDescLen returns a new EncodedSecurityListDescLenField initialized with val
func NewEncodedSecurityListDescLen(val int) EncodedSecurityListDescLenField {
	return EncodedSecurityListDescLenField{quickfix.FIXInt(val)}
}

func (f EncodedSecurityListDescLenField) Value() int { return f.Int() }

//EncodedSubjectField is a DATA field
type EncodedSubjectField struct{ quickfix.FIXString }

//Tag returns tag.EncodedSubject (357)
func (f EncodedSubjectField) Tag() quickfix.Tag { return tag.EncodedSubject }

//NewEncodedSubject returns a new EncodedSubjectField initialized with val
func NewEncodedSubject(val string) EncodedSubjectField {
	return EncodedSubjectField{quickfix.FIXString(val)}
}

func (f EncodedSubjectField) Value() string { return f.String() }

//EncodedSubjectLenField is a LENGTH field
type EncodedSubjectLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedSubjectLen (356)
func (f EncodedSubjectLenField) Tag() quickfix.Tag { return tag.EncodedSubjectLen }

//NewEncodedSubjectLen returns a new EncodedSubjectLenField initialized with val
func NewEncodedSubjectLen(val int) EncodedSubjectLenField {
	return EncodedSubjectLenField{quickfix.FIXInt(val)}
}

func (f EncodedSubjectLenField) Value() int { return f.Int() }

//EncodedSymbolField is a DATA field
type EncodedSymbolField struct{ quickfix.FIXString }

//Tag returns tag.EncodedSymbol (1360)
func (f EncodedSymbolField) Tag() quickfix.Tag { return tag.EncodedSymbol }

//NewEncodedSymbol returns a new EncodedSymbolField initialized with val
func NewEncodedSymbol(val string) EncodedSymbolField {
	return EncodedSymbolField{quickfix.FIXString(val)}
}

func (f EncodedSymbolField) Value() string { return f.String() }

//EncodedSymbolLenField is a LENGTH field
type EncodedSymbolLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedSymbolLen (1359)
func (f EncodedSymbolLenField) Tag() quickfix.Tag { return tag.EncodedSymbolLen }

//NewEncodedSymbolLen returns a new EncodedSymbolLenField initialized with val
func NewEncodedSymbolLen(val int) EncodedSymbolLenField {
	return EncodedSymbolLenField{quickfix.FIXInt(val)}
}

func (f EncodedSymbolLenField) Value() int { return f.Int() }

//EncodedTextField is a DATA field
type EncodedTextField struct{ quickfix.FIXString }

//Tag returns tag.EncodedText (355)
func (f EncodedTextField) Tag() quickfix.Tag { return tag.EncodedText }

//NewEncodedText returns a new EncodedTextField initialized with val
func NewEncodedText(val string) EncodedTextField {
	return EncodedTextField{quickfix.FIXString(val)}
}

func (f EncodedTextField) Value() string { return f.String() }

//EncodedTextLenField is a LENGTH field
type EncodedTextLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedTextLen (354)
func (f EncodedTextLenField) Tag() quickfix.Tag { return tag.EncodedTextLen }

//NewEncodedTextLen returns a new EncodedTextLenField initialized with val
func NewEncodedTextLen(val int) EncodedTextLenField {
	return EncodedTextLenField{quickfix.FIXInt(val)}
}

func (f EncodedTextLenField) Value() int { return f.Int() }

//EncodedUnderlyingIssuerField is a DATA field
type EncodedUnderlyingIssuerField struct{ quickfix.FIXString }

//Tag returns tag.EncodedUnderlyingIssuer (363)
func (f EncodedUnderlyingIssuerField) Tag() quickfix.Tag { return tag.EncodedUnderlyingIssuer }

//NewEncodedUnderlyingIssuer returns a new EncodedUnderlyingIssuerField initialized with val
func NewEncodedUnderlyingIssuer(val string) EncodedUnderlyingIssuerField {
	return EncodedUnderlyingIssuerField{quickfix.FIXString(val)}
}

func (f EncodedUnderlyingIssuerField) Value() string { return f.String() }

//EncodedUnderlyingIssuerLenField is a LENGTH field
type EncodedUnderlyingIssuerLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedUnderlyingIssuerLen (362)
func (f EncodedUnderlyingIssuerLenField) Tag() quickfix.Tag { return tag.EncodedUnderlyingIssuerLen }

//NewEncodedUnderlyingIssuerLen returns a new EncodedUnderlyingIssuerLenField initialized with val
func NewEncodedUnderlyingIssuerLen(val int) EncodedUnderlyingIssuerLenField {
	return EncodedUnderlyingIssuerLenField{quickfix.FIXInt(val)}
}

func (f EncodedUnderlyingIssuerLenField) Value() int { return f.Int() }

//EncodedUnderlyingSecurityDescField is a DATA field
type EncodedUnderlyingSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.EncodedUnderlyingSecurityDesc (365)
func (f EncodedUnderlyingSecurityDescField) Tag() quickfix.Tag {
	return tag.EncodedUnderlyingSecurityDesc
}

//NewEncodedUnderlyingSecurityDesc returns a new EncodedUnderlyingSecurityDescField initialized with val
func NewEncodedUnderlyingSecurityDesc(val string) EncodedUnderlyingSecurityDescField {
	return EncodedUnderlyingSecurityDescField{quickfix.FIXString(val)}
}

func (f EncodedUnderlyingSecurityDescField) Value() string { return f.String() }

//EncodedUnderlyingSecurityDescLenField is a LENGTH field
type EncodedUnderlyingSecurityDescLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedUnderlyingSecurityDescLen (364)
func (f EncodedUnderlyingSecurityDescLenField) Tag() quickfix.Tag {
	return tag.EncodedUnderlyingSecurityDescLen
}

//NewEncodedUnderlyingSecurityDescLen returns a new EncodedUnderlyingSecurityDescLenField initialized with val
func NewEncodedUnderlyingSecurityDescLen(val int) EncodedUnderlyingSecurityDescLenField {
	return EncodedUnderlyingSecurityDescLenField{quickfix.FIXInt(val)}
}

func (f EncodedUnderlyingSecurityDescLenField) Value() int { return f.Int() }

//EncryptMethodField is a enum.EncryptMethod field
type EncryptMethodField struct{ quickfix.FIXString }

//Tag returns tag.EncryptMethod (98)
func (f EncryptMethodField) Tag() quickfix.Tag { return tag.EncryptMethod }

func NewEncryptMethod(val enum.EncryptMethod) EncryptMethodField {
	return EncryptMethodField{quickfix.FIXString(val)}
}

func (f EncryptMethodField) Value() enum.EncryptMethod { return enum.EncryptMethod(f.String()) }

//EncryptedNewPasswordField is a DATA field
type EncryptedNewPasswordField struct{ quickfix.FIXString }

//Tag returns tag.EncryptedNewPassword (1404)
func (f EncryptedNewPasswordField) Tag() quickfix.Tag { return tag.EncryptedNewPassword }

//NewEncryptedNewPassword returns a new EncryptedNewPasswordField initialized with val
func NewEncryptedNewPassword(val string) EncryptedNewPasswordField {
	return EncryptedNewPasswordField{quickfix.FIXString(val)}
}

func (f EncryptedNewPasswordField) Value() string { return f.String() }

//EncryptedNewPasswordLenField is a LENGTH field
type EncryptedNewPasswordLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncryptedNewPasswordLen (1403)
func (f EncryptedNewPasswordLenField) Tag() quickfix.Tag { return tag.EncryptedNewPasswordLen }

//NewEncryptedNewPasswordLen returns a new EncryptedNewPasswordLenField initialized with val
func NewEncryptedNewPasswordLen(val int) EncryptedNewPasswordLenField {
	return EncryptedNewPasswordLenField{quickfix.FIXInt(val)}
}

func (f EncryptedNewPasswordLenField) Value() int { return f.Int() }

//EncryptedPasswordField is a DATA field
type EncryptedPasswordField struct{ quickfix.FIXString }

//Tag returns tag.EncryptedPassword (1402)
func (f EncryptedPasswordField) Tag() quickfix.Tag { return tag.EncryptedPassword }

//NewEncryptedPassword returns a new EncryptedPasswordField initialized with val
func NewEncryptedPassword(val string) EncryptedPasswordField {
	return EncryptedPasswordField{quickfix.FIXString(val)}
}

func (f EncryptedPasswordField) Value() string { return f.String() }

//EncryptedPasswordLenField is a LENGTH field
type EncryptedPasswordLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncryptedPasswordLen (1401)
func (f EncryptedPasswordLenField) Tag() quickfix.Tag { return tag.EncryptedPasswordLen }

//NewEncryptedPasswordLen returns a new EncryptedPasswordLenField initialized with val
func NewEncryptedPasswordLen(val int) EncryptedPasswordLenField {
	return EncryptedPasswordLenField{quickfix.FIXInt(val)}
}

func (f EncryptedPasswordLenField) Value() int { return f.Int() }

//EncryptedPasswordMethodField is a INT field
type EncryptedPasswordMethodField struct{ quickfix.FIXInt }

//Tag returns tag.EncryptedPasswordMethod (1400)
func (f EncryptedPasswordMethodField) Tag() quickfix.Tag { return tag.EncryptedPasswordMethod }

//NewEncryptedPasswordMethod returns a new EncryptedPasswordMethodField initialized with val
func NewEncryptedPasswordMethod(val int) EncryptedPasswordMethodField {
	return EncryptedPasswordMethodField{quickfix.FIXInt(val)}
}

func (f EncryptedPasswordMethodField) Value() int { return f.Int() }

//EndAccruedInterestAmtField is a AMT field
type EndAccruedInterestAmtField struct{ quickfix.FIXDecimal }

//Tag returns tag.EndAccruedInterestAmt (920)
func (f EndAccruedInterestAmtField) Tag() quickfix.Tag { return tag.EndAccruedInterestAmt }

//NewEndAccruedInterestAmt returns a new EndAccruedInterestAmtField initialized with val and scale
func NewEndAccruedInterestAmt(val decimal.Decimal, scale int32) EndAccruedInterestAmtField {
	return EndAccruedInterestAmtField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f EndAccruedInterestAmtField) Value() (val decimal.Decimal) { return f.Decimal }

//EndCashField is a AMT field
type EndCashField struct{ quickfix.FIXDecimal }

//Tag returns tag.EndCash (922)
func (f EndCashField) Tag() quickfix.Tag { return tag.EndCash }

//NewEndCash returns a new EndCashField initialized with val and scale
func NewEndCash(val decimal.Decimal, scale int32) EndCashField {
	return EndCashField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f EndCashField) Value() (val decimal.Decimal) { return f.Decimal }

//EndDateField is a LOCALMKTDATE field
type EndDateField struct{ quickfix.FIXString }

//Tag returns tag.EndDate (917)
func (f EndDateField) Tag() quickfix.Tag { return tag.EndDate }

//NewEndDate returns a new EndDateField initialized with val
func NewEndDate(val string) EndDateField {
	return EndDateField{quickfix.FIXString(val)}
}

func (f EndDateField) Value() string { return f.String() }

//EndMaturityMonthYearField is a MONTHYEAR field
type EndMaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.EndMaturityMonthYear (1226)
func (f EndMaturityMonthYearField) Tag() quickfix.Tag { return tag.EndMaturityMonthYear }

//NewEndMaturityMonthYear returns a new EndMaturityMonthYearField initialized with val
func NewEndMaturityMonthYear(val string) EndMaturityMonthYearField {
	return EndMaturityMonthYearField{quickfix.FIXString(val)}
}

func (f EndMaturityMonthYearField) Value() string { return f.String() }

//EndSeqNoField is a SEQNUM field
type EndSeqNoField struct{ quickfix.FIXInt }

//Tag returns tag.EndSeqNo (16)
func (f EndSeqNoField) Tag() quickfix.Tag { return tag.EndSeqNo }

//NewEndSeqNo returns a new EndSeqNoField initialized with val
func NewEndSeqNo(val int) EndSeqNoField {
	return EndSeqNoField{quickfix.FIXInt(val)}
}

func (f EndSeqNoField) Value() int { return f.Int() }

//EndStrikePxRangeField is a PRICE field
type EndStrikePxRangeField struct{ quickfix.FIXDecimal }

//Tag returns tag.EndStrikePxRange (1203)
func (f EndStrikePxRangeField) Tag() quickfix.Tag { return tag.EndStrikePxRange }

//NewEndStrikePxRange returns a new EndStrikePxRangeField initialized with val and scale
func NewEndStrikePxRange(val decimal.Decimal, scale int32) EndStrikePxRangeField {
	return EndStrikePxRangeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f EndStrikePxRangeField) Value() (val decimal.Decimal) { return f.Decimal }

//EndTickPriceRangeField is a PRICE field
type EndTickPriceRangeField struct{ quickfix.FIXDecimal }

//Tag returns tag.EndTickPriceRange (1207)
func (f EndTickPriceRangeField) Tag() quickfix.Tag { return tag.EndTickPriceRange }

//NewEndTickPriceRange returns a new EndTickPriceRangeField initialized with val and scale
func NewEndTickPriceRange(val decimal.Decimal, scale int32) EndTickPriceRangeField {
	return EndTickPriceRangeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f EndTickPriceRangeField) Value() (val decimal.Decimal) { return f.Decimal }

//EventDateField is a LOCALMKTDATE field
type EventDateField struct{ quickfix.FIXString }

//Tag returns tag.EventDate (866)
func (f EventDateField) Tag() quickfix.Tag { return tag.EventDate }

//NewEventDate returns a new EventDateField initialized with val
func NewEventDate(val string) EventDateField {
	return EventDateField{quickfix.FIXString(val)}
}

func (f EventDateField) Value() string { return f.String() }

//EventPxField is a PRICE field
type EventPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.EventPx (867)
func (f EventPxField) Tag() quickfix.Tag { return tag.EventPx }

//NewEventPx returns a new EventPxField initialized with val and scale
func NewEventPx(val decimal.Decimal, scale int32) EventPxField {
	return EventPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f EventPxField) Value() (val decimal.Decimal) { return f.Decimal }

//EventTextField is a STRING field
type EventTextField struct{ quickfix.FIXString }

//Tag returns tag.EventText (868)
func (f EventTextField) Tag() quickfix.Tag { return tag.EventText }

//NewEventText returns a new EventTextField initialized with val
func NewEventText(val string) EventTextField {
	return EventTextField{quickfix.FIXString(val)}
}

func (f EventTextField) Value() string { return f.String() }

//EventTimeField is a UTCTIMESTAMP field
type EventTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.EventTime (1145)
func (f EventTimeField) Tag() quickfix.Tag { return tag.EventTime }

//NewEventTime returns a new EventTimeField initialized with val
func NewEventTime(val time.Time) EventTimeField {
	return EventTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewEventTimeNoMillis returns a new EventTimeField initialized with val without millisecs
func NewEventTimeNoMillis(val time.Time) EventTimeField {
	return EventTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f EventTimeField) Value() time.Time { return f.Time }

//EventTypeField is a enum.EventType field
type EventTypeField struct{ quickfix.FIXString }

//Tag returns tag.EventType (865)
func (f EventTypeField) Tag() quickfix.Tag { return tag.EventType }

func NewEventType(val enum.EventType) EventTypeField {
	return EventTypeField{quickfix.FIXString(val)}
}

func (f EventTypeField) Value() enum.EventType { return enum.EventType(f.String()) }

//ExDateField is a LOCALMKTDATE field
type ExDateField struct{ quickfix.FIXString }

//Tag returns tag.ExDate (230)
func (f ExDateField) Tag() quickfix.Tag { return tag.ExDate }

//NewExDate returns a new ExDateField initialized with val
func NewExDate(val string) ExDateField {
	return ExDateField{quickfix.FIXString(val)}
}

func (f ExDateField) Value() string { return f.String() }

//ExDestinationField is a enum.ExDestination field
type ExDestinationField struct{ quickfix.FIXString }

//Tag returns tag.ExDestination (100)
func (f ExDestinationField) Tag() quickfix.Tag { return tag.ExDestination }

func NewExDestination(val enum.ExDestination) ExDestinationField {
	return ExDestinationField{quickfix.FIXString(val)}
}

func (f ExDestinationField) Value() enum.ExDestination { return enum.ExDestination(f.String()) }

//ExDestinationIDSourceField is a enum.ExDestinationIDSource field
type ExDestinationIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.ExDestinationIDSource (1133)
func (f ExDestinationIDSourceField) Tag() quickfix.Tag { return tag.ExDestinationIDSource }

func NewExDestinationIDSource(val enum.ExDestinationIDSource) ExDestinationIDSourceField {
	return ExDestinationIDSourceField{quickfix.FIXString(val)}
}

func (f ExDestinationIDSourceField) Value() enum.ExDestinationIDSource {
	return enum.ExDestinationIDSource(f.String())
}

//ExchangeForPhysicalField is a BOOLEAN field
type ExchangeForPhysicalField struct{ quickfix.FIXBoolean }

//Tag returns tag.ExchangeForPhysical (411)
func (f ExchangeForPhysicalField) Tag() quickfix.Tag { return tag.ExchangeForPhysical }

//NewExchangeForPhysical returns a new ExchangeForPhysicalField initialized with val
func NewExchangeForPhysical(val bool) ExchangeForPhysicalField {
	return ExchangeForPhysicalField{quickfix.FIXBoolean(val)}
}

func (f ExchangeForPhysicalField) Value() bool { return f.Bool() }

//ExchangeRuleField is a STRING field
type ExchangeRuleField struct{ quickfix.FIXString }

//Tag returns tag.ExchangeRule (825)
func (f ExchangeRuleField) Tag() quickfix.Tag { return tag.ExchangeRule }

//NewExchangeRule returns a new ExchangeRuleField initialized with val
func NewExchangeRule(val string) ExchangeRuleField {
	return ExchangeRuleField{quickfix.FIXString(val)}
}

func (f ExchangeRuleField) Value() string { return f.String() }

//ExchangeSpecialInstructionsField is a STRING field
type ExchangeSpecialInstructionsField struct{ quickfix.FIXString }

//Tag returns tag.ExchangeSpecialInstructions (1139)
func (f ExchangeSpecialInstructionsField) Tag() quickfix.Tag { return tag.ExchangeSpecialInstructions }

//NewExchangeSpecialInstructions returns a new ExchangeSpecialInstructionsField initialized with val
func NewExchangeSpecialInstructions(val string) ExchangeSpecialInstructionsField {
	return ExchangeSpecialInstructionsField{quickfix.FIXString(val)}
}

func (f ExchangeSpecialInstructionsField) Value() string { return f.String() }

//ExecAckStatusField is a enum.ExecAckStatus field
type ExecAckStatusField struct{ quickfix.FIXString }

//Tag returns tag.ExecAckStatus (1036)
func (f ExecAckStatusField) Tag() quickfix.Tag { return tag.ExecAckStatus }

func NewExecAckStatus(val enum.ExecAckStatus) ExecAckStatusField {
	return ExecAckStatusField{quickfix.FIXString(val)}
}

func (f ExecAckStatusField) Value() enum.ExecAckStatus { return enum.ExecAckStatus(f.String()) }

//ExecBrokerField is a STRING field
type ExecBrokerField struct{ quickfix.FIXString }

//Tag returns tag.ExecBroker (76)
func (f ExecBrokerField) Tag() quickfix.Tag { return tag.ExecBroker }

//NewExecBroker returns a new ExecBrokerField initialized with val
func NewExecBroker(val string) ExecBrokerField {
	return ExecBrokerField{quickfix.FIXString(val)}
}

func (f ExecBrokerField) Value() string { return f.String() }

//ExecIDField is a STRING field
type ExecIDField struct{ quickfix.FIXString }

//Tag returns tag.ExecID (17)
func (f ExecIDField) Tag() quickfix.Tag { return tag.ExecID }

//NewExecID returns a new ExecIDField initialized with val
func NewExecID(val string) ExecIDField {
	return ExecIDField{quickfix.FIXString(val)}
}

func (f ExecIDField) Value() string { return f.String() }

//ExecInstField is a enum.ExecInst field
type ExecInstField struct{ quickfix.FIXString }

//Tag returns tag.ExecInst (18)
func (f ExecInstField) Tag() quickfix.Tag { return tag.ExecInst }

func NewExecInst(val enum.ExecInst) ExecInstField {
	return ExecInstField{quickfix.FIXString(val)}
}

func (f ExecInstField) Value() enum.ExecInst { return enum.ExecInst(f.String()) }

//ExecInstValueField is a CHAR field
type ExecInstValueField struct{ quickfix.FIXString }

//Tag returns tag.ExecInstValue (1308)
func (f ExecInstValueField) Tag() quickfix.Tag { return tag.ExecInstValue }

//NewExecInstValue returns a new ExecInstValueField initialized with val
func NewExecInstValue(val string) ExecInstValueField {
	return ExecInstValueField{quickfix.FIXString(val)}
}

func (f ExecInstValueField) Value() string { return f.String() }

//ExecPriceAdjustmentField is a FLOAT field
type ExecPriceAdjustmentField struct{ quickfix.FIXDecimal }

//Tag returns tag.ExecPriceAdjustment (485)
func (f ExecPriceAdjustmentField) Tag() quickfix.Tag { return tag.ExecPriceAdjustment }

//NewExecPriceAdjustment returns a new ExecPriceAdjustmentField initialized with val and scale
func NewExecPriceAdjustment(val decimal.Decimal, scale int32) ExecPriceAdjustmentField {
	return ExecPriceAdjustmentField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f ExecPriceAdjustmentField) Value() (val decimal.Decimal) { return f.Decimal }

//ExecPriceTypeField is a enum.ExecPriceType field
type ExecPriceTypeField struct{ quickfix.FIXString }

//Tag returns tag.ExecPriceType (484)
func (f ExecPriceTypeField) Tag() quickfix.Tag { return tag.ExecPriceType }

func NewExecPriceType(val enum.ExecPriceType) ExecPriceTypeField {
	return ExecPriceTypeField{quickfix.FIXString(val)}
}

func (f ExecPriceTypeField) Value() enum.ExecPriceType { return enum.ExecPriceType(f.String()) }

//ExecRefIDField is a STRING field
type ExecRefIDField struct{ quickfix.FIXString }

//Tag returns tag.ExecRefID (19)
func (f ExecRefIDField) Tag() quickfix.Tag { return tag.ExecRefID }

//NewExecRefID returns a new ExecRefIDField initialized with val
func NewExecRefID(val string) ExecRefIDField {
	return ExecRefIDField{quickfix.FIXString(val)}
}

func (f ExecRefIDField) Value() string { return f.String() }

//ExecRestatementReasonField is a enum.ExecRestatementReason field
type ExecRestatementReasonField struct{ quickfix.FIXString }

//Tag returns tag.ExecRestatementReason (378)
func (f ExecRestatementReasonField) Tag() quickfix.Tag { return tag.ExecRestatementReason }

func NewExecRestatementReason(val enum.ExecRestatementReason) ExecRestatementReasonField {
	return ExecRestatementReasonField{quickfix.FIXString(val)}
}

func (f ExecRestatementReasonField) Value() enum.ExecRestatementReason {
	return enum.ExecRestatementReason(f.String())
}

//ExecTransTypeField is a enum.ExecTransType field
type ExecTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.ExecTransType (20)
func (f ExecTransTypeField) Tag() quickfix.Tag { return tag.ExecTransType }

func NewExecTransType(val enum.ExecTransType) ExecTransTypeField {
	return ExecTransTypeField{quickfix.FIXString(val)}
}

func (f ExecTransTypeField) Value() enum.ExecTransType { return enum.ExecTransType(f.String()) }

//ExecTypeField is a enum.ExecType field
type ExecTypeField struct{ quickfix.FIXString }

//Tag returns tag.ExecType (150)
func (f ExecTypeField) Tag() quickfix.Tag { return tag.ExecType }

func NewExecType(val enum.ExecType) ExecTypeField {
	return ExecTypeField{quickfix.FIXString(val)}
}

func (f ExecTypeField) Value() enum.ExecType { return enum.ExecType(f.String()) }

//ExecValuationPointField is a UTCTIMESTAMP field
type ExecValuationPointField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.ExecValuationPoint (515)
func (f ExecValuationPointField) Tag() quickfix.Tag { return tag.ExecValuationPoint }

//NewExecValuationPoint returns a new ExecValuationPointField initialized with val
func NewExecValuationPoint(val time.Time) ExecValuationPointField {
	return ExecValuationPointField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewExecValuationPointNoMillis returns a new ExecValuationPointField initialized with val without millisecs
func NewExecValuationPointNoMillis(val time.Time) ExecValuationPointField {
	return ExecValuationPointField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f ExecValuationPointField) Value() time.Time { return f.Time }

//ExerciseMethodField is a enum.ExerciseMethod field
type ExerciseMethodField struct{ quickfix.FIXString }

//Tag returns tag.ExerciseMethod (747)
func (f ExerciseMethodField) Tag() quickfix.Tag { return tag.ExerciseMethod }

func NewExerciseMethod(val enum.ExerciseMethod) ExerciseMethodField {
	return ExerciseMethodField{quickfix.FIXString(val)}
}

func (f ExerciseMethodField) Value() enum.ExerciseMethod { return enum.ExerciseMethod(f.String()) }

//ExerciseStyleField is a enum.ExerciseStyle field
type ExerciseStyleField struct{ quickfix.FIXString }

//Tag returns tag.ExerciseStyle (1194)
func (f ExerciseStyleField) Tag() quickfix.Tag { return tag.ExerciseStyle }

func NewExerciseStyle(val enum.ExerciseStyle) ExerciseStyleField {
	return ExerciseStyleField{quickfix.FIXString(val)}
}

func (f ExerciseStyleField) Value() enum.ExerciseStyle { return enum.ExerciseStyle(f.String()) }

//ExpQtyField is a QTY field
type ExpQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.ExpQty (983)
func (f ExpQtyField) Tag() quickfix.Tag { return tag.ExpQty }

//NewExpQty returns a new ExpQtyField initialized with val and scale
func NewExpQty(val decimal.Decimal, scale int32) ExpQtyField {
	return ExpQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f ExpQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//ExpTypeField is a enum.ExpType field
type ExpTypeField struct{ quickfix.FIXString }

//Tag returns tag.ExpType (982)
func (f ExpTypeField) Tag() quickfix.Tag { return tag.ExpType }

func NewExpType(val enum.ExpType) ExpTypeField {
	return ExpTypeField{quickfix.FIXString(val)}
}

func (f ExpTypeField) Value() enum.ExpType { return enum.ExpType(f.String()) }

//ExpirationCycleField is a enum.ExpirationCycle field
type ExpirationCycleField struct{ quickfix.FIXString }

//Tag returns tag.ExpirationCycle (827)
func (f ExpirationCycleField) Tag() quickfix.Tag { return tag.ExpirationCycle }

func NewExpirationCycle(val enum.ExpirationCycle) ExpirationCycleField {
	return ExpirationCycleField{quickfix.FIXString(val)}
}

func (f ExpirationCycleField) Value() enum.ExpirationCycle { return enum.ExpirationCycle(f.String()) }

//ExpirationQtyTypeField is a enum.ExpirationQtyType field
type ExpirationQtyTypeField struct{ quickfix.FIXString }

//Tag returns tag.ExpirationQtyType (982)
func (f ExpirationQtyTypeField) Tag() quickfix.Tag { return tag.ExpirationQtyType }

func NewExpirationQtyType(val enum.ExpirationQtyType) ExpirationQtyTypeField {
	return ExpirationQtyTypeField{quickfix.FIXString(val)}
}

func (f ExpirationQtyTypeField) Value() enum.ExpirationQtyType {
	return enum.ExpirationQtyType(f.String())
}

//ExpireDateField is a LOCALMKTDATE field
type ExpireDateField struct{ quickfix.FIXString }

//Tag returns tag.ExpireDate (432)
func (f ExpireDateField) Tag() quickfix.Tag { return tag.ExpireDate }

//NewExpireDate returns a new ExpireDateField initialized with val
func NewExpireDate(val string) ExpireDateField {
	return ExpireDateField{quickfix.FIXString(val)}
}

func (f ExpireDateField) Value() string { return f.String() }

//ExpireTimeField is a UTCTIMESTAMP field
type ExpireTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.ExpireTime (126)
func (f ExpireTimeField) Tag() quickfix.Tag { return tag.ExpireTime }

//NewExpireTime returns a new ExpireTimeField initialized with val
func NewExpireTime(val time.Time) ExpireTimeField {
	return ExpireTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewExpireTimeNoMillis returns a new ExpireTimeField initialized with val without millisecs
func NewExpireTimeNoMillis(val time.Time) ExpireTimeField {
	return ExpireTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f ExpireTimeField) Value() time.Time { return f.Time }

//FactorField is a FLOAT field
type FactorField struct{ quickfix.FIXDecimal }

//Tag returns tag.Factor (228)
func (f FactorField) Tag() quickfix.Tag { return tag.Factor }

//NewFactor returns a new FactorField initialized with val and scale
func NewFactor(val decimal.Decimal, scale int32) FactorField {
	return FactorField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f FactorField) Value() (val decimal.Decimal) { return f.Decimal }

//FairValueField is a AMT field
type FairValueField struct{ quickfix.FIXDecimal }

//Tag returns tag.FairValue (406)
func (f FairValueField) Tag() quickfix.Tag { return tag.FairValue }

//NewFairValue returns a new FairValueField initialized with val and scale
func NewFairValue(val decimal.Decimal, scale int32) FairValueField {
	return FairValueField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f FairValueField) Value() (val decimal.Decimal) { return f.Decimal }

//FeeMultiplierField is a FLOAT field
type FeeMultiplierField struct{ quickfix.FIXDecimal }

//Tag returns tag.FeeMultiplier (1329)
func (f FeeMultiplierField) Tag() quickfix.Tag { return tag.FeeMultiplier }

//NewFeeMultiplier returns a new FeeMultiplierField initialized with val and scale
func NewFeeMultiplier(val decimal.Decimal, scale int32) FeeMultiplierField {
	return FeeMultiplierField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f FeeMultiplierField) Value() (val decimal.Decimal) { return f.Decimal }

//FillExecIDField is a STRING field
type FillExecIDField struct{ quickfix.FIXString }

//Tag returns tag.FillExecID (1363)
func (f FillExecIDField) Tag() quickfix.Tag { return tag.FillExecID }

//NewFillExecID returns a new FillExecIDField initialized with val
func NewFillExecID(val string) FillExecIDField {
	return FillExecIDField{quickfix.FIXString(val)}
}

func (f FillExecIDField) Value() string { return f.String() }

//FillLiquidityIndField is a INT field
type FillLiquidityIndField struct{ quickfix.FIXInt }

//Tag returns tag.FillLiquidityInd (1443)
func (f FillLiquidityIndField) Tag() quickfix.Tag { return tag.FillLiquidityInd }

//NewFillLiquidityInd returns a new FillLiquidityIndField initialized with val
func NewFillLiquidityInd(val int) FillLiquidityIndField {
	return FillLiquidityIndField{quickfix.FIXInt(val)}
}

func (f FillLiquidityIndField) Value() int { return f.Int() }

//FillPxField is a PRICE field
type FillPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.FillPx (1364)
func (f FillPxField) Tag() quickfix.Tag { return tag.FillPx }

//NewFillPx returns a new FillPxField initialized with val and scale
func NewFillPx(val decimal.Decimal, scale int32) FillPxField {
	return FillPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f FillPxField) Value() (val decimal.Decimal) { return f.Decimal }

//FillQtyField is a QTY field
type FillQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.FillQty (1365)
func (f FillQtyField) Tag() quickfix.Tag { return tag.FillQty }

//NewFillQty returns a new FillQtyField initialized with val and scale
func NewFillQty(val decimal.Decimal, scale int32) FillQtyField {
	return FillQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f FillQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//FinancialStatusField is a enum.FinancialStatus field
type FinancialStatusField struct{ quickfix.FIXString }

//Tag returns tag.FinancialStatus (291)
func (f FinancialStatusField) Tag() quickfix.Tag { return tag.FinancialStatus }

func NewFinancialStatus(val enum.FinancialStatus) FinancialStatusField {
	return FinancialStatusField{quickfix.FIXString(val)}
}

func (f FinancialStatusField) Value() enum.FinancialStatus { return enum.FinancialStatus(f.String()) }

//FirmTradeIDField is a STRING field
type FirmTradeIDField struct{ quickfix.FIXString }

//Tag returns tag.FirmTradeID (1041)
func (f FirmTradeIDField) Tag() quickfix.Tag { return tag.FirmTradeID }

//NewFirmTradeID returns a new FirmTradeIDField initialized with val
func NewFirmTradeID(val string) FirmTradeIDField {
	return FirmTradeIDField{quickfix.FIXString(val)}
}

func (f FirmTradeIDField) Value() string { return f.String() }

//FirstPxField is a PRICE field
type FirstPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.FirstPx (1025)
func (f FirstPxField) Tag() quickfix.Tag { return tag.FirstPx }

//NewFirstPx returns a new FirstPxField initialized with val and scale
func NewFirstPx(val decimal.Decimal, scale int32) FirstPxField {
	return FirstPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f FirstPxField) Value() (val decimal.Decimal) { return f.Decimal }

//FlexProductEligibilityIndicatorField is a BOOLEAN field
type FlexProductEligibilityIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.FlexProductEligibilityIndicator (1242)
func (f FlexProductEligibilityIndicatorField) Tag() quickfix.Tag {
	return tag.FlexProductEligibilityIndicator
}

//NewFlexProductEligibilityIndicator returns a new FlexProductEligibilityIndicatorField initialized with val
func NewFlexProductEligibilityIndicator(val bool) FlexProductEligibilityIndicatorField {
	return FlexProductEligibilityIndicatorField{quickfix.FIXBoolean(val)}
}

func (f FlexProductEligibilityIndicatorField) Value() bool { return f.Bool() }

//FlexibleIndicatorField is a BOOLEAN field
type FlexibleIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.FlexibleIndicator (1244)
func (f FlexibleIndicatorField) Tag() quickfix.Tag { return tag.FlexibleIndicator }

//NewFlexibleIndicator returns a new FlexibleIndicatorField initialized with val
func NewFlexibleIndicator(val bool) FlexibleIndicatorField {
	return FlexibleIndicatorField{quickfix.FIXBoolean(val)}
}

func (f FlexibleIndicatorField) Value() bool { return f.Bool() }

//FloorPriceField is a PRICE field
type FloorPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.FloorPrice (1200)
func (f FloorPriceField) Tag() quickfix.Tag { return tag.FloorPrice }

//NewFloorPrice returns a new FloorPriceField initialized with val and scale
func NewFloorPrice(val decimal.Decimal, scale int32) FloorPriceField {
	return FloorPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f FloorPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//FlowScheduleTypeField is a enum.FlowScheduleType field
type FlowScheduleTypeField struct{ quickfix.FIXString }

//Tag returns tag.FlowScheduleType (1439)
func (f FlowScheduleTypeField) Tag() quickfix.Tag { return tag.FlowScheduleType }

func NewFlowScheduleType(val enum.FlowScheduleType) FlowScheduleTypeField {
	return FlowScheduleTypeField{quickfix.FIXString(val)}
}

func (f FlowScheduleTypeField) Value() enum.FlowScheduleType { return enum.FlowScheduleType(f.String()) }

//ForexReqField is a BOOLEAN field
type ForexReqField struct{ quickfix.FIXBoolean }

//Tag returns tag.ForexReq (121)
func (f ForexReqField) Tag() quickfix.Tag { return tag.ForexReq }

//NewForexReq returns a new ForexReqField initialized with val
func NewForexReq(val bool) ForexReqField {
	return ForexReqField{quickfix.FIXBoolean(val)}
}

func (f ForexReqField) Value() bool { return f.Bool() }

//FundRenewWaivField is a enum.FundRenewWaiv field
type FundRenewWaivField struct{ quickfix.FIXString }

//Tag returns tag.FundRenewWaiv (497)
func (f FundRenewWaivField) Tag() quickfix.Tag { return tag.FundRenewWaiv }

func NewFundRenewWaiv(val enum.FundRenewWaiv) FundRenewWaivField {
	return FundRenewWaivField{quickfix.FIXString(val)}
}

func (f FundRenewWaivField) Value() enum.FundRenewWaiv { return enum.FundRenewWaiv(f.String()) }

//FutSettDateField is a LOCALMKTDATE field
type FutSettDateField struct{ quickfix.FIXString }

//Tag returns tag.FutSettDate (64)
func (f FutSettDateField) Tag() quickfix.Tag { return tag.FutSettDate }

//NewFutSettDate returns a new FutSettDateField initialized with val
func NewFutSettDate(val string) FutSettDateField {
	return FutSettDateField{quickfix.FIXString(val)}
}

func (f FutSettDateField) Value() string { return f.String() }

//FutSettDate2Field is a LOCALMKTDATE field
type FutSettDate2Field struct{ quickfix.FIXString }

//Tag returns tag.FutSettDate2 (193)
func (f FutSettDate2Field) Tag() quickfix.Tag { return tag.FutSettDate2 }

//NewFutSettDate2 returns a new FutSettDate2Field initialized with val
func NewFutSettDate2(val string) FutSettDate2Field {
	return FutSettDate2Field{quickfix.FIXString(val)}
}

func (f FutSettDate2Field) Value() string { return f.String() }

//FuturesValuationMethodField is a enum.FuturesValuationMethod field
type FuturesValuationMethodField struct{ quickfix.FIXString }

//Tag returns tag.FuturesValuationMethod (1197)
func (f FuturesValuationMethodField) Tag() quickfix.Tag { return tag.FuturesValuationMethod }

func NewFuturesValuationMethod(val enum.FuturesValuationMethod) FuturesValuationMethodField {
	return FuturesValuationMethodField{quickfix.FIXString(val)}
}

func (f FuturesValuationMethodField) Value() enum.FuturesValuationMethod {
	return enum.FuturesValuationMethod(f.String())
}

//GTBookingInstField is a enum.GTBookingInst field
type GTBookingInstField struct{ quickfix.FIXString }

//Tag returns tag.GTBookingInst (427)
func (f GTBookingInstField) Tag() quickfix.Tag { return tag.GTBookingInst }

func NewGTBookingInst(val enum.GTBookingInst) GTBookingInstField {
	return GTBookingInstField{quickfix.FIXString(val)}
}

func (f GTBookingInstField) Value() enum.GTBookingInst { return enum.GTBookingInst(f.String()) }

//GapFillFlagField is a BOOLEAN field
type GapFillFlagField struct{ quickfix.FIXBoolean }

//Tag returns tag.GapFillFlag (123)
func (f GapFillFlagField) Tag() quickfix.Tag { return tag.GapFillFlag }

//NewGapFillFlag returns a new GapFillFlagField initialized with val
func NewGapFillFlag(val bool) GapFillFlagField {
	return GapFillFlagField{quickfix.FIXBoolean(val)}
}

func (f GapFillFlagField) Value() bool { return f.Bool() }

//GrossTradeAmtField is a AMT field
type GrossTradeAmtField struct{ quickfix.FIXDecimal }

//Tag returns tag.GrossTradeAmt (381)
func (f GrossTradeAmtField) Tag() quickfix.Tag { return tag.GrossTradeAmt }

//NewGrossTradeAmt returns a new GrossTradeAmtField initialized with val and scale
func NewGrossTradeAmt(val decimal.Decimal, scale int32) GrossTradeAmtField {
	return GrossTradeAmtField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f GrossTradeAmtField) Value() (val decimal.Decimal) { return f.Decimal }

//HaltReasonCharField is a enum.HaltReasonChar field
type HaltReasonCharField struct{ quickfix.FIXString }

//Tag returns tag.HaltReasonChar (327)
func (f HaltReasonCharField) Tag() quickfix.Tag { return tag.HaltReasonChar }

func NewHaltReasonChar(val enum.HaltReasonChar) HaltReasonCharField {
	return HaltReasonCharField{quickfix.FIXString(val)}
}

func (f HaltReasonCharField) Value() enum.HaltReasonChar { return enum.HaltReasonChar(f.String()) }

//HaltReasonIntField is a enum.HaltReasonInt field
type HaltReasonIntField struct{ quickfix.FIXString }

//Tag returns tag.HaltReasonInt (327)
func (f HaltReasonIntField) Tag() quickfix.Tag { return tag.HaltReasonInt }

func NewHaltReasonInt(val enum.HaltReasonInt) HaltReasonIntField {
	return HaltReasonIntField{quickfix.FIXString(val)}
}

func (f HaltReasonIntField) Value() enum.HaltReasonInt { return enum.HaltReasonInt(f.String()) }

//HandlInstField is a enum.HandlInst field
type HandlInstField struct{ quickfix.FIXString }

//Tag returns tag.HandlInst (21)
func (f HandlInstField) Tag() quickfix.Tag { return tag.HandlInst }

func NewHandlInst(val enum.HandlInst) HandlInstField {
	return HandlInstField{quickfix.FIXString(val)}
}

func (f HandlInstField) Value() enum.HandlInst { return enum.HandlInst(f.String()) }

//HeadlineField is a STRING field
type HeadlineField struct{ quickfix.FIXString }

//Tag returns tag.Headline (148)
func (f HeadlineField) Tag() quickfix.Tag { return tag.Headline }

//NewHeadline returns a new HeadlineField initialized with val
func NewHeadline(val string) HeadlineField {
	return HeadlineField{quickfix.FIXString(val)}
}

func (f HeadlineField) Value() string { return f.String() }

//HeartBtIntField is a INT field
type HeartBtIntField struct{ quickfix.FIXInt }

//Tag returns tag.HeartBtInt (108)
func (f HeartBtIntField) Tag() quickfix.Tag { return tag.HeartBtInt }

//NewHeartBtInt returns a new HeartBtIntField initialized with val
func NewHeartBtInt(val int) HeartBtIntField {
	return HeartBtIntField{quickfix.FIXInt(val)}
}

func (f HeartBtIntField) Value() int { return f.Int() }

//HighLimitPriceField is a PRICE field
type HighLimitPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.HighLimitPrice (1149)
func (f HighLimitPriceField) Tag() quickfix.Tag { return tag.HighLimitPrice }

//NewHighLimitPrice returns a new HighLimitPriceField initialized with val and scale
func NewHighLimitPrice(val decimal.Decimal, scale int32) HighLimitPriceField {
	return HighLimitPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f HighLimitPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//HighPxField is a PRICE field
type HighPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.HighPx (332)
func (f HighPxField) Tag() quickfix.Tag { return tag.HighPx }

//NewHighPx returns a new HighPxField initialized with val and scale
func NewHighPx(val decimal.Decimal, scale int32) HighPxField {
	return HighPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f HighPxField) Value() (val decimal.Decimal) { return f.Decimal }

//HopCompIDField is a STRING field
type HopCompIDField struct{ quickfix.FIXString }

//Tag returns tag.HopCompID (628)
func (f HopCompIDField) Tag() quickfix.Tag { return tag.HopCompID }

//NewHopCompID returns a new HopCompIDField initialized with val
func NewHopCompID(val string) HopCompIDField {
	return HopCompIDField{quickfix.FIXString(val)}
}

func (f HopCompIDField) Value() string { return f.String() }

//HopRefIDField is a SEQNUM field
type HopRefIDField struct{ quickfix.FIXInt }

//Tag returns tag.HopRefID (630)
func (f HopRefIDField) Tag() quickfix.Tag { return tag.HopRefID }

//NewHopRefID returns a new HopRefIDField initialized with val
func NewHopRefID(val int) HopRefIDField {
	return HopRefIDField{quickfix.FIXInt(val)}
}

func (f HopRefIDField) Value() int { return f.Int() }

//HopSendingTimeField is a UTCTIMESTAMP field
type HopSendingTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.HopSendingTime (629)
func (f HopSendingTimeField) Tag() quickfix.Tag { return tag.HopSendingTime }

//NewHopSendingTime returns a new HopSendingTimeField initialized with val
func NewHopSendingTime(val time.Time) HopSendingTimeField {
	return HopSendingTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewHopSendingTimeNoMillis returns a new HopSendingTimeField initialized with val without millisecs
func NewHopSendingTimeNoMillis(val time.Time) HopSendingTimeField {
	return HopSendingTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f HopSendingTimeField) Value() time.Time { return f.Time }

//HostCrossIDField is a STRING field
type HostCrossIDField struct{ quickfix.FIXString }

//Tag returns tag.HostCrossID (961)
func (f HostCrossIDField) Tag() quickfix.Tag { return tag.HostCrossID }

//NewHostCrossID returns a new HostCrossIDField initialized with val
func NewHostCrossID(val string) HostCrossIDField {
	return HostCrossIDField{quickfix.FIXString(val)}
}

func (f HostCrossIDField) Value() string { return f.String() }

//IDSourceField is a enum.IDSource field
type IDSourceField struct{ quickfix.FIXString }

//Tag returns tag.IDSource (22)
func (f IDSourceField) Tag() quickfix.Tag { return tag.IDSource }

func NewIDSource(val enum.IDSource) IDSourceField {
	return IDSourceField{quickfix.FIXString(val)}
}

func (f IDSourceField) Value() enum.IDSource { return enum.IDSource(f.String()) }

//IOIIDField is a STRING field
type IOIIDField struct{ quickfix.FIXString }

//Tag returns tag.IOIID (23)
func (f IOIIDField) Tag() quickfix.Tag { return tag.IOIID }

//NewIOIID returns a new IOIIDField initialized with val
func NewIOIID(val string) IOIIDField {
	return IOIIDField{quickfix.FIXString(val)}
}

func (f IOIIDField) Value() string { return f.String() }

//IOINaturalFlagField is a BOOLEAN field
type IOINaturalFlagField struct{ quickfix.FIXBoolean }

//Tag returns tag.IOINaturalFlag (130)
func (f IOINaturalFlagField) Tag() quickfix.Tag { return tag.IOINaturalFlag }

//NewIOINaturalFlag returns a new IOINaturalFlagField initialized with val
func NewIOINaturalFlag(val bool) IOINaturalFlagField {
	return IOINaturalFlagField{quickfix.FIXBoolean(val)}
}

func (f IOINaturalFlagField) Value() bool { return f.Bool() }

//IOIOthSvcField is a enum.IOIOthSvc field
type IOIOthSvcField struct{ quickfix.FIXString }

//Tag returns tag.IOIOthSvc (24)
func (f IOIOthSvcField) Tag() quickfix.Tag { return tag.IOIOthSvc }

func NewIOIOthSvc(val enum.IOIOthSvc) IOIOthSvcField {
	return IOIOthSvcField{quickfix.FIXString(val)}
}

func (f IOIOthSvcField) Value() enum.IOIOthSvc { return enum.IOIOthSvc(f.String()) }

//IOIQltyIndField is a enum.IOIQltyInd field
type IOIQltyIndField struct{ quickfix.FIXString }

//Tag returns tag.IOIQltyInd (25)
func (f IOIQltyIndField) Tag() quickfix.Tag { return tag.IOIQltyInd }

func NewIOIQltyInd(val enum.IOIQltyInd) IOIQltyIndField {
	return IOIQltyIndField{quickfix.FIXString(val)}
}

func (f IOIQltyIndField) Value() enum.IOIQltyInd { return enum.IOIQltyInd(f.String()) }

//IOIQtyField is a enum.IOIQty field
type IOIQtyField struct{ quickfix.FIXString }

//Tag returns tag.IOIQty (27)
func (f IOIQtyField) Tag() quickfix.Tag { return tag.IOIQty }

func NewIOIQty(val enum.IOIQty) IOIQtyField {
	return IOIQtyField{quickfix.FIXString(val)}
}

func (f IOIQtyField) Value() enum.IOIQty { return enum.IOIQty(f.String()) }

//IOIQualifierField is a enum.IOIQualifier field
type IOIQualifierField struct{ quickfix.FIXString }

//Tag returns tag.IOIQualifier (104)
func (f IOIQualifierField) Tag() quickfix.Tag { return tag.IOIQualifier }

func NewIOIQualifier(val enum.IOIQualifier) IOIQualifierField {
	return IOIQualifierField{quickfix.FIXString(val)}
}

func (f IOIQualifierField) Value() enum.IOIQualifier { return enum.IOIQualifier(f.String()) }

//IOIRefIDField is a STRING field
type IOIRefIDField struct{ quickfix.FIXString }

//Tag returns tag.IOIRefID (26)
func (f IOIRefIDField) Tag() quickfix.Tag { return tag.IOIRefID }

//NewIOIRefID returns a new IOIRefIDField initialized with val
func NewIOIRefID(val string) IOIRefIDField {
	return IOIRefIDField{quickfix.FIXString(val)}
}

func (f IOIRefIDField) Value() string { return f.String() }

//IOISharesField is a enum.IOIShares field
type IOISharesField struct{ quickfix.FIXString }

//Tag returns tag.IOIShares (27)
func (f IOISharesField) Tag() quickfix.Tag { return tag.IOIShares }

func NewIOIShares(val enum.IOIShares) IOISharesField {
	return IOISharesField{quickfix.FIXString(val)}
}

func (f IOISharesField) Value() enum.IOIShares { return enum.IOIShares(f.String()) }

//IOITransTypeField is a enum.IOITransType field
type IOITransTypeField struct{ quickfix.FIXString }

//Tag returns tag.IOITransType (28)
func (f IOITransTypeField) Tag() quickfix.Tag { return tag.IOITransType }

func NewIOITransType(val enum.IOITransType) IOITransTypeField {
	return IOITransTypeField{quickfix.FIXString(val)}
}

func (f IOITransTypeField) Value() enum.IOITransType { return enum.IOITransType(f.String()) }

//IOIidField is a STRING field
type IOIidField struct{ quickfix.FIXString }

//Tag returns tag.IOIid (23)
func (f IOIidField) Tag() quickfix.Tag { return tag.IOIid }

//NewIOIid returns a new IOIidField initialized with val
func NewIOIid(val string) IOIidField {
	return IOIidField{quickfix.FIXString(val)}
}

func (f IOIidField) Value() string { return f.String() }

//ImpliedMarketIndicatorField is a enum.ImpliedMarketIndicator field
type ImpliedMarketIndicatorField struct{ quickfix.FIXString }

//Tag returns tag.ImpliedMarketIndicator (1144)
func (f ImpliedMarketIndicatorField) Tag() quickfix.Tag { return tag.ImpliedMarketIndicator }

func NewImpliedMarketIndicator(val enum.ImpliedMarketIndicator) ImpliedMarketIndicatorField {
	return ImpliedMarketIndicatorField{quickfix.FIXString(val)}
}

func (f ImpliedMarketIndicatorField) Value() enum.ImpliedMarketIndicator {
	return enum.ImpliedMarketIndicator(f.String())
}

//InViewOfCommonField is a BOOLEAN field
type InViewOfCommonField struct{ quickfix.FIXBoolean }

//Tag returns tag.InViewOfCommon (328)
func (f InViewOfCommonField) Tag() quickfix.Tag { return tag.InViewOfCommon }

//NewInViewOfCommon returns a new InViewOfCommonField initialized with val
func NewInViewOfCommon(val bool) InViewOfCommonField {
	return InViewOfCommonField{quickfix.FIXBoolean(val)}
}

func (f InViewOfCommonField) Value() bool { return f.Bool() }

//IncTaxIndField is a enum.IncTaxInd field
type IncTaxIndField struct{ quickfix.FIXString }

//Tag returns tag.IncTaxInd (416)
func (f IncTaxIndField) Tag() quickfix.Tag { return tag.IncTaxInd }

func NewIncTaxInd(val enum.IncTaxInd) IncTaxIndField {
	return IncTaxIndField{quickfix.FIXString(val)}
}

func (f IncTaxIndField) Value() enum.IncTaxInd { return enum.IncTaxInd(f.String()) }

//IndividualAllocIDField is a STRING field
type IndividualAllocIDField struct{ quickfix.FIXString }

//Tag returns tag.IndividualAllocID (467)
func (f IndividualAllocIDField) Tag() quickfix.Tag { return tag.IndividualAllocID }

//NewIndividualAllocID returns a new IndividualAllocIDField initialized with val
func NewIndividualAllocID(val string) IndividualAllocIDField {
	return IndividualAllocIDField{quickfix.FIXString(val)}
}

func (f IndividualAllocIDField) Value() string { return f.String() }

//IndividualAllocRejCodeField is a INT field
type IndividualAllocRejCodeField struct{ quickfix.FIXInt }

//Tag returns tag.IndividualAllocRejCode (776)
func (f IndividualAllocRejCodeField) Tag() quickfix.Tag { return tag.IndividualAllocRejCode }

//NewIndividualAllocRejCode returns a new IndividualAllocRejCodeField initialized with val
func NewIndividualAllocRejCode(val int) IndividualAllocRejCodeField {
	return IndividualAllocRejCodeField{quickfix.FIXInt(val)}
}

func (f IndividualAllocRejCodeField) Value() int { return f.Int() }

//IndividualAllocTypeField is a enum.IndividualAllocType field
type IndividualAllocTypeField struct{ quickfix.FIXString }

//Tag returns tag.IndividualAllocType (992)
func (f IndividualAllocTypeField) Tag() quickfix.Tag { return tag.IndividualAllocType }

func NewIndividualAllocType(val enum.IndividualAllocType) IndividualAllocTypeField {
	return IndividualAllocTypeField{quickfix.FIXString(val)}
}

func (f IndividualAllocTypeField) Value() enum.IndividualAllocType {
	return enum.IndividualAllocType(f.String())
}

//InputSourceField is a STRING field
type InputSourceField struct{ quickfix.FIXString }

//Tag returns tag.InputSource (979)
func (f InputSourceField) Tag() quickfix.Tag { return tag.InputSource }

//NewInputSource returns a new InputSourceField initialized with val
func NewInputSource(val string) InputSourceField {
	return InputSourceField{quickfix.FIXString(val)}
}

func (f InputSourceField) Value() string { return f.String() }

//InstrAttribTypeField is a enum.InstrAttribType field
type InstrAttribTypeField struct{ quickfix.FIXString }

//Tag returns tag.InstrAttribType (871)
func (f InstrAttribTypeField) Tag() quickfix.Tag { return tag.InstrAttribType }

func NewInstrAttribType(val enum.InstrAttribType) InstrAttribTypeField {
	return InstrAttribTypeField{quickfix.FIXString(val)}
}

func (f InstrAttribTypeField) Value() enum.InstrAttribType { return enum.InstrAttribType(f.String()) }

//InstrAttribValueField is a STRING field
type InstrAttribValueField struct{ quickfix.FIXString }

//Tag returns tag.InstrAttribValue (872)
func (f InstrAttribValueField) Tag() quickfix.Tag { return tag.InstrAttribValue }

//NewInstrAttribValue returns a new InstrAttribValueField initialized with val
func NewInstrAttribValue(val string) InstrAttribValueField {
	return InstrAttribValueField{quickfix.FIXString(val)}
}

func (f InstrAttribValueField) Value() string { return f.String() }

//InstrRegistryField is a enum.InstrRegistry field
type InstrRegistryField struct{ quickfix.FIXString }

//Tag returns tag.InstrRegistry (543)
func (f InstrRegistryField) Tag() quickfix.Tag { return tag.InstrRegistry }

func NewInstrRegistry(val enum.InstrRegistry) InstrRegistryField {
	return InstrRegistryField{quickfix.FIXString(val)}
}

func (f InstrRegistryField) Value() enum.InstrRegistry { return enum.InstrRegistry(f.String()) }

//InstrmtAssignmentMethodField is a CHAR field
type InstrmtAssignmentMethodField struct{ quickfix.FIXString }

//Tag returns tag.InstrmtAssignmentMethod (1049)
func (f InstrmtAssignmentMethodField) Tag() quickfix.Tag { return tag.InstrmtAssignmentMethod }

//NewInstrmtAssignmentMethod returns a new InstrmtAssignmentMethodField initialized with val
func NewInstrmtAssignmentMethod(val string) InstrmtAssignmentMethodField {
	return InstrmtAssignmentMethodField{quickfix.FIXString(val)}
}

func (f InstrmtAssignmentMethodField) Value() string { return f.String() }

//InstrumentPartyIDField is a STRING field
type InstrumentPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.InstrumentPartyID (1019)
func (f InstrumentPartyIDField) Tag() quickfix.Tag { return tag.InstrumentPartyID }

//NewInstrumentPartyID returns a new InstrumentPartyIDField initialized with val
func NewInstrumentPartyID(val string) InstrumentPartyIDField {
	return InstrumentPartyIDField{quickfix.FIXString(val)}
}

func (f InstrumentPartyIDField) Value() string { return f.String() }

//InstrumentPartyIDSourceField is a CHAR field
type InstrumentPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.InstrumentPartyIDSource (1050)
func (f InstrumentPartyIDSourceField) Tag() quickfix.Tag { return tag.InstrumentPartyIDSource }

//NewInstrumentPartyIDSource returns a new InstrumentPartyIDSourceField initialized with val
func NewInstrumentPartyIDSource(val string) InstrumentPartyIDSourceField {
	return InstrumentPartyIDSourceField{quickfix.FIXString(val)}
}

func (f InstrumentPartyIDSourceField) Value() string { return f.String() }

//InstrumentPartyRoleField is a INT field
type InstrumentPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.InstrumentPartyRole (1051)
func (f InstrumentPartyRoleField) Tag() quickfix.Tag { return tag.InstrumentPartyRole }

//NewInstrumentPartyRole returns a new InstrumentPartyRoleField initialized with val
func NewInstrumentPartyRole(val int) InstrumentPartyRoleField {
	return InstrumentPartyRoleField{quickfix.FIXInt(val)}
}

func (f InstrumentPartyRoleField) Value() int { return f.Int() }

//InstrumentPartySubIDField is a STRING field
type InstrumentPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.InstrumentPartySubID (1053)
func (f InstrumentPartySubIDField) Tag() quickfix.Tag { return tag.InstrumentPartySubID }

//NewInstrumentPartySubID returns a new InstrumentPartySubIDField initialized with val
func NewInstrumentPartySubID(val string) InstrumentPartySubIDField {
	return InstrumentPartySubIDField{quickfix.FIXString(val)}
}

func (f InstrumentPartySubIDField) Value() string { return f.String() }

//InstrumentPartySubIDTypeField is a INT field
type InstrumentPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.InstrumentPartySubIDType (1054)
func (f InstrumentPartySubIDTypeField) Tag() quickfix.Tag { return tag.InstrumentPartySubIDType }

//NewInstrumentPartySubIDType returns a new InstrumentPartySubIDTypeField initialized with val
func NewInstrumentPartySubIDType(val int) InstrumentPartySubIDTypeField {
	return InstrumentPartySubIDTypeField{quickfix.FIXInt(val)}
}

func (f InstrumentPartySubIDTypeField) Value() int { return f.Int() }

//InterestAccrualDateField is a LOCALMKTDATE field
type InterestAccrualDateField struct{ quickfix.FIXString }

//Tag returns tag.InterestAccrualDate (874)
func (f InterestAccrualDateField) Tag() quickfix.Tag { return tag.InterestAccrualDate }

//NewInterestAccrualDate returns a new InterestAccrualDateField initialized with val
func NewInterestAccrualDate(val string) InterestAccrualDateField {
	return InterestAccrualDateField{quickfix.FIXString(val)}
}

func (f InterestAccrualDateField) Value() string { return f.String() }

//InterestAtMaturityField is a AMT field
type InterestAtMaturityField struct{ quickfix.FIXDecimal }

//Tag returns tag.InterestAtMaturity (738)
func (f InterestAtMaturityField) Tag() quickfix.Tag { return tag.InterestAtMaturity }

//NewInterestAtMaturity returns a new InterestAtMaturityField initialized with val and scale
func NewInterestAtMaturity(val decimal.Decimal, scale int32) InterestAtMaturityField {
	return InterestAtMaturityField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f InterestAtMaturityField) Value() (val decimal.Decimal) { return f.Decimal }

//InvestorCountryOfResidenceField is a COUNTRY field
type InvestorCountryOfResidenceField struct{ quickfix.FIXString }

//Tag returns tag.InvestorCountryOfResidence (475)
func (f InvestorCountryOfResidenceField) Tag() quickfix.Tag { return tag.InvestorCountryOfResidence }

//NewInvestorCountryOfResidence returns a new InvestorCountryOfResidenceField initialized with val
func NewInvestorCountryOfResidence(val string) InvestorCountryOfResidenceField {
	return InvestorCountryOfResidenceField{quickfix.FIXString(val)}
}

func (f InvestorCountryOfResidenceField) Value() string { return f.String() }

//IssueDateField is a LOCALMKTDATE field
type IssueDateField struct{ quickfix.FIXString }

//Tag returns tag.IssueDate (225)
func (f IssueDateField) Tag() quickfix.Tag { return tag.IssueDate }

//NewIssueDate returns a new IssueDateField initialized with val
func NewIssueDate(val string) IssueDateField {
	return IssueDateField{quickfix.FIXString(val)}
}

func (f IssueDateField) Value() string { return f.String() }

//IssuerField is a STRING field
type IssuerField struct{ quickfix.FIXString }

//Tag returns tag.Issuer (106)
func (f IssuerField) Tag() quickfix.Tag { return tag.Issuer }

//NewIssuer returns a new IssuerField initialized with val
func NewIssuer(val string) IssuerField {
	return IssuerField{quickfix.FIXString(val)}
}

func (f IssuerField) Value() string { return f.String() }

//LanguageCodeField is a LANGUAGE field
type LanguageCodeField struct{ quickfix.FIXString }

//Tag returns tag.LanguageCode (1474)
func (f LanguageCodeField) Tag() quickfix.Tag { return tag.LanguageCode }

//NewLanguageCode returns a new LanguageCodeField initialized with val
func NewLanguageCode(val string) LanguageCodeField {
	return LanguageCodeField{quickfix.FIXString(val)}
}

func (f LanguageCodeField) Value() string { return f.String() }

//LastCapacityField is a enum.LastCapacity field
type LastCapacityField struct{ quickfix.FIXString }

//Tag returns tag.LastCapacity (29)
func (f LastCapacityField) Tag() quickfix.Tag { return tag.LastCapacity }

func NewLastCapacity(val enum.LastCapacity) LastCapacityField {
	return LastCapacityField{quickfix.FIXString(val)}
}

func (f LastCapacityField) Value() enum.LastCapacity { return enum.LastCapacity(f.String()) }

//LastForwardPointsField is a PRICEOFFSET field
type LastForwardPointsField struct{ quickfix.FIXDecimal }

//Tag returns tag.LastForwardPoints (195)
func (f LastForwardPointsField) Tag() quickfix.Tag { return tag.LastForwardPoints }

//NewLastForwardPoints returns a new LastForwardPointsField initialized with val and scale
func NewLastForwardPoints(val decimal.Decimal, scale int32) LastForwardPointsField {
	return LastForwardPointsField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LastForwardPointsField) Value() (val decimal.Decimal) { return f.Decimal }

//LastForwardPoints2Field is a PRICEOFFSET field
type LastForwardPoints2Field struct{ quickfix.FIXDecimal }

//Tag returns tag.LastForwardPoints2 (641)
func (f LastForwardPoints2Field) Tag() quickfix.Tag { return tag.LastForwardPoints2 }

//NewLastForwardPoints2 returns a new LastForwardPoints2Field initialized with val and scale
func NewLastForwardPoints2(val decimal.Decimal, scale int32) LastForwardPoints2Field {
	return LastForwardPoints2Field{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LastForwardPoints2Field) Value() (val decimal.Decimal) { return f.Decimal }

//LastFragmentField is a BOOLEAN field
type LastFragmentField struct{ quickfix.FIXBoolean }

//Tag returns tag.LastFragment (893)
func (f LastFragmentField) Tag() quickfix.Tag { return tag.LastFragment }

//NewLastFragment returns a new LastFragmentField initialized with val
func NewLastFragment(val bool) LastFragmentField {
	return LastFragmentField{quickfix.FIXBoolean(val)}
}

func (f LastFragmentField) Value() bool { return f.Bool() }

//LastLiquidityIndField is a enum.LastLiquidityInd field
type LastLiquidityIndField struct{ quickfix.FIXString }

//Tag returns tag.LastLiquidityInd (851)
func (f LastLiquidityIndField) Tag() quickfix.Tag { return tag.LastLiquidityInd }

func NewLastLiquidityInd(val enum.LastLiquidityInd) LastLiquidityIndField {
	return LastLiquidityIndField{quickfix.FIXString(val)}
}

func (f LastLiquidityIndField) Value() enum.LastLiquidityInd { return enum.LastLiquidityInd(f.String()) }

//LastMktField is a EXCHANGE field
type LastMktField struct{ quickfix.FIXString }

//Tag returns tag.LastMkt (30)
func (f LastMktField) Tag() quickfix.Tag { return tag.LastMkt }

//NewLastMkt returns a new LastMktField initialized with val
func NewLastMkt(val string) LastMktField {
	return LastMktField{quickfix.FIXString(val)}
}

func (f LastMktField) Value() string { return f.String() }

//LastMsgSeqNumProcessedField is a SEQNUM field
type LastMsgSeqNumProcessedField struct{ quickfix.FIXInt }

//Tag returns tag.LastMsgSeqNumProcessed (369)
func (f LastMsgSeqNumProcessedField) Tag() quickfix.Tag { return tag.LastMsgSeqNumProcessed }

//NewLastMsgSeqNumProcessed returns a new LastMsgSeqNumProcessedField initialized with val
func NewLastMsgSeqNumProcessed(val int) LastMsgSeqNumProcessedField {
	return LastMsgSeqNumProcessedField{quickfix.FIXInt(val)}
}

func (f LastMsgSeqNumProcessedField) Value() int { return f.Int() }

//LastNetworkResponseIDField is a STRING field
type LastNetworkResponseIDField struct{ quickfix.FIXString }

//Tag returns tag.LastNetworkResponseID (934)
func (f LastNetworkResponseIDField) Tag() quickfix.Tag { return tag.LastNetworkResponseID }

//NewLastNetworkResponseID returns a new LastNetworkResponseIDField initialized with val
func NewLastNetworkResponseID(val string) LastNetworkResponseIDField {
	return LastNetworkResponseIDField{quickfix.FIXString(val)}
}

func (f LastNetworkResponseIDField) Value() string { return f.String() }

//LastParPxField is a PRICE field
type LastParPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.LastParPx (669)
func (f LastParPxField) Tag() quickfix.Tag { return tag.LastParPx }

//NewLastParPx returns a new LastParPxField initialized with val and scale
func NewLastParPx(val decimal.Decimal, scale int32) LastParPxField {
	return LastParPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LastParPxField) Value() (val decimal.Decimal) { return f.Decimal }

//LastPxField is a PRICE field
type LastPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.LastPx (31)
func (f LastPxField) Tag() quickfix.Tag { return tag.LastPx }

//NewLastPx returns a new LastPxField initialized with val and scale
func NewLastPx(val decimal.Decimal, scale int32) LastPxField {
	return LastPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LastPxField) Value() (val decimal.Decimal) { return f.Decimal }

//LastQtyField is a QTY field
type LastQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.LastQty (32)
func (f LastQtyField) Tag() quickfix.Tag { return tag.LastQty }

//NewLastQty returns a new LastQtyField initialized with val and scale
func NewLastQty(val decimal.Decimal, scale int32) LastQtyField {
	return LastQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LastQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//LastRptRequestedField is a BOOLEAN field
type LastRptRequestedField struct{ quickfix.FIXBoolean }

//Tag returns tag.LastRptRequested (912)
func (f LastRptRequestedField) Tag() quickfix.Tag { return tag.LastRptRequested }

//NewLastRptRequested returns a new LastRptRequestedField initialized with val
func NewLastRptRequested(val bool) LastRptRequestedField {
	return LastRptRequestedField{quickfix.FIXBoolean(val)}
}

func (f LastRptRequestedField) Value() bool { return f.Bool() }

//LastSharesField is a QTY field
type LastSharesField struct{ quickfix.FIXDecimal }

//Tag returns tag.LastShares (32)
func (f LastSharesField) Tag() quickfix.Tag { return tag.LastShares }

//NewLastShares returns a new LastSharesField initialized with val and scale
func NewLastShares(val decimal.Decimal, scale int32) LastSharesField {
	return LastSharesField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LastSharesField) Value() (val decimal.Decimal) { return f.Decimal }

//LastSpotRateField is a PRICE field
type LastSpotRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.LastSpotRate (194)
func (f LastSpotRateField) Tag() quickfix.Tag { return tag.LastSpotRate }

//NewLastSpotRate returns a new LastSpotRateField initialized with val and scale
func NewLastSpotRate(val decimal.Decimal, scale int32) LastSpotRateField {
	return LastSpotRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LastSpotRateField) Value() (val decimal.Decimal) { return f.Decimal }

//LastSwapPointsField is a PRICEOFFSET field
type LastSwapPointsField struct{ quickfix.FIXDecimal }

//Tag returns tag.LastSwapPoints (1071)
func (f LastSwapPointsField) Tag() quickfix.Tag { return tag.LastSwapPoints }

//NewLastSwapPoints returns a new LastSwapPointsField initialized with val and scale
func NewLastSwapPoints(val decimal.Decimal, scale int32) LastSwapPointsField {
	return LastSwapPointsField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LastSwapPointsField) Value() (val decimal.Decimal) { return f.Decimal }

//LastUpdateTimeField is a UTCTIMESTAMP field
type LastUpdateTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.LastUpdateTime (779)
func (f LastUpdateTimeField) Tag() quickfix.Tag { return tag.LastUpdateTime }

//NewLastUpdateTime returns a new LastUpdateTimeField initialized with val
func NewLastUpdateTime(val time.Time) LastUpdateTimeField {
	return LastUpdateTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewLastUpdateTimeNoMillis returns a new LastUpdateTimeField initialized with val without millisecs
func NewLastUpdateTimeNoMillis(val time.Time) LastUpdateTimeField {
	return LastUpdateTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f LastUpdateTimeField) Value() time.Time { return f.Time }

//LateIndicatorField is a BOOLEAN field
type LateIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.LateIndicator (978)
func (f LateIndicatorField) Tag() quickfix.Tag { return tag.LateIndicator }

//NewLateIndicator returns a new LateIndicatorField initialized with val
func NewLateIndicator(val bool) LateIndicatorField {
	return LateIndicatorField{quickfix.FIXBoolean(val)}
}

func (f LateIndicatorField) Value() bool { return f.Bool() }

//LeavesQtyField is a QTY field
type LeavesQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.LeavesQty (151)
func (f LeavesQtyField) Tag() quickfix.Tag { return tag.LeavesQty }

//NewLeavesQty returns a new LeavesQtyField initialized with val and scale
func NewLeavesQty(val decimal.Decimal, scale int32) LeavesQtyField {
	return LeavesQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LeavesQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//LegAllocAccountField is a STRING field
type LegAllocAccountField struct{ quickfix.FIXString }

//Tag returns tag.LegAllocAccount (671)
func (f LegAllocAccountField) Tag() quickfix.Tag { return tag.LegAllocAccount }

//NewLegAllocAccount returns a new LegAllocAccountField initialized with val
func NewLegAllocAccount(val string) LegAllocAccountField {
	return LegAllocAccountField{quickfix.FIXString(val)}
}

func (f LegAllocAccountField) Value() string { return f.String() }

//LegAllocAcctIDSourceField is a STRING field
type LegAllocAcctIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.LegAllocAcctIDSource (674)
func (f LegAllocAcctIDSourceField) Tag() quickfix.Tag { return tag.LegAllocAcctIDSource }

//NewLegAllocAcctIDSource returns a new LegAllocAcctIDSourceField initialized with val
func NewLegAllocAcctIDSource(val string) LegAllocAcctIDSourceField {
	return LegAllocAcctIDSourceField{quickfix.FIXString(val)}
}

func (f LegAllocAcctIDSourceField) Value() string { return f.String() }

//LegAllocIDField is a STRING field
type LegAllocIDField struct{ quickfix.FIXString }

//Tag returns tag.LegAllocID (1366)
func (f LegAllocIDField) Tag() quickfix.Tag { return tag.LegAllocID }

//NewLegAllocID returns a new LegAllocIDField initialized with val
func NewLegAllocID(val string) LegAllocIDField {
	return LegAllocIDField{quickfix.FIXString(val)}
}

func (f LegAllocIDField) Value() string { return f.String() }

//LegAllocQtyField is a QTY field
type LegAllocQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegAllocQty (673)
func (f LegAllocQtyField) Tag() quickfix.Tag { return tag.LegAllocQty }

//NewLegAllocQty returns a new LegAllocQtyField initialized with val and scale
func NewLegAllocQty(val decimal.Decimal, scale int32) LegAllocQtyField {
	return LegAllocQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegAllocQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//LegAllocSettlCurrencyField is a CURRENCY field
type LegAllocSettlCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.LegAllocSettlCurrency (1367)
func (f LegAllocSettlCurrencyField) Tag() quickfix.Tag { return tag.LegAllocSettlCurrency }

//NewLegAllocSettlCurrency returns a new LegAllocSettlCurrencyField initialized with val
func NewLegAllocSettlCurrency(val string) LegAllocSettlCurrencyField {
	return LegAllocSettlCurrencyField{quickfix.FIXString(val)}
}

func (f LegAllocSettlCurrencyField) Value() string { return f.String() }

//LegBenchmarkCurveCurrencyField is a CURRENCY field
type LegBenchmarkCurveCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.LegBenchmarkCurveCurrency (676)
func (f LegBenchmarkCurveCurrencyField) Tag() quickfix.Tag { return tag.LegBenchmarkCurveCurrency }

//NewLegBenchmarkCurveCurrency returns a new LegBenchmarkCurveCurrencyField initialized with val
func NewLegBenchmarkCurveCurrency(val string) LegBenchmarkCurveCurrencyField {
	return LegBenchmarkCurveCurrencyField{quickfix.FIXString(val)}
}

func (f LegBenchmarkCurveCurrencyField) Value() string { return f.String() }

//LegBenchmarkCurveNameField is a STRING field
type LegBenchmarkCurveNameField struct{ quickfix.FIXString }

//Tag returns tag.LegBenchmarkCurveName (677)
func (f LegBenchmarkCurveNameField) Tag() quickfix.Tag { return tag.LegBenchmarkCurveName }

//NewLegBenchmarkCurveName returns a new LegBenchmarkCurveNameField initialized with val
func NewLegBenchmarkCurveName(val string) LegBenchmarkCurveNameField {
	return LegBenchmarkCurveNameField{quickfix.FIXString(val)}
}

func (f LegBenchmarkCurveNameField) Value() string { return f.String() }

//LegBenchmarkCurvePointField is a STRING field
type LegBenchmarkCurvePointField struct{ quickfix.FIXString }

//Tag returns tag.LegBenchmarkCurvePoint (678)
func (f LegBenchmarkCurvePointField) Tag() quickfix.Tag { return tag.LegBenchmarkCurvePoint }

//NewLegBenchmarkCurvePoint returns a new LegBenchmarkCurvePointField initialized with val
func NewLegBenchmarkCurvePoint(val string) LegBenchmarkCurvePointField {
	return LegBenchmarkCurvePointField{quickfix.FIXString(val)}
}

func (f LegBenchmarkCurvePointField) Value() string { return f.String() }

//LegBenchmarkPriceField is a PRICE field
type LegBenchmarkPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegBenchmarkPrice (679)
func (f LegBenchmarkPriceField) Tag() quickfix.Tag { return tag.LegBenchmarkPrice }

//NewLegBenchmarkPrice returns a new LegBenchmarkPriceField initialized with val and scale
func NewLegBenchmarkPrice(val decimal.Decimal, scale int32) LegBenchmarkPriceField {
	return LegBenchmarkPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegBenchmarkPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//LegBenchmarkPriceTypeField is a INT field
type LegBenchmarkPriceTypeField struct{ quickfix.FIXInt }

//Tag returns tag.LegBenchmarkPriceType (680)
func (f LegBenchmarkPriceTypeField) Tag() quickfix.Tag { return tag.LegBenchmarkPriceType }

//NewLegBenchmarkPriceType returns a new LegBenchmarkPriceTypeField initialized with val
func NewLegBenchmarkPriceType(val int) LegBenchmarkPriceTypeField {
	return LegBenchmarkPriceTypeField{quickfix.FIXInt(val)}
}

func (f LegBenchmarkPriceTypeField) Value() int { return f.Int() }

//LegBidForwardPointsField is a PRICEOFFSET field
type LegBidForwardPointsField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegBidForwardPoints (1067)
func (f LegBidForwardPointsField) Tag() quickfix.Tag { return tag.LegBidForwardPoints }

//NewLegBidForwardPoints returns a new LegBidForwardPointsField initialized with val and scale
func NewLegBidForwardPoints(val decimal.Decimal, scale int32) LegBidForwardPointsField {
	return LegBidForwardPointsField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegBidForwardPointsField) Value() (val decimal.Decimal) { return f.Decimal }

//LegBidPxField is a PRICE field
type LegBidPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegBidPx (681)
func (f LegBidPxField) Tag() quickfix.Tag { return tag.LegBidPx }

//NewLegBidPx returns a new LegBidPxField initialized with val and scale
func NewLegBidPx(val decimal.Decimal, scale int32) LegBidPxField {
	return LegBidPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegBidPxField) Value() (val decimal.Decimal) { return f.Decimal }

//LegCFICodeField is a STRING field
type LegCFICodeField struct{ quickfix.FIXString }

//Tag returns tag.LegCFICode (608)
func (f LegCFICodeField) Tag() quickfix.Tag { return tag.LegCFICode }

//NewLegCFICode returns a new LegCFICodeField initialized with val
func NewLegCFICode(val string) LegCFICodeField {
	return LegCFICodeField{quickfix.FIXString(val)}
}

func (f LegCFICodeField) Value() string { return f.String() }

//LegCalculatedCcyLastQtyField is a QTY field
type LegCalculatedCcyLastQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegCalculatedCcyLastQty (1074)
func (f LegCalculatedCcyLastQtyField) Tag() quickfix.Tag { return tag.LegCalculatedCcyLastQty }

//NewLegCalculatedCcyLastQty returns a new LegCalculatedCcyLastQtyField initialized with val and scale
func NewLegCalculatedCcyLastQty(val decimal.Decimal, scale int32) LegCalculatedCcyLastQtyField {
	return LegCalculatedCcyLastQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegCalculatedCcyLastQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//LegContractMultiplierField is a FLOAT field
type LegContractMultiplierField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegContractMultiplier (614)
func (f LegContractMultiplierField) Tag() quickfix.Tag { return tag.LegContractMultiplier }

//NewLegContractMultiplier returns a new LegContractMultiplierField initialized with val and scale
func NewLegContractMultiplier(val decimal.Decimal, scale int32) LegContractMultiplierField {
	return LegContractMultiplierField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegContractMultiplierField) Value() (val decimal.Decimal) { return f.Decimal }

//LegContractMultiplierUnitField is a INT field
type LegContractMultiplierUnitField struct{ quickfix.FIXInt }

//Tag returns tag.LegContractMultiplierUnit (1436)
func (f LegContractMultiplierUnitField) Tag() quickfix.Tag { return tag.LegContractMultiplierUnit }

//NewLegContractMultiplierUnit returns a new LegContractMultiplierUnitField initialized with val
func NewLegContractMultiplierUnit(val int) LegContractMultiplierUnitField {
	return LegContractMultiplierUnitField{quickfix.FIXInt(val)}
}

func (f LegContractMultiplierUnitField) Value() int { return f.Int() }

//LegContractSettlMonthField is a MONTHYEAR field
type LegContractSettlMonthField struct{ quickfix.FIXString }

//Tag returns tag.LegContractSettlMonth (955)
func (f LegContractSettlMonthField) Tag() quickfix.Tag { return tag.LegContractSettlMonth }

//NewLegContractSettlMonth returns a new LegContractSettlMonthField initialized with val
func NewLegContractSettlMonth(val string) LegContractSettlMonthField {
	return LegContractSettlMonthField{quickfix.FIXString(val)}
}

func (f LegContractSettlMonthField) Value() string { return f.String() }

//LegCountryOfIssueField is a COUNTRY field
type LegCountryOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.LegCountryOfIssue (596)
func (f LegCountryOfIssueField) Tag() quickfix.Tag { return tag.LegCountryOfIssue }

//NewLegCountryOfIssue returns a new LegCountryOfIssueField initialized with val
func NewLegCountryOfIssue(val string) LegCountryOfIssueField {
	return LegCountryOfIssueField{quickfix.FIXString(val)}
}

func (f LegCountryOfIssueField) Value() string { return f.String() }

//LegCouponPaymentDateField is a LOCALMKTDATE field
type LegCouponPaymentDateField struct{ quickfix.FIXString }

//Tag returns tag.LegCouponPaymentDate (248)
func (f LegCouponPaymentDateField) Tag() quickfix.Tag { return tag.LegCouponPaymentDate }

//NewLegCouponPaymentDate returns a new LegCouponPaymentDateField initialized with val
func NewLegCouponPaymentDate(val string) LegCouponPaymentDateField {
	return LegCouponPaymentDateField{quickfix.FIXString(val)}
}

func (f LegCouponPaymentDateField) Value() string { return f.String() }

//LegCouponRateField is a PERCENTAGE field
type LegCouponRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegCouponRate (615)
func (f LegCouponRateField) Tag() quickfix.Tag { return tag.LegCouponRate }

//NewLegCouponRate returns a new LegCouponRateField initialized with val and scale
func NewLegCouponRate(val decimal.Decimal, scale int32) LegCouponRateField {
	return LegCouponRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegCouponRateField) Value() (val decimal.Decimal) { return f.Decimal }

//LegCoveredOrUncoveredField is a INT field
type LegCoveredOrUncoveredField struct{ quickfix.FIXInt }

//Tag returns tag.LegCoveredOrUncovered (565)
func (f LegCoveredOrUncoveredField) Tag() quickfix.Tag { return tag.LegCoveredOrUncovered }

//NewLegCoveredOrUncovered returns a new LegCoveredOrUncoveredField initialized with val
func NewLegCoveredOrUncovered(val int) LegCoveredOrUncoveredField {
	return LegCoveredOrUncoveredField{quickfix.FIXInt(val)}
}

func (f LegCoveredOrUncoveredField) Value() int { return f.Int() }

//LegCreditRatingField is a STRING field
type LegCreditRatingField struct{ quickfix.FIXString }

//Tag returns tag.LegCreditRating (257)
func (f LegCreditRatingField) Tag() quickfix.Tag { return tag.LegCreditRating }

//NewLegCreditRating returns a new LegCreditRatingField initialized with val
func NewLegCreditRating(val string) LegCreditRatingField {
	return LegCreditRatingField{quickfix.FIXString(val)}
}

func (f LegCreditRatingField) Value() string { return f.String() }

//LegCurrencyField is a CURRENCY field
type LegCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.LegCurrency (556)
func (f LegCurrencyField) Tag() quickfix.Tag { return tag.LegCurrency }

//NewLegCurrency returns a new LegCurrencyField initialized with val
func NewLegCurrency(val string) LegCurrencyField {
	return LegCurrencyField{quickfix.FIXString(val)}
}

func (f LegCurrencyField) Value() string { return f.String() }

//LegCurrencyRatioField is a FLOAT field
type LegCurrencyRatioField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegCurrencyRatio (1383)
func (f LegCurrencyRatioField) Tag() quickfix.Tag { return tag.LegCurrencyRatio }

//NewLegCurrencyRatio returns a new LegCurrencyRatioField initialized with val and scale
func NewLegCurrencyRatio(val decimal.Decimal, scale int32) LegCurrencyRatioField {
	return LegCurrencyRatioField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegCurrencyRatioField) Value() (val decimal.Decimal) { return f.Decimal }

//LegDatedDateField is a LOCALMKTDATE field
type LegDatedDateField struct{ quickfix.FIXString }

//Tag returns tag.LegDatedDate (739)
func (f LegDatedDateField) Tag() quickfix.Tag { return tag.LegDatedDate }

//NewLegDatedDate returns a new LegDatedDateField initialized with val
func NewLegDatedDate(val string) LegDatedDateField {
	return LegDatedDateField{quickfix.FIXString(val)}
}

func (f LegDatedDateField) Value() string { return f.String() }

//LegDividendYieldField is a PERCENTAGE field
type LegDividendYieldField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegDividendYield (1381)
func (f LegDividendYieldField) Tag() quickfix.Tag { return tag.LegDividendYield }

//NewLegDividendYield returns a new LegDividendYieldField initialized with val and scale
func NewLegDividendYield(val decimal.Decimal, scale int32) LegDividendYieldField {
	return LegDividendYieldField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegDividendYieldField) Value() (val decimal.Decimal) { return f.Decimal }

//LegExecInstField is a MULTIPLECHARVALUE field
type LegExecInstField struct{ quickfix.FIXString }

//Tag returns tag.LegExecInst (1384)
func (f LegExecInstField) Tag() quickfix.Tag { return tag.LegExecInst }

//NewLegExecInst returns a new LegExecInstField initialized with val
func NewLegExecInst(val string) LegExecInstField {
	return LegExecInstField{quickfix.FIXString(val)}
}

func (f LegExecInstField) Value() string { return f.String() }

//LegExerciseStyleField is a INT field
type LegExerciseStyleField struct{ quickfix.FIXInt }

//Tag returns tag.LegExerciseStyle (1420)
func (f LegExerciseStyleField) Tag() quickfix.Tag { return tag.LegExerciseStyle }

//NewLegExerciseStyle returns a new LegExerciseStyleField initialized with val
func NewLegExerciseStyle(val int) LegExerciseStyleField {
	return LegExerciseStyleField{quickfix.FIXInt(val)}
}

func (f LegExerciseStyleField) Value() int { return f.Int() }

//LegFactorField is a FLOAT field
type LegFactorField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegFactor (253)
func (f LegFactorField) Tag() quickfix.Tag { return tag.LegFactor }

//NewLegFactor returns a new LegFactorField initialized with val and scale
func NewLegFactor(val decimal.Decimal, scale int32) LegFactorField {
	return LegFactorField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegFactorField) Value() (val decimal.Decimal) { return f.Decimal }

//LegFlowScheduleTypeField is a INT field
type LegFlowScheduleTypeField struct{ quickfix.FIXInt }

//Tag returns tag.LegFlowScheduleType (1440)
func (f LegFlowScheduleTypeField) Tag() quickfix.Tag { return tag.LegFlowScheduleType }

//NewLegFlowScheduleType returns a new LegFlowScheduleTypeField initialized with val
func NewLegFlowScheduleType(val int) LegFlowScheduleTypeField {
	return LegFlowScheduleTypeField{quickfix.FIXInt(val)}
}

func (f LegFlowScheduleTypeField) Value() int { return f.Int() }

//LegFutSettDateField is a LOCALMKTDATE field
type LegFutSettDateField struct{ quickfix.FIXString }

//Tag returns tag.LegFutSettDate (588)
func (f LegFutSettDateField) Tag() quickfix.Tag { return tag.LegFutSettDate }

//NewLegFutSettDate returns a new LegFutSettDateField initialized with val
func NewLegFutSettDate(val string) LegFutSettDateField {
	return LegFutSettDateField{quickfix.FIXString(val)}
}

func (f LegFutSettDateField) Value() string { return f.String() }

//LegGrossTradeAmtField is a AMT field
type LegGrossTradeAmtField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegGrossTradeAmt (1075)
func (f LegGrossTradeAmtField) Tag() quickfix.Tag { return tag.LegGrossTradeAmt }

//NewLegGrossTradeAmt returns a new LegGrossTradeAmtField initialized with val and scale
func NewLegGrossTradeAmt(val decimal.Decimal, scale int32) LegGrossTradeAmtField {
	return LegGrossTradeAmtField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegGrossTradeAmtField) Value() (val decimal.Decimal) { return f.Decimal }

//LegIOIQtyField is a STRING field
type LegIOIQtyField struct{ quickfix.FIXString }

//Tag returns tag.LegIOIQty (682)
func (f LegIOIQtyField) Tag() quickfix.Tag { return tag.LegIOIQty }

//NewLegIOIQty returns a new LegIOIQtyField initialized with val
func NewLegIOIQty(val string) LegIOIQtyField {
	return LegIOIQtyField{quickfix.FIXString(val)}
}

func (f LegIOIQtyField) Value() string { return f.String() }

//LegIndividualAllocIDField is a STRING field
type LegIndividualAllocIDField struct{ quickfix.FIXString }

//Tag returns tag.LegIndividualAllocID (672)
func (f LegIndividualAllocIDField) Tag() quickfix.Tag { return tag.LegIndividualAllocID }

//NewLegIndividualAllocID returns a new LegIndividualAllocIDField initialized with val
func NewLegIndividualAllocID(val string) LegIndividualAllocIDField {
	return LegIndividualAllocIDField{quickfix.FIXString(val)}
}

func (f LegIndividualAllocIDField) Value() string { return f.String() }

//LegInstrRegistryField is a STRING field
type LegInstrRegistryField struct{ quickfix.FIXString }

//Tag returns tag.LegInstrRegistry (599)
func (f LegInstrRegistryField) Tag() quickfix.Tag { return tag.LegInstrRegistry }

//NewLegInstrRegistry returns a new LegInstrRegistryField initialized with val
func NewLegInstrRegistry(val string) LegInstrRegistryField {
	return LegInstrRegistryField{quickfix.FIXString(val)}
}

func (f LegInstrRegistryField) Value() string { return f.String() }

//LegInterestAccrualDateField is a LOCALMKTDATE field
type LegInterestAccrualDateField struct{ quickfix.FIXString }

//Tag returns tag.LegInterestAccrualDate (956)
func (f LegInterestAccrualDateField) Tag() quickfix.Tag { return tag.LegInterestAccrualDate }

//NewLegInterestAccrualDate returns a new LegInterestAccrualDateField initialized with val
func NewLegInterestAccrualDate(val string) LegInterestAccrualDateField {
	return LegInterestAccrualDateField{quickfix.FIXString(val)}
}

func (f LegInterestAccrualDateField) Value() string { return f.String() }

//LegIssueDateField is a LOCALMKTDATE field
type LegIssueDateField struct{ quickfix.FIXString }

//Tag returns tag.LegIssueDate (249)
func (f LegIssueDateField) Tag() quickfix.Tag { return tag.LegIssueDate }

//NewLegIssueDate returns a new LegIssueDateField initialized with val
func NewLegIssueDate(val string) LegIssueDateField {
	return LegIssueDateField{quickfix.FIXString(val)}
}

func (f LegIssueDateField) Value() string { return f.String() }

//LegIssuerField is a STRING field
type LegIssuerField struct{ quickfix.FIXString }

//Tag returns tag.LegIssuer (617)
func (f LegIssuerField) Tag() quickfix.Tag { return tag.LegIssuer }

//NewLegIssuer returns a new LegIssuerField initialized with val
func NewLegIssuer(val string) LegIssuerField {
	return LegIssuerField{quickfix.FIXString(val)}
}

func (f LegIssuerField) Value() string { return f.String() }

//LegLastForwardPointsField is a PRICEOFFSET field
type LegLastForwardPointsField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegLastForwardPoints (1073)
func (f LegLastForwardPointsField) Tag() quickfix.Tag { return tag.LegLastForwardPoints }

//NewLegLastForwardPoints returns a new LegLastForwardPointsField initialized with val and scale
func NewLegLastForwardPoints(val decimal.Decimal, scale int32) LegLastForwardPointsField {
	return LegLastForwardPointsField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegLastForwardPointsField) Value() (val decimal.Decimal) { return f.Decimal }

//LegLastPxField is a PRICE field
type LegLastPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegLastPx (637)
func (f LegLastPxField) Tag() quickfix.Tag { return tag.LegLastPx }

//NewLegLastPx returns a new LegLastPxField initialized with val and scale
func NewLegLastPx(val decimal.Decimal, scale int32) LegLastPxField {
	return LegLastPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegLastPxField) Value() (val decimal.Decimal) { return f.Decimal }

//LegLastQtyField is a QTY field
type LegLastQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegLastQty (1418)
func (f LegLastQtyField) Tag() quickfix.Tag { return tag.LegLastQty }

//NewLegLastQty returns a new LegLastQtyField initialized with val and scale
func NewLegLastQty(val decimal.Decimal, scale int32) LegLastQtyField {
	return LegLastQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegLastQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//LegLocaleOfIssueField is a STRING field
type LegLocaleOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.LegLocaleOfIssue (598)
func (f LegLocaleOfIssueField) Tag() quickfix.Tag { return tag.LegLocaleOfIssue }

//NewLegLocaleOfIssue returns a new LegLocaleOfIssueField initialized with val
func NewLegLocaleOfIssue(val string) LegLocaleOfIssueField {
	return LegLocaleOfIssueField{quickfix.FIXString(val)}
}

func (f LegLocaleOfIssueField) Value() string { return f.String() }

//LegMaturityDateField is a LOCALMKTDATE field
type LegMaturityDateField struct{ quickfix.FIXString }

//Tag returns tag.LegMaturityDate (611)
func (f LegMaturityDateField) Tag() quickfix.Tag { return tag.LegMaturityDate }

//NewLegMaturityDate returns a new LegMaturityDateField initialized with val
func NewLegMaturityDate(val string) LegMaturityDateField {
	return LegMaturityDateField{quickfix.FIXString(val)}
}

func (f LegMaturityDateField) Value() string { return f.String() }

//LegMaturityMonthYearField is a MONTHYEAR field
type LegMaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.LegMaturityMonthYear (610)
func (f LegMaturityMonthYearField) Tag() quickfix.Tag { return tag.LegMaturityMonthYear }

//NewLegMaturityMonthYear returns a new LegMaturityMonthYearField initialized with val
func NewLegMaturityMonthYear(val string) LegMaturityMonthYearField {
	return LegMaturityMonthYearField{quickfix.FIXString(val)}
}

func (f LegMaturityMonthYearField) Value() string { return f.String() }

//LegMaturityTimeField is a TZTIMEONLY field
type LegMaturityTimeField struct{ quickfix.FIXString }

//Tag returns tag.LegMaturityTime (1212)
func (f LegMaturityTimeField) Tag() quickfix.Tag { return tag.LegMaturityTime }

//NewLegMaturityTime returns a new LegMaturityTimeField initialized with val
func NewLegMaturityTime(val string) LegMaturityTimeField {
	return LegMaturityTimeField{quickfix.FIXString(val)}
}

func (f LegMaturityTimeField) Value() string { return f.String() }

//LegNumberField is a INT field
type LegNumberField struct{ quickfix.FIXInt }

//Tag returns tag.LegNumber (1152)
func (f LegNumberField) Tag() quickfix.Tag { return tag.LegNumber }

//NewLegNumber returns a new LegNumberField initialized with val
func NewLegNumber(val int) LegNumberField {
	return LegNumberField{quickfix.FIXInt(val)}
}

func (f LegNumberField) Value() int { return f.Int() }

//LegOfferForwardPointsField is a PRICEOFFSET field
type LegOfferForwardPointsField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegOfferForwardPoints (1068)
func (f LegOfferForwardPointsField) Tag() quickfix.Tag { return tag.LegOfferForwardPoints }

//NewLegOfferForwardPoints returns a new LegOfferForwardPointsField initialized with val and scale
func NewLegOfferForwardPoints(val decimal.Decimal, scale int32) LegOfferForwardPointsField {
	return LegOfferForwardPointsField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegOfferForwardPointsField) Value() (val decimal.Decimal) { return f.Decimal }

//LegOfferPxField is a PRICE field
type LegOfferPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegOfferPx (684)
func (f LegOfferPxField) Tag() quickfix.Tag { return tag.LegOfferPx }

//NewLegOfferPx returns a new LegOfferPxField initialized with val and scale
func NewLegOfferPx(val decimal.Decimal, scale int32) LegOfferPxField {
	return LegOfferPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegOfferPxField) Value() (val decimal.Decimal) { return f.Decimal }

//LegOptAttributeField is a CHAR field
type LegOptAttributeField struct{ quickfix.FIXString }

//Tag returns tag.LegOptAttribute (613)
func (f LegOptAttributeField) Tag() quickfix.Tag { return tag.LegOptAttribute }

//NewLegOptAttribute returns a new LegOptAttributeField initialized with val
func NewLegOptAttribute(val string) LegOptAttributeField {
	return LegOptAttributeField{quickfix.FIXString(val)}
}

func (f LegOptAttributeField) Value() string { return f.String() }

//LegOptionRatioField is a FLOAT field
type LegOptionRatioField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegOptionRatio (1017)
func (f LegOptionRatioField) Tag() quickfix.Tag { return tag.LegOptionRatio }

//NewLegOptionRatio returns a new LegOptionRatioField initialized with val and scale
func NewLegOptionRatio(val decimal.Decimal, scale int32) LegOptionRatioField {
	return LegOptionRatioField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegOptionRatioField) Value() (val decimal.Decimal) { return f.Decimal }

//LegOrderQtyField is a QTY field
type LegOrderQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegOrderQty (685)
func (f LegOrderQtyField) Tag() quickfix.Tag { return tag.LegOrderQty }

//NewLegOrderQty returns a new LegOrderQtyField initialized with val and scale
func NewLegOrderQty(val decimal.Decimal, scale int32) LegOrderQtyField {
	return LegOrderQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegOrderQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//LegPoolField is a STRING field
type LegPoolField struct{ quickfix.FIXString }

//Tag returns tag.LegPool (740)
func (f LegPoolField) Tag() quickfix.Tag { return tag.LegPool }

//NewLegPool returns a new LegPoolField initialized with val
func NewLegPool(val string) LegPoolField {
	return LegPoolField{quickfix.FIXString(val)}
}

func (f LegPoolField) Value() string { return f.String() }

//LegPositionEffectField is a CHAR field
type LegPositionEffectField struct{ quickfix.FIXString }

//Tag returns tag.LegPositionEffect (564)
func (f LegPositionEffectField) Tag() quickfix.Tag { return tag.LegPositionEffect }

//NewLegPositionEffect returns a new LegPositionEffectField initialized with val
func NewLegPositionEffect(val string) LegPositionEffectField {
	return LegPositionEffectField{quickfix.FIXString(val)}
}

func (f LegPositionEffectField) Value() string { return f.String() }

//LegPriceField is a PRICE field
type LegPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegPrice (566)
func (f LegPriceField) Tag() quickfix.Tag { return tag.LegPrice }

//NewLegPrice returns a new LegPriceField initialized with val and scale
func NewLegPrice(val decimal.Decimal, scale int32) LegPriceField {
	return LegPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//LegPriceTypeField is a INT field
type LegPriceTypeField struct{ quickfix.FIXInt }

//Tag returns tag.LegPriceType (686)
func (f LegPriceTypeField) Tag() quickfix.Tag { return tag.LegPriceType }

//NewLegPriceType returns a new LegPriceTypeField initialized with val
func NewLegPriceType(val int) LegPriceTypeField {
	return LegPriceTypeField{quickfix.FIXInt(val)}
}

func (f LegPriceTypeField) Value() int { return f.Int() }

//LegPriceUnitOfMeasureField is a STRING field
type LegPriceUnitOfMeasureField struct{ quickfix.FIXString }

//Tag returns tag.LegPriceUnitOfMeasure (1421)
func (f LegPriceUnitOfMeasureField) Tag() quickfix.Tag { return tag.LegPriceUnitOfMeasure }

//NewLegPriceUnitOfMeasure returns a new LegPriceUnitOfMeasureField initialized with val
func NewLegPriceUnitOfMeasure(val string) LegPriceUnitOfMeasureField {
	return LegPriceUnitOfMeasureField{quickfix.FIXString(val)}
}

func (f LegPriceUnitOfMeasureField) Value() string { return f.String() }

//LegPriceUnitOfMeasureQtyField is a QTY field
type LegPriceUnitOfMeasureQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegPriceUnitOfMeasureQty (1422)
func (f LegPriceUnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.LegPriceUnitOfMeasureQty }

//NewLegPriceUnitOfMeasureQty returns a new LegPriceUnitOfMeasureQtyField initialized with val and scale
func NewLegPriceUnitOfMeasureQty(val decimal.Decimal, scale int32) LegPriceUnitOfMeasureQtyField {
	return LegPriceUnitOfMeasureQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegPriceUnitOfMeasureQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//LegProductField is a INT field
type LegProductField struct{ quickfix.FIXInt }

//Tag returns tag.LegProduct (607)
func (f LegProductField) Tag() quickfix.Tag { return tag.LegProduct }

//NewLegProduct returns a new LegProductField initialized with val
func NewLegProduct(val int) LegProductField {
	return LegProductField{quickfix.FIXInt(val)}
}

func (f LegProductField) Value() int { return f.Int() }

//LegPutOrCallField is a INT field
type LegPutOrCallField struct{ quickfix.FIXInt }

//Tag returns tag.LegPutOrCall (1358)
func (f LegPutOrCallField) Tag() quickfix.Tag { return tag.LegPutOrCall }

//NewLegPutOrCall returns a new LegPutOrCallField initialized with val
func NewLegPutOrCall(val int) LegPutOrCallField {
	return LegPutOrCallField{quickfix.FIXInt(val)}
}

func (f LegPutOrCallField) Value() int { return f.Int() }

//LegQtyField is a QTY field
type LegQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegQty (687)
func (f LegQtyField) Tag() quickfix.Tag { return tag.LegQty }

//NewLegQty returns a new LegQtyField initialized with val and scale
func NewLegQty(val decimal.Decimal, scale int32) LegQtyField {
	return LegQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//LegRatioQtyField is a FLOAT field
type LegRatioQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegRatioQty (623)
func (f LegRatioQtyField) Tag() quickfix.Tag { return tag.LegRatioQty }

//NewLegRatioQty returns a new LegRatioQtyField initialized with val and scale
func NewLegRatioQty(val decimal.Decimal, scale int32) LegRatioQtyField {
	return LegRatioQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegRatioQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//LegRedemptionDateField is a LOCALMKTDATE field
type LegRedemptionDateField struct{ quickfix.FIXString }

//Tag returns tag.LegRedemptionDate (254)
func (f LegRedemptionDateField) Tag() quickfix.Tag { return tag.LegRedemptionDate }

//NewLegRedemptionDate returns a new LegRedemptionDateField initialized with val
func NewLegRedemptionDate(val string) LegRedemptionDateField {
	return LegRedemptionDateField{quickfix.FIXString(val)}
}

func (f LegRedemptionDateField) Value() string { return f.String() }

//LegRefIDField is a STRING field
type LegRefIDField struct{ quickfix.FIXString }

//Tag returns tag.LegRefID (654)
func (f LegRefIDField) Tag() quickfix.Tag { return tag.LegRefID }

//NewLegRefID returns a new LegRefIDField initialized with val
func NewLegRefID(val string) LegRefIDField {
	return LegRefIDField{quickfix.FIXString(val)}
}

func (f LegRefIDField) Value() string { return f.String() }

//LegRepoCollateralSecurityTypeField is a INT field
type LegRepoCollateralSecurityTypeField struct{ quickfix.FIXInt }

//Tag returns tag.LegRepoCollateralSecurityType (250)
func (f LegRepoCollateralSecurityTypeField) Tag() quickfix.Tag {
	return tag.LegRepoCollateralSecurityType
}

//NewLegRepoCollateralSecurityType returns a new LegRepoCollateralSecurityTypeField initialized with val
func NewLegRepoCollateralSecurityType(val int) LegRepoCollateralSecurityTypeField {
	return LegRepoCollateralSecurityTypeField{quickfix.FIXInt(val)}
}

func (f LegRepoCollateralSecurityTypeField) Value() int { return f.Int() }

//LegReportIDField is a STRING field
type LegReportIDField struct{ quickfix.FIXString }

//Tag returns tag.LegReportID (990)
func (f LegReportIDField) Tag() quickfix.Tag { return tag.LegReportID }

//NewLegReportID returns a new LegReportIDField initialized with val
func NewLegReportID(val string) LegReportIDField {
	return LegReportIDField{quickfix.FIXString(val)}
}

func (f LegReportIDField) Value() string { return f.String() }

//LegRepurchaseRateField is a PERCENTAGE field
type LegRepurchaseRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegRepurchaseRate (252)
func (f LegRepurchaseRateField) Tag() quickfix.Tag { return tag.LegRepurchaseRate }

//NewLegRepurchaseRate returns a new LegRepurchaseRateField initialized with val and scale
func NewLegRepurchaseRate(val decimal.Decimal, scale int32) LegRepurchaseRateField {
	return LegRepurchaseRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegRepurchaseRateField) Value() (val decimal.Decimal) { return f.Decimal }

//LegRepurchaseTermField is a INT field
type LegRepurchaseTermField struct{ quickfix.FIXInt }

//Tag returns tag.LegRepurchaseTerm (251)
func (f LegRepurchaseTermField) Tag() quickfix.Tag { return tag.LegRepurchaseTerm }

//NewLegRepurchaseTerm returns a new LegRepurchaseTermField initialized with val
func NewLegRepurchaseTerm(val int) LegRepurchaseTermField {
	return LegRepurchaseTermField{quickfix.FIXInt(val)}
}

func (f LegRepurchaseTermField) Value() int { return f.Int() }

//LegSecurityAltIDField is a STRING field
type LegSecurityAltIDField struct{ quickfix.FIXString }

//Tag returns tag.LegSecurityAltID (605)
func (f LegSecurityAltIDField) Tag() quickfix.Tag { return tag.LegSecurityAltID }

//NewLegSecurityAltID returns a new LegSecurityAltIDField initialized with val
func NewLegSecurityAltID(val string) LegSecurityAltIDField {
	return LegSecurityAltIDField{quickfix.FIXString(val)}
}

func (f LegSecurityAltIDField) Value() string { return f.String() }

//LegSecurityAltIDSourceField is a STRING field
type LegSecurityAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.LegSecurityAltIDSource (606)
func (f LegSecurityAltIDSourceField) Tag() quickfix.Tag { return tag.LegSecurityAltIDSource }

//NewLegSecurityAltIDSource returns a new LegSecurityAltIDSourceField initialized with val
func NewLegSecurityAltIDSource(val string) LegSecurityAltIDSourceField {
	return LegSecurityAltIDSourceField{quickfix.FIXString(val)}
}

func (f LegSecurityAltIDSourceField) Value() string { return f.String() }

//LegSecurityDescField is a STRING field
type LegSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.LegSecurityDesc (620)
func (f LegSecurityDescField) Tag() quickfix.Tag { return tag.LegSecurityDesc }

//NewLegSecurityDesc returns a new LegSecurityDescField initialized with val
func NewLegSecurityDesc(val string) LegSecurityDescField {
	return LegSecurityDescField{quickfix.FIXString(val)}
}

func (f LegSecurityDescField) Value() string { return f.String() }

//LegSecurityExchangeField is a EXCHANGE field
type LegSecurityExchangeField struct{ quickfix.FIXString }

//Tag returns tag.LegSecurityExchange (616)
func (f LegSecurityExchangeField) Tag() quickfix.Tag { return tag.LegSecurityExchange }

//NewLegSecurityExchange returns a new LegSecurityExchangeField initialized with val
func NewLegSecurityExchange(val string) LegSecurityExchangeField {
	return LegSecurityExchangeField{quickfix.FIXString(val)}
}

func (f LegSecurityExchangeField) Value() string { return f.String() }

//LegSecurityIDField is a STRING field
type LegSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.LegSecurityID (602)
func (f LegSecurityIDField) Tag() quickfix.Tag { return tag.LegSecurityID }

//NewLegSecurityID returns a new LegSecurityIDField initialized with val
func NewLegSecurityID(val string) LegSecurityIDField {
	return LegSecurityIDField{quickfix.FIXString(val)}
}

func (f LegSecurityIDField) Value() string { return f.String() }

//LegSecurityIDSourceField is a STRING field
type LegSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.LegSecurityIDSource (603)
func (f LegSecurityIDSourceField) Tag() quickfix.Tag { return tag.LegSecurityIDSource }

//NewLegSecurityIDSource returns a new LegSecurityIDSourceField initialized with val
func NewLegSecurityIDSource(val string) LegSecurityIDSourceField {
	return LegSecurityIDSourceField{quickfix.FIXString(val)}
}

func (f LegSecurityIDSourceField) Value() string { return f.String() }

//LegSecuritySubTypeField is a STRING field
type LegSecuritySubTypeField struct{ quickfix.FIXString }

//Tag returns tag.LegSecuritySubType (764)
func (f LegSecuritySubTypeField) Tag() quickfix.Tag { return tag.LegSecuritySubType }

//NewLegSecuritySubType returns a new LegSecuritySubTypeField initialized with val
func NewLegSecuritySubType(val string) LegSecuritySubTypeField {
	return LegSecuritySubTypeField{quickfix.FIXString(val)}
}

func (f LegSecuritySubTypeField) Value() string { return f.String() }

//LegSecurityTypeField is a STRING field
type LegSecurityTypeField struct{ quickfix.FIXString }

//Tag returns tag.LegSecurityType (609)
func (f LegSecurityTypeField) Tag() quickfix.Tag { return tag.LegSecurityType }

//NewLegSecurityType returns a new LegSecurityTypeField initialized with val
func NewLegSecurityType(val string) LegSecurityTypeField {
	return LegSecurityTypeField{quickfix.FIXString(val)}
}

func (f LegSecurityTypeField) Value() string { return f.String() }

//LegSettlCurrencyField is a CURRENCY field
type LegSettlCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.LegSettlCurrency (675)
func (f LegSettlCurrencyField) Tag() quickfix.Tag { return tag.LegSettlCurrency }

//NewLegSettlCurrency returns a new LegSettlCurrencyField initialized with val
func NewLegSettlCurrency(val string) LegSettlCurrencyField {
	return LegSettlCurrencyField{quickfix.FIXString(val)}
}

func (f LegSettlCurrencyField) Value() string { return f.String() }

//LegSettlDateField is a LOCALMKTDATE field
type LegSettlDateField struct{ quickfix.FIXString }

//Tag returns tag.LegSettlDate (588)
func (f LegSettlDateField) Tag() quickfix.Tag { return tag.LegSettlDate }

//NewLegSettlDate returns a new LegSettlDateField initialized with val
func NewLegSettlDate(val string) LegSettlDateField {
	return LegSettlDateField{quickfix.FIXString(val)}
}

func (f LegSettlDateField) Value() string { return f.String() }

//LegSettlTypeField is a CHAR field
type LegSettlTypeField struct{ quickfix.FIXString }

//Tag returns tag.LegSettlType (587)
func (f LegSettlTypeField) Tag() quickfix.Tag { return tag.LegSettlType }

//NewLegSettlType returns a new LegSettlTypeField initialized with val
func NewLegSettlType(val string) LegSettlTypeField {
	return LegSettlTypeField{quickfix.FIXString(val)}
}

func (f LegSettlTypeField) Value() string { return f.String() }

//LegSettlmntTypField is a CHAR field
type LegSettlmntTypField struct{ quickfix.FIXString }

//Tag returns tag.LegSettlmntTyp (587)
func (f LegSettlmntTypField) Tag() quickfix.Tag { return tag.LegSettlmntTyp }

//NewLegSettlmntTyp returns a new LegSettlmntTypField initialized with val
func NewLegSettlmntTyp(val string) LegSettlmntTypField {
	return LegSettlmntTypField{quickfix.FIXString(val)}
}

func (f LegSettlmntTypField) Value() string { return f.String() }

//LegSideField is a CHAR field
type LegSideField struct{ quickfix.FIXString }

//Tag returns tag.LegSide (624)
func (f LegSideField) Tag() quickfix.Tag { return tag.LegSide }

//NewLegSide returns a new LegSideField initialized with val
func NewLegSide(val string) LegSideField {
	return LegSideField{quickfix.FIXString(val)}
}

func (f LegSideField) Value() string { return f.String() }

//LegStateOrProvinceOfIssueField is a STRING field
type LegStateOrProvinceOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.LegStateOrProvinceOfIssue (597)
func (f LegStateOrProvinceOfIssueField) Tag() quickfix.Tag { return tag.LegStateOrProvinceOfIssue }

//NewLegStateOrProvinceOfIssue returns a new LegStateOrProvinceOfIssueField initialized with val
func NewLegStateOrProvinceOfIssue(val string) LegStateOrProvinceOfIssueField {
	return LegStateOrProvinceOfIssueField{quickfix.FIXString(val)}
}

func (f LegStateOrProvinceOfIssueField) Value() string { return f.String() }

//LegStipulationTypeField is a STRING field
type LegStipulationTypeField struct{ quickfix.FIXString }

//Tag returns tag.LegStipulationType (688)
func (f LegStipulationTypeField) Tag() quickfix.Tag { return tag.LegStipulationType }

//NewLegStipulationType returns a new LegStipulationTypeField initialized with val
func NewLegStipulationType(val string) LegStipulationTypeField {
	return LegStipulationTypeField{quickfix.FIXString(val)}
}

func (f LegStipulationTypeField) Value() string { return f.String() }

//LegStipulationValueField is a STRING field
type LegStipulationValueField struct{ quickfix.FIXString }

//Tag returns tag.LegStipulationValue (689)
func (f LegStipulationValueField) Tag() quickfix.Tag { return tag.LegStipulationValue }

//NewLegStipulationValue returns a new LegStipulationValueField initialized with val
func NewLegStipulationValue(val string) LegStipulationValueField {
	return LegStipulationValueField{quickfix.FIXString(val)}
}

func (f LegStipulationValueField) Value() string { return f.String() }

//LegStrikeCurrencyField is a CURRENCY field
type LegStrikeCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.LegStrikeCurrency (942)
func (f LegStrikeCurrencyField) Tag() quickfix.Tag { return tag.LegStrikeCurrency }

//NewLegStrikeCurrency returns a new LegStrikeCurrencyField initialized with val
func NewLegStrikeCurrency(val string) LegStrikeCurrencyField {
	return LegStrikeCurrencyField{quickfix.FIXString(val)}
}

func (f LegStrikeCurrencyField) Value() string { return f.String() }

//LegStrikePriceField is a PRICE field
type LegStrikePriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegStrikePrice (612)
func (f LegStrikePriceField) Tag() quickfix.Tag { return tag.LegStrikePrice }

//NewLegStrikePrice returns a new LegStrikePriceField initialized with val and scale
func NewLegStrikePrice(val decimal.Decimal, scale int32) LegStrikePriceField {
	return LegStrikePriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegStrikePriceField) Value() (val decimal.Decimal) { return f.Decimal }

//LegSwapTypeField is a enum.LegSwapType field
type LegSwapTypeField struct{ quickfix.FIXString }

//Tag returns tag.LegSwapType (690)
func (f LegSwapTypeField) Tag() quickfix.Tag { return tag.LegSwapType }

func NewLegSwapType(val enum.LegSwapType) LegSwapTypeField {
	return LegSwapTypeField{quickfix.FIXString(val)}
}

func (f LegSwapTypeField) Value() enum.LegSwapType { return enum.LegSwapType(f.String()) }

//LegSymbolField is a STRING field
type LegSymbolField struct{ quickfix.FIXString }

//Tag returns tag.LegSymbol (600)
func (f LegSymbolField) Tag() quickfix.Tag { return tag.LegSymbol }

//NewLegSymbol returns a new LegSymbolField initialized with val
func NewLegSymbol(val string) LegSymbolField {
	return LegSymbolField{quickfix.FIXString(val)}
}

func (f LegSymbolField) Value() string { return f.String() }

//LegSymbolSfxField is a STRING field
type LegSymbolSfxField struct{ quickfix.FIXString }

//Tag returns tag.LegSymbolSfx (601)
func (f LegSymbolSfxField) Tag() quickfix.Tag { return tag.LegSymbolSfx }

//NewLegSymbolSfx returns a new LegSymbolSfxField initialized with val
func NewLegSymbolSfx(val string) LegSymbolSfxField {
	return LegSymbolSfxField{quickfix.FIXString(val)}
}

func (f LegSymbolSfxField) Value() string { return f.String() }

//LegTimeUnitField is a STRING field
type LegTimeUnitField struct{ quickfix.FIXString }

//Tag returns tag.LegTimeUnit (1001)
func (f LegTimeUnitField) Tag() quickfix.Tag { return tag.LegTimeUnit }

//NewLegTimeUnit returns a new LegTimeUnitField initialized with val
func NewLegTimeUnit(val string) LegTimeUnitField {
	return LegTimeUnitField{quickfix.FIXString(val)}
}

func (f LegTimeUnitField) Value() string { return f.String() }

//LegUnitOfMeasureField is a STRING field
type LegUnitOfMeasureField struct{ quickfix.FIXString }

//Tag returns tag.LegUnitOfMeasure (999)
func (f LegUnitOfMeasureField) Tag() quickfix.Tag { return tag.LegUnitOfMeasure }

//NewLegUnitOfMeasure returns a new LegUnitOfMeasureField initialized with val
func NewLegUnitOfMeasure(val string) LegUnitOfMeasureField {
	return LegUnitOfMeasureField{quickfix.FIXString(val)}
}

func (f LegUnitOfMeasureField) Value() string { return f.String() }

//LegUnitOfMeasureQtyField is a QTY field
type LegUnitOfMeasureQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegUnitOfMeasureQty (1224)
func (f LegUnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.LegUnitOfMeasureQty }

//NewLegUnitOfMeasureQty returns a new LegUnitOfMeasureQtyField initialized with val and scale
func NewLegUnitOfMeasureQty(val decimal.Decimal, scale int32) LegUnitOfMeasureQtyField {
	return LegUnitOfMeasureQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegUnitOfMeasureQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//LegVolatilityField is a FLOAT field
type LegVolatilityField struct{ quickfix.FIXDecimal }

//Tag returns tag.LegVolatility (1379)
func (f LegVolatilityField) Tag() quickfix.Tag { return tag.LegVolatility }

//NewLegVolatility returns a new LegVolatilityField initialized with val and scale
func NewLegVolatility(val decimal.Decimal, scale int32) LegVolatilityField {
	return LegVolatilityField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LegVolatilityField) Value() (val decimal.Decimal) { return f.Decimal }

//LegalConfirmField is a BOOLEAN field
type LegalConfirmField struct{ quickfix.FIXBoolean }

//Tag returns tag.LegalConfirm (650)
func (f LegalConfirmField) Tag() quickfix.Tag { return tag.LegalConfirm }

//NewLegalConfirm returns a new LegalConfirmField initialized with val
func NewLegalConfirm(val bool) LegalConfirmField {
	return LegalConfirmField{quickfix.FIXBoolean(val)}
}

func (f LegalConfirmField) Value() bool { return f.Bool() }

//LinesOfTextField is a NUMINGROUP field
type LinesOfTextField struct{ quickfix.FIXInt }

//Tag returns tag.LinesOfText (33)
func (f LinesOfTextField) Tag() quickfix.Tag { return tag.LinesOfText }

//NewLinesOfText returns a new LinesOfTextField initialized with val
func NewLinesOfText(val int) LinesOfTextField {
	return LinesOfTextField{quickfix.FIXInt(val)}
}

func (f LinesOfTextField) Value() int { return f.Int() }

//LiquidityIndTypeField is a enum.LiquidityIndType field
type LiquidityIndTypeField struct{ quickfix.FIXString }

//Tag returns tag.LiquidityIndType (409)
func (f LiquidityIndTypeField) Tag() quickfix.Tag { return tag.LiquidityIndType }

func NewLiquidityIndType(val enum.LiquidityIndType) LiquidityIndTypeField {
	return LiquidityIndTypeField{quickfix.FIXString(val)}
}

func (f LiquidityIndTypeField) Value() enum.LiquidityIndType { return enum.LiquidityIndType(f.String()) }

//LiquidityNumSecuritiesField is a INT field
type LiquidityNumSecuritiesField struct{ quickfix.FIXInt }

//Tag returns tag.LiquidityNumSecurities (441)
func (f LiquidityNumSecuritiesField) Tag() quickfix.Tag { return tag.LiquidityNumSecurities }

//NewLiquidityNumSecurities returns a new LiquidityNumSecuritiesField initialized with val
func NewLiquidityNumSecurities(val int) LiquidityNumSecuritiesField {
	return LiquidityNumSecuritiesField{quickfix.FIXInt(val)}
}

func (f LiquidityNumSecuritiesField) Value() int { return f.Int() }

//LiquidityPctHighField is a PERCENTAGE field
type LiquidityPctHighField struct{ quickfix.FIXDecimal }

//Tag returns tag.LiquidityPctHigh (403)
func (f LiquidityPctHighField) Tag() quickfix.Tag { return tag.LiquidityPctHigh }

//NewLiquidityPctHigh returns a new LiquidityPctHighField initialized with val and scale
func NewLiquidityPctHigh(val decimal.Decimal, scale int32) LiquidityPctHighField {
	return LiquidityPctHighField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LiquidityPctHighField) Value() (val decimal.Decimal) { return f.Decimal }

//LiquidityPctLowField is a PERCENTAGE field
type LiquidityPctLowField struct{ quickfix.FIXDecimal }

//Tag returns tag.LiquidityPctLow (402)
func (f LiquidityPctLowField) Tag() quickfix.Tag { return tag.LiquidityPctLow }

//NewLiquidityPctLow returns a new LiquidityPctLowField initialized with val and scale
func NewLiquidityPctLow(val decimal.Decimal, scale int32) LiquidityPctLowField {
	return LiquidityPctLowField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LiquidityPctLowField) Value() (val decimal.Decimal) { return f.Decimal }

//LiquidityValueField is a AMT field
type LiquidityValueField struct{ quickfix.FIXDecimal }

//Tag returns tag.LiquidityValue (404)
func (f LiquidityValueField) Tag() quickfix.Tag { return tag.LiquidityValue }

//NewLiquidityValue returns a new LiquidityValueField initialized with val and scale
func NewLiquidityValue(val decimal.Decimal, scale int32) LiquidityValueField {
	return LiquidityValueField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LiquidityValueField) Value() (val decimal.Decimal) { return f.Decimal }

//ListExecInstField is a STRING field
type ListExecInstField struct{ quickfix.FIXString }

//Tag returns tag.ListExecInst (69)
func (f ListExecInstField) Tag() quickfix.Tag { return tag.ListExecInst }

//NewListExecInst returns a new ListExecInstField initialized with val
func NewListExecInst(val string) ListExecInstField {
	return ListExecInstField{quickfix.FIXString(val)}
}

func (f ListExecInstField) Value() string { return f.String() }

//ListExecInstTypeField is a enum.ListExecInstType field
type ListExecInstTypeField struct{ quickfix.FIXString }

//Tag returns tag.ListExecInstType (433)
func (f ListExecInstTypeField) Tag() quickfix.Tag { return tag.ListExecInstType }

func NewListExecInstType(val enum.ListExecInstType) ListExecInstTypeField {
	return ListExecInstTypeField{quickfix.FIXString(val)}
}

func (f ListExecInstTypeField) Value() enum.ListExecInstType { return enum.ListExecInstType(f.String()) }

//ListIDField is a STRING field
type ListIDField struct{ quickfix.FIXString }

//Tag returns tag.ListID (66)
func (f ListIDField) Tag() quickfix.Tag { return tag.ListID }

//NewListID returns a new ListIDField initialized with val
func NewListID(val string) ListIDField {
	return ListIDField{quickfix.FIXString(val)}
}

func (f ListIDField) Value() string { return f.String() }

//ListMethodField is a enum.ListMethod field
type ListMethodField struct{ quickfix.FIXString }

//Tag returns tag.ListMethod (1198)
func (f ListMethodField) Tag() quickfix.Tag { return tag.ListMethod }

func NewListMethod(val enum.ListMethod) ListMethodField {
	return ListMethodField{quickfix.FIXString(val)}
}

func (f ListMethodField) Value() enum.ListMethod { return enum.ListMethod(f.String()) }

//ListNameField is a STRING field
type ListNameField struct{ quickfix.FIXString }

//Tag returns tag.ListName (392)
func (f ListNameField) Tag() quickfix.Tag { return tag.ListName }

//NewListName returns a new ListNameField initialized with val
func NewListName(val string) ListNameField {
	return ListNameField{quickfix.FIXString(val)}
}

func (f ListNameField) Value() string { return f.String() }

//ListNoOrdsField is a INT field
type ListNoOrdsField struct{ quickfix.FIXInt }

//Tag returns tag.ListNoOrds (68)
func (f ListNoOrdsField) Tag() quickfix.Tag { return tag.ListNoOrds }

//NewListNoOrds returns a new ListNoOrdsField initialized with val
func NewListNoOrds(val int) ListNoOrdsField {
	return ListNoOrdsField{quickfix.FIXInt(val)}
}

func (f ListNoOrdsField) Value() int { return f.Int() }

//ListOrderStatusField is a enum.ListOrderStatus field
type ListOrderStatusField struct{ quickfix.FIXString }

//Tag returns tag.ListOrderStatus (431)
func (f ListOrderStatusField) Tag() quickfix.Tag { return tag.ListOrderStatus }

func NewListOrderStatus(val enum.ListOrderStatus) ListOrderStatusField {
	return ListOrderStatusField{quickfix.FIXString(val)}
}

func (f ListOrderStatusField) Value() enum.ListOrderStatus { return enum.ListOrderStatus(f.String()) }

//ListRejectReasonField is a enum.ListRejectReason field
type ListRejectReasonField struct{ quickfix.FIXString }

//Tag returns tag.ListRejectReason (1386)
func (f ListRejectReasonField) Tag() quickfix.Tag { return tag.ListRejectReason }

func NewListRejectReason(val enum.ListRejectReason) ListRejectReasonField {
	return ListRejectReasonField{quickfix.FIXString(val)}
}

func (f ListRejectReasonField) Value() enum.ListRejectReason { return enum.ListRejectReason(f.String()) }

//ListSeqNoField is a INT field
type ListSeqNoField struct{ quickfix.FIXInt }

//Tag returns tag.ListSeqNo (67)
func (f ListSeqNoField) Tag() quickfix.Tag { return tag.ListSeqNo }

//NewListSeqNo returns a new ListSeqNoField initialized with val
func NewListSeqNo(val int) ListSeqNoField {
	return ListSeqNoField{quickfix.FIXInt(val)}
}

func (f ListSeqNoField) Value() int { return f.Int() }

//ListStatusTextField is a STRING field
type ListStatusTextField struct{ quickfix.FIXString }

//Tag returns tag.ListStatusText (444)
func (f ListStatusTextField) Tag() quickfix.Tag { return tag.ListStatusText }

//NewListStatusText returns a new ListStatusTextField initialized with val
func NewListStatusText(val string) ListStatusTextField {
	return ListStatusTextField{quickfix.FIXString(val)}
}

func (f ListStatusTextField) Value() string { return f.String() }

//ListStatusTypeField is a enum.ListStatusType field
type ListStatusTypeField struct{ quickfix.FIXString }

//Tag returns tag.ListStatusType (429)
func (f ListStatusTypeField) Tag() quickfix.Tag { return tag.ListStatusType }

func NewListStatusType(val enum.ListStatusType) ListStatusTypeField {
	return ListStatusTypeField{quickfix.FIXString(val)}
}

func (f ListStatusTypeField) Value() enum.ListStatusType { return enum.ListStatusType(f.String()) }

//ListUpdateActionField is a CHAR field
type ListUpdateActionField struct{ quickfix.FIXString }

//Tag returns tag.ListUpdateAction (1324)
func (f ListUpdateActionField) Tag() quickfix.Tag { return tag.ListUpdateAction }

//NewListUpdateAction returns a new ListUpdateActionField initialized with val
func NewListUpdateAction(val string) ListUpdateActionField {
	return ListUpdateActionField{quickfix.FIXString(val)}
}

func (f ListUpdateActionField) Value() string { return f.String() }

//LocaleOfIssueField is a STRING field
type LocaleOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.LocaleOfIssue (472)
func (f LocaleOfIssueField) Tag() quickfix.Tag { return tag.LocaleOfIssue }

//NewLocaleOfIssue returns a new LocaleOfIssueField initialized with val
func NewLocaleOfIssue(val string) LocaleOfIssueField {
	return LocaleOfIssueField{quickfix.FIXString(val)}
}

func (f LocaleOfIssueField) Value() string { return f.String() }

//LocateReqdField is a BOOLEAN field
type LocateReqdField struct{ quickfix.FIXBoolean }

//Tag returns tag.LocateReqd (114)
func (f LocateReqdField) Tag() quickfix.Tag { return tag.LocateReqd }

//NewLocateReqd returns a new LocateReqdField initialized with val
func NewLocateReqd(val bool) LocateReqdField {
	return LocateReqdField{quickfix.FIXBoolean(val)}
}

func (f LocateReqdField) Value() bool { return f.Bool() }

//LocationIDField is a STRING field
type LocationIDField struct{ quickfix.FIXString }

//Tag returns tag.LocationID (283)
func (f LocationIDField) Tag() quickfix.Tag { return tag.LocationID }

//NewLocationID returns a new LocationIDField initialized with val
func NewLocationID(val string) LocationIDField {
	return LocationIDField{quickfix.FIXString(val)}
}

func (f LocationIDField) Value() string { return f.String() }

//LongQtyField is a QTY field
type LongQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.LongQty (704)
func (f LongQtyField) Tag() quickfix.Tag { return tag.LongQty }

//NewLongQty returns a new LongQtyField initialized with val and scale
func NewLongQty(val decimal.Decimal, scale int32) LongQtyField {
	return LongQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LongQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//LotTypeField is a enum.LotType field
type LotTypeField struct{ quickfix.FIXString }

//Tag returns tag.LotType (1093)
func (f LotTypeField) Tag() quickfix.Tag { return tag.LotType }

func NewLotType(val enum.LotType) LotTypeField {
	return LotTypeField{quickfix.FIXString(val)}
}

func (f LotTypeField) Value() enum.LotType { return enum.LotType(f.String()) }

//LowLimitPriceField is a PRICE field
type LowLimitPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.LowLimitPrice (1148)
func (f LowLimitPriceField) Tag() quickfix.Tag { return tag.LowLimitPrice }

//NewLowLimitPrice returns a new LowLimitPriceField initialized with val and scale
func NewLowLimitPrice(val decimal.Decimal, scale int32) LowLimitPriceField {
	return LowLimitPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LowLimitPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//LowPxField is a PRICE field
type LowPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.LowPx (333)
func (f LowPxField) Tag() quickfix.Tag { return tag.LowPx }

//NewLowPx returns a new LowPxField initialized with val and scale
func NewLowPx(val decimal.Decimal, scale int32) LowPxField {
	return LowPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LowPxField) Value() (val decimal.Decimal) { return f.Decimal }

//MDBookTypeField is a enum.MDBookType field
type MDBookTypeField struct{ quickfix.FIXString }

//Tag returns tag.MDBookType (1021)
func (f MDBookTypeField) Tag() quickfix.Tag { return tag.MDBookType }

func NewMDBookType(val enum.MDBookType) MDBookTypeField {
	return MDBookTypeField{quickfix.FIXString(val)}
}

func (f MDBookTypeField) Value() enum.MDBookType { return enum.MDBookType(f.String()) }

//MDEntryBuyerField is a STRING field
type MDEntryBuyerField struct{ quickfix.FIXString }

//Tag returns tag.MDEntryBuyer (288)
func (f MDEntryBuyerField) Tag() quickfix.Tag { return tag.MDEntryBuyer }

//NewMDEntryBuyer returns a new MDEntryBuyerField initialized with val
func NewMDEntryBuyer(val string) MDEntryBuyerField {
	return MDEntryBuyerField{quickfix.FIXString(val)}
}

func (f MDEntryBuyerField) Value() string { return f.String() }

//MDEntryDateField is a UTCDATEONLY field
type MDEntryDateField struct{ quickfix.FIXString }

//Tag returns tag.MDEntryDate (272)
func (f MDEntryDateField) Tag() quickfix.Tag { return tag.MDEntryDate }

//NewMDEntryDate returns a new MDEntryDateField initialized with val
func NewMDEntryDate(val string) MDEntryDateField {
	return MDEntryDateField{quickfix.FIXString(val)}
}

func (f MDEntryDateField) Value() string { return f.String() }

//MDEntryForwardPointsField is a PRICEOFFSET field
type MDEntryForwardPointsField struct{ quickfix.FIXDecimal }

//Tag returns tag.MDEntryForwardPoints (1027)
func (f MDEntryForwardPointsField) Tag() quickfix.Tag { return tag.MDEntryForwardPoints }

//NewMDEntryForwardPoints returns a new MDEntryForwardPointsField initialized with val and scale
func NewMDEntryForwardPoints(val decimal.Decimal, scale int32) MDEntryForwardPointsField {
	return MDEntryForwardPointsField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MDEntryForwardPointsField) Value() (val decimal.Decimal) { return f.Decimal }

//MDEntryIDField is a STRING field
type MDEntryIDField struct{ quickfix.FIXString }

//Tag returns tag.MDEntryID (278)
func (f MDEntryIDField) Tag() quickfix.Tag { return tag.MDEntryID }

//NewMDEntryID returns a new MDEntryIDField initialized with val
func NewMDEntryID(val string) MDEntryIDField {
	return MDEntryIDField{quickfix.FIXString(val)}
}

func (f MDEntryIDField) Value() string { return f.String() }

//MDEntryOriginatorField is a STRING field
type MDEntryOriginatorField struct{ quickfix.FIXString }

//Tag returns tag.MDEntryOriginator (282)
func (f MDEntryOriginatorField) Tag() quickfix.Tag { return tag.MDEntryOriginator }

//NewMDEntryOriginator returns a new MDEntryOriginatorField initialized with val
func NewMDEntryOriginator(val string) MDEntryOriginatorField {
	return MDEntryOriginatorField{quickfix.FIXString(val)}
}

func (f MDEntryOriginatorField) Value() string { return f.String() }

//MDEntryPositionNoField is a INT field
type MDEntryPositionNoField struct{ quickfix.FIXInt }

//Tag returns tag.MDEntryPositionNo (290)
func (f MDEntryPositionNoField) Tag() quickfix.Tag { return tag.MDEntryPositionNo }

//NewMDEntryPositionNo returns a new MDEntryPositionNoField initialized with val
func NewMDEntryPositionNo(val int) MDEntryPositionNoField {
	return MDEntryPositionNoField{quickfix.FIXInt(val)}
}

func (f MDEntryPositionNoField) Value() int { return f.Int() }

//MDEntryPxField is a PRICE field
type MDEntryPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.MDEntryPx (270)
func (f MDEntryPxField) Tag() quickfix.Tag { return tag.MDEntryPx }

//NewMDEntryPx returns a new MDEntryPxField initialized with val and scale
func NewMDEntryPx(val decimal.Decimal, scale int32) MDEntryPxField {
	return MDEntryPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MDEntryPxField) Value() (val decimal.Decimal) { return f.Decimal }

//MDEntryRefIDField is a STRING field
type MDEntryRefIDField struct{ quickfix.FIXString }

//Tag returns tag.MDEntryRefID (280)
func (f MDEntryRefIDField) Tag() quickfix.Tag { return tag.MDEntryRefID }

//NewMDEntryRefID returns a new MDEntryRefIDField initialized with val
func NewMDEntryRefID(val string) MDEntryRefIDField {
	return MDEntryRefIDField{quickfix.FIXString(val)}
}

func (f MDEntryRefIDField) Value() string { return f.String() }

//MDEntrySellerField is a STRING field
type MDEntrySellerField struct{ quickfix.FIXString }

//Tag returns tag.MDEntrySeller (289)
func (f MDEntrySellerField) Tag() quickfix.Tag { return tag.MDEntrySeller }

//NewMDEntrySeller returns a new MDEntrySellerField initialized with val
func NewMDEntrySeller(val string) MDEntrySellerField {
	return MDEntrySellerField{quickfix.FIXString(val)}
}

func (f MDEntrySellerField) Value() string { return f.String() }

//MDEntrySizeField is a QTY field
type MDEntrySizeField struct{ quickfix.FIXDecimal }

//Tag returns tag.MDEntrySize (271)
func (f MDEntrySizeField) Tag() quickfix.Tag { return tag.MDEntrySize }

//NewMDEntrySize returns a new MDEntrySizeField initialized with val and scale
func NewMDEntrySize(val decimal.Decimal, scale int32) MDEntrySizeField {
	return MDEntrySizeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MDEntrySizeField) Value() (val decimal.Decimal) { return f.Decimal }

//MDEntrySpotRateField is a FLOAT field
type MDEntrySpotRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.MDEntrySpotRate (1026)
func (f MDEntrySpotRateField) Tag() quickfix.Tag { return tag.MDEntrySpotRate }

//NewMDEntrySpotRate returns a new MDEntrySpotRateField initialized with val and scale
func NewMDEntrySpotRate(val decimal.Decimal, scale int32) MDEntrySpotRateField {
	return MDEntrySpotRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MDEntrySpotRateField) Value() (val decimal.Decimal) { return f.Decimal }

//MDEntryTimeField is a UTCTIMEONLY field
type MDEntryTimeField struct{ quickfix.FIXString }

//Tag returns tag.MDEntryTime (273)
func (f MDEntryTimeField) Tag() quickfix.Tag { return tag.MDEntryTime }

//NewMDEntryTime returns a new MDEntryTimeField initialized with val
func NewMDEntryTime(val string) MDEntryTimeField {
	return MDEntryTimeField{quickfix.FIXString(val)}
}

func (f MDEntryTimeField) Value() string { return f.String() }

//MDEntryTypeField is a enum.MDEntryType field
type MDEntryTypeField struct{ quickfix.FIXString }

//Tag returns tag.MDEntryType (269)
func (f MDEntryTypeField) Tag() quickfix.Tag { return tag.MDEntryType }

func NewMDEntryType(val enum.MDEntryType) MDEntryTypeField {
	return MDEntryTypeField{quickfix.FIXString(val)}
}

func (f MDEntryTypeField) Value() enum.MDEntryType { return enum.MDEntryType(f.String()) }

//MDFeedTypeField is a STRING field
type MDFeedTypeField struct{ quickfix.FIXString }

//Tag returns tag.MDFeedType (1022)
func (f MDFeedTypeField) Tag() quickfix.Tag { return tag.MDFeedType }

//NewMDFeedType returns a new MDFeedTypeField initialized with val
func NewMDFeedType(val string) MDFeedTypeField {
	return MDFeedTypeField{quickfix.FIXString(val)}
}

func (f MDFeedTypeField) Value() string { return f.String() }

//MDImplicitDeleteField is a BOOLEAN field
type MDImplicitDeleteField struct{ quickfix.FIXBoolean }

//Tag returns tag.MDImplicitDelete (547)
func (f MDImplicitDeleteField) Tag() quickfix.Tag { return tag.MDImplicitDelete }

//NewMDImplicitDelete returns a new MDImplicitDeleteField initialized with val
func NewMDImplicitDelete(val bool) MDImplicitDeleteField {
	return MDImplicitDeleteField{quickfix.FIXBoolean(val)}
}

func (f MDImplicitDeleteField) Value() bool { return f.Bool() }

//MDMktField is a EXCHANGE field
type MDMktField struct{ quickfix.FIXString }

//Tag returns tag.MDMkt (275)
func (f MDMktField) Tag() quickfix.Tag { return tag.MDMkt }

//NewMDMkt returns a new MDMktField initialized with val
func NewMDMkt(val string) MDMktField {
	return MDMktField{quickfix.FIXString(val)}
}

func (f MDMktField) Value() string { return f.String() }

//MDOriginTypeField is a enum.MDOriginType field
type MDOriginTypeField struct{ quickfix.FIXString }

//Tag returns tag.MDOriginType (1024)
func (f MDOriginTypeField) Tag() quickfix.Tag { return tag.MDOriginType }

func NewMDOriginType(val enum.MDOriginType) MDOriginTypeField {
	return MDOriginTypeField{quickfix.FIXString(val)}
}

func (f MDOriginTypeField) Value() enum.MDOriginType { return enum.MDOriginType(f.String()) }

//MDPriceLevelField is a INT field
type MDPriceLevelField struct{ quickfix.FIXInt }

//Tag returns tag.MDPriceLevel (1023)
func (f MDPriceLevelField) Tag() quickfix.Tag { return tag.MDPriceLevel }

//NewMDPriceLevel returns a new MDPriceLevelField initialized with val
func NewMDPriceLevel(val int) MDPriceLevelField {
	return MDPriceLevelField{quickfix.FIXInt(val)}
}

func (f MDPriceLevelField) Value() int { return f.Int() }

//MDQuoteTypeField is a enum.MDQuoteType field
type MDQuoteTypeField struct{ quickfix.FIXString }

//Tag returns tag.MDQuoteType (1070)
func (f MDQuoteTypeField) Tag() quickfix.Tag { return tag.MDQuoteType }

func NewMDQuoteType(val enum.MDQuoteType) MDQuoteTypeField {
	return MDQuoteTypeField{quickfix.FIXString(val)}
}

func (f MDQuoteTypeField) Value() enum.MDQuoteType { return enum.MDQuoteType(f.String()) }

//MDReportIDField is a INT field
type MDReportIDField struct{ quickfix.FIXInt }

//Tag returns tag.MDReportID (963)
func (f MDReportIDField) Tag() quickfix.Tag { return tag.MDReportID }

//NewMDReportID returns a new MDReportIDField initialized with val
func NewMDReportID(val int) MDReportIDField {
	return MDReportIDField{quickfix.FIXInt(val)}
}

func (f MDReportIDField) Value() int { return f.Int() }

//MDReqIDField is a STRING field
type MDReqIDField struct{ quickfix.FIXString }

//Tag returns tag.MDReqID (262)
func (f MDReqIDField) Tag() quickfix.Tag { return tag.MDReqID }

//NewMDReqID returns a new MDReqIDField initialized with val
func NewMDReqID(val string) MDReqIDField {
	return MDReqIDField{quickfix.FIXString(val)}
}

func (f MDReqIDField) Value() string { return f.String() }

//MDReqRejReasonField is a enum.MDReqRejReason field
type MDReqRejReasonField struct{ quickfix.FIXString }

//Tag returns tag.MDReqRejReason (281)
func (f MDReqRejReasonField) Tag() quickfix.Tag { return tag.MDReqRejReason }

func NewMDReqRejReason(val enum.MDReqRejReason) MDReqRejReasonField {
	return MDReqRejReasonField{quickfix.FIXString(val)}
}

func (f MDReqRejReasonField) Value() enum.MDReqRejReason { return enum.MDReqRejReason(f.String()) }

//MDSecSizeField is a QTY field
type MDSecSizeField struct{ quickfix.FIXDecimal }

//Tag returns tag.MDSecSize (1179)
func (f MDSecSizeField) Tag() quickfix.Tag { return tag.MDSecSize }

//NewMDSecSize returns a new MDSecSizeField initialized with val and scale
func NewMDSecSize(val decimal.Decimal, scale int32) MDSecSizeField {
	return MDSecSizeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MDSecSizeField) Value() (val decimal.Decimal) { return f.Decimal }

//MDSecSizeTypeField is a enum.MDSecSizeType field
type MDSecSizeTypeField struct{ quickfix.FIXString }

//Tag returns tag.MDSecSizeType (1178)
func (f MDSecSizeTypeField) Tag() quickfix.Tag { return tag.MDSecSizeType }

func NewMDSecSizeType(val enum.MDSecSizeType) MDSecSizeTypeField {
	return MDSecSizeTypeField{quickfix.FIXString(val)}
}

func (f MDSecSizeTypeField) Value() enum.MDSecSizeType { return enum.MDSecSizeType(f.String()) }

//MDStreamIDField is a STRING field
type MDStreamIDField struct{ quickfix.FIXString }

//Tag returns tag.MDStreamID (1500)
func (f MDStreamIDField) Tag() quickfix.Tag { return tag.MDStreamID }

//NewMDStreamID returns a new MDStreamIDField initialized with val
func NewMDStreamID(val string) MDStreamIDField {
	return MDStreamIDField{quickfix.FIXString(val)}
}

func (f MDStreamIDField) Value() string { return f.String() }

//MDSubBookTypeField is a INT field
type MDSubBookTypeField struct{ quickfix.FIXInt }

//Tag returns tag.MDSubBookType (1173)
func (f MDSubBookTypeField) Tag() quickfix.Tag { return tag.MDSubBookType }

//NewMDSubBookType returns a new MDSubBookTypeField initialized with val
func NewMDSubBookType(val int) MDSubBookTypeField {
	return MDSubBookTypeField{quickfix.FIXInt(val)}
}

func (f MDSubBookTypeField) Value() int { return f.Int() }

//MDUpdateActionField is a enum.MDUpdateAction field
type MDUpdateActionField struct{ quickfix.FIXString }

//Tag returns tag.MDUpdateAction (279)
func (f MDUpdateActionField) Tag() quickfix.Tag { return tag.MDUpdateAction }

func NewMDUpdateAction(val enum.MDUpdateAction) MDUpdateActionField {
	return MDUpdateActionField{quickfix.FIXString(val)}
}

func (f MDUpdateActionField) Value() enum.MDUpdateAction { return enum.MDUpdateAction(f.String()) }

//MDUpdateTypeField is a enum.MDUpdateType field
type MDUpdateTypeField struct{ quickfix.FIXString }

//Tag returns tag.MDUpdateType (265)
func (f MDUpdateTypeField) Tag() quickfix.Tag { return tag.MDUpdateType }

func NewMDUpdateType(val enum.MDUpdateType) MDUpdateTypeField {
	return MDUpdateTypeField{quickfix.FIXString(val)}
}

func (f MDUpdateTypeField) Value() enum.MDUpdateType { return enum.MDUpdateType(f.String()) }

//MailingDtlsField is a STRING field
type MailingDtlsField struct{ quickfix.FIXString }

//Tag returns tag.MailingDtls (474)
func (f MailingDtlsField) Tag() quickfix.Tag { return tag.MailingDtls }

//NewMailingDtls returns a new MailingDtlsField initialized with val
func NewMailingDtls(val string) MailingDtlsField {
	return MailingDtlsField{quickfix.FIXString(val)}
}

func (f MailingDtlsField) Value() string { return f.String() }

//MailingInstField is a STRING field
type MailingInstField struct{ quickfix.FIXString }

//Tag returns tag.MailingInst (482)
func (f MailingInstField) Tag() quickfix.Tag { return tag.MailingInst }

//NewMailingInst returns a new MailingInstField initialized with val
func NewMailingInst(val string) MailingInstField {
	return MailingInstField{quickfix.FIXString(val)}
}

func (f MailingInstField) Value() string { return f.String() }

//ManualOrderIndicatorField is a BOOLEAN field
type ManualOrderIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.ManualOrderIndicator (1028)
func (f ManualOrderIndicatorField) Tag() quickfix.Tag { return tag.ManualOrderIndicator }

//NewManualOrderIndicator returns a new ManualOrderIndicatorField initialized with val
func NewManualOrderIndicator(val bool) ManualOrderIndicatorField {
	return ManualOrderIndicatorField{quickfix.FIXBoolean(val)}
}

func (f ManualOrderIndicatorField) Value() bool { return f.Bool() }

//MarginExcessField is a AMT field
type MarginExcessField struct{ quickfix.FIXDecimal }

//Tag returns tag.MarginExcess (899)
func (f MarginExcessField) Tag() quickfix.Tag { return tag.MarginExcess }

//NewMarginExcess returns a new MarginExcessField initialized with val and scale
func NewMarginExcess(val decimal.Decimal, scale int32) MarginExcessField {
	return MarginExcessField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MarginExcessField) Value() (val decimal.Decimal) { return f.Decimal }

//MarginRatioField is a PERCENTAGE field
type MarginRatioField struct{ quickfix.FIXDecimal }

//Tag returns tag.MarginRatio (898)
func (f MarginRatioField) Tag() quickfix.Tag { return tag.MarginRatio }

//NewMarginRatio returns a new MarginRatioField initialized with val and scale
func NewMarginRatio(val decimal.Decimal, scale int32) MarginRatioField {
	return MarginRatioField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MarginRatioField) Value() (val decimal.Decimal) { return f.Decimal }

//MarketDepthField is a INT field
type MarketDepthField struct{ quickfix.FIXInt }

//Tag returns tag.MarketDepth (264)
func (f MarketDepthField) Tag() quickfix.Tag { return tag.MarketDepth }

//NewMarketDepth returns a new MarketDepthField initialized with val
func NewMarketDepth(val int) MarketDepthField {
	return MarketDepthField{quickfix.FIXInt(val)}
}

func (f MarketDepthField) Value() int { return f.Int() }

//MarketIDField is a EXCHANGE field
type MarketIDField struct{ quickfix.FIXString }

//Tag returns tag.MarketID (1301)
func (f MarketIDField) Tag() quickfix.Tag { return tag.MarketID }

//NewMarketID returns a new MarketIDField initialized with val
func NewMarketID(val string) MarketIDField {
	return MarketIDField{quickfix.FIXString(val)}
}

func (f MarketIDField) Value() string { return f.String() }

//MarketReportIDField is a STRING field
type MarketReportIDField struct{ quickfix.FIXString }

//Tag returns tag.MarketReportID (1394)
func (f MarketReportIDField) Tag() quickfix.Tag { return tag.MarketReportID }

//NewMarketReportID returns a new MarketReportIDField initialized with val
func NewMarketReportID(val string) MarketReportIDField {
	return MarketReportIDField{quickfix.FIXString(val)}
}

func (f MarketReportIDField) Value() string { return f.String() }

//MarketReqIDField is a STRING field
type MarketReqIDField struct{ quickfix.FIXString }

//Tag returns tag.MarketReqID (1393)
func (f MarketReqIDField) Tag() quickfix.Tag { return tag.MarketReqID }

//NewMarketReqID returns a new MarketReqIDField initialized with val
func NewMarketReqID(val string) MarketReqIDField {
	return MarketReqIDField{quickfix.FIXString(val)}
}

func (f MarketReqIDField) Value() string { return f.String() }

//MarketSegmentDescField is a STRING field
type MarketSegmentDescField struct{ quickfix.FIXString }

//Tag returns tag.MarketSegmentDesc (1396)
func (f MarketSegmentDescField) Tag() quickfix.Tag { return tag.MarketSegmentDesc }

//NewMarketSegmentDesc returns a new MarketSegmentDescField initialized with val
func NewMarketSegmentDesc(val string) MarketSegmentDescField {
	return MarketSegmentDescField{quickfix.FIXString(val)}
}

func (f MarketSegmentDescField) Value() string { return f.String() }

//MarketSegmentIDField is a STRING field
type MarketSegmentIDField struct{ quickfix.FIXString }

//Tag returns tag.MarketSegmentID (1300)
func (f MarketSegmentIDField) Tag() quickfix.Tag { return tag.MarketSegmentID }

//NewMarketSegmentID returns a new MarketSegmentIDField initialized with val
func NewMarketSegmentID(val string) MarketSegmentIDField {
	return MarketSegmentIDField{quickfix.FIXString(val)}
}

func (f MarketSegmentIDField) Value() string { return f.String() }

//MarketUpdateActionField is a enum.MarketUpdateAction field
type MarketUpdateActionField struct{ quickfix.FIXString }

//Tag returns tag.MarketUpdateAction (1395)
func (f MarketUpdateActionField) Tag() quickfix.Tag { return tag.MarketUpdateAction }

func NewMarketUpdateAction(val enum.MarketUpdateAction) MarketUpdateActionField {
	return MarketUpdateActionField{quickfix.FIXString(val)}
}

func (f MarketUpdateActionField) Value() enum.MarketUpdateAction {
	return enum.MarketUpdateAction(f.String())
}

//MassActionRejectReasonField is a enum.MassActionRejectReason field
type MassActionRejectReasonField struct{ quickfix.FIXString }

//Tag returns tag.MassActionRejectReason (1376)
func (f MassActionRejectReasonField) Tag() quickfix.Tag { return tag.MassActionRejectReason }

func NewMassActionRejectReason(val enum.MassActionRejectReason) MassActionRejectReasonField {
	return MassActionRejectReasonField{quickfix.FIXString(val)}
}

func (f MassActionRejectReasonField) Value() enum.MassActionRejectReason {
	return enum.MassActionRejectReason(f.String())
}

//MassActionReportIDField is a STRING field
type MassActionReportIDField struct{ quickfix.FIXString }

//Tag returns tag.MassActionReportID (1369)
func (f MassActionReportIDField) Tag() quickfix.Tag { return tag.MassActionReportID }

//NewMassActionReportID returns a new MassActionReportIDField initialized with val
func NewMassActionReportID(val string) MassActionReportIDField {
	return MassActionReportIDField{quickfix.FIXString(val)}
}

func (f MassActionReportIDField) Value() string { return f.String() }

//MassActionResponseField is a enum.MassActionResponse field
type MassActionResponseField struct{ quickfix.FIXString }

//Tag returns tag.MassActionResponse (1375)
func (f MassActionResponseField) Tag() quickfix.Tag { return tag.MassActionResponse }

func NewMassActionResponse(val enum.MassActionResponse) MassActionResponseField {
	return MassActionResponseField{quickfix.FIXString(val)}
}

func (f MassActionResponseField) Value() enum.MassActionResponse {
	return enum.MassActionResponse(f.String())
}

//MassActionScopeField is a enum.MassActionScope field
type MassActionScopeField struct{ quickfix.FIXString }

//Tag returns tag.MassActionScope (1374)
func (f MassActionScopeField) Tag() quickfix.Tag { return tag.MassActionScope }

func NewMassActionScope(val enum.MassActionScope) MassActionScopeField {
	return MassActionScopeField{quickfix.FIXString(val)}
}

func (f MassActionScopeField) Value() enum.MassActionScope { return enum.MassActionScope(f.String()) }

//MassActionTypeField is a enum.MassActionType field
type MassActionTypeField struct{ quickfix.FIXString }

//Tag returns tag.MassActionType (1373)
func (f MassActionTypeField) Tag() quickfix.Tag { return tag.MassActionType }

func NewMassActionType(val enum.MassActionType) MassActionTypeField {
	return MassActionTypeField{quickfix.FIXString(val)}
}

func (f MassActionTypeField) Value() enum.MassActionType { return enum.MassActionType(f.String()) }

//MassCancelRejectReasonField is a enum.MassCancelRejectReason field
type MassCancelRejectReasonField struct{ quickfix.FIXString }

//Tag returns tag.MassCancelRejectReason (532)
func (f MassCancelRejectReasonField) Tag() quickfix.Tag { return tag.MassCancelRejectReason }

func NewMassCancelRejectReason(val enum.MassCancelRejectReason) MassCancelRejectReasonField {
	return MassCancelRejectReasonField{quickfix.FIXString(val)}
}

func (f MassCancelRejectReasonField) Value() enum.MassCancelRejectReason {
	return enum.MassCancelRejectReason(f.String())
}

//MassCancelRequestTypeField is a enum.MassCancelRequestType field
type MassCancelRequestTypeField struct{ quickfix.FIXString }

//Tag returns tag.MassCancelRequestType (530)
func (f MassCancelRequestTypeField) Tag() quickfix.Tag { return tag.MassCancelRequestType }

func NewMassCancelRequestType(val enum.MassCancelRequestType) MassCancelRequestTypeField {
	return MassCancelRequestTypeField{quickfix.FIXString(val)}
}

func (f MassCancelRequestTypeField) Value() enum.MassCancelRequestType {
	return enum.MassCancelRequestType(f.String())
}

//MassCancelResponseField is a enum.MassCancelResponse field
type MassCancelResponseField struct{ quickfix.FIXString }

//Tag returns tag.MassCancelResponse (531)
func (f MassCancelResponseField) Tag() quickfix.Tag { return tag.MassCancelResponse }

func NewMassCancelResponse(val enum.MassCancelResponse) MassCancelResponseField {
	return MassCancelResponseField{quickfix.FIXString(val)}
}

func (f MassCancelResponseField) Value() enum.MassCancelResponse {
	return enum.MassCancelResponse(f.String())
}

//MassStatusReqIDField is a STRING field
type MassStatusReqIDField struct{ quickfix.FIXString }

//Tag returns tag.MassStatusReqID (584)
func (f MassStatusReqIDField) Tag() quickfix.Tag { return tag.MassStatusReqID }

//NewMassStatusReqID returns a new MassStatusReqIDField initialized with val
func NewMassStatusReqID(val string) MassStatusReqIDField {
	return MassStatusReqIDField{quickfix.FIXString(val)}
}

func (f MassStatusReqIDField) Value() string { return f.String() }

//MassStatusReqTypeField is a enum.MassStatusReqType field
type MassStatusReqTypeField struct{ quickfix.FIXString }

//Tag returns tag.MassStatusReqType (585)
func (f MassStatusReqTypeField) Tag() quickfix.Tag { return tag.MassStatusReqType }

func NewMassStatusReqType(val enum.MassStatusReqType) MassStatusReqTypeField {
	return MassStatusReqTypeField{quickfix.FIXString(val)}
}

func (f MassStatusReqTypeField) Value() enum.MassStatusReqType {
	return enum.MassStatusReqType(f.String())
}

//MatchAlgorithmField is a STRING field
type MatchAlgorithmField struct{ quickfix.FIXString }

//Tag returns tag.MatchAlgorithm (1142)
func (f MatchAlgorithmField) Tag() quickfix.Tag { return tag.MatchAlgorithm }

//NewMatchAlgorithm returns a new MatchAlgorithmField initialized with val
func NewMatchAlgorithm(val string) MatchAlgorithmField {
	return MatchAlgorithmField{quickfix.FIXString(val)}
}

func (f MatchAlgorithmField) Value() string { return f.String() }

//MatchIncrementField is a QTY field
type MatchIncrementField struct{ quickfix.FIXDecimal }

//Tag returns tag.MatchIncrement (1089)
func (f MatchIncrementField) Tag() quickfix.Tag { return tag.MatchIncrement }

//NewMatchIncrement returns a new MatchIncrementField initialized with val and scale
func NewMatchIncrement(val decimal.Decimal, scale int32) MatchIncrementField {
	return MatchIncrementField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MatchIncrementField) Value() (val decimal.Decimal) { return f.Decimal }

//MatchStatusField is a enum.MatchStatus field
type MatchStatusField struct{ quickfix.FIXString }

//Tag returns tag.MatchStatus (573)
func (f MatchStatusField) Tag() quickfix.Tag { return tag.MatchStatus }

func NewMatchStatus(val enum.MatchStatus) MatchStatusField {
	return MatchStatusField{quickfix.FIXString(val)}
}

func (f MatchStatusField) Value() enum.MatchStatus { return enum.MatchStatus(f.String()) }

//MatchTypeField is a enum.MatchType field
type MatchTypeField struct{ quickfix.FIXString }

//Tag returns tag.MatchType (574)
func (f MatchTypeField) Tag() quickfix.Tag { return tag.MatchType }

func NewMatchType(val enum.MatchType) MatchTypeField {
	return MatchTypeField{quickfix.FIXString(val)}
}

func (f MatchTypeField) Value() enum.MatchType { return enum.MatchType(f.String()) }

//MaturityDateField is a LOCALMKTDATE field
type MaturityDateField struct{ quickfix.FIXString }

//Tag returns tag.MaturityDate (541)
func (f MaturityDateField) Tag() quickfix.Tag { return tag.MaturityDate }

//NewMaturityDate returns a new MaturityDateField initialized with val
func NewMaturityDate(val string) MaturityDateField {
	return MaturityDateField{quickfix.FIXString(val)}
}

func (f MaturityDateField) Value() string { return f.String() }

//MaturityDayField is a DAYOFMONTH field
type MaturityDayField struct{ quickfix.FIXInt }

//Tag returns tag.MaturityDay (205)
func (f MaturityDayField) Tag() quickfix.Tag { return tag.MaturityDay }

//NewMaturityDay returns a new MaturityDayField initialized with val
func NewMaturityDay(val int) MaturityDayField {
	return MaturityDayField{quickfix.FIXInt(val)}
}

func (f MaturityDayField) Value() int { return f.Int() }

//MaturityMonthYearField is a MONTHYEAR field
type MaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.MaturityMonthYear (200)
func (f MaturityMonthYearField) Tag() quickfix.Tag { return tag.MaturityMonthYear }

//NewMaturityMonthYear returns a new MaturityMonthYearField initialized with val
func NewMaturityMonthYear(val string) MaturityMonthYearField {
	return MaturityMonthYearField{quickfix.FIXString(val)}
}

func (f MaturityMonthYearField) Value() string { return f.String() }

//MaturityMonthYearFormatField is a enum.MaturityMonthYearFormat field
type MaturityMonthYearFormatField struct{ quickfix.FIXString }

//Tag returns tag.MaturityMonthYearFormat (1303)
func (f MaturityMonthYearFormatField) Tag() quickfix.Tag { return tag.MaturityMonthYearFormat }

func NewMaturityMonthYearFormat(val enum.MaturityMonthYearFormat) MaturityMonthYearFormatField {
	return MaturityMonthYearFormatField{quickfix.FIXString(val)}
}

func (f MaturityMonthYearFormatField) Value() enum.MaturityMonthYearFormat {
	return enum.MaturityMonthYearFormat(f.String())
}

//MaturityMonthYearIncrementField is a INT field
type MaturityMonthYearIncrementField struct{ quickfix.FIXInt }

//Tag returns tag.MaturityMonthYearIncrement (1229)
func (f MaturityMonthYearIncrementField) Tag() quickfix.Tag { return tag.MaturityMonthYearIncrement }

//NewMaturityMonthYearIncrement returns a new MaturityMonthYearIncrementField initialized with val
func NewMaturityMonthYearIncrement(val int) MaturityMonthYearIncrementField {
	return MaturityMonthYearIncrementField{quickfix.FIXInt(val)}
}

func (f MaturityMonthYearIncrementField) Value() int { return f.Int() }

//MaturityMonthYearIncrementUnitsField is a enum.MaturityMonthYearIncrementUnits field
type MaturityMonthYearIncrementUnitsField struct{ quickfix.FIXString }

//Tag returns tag.MaturityMonthYearIncrementUnits (1302)
func (f MaturityMonthYearIncrementUnitsField) Tag() quickfix.Tag {
	return tag.MaturityMonthYearIncrementUnits
}

func NewMaturityMonthYearIncrementUnits(val enum.MaturityMonthYearIncrementUnits) MaturityMonthYearIncrementUnitsField {
	return MaturityMonthYearIncrementUnitsField{quickfix.FIXString(val)}
}

func (f MaturityMonthYearIncrementUnitsField) Value() enum.MaturityMonthYearIncrementUnits {
	return enum.MaturityMonthYearIncrementUnits(f.String())
}

//MaturityNetMoneyField is a AMT field
type MaturityNetMoneyField struct{ quickfix.FIXDecimal }

//Tag returns tag.MaturityNetMoney (890)
func (f MaturityNetMoneyField) Tag() quickfix.Tag { return tag.MaturityNetMoney }

//NewMaturityNetMoney returns a new MaturityNetMoneyField initialized with val and scale
func NewMaturityNetMoney(val decimal.Decimal, scale int32) MaturityNetMoneyField {
	return MaturityNetMoneyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MaturityNetMoneyField) Value() (val decimal.Decimal) { return f.Decimal }

//MaturityRuleIDField is a STRING field
type MaturityRuleIDField struct{ quickfix.FIXString }

//Tag returns tag.MaturityRuleID (1222)
func (f MaturityRuleIDField) Tag() quickfix.Tag { return tag.MaturityRuleID }

//NewMaturityRuleID returns a new MaturityRuleIDField initialized with val
func NewMaturityRuleID(val string) MaturityRuleIDField {
	return MaturityRuleIDField{quickfix.FIXString(val)}
}

func (f MaturityRuleIDField) Value() string { return f.String() }

//MaturityTimeField is a TZTIMEONLY field
type MaturityTimeField struct{ quickfix.FIXString }

//Tag returns tag.MaturityTime (1079)
func (f MaturityTimeField) Tag() quickfix.Tag { return tag.MaturityTime }

//NewMaturityTime returns a new MaturityTimeField initialized with val
func NewMaturityTime(val string) MaturityTimeField {
	return MaturityTimeField{quickfix.FIXString(val)}
}

func (f MaturityTimeField) Value() string { return f.String() }

//MaxFloorField is a QTY field
type MaxFloorField struct{ quickfix.FIXDecimal }

//Tag returns tag.MaxFloor (111)
func (f MaxFloorField) Tag() quickfix.Tag { return tag.MaxFloor }

//NewMaxFloor returns a new MaxFloorField initialized with val and scale
func NewMaxFloor(val decimal.Decimal, scale int32) MaxFloorField {
	return MaxFloorField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MaxFloorField) Value() (val decimal.Decimal) { return f.Decimal }

//MaxMessageSizeField is a LENGTH field
type MaxMessageSizeField struct{ quickfix.FIXInt }

//Tag returns tag.MaxMessageSize (383)
func (f MaxMessageSizeField) Tag() quickfix.Tag { return tag.MaxMessageSize }

//NewMaxMessageSize returns a new MaxMessageSizeField initialized with val
func NewMaxMessageSize(val int) MaxMessageSizeField {
	return MaxMessageSizeField{quickfix.FIXInt(val)}
}

func (f MaxMessageSizeField) Value() int { return f.Int() }

//MaxPriceLevelsField is a INT field
type MaxPriceLevelsField struct{ quickfix.FIXInt }

//Tag returns tag.MaxPriceLevels (1090)
func (f MaxPriceLevelsField) Tag() quickfix.Tag { return tag.MaxPriceLevels }

//NewMaxPriceLevels returns a new MaxPriceLevelsField initialized with val
func NewMaxPriceLevels(val int) MaxPriceLevelsField {
	return MaxPriceLevelsField{quickfix.FIXInt(val)}
}

func (f MaxPriceLevelsField) Value() int { return f.Int() }

//MaxPriceVariationField is a FLOAT field
type MaxPriceVariationField struct{ quickfix.FIXDecimal }

//Tag returns tag.MaxPriceVariation (1143)
func (f MaxPriceVariationField) Tag() quickfix.Tag { return tag.MaxPriceVariation }

//NewMaxPriceVariation returns a new MaxPriceVariationField initialized with val and scale
func NewMaxPriceVariation(val decimal.Decimal, scale int32) MaxPriceVariationField {
	return MaxPriceVariationField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MaxPriceVariationField) Value() (val decimal.Decimal) { return f.Decimal }

//MaxShowField is a QTY field
type MaxShowField struct{ quickfix.FIXDecimal }

//Tag returns tag.MaxShow (210)
func (f MaxShowField) Tag() quickfix.Tag { return tag.MaxShow }

//NewMaxShow returns a new MaxShowField initialized with val and scale
func NewMaxShow(val decimal.Decimal, scale int32) MaxShowField {
	return MaxShowField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MaxShowField) Value() (val decimal.Decimal) { return f.Decimal }

//MaxTradeVolField is a QTY field
type MaxTradeVolField struct{ quickfix.FIXDecimal }

//Tag returns tag.MaxTradeVol (1140)
func (f MaxTradeVolField) Tag() quickfix.Tag { return tag.MaxTradeVol }

//NewMaxTradeVol returns a new MaxTradeVolField initialized with val and scale
func NewMaxTradeVol(val decimal.Decimal, scale int32) MaxTradeVolField {
	return MaxTradeVolField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MaxTradeVolField) Value() (val decimal.Decimal) { return f.Decimal }

//MessageEncodingField is a enum.MessageEncoding field
type MessageEncodingField struct{ quickfix.FIXString }

//Tag returns tag.MessageEncoding (347)
func (f MessageEncodingField) Tag() quickfix.Tag { return tag.MessageEncoding }

func NewMessageEncoding(val enum.MessageEncoding) MessageEncodingField {
	return MessageEncodingField{quickfix.FIXString(val)}
}

func (f MessageEncodingField) Value() enum.MessageEncoding { return enum.MessageEncoding(f.String()) }

//MessageEventSourceField is a STRING field
type MessageEventSourceField struct{ quickfix.FIXString }

//Tag returns tag.MessageEventSource (1011)
func (f MessageEventSourceField) Tag() quickfix.Tag { return tag.MessageEventSource }

//NewMessageEventSource returns a new MessageEventSourceField initialized with val
func NewMessageEventSource(val string) MessageEventSourceField {
	return MessageEventSourceField{quickfix.FIXString(val)}
}

func (f MessageEventSourceField) Value() string { return f.String() }

//MidPxField is a PRICE field
type MidPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.MidPx (631)
func (f MidPxField) Tag() quickfix.Tag { return tag.MidPx }

//NewMidPx returns a new MidPxField initialized with val and scale
func NewMidPx(val decimal.Decimal, scale int32) MidPxField {
	return MidPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MidPxField) Value() (val decimal.Decimal) { return f.Decimal }

//MidYieldField is a PERCENTAGE field
type MidYieldField struct{ quickfix.FIXDecimal }

//Tag returns tag.MidYield (633)
func (f MidYieldField) Tag() quickfix.Tag { return tag.MidYield }

//NewMidYield returns a new MidYieldField initialized with val and scale
func NewMidYield(val decimal.Decimal, scale int32) MidYieldField {
	return MidYieldField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MidYieldField) Value() (val decimal.Decimal) { return f.Decimal }

//MinBidSizeField is a QTY field
type MinBidSizeField struct{ quickfix.FIXDecimal }

//Tag returns tag.MinBidSize (647)
func (f MinBidSizeField) Tag() quickfix.Tag { return tag.MinBidSize }

//NewMinBidSize returns a new MinBidSizeField initialized with val and scale
func NewMinBidSize(val decimal.Decimal, scale int32) MinBidSizeField {
	return MinBidSizeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MinBidSizeField) Value() (val decimal.Decimal) { return f.Decimal }

//MinLotSizeField is a QTY field
type MinLotSizeField struct{ quickfix.FIXDecimal }

//Tag returns tag.MinLotSize (1231)
func (f MinLotSizeField) Tag() quickfix.Tag { return tag.MinLotSize }

//NewMinLotSize returns a new MinLotSizeField initialized with val and scale
func NewMinLotSize(val decimal.Decimal, scale int32) MinLotSizeField {
	return MinLotSizeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MinLotSizeField) Value() (val decimal.Decimal) { return f.Decimal }

//MinOfferSizeField is a QTY field
type MinOfferSizeField struct{ quickfix.FIXDecimal }

//Tag returns tag.MinOfferSize (648)
func (f MinOfferSizeField) Tag() quickfix.Tag { return tag.MinOfferSize }

//NewMinOfferSize returns a new MinOfferSizeField initialized with val and scale
func NewMinOfferSize(val decimal.Decimal, scale int32) MinOfferSizeField {
	return MinOfferSizeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MinOfferSizeField) Value() (val decimal.Decimal) { return f.Decimal }

//MinPriceIncrementField is a FLOAT field
type MinPriceIncrementField struct{ quickfix.FIXDecimal }

//Tag returns tag.MinPriceIncrement (969)
func (f MinPriceIncrementField) Tag() quickfix.Tag { return tag.MinPriceIncrement }

//NewMinPriceIncrement returns a new MinPriceIncrementField initialized with val and scale
func NewMinPriceIncrement(val decimal.Decimal, scale int32) MinPriceIncrementField {
	return MinPriceIncrementField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MinPriceIncrementField) Value() (val decimal.Decimal) { return f.Decimal }

//MinPriceIncrementAmountField is a AMT field
type MinPriceIncrementAmountField struct{ quickfix.FIXDecimal }

//Tag returns tag.MinPriceIncrementAmount (1146)
func (f MinPriceIncrementAmountField) Tag() quickfix.Tag { return tag.MinPriceIncrementAmount }

//NewMinPriceIncrementAmount returns a new MinPriceIncrementAmountField initialized with val and scale
func NewMinPriceIncrementAmount(val decimal.Decimal, scale int32) MinPriceIncrementAmountField {
	return MinPriceIncrementAmountField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MinPriceIncrementAmountField) Value() (val decimal.Decimal) { return f.Decimal }

//MinQtyField is a QTY field
type MinQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.MinQty (110)
func (f MinQtyField) Tag() quickfix.Tag { return tag.MinQty }

//NewMinQty returns a new MinQtyField initialized with val and scale
func NewMinQty(val decimal.Decimal, scale int32) MinQtyField {
	return MinQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MinQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//MinTradeVolField is a QTY field
type MinTradeVolField struct{ quickfix.FIXDecimal }

//Tag returns tag.MinTradeVol (562)
func (f MinTradeVolField) Tag() quickfix.Tag { return tag.MinTradeVol }

//NewMinTradeVol returns a new MinTradeVolField initialized with val and scale
func NewMinTradeVol(val decimal.Decimal, scale int32) MinTradeVolField {
	return MinTradeVolField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MinTradeVolField) Value() (val decimal.Decimal) { return f.Decimal }

//MiscFeeAmtField is a AMT field
type MiscFeeAmtField struct{ quickfix.FIXDecimal }

//Tag returns tag.MiscFeeAmt (137)
func (f MiscFeeAmtField) Tag() quickfix.Tag { return tag.MiscFeeAmt }

//NewMiscFeeAmt returns a new MiscFeeAmtField initialized with val and scale
func NewMiscFeeAmt(val decimal.Decimal, scale int32) MiscFeeAmtField {
	return MiscFeeAmtField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MiscFeeAmtField) Value() (val decimal.Decimal) { return f.Decimal }

//MiscFeeBasisField is a enum.MiscFeeBasis field
type MiscFeeBasisField struct{ quickfix.FIXString }

//Tag returns tag.MiscFeeBasis (891)
func (f MiscFeeBasisField) Tag() quickfix.Tag { return tag.MiscFeeBasis }

func NewMiscFeeBasis(val enum.MiscFeeBasis) MiscFeeBasisField {
	return MiscFeeBasisField{quickfix.FIXString(val)}
}

func (f MiscFeeBasisField) Value() enum.MiscFeeBasis { return enum.MiscFeeBasis(f.String()) }

//MiscFeeCurrField is a CURRENCY field
type MiscFeeCurrField struct{ quickfix.FIXString }

//Tag returns tag.MiscFeeCurr (138)
func (f MiscFeeCurrField) Tag() quickfix.Tag { return tag.MiscFeeCurr }

//NewMiscFeeCurr returns a new MiscFeeCurrField initialized with val
func NewMiscFeeCurr(val string) MiscFeeCurrField {
	return MiscFeeCurrField{quickfix.FIXString(val)}
}

func (f MiscFeeCurrField) Value() string { return f.String() }

//MiscFeeTypeField is a enum.MiscFeeType field
type MiscFeeTypeField struct{ quickfix.FIXString }

//Tag returns tag.MiscFeeType (139)
func (f MiscFeeTypeField) Tag() quickfix.Tag { return tag.MiscFeeType }

func NewMiscFeeType(val enum.MiscFeeType) MiscFeeTypeField {
	return MiscFeeTypeField{quickfix.FIXString(val)}
}

func (f MiscFeeTypeField) Value() enum.MiscFeeType { return enum.MiscFeeType(f.String()) }

//MktBidPxField is a PRICE field
type MktBidPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.MktBidPx (645)
func (f MktBidPxField) Tag() quickfix.Tag { return tag.MktBidPx }

//NewMktBidPx returns a new MktBidPxField initialized with val and scale
func NewMktBidPx(val decimal.Decimal, scale int32) MktBidPxField {
	return MktBidPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MktBidPxField) Value() (val decimal.Decimal) { return f.Decimal }

//MktOfferPxField is a PRICE field
type MktOfferPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.MktOfferPx (646)
func (f MktOfferPxField) Tag() quickfix.Tag { return tag.MktOfferPx }

//NewMktOfferPx returns a new MktOfferPxField initialized with val and scale
func NewMktOfferPx(val decimal.Decimal, scale int32) MktOfferPxField {
	return MktOfferPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MktOfferPxField) Value() (val decimal.Decimal) { return f.Decimal }

//ModelTypeField is a enum.ModelType field
type ModelTypeField struct{ quickfix.FIXString }

//Tag returns tag.ModelType (1434)
func (f ModelTypeField) Tag() quickfix.Tag { return tag.ModelType }

func NewModelType(val enum.ModelType) ModelTypeField {
	return ModelTypeField{quickfix.FIXString(val)}
}

func (f ModelTypeField) Value() enum.ModelType { return enum.ModelType(f.String()) }

//MoneyLaunderingStatusField is a enum.MoneyLaunderingStatus field
type MoneyLaunderingStatusField struct{ quickfix.FIXString }

//Tag returns tag.MoneyLaunderingStatus (481)
func (f MoneyLaunderingStatusField) Tag() quickfix.Tag { return tag.MoneyLaunderingStatus }

func NewMoneyLaunderingStatus(val enum.MoneyLaunderingStatus) MoneyLaunderingStatusField {
	return MoneyLaunderingStatusField{quickfix.FIXString(val)}
}

func (f MoneyLaunderingStatusField) Value() enum.MoneyLaunderingStatus {
	return enum.MoneyLaunderingStatus(f.String())
}

//MsgDirectionField is a enum.MsgDirection field
type MsgDirectionField struct{ quickfix.FIXString }

//Tag returns tag.MsgDirection (385)
func (f MsgDirectionField) Tag() quickfix.Tag { return tag.MsgDirection }

func NewMsgDirection(val enum.MsgDirection) MsgDirectionField {
	return MsgDirectionField{quickfix.FIXString(val)}
}

func (f MsgDirectionField) Value() enum.MsgDirection { return enum.MsgDirection(f.String()) }

//MsgSeqNumField is a SEQNUM field
type MsgSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.MsgSeqNum (34)
func (f MsgSeqNumField) Tag() quickfix.Tag { return tag.MsgSeqNum }

//NewMsgSeqNum returns a new MsgSeqNumField initialized with val
func NewMsgSeqNum(val int) MsgSeqNumField {
	return MsgSeqNumField{quickfix.FIXInt(val)}
}

func (f MsgSeqNumField) Value() int { return f.Int() }

//MsgTypeField is a enum.MsgType field
type MsgTypeField struct{ quickfix.FIXString }

//Tag returns tag.MsgType (35)
func (f MsgTypeField) Tag() quickfix.Tag { return tag.MsgType }

func NewMsgType(val enum.MsgType) MsgTypeField {
	return MsgTypeField{quickfix.FIXString(val)}
}

func (f MsgTypeField) Value() enum.MsgType { return enum.MsgType(f.String()) }

//MultiLegReportingTypeField is a enum.MultiLegReportingType field
type MultiLegReportingTypeField struct{ quickfix.FIXString }

//Tag returns tag.MultiLegReportingType (442)
func (f MultiLegReportingTypeField) Tag() quickfix.Tag { return tag.MultiLegReportingType }

func NewMultiLegReportingType(val enum.MultiLegReportingType) MultiLegReportingTypeField {
	return MultiLegReportingTypeField{quickfix.FIXString(val)}
}

func (f MultiLegReportingTypeField) Value() enum.MultiLegReportingType {
	return enum.MultiLegReportingType(f.String())
}

//MultiLegRptTypeReqField is a enum.MultiLegRptTypeReq field
type MultiLegRptTypeReqField struct{ quickfix.FIXString }

//Tag returns tag.MultiLegRptTypeReq (563)
func (f MultiLegRptTypeReqField) Tag() quickfix.Tag { return tag.MultiLegRptTypeReq }

func NewMultiLegRptTypeReq(val enum.MultiLegRptTypeReq) MultiLegRptTypeReqField {
	return MultiLegRptTypeReqField{quickfix.FIXString(val)}
}

func (f MultiLegRptTypeReqField) Value() enum.MultiLegRptTypeReq {
	return enum.MultiLegRptTypeReq(f.String())
}

//MultilegModelField is a enum.MultilegModel field
type MultilegModelField struct{ quickfix.FIXString }

//Tag returns tag.MultilegModel (1377)
func (f MultilegModelField) Tag() quickfix.Tag { return tag.MultilegModel }

func NewMultilegModel(val enum.MultilegModel) MultilegModelField {
	return MultilegModelField{quickfix.FIXString(val)}
}

func (f MultilegModelField) Value() enum.MultilegModel { return enum.MultilegModel(f.String()) }

//MultilegPriceMethodField is a enum.MultilegPriceMethod field
type MultilegPriceMethodField struct{ quickfix.FIXString }

//Tag returns tag.MultilegPriceMethod (1378)
func (f MultilegPriceMethodField) Tag() quickfix.Tag { return tag.MultilegPriceMethod }

func NewMultilegPriceMethod(val enum.MultilegPriceMethod) MultilegPriceMethodField {
	return MultilegPriceMethodField{quickfix.FIXString(val)}
}

func (f MultilegPriceMethodField) Value() enum.MultilegPriceMethod {
	return enum.MultilegPriceMethod(f.String())
}

//NTPositionLimitField is a INT field
type NTPositionLimitField struct{ quickfix.FIXInt }

//Tag returns tag.NTPositionLimit (971)
func (f NTPositionLimitField) Tag() quickfix.Tag { return tag.NTPositionLimit }

//NewNTPositionLimit returns a new NTPositionLimitField initialized with val
func NewNTPositionLimit(val int) NTPositionLimitField {
	return NTPositionLimitField{quickfix.FIXInt(val)}
}

func (f NTPositionLimitField) Value() int { return f.Int() }

//Nested2PartyIDField is a STRING field
type Nested2PartyIDField struct{ quickfix.FIXString }

//Tag returns tag.Nested2PartyID (757)
func (f Nested2PartyIDField) Tag() quickfix.Tag { return tag.Nested2PartyID }

//NewNested2PartyID returns a new Nested2PartyIDField initialized with val
func NewNested2PartyID(val string) Nested2PartyIDField {
	return Nested2PartyIDField{quickfix.FIXString(val)}
}

func (f Nested2PartyIDField) Value() string { return f.String() }

//Nested2PartyIDSourceField is a CHAR field
type Nested2PartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.Nested2PartyIDSource (758)
func (f Nested2PartyIDSourceField) Tag() quickfix.Tag { return tag.Nested2PartyIDSource }

//NewNested2PartyIDSource returns a new Nested2PartyIDSourceField initialized with val
func NewNested2PartyIDSource(val string) Nested2PartyIDSourceField {
	return Nested2PartyIDSourceField{quickfix.FIXString(val)}
}

func (f Nested2PartyIDSourceField) Value() string { return f.String() }

//Nested2PartyRoleField is a INT field
type Nested2PartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.Nested2PartyRole (759)
func (f Nested2PartyRoleField) Tag() quickfix.Tag { return tag.Nested2PartyRole }

//NewNested2PartyRole returns a new Nested2PartyRoleField initialized with val
func NewNested2PartyRole(val int) Nested2PartyRoleField {
	return Nested2PartyRoleField{quickfix.FIXInt(val)}
}

func (f Nested2PartyRoleField) Value() int { return f.Int() }

//Nested2PartySubIDField is a STRING field
type Nested2PartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.Nested2PartySubID (760)
func (f Nested2PartySubIDField) Tag() quickfix.Tag { return tag.Nested2PartySubID }

//NewNested2PartySubID returns a new Nested2PartySubIDField initialized with val
func NewNested2PartySubID(val string) Nested2PartySubIDField {
	return Nested2PartySubIDField{quickfix.FIXString(val)}
}

func (f Nested2PartySubIDField) Value() string { return f.String() }

//Nested2PartySubIDTypeField is a INT field
type Nested2PartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.Nested2PartySubIDType (807)
func (f Nested2PartySubIDTypeField) Tag() quickfix.Tag { return tag.Nested2PartySubIDType }

//NewNested2PartySubIDType returns a new Nested2PartySubIDTypeField initialized with val
func NewNested2PartySubIDType(val int) Nested2PartySubIDTypeField {
	return Nested2PartySubIDTypeField{quickfix.FIXInt(val)}
}

func (f Nested2PartySubIDTypeField) Value() int { return f.Int() }

//Nested3PartyIDField is a STRING field
type Nested3PartyIDField struct{ quickfix.FIXString }

//Tag returns tag.Nested3PartyID (949)
func (f Nested3PartyIDField) Tag() quickfix.Tag { return tag.Nested3PartyID }

//NewNested3PartyID returns a new Nested3PartyIDField initialized with val
func NewNested3PartyID(val string) Nested3PartyIDField {
	return Nested3PartyIDField{quickfix.FIXString(val)}
}

func (f Nested3PartyIDField) Value() string { return f.String() }

//Nested3PartyIDSourceField is a CHAR field
type Nested3PartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.Nested3PartyIDSource (950)
func (f Nested3PartyIDSourceField) Tag() quickfix.Tag { return tag.Nested3PartyIDSource }

//NewNested3PartyIDSource returns a new Nested3PartyIDSourceField initialized with val
func NewNested3PartyIDSource(val string) Nested3PartyIDSourceField {
	return Nested3PartyIDSourceField{quickfix.FIXString(val)}
}

func (f Nested3PartyIDSourceField) Value() string { return f.String() }

//Nested3PartyRoleField is a INT field
type Nested3PartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.Nested3PartyRole (951)
func (f Nested3PartyRoleField) Tag() quickfix.Tag { return tag.Nested3PartyRole }

//NewNested3PartyRole returns a new Nested3PartyRoleField initialized with val
func NewNested3PartyRole(val int) Nested3PartyRoleField {
	return Nested3PartyRoleField{quickfix.FIXInt(val)}
}

func (f Nested3PartyRoleField) Value() int { return f.Int() }

//Nested3PartySubIDField is a STRING field
type Nested3PartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.Nested3PartySubID (953)
func (f Nested3PartySubIDField) Tag() quickfix.Tag { return tag.Nested3PartySubID }

//NewNested3PartySubID returns a new Nested3PartySubIDField initialized with val
func NewNested3PartySubID(val string) Nested3PartySubIDField {
	return Nested3PartySubIDField{quickfix.FIXString(val)}
}

func (f Nested3PartySubIDField) Value() string { return f.String() }

//Nested3PartySubIDTypeField is a INT field
type Nested3PartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.Nested3PartySubIDType (954)
func (f Nested3PartySubIDTypeField) Tag() quickfix.Tag { return tag.Nested3PartySubIDType }

//NewNested3PartySubIDType returns a new Nested3PartySubIDTypeField initialized with val
func NewNested3PartySubIDType(val int) Nested3PartySubIDTypeField {
	return Nested3PartySubIDTypeField{quickfix.FIXInt(val)}
}

func (f Nested3PartySubIDTypeField) Value() int { return f.Int() }

//Nested4PartyIDField is a STRING field
type Nested4PartyIDField struct{ quickfix.FIXString }

//Tag returns tag.Nested4PartyID (1415)
func (f Nested4PartyIDField) Tag() quickfix.Tag { return tag.Nested4PartyID }

//NewNested4PartyID returns a new Nested4PartyIDField initialized with val
func NewNested4PartyID(val string) Nested4PartyIDField {
	return Nested4PartyIDField{quickfix.FIXString(val)}
}

func (f Nested4PartyIDField) Value() string { return f.String() }

//Nested4PartyIDSourceField is a CHAR field
type Nested4PartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.Nested4PartyIDSource (1416)
func (f Nested4PartyIDSourceField) Tag() quickfix.Tag { return tag.Nested4PartyIDSource }

//NewNested4PartyIDSource returns a new Nested4PartyIDSourceField initialized with val
func NewNested4PartyIDSource(val string) Nested4PartyIDSourceField {
	return Nested4PartyIDSourceField{quickfix.FIXString(val)}
}

func (f Nested4PartyIDSourceField) Value() string { return f.String() }

//Nested4PartyRoleField is a INT field
type Nested4PartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.Nested4PartyRole (1417)
func (f Nested4PartyRoleField) Tag() quickfix.Tag { return tag.Nested4PartyRole }

//NewNested4PartyRole returns a new Nested4PartyRoleField initialized with val
func NewNested4PartyRole(val int) Nested4PartyRoleField {
	return Nested4PartyRoleField{quickfix.FIXInt(val)}
}

func (f Nested4PartyRoleField) Value() int { return f.Int() }

//Nested4PartySubIDField is a STRING field
type Nested4PartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.Nested4PartySubID (1412)
func (f Nested4PartySubIDField) Tag() quickfix.Tag { return tag.Nested4PartySubID }

//NewNested4PartySubID returns a new Nested4PartySubIDField initialized with val
func NewNested4PartySubID(val string) Nested4PartySubIDField {
	return Nested4PartySubIDField{quickfix.FIXString(val)}
}

func (f Nested4PartySubIDField) Value() string { return f.String() }

//Nested4PartySubIDTypeField is a INT field
type Nested4PartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.Nested4PartySubIDType (1411)
func (f Nested4PartySubIDTypeField) Tag() quickfix.Tag { return tag.Nested4PartySubIDType }

//NewNested4PartySubIDType returns a new Nested4PartySubIDTypeField initialized with val
func NewNested4PartySubIDType(val int) Nested4PartySubIDTypeField {
	return Nested4PartySubIDTypeField{quickfix.FIXInt(val)}
}

func (f Nested4PartySubIDTypeField) Value() int { return f.Int() }

//NestedInstrAttribTypeField is a INT field
type NestedInstrAttribTypeField struct{ quickfix.FIXInt }

//Tag returns tag.NestedInstrAttribType (1210)
func (f NestedInstrAttribTypeField) Tag() quickfix.Tag { return tag.NestedInstrAttribType }

//NewNestedInstrAttribType returns a new NestedInstrAttribTypeField initialized with val
func NewNestedInstrAttribType(val int) NestedInstrAttribTypeField {
	return NestedInstrAttribTypeField{quickfix.FIXInt(val)}
}

func (f NestedInstrAttribTypeField) Value() int { return f.Int() }

//NestedInstrAttribValueField is a STRING field
type NestedInstrAttribValueField struct{ quickfix.FIXString }

//Tag returns tag.NestedInstrAttribValue (1211)
func (f NestedInstrAttribValueField) Tag() quickfix.Tag { return tag.NestedInstrAttribValue }

//NewNestedInstrAttribValue returns a new NestedInstrAttribValueField initialized with val
func NewNestedInstrAttribValue(val string) NestedInstrAttribValueField {
	return NestedInstrAttribValueField{quickfix.FIXString(val)}
}

func (f NestedInstrAttribValueField) Value() string { return f.String() }

//NestedPartyIDField is a STRING field
type NestedPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.NestedPartyID (524)
func (f NestedPartyIDField) Tag() quickfix.Tag { return tag.NestedPartyID }

//NewNestedPartyID returns a new NestedPartyIDField initialized with val
func NewNestedPartyID(val string) NestedPartyIDField {
	return NestedPartyIDField{quickfix.FIXString(val)}
}

func (f NestedPartyIDField) Value() string { return f.String() }

//NestedPartyIDSourceField is a CHAR field
type NestedPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.NestedPartyIDSource (525)
func (f NestedPartyIDSourceField) Tag() quickfix.Tag { return tag.NestedPartyIDSource }

//NewNestedPartyIDSource returns a new NestedPartyIDSourceField initialized with val
func NewNestedPartyIDSource(val string) NestedPartyIDSourceField {
	return NestedPartyIDSourceField{quickfix.FIXString(val)}
}

func (f NestedPartyIDSourceField) Value() string { return f.String() }

//NestedPartyRoleField is a INT field
type NestedPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.NestedPartyRole (538)
func (f NestedPartyRoleField) Tag() quickfix.Tag { return tag.NestedPartyRole }

//NewNestedPartyRole returns a new NestedPartyRoleField initialized with val
func NewNestedPartyRole(val int) NestedPartyRoleField {
	return NestedPartyRoleField{quickfix.FIXInt(val)}
}

func (f NestedPartyRoleField) Value() int { return f.Int() }

//NestedPartySubIDField is a STRING field
type NestedPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.NestedPartySubID (545)
func (f NestedPartySubIDField) Tag() quickfix.Tag { return tag.NestedPartySubID }

//NewNestedPartySubID returns a new NestedPartySubIDField initialized with val
func NewNestedPartySubID(val string) NestedPartySubIDField {
	return NestedPartySubIDField{quickfix.FIXString(val)}
}

func (f NestedPartySubIDField) Value() string { return f.String() }

//NestedPartySubIDTypeField is a INT field
type NestedPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.NestedPartySubIDType (805)
func (f NestedPartySubIDTypeField) Tag() quickfix.Tag { return tag.NestedPartySubIDType }

//NewNestedPartySubIDType returns a new NestedPartySubIDTypeField initialized with val
func NewNestedPartySubIDType(val int) NestedPartySubIDTypeField {
	return NestedPartySubIDTypeField{quickfix.FIXInt(val)}
}

func (f NestedPartySubIDTypeField) Value() int { return f.Int() }

//NetChgPrevDayField is a PRICEOFFSET field
type NetChgPrevDayField struct{ quickfix.FIXDecimal }

//Tag returns tag.NetChgPrevDay (451)
func (f NetChgPrevDayField) Tag() quickfix.Tag { return tag.NetChgPrevDay }

//NewNetChgPrevDay returns a new NetChgPrevDayField initialized with val and scale
func NewNetChgPrevDay(val decimal.Decimal, scale int32) NetChgPrevDayField {
	return NetChgPrevDayField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f NetChgPrevDayField) Value() (val decimal.Decimal) { return f.Decimal }

//NetGrossIndField is a enum.NetGrossInd field
type NetGrossIndField struct{ quickfix.FIXString }

//Tag returns tag.NetGrossInd (430)
func (f NetGrossIndField) Tag() quickfix.Tag { return tag.NetGrossInd }

func NewNetGrossInd(val enum.NetGrossInd) NetGrossIndField {
	return NetGrossIndField{quickfix.FIXString(val)}
}

func (f NetGrossIndField) Value() enum.NetGrossInd { return enum.NetGrossInd(f.String()) }

//NetMoneyField is a AMT field
type NetMoneyField struct{ quickfix.FIXDecimal }

//Tag returns tag.NetMoney (118)
func (f NetMoneyField) Tag() quickfix.Tag { return tag.NetMoney }

//NewNetMoney returns a new NetMoneyField initialized with val and scale
func NewNetMoney(val decimal.Decimal, scale int32) NetMoneyField {
	return NetMoneyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f NetMoneyField) Value() (val decimal.Decimal) { return f.Decimal }

//NetworkRequestIDField is a STRING field
type NetworkRequestIDField struct{ quickfix.FIXString }

//Tag returns tag.NetworkRequestID (933)
func (f NetworkRequestIDField) Tag() quickfix.Tag { return tag.NetworkRequestID }

//NewNetworkRequestID returns a new NetworkRequestIDField initialized with val
func NewNetworkRequestID(val string) NetworkRequestIDField {
	return NetworkRequestIDField{quickfix.FIXString(val)}
}

func (f NetworkRequestIDField) Value() string { return f.String() }

//NetworkRequestTypeField is a enum.NetworkRequestType field
type NetworkRequestTypeField struct{ quickfix.FIXString }

//Tag returns tag.NetworkRequestType (935)
func (f NetworkRequestTypeField) Tag() quickfix.Tag { return tag.NetworkRequestType }

func NewNetworkRequestType(val enum.NetworkRequestType) NetworkRequestTypeField {
	return NetworkRequestTypeField{quickfix.FIXString(val)}
}

func (f NetworkRequestTypeField) Value() enum.NetworkRequestType {
	return enum.NetworkRequestType(f.String())
}

//NetworkResponseIDField is a STRING field
type NetworkResponseIDField struct{ quickfix.FIXString }

//Tag returns tag.NetworkResponseID (932)
func (f NetworkResponseIDField) Tag() quickfix.Tag { return tag.NetworkResponseID }

//NewNetworkResponseID returns a new NetworkResponseIDField initialized with val
func NewNetworkResponseID(val string) NetworkResponseIDField {
	return NetworkResponseIDField{quickfix.FIXString(val)}
}

func (f NetworkResponseIDField) Value() string { return f.String() }

//NetworkStatusResponseTypeField is a enum.NetworkStatusResponseType field
type NetworkStatusResponseTypeField struct{ quickfix.FIXString }

//Tag returns tag.NetworkStatusResponseType (937)
func (f NetworkStatusResponseTypeField) Tag() quickfix.Tag { return tag.NetworkStatusResponseType }

func NewNetworkStatusResponseType(val enum.NetworkStatusResponseType) NetworkStatusResponseTypeField {
	return NetworkStatusResponseTypeField{quickfix.FIXString(val)}
}

func (f NetworkStatusResponseTypeField) Value() enum.NetworkStatusResponseType {
	return enum.NetworkStatusResponseType(f.String())
}

//NewPasswordField is a STRING field
type NewPasswordField struct{ quickfix.FIXString }

//Tag returns tag.NewPassword (925)
func (f NewPasswordField) Tag() quickfix.Tag { return tag.NewPassword }

//NewNewPassword returns a new NewPasswordField initialized with val
func NewNewPassword(val string) NewPasswordField {
	return NewPasswordField{quickfix.FIXString(val)}
}

func (f NewPasswordField) Value() string { return f.String() }

//NewSeqNoField is a SEQNUM field
type NewSeqNoField struct{ quickfix.FIXInt }

//Tag returns tag.NewSeqNo (36)
func (f NewSeqNoField) Tag() quickfix.Tag { return tag.NewSeqNo }

//NewNewSeqNo returns a new NewSeqNoField initialized with val
func NewNewSeqNo(val int) NewSeqNoField {
	return NewSeqNoField{quickfix.FIXInt(val)}
}

func (f NewSeqNoField) Value() int { return f.Int() }

//NewsCategoryField is a enum.NewsCategory field
type NewsCategoryField struct{ quickfix.FIXString }

//Tag returns tag.NewsCategory (1473)
func (f NewsCategoryField) Tag() quickfix.Tag { return tag.NewsCategory }

func NewNewsCategory(val enum.NewsCategory) NewsCategoryField {
	return NewsCategoryField{quickfix.FIXString(val)}
}

func (f NewsCategoryField) Value() enum.NewsCategory { return enum.NewsCategory(f.String()) }

//NewsIDField is a STRING field
type NewsIDField struct{ quickfix.FIXString }

//Tag returns tag.NewsID (1472)
func (f NewsIDField) Tag() quickfix.Tag { return tag.NewsID }

//NewNewsID returns a new NewsIDField initialized with val
func NewNewsID(val string) NewsIDField {
	return NewsIDField{quickfix.FIXString(val)}
}

func (f NewsIDField) Value() string { return f.String() }

//NewsRefIDField is a STRING field
type NewsRefIDField struct{ quickfix.FIXString }

//Tag returns tag.NewsRefID (1476)
func (f NewsRefIDField) Tag() quickfix.Tag { return tag.NewsRefID }

//NewNewsRefID returns a new NewsRefIDField initialized with val
func NewNewsRefID(val string) NewsRefIDField {
	return NewsRefIDField{quickfix.FIXString(val)}
}

func (f NewsRefIDField) Value() string { return f.String() }

//NewsRefTypeField is a enum.NewsRefType field
type NewsRefTypeField struct{ quickfix.FIXString }

//Tag returns tag.NewsRefType (1477)
func (f NewsRefTypeField) Tag() quickfix.Tag { return tag.NewsRefType }

func NewNewsRefType(val enum.NewsRefType) NewsRefTypeField {
	return NewsRefTypeField{quickfix.FIXString(val)}
}

func (f NewsRefTypeField) Value() enum.NewsRefType { return enum.NewsRefType(f.String()) }

//NextExpectedMsgSeqNumField is a SEQNUM field
type NextExpectedMsgSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.NextExpectedMsgSeqNum (789)
func (f NextExpectedMsgSeqNumField) Tag() quickfix.Tag { return tag.NextExpectedMsgSeqNum }

//NewNextExpectedMsgSeqNum returns a new NextExpectedMsgSeqNumField initialized with val
func NewNextExpectedMsgSeqNum(val int) NextExpectedMsgSeqNumField {
	return NextExpectedMsgSeqNumField{quickfix.FIXInt(val)}
}

func (f NextExpectedMsgSeqNumField) Value() int { return f.Int() }

//NoAffectedOrdersField is a NUMINGROUP field
type NoAffectedOrdersField struct{ quickfix.FIXInt }

//Tag returns tag.NoAffectedOrders (534)
func (f NoAffectedOrdersField) Tag() quickfix.Tag { return tag.NoAffectedOrders }

//NewNoAffectedOrders returns a new NoAffectedOrdersField initialized with val
func NewNoAffectedOrders(val int) NoAffectedOrdersField {
	return NoAffectedOrdersField{quickfix.FIXInt(val)}
}

func (f NoAffectedOrdersField) Value() int { return f.Int() }

//NoAllocsField is a NUMINGROUP field
type NoAllocsField struct{ quickfix.FIXInt }

//Tag returns tag.NoAllocs (78)
func (f NoAllocsField) Tag() quickfix.Tag { return tag.NoAllocs }

//NewNoAllocs returns a new NoAllocsField initialized with val
func NewNoAllocs(val int) NoAllocsField {
	return NoAllocsField{quickfix.FIXInt(val)}
}

func (f NoAllocsField) Value() int { return f.Int() }

//NoAltMDSourceField is a NUMINGROUP field
type NoAltMDSourceField struct{ quickfix.FIXInt }

//Tag returns tag.NoAltMDSource (816)
func (f NoAltMDSourceField) Tag() quickfix.Tag { return tag.NoAltMDSource }

//NewNoAltMDSource returns a new NoAltMDSourceField initialized with val
func NewNoAltMDSource(val int) NoAltMDSourceField {
	return NoAltMDSourceField{quickfix.FIXInt(val)}
}

func (f NoAltMDSourceField) Value() int { return f.Int() }

//NoApplIDsField is a NUMINGROUP field
type NoApplIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoApplIDs (1351)
func (f NoApplIDsField) Tag() quickfix.Tag { return tag.NoApplIDs }

//NewNoApplIDs returns a new NoApplIDsField initialized with val
func NewNoApplIDs(val int) NoApplIDsField {
	return NoApplIDsField{quickfix.FIXInt(val)}
}

func (f NoApplIDsField) Value() int { return f.Int() }

//NoAsgnReqsField is a NUMINGROUP field
type NoAsgnReqsField struct{ quickfix.FIXInt }

//Tag returns tag.NoAsgnReqs (1499)
func (f NoAsgnReqsField) Tag() quickfix.Tag { return tag.NoAsgnReqs }

//NewNoAsgnReqs returns a new NoAsgnReqsField initialized with val
func NewNoAsgnReqs(val int) NoAsgnReqsField {
	return NoAsgnReqsField{quickfix.FIXInt(val)}
}

func (f NoAsgnReqsField) Value() int { return f.Int() }

//NoBidComponentsField is a NUMINGROUP field
type NoBidComponentsField struct{ quickfix.FIXInt }

//Tag returns tag.NoBidComponents (420)
func (f NoBidComponentsField) Tag() quickfix.Tag { return tag.NoBidComponents }

//NewNoBidComponents returns a new NoBidComponentsField initialized with val
func NewNoBidComponents(val int) NoBidComponentsField {
	return NoBidComponentsField{quickfix.FIXInt(val)}
}

func (f NoBidComponentsField) Value() int { return f.Int() }

//NoBidDescriptorsField is a NUMINGROUP field
type NoBidDescriptorsField struct{ quickfix.FIXInt }

//Tag returns tag.NoBidDescriptors (398)
func (f NoBidDescriptorsField) Tag() quickfix.Tag { return tag.NoBidDescriptors }

//NewNoBidDescriptors returns a new NoBidDescriptorsField initialized with val
func NewNoBidDescriptors(val int) NoBidDescriptorsField {
	return NoBidDescriptorsField{quickfix.FIXInt(val)}
}

func (f NoBidDescriptorsField) Value() int { return f.Int() }

//NoCapacitiesField is a NUMINGROUP field
type NoCapacitiesField struct{ quickfix.FIXInt }

//Tag returns tag.NoCapacities (862)
func (f NoCapacitiesField) Tag() quickfix.Tag { return tag.NoCapacities }

//NewNoCapacities returns a new NoCapacitiesField initialized with val
func NewNoCapacities(val int) NoCapacitiesField {
	return NoCapacitiesField{quickfix.FIXInt(val)}
}

func (f NoCapacitiesField) Value() int { return f.Int() }

//NoClearingInstructionsField is a NUMINGROUP field
type NoClearingInstructionsField struct{ quickfix.FIXInt }

//Tag returns tag.NoClearingInstructions (576)
func (f NoClearingInstructionsField) Tag() quickfix.Tag { return tag.NoClearingInstructions }

//NewNoClearingInstructions returns a new NoClearingInstructionsField initialized with val
func NewNoClearingInstructions(val int) NoClearingInstructionsField {
	return NoClearingInstructionsField{quickfix.FIXInt(val)}
}

func (f NoClearingInstructionsField) Value() int { return f.Int() }

//NoCollInquiryQualifierField is a NUMINGROUP field
type NoCollInquiryQualifierField struct{ quickfix.FIXInt }

//Tag returns tag.NoCollInquiryQualifier (938)
func (f NoCollInquiryQualifierField) Tag() quickfix.Tag { return tag.NoCollInquiryQualifier }

//NewNoCollInquiryQualifier returns a new NoCollInquiryQualifierField initialized with val
func NewNoCollInquiryQualifier(val int) NoCollInquiryQualifierField {
	return NoCollInquiryQualifierField{quickfix.FIXInt(val)}
}

func (f NoCollInquiryQualifierField) Value() int { return f.Int() }

//NoCompIDsField is a NUMINGROUP field
type NoCompIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoCompIDs (936)
func (f NoCompIDsField) Tag() quickfix.Tag { return tag.NoCompIDs }

//NewNoCompIDs returns a new NoCompIDsField initialized with val
func NewNoCompIDs(val int) NoCompIDsField {
	return NoCompIDsField{quickfix.FIXInt(val)}
}

func (f NoCompIDsField) Value() int { return f.Int() }

//NoComplexEventDatesField is a NUMINGROUP field
type NoComplexEventDatesField struct{ quickfix.FIXInt }

//Tag returns tag.NoComplexEventDates (1491)
func (f NoComplexEventDatesField) Tag() quickfix.Tag { return tag.NoComplexEventDates }

//NewNoComplexEventDates returns a new NoComplexEventDatesField initialized with val
func NewNoComplexEventDates(val int) NoComplexEventDatesField {
	return NoComplexEventDatesField{quickfix.FIXInt(val)}
}

func (f NoComplexEventDatesField) Value() int { return f.Int() }

//NoComplexEventTimesField is a NUMINGROUP field
type NoComplexEventTimesField struct{ quickfix.FIXInt }

//Tag returns tag.NoComplexEventTimes (1494)
func (f NoComplexEventTimesField) Tag() quickfix.Tag { return tag.NoComplexEventTimes }

//NewNoComplexEventTimes returns a new NoComplexEventTimesField initialized with val
func NewNoComplexEventTimes(val int) NoComplexEventTimesField {
	return NoComplexEventTimesField{quickfix.FIXInt(val)}
}

func (f NoComplexEventTimesField) Value() int { return f.Int() }

//NoComplexEventsField is a NUMINGROUP field
type NoComplexEventsField struct{ quickfix.FIXInt }

//Tag returns tag.NoComplexEvents (1483)
func (f NoComplexEventsField) Tag() quickfix.Tag { return tag.NoComplexEvents }

//NewNoComplexEvents returns a new NoComplexEventsField initialized with val
func NewNoComplexEvents(val int) NoComplexEventsField {
	return NoComplexEventsField{quickfix.FIXInt(val)}
}

func (f NoComplexEventsField) Value() int { return f.Int() }

//NoContAmtsField is a NUMINGROUP field
type NoContAmtsField struct{ quickfix.FIXInt }

//Tag returns tag.NoContAmts (518)
func (f NoContAmtsField) Tag() quickfix.Tag { return tag.NoContAmts }

//NewNoContAmts returns a new NoContAmtsField initialized with val
func NewNoContAmts(val int) NoContAmtsField {
	return NoContAmtsField{quickfix.FIXInt(val)}
}

func (f NoContAmtsField) Value() int { return f.Int() }

//NoContextPartyIDsField is a NUMINGROUP field
type NoContextPartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoContextPartyIDs (1522)
func (f NoContextPartyIDsField) Tag() quickfix.Tag { return tag.NoContextPartyIDs }

//NewNoContextPartyIDs returns a new NoContextPartyIDsField initialized with val
func NewNoContextPartyIDs(val int) NoContextPartyIDsField {
	return NoContextPartyIDsField{quickfix.FIXInt(val)}
}

func (f NoContextPartyIDsField) Value() int { return f.Int() }

//NoContextPartySubIDsField is a NUMINGROUP field
type NoContextPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoContextPartySubIDs (1526)
func (f NoContextPartySubIDsField) Tag() quickfix.Tag { return tag.NoContextPartySubIDs }

//NewNoContextPartySubIDs returns a new NoContextPartySubIDsField initialized with val
func NewNoContextPartySubIDs(val int) NoContextPartySubIDsField {
	return NoContextPartySubIDsField{quickfix.FIXInt(val)}
}

func (f NoContextPartySubIDsField) Value() int { return f.Int() }

//NoContraBrokersField is a NUMINGROUP field
type NoContraBrokersField struct{ quickfix.FIXInt }

//Tag returns tag.NoContraBrokers (382)
func (f NoContraBrokersField) Tag() quickfix.Tag { return tag.NoContraBrokers }

//NewNoContraBrokers returns a new NoContraBrokersField initialized with val
func NewNoContraBrokers(val int) NoContraBrokersField {
	return NoContraBrokersField{quickfix.FIXInt(val)}
}

func (f NoContraBrokersField) Value() int { return f.Int() }

//NoDatesField is a INT field
type NoDatesField struct{ quickfix.FIXInt }

//Tag returns tag.NoDates (580)
func (f NoDatesField) Tag() quickfix.Tag { return tag.NoDates }

//NewNoDates returns a new NoDatesField initialized with val
func NewNoDates(val int) NoDatesField {
	return NoDatesField{quickfix.FIXInt(val)}
}

func (f NoDatesField) Value() int { return f.Int() }

//NoDerivativeEventsField is a NUMINGROUP field
type NoDerivativeEventsField struct{ quickfix.FIXInt }

//Tag returns tag.NoDerivativeEvents (1286)
func (f NoDerivativeEventsField) Tag() quickfix.Tag { return tag.NoDerivativeEvents }

//NewNoDerivativeEvents returns a new NoDerivativeEventsField initialized with val
func NewNoDerivativeEvents(val int) NoDerivativeEventsField {
	return NoDerivativeEventsField{quickfix.FIXInt(val)}
}

func (f NoDerivativeEventsField) Value() int { return f.Int() }

//NoDerivativeInstrAttribField is a NUMINGROUP field
type NoDerivativeInstrAttribField struct{ quickfix.FIXInt }

//Tag returns tag.NoDerivativeInstrAttrib (1311)
func (f NoDerivativeInstrAttribField) Tag() quickfix.Tag { return tag.NoDerivativeInstrAttrib }

//NewNoDerivativeInstrAttrib returns a new NoDerivativeInstrAttribField initialized with val
func NewNoDerivativeInstrAttrib(val int) NoDerivativeInstrAttribField {
	return NoDerivativeInstrAttribField{quickfix.FIXInt(val)}
}

func (f NoDerivativeInstrAttribField) Value() int { return f.Int() }

//NoDerivativeInstrumentPartiesField is a NUMINGROUP field
type NoDerivativeInstrumentPartiesField struct{ quickfix.FIXInt }

//Tag returns tag.NoDerivativeInstrumentParties (1292)
func (f NoDerivativeInstrumentPartiesField) Tag() quickfix.Tag {
	return tag.NoDerivativeInstrumentParties
}

//NewNoDerivativeInstrumentParties returns a new NoDerivativeInstrumentPartiesField initialized with val
func NewNoDerivativeInstrumentParties(val int) NoDerivativeInstrumentPartiesField {
	return NoDerivativeInstrumentPartiesField{quickfix.FIXInt(val)}
}

func (f NoDerivativeInstrumentPartiesField) Value() int { return f.Int() }

//NoDerivativeInstrumentPartySubIDsField is a NUMINGROUP field
type NoDerivativeInstrumentPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoDerivativeInstrumentPartySubIDs (1296)
func (f NoDerivativeInstrumentPartySubIDsField) Tag() quickfix.Tag {
	return tag.NoDerivativeInstrumentPartySubIDs
}

//NewNoDerivativeInstrumentPartySubIDs returns a new NoDerivativeInstrumentPartySubIDsField initialized with val
func NewNoDerivativeInstrumentPartySubIDs(val int) NoDerivativeInstrumentPartySubIDsField {
	return NoDerivativeInstrumentPartySubIDsField{quickfix.FIXInt(val)}
}

func (f NoDerivativeInstrumentPartySubIDsField) Value() int { return f.Int() }

//NoDerivativeSecurityAltIDField is a NUMINGROUP field
type NoDerivativeSecurityAltIDField struct{ quickfix.FIXInt }

//Tag returns tag.NoDerivativeSecurityAltID (1218)
func (f NoDerivativeSecurityAltIDField) Tag() quickfix.Tag { return tag.NoDerivativeSecurityAltID }

//NewNoDerivativeSecurityAltID returns a new NoDerivativeSecurityAltIDField initialized with val
func NewNoDerivativeSecurityAltID(val int) NoDerivativeSecurityAltIDField {
	return NoDerivativeSecurityAltIDField{quickfix.FIXInt(val)}
}

func (f NoDerivativeSecurityAltIDField) Value() int { return f.Int() }

//NoDistribInstsField is a NUMINGROUP field
type NoDistribInstsField struct{ quickfix.FIXInt }

//Tag returns tag.NoDistribInsts (510)
func (f NoDistribInstsField) Tag() quickfix.Tag { return tag.NoDistribInsts }

//NewNoDistribInsts returns a new NoDistribInstsField initialized with val
func NewNoDistribInsts(val int) NoDistribInstsField {
	return NoDistribInstsField{quickfix.FIXInt(val)}
}

func (f NoDistribInstsField) Value() int { return f.Int() }

//NoDlvyInstField is a NUMINGROUP field
type NoDlvyInstField struct{ quickfix.FIXInt }

//Tag returns tag.NoDlvyInst (85)
func (f NoDlvyInstField) Tag() quickfix.Tag { return tag.NoDlvyInst }

//NewNoDlvyInst returns a new NoDlvyInstField initialized with val
func NewNoDlvyInst(val int) NoDlvyInstField {
	return NoDlvyInstField{quickfix.FIXInt(val)}
}

func (f NoDlvyInstField) Value() int { return f.Int() }

//NoEventsField is a NUMINGROUP field
type NoEventsField struct{ quickfix.FIXInt }

//Tag returns tag.NoEvents (864)
func (f NoEventsField) Tag() quickfix.Tag { return tag.NoEvents }

//NewNoEvents returns a new NoEventsField initialized with val
func NewNoEvents(val int) NoEventsField {
	return NoEventsField{quickfix.FIXInt(val)}
}

func (f NoEventsField) Value() int { return f.Int() }

//NoExecInstRulesField is a NUMINGROUP field
type NoExecInstRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoExecInstRules (1232)
func (f NoExecInstRulesField) Tag() quickfix.Tag { return tag.NoExecInstRules }

//NewNoExecInstRules returns a new NoExecInstRulesField initialized with val
func NewNoExecInstRules(val int) NoExecInstRulesField {
	return NoExecInstRulesField{quickfix.FIXInt(val)}
}

func (f NoExecInstRulesField) Value() int { return f.Int() }

//NoExecsField is a NUMINGROUP field
type NoExecsField struct{ quickfix.FIXInt }

//Tag returns tag.NoExecs (124)
func (f NoExecsField) Tag() quickfix.Tag { return tag.NoExecs }

//NewNoExecs returns a new NoExecsField initialized with val
func NewNoExecs(val int) NoExecsField {
	return NoExecsField{quickfix.FIXInt(val)}
}

func (f NoExecsField) Value() int { return f.Int() }

//NoExpirationField is a NUMINGROUP field
type NoExpirationField struct{ quickfix.FIXInt }

//Tag returns tag.NoExpiration (981)
func (f NoExpirationField) Tag() quickfix.Tag { return tag.NoExpiration }

//NewNoExpiration returns a new NoExpirationField initialized with val
func NewNoExpiration(val int) NoExpirationField {
	return NoExpirationField{quickfix.FIXInt(val)}
}

func (f NoExpirationField) Value() int { return f.Int() }

//NoFillsField is a NUMINGROUP field
type NoFillsField struct{ quickfix.FIXInt }

//Tag returns tag.NoFills (1362)
func (f NoFillsField) Tag() quickfix.Tag { return tag.NoFills }

//NewNoFills returns a new NoFillsField initialized with val
func NewNoFills(val int) NoFillsField {
	return NoFillsField{quickfix.FIXInt(val)}
}

func (f NoFillsField) Value() int { return f.Int() }

//NoHopsField is a NUMINGROUP field
type NoHopsField struct{ quickfix.FIXInt }

//Tag returns tag.NoHops (627)
func (f NoHopsField) Tag() quickfix.Tag { return tag.NoHops }

//NewNoHops returns a new NoHopsField initialized with val
func NewNoHops(val int) NoHopsField {
	return NoHopsField{quickfix.FIXInt(val)}
}

func (f NoHopsField) Value() int { return f.Int() }

//NoIOIQualifiersField is a NUMINGROUP field
type NoIOIQualifiersField struct{ quickfix.FIXInt }

//Tag returns tag.NoIOIQualifiers (199)
func (f NoIOIQualifiersField) Tag() quickfix.Tag { return tag.NoIOIQualifiers }

//NewNoIOIQualifiers returns a new NoIOIQualifiersField initialized with val
func NewNoIOIQualifiers(val int) NoIOIQualifiersField {
	return NoIOIQualifiersField{quickfix.FIXInt(val)}
}

func (f NoIOIQualifiersField) Value() int { return f.Int() }

//NoInstrAttribField is a NUMINGROUP field
type NoInstrAttribField struct{ quickfix.FIXInt }

//Tag returns tag.NoInstrAttrib (870)
func (f NoInstrAttribField) Tag() quickfix.Tag { return tag.NoInstrAttrib }

//NewNoInstrAttrib returns a new NoInstrAttribField initialized with val
func NewNoInstrAttrib(val int) NoInstrAttribField {
	return NoInstrAttribField{quickfix.FIXInt(val)}
}

func (f NoInstrAttribField) Value() int { return f.Int() }

//NoInstrumentPartiesField is a NUMINGROUP field
type NoInstrumentPartiesField struct{ quickfix.FIXInt }

//Tag returns tag.NoInstrumentParties (1018)
func (f NoInstrumentPartiesField) Tag() quickfix.Tag { return tag.NoInstrumentParties }

//NewNoInstrumentParties returns a new NoInstrumentPartiesField initialized with val
func NewNoInstrumentParties(val int) NoInstrumentPartiesField {
	return NoInstrumentPartiesField{quickfix.FIXInt(val)}
}

func (f NoInstrumentPartiesField) Value() int { return f.Int() }

//NoInstrumentPartySubIDsField is a NUMINGROUP field
type NoInstrumentPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoInstrumentPartySubIDs (1052)
func (f NoInstrumentPartySubIDsField) Tag() quickfix.Tag { return tag.NoInstrumentPartySubIDs }

//NewNoInstrumentPartySubIDs returns a new NoInstrumentPartySubIDsField initialized with val
func NewNoInstrumentPartySubIDs(val int) NoInstrumentPartySubIDsField {
	return NoInstrumentPartySubIDsField{quickfix.FIXInt(val)}
}

func (f NoInstrumentPartySubIDsField) Value() int { return f.Int() }

//NoLegAllocsField is a NUMINGROUP field
type NoLegAllocsField struct{ quickfix.FIXInt }

//Tag returns tag.NoLegAllocs (670)
func (f NoLegAllocsField) Tag() quickfix.Tag { return tag.NoLegAllocs }

//NewNoLegAllocs returns a new NoLegAllocsField initialized with val
func NewNoLegAllocs(val int) NoLegAllocsField {
	return NoLegAllocsField{quickfix.FIXInt(val)}
}

func (f NoLegAllocsField) Value() int { return f.Int() }

//NoLegSecurityAltIDField is a STRING field
type NoLegSecurityAltIDField struct{ quickfix.FIXString }

//Tag returns tag.NoLegSecurityAltID (604)
func (f NoLegSecurityAltIDField) Tag() quickfix.Tag { return tag.NoLegSecurityAltID }

//NewNoLegSecurityAltID returns a new NoLegSecurityAltIDField initialized with val
func NewNoLegSecurityAltID(val string) NoLegSecurityAltIDField {
	return NoLegSecurityAltIDField{quickfix.FIXString(val)}
}

func (f NoLegSecurityAltIDField) Value() string { return f.String() }

//NoLegStipulationsField is a NUMINGROUP field
type NoLegStipulationsField struct{ quickfix.FIXInt }

//Tag returns tag.NoLegStipulations (683)
func (f NoLegStipulationsField) Tag() quickfix.Tag { return tag.NoLegStipulations }

//NewNoLegStipulations returns a new NoLegStipulationsField initialized with val
func NewNoLegStipulations(val int) NoLegStipulationsField {
	return NoLegStipulationsField{quickfix.FIXInt(val)}
}

func (f NoLegStipulationsField) Value() int { return f.Int() }

//NoLegsField is a NUMINGROUP field
type NoLegsField struct{ quickfix.FIXInt }

//Tag returns tag.NoLegs (555)
func (f NoLegsField) Tag() quickfix.Tag { return tag.NoLegs }

//NewNoLegs returns a new NoLegsField initialized with val
func NewNoLegs(val int) NoLegsField {
	return NoLegsField{quickfix.FIXInt(val)}
}

func (f NoLegsField) Value() int { return f.Int() }

//NoLinesOfTextField is a NUMINGROUP field
type NoLinesOfTextField struct{ quickfix.FIXInt }

//Tag returns tag.NoLinesOfText (33)
func (f NoLinesOfTextField) Tag() quickfix.Tag { return tag.NoLinesOfText }

//NewNoLinesOfText returns a new NoLinesOfTextField initialized with val
func NewNoLinesOfText(val int) NoLinesOfTextField {
	return NoLinesOfTextField{quickfix.FIXInt(val)}
}

func (f NoLinesOfTextField) Value() int { return f.Int() }

//NoLotTypeRulesField is a NUMINGROUP field
type NoLotTypeRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoLotTypeRules (1234)
func (f NoLotTypeRulesField) Tag() quickfix.Tag { return tag.NoLotTypeRules }

//NewNoLotTypeRules returns a new NoLotTypeRulesField initialized with val
func NewNoLotTypeRules(val int) NoLotTypeRulesField {
	return NoLotTypeRulesField{quickfix.FIXInt(val)}
}

func (f NoLotTypeRulesField) Value() int { return f.Int() }

//NoMDEntriesField is a NUMINGROUP field
type NoMDEntriesField struct{ quickfix.FIXInt }

//Tag returns tag.NoMDEntries (268)
func (f NoMDEntriesField) Tag() quickfix.Tag { return tag.NoMDEntries }

//NewNoMDEntries returns a new NoMDEntriesField initialized with val
func NewNoMDEntries(val int) NoMDEntriesField {
	return NoMDEntriesField{quickfix.FIXInt(val)}
}

func (f NoMDEntriesField) Value() int { return f.Int() }

//NoMDEntryTypesField is a NUMINGROUP field
type NoMDEntryTypesField struct{ quickfix.FIXInt }

//Tag returns tag.NoMDEntryTypes (267)
func (f NoMDEntryTypesField) Tag() quickfix.Tag { return tag.NoMDEntryTypes }

//NewNoMDEntryTypes returns a new NoMDEntryTypesField initialized with val
func NewNoMDEntryTypes(val int) NoMDEntryTypesField {
	return NoMDEntryTypesField{quickfix.FIXInt(val)}
}

func (f NoMDEntryTypesField) Value() int { return f.Int() }

//NoMDFeedTypesField is a NUMINGROUP field
type NoMDFeedTypesField struct{ quickfix.FIXInt }

//Tag returns tag.NoMDFeedTypes (1141)
func (f NoMDFeedTypesField) Tag() quickfix.Tag { return tag.NoMDFeedTypes }

//NewNoMDFeedTypes returns a new NoMDFeedTypesField initialized with val
func NewNoMDFeedTypes(val int) NoMDFeedTypesField {
	return NoMDFeedTypesField{quickfix.FIXInt(val)}
}

func (f NoMDFeedTypesField) Value() int { return f.Int() }

//NoMarketSegmentsField is a NUMINGROUP field
type NoMarketSegmentsField struct{ quickfix.FIXInt }

//Tag returns tag.NoMarketSegments (1310)
func (f NoMarketSegmentsField) Tag() quickfix.Tag { return tag.NoMarketSegments }

//NewNoMarketSegments returns a new NoMarketSegmentsField initialized with val
func NewNoMarketSegments(val int) NoMarketSegmentsField {
	return NoMarketSegmentsField{quickfix.FIXInt(val)}
}

func (f NoMarketSegmentsField) Value() int { return f.Int() }

//NoMatchRulesField is a NUMINGROUP field
type NoMatchRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoMatchRules (1235)
func (f NoMatchRulesField) Tag() quickfix.Tag { return tag.NoMatchRules }

//NewNoMatchRules returns a new NoMatchRulesField initialized with val
func NewNoMatchRules(val int) NoMatchRulesField {
	return NoMatchRulesField{quickfix.FIXInt(val)}
}

func (f NoMatchRulesField) Value() int { return f.Int() }

//NoMaturityRulesField is a NUMINGROUP field
type NoMaturityRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoMaturityRules (1236)
func (f NoMaturityRulesField) Tag() quickfix.Tag { return tag.NoMaturityRules }

//NewNoMaturityRules returns a new NoMaturityRulesField initialized with val
func NewNoMaturityRules(val int) NoMaturityRulesField {
	return NoMaturityRulesField{quickfix.FIXInt(val)}
}

func (f NoMaturityRulesField) Value() int { return f.Int() }

//NoMiscFeesField is a NUMINGROUP field
type NoMiscFeesField struct{ quickfix.FIXInt }

//Tag returns tag.NoMiscFees (136)
func (f NoMiscFeesField) Tag() quickfix.Tag { return tag.NoMiscFees }

//NewNoMiscFees returns a new NoMiscFeesField initialized with val
func NewNoMiscFees(val int) NoMiscFeesField {
	return NoMiscFeesField{quickfix.FIXInt(val)}
}

func (f NoMiscFeesField) Value() int { return f.Int() }

//NoMsgTypesField is a NUMINGROUP field
type NoMsgTypesField struct{ quickfix.FIXInt }

//Tag returns tag.NoMsgTypes (384)
func (f NoMsgTypesField) Tag() quickfix.Tag { return tag.NoMsgTypes }

//NewNoMsgTypes returns a new NoMsgTypesField initialized with val
func NewNoMsgTypes(val int) NoMsgTypesField {
	return NoMsgTypesField{quickfix.FIXInt(val)}
}

func (f NoMsgTypesField) Value() int { return f.Int() }

//NoNested2PartyIDsField is a NUMINGROUP field
type NoNested2PartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNested2PartyIDs (756)
func (f NoNested2PartyIDsField) Tag() quickfix.Tag { return tag.NoNested2PartyIDs }

//NewNoNested2PartyIDs returns a new NoNested2PartyIDsField initialized with val
func NewNoNested2PartyIDs(val int) NoNested2PartyIDsField {
	return NoNested2PartyIDsField{quickfix.FIXInt(val)}
}

func (f NoNested2PartyIDsField) Value() int { return f.Int() }

//NoNested2PartySubIDsField is a NUMINGROUP field
type NoNested2PartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNested2PartySubIDs (806)
func (f NoNested2PartySubIDsField) Tag() quickfix.Tag { return tag.NoNested2PartySubIDs }

//NewNoNested2PartySubIDs returns a new NoNested2PartySubIDsField initialized with val
func NewNoNested2PartySubIDs(val int) NoNested2PartySubIDsField {
	return NoNested2PartySubIDsField{quickfix.FIXInt(val)}
}

func (f NoNested2PartySubIDsField) Value() int { return f.Int() }

//NoNested3PartyIDsField is a NUMINGROUP field
type NoNested3PartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNested3PartyIDs (948)
func (f NoNested3PartyIDsField) Tag() quickfix.Tag { return tag.NoNested3PartyIDs }

//NewNoNested3PartyIDs returns a new NoNested3PartyIDsField initialized with val
func NewNoNested3PartyIDs(val int) NoNested3PartyIDsField {
	return NoNested3PartyIDsField{quickfix.FIXInt(val)}
}

func (f NoNested3PartyIDsField) Value() int { return f.Int() }

//NoNested3PartySubIDsField is a NUMINGROUP field
type NoNested3PartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNested3PartySubIDs (952)
func (f NoNested3PartySubIDsField) Tag() quickfix.Tag { return tag.NoNested3PartySubIDs }

//NewNoNested3PartySubIDs returns a new NoNested3PartySubIDsField initialized with val
func NewNoNested3PartySubIDs(val int) NoNested3PartySubIDsField {
	return NoNested3PartySubIDsField{quickfix.FIXInt(val)}
}

func (f NoNested3PartySubIDsField) Value() int { return f.Int() }

//NoNested4PartyIDsField is a NUMINGROUP field
type NoNested4PartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNested4PartyIDs (1414)
func (f NoNested4PartyIDsField) Tag() quickfix.Tag { return tag.NoNested4PartyIDs }

//NewNoNested4PartyIDs returns a new NoNested4PartyIDsField initialized with val
func NewNoNested4PartyIDs(val int) NoNested4PartyIDsField {
	return NoNested4PartyIDsField{quickfix.FIXInt(val)}
}

func (f NoNested4PartyIDsField) Value() int { return f.Int() }

//NoNested4PartySubIDsField is a NUMINGROUP field
type NoNested4PartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNested4PartySubIDs (1413)
func (f NoNested4PartySubIDsField) Tag() quickfix.Tag { return tag.NoNested4PartySubIDs }

//NewNoNested4PartySubIDs returns a new NoNested4PartySubIDsField initialized with val
func NewNoNested4PartySubIDs(val int) NoNested4PartySubIDsField {
	return NoNested4PartySubIDsField{quickfix.FIXInt(val)}
}

func (f NoNested4PartySubIDsField) Value() int { return f.Int() }

//NoNestedInstrAttribField is a NUMINGROUP field
type NoNestedInstrAttribField struct{ quickfix.FIXInt }

//Tag returns tag.NoNestedInstrAttrib (1312)
func (f NoNestedInstrAttribField) Tag() quickfix.Tag { return tag.NoNestedInstrAttrib }

//NewNoNestedInstrAttrib returns a new NoNestedInstrAttribField initialized with val
func NewNoNestedInstrAttrib(val int) NoNestedInstrAttribField {
	return NoNestedInstrAttribField{quickfix.FIXInt(val)}
}

func (f NoNestedInstrAttribField) Value() int { return f.Int() }

//NoNestedPartyIDsField is a NUMINGROUP field
type NoNestedPartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNestedPartyIDs (539)
func (f NoNestedPartyIDsField) Tag() quickfix.Tag { return tag.NoNestedPartyIDs }

//NewNoNestedPartyIDs returns a new NoNestedPartyIDsField initialized with val
func NewNoNestedPartyIDs(val int) NoNestedPartyIDsField {
	return NoNestedPartyIDsField{quickfix.FIXInt(val)}
}

func (f NoNestedPartyIDsField) Value() int { return f.Int() }

//NoNestedPartySubIDsField is a NUMINGROUP field
type NoNestedPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNestedPartySubIDs (804)
func (f NoNestedPartySubIDsField) Tag() quickfix.Tag { return tag.NoNestedPartySubIDs }

//NewNoNestedPartySubIDs returns a new NoNestedPartySubIDsField initialized with val
func NewNoNestedPartySubIDs(val int) NoNestedPartySubIDsField {
	return NoNestedPartySubIDsField{quickfix.FIXInt(val)}
}

func (f NoNestedPartySubIDsField) Value() int { return f.Int() }

//NoNewsRefIDsField is a NUMINGROUP field
type NoNewsRefIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNewsRefIDs (1475)
func (f NoNewsRefIDsField) Tag() quickfix.Tag { return tag.NoNewsRefIDs }

//NewNoNewsRefIDs returns a new NoNewsRefIDsField initialized with val
func NewNoNewsRefIDs(val int) NoNewsRefIDsField {
	return NoNewsRefIDsField{quickfix.FIXInt(val)}
}

func (f NoNewsRefIDsField) Value() int { return f.Int() }

//NoNotAffectedOrdersField is a NUMINGROUP field
type NoNotAffectedOrdersField struct{ quickfix.FIXInt }

//Tag returns tag.NoNotAffectedOrders (1370)
func (f NoNotAffectedOrdersField) Tag() quickfix.Tag { return tag.NoNotAffectedOrders }

//NewNoNotAffectedOrders returns a new NoNotAffectedOrdersField initialized with val
func NewNoNotAffectedOrders(val int) NoNotAffectedOrdersField {
	return NoNotAffectedOrdersField{quickfix.FIXInt(val)}
}

func (f NoNotAffectedOrdersField) Value() int { return f.Int() }

//NoOfLegUnderlyingsField is a NUMINGROUP field
type NoOfLegUnderlyingsField struct{ quickfix.FIXInt }

//Tag returns tag.NoOfLegUnderlyings (1342)
func (f NoOfLegUnderlyingsField) Tag() quickfix.Tag { return tag.NoOfLegUnderlyings }

//NewNoOfLegUnderlyings returns a new NoOfLegUnderlyingsField initialized with val
func NewNoOfLegUnderlyings(val int) NoOfLegUnderlyingsField {
	return NoOfLegUnderlyingsField{quickfix.FIXInt(val)}
}

func (f NoOfLegUnderlyingsField) Value() int { return f.Int() }

//NoOfSecSizesField is a NUMINGROUP field
type NoOfSecSizesField struct{ quickfix.FIXInt }

//Tag returns tag.NoOfSecSizes (1177)
func (f NoOfSecSizesField) Tag() quickfix.Tag { return tag.NoOfSecSizes }

//NewNoOfSecSizes returns a new NoOfSecSizesField initialized with val
func NewNoOfSecSizes(val int) NoOfSecSizesField {
	return NoOfSecSizesField{quickfix.FIXInt(val)}
}

func (f NoOfSecSizesField) Value() int { return f.Int() }

//NoOrdTypeRulesField is a NUMINGROUP field
type NoOrdTypeRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoOrdTypeRules (1237)
func (f NoOrdTypeRulesField) Tag() quickfix.Tag { return tag.NoOrdTypeRules }

//NewNoOrdTypeRules returns a new NoOrdTypeRulesField initialized with val
func NewNoOrdTypeRules(val int) NoOrdTypeRulesField {
	return NoOrdTypeRulesField{quickfix.FIXInt(val)}
}

func (f NoOrdTypeRulesField) Value() int { return f.Int() }

//NoOrdersField is a NUMINGROUP field
type NoOrdersField struct{ quickfix.FIXInt }

//Tag returns tag.NoOrders (73)
func (f NoOrdersField) Tag() quickfix.Tag { return tag.NoOrders }

//NewNoOrders returns a new NoOrdersField initialized with val
func NewNoOrders(val int) NoOrdersField {
	return NoOrdersField{quickfix.FIXInt(val)}
}

func (f NoOrdersField) Value() int { return f.Int() }

//NoPartyAltIDsField is a NUMINGROUP field
type NoPartyAltIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoPartyAltIDs (1516)
func (f NoPartyAltIDsField) Tag() quickfix.Tag { return tag.NoPartyAltIDs }

//NewNoPartyAltIDs returns a new NoPartyAltIDsField initialized with val
func NewNoPartyAltIDs(val int) NoPartyAltIDsField {
	return NoPartyAltIDsField{quickfix.FIXInt(val)}
}

func (f NoPartyAltIDsField) Value() int { return f.Int() }

//NoPartyAltSubIDsField is a NUMINGROUP field
type NoPartyAltSubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoPartyAltSubIDs (1519)
func (f NoPartyAltSubIDsField) Tag() quickfix.Tag { return tag.NoPartyAltSubIDs }

//NewNoPartyAltSubIDs returns a new NoPartyAltSubIDsField initialized with val
func NewNoPartyAltSubIDs(val int) NoPartyAltSubIDsField {
	return NoPartyAltSubIDsField{quickfix.FIXInt(val)}
}

func (f NoPartyAltSubIDsField) Value() int { return f.Int() }

//NoPartyIDsField is a NUMINGROUP field
type NoPartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoPartyIDs (453)
func (f NoPartyIDsField) Tag() quickfix.Tag { return tag.NoPartyIDs }

//NewNoPartyIDs returns a new NoPartyIDsField initialized with val
func NewNoPartyIDs(val int) NoPartyIDsField {
	return NoPartyIDsField{quickfix.FIXInt(val)}
}

func (f NoPartyIDsField) Value() int { return f.Int() }

//NoPartyListField is a NUMINGROUP field
type NoPartyListField struct{ quickfix.FIXInt }

//Tag returns tag.NoPartyList (1513)
func (f NoPartyListField) Tag() quickfix.Tag { return tag.NoPartyList }

//NewNoPartyList returns a new NoPartyListField initialized with val
func NewNoPartyList(val int) NoPartyListField {
	return NoPartyListField{quickfix.FIXInt(val)}
}

func (f NoPartyListField) Value() int { return f.Int() }

//NoPartyListResponseTypesField is a NUMINGROUP field
type NoPartyListResponseTypesField struct{ quickfix.FIXInt }

//Tag returns tag.NoPartyListResponseTypes (1506)
func (f NoPartyListResponseTypesField) Tag() quickfix.Tag { return tag.NoPartyListResponseTypes }

//NewNoPartyListResponseTypes returns a new NoPartyListResponseTypesField initialized with val
func NewNoPartyListResponseTypes(val int) NoPartyListResponseTypesField {
	return NoPartyListResponseTypesField{quickfix.FIXInt(val)}
}

func (f NoPartyListResponseTypesField) Value() int { return f.Int() }

//NoPartyRelationshipsField is a NUMINGROUP field
type NoPartyRelationshipsField struct{ quickfix.FIXInt }

//Tag returns tag.NoPartyRelationships (1514)
func (f NoPartyRelationshipsField) Tag() quickfix.Tag { return tag.NoPartyRelationships }

//NewNoPartyRelationships returns a new NoPartyRelationshipsField initialized with val
func NewNoPartyRelationships(val int) NoPartyRelationshipsField {
	return NoPartyRelationshipsField{quickfix.FIXInt(val)}
}

func (f NoPartyRelationshipsField) Value() int { return f.Int() }

//NoPartySubIDsField is a NUMINGROUP field
type NoPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoPartySubIDs (802)
func (f NoPartySubIDsField) Tag() quickfix.Tag { return tag.NoPartySubIDs }

//NewNoPartySubIDs returns a new NoPartySubIDsField initialized with val
func NewNoPartySubIDs(val int) NoPartySubIDsField {
	return NoPartySubIDsField{quickfix.FIXInt(val)}
}

func (f NoPartySubIDsField) Value() int { return f.Int() }

//NoPosAmtField is a NUMINGROUP field
type NoPosAmtField struct{ quickfix.FIXInt }

//Tag returns tag.NoPosAmt (753)
func (f NoPosAmtField) Tag() quickfix.Tag { return tag.NoPosAmt }

//NewNoPosAmt returns a new NoPosAmtField initialized with val
func NewNoPosAmt(val int) NoPosAmtField {
	return NoPosAmtField{quickfix.FIXInt(val)}
}

func (f NoPosAmtField) Value() int { return f.Int() }

//NoPositionsField is a NUMINGROUP field
type NoPositionsField struct{ quickfix.FIXInt }

//Tag returns tag.NoPositions (702)
func (f NoPositionsField) Tag() quickfix.Tag { return tag.NoPositions }

//NewNoPositions returns a new NoPositionsField initialized with val
func NewNoPositions(val int) NoPositionsField {
	return NoPositionsField{quickfix.FIXInt(val)}
}

func (f NoPositionsField) Value() int { return f.Int() }

//NoQuoteEntriesField is a NUMINGROUP field
type NoQuoteEntriesField struct{ quickfix.FIXInt }

//Tag returns tag.NoQuoteEntries (295)
func (f NoQuoteEntriesField) Tag() quickfix.Tag { return tag.NoQuoteEntries }

//NewNoQuoteEntries returns a new NoQuoteEntriesField initialized with val
func NewNoQuoteEntries(val int) NoQuoteEntriesField {
	return NoQuoteEntriesField{quickfix.FIXInt(val)}
}

func (f NoQuoteEntriesField) Value() int { return f.Int() }

//NoQuoteQualifiersField is a NUMINGROUP field
type NoQuoteQualifiersField struct{ quickfix.FIXInt }

//Tag returns tag.NoQuoteQualifiers (735)
func (f NoQuoteQualifiersField) Tag() quickfix.Tag { return tag.NoQuoteQualifiers }

//NewNoQuoteQualifiers returns a new NoQuoteQualifiersField initialized with val
func NewNoQuoteQualifiers(val int) NoQuoteQualifiersField {
	return NoQuoteQualifiersField{quickfix.FIXInt(val)}
}

func (f NoQuoteQualifiersField) Value() int { return f.Int() }

//NoQuoteSetsField is a NUMINGROUP field
type NoQuoteSetsField struct{ quickfix.FIXInt }

//Tag returns tag.NoQuoteSets (296)
func (f NoQuoteSetsField) Tag() quickfix.Tag { return tag.NoQuoteSets }

//NewNoQuoteSets returns a new NoQuoteSetsField initialized with val
func NewNoQuoteSets(val int) NoQuoteSetsField {
	return NoQuoteSetsField{quickfix.FIXInt(val)}
}

func (f NoQuoteSetsField) Value() int { return f.Int() }

//NoRateSourcesField is a NUMINGROUP field
type NoRateSourcesField struct{ quickfix.FIXInt }

//Tag returns tag.NoRateSources (1445)
func (f NoRateSourcesField) Tag() quickfix.Tag { return tag.NoRateSources }

//NewNoRateSources returns a new NoRateSourcesField initialized with val
func NewNoRateSources(val int) NoRateSourcesField {
	return NoRateSourcesField{quickfix.FIXInt(val)}
}

func (f NoRateSourcesField) Value() int { return f.Int() }

//NoRegistDtlsField is a NUMINGROUP field
type NoRegistDtlsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRegistDtls (473)
func (f NoRegistDtlsField) Tag() quickfix.Tag { return tag.NoRegistDtls }

//NewNoRegistDtls returns a new NoRegistDtlsField initialized with val
func NewNoRegistDtls(val int) NoRegistDtlsField {
	return NoRegistDtlsField{quickfix.FIXInt(val)}
}

func (f NoRegistDtlsField) Value() int { return f.Int() }

//NoRelatedContextPartyIDsField is a NUMINGROUP field
type NoRelatedContextPartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelatedContextPartyIDs (1575)
func (f NoRelatedContextPartyIDsField) Tag() quickfix.Tag { return tag.NoRelatedContextPartyIDs }

//NewNoRelatedContextPartyIDs returns a new NoRelatedContextPartyIDsField initialized with val
func NewNoRelatedContextPartyIDs(val int) NoRelatedContextPartyIDsField {
	return NoRelatedContextPartyIDsField{quickfix.FIXInt(val)}
}

func (f NoRelatedContextPartyIDsField) Value() int { return f.Int() }

//NoRelatedContextPartySubIDsField is a NUMINGROUP field
type NoRelatedContextPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelatedContextPartySubIDs (1579)
func (f NoRelatedContextPartySubIDsField) Tag() quickfix.Tag { return tag.NoRelatedContextPartySubIDs }

//NewNoRelatedContextPartySubIDs returns a new NoRelatedContextPartySubIDsField initialized with val
func NewNoRelatedContextPartySubIDs(val int) NoRelatedContextPartySubIDsField {
	return NoRelatedContextPartySubIDsField{quickfix.FIXInt(val)}
}

func (f NoRelatedContextPartySubIDsField) Value() int { return f.Int() }

//NoRelatedPartyAltIDsField is a NUMINGROUP field
type NoRelatedPartyAltIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelatedPartyAltIDs (1569)
func (f NoRelatedPartyAltIDsField) Tag() quickfix.Tag { return tag.NoRelatedPartyAltIDs }

//NewNoRelatedPartyAltIDs returns a new NoRelatedPartyAltIDsField initialized with val
func NewNoRelatedPartyAltIDs(val int) NoRelatedPartyAltIDsField {
	return NoRelatedPartyAltIDsField{quickfix.FIXInt(val)}
}

func (f NoRelatedPartyAltIDsField) Value() int { return f.Int() }

//NoRelatedPartyAltSubIDsField is a NUMINGROUP field
type NoRelatedPartyAltSubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelatedPartyAltSubIDs (1572)
func (f NoRelatedPartyAltSubIDsField) Tag() quickfix.Tag { return tag.NoRelatedPartyAltSubIDs }

//NewNoRelatedPartyAltSubIDs returns a new NoRelatedPartyAltSubIDsField initialized with val
func NewNoRelatedPartyAltSubIDs(val int) NoRelatedPartyAltSubIDsField {
	return NoRelatedPartyAltSubIDsField{quickfix.FIXInt(val)}
}

func (f NoRelatedPartyAltSubIDsField) Value() int { return f.Int() }

//NoRelatedPartyIDsField is a NUMINGROUP field
type NoRelatedPartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelatedPartyIDs (1562)
func (f NoRelatedPartyIDsField) Tag() quickfix.Tag { return tag.NoRelatedPartyIDs }

//NewNoRelatedPartyIDs returns a new NoRelatedPartyIDsField initialized with val
func NewNoRelatedPartyIDs(val int) NoRelatedPartyIDsField {
	return NoRelatedPartyIDsField{quickfix.FIXInt(val)}
}

func (f NoRelatedPartyIDsField) Value() int { return f.Int() }

//NoRelatedPartySubIDsField is a NUMINGROUP field
type NoRelatedPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelatedPartySubIDs (1566)
func (f NoRelatedPartySubIDsField) Tag() quickfix.Tag { return tag.NoRelatedPartySubIDs }

//NewNoRelatedPartySubIDs returns a new NoRelatedPartySubIDsField initialized with val
func NewNoRelatedPartySubIDs(val int) NoRelatedPartySubIDsField {
	return NoRelatedPartySubIDsField{quickfix.FIXInt(val)}
}

func (f NoRelatedPartySubIDsField) Value() int { return f.Int() }

//NoRelatedSymField is a NUMINGROUP field
type NoRelatedSymField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelatedSym (146)
func (f NoRelatedSymField) Tag() quickfix.Tag { return tag.NoRelatedSym }

//NewNoRelatedSym returns a new NoRelatedSymField initialized with val
func NewNoRelatedSym(val int) NoRelatedSymField {
	return NoRelatedSymField{quickfix.FIXInt(val)}
}

func (f NoRelatedSymField) Value() int { return f.Int() }

//NoRelationshipRiskInstrumentsField is a NUMINGROUP field
type NoRelationshipRiskInstrumentsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelationshipRiskInstruments (1587)
func (f NoRelationshipRiskInstrumentsField) Tag() quickfix.Tag {
	return tag.NoRelationshipRiskInstruments
}

//NewNoRelationshipRiskInstruments returns a new NoRelationshipRiskInstrumentsField initialized with val
func NewNoRelationshipRiskInstruments(val int) NoRelationshipRiskInstrumentsField {
	return NoRelationshipRiskInstrumentsField{quickfix.FIXInt(val)}
}

func (f NoRelationshipRiskInstrumentsField) Value() int { return f.Int() }

//NoRelationshipRiskLimitsField is a NUMINGROUP field
type NoRelationshipRiskLimitsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelationshipRiskLimits (1582)
func (f NoRelationshipRiskLimitsField) Tag() quickfix.Tag { return tag.NoRelationshipRiskLimits }

//NewNoRelationshipRiskLimits returns a new NoRelationshipRiskLimitsField initialized with val
func NewNoRelationshipRiskLimits(val int) NoRelationshipRiskLimitsField {
	return NoRelationshipRiskLimitsField{quickfix.FIXInt(val)}
}

func (f NoRelationshipRiskLimitsField) Value() int { return f.Int() }

//NoRelationshipRiskSecurityAltIDField is a NUMINGROUP field
type NoRelationshipRiskSecurityAltIDField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelationshipRiskSecurityAltID (1593)
func (f NoRelationshipRiskSecurityAltIDField) Tag() quickfix.Tag {
	return tag.NoRelationshipRiskSecurityAltID
}

//NewNoRelationshipRiskSecurityAltID returns a new NoRelationshipRiskSecurityAltIDField initialized with val
func NewNoRelationshipRiskSecurityAltID(val int) NoRelationshipRiskSecurityAltIDField {
	return NoRelationshipRiskSecurityAltIDField{quickfix.FIXInt(val)}
}

func (f NoRelationshipRiskSecurityAltIDField) Value() int { return f.Int() }

//NoRelationshipRiskWarningLevelsField is a NUMINGROUP field
type NoRelationshipRiskWarningLevelsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelationshipRiskWarningLevels (1613)
func (f NoRelationshipRiskWarningLevelsField) Tag() quickfix.Tag {
	return tag.NoRelationshipRiskWarningLevels
}

//NewNoRelationshipRiskWarningLevels returns a new NoRelationshipRiskWarningLevelsField initialized with val
func NewNoRelationshipRiskWarningLevels(val int) NoRelationshipRiskWarningLevelsField {
	return NoRelationshipRiskWarningLevelsField{quickfix.FIXInt(val)}
}

func (f NoRelationshipRiskWarningLevelsField) Value() int { return f.Int() }

//NoRequestedPartyRolesField is a NUMINGROUP field
type NoRequestedPartyRolesField struct{ quickfix.FIXInt }

//Tag returns tag.NoRequestedPartyRoles (1508)
func (f NoRequestedPartyRolesField) Tag() quickfix.Tag { return tag.NoRequestedPartyRoles }

//NewNoRequestedPartyRoles returns a new NoRequestedPartyRolesField initialized with val
func NewNoRequestedPartyRoles(val int) NoRequestedPartyRolesField {
	return NoRequestedPartyRolesField{quickfix.FIXInt(val)}
}

func (f NoRequestedPartyRolesField) Value() int { return f.Int() }

//NoRiskInstrumentsField is a NUMINGROUP field
type NoRiskInstrumentsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRiskInstruments (1534)
func (f NoRiskInstrumentsField) Tag() quickfix.Tag { return tag.NoRiskInstruments }

//NewNoRiskInstruments returns a new NoRiskInstrumentsField initialized with val
func NewNoRiskInstruments(val int) NoRiskInstrumentsField {
	return NoRiskInstrumentsField{quickfix.FIXInt(val)}
}

func (f NoRiskInstrumentsField) Value() int { return f.Int() }

//NoRiskLimitsField is a NUMINGROUP field
type NoRiskLimitsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRiskLimits (1529)
func (f NoRiskLimitsField) Tag() quickfix.Tag { return tag.NoRiskLimits }

//NewNoRiskLimits returns a new NoRiskLimitsField initialized with val
func NewNoRiskLimits(val int) NoRiskLimitsField {
	return NoRiskLimitsField{quickfix.FIXInt(val)}
}

func (f NoRiskLimitsField) Value() int { return f.Int() }

//NoRiskSecurityAltIDField is a NUMINGROUP field
type NoRiskSecurityAltIDField struct{ quickfix.FIXInt }

//Tag returns tag.NoRiskSecurityAltID (1540)
func (f NoRiskSecurityAltIDField) Tag() quickfix.Tag { return tag.NoRiskSecurityAltID }

//NewNoRiskSecurityAltID returns a new NoRiskSecurityAltIDField initialized with val
func NewNoRiskSecurityAltID(val int) NoRiskSecurityAltIDField {
	return NoRiskSecurityAltIDField{quickfix.FIXInt(val)}
}

func (f NoRiskSecurityAltIDField) Value() int { return f.Int() }

//NoRiskWarningLevelsField is a NUMINGROUP field
type NoRiskWarningLevelsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRiskWarningLevels (1559)
func (f NoRiskWarningLevelsField) Tag() quickfix.Tag { return tag.NoRiskWarningLevels }

//NewNoRiskWarningLevels returns a new NoRiskWarningLevelsField initialized with val
func NewNoRiskWarningLevels(val int) NoRiskWarningLevelsField {
	return NoRiskWarningLevelsField{quickfix.FIXInt(val)}
}

func (f NoRiskWarningLevelsField) Value() int { return f.Int() }

//NoRootPartyIDsField is a NUMINGROUP field
type NoRootPartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRootPartyIDs (1116)
func (f NoRootPartyIDsField) Tag() quickfix.Tag { return tag.NoRootPartyIDs }

//NewNoRootPartyIDs returns a new NoRootPartyIDsField initialized with val
func NewNoRootPartyIDs(val int) NoRootPartyIDsField {
	return NoRootPartyIDsField{quickfix.FIXInt(val)}
}

func (f NoRootPartyIDsField) Value() int { return f.Int() }

//NoRootPartySubIDsField is a NUMINGROUP field
type NoRootPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRootPartySubIDs (1120)
func (f NoRootPartySubIDsField) Tag() quickfix.Tag { return tag.NoRootPartySubIDs }

//NewNoRootPartySubIDs returns a new NoRootPartySubIDsField initialized with val
func NewNoRootPartySubIDs(val int) NoRootPartySubIDsField {
	return NoRootPartySubIDsField{quickfix.FIXInt(val)}
}

func (f NoRootPartySubIDsField) Value() int { return f.Int() }

//NoRoutingIDsField is a NUMINGROUP field
type NoRoutingIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRoutingIDs (215)
func (f NoRoutingIDsField) Tag() quickfix.Tag { return tag.NoRoutingIDs }

//NewNoRoutingIDs returns a new NoRoutingIDsField initialized with val
func NewNoRoutingIDs(val int) NoRoutingIDsField {
	return NoRoutingIDsField{quickfix.FIXInt(val)}
}

func (f NoRoutingIDsField) Value() int { return f.Int() }

//NoRptsField is a INT field
type NoRptsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRpts (82)
func (f NoRptsField) Tag() quickfix.Tag { return tag.NoRpts }

//NewNoRpts returns a new NoRptsField initialized with val
func NewNoRpts(val int) NoRptsField {
	return NoRptsField{quickfix.FIXInt(val)}
}

func (f NoRptsField) Value() int { return f.Int() }

//NoSecurityAltIDField is a NUMINGROUP field
type NoSecurityAltIDField struct{ quickfix.FIXInt }

//Tag returns tag.NoSecurityAltID (454)
func (f NoSecurityAltIDField) Tag() quickfix.Tag { return tag.NoSecurityAltID }

//NewNoSecurityAltID returns a new NoSecurityAltIDField initialized with val
func NewNoSecurityAltID(val int) NoSecurityAltIDField {
	return NoSecurityAltIDField{quickfix.FIXInt(val)}
}

func (f NoSecurityAltIDField) Value() int { return f.Int() }

//NoSecurityTypesField is a NUMINGROUP field
type NoSecurityTypesField struct{ quickfix.FIXInt }

//Tag returns tag.NoSecurityTypes (558)
func (f NoSecurityTypesField) Tag() quickfix.Tag { return tag.NoSecurityTypes }

//NewNoSecurityTypes returns a new NoSecurityTypesField initialized with val
func NewNoSecurityTypes(val int) NoSecurityTypesField {
	return NoSecurityTypesField{quickfix.FIXInt(val)}
}

func (f NoSecurityTypesField) Value() int { return f.Int() }

//NoSettlDetailsField is a NUMINGROUP field
type NoSettlDetailsField struct{ quickfix.FIXInt }

//Tag returns tag.NoSettlDetails (1158)
func (f NoSettlDetailsField) Tag() quickfix.Tag { return tag.NoSettlDetails }

//NewNoSettlDetails returns a new NoSettlDetailsField initialized with val
func NewNoSettlDetails(val int) NoSettlDetailsField {
	return NoSettlDetailsField{quickfix.FIXInt(val)}
}

func (f NoSettlDetailsField) Value() int { return f.Int() }

//NoSettlInstField is a NUMINGROUP field
type NoSettlInstField struct{ quickfix.FIXInt }

//Tag returns tag.NoSettlInst (778)
func (f NoSettlInstField) Tag() quickfix.Tag { return tag.NoSettlInst }

//NewNoSettlInst returns a new NoSettlInstField initialized with val
func NewNoSettlInst(val int) NoSettlInstField {
	return NoSettlInstField{quickfix.FIXInt(val)}
}

func (f NoSettlInstField) Value() int { return f.Int() }

//NoSettlObligField is a NUMINGROUP field
type NoSettlObligField struct{ quickfix.FIXInt }

//Tag returns tag.NoSettlOblig (1165)
func (f NoSettlObligField) Tag() quickfix.Tag { return tag.NoSettlOblig }

//NewNoSettlOblig returns a new NoSettlObligField initialized with val
func NewNoSettlOblig(val int) NoSettlObligField {
	return NoSettlObligField{quickfix.FIXInt(val)}
}

func (f NoSettlObligField) Value() int { return f.Int() }

//NoSettlPartyIDsField is a NUMINGROUP field
type NoSettlPartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoSettlPartyIDs (781)
func (f NoSettlPartyIDsField) Tag() quickfix.Tag { return tag.NoSettlPartyIDs }

//NewNoSettlPartyIDs returns a new NoSettlPartyIDsField initialized with val
func NewNoSettlPartyIDs(val int) NoSettlPartyIDsField {
	return NoSettlPartyIDsField{quickfix.FIXInt(val)}
}

func (f NoSettlPartyIDsField) Value() int { return f.Int() }

//NoSettlPartySubIDsField is a NUMINGROUP field
type NoSettlPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoSettlPartySubIDs (801)
func (f NoSettlPartySubIDsField) Tag() quickfix.Tag { return tag.NoSettlPartySubIDs }

//NewNoSettlPartySubIDs returns a new NoSettlPartySubIDsField initialized with val
func NewNoSettlPartySubIDs(val int) NoSettlPartySubIDsField {
	return NoSettlPartySubIDsField{quickfix.FIXInt(val)}
}

func (f NoSettlPartySubIDsField) Value() int { return f.Int() }

//NoSideTrdRegTSField is a NUMINGROUP field
type NoSideTrdRegTSField struct{ quickfix.FIXInt }

//Tag returns tag.NoSideTrdRegTS (1016)
func (f NoSideTrdRegTSField) Tag() quickfix.Tag { return tag.NoSideTrdRegTS }

//NewNoSideTrdRegTS returns a new NoSideTrdRegTSField initialized with val
func NewNoSideTrdRegTS(val int) NoSideTrdRegTSField {
	return NoSideTrdRegTSField{quickfix.FIXInt(val)}
}

func (f NoSideTrdRegTSField) Value() int { return f.Int() }

//NoSidesField is a enum.NoSides field
type NoSidesField struct{ quickfix.FIXString }

//Tag returns tag.NoSides (552)
func (f NoSidesField) Tag() quickfix.Tag { return tag.NoSides }

func NewNoSides(val enum.NoSides) NoSidesField {
	return NoSidesField{quickfix.FIXString(val)}
}

func (f NoSidesField) Value() enum.NoSides { return enum.NoSides(f.String()) }

//NoStatsIndicatorsField is a NUMINGROUP field
type NoStatsIndicatorsField struct{ quickfix.FIXInt }

//Tag returns tag.NoStatsIndicators (1175)
func (f NoStatsIndicatorsField) Tag() quickfix.Tag { return tag.NoStatsIndicators }

//NewNoStatsIndicators returns a new NoStatsIndicatorsField initialized with val
func NewNoStatsIndicators(val int) NoStatsIndicatorsField {
	return NoStatsIndicatorsField{quickfix.FIXInt(val)}
}

func (f NoStatsIndicatorsField) Value() int { return f.Int() }

//NoStipulationsField is a NUMINGROUP field
type NoStipulationsField struct{ quickfix.FIXInt }

//Tag returns tag.NoStipulations (232)
func (f NoStipulationsField) Tag() quickfix.Tag { return tag.NoStipulations }

//NewNoStipulations returns a new NoStipulationsField initialized with val
func NewNoStipulations(val int) NoStipulationsField {
	return NoStipulationsField{quickfix.FIXInt(val)}
}

func (f NoStipulationsField) Value() int { return f.Int() }

//NoStrategyParametersField is a NUMINGROUP field
type NoStrategyParametersField struct{ quickfix.FIXInt }

//Tag returns tag.NoStrategyParameters (957)
func (f NoStrategyParametersField) Tag() quickfix.Tag { return tag.NoStrategyParameters }

//NewNoStrategyParameters returns a new NoStrategyParametersField initialized with val
func NewNoStrategyParameters(val int) NoStrategyParametersField {
	return NoStrategyParametersField{quickfix.FIXInt(val)}
}

func (f NoStrategyParametersField) Value() int { return f.Int() }

//NoStrikeRulesField is a NUMINGROUP field
type NoStrikeRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoStrikeRules (1201)
func (f NoStrikeRulesField) Tag() quickfix.Tag { return tag.NoStrikeRules }

//NewNoStrikeRules returns a new NoStrikeRulesField initialized with val
func NewNoStrikeRules(val int) NoStrikeRulesField {
	return NoStrikeRulesField{quickfix.FIXInt(val)}
}

func (f NoStrikeRulesField) Value() int { return f.Int() }

//NoStrikesField is a NUMINGROUP field
type NoStrikesField struct{ quickfix.FIXInt }

//Tag returns tag.NoStrikes (428)
func (f NoStrikesField) Tag() quickfix.Tag { return tag.NoStrikes }

//NewNoStrikes returns a new NoStrikesField initialized with val
func NewNoStrikes(val int) NoStrikesField {
	return NoStrikesField{quickfix.FIXInt(val)}
}

func (f NoStrikesField) Value() int { return f.Int() }

//NoTargetPartyIDsField is a NUMINGROUP field
type NoTargetPartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoTargetPartyIDs (1461)
func (f NoTargetPartyIDsField) Tag() quickfix.Tag { return tag.NoTargetPartyIDs }

//NewNoTargetPartyIDs returns a new NoTargetPartyIDsField initialized with val
func NewNoTargetPartyIDs(val int) NoTargetPartyIDsField {
	return NoTargetPartyIDsField{quickfix.FIXInt(val)}
}

func (f NoTargetPartyIDsField) Value() int { return f.Int() }

//NoTickRulesField is a NUMINGROUP field
type NoTickRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoTickRules (1205)
func (f NoTickRulesField) Tag() quickfix.Tag { return tag.NoTickRules }

//NewNoTickRules returns a new NoTickRulesField initialized with val
func NewNoTickRules(val int) NoTickRulesField {
	return NoTickRulesField{quickfix.FIXInt(val)}
}

func (f NoTickRulesField) Value() int { return f.Int() }

//NoTimeInForceRulesField is a NUMINGROUP field
type NoTimeInForceRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoTimeInForceRules (1239)
func (f NoTimeInForceRulesField) Tag() quickfix.Tag { return tag.NoTimeInForceRules }

//NewNoTimeInForceRules returns a new NoTimeInForceRulesField initialized with val
func NewNoTimeInForceRules(val int) NoTimeInForceRulesField {
	return NoTimeInForceRulesField{quickfix.FIXInt(val)}
}

func (f NoTimeInForceRulesField) Value() int { return f.Int() }

//NoTradesField is a NUMINGROUP field
type NoTradesField struct{ quickfix.FIXInt }

//Tag returns tag.NoTrades (897)
func (f NoTradesField) Tag() quickfix.Tag { return tag.NoTrades }

//NewNoTrades returns a new NoTradesField initialized with val
func NewNoTrades(val int) NoTradesField {
	return NoTradesField{quickfix.FIXInt(val)}
}

func (f NoTradesField) Value() int { return f.Int() }

//NoTradingSessionRulesField is a NUMINGROUP field
type NoTradingSessionRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoTradingSessionRules (1309)
func (f NoTradingSessionRulesField) Tag() quickfix.Tag { return tag.NoTradingSessionRules }

//NewNoTradingSessionRules returns a new NoTradingSessionRulesField initialized with val
func NewNoTradingSessionRules(val int) NoTradingSessionRulesField {
	return NoTradingSessionRulesField{quickfix.FIXInt(val)}
}

func (f NoTradingSessionRulesField) Value() int { return f.Int() }

//NoTradingSessionsField is a NUMINGROUP field
type NoTradingSessionsField struct{ quickfix.FIXInt }

//Tag returns tag.NoTradingSessions (386)
func (f NoTradingSessionsField) Tag() quickfix.Tag { return tag.NoTradingSessions }

//NewNoTradingSessions returns a new NoTradingSessionsField initialized with val
func NewNoTradingSessions(val int) NoTradingSessionsField {
	return NoTradingSessionsField{quickfix.FIXInt(val)}
}

func (f NoTradingSessionsField) Value() int { return f.Int() }

//NoTrdRegTimestampsField is a NUMINGROUP field
type NoTrdRegTimestampsField struct{ quickfix.FIXInt }

//Tag returns tag.NoTrdRegTimestamps (768)
func (f NoTrdRegTimestampsField) Tag() quickfix.Tag { return tag.NoTrdRegTimestamps }

//NewNoTrdRegTimestamps returns a new NoTrdRegTimestampsField initialized with val
func NewNoTrdRegTimestamps(val int) NoTrdRegTimestampsField {
	return NoTrdRegTimestampsField{quickfix.FIXInt(val)}
}

func (f NoTrdRegTimestampsField) Value() int { return f.Int() }

//NoTrdRepIndicatorsField is a NUMINGROUP field
type NoTrdRepIndicatorsField struct{ quickfix.FIXInt }

//Tag returns tag.NoTrdRepIndicators (1387)
func (f NoTrdRepIndicatorsField) Tag() quickfix.Tag { return tag.NoTrdRepIndicators }

//NewNoTrdRepIndicators returns a new NoTrdRepIndicatorsField initialized with val
func NewNoTrdRepIndicators(val int) NoTrdRepIndicatorsField {
	return NoTrdRepIndicatorsField{quickfix.FIXInt(val)}
}

func (f NoTrdRepIndicatorsField) Value() int { return f.Int() }

//NoUnderlyingAmountsField is a NUMINGROUP field
type NoUnderlyingAmountsField struct{ quickfix.FIXInt }

//Tag returns tag.NoUnderlyingAmounts (984)
func (f NoUnderlyingAmountsField) Tag() quickfix.Tag { return tag.NoUnderlyingAmounts }

//NewNoUnderlyingAmounts returns a new NoUnderlyingAmountsField initialized with val
func NewNoUnderlyingAmounts(val int) NoUnderlyingAmountsField {
	return NoUnderlyingAmountsField{quickfix.FIXInt(val)}
}

func (f NoUnderlyingAmountsField) Value() int { return f.Int() }

//NoUnderlyingLegSecurityAltIDField is a NUMINGROUP field
type NoUnderlyingLegSecurityAltIDField struct{ quickfix.FIXInt }

//Tag returns tag.NoUnderlyingLegSecurityAltID (1334)
func (f NoUnderlyingLegSecurityAltIDField) Tag() quickfix.Tag { return tag.NoUnderlyingLegSecurityAltID }

//NewNoUnderlyingLegSecurityAltID returns a new NoUnderlyingLegSecurityAltIDField initialized with val
func NewNoUnderlyingLegSecurityAltID(val int) NoUnderlyingLegSecurityAltIDField {
	return NoUnderlyingLegSecurityAltIDField{quickfix.FIXInt(val)}
}

func (f NoUnderlyingLegSecurityAltIDField) Value() int { return f.Int() }

//NoUnderlyingSecurityAltIDField is a NUMINGROUP field
type NoUnderlyingSecurityAltIDField struct{ quickfix.FIXInt }

//Tag returns tag.NoUnderlyingSecurityAltID (457)
func (f NoUnderlyingSecurityAltIDField) Tag() quickfix.Tag { return tag.NoUnderlyingSecurityAltID }

//NewNoUnderlyingSecurityAltID returns a new NoUnderlyingSecurityAltIDField initialized with val
func NewNoUnderlyingSecurityAltID(val int) NoUnderlyingSecurityAltIDField {
	return NoUnderlyingSecurityAltIDField{quickfix.FIXInt(val)}
}

func (f NoUnderlyingSecurityAltIDField) Value() int { return f.Int() }

//NoUnderlyingStipsField is a NUMINGROUP field
type NoUnderlyingStipsField struct{ quickfix.FIXInt }

//Tag returns tag.NoUnderlyingStips (887)
func (f NoUnderlyingStipsField) Tag() quickfix.Tag { return tag.NoUnderlyingStips }

//NewNoUnderlyingStips returns a new NoUnderlyingStipsField initialized with val
func NewNoUnderlyingStips(val int) NoUnderlyingStipsField {
	return NoUnderlyingStipsField{quickfix.FIXInt(val)}
}

func (f NoUnderlyingStipsField) Value() int { return f.Int() }

//NoUnderlyingsField is a NUMINGROUP field
type NoUnderlyingsField struct{ quickfix.FIXInt }

//Tag returns tag.NoUnderlyings (711)
func (f NoUnderlyingsField) Tag() quickfix.Tag { return tag.NoUnderlyings }

//NewNoUnderlyings returns a new NoUnderlyingsField initialized with val
func NewNoUnderlyings(val int) NoUnderlyingsField {
	return NoUnderlyingsField{quickfix.FIXInt(val)}
}

func (f NoUnderlyingsField) Value() int { return f.Int() }

//NoUndlyInstrumentPartiesField is a NUMINGROUP field
type NoUndlyInstrumentPartiesField struct{ quickfix.FIXInt }

//Tag returns tag.NoUndlyInstrumentParties (1058)
func (f NoUndlyInstrumentPartiesField) Tag() quickfix.Tag { return tag.NoUndlyInstrumentParties }

//NewNoUndlyInstrumentParties returns a new NoUndlyInstrumentPartiesField initialized with val
func NewNoUndlyInstrumentParties(val int) NoUndlyInstrumentPartiesField {
	return NoUndlyInstrumentPartiesField{quickfix.FIXInt(val)}
}

func (f NoUndlyInstrumentPartiesField) Value() int { return f.Int() }

//NoUndlyInstrumentPartySubIDsField is a NUMINGROUP field
type NoUndlyInstrumentPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoUndlyInstrumentPartySubIDs (1062)
func (f NoUndlyInstrumentPartySubIDsField) Tag() quickfix.Tag { return tag.NoUndlyInstrumentPartySubIDs }

//NewNoUndlyInstrumentPartySubIDs returns a new NoUndlyInstrumentPartySubIDsField initialized with val
func NewNoUndlyInstrumentPartySubIDs(val int) NoUndlyInstrumentPartySubIDsField {
	return NoUndlyInstrumentPartySubIDsField{quickfix.FIXInt(val)}
}

func (f NoUndlyInstrumentPartySubIDsField) Value() int { return f.Int() }

//NotAffOrigClOrdIDField is a STRING field
type NotAffOrigClOrdIDField struct{ quickfix.FIXString }

//Tag returns tag.NotAffOrigClOrdID (1372)
func (f NotAffOrigClOrdIDField) Tag() quickfix.Tag { return tag.NotAffOrigClOrdID }

//NewNotAffOrigClOrdID returns a new NotAffOrigClOrdIDField initialized with val
func NewNotAffOrigClOrdID(val string) NotAffOrigClOrdIDField {
	return NotAffOrigClOrdIDField{quickfix.FIXString(val)}
}

func (f NotAffOrigClOrdIDField) Value() string { return f.String() }

//NotAffectedOrderIDField is a STRING field
type NotAffectedOrderIDField struct{ quickfix.FIXString }

//Tag returns tag.NotAffectedOrderID (1371)
func (f NotAffectedOrderIDField) Tag() quickfix.Tag { return tag.NotAffectedOrderID }

//NewNotAffectedOrderID returns a new NotAffectedOrderIDField initialized with val
func NewNotAffectedOrderID(val string) NotAffectedOrderIDField {
	return NotAffectedOrderIDField{quickfix.FIXString(val)}
}

func (f NotAffectedOrderIDField) Value() string { return f.String() }

//NotifyBrokerOfCreditField is a BOOLEAN field
type NotifyBrokerOfCreditField struct{ quickfix.FIXBoolean }

//Tag returns tag.NotifyBrokerOfCredit (208)
func (f NotifyBrokerOfCreditField) Tag() quickfix.Tag { return tag.NotifyBrokerOfCredit }

//NewNotifyBrokerOfCredit returns a new NotifyBrokerOfCreditField initialized with val
func NewNotifyBrokerOfCredit(val bool) NotifyBrokerOfCreditField {
	return NotifyBrokerOfCreditField{quickfix.FIXBoolean(val)}
}

func (f NotifyBrokerOfCreditField) Value() bool { return f.Bool() }

//NotionalPercentageOutstandingField is a PERCENTAGE field
type NotionalPercentageOutstandingField struct{ quickfix.FIXDecimal }

//Tag returns tag.NotionalPercentageOutstanding (1451)
func (f NotionalPercentageOutstandingField) Tag() quickfix.Tag {
	return tag.NotionalPercentageOutstanding
}

//NewNotionalPercentageOutstanding returns a new NotionalPercentageOutstandingField initialized with val and scale
func NewNotionalPercentageOutstanding(val decimal.Decimal, scale int32) NotionalPercentageOutstandingField {
	return NotionalPercentageOutstandingField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f NotionalPercentageOutstandingField) Value() (val decimal.Decimal) { return f.Decimal }

//NumBiddersField is a INT field
type NumBiddersField struct{ quickfix.FIXInt }

//Tag returns tag.NumBidders (417)
func (f NumBiddersField) Tag() quickfix.Tag { return tag.NumBidders }

//NewNumBidders returns a new NumBiddersField initialized with val
func NewNumBidders(val int) NumBiddersField {
	return NumBiddersField{quickfix.FIXInt(val)}
}

func (f NumBiddersField) Value() int { return f.Int() }

//NumDaysInterestField is a INT field
type NumDaysInterestField struct{ quickfix.FIXInt }

//Tag returns tag.NumDaysInterest (157)
func (f NumDaysInterestField) Tag() quickfix.Tag { return tag.NumDaysInterest }

//NewNumDaysInterest returns a new NumDaysInterestField initialized with val
func NewNumDaysInterest(val int) NumDaysInterestField {
	return NumDaysInterestField{quickfix.FIXInt(val)}
}

func (f NumDaysInterestField) Value() int { return f.Int() }

//NumTicketsField is a INT field
type NumTicketsField struct{ quickfix.FIXInt }

//Tag returns tag.NumTickets (395)
func (f NumTicketsField) Tag() quickfix.Tag { return tag.NumTickets }

//NewNumTickets returns a new NumTicketsField initialized with val
func NewNumTickets(val int) NumTicketsField {
	return NumTicketsField{quickfix.FIXInt(val)}
}

func (f NumTicketsField) Value() int { return f.Int() }

//NumberOfOrdersField is a INT field
type NumberOfOrdersField struct{ quickfix.FIXInt }

//Tag returns tag.NumberOfOrders (346)
func (f NumberOfOrdersField) Tag() quickfix.Tag { return tag.NumberOfOrders }

//NewNumberOfOrders returns a new NumberOfOrdersField initialized with val
func NewNumberOfOrders(val int) NumberOfOrdersField {
	return NumberOfOrdersField{quickfix.FIXInt(val)}
}

func (f NumberOfOrdersField) Value() int { return f.Int() }

//OddLotField is a BOOLEAN field
type OddLotField struct{ quickfix.FIXBoolean }

//Tag returns tag.OddLot (575)
func (f OddLotField) Tag() quickfix.Tag { return tag.OddLot }

//NewOddLot returns a new OddLotField initialized with val
func NewOddLot(val bool) OddLotField {
	return OddLotField{quickfix.FIXBoolean(val)}
}

func (f OddLotField) Value() bool { return f.Bool() }

//OfferForwardPointsField is a PRICEOFFSET field
type OfferForwardPointsField struct{ quickfix.FIXDecimal }

//Tag returns tag.OfferForwardPoints (191)
func (f OfferForwardPointsField) Tag() quickfix.Tag { return tag.OfferForwardPoints }

//NewOfferForwardPoints returns a new OfferForwardPointsField initialized with val and scale
func NewOfferForwardPoints(val decimal.Decimal, scale int32) OfferForwardPointsField {
	return OfferForwardPointsField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OfferForwardPointsField) Value() (val decimal.Decimal) { return f.Decimal }

//OfferForwardPoints2Field is a PRICEOFFSET field
type OfferForwardPoints2Field struct{ quickfix.FIXDecimal }

//Tag returns tag.OfferForwardPoints2 (643)
func (f OfferForwardPoints2Field) Tag() quickfix.Tag { return tag.OfferForwardPoints2 }

//NewOfferForwardPoints2 returns a new OfferForwardPoints2Field initialized with val and scale
func NewOfferForwardPoints2(val decimal.Decimal, scale int32) OfferForwardPoints2Field {
	return OfferForwardPoints2Field{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OfferForwardPoints2Field) Value() (val decimal.Decimal) { return f.Decimal }

//OfferPxField is a PRICE field
type OfferPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.OfferPx (133)
func (f OfferPxField) Tag() quickfix.Tag { return tag.OfferPx }

//NewOfferPx returns a new OfferPxField initialized with val and scale
func NewOfferPx(val decimal.Decimal, scale int32) OfferPxField {
	return OfferPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OfferPxField) Value() (val decimal.Decimal) { return f.Decimal }

//OfferSizeField is a QTY field
type OfferSizeField struct{ quickfix.FIXDecimal }

//Tag returns tag.OfferSize (135)
func (f OfferSizeField) Tag() quickfix.Tag { return tag.OfferSize }

//NewOfferSize returns a new OfferSizeField initialized with val and scale
func NewOfferSize(val decimal.Decimal, scale int32) OfferSizeField {
	return OfferSizeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OfferSizeField) Value() (val decimal.Decimal) { return f.Decimal }

//OfferSpotRateField is a PRICE field
type OfferSpotRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.OfferSpotRate (190)
func (f OfferSpotRateField) Tag() quickfix.Tag { return tag.OfferSpotRate }

//NewOfferSpotRate returns a new OfferSpotRateField initialized with val and scale
func NewOfferSpotRate(val decimal.Decimal, scale int32) OfferSpotRateField {
	return OfferSpotRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OfferSpotRateField) Value() (val decimal.Decimal) { return f.Decimal }

//OfferSwapPointsField is a PRICEOFFSET field
type OfferSwapPointsField struct{ quickfix.FIXDecimal }

//Tag returns tag.OfferSwapPoints (1066)
func (f OfferSwapPointsField) Tag() quickfix.Tag { return tag.OfferSwapPoints }

//NewOfferSwapPoints returns a new OfferSwapPointsField initialized with val and scale
func NewOfferSwapPoints(val decimal.Decimal, scale int32) OfferSwapPointsField {
	return OfferSwapPointsField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OfferSwapPointsField) Value() (val decimal.Decimal) { return f.Decimal }

//OfferYieldField is a PERCENTAGE field
type OfferYieldField struct{ quickfix.FIXDecimal }

//Tag returns tag.OfferYield (634)
func (f OfferYieldField) Tag() quickfix.Tag { return tag.OfferYield }

//NewOfferYield returns a new OfferYieldField initialized with val and scale
func NewOfferYield(val decimal.Decimal, scale int32) OfferYieldField {
	return OfferYieldField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OfferYieldField) Value() (val decimal.Decimal) { return f.Decimal }

//OnBehalfOfCompIDField is a STRING field
type OnBehalfOfCompIDField struct{ quickfix.FIXString }

//Tag returns tag.OnBehalfOfCompID (115)
func (f OnBehalfOfCompIDField) Tag() quickfix.Tag { return tag.OnBehalfOfCompID }

//NewOnBehalfOfCompID returns a new OnBehalfOfCompIDField initialized with val
func NewOnBehalfOfCompID(val string) OnBehalfOfCompIDField {
	return OnBehalfOfCompIDField{quickfix.FIXString(val)}
}

func (f OnBehalfOfCompIDField) Value() string { return f.String() }

//OnBehalfOfLocationIDField is a STRING field
type OnBehalfOfLocationIDField struct{ quickfix.FIXString }

//Tag returns tag.OnBehalfOfLocationID (144)
func (f OnBehalfOfLocationIDField) Tag() quickfix.Tag { return tag.OnBehalfOfLocationID }

//NewOnBehalfOfLocationID returns a new OnBehalfOfLocationIDField initialized with val
func NewOnBehalfOfLocationID(val string) OnBehalfOfLocationIDField {
	return OnBehalfOfLocationIDField{quickfix.FIXString(val)}
}

func (f OnBehalfOfLocationIDField) Value() string { return f.String() }

//OnBehalfOfSendingTimeField is a UTCTIMESTAMP field
type OnBehalfOfSendingTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.OnBehalfOfSendingTime (370)
func (f OnBehalfOfSendingTimeField) Tag() quickfix.Tag { return tag.OnBehalfOfSendingTime }

//NewOnBehalfOfSendingTime returns a new OnBehalfOfSendingTimeField initialized with val
func NewOnBehalfOfSendingTime(val time.Time) OnBehalfOfSendingTimeField {
	return OnBehalfOfSendingTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewOnBehalfOfSendingTimeNoMillis returns a new OnBehalfOfSendingTimeField initialized with val without millisecs
func NewOnBehalfOfSendingTimeNoMillis(val time.Time) OnBehalfOfSendingTimeField {
	return OnBehalfOfSendingTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f OnBehalfOfSendingTimeField) Value() time.Time { return f.Time }

//OnBehalfOfSubIDField is a STRING field
type OnBehalfOfSubIDField struct{ quickfix.FIXString }

//Tag returns tag.OnBehalfOfSubID (116)
func (f OnBehalfOfSubIDField) Tag() quickfix.Tag { return tag.OnBehalfOfSubID }

//NewOnBehalfOfSubID returns a new OnBehalfOfSubIDField initialized with val
func NewOnBehalfOfSubID(val string) OnBehalfOfSubIDField {
	return OnBehalfOfSubIDField{quickfix.FIXString(val)}
}

func (f OnBehalfOfSubIDField) Value() string { return f.String() }

//OpenCloseField is a enum.OpenClose field
type OpenCloseField struct{ quickfix.FIXString }

//Tag returns tag.OpenClose (77)
func (f OpenCloseField) Tag() quickfix.Tag { return tag.OpenClose }

func NewOpenClose(val enum.OpenClose) OpenCloseField {
	return OpenCloseField{quickfix.FIXString(val)}
}

func (f OpenCloseField) Value() enum.OpenClose { return enum.OpenClose(f.String()) }

//OpenCloseSettlFlagField is a enum.OpenCloseSettlFlag field
type OpenCloseSettlFlagField struct{ quickfix.FIXString }

//Tag returns tag.OpenCloseSettlFlag (286)
func (f OpenCloseSettlFlagField) Tag() quickfix.Tag { return tag.OpenCloseSettlFlag }

func NewOpenCloseSettlFlag(val enum.OpenCloseSettlFlag) OpenCloseSettlFlagField {
	return OpenCloseSettlFlagField{quickfix.FIXString(val)}
}

func (f OpenCloseSettlFlagField) Value() enum.OpenCloseSettlFlag {
	return enum.OpenCloseSettlFlag(f.String())
}

//OpenCloseSettleFlagField is a enum.OpenCloseSettleFlag field
type OpenCloseSettleFlagField struct{ quickfix.FIXString }

//Tag returns tag.OpenCloseSettleFlag (286)
func (f OpenCloseSettleFlagField) Tag() quickfix.Tag { return tag.OpenCloseSettleFlag }

func NewOpenCloseSettleFlag(val enum.OpenCloseSettleFlag) OpenCloseSettleFlagField {
	return OpenCloseSettleFlagField{quickfix.FIXString(val)}
}

func (f OpenCloseSettleFlagField) Value() enum.OpenCloseSettleFlag {
	return enum.OpenCloseSettleFlag(f.String())
}

//OpenInterestField is a AMT field
type OpenInterestField struct{ quickfix.FIXDecimal }

//Tag returns tag.OpenInterest (746)
func (f OpenInterestField) Tag() quickfix.Tag { return tag.OpenInterest }

//NewOpenInterest returns a new OpenInterestField initialized with val and scale
func NewOpenInterest(val decimal.Decimal, scale int32) OpenInterestField {
	return OpenInterestField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OpenInterestField) Value() (val decimal.Decimal) { return f.Decimal }

//OptAttributeField is a CHAR field
type OptAttributeField struct{ quickfix.FIXString }

//Tag returns tag.OptAttribute (206)
func (f OptAttributeField) Tag() quickfix.Tag { return tag.OptAttribute }

//NewOptAttribute returns a new OptAttributeField initialized with val
func NewOptAttribute(val string) OptAttributeField {
	return OptAttributeField{quickfix.FIXString(val)}
}

func (f OptAttributeField) Value() string { return f.String() }

//OptPayAmountField is a AMT field
type OptPayAmountField struct{ quickfix.FIXDecimal }

//Tag returns tag.OptPayAmount (1195)
func (f OptPayAmountField) Tag() quickfix.Tag { return tag.OptPayAmount }

//NewOptPayAmount returns a new OptPayAmountField initialized with val and scale
func NewOptPayAmount(val decimal.Decimal, scale int32) OptPayAmountField {
	return OptPayAmountField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OptPayAmountField) Value() (val decimal.Decimal) { return f.Decimal }

//OptPayoutAmountField is a AMT field
type OptPayoutAmountField struct{ quickfix.FIXDecimal }

//Tag returns tag.OptPayoutAmount (1195)
func (f OptPayoutAmountField) Tag() quickfix.Tag { return tag.OptPayoutAmount }

//NewOptPayoutAmount returns a new OptPayoutAmountField initialized with val and scale
func NewOptPayoutAmount(val decimal.Decimal, scale int32) OptPayoutAmountField {
	return OptPayoutAmountField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OptPayoutAmountField) Value() (val decimal.Decimal) { return f.Decimal }

//OptPayoutTypeField is a enum.OptPayoutType field
type OptPayoutTypeField struct{ quickfix.FIXString }

//Tag returns tag.OptPayoutType (1482)
func (f OptPayoutTypeField) Tag() quickfix.Tag { return tag.OptPayoutType }

func NewOptPayoutType(val enum.OptPayoutType) OptPayoutTypeField {
	return OptPayoutTypeField{quickfix.FIXString(val)}
}

func (f OptPayoutTypeField) Value() enum.OptPayoutType { return enum.OptPayoutType(f.String()) }

//OrdRejReasonField is a enum.OrdRejReason field
type OrdRejReasonField struct{ quickfix.FIXString }

//Tag returns tag.OrdRejReason (103)
func (f OrdRejReasonField) Tag() quickfix.Tag { return tag.OrdRejReason }

func NewOrdRejReason(val enum.OrdRejReason) OrdRejReasonField {
	return OrdRejReasonField{quickfix.FIXString(val)}
}

func (f OrdRejReasonField) Value() enum.OrdRejReason { return enum.OrdRejReason(f.String()) }

//OrdStatusField is a enum.OrdStatus field
type OrdStatusField struct{ quickfix.FIXString }

//Tag returns tag.OrdStatus (39)
func (f OrdStatusField) Tag() quickfix.Tag { return tag.OrdStatus }

func NewOrdStatus(val enum.OrdStatus) OrdStatusField {
	return OrdStatusField{quickfix.FIXString(val)}
}

func (f OrdStatusField) Value() enum.OrdStatus { return enum.OrdStatus(f.String()) }

//OrdStatusReqIDField is a STRING field
type OrdStatusReqIDField struct{ quickfix.FIXString }

//Tag returns tag.OrdStatusReqID (790)
func (f OrdStatusReqIDField) Tag() quickfix.Tag { return tag.OrdStatusReqID }

//NewOrdStatusReqID returns a new OrdStatusReqIDField initialized with val
func NewOrdStatusReqID(val string) OrdStatusReqIDField {
	return OrdStatusReqIDField{quickfix.FIXString(val)}
}

func (f OrdStatusReqIDField) Value() string { return f.String() }

//OrdTypeField is a enum.OrdType field
type OrdTypeField struct{ quickfix.FIXString }

//Tag returns tag.OrdType (40)
func (f OrdTypeField) Tag() quickfix.Tag { return tag.OrdType }

func NewOrdType(val enum.OrdType) OrdTypeField {
	return OrdTypeField{quickfix.FIXString(val)}
}

func (f OrdTypeField) Value() enum.OrdType { return enum.OrdType(f.String()) }

//OrderAvgPxField is a PRICE field
type OrderAvgPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.OrderAvgPx (799)
func (f OrderAvgPxField) Tag() quickfix.Tag { return tag.OrderAvgPx }

//NewOrderAvgPx returns a new OrderAvgPxField initialized with val and scale
func NewOrderAvgPx(val decimal.Decimal, scale int32) OrderAvgPxField {
	return OrderAvgPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OrderAvgPxField) Value() (val decimal.Decimal) { return f.Decimal }

//OrderBookingQtyField is a QTY field
type OrderBookingQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.OrderBookingQty (800)
func (f OrderBookingQtyField) Tag() quickfix.Tag { return tag.OrderBookingQty }

//NewOrderBookingQty returns a new OrderBookingQtyField initialized with val and scale
func NewOrderBookingQty(val decimal.Decimal, scale int32) OrderBookingQtyField {
	return OrderBookingQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OrderBookingQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//OrderCapacityField is a enum.OrderCapacity field
type OrderCapacityField struct{ quickfix.FIXString }

//Tag returns tag.OrderCapacity (528)
func (f OrderCapacityField) Tag() quickfix.Tag { return tag.OrderCapacity }

func NewOrderCapacity(val enum.OrderCapacity) OrderCapacityField {
	return OrderCapacityField{quickfix.FIXString(val)}
}

func (f OrderCapacityField) Value() enum.OrderCapacity { return enum.OrderCapacity(f.String()) }

//OrderCapacityQtyField is a QTY field
type OrderCapacityQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.OrderCapacityQty (863)
func (f OrderCapacityQtyField) Tag() quickfix.Tag { return tag.OrderCapacityQty }

//NewOrderCapacityQty returns a new OrderCapacityQtyField initialized with val and scale
func NewOrderCapacityQty(val decimal.Decimal, scale int32) OrderCapacityQtyField {
	return OrderCapacityQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OrderCapacityQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//OrderCategoryField is a enum.OrderCategory field
type OrderCategoryField struct{ quickfix.FIXString }

//Tag returns tag.OrderCategory (1115)
func (f OrderCategoryField) Tag() quickfix.Tag { return tag.OrderCategory }

func NewOrderCategory(val enum.OrderCategory) OrderCategoryField {
	return OrderCategoryField{quickfix.FIXString(val)}
}

func (f OrderCategoryField) Value() enum.OrderCategory { return enum.OrderCategory(f.String()) }

//OrderDelayField is a INT field
type OrderDelayField struct{ quickfix.FIXInt }

//Tag returns tag.OrderDelay (1428)
func (f OrderDelayField) Tag() quickfix.Tag { return tag.OrderDelay }

//NewOrderDelay returns a new OrderDelayField initialized with val
func NewOrderDelay(val int) OrderDelayField {
	return OrderDelayField{quickfix.FIXInt(val)}
}

func (f OrderDelayField) Value() int { return f.Int() }

//OrderDelayUnitField is a enum.OrderDelayUnit field
type OrderDelayUnitField struct{ quickfix.FIXString }

//Tag returns tag.OrderDelayUnit (1429)
func (f OrderDelayUnitField) Tag() quickfix.Tag { return tag.OrderDelayUnit }

func NewOrderDelayUnit(val enum.OrderDelayUnit) OrderDelayUnitField {
	return OrderDelayUnitField{quickfix.FIXString(val)}
}

func (f OrderDelayUnitField) Value() enum.OrderDelayUnit { return enum.OrderDelayUnit(f.String()) }

//OrderHandlingInstSourceField is a enum.OrderHandlingInstSource field
type OrderHandlingInstSourceField struct{ quickfix.FIXString }

//Tag returns tag.OrderHandlingInstSource (1032)
func (f OrderHandlingInstSourceField) Tag() quickfix.Tag { return tag.OrderHandlingInstSource }

func NewOrderHandlingInstSource(val enum.OrderHandlingInstSource) OrderHandlingInstSourceField {
	return OrderHandlingInstSourceField{quickfix.FIXString(val)}
}

func (f OrderHandlingInstSourceField) Value() enum.OrderHandlingInstSource {
	return enum.OrderHandlingInstSource(f.String())
}

//OrderIDField is a STRING field
type OrderIDField struct{ quickfix.FIXString }

//Tag returns tag.OrderID (37)
func (f OrderIDField) Tag() quickfix.Tag { return tag.OrderID }

//NewOrderID returns a new OrderIDField initialized with val
func NewOrderID(val string) OrderIDField {
	return OrderIDField{quickfix.FIXString(val)}
}

func (f OrderIDField) Value() string { return f.String() }

//OrderInputDeviceField is a STRING field
type OrderInputDeviceField struct{ quickfix.FIXString }

//Tag returns tag.OrderInputDevice (821)
func (f OrderInputDeviceField) Tag() quickfix.Tag { return tag.OrderInputDevice }

//NewOrderInputDevice returns a new OrderInputDeviceField initialized with val
func NewOrderInputDevice(val string) OrderInputDeviceField {
	return OrderInputDeviceField{quickfix.FIXString(val)}
}

func (f OrderInputDeviceField) Value() string { return f.String() }

//OrderPercentField is a PERCENTAGE field
type OrderPercentField struct{ quickfix.FIXDecimal }

//Tag returns tag.OrderPercent (516)
func (f OrderPercentField) Tag() quickfix.Tag { return tag.OrderPercent }

//NewOrderPercent returns a new OrderPercentField initialized with val and scale
func NewOrderPercent(val decimal.Decimal, scale int32) OrderPercentField {
	return OrderPercentField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OrderPercentField) Value() (val decimal.Decimal) { return f.Decimal }

//OrderQtyField is a QTY field
type OrderQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.OrderQty (38)
func (f OrderQtyField) Tag() quickfix.Tag { return tag.OrderQty }

//NewOrderQty returns a new OrderQtyField initialized with val and scale
func NewOrderQty(val decimal.Decimal, scale int32) OrderQtyField {
	return OrderQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OrderQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//OrderQty2Field is a QTY field
type OrderQty2Field struct{ quickfix.FIXDecimal }

//Tag returns tag.OrderQty2 (192)
func (f OrderQty2Field) Tag() quickfix.Tag { return tag.OrderQty2 }

//NewOrderQty2 returns a new OrderQty2Field initialized with val and scale
func NewOrderQty2(val decimal.Decimal, scale int32) OrderQty2Field {
	return OrderQty2Field{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OrderQty2Field) Value() (val decimal.Decimal) { return f.Decimal }

//OrderRestrictionsField is a enum.OrderRestrictions field
type OrderRestrictionsField struct{ quickfix.FIXString }

//Tag returns tag.OrderRestrictions (529)
func (f OrderRestrictionsField) Tag() quickfix.Tag { return tag.OrderRestrictions }

func NewOrderRestrictions(val enum.OrderRestrictions) OrderRestrictionsField {
	return OrderRestrictionsField{quickfix.FIXString(val)}
}

func (f OrderRestrictionsField) Value() enum.OrderRestrictions {
	return enum.OrderRestrictions(f.String())
}

//OrigClOrdIDField is a STRING field
type OrigClOrdIDField struct{ quickfix.FIXString }

//Tag returns tag.OrigClOrdID (41)
func (f OrigClOrdIDField) Tag() quickfix.Tag { return tag.OrigClOrdID }

//NewOrigClOrdID returns a new OrigClOrdIDField initialized with val
func NewOrigClOrdID(val string) OrigClOrdIDField {
	return OrigClOrdIDField{quickfix.FIXString(val)}
}

func (f OrigClOrdIDField) Value() string { return f.String() }

//OrigCrossIDField is a STRING field
type OrigCrossIDField struct{ quickfix.FIXString }

//Tag returns tag.OrigCrossID (551)
func (f OrigCrossIDField) Tag() quickfix.Tag { return tag.OrigCrossID }

//NewOrigCrossID returns a new OrigCrossIDField initialized with val
func NewOrigCrossID(val string) OrigCrossIDField {
	return OrigCrossIDField{quickfix.FIXString(val)}
}

func (f OrigCrossIDField) Value() string { return f.String() }

//OrigCustOrderCapacityField is a enum.OrigCustOrderCapacity field
type OrigCustOrderCapacityField struct{ quickfix.FIXString }

//Tag returns tag.OrigCustOrderCapacity (1432)
func (f OrigCustOrderCapacityField) Tag() quickfix.Tag { return tag.OrigCustOrderCapacity }

func NewOrigCustOrderCapacity(val enum.OrigCustOrderCapacity) OrigCustOrderCapacityField {
	return OrigCustOrderCapacityField{quickfix.FIXString(val)}
}

func (f OrigCustOrderCapacityField) Value() enum.OrigCustOrderCapacity {
	return enum.OrigCustOrderCapacity(f.String())
}

//OrigOrdModTimeField is a UTCTIMESTAMP field
type OrigOrdModTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.OrigOrdModTime (586)
func (f OrigOrdModTimeField) Tag() quickfix.Tag { return tag.OrigOrdModTime }

//NewOrigOrdModTime returns a new OrigOrdModTimeField initialized with val
func NewOrigOrdModTime(val time.Time) OrigOrdModTimeField {
	return OrigOrdModTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewOrigOrdModTimeNoMillis returns a new OrigOrdModTimeField initialized with val without millisecs
func NewOrigOrdModTimeNoMillis(val time.Time) OrigOrdModTimeField {
	return OrigOrdModTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f OrigOrdModTimeField) Value() time.Time { return f.Time }

//OrigPosReqRefIDField is a STRING field
type OrigPosReqRefIDField struct{ quickfix.FIXString }

//Tag returns tag.OrigPosReqRefID (713)
func (f OrigPosReqRefIDField) Tag() quickfix.Tag { return tag.OrigPosReqRefID }

//NewOrigPosReqRefID returns a new OrigPosReqRefIDField initialized with val
func NewOrigPosReqRefID(val string) OrigPosReqRefIDField {
	return OrigPosReqRefIDField{quickfix.FIXString(val)}
}

func (f OrigPosReqRefIDField) Value() string { return f.String() }

//OrigSecondaryTradeIDField is a STRING field
type OrigSecondaryTradeIDField struct{ quickfix.FIXString }

//Tag returns tag.OrigSecondaryTradeID (1127)
func (f OrigSecondaryTradeIDField) Tag() quickfix.Tag { return tag.OrigSecondaryTradeID }

//NewOrigSecondaryTradeID returns a new OrigSecondaryTradeIDField initialized with val
func NewOrigSecondaryTradeID(val string) OrigSecondaryTradeIDField {
	return OrigSecondaryTradeIDField{quickfix.FIXString(val)}
}

func (f OrigSecondaryTradeIDField) Value() string { return f.String() }

//OrigSendingTimeField is a UTCTIMESTAMP field
type OrigSendingTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.OrigSendingTime (122)
func (f OrigSendingTimeField) Tag() quickfix.Tag { return tag.OrigSendingTime }

//NewOrigSendingTime returns a new OrigSendingTimeField initialized with val
func NewOrigSendingTime(val time.Time) OrigSendingTimeField {
	return OrigSendingTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewOrigSendingTimeNoMillis returns a new OrigSendingTimeField initialized with val without millisecs
func NewOrigSendingTimeNoMillis(val time.Time) OrigSendingTimeField {
	return OrigSendingTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f OrigSendingTimeField) Value() time.Time { return f.Time }

//OrigTimeField is a UTCTIMESTAMP field
type OrigTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.OrigTime (42)
func (f OrigTimeField) Tag() quickfix.Tag { return tag.OrigTime }

//NewOrigTime returns a new OrigTimeField initialized with val
func NewOrigTime(val time.Time) OrigTimeField {
	return OrigTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewOrigTimeNoMillis returns a new OrigTimeField initialized with val without millisecs
func NewOrigTimeNoMillis(val time.Time) OrigTimeField {
	return OrigTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f OrigTimeField) Value() time.Time { return f.Time }

//OrigTradeDateField is a LOCALMKTDATE field
type OrigTradeDateField struct{ quickfix.FIXString }

//Tag returns tag.OrigTradeDate (1125)
func (f OrigTradeDateField) Tag() quickfix.Tag { return tag.OrigTradeDate }

//NewOrigTradeDate returns a new OrigTradeDateField initialized with val
func NewOrigTradeDate(val string) OrigTradeDateField {
	return OrigTradeDateField{quickfix.FIXString(val)}
}

func (f OrigTradeDateField) Value() string { return f.String() }

//OrigTradeHandlingInstrField is a CHAR field
type OrigTradeHandlingInstrField struct{ quickfix.FIXString }

//Tag returns tag.OrigTradeHandlingInstr (1124)
func (f OrigTradeHandlingInstrField) Tag() quickfix.Tag { return tag.OrigTradeHandlingInstr }

//NewOrigTradeHandlingInstr returns a new OrigTradeHandlingInstrField initialized with val
func NewOrigTradeHandlingInstr(val string) OrigTradeHandlingInstrField {
	return OrigTradeHandlingInstrField{quickfix.FIXString(val)}
}

func (f OrigTradeHandlingInstrField) Value() string { return f.String() }

//OrigTradeIDField is a STRING field
type OrigTradeIDField struct{ quickfix.FIXString }

//Tag returns tag.OrigTradeID (1126)
func (f OrigTradeIDField) Tag() quickfix.Tag { return tag.OrigTradeID }

//NewOrigTradeID returns a new OrigTradeIDField initialized with val
func NewOrigTradeID(val string) OrigTradeIDField {
	return OrigTradeIDField{quickfix.FIXString(val)}
}

func (f OrigTradeIDField) Value() string { return f.String() }

//OriginalNotionalPercentageOutstandingField is a PERCENTAGE field
type OriginalNotionalPercentageOutstandingField struct{ quickfix.FIXDecimal }

//Tag returns tag.OriginalNotionalPercentageOutstanding (1452)
func (f OriginalNotionalPercentageOutstandingField) Tag() quickfix.Tag {
	return tag.OriginalNotionalPercentageOutstanding
}

//NewOriginalNotionalPercentageOutstanding returns a new OriginalNotionalPercentageOutstandingField initialized with val and scale
func NewOriginalNotionalPercentageOutstanding(val decimal.Decimal, scale int32) OriginalNotionalPercentageOutstandingField {
	return OriginalNotionalPercentageOutstandingField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OriginalNotionalPercentageOutstandingField) Value() (val decimal.Decimal) { return f.Decimal }

//OutMainCntryUIndexField is a AMT field
type OutMainCntryUIndexField struct{ quickfix.FIXDecimal }

//Tag returns tag.OutMainCntryUIndex (412)
func (f OutMainCntryUIndexField) Tag() quickfix.Tag { return tag.OutMainCntryUIndex }

//NewOutMainCntryUIndex returns a new OutMainCntryUIndexField initialized with val and scale
func NewOutMainCntryUIndex(val decimal.Decimal, scale int32) OutMainCntryUIndexField {
	return OutMainCntryUIndexField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OutMainCntryUIndexField) Value() (val decimal.Decimal) { return f.Decimal }

//OutsideIndexPctField is a PERCENTAGE field
type OutsideIndexPctField struct{ quickfix.FIXDecimal }

//Tag returns tag.OutsideIndexPct (407)
func (f OutsideIndexPctField) Tag() quickfix.Tag { return tag.OutsideIndexPct }

//NewOutsideIndexPct returns a new OutsideIndexPctField initialized with val and scale
func NewOutsideIndexPct(val decimal.Decimal, scale int32) OutsideIndexPctField {
	return OutsideIndexPctField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OutsideIndexPctField) Value() (val decimal.Decimal) { return f.Decimal }

//OwnerTypeField is a enum.OwnerType field
type OwnerTypeField struct{ quickfix.FIXString }

//Tag returns tag.OwnerType (522)
func (f OwnerTypeField) Tag() quickfix.Tag { return tag.OwnerType }

func NewOwnerType(val enum.OwnerType) OwnerTypeField {
	return OwnerTypeField{quickfix.FIXString(val)}
}

func (f OwnerTypeField) Value() enum.OwnerType { return enum.OwnerType(f.String()) }

//OwnershipTypeField is a enum.OwnershipType field
type OwnershipTypeField struct{ quickfix.FIXString }

//Tag returns tag.OwnershipType (517)
func (f OwnershipTypeField) Tag() quickfix.Tag { return tag.OwnershipType }

func NewOwnershipType(val enum.OwnershipType) OwnershipTypeField {
	return OwnershipTypeField{quickfix.FIXString(val)}
}

func (f OwnershipTypeField) Value() enum.OwnershipType { return enum.OwnershipType(f.String()) }

//ParentMktSegmIDField is a STRING field
type ParentMktSegmIDField struct{ quickfix.FIXString }

//Tag returns tag.ParentMktSegmID (1325)
func (f ParentMktSegmIDField) Tag() quickfix.Tag { return tag.ParentMktSegmID }

//NewParentMktSegmID returns a new ParentMktSegmIDField initialized with val
func NewParentMktSegmID(val string) ParentMktSegmIDField {
	return ParentMktSegmIDField{quickfix.FIXString(val)}
}

func (f ParentMktSegmIDField) Value() string { return f.String() }

//ParticipationRateField is a PERCENTAGE field
type ParticipationRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.ParticipationRate (849)
func (f ParticipationRateField) Tag() quickfix.Tag { return tag.ParticipationRate }

//NewParticipationRate returns a new ParticipationRateField initialized with val and scale
func NewParticipationRate(val decimal.Decimal, scale int32) ParticipationRateField {
	return ParticipationRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f ParticipationRateField) Value() (val decimal.Decimal) { return f.Decimal }

//PartyAltIDField is a STRING field
type PartyAltIDField struct{ quickfix.FIXString }

//Tag returns tag.PartyAltID (1517)
func (f PartyAltIDField) Tag() quickfix.Tag { return tag.PartyAltID }

//NewPartyAltID returns a new PartyAltIDField initialized with val
func NewPartyAltID(val string) PartyAltIDField {
	return PartyAltIDField{quickfix.FIXString(val)}
}

func (f PartyAltIDField) Value() string { return f.String() }

//PartyAltIDSourceField is a CHAR field
type PartyAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.PartyAltIDSource (1518)
func (f PartyAltIDSourceField) Tag() quickfix.Tag { return tag.PartyAltIDSource }

//NewPartyAltIDSource returns a new PartyAltIDSourceField initialized with val
func NewPartyAltIDSource(val string) PartyAltIDSourceField {
	return PartyAltIDSourceField{quickfix.FIXString(val)}
}

func (f PartyAltIDSourceField) Value() string { return f.String() }

//PartyAltSubIDField is a STRING field
type PartyAltSubIDField struct{ quickfix.FIXString }

//Tag returns tag.PartyAltSubID (1520)
func (f PartyAltSubIDField) Tag() quickfix.Tag { return tag.PartyAltSubID }

//NewPartyAltSubID returns a new PartyAltSubIDField initialized with val
func NewPartyAltSubID(val string) PartyAltSubIDField {
	return PartyAltSubIDField{quickfix.FIXString(val)}
}

func (f PartyAltSubIDField) Value() string { return f.String() }

//PartyAltSubIDTypeField is a INT field
type PartyAltSubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.PartyAltSubIDType (1521)
func (f PartyAltSubIDTypeField) Tag() quickfix.Tag { return tag.PartyAltSubIDType }

//NewPartyAltSubIDType returns a new PartyAltSubIDTypeField initialized with val
func NewPartyAltSubIDType(val int) PartyAltSubIDTypeField {
	return PartyAltSubIDTypeField{quickfix.FIXInt(val)}
}

func (f PartyAltSubIDTypeField) Value() int { return f.Int() }

//PartyDetailsListReportIDField is a STRING field
type PartyDetailsListReportIDField struct{ quickfix.FIXString }

//Tag returns tag.PartyDetailsListReportID (1510)
func (f PartyDetailsListReportIDField) Tag() quickfix.Tag { return tag.PartyDetailsListReportID }

//NewPartyDetailsListReportID returns a new PartyDetailsListReportIDField initialized with val
func NewPartyDetailsListReportID(val string) PartyDetailsListReportIDField {
	return PartyDetailsListReportIDField{quickfix.FIXString(val)}
}

func (f PartyDetailsListReportIDField) Value() string { return f.String() }

//PartyDetailsListRequestIDField is a STRING field
type PartyDetailsListRequestIDField struct{ quickfix.FIXString }

//Tag returns tag.PartyDetailsListRequestID (1505)
func (f PartyDetailsListRequestIDField) Tag() quickfix.Tag { return tag.PartyDetailsListRequestID }

//NewPartyDetailsListRequestID returns a new PartyDetailsListRequestIDField initialized with val
func NewPartyDetailsListRequestID(val string) PartyDetailsListRequestIDField {
	return PartyDetailsListRequestIDField{quickfix.FIXString(val)}
}

func (f PartyDetailsListRequestIDField) Value() string { return f.String() }

//PartyDetailsRequestResultField is a enum.PartyDetailsRequestResult field
type PartyDetailsRequestResultField struct{ quickfix.FIXString }

//Tag returns tag.PartyDetailsRequestResult (1511)
func (f PartyDetailsRequestResultField) Tag() quickfix.Tag { return tag.PartyDetailsRequestResult }

func NewPartyDetailsRequestResult(val enum.PartyDetailsRequestResult) PartyDetailsRequestResultField {
	return PartyDetailsRequestResultField{quickfix.FIXString(val)}
}

func (f PartyDetailsRequestResultField) Value() enum.PartyDetailsRequestResult {
	return enum.PartyDetailsRequestResult(f.String())
}

//PartyIDField is a STRING field
type PartyIDField struct{ quickfix.FIXString }

//Tag returns tag.PartyID (448)
func (f PartyIDField) Tag() quickfix.Tag { return tag.PartyID }

//NewPartyID returns a new PartyIDField initialized with val
func NewPartyID(val string) PartyIDField {
	return PartyIDField{quickfix.FIXString(val)}
}

func (f PartyIDField) Value() string { return f.String() }

//PartyIDSourceField is a enum.PartyIDSource field
type PartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.PartyIDSource (447)
func (f PartyIDSourceField) Tag() quickfix.Tag { return tag.PartyIDSource }

func NewPartyIDSource(val enum.PartyIDSource) PartyIDSourceField {
	return PartyIDSourceField{quickfix.FIXString(val)}
}

func (f PartyIDSourceField) Value() enum.PartyIDSource { return enum.PartyIDSource(f.String()) }

//PartyListResponseTypeField is a enum.PartyListResponseType field
type PartyListResponseTypeField struct{ quickfix.FIXString }

//Tag returns tag.PartyListResponseType (1507)
func (f PartyListResponseTypeField) Tag() quickfix.Tag { return tag.PartyListResponseType }

func NewPartyListResponseType(val enum.PartyListResponseType) PartyListResponseTypeField {
	return PartyListResponseTypeField{quickfix.FIXString(val)}
}

func (f PartyListResponseTypeField) Value() enum.PartyListResponseType {
	return enum.PartyListResponseType(f.String())
}

//PartyRelationshipField is a enum.PartyRelationship field
type PartyRelationshipField struct{ quickfix.FIXString }

//Tag returns tag.PartyRelationship (1515)
func (f PartyRelationshipField) Tag() quickfix.Tag { return tag.PartyRelationship }

func NewPartyRelationship(val enum.PartyRelationship) PartyRelationshipField {
	return PartyRelationshipField{quickfix.FIXString(val)}
}

func (f PartyRelationshipField) Value() enum.PartyRelationship {
	return enum.PartyRelationship(f.String())
}

//PartyRoleField is a enum.PartyRole field
type PartyRoleField struct{ quickfix.FIXString }

//Tag returns tag.PartyRole (452)
func (f PartyRoleField) Tag() quickfix.Tag { return tag.PartyRole }

func NewPartyRole(val enum.PartyRole) PartyRoleField {
	return PartyRoleField{quickfix.FIXString(val)}
}

func (f PartyRoleField) Value() enum.PartyRole { return enum.PartyRole(f.String()) }

//PartySubIDField is a STRING field
type PartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.PartySubID (523)
func (f PartySubIDField) Tag() quickfix.Tag { return tag.PartySubID }

//NewPartySubID returns a new PartySubIDField initialized with val
func NewPartySubID(val string) PartySubIDField {
	return PartySubIDField{quickfix.FIXString(val)}
}

func (f PartySubIDField) Value() string { return f.String() }

//PartySubIDTypeField is a enum.PartySubIDType field
type PartySubIDTypeField struct{ quickfix.FIXString }

//Tag returns tag.PartySubIDType (803)
func (f PartySubIDTypeField) Tag() quickfix.Tag { return tag.PartySubIDType }

func NewPartySubIDType(val enum.PartySubIDType) PartySubIDTypeField {
	return PartySubIDTypeField{quickfix.FIXString(val)}
}

func (f PartySubIDTypeField) Value() enum.PartySubIDType { return enum.PartySubIDType(f.String()) }

//PasswordField is a STRING field
type PasswordField struct{ quickfix.FIXString }

//Tag returns tag.Password (554)
func (f PasswordField) Tag() quickfix.Tag { return tag.Password }

//NewPassword returns a new PasswordField initialized with val
func NewPassword(val string) PasswordField {
	return PasswordField{quickfix.FIXString(val)}
}

func (f PasswordField) Value() string { return f.String() }

//PaymentDateField is a LOCALMKTDATE field
type PaymentDateField struct{ quickfix.FIXString }

//Tag returns tag.PaymentDate (504)
func (f PaymentDateField) Tag() quickfix.Tag { return tag.PaymentDate }

//NewPaymentDate returns a new PaymentDateField initialized with val
func NewPaymentDate(val string) PaymentDateField {
	return PaymentDateField{quickfix.FIXString(val)}
}

func (f PaymentDateField) Value() string { return f.String() }

//PaymentMethodField is a enum.PaymentMethod field
type PaymentMethodField struct{ quickfix.FIXString }

//Tag returns tag.PaymentMethod (492)
func (f PaymentMethodField) Tag() quickfix.Tag { return tag.PaymentMethod }

func NewPaymentMethod(val enum.PaymentMethod) PaymentMethodField {
	return PaymentMethodField{quickfix.FIXString(val)}
}

func (f PaymentMethodField) Value() enum.PaymentMethod { return enum.PaymentMethod(f.String()) }

//PaymentRefField is a STRING field
type PaymentRefField struct{ quickfix.FIXString }

//Tag returns tag.PaymentRef (476)
func (f PaymentRefField) Tag() quickfix.Tag { return tag.PaymentRef }

//NewPaymentRef returns a new PaymentRefField initialized with val
func NewPaymentRef(val string) PaymentRefField {
	return PaymentRefField{quickfix.FIXString(val)}
}

func (f PaymentRefField) Value() string { return f.String() }

//PaymentRemitterIDField is a STRING field
type PaymentRemitterIDField struct{ quickfix.FIXString }

//Tag returns tag.PaymentRemitterID (505)
func (f PaymentRemitterIDField) Tag() quickfix.Tag { return tag.PaymentRemitterID }

//NewPaymentRemitterID returns a new PaymentRemitterIDField initialized with val
func NewPaymentRemitterID(val string) PaymentRemitterIDField {
	return PaymentRemitterIDField{quickfix.FIXString(val)}
}

func (f PaymentRemitterIDField) Value() string { return f.String() }

//PctAtRiskField is a PERCENTAGE field
type PctAtRiskField struct{ quickfix.FIXDecimal }

//Tag returns tag.PctAtRisk (869)
func (f PctAtRiskField) Tag() quickfix.Tag { return tag.PctAtRisk }

//NewPctAtRisk returns a new PctAtRiskField initialized with val and scale
func NewPctAtRisk(val decimal.Decimal, scale int32) PctAtRiskField {
	return PctAtRiskField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f PctAtRiskField) Value() (val decimal.Decimal) { return f.Decimal }

//PegDifferenceField is a PRICEOFFSET field
type PegDifferenceField struct{ quickfix.FIXDecimal }

//Tag returns tag.PegDifference (211)
func (f PegDifferenceField) Tag() quickfix.Tag { return tag.PegDifference }

//NewPegDifference returns a new PegDifferenceField initialized with val and scale
func NewPegDifference(val decimal.Decimal, scale int32) PegDifferenceField {
	return PegDifferenceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f PegDifferenceField) Value() (val decimal.Decimal) { return f.Decimal }

//PegLimitTypeField is a enum.PegLimitType field
type PegLimitTypeField struct{ quickfix.FIXString }

//Tag returns tag.PegLimitType (837)
func (f PegLimitTypeField) Tag() quickfix.Tag { return tag.PegLimitType }

func NewPegLimitType(val enum.PegLimitType) PegLimitTypeField {
	return PegLimitTypeField{quickfix.FIXString(val)}
}

func (f PegLimitTypeField) Value() enum.PegLimitType { return enum.PegLimitType(f.String()) }

//PegMoveTypeField is a enum.PegMoveType field
type PegMoveTypeField struct{ quickfix.FIXString }

//Tag returns tag.PegMoveType (835)
func (f PegMoveTypeField) Tag() quickfix.Tag { return tag.PegMoveType }

func NewPegMoveType(val enum.PegMoveType) PegMoveTypeField {
	return PegMoveTypeField{quickfix.FIXString(val)}
}

func (f PegMoveTypeField) Value() enum.PegMoveType { return enum.PegMoveType(f.String()) }

//PegOffsetTypeField is a enum.PegOffsetType field
type PegOffsetTypeField struct{ quickfix.FIXString }

//Tag returns tag.PegOffsetType (836)
func (f PegOffsetTypeField) Tag() quickfix.Tag { return tag.PegOffsetType }

func NewPegOffsetType(val enum.PegOffsetType) PegOffsetTypeField {
	return PegOffsetTypeField{quickfix.FIXString(val)}
}

func (f PegOffsetTypeField) Value() enum.PegOffsetType { return enum.PegOffsetType(f.String()) }

//PegOffsetValueField is a FLOAT field
type PegOffsetValueField struct{ quickfix.FIXDecimal }

//Tag returns tag.PegOffsetValue (211)
func (f PegOffsetValueField) Tag() quickfix.Tag { return tag.PegOffsetValue }

//NewPegOffsetValue returns a new PegOffsetValueField initialized with val and scale
func NewPegOffsetValue(val decimal.Decimal, scale int32) PegOffsetValueField {
	return PegOffsetValueField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f PegOffsetValueField) Value() (val decimal.Decimal) { return f.Decimal }

//PegPriceTypeField is a enum.PegPriceType field
type PegPriceTypeField struct{ quickfix.FIXString }

//Tag returns tag.PegPriceType (1094)
func (f PegPriceTypeField) Tag() quickfix.Tag { return tag.PegPriceType }

func NewPegPriceType(val enum.PegPriceType) PegPriceTypeField {
	return PegPriceTypeField{quickfix.FIXString(val)}
}

func (f PegPriceTypeField) Value() enum.PegPriceType { return enum.PegPriceType(f.String()) }

//PegRoundDirectionField is a enum.PegRoundDirection field
type PegRoundDirectionField struct{ quickfix.FIXString }

//Tag returns tag.PegRoundDirection (838)
func (f PegRoundDirectionField) Tag() quickfix.Tag { return tag.PegRoundDirection }

func NewPegRoundDirection(val enum.PegRoundDirection) PegRoundDirectionField {
	return PegRoundDirectionField{quickfix.FIXString(val)}
}

func (f PegRoundDirectionField) Value() enum.PegRoundDirection {
	return enum.PegRoundDirection(f.String())
}

//PegScopeField is a enum.PegScope field
type PegScopeField struct{ quickfix.FIXString }

//Tag returns tag.PegScope (840)
func (f PegScopeField) Tag() quickfix.Tag { return tag.PegScope }

func NewPegScope(val enum.PegScope) PegScopeField {
	return PegScopeField{quickfix.FIXString(val)}
}

func (f PegScopeField) Value() enum.PegScope { return enum.PegScope(f.String()) }

//PegSecurityDescField is a STRING field
type PegSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.PegSecurityDesc (1099)
func (f PegSecurityDescField) Tag() quickfix.Tag { return tag.PegSecurityDesc }

//NewPegSecurityDesc returns a new PegSecurityDescField initialized with val
func NewPegSecurityDesc(val string) PegSecurityDescField {
	return PegSecurityDescField{quickfix.FIXString(val)}
}

func (f PegSecurityDescField) Value() string { return f.String() }

//PegSecurityIDField is a STRING field
type PegSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.PegSecurityID (1097)
func (f PegSecurityIDField) Tag() quickfix.Tag { return tag.PegSecurityID }

//NewPegSecurityID returns a new PegSecurityIDField initialized with val
func NewPegSecurityID(val string) PegSecurityIDField {
	return PegSecurityIDField{quickfix.FIXString(val)}
}

func (f PegSecurityIDField) Value() string { return f.String() }

//PegSecurityIDSourceField is a STRING field
type PegSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.PegSecurityIDSource (1096)
func (f PegSecurityIDSourceField) Tag() quickfix.Tag { return tag.PegSecurityIDSource }

//NewPegSecurityIDSource returns a new PegSecurityIDSourceField initialized with val
func NewPegSecurityIDSource(val string) PegSecurityIDSourceField {
	return PegSecurityIDSourceField{quickfix.FIXString(val)}
}

func (f PegSecurityIDSourceField) Value() string { return f.String() }

//PegSymbolField is a STRING field
type PegSymbolField struct{ quickfix.FIXString }

//Tag returns tag.PegSymbol (1098)
func (f PegSymbolField) Tag() quickfix.Tag { return tag.PegSymbol }

//NewPegSymbol returns a new PegSymbolField initialized with val
func NewPegSymbol(val string) PegSymbolField {
	return PegSymbolField{quickfix.FIXString(val)}
}

func (f PegSymbolField) Value() string { return f.String() }

//PeggedPriceField is a PRICE field
type PeggedPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.PeggedPrice (839)
func (f PeggedPriceField) Tag() quickfix.Tag { return tag.PeggedPrice }

//NewPeggedPrice returns a new PeggedPriceField initialized with val and scale
func NewPeggedPrice(val decimal.Decimal, scale int32) PeggedPriceField {
	return PeggedPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f PeggedPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//PeggedRefPriceField is a PRICE field
type PeggedRefPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.PeggedRefPrice (1095)
func (f PeggedRefPriceField) Tag() quickfix.Tag { return tag.PeggedRefPrice }

//NewPeggedRefPrice returns a new PeggedRefPriceField initialized with val and scale
func NewPeggedRefPrice(val decimal.Decimal, scale int32) PeggedRefPriceField {
	return PeggedRefPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f PeggedRefPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//PoolField is a STRING field
type PoolField struct{ quickfix.FIXString }

//Tag returns tag.Pool (691)
func (f PoolField) Tag() quickfix.Tag { return tag.Pool }

//NewPool returns a new PoolField initialized with val
func NewPool(val string) PoolField {
	return PoolField{quickfix.FIXString(val)}
}

func (f PoolField) Value() string { return f.String() }

//PosAmtField is a AMT field
type PosAmtField struct{ quickfix.FIXDecimal }

//Tag returns tag.PosAmt (708)
func (f PosAmtField) Tag() quickfix.Tag { return tag.PosAmt }

//NewPosAmt returns a new PosAmtField initialized with val and scale
func NewPosAmt(val decimal.Decimal, scale int32) PosAmtField {
	return PosAmtField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f PosAmtField) Value() (val decimal.Decimal) { return f.Decimal }

//PosAmtTypeField is a enum.PosAmtType field
type PosAmtTypeField struct{ quickfix.FIXString }

//Tag returns tag.PosAmtType (707)
func (f PosAmtTypeField) Tag() quickfix.Tag { return tag.PosAmtType }

func NewPosAmtType(val enum.PosAmtType) PosAmtTypeField {
	return PosAmtTypeField{quickfix.FIXString(val)}
}

func (f PosAmtTypeField) Value() enum.PosAmtType { return enum.PosAmtType(f.String()) }

//PosMaintActionField is a enum.PosMaintAction field
type PosMaintActionField struct{ quickfix.FIXString }

//Tag returns tag.PosMaintAction (712)
func (f PosMaintActionField) Tag() quickfix.Tag { return tag.PosMaintAction }

func NewPosMaintAction(val enum.PosMaintAction) PosMaintActionField {
	return PosMaintActionField{quickfix.FIXString(val)}
}

func (f PosMaintActionField) Value() enum.PosMaintAction { return enum.PosMaintAction(f.String()) }

//PosMaintResultField is a enum.PosMaintResult field
type PosMaintResultField struct{ quickfix.FIXString }

//Tag returns tag.PosMaintResult (723)
func (f PosMaintResultField) Tag() quickfix.Tag { return tag.PosMaintResult }

func NewPosMaintResult(val enum.PosMaintResult) PosMaintResultField {
	return PosMaintResultField{quickfix.FIXString(val)}
}

func (f PosMaintResultField) Value() enum.PosMaintResult { return enum.PosMaintResult(f.String()) }

//PosMaintRptIDField is a STRING field
type PosMaintRptIDField struct{ quickfix.FIXString }

//Tag returns tag.PosMaintRptID (721)
func (f PosMaintRptIDField) Tag() quickfix.Tag { return tag.PosMaintRptID }

//NewPosMaintRptID returns a new PosMaintRptIDField initialized with val
func NewPosMaintRptID(val string) PosMaintRptIDField {
	return PosMaintRptIDField{quickfix.FIXString(val)}
}

func (f PosMaintRptIDField) Value() string { return f.String() }

//PosMaintRptRefIDField is a STRING field
type PosMaintRptRefIDField struct{ quickfix.FIXString }

//Tag returns tag.PosMaintRptRefID (714)
func (f PosMaintRptRefIDField) Tag() quickfix.Tag { return tag.PosMaintRptRefID }

//NewPosMaintRptRefID returns a new PosMaintRptRefIDField initialized with val
func NewPosMaintRptRefID(val string) PosMaintRptRefIDField {
	return PosMaintRptRefIDField{quickfix.FIXString(val)}
}

func (f PosMaintRptRefIDField) Value() string { return f.String() }

//PosMaintStatusField is a enum.PosMaintStatus field
type PosMaintStatusField struct{ quickfix.FIXString }

//Tag returns tag.PosMaintStatus (722)
func (f PosMaintStatusField) Tag() quickfix.Tag { return tag.PosMaintStatus }

func NewPosMaintStatus(val enum.PosMaintStatus) PosMaintStatusField {
	return PosMaintStatusField{quickfix.FIXString(val)}
}

func (f PosMaintStatusField) Value() enum.PosMaintStatus { return enum.PosMaintStatus(f.String()) }

//PosQtyStatusField is a enum.PosQtyStatus field
type PosQtyStatusField struct{ quickfix.FIXString }

//Tag returns tag.PosQtyStatus (706)
func (f PosQtyStatusField) Tag() quickfix.Tag { return tag.PosQtyStatus }

func NewPosQtyStatus(val enum.PosQtyStatus) PosQtyStatusField {
	return PosQtyStatusField{quickfix.FIXString(val)}
}

func (f PosQtyStatusField) Value() enum.PosQtyStatus { return enum.PosQtyStatus(f.String()) }

//PosReqIDField is a STRING field
type PosReqIDField struct{ quickfix.FIXString }

//Tag returns tag.PosReqID (710)
func (f PosReqIDField) Tag() quickfix.Tag { return tag.PosReqID }

//NewPosReqID returns a new PosReqIDField initialized with val
func NewPosReqID(val string) PosReqIDField {
	return PosReqIDField{quickfix.FIXString(val)}
}

func (f PosReqIDField) Value() string { return f.String() }

//PosReqResultField is a enum.PosReqResult field
type PosReqResultField struct{ quickfix.FIXString }

//Tag returns tag.PosReqResult (728)
func (f PosReqResultField) Tag() quickfix.Tag { return tag.PosReqResult }

func NewPosReqResult(val enum.PosReqResult) PosReqResultField {
	return PosReqResultField{quickfix.FIXString(val)}
}

func (f PosReqResultField) Value() enum.PosReqResult { return enum.PosReqResult(f.String()) }

//PosReqStatusField is a enum.PosReqStatus field
type PosReqStatusField struct{ quickfix.FIXString }

//Tag returns tag.PosReqStatus (729)
func (f PosReqStatusField) Tag() quickfix.Tag { return tag.PosReqStatus }

func NewPosReqStatus(val enum.PosReqStatus) PosReqStatusField {
	return PosReqStatusField{quickfix.FIXString(val)}
}

func (f PosReqStatusField) Value() enum.PosReqStatus { return enum.PosReqStatus(f.String()) }

//PosReqTypeField is a enum.PosReqType field
type PosReqTypeField struct{ quickfix.FIXString }

//Tag returns tag.PosReqType (724)
func (f PosReqTypeField) Tag() quickfix.Tag { return tag.PosReqType }

func NewPosReqType(val enum.PosReqType) PosReqTypeField {
	return PosReqTypeField{quickfix.FIXString(val)}
}

func (f PosReqTypeField) Value() enum.PosReqType { return enum.PosReqType(f.String()) }

//PosTransTypeField is a enum.PosTransType field
type PosTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.PosTransType (709)
func (f PosTransTypeField) Tag() quickfix.Tag { return tag.PosTransType }

func NewPosTransType(val enum.PosTransType) PosTransTypeField {
	return PosTransTypeField{quickfix.FIXString(val)}
}

func (f PosTransTypeField) Value() enum.PosTransType { return enum.PosTransType(f.String()) }

//PosTypeField is a enum.PosType field
type PosTypeField struct{ quickfix.FIXString }

//Tag returns tag.PosType (703)
func (f PosTypeField) Tag() quickfix.Tag { return tag.PosType }

func NewPosType(val enum.PosType) PosTypeField {
	return PosTypeField{quickfix.FIXString(val)}
}

func (f PosTypeField) Value() enum.PosType { return enum.PosType(f.String()) }

//PositionCurrencyField is a STRING field
type PositionCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.PositionCurrency (1055)
func (f PositionCurrencyField) Tag() quickfix.Tag { return tag.PositionCurrency }

//NewPositionCurrency returns a new PositionCurrencyField initialized with val
func NewPositionCurrency(val string) PositionCurrencyField {
	return PositionCurrencyField{quickfix.FIXString(val)}
}

func (f PositionCurrencyField) Value() string { return f.String() }

//PositionEffectField is a enum.PositionEffect field
type PositionEffectField struct{ quickfix.FIXString }

//Tag returns tag.PositionEffect (77)
func (f PositionEffectField) Tag() quickfix.Tag { return tag.PositionEffect }

func NewPositionEffect(val enum.PositionEffect) PositionEffectField {
	return PositionEffectField{quickfix.FIXString(val)}
}

func (f PositionEffectField) Value() enum.PositionEffect { return enum.PositionEffect(f.String()) }

//PositionLimitField is a INT field
type PositionLimitField struct{ quickfix.FIXInt }

//Tag returns tag.PositionLimit (970)
func (f PositionLimitField) Tag() quickfix.Tag { return tag.PositionLimit }

//NewPositionLimit returns a new PositionLimitField initialized with val
func NewPositionLimit(val int) PositionLimitField {
	return PositionLimitField{quickfix.FIXInt(val)}
}

func (f PositionLimitField) Value() int { return f.Int() }

//PossDupFlagField is a BOOLEAN field
type PossDupFlagField struct{ quickfix.FIXBoolean }

//Tag returns tag.PossDupFlag (43)
func (f PossDupFlagField) Tag() quickfix.Tag { return tag.PossDupFlag }

//NewPossDupFlag returns a new PossDupFlagField initialized with val
func NewPossDupFlag(val bool) PossDupFlagField {
	return PossDupFlagField{quickfix.FIXBoolean(val)}
}

func (f PossDupFlagField) Value() bool { return f.Bool() }

//PossResendField is a BOOLEAN field
type PossResendField struct{ quickfix.FIXBoolean }

//Tag returns tag.PossResend (97)
func (f PossResendField) Tag() quickfix.Tag { return tag.PossResend }

//NewPossResend returns a new PossResendField initialized with val
func NewPossResend(val bool) PossResendField {
	return PossResendField{quickfix.FIXBoolean(val)}
}

func (f PossResendField) Value() bool { return f.Bool() }

//PreTradeAnonymityField is a BOOLEAN field
type PreTradeAnonymityField struct{ quickfix.FIXBoolean }

//Tag returns tag.PreTradeAnonymity (1091)
func (f PreTradeAnonymityField) Tag() quickfix.Tag { return tag.PreTradeAnonymity }

//NewPreTradeAnonymity returns a new PreTradeAnonymityField initialized with val
func NewPreTradeAnonymity(val bool) PreTradeAnonymityField {
	return PreTradeAnonymityField{quickfix.FIXBoolean(val)}
}

func (f PreTradeAnonymityField) Value() bool { return f.Bool() }

//PreallocMethodField is a enum.PreallocMethod field
type PreallocMethodField struct{ quickfix.FIXString }

//Tag returns tag.PreallocMethod (591)
func (f PreallocMethodField) Tag() quickfix.Tag { return tag.PreallocMethod }

func NewPreallocMethod(val enum.PreallocMethod) PreallocMethodField {
	return PreallocMethodField{quickfix.FIXString(val)}
}

func (f PreallocMethodField) Value() enum.PreallocMethod { return enum.PreallocMethod(f.String()) }

//PrevClosePxField is a PRICE field
type PrevClosePxField struct{ quickfix.FIXDecimal }

//Tag returns tag.PrevClosePx (140)
func (f PrevClosePxField) Tag() quickfix.Tag { return tag.PrevClosePx }

//NewPrevClosePx returns a new PrevClosePxField initialized with val and scale
func NewPrevClosePx(val decimal.Decimal, scale int32) PrevClosePxField {
	return PrevClosePxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f PrevClosePxField) Value() (val decimal.Decimal) { return f.Decimal }

//PreviouslyReportedField is a BOOLEAN field
type PreviouslyReportedField struct{ quickfix.FIXBoolean }

//Tag returns tag.PreviouslyReported (570)
func (f PreviouslyReportedField) Tag() quickfix.Tag { return tag.PreviouslyReported }

//NewPreviouslyReported returns a new PreviouslyReportedField initialized with val
func NewPreviouslyReported(val bool) PreviouslyReportedField {
	return PreviouslyReportedField{quickfix.FIXBoolean(val)}
}

func (f PreviouslyReportedField) Value() bool { return f.Bool() }

//PriceField is a PRICE field
type PriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.Price (44)
func (f PriceField) Tag() quickfix.Tag { return tag.Price }

//NewPrice returns a new PriceField initialized with val and scale
func NewPrice(val decimal.Decimal, scale int32) PriceField {
	return PriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f PriceField) Value() (val decimal.Decimal) { return f.Decimal }

//Price2Field is a PRICE field
type Price2Field struct{ quickfix.FIXDecimal }

//Tag returns tag.Price2 (640)
func (f Price2Field) Tag() quickfix.Tag { return tag.Price2 }

//NewPrice2 returns a new Price2Field initialized with val and scale
func NewPrice2(val decimal.Decimal, scale int32) Price2Field {
	return Price2Field{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f Price2Field) Value() (val decimal.Decimal) { return f.Decimal }

//PriceDeltaField is a FLOAT field
type PriceDeltaField struct{ quickfix.FIXDecimal }

//Tag returns tag.PriceDelta (811)
func (f PriceDeltaField) Tag() quickfix.Tag { return tag.PriceDelta }

//NewPriceDelta returns a new PriceDeltaField initialized with val and scale
func NewPriceDelta(val decimal.Decimal, scale int32) PriceDeltaField {
	return PriceDeltaField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f PriceDeltaField) Value() (val decimal.Decimal) { return f.Decimal }

//PriceImprovementField is a PRICEOFFSET field
type PriceImprovementField struct{ quickfix.FIXDecimal }

//Tag returns tag.PriceImprovement (639)
func (f PriceImprovementField) Tag() quickfix.Tag { return tag.PriceImprovement }

//NewPriceImprovement returns a new PriceImprovementField initialized with val and scale
func NewPriceImprovement(val decimal.Decimal, scale int32) PriceImprovementField {
	return PriceImprovementField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f PriceImprovementField) Value() (val decimal.Decimal) { return f.Decimal }

//PriceLimitTypeField is a enum.PriceLimitType field
type PriceLimitTypeField struct{ quickfix.FIXString }

//Tag returns tag.PriceLimitType (1306)
func (f PriceLimitTypeField) Tag() quickfix.Tag { return tag.PriceLimitType }

func NewPriceLimitType(val enum.PriceLimitType) PriceLimitTypeField {
	return PriceLimitTypeField{quickfix.FIXString(val)}
}

func (f PriceLimitTypeField) Value() enum.PriceLimitType { return enum.PriceLimitType(f.String()) }

//PriceProtectionScopeField is a enum.PriceProtectionScope field
type PriceProtectionScopeField struct{ quickfix.FIXString }

//Tag returns tag.PriceProtectionScope (1092)
func (f PriceProtectionScopeField) Tag() quickfix.Tag { return tag.PriceProtectionScope }

func NewPriceProtectionScope(val enum.PriceProtectionScope) PriceProtectionScopeField {
	return PriceProtectionScopeField{quickfix.FIXString(val)}
}

func (f PriceProtectionScopeField) Value() enum.PriceProtectionScope {
	return enum.PriceProtectionScope(f.String())
}

//PriceQuoteMethodField is a enum.PriceQuoteMethod field
type PriceQuoteMethodField struct{ quickfix.FIXString }

//Tag returns tag.PriceQuoteMethod (1196)
func (f PriceQuoteMethodField) Tag() quickfix.Tag { return tag.PriceQuoteMethod }

func NewPriceQuoteMethod(val enum.PriceQuoteMethod) PriceQuoteMethodField {
	return PriceQuoteMethodField{quickfix.FIXString(val)}
}

func (f PriceQuoteMethodField) Value() enum.PriceQuoteMethod { return enum.PriceQuoteMethod(f.String()) }

//PriceTypeField is a enum.PriceType field
type PriceTypeField struct{ quickfix.FIXString }

//Tag returns tag.PriceType (423)
func (f PriceTypeField) Tag() quickfix.Tag { return tag.PriceType }

func NewPriceType(val enum.PriceType) PriceTypeField {
	return PriceTypeField{quickfix.FIXString(val)}
}

func (f PriceTypeField) Value() enum.PriceType { return enum.PriceType(f.String()) }

//PriceUnitOfMeasureField is a STRING field
type PriceUnitOfMeasureField struct{ quickfix.FIXString }

//Tag returns tag.PriceUnitOfMeasure (1191)
func (f PriceUnitOfMeasureField) Tag() quickfix.Tag { return tag.PriceUnitOfMeasure }

//NewPriceUnitOfMeasure returns a new PriceUnitOfMeasureField initialized with val
func NewPriceUnitOfMeasure(val string) PriceUnitOfMeasureField {
	return PriceUnitOfMeasureField{quickfix.FIXString(val)}
}

func (f PriceUnitOfMeasureField) Value() string { return f.String() }

//PriceUnitOfMeasureQtyField is a QTY field
type PriceUnitOfMeasureQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.PriceUnitOfMeasureQty (1192)
func (f PriceUnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.PriceUnitOfMeasureQty }

//NewPriceUnitOfMeasureQty returns a new PriceUnitOfMeasureQtyField initialized with val and scale
func NewPriceUnitOfMeasureQty(val decimal.Decimal, scale int32) PriceUnitOfMeasureQtyField {
	return PriceUnitOfMeasureQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f PriceUnitOfMeasureQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//PriorSettlPriceField is a PRICE field
type PriorSettlPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.PriorSettlPrice (734)
func (f PriorSettlPriceField) Tag() quickfix.Tag { return tag.PriorSettlPrice }

//NewPriorSettlPrice returns a new PriorSettlPriceField initialized with val and scale
func NewPriorSettlPrice(val decimal.Decimal, scale int32) PriorSettlPriceField {
	return PriorSettlPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f PriorSettlPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//PriorSpreadIndicatorField is a BOOLEAN field
type PriorSpreadIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.PriorSpreadIndicator (720)
func (f PriorSpreadIndicatorField) Tag() quickfix.Tag { return tag.PriorSpreadIndicator }

//NewPriorSpreadIndicator returns a new PriorSpreadIndicatorField initialized with val
func NewPriorSpreadIndicator(val bool) PriorSpreadIndicatorField {
	return PriorSpreadIndicatorField{quickfix.FIXBoolean(val)}
}

func (f PriorSpreadIndicatorField) Value() bool { return f.Bool() }

//PriorityIndicatorField is a enum.PriorityIndicator field
type PriorityIndicatorField struct{ quickfix.FIXString }

//Tag returns tag.PriorityIndicator (638)
func (f PriorityIndicatorField) Tag() quickfix.Tag { return tag.PriorityIndicator }

func NewPriorityIndicator(val enum.PriorityIndicator) PriorityIndicatorField {
	return PriorityIndicatorField{quickfix.FIXString(val)}
}

func (f PriorityIndicatorField) Value() enum.PriorityIndicator {
	return enum.PriorityIndicator(f.String())
}

//PrivateQuoteField is a BOOLEAN field
type PrivateQuoteField struct{ quickfix.FIXBoolean }

//Tag returns tag.PrivateQuote (1171)
func (f PrivateQuoteField) Tag() quickfix.Tag { return tag.PrivateQuote }

//NewPrivateQuote returns a new PrivateQuoteField initialized with val
func NewPrivateQuote(val bool) PrivateQuoteField {
	return PrivateQuoteField{quickfix.FIXBoolean(val)}
}

func (f PrivateQuoteField) Value() bool { return f.Bool() }

//ProcessCodeField is a enum.ProcessCode field
type ProcessCodeField struct{ quickfix.FIXString }

//Tag returns tag.ProcessCode (81)
func (f ProcessCodeField) Tag() quickfix.Tag { return tag.ProcessCode }

func NewProcessCode(val enum.ProcessCode) ProcessCodeField {
	return ProcessCodeField{quickfix.FIXString(val)}
}

func (f ProcessCodeField) Value() enum.ProcessCode { return enum.ProcessCode(f.String()) }

//ProductField is a enum.Product field
type ProductField struct{ quickfix.FIXString }

//Tag returns tag.Product (460)
func (f ProductField) Tag() quickfix.Tag { return tag.Product }

func NewProduct(val enum.Product) ProductField {
	return ProductField{quickfix.FIXString(val)}
}

func (f ProductField) Value() enum.Product { return enum.Product(f.String()) }

//ProductComplexField is a STRING field
type ProductComplexField struct{ quickfix.FIXString }

//Tag returns tag.ProductComplex (1227)
func (f ProductComplexField) Tag() quickfix.Tag { return tag.ProductComplex }

//NewProductComplex returns a new ProductComplexField initialized with val
func NewProductComplex(val string) ProductComplexField {
	return ProductComplexField{quickfix.FIXString(val)}
}

func (f ProductComplexField) Value() string { return f.String() }

//ProgPeriodIntervalField is a INT field
type ProgPeriodIntervalField struct{ quickfix.FIXInt }

//Tag returns tag.ProgPeriodInterval (415)
func (f ProgPeriodIntervalField) Tag() quickfix.Tag { return tag.ProgPeriodInterval }

//NewProgPeriodInterval returns a new ProgPeriodIntervalField initialized with val
func NewProgPeriodInterval(val int) ProgPeriodIntervalField {
	return ProgPeriodIntervalField{quickfix.FIXInt(val)}
}

func (f ProgPeriodIntervalField) Value() int { return f.Int() }

//ProgRptReqsField is a enum.ProgRptReqs field
type ProgRptReqsField struct{ quickfix.FIXString }

//Tag returns tag.ProgRptReqs (414)
func (f ProgRptReqsField) Tag() quickfix.Tag { return tag.ProgRptReqs }

func NewProgRptReqs(val enum.ProgRptReqs) ProgRptReqsField {
	return ProgRptReqsField{quickfix.FIXString(val)}
}

func (f ProgRptReqsField) Value() enum.ProgRptReqs { return enum.ProgRptReqs(f.String()) }

//PublishTrdIndicatorField is a BOOLEAN field
type PublishTrdIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.PublishTrdIndicator (852)
func (f PublishTrdIndicatorField) Tag() quickfix.Tag { return tag.PublishTrdIndicator }

//NewPublishTrdIndicator returns a new PublishTrdIndicatorField initialized with val
func NewPublishTrdIndicator(val bool) PublishTrdIndicatorField {
	return PublishTrdIndicatorField{quickfix.FIXBoolean(val)}
}

func (f PublishTrdIndicatorField) Value() bool { return f.Bool() }

//PutOrCallField is a enum.PutOrCall field
type PutOrCallField struct{ quickfix.FIXString }

//Tag returns tag.PutOrCall (201)
func (f PutOrCallField) Tag() quickfix.Tag { return tag.PutOrCall }

func NewPutOrCall(val enum.PutOrCall) PutOrCallField {
	return PutOrCallField{quickfix.FIXString(val)}
}

func (f PutOrCallField) Value() enum.PutOrCall { return enum.PutOrCall(f.String()) }

//QtyTypeField is a enum.QtyType field
type QtyTypeField struct{ quickfix.FIXString }

//Tag returns tag.QtyType (854)
func (f QtyTypeField) Tag() quickfix.Tag { return tag.QtyType }

func NewQtyType(val enum.QtyType) QtyTypeField {
	return QtyTypeField{quickfix.FIXString(val)}
}

func (f QtyTypeField) Value() enum.QtyType { return enum.QtyType(f.String()) }

//QuantityField is a QTY field
type QuantityField struct{ quickfix.FIXDecimal }

//Tag returns tag.Quantity (53)
func (f QuantityField) Tag() quickfix.Tag { return tag.Quantity }

//NewQuantity returns a new QuantityField initialized with val and scale
func NewQuantity(val decimal.Decimal, scale int32) QuantityField {
	return QuantityField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f QuantityField) Value() (val decimal.Decimal) { return f.Decimal }

//QuantityDateField is a LOCALMKTDATE field
type QuantityDateField struct{ quickfix.FIXString }

//Tag returns tag.QuantityDate (976)
func (f QuantityDateField) Tag() quickfix.Tag { return tag.QuantityDate }

//NewQuantityDate returns a new QuantityDateField initialized with val
func NewQuantityDate(val string) QuantityDateField {
	return QuantityDateField{quickfix.FIXString(val)}
}

func (f QuantityDateField) Value() string { return f.String() }

//QuantityTypeField is a enum.QuantityType field
type QuantityTypeField struct{ quickfix.FIXString }

//Tag returns tag.QuantityType (465)
func (f QuantityTypeField) Tag() quickfix.Tag { return tag.QuantityType }

func NewQuantityType(val enum.QuantityType) QuantityTypeField {
	return QuantityTypeField{quickfix.FIXString(val)}
}

func (f QuantityTypeField) Value() enum.QuantityType { return enum.QuantityType(f.String()) }

//QuoteAckStatusField is a enum.QuoteAckStatus field
type QuoteAckStatusField struct{ quickfix.FIXString }

//Tag returns tag.QuoteAckStatus (297)
func (f QuoteAckStatusField) Tag() quickfix.Tag { return tag.QuoteAckStatus }

func NewQuoteAckStatus(val enum.QuoteAckStatus) QuoteAckStatusField {
	return QuoteAckStatusField{quickfix.FIXString(val)}
}

func (f QuoteAckStatusField) Value() enum.QuoteAckStatus { return enum.QuoteAckStatus(f.String()) }

//QuoteCancelTypeField is a enum.QuoteCancelType field
type QuoteCancelTypeField struct{ quickfix.FIXString }

//Tag returns tag.QuoteCancelType (298)
func (f QuoteCancelTypeField) Tag() quickfix.Tag { return tag.QuoteCancelType }

func NewQuoteCancelType(val enum.QuoteCancelType) QuoteCancelTypeField {
	return QuoteCancelTypeField{quickfix.FIXString(val)}
}

func (f QuoteCancelTypeField) Value() enum.QuoteCancelType { return enum.QuoteCancelType(f.String()) }

//QuoteConditionField is a enum.QuoteCondition field
type QuoteConditionField struct{ quickfix.FIXString }

//Tag returns tag.QuoteCondition (276)
func (f QuoteConditionField) Tag() quickfix.Tag { return tag.QuoteCondition }

func NewQuoteCondition(val enum.QuoteCondition) QuoteConditionField {
	return QuoteConditionField{quickfix.FIXString(val)}
}

func (f QuoteConditionField) Value() enum.QuoteCondition { return enum.QuoteCondition(f.String()) }

//QuoteEntryIDField is a STRING field
type QuoteEntryIDField struct{ quickfix.FIXString }

//Tag returns tag.QuoteEntryID (299)
func (f QuoteEntryIDField) Tag() quickfix.Tag { return tag.QuoteEntryID }

//NewQuoteEntryID returns a new QuoteEntryIDField initialized with val
func NewQuoteEntryID(val string) QuoteEntryIDField {
	return QuoteEntryIDField{quickfix.FIXString(val)}
}

func (f QuoteEntryIDField) Value() string { return f.String() }

//QuoteEntryRejectReasonField is a enum.QuoteEntryRejectReason field
type QuoteEntryRejectReasonField struct{ quickfix.FIXString }

//Tag returns tag.QuoteEntryRejectReason (368)
func (f QuoteEntryRejectReasonField) Tag() quickfix.Tag { return tag.QuoteEntryRejectReason }

func NewQuoteEntryRejectReason(val enum.QuoteEntryRejectReason) QuoteEntryRejectReasonField {
	return QuoteEntryRejectReasonField{quickfix.FIXString(val)}
}

func (f QuoteEntryRejectReasonField) Value() enum.QuoteEntryRejectReason {
	return enum.QuoteEntryRejectReason(f.String())
}

//QuoteEntryStatusField is a enum.QuoteEntryStatus field
type QuoteEntryStatusField struct{ quickfix.FIXString }

//Tag returns tag.QuoteEntryStatus (1167)
func (f QuoteEntryStatusField) Tag() quickfix.Tag { return tag.QuoteEntryStatus }

func NewQuoteEntryStatus(val enum.QuoteEntryStatus) QuoteEntryStatusField {
	return QuoteEntryStatusField{quickfix.FIXString(val)}
}

func (f QuoteEntryStatusField) Value() enum.QuoteEntryStatus { return enum.QuoteEntryStatus(f.String()) }

//QuoteIDField is a STRING field
type QuoteIDField struct{ quickfix.FIXString }

//Tag returns tag.QuoteID (117)
func (f QuoteIDField) Tag() quickfix.Tag { return tag.QuoteID }

//NewQuoteID returns a new QuoteIDField initialized with val
func NewQuoteID(val string) QuoteIDField {
	return QuoteIDField{quickfix.FIXString(val)}
}

func (f QuoteIDField) Value() string { return f.String() }

//QuoteMsgIDField is a STRING field
type QuoteMsgIDField struct{ quickfix.FIXString }

//Tag returns tag.QuoteMsgID (1166)
func (f QuoteMsgIDField) Tag() quickfix.Tag { return tag.QuoteMsgID }

//NewQuoteMsgID returns a new QuoteMsgIDField initialized with val
func NewQuoteMsgID(val string) QuoteMsgIDField {
	return QuoteMsgIDField{quickfix.FIXString(val)}
}

func (f QuoteMsgIDField) Value() string { return f.String() }

//QuotePriceTypeField is a enum.QuotePriceType field
type QuotePriceTypeField struct{ quickfix.FIXString }

//Tag returns tag.QuotePriceType (692)
func (f QuotePriceTypeField) Tag() quickfix.Tag { return tag.QuotePriceType }

func NewQuotePriceType(val enum.QuotePriceType) QuotePriceTypeField {
	return QuotePriceTypeField{quickfix.FIXString(val)}
}

func (f QuotePriceTypeField) Value() enum.QuotePriceType { return enum.QuotePriceType(f.String()) }

//QuoteQualifierField is a CHAR field
type QuoteQualifierField struct{ quickfix.FIXString }

//Tag returns tag.QuoteQualifier (695)
func (f QuoteQualifierField) Tag() quickfix.Tag { return tag.QuoteQualifier }

//NewQuoteQualifier returns a new QuoteQualifierField initialized with val
func NewQuoteQualifier(val string) QuoteQualifierField {
	return QuoteQualifierField{quickfix.FIXString(val)}
}

func (f QuoteQualifierField) Value() string { return f.String() }

//QuoteRejectReasonField is a enum.QuoteRejectReason field
type QuoteRejectReasonField struct{ quickfix.FIXString }

//Tag returns tag.QuoteRejectReason (300)
func (f QuoteRejectReasonField) Tag() quickfix.Tag { return tag.QuoteRejectReason }

func NewQuoteRejectReason(val enum.QuoteRejectReason) QuoteRejectReasonField {
	return QuoteRejectReasonField{quickfix.FIXString(val)}
}

func (f QuoteRejectReasonField) Value() enum.QuoteRejectReason {
	return enum.QuoteRejectReason(f.String())
}

//QuoteReqIDField is a STRING field
type QuoteReqIDField struct{ quickfix.FIXString }

//Tag returns tag.QuoteReqID (131)
func (f QuoteReqIDField) Tag() quickfix.Tag { return tag.QuoteReqID }

//NewQuoteReqID returns a new QuoteReqIDField initialized with val
func NewQuoteReqID(val string) QuoteReqIDField {
	return QuoteReqIDField{quickfix.FIXString(val)}
}

func (f QuoteReqIDField) Value() string { return f.String() }

//QuoteRequestRejectReasonField is a enum.QuoteRequestRejectReason field
type QuoteRequestRejectReasonField struct{ quickfix.FIXString }

//Tag returns tag.QuoteRequestRejectReason (658)
func (f QuoteRequestRejectReasonField) Tag() quickfix.Tag { return tag.QuoteRequestRejectReason }

func NewQuoteRequestRejectReason(val enum.QuoteRequestRejectReason) QuoteRequestRejectReasonField {
	return QuoteRequestRejectReasonField{quickfix.FIXString(val)}
}

func (f QuoteRequestRejectReasonField) Value() enum.QuoteRequestRejectReason {
	return enum.QuoteRequestRejectReason(f.String())
}

//QuoteRequestTypeField is a enum.QuoteRequestType field
type QuoteRequestTypeField struct{ quickfix.FIXString }

//Tag returns tag.QuoteRequestType (303)
func (f QuoteRequestTypeField) Tag() quickfix.Tag { return tag.QuoteRequestType }

func NewQuoteRequestType(val enum.QuoteRequestType) QuoteRequestTypeField {
	return QuoteRequestTypeField{quickfix.FIXString(val)}
}

func (f QuoteRequestTypeField) Value() enum.QuoteRequestType { return enum.QuoteRequestType(f.String()) }

//QuoteRespIDField is a STRING field
type QuoteRespIDField struct{ quickfix.FIXString }

//Tag returns tag.QuoteRespID (693)
func (f QuoteRespIDField) Tag() quickfix.Tag { return tag.QuoteRespID }

//NewQuoteRespID returns a new QuoteRespIDField initialized with val
func NewQuoteRespID(val string) QuoteRespIDField {
	return QuoteRespIDField{quickfix.FIXString(val)}
}

func (f QuoteRespIDField) Value() string { return f.String() }

//QuoteRespTypeField is a enum.QuoteRespType field
type QuoteRespTypeField struct{ quickfix.FIXString }

//Tag returns tag.QuoteRespType (694)
func (f QuoteRespTypeField) Tag() quickfix.Tag { return tag.QuoteRespType }

func NewQuoteRespType(val enum.QuoteRespType) QuoteRespTypeField {
	return QuoteRespTypeField{quickfix.FIXString(val)}
}

func (f QuoteRespTypeField) Value() enum.QuoteRespType { return enum.QuoteRespType(f.String()) }

//QuoteResponseLevelField is a enum.QuoteResponseLevel field
type QuoteResponseLevelField struct{ quickfix.FIXString }

//Tag returns tag.QuoteResponseLevel (301)
func (f QuoteResponseLevelField) Tag() quickfix.Tag { return tag.QuoteResponseLevel }

func NewQuoteResponseLevel(val enum.QuoteResponseLevel) QuoteResponseLevelField {
	return QuoteResponseLevelField{quickfix.FIXString(val)}
}

func (f QuoteResponseLevelField) Value() enum.QuoteResponseLevel {
	return enum.QuoteResponseLevel(f.String())
}

//QuoteSetIDField is a STRING field
type QuoteSetIDField struct{ quickfix.FIXString }

//Tag returns tag.QuoteSetID (302)
func (f QuoteSetIDField) Tag() quickfix.Tag { return tag.QuoteSetID }

//NewQuoteSetID returns a new QuoteSetIDField initialized with val
func NewQuoteSetID(val string) QuoteSetIDField {
	return QuoteSetIDField{quickfix.FIXString(val)}
}

func (f QuoteSetIDField) Value() string { return f.String() }

//QuoteSetValidUntilTimeField is a UTCTIMESTAMP field
type QuoteSetValidUntilTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.QuoteSetValidUntilTime (367)
func (f QuoteSetValidUntilTimeField) Tag() quickfix.Tag { return tag.QuoteSetValidUntilTime }

//NewQuoteSetValidUntilTime returns a new QuoteSetValidUntilTimeField initialized with val
func NewQuoteSetValidUntilTime(val time.Time) QuoteSetValidUntilTimeField {
	return QuoteSetValidUntilTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewQuoteSetValidUntilTimeNoMillis returns a new QuoteSetValidUntilTimeField initialized with val without millisecs
func NewQuoteSetValidUntilTimeNoMillis(val time.Time) QuoteSetValidUntilTimeField {
	return QuoteSetValidUntilTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f QuoteSetValidUntilTimeField) Value() time.Time { return f.Time }

//QuoteStatusField is a enum.QuoteStatus field
type QuoteStatusField struct{ quickfix.FIXString }

//Tag returns tag.QuoteStatus (297)
func (f QuoteStatusField) Tag() quickfix.Tag { return tag.QuoteStatus }

func NewQuoteStatus(val enum.QuoteStatus) QuoteStatusField {
	return QuoteStatusField{quickfix.FIXString(val)}
}

func (f QuoteStatusField) Value() enum.QuoteStatus { return enum.QuoteStatus(f.String()) }

//QuoteStatusReqIDField is a STRING field
type QuoteStatusReqIDField struct{ quickfix.FIXString }

//Tag returns tag.QuoteStatusReqID (649)
func (f QuoteStatusReqIDField) Tag() quickfix.Tag { return tag.QuoteStatusReqID }

//NewQuoteStatusReqID returns a new QuoteStatusReqIDField initialized with val
func NewQuoteStatusReqID(val string) QuoteStatusReqIDField {
	return QuoteStatusReqIDField{quickfix.FIXString(val)}
}

func (f QuoteStatusReqIDField) Value() string { return f.String() }

//QuoteTypeField is a enum.QuoteType field
type QuoteTypeField struct{ quickfix.FIXString }

//Tag returns tag.QuoteType (537)
func (f QuoteTypeField) Tag() quickfix.Tag { return tag.QuoteType }

func NewQuoteType(val enum.QuoteType) QuoteTypeField {
	return QuoteTypeField{quickfix.FIXString(val)}
}

func (f QuoteTypeField) Value() enum.QuoteType { return enum.QuoteType(f.String()) }

//RFQReqIDField is a STRING field
type RFQReqIDField struct{ quickfix.FIXString }

//Tag returns tag.RFQReqID (644)
func (f RFQReqIDField) Tag() quickfix.Tag { return tag.RFQReqID }

//NewRFQReqID returns a new RFQReqIDField initialized with val
func NewRFQReqID(val string) RFQReqIDField {
	return RFQReqIDField{quickfix.FIXString(val)}
}

func (f RFQReqIDField) Value() string { return f.String() }

//RateSourceField is a enum.RateSource field
type RateSourceField struct{ quickfix.FIXString }

//Tag returns tag.RateSource (1446)
func (f RateSourceField) Tag() quickfix.Tag { return tag.RateSource }

func NewRateSource(val enum.RateSource) RateSourceField {
	return RateSourceField{quickfix.FIXString(val)}
}

func (f RateSourceField) Value() enum.RateSource { return enum.RateSource(f.String()) }

//RateSourceTypeField is a enum.RateSourceType field
type RateSourceTypeField struct{ quickfix.FIXString }

//Tag returns tag.RateSourceType (1447)
func (f RateSourceTypeField) Tag() quickfix.Tag { return tag.RateSourceType }

func NewRateSourceType(val enum.RateSourceType) RateSourceTypeField {
	return RateSourceTypeField{quickfix.FIXString(val)}
}

func (f RateSourceTypeField) Value() enum.RateSourceType { return enum.RateSourceType(f.String()) }

//RatioQtyField is a QTY field
type RatioQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.RatioQty (319)
func (f RatioQtyField) Tag() quickfix.Tag { return tag.RatioQty }

//NewRatioQty returns a new RatioQtyField initialized with val and scale
func NewRatioQty(val decimal.Decimal, scale int32) RatioQtyField {
	return RatioQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f RatioQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//RawDataField is a DATA field
type RawDataField struct{ quickfix.FIXString }

//Tag returns tag.RawData (96)
func (f RawDataField) Tag() quickfix.Tag { return tag.RawData }

//NewRawData returns a new RawDataField initialized with val
func NewRawData(val string) RawDataField {
	return RawDataField{quickfix.FIXString(val)}
}

func (f RawDataField) Value() string { return f.String() }

//RawDataLengthField is a LENGTH field
type RawDataLengthField struct{ quickfix.FIXInt }

//Tag returns tag.RawDataLength (95)
func (f RawDataLengthField) Tag() quickfix.Tag { return tag.RawDataLength }

//NewRawDataLength returns a new RawDataLengthField initialized with val
func NewRawDataLength(val int) RawDataLengthField {
	return RawDataLengthField{quickfix.FIXInt(val)}
}

func (f RawDataLengthField) Value() int { return f.Int() }

//ReceivedDeptIDField is a STRING field
type ReceivedDeptIDField struct{ quickfix.FIXString }

//Tag returns tag.ReceivedDeptID (1030)
func (f ReceivedDeptIDField) Tag() quickfix.Tag { return tag.ReceivedDeptID }

//NewReceivedDeptID returns a new ReceivedDeptIDField initialized with val
func NewReceivedDeptID(val string) ReceivedDeptIDField {
	return ReceivedDeptIDField{quickfix.FIXString(val)}
}

func (f ReceivedDeptIDField) Value() string { return f.String() }

//RedemptionDateField is a LOCALMKTDATE field
type RedemptionDateField struct{ quickfix.FIXString }

//Tag returns tag.RedemptionDate (240)
func (f RedemptionDateField) Tag() quickfix.Tag { return tag.RedemptionDate }

//NewRedemptionDate returns a new RedemptionDateField initialized with val
func NewRedemptionDate(val string) RedemptionDateField {
	return RedemptionDateField{quickfix.FIXString(val)}
}

func (f RedemptionDateField) Value() string { return f.String() }

//RefAllocIDField is a STRING field
type RefAllocIDField struct{ quickfix.FIXString }

//Tag returns tag.RefAllocID (72)
func (f RefAllocIDField) Tag() quickfix.Tag { return tag.RefAllocID }

//NewRefAllocID returns a new RefAllocIDField initialized with val
func NewRefAllocID(val string) RefAllocIDField {
	return RefAllocIDField{quickfix.FIXString(val)}
}

func (f RefAllocIDField) Value() string { return f.String() }

//RefApplExtIDField is a INT field
type RefApplExtIDField struct{ quickfix.FIXInt }

//Tag returns tag.RefApplExtID (1406)
func (f RefApplExtIDField) Tag() quickfix.Tag { return tag.RefApplExtID }

//NewRefApplExtID returns a new RefApplExtIDField initialized with val
func NewRefApplExtID(val int) RefApplExtIDField {
	return RefApplExtIDField{quickfix.FIXInt(val)}
}

func (f RefApplExtIDField) Value() int { return f.Int() }

//RefApplIDField is a STRING field
type RefApplIDField struct{ quickfix.FIXString }

//Tag returns tag.RefApplID (1355)
func (f RefApplIDField) Tag() quickfix.Tag { return tag.RefApplID }

//NewRefApplID returns a new RefApplIDField initialized with val
func NewRefApplID(val string) RefApplIDField {
	return RefApplIDField{quickfix.FIXString(val)}
}

func (f RefApplIDField) Value() string { return f.String() }

//RefApplLastSeqNumField is a SEQNUM field
type RefApplLastSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.RefApplLastSeqNum (1357)
func (f RefApplLastSeqNumField) Tag() quickfix.Tag { return tag.RefApplLastSeqNum }

//NewRefApplLastSeqNum returns a new RefApplLastSeqNumField initialized with val
func NewRefApplLastSeqNum(val int) RefApplLastSeqNumField {
	return RefApplLastSeqNumField{quickfix.FIXInt(val)}
}

func (f RefApplLastSeqNumField) Value() int { return f.Int() }

//RefApplReqIDField is a STRING field
type RefApplReqIDField struct{ quickfix.FIXString }

//Tag returns tag.RefApplReqID (1433)
func (f RefApplReqIDField) Tag() quickfix.Tag { return tag.RefApplReqID }

//NewRefApplReqID returns a new RefApplReqIDField initialized with val
func NewRefApplReqID(val string) RefApplReqIDField {
	return RefApplReqIDField{quickfix.FIXString(val)}
}

func (f RefApplReqIDField) Value() string { return f.String() }

//RefApplVerIDField is a STRING field
type RefApplVerIDField struct{ quickfix.FIXString }

//Tag returns tag.RefApplVerID (1130)
func (f RefApplVerIDField) Tag() quickfix.Tag { return tag.RefApplVerID }

//NewRefApplVerID returns a new RefApplVerIDField initialized with val
func NewRefApplVerID(val string) RefApplVerIDField {
	return RefApplVerIDField{quickfix.FIXString(val)}
}

func (f RefApplVerIDField) Value() string { return f.String() }

//RefCompIDField is a STRING field
type RefCompIDField struct{ quickfix.FIXString }

//Tag returns tag.RefCompID (930)
func (f RefCompIDField) Tag() quickfix.Tag { return tag.RefCompID }

//NewRefCompID returns a new RefCompIDField initialized with val
func NewRefCompID(val string) RefCompIDField {
	return RefCompIDField{quickfix.FIXString(val)}
}

func (f RefCompIDField) Value() string { return f.String() }

//RefCstmApplVerIDField is a STRING field
type RefCstmApplVerIDField struct{ quickfix.FIXString }

//Tag returns tag.RefCstmApplVerID (1131)
func (f RefCstmApplVerIDField) Tag() quickfix.Tag { return tag.RefCstmApplVerID }

//NewRefCstmApplVerID returns a new RefCstmApplVerIDField initialized with val
func NewRefCstmApplVerID(val string) RefCstmApplVerIDField {
	return RefCstmApplVerIDField{quickfix.FIXString(val)}
}

func (f RefCstmApplVerIDField) Value() string { return f.String() }

//RefMsgTypeField is a STRING field
type RefMsgTypeField struct{ quickfix.FIXString }

//Tag returns tag.RefMsgType (372)
func (f RefMsgTypeField) Tag() quickfix.Tag { return tag.RefMsgType }

//NewRefMsgType returns a new RefMsgTypeField initialized with val
func NewRefMsgType(val string) RefMsgTypeField {
	return RefMsgTypeField{quickfix.FIXString(val)}
}

func (f RefMsgTypeField) Value() string { return f.String() }

//RefOrdIDReasonField is a enum.RefOrdIDReason field
type RefOrdIDReasonField struct{ quickfix.FIXString }

//Tag returns tag.RefOrdIDReason (1431)
func (f RefOrdIDReasonField) Tag() quickfix.Tag { return tag.RefOrdIDReason }

func NewRefOrdIDReason(val enum.RefOrdIDReason) RefOrdIDReasonField {
	return RefOrdIDReasonField{quickfix.FIXString(val)}
}

func (f RefOrdIDReasonField) Value() enum.RefOrdIDReason { return enum.RefOrdIDReason(f.String()) }

//RefOrderIDField is a STRING field
type RefOrderIDField struct{ quickfix.FIXString }

//Tag returns tag.RefOrderID (1080)
func (f RefOrderIDField) Tag() quickfix.Tag { return tag.RefOrderID }

//NewRefOrderID returns a new RefOrderIDField initialized with val
func NewRefOrderID(val string) RefOrderIDField {
	return RefOrderIDField{quickfix.FIXString(val)}
}

func (f RefOrderIDField) Value() string { return f.String() }

//RefOrderIDSourceField is a enum.RefOrderIDSource field
type RefOrderIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RefOrderIDSource (1081)
func (f RefOrderIDSourceField) Tag() quickfix.Tag { return tag.RefOrderIDSource }

func NewRefOrderIDSource(val enum.RefOrderIDSource) RefOrderIDSourceField {
	return RefOrderIDSourceField{quickfix.FIXString(val)}
}

func (f RefOrderIDSourceField) Value() enum.RefOrderIDSource { return enum.RefOrderIDSource(f.String()) }

//RefSeqNumField is a SEQNUM field
type RefSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.RefSeqNum (45)
func (f RefSeqNumField) Tag() quickfix.Tag { return tag.RefSeqNum }

//NewRefSeqNum returns a new RefSeqNumField initialized with val
func NewRefSeqNum(val int) RefSeqNumField {
	return RefSeqNumField{quickfix.FIXInt(val)}
}

func (f RefSeqNumField) Value() int { return f.Int() }

//RefSubIDField is a STRING field
type RefSubIDField struct{ quickfix.FIXString }

//Tag returns tag.RefSubID (931)
func (f RefSubIDField) Tag() quickfix.Tag { return tag.RefSubID }

//NewRefSubID returns a new RefSubIDField initialized with val
func NewRefSubID(val string) RefSubIDField {
	return RefSubIDField{quickfix.FIXString(val)}
}

func (f RefSubIDField) Value() string { return f.String() }

//RefTagIDField is a INT field
type RefTagIDField struct{ quickfix.FIXInt }

//Tag returns tag.RefTagID (371)
func (f RefTagIDField) Tag() quickfix.Tag { return tag.RefTagID }

//NewRefTagID returns a new RefTagIDField initialized with val
func NewRefTagID(val int) RefTagIDField {
	return RefTagIDField{quickfix.FIXInt(val)}
}

func (f RefTagIDField) Value() int { return f.Int() }

//ReferencePageField is a STRING field
type ReferencePageField struct{ quickfix.FIXString }

//Tag returns tag.ReferencePage (1448)
func (f ReferencePageField) Tag() quickfix.Tag { return tag.ReferencePage }

//NewReferencePage returns a new ReferencePageField initialized with val
func NewReferencePage(val string) ReferencePageField {
	return ReferencePageField{quickfix.FIXString(val)}
}

func (f ReferencePageField) Value() string { return f.String() }

//RefreshIndicatorField is a BOOLEAN field
type RefreshIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.RefreshIndicator (1187)
func (f RefreshIndicatorField) Tag() quickfix.Tag { return tag.RefreshIndicator }

//NewRefreshIndicator returns a new RefreshIndicatorField initialized with val
func NewRefreshIndicator(val bool) RefreshIndicatorField {
	return RefreshIndicatorField{quickfix.FIXBoolean(val)}
}

func (f RefreshIndicatorField) Value() bool { return f.Bool() }

//RefreshQtyField is a QTY field
type RefreshQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.RefreshQty (1088)
func (f RefreshQtyField) Tag() quickfix.Tag { return tag.RefreshQty }

//NewRefreshQty returns a new RefreshQtyField initialized with val and scale
func NewRefreshQty(val decimal.Decimal, scale int32) RefreshQtyField {
	return RefreshQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f RefreshQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//RegistAcctTypeField is a STRING field
type RegistAcctTypeField struct{ quickfix.FIXString }

//Tag returns tag.RegistAcctType (493)
func (f RegistAcctTypeField) Tag() quickfix.Tag { return tag.RegistAcctType }

//NewRegistAcctType returns a new RegistAcctTypeField initialized with val
func NewRegistAcctType(val string) RegistAcctTypeField {
	return RegistAcctTypeField{quickfix.FIXString(val)}
}

func (f RegistAcctTypeField) Value() string { return f.String() }

//RegistDetlsField is a STRING field
type RegistDetlsField struct{ quickfix.FIXString }

//Tag returns tag.RegistDetls (509)
func (f RegistDetlsField) Tag() quickfix.Tag { return tag.RegistDetls }

//NewRegistDetls returns a new RegistDetlsField initialized with val
func NewRegistDetls(val string) RegistDetlsField {
	return RegistDetlsField{quickfix.FIXString(val)}
}

func (f RegistDetlsField) Value() string { return f.String() }

//RegistDtlsField is a STRING field
type RegistDtlsField struct{ quickfix.FIXString }

//Tag returns tag.RegistDtls (509)
func (f RegistDtlsField) Tag() quickfix.Tag { return tag.RegistDtls }

//NewRegistDtls returns a new RegistDtlsField initialized with val
func NewRegistDtls(val string) RegistDtlsField {
	return RegistDtlsField{quickfix.FIXString(val)}
}

func (f RegistDtlsField) Value() string { return f.String() }

//RegistEmailField is a STRING field
type RegistEmailField struct{ quickfix.FIXString }

//Tag returns tag.RegistEmail (511)
func (f RegistEmailField) Tag() quickfix.Tag { return tag.RegistEmail }

//NewRegistEmail returns a new RegistEmailField initialized with val
func NewRegistEmail(val string) RegistEmailField {
	return RegistEmailField{quickfix.FIXString(val)}
}

func (f RegistEmailField) Value() string { return f.String() }

//RegistIDField is a STRING field
type RegistIDField struct{ quickfix.FIXString }

//Tag returns tag.RegistID (513)
func (f RegistIDField) Tag() quickfix.Tag { return tag.RegistID }

//NewRegistID returns a new RegistIDField initialized with val
func NewRegistID(val string) RegistIDField {
	return RegistIDField{quickfix.FIXString(val)}
}

func (f RegistIDField) Value() string { return f.String() }

//RegistRefIDField is a STRING field
type RegistRefIDField struct{ quickfix.FIXString }

//Tag returns tag.RegistRefID (508)
func (f RegistRefIDField) Tag() quickfix.Tag { return tag.RegistRefID }

//NewRegistRefID returns a new RegistRefIDField initialized with val
func NewRegistRefID(val string) RegistRefIDField {
	return RegistRefIDField{quickfix.FIXString(val)}
}

func (f RegistRefIDField) Value() string { return f.String() }

//RegistRejReasonCodeField is a enum.RegistRejReasonCode field
type RegistRejReasonCodeField struct{ quickfix.FIXString }

//Tag returns tag.RegistRejReasonCode (507)
func (f RegistRejReasonCodeField) Tag() quickfix.Tag { return tag.RegistRejReasonCode }

func NewRegistRejReasonCode(val enum.RegistRejReasonCode) RegistRejReasonCodeField {
	return RegistRejReasonCodeField{quickfix.FIXString(val)}
}

func (f RegistRejReasonCodeField) Value() enum.RegistRejReasonCode {
	return enum.RegistRejReasonCode(f.String())
}

//RegistRejReasonTextField is a STRING field
type RegistRejReasonTextField struct{ quickfix.FIXString }

//Tag returns tag.RegistRejReasonText (496)
func (f RegistRejReasonTextField) Tag() quickfix.Tag { return tag.RegistRejReasonText }

//NewRegistRejReasonText returns a new RegistRejReasonTextField initialized with val
func NewRegistRejReasonText(val string) RegistRejReasonTextField {
	return RegistRejReasonTextField{quickfix.FIXString(val)}
}

func (f RegistRejReasonTextField) Value() string { return f.String() }

//RegistStatusField is a enum.RegistStatus field
type RegistStatusField struct{ quickfix.FIXString }

//Tag returns tag.RegistStatus (506)
func (f RegistStatusField) Tag() quickfix.Tag { return tag.RegistStatus }

func NewRegistStatus(val enum.RegistStatus) RegistStatusField {
	return RegistStatusField{quickfix.FIXString(val)}
}

func (f RegistStatusField) Value() enum.RegistStatus { return enum.RegistStatus(f.String()) }

//RegistTransTypeField is a enum.RegistTransType field
type RegistTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.RegistTransType (514)
func (f RegistTransTypeField) Tag() quickfix.Tag { return tag.RegistTransType }

func NewRegistTransType(val enum.RegistTransType) RegistTransTypeField {
	return RegistTransTypeField{quickfix.FIXString(val)}
}

func (f RegistTransTypeField) Value() enum.RegistTransType { return enum.RegistTransType(f.String()) }

//RejectTextField is a STRING field
type RejectTextField struct{ quickfix.FIXString }

//Tag returns tag.RejectText (1328)
func (f RejectTextField) Tag() quickfix.Tag { return tag.RejectText }

//NewRejectText returns a new RejectTextField initialized with val
func NewRejectText(val string) RejectTextField {
	return RejectTextField{quickfix.FIXString(val)}
}

func (f RejectTextField) Value() string { return f.String() }

//RelSymTransactTimeField is a UTCTIMESTAMP field
type RelSymTransactTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.RelSymTransactTime (1504)
func (f RelSymTransactTimeField) Tag() quickfix.Tag { return tag.RelSymTransactTime }

//NewRelSymTransactTime returns a new RelSymTransactTimeField initialized with val
func NewRelSymTransactTime(val time.Time) RelSymTransactTimeField {
	return RelSymTransactTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewRelSymTransactTimeNoMillis returns a new RelSymTransactTimeField initialized with val without millisecs
func NewRelSymTransactTimeNoMillis(val time.Time) RelSymTransactTimeField {
	return RelSymTransactTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f RelSymTransactTimeField) Value() time.Time { return f.Time }

//RelatdSymField is a STRING field
type RelatdSymField struct{ quickfix.FIXString }

//Tag returns tag.RelatdSym (46)
func (f RelatdSymField) Tag() quickfix.Tag { return tag.RelatdSym }

//NewRelatdSym returns a new RelatdSymField initialized with val
func NewRelatdSym(val string) RelatdSymField {
	return RelatdSymField{quickfix.FIXString(val)}
}

func (f RelatdSymField) Value() string { return f.String() }

//RelatedContextPartyIDField is a STRING field
type RelatedContextPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.RelatedContextPartyID (1576)
func (f RelatedContextPartyIDField) Tag() quickfix.Tag { return tag.RelatedContextPartyID }

//NewRelatedContextPartyID returns a new RelatedContextPartyIDField initialized with val
func NewRelatedContextPartyID(val string) RelatedContextPartyIDField {
	return RelatedContextPartyIDField{quickfix.FIXString(val)}
}

func (f RelatedContextPartyIDField) Value() string { return f.String() }

//RelatedContextPartyIDSourceField is a CHAR field
type RelatedContextPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RelatedContextPartyIDSource (1577)
func (f RelatedContextPartyIDSourceField) Tag() quickfix.Tag { return tag.RelatedContextPartyIDSource }

//NewRelatedContextPartyIDSource returns a new RelatedContextPartyIDSourceField initialized with val
func NewRelatedContextPartyIDSource(val string) RelatedContextPartyIDSourceField {
	return RelatedContextPartyIDSourceField{quickfix.FIXString(val)}
}

func (f RelatedContextPartyIDSourceField) Value() string { return f.String() }

//RelatedContextPartyRoleField is a INT field
type RelatedContextPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.RelatedContextPartyRole (1578)
func (f RelatedContextPartyRoleField) Tag() quickfix.Tag { return tag.RelatedContextPartyRole }

//NewRelatedContextPartyRole returns a new RelatedContextPartyRoleField initialized with val
func NewRelatedContextPartyRole(val int) RelatedContextPartyRoleField {
	return RelatedContextPartyRoleField{quickfix.FIXInt(val)}
}

func (f RelatedContextPartyRoleField) Value() int { return f.Int() }

//RelatedContextPartySubIDField is a STRING field
type RelatedContextPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.RelatedContextPartySubID (1580)
func (f RelatedContextPartySubIDField) Tag() quickfix.Tag { return tag.RelatedContextPartySubID }

//NewRelatedContextPartySubID returns a new RelatedContextPartySubIDField initialized with val
func NewRelatedContextPartySubID(val string) RelatedContextPartySubIDField {
	return RelatedContextPartySubIDField{quickfix.FIXString(val)}
}

func (f RelatedContextPartySubIDField) Value() string { return f.String() }

//RelatedContextPartySubIDTypeField is a INT field
type RelatedContextPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.RelatedContextPartySubIDType (1581)
func (f RelatedContextPartySubIDTypeField) Tag() quickfix.Tag { return tag.RelatedContextPartySubIDType }

//NewRelatedContextPartySubIDType returns a new RelatedContextPartySubIDTypeField initialized with val
func NewRelatedContextPartySubIDType(val int) RelatedContextPartySubIDTypeField {
	return RelatedContextPartySubIDTypeField{quickfix.FIXInt(val)}
}

func (f RelatedContextPartySubIDTypeField) Value() int { return f.Int() }

//RelatedPartyAltIDField is a STRING field
type RelatedPartyAltIDField struct{ quickfix.FIXString }

//Tag returns tag.RelatedPartyAltID (1570)
func (f RelatedPartyAltIDField) Tag() quickfix.Tag { return tag.RelatedPartyAltID }

//NewRelatedPartyAltID returns a new RelatedPartyAltIDField initialized with val
func NewRelatedPartyAltID(val string) RelatedPartyAltIDField {
	return RelatedPartyAltIDField{quickfix.FIXString(val)}
}

func (f RelatedPartyAltIDField) Value() string { return f.String() }

//RelatedPartyAltIDSourceField is a CHAR field
type RelatedPartyAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RelatedPartyAltIDSource (1571)
func (f RelatedPartyAltIDSourceField) Tag() quickfix.Tag { return tag.RelatedPartyAltIDSource }

//NewRelatedPartyAltIDSource returns a new RelatedPartyAltIDSourceField initialized with val
func NewRelatedPartyAltIDSource(val string) RelatedPartyAltIDSourceField {
	return RelatedPartyAltIDSourceField{quickfix.FIXString(val)}
}

func (f RelatedPartyAltIDSourceField) Value() string { return f.String() }

//RelatedPartyAltSubIDField is a STRING field
type RelatedPartyAltSubIDField struct{ quickfix.FIXString }

//Tag returns tag.RelatedPartyAltSubID (1573)
func (f RelatedPartyAltSubIDField) Tag() quickfix.Tag { return tag.RelatedPartyAltSubID }

//NewRelatedPartyAltSubID returns a new RelatedPartyAltSubIDField initialized with val
func NewRelatedPartyAltSubID(val string) RelatedPartyAltSubIDField {
	return RelatedPartyAltSubIDField{quickfix.FIXString(val)}
}

func (f RelatedPartyAltSubIDField) Value() string { return f.String() }

//RelatedPartyAltSubIDTypeField is a INT field
type RelatedPartyAltSubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.RelatedPartyAltSubIDType (1574)
func (f RelatedPartyAltSubIDTypeField) Tag() quickfix.Tag { return tag.RelatedPartyAltSubIDType }

//NewRelatedPartyAltSubIDType returns a new RelatedPartyAltSubIDTypeField initialized with val
func NewRelatedPartyAltSubIDType(val int) RelatedPartyAltSubIDTypeField {
	return RelatedPartyAltSubIDTypeField{quickfix.FIXInt(val)}
}

func (f RelatedPartyAltSubIDTypeField) Value() int { return f.Int() }

//RelatedPartyIDField is a STRING field
type RelatedPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.RelatedPartyID (1563)
func (f RelatedPartyIDField) Tag() quickfix.Tag { return tag.RelatedPartyID }

//NewRelatedPartyID returns a new RelatedPartyIDField initialized with val
func NewRelatedPartyID(val string) RelatedPartyIDField {
	return RelatedPartyIDField{quickfix.FIXString(val)}
}

func (f RelatedPartyIDField) Value() string { return f.String() }

//RelatedPartyIDSourceField is a CHAR field
type RelatedPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RelatedPartyIDSource (1564)
func (f RelatedPartyIDSourceField) Tag() quickfix.Tag { return tag.RelatedPartyIDSource }

//NewRelatedPartyIDSource returns a new RelatedPartyIDSourceField initialized with val
func NewRelatedPartyIDSource(val string) RelatedPartyIDSourceField {
	return RelatedPartyIDSourceField{quickfix.FIXString(val)}
}

func (f RelatedPartyIDSourceField) Value() string { return f.String() }

//RelatedPartyRoleField is a INT field
type RelatedPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.RelatedPartyRole (1565)
func (f RelatedPartyRoleField) Tag() quickfix.Tag { return tag.RelatedPartyRole }

//NewRelatedPartyRole returns a new RelatedPartyRoleField initialized with val
func NewRelatedPartyRole(val int) RelatedPartyRoleField {
	return RelatedPartyRoleField{quickfix.FIXInt(val)}
}

func (f RelatedPartyRoleField) Value() int { return f.Int() }

//RelatedPartySubIDField is a STRING field
type RelatedPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.RelatedPartySubID (1567)
func (f RelatedPartySubIDField) Tag() quickfix.Tag { return tag.RelatedPartySubID }

//NewRelatedPartySubID returns a new RelatedPartySubIDField initialized with val
func NewRelatedPartySubID(val string) RelatedPartySubIDField {
	return RelatedPartySubIDField{quickfix.FIXString(val)}
}

func (f RelatedPartySubIDField) Value() string { return f.String() }

//RelatedPartySubIDTypeField is a INT field
type RelatedPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.RelatedPartySubIDType (1568)
func (f RelatedPartySubIDTypeField) Tag() quickfix.Tag { return tag.RelatedPartySubIDType }

//NewRelatedPartySubIDType returns a new RelatedPartySubIDTypeField initialized with val
func NewRelatedPartySubIDType(val int) RelatedPartySubIDTypeField {
	return RelatedPartySubIDTypeField{quickfix.FIXInt(val)}
}

func (f RelatedPartySubIDTypeField) Value() int { return f.Int() }

//RelationshipRiskCFICodeField is a STRING field
type RelationshipRiskCFICodeField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskCFICode (1599)
func (f RelationshipRiskCFICodeField) Tag() quickfix.Tag { return tag.RelationshipRiskCFICode }

//NewRelationshipRiskCFICode returns a new RelationshipRiskCFICodeField initialized with val
func NewRelationshipRiskCFICode(val string) RelationshipRiskCFICodeField {
	return RelationshipRiskCFICodeField{quickfix.FIXString(val)}
}

func (f RelationshipRiskCFICodeField) Value() string { return f.String() }

//RelationshipRiskCouponRateField is a PERCENTAGE field
type RelationshipRiskCouponRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.RelationshipRiskCouponRate (1608)
func (f RelationshipRiskCouponRateField) Tag() quickfix.Tag { return tag.RelationshipRiskCouponRate }

//NewRelationshipRiskCouponRate returns a new RelationshipRiskCouponRateField initialized with val and scale
func NewRelationshipRiskCouponRate(val decimal.Decimal, scale int32) RelationshipRiskCouponRateField {
	return RelationshipRiskCouponRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f RelationshipRiskCouponRateField) Value() (val decimal.Decimal) { return f.Decimal }

//RelationshipRiskEncodedSecurityDescField is a DATA field
type RelationshipRiskEncodedSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskEncodedSecurityDesc (1619)
func (f RelationshipRiskEncodedSecurityDescField) Tag() quickfix.Tag {
	return tag.RelationshipRiskEncodedSecurityDesc
}

//NewRelationshipRiskEncodedSecurityDesc returns a new RelationshipRiskEncodedSecurityDescField initialized with val
func NewRelationshipRiskEncodedSecurityDesc(val string) RelationshipRiskEncodedSecurityDescField {
	return RelationshipRiskEncodedSecurityDescField{quickfix.FIXString(val)}
}

func (f RelationshipRiskEncodedSecurityDescField) Value() string { return f.String() }

//RelationshipRiskEncodedSecurityDescLenField is a LENGTH field
type RelationshipRiskEncodedSecurityDescLenField struct{ quickfix.FIXInt }

//Tag returns tag.RelationshipRiskEncodedSecurityDescLen (1618)
func (f RelationshipRiskEncodedSecurityDescLenField) Tag() quickfix.Tag {
	return tag.RelationshipRiskEncodedSecurityDescLen
}

//NewRelationshipRiskEncodedSecurityDescLen returns a new RelationshipRiskEncodedSecurityDescLenField initialized with val
func NewRelationshipRiskEncodedSecurityDescLen(val int) RelationshipRiskEncodedSecurityDescLenField {
	return RelationshipRiskEncodedSecurityDescLenField{quickfix.FIXInt(val)}
}

func (f RelationshipRiskEncodedSecurityDescLenField) Value() int { return f.Int() }

//RelationshipRiskFlexibleIndicatorField is a BOOLEAN field
type RelationshipRiskFlexibleIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.RelationshipRiskFlexibleIndicator (1607)
func (f RelationshipRiskFlexibleIndicatorField) Tag() quickfix.Tag {
	return tag.RelationshipRiskFlexibleIndicator
}

//NewRelationshipRiskFlexibleIndicator returns a new RelationshipRiskFlexibleIndicatorField initialized with val
func NewRelationshipRiskFlexibleIndicator(val bool) RelationshipRiskFlexibleIndicatorField {
	return RelationshipRiskFlexibleIndicatorField{quickfix.FIXBoolean(val)}
}

func (f RelationshipRiskFlexibleIndicatorField) Value() bool { return f.Bool() }

//RelationshipRiskInstrumentMultiplierField is a FLOAT field
type RelationshipRiskInstrumentMultiplierField struct{ quickfix.FIXDecimal }

//Tag returns tag.RelationshipRiskInstrumentMultiplier (1612)
func (f RelationshipRiskInstrumentMultiplierField) Tag() quickfix.Tag {
	return tag.RelationshipRiskInstrumentMultiplier
}

//NewRelationshipRiskInstrumentMultiplier returns a new RelationshipRiskInstrumentMultiplierField initialized with val and scale
func NewRelationshipRiskInstrumentMultiplier(val decimal.Decimal, scale int32) RelationshipRiskInstrumentMultiplierField {
	return RelationshipRiskInstrumentMultiplierField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f RelationshipRiskInstrumentMultiplierField) Value() (val decimal.Decimal) { return f.Decimal }

//RelationshipRiskInstrumentOperatorField is a INT field
type RelationshipRiskInstrumentOperatorField struct{ quickfix.FIXInt }

//Tag returns tag.RelationshipRiskInstrumentOperator (1588)
func (f RelationshipRiskInstrumentOperatorField) Tag() quickfix.Tag {
	return tag.RelationshipRiskInstrumentOperator
}

//NewRelationshipRiskInstrumentOperator returns a new RelationshipRiskInstrumentOperatorField initialized with val
func NewRelationshipRiskInstrumentOperator(val int) RelationshipRiskInstrumentOperatorField {
	return RelationshipRiskInstrumentOperatorField{quickfix.FIXInt(val)}
}

func (f RelationshipRiskInstrumentOperatorField) Value() int { return f.Int() }

//RelationshipRiskInstrumentSettlTypeField is a STRING field
type RelationshipRiskInstrumentSettlTypeField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskInstrumentSettlType (1611)
func (f RelationshipRiskInstrumentSettlTypeField) Tag() quickfix.Tag {
	return tag.RelationshipRiskInstrumentSettlType
}

//NewRelationshipRiskInstrumentSettlType returns a new RelationshipRiskInstrumentSettlTypeField initialized with val
func NewRelationshipRiskInstrumentSettlType(val string) RelationshipRiskInstrumentSettlTypeField {
	return RelationshipRiskInstrumentSettlTypeField{quickfix.FIXString(val)}
}

func (f RelationshipRiskInstrumentSettlTypeField) Value() string { return f.String() }

//RelationshipRiskLimitAmountField is a AMT field
type RelationshipRiskLimitAmountField struct{ quickfix.FIXDecimal }

//Tag returns tag.RelationshipRiskLimitAmount (1584)
func (f RelationshipRiskLimitAmountField) Tag() quickfix.Tag { return tag.RelationshipRiskLimitAmount }

//NewRelationshipRiskLimitAmount returns a new RelationshipRiskLimitAmountField initialized with val and scale
func NewRelationshipRiskLimitAmount(val decimal.Decimal, scale int32) RelationshipRiskLimitAmountField {
	return RelationshipRiskLimitAmountField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f RelationshipRiskLimitAmountField) Value() (val decimal.Decimal) { return f.Decimal }

//RelationshipRiskLimitCurrencyField is a CURRENCY field
type RelationshipRiskLimitCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskLimitCurrency (1585)
func (f RelationshipRiskLimitCurrencyField) Tag() quickfix.Tag {
	return tag.RelationshipRiskLimitCurrency
}

//NewRelationshipRiskLimitCurrency returns a new RelationshipRiskLimitCurrencyField initialized with val
func NewRelationshipRiskLimitCurrency(val string) RelationshipRiskLimitCurrencyField {
	return RelationshipRiskLimitCurrencyField{quickfix.FIXString(val)}
}

func (f RelationshipRiskLimitCurrencyField) Value() string { return f.String() }

//RelationshipRiskLimitPlatformField is a STRING field
type RelationshipRiskLimitPlatformField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskLimitPlatform (1586)
func (f RelationshipRiskLimitPlatformField) Tag() quickfix.Tag {
	return tag.RelationshipRiskLimitPlatform
}

//NewRelationshipRiskLimitPlatform returns a new RelationshipRiskLimitPlatformField initialized with val
func NewRelationshipRiskLimitPlatform(val string) RelationshipRiskLimitPlatformField {
	return RelationshipRiskLimitPlatformField{quickfix.FIXString(val)}
}

func (f RelationshipRiskLimitPlatformField) Value() string { return f.String() }

//RelationshipRiskLimitTypeField is a INT field
type RelationshipRiskLimitTypeField struct{ quickfix.FIXInt }

//Tag returns tag.RelationshipRiskLimitType (1583)
func (f RelationshipRiskLimitTypeField) Tag() quickfix.Tag { return tag.RelationshipRiskLimitType }

//NewRelationshipRiskLimitType returns a new RelationshipRiskLimitTypeField initialized with val
func NewRelationshipRiskLimitType(val int) RelationshipRiskLimitTypeField {
	return RelationshipRiskLimitTypeField{quickfix.FIXInt(val)}
}

func (f RelationshipRiskLimitTypeField) Value() int { return f.Int() }

//RelationshipRiskMaturityMonthYearField is a MONTHYEAR field
type RelationshipRiskMaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskMaturityMonthYear (1602)
func (f RelationshipRiskMaturityMonthYearField) Tag() quickfix.Tag {
	return tag.RelationshipRiskMaturityMonthYear
}

//NewRelationshipRiskMaturityMonthYear returns a new RelationshipRiskMaturityMonthYearField initialized with val
func NewRelationshipRiskMaturityMonthYear(val string) RelationshipRiskMaturityMonthYearField {
	return RelationshipRiskMaturityMonthYearField{quickfix.FIXString(val)}
}

func (f RelationshipRiskMaturityMonthYearField) Value() string { return f.String() }

//RelationshipRiskMaturityTimeField is a TZTIMEONLY field
type RelationshipRiskMaturityTimeField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskMaturityTime (1603)
func (f RelationshipRiskMaturityTimeField) Tag() quickfix.Tag { return tag.RelationshipRiskMaturityTime }

//NewRelationshipRiskMaturityTime returns a new RelationshipRiskMaturityTimeField initialized with val
func NewRelationshipRiskMaturityTime(val string) RelationshipRiskMaturityTimeField {
	return RelationshipRiskMaturityTimeField{quickfix.FIXString(val)}
}

func (f RelationshipRiskMaturityTimeField) Value() string { return f.String() }

//RelationshipRiskProductField is a INT field
type RelationshipRiskProductField struct{ quickfix.FIXInt }

//Tag returns tag.RelationshipRiskProduct (1596)
func (f RelationshipRiskProductField) Tag() quickfix.Tag { return tag.RelationshipRiskProduct }

//NewRelationshipRiskProduct returns a new RelationshipRiskProductField initialized with val
func NewRelationshipRiskProduct(val int) RelationshipRiskProductField {
	return RelationshipRiskProductField{quickfix.FIXInt(val)}
}

func (f RelationshipRiskProductField) Value() int { return f.Int() }

//RelationshipRiskProductComplexField is a STRING field
type RelationshipRiskProductComplexField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskProductComplex (1597)
func (f RelationshipRiskProductComplexField) Tag() quickfix.Tag {
	return tag.RelationshipRiskProductComplex
}

//NewRelationshipRiskProductComplex returns a new RelationshipRiskProductComplexField initialized with val
func NewRelationshipRiskProductComplex(val string) RelationshipRiskProductComplexField {
	return RelationshipRiskProductComplexField{quickfix.FIXString(val)}
}

func (f RelationshipRiskProductComplexField) Value() string { return f.String() }

//RelationshipRiskPutOrCallField is a INT field
type RelationshipRiskPutOrCallField struct{ quickfix.FIXInt }

//Tag returns tag.RelationshipRiskPutOrCall (1606)
func (f RelationshipRiskPutOrCallField) Tag() quickfix.Tag { return tag.RelationshipRiskPutOrCall }

//NewRelationshipRiskPutOrCall returns a new RelationshipRiskPutOrCallField initialized with val
func NewRelationshipRiskPutOrCall(val int) RelationshipRiskPutOrCallField {
	return RelationshipRiskPutOrCallField{quickfix.FIXInt(val)}
}

func (f RelationshipRiskPutOrCallField) Value() int { return f.Int() }

//RelationshipRiskRestructuringTypeField is a STRING field
type RelationshipRiskRestructuringTypeField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskRestructuringType (1604)
func (f RelationshipRiskRestructuringTypeField) Tag() quickfix.Tag {
	return tag.RelationshipRiskRestructuringType
}

//NewRelationshipRiskRestructuringType returns a new RelationshipRiskRestructuringTypeField initialized with val
func NewRelationshipRiskRestructuringType(val string) RelationshipRiskRestructuringTypeField {
	return RelationshipRiskRestructuringTypeField{quickfix.FIXString(val)}
}

func (f RelationshipRiskRestructuringTypeField) Value() string { return f.String() }

//RelationshipRiskSecurityAltIDField is a STRING field
type RelationshipRiskSecurityAltIDField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecurityAltID (1594)
func (f RelationshipRiskSecurityAltIDField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecurityAltID
}

//NewRelationshipRiskSecurityAltID returns a new RelationshipRiskSecurityAltIDField initialized with val
func NewRelationshipRiskSecurityAltID(val string) RelationshipRiskSecurityAltIDField {
	return RelationshipRiskSecurityAltIDField{quickfix.FIXString(val)}
}

func (f RelationshipRiskSecurityAltIDField) Value() string { return f.String() }

//RelationshipRiskSecurityAltIDSourceField is a STRING field
type RelationshipRiskSecurityAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecurityAltIDSource (1595)
func (f RelationshipRiskSecurityAltIDSourceField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecurityAltIDSource
}

//NewRelationshipRiskSecurityAltIDSource returns a new RelationshipRiskSecurityAltIDSourceField initialized with val
func NewRelationshipRiskSecurityAltIDSource(val string) RelationshipRiskSecurityAltIDSourceField {
	return RelationshipRiskSecurityAltIDSourceField{quickfix.FIXString(val)}
}

func (f RelationshipRiskSecurityAltIDSourceField) Value() string { return f.String() }

//RelationshipRiskSecurityDescField is a STRING field
type RelationshipRiskSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecurityDesc (1610)
func (f RelationshipRiskSecurityDescField) Tag() quickfix.Tag { return tag.RelationshipRiskSecurityDesc }

//NewRelationshipRiskSecurityDesc returns a new RelationshipRiskSecurityDescField initialized with val
func NewRelationshipRiskSecurityDesc(val string) RelationshipRiskSecurityDescField {
	return RelationshipRiskSecurityDescField{quickfix.FIXString(val)}
}

func (f RelationshipRiskSecurityDescField) Value() string { return f.String() }

//RelationshipRiskSecurityExchangeField is a EXCHANGE field
type RelationshipRiskSecurityExchangeField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecurityExchange (1609)
func (f RelationshipRiskSecurityExchangeField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecurityExchange
}

//NewRelationshipRiskSecurityExchange returns a new RelationshipRiskSecurityExchangeField initialized with val
func NewRelationshipRiskSecurityExchange(val string) RelationshipRiskSecurityExchangeField {
	return RelationshipRiskSecurityExchangeField{quickfix.FIXString(val)}
}

func (f RelationshipRiskSecurityExchangeField) Value() string { return f.String() }

//RelationshipRiskSecurityGroupField is a STRING field
type RelationshipRiskSecurityGroupField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecurityGroup (1598)
func (f RelationshipRiskSecurityGroupField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecurityGroup
}

//NewRelationshipRiskSecurityGroup returns a new RelationshipRiskSecurityGroupField initialized with val
func NewRelationshipRiskSecurityGroup(val string) RelationshipRiskSecurityGroupField {
	return RelationshipRiskSecurityGroupField{quickfix.FIXString(val)}
}

func (f RelationshipRiskSecurityGroupField) Value() string { return f.String() }

//RelationshipRiskSecurityIDField is a STRING field
type RelationshipRiskSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecurityID (1591)
func (f RelationshipRiskSecurityIDField) Tag() quickfix.Tag { return tag.RelationshipRiskSecurityID }

//NewRelationshipRiskSecurityID returns a new RelationshipRiskSecurityIDField initialized with val
func NewRelationshipRiskSecurityID(val string) RelationshipRiskSecurityIDField {
	return RelationshipRiskSecurityIDField{quickfix.FIXString(val)}
}

func (f RelationshipRiskSecurityIDField) Value() string { return f.String() }

//RelationshipRiskSecurityIDSourceField is a STRING field
type RelationshipRiskSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecurityIDSource (1592)
func (f RelationshipRiskSecurityIDSourceField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecurityIDSource
}

//NewRelationshipRiskSecurityIDSource returns a new RelationshipRiskSecurityIDSourceField initialized with val
func NewRelationshipRiskSecurityIDSource(val string) RelationshipRiskSecurityIDSourceField {
	return RelationshipRiskSecurityIDSourceField{quickfix.FIXString(val)}
}

func (f RelationshipRiskSecurityIDSourceField) Value() string { return f.String() }

//RelationshipRiskSecuritySubTypeField is a STRING field
type RelationshipRiskSecuritySubTypeField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecuritySubType (1601)
func (f RelationshipRiskSecuritySubTypeField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecuritySubType
}

//NewRelationshipRiskSecuritySubType returns a new RelationshipRiskSecuritySubTypeField initialized with val
func NewRelationshipRiskSecuritySubType(val string) RelationshipRiskSecuritySubTypeField {
	return RelationshipRiskSecuritySubTypeField{quickfix.FIXString(val)}
}

func (f RelationshipRiskSecuritySubTypeField) Value() string { return f.String() }

//RelationshipRiskSecurityTypeField is a STRING field
type RelationshipRiskSecurityTypeField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecurityType (1600)
func (f RelationshipRiskSecurityTypeField) Tag() quickfix.Tag { return tag.RelationshipRiskSecurityType }

//NewRelationshipRiskSecurityType returns a new RelationshipRiskSecurityTypeField initialized with val
func NewRelationshipRiskSecurityType(val string) RelationshipRiskSecurityTypeField {
	return RelationshipRiskSecurityTypeField{quickfix.FIXString(val)}
}

func (f RelationshipRiskSecurityTypeField) Value() string { return f.String() }

//RelationshipRiskSeniorityField is a STRING field
type RelationshipRiskSeniorityField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSeniority (1605)
func (f RelationshipRiskSeniorityField) Tag() quickfix.Tag { return tag.RelationshipRiskSeniority }

//NewRelationshipRiskSeniority returns a new RelationshipRiskSeniorityField initialized with val
func NewRelationshipRiskSeniority(val string) RelationshipRiskSeniorityField {
	return RelationshipRiskSeniorityField{quickfix.FIXString(val)}
}

func (f RelationshipRiskSeniorityField) Value() string { return f.String() }

//RelationshipRiskSymbolField is a STRING field
type RelationshipRiskSymbolField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSymbol (1589)
func (f RelationshipRiskSymbolField) Tag() quickfix.Tag { return tag.RelationshipRiskSymbol }

//NewRelationshipRiskSymbol returns a new RelationshipRiskSymbolField initialized with val
func NewRelationshipRiskSymbol(val string) RelationshipRiskSymbolField {
	return RelationshipRiskSymbolField{quickfix.FIXString(val)}
}

func (f RelationshipRiskSymbolField) Value() string { return f.String() }

//RelationshipRiskSymbolSfxField is a STRING field
type RelationshipRiskSymbolSfxField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSymbolSfx (1590)
func (f RelationshipRiskSymbolSfxField) Tag() quickfix.Tag { return tag.RelationshipRiskSymbolSfx }

//NewRelationshipRiskSymbolSfx returns a new RelationshipRiskSymbolSfxField initialized with val
func NewRelationshipRiskSymbolSfx(val string) RelationshipRiskSymbolSfxField {
	return RelationshipRiskSymbolSfxField{quickfix.FIXString(val)}
}

func (f RelationshipRiskSymbolSfxField) Value() string { return f.String() }

//RelationshipRiskWarningLevelNameField is a STRING field
type RelationshipRiskWarningLevelNameField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskWarningLevelName (1615)
func (f RelationshipRiskWarningLevelNameField) Tag() quickfix.Tag {
	return tag.RelationshipRiskWarningLevelName
}

//NewRelationshipRiskWarningLevelName returns a new RelationshipRiskWarningLevelNameField initialized with val
func NewRelationshipRiskWarningLevelName(val string) RelationshipRiskWarningLevelNameField {
	return RelationshipRiskWarningLevelNameField{quickfix.FIXString(val)}
}

func (f RelationshipRiskWarningLevelNameField) Value() string { return f.String() }

//RelationshipRiskWarningLevelPercentField is a PERCENTAGE field
type RelationshipRiskWarningLevelPercentField struct{ quickfix.FIXDecimal }

//Tag returns tag.RelationshipRiskWarningLevelPercent (1614)
func (f RelationshipRiskWarningLevelPercentField) Tag() quickfix.Tag {
	return tag.RelationshipRiskWarningLevelPercent
}

//NewRelationshipRiskWarningLevelPercent returns a new RelationshipRiskWarningLevelPercentField initialized with val and scale
func NewRelationshipRiskWarningLevelPercent(val decimal.Decimal, scale int32) RelationshipRiskWarningLevelPercentField {
	return RelationshipRiskWarningLevelPercentField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f RelationshipRiskWarningLevelPercentField) Value() (val decimal.Decimal) { return f.Decimal }

//RepoCollateralSecurityTypeField is a INT field
type RepoCollateralSecurityTypeField struct{ quickfix.FIXInt }

//Tag returns tag.RepoCollateralSecurityType (239)
func (f RepoCollateralSecurityTypeField) Tag() quickfix.Tag { return tag.RepoCollateralSecurityType }

//NewRepoCollateralSecurityType returns a new RepoCollateralSecurityTypeField initialized with val
func NewRepoCollateralSecurityType(val int) RepoCollateralSecurityTypeField {
	return RepoCollateralSecurityTypeField{quickfix.FIXInt(val)}
}

func (f RepoCollateralSecurityTypeField) Value() int { return f.Int() }

//ReportToExchField is a BOOLEAN field
type ReportToExchField struct{ quickfix.FIXBoolean }

//Tag returns tag.ReportToExch (113)
func (f ReportToExchField) Tag() quickfix.Tag { return tag.ReportToExch }

//NewReportToExch returns a new ReportToExchField initialized with val
func NewReportToExch(val bool) ReportToExchField {
	return ReportToExchField{quickfix.FIXBoolean(val)}
}

func (f ReportToExchField) Value() bool { return f.Bool() }

//ReportedPxField is a PRICE field
type ReportedPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.ReportedPx (861)
func (f ReportedPxField) Tag() quickfix.Tag { return tag.ReportedPx }

//NewReportedPx returns a new ReportedPxField initialized with val and scale
func NewReportedPx(val decimal.Decimal, scale int32) ReportedPxField {
	return ReportedPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f ReportedPxField) Value() (val decimal.Decimal) { return f.Decimal }

//ReportedPxDiffField is a BOOLEAN field
type ReportedPxDiffField struct{ quickfix.FIXBoolean }

//Tag returns tag.ReportedPxDiff (1134)
func (f ReportedPxDiffField) Tag() quickfix.Tag { return tag.ReportedPxDiff }

//NewReportedPxDiff returns a new ReportedPxDiffField initialized with val
func NewReportedPxDiff(val bool) ReportedPxDiffField {
	return ReportedPxDiffField{quickfix.FIXBoolean(val)}
}

func (f ReportedPxDiffField) Value() bool { return f.Bool() }

//RepurchaseRateField is a PERCENTAGE field
type RepurchaseRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.RepurchaseRate (227)
func (f RepurchaseRateField) Tag() quickfix.Tag { return tag.RepurchaseRate }

//NewRepurchaseRate returns a new RepurchaseRateField initialized with val and scale
func NewRepurchaseRate(val decimal.Decimal, scale int32) RepurchaseRateField {
	return RepurchaseRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f RepurchaseRateField) Value() (val decimal.Decimal) { return f.Decimal }

//RepurchaseTermField is a INT field
type RepurchaseTermField struct{ quickfix.FIXInt }

//Tag returns tag.RepurchaseTerm (226)
func (f RepurchaseTermField) Tag() quickfix.Tag { return tag.RepurchaseTerm }

//NewRepurchaseTerm returns a new RepurchaseTermField initialized with val
func NewRepurchaseTerm(val int) RepurchaseTermField {
	return RepurchaseTermField{quickfix.FIXInt(val)}
}

func (f RepurchaseTermField) Value() int { return f.Int() }

//RequestedPartyRoleField is a INT field
type RequestedPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.RequestedPartyRole (1509)
func (f RequestedPartyRoleField) Tag() quickfix.Tag { return tag.RequestedPartyRole }

//NewRequestedPartyRole returns a new RequestedPartyRoleField initialized with val
func NewRequestedPartyRole(val int) RequestedPartyRoleField {
	return RequestedPartyRoleField{quickfix.FIXInt(val)}
}

func (f RequestedPartyRoleField) Value() int { return f.Int() }

//ResetSeqNumFlagField is a BOOLEAN field
type ResetSeqNumFlagField struct{ quickfix.FIXBoolean }

//Tag returns tag.ResetSeqNumFlag (141)
func (f ResetSeqNumFlagField) Tag() quickfix.Tag { return tag.ResetSeqNumFlag }

//NewResetSeqNumFlag returns a new ResetSeqNumFlagField initialized with val
func NewResetSeqNumFlag(val bool) ResetSeqNumFlagField {
	return ResetSeqNumFlagField{quickfix.FIXBoolean(val)}
}

func (f ResetSeqNumFlagField) Value() bool { return f.Bool() }

//RespondentTypeField is a enum.RespondentType field
type RespondentTypeField struct{ quickfix.FIXString }

//Tag returns tag.RespondentType (1172)
func (f RespondentTypeField) Tag() quickfix.Tag { return tag.RespondentType }

func NewRespondentType(val enum.RespondentType) RespondentTypeField {
	return RespondentTypeField{quickfix.FIXString(val)}
}

func (f RespondentTypeField) Value() enum.RespondentType { return enum.RespondentType(f.String()) }

//ResponseDestinationField is a STRING field
type ResponseDestinationField struct{ quickfix.FIXString }

//Tag returns tag.ResponseDestination (726)
func (f ResponseDestinationField) Tag() quickfix.Tag { return tag.ResponseDestination }

//NewResponseDestination returns a new ResponseDestinationField initialized with val
func NewResponseDestination(val string) ResponseDestinationField {
	return ResponseDestinationField{quickfix.FIXString(val)}
}

func (f ResponseDestinationField) Value() string { return f.String() }

//ResponseTransportTypeField is a enum.ResponseTransportType field
type ResponseTransportTypeField struct{ quickfix.FIXString }

//Tag returns tag.ResponseTransportType (725)
func (f ResponseTransportTypeField) Tag() quickfix.Tag { return tag.ResponseTransportType }

func NewResponseTransportType(val enum.ResponseTransportType) ResponseTransportTypeField {
	return ResponseTransportTypeField{quickfix.FIXString(val)}
}

func (f ResponseTransportTypeField) Value() enum.ResponseTransportType {
	return enum.ResponseTransportType(f.String())
}

//RestructuringTypeField is a enum.RestructuringType field
type RestructuringTypeField struct{ quickfix.FIXString }

//Tag returns tag.RestructuringType (1449)
func (f RestructuringTypeField) Tag() quickfix.Tag { return tag.RestructuringType }

func NewRestructuringType(val enum.RestructuringType) RestructuringTypeField {
	return RestructuringTypeField{quickfix.FIXString(val)}
}

func (f RestructuringTypeField) Value() enum.RestructuringType {
	return enum.RestructuringType(f.String())
}

//ReversalIndicatorField is a BOOLEAN field
type ReversalIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.ReversalIndicator (700)
func (f ReversalIndicatorField) Tag() quickfix.Tag { return tag.ReversalIndicator }

//NewReversalIndicator returns a new ReversalIndicatorField initialized with val
func NewReversalIndicator(val bool) ReversalIndicatorField {
	return ReversalIndicatorField{quickfix.FIXBoolean(val)}
}

func (f ReversalIndicatorField) Value() bool { return f.Bool() }

//RiskCFICodeField is a STRING field
type RiskCFICodeField struct{ quickfix.FIXString }

//Tag returns tag.RiskCFICode (1546)
func (f RiskCFICodeField) Tag() quickfix.Tag { return tag.RiskCFICode }

//NewRiskCFICode returns a new RiskCFICodeField initialized with val
func NewRiskCFICode(val string) RiskCFICodeField {
	return RiskCFICodeField{quickfix.FIXString(val)}
}

func (f RiskCFICodeField) Value() string { return f.String() }

//RiskCouponRateField is a PERCENTAGE field
type RiskCouponRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.RiskCouponRate (1555)
func (f RiskCouponRateField) Tag() quickfix.Tag { return tag.RiskCouponRate }

//NewRiskCouponRate returns a new RiskCouponRateField initialized with val and scale
func NewRiskCouponRate(val decimal.Decimal, scale int32) RiskCouponRateField {
	return RiskCouponRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f RiskCouponRateField) Value() (val decimal.Decimal) { return f.Decimal }

//RiskEncodedSecurityDescField is a DATA field
type RiskEncodedSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.RiskEncodedSecurityDesc (1621)
func (f RiskEncodedSecurityDescField) Tag() quickfix.Tag { return tag.RiskEncodedSecurityDesc }

//NewRiskEncodedSecurityDesc returns a new RiskEncodedSecurityDescField initialized with val
func NewRiskEncodedSecurityDesc(val string) RiskEncodedSecurityDescField {
	return RiskEncodedSecurityDescField{quickfix.FIXString(val)}
}

func (f RiskEncodedSecurityDescField) Value() string { return f.String() }

//RiskEncodedSecurityDescLenField is a LENGTH field
type RiskEncodedSecurityDescLenField struct{ quickfix.FIXInt }

//Tag returns tag.RiskEncodedSecurityDescLen (1620)
func (f RiskEncodedSecurityDescLenField) Tag() quickfix.Tag { return tag.RiskEncodedSecurityDescLen }

//NewRiskEncodedSecurityDescLen returns a new RiskEncodedSecurityDescLenField initialized with val
func NewRiskEncodedSecurityDescLen(val int) RiskEncodedSecurityDescLenField {
	return RiskEncodedSecurityDescLenField{quickfix.FIXInt(val)}
}

func (f RiskEncodedSecurityDescLenField) Value() int { return f.Int() }

//RiskFlexibleIndicatorField is a BOOLEAN field
type RiskFlexibleIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.RiskFlexibleIndicator (1554)
func (f RiskFlexibleIndicatorField) Tag() quickfix.Tag { return tag.RiskFlexibleIndicator }

//NewRiskFlexibleIndicator returns a new RiskFlexibleIndicatorField initialized with val
func NewRiskFlexibleIndicator(val bool) RiskFlexibleIndicatorField {
	return RiskFlexibleIndicatorField{quickfix.FIXBoolean(val)}
}

func (f RiskFlexibleIndicatorField) Value() bool { return f.Bool() }

//RiskFreeRateField is a FLOAT field
type RiskFreeRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.RiskFreeRate (1190)
func (f RiskFreeRateField) Tag() quickfix.Tag { return tag.RiskFreeRate }

//NewRiskFreeRate returns a new RiskFreeRateField initialized with val and scale
func NewRiskFreeRate(val decimal.Decimal, scale int32) RiskFreeRateField {
	return RiskFreeRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f RiskFreeRateField) Value() (val decimal.Decimal) { return f.Decimal }

//RiskInstrumentMultiplierField is a FLOAT field
type RiskInstrumentMultiplierField struct{ quickfix.FIXDecimal }

//Tag returns tag.RiskInstrumentMultiplier (1558)
func (f RiskInstrumentMultiplierField) Tag() quickfix.Tag { return tag.RiskInstrumentMultiplier }

//NewRiskInstrumentMultiplier returns a new RiskInstrumentMultiplierField initialized with val and scale
func NewRiskInstrumentMultiplier(val decimal.Decimal, scale int32) RiskInstrumentMultiplierField {
	return RiskInstrumentMultiplierField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f RiskInstrumentMultiplierField) Value() (val decimal.Decimal) { return f.Decimal }

//RiskInstrumentOperatorField is a enum.RiskInstrumentOperator field
type RiskInstrumentOperatorField struct{ quickfix.FIXString }

//Tag returns tag.RiskInstrumentOperator (1535)
func (f RiskInstrumentOperatorField) Tag() quickfix.Tag { return tag.RiskInstrumentOperator }

func NewRiskInstrumentOperator(val enum.RiskInstrumentOperator) RiskInstrumentOperatorField {
	return RiskInstrumentOperatorField{quickfix.FIXString(val)}
}

func (f RiskInstrumentOperatorField) Value() enum.RiskInstrumentOperator {
	return enum.RiskInstrumentOperator(f.String())
}

//RiskInstrumentSettlTypeField is a STRING field
type RiskInstrumentSettlTypeField struct{ quickfix.FIXString }

//Tag returns tag.RiskInstrumentSettlType (1557)
func (f RiskInstrumentSettlTypeField) Tag() quickfix.Tag { return tag.RiskInstrumentSettlType }

//NewRiskInstrumentSettlType returns a new RiskInstrumentSettlTypeField initialized with val
func NewRiskInstrumentSettlType(val string) RiskInstrumentSettlTypeField {
	return RiskInstrumentSettlTypeField{quickfix.FIXString(val)}
}

func (f RiskInstrumentSettlTypeField) Value() string { return f.String() }

//RiskLimitAmountField is a AMT field
type RiskLimitAmountField struct{ quickfix.FIXDecimal }

//Tag returns tag.RiskLimitAmount (1531)
func (f RiskLimitAmountField) Tag() quickfix.Tag { return tag.RiskLimitAmount }

//NewRiskLimitAmount returns a new RiskLimitAmountField initialized with val and scale
func NewRiskLimitAmount(val decimal.Decimal, scale int32) RiskLimitAmountField {
	return RiskLimitAmountField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f RiskLimitAmountField) Value() (val decimal.Decimal) { return f.Decimal }

//RiskLimitCurrencyField is a CURRENCY field
type RiskLimitCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.RiskLimitCurrency (1532)
func (f RiskLimitCurrencyField) Tag() quickfix.Tag { return tag.RiskLimitCurrency }

//NewRiskLimitCurrency returns a new RiskLimitCurrencyField initialized with val
func NewRiskLimitCurrency(val string) RiskLimitCurrencyField {
	return RiskLimitCurrencyField{quickfix.FIXString(val)}
}

func (f RiskLimitCurrencyField) Value() string { return f.String() }

//RiskLimitPlatformField is a STRING field
type RiskLimitPlatformField struct{ quickfix.FIXString }

//Tag returns tag.RiskLimitPlatform (1533)
func (f RiskLimitPlatformField) Tag() quickfix.Tag { return tag.RiskLimitPlatform }

//NewRiskLimitPlatform returns a new RiskLimitPlatformField initialized with val
func NewRiskLimitPlatform(val string) RiskLimitPlatformField {
	return RiskLimitPlatformField{quickfix.FIXString(val)}
}

func (f RiskLimitPlatformField) Value() string { return f.String() }

//RiskLimitTypeField is a enum.RiskLimitType field
type RiskLimitTypeField struct{ quickfix.FIXString }

//Tag returns tag.RiskLimitType (1530)
func (f RiskLimitTypeField) Tag() quickfix.Tag { return tag.RiskLimitType }

func NewRiskLimitType(val enum.RiskLimitType) RiskLimitTypeField {
	return RiskLimitTypeField{quickfix.FIXString(val)}
}

func (f RiskLimitTypeField) Value() enum.RiskLimitType { return enum.RiskLimitType(f.String()) }

//RiskMaturityMonthYearField is a MONTHYEAR field
type RiskMaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.RiskMaturityMonthYear (1549)
func (f RiskMaturityMonthYearField) Tag() quickfix.Tag { return tag.RiskMaturityMonthYear }

//NewRiskMaturityMonthYear returns a new RiskMaturityMonthYearField initialized with val
func NewRiskMaturityMonthYear(val string) RiskMaturityMonthYearField {
	return RiskMaturityMonthYearField{quickfix.FIXString(val)}
}

func (f RiskMaturityMonthYearField) Value() string { return f.String() }

//RiskMaturityTimeField is a TZTIMEONLY field
type RiskMaturityTimeField struct{ quickfix.FIXString }

//Tag returns tag.RiskMaturityTime (1550)
func (f RiskMaturityTimeField) Tag() quickfix.Tag { return tag.RiskMaturityTime }

//NewRiskMaturityTime returns a new RiskMaturityTimeField initialized with val
func NewRiskMaturityTime(val string) RiskMaturityTimeField {
	return RiskMaturityTimeField{quickfix.FIXString(val)}
}

func (f RiskMaturityTimeField) Value() string { return f.String() }

//RiskProductField is a INT field
type RiskProductField struct{ quickfix.FIXInt }

//Tag returns tag.RiskProduct (1543)
func (f RiskProductField) Tag() quickfix.Tag { return tag.RiskProduct }

//NewRiskProduct returns a new RiskProductField initialized with val
func NewRiskProduct(val int) RiskProductField {
	return RiskProductField{quickfix.FIXInt(val)}
}

func (f RiskProductField) Value() int { return f.Int() }

//RiskProductComplexField is a STRING field
type RiskProductComplexField struct{ quickfix.FIXString }

//Tag returns tag.RiskProductComplex (1544)
func (f RiskProductComplexField) Tag() quickfix.Tag { return tag.RiskProductComplex }

//NewRiskProductComplex returns a new RiskProductComplexField initialized with val
func NewRiskProductComplex(val string) RiskProductComplexField {
	return RiskProductComplexField{quickfix.FIXString(val)}
}

func (f RiskProductComplexField) Value() string { return f.String() }

//RiskPutOrCallField is a INT field
type RiskPutOrCallField struct{ quickfix.FIXInt }

//Tag returns tag.RiskPutOrCall (1553)
func (f RiskPutOrCallField) Tag() quickfix.Tag { return tag.RiskPutOrCall }

//NewRiskPutOrCall returns a new RiskPutOrCallField initialized with val
func NewRiskPutOrCall(val int) RiskPutOrCallField {
	return RiskPutOrCallField{quickfix.FIXInt(val)}
}

func (f RiskPutOrCallField) Value() int { return f.Int() }

//RiskRestructuringTypeField is a STRING field
type RiskRestructuringTypeField struct{ quickfix.FIXString }

//Tag returns tag.RiskRestructuringType (1551)
func (f RiskRestructuringTypeField) Tag() quickfix.Tag { return tag.RiskRestructuringType }

//NewRiskRestructuringType returns a new RiskRestructuringTypeField initialized with val
func NewRiskRestructuringType(val string) RiskRestructuringTypeField {
	return RiskRestructuringTypeField{quickfix.FIXString(val)}
}

func (f RiskRestructuringTypeField) Value() string { return f.String() }

//RiskSecurityAltIDField is a STRING field
type RiskSecurityAltIDField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecurityAltID (1541)
func (f RiskSecurityAltIDField) Tag() quickfix.Tag { return tag.RiskSecurityAltID }

//NewRiskSecurityAltID returns a new RiskSecurityAltIDField initialized with val
func NewRiskSecurityAltID(val string) RiskSecurityAltIDField {
	return RiskSecurityAltIDField{quickfix.FIXString(val)}
}

func (f RiskSecurityAltIDField) Value() string { return f.String() }

//RiskSecurityAltIDSourceField is a STRING field
type RiskSecurityAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecurityAltIDSource (1542)
func (f RiskSecurityAltIDSourceField) Tag() quickfix.Tag { return tag.RiskSecurityAltIDSource }

//NewRiskSecurityAltIDSource returns a new RiskSecurityAltIDSourceField initialized with val
func NewRiskSecurityAltIDSource(val string) RiskSecurityAltIDSourceField {
	return RiskSecurityAltIDSourceField{quickfix.FIXString(val)}
}

func (f RiskSecurityAltIDSourceField) Value() string { return f.String() }

//RiskSecurityDescField is a STRING field
type RiskSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecurityDesc (1556)
func (f RiskSecurityDescField) Tag() quickfix.Tag { return tag.RiskSecurityDesc }

//NewRiskSecurityDesc returns a new RiskSecurityDescField initialized with val
func NewRiskSecurityDesc(val string) RiskSecurityDescField {
	return RiskSecurityDescField{quickfix.FIXString(val)}
}

func (f RiskSecurityDescField) Value() string { return f.String() }

//RiskSecurityExchangeField is a EXCHANGE field
type RiskSecurityExchangeField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecurityExchange (1616)
func (f RiskSecurityExchangeField) Tag() quickfix.Tag { return tag.RiskSecurityExchange }

//NewRiskSecurityExchange returns a new RiskSecurityExchangeField initialized with val
func NewRiskSecurityExchange(val string) RiskSecurityExchangeField {
	return RiskSecurityExchangeField{quickfix.FIXString(val)}
}

func (f RiskSecurityExchangeField) Value() string { return f.String() }

//RiskSecurityGroupField is a STRING field
type RiskSecurityGroupField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecurityGroup (1545)
func (f RiskSecurityGroupField) Tag() quickfix.Tag { return tag.RiskSecurityGroup }

//NewRiskSecurityGroup returns a new RiskSecurityGroupField initialized with val
func NewRiskSecurityGroup(val string) RiskSecurityGroupField {
	return RiskSecurityGroupField{quickfix.FIXString(val)}
}

func (f RiskSecurityGroupField) Value() string { return f.String() }

//RiskSecurityIDField is a STRING field
type RiskSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecurityID (1538)
func (f RiskSecurityIDField) Tag() quickfix.Tag { return tag.RiskSecurityID }

//NewRiskSecurityID returns a new RiskSecurityIDField initialized with val
func NewRiskSecurityID(val string) RiskSecurityIDField {
	return RiskSecurityIDField{quickfix.FIXString(val)}
}

func (f RiskSecurityIDField) Value() string { return f.String() }

//RiskSecurityIDSourceField is a STRING field
type RiskSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecurityIDSource (1539)
func (f RiskSecurityIDSourceField) Tag() quickfix.Tag { return tag.RiskSecurityIDSource }

//NewRiskSecurityIDSource returns a new RiskSecurityIDSourceField initialized with val
func NewRiskSecurityIDSource(val string) RiskSecurityIDSourceField {
	return RiskSecurityIDSourceField{quickfix.FIXString(val)}
}

func (f RiskSecurityIDSourceField) Value() string { return f.String() }

//RiskSecuritySubTypeField is a STRING field
type RiskSecuritySubTypeField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecuritySubType (1548)
func (f RiskSecuritySubTypeField) Tag() quickfix.Tag { return tag.RiskSecuritySubType }

//NewRiskSecuritySubType returns a new RiskSecuritySubTypeField initialized with val
func NewRiskSecuritySubType(val string) RiskSecuritySubTypeField {
	return RiskSecuritySubTypeField{quickfix.FIXString(val)}
}

func (f RiskSecuritySubTypeField) Value() string { return f.String() }

//RiskSecurityTypeField is a STRING field
type RiskSecurityTypeField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecurityType (1547)
func (f RiskSecurityTypeField) Tag() quickfix.Tag { return tag.RiskSecurityType }

//NewRiskSecurityType returns a new RiskSecurityTypeField initialized with val
func NewRiskSecurityType(val string) RiskSecurityTypeField {
	return RiskSecurityTypeField{quickfix.FIXString(val)}
}

func (f RiskSecurityTypeField) Value() string { return f.String() }

//RiskSeniorityField is a STRING field
type RiskSeniorityField struct{ quickfix.FIXString }

//Tag returns tag.RiskSeniority (1552)
func (f RiskSeniorityField) Tag() quickfix.Tag { return tag.RiskSeniority }

//NewRiskSeniority returns a new RiskSeniorityField initialized with val
func NewRiskSeniority(val string) RiskSeniorityField {
	return RiskSeniorityField{quickfix.FIXString(val)}
}

func (f RiskSeniorityField) Value() string { return f.String() }

//RiskSymbolField is a STRING field
type RiskSymbolField struct{ quickfix.FIXString }

//Tag returns tag.RiskSymbol (1536)
func (f RiskSymbolField) Tag() quickfix.Tag { return tag.RiskSymbol }

//NewRiskSymbol returns a new RiskSymbolField initialized with val
func NewRiskSymbol(val string) RiskSymbolField {
	return RiskSymbolField{quickfix.FIXString(val)}
}

func (f RiskSymbolField) Value() string { return f.String() }

//RiskSymbolSfxField is a STRING field
type RiskSymbolSfxField struct{ quickfix.FIXString }

//Tag returns tag.RiskSymbolSfx (1537)
func (f RiskSymbolSfxField) Tag() quickfix.Tag { return tag.RiskSymbolSfx }

//NewRiskSymbolSfx returns a new RiskSymbolSfxField initialized with val
func NewRiskSymbolSfx(val string) RiskSymbolSfxField {
	return RiskSymbolSfxField{quickfix.FIXString(val)}
}

func (f RiskSymbolSfxField) Value() string { return f.String() }

//RiskWarningLevelNameField is a STRING field
type RiskWarningLevelNameField struct{ quickfix.FIXString }

//Tag returns tag.RiskWarningLevelName (1561)
func (f RiskWarningLevelNameField) Tag() quickfix.Tag { return tag.RiskWarningLevelName }

//NewRiskWarningLevelName returns a new RiskWarningLevelNameField initialized with val
func NewRiskWarningLevelName(val string) RiskWarningLevelNameField {
	return RiskWarningLevelNameField{quickfix.FIXString(val)}
}

func (f RiskWarningLevelNameField) Value() string { return f.String() }

//RiskWarningLevelPercentField is a PERCENTAGE field
type RiskWarningLevelPercentField struct{ quickfix.FIXDecimal }

//Tag returns tag.RiskWarningLevelPercent (1560)
func (f RiskWarningLevelPercentField) Tag() quickfix.Tag { return tag.RiskWarningLevelPercent }

//NewRiskWarningLevelPercent returns a new RiskWarningLevelPercentField initialized with val and scale
func NewRiskWarningLevelPercent(val decimal.Decimal, scale int32) RiskWarningLevelPercentField {
	return RiskWarningLevelPercentField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f RiskWarningLevelPercentField) Value() (val decimal.Decimal) { return f.Decimal }

//RndPxField is a PRICE field
type RndPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.RndPx (991)
func (f RndPxField) Tag() quickfix.Tag { return tag.RndPx }

//NewRndPx returns a new RndPxField initialized with val and scale
func NewRndPx(val decimal.Decimal, scale int32) RndPxField {
	return RndPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f RndPxField) Value() (val decimal.Decimal) { return f.Decimal }

//RootPartyIDField is a STRING field
type RootPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.RootPartyID (1117)
func (f RootPartyIDField) Tag() quickfix.Tag { return tag.RootPartyID }

//NewRootPartyID returns a new RootPartyIDField initialized with val
func NewRootPartyID(val string) RootPartyIDField {
	return RootPartyIDField{quickfix.FIXString(val)}
}

func (f RootPartyIDField) Value() string { return f.String() }

//RootPartyIDSourceField is a CHAR field
type RootPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RootPartyIDSource (1118)
func (f RootPartyIDSourceField) Tag() quickfix.Tag { return tag.RootPartyIDSource }

//NewRootPartyIDSource returns a new RootPartyIDSourceField initialized with val
func NewRootPartyIDSource(val string) RootPartyIDSourceField {
	return RootPartyIDSourceField{quickfix.FIXString(val)}
}

func (f RootPartyIDSourceField) Value() string { return f.String() }

//RootPartyRoleField is a INT field
type RootPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.RootPartyRole (1119)
func (f RootPartyRoleField) Tag() quickfix.Tag { return tag.RootPartyRole }

//NewRootPartyRole returns a new RootPartyRoleField initialized with val
func NewRootPartyRole(val int) RootPartyRoleField {
	return RootPartyRoleField{quickfix.FIXInt(val)}
}

func (f RootPartyRoleField) Value() int { return f.Int() }

//RootPartySubIDField is a STRING field
type RootPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.RootPartySubID (1121)
func (f RootPartySubIDField) Tag() quickfix.Tag { return tag.RootPartySubID }

//NewRootPartySubID returns a new RootPartySubIDField initialized with val
func NewRootPartySubID(val string) RootPartySubIDField {
	return RootPartySubIDField{quickfix.FIXString(val)}
}

func (f RootPartySubIDField) Value() string { return f.String() }

//RootPartySubIDTypeField is a INT field
type RootPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.RootPartySubIDType (1122)
func (f RootPartySubIDTypeField) Tag() quickfix.Tag { return tag.RootPartySubIDType }

//NewRootPartySubIDType returns a new RootPartySubIDTypeField initialized with val
func NewRootPartySubIDType(val int) RootPartySubIDTypeField {
	return RootPartySubIDTypeField{quickfix.FIXInt(val)}
}

func (f RootPartySubIDTypeField) Value() int { return f.Int() }

//RoundLotField is a QTY field
type RoundLotField struct{ quickfix.FIXDecimal }

//Tag returns tag.RoundLot (561)
func (f RoundLotField) Tag() quickfix.Tag { return tag.RoundLot }

//NewRoundLot returns a new RoundLotField initialized with val and scale
func NewRoundLot(val decimal.Decimal, scale int32) RoundLotField {
	return RoundLotField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f RoundLotField) Value() (val decimal.Decimal) { return f.Decimal }

//RoundingDirectionField is a enum.RoundingDirection field
type RoundingDirectionField struct{ quickfix.FIXString }

//Tag returns tag.RoundingDirection (468)
func (f RoundingDirectionField) Tag() quickfix.Tag { return tag.RoundingDirection }

func NewRoundingDirection(val enum.RoundingDirection) RoundingDirectionField {
	return RoundingDirectionField{quickfix.FIXString(val)}
}

func (f RoundingDirectionField) Value() enum.RoundingDirection {
	return enum.RoundingDirection(f.String())
}

//RoundingModulusField is a FLOAT field
type RoundingModulusField struct{ quickfix.FIXDecimal }

//Tag returns tag.RoundingModulus (469)
func (f RoundingModulusField) Tag() quickfix.Tag { return tag.RoundingModulus }

//NewRoundingModulus returns a new RoundingModulusField initialized with val and scale
func NewRoundingModulus(val decimal.Decimal, scale int32) RoundingModulusField {
	return RoundingModulusField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f RoundingModulusField) Value() (val decimal.Decimal) { return f.Decimal }

//RoutingIDField is a STRING field
type RoutingIDField struct{ quickfix.FIXString }

//Tag returns tag.RoutingID (217)
func (f RoutingIDField) Tag() quickfix.Tag { return tag.RoutingID }

//NewRoutingID returns a new RoutingIDField initialized with val
func NewRoutingID(val string) RoutingIDField {
	return RoutingIDField{quickfix.FIXString(val)}
}

func (f RoutingIDField) Value() string { return f.String() }

//RoutingTypeField is a enum.RoutingType field
type RoutingTypeField struct{ quickfix.FIXString }

//Tag returns tag.RoutingType (216)
func (f RoutingTypeField) Tag() quickfix.Tag { return tag.RoutingType }

func NewRoutingType(val enum.RoutingType) RoutingTypeField {
	return RoutingTypeField{quickfix.FIXString(val)}
}

func (f RoutingTypeField) Value() enum.RoutingType { return enum.RoutingType(f.String()) }

//RptSeqField is a INT field
type RptSeqField struct{ quickfix.FIXInt }

//Tag returns tag.RptSeq (83)
func (f RptSeqField) Tag() quickfix.Tag { return tag.RptSeq }

//NewRptSeq returns a new RptSeqField initialized with val
func NewRptSeq(val int) RptSeqField {
	return RptSeqField{quickfix.FIXInt(val)}
}

func (f RptSeqField) Value() int { return f.Int() }

//RptSysField is a STRING field
type RptSysField struct{ quickfix.FIXString }

//Tag returns tag.RptSys (1135)
func (f RptSysField) Tag() quickfix.Tag { return tag.RptSys }

//NewRptSys returns a new RptSysField initialized with val
func NewRptSys(val string) RptSysField {
	return RptSysField{quickfix.FIXString(val)}
}

func (f RptSysField) Value() string { return f.String() }

//Rule80AField is a enum.Rule80A field
type Rule80AField struct{ quickfix.FIXString }

//Tag returns tag.Rule80A (47)
func (f Rule80AField) Tag() quickfix.Tag { return tag.Rule80A }

func NewRule80A(val enum.Rule80A) Rule80AField {
	return Rule80AField{quickfix.FIXString(val)}
}

func (f Rule80AField) Value() enum.Rule80A { return enum.Rule80A(f.String()) }

//ScopeField is a enum.Scope field
type ScopeField struct{ quickfix.FIXString }

//Tag returns tag.Scope (546)
func (f ScopeField) Tag() quickfix.Tag { return tag.Scope }

func NewScope(val enum.Scope) ScopeField {
	return ScopeField{quickfix.FIXString(val)}
}

func (f ScopeField) Value() enum.Scope { return enum.Scope(f.String()) }

//SecDefStatusField is a enum.SecDefStatus field
type SecDefStatusField struct{ quickfix.FIXString }

//Tag returns tag.SecDefStatus (653)
func (f SecDefStatusField) Tag() quickfix.Tag { return tag.SecDefStatus }

func NewSecDefStatus(val enum.SecDefStatus) SecDefStatusField {
	return SecDefStatusField{quickfix.FIXString(val)}
}

func (f SecDefStatusField) Value() enum.SecDefStatus { return enum.SecDefStatus(f.String()) }

//SecondaryAllocIDField is a STRING field
type SecondaryAllocIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryAllocID (793)
func (f SecondaryAllocIDField) Tag() quickfix.Tag { return tag.SecondaryAllocID }

//NewSecondaryAllocID returns a new SecondaryAllocIDField initialized with val
func NewSecondaryAllocID(val string) SecondaryAllocIDField {
	return SecondaryAllocIDField{quickfix.FIXString(val)}
}

func (f SecondaryAllocIDField) Value() string { return f.String() }

//SecondaryClOrdIDField is a STRING field
type SecondaryClOrdIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryClOrdID (526)
func (f SecondaryClOrdIDField) Tag() quickfix.Tag { return tag.SecondaryClOrdID }

//NewSecondaryClOrdID returns a new SecondaryClOrdIDField initialized with val
func NewSecondaryClOrdID(val string) SecondaryClOrdIDField {
	return SecondaryClOrdIDField{quickfix.FIXString(val)}
}

func (f SecondaryClOrdIDField) Value() string { return f.String() }

//SecondaryDisplayQtyField is a QTY field
type SecondaryDisplayQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.SecondaryDisplayQty (1082)
func (f SecondaryDisplayQtyField) Tag() quickfix.Tag { return tag.SecondaryDisplayQty }

//NewSecondaryDisplayQty returns a new SecondaryDisplayQtyField initialized with val and scale
func NewSecondaryDisplayQty(val decimal.Decimal, scale int32) SecondaryDisplayQtyField {
	return SecondaryDisplayQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SecondaryDisplayQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//SecondaryExecIDField is a STRING field
type SecondaryExecIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryExecID (527)
func (f SecondaryExecIDField) Tag() quickfix.Tag { return tag.SecondaryExecID }

//NewSecondaryExecID returns a new SecondaryExecIDField initialized with val
func NewSecondaryExecID(val string) SecondaryExecIDField {
	return SecondaryExecIDField{quickfix.FIXString(val)}
}

func (f SecondaryExecIDField) Value() string { return f.String() }

//SecondaryFirmTradeIDField is a STRING field
type SecondaryFirmTradeIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryFirmTradeID (1042)
func (f SecondaryFirmTradeIDField) Tag() quickfix.Tag { return tag.SecondaryFirmTradeID }

//NewSecondaryFirmTradeID returns a new SecondaryFirmTradeIDField initialized with val
func NewSecondaryFirmTradeID(val string) SecondaryFirmTradeIDField {
	return SecondaryFirmTradeIDField{quickfix.FIXString(val)}
}

func (f SecondaryFirmTradeIDField) Value() string { return f.String() }

//SecondaryHighLimitPriceField is a PRICE field
type SecondaryHighLimitPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.SecondaryHighLimitPrice (1230)
func (f SecondaryHighLimitPriceField) Tag() quickfix.Tag { return tag.SecondaryHighLimitPrice }

//NewSecondaryHighLimitPrice returns a new SecondaryHighLimitPriceField initialized with val and scale
func NewSecondaryHighLimitPrice(val decimal.Decimal, scale int32) SecondaryHighLimitPriceField {
	return SecondaryHighLimitPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SecondaryHighLimitPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//SecondaryIndividualAllocIDField is a STRING field
type SecondaryIndividualAllocIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryIndividualAllocID (989)
func (f SecondaryIndividualAllocIDField) Tag() quickfix.Tag { return tag.SecondaryIndividualAllocID }

//NewSecondaryIndividualAllocID returns a new SecondaryIndividualAllocIDField initialized with val
func NewSecondaryIndividualAllocID(val string) SecondaryIndividualAllocIDField {
	return SecondaryIndividualAllocIDField{quickfix.FIXString(val)}
}

func (f SecondaryIndividualAllocIDField) Value() string { return f.String() }

//SecondaryLowLimitPriceField is a PRICE field
type SecondaryLowLimitPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.SecondaryLowLimitPrice (1221)
func (f SecondaryLowLimitPriceField) Tag() quickfix.Tag { return tag.SecondaryLowLimitPrice }

//NewSecondaryLowLimitPrice returns a new SecondaryLowLimitPriceField initialized with val and scale
func NewSecondaryLowLimitPrice(val decimal.Decimal, scale int32) SecondaryLowLimitPriceField {
	return SecondaryLowLimitPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SecondaryLowLimitPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//SecondaryOrderIDField is a STRING field
type SecondaryOrderIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryOrderID (198)
func (f SecondaryOrderIDField) Tag() quickfix.Tag { return tag.SecondaryOrderID }

//NewSecondaryOrderID returns a new SecondaryOrderIDField initialized with val
func NewSecondaryOrderID(val string) SecondaryOrderIDField {
	return SecondaryOrderIDField{quickfix.FIXString(val)}
}

func (f SecondaryOrderIDField) Value() string { return f.String() }

//SecondaryPriceLimitTypeField is a INT field
type SecondaryPriceLimitTypeField struct{ quickfix.FIXInt }

//Tag returns tag.SecondaryPriceLimitType (1305)
func (f SecondaryPriceLimitTypeField) Tag() quickfix.Tag { return tag.SecondaryPriceLimitType }

//NewSecondaryPriceLimitType returns a new SecondaryPriceLimitTypeField initialized with val
func NewSecondaryPriceLimitType(val int) SecondaryPriceLimitTypeField {
	return SecondaryPriceLimitTypeField{quickfix.FIXInt(val)}
}

func (f SecondaryPriceLimitTypeField) Value() int { return f.Int() }

//SecondaryTradeIDField is a STRING field
type SecondaryTradeIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryTradeID (1040)
func (f SecondaryTradeIDField) Tag() quickfix.Tag { return tag.SecondaryTradeID }

//NewSecondaryTradeID returns a new SecondaryTradeIDField initialized with val
func NewSecondaryTradeID(val string) SecondaryTradeIDField {
	return SecondaryTradeIDField{quickfix.FIXString(val)}
}

func (f SecondaryTradeIDField) Value() string { return f.String() }

//SecondaryTradeReportIDField is a STRING field
type SecondaryTradeReportIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryTradeReportID (818)
func (f SecondaryTradeReportIDField) Tag() quickfix.Tag { return tag.SecondaryTradeReportID }

//NewSecondaryTradeReportID returns a new SecondaryTradeReportIDField initialized with val
func NewSecondaryTradeReportID(val string) SecondaryTradeReportIDField {
	return SecondaryTradeReportIDField{quickfix.FIXString(val)}
}

func (f SecondaryTradeReportIDField) Value() string { return f.String() }

//SecondaryTradeReportRefIDField is a STRING field
type SecondaryTradeReportRefIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryTradeReportRefID (881)
func (f SecondaryTradeReportRefIDField) Tag() quickfix.Tag { return tag.SecondaryTradeReportRefID }

//NewSecondaryTradeReportRefID returns a new SecondaryTradeReportRefIDField initialized with val
func NewSecondaryTradeReportRefID(val string) SecondaryTradeReportRefIDField {
	return SecondaryTradeReportRefIDField{quickfix.FIXString(val)}
}

func (f SecondaryTradeReportRefIDField) Value() string { return f.String() }

//SecondaryTradingReferencePriceField is a PRICE field
type SecondaryTradingReferencePriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.SecondaryTradingReferencePrice (1240)
func (f SecondaryTradingReferencePriceField) Tag() quickfix.Tag {
	return tag.SecondaryTradingReferencePrice
}

//NewSecondaryTradingReferencePrice returns a new SecondaryTradingReferencePriceField initialized with val and scale
func NewSecondaryTradingReferencePrice(val decimal.Decimal, scale int32) SecondaryTradingReferencePriceField {
	return SecondaryTradingReferencePriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SecondaryTradingReferencePriceField) Value() (val decimal.Decimal) { return f.Decimal }

//SecondaryTrdTypeField is a INT field
type SecondaryTrdTypeField struct{ quickfix.FIXInt }

//Tag returns tag.SecondaryTrdType (855)
func (f SecondaryTrdTypeField) Tag() quickfix.Tag { return tag.SecondaryTrdType }

//NewSecondaryTrdType returns a new SecondaryTrdTypeField initialized with val
func NewSecondaryTrdType(val int) SecondaryTrdTypeField {
	return SecondaryTrdTypeField{quickfix.FIXInt(val)}
}

func (f SecondaryTrdTypeField) Value() int { return f.Int() }

//SecureDataField is a DATA field
type SecureDataField struct{ quickfix.FIXString }

//Tag returns tag.SecureData (91)
func (f SecureDataField) Tag() quickfix.Tag { return tag.SecureData }

//NewSecureData returns a new SecureDataField initialized with val
func NewSecureData(val string) SecureDataField {
	return SecureDataField{quickfix.FIXString(val)}
}

func (f SecureDataField) Value() string { return f.String() }

//SecureDataLenField is a LENGTH field
type SecureDataLenField struct{ quickfix.FIXInt }

//Tag returns tag.SecureDataLen (90)
func (f SecureDataLenField) Tag() quickfix.Tag { return tag.SecureDataLen }

//NewSecureDataLen returns a new SecureDataLenField initialized with val
func NewSecureDataLen(val int) SecureDataLenField {
	return SecureDataLenField{quickfix.FIXInt(val)}
}

func (f SecureDataLenField) Value() int { return f.Int() }

//SecurityAltIDField is a STRING field
type SecurityAltIDField struct{ quickfix.FIXString }

//Tag returns tag.SecurityAltID (455)
func (f SecurityAltIDField) Tag() quickfix.Tag { return tag.SecurityAltID }

//NewSecurityAltID returns a new SecurityAltIDField initialized with val
func NewSecurityAltID(val string) SecurityAltIDField {
	return SecurityAltIDField{quickfix.FIXString(val)}
}

func (f SecurityAltIDField) Value() string { return f.String() }

//SecurityAltIDSourceField is a STRING field
type SecurityAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.SecurityAltIDSource (456)
func (f SecurityAltIDSourceField) Tag() quickfix.Tag { return tag.SecurityAltIDSource }

//NewSecurityAltIDSource returns a new SecurityAltIDSourceField initialized with val
func NewSecurityAltIDSource(val string) SecurityAltIDSourceField {
	return SecurityAltIDSourceField{quickfix.FIXString(val)}
}

func (f SecurityAltIDSourceField) Value() string { return f.String() }

//SecurityDescField is a STRING field
type SecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.SecurityDesc (107)
func (f SecurityDescField) Tag() quickfix.Tag { return tag.SecurityDesc }

//NewSecurityDesc returns a new SecurityDescField initialized with val
func NewSecurityDesc(val string) SecurityDescField {
	return SecurityDescField{quickfix.FIXString(val)}
}

func (f SecurityDescField) Value() string { return f.String() }

//SecurityExchangeField is a EXCHANGE field
type SecurityExchangeField struct{ quickfix.FIXString }

//Tag returns tag.SecurityExchange (207)
func (f SecurityExchangeField) Tag() quickfix.Tag { return tag.SecurityExchange }

//NewSecurityExchange returns a new SecurityExchangeField initialized with val
func NewSecurityExchange(val string) SecurityExchangeField {
	return SecurityExchangeField{quickfix.FIXString(val)}
}

func (f SecurityExchangeField) Value() string { return f.String() }

//SecurityGroupField is a STRING field
type SecurityGroupField struct{ quickfix.FIXString }

//Tag returns tag.SecurityGroup (1151)
func (f SecurityGroupField) Tag() quickfix.Tag { return tag.SecurityGroup }

//NewSecurityGroup returns a new SecurityGroupField initialized with val
func NewSecurityGroup(val string) SecurityGroupField {
	return SecurityGroupField{quickfix.FIXString(val)}
}

func (f SecurityGroupField) Value() string { return f.String() }

//SecurityIDField is a STRING field
type SecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.SecurityID (48)
func (f SecurityIDField) Tag() quickfix.Tag { return tag.SecurityID }

//NewSecurityID returns a new SecurityIDField initialized with val
func NewSecurityID(val string) SecurityIDField {
	return SecurityIDField{quickfix.FIXString(val)}
}

func (f SecurityIDField) Value() string { return f.String() }

//SecurityIDSourceField is a enum.SecurityIDSource field
type SecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.SecurityIDSource (22)
func (f SecurityIDSourceField) Tag() quickfix.Tag { return tag.SecurityIDSource }

func NewSecurityIDSource(val enum.SecurityIDSource) SecurityIDSourceField {
	return SecurityIDSourceField{quickfix.FIXString(val)}
}

func (f SecurityIDSourceField) Value() enum.SecurityIDSource { return enum.SecurityIDSource(f.String()) }

//SecurityListDescField is a STRING field
type SecurityListDescField struct{ quickfix.FIXString }

//Tag returns tag.SecurityListDesc (1467)
func (f SecurityListDescField) Tag() quickfix.Tag { return tag.SecurityListDesc }

//NewSecurityListDesc returns a new SecurityListDescField initialized with val
func NewSecurityListDesc(val string) SecurityListDescField {
	return SecurityListDescField{quickfix.FIXString(val)}
}

func (f SecurityListDescField) Value() string { return f.String() }

//SecurityListIDField is a STRING field
type SecurityListIDField struct{ quickfix.FIXString }

//Tag returns tag.SecurityListID (1465)
func (f SecurityListIDField) Tag() quickfix.Tag { return tag.SecurityListID }

//NewSecurityListID returns a new SecurityListIDField initialized with val
func NewSecurityListID(val string) SecurityListIDField {
	return SecurityListIDField{quickfix.FIXString(val)}
}

func (f SecurityListIDField) Value() string { return f.String() }

//SecurityListRefIDField is a STRING field
type SecurityListRefIDField struct{ quickfix.FIXString }

//Tag returns tag.SecurityListRefID (1466)
func (f SecurityListRefIDField) Tag() quickfix.Tag { return tag.SecurityListRefID }

//NewSecurityListRefID returns a new SecurityListRefIDField initialized with val
func NewSecurityListRefID(val string) SecurityListRefIDField {
	return SecurityListRefIDField{quickfix.FIXString(val)}
}

func (f SecurityListRefIDField) Value() string { return f.String() }

//SecurityListRequestTypeField is a enum.SecurityListRequestType field
type SecurityListRequestTypeField struct{ quickfix.FIXString }

//Tag returns tag.SecurityListRequestType (559)
func (f SecurityListRequestTypeField) Tag() quickfix.Tag { return tag.SecurityListRequestType }

func NewSecurityListRequestType(val enum.SecurityListRequestType) SecurityListRequestTypeField {
	return SecurityListRequestTypeField{quickfix.FIXString(val)}
}

func (f SecurityListRequestTypeField) Value() enum.SecurityListRequestType {
	return enum.SecurityListRequestType(f.String())
}

//SecurityListTypeField is a enum.SecurityListType field
type SecurityListTypeField struct{ quickfix.FIXString }

//Tag returns tag.SecurityListType (1470)
func (f SecurityListTypeField) Tag() quickfix.Tag { return tag.SecurityListType }

func NewSecurityListType(val enum.SecurityListType) SecurityListTypeField {
	return SecurityListTypeField{quickfix.FIXString(val)}
}

func (f SecurityListTypeField) Value() enum.SecurityListType { return enum.SecurityListType(f.String()) }

//SecurityListTypeSourceField is a enum.SecurityListTypeSource field
type SecurityListTypeSourceField struct{ quickfix.FIXString }

//Tag returns tag.SecurityListTypeSource (1471)
func (f SecurityListTypeSourceField) Tag() quickfix.Tag { return tag.SecurityListTypeSource }

func NewSecurityListTypeSource(val enum.SecurityListTypeSource) SecurityListTypeSourceField {
	return SecurityListTypeSourceField{quickfix.FIXString(val)}
}

func (f SecurityListTypeSourceField) Value() enum.SecurityListTypeSource {
	return enum.SecurityListTypeSource(f.String())
}

//SecurityReportIDField is a INT field
type SecurityReportIDField struct{ quickfix.FIXInt }

//Tag returns tag.SecurityReportID (964)
func (f SecurityReportIDField) Tag() quickfix.Tag { return tag.SecurityReportID }

//NewSecurityReportID returns a new SecurityReportIDField initialized with val
func NewSecurityReportID(val int) SecurityReportIDField {
	return SecurityReportIDField{quickfix.FIXInt(val)}
}

func (f SecurityReportIDField) Value() int { return f.Int() }

//SecurityReqIDField is a STRING field
type SecurityReqIDField struct{ quickfix.FIXString }

//Tag returns tag.SecurityReqID (320)
func (f SecurityReqIDField) Tag() quickfix.Tag { return tag.SecurityReqID }

//NewSecurityReqID returns a new SecurityReqIDField initialized with val
func NewSecurityReqID(val string) SecurityReqIDField {
	return SecurityReqIDField{quickfix.FIXString(val)}
}

func (f SecurityReqIDField) Value() string { return f.String() }

//SecurityRequestResultField is a enum.SecurityRequestResult field
type SecurityRequestResultField struct{ quickfix.FIXString }

//Tag returns tag.SecurityRequestResult (560)
func (f SecurityRequestResultField) Tag() quickfix.Tag { return tag.SecurityRequestResult }

func NewSecurityRequestResult(val enum.SecurityRequestResult) SecurityRequestResultField {
	return SecurityRequestResultField{quickfix.FIXString(val)}
}

func (f SecurityRequestResultField) Value() enum.SecurityRequestResult {
	return enum.SecurityRequestResult(f.String())
}

//SecurityRequestTypeField is a enum.SecurityRequestType field
type SecurityRequestTypeField struct{ quickfix.FIXString }

//Tag returns tag.SecurityRequestType (321)
func (f SecurityRequestTypeField) Tag() quickfix.Tag { return tag.SecurityRequestType }

func NewSecurityRequestType(val enum.SecurityRequestType) SecurityRequestTypeField {
	return SecurityRequestTypeField{quickfix.FIXString(val)}
}

func (f SecurityRequestTypeField) Value() enum.SecurityRequestType {
	return enum.SecurityRequestType(f.String())
}

//SecurityResponseIDField is a STRING field
type SecurityResponseIDField struct{ quickfix.FIXString }

//Tag returns tag.SecurityResponseID (322)
func (f SecurityResponseIDField) Tag() quickfix.Tag { return tag.SecurityResponseID }

//NewSecurityResponseID returns a new SecurityResponseIDField initialized with val
func NewSecurityResponseID(val string) SecurityResponseIDField {
	return SecurityResponseIDField{quickfix.FIXString(val)}
}

func (f SecurityResponseIDField) Value() string { return f.String() }

//SecurityResponseTypeField is a enum.SecurityResponseType field
type SecurityResponseTypeField struct{ quickfix.FIXString }

//Tag returns tag.SecurityResponseType (323)
func (f SecurityResponseTypeField) Tag() quickfix.Tag { return tag.SecurityResponseType }

func NewSecurityResponseType(val enum.SecurityResponseType) SecurityResponseTypeField {
	return SecurityResponseTypeField{quickfix.FIXString(val)}
}

func (f SecurityResponseTypeField) Value() enum.SecurityResponseType {
	return enum.SecurityResponseType(f.String())
}

//SecuritySettlAgentAcctNameField is a STRING field
type SecuritySettlAgentAcctNameField struct{ quickfix.FIXString }

//Tag returns tag.SecuritySettlAgentAcctName (179)
func (f SecuritySettlAgentAcctNameField) Tag() quickfix.Tag { return tag.SecuritySettlAgentAcctName }

//NewSecuritySettlAgentAcctName returns a new SecuritySettlAgentAcctNameField initialized with val
func NewSecuritySettlAgentAcctName(val string) SecuritySettlAgentAcctNameField {
	return SecuritySettlAgentAcctNameField{quickfix.FIXString(val)}
}

func (f SecuritySettlAgentAcctNameField) Value() string { return f.String() }

//SecuritySettlAgentAcctNumField is a STRING field
type SecuritySettlAgentAcctNumField struct{ quickfix.FIXString }

//Tag returns tag.SecuritySettlAgentAcctNum (178)
func (f SecuritySettlAgentAcctNumField) Tag() quickfix.Tag { return tag.SecuritySettlAgentAcctNum }

//NewSecuritySettlAgentAcctNum returns a new SecuritySettlAgentAcctNumField initialized with val
func NewSecuritySettlAgentAcctNum(val string) SecuritySettlAgentAcctNumField {
	return SecuritySettlAgentAcctNumField{quickfix.FIXString(val)}
}

func (f SecuritySettlAgentAcctNumField) Value() string { return f.String() }

//SecuritySettlAgentCodeField is a STRING field
type SecuritySettlAgentCodeField struct{ quickfix.FIXString }

//Tag returns tag.SecuritySettlAgentCode (177)
func (f SecuritySettlAgentCodeField) Tag() quickfix.Tag { return tag.SecuritySettlAgentCode }

//NewSecuritySettlAgentCode returns a new SecuritySettlAgentCodeField initialized with val
func NewSecuritySettlAgentCode(val string) SecuritySettlAgentCodeField {
	return SecuritySettlAgentCodeField{quickfix.FIXString(val)}
}

func (f SecuritySettlAgentCodeField) Value() string { return f.String() }

//SecuritySettlAgentContactNameField is a STRING field
type SecuritySettlAgentContactNameField struct{ quickfix.FIXString }

//Tag returns tag.SecuritySettlAgentContactName (180)
func (f SecuritySettlAgentContactNameField) Tag() quickfix.Tag {
	return tag.SecuritySettlAgentContactName
}

//NewSecuritySettlAgentContactName returns a new SecuritySettlAgentContactNameField initialized with val
func NewSecuritySettlAgentContactName(val string) SecuritySettlAgentContactNameField {
	return SecuritySettlAgentContactNameField{quickfix.FIXString(val)}
}

func (f SecuritySettlAgentContactNameField) Value() string { return f.String() }

//SecuritySettlAgentContactPhoneField is a STRING field
type SecuritySettlAgentContactPhoneField struct{ quickfix.FIXString }

//Tag returns tag.SecuritySettlAgentContactPhone (181)
func (f SecuritySettlAgentContactPhoneField) Tag() quickfix.Tag {
	return tag.SecuritySettlAgentContactPhone
}

//NewSecuritySettlAgentContactPhone returns a new SecuritySettlAgentContactPhoneField initialized with val
func NewSecuritySettlAgentContactPhone(val string) SecuritySettlAgentContactPhoneField {
	return SecuritySettlAgentContactPhoneField{quickfix.FIXString(val)}
}

func (f SecuritySettlAgentContactPhoneField) Value() string { return f.String() }

//SecuritySettlAgentNameField is a STRING field
type SecuritySettlAgentNameField struct{ quickfix.FIXString }

//Tag returns tag.SecuritySettlAgentName (176)
func (f SecuritySettlAgentNameField) Tag() quickfix.Tag { return tag.SecuritySettlAgentName }

//NewSecuritySettlAgentName returns a new SecuritySettlAgentNameField initialized with val
func NewSecuritySettlAgentName(val string) SecuritySettlAgentNameField {
	return SecuritySettlAgentNameField{quickfix.FIXString(val)}
}

func (f SecuritySettlAgentNameField) Value() string { return f.String() }

//SecurityStatusField is a enum.SecurityStatus field
type SecurityStatusField struct{ quickfix.FIXString }

//Tag returns tag.SecurityStatus (965)
func (f SecurityStatusField) Tag() quickfix.Tag { return tag.SecurityStatus }

func NewSecurityStatus(val enum.SecurityStatus) SecurityStatusField {
	return SecurityStatusField{quickfix.FIXString(val)}
}

func (f SecurityStatusField) Value() enum.SecurityStatus { return enum.SecurityStatus(f.String()) }

//SecurityStatusReqIDField is a STRING field
type SecurityStatusReqIDField struct{ quickfix.FIXString }

//Tag returns tag.SecurityStatusReqID (324)
func (f SecurityStatusReqIDField) Tag() quickfix.Tag { return tag.SecurityStatusReqID }

//NewSecurityStatusReqID returns a new SecurityStatusReqIDField initialized with val
func NewSecurityStatusReqID(val string) SecurityStatusReqIDField {
	return SecurityStatusReqIDField{quickfix.FIXString(val)}
}

func (f SecurityStatusReqIDField) Value() string { return f.String() }

//SecuritySubTypeField is a STRING field
type SecuritySubTypeField struct{ quickfix.FIXString }

//Tag returns tag.SecuritySubType (762)
func (f SecuritySubTypeField) Tag() quickfix.Tag { return tag.SecuritySubType }

//NewSecuritySubType returns a new SecuritySubTypeField initialized with val
func NewSecuritySubType(val string) SecuritySubTypeField {
	return SecuritySubTypeField{quickfix.FIXString(val)}
}

func (f SecuritySubTypeField) Value() string { return f.String() }

//SecurityTradingEventField is a enum.SecurityTradingEvent field
type SecurityTradingEventField struct{ quickfix.FIXString }

//Tag returns tag.SecurityTradingEvent (1174)
func (f SecurityTradingEventField) Tag() quickfix.Tag { return tag.SecurityTradingEvent }

func NewSecurityTradingEvent(val enum.SecurityTradingEvent) SecurityTradingEventField {
	return SecurityTradingEventField{quickfix.FIXString(val)}
}

func (f SecurityTradingEventField) Value() enum.SecurityTradingEvent {
	return enum.SecurityTradingEvent(f.String())
}

//SecurityTradingStatusField is a enum.SecurityTradingStatus field
type SecurityTradingStatusField struct{ quickfix.FIXString }

//Tag returns tag.SecurityTradingStatus (326)
func (f SecurityTradingStatusField) Tag() quickfix.Tag { return tag.SecurityTradingStatus }

func NewSecurityTradingStatus(val enum.SecurityTradingStatus) SecurityTradingStatusField {
	return SecurityTradingStatusField{quickfix.FIXString(val)}
}

func (f SecurityTradingStatusField) Value() enum.SecurityTradingStatus {
	return enum.SecurityTradingStatus(f.String())
}

//SecurityTypeField is a enum.SecurityType field
type SecurityTypeField struct{ quickfix.FIXString }

//Tag returns tag.SecurityType (167)
func (f SecurityTypeField) Tag() quickfix.Tag { return tag.SecurityType }

func NewSecurityType(val enum.SecurityType) SecurityTypeField {
	return SecurityTypeField{quickfix.FIXString(val)}
}

func (f SecurityTypeField) Value() enum.SecurityType { return enum.SecurityType(f.String()) }

//SecurityUpdateActionField is a enum.SecurityUpdateAction field
type SecurityUpdateActionField struct{ quickfix.FIXString }

//Tag returns tag.SecurityUpdateAction (980)
func (f SecurityUpdateActionField) Tag() quickfix.Tag { return tag.SecurityUpdateAction }

func NewSecurityUpdateAction(val enum.SecurityUpdateAction) SecurityUpdateActionField {
	return SecurityUpdateActionField{quickfix.FIXString(val)}
}

func (f SecurityUpdateActionField) Value() enum.SecurityUpdateAction {
	return enum.SecurityUpdateAction(f.String())
}

//SecurityXMLField is a XMLDATA field
type SecurityXMLField struct{ quickfix.FIXString }

//Tag returns tag.SecurityXML (1185)
func (f SecurityXMLField) Tag() quickfix.Tag { return tag.SecurityXML }

//NewSecurityXML returns a new SecurityXMLField initialized with val
func NewSecurityXML(val string) SecurityXMLField {
	return SecurityXMLField{quickfix.FIXString(val)}
}

func (f SecurityXMLField) Value() string { return f.String() }

//SecurityXMLLenField is a LENGTH field
type SecurityXMLLenField struct{ quickfix.FIXInt }

//Tag returns tag.SecurityXMLLen (1184)
func (f SecurityXMLLenField) Tag() quickfix.Tag { return tag.SecurityXMLLen }

//NewSecurityXMLLen returns a new SecurityXMLLenField initialized with val
func NewSecurityXMLLen(val int) SecurityXMLLenField {
	return SecurityXMLLenField{quickfix.FIXInt(val)}
}

func (f SecurityXMLLenField) Value() int { return f.Int() }

//SecurityXMLSchemaField is a STRING field
type SecurityXMLSchemaField struct{ quickfix.FIXString }

//Tag returns tag.SecurityXMLSchema (1186)
func (f SecurityXMLSchemaField) Tag() quickfix.Tag { return tag.SecurityXMLSchema }

//NewSecurityXMLSchema returns a new SecurityXMLSchemaField initialized with val
func NewSecurityXMLSchema(val string) SecurityXMLSchemaField {
	return SecurityXMLSchemaField{quickfix.FIXString(val)}
}

func (f SecurityXMLSchemaField) Value() string { return f.String() }

//SellVolumeField is a QTY field
type SellVolumeField struct{ quickfix.FIXDecimal }

//Tag returns tag.SellVolume (331)
func (f SellVolumeField) Tag() quickfix.Tag { return tag.SellVolume }

//NewSellVolume returns a new SellVolumeField initialized with val and scale
func NewSellVolume(val decimal.Decimal, scale int32) SellVolumeField {
	return SellVolumeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SellVolumeField) Value() (val decimal.Decimal) { return f.Decimal }

//SellerDaysField is a INT field
type SellerDaysField struct{ quickfix.FIXInt }

//Tag returns tag.SellerDays (287)
func (f SellerDaysField) Tag() quickfix.Tag { return tag.SellerDays }

//NewSellerDays returns a new SellerDaysField initialized with val
func NewSellerDays(val int) SellerDaysField {
	return SellerDaysField{quickfix.FIXInt(val)}
}

func (f SellerDaysField) Value() int { return f.Int() }

//SenderCompIDField is a STRING field
type SenderCompIDField struct{ quickfix.FIXString }

//Tag returns tag.SenderCompID (49)
func (f SenderCompIDField) Tag() quickfix.Tag { return tag.SenderCompID }

//NewSenderCompID returns a new SenderCompIDField initialized with val
func NewSenderCompID(val string) SenderCompIDField {
	return SenderCompIDField{quickfix.FIXString(val)}
}

func (f SenderCompIDField) Value() string { return f.String() }

//SenderLocationIDField is a STRING field
type SenderLocationIDField struct{ quickfix.FIXString }

//Tag returns tag.SenderLocationID (142)
func (f SenderLocationIDField) Tag() quickfix.Tag { return tag.SenderLocationID }

//NewSenderLocationID returns a new SenderLocationIDField initialized with val
func NewSenderLocationID(val string) SenderLocationIDField {
	return SenderLocationIDField{quickfix.FIXString(val)}
}

func (f SenderLocationIDField) Value() string { return f.String() }

//SenderSubIDField is a STRING field
type SenderSubIDField struct{ quickfix.FIXString }

//Tag returns tag.SenderSubID (50)
func (f SenderSubIDField) Tag() quickfix.Tag { return tag.SenderSubID }

//NewSenderSubID returns a new SenderSubIDField initialized with val
func NewSenderSubID(val string) SenderSubIDField {
	return SenderSubIDField{quickfix.FIXString(val)}
}

func (f SenderSubIDField) Value() string { return f.String() }

//SendingDateField is a LOCALMKTDATE field
type SendingDateField struct{ quickfix.FIXString }

//Tag returns tag.SendingDate (51)
func (f SendingDateField) Tag() quickfix.Tag { return tag.SendingDate }

//NewSendingDate returns a new SendingDateField initialized with val
func NewSendingDate(val string) SendingDateField {
	return SendingDateField{quickfix.FIXString(val)}
}

func (f SendingDateField) Value() string { return f.String() }

//SendingTimeField is a UTCTIMESTAMP field
type SendingTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.SendingTime (52)
func (f SendingTimeField) Tag() quickfix.Tag { return tag.SendingTime }

//NewSendingTime returns a new SendingTimeField initialized with val
func NewSendingTime(val time.Time) SendingTimeField {
	return SendingTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewSendingTimeNoMillis returns a new SendingTimeField initialized with val without millisecs
func NewSendingTimeNoMillis(val time.Time) SendingTimeField {
	return SendingTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f SendingTimeField) Value() time.Time { return f.Time }

//SeniorityField is a enum.Seniority field
type SeniorityField struct{ quickfix.FIXString }

//Tag returns tag.Seniority (1450)
func (f SeniorityField) Tag() quickfix.Tag { return tag.Seniority }

func NewSeniority(val enum.Seniority) SeniorityField {
	return SeniorityField{quickfix.FIXString(val)}
}

func (f SeniorityField) Value() enum.Seniority { return enum.Seniority(f.String()) }

//SessionRejectReasonField is a enum.SessionRejectReason field
type SessionRejectReasonField struct{ quickfix.FIXString }

//Tag returns tag.SessionRejectReason (373)
func (f SessionRejectReasonField) Tag() quickfix.Tag { return tag.SessionRejectReason }

func NewSessionRejectReason(val enum.SessionRejectReason) SessionRejectReasonField {
	return SessionRejectReasonField{quickfix.FIXString(val)}
}

func (f SessionRejectReasonField) Value() enum.SessionRejectReason {
	return enum.SessionRejectReason(f.String())
}

//SessionStatusField is a enum.SessionStatus field
type SessionStatusField struct{ quickfix.FIXString }

//Tag returns tag.SessionStatus (1409)
func (f SessionStatusField) Tag() quickfix.Tag { return tag.SessionStatus }

func NewSessionStatus(val enum.SessionStatus) SessionStatusField {
	return SessionStatusField{quickfix.FIXString(val)}
}

func (f SessionStatusField) Value() enum.SessionStatus { return enum.SessionStatus(f.String()) }

//SettlBrkrCodeField is a STRING field
type SettlBrkrCodeField struct{ quickfix.FIXString }

//Tag returns tag.SettlBrkrCode (174)
func (f SettlBrkrCodeField) Tag() quickfix.Tag { return tag.SettlBrkrCode }

//NewSettlBrkrCode returns a new SettlBrkrCodeField initialized with val
func NewSettlBrkrCode(val string) SettlBrkrCodeField {
	return SettlBrkrCodeField{quickfix.FIXString(val)}
}

func (f SettlBrkrCodeField) Value() string { return f.String() }

//SettlCurrAmtField is a AMT field
type SettlCurrAmtField struct{ quickfix.FIXDecimal }

//Tag returns tag.SettlCurrAmt (119)
func (f SettlCurrAmtField) Tag() quickfix.Tag { return tag.SettlCurrAmt }

//NewSettlCurrAmt returns a new SettlCurrAmtField initialized with val and scale
func NewSettlCurrAmt(val decimal.Decimal, scale int32) SettlCurrAmtField {
	return SettlCurrAmtField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SettlCurrAmtField) Value() (val decimal.Decimal) { return f.Decimal }

//SettlCurrBidFxRateField is a FLOAT field
type SettlCurrBidFxRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.SettlCurrBidFxRate (656)
func (f SettlCurrBidFxRateField) Tag() quickfix.Tag { return tag.SettlCurrBidFxRate }

//NewSettlCurrBidFxRate returns a new SettlCurrBidFxRateField initialized with val and scale
func NewSettlCurrBidFxRate(val decimal.Decimal, scale int32) SettlCurrBidFxRateField {
	return SettlCurrBidFxRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SettlCurrBidFxRateField) Value() (val decimal.Decimal) { return f.Decimal }

//SettlCurrFxRateField is a FLOAT field
type SettlCurrFxRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.SettlCurrFxRate (155)
func (f SettlCurrFxRateField) Tag() quickfix.Tag { return tag.SettlCurrFxRate }

//NewSettlCurrFxRate returns a new SettlCurrFxRateField initialized with val and scale
func NewSettlCurrFxRate(val decimal.Decimal, scale int32) SettlCurrFxRateField {
	return SettlCurrFxRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SettlCurrFxRateField) Value() (val decimal.Decimal) { return f.Decimal }

//SettlCurrFxRateCalcField is a enum.SettlCurrFxRateCalc field
type SettlCurrFxRateCalcField struct{ quickfix.FIXString }

//Tag returns tag.SettlCurrFxRateCalc (156)
func (f SettlCurrFxRateCalcField) Tag() quickfix.Tag { return tag.SettlCurrFxRateCalc }

func NewSettlCurrFxRateCalc(val enum.SettlCurrFxRateCalc) SettlCurrFxRateCalcField {
	return SettlCurrFxRateCalcField{quickfix.FIXString(val)}
}

func (f SettlCurrFxRateCalcField) Value() enum.SettlCurrFxRateCalc {
	return enum.SettlCurrFxRateCalc(f.String())
}

//SettlCurrOfferFxRateField is a FLOAT field
type SettlCurrOfferFxRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.SettlCurrOfferFxRate (657)
func (f SettlCurrOfferFxRateField) Tag() quickfix.Tag { return tag.SettlCurrOfferFxRate }

//NewSettlCurrOfferFxRate returns a new SettlCurrOfferFxRateField initialized with val and scale
func NewSettlCurrOfferFxRate(val decimal.Decimal, scale int32) SettlCurrOfferFxRateField {
	return SettlCurrOfferFxRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SettlCurrOfferFxRateField) Value() (val decimal.Decimal) { return f.Decimal }

//SettlCurrencyField is a CURRENCY field
type SettlCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.SettlCurrency (120)
func (f SettlCurrencyField) Tag() quickfix.Tag { return tag.SettlCurrency }

//NewSettlCurrency returns a new SettlCurrencyField initialized with val
func NewSettlCurrency(val string) SettlCurrencyField {
	return SettlCurrencyField{quickfix.FIXString(val)}
}

func (f SettlCurrencyField) Value() string { return f.String() }

//SettlDateField is a LOCALMKTDATE field
type SettlDateField struct{ quickfix.FIXString }

//Tag returns tag.SettlDate (64)
func (f SettlDateField) Tag() quickfix.Tag { return tag.SettlDate }

//NewSettlDate returns a new SettlDateField initialized with val
func NewSettlDate(val string) SettlDateField {
	return SettlDateField{quickfix.FIXString(val)}
}

func (f SettlDateField) Value() string { return f.String() }

//SettlDate2Field is a LOCALMKTDATE field
type SettlDate2Field struct{ quickfix.FIXString }

//Tag returns tag.SettlDate2 (193)
func (f SettlDate2Field) Tag() quickfix.Tag { return tag.SettlDate2 }

//NewSettlDate2 returns a new SettlDate2Field initialized with val
func NewSettlDate2(val string) SettlDate2Field {
	return SettlDate2Field{quickfix.FIXString(val)}
}

func (f SettlDate2Field) Value() string { return f.String() }

//SettlDeliveryTypeField is a enum.SettlDeliveryType field
type SettlDeliveryTypeField struct{ quickfix.FIXString }

//Tag returns tag.SettlDeliveryType (172)
func (f SettlDeliveryTypeField) Tag() quickfix.Tag { return tag.SettlDeliveryType }

func NewSettlDeliveryType(val enum.SettlDeliveryType) SettlDeliveryTypeField {
	return SettlDeliveryTypeField{quickfix.FIXString(val)}
}

func (f SettlDeliveryTypeField) Value() enum.SettlDeliveryType {
	return enum.SettlDeliveryType(f.String())
}

//SettlDepositoryCodeField is a STRING field
type SettlDepositoryCodeField struct{ quickfix.FIXString }

//Tag returns tag.SettlDepositoryCode (173)
func (f SettlDepositoryCodeField) Tag() quickfix.Tag { return tag.SettlDepositoryCode }

//NewSettlDepositoryCode returns a new SettlDepositoryCodeField initialized with val
func NewSettlDepositoryCode(val string) SettlDepositoryCodeField {
	return SettlDepositoryCodeField{quickfix.FIXString(val)}
}

func (f SettlDepositoryCodeField) Value() string { return f.String() }

//SettlInstCodeField is a STRING field
type SettlInstCodeField struct{ quickfix.FIXString }

//Tag returns tag.SettlInstCode (175)
func (f SettlInstCodeField) Tag() quickfix.Tag { return tag.SettlInstCode }

//NewSettlInstCode returns a new SettlInstCodeField initialized with val
func NewSettlInstCode(val string) SettlInstCodeField {
	return SettlInstCodeField{quickfix.FIXString(val)}
}

func (f SettlInstCodeField) Value() string { return f.String() }

//SettlInstIDField is a STRING field
type SettlInstIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlInstID (162)
func (f SettlInstIDField) Tag() quickfix.Tag { return tag.SettlInstID }

//NewSettlInstID returns a new SettlInstIDField initialized with val
func NewSettlInstID(val string) SettlInstIDField {
	return SettlInstIDField{quickfix.FIXString(val)}
}

func (f SettlInstIDField) Value() string { return f.String() }

//SettlInstModeField is a enum.SettlInstMode field
type SettlInstModeField struct{ quickfix.FIXString }

//Tag returns tag.SettlInstMode (160)
func (f SettlInstModeField) Tag() quickfix.Tag { return tag.SettlInstMode }

func NewSettlInstMode(val enum.SettlInstMode) SettlInstModeField {
	return SettlInstModeField{quickfix.FIXString(val)}
}

func (f SettlInstModeField) Value() enum.SettlInstMode { return enum.SettlInstMode(f.String()) }

//SettlInstMsgIDField is a STRING field
type SettlInstMsgIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlInstMsgID (777)
func (f SettlInstMsgIDField) Tag() quickfix.Tag { return tag.SettlInstMsgID }

//NewSettlInstMsgID returns a new SettlInstMsgIDField initialized with val
func NewSettlInstMsgID(val string) SettlInstMsgIDField {
	return SettlInstMsgIDField{quickfix.FIXString(val)}
}

func (f SettlInstMsgIDField) Value() string { return f.String() }

//SettlInstRefIDField is a STRING field
type SettlInstRefIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlInstRefID (214)
func (f SettlInstRefIDField) Tag() quickfix.Tag { return tag.SettlInstRefID }

//NewSettlInstRefID returns a new SettlInstRefIDField initialized with val
func NewSettlInstRefID(val string) SettlInstRefIDField {
	return SettlInstRefIDField{quickfix.FIXString(val)}
}

func (f SettlInstRefIDField) Value() string { return f.String() }

//SettlInstReqIDField is a STRING field
type SettlInstReqIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlInstReqID (791)
func (f SettlInstReqIDField) Tag() quickfix.Tag { return tag.SettlInstReqID }

//NewSettlInstReqID returns a new SettlInstReqIDField initialized with val
func NewSettlInstReqID(val string) SettlInstReqIDField {
	return SettlInstReqIDField{quickfix.FIXString(val)}
}

func (f SettlInstReqIDField) Value() string { return f.String() }

//SettlInstReqRejCodeField is a enum.SettlInstReqRejCode field
type SettlInstReqRejCodeField struct{ quickfix.FIXString }

//Tag returns tag.SettlInstReqRejCode (792)
func (f SettlInstReqRejCodeField) Tag() quickfix.Tag { return tag.SettlInstReqRejCode }

func NewSettlInstReqRejCode(val enum.SettlInstReqRejCode) SettlInstReqRejCodeField {
	return SettlInstReqRejCodeField{quickfix.FIXString(val)}
}

func (f SettlInstReqRejCodeField) Value() enum.SettlInstReqRejCode {
	return enum.SettlInstReqRejCode(f.String())
}

//SettlInstSourceField is a enum.SettlInstSource field
type SettlInstSourceField struct{ quickfix.FIXString }

//Tag returns tag.SettlInstSource (165)
func (f SettlInstSourceField) Tag() quickfix.Tag { return tag.SettlInstSource }

func NewSettlInstSource(val enum.SettlInstSource) SettlInstSourceField {
	return SettlInstSourceField{quickfix.FIXString(val)}
}

func (f SettlInstSourceField) Value() enum.SettlInstSource { return enum.SettlInstSource(f.String()) }

//SettlInstTransTypeField is a enum.SettlInstTransType field
type SettlInstTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.SettlInstTransType (163)
func (f SettlInstTransTypeField) Tag() quickfix.Tag { return tag.SettlInstTransType }

func NewSettlInstTransType(val enum.SettlInstTransType) SettlInstTransTypeField {
	return SettlInstTransTypeField{quickfix.FIXString(val)}
}

func (f SettlInstTransTypeField) Value() enum.SettlInstTransType {
	return enum.SettlInstTransType(f.String())
}

//SettlLocationField is a enum.SettlLocation field
type SettlLocationField struct{ quickfix.FIXString }

//Tag returns tag.SettlLocation (166)
func (f SettlLocationField) Tag() quickfix.Tag { return tag.SettlLocation }

func NewSettlLocation(val enum.SettlLocation) SettlLocationField {
	return SettlLocationField{quickfix.FIXString(val)}
}

func (f SettlLocationField) Value() enum.SettlLocation { return enum.SettlLocation(f.String()) }

//SettlMethodField is a enum.SettlMethod field
type SettlMethodField struct{ quickfix.FIXString }

//Tag returns tag.SettlMethod (1193)
func (f SettlMethodField) Tag() quickfix.Tag { return tag.SettlMethod }

func NewSettlMethod(val enum.SettlMethod) SettlMethodField {
	return SettlMethodField{quickfix.FIXString(val)}
}

func (f SettlMethodField) Value() enum.SettlMethod { return enum.SettlMethod(f.String()) }

//SettlObligIDField is a STRING field
type SettlObligIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlObligID (1161)
func (f SettlObligIDField) Tag() quickfix.Tag { return tag.SettlObligID }

//NewSettlObligID returns a new SettlObligIDField initialized with val
func NewSettlObligID(val string) SettlObligIDField {
	return SettlObligIDField{quickfix.FIXString(val)}
}

func (f SettlObligIDField) Value() string { return f.String() }

//SettlObligModeField is a enum.SettlObligMode field
type SettlObligModeField struct{ quickfix.FIXString }

//Tag returns tag.SettlObligMode (1159)
func (f SettlObligModeField) Tag() quickfix.Tag { return tag.SettlObligMode }

func NewSettlObligMode(val enum.SettlObligMode) SettlObligModeField {
	return SettlObligModeField{quickfix.FIXString(val)}
}

func (f SettlObligModeField) Value() enum.SettlObligMode { return enum.SettlObligMode(f.String()) }

//SettlObligMsgIDField is a STRING field
type SettlObligMsgIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlObligMsgID (1160)
func (f SettlObligMsgIDField) Tag() quickfix.Tag { return tag.SettlObligMsgID }

//NewSettlObligMsgID returns a new SettlObligMsgIDField initialized with val
func NewSettlObligMsgID(val string) SettlObligMsgIDField {
	return SettlObligMsgIDField{quickfix.FIXString(val)}
}

func (f SettlObligMsgIDField) Value() string { return f.String() }

//SettlObligRefIDField is a STRING field
type SettlObligRefIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlObligRefID (1163)
func (f SettlObligRefIDField) Tag() quickfix.Tag { return tag.SettlObligRefID }

//NewSettlObligRefID returns a new SettlObligRefIDField initialized with val
func NewSettlObligRefID(val string) SettlObligRefIDField {
	return SettlObligRefIDField{quickfix.FIXString(val)}
}

func (f SettlObligRefIDField) Value() string { return f.String() }

//SettlObligSourceField is a enum.SettlObligSource field
type SettlObligSourceField struct{ quickfix.FIXString }

//Tag returns tag.SettlObligSource (1164)
func (f SettlObligSourceField) Tag() quickfix.Tag { return tag.SettlObligSource }

func NewSettlObligSource(val enum.SettlObligSource) SettlObligSourceField {
	return SettlObligSourceField{quickfix.FIXString(val)}
}

func (f SettlObligSourceField) Value() enum.SettlObligSource { return enum.SettlObligSource(f.String()) }

//SettlObligTransTypeField is a enum.SettlObligTransType field
type SettlObligTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.SettlObligTransType (1162)
func (f SettlObligTransTypeField) Tag() quickfix.Tag { return tag.SettlObligTransType }

func NewSettlObligTransType(val enum.SettlObligTransType) SettlObligTransTypeField {
	return SettlObligTransTypeField{quickfix.FIXString(val)}
}

func (f SettlObligTransTypeField) Value() enum.SettlObligTransType {
	return enum.SettlObligTransType(f.String())
}

//SettlPartyIDField is a STRING field
type SettlPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlPartyID (782)
func (f SettlPartyIDField) Tag() quickfix.Tag { return tag.SettlPartyID }

//NewSettlPartyID returns a new SettlPartyIDField initialized with val
func NewSettlPartyID(val string) SettlPartyIDField {
	return SettlPartyIDField{quickfix.FIXString(val)}
}

func (f SettlPartyIDField) Value() string { return f.String() }

//SettlPartyIDSourceField is a CHAR field
type SettlPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.SettlPartyIDSource (783)
func (f SettlPartyIDSourceField) Tag() quickfix.Tag { return tag.SettlPartyIDSource }

//NewSettlPartyIDSource returns a new SettlPartyIDSourceField initialized with val
func NewSettlPartyIDSource(val string) SettlPartyIDSourceField {
	return SettlPartyIDSourceField{quickfix.FIXString(val)}
}

func (f SettlPartyIDSourceField) Value() string { return f.String() }

//SettlPartyRoleField is a INT field
type SettlPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.SettlPartyRole (784)
func (f SettlPartyRoleField) Tag() quickfix.Tag { return tag.SettlPartyRole }

//NewSettlPartyRole returns a new SettlPartyRoleField initialized with val
func NewSettlPartyRole(val int) SettlPartyRoleField {
	return SettlPartyRoleField{quickfix.FIXInt(val)}
}

func (f SettlPartyRoleField) Value() int { return f.Int() }

//SettlPartySubIDField is a STRING field
type SettlPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlPartySubID (785)
func (f SettlPartySubIDField) Tag() quickfix.Tag { return tag.SettlPartySubID }

//NewSettlPartySubID returns a new SettlPartySubIDField initialized with val
func NewSettlPartySubID(val string) SettlPartySubIDField {
	return SettlPartySubIDField{quickfix.FIXString(val)}
}

func (f SettlPartySubIDField) Value() string { return f.String() }

//SettlPartySubIDTypeField is a INT field
type SettlPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.SettlPartySubIDType (786)
func (f SettlPartySubIDTypeField) Tag() quickfix.Tag { return tag.SettlPartySubIDType }

//NewSettlPartySubIDType returns a new SettlPartySubIDTypeField initialized with val
func NewSettlPartySubIDType(val int) SettlPartySubIDTypeField {
	return SettlPartySubIDTypeField{quickfix.FIXInt(val)}
}

func (f SettlPartySubIDTypeField) Value() int { return f.Int() }

//SettlPriceField is a PRICE field
type SettlPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.SettlPrice (730)
func (f SettlPriceField) Tag() quickfix.Tag { return tag.SettlPrice }

//NewSettlPrice returns a new SettlPriceField initialized with val and scale
func NewSettlPrice(val decimal.Decimal, scale int32) SettlPriceField {
	return SettlPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SettlPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//SettlPriceTypeField is a enum.SettlPriceType field
type SettlPriceTypeField struct{ quickfix.FIXString }

//Tag returns tag.SettlPriceType (731)
func (f SettlPriceTypeField) Tag() quickfix.Tag { return tag.SettlPriceType }

func NewSettlPriceType(val enum.SettlPriceType) SettlPriceTypeField {
	return SettlPriceTypeField{quickfix.FIXString(val)}
}

func (f SettlPriceTypeField) Value() enum.SettlPriceType { return enum.SettlPriceType(f.String()) }

//SettlSessIDField is a enum.SettlSessID field
type SettlSessIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlSessID (716)
func (f SettlSessIDField) Tag() quickfix.Tag { return tag.SettlSessID }

func NewSettlSessID(val enum.SettlSessID) SettlSessIDField {
	return SettlSessIDField{quickfix.FIXString(val)}
}

func (f SettlSessIDField) Value() enum.SettlSessID { return enum.SettlSessID(f.String()) }

//SettlSessSubIDField is a STRING field
type SettlSessSubIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlSessSubID (717)
func (f SettlSessSubIDField) Tag() quickfix.Tag { return tag.SettlSessSubID }

//NewSettlSessSubID returns a new SettlSessSubIDField initialized with val
func NewSettlSessSubID(val string) SettlSessSubIDField {
	return SettlSessSubIDField{quickfix.FIXString(val)}
}

func (f SettlSessSubIDField) Value() string { return f.String() }

//SettlTypeField is a enum.SettlType field
type SettlTypeField struct{ quickfix.FIXString }

//Tag returns tag.SettlType (63)
func (f SettlTypeField) Tag() quickfix.Tag { return tag.SettlType }

func NewSettlType(val enum.SettlType) SettlTypeField {
	return SettlTypeField{quickfix.FIXString(val)}
}

func (f SettlTypeField) Value() enum.SettlType { return enum.SettlType(f.String()) }

//SettleOnOpenFlagField is a STRING field
type SettleOnOpenFlagField struct{ quickfix.FIXString }

//Tag returns tag.SettleOnOpenFlag (966)
func (f SettleOnOpenFlagField) Tag() quickfix.Tag { return tag.SettleOnOpenFlag }

//NewSettleOnOpenFlag returns a new SettleOnOpenFlagField initialized with val
func NewSettleOnOpenFlag(val string) SettleOnOpenFlagField {
	return SettleOnOpenFlagField{quickfix.FIXString(val)}
}

func (f SettleOnOpenFlagField) Value() string { return f.String() }

//SettlementCycleNoField is a INT field
type SettlementCycleNoField struct{ quickfix.FIXInt }

//Tag returns tag.SettlementCycleNo (1153)
func (f SettlementCycleNoField) Tag() quickfix.Tag { return tag.SettlementCycleNo }

//NewSettlementCycleNo returns a new SettlementCycleNoField initialized with val
func NewSettlementCycleNo(val int) SettlementCycleNoField {
	return SettlementCycleNoField{quickfix.FIXInt(val)}
}

func (f SettlementCycleNoField) Value() int { return f.Int() }

//SettlmntTypField is a enum.SettlmntTyp field
type SettlmntTypField struct{ quickfix.FIXString }

//Tag returns tag.SettlmntTyp (63)
func (f SettlmntTypField) Tag() quickfix.Tag { return tag.SettlmntTyp }

func NewSettlmntTyp(val enum.SettlmntTyp) SettlmntTypField {
	return SettlmntTypField{quickfix.FIXString(val)}
}

func (f SettlmntTypField) Value() enum.SettlmntTyp { return enum.SettlmntTyp(f.String()) }

//SharedCommissionField is a AMT field
type SharedCommissionField struct{ quickfix.FIXDecimal }

//Tag returns tag.SharedCommission (858)
func (f SharedCommissionField) Tag() quickfix.Tag { return tag.SharedCommission }

//NewSharedCommission returns a new SharedCommissionField initialized with val and scale
func NewSharedCommission(val decimal.Decimal, scale int32) SharedCommissionField {
	return SharedCommissionField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SharedCommissionField) Value() (val decimal.Decimal) { return f.Decimal }

//SharesField is a QTY field
type SharesField struct{ quickfix.FIXDecimal }

//Tag returns tag.Shares (53)
func (f SharesField) Tag() quickfix.Tag { return tag.Shares }

//NewShares returns a new SharesField initialized with val and scale
func NewShares(val decimal.Decimal, scale int32) SharesField {
	return SharesField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SharesField) Value() (val decimal.Decimal) { return f.Decimal }

//ShortQtyField is a QTY field
type ShortQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.ShortQty (705)
func (f ShortQtyField) Tag() quickfix.Tag { return tag.ShortQty }

//NewShortQty returns a new ShortQtyField initialized with val and scale
func NewShortQty(val decimal.Decimal, scale int32) ShortQtyField {
	return ShortQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f ShortQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//ShortSaleReasonField is a enum.ShortSaleReason field
type ShortSaleReasonField struct{ quickfix.FIXString }

//Tag returns tag.ShortSaleReason (853)
func (f ShortSaleReasonField) Tag() quickfix.Tag { return tag.ShortSaleReason }

func NewShortSaleReason(val enum.ShortSaleReason) ShortSaleReasonField {
	return ShortSaleReasonField{quickfix.FIXString(val)}
}

func (f ShortSaleReasonField) Value() enum.ShortSaleReason { return enum.ShortSaleReason(f.String()) }

//SideField is a enum.Side field
type SideField struct{ quickfix.FIXString }

//Tag returns tag.Side (54)
func (f SideField) Tag() quickfix.Tag { return tag.Side }

func NewSide(val enum.Side) SideField {
	return SideField{quickfix.FIXString(val)}
}

func (f SideField) Value() enum.Side { return enum.Side(f.String()) }

//SideComplianceIDField is a STRING field
type SideComplianceIDField struct{ quickfix.FIXString }

//Tag returns tag.SideComplianceID (659)
func (f SideComplianceIDField) Tag() quickfix.Tag { return tag.SideComplianceID }

//NewSideComplianceID returns a new SideComplianceIDField initialized with val
func NewSideComplianceID(val string) SideComplianceIDField {
	return SideComplianceIDField{quickfix.FIXString(val)}
}

func (f SideComplianceIDField) Value() string { return f.String() }

//SideCurrencyField is a CURRENCY field
type SideCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.SideCurrency (1154)
func (f SideCurrencyField) Tag() quickfix.Tag { return tag.SideCurrency }

//NewSideCurrency returns a new SideCurrencyField initialized with val
func NewSideCurrency(val string) SideCurrencyField {
	return SideCurrencyField{quickfix.FIXString(val)}
}

func (f SideCurrencyField) Value() string { return f.String() }

//SideExecIDField is a STRING field
type SideExecIDField struct{ quickfix.FIXString }

//Tag returns tag.SideExecID (1427)
func (f SideExecIDField) Tag() quickfix.Tag { return tag.SideExecID }

//NewSideExecID returns a new SideExecIDField initialized with val
func NewSideExecID(val string) SideExecIDField {
	return SideExecIDField{quickfix.FIXString(val)}
}

func (f SideExecIDField) Value() string { return f.String() }

//SideFillStationCdField is a STRING field
type SideFillStationCdField struct{ quickfix.FIXString }

//Tag returns tag.SideFillStationCd (1006)
func (f SideFillStationCdField) Tag() quickfix.Tag { return tag.SideFillStationCd }

//NewSideFillStationCd returns a new SideFillStationCdField initialized with val
func NewSideFillStationCd(val string) SideFillStationCdField {
	return SideFillStationCdField{quickfix.FIXString(val)}
}

func (f SideFillStationCdField) Value() string { return f.String() }

//SideGrossTradeAmtField is a AMT field
type SideGrossTradeAmtField struct{ quickfix.FIXDecimal }

//Tag returns tag.SideGrossTradeAmt (1072)
func (f SideGrossTradeAmtField) Tag() quickfix.Tag { return tag.SideGrossTradeAmt }

//NewSideGrossTradeAmt returns a new SideGrossTradeAmtField initialized with val and scale
func NewSideGrossTradeAmt(val decimal.Decimal, scale int32) SideGrossTradeAmtField {
	return SideGrossTradeAmtField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SideGrossTradeAmtField) Value() (val decimal.Decimal) { return f.Decimal }

//SideLastQtyField is a INT field
type SideLastQtyField struct{ quickfix.FIXInt }

//Tag returns tag.SideLastQty (1009)
func (f SideLastQtyField) Tag() quickfix.Tag { return tag.SideLastQty }

//NewSideLastQty returns a new SideLastQtyField initialized with val
func NewSideLastQty(val int) SideLastQtyField {
	return SideLastQtyField{quickfix.FIXInt(val)}
}

func (f SideLastQtyField) Value() int { return f.Int() }

//SideLiquidityIndField is a INT field
type SideLiquidityIndField struct{ quickfix.FIXInt }

//Tag returns tag.SideLiquidityInd (1444)
func (f SideLiquidityIndField) Tag() quickfix.Tag { return tag.SideLiquidityInd }

//NewSideLiquidityInd returns a new SideLiquidityIndField initialized with val
func NewSideLiquidityInd(val int) SideLiquidityIndField {
	return SideLiquidityIndField{quickfix.FIXInt(val)}
}

func (f SideLiquidityIndField) Value() int { return f.Int() }

//SideMultiLegReportingTypeField is a enum.SideMultiLegReportingType field
type SideMultiLegReportingTypeField struct{ quickfix.FIXString }

//Tag returns tag.SideMultiLegReportingType (752)
func (f SideMultiLegReportingTypeField) Tag() quickfix.Tag { return tag.SideMultiLegReportingType }

func NewSideMultiLegReportingType(val enum.SideMultiLegReportingType) SideMultiLegReportingTypeField {
	return SideMultiLegReportingTypeField{quickfix.FIXString(val)}
}

func (f SideMultiLegReportingTypeField) Value() enum.SideMultiLegReportingType {
	return enum.SideMultiLegReportingType(f.String())
}

//SideQtyField is a INT field
type SideQtyField struct{ quickfix.FIXInt }

//Tag returns tag.SideQty (1009)
func (f SideQtyField) Tag() quickfix.Tag { return tag.SideQty }

//NewSideQty returns a new SideQtyField initialized with val
func NewSideQty(val int) SideQtyField {
	return SideQtyField{quickfix.FIXInt(val)}
}

func (f SideQtyField) Value() int { return f.Int() }

//SideReasonCdField is a STRING field
type SideReasonCdField struct{ quickfix.FIXString }

//Tag returns tag.SideReasonCd (1007)
func (f SideReasonCdField) Tag() quickfix.Tag { return tag.SideReasonCd }

//NewSideReasonCd returns a new SideReasonCdField initialized with val
func NewSideReasonCd(val string) SideReasonCdField {
	return SideReasonCdField{quickfix.FIXString(val)}
}

func (f SideReasonCdField) Value() string { return f.String() }

//SideSettlCurrencyField is a CURRENCY field
type SideSettlCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.SideSettlCurrency (1155)
func (f SideSettlCurrencyField) Tag() quickfix.Tag { return tag.SideSettlCurrency }

//NewSideSettlCurrency returns a new SideSettlCurrencyField initialized with val
func NewSideSettlCurrency(val string) SideSettlCurrencyField {
	return SideSettlCurrencyField{quickfix.FIXString(val)}
}

func (f SideSettlCurrencyField) Value() string { return f.String() }

//SideTimeInForceField is a UTCTIMESTAMP field
type SideTimeInForceField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.SideTimeInForce (962)
func (f SideTimeInForceField) Tag() quickfix.Tag { return tag.SideTimeInForce }

//NewSideTimeInForce returns a new SideTimeInForceField initialized with val
func NewSideTimeInForce(val time.Time) SideTimeInForceField {
	return SideTimeInForceField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewSideTimeInForceNoMillis returns a new SideTimeInForceField initialized with val without millisecs
func NewSideTimeInForceNoMillis(val time.Time) SideTimeInForceField {
	return SideTimeInForceField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f SideTimeInForceField) Value() time.Time { return f.Time }

//SideTradeReportIDField is a STRING field
type SideTradeReportIDField struct{ quickfix.FIXString }

//Tag returns tag.SideTradeReportID (1005)
func (f SideTradeReportIDField) Tag() quickfix.Tag { return tag.SideTradeReportID }

//NewSideTradeReportID returns a new SideTradeReportIDField initialized with val
func NewSideTradeReportID(val string) SideTradeReportIDField {
	return SideTradeReportIDField{quickfix.FIXString(val)}
}

func (f SideTradeReportIDField) Value() string { return f.String() }

//SideTrdRegTimestampField is a UTCTIMESTAMP field
type SideTrdRegTimestampField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.SideTrdRegTimestamp (1012)
func (f SideTrdRegTimestampField) Tag() quickfix.Tag { return tag.SideTrdRegTimestamp }

//NewSideTrdRegTimestamp returns a new SideTrdRegTimestampField initialized with val
func NewSideTrdRegTimestamp(val time.Time) SideTrdRegTimestampField {
	return SideTrdRegTimestampField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewSideTrdRegTimestampNoMillis returns a new SideTrdRegTimestampField initialized with val without millisecs
func NewSideTrdRegTimestampNoMillis(val time.Time) SideTrdRegTimestampField {
	return SideTrdRegTimestampField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f SideTrdRegTimestampField) Value() time.Time { return f.Time }

//SideTrdRegTimestampSrcField is a STRING field
type SideTrdRegTimestampSrcField struct{ quickfix.FIXString }

//Tag returns tag.SideTrdRegTimestampSrc (1014)
func (f SideTrdRegTimestampSrcField) Tag() quickfix.Tag { return tag.SideTrdRegTimestampSrc }

//NewSideTrdRegTimestampSrc returns a new SideTrdRegTimestampSrcField initialized with val
func NewSideTrdRegTimestampSrc(val string) SideTrdRegTimestampSrcField {
	return SideTrdRegTimestampSrcField{quickfix.FIXString(val)}
}

func (f SideTrdRegTimestampSrcField) Value() string { return f.String() }

//SideTrdRegTimestampTypeField is a INT field
type SideTrdRegTimestampTypeField struct{ quickfix.FIXInt }

//Tag returns tag.SideTrdRegTimestampType (1013)
func (f SideTrdRegTimestampTypeField) Tag() quickfix.Tag { return tag.SideTrdRegTimestampType }

//NewSideTrdRegTimestampType returns a new SideTrdRegTimestampTypeField initialized with val
func NewSideTrdRegTimestampType(val int) SideTrdRegTimestampTypeField {
	return SideTrdRegTimestampTypeField{quickfix.FIXInt(val)}
}

func (f SideTrdRegTimestampTypeField) Value() int { return f.Int() }

//SideTrdSubTypField is a enum.SideTrdSubTyp field
type SideTrdSubTypField struct{ quickfix.FIXString }

//Tag returns tag.SideTrdSubTyp (1008)
func (f SideTrdSubTypField) Tag() quickfix.Tag { return tag.SideTrdSubTyp }

func NewSideTrdSubTyp(val enum.SideTrdSubTyp) SideTrdSubTypField {
	return SideTrdSubTypField{quickfix.FIXString(val)}
}

func (f SideTrdSubTypField) Value() enum.SideTrdSubTyp { return enum.SideTrdSubTyp(f.String()) }

//SideValue1Field is a AMT field
type SideValue1Field struct{ quickfix.FIXDecimal }

//Tag returns tag.SideValue1 (396)
func (f SideValue1Field) Tag() quickfix.Tag { return tag.SideValue1 }

//NewSideValue1 returns a new SideValue1Field initialized with val and scale
func NewSideValue1(val decimal.Decimal, scale int32) SideValue1Field {
	return SideValue1Field{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SideValue1Field) Value() (val decimal.Decimal) { return f.Decimal }

//SideValue2Field is a AMT field
type SideValue2Field struct{ quickfix.FIXDecimal }

//Tag returns tag.SideValue2 (397)
func (f SideValue2Field) Tag() quickfix.Tag { return tag.SideValue2 }

//NewSideValue2 returns a new SideValue2Field initialized with val and scale
func NewSideValue2(val decimal.Decimal, scale int32) SideValue2Field {
	return SideValue2Field{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SideValue2Field) Value() (val decimal.Decimal) { return f.Decimal }

//SideValueIndField is a enum.SideValueInd field
type SideValueIndField struct{ quickfix.FIXString }

//Tag returns tag.SideValueInd (401)
func (f SideValueIndField) Tag() quickfix.Tag { return tag.SideValueInd }

func NewSideValueInd(val enum.SideValueInd) SideValueIndField {
	return SideValueIndField{quickfix.FIXString(val)}
}

func (f SideValueIndField) Value() enum.SideValueInd { return enum.SideValueInd(f.String()) }

//SignatureField is a DATA field
type SignatureField struct{ quickfix.FIXString }

//Tag returns tag.Signature (89)
func (f SignatureField) Tag() quickfix.Tag { return tag.Signature }

//NewSignature returns a new SignatureField initialized with val
func NewSignature(val string) SignatureField {
	return SignatureField{quickfix.FIXString(val)}
}

func (f SignatureField) Value() string { return f.String() }

//SignatureLengthField is a LENGTH field
type SignatureLengthField struct{ quickfix.FIXInt }

//Tag returns tag.SignatureLength (93)
func (f SignatureLengthField) Tag() quickfix.Tag { return tag.SignatureLength }

//NewSignatureLength returns a new SignatureLengthField initialized with val
func NewSignatureLength(val int) SignatureLengthField {
	return SignatureLengthField{quickfix.FIXInt(val)}
}

func (f SignatureLengthField) Value() int { return f.Int() }

//SolicitedFlagField is a BOOLEAN field
type SolicitedFlagField struct{ quickfix.FIXBoolean }

//Tag returns tag.SolicitedFlag (377)
func (f SolicitedFlagField) Tag() quickfix.Tag { return tag.SolicitedFlag }

//NewSolicitedFlag returns a new SolicitedFlagField initialized with val
func NewSolicitedFlag(val bool) SolicitedFlagField {
	return SolicitedFlagField{quickfix.FIXBoolean(val)}
}

func (f SolicitedFlagField) Value() bool { return f.Bool() }

//SpreadField is a PRICEOFFSET field
type SpreadField struct{ quickfix.FIXDecimal }

//Tag returns tag.Spread (218)
func (f SpreadField) Tag() quickfix.Tag { return tag.Spread }

//NewSpread returns a new SpreadField initialized with val and scale
func NewSpread(val decimal.Decimal, scale int32) SpreadField {
	return SpreadField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SpreadField) Value() (val decimal.Decimal) { return f.Decimal }

//SpreadToBenchmarkField is a PRICEOFFSET field
type SpreadToBenchmarkField struct{ quickfix.FIXDecimal }

//Tag returns tag.SpreadToBenchmark (218)
func (f SpreadToBenchmarkField) Tag() quickfix.Tag { return tag.SpreadToBenchmark }

//NewSpreadToBenchmark returns a new SpreadToBenchmarkField initialized with val and scale
func NewSpreadToBenchmark(val decimal.Decimal, scale int32) SpreadToBenchmarkField {
	return SpreadToBenchmarkField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SpreadToBenchmarkField) Value() (val decimal.Decimal) { return f.Decimal }

//StandInstDbIDField is a STRING field
type StandInstDbIDField struct{ quickfix.FIXString }

//Tag returns tag.StandInstDbID (171)
func (f StandInstDbIDField) Tag() quickfix.Tag { return tag.StandInstDbID }

//NewStandInstDbID returns a new StandInstDbIDField initialized with val
func NewStandInstDbID(val string) StandInstDbIDField {
	return StandInstDbIDField{quickfix.FIXString(val)}
}

func (f StandInstDbIDField) Value() string { return f.String() }

//StandInstDbNameField is a STRING field
type StandInstDbNameField struct{ quickfix.FIXString }

//Tag returns tag.StandInstDbName (170)
func (f StandInstDbNameField) Tag() quickfix.Tag { return tag.StandInstDbName }

//NewStandInstDbName returns a new StandInstDbNameField initialized with val
func NewStandInstDbName(val string) StandInstDbNameField {
	return StandInstDbNameField{quickfix.FIXString(val)}
}

func (f StandInstDbNameField) Value() string { return f.String() }

//StandInstDbTypeField is a enum.StandInstDbType field
type StandInstDbTypeField struct{ quickfix.FIXString }

//Tag returns tag.StandInstDbType (169)
func (f StandInstDbTypeField) Tag() quickfix.Tag { return tag.StandInstDbType }

func NewStandInstDbType(val enum.StandInstDbType) StandInstDbTypeField {
	return StandInstDbTypeField{quickfix.FIXString(val)}
}

func (f StandInstDbTypeField) Value() enum.StandInstDbType { return enum.StandInstDbType(f.String()) }

//StartCashField is a AMT field
type StartCashField struct{ quickfix.FIXDecimal }

//Tag returns tag.StartCash (921)
func (f StartCashField) Tag() quickfix.Tag { return tag.StartCash }

//NewStartCash returns a new StartCashField initialized with val and scale
func NewStartCash(val decimal.Decimal, scale int32) StartCashField {
	return StartCashField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f StartCashField) Value() (val decimal.Decimal) { return f.Decimal }

//StartDateField is a LOCALMKTDATE field
type StartDateField struct{ quickfix.FIXString }

//Tag returns tag.StartDate (916)
func (f StartDateField) Tag() quickfix.Tag { return tag.StartDate }

//NewStartDate returns a new StartDateField initialized with val
func NewStartDate(val string) StartDateField {
	return StartDateField{quickfix.FIXString(val)}
}

func (f StartDateField) Value() string { return f.String() }

//StartMaturityMonthYearField is a MONTHYEAR field
type StartMaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.StartMaturityMonthYear (1241)
func (f StartMaturityMonthYearField) Tag() quickfix.Tag { return tag.StartMaturityMonthYear }

//NewStartMaturityMonthYear returns a new StartMaturityMonthYearField initialized with val
func NewStartMaturityMonthYear(val string) StartMaturityMonthYearField {
	return StartMaturityMonthYearField{quickfix.FIXString(val)}
}

func (f StartMaturityMonthYearField) Value() string { return f.String() }

//StartStrikePxRangeField is a PRICE field
type StartStrikePxRangeField struct{ quickfix.FIXDecimal }

//Tag returns tag.StartStrikePxRange (1202)
func (f StartStrikePxRangeField) Tag() quickfix.Tag { return tag.StartStrikePxRange }

//NewStartStrikePxRange returns a new StartStrikePxRangeField initialized with val and scale
func NewStartStrikePxRange(val decimal.Decimal, scale int32) StartStrikePxRangeField {
	return StartStrikePxRangeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f StartStrikePxRangeField) Value() (val decimal.Decimal) { return f.Decimal }

//StartTickPriceRangeField is a PRICE field
type StartTickPriceRangeField struct{ quickfix.FIXDecimal }

//Tag returns tag.StartTickPriceRange (1206)
func (f StartTickPriceRangeField) Tag() quickfix.Tag { return tag.StartTickPriceRange }

//NewStartTickPriceRange returns a new StartTickPriceRangeField initialized with val and scale
func NewStartTickPriceRange(val decimal.Decimal, scale int32) StartTickPriceRangeField {
	return StartTickPriceRangeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f StartTickPriceRangeField) Value() (val decimal.Decimal) { return f.Decimal }

//StateOrProvinceOfIssueField is a STRING field
type StateOrProvinceOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.StateOrProvinceOfIssue (471)
func (f StateOrProvinceOfIssueField) Tag() quickfix.Tag { return tag.StateOrProvinceOfIssue }

//NewStateOrProvinceOfIssue returns a new StateOrProvinceOfIssueField initialized with val
func NewStateOrProvinceOfIssue(val string) StateOrProvinceOfIssueField {
	return StateOrProvinceOfIssueField{quickfix.FIXString(val)}
}

func (f StateOrProvinceOfIssueField) Value() string { return f.String() }

//StatsTypeField is a enum.StatsType field
type StatsTypeField struct{ quickfix.FIXString }

//Tag returns tag.StatsType (1176)
func (f StatsTypeField) Tag() quickfix.Tag { return tag.StatsType }

func NewStatsType(val enum.StatsType) StatsTypeField {
	return StatsTypeField{quickfix.FIXString(val)}
}

func (f StatsTypeField) Value() enum.StatsType { return enum.StatsType(f.String()) }

//StatusTextField is a STRING field
type StatusTextField struct{ quickfix.FIXString }

//Tag returns tag.StatusText (929)
func (f StatusTextField) Tag() quickfix.Tag { return tag.StatusText }

//NewStatusText returns a new StatusTextField initialized with val
func NewStatusText(val string) StatusTextField {
	return StatusTextField{quickfix.FIXString(val)}
}

func (f StatusTextField) Value() string { return f.String() }

//StatusValueField is a enum.StatusValue field
type StatusValueField struct{ quickfix.FIXString }

//Tag returns tag.StatusValue (928)
func (f StatusValueField) Tag() quickfix.Tag { return tag.StatusValue }

func NewStatusValue(val enum.StatusValue) StatusValueField {
	return StatusValueField{quickfix.FIXString(val)}
}

func (f StatusValueField) Value() enum.StatusValue { return enum.StatusValue(f.String()) }

//StipulationTypeField is a enum.StipulationType field
type StipulationTypeField struct{ quickfix.FIXString }

//Tag returns tag.StipulationType (233)
func (f StipulationTypeField) Tag() quickfix.Tag { return tag.StipulationType }

func NewStipulationType(val enum.StipulationType) StipulationTypeField {
	return StipulationTypeField{quickfix.FIXString(val)}
}

func (f StipulationTypeField) Value() enum.StipulationType { return enum.StipulationType(f.String()) }

//StipulationValueField is a STRING field
type StipulationValueField struct{ quickfix.FIXString }

//Tag returns tag.StipulationValue (234)
func (f StipulationValueField) Tag() quickfix.Tag { return tag.StipulationValue }

//NewStipulationValue returns a new StipulationValueField initialized with val
func NewStipulationValue(val string) StipulationValueField {
	return StipulationValueField{quickfix.FIXString(val)}
}

func (f StipulationValueField) Value() string { return f.String() }

//StopPxField is a PRICE field
type StopPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.StopPx (99)
func (f StopPxField) Tag() quickfix.Tag { return tag.StopPx }

//NewStopPx returns a new StopPxField initialized with val and scale
func NewStopPx(val decimal.Decimal, scale int32) StopPxField {
	return StopPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f StopPxField) Value() (val decimal.Decimal) { return f.Decimal }

//StrategyParameterNameField is a STRING field
type StrategyParameterNameField struct{ quickfix.FIXString }

//Tag returns tag.StrategyParameterName (958)
func (f StrategyParameterNameField) Tag() quickfix.Tag { return tag.StrategyParameterName }

//NewStrategyParameterName returns a new StrategyParameterNameField initialized with val
func NewStrategyParameterName(val string) StrategyParameterNameField {
	return StrategyParameterNameField{quickfix.FIXString(val)}
}

func (f StrategyParameterNameField) Value() string { return f.String() }

//StrategyParameterTypeField is a enum.StrategyParameterType field
type StrategyParameterTypeField struct{ quickfix.FIXString }

//Tag returns tag.StrategyParameterType (959)
func (f StrategyParameterTypeField) Tag() quickfix.Tag { return tag.StrategyParameterType }

func NewStrategyParameterType(val enum.StrategyParameterType) StrategyParameterTypeField {
	return StrategyParameterTypeField{quickfix.FIXString(val)}
}

func (f StrategyParameterTypeField) Value() enum.StrategyParameterType {
	return enum.StrategyParameterType(f.String())
}

//StrategyParameterValueField is a STRING field
type StrategyParameterValueField struct{ quickfix.FIXString }

//Tag returns tag.StrategyParameterValue (960)
func (f StrategyParameterValueField) Tag() quickfix.Tag { return tag.StrategyParameterValue }

//NewStrategyParameterValue returns a new StrategyParameterValueField initialized with val
func NewStrategyParameterValue(val string) StrategyParameterValueField {
	return StrategyParameterValueField{quickfix.FIXString(val)}
}

func (f StrategyParameterValueField) Value() string { return f.String() }

//StreamAsgnAckTypeField is a enum.StreamAsgnAckType field
type StreamAsgnAckTypeField struct{ quickfix.FIXString }

//Tag returns tag.StreamAsgnAckType (1503)
func (f StreamAsgnAckTypeField) Tag() quickfix.Tag { return tag.StreamAsgnAckType }

func NewStreamAsgnAckType(val enum.StreamAsgnAckType) StreamAsgnAckTypeField {
	return StreamAsgnAckTypeField{quickfix.FIXString(val)}
}

func (f StreamAsgnAckTypeField) Value() enum.StreamAsgnAckType {
	return enum.StreamAsgnAckType(f.String())
}

//StreamAsgnRejReasonField is a enum.StreamAsgnRejReason field
type StreamAsgnRejReasonField struct{ quickfix.FIXString }

//Tag returns tag.StreamAsgnRejReason (1502)
func (f StreamAsgnRejReasonField) Tag() quickfix.Tag { return tag.StreamAsgnRejReason }

func NewStreamAsgnRejReason(val enum.StreamAsgnRejReason) StreamAsgnRejReasonField {
	return StreamAsgnRejReasonField{quickfix.FIXString(val)}
}

func (f StreamAsgnRejReasonField) Value() enum.StreamAsgnRejReason {
	return enum.StreamAsgnRejReason(f.String())
}

//StreamAsgnReqIDField is a STRING field
type StreamAsgnReqIDField struct{ quickfix.FIXString }

//Tag returns tag.StreamAsgnReqID (1497)
func (f StreamAsgnReqIDField) Tag() quickfix.Tag { return tag.StreamAsgnReqID }

//NewStreamAsgnReqID returns a new StreamAsgnReqIDField initialized with val
func NewStreamAsgnReqID(val string) StreamAsgnReqIDField {
	return StreamAsgnReqIDField{quickfix.FIXString(val)}
}

func (f StreamAsgnReqIDField) Value() string { return f.String() }

//StreamAsgnReqTypeField is a enum.StreamAsgnReqType field
type StreamAsgnReqTypeField struct{ quickfix.FIXString }

//Tag returns tag.StreamAsgnReqType (1498)
func (f StreamAsgnReqTypeField) Tag() quickfix.Tag { return tag.StreamAsgnReqType }

func NewStreamAsgnReqType(val enum.StreamAsgnReqType) StreamAsgnReqTypeField {
	return StreamAsgnReqTypeField{quickfix.FIXString(val)}
}

func (f StreamAsgnReqTypeField) Value() enum.StreamAsgnReqType {
	return enum.StreamAsgnReqType(f.String())
}

//StreamAsgnRptIDField is a STRING field
type StreamAsgnRptIDField struct{ quickfix.FIXString }

//Tag returns tag.StreamAsgnRptID (1501)
func (f StreamAsgnRptIDField) Tag() quickfix.Tag { return tag.StreamAsgnRptID }

//NewStreamAsgnRptID returns a new StreamAsgnRptIDField initialized with val
func NewStreamAsgnRptID(val string) StreamAsgnRptIDField {
	return StreamAsgnRptIDField{quickfix.FIXString(val)}
}

func (f StreamAsgnRptIDField) Value() string { return f.String() }

//StreamAsgnTypeField is a enum.StreamAsgnType field
type StreamAsgnTypeField struct{ quickfix.FIXString }

//Tag returns tag.StreamAsgnType (1617)
func (f StreamAsgnTypeField) Tag() quickfix.Tag { return tag.StreamAsgnType }

func NewStreamAsgnType(val enum.StreamAsgnType) StreamAsgnTypeField {
	return StreamAsgnTypeField{quickfix.FIXString(val)}
}

func (f StreamAsgnTypeField) Value() enum.StreamAsgnType { return enum.StreamAsgnType(f.String()) }

//StrikeCurrencyField is a CURRENCY field
type StrikeCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.StrikeCurrency (947)
func (f StrikeCurrencyField) Tag() quickfix.Tag { return tag.StrikeCurrency }

//NewStrikeCurrency returns a new StrikeCurrencyField initialized with val
func NewStrikeCurrency(val string) StrikeCurrencyField {
	return StrikeCurrencyField{quickfix.FIXString(val)}
}

func (f StrikeCurrencyField) Value() string { return f.String() }

//StrikeExerciseStyleField is a INT field
type StrikeExerciseStyleField struct{ quickfix.FIXInt }

//Tag returns tag.StrikeExerciseStyle (1304)
func (f StrikeExerciseStyleField) Tag() quickfix.Tag { return tag.StrikeExerciseStyle }

//NewStrikeExerciseStyle returns a new StrikeExerciseStyleField initialized with val
func NewStrikeExerciseStyle(val int) StrikeExerciseStyleField {
	return StrikeExerciseStyleField{quickfix.FIXInt(val)}
}

func (f StrikeExerciseStyleField) Value() int { return f.Int() }

//StrikeIncrementField is a FLOAT field
type StrikeIncrementField struct{ quickfix.FIXDecimal }

//Tag returns tag.StrikeIncrement (1204)
func (f StrikeIncrementField) Tag() quickfix.Tag { return tag.StrikeIncrement }

//NewStrikeIncrement returns a new StrikeIncrementField initialized with val and scale
func NewStrikeIncrement(val decimal.Decimal, scale int32) StrikeIncrementField {
	return StrikeIncrementField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f StrikeIncrementField) Value() (val decimal.Decimal) { return f.Decimal }

//StrikeMultiplierField is a FLOAT field
type StrikeMultiplierField struct{ quickfix.FIXDecimal }

//Tag returns tag.StrikeMultiplier (967)
func (f StrikeMultiplierField) Tag() quickfix.Tag { return tag.StrikeMultiplier }

//NewStrikeMultiplier returns a new StrikeMultiplierField initialized with val and scale
func NewStrikeMultiplier(val decimal.Decimal, scale int32) StrikeMultiplierField {
	return StrikeMultiplierField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f StrikeMultiplierField) Value() (val decimal.Decimal) { return f.Decimal }

//StrikePriceField is a PRICE field
type StrikePriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.StrikePrice (202)
func (f StrikePriceField) Tag() quickfix.Tag { return tag.StrikePrice }

//NewStrikePrice returns a new StrikePriceField initialized with val and scale
func NewStrikePrice(val decimal.Decimal, scale int32) StrikePriceField {
	return StrikePriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f StrikePriceField) Value() (val decimal.Decimal) { return f.Decimal }

//StrikePriceBoundaryMethodField is a enum.StrikePriceBoundaryMethod field
type StrikePriceBoundaryMethodField struct{ quickfix.FIXString }

//Tag returns tag.StrikePriceBoundaryMethod (1479)
func (f StrikePriceBoundaryMethodField) Tag() quickfix.Tag { return tag.StrikePriceBoundaryMethod }

func NewStrikePriceBoundaryMethod(val enum.StrikePriceBoundaryMethod) StrikePriceBoundaryMethodField {
	return StrikePriceBoundaryMethodField{quickfix.FIXString(val)}
}

func (f StrikePriceBoundaryMethodField) Value() enum.StrikePriceBoundaryMethod {
	return enum.StrikePriceBoundaryMethod(f.String())
}

//StrikePriceBoundaryPrecisionField is a PERCENTAGE field
type StrikePriceBoundaryPrecisionField struct{ quickfix.FIXDecimal }

//Tag returns tag.StrikePriceBoundaryPrecision (1480)
func (f StrikePriceBoundaryPrecisionField) Tag() quickfix.Tag { return tag.StrikePriceBoundaryPrecision }

//NewStrikePriceBoundaryPrecision returns a new StrikePriceBoundaryPrecisionField initialized with val and scale
func NewStrikePriceBoundaryPrecision(val decimal.Decimal, scale int32) StrikePriceBoundaryPrecisionField {
	return StrikePriceBoundaryPrecisionField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f StrikePriceBoundaryPrecisionField) Value() (val decimal.Decimal) { return f.Decimal }

//StrikePriceDeterminationMethodField is a enum.StrikePriceDeterminationMethod field
type StrikePriceDeterminationMethodField struct{ quickfix.FIXString }

//Tag returns tag.StrikePriceDeterminationMethod (1478)
func (f StrikePriceDeterminationMethodField) Tag() quickfix.Tag {
	return tag.StrikePriceDeterminationMethod
}

func NewStrikePriceDeterminationMethod(val enum.StrikePriceDeterminationMethod) StrikePriceDeterminationMethodField {
	return StrikePriceDeterminationMethodField{quickfix.FIXString(val)}
}

func (f StrikePriceDeterminationMethodField) Value() enum.StrikePriceDeterminationMethod {
	return enum.StrikePriceDeterminationMethod(f.String())
}

//StrikeRuleIDField is a STRING field
type StrikeRuleIDField struct{ quickfix.FIXString }

//Tag returns tag.StrikeRuleID (1223)
func (f StrikeRuleIDField) Tag() quickfix.Tag { return tag.StrikeRuleID }

//NewStrikeRuleID returns a new StrikeRuleIDField initialized with val
func NewStrikeRuleID(val string) StrikeRuleIDField {
	return StrikeRuleIDField{quickfix.FIXString(val)}
}

func (f StrikeRuleIDField) Value() string { return f.String() }

//StrikeTimeField is a UTCTIMESTAMP field
type StrikeTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.StrikeTime (443)
func (f StrikeTimeField) Tag() quickfix.Tag { return tag.StrikeTime }

//NewStrikeTime returns a new StrikeTimeField initialized with val
func NewStrikeTime(val time.Time) StrikeTimeField {
	return StrikeTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewStrikeTimeNoMillis returns a new StrikeTimeField initialized with val without millisecs
func NewStrikeTimeNoMillis(val time.Time) StrikeTimeField {
	return StrikeTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f StrikeTimeField) Value() time.Time { return f.Time }

//StrikeValueField is a FLOAT field
type StrikeValueField struct{ quickfix.FIXDecimal }

//Tag returns tag.StrikeValue (968)
func (f StrikeValueField) Tag() quickfix.Tag { return tag.StrikeValue }

//NewStrikeValue returns a new StrikeValueField initialized with val and scale
func NewStrikeValue(val decimal.Decimal, scale int32) StrikeValueField {
	return StrikeValueField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f StrikeValueField) Value() (val decimal.Decimal) { return f.Decimal }

//SubjectField is a STRING field
type SubjectField struct{ quickfix.FIXString }

//Tag returns tag.Subject (147)
func (f SubjectField) Tag() quickfix.Tag { return tag.Subject }

//NewSubject returns a new SubjectField initialized with val
func NewSubject(val string) SubjectField {
	return SubjectField{quickfix.FIXString(val)}
}

func (f SubjectField) Value() string { return f.String() }

//SubscriptionRequestTypeField is a enum.SubscriptionRequestType field
type SubscriptionRequestTypeField struct{ quickfix.FIXString }

//Tag returns tag.SubscriptionRequestType (263)
func (f SubscriptionRequestTypeField) Tag() quickfix.Tag { return tag.SubscriptionRequestType }

func NewSubscriptionRequestType(val enum.SubscriptionRequestType) SubscriptionRequestTypeField {
	return SubscriptionRequestTypeField{quickfix.FIXString(val)}
}

func (f SubscriptionRequestTypeField) Value() enum.SubscriptionRequestType {
	return enum.SubscriptionRequestType(f.String())
}

//SwapPointsField is a PRICEOFFSET field
type SwapPointsField struct{ quickfix.FIXDecimal }

//Tag returns tag.SwapPoints (1069)
func (f SwapPointsField) Tag() quickfix.Tag { return tag.SwapPoints }

//NewSwapPoints returns a new SwapPointsField initialized with val and scale
func NewSwapPoints(val decimal.Decimal, scale int32) SwapPointsField {
	return SwapPointsField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f SwapPointsField) Value() (val decimal.Decimal) { return f.Decimal }

//SymbolField is a STRING field
type SymbolField struct{ quickfix.FIXString }

//Tag returns tag.Symbol (55)
func (f SymbolField) Tag() quickfix.Tag { return tag.Symbol }

//NewSymbol returns a new SymbolField initialized with val
func NewSymbol(val string) SymbolField {
	return SymbolField{quickfix.FIXString(val)}
}

func (f SymbolField) Value() string { return f.String() }

//SymbolSfxField is a enum.SymbolSfx field
type SymbolSfxField struct{ quickfix.FIXString }

//Tag returns tag.SymbolSfx (65)
func (f SymbolSfxField) Tag() quickfix.Tag { return tag.SymbolSfx }

func NewSymbolSfx(val enum.SymbolSfx) SymbolSfxField {
	return SymbolSfxField{quickfix.FIXString(val)}
}

func (f SymbolSfxField) Value() enum.SymbolSfx { return enum.SymbolSfx(f.String()) }

//TZTransactTimeField is a TZTIMESTAMP field
type TZTransactTimeField struct{ quickfix.FIXString }

//Tag returns tag.TZTransactTime (1132)
func (f TZTransactTimeField) Tag() quickfix.Tag { return tag.TZTransactTime }

//NewTZTransactTime returns a new TZTransactTimeField initialized with val
func NewTZTransactTime(val string) TZTransactTimeField {
	return TZTransactTimeField{quickfix.FIXString(val)}
}

func (f TZTransactTimeField) Value() string { return f.String() }

//TargetCompIDField is a STRING field
type TargetCompIDField struct{ quickfix.FIXString }

//Tag returns tag.TargetCompID (56)
func (f TargetCompIDField) Tag() quickfix.Tag { return tag.TargetCompID }

//NewTargetCompID returns a new TargetCompIDField initialized with val
func NewTargetCompID(val string) TargetCompIDField {
	return TargetCompIDField{quickfix.FIXString(val)}
}

func (f TargetCompIDField) Value() string { return f.String() }

//TargetLocationIDField is a STRING field
type TargetLocationIDField struct{ quickfix.FIXString }

//Tag returns tag.TargetLocationID (143)
func (f TargetLocationIDField) Tag() quickfix.Tag { return tag.TargetLocationID }

//NewTargetLocationID returns a new TargetLocationIDField initialized with val
func NewTargetLocationID(val string) TargetLocationIDField {
	return TargetLocationIDField{quickfix.FIXString(val)}
}

func (f TargetLocationIDField) Value() string { return f.String() }

//TargetPartyIDField is a STRING field
type TargetPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.TargetPartyID (1462)
func (f TargetPartyIDField) Tag() quickfix.Tag { return tag.TargetPartyID }

//NewTargetPartyID returns a new TargetPartyIDField initialized with val
func NewTargetPartyID(val string) TargetPartyIDField {
	return TargetPartyIDField{quickfix.FIXString(val)}
}

func (f TargetPartyIDField) Value() string { return f.String() }

//TargetPartyIDSourceField is a CHAR field
type TargetPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.TargetPartyIDSource (1463)
func (f TargetPartyIDSourceField) Tag() quickfix.Tag { return tag.TargetPartyIDSource }

//NewTargetPartyIDSource returns a new TargetPartyIDSourceField initialized with val
func NewTargetPartyIDSource(val string) TargetPartyIDSourceField {
	return TargetPartyIDSourceField{quickfix.FIXString(val)}
}

func (f TargetPartyIDSourceField) Value() string { return f.String() }

//TargetPartyRoleField is a INT field
type TargetPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.TargetPartyRole (1464)
func (f TargetPartyRoleField) Tag() quickfix.Tag { return tag.TargetPartyRole }

//NewTargetPartyRole returns a new TargetPartyRoleField initialized with val
func NewTargetPartyRole(val int) TargetPartyRoleField {
	return TargetPartyRoleField{quickfix.FIXInt(val)}
}

func (f TargetPartyRoleField) Value() int { return f.Int() }

//TargetStrategyField is a enum.TargetStrategy field
type TargetStrategyField struct{ quickfix.FIXString }

//Tag returns tag.TargetStrategy (847)
func (f TargetStrategyField) Tag() quickfix.Tag { return tag.TargetStrategy }

func NewTargetStrategy(val enum.TargetStrategy) TargetStrategyField {
	return TargetStrategyField{quickfix.FIXString(val)}
}

func (f TargetStrategyField) Value() enum.TargetStrategy { return enum.TargetStrategy(f.String()) }

//TargetStrategyParametersField is a STRING field
type TargetStrategyParametersField struct{ quickfix.FIXString }

//Tag returns tag.TargetStrategyParameters (848)
func (f TargetStrategyParametersField) Tag() quickfix.Tag { return tag.TargetStrategyParameters }

//NewTargetStrategyParameters returns a new TargetStrategyParametersField initialized with val
func NewTargetStrategyParameters(val string) TargetStrategyParametersField {
	return TargetStrategyParametersField{quickfix.FIXString(val)}
}

func (f TargetStrategyParametersField) Value() string { return f.String() }

//TargetStrategyPerformanceField is a FLOAT field
type TargetStrategyPerformanceField struct{ quickfix.FIXDecimal }

//Tag returns tag.TargetStrategyPerformance (850)
func (f TargetStrategyPerformanceField) Tag() quickfix.Tag { return tag.TargetStrategyPerformance }

//NewTargetStrategyPerformance returns a new TargetStrategyPerformanceField initialized with val and scale
func NewTargetStrategyPerformance(val decimal.Decimal, scale int32) TargetStrategyPerformanceField {
	return TargetStrategyPerformanceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f TargetStrategyPerformanceField) Value() (val decimal.Decimal) { return f.Decimal }

//TargetSubIDField is a STRING field
type TargetSubIDField struct{ quickfix.FIXString }

//Tag returns tag.TargetSubID (57)
func (f TargetSubIDField) Tag() quickfix.Tag { return tag.TargetSubID }

//NewTargetSubID returns a new TargetSubIDField initialized with val
func NewTargetSubID(val string) TargetSubIDField {
	return TargetSubIDField{quickfix.FIXString(val)}
}

func (f TargetSubIDField) Value() string { return f.String() }

//TaxAdvantageTypeField is a enum.TaxAdvantageType field
type TaxAdvantageTypeField struct{ quickfix.FIXString }

//Tag returns tag.TaxAdvantageType (495)
func (f TaxAdvantageTypeField) Tag() quickfix.Tag { return tag.TaxAdvantageType }

func NewTaxAdvantageType(val enum.TaxAdvantageType) TaxAdvantageTypeField {
	return TaxAdvantageTypeField{quickfix.FIXString(val)}
}

func (f TaxAdvantageTypeField) Value() enum.TaxAdvantageType { return enum.TaxAdvantageType(f.String()) }

//TerminationTypeField is a enum.TerminationType field
type TerminationTypeField struct{ quickfix.FIXString }

//Tag returns tag.TerminationType (788)
func (f TerminationTypeField) Tag() quickfix.Tag { return tag.TerminationType }

func NewTerminationType(val enum.TerminationType) TerminationTypeField {
	return TerminationTypeField{quickfix.FIXString(val)}
}

func (f TerminationTypeField) Value() enum.TerminationType { return enum.TerminationType(f.String()) }

//TestMessageIndicatorField is a BOOLEAN field
type TestMessageIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.TestMessageIndicator (464)
func (f TestMessageIndicatorField) Tag() quickfix.Tag { return tag.TestMessageIndicator }

//NewTestMessageIndicator returns a new TestMessageIndicatorField initialized with val
func NewTestMessageIndicator(val bool) TestMessageIndicatorField {
	return TestMessageIndicatorField{quickfix.FIXBoolean(val)}
}

func (f TestMessageIndicatorField) Value() bool { return f.Bool() }

//TestReqIDField is a STRING field
type TestReqIDField struct{ quickfix.FIXString }

//Tag returns tag.TestReqID (112)
func (f TestReqIDField) Tag() quickfix.Tag { return tag.TestReqID }

//NewTestReqID returns a new TestReqIDField initialized with val
func NewTestReqID(val string) TestReqIDField {
	return TestReqIDField{quickfix.FIXString(val)}
}

func (f TestReqIDField) Value() string { return f.String() }

//TextField is a STRING field
type TextField struct{ quickfix.FIXString }

//Tag returns tag.Text (58)
func (f TextField) Tag() quickfix.Tag { return tag.Text }

//NewText returns a new TextField initialized with val
func NewText(val string) TextField {
	return TextField{quickfix.FIXString(val)}
}

func (f TextField) Value() string { return f.String() }

//ThresholdAmountField is a PRICEOFFSET field
type ThresholdAmountField struct{ quickfix.FIXDecimal }

//Tag returns tag.ThresholdAmount (834)
func (f ThresholdAmountField) Tag() quickfix.Tag { return tag.ThresholdAmount }

//NewThresholdAmount returns a new ThresholdAmountField initialized with val and scale
func NewThresholdAmount(val decimal.Decimal, scale int32) ThresholdAmountField {
	return ThresholdAmountField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f ThresholdAmountField) Value() (val decimal.Decimal) { return f.Decimal }

//TickDirectionField is a enum.TickDirection field
type TickDirectionField struct{ quickfix.FIXString }

//Tag returns tag.TickDirection (274)
func (f TickDirectionField) Tag() quickfix.Tag { return tag.TickDirection }

func NewTickDirection(val enum.TickDirection) TickDirectionField {
	return TickDirectionField{quickfix.FIXString(val)}
}

func (f TickDirectionField) Value() enum.TickDirection { return enum.TickDirection(f.String()) }

//TickIncrementField is a PRICE field
type TickIncrementField struct{ quickfix.FIXDecimal }

//Tag returns tag.TickIncrement (1208)
func (f TickIncrementField) Tag() quickfix.Tag { return tag.TickIncrement }

//NewTickIncrement returns a new TickIncrementField initialized with val and scale
func NewTickIncrement(val decimal.Decimal, scale int32) TickIncrementField {
	return TickIncrementField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f TickIncrementField) Value() (val decimal.Decimal) { return f.Decimal }

//TickRuleTypeField is a enum.TickRuleType field
type TickRuleTypeField struct{ quickfix.FIXString }

//Tag returns tag.TickRuleType (1209)
func (f TickRuleTypeField) Tag() quickfix.Tag { return tag.TickRuleType }

func NewTickRuleType(val enum.TickRuleType) TickRuleTypeField {
	return TickRuleTypeField{quickfix.FIXString(val)}
}

func (f TickRuleTypeField) Value() enum.TickRuleType { return enum.TickRuleType(f.String()) }

//TierCodeField is a STRING field
type TierCodeField struct{ quickfix.FIXString }

//Tag returns tag.TierCode (994)
func (f TierCodeField) Tag() quickfix.Tag { return tag.TierCode }

//NewTierCode returns a new TierCodeField initialized with val
func NewTierCode(val string) TierCodeField {
	return TierCodeField{quickfix.FIXString(val)}
}

func (f TierCodeField) Value() string { return f.String() }

//TimeBracketField is a STRING field
type TimeBracketField struct{ quickfix.FIXString }

//Tag returns tag.TimeBracket (943)
func (f TimeBracketField) Tag() quickfix.Tag { return tag.TimeBracket }

//NewTimeBracket returns a new TimeBracketField initialized with val
func NewTimeBracket(val string) TimeBracketField {
	return TimeBracketField{quickfix.FIXString(val)}
}

func (f TimeBracketField) Value() string { return f.String() }

//TimeInForceField is a enum.TimeInForce field
type TimeInForceField struct{ quickfix.FIXString }

//Tag returns tag.TimeInForce (59)
func (f TimeInForceField) Tag() quickfix.Tag { return tag.TimeInForce }

func NewTimeInForce(val enum.TimeInForce) TimeInForceField {
	return TimeInForceField{quickfix.FIXString(val)}
}

func (f TimeInForceField) Value() enum.TimeInForce { return enum.TimeInForce(f.String()) }

//TimeToExpirationField is a FLOAT field
type TimeToExpirationField struct{ quickfix.FIXDecimal }

//Tag returns tag.TimeToExpiration (1189)
func (f TimeToExpirationField) Tag() quickfix.Tag { return tag.TimeToExpiration }

//NewTimeToExpiration returns a new TimeToExpirationField initialized with val and scale
func NewTimeToExpiration(val decimal.Decimal, scale int32) TimeToExpirationField {
	return TimeToExpirationField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f TimeToExpirationField) Value() (val decimal.Decimal) { return f.Decimal }

//TimeUnitField is a enum.TimeUnit field
type TimeUnitField struct{ quickfix.FIXString }

//Tag returns tag.TimeUnit (997)
func (f TimeUnitField) Tag() quickfix.Tag { return tag.TimeUnit }

func NewTimeUnit(val enum.TimeUnit) TimeUnitField {
	return TimeUnitField{quickfix.FIXString(val)}
}

func (f TimeUnitField) Value() enum.TimeUnit { return enum.TimeUnit(f.String()) }

//TotNoAccQuotesField is a INT field
type TotNoAccQuotesField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoAccQuotes (1169)
func (f TotNoAccQuotesField) Tag() quickfix.Tag { return tag.TotNoAccQuotes }

//NewTotNoAccQuotes returns a new TotNoAccQuotesField initialized with val
func NewTotNoAccQuotes(val int) TotNoAccQuotesField {
	return TotNoAccQuotesField{quickfix.FIXInt(val)}
}

func (f TotNoAccQuotesField) Value() int { return f.Int() }

//TotNoAllocsField is a INT field
type TotNoAllocsField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoAllocs (892)
func (f TotNoAllocsField) Tag() quickfix.Tag { return tag.TotNoAllocs }

//NewTotNoAllocs returns a new TotNoAllocsField initialized with val
func NewTotNoAllocs(val int) TotNoAllocsField {
	return TotNoAllocsField{quickfix.FIXInt(val)}
}

func (f TotNoAllocsField) Value() int { return f.Int() }

//TotNoCxldQuotesField is a INT field
type TotNoCxldQuotesField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoCxldQuotes (1168)
func (f TotNoCxldQuotesField) Tag() quickfix.Tag { return tag.TotNoCxldQuotes }

//NewTotNoCxldQuotes returns a new TotNoCxldQuotesField initialized with val
func NewTotNoCxldQuotes(val int) TotNoCxldQuotesField {
	return TotNoCxldQuotesField{quickfix.FIXInt(val)}
}

func (f TotNoCxldQuotesField) Value() int { return f.Int() }

//TotNoFillsField is a INT field
type TotNoFillsField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoFills (1361)
func (f TotNoFillsField) Tag() quickfix.Tag { return tag.TotNoFills }

//NewTotNoFills returns a new TotNoFillsField initialized with val
func NewTotNoFills(val int) TotNoFillsField {
	return TotNoFillsField{quickfix.FIXInt(val)}
}

func (f TotNoFillsField) Value() int { return f.Int() }

//TotNoOrdersField is a INT field
type TotNoOrdersField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoOrders (68)
func (f TotNoOrdersField) Tag() quickfix.Tag { return tag.TotNoOrders }

//NewTotNoOrders returns a new TotNoOrdersField initialized with val
func NewTotNoOrders(val int) TotNoOrdersField {
	return TotNoOrdersField{quickfix.FIXInt(val)}
}

func (f TotNoOrdersField) Value() int { return f.Int() }

//TotNoPartyListField is a INT field
type TotNoPartyListField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoPartyList (1512)
func (f TotNoPartyListField) Tag() quickfix.Tag { return tag.TotNoPartyList }

//NewTotNoPartyList returns a new TotNoPartyListField initialized with val
func NewTotNoPartyList(val int) TotNoPartyListField {
	return TotNoPartyListField{quickfix.FIXInt(val)}
}

func (f TotNoPartyListField) Value() int { return f.Int() }

//TotNoQuoteEntriesField is a INT field
type TotNoQuoteEntriesField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoQuoteEntries (304)
func (f TotNoQuoteEntriesField) Tag() quickfix.Tag { return tag.TotNoQuoteEntries }

//NewTotNoQuoteEntries returns a new TotNoQuoteEntriesField initialized with val
func NewTotNoQuoteEntries(val int) TotNoQuoteEntriesField {
	return TotNoQuoteEntriesField{quickfix.FIXInt(val)}
}

func (f TotNoQuoteEntriesField) Value() int { return f.Int() }

//TotNoRejQuotesField is a INT field
type TotNoRejQuotesField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoRejQuotes (1170)
func (f TotNoRejQuotesField) Tag() quickfix.Tag { return tag.TotNoRejQuotes }

//NewTotNoRejQuotes returns a new TotNoRejQuotesField initialized with val
func NewTotNoRejQuotes(val int) TotNoRejQuotesField {
	return TotNoRejQuotesField{quickfix.FIXInt(val)}
}

func (f TotNoRejQuotesField) Value() int { return f.Int() }

//TotNoRelatedSymField is a INT field
type TotNoRelatedSymField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoRelatedSym (393)
func (f TotNoRelatedSymField) Tag() quickfix.Tag { return tag.TotNoRelatedSym }

//NewTotNoRelatedSym returns a new TotNoRelatedSymField initialized with val
func NewTotNoRelatedSym(val int) TotNoRelatedSymField {
	return TotNoRelatedSymField{quickfix.FIXInt(val)}
}

func (f TotNoRelatedSymField) Value() int { return f.Int() }

//TotNoSecurityTypesField is a INT field
type TotNoSecurityTypesField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoSecurityTypes (557)
func (f TotNoSecurityTypesField) Tag() quickfix.Tag { return tag.TotNoSecurityTypes }

//NewTotNoSecurityTypes returns a new TotNoSecurityTypesField initialized with val
func NewTotNoSecurityTypes(val int) TotNoSecurityTypesField {
	return TotNoSecurityTypesField{quickfix.FIXInt(val)}
}

func (f TotNoSecurityTypesField) Value() int { return f.Int() }

//TotNoStrikesField is a INT field
type TotNoStrikesField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoStrikes (422)
func (f TotNoStrikesField) Tag() quickfix.Tag { return tag.TotNoStrikes }

//NewTotNoStrikes returns a new TotNoStrikesField initialized with val
func NewTotNoStrikes(val int) TotNoStrikesField {
	return TotNoStrikesField{quickfix.FIXInt(val)}
}

func (f TotNoStrikesField) Value() int { return f.Int() }

//TotNumAssignmentReportsField is a INT field
type TotNumAssignmentReportsField struct{ quickfix.FIXInt }

//Tag returns tag.TotNumAssignmentReports (832)
func (f TotNumAssignmentReportsField) Tag() quickfix.Tag { return tag.TotNumAssignmentReports }

//NewTotNumAssignmentReports returns a new TotNumAssignmentReportsField initialized with val
func NewTotNumAssignmentReports(val int) TotNumAssignmentReportsField {
	return TotNumAssignmentReportsField{quickfix.FIXInt(val)}
}

func (f TotNumAssignmentReportsField) Value() int { return f.Int() }

//TotNumReportsField is a INT field
type TotNumReportsField struct{ quickfix.FIXInt }

//Tag returns tag.TotNumReports (911)
func (f TotNumReportsField) Tag() quickfix.Tag { return tag.TotNumReports }

//NewTotNumReports returns a new TotNumReportsField initialized with val
func NewTotNumReports(val int) TotNumReportsField {
	return TotNumReportsField{quickfix.FIXInt(val)}
}

func (f TotNumReportsField) Value() int { return f.Int() }

//TotNumTradeReportsField is a INT field
type TotNumTradeReportsField struct{ quickfix.FIXInt }

//Tag returns tag.TotNumTradeReports (748)
func (f TotNumTradeReportsField) Tag() quickfix.Tag { return tag.TotNumTradeReports }

//NewTotNumTradeReports returns a new TotNumTradeReportsField initialized with val
func NewTotNumTradeReports(val int) TotNumTradeReportsField {
	return TotNumTradeReportsField{quickfix.FIXInt(val)}
}

func (f TotNumTradeReportsField) Value() int { return f.Int() }

//TotQuoteEntriesField is a INT field
type TotQuoteEntriesField struct{ quickfix.FIXInt }

//Tag returns tag.TotQuoteEntries (304)
func (f TotQuoteEntriesField) Tag() quickfix.Tag { return tag.TotQuoteEntries }

//NewTotQuoteEntries returns a new TotQuoteEntriesField initialized with val
func NewTotQuoteEntries(val int) TotQuoteEntriesField {
	return TotQuoteEntriesField{quickfix.FIXInt(val)}
}

func (f TotQuoteEntriesField) Value() int { return f.Int() }

//TotalAccruedInterestAmtField is a AMT field
type TotalAccruedInterestAmtField struct{ quickfix.FIXDecimal }

//Tag returns tag.TotalAccruedInterestAmt (540)
func (f TotalAccruedInterestAmtField) Tag() quickfix.Tag { return tag.TotalAccruedInterestAmt }

//NewTotalAccruedInterestAmt returns a new TotalAccruedInterestAmtField initialized with val and scale
func NewTotalAccruedInterestAmt(val decimal.Decimal, scale int32) TotalAccruedInterestAmtField {
	return TotalAccruedInterestAmtField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f TotalAccruedInterestAmtField) Value() (val decimal.Decimal) { return f.Decimal }

//TotalAffectedOrdersField is a INT field
type TotalAffectedOrdersField struct{ quickfix.FIXInt }

//Tag returns tag.TotalAffectedOrders (533)
func (f TotalAffectedOrdersField) Tag() quickfix.Tag { return tag.TotalAffectedOrders }

//NewTotalAffectedOrders returns a new TotalAffectedOrdersField initialized with val
func NewTotalAffectedOrders(val int) TotalAffectedOrdersField {
	return TotalAffectedOrdersField{quickfix.FIXInt(val)}
}

func (f TotalAffectedOrdersField) Value() int { return f.Int() }

//TotalNetValueField is a AMT field
type TotalNetValueField struct{ quickfix.FIXDecimal }

//Tag returns tag.TotalNetValue (900)
func (f TotalNetValueField) Tag() quickfix.Tag { return tag.TotalNetValue }

//NewTotalNetValue returns a new TotalNetValueField initialized with val and scale
func NewTotalNetValue(val decimal.Decimal, scale int32) TotalNetValueField {
	return TotalNetValueField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f TotalNetValueField) Value() (val decimal.Decimal) { return f.Decimal }

//TotalNumPosReportsField is a INT field
type TotalNumPosReportsField struct{ quickfix.FIXInt }

//Tag returns tag.TotalNumPosReports (727)
func (f TotalNumPosReportsField) Tag() quickfix.Tag { return tag.TotalNumPosReports }

//NewTotalNumPosReports returns a new TotalNumPosReportsField initialized with val
func NewTotalNumPosReports(val int) TotalNumPosReportsField {
	return TotalNumPosReportsField{quickfix.FIXInt(val)}
}

func (f TotalNumPosReportsField) Value() int { return f.Int() }

//TotalNumSecuritiesField is a INT field
type TotalNumSecuritiesField struct{ quickfix.FIXInt }

//Tag returns tag.TotalNumSecurities (393)
func (f TotalNumSecuritiesField) Tag() quickfix.Tag { return tag.TotalNumSecurities }

//NewTotalNumSecurities returns a new TotalNumSecuritiesField initialized with val
func NewTotalNumSecurities(val int) TotalNumSecuritiesField {
	return TotalNumSecuritiesField{quickfix.FIXInt(val)}
}

func (f TotalNumSecuritiesField) Value() int { return f.Int() }

//TotalNumSecurityTypesField is a INT field
type TotalNumSecurityTypesField struct{ quickfix.FIXInt }

//Tag returns tag.TotalNumSecurityTypes (557)
func (f TotalNumSecurityTypesField) Tag() quickfix.Tag { return tag.TotalNumSecurityTypes }

//NewTotalNumSecurityTypes returns a new TotalNumSecurityTypesField initialized with val
func NewTotalNumSecurityTypes(val int) TotalNumSecurityTypesField {
	return TotalNumSecurityTypesField{quickfix.FIXInt(val)}
}

func (f TotalNumSecurityTypesField) Value() int { return f.Int() }

//TotalTakedownField is a AMT field
type TotalTakedownField struct{ quickfix.FIXDecimal }

//Tag returns tag.TotalTakedown (237)
func (f TotalTakedownField) Tag() quickfix.Tag { return tag.TotalTakedown }

//NewTotalTakedown returns a new TotalTakedownField initialized with val and scale
func NewTotalTakedown(val decimal.Decimal, scale int32) TotalTakedownField {
	return TotalTakedownField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f TotalTakedownField) Value() (val decimal.Decimal) { return f.Decimal }

//TotalVolumeTradedField is a QTY field
type TotalVolumeTradedField struct{ quickfix.FIXDecimal }

//Tag returns tag.TotalVolumeTraded (387)
func (f TotalVolumeTradedField) Tag() quickfix.Tag { return tag.TotalVolumeTraded }

//NewTotalVolumeTraded returns a new TotalVolumeTradedField initialized with val and scale
func NewTotalVolumeTraded(val decimal.Decimal, scale int32) TotalVolumeTradedField {
	return TotalVolumeTradedField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f TotalVolumeTradedField) Value() (val decimal.Decimal) { return f.Decimal }

//TotalVolumeTradedDateField is a UTCDATEONLY field
type TotalVolumeTradedDateField struct{ quickfix.FIXString }

//Tag returns tag.TotalVolumeTradedDate (449)
func (f TotalVolumeTradedDateField) Tag() quickfix.Tag { return tag.TotalVolumeTradedDate }

//NewTotalVolumeTradedDate returns a new TotalVolumeTradedDateField initialized with val
func NewTotalVolumeTradedDate(val string) TotalVolumeTradedDateField {
	return TotalVolumeTradedDateField{quickfix.FIXString(val)}
}

func (f TotalVolumeTradedDateField) Value() string { return f.String() }

//TotalVolumeTradedTimeField is a UTCTIMEONLY field
type TotalVolumeTradedTimeField struct{ quickfix.FIXString }

//Tag returns tag.TotalVolumeTradedTime (450)
func (f TotalVolumeTradedTimeField) Tag() quickfix.Tag { return tag.TotalVolumeTradedTime }

//NewTotalVolumeTradedTime returns a new TotalVolumeTradedTimeField initialized with val
func NewTotalVolumeTradedTime(val string) TotalVolumeTradedTimeField {
	return TotalVolumeTradedTimeField{quickfix.FIXString(val)}
}

func (f TotalVolumeTradedTimeField) Value() string { return f.String() }

//TradSesCloseTimeField is a UTCTIMESTAMP field
type TradSesCloseTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.TradSesCloseTime (344)
func (f TradSesCloseTimeField) Tag() quickfix.Tag { return tag.TradSesCloseTime }

//NewTradSesCloseTime returns a new TradSesCloseTimeField initialized with val
func NewTradSesCloseTime(val time.Time) TradSesCloseTimeField {
	return TradSesCloseTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewTradSesCloseTimeNoMillis returns a new TradSesCloseTimeField initialized with val without millisecs
func NewTradSesCloseTimeNoMillis(val time.Time) TradSesCloseTimeField {
	return TradSesCloseTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f TradSesCloseTimeField) Value() time.Time { return f.Time }

//TradSesEndTimeField is a UTCTIMESTAMP field
type TradSesEndTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.TradSesEndTime (345)
func (f TradSesEndTimeField) Tag() quickfix.Tag { return tag.TradSesEndTime }

//NewTradSesEndTime returns a new TradSesEndTimeField initialized with val
func NewTradSesEndTime(val time.Time) TradSesEndTimeField {
	return TradSesEndTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewTradSesEndTimeNoMillis returns a new TradSesEndTimeField initialized with val without millisecs
func NewTradSesEndTimeNoMillis(val time.Time) TradSesEndTimeField {
	return TradSesEndTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f TradSesEndTimeField) Value() time.Time { return f.Time }

//TradSesEventField is a enum.TradSesEvent field
type TradSesEventField struct{ quickfix.FIXString }

//Tag returns tag.TradSesEvent (1368)
func (f TradSesEventField) Tag() quickfix.Tag { return tag.TradSesEvent }

func NewTradSesEvent(val enum.TradSesEvent) TradSesEventField {
	return TradSesEventField{quickfix.FIXString(val)}
}

func (f TradSesEventField) Value() enum.TradSesEvent { return enum.TradSesEvent(f.String()) }

//TradSesMethodField is a enum.TradSesMethod field
type TradSesMethodField struct{ quickfix.FIXString }

//Tag returns tag.TradSesMethod (338)
func (f TradSesMethodField) Tag() quickfix.Tag { return tag.TradSesMethod }

func NewTradSesMethod(val enum.TradSesMethod) TradSesMethodField {
	return TradSesMethodField{quickfix.FIXString(val)}
}

func (f TradSesMethodField) Value() enum.TradSesMethod { return enum.TradSesMethod(f.String()) }

//TradSesModeField is a enum.TradSesMode field
type TradSesModeField struct{ quickfix.FIXString }

//Tag returns tag.TradSesMode (339)
func (f TradSesModeField) Tag() quickfix.Tag { return tag.TradSesMode }

func NewTradSesMode(val enum.TradSesMode) TradSesModeField {
	return TradSesModeField{quickfix.FIXString(val)}
}

func (f TradSesModeField) Value() enum.TradSesMode { return enum.TradSesMode(f.String()) }

//TradSesOpenTimeField is a UTCTIMESTAMP field
type TradSesOpenTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.TradSesOpenTime (342)
func (f TradSesOpenTimeField) Tag() quickfix.Tag { return tag.TradSesOpenTime }

//NewTradSesOpenTime returns a new TradSesOpenTimeField initialized with val
func NewTradSesOpenTime(val time.Time) TradSesOpenTimeField {
	return TradSesOpenTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewTradSesOpenTimeNoMillis returns a new TradSesOpenTimeField initialized with val without millisecs
func NewTradSesOpenTimeNoMillis(val time.Time) TradSesOpenTimeField {
	return TradSesOpenTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f TradSesOpenTimeField) Value() time.Time { return f.Time }

//TradSesPreCloseTimeField is a UTCTIMESTAMP field
type TradSesPreCloseTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.TradSesPreCloseTime (343)
func (f TradSesPreCloseTimeField) Tag() quickfix.Tag { return tag.TradSesPreCloseTime }

//NewTradSesPreCloseTime returns a new TradSesPreCloseTimeField initialized with val
func NewTradSesPreCloseTime(val time.Time) TradSesPreCloseTimeField {
	return TradSesPreCloseTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewTradSesPreCloseTimeNoMillis returns a new TradSesPreCloseTimeField initialized with val without millisecs
func NewTradSesPreCloseTimeNoMillis(val time.Time) TradSesPreCloseTimeField {
	return TradSesPreCloseTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f TradSesPreCloseTimeField) Value() time.Time { return f.Time }

//TradSesReqIDField is a STRING field
type TradSesReqIDField struct{ quickfix.FIXString }

//Tag returns tag.TradSesReqID (335)
func (f TradSesReqIDField) Tag() quickfix.Tag { return tag.TradSesReqID }

//NewTradSesReqID returns a new TradSesReqIDField initialized with val
func NewTradSesReqID(val string) TradSesReqIDField {
	return TradSesReqIDField{quickfix.FIXString(val)}
}

func (f TradSesReqIDField) Value() string { return f.String() }

//TradSesStartTimeField is a UTCTIMESTAMP field
type TradSesStartTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.TradSesStartTime (341)
func (f TradSesStartTimeField) Tag() quickfix.Tag { return tag.TradSesStartTime }

//NewTradSesStartTime returns a new TradSesStartTimeField initialized with val
func NewTradSesStartTime(val time.Time) TradSesStartTimeField {
	return TradSesStartTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewTradSesStartTimeNoMillis returns a new TradSesStartTimeField initialized with val without millisecs
func NewTradSesStartTimeNoMillis(val time.Time) TradSesStartTimeField {
	return TradSesStartTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f TradSesStartTimeField) Value() time.Time { return f.Time }

//TradSesStatusField is a enum.TradSesStatus field
type TradSesStatusField struct{ quickfix.FIXString }

//Tag returns tag.TradSesStatus (340)
func (f TradSesStatusField) Tag() quickfix.Tag { return tag.TradSesStatus }

func NewTradSesStatus(val enum.TradSesStatus) TradSesStatusField {
	return TradSesStatusField{quickfix.FIXString(val)}
}

func (f TradSesStatusField) Value() enum.TradSesStatus { return enum.TradSesStatus(f.String()) }

//TradSesStatusRejReasonField is a enum.TradSesStatusRejReason field
type TradSesStatusRejReasonField struct{ quickfix.FIXString }

//Tag returns tag.TradSesStatusRejReason (567)
func (f TradSesStatusRejReasonField) Tag() quickfix.Tag { return tag.TradSesStatusRejReason }

func NewTradSesStatusRejReason(val enum.TradSesStatusRejReason) TradSesStatusRejReasonField {
	return TradSesStatusRejReasonField{quickfix.FIXString(val)}
}

func (f TradSesStatusRejReasonField) Value() enum.TradSesStatusRejReason {
	return enum.TradSesStatusRejReason(f.String())
}

//TradSesUpdateActionField is a CHAR field
type TradSesUpdateActionField struct{ quickfix.FIXString }

//Tag returns tag.TradSesUpdateAction (1327)
func (f TradSesUpdateActionField) Tag() quickfix.Tag { return tag.TradSesUpdateAction }

//NewTradSesUpdateAction returns a new TradSesUpdateActionField initialized with val
func NewTradSesUpdateAction(val string) TradSesUpdateActionField {
	return TradSesUpdateActionField{quickfix.FIXString(val)}
}

func (f TradSesUpdateActionField) Value() string { return f.String() }

//TradeAllocIndicatorField is a enum.TradeAllocIndicator field
type TradeAllocIndicatorField struct{ quickfix.FIXString }

//Tag returns tag.TradeAllocIndicator (826)
func (f TradeAllocIndicatorField) Tag() quickfix.Tag { return tag.TradeAllocIndicator }

func NewTradeAllocIndicator(val enum.TradeAllocIndicator) TradeAllocIndicatorField {
	return TradeAllocIndicatorField{quickfix.FIXString(val)}
}

func (f TradeAllocIndicatorField) Value() enum.TradeAllocIndicator {
	return enum.TradeAllocIndicator(f.String())
}

//TradeConditionField is a enum.TradeCondition field
type TradeConditionField struct{ quickfix.FIXString }

//Tag returns tag.TradeCondition (277)
func (f TradeConditionField) Tag() quickfix.Tag { return tag.TradeCondition }

func NewTradeCondition(val enum.TradeCondition) TradeConditionField {
	return TradeConditionField{quickfix.FIXString(val)}
}

func (f TradeConditionField) Value() enum.TradeCondition { return enum.TradeCondition(f.String()) }

//TradeDateField is a LOCALMKTDATE field
type TradeDateField struct{ quickfix.FIXString }

//Tag returns tag.TradeDate (75)
func (f TradeDateField) Tag() quickfix.Tag { return tag.TradeDate }

//NewTradeDate returns a new TradeDateField initialized with val
func NewTradeDate(val string) TradeDateField {
	return TradeDateField{quickfix.FIXString(val)}
}

func (f TradeDateField) Value() string { return f.String() }

//TradeHandlingInstrField is a enum.TradeHandlingInstr field
type TradeHandlingInstrField struct{ quickfix.FIXString }

//Tag returns tag.TradeHandlingInstr (1123)
func (f TradeHandlingInstrField) Tag() quickfix.Tag { return tag.TradeHandlingInstr }

func NewTradeHandlingInstr(val enum.TradeHandlingInstr) TradeHandlingInstrField {
	return TradeHandlingInstrField{quickfix.FIXString(val)}
}

func (f TradeHandlingInstrField) Value() enum.TradeHandlingInstr {
	return enum.TradeHandlingInstr(f.String())
}

//TradeIDField is a STRING field
type TradeIDField struct{ quickfix.FIXString }

//Tag returns tag.TradeID (1003)
func (f TradeIDField) Tag() quickfix.Tag { return tag.TradeID }

//NewTradeID returns a new TradeIDField initialized with val
func NewTradeID(val string) TradeIDField {
	return TradeIDField{quickfix.FIXString(val)}
}

func (f TradeIDField) Value() string { return f.String() }

//TradeInputDeviceField is a STRING field
type TradeInputDeviceField struct{ quickfix.FIXString }

//Tag returns tag.TradeInputDevice (579)
func (f TradeInputDeviceField) Tag() quickfix.Tag { return tag.TradeInputDevice }

//NewTradeInputDevice returns a new TradeInputDeviceField initialized with val
func NewTradeInputDevice(val string) TradeInputDeviceField {
	return TradeInputDeviceField{quickfix.FIXString(val)}
}

func (f TradeInputDeviceField) Value() string { return f.String() }

//TradeInputSourceField is a STRING field
type TradeInputSourceField struct{ quickfix.FIXString }

//Tag returns tag.TradeInputSource (578)
func (f TradeInputSourceField) Tag() quickfix.Tag { return tag.TradeInputSource }

//NewTradeInputSource returns a new TradeInputSourceField initialized with val
func NewTradeInputSource(val string) TradeInputSourceField {
	return TradeInputSourceField{quickfix.FIXString(val)}
}

func (f TradeInputSourceField) Value() string { return f.String() }

//TradeLegRefIDField is a STRING field
type TradeLegRefIDField struct{ quickfix.FIXString }

//Tag returns tag.TradeLegRefID (824)
func (f TradeLegRefIDField) Tag() quickfix.Tag { return tag.TradeLegRefID }

//NewTradeLegRefID returns a new TradeLegRefIDField initialized with val
func NewTradeLegRefID(val string) TradeLegRefIDField {
	return TradeLegRefIDField{quickfix.FIXString(val)}
}

func (f TradeLegRefIDField) Value() string { return f.String() }

//TradeLinkIDField is a STRING field
type TradeLinkIDField struct{ quickfix.FIXString }

//Tag returns tag.TradeLinkID (820)
func (f TradeLinkIDField) Tag() quickfix.Tag { return tag.TradeLinkID }

//NewTradeLinkID returns a new TradeLinkIDField initialized with val
func NewTradeLinkID(val string) TradeLinkIDField {
	return TradeLinkIDField{quickfix.FIXString(val)}
}

func (f TradeLinkIDField) Value() string { return f.String() }

//TradeOriginationDateField is a LOCALMKTDATE field
type TradeOriginationDateField struct{ quickfix.FIXString }

//Tag returns tag.TradeOriginationDate (229)
func (f TradeOriginationDateField) Tag() quickfix.Tag { return tag.TradeOriginationDate }

//NewTradeOriginationDate returns a new TradeOriginationDateField initialized with val
func NewTradeOriginationDate(val string) TradeOriginationDateField {
	return TradeOriginationDateField{quickfix.FIXString(val)}
}

func (f TradeOriginationDateField) Value() string { return f.String() }

//TradePublishIndicatorField is a enum.TradePublishIndicator field
type TradePublishIndicatorField struct{ quickfix.FIXString }

//Tag returns tag.TradePublishIndicator (1390)
func (f TradePublishIndicatorField) Tag() quickfix.Tag { return tag.TradePublishIndicator }

func NewTradePublishIndicator(val enum.TradePublishIndicator) TradePublishIndicatorField {
	return TradePublishIndicatorField{quickfix.FIXString(val)}
}

func (f TradePublishIndicatorField) Value() enum.TradePublishIndicator {
	return enum.TradePublishIndicator(f.String())
}

//TradeReportIDField is a STRING field
type TradeReportIDField struct{ quickfix.FIXString }

//Tag returns tag.TradeReportID (571)
func (f TradeReportIDField) Tag() quickfix.Tag { return tag.TradeReportID }

//NewTradeReportID returns a new TradeReportIDField initialized with val
func NewTradeReportID(val string) TradeReportIDField {
	return TradeReportIDField{quickfix.FIXString(val)}
}

func (f TradeReportIDField) Value() string { return f.String() }

//TradeReportRefIDField is a STRING field
type TradeReportRefIDField struct{ quickfix.FIXString }

//Tag returns tag.TradeReportRefID (572)
func (f TradeReportRefIDField) Tag() quickfix.Tag { return tag.TradeReportRefID }

//NewTradeReportRefID returns a new TradeReportRefIDField initialized with val
func NewTradeReportRefID(val string) TradeReportRefIDField {
	return TradeReportRefIDField{quickfix.FIXString(val)}
}

func (f TradeReportRefIDField) Value() string { return f.String() }

//TradeReportRejectReasonField is a enum.TradeReportRejectReason field
type TradeReportRejectReasonField struct{ quickfix.FIXString }

//Tag returns tag.TradeReportRejectReason (751)
func (f TradeReportRejectReasonField) Tag() quickfix.Tag { return tag.TradeReportRejectReason }

func NewTradeReportRejectReason(val enum.TradeReportRejectReason) TradeReportRejectReasonField {
	return TradeReportRejectReasonField{quickfix.FIXString(val)}
}

func (f TradeReportRejectReasonField) Value() enum.TradeReportRejectReason {
	return enum.TradeReportRejectReason(f.String())
}

//TradeReportTransTypeField is a enum.TradeReportTransType field
type TradeReportTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.TradeReportTransType (487)
func (f TradeReportTransTypeField) Tag() quickfix.Tag { return tag.TradeReportTransType }

func NewTradeReportTransType(val enum.TradeReportTransType) TradeReportTransTypeField {
	return TradeReportTransTypeField{quickfix.FIXString(val)}
}

func (f TradeReportTransTypeField) Value() enum.TradeReportTransType {
	return enum.TradeReportTransType(f.String())
}

//TradeReportTypeField is a enum.TradeReportType field
type TradeReportTypeField struct{ quickfix.FIXString }

//Tag returns tag.TradeReportType (856)
func (f TradeReportTypeField) Tag() quickfix.Tag { return tag.TradeReportType }

func NewTradeReportType(val enum.TradeReportType) TradeReportTypeField {
	return TradeReportTypeField{quickfix.FIXString(val)}
}

func (f TradeReportTypeField) Value() enum.TradeReportType { return enum.TradeReportType(f.String()) }

//TradeRequestIDField is a STRING field
type TradeRequestIDField struct{ quickfix.FIXString }

//Tag returns tag.TradeRequestID (568)
func (f TradeRequestIDField) Tag() quickfix.Tag { return tag.TradeRequestID }

//NewTradeRequestID returns a new TradeRequestIDField initialized with val
func NewTradeRequestID(val string) TradeRequestIDField {
	return TradeRequestIDField{quickfix.FIXString(val)}
}

func (f TradeRequestIDField) Value() string { return f.String() }

//TradeRequestResultField is a enum.TradeRequestResult field
type TradeRequestResultField struct{ quickfix.FIXString }

//Tag returns tag.TradeRequestResult (749)
func (f TradeRequestResultField) Tag() quickfix.Tag { return tag.TradeRequestResult }

func NewTradeRequestResult(val enum.TradeRequestResult) TradeRequestResultField {
	return TradeRequestResultField{quickfix.FIXString(val)}
}

func (f TradeRequestResultField) Value() enum.TradeRequestResult {
	return enum.TradeRequestResult(f.String())
}

//TradeRequestStatusField is a enum.TradeRequestStatus field
type TradeRequestStatusField struct{ quickfix.FIXString }

//Tag returns tag.TradeRequestStatus (750)
func (f TradeRequestStatusField) Tag() quickfix.Tag { return tag.TradeRequestStatus }

func NewTradeRequestStatus(val enum.TradeRequestStatus) TradeRequestStatusField {
	return TradeRequestStatusField{quickfix.FIXString(val)}
}

func (f TradeRequestStatusField) Value() enum.TradeRequestStatus {
	return enum.TradeRequestStatus(f.String())
}

//TradeRequestTypeField is a enum.TradeRequestType field
type TradeRequestTypeField struct{ quickfix.FIXString }

//Tag returns tag.TradeRequestType (569)
func (f TradeRequestTypeField) Tag() quickfix.Tag { return tag.TradeRequestType }

func NewTradeRequestType(val enum.TradeRequestType) TradeRequestTypeField {
	return TradeRequestTypeField{quickfix.FIXString(val)}
}

func (f TradeRequestTypeField) Value() enum.TradeRequestType { return enum.TradeRequestType(f.String()) }

//TradeTypeField is a enum.TradeType field
type TradeTypeField struct{ quickfix.FIXString }

//Tag returns tag.TradeType (418)
func (f TradeTypeField) Tag() quickfix.Tag { return tag.TradeType }

func NewTradeType(val enum.TradeType) TradeTypeField {
	return TradeTypeField{quickfix.FIXString(val)}
}

func (f TradeTypeField) Value() enum.TradeType { return enum.TradeType(f.String()) }

//TradeVolumeField is a QTY field
type TradeVolumeField struct{ quickfix.FIXDecimal }

//Tag returns tag.TradeVolume (1020)
func (f TradeVolumeField) Tag() quickfix.Tag { return tag.TradeVolume }

//NewTradeVolume returns a new TradeVolumeField initialized with val and scale
func NewTradeVolume(val decimal.Decimal, scale int32) TradeVolumeField {
	return TradeVolumeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f TradeVolumeField) Value() (val decimal.Decimal) { return f.Decimal }

//TradedFlatSwitchField is a BOOLEAN field
type TradedFlatSwitchField struct{ quickfix.FIXBoolean }

//Tag returns tag.TradedFlatSwitch (258)
func (f TradedFlatSwitchField) Tag() quickfix.Tag { return tag.TradedFlatSwitch }

//NewTradedFlatSwitch returns a new TradedFlatSwitchField initialized with val
func NewTradedFlatSwitch(val bool) TradedFlatSwitchField {
	return TradedFlatSwitchField{quickfix.FIXBoolean(val)}
}

func (f TradedFlatSwitchField) Value() bool { return f.Bool() }

//TradingCurrencyField is a CURRENCY field
type TradingCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.TradingCurrency (1245)
func (f TradingCurrencyField) Tag() quickfix.Tag { return tag.TradingCurrency }

//NewTradingCurrency returns a new TradingCurrencyField initialized with val
func NewTradingCurrency(val string) TradingCurrencyField {
	return TradingCurrencyField{quickfix.FIXString(val)}
}

func (f TradingCurrencyField) Value() string { return f.String() }

//TradingReferencePriceField is a PRICE field
type TradingReferencePriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.TradingReferencePrice (1150)
func (f TradingReferencePriceField) Tag() quickfix.Tag { return tag.TradingReferencePrice }

//NewTradingReferencePrice returns a new TradingReferencePriceField initialized with val and scale
func NewTradingReferencePrice(val decimal.Decimal, scale int32) TradingReferencePriceField {
	return TradingReferencePriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f TradingReferencePriceField) Value() (val decimal.Decimal) { return f.Decimal }

//TradingSessionDescField is a STRING field
type TradingSessionDescField struct{ quickfix.FIXString }

//Tag returns tag.TradingSessionDesc (1326)
func (f TradingSessionDescField) Tag() quickfix.Tag { return tag.TradingSessionDesc }

//NewTradingSessionDesc returns a new TradingSessionDescField initialized with val
func NewTradingSessionDesc(val string) TradingSessionDescField {
	return TradingSessionDescField{quickfix.FIXString(val)}
}

func (f TradingSessionDescField) Value() string { return f.String() }

//TradingSessionIDField is a enum.TradingSessionID field
type TradingSessionIDField struct{ quickfix.FIXString }

//Tag returns tag.TradingSessionID (336)
func (f TradingSessionIDField) Tag() quickfix.Tag { return tag.TradingSessionID }

func NewTradingSessionID(val enum.TradingSessionID) TradingSessionIDField {
	return TradingSessionIDField{quickfix.FIXString(val)}
}

func (f TradingSessionIDField) Value() enum.TradingSessionID { return enum.TradingSessionID(f.String()) }

//TradingSessionSubIDField is a enum.TradingSessionSubID field
type TradingSessionSubIDField struct{ quickfix.FIXString }

//Tag returns tag.TradingSessionSubID (625)
func (f TradingSessionSubIDField) Tag() quickfix.Tag { return tag.TradingSessionSubID }

func NewTradingSessionSubID(val enum.TradingSessionSubID) TradingSessionSubIDField {
	return TradingSessionSubIDField{quickfix.FIXString(val)}
}

func (f TradingSessionSubIDField) Value() enum.TradingSessionSubID {
	return enum.TradingSessionSubID(f.String())
}

//TransBkdTimeField is a UTCTIMESTAMP field
type TransBkdTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.TransBkdTime (483)
func (f TransBkdTimeField) Tag() quickfix.Tag { return tag.TransBkdTime }

//NewTransBkdTime returns a new TransBkdTimeField initialized with val
func NewTransBkdTime(val time.Time) TransBkdTimeField {
	return TransBkdTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewTransBkdTimeNoMillis returns a new TransBkdTimeField initialized with val without millisecs
func NewTransBkdTimeNoMillis(val time.Time) TransBkdTimeField {
	return TransBkdTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f TransBkdTimeField) Value() time.Time { return f.Time }

//TransactTimeField is a UTCTIMESTAMP field
type TransactTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.TransactTime (60)
func (f TransactTimeField) Tag() quickfix.Tag { return tag.TransactTime }

//NewTransactTime returns a new TransactTimeField initialized with val
func NewTransactTime(val time.Time) TransactTimeField {
	return TransactTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewTransactTimeNoMillis returns a new TransactTimeField initialized with val without millisecs
func NewTransactTimeNoMillis(val time.Time) TransactTimeField {
	return TransactTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f TransactTimeField) Value() time.Time { return f.Time }

//TransferReasonField is a STRING field
type TransferReasonField struct{ quickfix.FIXString }

//Tag returns tag.TransferReason (830)
func (f TransferReasonField) Tag() quickfix.Tag { return tag.TransferReason }

//NewTransferReason returns a new TransferReasonField initialized with val
func NewTransferReason(val string) TransferReasonField {
	return TransferReasonField{quickfix.FIXString(val)}
}

func (f TransferReasonField) Value() string { return f.String() }

//TrdMatchIDField is a STRING field
type TrdMatchIDField struct{ quickfix.FIXString }

//Tag returns tag.TrdMatchID (880)
func (f TrdMatchIDField) Tag() quickfix.Tag { return tag.TrdMatchID }

//NewTrdMatchID returns a new TrdMatchIDField initialized with val
func NewTrdMatchID(val string) TrdMatchIDField {
	return TrdMatchIDField{quickfix.FIXString(val)}
}

func (f TrdMatchIDField) Value() string { return f.String() }

//TrdRegTimestampField is a UTCTIMESTAMP field
type TrdRegTimestampField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.TrdRegTimestamp (769)
func (f TrdRegTimestampField) Tag() quickfix.Tag { return tag.TrdRegTimestamp }

//NewTrdRegTimestamp returns a new TrdRegTimestampField initialized with val
func NewTrdRegTimestamp(val time.Time) TrdRegTimestampField {
	return TrdRegTimestampField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewTrdRegTimestampNoMillis returns a new TrdRegTimestampField initialized with val without millisecs
func NewTrdRegTimestampNoMillis(val time.Time) TrdRegTimestampField {
	return TrdRegTimestampField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f TrdRegTimestampField) Value() time.Time { return f.Time }

//TrdRegTimestampOriginField is a STRING field
type TrdRegTimestampOriginField struct{ quickfix.FIXString }

//Tag returns tag.TrdRegTimestampOrigin (771)
func (f TrdRegTimestampOriginField) Tag() quickfix.Tag { return tag.TrdRegTimestampOrigin }

//NewTrdRegTimestampOrigin returns a new TrdRegTimestampOriginField initialized with val
func NewTrdRegTimestampOrigin(val string) TrdRegTimestampOriginField {
	return TrdRegTimestampOriginField{quickfix.FIXString(val)}
}

func (f TrdRegTimestampOriginField) Value() string { return f.String() }

//TrdRegTimestampTypeField is a enum.TrdRegTimestampType field
type TrdRegTimestampTypeField struct{ quickfix.FIXString }

//Tag returns tag.TrdRegTimestampType (770)
func (f TrdRegTimestampTypeField) Tag() quickfix.Tag { return tag.TrdRegTimestampType }

func NewTrdRegTimestampType(val enum.TrdRegTimestampType) TrdRegTimestampTypeField {
	return TrdRegTimestampTypeField{quickfix.FIXString(val)}
}

func (f TrdRegTimestampTypeField) Value() enum.TrdRegTimestampType {
	return enum.TrdRegTimestampType(f.String())
}

//TrdRepIndicatorField is a BOOLEAN field
type TrdRepIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.TrdRepIndicator (1389)
func (f TrdRepIndicatorField) Tag() quickfix.Tag { return tag.TrdRepIndicator }

//NewTrdRepIndicator returns a new TrdRepIndicatorField initialized with val
func NewTrdRepIndicator(val bool) TrdRepIndicatorField {
	return TrdRepIndicatorField{quickfix.FIXBoolean(val)}
}

func (f TrdRepIndicatorField) Value() bool { return f.Bool() }

//TrdRepPartyRoleField is a INT field
type TrdRepPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.TrdRepPartyRole (1388)
func (f TrdRepPartyRoleField) Tag() quickfix.Tag { return tag.TrdRepPartyRole }

//NewTrdRepPartyRole returns a new TrdRepPartyRoleField initialized with val
func NewTrdRepPartyRole(val int) TrdRepPartyRoleField {
	return TrdRepPartyRoleField{quickfix.FIXInt(val)}
}

func (f TrdRepPartyRoleField) Value() int { return f.Int() }

//TrdRptStatusField is a enum.TrdRptStatus field
type TrdRptStatusField struct{ quickfix.FIXString }

//Tag returns tag.TrdRptStatus (939)
func (f TrdRptStatusField) Tag() quickfix.Tag { return tag.TrdRptStatus }

func NewTrdRptStatus(val enum.TrdRptStatus) TrdRptStatusField {
	return TrdRptStatusField{quickfix.FIXString(val)}
}

func (f TrdRptStatusField) Value() enum.TrdRptStatus { return enum.TrdRptStatus(f.String()) }

//TrdSubTypeField is a enum.TrdSubType field
type TrdSubTypeField struct{ quickfix.FIXString }

//Tag returns tag.TrdSubType (829)
func (f TrdSubTypeField) Tag() quickfix.Tag { return tag.TrdSubType }

func NewTrdSubType(val enum.TrdSubType) TrdSubTypeField {
	return TrdSubTypeField{quickfix.FIXString(val)}
}

func (f TrdSubTypeField) Value() enum.TrdSubType { return enum.TrdSubType(f.String()) }

//TrdTypeField is a enum.TrdType field
type TrdTypeField struct{ quickfix.FIXString }

//Tag returns tag.TrdType (828)
func (f TrdTypeField) Tag() quickfix.Tag { return tag.TrdType }

func NewTrdType(val enum.TrdType) TrdTypeField {
	return TrdTypeField{quickfix.FIXString(val)}
}

func (f TrdTypeField) Value() enum.TrdType { return enum.TrdType(f.String()) }

//TriggerActionField is a enum.TriggerAction field
type TriggerActionField struct{ quickfix.FIXString }

//Tag returns tag.TriggerAction (1101)
func (f TriggerActionField) Tag() quickfix.Tag { return tag.TriggerAction }

func NewTriggerAction(val enum.TriggerAction) TriggerActionField {
	return TriggerActionField{quickfix.FIXString(val)}
}

func (f TriggerActionField) Value() enum.TriggerAction { return enum.TriggerAction(f.String()) }

//TriggerNewPriceField is a PRICE field
type TriggerNewPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.TriggerNewPrice (1110)
func (f TriggerNewPriceField) Tag() quickfix.Tag { return tag.TriggerNewPrice }

//NewTriggerNewPrice returns a new TriggerNewPriceField initialized with val and scale
func NewTriggerNewPrice(val decimal.Decimal, scale int32) TriggerNewPriceField {
	return TriggerNewPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f TriggerNewPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//TriggerNewQtyField is a QTY field
type TriggerNewQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.TriggerNewQty (1112)
func (f TriggerNewQtyField) Tag() quickfix.Tag { return tag.TriggerNewQty }

//NewTriggerNewQty returns a new TriggerNewQtyField initialized with val and scale
func NewTriggerNewQty(val decimal.Decimal, scale int32) TriggerNewQtyField {
	return TriggerNewQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f TriggerNewQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//TriggerOrderTypeField is a enum.TriggerOrderType field
type TriggerOrderTypeField struct{ quickfix.FIXString }

//Tag returns tag.TriggerOrderType (1111)
func (f TriggerOrderTypeField) Tag() quickfix.Tag { return tag.TriggerOrderType }

func NewTriggerOrderType(val enum.TriggerOrderType) TriggerOrderTypeField {
	return TriggerOrderTypeField{quickfix.FIXString(val)}
}

func (f TriggerOrderTypeField) Value() enum.TriggerOrderType { return enum.TriggerOrderType(f.String()) }

//TriggerPriceField is a PRICE field
type TriggerPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.TriggerPrice (1102)
func (f TriggerPriceField) Tag() quickfix.Tag { return tag.TriggerPrice }

//NewTriggerPrice returns a new TriggerPriceField initialized with val and scale
func NewTriggerPrice(val decimal.Decimal, scale int32) TriggerPriceField {
	return TriggerPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f TriggerPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//TriggerPriceDirectionField is a enum.TriggerPriceDirection field
type TriggerPriceDirectionField struct{ quickfix.FIXString }

//Tag returns tag.TriggerPriceDirection (1109)
func (f TriggerPriceDirectionField) Tag() quickfix.Tag { return tag.TriggerPriceDirection }

func NewTriggerPriceDirection(val enum.TriggerPriceDirection) TriggerPriceDirectionField {
	return TriggerPriceDirectionField{quickfix.FIXString(val)}
}

func (f TriggerPriceDirectionField) Value() enum.TriggerPriceDirection {
	return enum.TriggerPriceDirection(f.String())
}

//TriggerPriceTypeField is a enum.TriggerPriceType field
type TriggerPriceTypeField struct{ quickfix.FIXString }

//Tag returns tag.TriggerPriceType (1107)
func (f TriggerPriceTypeField) Tag() quickfix.Tag { return tag.TriggerPriceType }

func NewTriggerPriceType(val enum.TriggerPriceType) TriggerPriceTypeField {
	return TriggerPriceTypeField{quickfix.FIXString(val)}
}

func (f TriggerPriceTypeField) Value() enum.TriggerPriceType { return enum.TriggerPriceType(f.String()) }

//TriggerPriceTypeScopeField is a enum.TriggerPriceTypeScope field
type TriggerPriceTypeScopeField struct{ quickfix.FIXString }

//Tag returns tag.TriggerPriceTypeScope (1108)
func (f TriggerPriceTypeScopeField) Tag() quickfix.Tag { return tag.TriggerPriceTypeScope }

func NewTriggerPriceTypeScope(val enum.TriggerPriceTypeScope) TriggerPriceTypeScopeField {
	return TriggerPriceTypeScopeField{quickfix.FIXString(val)}
}

func (f TriggerPriceTypeScopeField) Value() enum.TriggerPriceTypeScope {
	return enum.TriggerPriceTypeScope(f.String())
}

//TriggerSecurityDescField is a STRING field
type TriggerSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.TriggerSecurityDesc (1106)
func (f TriggerSecurityDescField) Tag() quickfix.Tag { return tag.TriggerSecurityDesc }

//NewTriggerSecurityDesc returns a new TriggerSecurityDescField initialized with val
func NewTriggerSecurityDesc(val string) TriggerSecurityDescField {
	return TriggerSecurityDescField{quickfix.FIXString(val)}
}

func (f TriggerSecurityDescField) Value() string { return f.String() }

//TriggerSecurityIDField is a STRING field
type TriggerSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.TriggerSecurityID (1104)
func (f TriggerSecurityIDField) Tag() quickfix.Tag { return tag.TriggerSecurityID }

//NewTriggerSecurityID returns a new TriggerSecurityIDField initialized with val
func NewTriggerSecurityID(val string) TriggerSecurityIDField {
	return TriggerSecurityIDField{quickfix.FIXString(val)}
}

func (f TriggerSecurityIDField) Value() string { return f.String() }

//TriggerSecurityIDSourceField is a STRING field
type TriggerSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.TriggerSecurityIDSource (1105)
func (f TriggerSecurityIDSourceField) Tag() quickfix.Tag { return tag.TriggerSecurityIDSource }

//NewTriggerSecurityIDSource returns a new TriggerSecurityIDSourceField initialized with val
func NewTriggerSecurityIDSource(val string) TriggerSecurityIDSourceField {
	return TriggerSecurityIDSourceField{quickfix.FIXString(val)}
}

func (f TriggerSecurityIDSourceField) Value() string { return f.String() }

//TriggerSymbolField is a STRING field
type TriggerSymbolField struct{ quickfix.FIXString }

//Tag returns tag.TriggerSymbol (1103)
func (f TriggerSymbolField) Tag() quickfix.Tag { return tag.TriggerSymbol }

//NewTriggerSymbol returns a new TriggerSymbolField initialized with val
func NewTriggerSymbol(val string) TriggerSymbolField {
	return TriggerSymbolField{quickfix.FIXString(val)}
}

func (f TriggerSymbolField) Value() string { return f.String() }

//TriggerTradingSessionIDField is a STRING field
type TriggerTradingSessionIDField struct{ quickfix.FIXString }

//Tag returns tag.TriggerTradingSessionID (1113)
func (f TriggerTradingSessionIDField) Tag() quickfix.Tag { return tag.TriggerTradingSessionID }

//NewTriggerTradingSessionID returns a new TriggerTradingSessionIDField initialized with val
func NewTriggerTradingSessionID(val string) TriggerTradingSessionIDField {
	return TriggerTradingSessionIDField{quickfix.FIXString(val)}
}

func (f TriggerTradingSessionIDField) Value() string { return f.String() }

//TriggerTradingSessionSubIDField is a STRING field
type TriggerTradingSessionSubIDField struct{ quickfix.FIXString }

//Tag returns tag.TriggerTradingSessionSubID (1114)
func (f TriggerTradingSessionSubIDField) Tag() quickfix.Tag { return tag.TriggerTradingSessionSubID }

//NewTriggerTradingSessionSubID returns a new TriggerTradingSessionSubIDField initialized with val
func NewTriggerTradingSessionSubID(val string) TriggerTradingSessionSubIDField {
	return TriggerTradingSessionSubIDField{quickfix.FIXString(val)}
}

func (f TriggerTradingSessionSubIDField) Value() string { return f.String() }

//TriggerTypeField is a enum.TriggerType field
type TriggerTypeField struct{ quickfix.FIXString }

//Tag returns tag.TriggerType (1100)
func (f TriggerTypeField) Tag() quickfix.Tag { return tag.TriggerType }

func NewTriggerType(val enum.TriggerType) TriggerTypeField {
	return TriggerTypeField{quickfix.FIXString(val)}
}

func (f TriggerTypeField) Value() enum.TriggerType { return enum.TriggerType(f.String()) }

//URLLinkField is a STRING field
type URLLinkField struct{ quickfix.FIXString }

//Tag returns tag.URLLink (149)
func (f URLLinkField) Tag() quickfix.Tag { return tag.URLLink }

//NewURLLink returns a new URLLinkField initialized with val
func NewURLLink(val string) URLLinkField {
	return URLLinkField{quickfix.FIXString(val)}
}

func (f URLLinkField) Value() string { return f.String() }

//UnderlyingAdjustedQuantityField is a QTY field
type UnderlyingAdjustedQuantityField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingAdjustedQuantity (1044)
func (f UnderlyingAdjustedQuantityField) Tag() quickfix.Tag { return tag.UnderlyingAdjustedQuantity }

//NewUnderlyingAdjustedQuantity returns a new UnderlyingAdjustedQuantityField initialized with val and scale
func NewUnderlyingAdjustedQuantity(val decimal.Decimal, scale int32) UnderlyingAdjustedQuantityField {
	return UnderlyingAdjustedQuantityField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingAdjustedQuantityField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingAllocationPercentField is a PERCENTAGE field
type UnderlyingAllocationPercentField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingAllocationPercent (972)
func (f UnderlyingAllocationPercentField) Tag() quickfix.Tag { return tag.UnderlyingAllocationPercent }

//NewUnderlyingAllocationPercent returns a new UnderlyingAllocationPercentField initialized with val and scale
func NewUnderlyingAllocationPercent(val decimal.Decimal, scale int32) UnderlyingAllocationPercentField {
	return UnderlyingAllocationPercentField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingAllocationPercentField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingAttachmentPointField is a PERCENTAGE field
type UnderlyingAttachmentPointField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingAttachmentPoint (1459)
func (f UnderlyingAttachmentPointField) Tag() quickfix.Tag { return tag.UnderlyingAttachmentPoint }

//NewUnderlyingAttachmentPoint returns a new UnderlyingAttachmentPointField initialized with val and scale
func NewUnderlyingAttachmentPoint(val decimal.Decimal, scale int32) UnderlyingAttachmentPointField {
	return UnderlyingAttachmentPointField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingAttachmentPointField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingCFICodeField is a STRING field
type UnderlyingCFICodeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingCFICode (463)
func (f UnderlyingCFICodeField) Tag() quickfix.Tag { return tag.UnderlyingCFICode }

//NewUnderlyingCFICode returns a new UnderlyingCFICodeField initialized with val
func NewUnderlyingCFICode(val string) UnderlyingCFICodeField {
	return UnderlyingCFICodeField{quickfix.FIXString(val)}
}

func (f UnderlyingCFICodeField) Value() string { return f.String() }

//UnderlyingCPProgramField is a STRING field
type UnderlyingCPProgramField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingCPProgram (877)
func (f UnderlyingCPProgramField) Tag() quickfix.Tag { return tag.UnderlyingCPProgram }

//NewUnderlyingCPProgram returns a new UnderlyingCPProgramField initialized with val
func NewUnderlyingCPProgram(val string) UnderlyingCPProgramField {
	return UnderlyingCPProgramField{quickfix.FIXString(val)}
}

func (f UnderlyingCPProgramField) Value() string { return f.String() }

//UnderlyingCPRegTypeField is a STRING field
type UnderlyingCPRegTypeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingCPRegType (878)
func (f UnderlyingCPRegTypeField) Tag() quickfix.Tag { return tag.UnderlyingCPRegType }

//NewUnderlyingCPRegType returns a new UnderlyingCPRegTypeField initialized with val
func NewUnderlyingCPRegType(val string) UnderlyingCPRegTypeField {
	return UnderlyingCPRegTypeField{quickfix.FIXString(val)}
}

func (f UnderlyingCPRegTypeField) Value() string { return f.String() }

//UnderlyingCapValueField is a AMT field
type UnderlyingCapValueField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingCapValue (1038)
func (f UnderlyingCapValueField) Tag() quickfix.Tag { return tag.UnderlyingCapValue }

//NewUnderlyingCapValue returns a new UnderlyingCapValueField initialized with val and scale
func NewUnderlyingCapValue(val decimal.Decimal, scale int32) UnderlyingCapValueField {
	return UnderlyingCapValueField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingCapValueField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingCashAmountField is a AMT field
type UnderlyingCashAmountField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingCashAmount (973)
func (f UnderlyingCashAmountField) Tag() quickfix.Tag { return tag.UnderlyingCashAmount }

//NewUnderlyingCashAmount returns a new UnderlyingCashAmountField initialized with val and scale
func NewUnderlyingCashAmount(val decimal.Decimal, scale int32) UnderlyingCashAmountField {
	return UnderlyingCashAmountField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingCashAmountField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingCashTypeField is a enum.UnderlyingCashType field
type UnderlyingCashTypeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingCashType (974)
func (f UnderlyingCashTypeField) Tag() quickfix.Tag { return tag.UnderlyingCashType }

func NewUnderlyingCashType(val enum.UnderlyingCashType) UnderlyingCashTypeField {
	return UnderlyingCashTypeField{quickfix.FIXString(val)}
}

func (f UnderlyingCashTypeField) Value() enum.UnderlyingCashType {
	return enum.UnderlyingCashType(f.String())
}

//UnderlyingCollectAmountField is a AMT field
type UnderlyingCollectAmountField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingCollectAmount (986)
func (f UnderlyingCollectAmountField) Tag() quickfix.Tag { return tag.UnderlyingCollectAmount }

//NewUnderlyingCollectAmount returns a new UnderlyingCollectAmountField initialized with val and scale
func NewUnderlyingCollectAmount(val decimal.Decimal, scale int32) UnderlyingCollectAmountField {
	return UnderlyingCollectAmountField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingCollectAmountField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingContractMultiplierField is a FLOAT field
type UnderlyingContractMultiplierField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingContractMultiplier (436)
func (f UnderlyingContractMultiplierField) Tag() quickfix.Tag { return tag.UnderlyingContractMultiplier }

//NewUnderlyingContractMultiplier returns a new UnderlyingContractMultiplierField initialized with val and scale
func NewUnderlyingContractMultiplier(val decimal.Decimal, scale int32) UnderlyingContractMultiplierField {
	return UnderlyingContractMultiplierField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingContractMultiplierField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingContractMultiplierUnitField is a INT field
type UnderlyingContractMultiplierUnitField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingContractMultiplierUnit (1437)
func (f UnderlyingContractMultiplierUnitField) Tag() quickfix.Tag {
	return tag.UnderlyingContractMultiplierUnit
}

//NewUnderlyingContractMultiplierUnit returns a new UnderlyingContractMultiplierUnitField initialized with val
func NewUnderlyingContractMultiplierUnit(val int) UnderlyingContractMultiplierUnitField {
	return UnderlyingContractMultiplierUnitField{quickfix.FIXInt(val)}
}

func (f UnderlyingContractMultiplierUnitField) Value() int { return f.Int() }

//UnderlyingCountryOfIssueField is a COUNTRY field
type UnderlyingCountryOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingCountryOfIssue (592)
func (f UnderlyingCountryOfIssueField) Tag() quickfix.Tag { return tag.UnderlyingCountryOfIssue }

//NewUnderlyingCountryOfIssue returns a new UnderlyingCountryOfIssueField initialized with val
func NewUnderlyingCountryOfIssue(val string) UnderlyingCountryOfIssueField {
	return UnderlyingCountryOfIssueField{quickfix.FIXString(val)}
}

func (f UnderlyingCountryOfIssueField) Value() string { return f.String() }

//UnderlyingCouponPaymentDateField is a LOCALMKTDATE field
type UnderlyingCouponPaymentDateField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingCouponPaymentDate (241)
func (f UnderlyingCouponPaymentDateField) Tag() quickfix.Tag { return tag.UnderlyingCouponPaymentDate }

//NewUnderlyingCouponPaymentDate returns a new UnderlyingCouponPaymentDateField initialized with val
func NewUnderlyingCouponPaymentDate(val string) UnderlyingCouponPaymentDateField {
	return UnderlyingCouponPaymentDateField{quickfix.FIXString(val)}
}

func (f UnderlyingCouponPaymentDateField) Value() string { return f.String() }

//UnderlyingCouponRateField is a PERCENTAGE field
type UnderlyingCouponRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingCouponRate (435)
func (f UnderlyingCouponRateField) Tag() quickfix.Tag { return tag.UnderlyingCouponRate }

//NewUnderlyingCouponRate returns a new UnderlyingCouponRateField initialized with val and scale
func NewUnderlyingCouponRate(val decimal.Decimal, scale int32) UnderlyingCouponRateField {
	return UnderlyingCouponRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingCouponRateField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingCreditRatingField is a STRING field
type UnderlyingCreditRatingField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingCreditRating (256)
func (f UnderlyingCreditRatingField) Tag() quickfix.Tag { return tag.UnderlyingCreditRating }

//NewUnderlyingCreditRating returns a new UnderlyingCreditRatingField initialized with val
func NewUnderlyingCreditRating(val string) UnderlyingCreditRatingField {
	return UnderlyingCreditRatingField{quickfix.FIXString(val)}
}

func (f UnderlyingCreditRatingField) Value() string { return f.String() }

//UnderlyingCurrencyField is a CURRENCY field
type UnderlyingCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingCurrency (318)
func (f UnderlyingCurrencyField) Tag() quickfix.Tag { return tag.UnderlyingCurrency }

//NewUnderlyingCurrency returns a new UnderlyingCurrencyField initialized with val
func NewUnderlyingCurrency(val string) UnderlyingCurrencyField {
	return UnderlyingCurrencyField{quickfix.FIXString(val)}
}

func (f UnderlyingCurrencyField) Value() string { return f.String() }

//UnderlyingCurrentValueField is a AMT field
type UnderlyingCurrentValueField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingCurrentValue (885)
func (f UnderlyingCurrentValueField) Tag() quickfix.Tag { return tag.UnderlyingCurrentValue }

//NewUnderlyingCurrentValue returns a new UnderlyingCurrentValueField initialized with val and scale
func NewUnderlyingCurrentValue(val decimal.Decimal, scale int32) UnderlyingCurrentValueField {
	return UnderlyingCurrentValueField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingCurrentValueField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingDeliveryAmountField is a AMT field
type UnderlyingDeliveryAmountField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingDeliveryAmount (1037)
func (f UnderlyingDeliveryAmountField) Tag() quickfix.Tag { return tag.UnderlyingDeliveryAmount }

//NewUnderlyingDeliveryAmount returns a new UnderlyingDeliveryAmountField initialized with val and scale
func NewUnderlyingDeliveryAmount(val decimal.Decimal, scale int32) UnderlyingDeliveryAmountField {
	return UnderlyingDeliveryAmountField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingDeliveryAmountField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingDetachmentPointField is a PERCENTAGE field
type UnderlyingDetachmentPointField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingDetachmentPoint (1460)
func (f UnderlyingDetachmentPointField) Tag() quickfix.Tag { return tag.UnderlyingDetachmentPoint }

//NewUnderlyingDetachmentPoint returns a new UnderlyingDetachmentPointField initialized with val and scale
func NewUnderlyingDetachmentPoint(val decimal.Decimal, scale int32) UnderlyingDetachmentPointField {
	return UnderlyingDetachmentPointField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingDetachmentPointField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingDirtyPriceField is a PRICE field
type UnderlyingDirtyPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingDirtyPrice (882)
func (f UnderlyingDirtyPriceField) Tag() quickfix.Tag { return tag.UnderlyingDirtyPrice }

//NewUnderlyingDirtyPrice returns a new UnderlyingDirtyPriceField initialized with val and scale
func NewUnderlyingDirtyPrice(val decimal.Decimal, scale int32) UnderlyingDirtyPriceField {
	return UnderlyingDirtyPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingDirtyPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingEndPriceField is a PRICE field
type UnderlyingEndPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingEndPrice (883)
func (f UnderlyingEndPriceField) Tag() quickfix.Tag { return tag.UnderlyingEndPrice }

//NewUnderlyingEndPrice returns a new UnderlyingEndPriceField initialized with val and scale
func NewUnderlyingEndPrice(val decimal.Decimal, scale int32) UnderlyingEndPriceField {
	return UnderlyingEndPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingEndPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingEndValueField is a AMT field
type UnderlyingEndValueField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingEndValue (886)
func (f UnderlyingEndValueField) Tag() quickfix.Tag { return tag.UnderlyingEndValue }

//NewUnderlyingEndValue returns a new UnderlyingEndValueField initialized with val and scale
func NewUnderlyingEndValue(val decimal.Decimal, scale int32) UnderlyingEndValueField {
	return UnderlyingEndValueField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingEndValueField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingExerciseStyleField is a INT field
type UnderlyingExerciseStyleField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingExerciseStyle (1419)
func (f UnderlyingExerciseStyleField) Tag() quickfix.Tag { return tag.UnderlyingExerciseStyle }

//NewUnderlyingExerciseStyle returns a new UnderlyingExerciseStyleField initialized with val
func NewUnderlyingExerciseStyle(val int) UnderlyingExerciseStyleField {
	return UnderlyingExerciseStyleField{quickfix.FIXInt(val)}
}

func (f UnderlyingExerciseStyleField) Value() int { return f.Int() }

//UnderlyingFXRateField is a FLOAT field
type UnderlyingFXRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingFXRate (1045)
func (f UnderlyingFXRateField) Tag() quickfix.Tag { return tag.UnderlyingFXRate }

//NewUnderlyingFXRate returns a new UnderlyingFXRateField initialized with val and scale
func NewUnderlyingFXRate(val decimal.Decimal, scale int32) UnderlyingFXRateField {
	return UnderlyingFXRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingFXRateField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingFXRateCalcField is a enum.UnderlyingFXRateCalc field
type UnderlyingFXRateCalcField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingFXRateCalc (1046)
func (f UnderlyingFXRateCalcField) Tag() quickfix.Tag { return tag.UnderlyingFXRateCalc }

func NewUnderlyingFXRateCalc(val enum.UnderlyingFXRateCalc) UnderlyingFXRateCalcField {
	return UnderlyingFXRateCalcField{quickfix.FIXString(val)}
}

func (f UnderlyingFXRateCalcField) Value() enum.UnderlyingFXRateCalc {
	return enum.UnderlyingFXRateCalc(f.String())
}

//UnderlyingFactorField is a FLOAT field
type UnderlyingFactorField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingFactor (246)
func (f UnderlyingFactorField) Tag() quickfix.Tag { return tag.UnderlyingFactor }

//NewUnderlyingFactor returns a new UnderlyingFactorField initialized with val and scale
func NewUnderlyingFactor(val decimal.Decimal, scale int32) UnderlyingFactorField {
	return UnderlyingFactorField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingFactorField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingFlowScheduleTypeField is a INT field
type UnderlyingFlowScheduleTypeField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingFlowScheduleType (1441)
func (f UnderlyingFlowScheduleTypeField) Tag() quickfix.Tag { return tag.UnderlyingFlowScheduleType }

//NewUnderlyingFlowScheduleType returns a new UnderlyingFlowScheduleTypeField initialized with val
func NewUnderlyingFlowScheduleType(val int) UnderlyingFlowScheduleTypeField {
	return UnderlyingFlowScheduleTypeField{quickfix.FIXInt(val)}
}

func (f UnderlyingFlowScheduleTypeField) Value() int { return f.Int() }

//UnderlyingIDSourceField is a STRING field
type UnderlyingIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingIDSource (305)
func (f UnderlyingIDSourceField) Tag() quickfix.Tag { return tag.UnderlyingIDSource }

//NewUnderlyingIDSource returns a new UnderlyingIDSourceField initialized with val
func NewUnderlyingIDSource(val string) UnderlyingIDSourceField {
	return UnderlyingIDSourceField{quickfix.FIXString(val)}
}

func (f UnderlyingIDSourceField) Value() string { return f.String() }

//UnderlyingInstrRegistryField is a STRING field
type UnderlyingInstrRegistryField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingInstrRegistry (595)
func (f UnderlyingInstrRegistryField) Tag() quickfix.Tag { return tag.UnderlyingInstrRegistry }

//NewUnderlyingInstrRegistry returns a new UnderlyingInstrRegistryField initialized with val
func NewUnderlyingInstrRegistry(val string) UnderlyingInstrRegistryField {
	return UnderlyingInstrRegistryField{quickfix.FIXString(val)}
}

func (f UnderlyingInstrRegistryField) Value() string { return f.String() }

//UnderlyingInstrumentPartyIDField is a STRING field
type UnderlyingInstrumentPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingInstrumentPartyID (1059)
func (f UnderlyingInstrumentPartyIDField) Tag() quickfix.Tag { return tag.UnderlyingInstrumentPartyID }

//NewUnderlyingInstrumentPartyID returns a new UnderlyingInstrumentPartyIDField initialized with val
func NewUnderlyingInstrumentPartyID(val string) UnderlyingInstrumentPartyIDField {
	return UnderlyingInstrumentPartyIDField{quickfix.FIXString(val)}
}

func (f UnderlyingInstrumentPartyIDField) Value() string { return f.String() }

//UnderlyingInstrumentPartyIDSourceField is a CHAR field
type UnderlyingInstrumentPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingInstrumentPartyIDSource (1060)
func (f UnderlyingInstrumentPartyIDSourceField) Tag() quickfix.Tag {
	return tag.UnderlyingInstrumentPartyIDSource
}

//NewUnderlyingInstrumentPartyIDSource returns a new UnderlyingInstrumentPartyIDSourceField initialized with val
func NewUnderlyingInstrumentPartyIDSource(val string) UnderlyingInstrumentPartyIDSourceField {
	return UnderlyingInstrumentPartyIDSourceField{quickfix.FIXString(val)}
}

func (f UnderlyingInstrumentPartyIDSourceField) Value() string { return f.String() }

//UnderlyingInstrumentPartyRoleField is a INT field
type UnderlyingInstrumentPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingInstrumentPartyRole (1061)
func (f UnderlyingInstrumentPartyRoleField) Tag() quickfix.Tag {
	return tag.UnderlyingInstrumentPartyRole
}

//NewUnderlyingInstrumentPartyRole returns a new UnderlyingInstrumentPartyRoleField initialized with val
func NewUnderlyingInstrumentPartyRole(val int) UnderlyingInstrumentPartyRoleField {
	return UnderlyingInstrumentPartyRoleField{quickfix.FIXInt(val)}
}

func (f UnderlyingInstrumentPartyRoleField) Value() int { return f.Int() }

//UnderlyingInstrumentPartySubIDField is a STRING field
type UnderlyingInstrumentPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingInstrumentPartySubID (1063)
func (f UnderlyingInstrumentPartySubIDField) Tag() quickfix.Tag {
	return tag.UnderlyingInstrumentPartySubID
}

//NewUnderlyingInstrumentPartySubID returns a new UnderlyingInstrumentPartySubIDField initialized with val
func NewUnderlyingInstrumentPartySubID(val string) UnderlyingInstrumentPartySubIDField {
	return UnderlyingInstrumentPartySubIDField{quickfix.FIXString(val)}
}

func (f UnderlyingInstrumentPartySubIDField) Value() string { return f.String() }

//UnderlyingInstrumentPartySubIDTypeField is a INT field
type UnderlyingInstrumentPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingInstrumentPartySubIDType (1064)
func (f UnderlyingInstrumentPartySubIDTypeField) Tag() quickfix.Tag {
	return tag.UnderlyingInstrumentPartySubIDType
}

//NewUnderlyingInstrumentPartySubIDType returns a new UnderlyingInstrumentPartySubIDTypeField initialized with val
func NewUnderlyingInstrumentPartySubIDType(val int) UnderlyingInstrumentPartySubIDTypeField {
	return UnderlyingInstrumentPartySubIDTypeField{quickfix.FIXInt(val)}
}

func (f UnderlyingInstrumentPartySubIDTypeField) Value() int { return f.Int() }

//UnderlyingIssueDateField is a LOCALMKTDATE field
type UnderlyingIssueDateField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingIssueDate (242)
func (f UnderlyingIssueDateField) Tag() quickfix.Tag { return tag.UnderlyingIssueDate }

//NewUnderlyingIssueDate returns a new UnderlyingIssueDateField initialized with val
func NewUnderlyingIssueDate(val string) UnderlyingIssueDateField {
	return UnderlyingIssueDateField{quickfix.FIXString(val)}
}

func (f UnderlyingIssueDateField) Value() string { return f.String() }

//UnderlyingIssuerField is a STRING field
type UnderlyingIssuerField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingIssuer (306)
func (f UnderlyingIssuerField) Tag() quickfix.Tag { return tag.UnderlyingIssuer }

//NewUnderlyingIssuer returns a new UnderlyingIssuerField initialized with val
func NewUnderlyingIssuer(val string) UnderlyingIssuerField {
	return UnderlyingIssuerField{quickfix.FIXString(val)}
}

func (f UnderlyingIssuerField) Value() string { return f.String() }

//UnderlyingLastPxField is a PRICE field
type UnderlyingLastPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingLastPx (651)
func (f UnderlyingLastPxField) Tag() quickfix.Tag { return tag.UnderlyingLastPx }

//NewUnderlyingLastPx returns a new UnderlyingLastPxField initialized with val and scale
func NewUnderlyingLastPx(val decimal.Decimal, scale int32) UnderlyingLastPxField {
	return UnderlyingLastPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingLastPxField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingLastQtyField is a QTY field
type UnderlyingLastQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingLastQty (652)
func (f UnderlyingLastQtyField) Tag() quickfix.Tag { return tag.UnderlyingLastQty }

//NewUnderlyingLastQty returns a new UnderlyingLastQtyField initialized with val and scale
func NewUnderlyingLastQty(val decimal.Decimal, scale int32) UnderlyingLastQtyField {
	return UnderlyingLastQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingLastQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingLegCFICodeField is a STRING field
type UnderlyingLegCFICodeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegCFICode (1344)
func (f UnderlyingLegCFICodeField) Tag() quickfix.Tag { return tag.UnderlyingLegCFICode }

//NewUnderlyingLegCFICode returns a new UnderlyingLegCFICodeField initialized with val
func NewUnderlyingLegCFICode(val string) UnderlyingLegCFICodeField {
	return UnderlyingLegCFICodeField{quickfix.FIXString(val)}
}

func (f UnderlyingLegCFICodeField) Value() string { return f.String() }

//UnderlyingLegMaturityDateField is a LOCALMKTDATE field
type UnderlyingLegMaturityDateField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegMaturityDate (1345)
func (f UnderlyingLegMaturityDateField) Tag() quickfix.Tag { return tag.UnderlyingLegMaturityDate }

//NewUnderlyingLegMaturityDate returns a new UnderlyingLegMaturityDateField initialized with val
func NewUnderlyingLegMaturityDate(val string) UnderlyingLegMaturityDateField {
	return UnderlyingLegMaturityDateField{quickfix.FIXString(val)}
}

func (f UnderlyingLegMaturityDateField) Value() string { return f.String() }

//UnderlyingLegMaturityMonthYearField is a MONTHYEAR field
type UnderlyingLegMaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegMaturityMonthYear (1339)
func (f UnderlyingLegMaturityMonthYearField) Tag() quickfix.Tag {
	return tag.UnderlyingLegMaturityMonthYear
}

//NewUnderlyingLegMaturityMonthYear returns a new UnderlyingLegMaturityMonthYearField initialized with val
func NewUnderlyingLegMaturityMonthYear(val string) UnderlyingLegMaturityMonthYearField {
	return UnderlyingLegMaturityMonthYearField{quickfix.FIXString(val)}
}

func (f UnderlyingLegMaturityMonthYearField) Value() string { return f.String() }

//UnderlyingLegMaturityTimeField is a TZTIMEONLY field
type UnderlyingLegMaturityTimeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegMaturityTime (1405)
func (f UnderlyingLegMaturityTimeField) Tag() quickfix.Tag { return tag.UnderlyingLegMaturityTime }

//NewUnderlyingLegMaturityTime returns a new UnderlyingLegMaturityTimeField initialized with val
func NewUnderlyingLegMaturityTime(val string) UnderlyingLegMaturityTimeField {
	return UnderlyingLegMaturityTimeField{quickfix.FIXString(val)}
}

func (f UnderlyingLegMaturityTimeField) Value() string { return f.String() }

//UnderlyingLegOptAttributeField is a CHAR field
type UnderlyingLegOptAttributeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegOptAttribute (1391)
func (f UnderlyingLegOptAttributeField) Tag() quickfix.Tag { return tag.UnderlyingLegOptAttribute }

//NewUnderlyingLegOptAttribute returns a new UnderlyingLegOptAttributeField initialized with val
func NewUnderlyingLegOptAttribute(val string) UnderlyingLegOptAttributeField {
	return UnderlyingLegOptAttributeField{quickfix.FIXString(val)}
}

func (f UnderlyingLegOptAttributeField) Value() string { return f.String() }

//UnderlyingLegPutOrCallField is a INT field
type UnderlyingLegPutOrCallField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingLegPutOrCall (1343)
func (f UnderlyingLegPutOrCallField) Tag() quickfix.Tag { return tag.UnderlyingLegPutOrCall }

//NewUnderlyingLegPutOrCall returns a new UnderlyingLegPutOrCallField initialized with val
func NewUnderlyingLegPutOrCall(val int) UnderlyingLegPutOrCallField {
	return UnderlyingLegPutOrCallField{quickfix.FIXInt(val)}
}

func (f UnderlyingLegPutOrCallField) Value() int { return f.Int() }

//UnderlyingLegSecurityAltIDField is a STRING field
type UnderlyingLegSecurityAltIDField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSecurityAltID (1335)
func (f UnderlyingLegSecurityAltIDField) Tag() quickfix.Tag { return tag.UnderlyingLegSecurityAltID }

//NewUnderlyingLegSecurityAltID returns a new UnderlyingLegSecurityAltIDField initialized with val
func NewUnderlyingLegSecurityAltID(val string) UnderlyingLegSecurityAltIDField {
	return UnderlyingLegSecurityAltIDField{quickfix.FIXString(val)}
}

func (f UnderlyingLegSecurityAltIDField) Value() string { return f.String() }

//UnderlyingLegSecurityAltIDSourceField is a STRING field
type UnderlyingLegSecurityAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSecurityAltIDSource (1336)
func (f UnderlyingLegSecurityAltIDSourceField) Tag() quickfix.Tag {
	return tag.UnderlyingLegSecurityAltIDSource
}

//NewUnderlyingLegSecurityAltIDSource returns a new UnderlyingLegSecurityAltIDSourceField initialized with val
func NewUnderlyingLegSecurityAltIDSource(val string) UnderlyingLegSecurityAltIDSourceField {
	return UnderlyingLegSecurityAltIDSourceField{quickfix.FIXString(val)}
}

func (f UnderlyingLegSecurityAltIDSourceField) Value() string { return f.String() }

//UnderlyingLegSecurityDescField is a STRING field
type UnderlyingLegSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSecurityDesc (1392)
func (f UnderlyingLegSecurityDescField) Tag() quickfix.Tag { return tag.UnderlyingLegSecurityDesc }

//NewUnderlyingLegSecurityDesc returns a new UnderlyingLegSecurityDescField initialized with val
func NewUnderlyingLegSecurityDesc(val string) UnderlyingLegSecurityDescField {
	return UnderlyingLegSecurityDescField{quickfix.FIXString(val)}
}

func (f UnderlyingLegSecurityDescField) Value() string { return f.String() }

//UnderlyingLegSecurityExchangeField is a STRING field
type UnderlyingLegSecurityExchangeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSecurityExchange (1341)
func (f UnderlyingLegSecurityExchangeField) Tag() quickfix.Tag {
	return tag.UnderlyingLegSecurityExchange
}

//NewUnderlyingLegSecurityExchange returns a new UnderlyingLegSecurityExchangeField initialized with val
func NewUnderlyingLegSecurityExchange(val string) UnderlyingLegSecurityExchangeField {
	return UnderlyingLegSecurityExchangeField{quickfix.FIXString(val)}
}

func (f UnderlyingLegSecurityExchangeField) Value() string { return f.String() }

//UnderlyingLegSecurityIDField is a STRING field
type UnderlyingLegSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSecurityID (1332)
func (f UnderlyingLegSecurityIDField) Tag() quickfix.Tag { return tag.UnderlyingLegSecurityID }

//NewUnderlyingLegSecurityID returns a new UnderlyingLegSecurityIDField initialized with val
func NewUnderlyingLegSecurityID(val string) UnderlyingLegSecurityIDField {
	return UnderlyingLegSecurityIDField{quickfix.FIXString(val)}
}

func (f UnderlyingLegSecurityIDField) Value() string { return f.String() }

//UnderlyingLegSecurityIDSourceField is a STRING field
type UnderlyingLegSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSecurityIDSource (1333)
func (f UnderlyingLegSecurityIDSourceField) Tag() quickfix.Tag {
	return tag.UnderlyingLegSecurityIDSource
}

//NewUnderlyingLegSecurityIDSource returns a new UnderlyingLegSecurityIDSourceField initialized with val
func NewUnderlyingLegSecurityIDSource(val string) UnderlyingLegSecurityIDSourceField {
	return UnderlyingLegSecurityIDSourceField{quickfix.FIXString(val)}
}

func (f UnderlyingLegSecurityIDSourceField) Value() string { return f.String() }

//UnderlyingLegSecuritySubTypeField is a STRING field
type UnderlyingLegSecuritySubTypeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSecuritySubType (1338)
func (f UnderlyingLegSecuritySubTypeField) Tag() quickfix.Tag { return tag.UnderlyingLegSecuritySubType }

//NewUnderlyingLegSecuritySubType returns a new UnderlyingLegSecuritySubTypeField initialized with val
func NewUnderlyingLegSecuritySubType(val string) UnderlyingLegSecuritySubTypeField {
	return UnderlyingLegSecuritySubTypeField{quickfix.FIXString(val)}
}

func (f UnderlyingLegSecuritySubTypeField) Value() string { return f.String() }

//UnderlyingLegSecurityTypeField is a STRING field
type UnderlyingLegSecurityTypeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSecurityType (1337)
func (f UnderlyingLegSecurityTypeField) Tag() quickfix.Tag { return tag.UnderlyingLegSecurityType }

//NewUnderlyingLegSecurityType returns a new UnderlyingLegSecurityTypeField initialized with val
func NewUnderlyingLegSecurityType(val string) UnderlyingLegSecurityTypeField {
	return UnderlyingLegSecurityTypeField{quickfix.FIXString(val)}
}

func (f UnderlyingLegSecurityTypeField) Value() string { return f.String() }

//UnderlyingLegStrikePriceField is a PRICE field
type UnderlyingLegStrikePriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingLegStrikePrice (1340)
func (f UnderlyingLegStrikePriceField) Tag() quickfix.Tag { return tag.UnderlyingLegStrikePrice }

//NewUnderlyingLegStrikePrice returns a new UnderlyingLegStrikePriceField initialized with val and scale
func NewUnderlyingLegStrikePrice(val decimal.Decimal, scale int32) UnderlyingLegStrikePriceField {
	return UnderlyingLegStrikePriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingLegStrikePriceField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingLegSymbolField is a STRING field
type UnderlyingLegSymbolField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSymbol (1330)
func (f UnderlyingLegSymbolField) Tag() quickfix.Tag { return tag.UnderlyingLegSymbol }

//NewUnderlyingLegSymbol returns a new UnderlyingLegSymbolField initialized with val
func NewUnderlyingLegSymbol(val string) UnderlyingLegSymbolField {
	return UnderlyingLegSymbolField{quickfix.FIXString(val)}
}

func (f UnderlyingLegSymbolField) Value() string { return f.String() }

//UnderlyingLegSymbolSfxField is a STRING field
type UnderlyingLegSymbolSfxField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSymbolSfx (1331)
func (f UnderlyingLegSymbolSfxField) Tag() quickfix.Tag { return tag.UnderlyingLegSymbolSfx }

//NewUnderlyingLegSymbolSfx returns a new UnderlyingLegSymbolSfxField initialized with val
func NewUnderlyingLegSymbolSfx(val string) UnderlyingLegSymbolSfxField {
	return UnderlyingLegSymbolSfxField{quickfix.FIXString(val)}
}

func (f UnderlyingLegSymbolSfxField) Value() string { return f.String() }

//UnderlyingLocaleOfIssueField is a STRING field
type UnderlyingLocaleOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLocaleOfIssue (594)
func (f UnderlyingLocaleOfIssueField) Tag() quickfix.Tag { return tag.UnderlyingLocaleOfIssue }

//NewUnderlyingLocaleOfIssue returns a new UnderlyingLocaleOfIssueField initialized with val
func NewUnderlyingLocaleOfIssue(val string) UnderlyingLocaleOfIssueField {
	return UnderlyingLocaleOfIssueField{quickfix.FIXString(val)}
}

func (f UnderlyingLocaleOfIssueField) Value() string { return f.String() }

//UnderlyingMaturityDateField is a LOCALMKTDATE field
type UnderlyingMaturityDateField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingMaturityDate (542)
func (f UnderlyingMaturityDateField) Tag() quickfix.Tag { return tag.UnderlyingMaturityDate }

//NewUnderlyingMaturityDate returns a new UnderlyingMaturityDateField initialized with val
func NewUnderlyingMaturityDate(val string) UnderlyingMaturityDateField {
	return UnderlyingMaturityDateField{quickfix.FIXString(val)}
}

func (f UnderlyingMaturityDateField) Value() string { return f.String() }

//UnderlyingMaturityDayField is a DAYOFMONTH field
type UnderlyingMaturityDayField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingMaturityDay (314)
func (f UnderlyingMaturityDayField) Tag() quickfix.Tag { return tag.UnderlyingMaturityDay }

//NewUnderlyingMaturityDay returns a new UnderlyingMaturityDayField initialized with val
func NewUnderlyingMaturityDay(val int) UnderlyingMaturityDayField {
	return UnderlyingMaturityDayField{quickfix.FIXInt(val)}
}

func (f UnderlyingMaturityDayField) Value() int { return f.Int() }

//UnderlyingMaturityMonthYearField is a MONTHYEAR field
type UnderlyingMaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingMaturityMonthYear (313)
func (f UnderlyingMaturityMonthYearField) Tag() quickfix.Tag { return tag.UnderlyingMaturityMonthYear }

//NewUnderlyingMaturityMonthYear returns a new UnderlyingMaturityMonthYearField initialized with val
func NewUnderlyingMaturityMonthYear(val string) UnderlyingMaturityMonthYearField {
	return UnderlyingMaturityMonthYearField{quickfix.FIXString(val)}
}

func (f UnderlyingMaturityMonthYearField) Value() string { return f.String() }

//UnderlyingMaturityTimeField is a TZTIMEONLY field
type UnderlyingMaturityTimeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingMaturityTime (1213)
func (f UnderlyingMaturityTimeField) Tag() quickfix.Tag { return tag.UnderlyingMaturityTime }

//NewUnderlyingMaturityTime returns a new UnderlyingMaturityTimeField initialized with val
func NewUnderlyingMaturityTime(val string) UnderlyingMaturityTimeField {
	return UnderlyingMaturityTimeField{quickfix.FIXString(val)}
}

func (f UnderlyingMaturityTimeField) Value() string { return f.String() }

//UnderlyingNotionalPercentageOutstandingField is a PERCENTAGE field
type UnderlyingNotionalPercentageOutstandingField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingNotionalPercentageOutstanding (1455)
func (f UnderlyingNotionalPercentageOutstandingField) Tag() quickfix.Tag {
	return tag.UnderlyingNotionalPercentageOutstanding
}

//NewUnderlyingNotionalPercentageOutstanding returns a new UnderlyingNotionalPercentageOutstandingField initialized with val and scale
func NewUnderlyingNotionalPercentageOutstanding(val decimal.Decimal, scale int32) UnderlyingNotionalPercentageOutstandingField {
	return UnderlyingNotionalPercentageOutstandingField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingNotionalPercentageOutstandingField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingOptAttributeField is a CHAR field
type UnderlyingOptAttributeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingOptAttribute (317)
func (f UnderlyingOptAttributeField) Tag() quickfix.Tag { return tag.UnderlyingOptAttribute }

//NewUnderlyingOptAttribute returns a new UnderlyingOptAttributeField initialized with val
func NewUnderlyingOptAttribute(val string) UnderlyingOptAttributeField {
	return UnderlyingOptAttributeField{quickfix.FIXString(val)}
}

func (f UnderlyingOptAttributeField) Value() string { return f.String() }

//UnderlyingOriginalNotionalPercentageOutstandingField is a PERCENTAGE field
type UnderlyingOriginalNotionalPercentageOutstandingField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingOriginalNotionalPercentageOutstanding (1456)
func (f UnderlyingOriginalNotionalPercentageOutstandingField) Tag() quickfix.Tag {
	return tag.UnderlyingOriginalNotionalPercentageOutstanding
}

//NewUnderlyingOriginalNotionalPercentageOutstanding returns a new UnderlyingOriginalNotionalPercentageOutstandingField initialized with val and scale
func NewUnderlyingOriginalNotionalPercentageOutstanding(val decimal.Decimal, scale int32) UnderlyingOriginalNotionalPercentageOutstandingField {
	return UnderlyingOriginalNotionalPercentageOutstandingField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingOriginalNotionalPercentageOutstandingField) Value() (val decimal.Decimal) {
	return f.Decimal
}

//UnderlyingPayAmountField is a AMT field
type UnderlyingPayAmountField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingPayAmount (985)
func (f UnderlyingPayAmountField) Tag() quickfix.Tag { return tag.UnderlyingPayAmount }

//NewUnderlyingPayAmount returns a new UnderlyingPayAmountField initialized with val and scale
func NewUnderlyingPayAmount(val decimal.Decimal, scale int32) UnderlyingPayAmountField {
	return UnderlyingPayAmountField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingPayAmountField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingPriceDeterminationMethodField is a enum.UnderlyingPriceDeterminationMethod field
type UnderlyingPriceDeterminationMethodField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingPriceDeterminationMethod (1481)
func (f UnderlyingPriceDeterminationMethodField) Tag() quickfix.Tag {
	return tag.UnderlyingPriceDeterminationMethod
}

func NewUnderlyingPriceDeterminationMethod(val enum.UnderlyingPriceDeterminationMethod) UnderlyingPriceDeterminationMethodField {
	return UnderlyingPriceDeterminationMethodField{quickfix.FIXString(val)}
}

func (f UnderlyingPriceDeterminationMethodField) Value() enum.UnderlyingPriceDeterminationMethod {
	return enum.UnderlyingPriceDeterminationMethod(f.String())
}

//UnderlyingPriceUnitOfMeasureField is a STRING field
type UnderlyingPriceUnitOfMeasureField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingPriceUnitOfMeasure (1424)
func (f UnderlyingPriceUnitOfMeasureField) Tag() quickfix.Tag { return tag.UnderlyingPriceUnitOfMeasure }

//NewUnderlyingPriceUnitOfMeasure returns a new UnderlyingPriceUnitOfMeasureField initialized with val
func NewUnderlyingPriceUnitOfMeasure(val string) UnderlyingPriceUnitOfMeasureField {
	return UnderlyingPriceUnitOfMeasureField{quickfix.FIXString(val)}
}

func (f UnderlyingPriceUnitOfMeasureField) Value() string { return f.String() }

//UnderlyingPriceUnitOfMeasureQtyField is a QTY field
type UnderlyingPriceUnitOfMeasureQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingPriceUnitOfMeasureQty (1425)
func (f UnderlyingPriceUnitOfMeasureQtyField) Tag() quickfix.Tag {
	return tag.UnderlyingPriceUnitOfMeasureQty
}

//NewUnderlyingPriceUnitOfMeasureQty returns a new UnderlyingPriceUnitOfMeasureQtyField initialized with val and scale
func NewUnderlyingPriceUnitOfMeasureQty(val decimal.Decimal, scale int32) UnderlyingPriceUnitOfMeasureQtyField {
	return UnderlyingPriceUnitOfMeasureQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingPriceUnitOfMeasureQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingProductField is a INT field
type UnderlyingProductField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingProduct (462)
func (f UnderlyingProductField) Tag() quickfix.Tag { return tag.UnderlyingProduct }

//NewUnderlyingProduct returns a new UnderlyingProductField initialized with val
func NewUnderlyingProduct(val int) UnderlyingProductField {
	return UnderlyingProductField{quickfix.FIXInt(val)}
}

func (f UnderlyingProductField) Value() int { return f.Int() }

//UnderlyingPutOrCallField is a INT field
type UnderlyingPutOrCallField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingPutOrCall (315)
func (f UnderlyingPutOrCallField) Tag() quickfix.Tag { return tag.UnderlyingPutOrCall }

//NewUnderlyingPutOrCall returns a new UnderlyingPutOrCallField initialized with val
func NewUnderlyingPutOrCall(val int) UnderlyingPutOrCallField {
	return UnderlyingPutOrCallField{quickfix.FIXInt(val)}
}

func (f UnderlyingPutOrCallField) Value() int { return f.Int() }

//UnderlyingPxField is a PRICE field
type UnderlyingPxField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingPx (810)
func (f UnderlyingPxField) Tag() quickfix.Tag { return tag.UnderlyingPx }

//NewUnderlyingPx returns a new UnderlyingPxField initialized with val and scale
func NewUnderlyingPx(val decimal.Decimal, scale int32) UnderlyingPxField {
	return UnderlyingPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingPxField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingQtyField is a QTY field
type UnderlyingQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingQty (879)
func (f UnderlyingQtyField) Tag() quickfix.Tag { return tag.UnderlyingQty }

//NewUnderlyingQty returns a new UnderlyingQtyField initialized with val and scale
func NewUnderlyingQty(val decimal.Decimal, scale int32) UnderlyingQtyField {
	return UnderlyingQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingRedemptionDateField is a LOCALMKTDATE field
type UnderlyingRedemptionDateField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingRedemptionDate (247)
func (f UnderlyingRedemptionDateField) Tag() quickfix.Tag { return tag.UnderlyingRedemptionDate }

//NewUnderlyingRedemptionDate returns a new UnderlyingRedemptionDateField initialized with val
func NewUnderlyingRedemptionDate(val string) UnderlyingRedemptionDateField {
	return UnderlyingRedemptionDateField{quickfix.FIXString(val)}
}

func (f UnderlyingRedemptionDateField) Value() string { return f.String() }

//UnderlyingRepoCollateralSecurityTypeField is a INT field
type UnderlyingRepoCollateralSecurityTypeField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingRepoCollateralSecurityType (243)
func (f UnderlyingRepoCollateralSecurityTypeField) Tag() quickfix.Tag {
	return tag.UnderlyingRepoCollateralSecurityType
}

//NewUnderlyingRepoCollateralSecurityType returns a new UnderlyingRepoCollateralSecurityTypeField initialized with val
func NewUnderlyingRepoCollateralSecurityType(val int) UnderlyingRepoCollateralSecurityTypeField {
	return UnderlyingRepoCollateralSecurityTypeField{quickfix.FIXInt(val)}
}

func (f UnderlyingRepoCollateralSecurityTypeField) Value() int { return f.Int() }

//UnderlyingRepurchaseRateField is a PERCENTAGE field
type UnderlyingRepurchaseRateField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingRepurchaseRate (245)
func (f UnderlyingRepurchaseRateField) Tag() quickfix.Tag { return tag.UnderlyingRepurchaseRate }

//NewUnderlyingRepurchaseRate returns a new UnderlyingRepurchaseRateField initialized with val and scale
func NewUnderlyingRepurchaseRate(val decimal.Decimal, scale int32) UnderlyingRepurchaseRateField {
	return UnderlyingRepurchaseRateField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingRepurchaseRateField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingRepurchaseTermField is a INT field
type UnderlyingRepurchaseTermField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingRepurchaseTerm (244)
func (f UnderlyingRepurchaseTermField) Tag() quickfix.Tag { return tag.UnderlyingRepurchaseTerm }

//NewUnderlyingRepurchaseTerm returns a new UnderlyingRepurchaseTermField initialized with val
func NewUnderlyingRepurchaseTerm(val int) UnderlyingRepurchaseTermField {
	return UnderlyingRepurchaseTermField{quickfix.FIXInt(val)}
}

func (f UnderlyingRepurchaseTermField) Value() int { return f.Int() }

//UnderlyingRestructuringTypeField is a STRING field
type UnderlyingRestructuringTypeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingRestructuringType (1453)
func (f UnderlyingRestructuringTypeField) Tag() quickfix.Tag { return tag.UnderlyingRestructuringType }

//NewUnderlyingRestructuringType returns a new UnderlyingRestructuringTypeField initialized with val
func NewUnderlyingRestructuringType(val string) UnderlyingRestructuringTypeField {
	return UnderlyingRestructuringTypeField{quickfix.FIXString(val)}
}

func (f UnderlyingRestructuringTypeField) Value() string { return f.String() }

//UnderlyingSecurityAltIDField is a STRING field
type UnderlyingSecurityAltIDField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSecurityAltID (458)
func (f UnderlyingSecurityAltIDField) Tag() quickfix.Tag { return tag.UnderlyingSecurityAltID }

//NewUnderlyingSecurityAltID returns a new UnderlyingSecurityAltIDField initialized with val
func NewUnderlyingSecurityAltID(val string) UnderlyingSecurityAltIDField {
	return UnderlyingSecurityAltIDField{quickfix.FIXString(val)}
}

func (f UnderlyingSecurityAltIDField) Value() string { return f.String() }

//UnderlyingSecurityAltIDSourceField is a STRING field
type UnderlyingSecurityAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSecurityAltIDSource (459)
func (f UnderlyingSecurityAltIDSourceField) Tag() quickfix.Tag {
	return tag.UnderlyingSecurityAltIDSource
}

//NewUnderlyingSecurityAltIDSource returns a new UnderlyingSecurityAltIDSourceField initialized with val
func NewUnderlyingSecurityAltIDSource(val string) UnderlyingSecurityAltIDSourceField {
	return UnderlyingSecurityAltIDSourceField{quickfix.FIXString(val)}
}

func (f UnderlyingSecurityAltIDSourceField) Value() string { return f.String() }

//UnderlyingSecurityDescField is a STRING field
type UnderlyingSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSecurityDesc (307)
func (f UnderlyingSecurityDescField) Tag() quickfix.Tag { return tag.UnderlyingSecurityDesc }

//NewUnderlyingSecurityDesc returns a new UnderlyingSecurityDescField initialized with val
func NewUnderlyingSecurityDesc(val string) UnderlyingSecurityDescField {
	return UnderlyingSecurityDescField{quickfix.FIXString(val)}
}

func (f UnderlyingSecurityDescField) Value() string { return f.String() }

//UnderlyingSecurityExchangeField is a EXCHANGE field
type UnderlyingSecurityExchangeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSecurityExchange (308)
func (f UnderlyingSecurityExchangeField) Tag() quickfix.Tag { return tag.UnderlyingSecurityExchange }

//NewUnderlyingSecurityExchange returns a new UnderlyingSecurityExchangeField initialized with val
func NewUnderlyingSecurityExchange(val string) UnderlyingSecurityExchangeField {
	return UnderlyingSecurityExchangeField{quickfix.FIXString(val)}
}

func (f UnderlyingSecurityExchangeField) Value() string { return f.String() }

//UnderlyingSecurityIDField is a STRING field
type UnderlyingSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSecurityID (309)
func (f UnderlyingSecurityIDField) Tag() quickfix.Tag { return tag.UnderlyingSecurityID }

//NewUnderlyingSecurityID returns a new UnderlyingSecurityIDField initialized with val
func NewUnderlyingSecurityID(val string) UnderlyingSecurityIDField {
	return UnderlyingSecurityIDField{quickfix.FIXString(val)}
}

func (f UnderlyingSecurityIDField) Value() string { return f.String() }

//UnderlyingSecurityIDSourceField is a STRING field
type UnderlyingSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSecurityIDSource (305)
func (f UnderlyingSecurityIDSourceField) Tag() quickfix.Tag { return tag.UnderlyingSecurityIDSource }

//NewUnderlyingSecurityIDSource returns a new UnderlyingSecurityIDSourceField initialized with val
func NewUnderlyingSecurityIDSource(val string) UnderlyingSecurityIDSourceField {
	return UnderlyingSecurityIDSourceField{quickfix.FIXString(val)}
}

func (f UnderlyingSecurityIDSourceField) Value() string { return f.String() }

//UnderlyingSecuritySubTypeField is a STRING field
type UnderlyingSecuritySubTypeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSecuritySubType (763)
func (f UnderlyingSecuritySubTypeField) Tag() quickfix.Tag { return tag.UnderlyingSecuritySubType }

//NewUnderlyingSecuritySubType returns a new UnderlyingSecuritySubTypeField initialized with val
func NewUnderlyingSecuritySubType(val string) UnderlyingSecuritySubTypeField {
	return UnderlyingSecuritySubTypeField{quickfix.FIXString(val)}
}

func (f UnderlyingSecuritySubTypeField) Value() string { return f.String() }

//UnderlyingSecurityTypeField is a STRING field
type UnderlyingSecurityTypeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSecurityType (310)
func (f UnderlyingSecurityTypeField) Tag() quickfix.Tag { return tag.UnderlyingSecurityType }

//NewUnderlyingSecurityType returns a new UnderlyingSecurityTypeField initialized with val
func NewUnderlyingSecurityType(val string) UnderlyingSecurityTypeField {
	return UnderlyingSecurityTypeField{quickfix.FIXString(val)}
}

func (f UnderlyingSecurityTypeField) Value() string { return f.String() }

//UnderlyingSeniorityField is a STRING field
type UnderlyingSeniorityField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSeniority (1454)
func (f UnderlyingSeniorityField) Tag() quickfix.Tag { return tag.UnderlyingSeniority }

//NewUnderlyingSeniority returns a new UnderlyingSeniorityField initialized with val
func NewUnderlyingSeniority(val string) UnderlyingSeniorityField {
	return UnderlyingSeniorityField{quickfix.FIXString(val)}
}

func (f UnderlyingSeniorityField) Value() string { return f.String() }

//UnderlyingSettlMethodField is a STRING field
type UnderlyingSettlMethodField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSettlMethod (1039)
func (f UnderlyingSettlMethodField) Tag() quickfix.Tag { return tag.UnderlyingSettlMethod }

//NewUnderlyingSettlMethod returns a new UnderlyingSettlMethodField initialized with val
func NewUnderlyingSettlMethod(val string) UnderlyingSettlMethodField {
	return UnderlyingSettlMethodField{quickfix.FIXString(val)}
}

func (f UnderlyingSettlMethodField) Value() string { return f.String() }

//UnderlyingSettlPriceField is a PRICE field
type UnderlyingSettlPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingSettlPrice (732)
func (f UnderlyingSettlPriceField) Tag() quickfix.Tag { return tag.UnderlyingSettlPrice }

//NewUnderlyingSettlPrice returns a new UnderlyingSettlPriceField initialized with val and scale
func NewUnderlyingSettlPrice(val decimal.Decimal, scale int32) UnderlyingSettlPriceField {
	return UnderlyingSettlPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingSettlPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingSettlPriceTypeField is a INT field
type UnderlyingSettlPriceTypeField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingSettlPriceType (733)
func (f UnderlyingSettlPriceTypeField) Tag() quickfix.Tag { return tag.UnderlyingSettlPriceType }

//NewUnderlyingSettlPriceType returns a new UnderlyingSettlPriceTypeField initialized with val
func NewUnderlyingSettlPriceType(val int) UnderlyingSettlPriceTypeField {
	return UnderlyingSettlPriceTypeField{quickfix.FIXInt(val)}
}

func (f UnderlyingSettlPriceTypeField) Value() int { return f.Int() }

//UnderlyingSettlementDateField is a LOCALMKTDATE field
type UnderlyingSettlementDateField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSettlementDate (987)
func (f UnderlyingSettlementDateField) Tag() quickfix.Tag { return tag.UnderlyingSettlementDate }

//NewUnderlyingSettlementDate returns a new UnderlyingSettlementDateField initialized with val
func NewUnderlyingSettlementDate(val string) UnderlyingSettlementDateField {
	return UnderlyingSettlementDateField{quickfix.FIXString(val)}
}

func (f UnderlyingSettlementDateField) Value() string { return f.String() }

//UnderlyingSettlementStatusField is a STRING field
type UnderlyingSettlementStatusField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSettlementStatus (988)
func (f UnderlyingSettlementStatusField) Tag() quickfix.Tag { return tag.UnderlyingSettlementStatus }

//NewUnderlyingSettlementStatus returns a new UnderlyingSettlementStatusField initialized with val
func NewUnderlyingSettlementStatus(val string) UnderlyingSettlementStatusField {
	return UnderlyingSettlementStatusField{quickfix.FIXString(val)}
}

func (f UnderlyingSettlementStatusField) Value() string { return f.String() }

//UnderlyingSettlementTypeField is a enum.UnderlyingSettlementType field
type UnderlyingSettlementTypeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSettlementType (975)
func (f UnderlyingSettlementTypeField) Tag() quickfix.Tag { return tag.UnderlyingSettlementType }

func NewUnderlyingSettlementType(val enum.UnderlyingSettlementType) UnderlyingSettlementTypeField {
	return UnderlyingSettlementTypeField{quickfix.FIXString(val)}
}

func (f UnderlyingSettlementTypeField) Value() enum.UnderlyingSettlementType {
	return enum.UnderlyingSettlementType(f.String())
}

//UnderlyingStartValueField is a AMT field
type UnderlyingStartValueField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingStartValue (884)
func (f UnderlyingStartValueField) Tag() quickfix.Tag { return tag.UnderlyingStartValue }

//NewUnderlyingStartValue returns a new UnderlyingStartValueField initialized with val and scale
func NewUnderlyingStartValue(val decimal.Decimal, scale int32) UnderlyingStartValueField {
	return UnderlyingStartValueField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingStartValueField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingStateOrProvinceOfIssueField is a STRING field
type UnderlyingStateOrProvinceOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingStateOrProvinceOfIssue (593)
func (f UnderlyingStateOrProvinceOfIssueField) Tag() quickfix.Tag {
	return tag.UnderlyingStateOrProvinceOfIssue
}

//NewUnderlyingStateOrProvinceOfIssue returns a new UnderlyingStateOrProvinceOfIssueField initialized with val
func NewUnderlyingStateOrProvinceOfIssue(val string) UnderlyingStateOrProvinceOfIssueField {
	return UnderlyingStateOrProvinceOfIssueField{quickfix.FIXString(val)}
}

func (f UnderlyingStateOrProvinceOfIssueField) Value() string { return f.String() }

//UnderlyingStipTypeField is a STRING field
type UnderlyingStipTypeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingStipType (888)
func (f UnderlyingStipTypeField) Tag() quickfix.Tag { return tag.UnderlyingStipType }

//NewUnderlyingStipType returns a new UnderlyingStipTypeField initialized with val
func NewUnderlyingStipType(val string) UnderlyingStipTypeField {
	return UnderlyingStipTypeField{quickfix.FIXString(val)}
}

func (f UnderlyingStipTypeField) Value() string { return f.String() }

//UnderlyingStipValueField is a STRING field
type UnderlyingStipValueField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingStipValue (889)
func (f UnderlyingStipValueField) Tag() quickfix.Tag { return tag.UnderlyingStipValue }

//NewUnderlyingStipValue returns a new UnderlyingStipValueField initialized with val
func NewUnderlyingStipValue(val string) UnderlyingStipValueField {
	return UnderlyingStipValueField{quickfix.FIXString(val)}
}

func (f UnderlyingStipValueField) Value() string { return f.String() }

//UnderlyingStrikeCurrencyField is a CURRENCY field
type UnderlyingStrikeCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingStrikeCurrency (941)
func (f UnderlyingStrikeCurrencyField) Tag() quickfix.Tag { return tag.UnderlyingStrikeCurrency }

//NewUnderlyingStrikeCurrency returns a new UnderlyingStrikeCurrencyField initialized with val
func NewUnderlyingStrikeCurrency(val string) UnderlyingStrikeCurrencyField {
	return UnderlyingStrikeCurrencyField{quickfix.FIXString(val)}
}

func (f UnderlyingStrikeCurrencyField) Value() string { return f.String() }

//UnderlyingStrikePriceField is a PRICE field
type UnderlyingStrikePriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingStrikePrice (316)
func (f UnderlyingStrikePriceField) Tag() quickfix.Tag { return tag.UnderlyingStrikePrice }

//NewUnderlyingStrikePrice returns a new UnderlyingStrikePriceField initialized with val and scale
func NewUnderlyingStrikePrice(val decimal.Decimal, scale int32) UnderlyingStrikePriceField {
	return UnderlyingStrikePriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingStrikePriceField) Value() (val decimal.Decimal) { return f.Decimal }

//UnderlyingSymbolField is a STRING field
type UnderlyingSymbolField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSymbol (311)
func (f UnderlyingSymbolField) Tag() quickfix.Tag { return tag.UnderlyingSymbol }

//NewUnderlyingSymbol returns a new UnderlyingSymbolField initialized with val
func NewUnderlyingSymbol(val string) UnderlyingSymbolField {
	return UnderlyingSymbolField{quickfix.FIXString(val)}
}

func (f UnderlyingSymbolField) Value() string { return f.String() }

//UnderlyingSymbolSfxField is a STRING field
type UnderlyingSymbolSfxField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSymbolSfx (312)
func (f UnderlyingSymbolSfxField) Tag() quickfix.Tag { return tag.UnderlyingSymbolSfx }

//NewUnderlyingSymbolSfx returns a new UnderlyingSymbolSfxField initialized with val
func NewUnderlyingSymbolSfx(val string) UnderlyingSymbolSfxField {
	return UnderlyingSymbolSfxField{quickfix.FIXString(val)}
}

func (f UnderlyingSymbolSfxField) Value() string { return f.String() }

//UnderlyingTimeUnitField is a STRING field
type UnderlyingTimeUnitField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingTimeUnit (1000)
func (f UnderlyingTimeUnitField) Tag() quickfix.Tag { return tag.UnderlyingTimeUnit }

//NewUnderlyingTimeUnit returns a new UnderlyingTimeUnitField initialized with val
func NewUnderlyingTimeUnit(val string) UnderlyingTimeUnitField {
	return UnderlyingTimeUnitField{quickfix.FIXString(val)}
}

func (f UnderlyingTimeUnitField) Value() string { return f.String() }

//UnderlyingTradingSessionIDField is a STRING field
type UnderlyingTradingSessionIDField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingTradingSessionID (822)
func (f UnderlyingTradingSessionIDField) Tag() quickfix.Tag { return tag.UnderlyingTradingSessionID }

//NewUnderlyingTradingSessionID returns a new UnderlyingTradingSessionIDField initialized with val
func NewUnderlyingTradingSessionID(val string) UnderlyingTradingSessionIDField {
	return UnderlyingTradingSessionIDField{quickfix.FIXString(val)}
}

func (f UnderlyingTradingSessionIDField) Value() string { return f.String() }

//UnderlyingTradingSessionSubIDField is a STRING field
type UnderlyingTradingSessionSubIDField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingTradingSessionSubID (823)
func (f UnderlyingTradingSessionSubIDField) Tag() quickfix.Tag {
	return tag.UnderlyingTradingSessionSubID
}

//NewUnderlyingTradingSessionSubID returns a new UnderlyingTradingSessionSubIDField initialized with val
func NewUnderlyingTradingSessionSubID(val string) UnderlyingTradingSessionSubIDField {
	return UnderlyingTradingSessionSubIDField{quickfix.FIXString(val)}
}

func (f UnderlyingTradingSessionSubIDField) Value() string { return f.String() }

//UnderlyingUnitOfMeasureField is a STRING field
type UnderlyingUnitOfMeasureField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingUnitOfMeasure (998)
func (f UnderlyingUnitOfMeasureField) Tag() quickfix.Tag { return tag.UnderlyingUnitOfMeasure }

//NewUnderlyingUnitOfMeasure returns a new UnderlyingUnitOfMeasureField initialized with val
func NewUnderlyingUnitOfMeasure(val string) UnderlyingUnitOfMeasureField {
	return UnderlyingUnitOfMeasureField{quickfix.FIXString(val)}
}

func (f UnderlyingUnitOfMeasureField) Value() string { return f.String() }

//UnderlyingUnitOfMeasureQtyField is a QTY field
type UnderlyingUnitOfMeasureQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnderlyingUnitOfMeasureQty (1423)
func (f UnderlyingUnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.UnderlyingUnitOfMeasureQty }

//NewUnderlyingUnitOfMeasureQty returns a new UnderlyingUnitOfMeasureQtyField initialized with val and scale
func NewUnderlyingUnitOfMeasureQty(val decimal.Decimal, scale int32) UnderlyingUnitOfMeasureQtyField {
	return UnderlyingUnitOfMeasureQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnderlyingUnitOfMeasureQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//UndlyInstrumentPartyIDField is a STRING field
type UndlyInstrumentPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.UndlyInstrumentPartyID (1059)
func (f UndlyInstrumentPartyIDField) Tag() quickfix.Tag { return tag.UndlyInstrumentPartyID }

//NewUndlyInstrumentPartyID returns a new UndlyInstrumentPartyIDField initialized with val
func NewUndlyInstrumentPartyID(val string) UndlyInstrumentPartyIDField {
	return UndlyInstrumentPartyIDField{quickfix.FIXString(val)}
}

func (f UndlyInstrumentPartyIDField) Value() string { return f.String() }

//UndlyInstrumentPartyIDSourceField is a CHAR field
type UndlyInstrumentPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.UndlyInstrumentPartyIDSource (1060)
func (f UndlyInstrumentPartyIDSourceField) Tag() quickfix.Tag { return tag.UndlyInstrumentPartyIDSource }

//NewUndlyInstrumentPartyIDSource returns a new UndlyInstrumentPartyIDSourceField initialized with val
func NewUndlyInstrumentPartyIDSource(val string) UndlyInstrumentPartyIDSourceField {
	return UndlyInstrumentPartyIDSourceField{quickfix.FIXString(val)}
}

func (f UndlyInstrumentPartyIDSourceField) Value() string { return f.String() }

//UndlyInstrumentPartyRoleField is a INT field
type UndlyInstrumentPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.UndlyInstrumentPartyRole (1061)
func (f UndlyInstrumentPartyRoleField) Tag() quickfix.Tag { return tag.UndlyInstrumentPartyRole }

//NewUndlyInstrumentPartyRole returns a new UndlyInstrumentPartyRoleField initialized with val
func NewUndlyInstrumentPartyRole(val int) UndlyInstrumentPartyRoleField {
	return UndlyInstrumentPartyRoleField{quickfix.FIXInt(val)}
}

func (f UndlyInstrumentPartyRoleField) Value() int { return f.Int() }

//UndlyInstrumentPartySubIDField is a STRING field
type UndlyInstrumentPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.UndlyInstrumentPartySubID (1063)
func (f UndlyInstrumentPartySubIDField) Tag() quickfix.Tag { return tag.UndlyInstrumentPartySubID }

//NewUndlyInstrumentPartySubID returns a new UndlyInstrumentPartySubIDField initialized with val
func NewUndlyInstrumentPartySubID(val string) UndlyInstrumentPartySubIDField {
	return UndlyInstrumentPartySubIDField{quickfix.FIXString(val)}
}

func (f UndlyInstrumentPartySubIDField) Value() string { return f.String() }

//UndlyInstrumentPartySubIDTypeField is a INT field
type UndlyInstrumentPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.UndlyInstrumentPartySubIDType (1064)
func (f UndlyInstrumentPartySubIDTypeField) Tag() quickfix.Tag {
	return tag.UndlyInstrumentPartySubIDType
}

//NewUndlyInstrumentPartySubIDType returns a new UndlyInstrumentPartySubIDTypeField initialized with val
func NewUndlyInstrumentPartySubIDType(val int) UndlyInstrumentPartySubIDTypeField {
	return UndlyInstrumentPartySubIDTypeField{quickfix.FIXInt(val)}
}

func (f UndlyInstrumentPartySubIDTypeField) Value() int { return f.Int() }

//UnitOfMeasureField is a enum.UnitOfMeasure field
type UnitOfMeasureField struct{ quickfix.FIXString }

//Tag returns tag.UnitOfMeasure (996)
func (f UnitOfMeasureField) Tag() quickfix.Tag { return tag.UnitOfMeasure }

func NewUnitOfMeasure(val enum.UnitOfMeasure) UnitOfMeasureField {
	return UnitOfMeasureField{quickfix.FIXString(val)}
}

func (f UnitOfMeasureField) Value() enum.UnitOfMeasure { return enum.UnitOfMeasure(f.String()) }

//UnitOfMeasureQtyField is a QTY field
type UnitOfMeasureQtyField struct{ quickfix.FIXDecimal }

//Tag returns tag.UnitOfMeasureQty (1147)
func (f UnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.UnitOfMeasureQty }

//NewUnitOfMeasureQty returns a new UnitOfMeasureQtyField initialized with val and scale
func NewUnitOfMeasureQty(val decimal.Decimal, scale int32) UnitOfMeasureQtyField {
	return UnitOfMeasureQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f UnitOfMeasureQtyField) Value() (val decimal.Decimal) { return f.Decimal }

//UnsolicitedIndicatorField is a BOOLEAN field
type UnsolicitedIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.UnsolicitedIndicator (325)
func (f UnsolicitedIndicatorField) Tag() quickfix.Tag { return tag.UnsolicitedIndicator }

//NewUnsolicitedIndicator returns a new UnsolicitedIndicatorField initialized with val
func NewUnsolicitedIndicator(val bool) UnsolicitedIndicatorField {
	return UnsolicitedIndicatorField{quickfix.FIXBoolean(val)}
}

func (f UnsolicitedIndicatorField) Value() bool { return f.Bool() }

//UrgencyField is a enum.Urgency field
type UrgencyField struct{ quickfix.FIXString }

//Tag returns tag.Urgency (61)
func (f UrgencyField) Tag() quickfix.Tag { return tag.Urgency }

func NewUrgency(val enum.Urgency) UrgencyField {
	return UrgencyField{quickfix.FIXString(val)}
}

func (f UrgencyField) Value() enum.Urgency { return enum.Urgency(f.String()) }

//UserRequestIDField is a STRING field
type UserRequestIDField struct{ quickfix.FIXString }

//Tag returns tag.UserRequestID (923)
func (f UserRequestIDField) Tag() quickfix.Tag { return tag.UserRequestID }

//NewUserRequestID returns a new UserRequestIDField initialized with val
func NewUserRequestID(val string) UserRequestIDField {
	return UserRequestIDField{quickfix.FIXString(val)}
}

func (f UserRequestIDField) Value() string { return f.String() }

//UserRequestTypeField is a enum.UserRequestType field
type UserRequestTypeField struct{ quickfix.FIXString }

//Tag returns tag.UserRequestType (924)
func (f UserRequestTypeField) Tag() quickfix.Tag { return tag.UserRequestType }

func NewUserRequestType(val enum.UserRequestType) UserRequestTypeField {
	return UserRequestTypeField{quickfix.FIXString(val)}
}

func (f UserRequestTypeField) Value() enum.UserRequestType { return enum.UserRequestType(f.String()) }

//UserStatusField is a enum.UserStatus field
type UserStatusField struct{ quickfix.FIXString }

//Tag returns tag.UserStatus (926)
func (f UserStatusField) Tag() quickfix.Tag { return tag.UserStatus }

func NewUserStatus(val enum.UserStatus) UserStatusField {
	return UserStatusField{quickfix.FIXString(val)}
}

func (f UserStatusField) Value() enum.UserStatus { return enum.UserStatus(f.String()) }

//UserStatusTextField is a STRING field
type UserStatusTextField struct{ quickfix.FIXString }

//Tag returns tag.UserStatusText (927)
func (f UserStatusTextField) Tag() quickfix.Tag { return tag.UserStatusText }

//NewUserStatusText returns a new UserStatusTextField initialized with val
func NewUserStatusText(val string) UserStatusTextField {
	return UserStatusTextField{quickfix.FIXString(val)}
}

func (f UserStatusTextField) Value() string { return f.String() }

//UsernameField is a STRING field
type UsernameField struct{ quickfix.FIXString }

//Tag returns tag.Username (553)
func (f UsernameField) Tag() quickfix.Tag { return tag.Username }

//NewUsername returns a new UsernameField initialized with val
func NewUsername(val string) UsernameField {
	return UsernameField{quickfix.FIXString(val)}
}

func (f UsernameField) Value() string { return f.String() }

//ValidUntilTimeField is a UTCTIMESTAMP field
type ValidUntilTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.ValidUntilTime (62)
func (f ValidUntilTimeField) Tag() quickfix.Tag { return tag.ValidUntilTime }

//NewValidUntilTime returns a new ValidUntilTimeField initialized with val
func NewValidUntilTime(val time.Time) ValidUntilTimeField {
	return ValidUntilTimeField{quickfix.FIXUTCTimestamp{Time: val}}
}

//NewValidUntilTimeNoMillis returns a new ValidUntilTimeField initialized with val without millisecs
func NewValidUntilTimeNoMillis(val time.Time) ValidUntilTimeField {
	return ValidUntilTimeField{quickfix.FIXUTCTimestamp{Time: val, NoMillis: true}}
}

func (f ValidUntilTimeField) Value() time.Time { return f.Time }

//ValuationMethodField is a enum.ValuationMethod field
type ValuationMethodField struct{ quickfix.FIXString }

//Tag returns tag.ValuationMethod (1197)
func (f ValuationMethodField) Tag() quickfix.Tag { return tag.ValuationMethod }

func NewValuationMethod(val enum.ValuationMethod) ValuationMethodField {
	return ValuationMethodField{quickfix.FIXString(val)}
}

func (f ValuationMethodField) Value() enum.ValuationMethod { return enum.ValuationMethod(f.String()) }

//ValueOfFuturesField is a AMT field
type ValueOfFuturesField struct{ quickfix.FIXDecimal }

//Tag returns tag.ValueOfFutures (408)
func (f ValueOfFuturesField) Tag() quickfix.Tag { return tag.ValueOfFutures }

//NewValueOfFutures returns a new ValueOfFuturesField initialized with val and scale
func NewValueOfFutures(val decimal.Decimal, scale int32) ValueOfFuturesField {
	return ValueOfFuturesField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f ValueOfFuturesField) Value() (val decimal.Decimal) { return f.Decimal }

//VenueTypeField is a enum.VenueType field
type VenueTypeField struct{ quickfix.FIXString }

//Tag returns tag.VenueType (1430)
func (f VenueTypeField) Tag() quickfix.Tag { return tag.VenueType }

func NewVenueType(val enum.VenueType) VenueTypeField {
	return VenueTypeField{quickfix.FIXString(val)}
}

func (f VenueTypeField) Value() enum.VenueType { return enum.VenueType(f.String()) }

//VolatilityField is a FLOAT field
type VolatilityField struct{ quickfix.FIXDecimal }

//Tag returns tag.Volatility (1188)
func (f VolatilityField) Tag() quickfix.Tag { return tag.Volatility }

//NewVolatility returns a new VolatilityField initialized with val and scale
func NewVolatility(val decimal.Decimal, scale int32) VolatilityField {
	return VolatilityField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f VolatilityField) Value() (val decimal.Decimal) { return f.Decimal }

//WaveNoField is a STRING field
type WaveNoField struct{ quickfix.FIXString }

//Tag returns tag.WaveNo (105)
func (f WaveNoField) Tag() quickfix.Tag { return tag.WaveNo }

//NewWaveNo returns a new WaveNoField initialized with val
func NewWaveNo(val string) WaveNoField {
	return WaveNoField{quickfix.FIXString(val)}
}

func (f WaveNoField) Value() string { return f.String() }

//WorkingIndicatorField is a BOOLEAN field
type WorkingIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.WorkingIndicator (636)
func (f WorkingIndicatorField) Tag() quickfix.Tag { return tag.WorkingIndicator }

//NewWorkingIndicator returns a new WorkingIndicatorField initialized with val
func NewWorkingIndicator(val bool) WorkingIndicatorField {
	return WorkingIndicatorField{quickfix.FIXBoolean(val)}
}

func (f WorkingIndicatorField) Value() bool { return f.Bool() }

//WtAverageLiquidityField is a PERCENTAGE field
type WtAverageLiquidityField struct{ quickfix.FIXDecimal }

//Tag returns tag.WtAverageLiquidity (410)
func (f WtAverageLiquidityField) Tag() quickfix.Tag { return tag.WtAverageLiquidity }

//NewWtAverageLiquidity returns a new WtAverageLiquidityField initialized with val and scale
func NewWtAverageLiquidity(val decimal.Decimal, scale int32) WtAverageLiquidityField {
	return WtAverageLiquidityField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f WtAverageLiquidityField) Value() (val decimal.Decimal) { return f.Decimal }

//XmlDataField is a DATA field
type XmlDataField struct{ quickfix.FIXString }

//Tag returns tag.XmlData (213)
func (f XmlDataField) Tag() quickfix.Tag { return tag.XmlData }

//NewXmlData returns a new XmlDataField initialized with val
func NewXmlData(val string) XmlDataField {
	return XmlDataField{quickfix.FIXString(val)}
}

func (f XmlDataField) Value() string { return f.String() }

//XmlDataLenField is a LENGTH field
type XmlDataLenField struct{ quickfix.FIXInt }

//Tag returns tag.XmlDataLen (212)
func (f XmlDataLenField) Tag() quickfix.Tag { return tag.XmlDataLen }

//NewXmlDataLen returns a new XmlDataLenField initialized with val
func NewXmlDataLen(val int) XmlDataLenField {
	return XmlDataLenField{quickfix.FIXInt(val)}
}

func (f XmlDataLenField) Value() int { return f.Int() }

//YieldField is a PERCENTAGE field
type YieldField struct{ quickfix.FIXDecimal }

//Tag returns tag.Yield (236)
func (f YieldField) Tag() quickfix.Tag { return tag.Yield }

//NewYield returns a new YieldField initialized with val and scale
func NewYield(val decimal.Decimal, scale int32) YieldField {
	return YieldField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f YieldField) Value() (val decimal.Decimal) { return f.Decimal }

//YieldCalcDateField is a LOCALMKTDATE field
type YieldCalcDateField struct{ quickfix.FIXString }

//Tag returns tag.YieldCalcDate (701)
func (f YieldCalcDateField) Tag() quickfix.Tag { return tag.YieldCalcDate }

//NewYieldCalcDate returns a new YieldCalcDateField initialized with val
func NewYieldCalcDate(val string) YieldCalcDateField {
	return YieldCalcDateField{quickfix.FIXString(val)}
}

func (f YieldCalcDateField) Value() string { return f.String() }

//YieldRedemptionDateField is a LOCALMKTDATE field
type YieldRedemptionDateField struct{ quickfix.FIXString }

//Tag returns tag.YieldRedemptionDate (696)
func (f YieldRedemptionDateField) Tag() quickfix.Tag { return tag.YieldRedemptionDate }

//NewYieldRedemptionDate returns a new YieldRedemptionDateField initialized with val
func NewYieldRedemptionDate(val string) YieldRedemptionDateField {
	return YieldRedemptionDateField{quickfix.FIXString(val)}
}

func (f YieldRedemptionDateField) Value() string { return f.String() }

//YieldRedemptionPriceField is a PRICE field
type YieldRedemptionPriceField struct{ quickfix.FIXDecimal }

//Tag returns tag.YieldRedemptionPrice (697)
func (f YieldRedemptionPriceField) Tag() quickfix.Tag { return tag.YieldRedemptionPrice }

//NewYieldRedemptionPrice returns a new YieldRedemptionPriceField initialized with val and scale
func NewYieldRedemptionPrice(val decimal.Decimal, scale int32) YieldRedemptionPriceField {
	return YieldRedemptionPriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f YieldRedemptionPriceField) Value() (val decimal.Decimal) { return f.Decimal }

//YieldRedemptionPriceTypeField is a INT field
type YieldRedemptionPriceTypeField struct{ quickfix.FIXInt }

//Tag returns tag.YieldRedemptionPriceType (698)
func (f YieldRedemptionPriceTypeField) Tag() quickfix.Tag { return tag.YieldRedemptionPriceType }

//NewYieldRedemptionPriceType returns a new YieldRedemptionPriceTypeField initialized with val
func NewYieldRedemptionPriceType(val int) YieldRedemptionPriceTypeField {
	return YieldRedemptionPriceTypeField{quickfix.FIXInt(val)}
}

func (f YieldRedemptionPriceTypeField) Value() int { return f.Int() }

//YieldTypeField is a enum.YieldType field
type YieldTypeField struct{ quickfix.FIXString }

//Tag returns tag.YieldType (235)
func (f YieldTypeField) Tag() quickfix.Tag { return tag.YieldType }

func NewYieldType(val enum.YieldType) YieldTypeField {
	return YieldTypeField{quickfix.FIXString(val)}
}

func (f YieldTypeField) Value() enum.YieldType { return enum.YieldType(f.String()) }
