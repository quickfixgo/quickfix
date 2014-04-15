package field

import (
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/tag"
	"github.com/quickfixgo/quickfix/message"
)

//Account is a STRING field
type Account struct{ message.StringValue }

//Tag returns tag.Account (1)
func (f Account) Tag() fix.Tag { return tag.Account }

//BuildAccount returns a new Account initialized with val
func BuildAccount(val string) Account {
	var field Account
	field.Value = val
	return field
}

//AccountType is a INT field
type AccountType struct{ message.IntValue }

//Tag returns tag.AccountType (581)
func (f AccountType) Tag() fix.Tag { return tag.AccountType }

//BuildAccountType returns a new AccountType initialized with val
func BuildAccountType(val int) AccountType {
	var field AccountType
	field.Value = val
	return field
}

//AccruedInterestAmt is a AMT field
type AccruedInterestAmt struct{ message.AmtValue }

//Tag returns tag.AccruedInterestAmt (159)
func (f AccruedInterestAmt) Tag() fix.Tag { return tag.AccruedInterestAmt }

//BuildAccruedInterestAmt returns a new AccruedInterestAmt initialized with val
func BuildAccruedInterestAmt(val float64) AccruedInterestAmt {
	var field AccruedInterestAmt
	field.Value = val
	return field
}

//AccruedInterestRate is a PERCENTAGE field
type AccruedInterestRate struct{ message.PercentageValue }

//Tag returns tag.AccruedInterestRate (158)
func (f AccruedInterestRate) Tag() fix.Tag { return tag.AccruedInterestRate }

//BuildAccruedInterestRate returns a new AccruedInterestRate initialized with val
func BuildAccruedInterestRate(val float64) AccruedInterestRate {
	var field AccruedInterestRate
	field.Value = val
	return field
}

//AcctIDSource is a INT field
type AcctIDSource struct{ message.IntValue }

//Tag returns tag.AcctIDSource (660)
func (f AcctIDSource) Tag() fix.Tag { return tag.AcctIDSource }

//BuildAcctIDSource returns a new AcctIDSource initialized with val
func BuildAcctIDSource(val int) AcctIDSource {
	var field AcctIDSource
	field.Value = val
	return field
}

//Adjustment is a INT field
type Adjustment struct{ message.IntValue }

//Tag returns tag.Adjustment (334)
func (f Adjustment) Tag() fix.Tag { return tag.Adjustment }

//BuildAdjustment returns a new Adjustment initialized with val
func BuildAdjustment(val int) Adjustment {
	var field Adjustment
	field.Value = val
	return field
}

//AdjustmentType is a INT field
type AdjustmentType struct{ message.IntValue }

//Tag returns tag.AdjustmentType (718)
func (f AdjustmentType) Tag() fix.Tag { return tag.AdjustmentType }

//BuildAdjustmentType returns a new AdjustmentType initialized with val
func BuildAdjustmentType(val int) AdjustmentType {
	var field AdjustmentType
	field.Value = val
	return field
}

//AdvId is a STRING field
type AdvId struct{ message.StringValue }

//Tag returns tag.AdvId (2)
func (f AdvId) Tag() fix.Tag { return tag.AdvId }

//BuildAdvId returns a new AdvId initialized with val
func BuildAdvId(val string) AdvId {
	var field AdvId
	field.Value = val
	return field
}

//AdvRefID is a STRING field
type AdvRefID struct{ message.StringValue }

//Tag returns tag.AdvRefID (3)
func (f AdvRefID) Tag() fix.Tag { return tag.AdvRefID }

//BuildAdvRefID returns a new AdvRefID initialized with val
func BuildAdvRefID(val string) AdvRefID {
	var field AdvRefID
	field.Value = val
	return field
}

//AdvSide is a CHAR field
type AdvSide struct{ message.CharValue }

//Tag returns tag.AdvSide (4)
func (f AdvSide) Tag() fix.Tag { return tag.AdvSide }

//BuildAdvSide returns a new AdvSide initialized with val
func BuildAdvSide(val string) AdvSide {
	var field AdvSide
	field.Value = val
	return field
}

//AdvTransType is a STRING field
type AdvTransType struct{ message.StringValue }

//Tag returns tag.AdvTransType (5)
func (f AdvTransType) Tag() fix.Tag { return tag.AdvTransType }

//BuildAdvTransType returns a new AdvTransType initialized with val
func BuildAdvTransType(val string) AdvTransType {
	var field AdvTransType
	field.Value = val
	return field
}

//AffectedOrderID is a STRING field
type AffectedOrderID struct{ message.StringValue }

//Tag returns tag.AffectedOrderID (535)
func (f AffectedOrderID) Tag() fix.Tag { return tag.AffectedOrderID }

//BuildAffectedOrderID returns a new AffectedOrderID initialized with val
func BuildAffectedOrderID(val string) AffectedOrderID {
	var field AffectedOrderID
	field.Value = val
	return field
}

//AffectedSecondaryOrderID is a STRING field
type AffectedSecondaryOrderID struct{ message.StringValue }

//Tag returns tag.AffectedSecondaryOrderID (536)
func (f AffectedSecondaryOrderID) Tag() fix.Tag { return tag.AffectedSecondaryOrderID }

//BuildAffectedSecondaryOrderID returns a new AffectedSecondaryOrderID initialized with val
func BuildAffectedSecondaryOrderID(val string) AffectedSecondaryOrderID {
	var field AffectedSecondaryOrderID
	field.Value = val
	return field
}

//AffirmStatus is a INT field
type AffirmStatus struct{ message.IntValue }

//Tag returns tag.AffirmStatus (940)
func (f AffirmStatus) Tag() fix.Tag { return tag.AffirmStatus }

//BuildAffirmStatus returns a new AffirmStatus initialized with val
func BuildAffirmStatus(val int) AffirmStatus {
	var field AffirmStatus
	field.Value = val
	return field
}

//AggregatedBook is a BOOLEAN field
type AggregatedBook struct{ message.BooleanValue }

//Tag returns tag.AggregatedBook (266)
func (f AggregatedBook) Tag() fix.Tag { return tag.AggregatedBook }

//BuildAggregatedBook returns a new AggregatedBook initialized with val
func BuildAggregatedBook(val bool) AggregatedBook {
	var field AggregatedBook
	field.Value = val
	return field
}

//AggressorIndicator is a BOOLEAN field
type AggressorIndicator struct{ message.BooleanValue }

//Tag returns tag.AggressorIndicator (1057)
func (f AggressorIndicator) Tag() fix.Tag { return tag.AggressorIndicator }

//BuildAggressorIndicator returns a new AggressorIndicator initialized with val
func BuildAggressorIndicator(val bool) AggressorIndicator {
	var field AggressorIndicator
	field.Value = val
	return field
}

//AgreementCurrency is a CURRENCY field
type AgreementCurrency struct{ message.CurrencyValue }

//Tag returns tag.AgreementCurrency (918)
func (f AgreementCurrency) Tag() fix.Tag { return tag.AgreementCurrency }

//BuildAgreementCurrency returns a new AgreementCurrency initialized with val
func BuildAgreementCurrency(val string) AgreementCurrency {
	var field AgreementCurrency
	field.Value = val
	return field
}

//AgreementDate is a LOCALMKTDATE field
type AgreementDate struct{ message.LocalMktDateValue }

//Tag returns tag.AgreementDate (915)
func (f AgreementDate) Tag() fix.Tag { return tag.AgreementDate }

//BuildAgreementDate returns a new AgreementDate initialized with val
func BuildAgreementDate(val string) AgreementDate {
	var field AgreementDate
	field.Value = val
	return field
}

//AgreementDesc is a STRING field
type AgreementDesc struct{ message.StringValue }

//Tag returns tag.AgreementDesc (913)
func (f AgreementDesc) Tag() fix.Tag { return tag.AgreementDesc }

//BuildAgreementDesc returns a new AgreementDesc initialized with val
func BuildAgreementDesc(val string) AgreementDesc {
	var field AgreementDesc
	field.Value = val
	return field
}

//AgreementID is a STRING field
type AgreementID struct{ message.StringValue }

//Tag returns tag.AgreementID (914)
func (f AgreementID) Tag() fix.Tag { return tag.AgreementID }

//BuildAgreementID returns a new AgreementID initialized with val
func BuildAgreementID(val string) AgreementID {
	var field AgreementID
	field.Value = val
	return field
}

//AllocAccount is a STRING field
type AllocAccount struct{ message.StringValue }

//Tag returns tag.AllocAccount (79)
func (f AllocAccount) Tag() fix.Tag { return tag.AllocAccount }

//BuildAllocAccount returns a new AllocAccount initialized with val
func BuildAllocAccount(val string) AllocAccount {
	var field AllocAccount
	field.Value = val
	return field
}

//AllocAccountType is a INT field
type AllocAccountType struct{ message.IntValue }

//Tag returns tag.AllocAccountType (798)
func (f AllocAccountType) Tag() fix.Tag { return tag.AllocAccountType }

//BuildAllocAccountType returns a new AllocAccountType initialized with val
func BuildAllocAccountType(val int) AllocAccountType {
	var field AllocAccountType
	field.Value = val
	return field
}

//AllocAccruedInterestAmt is a AMT field
type AllocAccruedInterestAmt struct{ message.AmtValue }

//Tag returns tag.AllocAccruedInterestAmt (742)
func (f AllocAccruedInterestAmt) Tag() fix.Tag { return tag.AllocAccruedInterestAmt }

//BuildAllocAccruedInterestAmt returns a new AllocAccruedInterestAmt initialized with val
func BuildAllocAccruedInterestAmt(val float64) AllocAccruedInterestAmt {
	var field AllocAccruedInterestAmt
	field.Value = val
	return field
}

//AllocAcctIDSource is a INT field
type AllocAcctIDSource struct{ message.IntValue }

//Tag returns tag.AllocAcctIDSource (661)
func (f AllocAcctIDSource) Tag() fix.Tag { return tag.AllocAcctIDSource }

//BuildAllocAcctIDSource returns a new AllocAcctIDSource initialized with val
func BuildAllocAcctIDSource(val int) AllocAcctIDSource {
	var field AllocAcctIDSource
	field.Value = val
	return field
}

//AllocAvgPx is a PRICE field
type AllocAvgPx struct{ message.PriceValue }

//Tag returns tag.AllocAvgPx (153)
func (f AllocAvgPx) Tag() fix.Tag { return tag.AllocAvgPx }

//BuildAllocAvgPx returns a new AllocAvgPx initialized with val
func BuildAllocAvgPx(val float64) AllocAvgPx {
	var field AllocAvgPx
	field.Value = val
	return field
}

//AllocCancReplaceReason is a INT field
type AllocCancReplaceReason struct{ message.IntValue }

//Tag returns tag.AllocCancReplaceReason (796)
func (f AllocCancReplaceReason) Tag() fix.Tag { return tag.AllocCancReplaceReason }

//BuildAllocCancReplaceReason returns a new AllocCancReplaceReason initialized with val
func BuildAllocCancReplaceReason(val int) AllocCancReplaceReason {
	var field AllocCancReplaceReason
	field.Value = val
	return field
}

//AllocClearingFeeIndicator is a STRING field
type AllocClearingFeeIndicator struct{ message.StringValue }

//Tag returns tag.AllocClearingFeeIndicator (1136)
func (f AllocClearingFeeIndicator) Tag() fix.Tag { return tag.AllocClearingFeeIndicator }

//BuildAllocClearingFeeIndicator returns a new AllocClearingFeeIndicator initialized with val
func BuildAllocClearingFeeIndicator(val string) AllocClearingFeeIndicator {
	var field AllocClearingFeeIndicator
	field.Value = val
	return field
}

//AllocCustomerCapacity is a STRING field
type AllocCustomerCapacity struct{ message.StringValue }

//Tag returns tag.AllocCustomerCapacity (993)
func (f AllocCustomerCapacity) Tag() fix.Tag { return tag.AllocCustomerCapacity }

//BuildAllocCustomerCapacity returns a new AllocCustomerCapacity initialized with val
func BuildAllocCustomerCapacity(val string) AllocCustomerCapacity {
	var field AllocCustomerCapacity
	field.Value = val
	return field
}

//AllocHandlInst is a INT field
type AllocHandlInst struct{ message.IntValue }

//Tag returns tag.AllocHandlInst (209)
func (f AllocHandlInst) Tag() fix.Tag { return tag.AllocHandlInst }

//BuildAllocHandlInst returns a new AllocHandlInst initialized with val
func BuildAllocHandlInst(val int) AllocHandlInst {
	var field AllocHandlInst
	field.Value = val
	return field
}

//AllocID is a STRING field
type AllocID struct{ message.StringValue }

//Tag returns tag.AllocID (70)
func (f AllocID) Tag() fix.Tag { return tag.AllocID }

//BuildAllocID returns a new AllocID initialized with val
func BuildAllocID(val string) AllocID {
	var field AllocID
	field.Value = val
	return field
}

//AllocInterestAtMaturity is a AMT field
type AllocInterestAtMaturity struct{ message.AmtValue }

//Tag returns tag.AllocInterestAtMaturity (741)
func (f AllocInterestAtMaturity) Tag() fix.Tag { return tag.AllocInterestAtMaturity }

//BuildAllocInterestAtMaturity returns a new AllocInterestAtMaturity initialized with val
func BuildAllocInterestAtMaturity(val float64) AllocInterestAtMaturity {
	var field AllocInterestAtMaturity
	field.Value = val
	return field
}

//AllocIntermedReqType is a INT field
type AllocIntermedReqType struct{ message.IntValue }

//Tag returns tag.AllocIntermedReqType (808)
func (f AllocIntermedReqType) Tag() fix.Tag { return tag.AllocIntermedReqType }

//BuildAllocIntermedReqType returns a new AllocIntermedReqType initialized with val
func BuildAllocIntermedReqType(val int) AllocIntermedReqType {
	var field AllocIntermedReqType
	field.Value = val
	return field
}

//AllocLinkID is a STRING field
type AllocLinkID struct{ message.StringValue }

//Tag returns tag.AllocLinkID (196)
func (f AllocLinkID) Tag() fix.Tag { return tag.AllocLinkID }

//BuildAllocLinkID returns a new AllocLinkID initialized with val
func BuildAllocLinkID(val string) AllocLinkID {
	var field AllocLinkID
	field.Value = val
	return field
}

//AllocLinkType is a INT field
type AllocLinkType struct{ message.IntValue }

//Tag returns tag.AllocLinkType (197)
func (f AllocLinkType) Tag() fix.Tag { return tag.AllocLinkType }

//BuildAllocLinkType returns a new AllocLinkType initialized with val
func BuildAllocLinkType(val int) AllocLinkType {
	var field AllocLinkType
	field.Value = val
	return field
}

//AllocMethod is a INT field
type AllocMethod struct{ message.IntValue }

//Tag returns tag.AllocMethod (1002)
func (f AllocMethod) Tag() fix.Tag { return tag.AllocMethod }

//BuildAllocMethod returns a new AllocMethod initialized with val
func BuildAllocMethod(val int) AllocMethod {
	var field AllocMethod
	field.Value = val
	return field
}

//AllocNetMoney is a AMT field
type AllocNetMoney struct{ message.AmtValue }

//Tag returns tag.AllocNetMoney (154)
func (f AllocNetMoney) Tag() fix.Tag { return tag.AllocNetMoney }

//BuildAllocNetMoney returns a new AllocNetMoney initialized with val
func BuildAllocNetMoney(val float64) AllocNetMoney {
	var field AllocNetMoney
	field.Value = val
	return field
}

//AllocNoOrdersType is a INT field
type AllocNoOrdersType struct{ message.IntValue }

//Tag returns tag.AllocNoOrdersType (857)
func (f AllocNoOrdersType) Tag() fix.Tag { return tag.AllocNoOrdersType }

//BuildAllocNoOrdersType returns a new AllocNoOrdersType initialized with val
func BuildAllocNoOrdersType(val int) AllocNoOrdersType {
	var field AllocNoOrdersType
	field.Value = val
	return field
}

//AllocPositionEffect is a CHAR field
type AllocPositionEffect struct{ message.CharValue }

//Tag returns tag.AllocPositionEffect (1047)
func (f AllocPositionEffect) Tag() fix.Tag { return tag.AllocPositionEffect }

//BuildAllocPositionEffect returns a new AllocPositionEffect initialized with val
func BuildAllocPositionEffect(val string) AllocPositionEffect {
	var field AllocPositionEffect
	field.Value = val
	return field
}

//AllocPrice is a PRICE field
type AllocPrice struct{ message.PriceValue }

//Tag returns tag.AllocPrice (366)
func (f AllocPrice) Tag() fix.Tag { return tag.AllocPrice }

//BuildAllocPrice returns a new AllocPrice initialized with val
func BuildAllocPrice(val float64) AllocPrice {
	var field AllocPrice
	field.Value = val
	return field
}

//AllocQty is a QTY field
type AllocQty struct{ message.QtyValue }

//Tag returns tag.AllocQty (80)
func (f AllocQty) Tag() fix.Tag { return tag.AllocQty }

//BuildAllocQty returns a new AllocQty initialized with val
func BuildAllocQty(val float64) AllocQty {
	var field AllocQty
	field.Value = val
	return field
}

//AllocRejCode is a INT field
type AllocRejCode struct{ message.IntValue }

//Tag returns tag.AllocRejCode (88)
func (f AllocRejCode) Tag() fix.Tag { return tag.AllocRejCode }

//BuildAllocRejCode returns a new AllocRejCode initialized with val
func BuildAllocRejCode(val int) AllocRejCode {
	var field AllocRejCode
	field.Value = val
	return field
}

//AllocReportID is a STRING field
type AllocReportID struct{ message.StringValue }

//Tag returns tag.AllocReportID (755)
func (f AllocReportID) Tag() fix.Tag { return tag.AllocReportID }

//BuildAllocReportID returns a new AllocReportID initialized with val
func BuildAllocReportID(val string) AllocReportID {
	var field AllocReportID
	field.Value = val
	return field
}

//AllocReportRefID is a STRING field
type AllocReportRefID struct{ message.StringValue }

//Tag returns tag.AllocReportRefID (795)
func (f AllocReportRefID) Tag() fix.Tag { return tag.AllocReportRefID }

//BuildAllocReportRefID returns a new AllocReportRefID initialized with val
func BuildAllocReportRefID(val string) AllocReportRefID {
	var field AllocReportRefID
	field.Value = val
	return field
}

//AllocReportType is a INT field
type AllocReportType struct{ message.IntValue }

//Tag returns tag.AllocReportType (794)
func (f AllocReportType) Tag() fix.Tag { return tag.AllocReportType }

//BuildAllocReportType returns a new AllocReportType initialized with val
func BuildAllocReportType(val int) AllocReportType {
	var field AllocReportType
	field.Value = val
	return field
}

//AllocSettlCurrAmt is a AMT field
type AllocSettlCurrAmt struct{ message.AmtValue }

//Tag returns tag.AllocSettlCurrAmt (737)
func (f AllocSettlCurrAmt) Tag() fix.Tag { return tag.AllocSettlCurrAmt }

//BuildAllocSettlCurrAmt returns a new AllocSettlCurrAmt initialized with val
func BuildAllocSettlCurrAmt(val float64) AllocSettlCurrAmt {
	var field AllocSettlCurrAmt
	field.Value = val
	return field
}

//AllocSettlCurrency is a CURRENCY field
type AllocSettlCurrency struct{ message.CurrencyValue }

//Tag returns tag.AllocSettlCurrency (736)
func (f AllocSettlCurrency) Tag() fix.Tag { return tag.AllocSettlCurrency }

//BuildAllocSettlCurrency returns a new AllocSettlCurrency initialized with val
func BuildAllocSettlCurrency(val string) AllocSettlCurrency {
	var field AllocSettlCurrency
	field.Value = val
	return field
}

//AllocSettlInstType is a INT field
type AllocSettlInstType struct{ message.IntValue }

//Tag returns tag.AllocSettlInstType (780)
func (f AllocSettlInstType) Tag() fix.Tag { return tag.AllocSettlInstType }

//BuildAllocSettlInstType returns a new AllocSettlInstType initialized with val
func BuildAllocSettlInstType(val int) AllocSettlInstType {
	var field AllocSettlInstType
	field.Value = val
	return field
}

//AllocShares is a QTY field
type AllocShares struct{ message.QtyValue }

//Tag returns tag.AllocShares (80)
func (f AllocShares) Tag() fix.Tag { return tag.AllocShares }

//BuildAllocShares returns a new AllocShares initialized with val
func BuildAllocShares(val float64) AllocShares {
	var field AllocShares
	field.Value = val
	return field
}

//AllocStatus is a INT field
type AllocStatus struct{ message.IntValue }

//Tag returns tag.AllocStatus (87)
func (f AllocStatus) Tag() fix.Tag { return tag.AllocStatus }

//BuildAllocStatus returns a new AllocStatus initialized with val
func BuildAllocStatus(val int) AllocStatus {
	var field AllocStatus
	field.Value = val
	return field
}

//AllocText is a STRING field
type AllocText struct{ message.StringValue }

//Tag returns tag.AllocText (161)
func (f AllocText) Tag() fix.Tag { return tag.AllocText }

//BuildAllocText returns a new AllocText initialized with val
func BuildAllocText(val string) AllocText {
	var field AllocText
	field.Value = val
	return field
}

//AllocTransType is a CHAR field
type AllocTransType struct{ message.CharValue }

//Tag returns tag.AllocTransType (71)
func (f AllocTransType) Tag() fix.Tag { return tag.AllocTransType }

//BuildAllocTransType returns a new AllocTransType initialized with val
func BuildAllocTransType(val string) AllocTransType {
	var field AllocTransType
	field.Value = val
	return field
}

//AllocType is a INT field
type AllocType struct{ message.IntValue }

//Tag returns tag.AllocType (626)
func (f AllocType) Tag() fix.Tag { return tag.AllocType }

//BuildAllocType returns a new AllocType initialized with val
func BuildAllocType(val int) AllocType {
	var field AllocType
	field.Value = val
	return field
}

//AllowableOneSidednessCurr is a CURRENCY field
type AllowableOneSidednessCurr struct{ message.CurrencyValue }

//Tag returns tag.AllowableOneSidednessCurr (767)
func (f AllowableOneSidednessCurr) Tag() fix.Tag { return tag.AllowableOneSidednessCurr }

//BuildAllowableOneSidednessCurr returns a new AllowableOneSidednessCurr initialized with val
func BuildAllowableOneSidednessCurr(val string) AllowableOneSidednessCurr {
	var field AllowableOneSidednessCurr
	field.Value = val
	return field
}

//AllowableOneSidednessPct is a PERCENTAGE field
type AllowableOneSidednessPct struct{ message.PercentageValue }

//Tag returns tag.AllowableOneSidednessPct (765)
func (f AllowableOneSidednessPct) Tag() fix.Tag { return tag.AllowableOneSidednessPct }

//BuildAllowableOneSidednessPct returns a new AllowableOneSidednessPct initialized with val
func BuildAllowableOneSidednessPct(val float64) AllowableOneSidednessPct {
	var field AllowableOneSidednessPct
	field.Value = val
	return field
}

//AllowableOneSidednessValue is a AMT field
type AllowableOneSidednessValue struct{ message.AmtValue }

//Tag returns tag.AllowableOneSidednessValue (766)
func (f AllowableOneSidednessValue) Tag() fix.Tag { return tag.AllowableOneSidednessValue }

//BuildAllowableOneSidednessValue returns a new AllowableOneSidednessValue initialized with val
func BuildAllowableOneSidednessValue(val float64) AllowableOneSidednessValue {
	var field AllowableOneSidednessValue
	field.Value = val
	return field
}

//AltMDSourceID is a STRING field
type AltMDSourceID struct{ message.StringValue }

//Tag returns tag.AltMDSourceID (817)
func (f AltMDSourceID) Tag() fix.Tag { return tag.AltMDSourceID }

//BuildAltMDSourceID returns a new AltMDSourceID initialized with val
func BuildAltMDSourceID(val string) AltMDSourceID {
	var field AltMDSourceID
	field.Value = val
	return field
}

//ApplBegSeqNum is a SEQNUM field
type ApplBegSeqNum struct{ message.SeqNumValue }

//Tag returns tag.ApplBegSeqNum (1182)
func (f ApplBegSeqNum) Tag() fix.Tag { return tag.ApplBegSeqNum }

//BuildApplBegSeqNum returns a new ApplBegSeqNum initialized with val
func BuildApplBegSeqNum(val int) ApplBegSeqNum {
	var field ApplBegSeqNum
	field.Value = val
	return field
}

//ApplEndSeqNum is a SEQNUM field
type ApplEndSeqNum struct{ message.SeqNumValue }

//Tag returns tag.ApplEndSeqNum (1183)
func (f ApplEndSeqNum) Tag() fix.Tag { return tag.ApplEndSeqNum }

//BuildApplEndSeqNum returns a new ApplEndSeqNum initialized with val
func BuildApplEndSeqNum(val int) ApplEndSeqNum {
	var field ApplEndSeqNum
	field.Value = val
	return field
}

//ApplExtID is a INT field
type ApplExtID struct{ message.IntValue }

//Tag returns tag.ApplExtID (1156)
func (f ApplExtID) Tag() fix.Tag { return tag.ApplExtID }

//BuildApplExtID returns a new ApplExtID initialized with val
func BuildApplExtID(val int) ApplExtID {
	var field ApplExtID
	field.Value = val
	return field
}

//ApplID is a STRING field
type ApplID struct{ message.StringValue }

//Tag returns tag.ApplID (1180)
func (f ApplID) Tag() fix.Tag { return tag.ApplID }

//BuildApplID returns a new ApplID initialized with val
func BuildApplID(val string) ApplID {
	var field ApplID
	field.Value = val
	return field
}

//ApplLastSeqNum is a SEQNUM field
type ApplLastSeqNum struct{ message.SeqNumValue }

//Tag returns tag.ApplLastSeqNum (1350)
func (f ApplLastSeqNum) Tag() fix.Tag { return tag.ApplLastSeqNum }

//BuildApplLastSeqNum returns a new ApplLastSeqNum initialized with val
func BuildApplLastSeqNum(val int) ApplLastSeqNum {
	var field ApplLastSeqNum
	field.Value = val
	return field
}

//ApplNewSeqNum is a SEQNUM field
type ApplNewSeqNum struct{ message.SeqNumValue }

//Tag returns tag.ApplNewSeqNum (1399)
func (f ApplNewSeqNum) Tag() fix.Tag { return tag.ApplNewSeqNum }

//BuildApplNewSeqNum returns a new ApplNewSeqNum initialized with val
func BuildApplNewSeqNum(val int) ApplNewSeqNum {
	var field ApplNewSeqNum
	field.Value = val
	return field
}

//ApplQueueAction is a INT field
type ApplQueueAction struct{ message.IntValue }

//Tag returns tag.ApplQueueAction (815)
func (f ApplQueueAction) Tag() fix.Tag { return tag.ApplQueueAction }

//BuildApplQueueAction returns a new ApplQueueAction initialized with val
func BuildApplQueueAction(val int) ApplQueueAction {
	var field ApplQueueAction
	field.Value = val
	return field
}

//ApplQueueDepth is a INT field
type ApplQueueDepth struct{ message.IntValue }

//Tag returns tag.ApplQueueDepth (813)
func (f ApplQueueDepth) Tag() fix.Tag { return tag.ApplQueueDepth }

//BuildApplQueueDepth returns a new ApplQueueDepth initialized with val
func BuildApplQueueDepth(val int) ApplQueueDepth {
	var field ApplQueueDepth
	field.Value = val
	return field
}

//ApplQueueMax is a INT field
type ApplQueueMax struct{ message.IntValue }

//Tag returns tag.ApplQueueMax (812)
func (f ApplQueueMax) Tag() fix.Tag { return tag.ApplQueueMax }

//BuildApplQueueMax returns a new ApplQueueMax initialized with val
func BuildApplQueueMax(val int) ApplQueueMax {
	var field ApplQueueMax
	field.Value = val
	return field
}

//ApplQueueResolution is a INT field
type ApplQueueResolution struct{ message.IntValue }

//Tag returns tag.ApplQueueResolution (814)
func (f ApplQueueResolution) Tag() fix.Tag { return tag.ApplQueueResolution }

//BuildApplQueueResolution returns a new ApplQueueResolution initialized with val
func BuildApplQueueResolution(val int) ApplQueueResolution {
	var field ApplQueueResolution
	field.Value = val
	return field
}

//ApplReportID is a STRING field
type ApplReportID struct{ message.StringValue }

//Tag returns tag.ApplReportID (1356)
func (f ApplReportID) Tag() fix.Tag { return tag.ApplReportID }

//BuildApplReportID returns a new ApplReportID initialized with val
func BuildApplReportID(val string) ApplReportID {
	var field ApplReportID
	field.Value = val
	return field
}

//ApplReportType is a INT field
type ApplReportType struct{ message.IntValue }

//Tag returns tag.ApplReportType (1426)
func (f ApplReportType) Tag() fix.Tag { return tag.ApplReportType }

//BuildApplReportType returns a new ApplReportType initialized with val
func BuildApplReportType(val int) ApplReportType {
	var field ApplReportType
	field.Value = val
	return field
}

//ApplReqID is a STRING field
type ApplReqID struct{ message.StringValue }

//Tag returns tag.ApplReqID (1346)
func (f ApplReqID) Tag() fix.Tag { return tag.ApplReqID }

//BuildApplReqID returns a new ApplReqID initialized with val
func BuildApplReqID(val string) ApplReqID {
	var field ApplReqID
	field.Value = val
	return field
}

//ApplReqType is a INT field
type ApplReqType struct{ message.IntValue }

//Tag returns tag.ApplReqType (1347)
func (f ApplReqType) Tag() fix.Tag { return tag.ApplReqType }

//BuildApplReqType returns a new ApplReqType initialized with val
func BuildApplReqType(val int) ApplReqType {
	var field ApplReqType
	field.Value = val
	return field
}

//ApplResendFlag is a BOOLEAN field
type ApplResendFlag struct{ message.BooleanValue }

//Tag returns tag.ApplResendFlag (1352)
func (f ApplResendFlag) Tag() fix.Tag { return tag.ApplResendFlag }

//BuildApplResendFlag returns a new ApplResendFlag initialized with val
func BuildApplResendFlag(val bool) ApplResendFlag {
	var field ApplResendFlag
	field.Value = val
	return field
}

//ApplResponseError is a INT field
type ApplResponseError struct{ message.IntValue }

//Tag returns tag.ApplResponseError (1354)
func (f ApplResponseError) Tag() fix.Tag { return tag.ApplResponseError }

//BuildApplResponseError returns a new ApplResponseError initialized with val
func BuildApplResponseError(val int) ApplResponseError {
	var field ApplResponseError
	field.Value = val
	return field
}

//ApplResponseID is a STRING field
type ApplResponseID struct{ message.StringValue }

//Tag returns tag.ApplResponseID (1353)
func (f ApplResponseID) Tag() fix.Tag { return tag.ApplResponseID }

//BuildApplResponseID returns a new ApplResponseID initialized with val
func BuildApplResponseID(val string) ApplResponseID {
	var field ApplResponseID
	field.Value = val
	return field
}

//ApplResponseType is a INT field
type ApplResponseType struct{ message.IntValue }

//Tag returns tag.ApplResponseType (1348)
func (f ApplResponseType) Tag() fix.Tag { return tag.ApplResponseType }

//BuildApplResponseType returns a new ApplResponseType initialized with val
func BuildApplResponseType(val int) ApplResponseType {
	var field ApplResponseType
	field.Value = val
	return field
}

//ApplSeqNum is a SEQNUM field
type ApplSeqNum struct{ message.SeqNumValue }

//Tag returns tag.ApplSeqNum (1181)
func (f ApplSeqNum) Tag() fix.Tag { return tag.ApplSeqNum }

//BuildApplSeqNum returns a new ApplSeqNum initialized with val
func BuildApplSeqNum(val int) ApplSeqNum {
	var field ApplSeqNum
	field.Value = val
	return field
}

//ApplTotalMessageCount is a INT field
type ApplTotalMessageCount struct{ message.IntValue }

//Tag returns tag.ApplTotalMessageCount (1349)
func (f ApplTotalMessageCount) Tag() fix.Tag { return tag.ApplTotalMessageCount }

//BuildApplTotalMessageCount returns a new ApplTotalMessageCount initialized with val
func BuildApplTotalMessageCount(val int) ApplTotalMessageCount {
	var field ApplTotalMessageCount
	field.Value = val
	return field
}

//ApplVerID is a STRING field
type ApplVerID struct{ message.StringValue }

//Tag returns tag.ApplVerID (1128)
func (f ApplVerID) Tag() fix.Tag { return tag.ApplVerID }

//BuildApplVerID returns a new ApplVerID initialized with val
func BuildApplVerID(val string) ApplVerID {
	var field ApplVerID
	field.Value = val
	return field
}

//AsOfIndicator is a CHAR field
type AsOfIndicator struct{ message.CharValue }

//Tag returns tag.AsOfIndicator (1015)
func (f AsOfIndicator) Tag() fix.Tag { return tag.AsOfIndicator }

//BuildAsOfIndicator returns a new AsOfIndicator initialized with val
func BuildAsOfIndicator(val string) AsOfIndicator {
	var field AsOfIndicator
	field.Value = val
	return field
}

//AsgnReqID is a STRING field
type AsgnReqID struct{ message.StringValue }

//Tag returns tag.AsgnReqID (831)
func (f AsgnReqID) Tag() fix.Tag { return tag.AsgnReqID }

//BuildAsgnReqID returns a new AsgnReqID initialized with val
func BuildAsgnReqID(val string) AsgnReqID {
	var field AsgnReqID
	field.Value = val
	return field
}

//AsgnRptID is a STRING field
type AsgnRptID struct{ message.StringValue }

//Tag returns tag.AsgnRptID (833)
func (f AsgnRptID) Tag() fix.Tag { return tag.AsgnRptID }

//BuildAsgnRptID returns a new AsgnRptID initialized with val
func BuildAsgnRptID(val string) AsgnRptID {
	var field AsgnRptID
	field.Value = val
	return field
}

//AssignmentMethod is a CHAR field
type AssignmentMethod struct{ message.CharValue }

//Tag returns tag.AssignmentMethod (744)
func (f AssignmentMethod) Tag() fix.Tag { return tag.AssignmentMethod }

//BuildAssignmentMethod returns a new AssignmentMethod initialized with val
func BuildAssignmentMethod(val string) AssignmentMethod {
	var field AssignmentMethod
	field.Value = val
	return field
}

//AssignmentUnit is a QTY field
type AssignmentUnit struct{ message.QtyValue }

//Tag returns tag.AssignmentUnit (745)
func (f AssignmentUnit) Tag() fix.Tag { return tag.AssignmentUnit }

//BuildAssignmentUnit returns a new AssignmentUnit initialized with val
func BuildAssignmentUnit(val float64) AssignmentUnit {
	var field AssignmentUnit
	field.Value = val
	return field
}

//AttachmentPoint is a PERCENTAGE field
type AttachmentPoint struct{ message.PercentageValue }

//Tag returns tag.AttachmentPoint (1457)
func (f AttachmentPoint) Tag() fix.Tag { return tag.AttachmentPoint }

//BuildAttachmentPoint returns a new AttachmentPoint initialized with val
func BuildAttachmentPoint(val float64) AttachmentPoint {
	var field AttachmentPoint
	field.Value = val
	return field
}

//AutoAcceptIndicator is a BOOLEAN field
type AutoAcceptIndicator struct{ message.BooleanValue }

//Tag returns tag.AutoAcceptIndicator (754)
func (f AutoAcceptIndicator) Tag() fix.Tag { return tag.AutoAcceptIndicator }

//BuildAutoAcceptIndicator returns a new AutoAcceptIndicator initialized with val
func BuildAutoAcceptIndicator(val bool) AutoAcceptIndicator {
	var field AutoAcceptIndicator
	field.Value = val
	return field
}

//AvgParPx is a PRICE field
type AvgParPx struct{ message.PriceValue }

//Tag returns tag.AvgParPx (860)
func (f AvgParPx) Tag() fix.Tag { return tag.AvgParPx }

//BuildAvgParPx returns a new AvgParPx initialized with val
func BuildAvgParPx(val float64) AvgParPx {
	var field AvgParPx
	field.Value = val
	return field
}

//AvgPrxPrecision is a INT field
type AvgPrxPrecision struct{ message.IntValue }

//Tag returns tag.AvgPrxPrecision (74)
func (f AvgPrxPrecision) Tag() fix.Tag { return tag.AvgPrxPrecision }

//BuildAvgPrxPrecision returns a new AvgPrxPrecision initialized with val
func BuildAvgPrxPrecision(val int) AvgPrxPrecision {
	var field AvgPrxPrecision
	field.Value = val
	return field
}

//AvgPx is a PRICE field
type AvgPx struct{ message.PriceValue }

//Tag returns tag.AvgPx (6)
func (f AvgPx) Tag() fix.Tag { return tag.AvgPx }

//BuildAvgPx returns a new AvgPx initialized with val
func BuildAvgPx(val float64) AvgPx {
	var field AvgPx
	field.Value = val
	return field
}

//AvgPxIndicator is a INT field
type AvgPxIndicator struct{ message.IntValue }

//Tag returns tag.AvgPxIndicator (819)
func (f AvgPxIndicator) Tag() fix.Tag { return tag.AvgPxIndicator }

//BuildAvgPxIndicator returns a new AvgPxIndicator initialized with val
func BuildAvgPxIndicator(val int) AvgPxIndicator {
	var field AvgPxIndicator
	field.Value = val
	return field
}

//AvgPxPrecision is a INT field
type AvgPxPrecision struct{ message.IntValue }

//Tag returns tag.AvgPxPrecision (74)
func (f AvgPxPrecision) Tag() fix.Tag { return tag.AvgPxPrecision }

//BuildAvgPxPrecision returns a new AvgPxPrecision initialized with val
func BuildAvgPxPrecision(val int) AvgPxPrecision {
	var field AvgPxPrecision
	field.Value = val
	return field
}

//BasisFeatureDate is a LOCALMKTDATE field
type BasisFeatureDate struct{ message.LocalMktDateValue }

//Tag returns tag.BasisFeatureDate (259)
func (f BasisFeatureDate) Tag() fix.Tag { return tag.BasisFeatureDate }

//BuildBasisFeatureDate returns a new BasisFeatureDate initialized with val
func BuildBasisFeatureDate(val string) BasisFeatureDate {
	var field BasisFeatureDate
	field.Value = val
	return field
}

//BasisFeaturePrice is a PRICE field
type BasisFeaturePrice struct{ message.PriceValue }

//Tag returns tag.BasisFeaturePrice (260)
func (f BasisFeaturePrice) Tag() fix.Tag { return tag.BasisFeaturePrice }

//BuildBasisFeaturePrice returns a new BasisFeaturePrice initialized with val
func BuildBasisFeaturePrice(val float64) BasisFeaturePrice {
	var field BasisFeaturePrice
	field.Value = val
	return field
}

//BasisPxType is a CHAR field
type BasisPxType struct{ message.CharValue }

//Tag returns tag.BasisPxType (419)
func (f BasisPxType) Tag() fix.Tag { return tag.BasisPxType }

//BuildBasisPxType returns a new BasisPxType initialized with val
func BuildBasisPxType(val string) BasisPxType {
	var field BasisPxType
	field.Value = val
	return field
}

//BeginSeqNo is a SEQNUM field
type BeginSeqNo struct{ message.SeqNumValue }

//Tag returns tag.BeginSeqNo (7)
func (f BeginSeqNo) Tag() fix.Tag { return tag.BeginSeqNo }

//BuildBeginSeqNo returns a new BeginSeqNo initialized with val
func BuildBeginSeqNo(val int) BeginSeqNo {
	var field BeginSeqNo
	field.Value = val
	return field
}

//BeginString is a STRING field
type BeginString struct{ message.StringValue }

//Tag returns tag.BeginString (8)
func (f BeginString) Tag() fix.Tag { return tag.BeginString }

//BuildBeginString returns a new BeginString initialized with val
func BuildBeginString(val string) BeginString {
	var field BeginString
	field.Value = val
	return field
}

//Benchmark is a CHAR field
type Benchmark struct{ message.CharValue }

//Tag returns tag.Benchmark (219)
func (f Benchmark) Tag() fix.Tag { return tag.Benchmark }

//BuildBenchmark returns a new Benchmark initialized with val
func BuildBenchmark(val string) Benchmark {
	var field Benchmark
	field.Value = val
	return field
}

//BenchmarkCurveCurrency is a CURRENCY field
type BenchmarkCurveCurrency struct{ message.CurrencyValue }

//Tag returns tag.BenchmarkCurveCurrency (220)
func (f BenchmarkCurveCurrency) Tag() fix.Tag { return tag.BenchmarkCurveCurrency }

//BuildBenchmarkCurveCurrency returns a new BenchmarkCurveCurrency initialized with val
func BuildBenchmarkCurveCurrency(val string) BenchmarkCurveCurrency {
	var field BenchmarkCurveCurrency
	field.Value = val
	return field
}

//BenchmarkCurveName is a STRING field
type BenchmarkCurveName struct{ message.StringValue }

//Tag returns tag.BenchmarkCurveName (221)
func (f BenchmarkCurveName) Tag() fix.Tag { return tag.BenchmarkCurveName }

//BuildBenchmarkCurveName returns a new BenchmarkCurveName initialized with val
func BuildBenchmarkCurveName(val string) BenchmarkCurveName {
	var field BenchmarkCurveName
	field.Value = val
	return field
}

//BenchmarkCurvePoint is a STRING field
type BenchmarkCurvePoint struct{ message.StringValue }

//Tag returns tag.BenchmarkCurvePoint (222)
func (f BenchmarkCurvePoint) Tag() fix.Tag { return tag.BenchmarkCurvePoint }

//BuildBenchmarkCurvePoint returns a new BenchmarkCurvePoint initialized with val
func BuildBenchmarkCurvePoint(val string) BenchmarkCurvePoint {
	var field BenchmarkCurvePoint
	field.Value = val
	return field
}

//BenchmarkPrice is a PRICE field
type BenchmarkPrice struct{ message.PriceValue }

//Tag returns tag.BenchmarkPrice (662)
func (f BenchmarkPrice) Tag() fix.Tag { return tag.BenchmarkPrice }

//BuildBenchmarkPrice returns a new BenchmarkPrice initialized with val
func BuildBenchmarkPrice(val float64) BenchmarkPrice {
	var field BenchmarkPrice
	field.Value = val
	return field
}

//BenchmarkPriceType is a INT field
type BenchmarkPriceType struct{ message.IntValue }

//Tag returns tag.BenchmarkPriceType (663)
func (f BenchmarkPriceType) Tag() fix.Tag { return tag.BenchmarkPriceType }

//BuildBenchmarkPriceType returns a new BenchmarkPriceType initialized with val
func BuildBenchmarkPriceType(val int) BenchmarkPriceType {
	var field BenchmarkPriceType
	field.Value = val
	return field
}

//BenchmarkSecurityID is a STRING field
type BenchmarkSecurityID struct{ message.StringValue }

//Tag returns tag.BenchmarkSecurityID (699)
func (f BenchmarkSecurityID) Tag() fix.Tag { return tag.BenchmarkSecurityID }

//BuildBenchmarkSecurityID returns a new BenchmarkSecurityID initialized with val
func BuildBenchmarkSecurityID(val string) BenchmarkSecurityID {
	var field BenchmarkSecurityID
	field.Value = val
	return field
}

//BenchmarkSecurityIDSource is a STRING field
type BenchmarkSecurityIDSource struct{ message.StringValue }

//Tag returns tag.BenchmarkSecurityIDSource (761)
func (f BenchmarkSecurityIDSource) Tag() fix.Tag { return tag.BenchmarkSecurityIDSource }

//BuildBenchmarkSecurityIDSource returns a new BenchmarkSecurityIDSource initialized with val
func BuildBenchmarkSecurityIDSource(val string) BenchmarkSecurityIDSource {
	var field BenchmarkSecurityIDSource
	field.Value = val
	return field
}

//BidDescriptor is a STRING field
type BidDescriptor struct{ message.StringValue }

//Tag returns tag.BidDescriptor (400)
func (f BidDescriptor) Tag() fix.Tag { return tag.BidDescriptor }

//BuildBidDescriptor returns a new BidDescriptor initialized with val
func BuildBidDescriptor(val string) BidDescriptor {
	var field BidDescriptor
	field.Value = val
	return field
}

//BidDescriptorType is a INT field
type BidDescriptorType struct{ message.IntValue }

//Tag returns tag.BidDescriptorType (399)
func (f BidDescriptorType) Tag() fix.Tag { return tag.BidDescriptorType }

//BuildBidDescriptorType returns a new BidDescriptorType initialized with val
func BuildBidDescriptorType(val int) BidDescriptorType {
	var field BidDescriptorType
	field.Value = val
	return field
}

//BidForwardPoints is a PRICEOFFSET field
type BidForwardPoints struct{ message.PriceOffsetValue }

//Tag returns tag.BidForwardPoints (189)
func (f BidForwardPoints) Tag() fix.Tag { return tag.BidForwardPoints }

//BuildBidForwardPoints returns a new BidForwardPoints initialized with val
func BuildBidForwardPoints(val float64) BidForwardPoints {
	var field BidForwardPoints
	field.Value = val
	return field
}

//BidForwardPoints2 is a PRICEOFFSET field
type BidForwardPoints2 struct{ message.PriceOffsetValue }

//Tag returns tag.BidForwardPoints2 (642)
func (f BidForwardPoints2) Tag() fix.Tag { return tag.BidForwardPoints2 }

//BuildBidForwardPoints2 returns a new BidForwardPoints2 initialized with val
func BuildBidForwardPoints2(val float64) BidForwardPoints2 {
	var field BidForwardPoints2
	field.Value = val
	return field
}

//BidID is a STRING field
type BidID struct{ message.StringValue }

//Tag returns tag.BidID (390)
func (f BidID) Tag() fix.Tag { return tag.BidID }

//BuildBidID returns a new BidID initialized with val
func BuildBidID(val string) BidID {
	var field BidID
	field.Value = val
	return field
}

//BidPx is a PRICE field
type BidPx struct{ message.PriceValue }

//Tag returns tag.BidPx (132)
func (f BidPx) Tag() fix.Tag { return tag.BidPx }

//BuildBidPx returns a new BidPx initialized with val
func BuildBidPx(val float64) BidPx {
	var field BidPx
	field.Value = val
	return field
}

//BidRequestTransType is a CHAR field
type BidRequestTransType struct{ message.CharValue }

//Tag returns tag.BidRequestTransType (374)
func (f BidRequestTransType) Tag() fix.Tag { return tag.BidRequestTransType }

//BuildBidRequestTransType returns a new BidRequestTransType initialized with val
func BuildBidRequestTransType(val string) BidRequestTransType {
	var field BidRequestTransType
	field.Value = val
	return field
}

//BidSize is a QTY field
type BidSize struct{ message.QtyValue }

//Tag returns tag.BidSize (134)
func (f BidSize) Tag() fix.Tag { return tag.BidSize }

//BuildBidSize returns a new BidSize initialized with val
func BuildBidSize(val float64) BidSize {
	var field BidSize
	field.Value = val
	return field
}

//BidSpotRate is a PRICE field
type BidSpotRate struct{ message.PriceValue }

//Tag returns tag.BidSpotRate (188)
func (f BidSpotRate) Tag() fix.Tag { return tag.BidSpotRate }

//BuildBidSpotRate returns a new BidSpotRate initialized with val
func BuildBidSpotRate(val float64) BidSpotRate {
	var field BidSpotRate
	field.Value = val
	return field
}

//BidSwapPoints is a PRICEOFFSET field
type BidSwapPoints struct{ message.PriceOffsetValue }

//Tag returns tag.BidSwapPoints (1065)
func (f BidSwapPoints) Tag() fix.Tag { return tag.BidSwapPoints }

//BuildBidSwapPoints returns a new BidSwapPoints initialized with val
func BuildBidSwapPoints(val float64) BidSwapPoints {
	var field BidSwapPoints
	field.Value = val
	return field
}

//BidTradeType is a CHAR field
type BidTradeType struct{ message.CharValue }

//Tag returns tag.BidTradeType (418)
func (f BidTradeType) Tag() fix.Tag { return tag.BidTradeType }

//BuildBidTradeType returns a new BidTradeType initialized with val
func BuildBidTradeType(val string) BidTradeType {
	var field BidTradeType
	field.Value = val
	return field
}

//BidType is a INT field
type BidType struct{ message.IntValue }

//Tag returns tag.BidType (394)
func (f BidType) Tag() fix.Tag { return tag.BidType }

//BuildBidType returns a new BidType initialized with val
func BuildBidType(val int) BidType {
	var field BidType
	field.Value = val
	return field
}

//BidYield is a PERCENTAGE field
type BidYield struct{ message.PercentageValue }

//Tag returns tag.BidYield (632)
func (f BidYield) Tag() fix.Tag { return tag.BidYield }

//BuildBidYield returns a new BidYield initialized with val
func BuildBidYield(val float64) BidYield {
	var field BidYield
	field.Value = val
	return field
}

//BodyLength is a LENGTH field
type BodyLength struct{ message.LengthValue }

//Tag returns tag.BodyLength (9)
func (f BodyLength) Tag() fix.Tag { return tag.BodyLength }

//BuildBodyLength returns a new BodyLength initialized with val
func BuildBodyLength(val int) BodyLength {
	var field BodyLength
	field.Value = val
	return field
}

//BookingRefID is a STRING field
type BookingRefID struct{ message.StringValue }

//Tag returns tag.BookingRefID (466)
func (f BookingRefID) Tag() fix.Tag { return tag.BookingRefID }

//BuildBookingRefID returns a new BookingRefID initialized with val
func BuildBookingRefID(val string) BookingRefID {
	var field BookingRefID
	field.Value = val
	return field
}

//BookingType is a INT field
type BookingType struct{ message.IntValue }

//Tag returns tag.BookingType (775)
func (f BookingType) Tag() fix.Tag { return tag.BookingType }

//BuildBookingType returns a new BookingType initialized with val
func BuildBookingType(val int) BookingType {
	var field BookingType
	field.Value = val
	return field
}

//BookingUnit is a CHAR field
type BookingUnit struct{ message.CharValue }

//Tag returns tag.BookingUnit (590)
func (f BookingUnit) Tag() fix.Tag { return tag.BookingUnit }

//BuildBookingUnit returns a new BookingUnit initialized with val
func BuildBookingUnit(val string) BookingUnit {
	var field BookingUnit
	field.Value = val
	return field
}

//BrokerOfCredit is a STRING field
type BrokerOfCredit struct{ message.StringValue }

//Tag returns tag.BrokerOfCredit (92)
func (f BrokerOfCredit) Tag() fix.Tag { return tag.BrokerOfCredit }

//BuildBrokerOfCredit returns a new BrokerOfCredit initialized with val
func BuildBrokerOfCredit(val string) BrokerOfCredit {
	var field BrokerOfCredit
	field.Value = val
	return field
}

//BusinessRejectReason is a INT field
type BusinessRejectReason struct{ message.IntValue }

//Tag returns tag.BusinessRejectReason (380)
func (f BusinessRejectReason) Tag() fix.Tag { return tag.BusinessRejectReason }

//BuildBusinessRejectReason returns a new BusinessRejectReason initialized with val
func BuildBusinessRejectReason(val int) BusinessRejectReason {
	var field BusinessRejectReason
	field.Value = val
	return field
}

//BusinessRejectRefID is a STRING field
type BusinessRejectRefID struct{ message.StringValue }

//Tag returns tag.BusinessRejectRefID (379)
func (f BusinessRejectRefID) Tag() fix.Tag { return tag.BusinessRejectRefID }

//BuildBusinessRejectRefID returns a new BusinessRejectRefID initialized with val
func BuildBusinessRejectRefID(val string) BusinessRejectRefID {
	var field BusinessRejectRefID
	field.Value = val
	return field
}

//BuyVolume is a QTY field
type BuyVolume struct{ message.QtyValue }

//Tag returns tag.BuyVolume (330)
func (f BuyVolume) Tag() fix.Tag { return tag.BuyVolume }

//BuildBuyVolume returns a new BuyVolume initialized with val
func BuildBuyVolume(val float64) BuyVolume {
	var field BuyVolume
	field.Value = val
	return field
}

//CFICode is a STRING field
type CFICode struct{ message.StringValue }

//Tag returns tag.CFICode (461)
func (f CFICode) Tag() fix.Tag { return tag.CFICode }

//BuildCFICode returns a new CFICode initialized with val
func BuildCFICode(val string) CFICode {
	var field CFICode
	field.Value = val
	return field
}

//CPProgram is a INT field
type CPProgram struct{ message.IntValue }

//Tag returns tag.CPProgram (875)
func (f CPProgram) Tag() fix.Tag { return tag.CPProgram }

//BuildCPProgram returns a new CPProgram initialized with val
func BuildCPProgram(val int) CPProgram {
	var field CPProgram
	field.Value = val
	return field
}

//CPRegType is a STRING field
type CPRegType struct{ message.StringValue }

//Tag returns tag.CPRegType (876)
func (f CPRegType) Tag() fix.Tag { return tag.CPRegType }

//BuildCPRegType returns a new CPRegType initialized with val
func BuildCPRegType(val string) CPRegType {
	var field CPRegType
	field.Value = val
	return field
}

//CalculatedCcyLastQty is a QTY field
type CalculatedCcyLastQty struct{ message.QtyValue }

//Tag returns tag.CalculatedCcyLastQty (1056)
func (f CalculatedCcyLastQty) Tag() fix.Tag { return tag.CalculatedCcyLastQty }

//BuildCalculatedCcyLastQty returns a new CalculatedCcyLastQty initialized with val
func BuildCalculatedCcyLastQty(val float64) CalculatedCcyLastQty {
	var field CalculatedCcyLastQty
	field.Value = val
	return field
}

//CancellationRights is a CHAR field
type CancellationRights struct{ message.CharValue }

//Tag returns tag.CancellationRights (480)
func (f CancellationRights) Tag() fix.Tag { return tag.CancellationRights }

//BuildCancellationRights returns a new CancellationRights initialized with val
func BuildCancellationRights(val string) CancellationRights {
	var field CancellationRights
	field.Value = val
	return field
}

//CapPrice is a PRICE field
type CapPrice struct{ message.PriceValue }

//Tag returns tag.CapPrice (1199)
func (f CapPrice) Tag() fix.Tag { return tag.CapPrice }

//BuildCapPrice returns a new CapPrice initialized with val
func BuildCapPrice(val float64) CapPrice {
	var field CapPrice
	field.Value = val
	return field
}

//CardExpDate is a LOCALMKTDATE field
type CardExpDate struct{ message.LocalMktDateValue }

//Tag returns tag.CardExpDate (490)
func (f CardExpDate) Tag() fix.Tag { return tag.CardExpDate }

//BuildCardExpDate returns a new CardExpDate initialized with val
func BuildCardExpDate(val string) CardExpDate {
	var field CardExpDate
	field.Value = val
	return field
}

//CardHolderName is a STRING field
type CardHolderName struct{ message.StringValue }

//Tag returns tag.CardHolderName (488)
func (f CardHolderName) Tag() fix.Tag { return tag.CardHolderName }

//BuildCardHolderName returns a new CardHolderName initialized with val
func BuildCardHolderName(val string) CardHolderName {
	var field CardHolderName
	field.Value = val
	return field
}

//CardIssNo is a STRING field
type CardIssNo struct{ message.StringValue }

//Tag returns tag.CardIssNo (491)
func (f CardIssNo) Tag() fix.Tag { return tag.CardIssNo }

//BuildCardIssNo returns a new CardIssNo initialized with val
func BuildCardIssNo(val string) CardIssNo {
	var field CardIssNo
	field.Value = val
	return field
}

//CardIssNum is a STRING field
type CardIssNum struct{ message.StringValue }

//Tag returns tag.CardIssNum (491)
func (f CardIssNum) Tag() fix.Tag { return tag.CardIssNum }

//BuildCardIssNum returns a new CardIssNum initialized with val
func BuildCardIssNum(val string) CardIssNum {
	var field CardIssNum
	field.Value = val
	return field
}

//CardNumber is a STRING field
type CardNumber struct{ message.StringValue }

//Tag returns tag.CardNumber (489)
func (f CardNumber) Tag() fix.Tag { return tag.CardNumber }

//BuildCardNumber returns a new CardNumber initialized with val
func BuildCardNumber(val string) CardNumber {
	var field CardNumber
	field.Value = val
	return field
}

//CardStartDate is a LOCALMKTDATE field
type CardStartDate struct{ message.LocalMktDateValue }

//Tag returns tag.CardStartDate (503)
func (f CardStartDate) Tag() fix.Tag { return tag.CardStartDate }

//BuildCardStartDate returns a new CardStartDate initialized with val
func BuildCardStartDate(val string) CardStartDate {
	var field CardStartDate
	field.Value = val
	return field
}

//CashDistribAgentAcctName is a STRING field
type CashDistribAgentAcctName struct{ message.StringValue }

//Tag returns tag.CashDistribAgentAcctName (502)
func (f CashDistribAgentAcctName) Tag() fix.Tag { return tag.CashDistribAgentAcctName }

//BuildCashDistribAgentAcctName returns a new CashDistribAgentAcctName initialized with val
func BuildCashDistribAgentAcctName(val string) CashDistribAgentAcctName {
	var field CashDistribAgentAcctName
	field.Value = val
	return field
}

//CashDistribAgentAcctNumber is a STRING field
type CashDistribAgentAcctNumber struct{ message.StringValue }

//Tag returns tag.CashDistribAgentAcctNumber (500)
func (f CashDistribAgentAcctNumber) Tag() fix.Tag { return tag.CashDistribAgentAcctNumber }

//BuildCashDistribAgentAcctNumber returns a new CashDistribAgentAcctNumber initialized with val
func BuildCashDistribAgentAcctNumber(val string) CashDistribAgentAcctNumber {
	var field CashDistribAgentAcctNumber
	field.Value = val
	return field
}

//CashDistribAgentCode is a STRING field
type CashDistribAgentCode struct{ message.StringValue }

//Tag returns tag.CashDistribAgentCode (499)
func (f CashDistribAgentCode) Tag() fix.Tag { return tag.CashDistribAgentCode }

//BuildCashDistribAgentCode returns a new CashDistribAgentCode initialized with val
func BuildCashDistribAgentCode(val string) CashDistribAgentCode {
	var field CashDistribAgentCode
	field.Value = val
	return field
}

//CashDistribAgentName is a STRING field
type CashDistribAgentName struct{ message.StringValue }

//Tag returns tag.CashDistribAgentName (498)
func (f CashDistribAgentName) Tag() fix.Tag { return tag.CashDistribAgentName }

//BuildCashDistribAgentName returns a new CashDistribAgentName initialized with val
func BuildCashDistribAgentName(val string) CashDistribAgentName {
	var field CashDistribAgentName
	field.Value = val
	return field
}

//CashDistribCurr is a CURRENCY field
type CashDistribCurr struct{ message.CurrencyValue }

//Tag returns tag.CashDistribCurr (478)
func (f CashDistribCurr) Tag() fix.Tag { return tag.CashDistribCurr }

//BuildCashDistribCurr returns a new CashDistribCurr initialized with val
func BuildCashDistribCurr(val string) CashDistribCurr {
	var field CashDistribCurr
	field.Value = val
	return field
}

//CashDistribPayRef is a STRING field
type CashDistribPayRef struct{ message.StringValue }

//Tag returns tag.CashDistribPayRef (501)
func (f CashDistribPayRef) Tag() fix.Tag { return tag.CashDistribPayRef }

//BuildCashDistribPayRef returns a new CashDistribPayRef initialized with val
func BuildCashDistribPayRef(val string) CashDistribPayRef {
	var field CashDistribPayRef
	field.Value = val
	return field
}

//CashMargin is a CHAR field
type CashMargin struct{ message.CharValue }

//Tag returns tag.CashMargin (544)
func (f CashMargin) Tag() fix.Tag { return tag.CashMargin }

//BuildCashMargin returns a new CashMargin initialized with val
func BuildCashMargin(val string) CashMargin {
	var field CashMargin
	field.Value = val
	return field
}

//CashOrderQty is a QTY field
type CashOrderQty struct{ message.QtyValue }

//Tag returns tag.CashOrderQty (152)
func (f CashOrderQty) Tag() fix.Tag { return tag.CashOrderQty }

//BuildCashOrderQty returns a new CashOrderQty initialized with val
func BuildCashOrderQty(val float64) CashOrderQty {
	var field CashOrderQty
	field.Value = val
	return field
}

//CashOutstanding is a AMT field
type CashOutstanding struct{ message.AmtValue }

//Tag returns tag.CashOutstanding (901)
func (f CashOutstanding) Tag() fix.Tag { return tag.CashOutstanding }

//BuildCashOutstanding returns a new CashOutstanding initialized with val
func BuildCashOutstanding(val float64) CashOutstanding {
	var field CashOutstanding
	field.Value = val
	return field
}

//CashSettlAgentAcctName is a STRING field
type CashSettlAgentAcctName struct{ message.StringValue }

//Tag returns tag.CashSettlAgentAcctName (185)
func (f CashSettlAgentAcctName) Tag() fix.Tag { return tag.CashSettlAgentAcctName }

//BuildCashSettlAgentAcctName returns a new CashSettlAgentAcctName initialized with val
func BuildCashSettlAgentAcctName(val string) CashSettlAgentAcctName {
	var field CashSettlAgentAcctName
	field.Value = val
	return field
}

//CashSettlAgentAcctNum is a STRING field
type CashSettlAgentAcctNum struct{ message.StringValue }

//Tag returns tag.CashSettlAgentAcctNum (184)
func (f CashSettlAgentAcctNum) Tag() fix.Tag { return tag.CashSettlAgentAcctNum }

//BuildCashSettlAgentAcctNum returns a new CashSettlAgentAcctNum initialized with val
func BuildCashSettlAgentAcctNum(val string) CashSettlAgentAcctNum {
	var field CashSettlAgentAcctNum
	field.Value = val
	return field
}

//CashSettlAgentCode is a STRING field
type CashSettlAgentCode struct{ message.StringValue }

//Tag returns tag.CashSettlAgentCode (183)
func (f CashSettlAgentCode) Tag() fix.Tag { return tag.CashSettlAgentCode }

//BuildCashSettlAgentCode returns a new CashSettlAgentCode initialized with val
func BuildCashSettlAgentCode(val string) CashSettlAgentCode {
	var field CashSettlAgentCode
	field.Value = val
	return field
}

//CashSettlAgentContactName is a STRING field
type CashSettlAgentContactName struct{ message.StringValue }

//Tag returns tag.CashSettlAgentContactName (186)
func (f CashSettlAgentContactName) Tag() fix.Tag { return tag.CashSettlAgentContactName }

//BuildCashSettlAgentContactName returns a new CashSettlAgentContactName initialized with val
func BuildCashSettlAgentContactName(val string) CashSettlAgentContactName {
	var field CashSettlAgentContactName
	field.Value = val
	return field
}

//CashSettlAgentContactPhone is a STRING field
type CashSettlAgentContactPhone struct{ message.StringValue }

//Tag returns tag.CashSettlAgentContactPhone (187)
func (f CashSettlAgentContactPhone) Tag() fix.Tag { return tag.CashSettlAgentContactPhone }

//BuildCashSettlAgentContactPhone returns a new CashSettlAgentContactPhone initialized with val
func BuildCashSettlAgentContactPhone(val string) CashSettlAgentContactPhone {
	var field CashSettlAgentContactPhone
	field.Value = val
	return field
}

//CashSettlAgentName is a STRING field
type CashSettlAgentName struct{ message.StringValue }

//Tag returns tag.CashSettlAgentName (182)
func (f CashSettlAgentName) Tag() fix.Tag { return tag.CashSettlAgentName }

//BuildCashSettlAgentName returns a new CashSettlAgentName initialized with val
func BuildCashSettlAgentName(val string) CashSettlAgentName {
	var field CashSettlAgentName
	field.Value = val
	return field
}

//CcyAmt is a AMT field
type CcyAmt struct{ message.AmtValue }

//Tag returns tag.CcyAmt (1157)
func (f CcyAmt) Tag() fix.Tag { return tag.CcyAmt }

//BuildCcyAmt returns a new CcyAmt initialized with val
func BuildCcyAmt(val float64) CcyAmt {
	var field CcyAmt
	field.Value = val
	return field
}

//CheckSum is a STRING field
type CheckSum struct{ message.StringValue }

//Tag returns tag.CheckSum (10)
func (f CheckSum) Tag() fix.Tag { return tag.CheckSum }

//BuildCheckSum returns a new CheckSum initialized with val
func BuildCheckSum(val string) CheckSum {
	var field CheckSum
	field.Value = val
	return field
}

//ClOrdID is a STRING field
type ClOrdID struct{ message.StringValue }

//Tag returns tag.ClOrdID (11)
func (f ClOrdID) Tag() fix.Tag { return tag.ClOrdID }

//BuildClOrdID returns a new ClOrdID initialized with val
func BuildClOrdID(val string) ClOrdID {
	var field ClOrdID
	field.Value = val
	return field
}

//ClOrdLinkID is a STRING field
type ClOrdLinkID struct{ message.StringValue }

//Tag returns tag.ClOrdLinkID (583)
func (f ClOrdLinkID) Tag() fix.Tag { return tag.ClOrdLinkID }

//BuildClOrdLinkID returns a new ClOrdLinkID initialized with val
func BuildClOrdLinkID(val string) ClOrdLinkID {
	var field ClOrdLinkID
	field.Value = val
	return field
}

//ClearingAccount is a STRING field
type ClearingAccount struct{ message.StringValue }

//Tag returns tag.ClearingAccount (440)
func (f ClearingAccount) Tag() fix.Tag { return tag.ClearingAccount }

//BuildClearingAccount returns a new ClearingAccount initialized with val
func BuildClearingAccount(val string) ClearingAccount {
	var field ClearingAccount
	field.Value = val
	return field
}

//ClearingBusinessDate is a LOCALMKTDATE field
type ClearingBusinessDate struct{ message.LocalMktDateValue }

//Tag returns tag.ClearingBusinessDate (715)
func (f ClearingBusinessDate) Tag() fix.Tag { return tag.ClearingBusinessDate }

//BuildClearingBusinessDate returns a new ClearingBusinessDate initialized with val
func BuildClearingBusinessDate(val string) ClearingBusinessDate {
	var field ClearingBusinessDate
	field.Value = val
	return field
}

//ClearingFeeIndicator is a STRING field
type ClearingFeeIndicator struct{ message.StringValue }

//Tag returns tag.ClearingFeeIndicator (635)
func (f ClearingFeeIndicator) Tag() fix.Tag { return tag.ClearingFeeIndicator }

//BuildClearingFeeIndicator returns a new ClearingFeeIndicator initialized with val
func BuildClearingFeeIndicator(val string) ClearingFeeIndicator {
	var field ClearingFeeIndicator
	field.Value = val
	return field
}

//ClearingFirm is a STRING field
type ClearingFirm struct{ message.StringValue }

//Tag returns tag.ClearingFirm (439)
func (f ClearingFirm) Tag() fix.Tag { return tag.ClearingFirm }

//BuildClearingFirm returns a new ClearingFirm initialized with val
func BuildClearingFirm(val string) ClearingFirm {
	var field ClearingFirm
	field.Value = val
	return field
}

//ClearingInstruction is a INT field
type ClearingInstruction struct{ message.IntValue }

//Tag returns tag.ClearingInstruction (577)
func (f ClearingInstruction) Tag() fix.Tag { return tag.ClearingInstruction }

//BuildClearingInstruction returns a new ClearingInstruction initialized with val
func BuildClearingInstruction(val int) ClearingInstruction {
	var field ClearingInstruction
	field.Value = val
	return field
}

//ClientBidID is a STRING field
type ClientBidID struct{ message.StringValue }

//Tag returns tag.ClientBidID (391)
func (f ClientBidID) Tag() fix.Tag { return tag.ClientBidID }

//BuildClientBidID returns a new ClientBidID initialized with val
func BuildClientBidID(val string) ClientBidID {
	var field ClientBidID
	field.Value = val
	return field
}

//ClientID is a STRING field
type ClientID struct{ message.StringValue }

//Tag returns tag.ClientID (109)
func (f ClientID) Tag() fix.Tag { return tag.ClientID }

//BuildClientID returns a new ClientID initialized with val
func BuildClientID(val string) ClientID {
	var field ClientID
	field.Value = val
	return field
}

//CollAction is a INT field
type CollAction struct{ message.IntValue }

//Tag returns tag.CollAction (944)
func (f CollAction) Tag() fix.Tag { return tag.CollAction }

//BuildCollAction returns a new CollAction initialized with val
func BuildCollAction(val int) CollAction {
	var field CollAction
	field.Value = val
	return field
}

//CollApplType is a INT field
type CollApplType struct{ message.IntValue }

//Tag returns tag.CollApplType (1043)
func (f CollApplType) Tag() fix.Tag { return tag.CollApplType }

//BuildCollApplType returns a new CollApplType initialized with val
func BuildCollApplType(val int) CollApplType {
	var field CollApplType
	field.Value = val
	return field
}

//CollAsgnID is a STRING field
type CollAsgnID struct{ message.StringValue }

//Tag returns tag.CollAsgnID (902)
func (f CollAsgnID) Tag() fix.Tag { return tag.CollAsgnID }

//BuildCollAsgnID returns a new CollAsgnID initialized with val
func BuildCollAsgnID(val string) CollAsgnID {
	var field CollAsgnID
	field.Value = val
	return field
}

//CollAsgnReason is a INT field
type CollAsgnReason struct{ message.IntValue }

//Tag returns tag.CollAsgnReason (895)
func (f CollAsgnReason) Tag() fix.Tag { return tag.CollAsgnReason }

//BuildCollAsgnReason returns a new CollAsgnReason initialized with val
func BuildCollAsgnReason(val int) CollAsgnReason {
	var field CollAsgnReason
	field.Value = val
	return field
}

//CollAsgnRefID is a STRING field
type CollAsgnRefID struct{ message.StringValue }

//Tag returns tag.CollAsgnRefID (907)
func (f CollAsgnRefID) Tag() fix.Tag { return tag.CollAsgnRefID }

//BuildCollAsgnRefID returns a new CollAsgnRefID initialized with val
func BuildCollAsgnRefID(val string) CollAsgnRefID {
	var field CollAsgnRefID
	field.Value = val
	return field
}

//CollAsgnRejectReason is a INT field
type CollAsgnRejectReason struct{ message.IntValue }

//Tag returns tag.CollAsgnRejectReason (906)
func (f CollAsgnRejectReason) Tag() fix.Tag { return tag.CollAsgnRejectReason }

//BuildCollAsgnRejectReason returns a new CollAsgnRejectReason initialized with val
func BuildCollAsgnRejectReason(val int) CollAsgnRejectReason {
	var field CollAsgnRejectReason
	field.Value = val
	return field
}

//CollAsgnRespType is a INT field
type CollAsgnRespType struct{ message.IntValue }

//Tag returns tag.CollAsgnRespType (905)
func (f CollAsgnRespType) Tag() fix.Tag { return tag.CollAsgnRespType }

//BuildCollAsgnRespType returns a new CollAsgnRespType initialized with val
func BuildCollAsgnRespType(val int) CollAsgnRespType {
	var field CollAsgnRespType
	field.Value = val
	return field
}

//CollAsgnTransType is a INT field
type CollAsgnTransType struct{ message.IntValue }

//Tag returns tag.CollAsgnTransType (903)
func (f CollAsgnTransType) Tag() fix.Tag { return tag.CollAsgnTransType }

//BuildCollAsgnTransType returns a new CollAsgnTransType initialized with val
func BuildCollAsgnTransType(val int) CollAsgnTransType {
	var field CollAsgnTransType
	field.Value = val
	return field
}

//CollInquiryID is a STRING field
type CollInquiryID struct{ message.StringValue }

//Tag returns tag.CollInquiryID (909)
func (f CollInquiryID) Tag() fix.Tag { return tag.CollInquiryID }

//BuildCollInquiryID returns a new CollInquiryID initialized with val
func BuildCollInquiryID(val string) CollInquiryID {
	var field CollInquiryID
	field.Value = val
	return field
}

//CollInquiryQualifier is a INT field
type CollInquiryQualifier struct{ message.IntValue }

//Tag returns tag.CollInquiryQualifier (896)
func (f CollInquiryQualifier) Tag() fix.Tag { return tag.CollInquiryQualifier }

//BuildCollInquiryQualifier returns a new CollInquiryQualifier initialized with val
func BuildCollInquiryQualifier(val int) CollInquiryQualifier {
	var field CollInquiryQualifier
	field.Value = val
	return field
}

//CollInquiryResult is a INT field
type CollInquiryResult struct{ message.IntValue }

//Tag returns tag.CollInquiryResult (946)
func (f CollInquiryResult) Tag() fix.Tag { return tag.CollInquiryResult }

//BuildCollInquiryResult returns a new CollInquiryResult initialized with val
func BuildCollInquiryResult(val int) CollInquiryResult {
	var field CollInquiryResult
	field.Value = val
	return field
}

//CollInquiryStatus is a INT field
type CollInquiryStatus struct{ message.IntValue }

//Tag returns tag.CollInquiryStatus (945)
func (f CollInquiryStatus) Tag() fix.Tag { return tag.CollInquiryStatus }

//BuildCollInquiryStatus returns a new CollInquiryStatus initialized with val
func BuildCollInquiryStatus(val int) CollInquiryStatus {
	var field CollInquiryStatus
	field.Value = val
	return field
}

//CollReqID is a STRING field
type CollReqID struct{ message.StringValue }

//Tag returns tag.CollReqID (894)
func (f CollReqID) Tag() fix.Tag { return tag.CollReqID }

//BuildCollReqID returns a new CollReqID initialized with val
func BuildCollReqID(val string) CollReqID {
	var field CollReqID
	field.Value = val
	return field
}

//CollRespID is a STRING field
type CollRespID struct{ message.StringValue }

//Tag returns tag.CollRespID (904)
func (f CollRespID) Tag() fix.Tag { return tag.CollRespID }

//BuildCollRespID returns a new CollRespID initialized with val
func BuildCollRespID(val string) CollRespID {
	var field CollRespID
	field.Value = val
	return field
}

//CollRptID is a STRING field
type CollRptID struct{ message.StringValue }

//Tag returns tag.CollRptID (908)
func (f CollRptID) Tag() fix.Tag { return tag.CollRptID }

//BuildCollRptID returns a new CollRptID initialized with val
func BuildCollRptID(val string) CollRptID {
	var field CollRptID
	field.Value = val
	return field
}

//CollStatus is a INT field
type CollStatus struct{ message.IntValue }

//Tag returns tag.CollStatus (910)
func (f CollStatus) Tag() fix.Tag { return tag.CollStatus }

//BuildCollStatus returns a new CollStatus initialized with val
func BuildCollStatus(val int) CollStatus {
	var field CollStatus
	field.Value = val
	return field
}

//CommCurrency is a CURRENCY field
type CommCurrency struct{ message.CurrencyValue }

//Tag returns tag.CommCurrency (479)
func (f CommCurrency) Tag() fix.Tag { return tag.CommCurrency }

//BuildCommCurrency returns a new CommCurrency initialized with val
func BuildCommCurrency(val string) CommCurrency {
	var field CommCurrency
	field.Value = val
	return field
}

//CommType is a CHAR field
type CommType struct{ message.CharValue }

//Tag returns tag.CommType (13)
func (f CommType) Tag() fix.Tag { return tag.CommType }

//BuildCommType returns a new CommType initialized with val
func BuildCommType(val string) CommType {
	var field CommType
	field.Value = val
	return field
}

//Commission is a AMT field
type Commission struct{ message.AmtValue }

//Tag returns tag.Commission (12)
func (f Commission) Tag() fix.Tag { return tag.Commission }

//BuildCommission returns a new Commission initialized with val
func BuildCommission(val float64) Commission {
	var field Commission
	field.Value = val
	return field
}

//ComplexEventCondition is a INT field
type ComplexEventCondition struct{ message.IntValue }

//Tag returns tag.ComplexEventCondition (1490)
func (f ComplexEventCondition) Tag() fix.Tag { return tag.ComplexEventCondition }

//BuildComplexEventCondition returns a new ComplexEventCondition initialized with val
func BuildComplexEventCondition(val int) ComplexEventCondition {
	var field ComplexEventCondition
	field.Value = val
	return field
}

//ComplexEventEndDate is a UTCTIMESTAMP field
type ComplexEventEndDate struct{ message.UTCTimestampValue }

//Tag returns tag.ComplexEventEndDate (1493)
func (f ComplexEventEndDate) Tag() fix.Tag { return tag.ComplexEventEndDate }

//ComplexEventEndTime is a UTCTIMEONLY field
type ComplexEventEndTime struct{ message.UTCTimeOnlyValue }

//Tag returns tag.ComplexEventEndTime (1496)
func (f ComplexEventEndTime) Tag() fix.Tag { return tag.ComplexEventEndTime }

//ComplexEventPrice is a PRICE field
type ComplexEventPrice struct{ message.PriceValue }

//Tag returns tag.ComplexEventPrice (1486)
func (f ComplexEventPrice) Tag() fix.Tag { return tag.ComplexEventPrice }

//BuildComplexEventPrice returns a new ComplexEventPrice initialized with val
func BuildComplexEventPrice(val float64) ComplexEventPrice {
	var field ComplexEventPrice
	field.Value = val
	return field
}

//ComplexEventPriceBoundaryMethod is a INT field
type ComplexEventPriceBoundaryMethod struct{ message.IntValue }

//Tag returns tag.ComplexEventPriceBoundaryMethod (1487)
func (f ComplexEventPriceBoundaryMethod) Tag() fix.Tag { return tag.ComplexEventPriceBoundaryMethod }

//BuildComplexEventPriceBoundaryMethod returns a new ComplexEventPriceBoundaryMethod initialized with val
func BuildComplexEventPriceBoundaryMethod(val int) ComplexEventPriceBoundaryMethod {
	var field ComplexEventPriceBoundaryMethod
	field.Value = val
	return field
}

//ComplexEventPriceBoundaryPrecision is a PERCENTAGE field
type ComplexEventPriceBoundaryPrecision struct{ message.PercentageValue }

//Tag returns tag.ComplexEventPriceBoundaryPrecision (1488)
func (f ComplexEventPriceBoundaryPrecision) Tag() fix.Tag {
	return tag.ComplexEventPriceBoundaryPrecision
}

//BuildComplexEventPriceBoundaryPrecision returns a new ComplexEventPriceBoundaryPrecision initialized with val
func BuildComplexEventPriceBoundaryPrecision(val float64) ComplexEventPriceBoundaryPrecision {
	var field ComplexEventPriceBoundaryPrecision
	field.Value = val
	return field
}

//ComplexEventPriceTimeType is a INT field
type ComplexEventPriceTimeType struct{ message.IntValue }

//Tag returns tag.ComplexEventPriceTimeType (1489)
func (f ComplexEventPriceTimeType) Tag() fix.Tag { return tag.ComplexEventPriceTimeType }

//BuildComplexEventPriceTimeType returns a new ComplexEventPriceTimeType initialized with val
func BuildComplexEventPriceTimeType(val int) ComplexEventPriceTimeType {
	var field ComplexEventPriceTimeType
	field.Value = val
	return field
}

//ComplexEventStartDate is a UTCTIMESTAMP field
type ComplexEventStartDate struct{ message.UTCTimestampValue }

//Tag returns tag.ComplexEventStartDate (1492)
func (f ComplexEventStartDate) Tag() fix.Tag { return tag.ComplexEventStartDate }

//ComplexEventStartTime is a UTCTIMEONLY field
type ComplexEventStartTime struct{ message.UTCTimeOnlyValue }

//Tag returns tag.ComplexEventStartTime (1495)
func (f ComplexEventStartTime) Tag() fix.Tag { return tag.ComplexEventStartTime }

//ComplexEventType is a INT field
type ComplexEventType struct{ message.IntValue }

//Tag returns tag.ComplexEventType (1484)
func (f ComplexEventType) Tag() fix.Tag { return tag.ComplexEventType }

//BuildComplexEventType returns a new ComplexEventType initialized with val
func BuildComplexEventType(val int) ComplexEventType {
	var field ComplexEventType
	field.Value = val
	return field
}

//ComplexOptPayoutAmount is a AMT field
type ComplexOptPayoutAmount struct{ message.AmtValue }

//Tag returns tag.ComplexOptPayoutAmount (1485)
func (f ComplexOptPayoutAmount) Tag() fix.Tag { return tag.ComplexOptPayoutAmount }

//BuildComplexOptPayoutAmount returns a new ComplexOptPayoutAmount initialized with val
func BuildComplexOptPayoutAmount(val float64) ComplexOptPayoutAmount {
	var field ComplexOptPayoutAmount
	field.Value = val
	return field
}

//ComplianceID is a STRING field
type ComplianceID struct{ message.StringValue }

//Tag returns tag.ComplianceID (376)
func (f ComplianceID) Tag() fix.Tag { return tag.ComplianceID }

//BuildComplianceID returns a new ComplianceID initialized with val
func BuildComplianceID(val string) ComplianceID {
	var field ComplianceID
	field.Value = val
	return field
}

//Concession is a AMT field
type Concession struct{ message.AmtValue }

//Tag returns tag.Concession (238)
func (f Concession) Tag() fix.Tag { return tag.Concession }

//BuildConcession returns a new Concession initialized with val
func BuildConcession(val float64) Concession {
	var field Concession
	field.Value = val
	return field
}

//ConfirmID is a STRING field
type ConfirmID struct{ message.StringValue }

//Tag returns tag.ConfirmID (664)
func (f ConfirmID) Tag() fix.Tag { return tag.ConfirmID }

//BuildConfirmID returns a new ConfirmID initialized with val
func BuildConfirmID(val string) ConfirmID {
	var field ConfirmID
	field.Value = val
	return field
}

//ConfirmRefID is a STRING field
type ConfirmRefID struct{ message.StringValue }

//Tag returns tag.ConfirmRefID (772)
func (f ConfirmRefID) Tag() fix.Tag { return tag.ConfirmRefID }

//BuildConfirmRefID returns a new ConfirmRefID initialized with val
func BuildConfirmRefID(val string) ConfirmRefID {
	var field ConfirmRefID
	field.Value = val
	return field
}

//ConfirmRejReason is a INT field
type ConfirmRejReason struct{ message.IntValue }

//Tag returns tag.ConfirmRejReason (774)
func (f ConfirmRejReason) Tag() fix.Tag { return tag.ConfirmRejReason }

//BuildConfirmRejReason returns a new ConfirmRejReason initialized with val
func BuildConfirmRejReason(val int) ConfirmRejReason {
	var field ConfirmRejReason
	field.Value = val
	return field
}

//ConfirmReqID is a STRING field
type ConfirmReqID struct{ message.StringValue }

//Tag returns tag.ConfirmReqID (859)
func (f ConfirmReqID) Tag() fix.Tag { return tag.ConfirmReqID }

//BuildConfirmReqID returns a new ConfirmReqID initialized with val
func BuildConfirmReqID(val string) ConfirmReqID {
	var field ConfirmReqID
	field.Value = val
	return field
}

//ConfirmStatus is a INT field
type ConfirmStatus struct{ message.IntValue }

//Tag returns tag.ConfirmStatus (665)
func (f ConfirmStatus) Tag() fix.Tag { return tag.ConfirmStatus }

//BuildConfirmStatus returns a new ConfirmStatus initialized with val
func BuildConfirmStatus(val int) ConfirmStatus {
	var field ConfirmStatus
	field.Value = val
	return field
}

//ConfirmTransType is a INT field
type ConfirmTransType struct{ message.IntValue }

//Tag returns tag.ConfirmTransType (666)
func (f ConfirmTransType) Tag() fix.Tag { return tag.ConfirmTransType }

//BuildConfirmTransType returns a new ConfirmTransType initialized with val
func BuildConfirmTransType(val int) ConfirmTransType {
	var field ConfirmTransType
	field.Value = val
	return field
}

//ConfirmType is a INT field
type ConfirmType struct{ message.IntValue }

//Tag returns tag.ConfirmType (773)
func (f ConfirmType) Tag() fix.Tag { return tag.ConfirmType }

//BuildConfirmType returns a new ConfirmType initialized with val
func BuildConfirmType(val int) ConfirmType {
	var field ConfirmType
	field.Value = val
	return field
}

//ContAmtCurr is a CURRENCY field
type ContAmtCurr struct{ message.CurrencyValue }

//Tag returns tag.ContAmtCurr (521)
func (f ContAmtCurr) Tag() fix.Tag { return tag.ContAmtCurr }

//BuildContAmtCurr returns a new ContAmtCurr initialized with val
func BuildContAmtCurr(val string) ContAmtCurr {
	var field ContAmtCurr
	field.Value = val
	return field
}

//ContAmtType is a INT field
type ContAmtType struct{ message.IntValue }

//Tag returns tag.ContAmtType (519)
func (f ContAmtType) Tag() fix.Tag { return tag.ContAmtType }

//BuildContAmtType returns a new ContAmtType initialized with val
func BuildContAmtType(val int) ContAmtType {
	var field ContAmtType
	field.Value = val
	return field
}

//ContAmtValue is a FLOAT field
type ContAmtValue struct{ message.FloatValue }

//Tag returns tag.ContAmtValue (520)
func (f ContAmtValue) Tag() fix.Tag { return tag.ContAmtValue }

//BuildContAmtValue returns a new ContAmtValue initialized with val
func BuildContAmtValue(val float64) ContAmtValue {
	var field ContAmtValue
	field.Value = val
	return field
}

//ContIntRptID is a STRING field
type ContIntRptID struct{ message.StringValue }

//Tag returns tag.ContIntRptID (977)
func (f ContIntRptID) Tag() fix.Tag { return tag.ContIntRptID }

//BuildContIntRptID returns a new ContIntRptID initialized with val
func BuildContIntRptID(val string) ContIntRptID {
	var field ContIntRptID
	field.Value = val
	return field
}

//ContextPartyID is a STRING field
type ContextPartyID struct{ message.StringValue }

//Tag returns tag.ContextPartyID (1523)
func (f ContextPartyID) Tag() fix.Tag { return tag.ContextPartyID }

//BuildContextPartyID returns a new ContextPartyID initialized with val
func BuildContextPartyID(val string) ContextPartyID {
	var field ContextPartyID
	field.Value = val
	return field
}

//ContextPartyIDSource is a CHAR field
type ContextPartyIDSource struct{ message.CharValue }

//Tag returns tag.ContextPartyIDSource (1524)
func (f ContextPartyIDSource) Tag() fix.Tag { return tag.ContextPartyIDSource }

//BuildContextPartyIDSource returns a new ContextPartyIDSource initialized with val
func BuildContextPartyIDSource(val string) ContextPartyIDSource {
	var field ContextPartyIDSource
	field.Value = val
	return field
}

//ContextPartyRole is a INT field
type ContextPartyRole struct{ message.IntValue }

//Tag returns tag.ContextPartyRole (1525)
func (f ContextPartyRole) Tag() fix.Tag { return tag.ContextPartyRole }

//BuildContextPartyRole returns a new ContextPartyRole initialized with val
func BuildContextPartyRole(val int) ContextPartyRole {
	var field ContextPartyRole
	field.Value = val
	return field
}

//ContextPartySubID is a STRING field
type ContextPartySubID struct{ message.StringValue }

//Tag returns tag.ContextPartySubID (1527)
func (f ContextPartySubID) Tag() fix.Tag { return tag.ContextPartySubID }

//BuildContextPartySubID returns a new ContextPartySubID initialized with val
func BuildContextPartySubID(val string) ContextPartySubID {
	var field ContextPartySubID
	field.Value = val
	return field
}

//ContextPartySubIDType is a INT field
type ContextPartySubIDType struct{ message.IntValue }

//Tag returns tag.ContextPartySubIDType (1528)
func (f ContextPartySubIDType) Tag() fix.Tag { return tag.ContextPartySubIDType }

//BuildContextPartySubIDType returns a new ContextPartySubIDType initialized with val
func BuildContextPartySubIDType(val int) ContextPartySubIDType {
	var field ContextPartySubIDType
	field.Value = val
	return field
}

//ContingencyType is a INT field
type ContingencyType struct{ message.IntValue }

//Tag returns tag.ContingencyType (1385)
func (f ContingencyType) Tag() fix.Tag { return tag.ContingencyType }

//BuildContingencyType returns a new ContingencyType initialized with val
func BuildContingencyType(val int) ContingencyType {
	var field ContingencyType
	field.Value = val
	return field
}

//ContraBroker is a STRING field
type ContraBroker struct{ message.StringValue }

//Tag returns tag.ContraBroker (375)
func (f ContraBroker) Tag() fix.Tag { return tag.ContraBroker }

//BuildContraBroker returns a new ContraBroker initialized with val
func BuildContraBroker(val string) ContraBroker {
	var field ContraBroker
	field.Value = val
	return field
}

//ContraLegRefID is a STRING field
type ContraLegRefID struct{ message.StringValue }

//Tag returns tag.ContraLegRefID (655)
func (f ContraLegRefID) Tag() fix.Tag { return tag.ContraLegRefID }

//BuildContraLegRefID returns a new ContraLegRefID initialized with val
func BuildContraLegRefID(val string) ContraLegRefID {
	var field ContraLegRefID
	field.Value = val
	return field
}

//ContraTradeQty is a QTY field
type ContraTradeQty struct{ message.QtyValue }

//Tag returns tag.ContraTradeQty (437)
func (f ContraTradeQty) Tag() fix.Tag { return tag.ContraTradeQty }

//BuildContraTradeQty returns a new ContraTradeQty initialized with val
func BuildContraTradeQty(val float64) ContraTradeQty {
	var field ContraTradeQty
	field.Value = val
	return field
}

//ContraTradeTime is a UTCTIMESTAMP field
type ContraTradeTime struct{ message.UTCTimestampValue }

//Tag returns tag.ContraTradeTime (438)
func (f ContraTradeTime) Tag() fix.Tag { return tag.ContraTradeTime }

//ContraTrader is a STRING field
type ContraTrader struct{ message.StringValue }

//Tag returns tag.ContraTrader (337)
func (f ContraTrader) Tag() fix.Tag { return tag.ContraTrader }

//BuildContraTrader returns a new ContraTrader initialized with val
func BuildContraTrader(val string) ContraTrader {
	var field ContraTrader
	field.Value = val
	return field
}

//ContractMultiplier is a FLOAT field
type ContractMultiplier struct{ message.FloatValue }

//Tag returns tag.ContractMultiplier (231)
func (f ContractMultiplier) Tag() fix.Tag { return tag.ContractMultiplier }

//BuildContractMultiplier returns a new ContractMultiplier initialized with val
func BuildContractMultiplier(val float64) ContractMultiplier {
	var field ContractMultiplier
	field.Value = val
	return field
}

//ContractMultiplierUnit is a INT field
type ContractMultiplierUnit struct{ message.IntValue }

//Tag returns tag.ContractMultiplierUnit (1435)
func (f ContractMultiplierUnit) Tag() fix.Tag { return tag.ContractMultiplierUnit }

//BuildContractMultiplierUnit returns a new ContractMultiplierUnit initialized with val
func BuildContractMultiplierUnit(val int) ContractMultiplierUnit {
	var field ContractMultiplierUnit
	field.Value = val
	return field
}

//ContractSettlMonth is a MONTHYEAR field
type ContractSettlMonth struct{ message.MonthYearValue }

//Tag returns tag.ContractSettlMonth (667)
func (f ContractSettlMonth) Tag() fix.Tag { return tag.ContractSettlMonth }

//BuildContractSettlMonth returns a new ContractSettlMonth initialized with val
func BuildContractSettlMonth(val string) ContractSettlMonth {
	var field ContractSettlMonth
	field.Value = val
	return field
}

//ContraryInstructionIndicator is a BOOLEAN field
type ContraryInstructionIndicator struct{ message.BooleanValue }

//Tag returns tag.ContraryInstructionIndicator (719)
func (f ContraryInstructionIndicator) Tag() fix.Tag { return tag.ContraryInstructionIndicator }

//BuildContraryInstructionIndicator returns a new ContraryInstructionIndicator initialized with val
func BuildContraryInstructionIndicator(val bool) ContraryInstructionIndicator {
	var field ContraryInstructionIndicator
	field.Value = val
	return field
}

//CopyMsgIndicator is a BOOLEAN field
type CopyMsgIndicator struct{ message.BooleanValue }

//Tag returns tag.CopyMsgIndicator (797)
func (f CopyMsgIndicator) Tag() fix.Tag { return tag.CopyMsgIndicator }

//BuildCopyMsgIndicator returns a new CopyMsgIndicator initialized with val
func BuildCopyMsgIndicator(val bool) CopyMsgIndicator {
	var field CopyMsgIndicator
	field.Value = val
	return field
}

//CorporateAction is a MULTIPLECHARVALUE field
type CorporateAction struct{ message.MultipleCharValue }

//Tag returns tag.CorporateAction (292)
func (f CorporateAction) Tag() fix.Tag { return tag.CorporateAction }

//BuildCorporateAction returns a new CorporateAction initialized with val
func BuildCorporateAction(val string) CorporateAction {
	var field CorporateAction
	field.Value = val
	return field
}

//Country is a COUNTRY field
type Country struct{ message.CountryValue }

//Tag returns tag.Country (421)
func (f Country) Tag() fix.Tag { return tag.Country }

//BuildCountry returns a new Country initialized with val
func BuildCountry(val string) Country {
	var field Country
	field.Value = val
	return field
}

//CountryOfIssue is a COUNTRY field
type CountryOfIssue struct{ message.CountryValue }

//Tag returns tag.CountryOfIssue (470)
func (f CountryOfIssue) Tag() fix.Tag { return tag.CountryOfIssue }

//BuildCountryOfIssue returns a new CountryOfIssue initialized with val
func BuildCountryOfIssue(val string) CountryOfIssue {
	var field CountryOfIssue
	field.Value = val
	return field
}

//CouponPaymentDate is a LOCALMKTDATE field
type CouponPaymentDate struct{ message.LocalMktDateValue }

//Tag returns tag.CouponPaymentDate (224)
func (f CouponPaymentDate) Tag() fix.Tag { return tag.CouponPaymentDate }

//BuildCouponPaymentDate returns a new CouponPaymentDate initialized with val
func BuildCouponPaymentDate(val string) CouponPaymentDate {
	var field CouponPaymentDate
	field.Value = val
	return field
}

//CouponRate is a PERCENTAGE field
type CouponRate struct{ message.PercentageValue }

//Tag returns tag.CouponRate (223)
func (f CouponRate) Tag() fix.Tag { return tag.CouponRate }

//BuildCouponRate returns a new CouponRate initialized with val
func BuildCouponRate(val float64) CouponRate {
	var field CouponRate
	field.Value = val
	return field
}

//CoveredOrUncovered is a INT field
type CoveredOrUncovered struct{ message.IntValue }

//Tag returns tag.CoveredOrUncovered (203)
func (f CoveredOrUncovered) Tag() fix.Tag { return tag.CoveredOrUncovered }

//BuildCoveredOrUncovered returns a new CoveredOrUncovered initialized with val
func BuildCoveredOrUncovered(val int) CoveredOrUncovered {
	var field CoveredOrUncovered
	field.Value = val
	return field
}

//CreditRating is a STRING field
type CreditRating struct{ message.StringValue }

//Tag returns tag.CreditRating (255)
func (f CreditRating) Tag() fix.Tag { return tag.CreditRating }

//BuildCreditRating returns a new CreditRating initialized with val
func BuildCreditRating(val string) CreditRating {
	var field CreditRating
	field.Value = val
	return field
}

//CrossID is a STRING field
type CrossID struct{ message.StringValue }

//Tag returns tag.CrossID (548)
func (f CrossID) Tag() fix.Tag { return tag.CrossID }

//BuildCrossID returns a new CrossID initialized with val
func BuildCrossID(val string) CrossID {
	var field CrossID
	field.Value = val
	return field
}

//CrossPercent is a PERCENTAGE field
type CrossPercent struct{ message.PercentageValue }

//Tag returns tag.CrossPercent (413)
func (f CrossPercent) Tag() fix.Tag { return tag.CrossPercent }

//BuildCrossPercent returns a new CrossPercent initialized with val
func BuildCrossPercent(val float64) CrossPercent {
	var field CrossPercent
	field.Value = val
	return field
}

//CrossPrioritization is a INT field
type CrossPrioritization struct{ message.IntValue }

//Tag returns tag.CrossPrioritization (550)
func (f CrossPrioritization) Tag() fix.Tag { return tag.CrossPrioritization }

//BuildCrossPrioritization returns a new CrossPrioritization initialized with val
func BuildCrossPrioritization(val int) CrossPrioritization {
	var field CrossPrioritization
	field.Value = val
	return field
}

//CrossType is a INT field
type CrossType struct{ message.IntValue }

//Tag returns tag.CrossType (549)
func (f CrossType) Tag() fix.Tag { return tag.CrossType }

//BuildCrossType returns a new CrossType initialized with val
func BuildCrossType(val int) CrossType {
	var field CrossType
	field.Value = val
	return field
}

//CstmApplVerID is a STRING field
type CstmApplVerID struct{ message.StringValue }

//Tag returns tag.CstmApplVerID (1129)
func (f CstmApplVerID) Tag() fix.Tag { return tag.CstmApplVerID }

//BuildCstmApplVerID returns a new CstmApplVerID initialized with val
func BuildCstmApplVerID(val string) CstmApplVerID {
	var field CstmApplVerID
	field.Value = val
	return field
}

//CumQty is a QTY field
type CumQty struct{ message.QtyValue }

//Tag returns tag.CumQty (14)
func (f CumQty) Tag() fix.Tag { return tag.CumQty }

//BuildCumQty returns a new CumQty initialized with val
func BuildCumQty(val float64) CumQty {
	var field CumQty
	field.Value = val
	return field
}

//Currency is a CURRENCY field
type Currency struct{ message.CurrencyValue }

//Tag returns tag.Currency (15)
func (f Currency) Tag() fix.Tag { return tag.Currency }

//BuildCurrency returns a new Currency initialized with val
func BuildCurrency(val string) Currency {
	var field Currency
	field.Value = val
	return field
}

//CurrencyRatio is a FLOAT field
type CurrencyRatio struct{ message.FloatValue }

//Tag returns tag.CurrencyRatio (1382)
func (f CurrencyRatio) Tag() fix.Tag { return tag.CurrencyRatio }

//BuildCurrencyRatio returns a new CurrencyRatio initialized with val
func BuildCurrencyRatio(val float64) CurrencyRatio {
	var field CurrencyRatio
	field.Value = val
	return field
}

//CustDirectedOrder is a BOOLEAN field
type CustDirectedOrder struct{ message.BooleanValue }

//Tag returns tag.CustDirectedOrder (1029)
func (f CustDirectedOrder) Tag() fix.Tag { return tag.CustDirectedOrder }

//BuildCustDirectedOrder returns a new CustDirectedOrder initialized with val
func BuildCustDirectedOrder(val bool) CustDirectedOrder {
	var field CustDirectedOrder
	field.Value = val
	return field
}

//CustOrderCapacity is a INT field
type CustOrderCapacity struct{ message.IntValue }

//Tag returns tag.CustOrderCapacity (582)
func (f CustOrderCapacity) Tag() fix.Tag { return tag.CustOrderCapacity }

//BuildCustOrderCapacity returns a new CustOrderCapacity initialized with val
func BuildCustOrderCapacity(val int) CustOrderCapacity {
	var field CustOrderCapacity
	field.Value = val
	return field
}

//CustOrderHandlingInst is a MULTIPLESTRINGVALUE field
type CustOrderHandlingInst struct{ message.MultipleStringValue }

//Tag returns tag.CustOrderHandlingInst (1031)
func (f CustOrderHandlingInst) Tag() fix.Tag { return tag.CustOrderHandlingInst }

//BuildCustOrderHandlingInst returns a new CustOrderHandlingInst initialized with val
func BuildCustOrderHandlingInst(val string) CustOrderHandlingInst {
	var field CustOrderHandlingInst
	field.Value = val
	return field
}

//CustomerOrFirm is a INT field
type CustomerOrFirm struct{ message.IntValue }

//Tag returns tag.CustomerOrFirm (204)
func (f CustomerOrFirm) Tag() fix.Tag { return tag.CustomerOrFirm }

//BuildCustomerOrFirm returns a new CustomerOrFirm initialized with val
func BuildCustomerOrFirm(val int) CustomerOrFirm {
	var field CustomerOrFirm
	field.Value = val
	return field
}

//CxlQty is a QTY field
type CxlQty struct{ message.QtyValue }

//Tag returns tag.CxlQty (84)
func (f CxlQty) Tag() fix.Tag { return tag.CxlQty }

//BuildCxlQty returns a new CxlQty initialized with val
func BuildCxlQty(val float64) CxlQty {
	var field CxlQty
	field.Value = val
	return field
}

//CxlRejReason is a INT field
type CxlRejReason struct{ message.IntValue }

//Tag returns tag.CxlRejReason (102)
func (f CxlRejReason) Tag() fix.Tag { return tag.CxlRejReason }

//BuildCxlRejReason returns a new CxlRejReason initialized with val
func BuildCxlRejReason(val int) CxlRejReason {
	var field CxlRejReason
	field.Value = val
	return field
}

//CxlRejResponseTo is a CHAR field
type CxlRejResponseTo struct{ message.CharValue }

//Tag returns tag.CxlRejResponseTo (434)
func (f CxlRejResponseTo) Tag() fix.Tag { return tag.CxlRejResponseTo }

//BuildCxlRejResponseTo returns a new CxlRejResponseTo initialized with val
func BuildCxlRejResponseTo(val string) CxlRejResponseTo {
	var field CxlRejResponseTo
	field.Value = val
	return field
}

//CxlType is a CHAR field
type CxlType struct{ message.CharValue }

//Tag returns tag.CxlType (125)
func (f CxlType) Tag() fix.Tag { return tag.CxlType }

//BuildCxlType returns a new CxlType initialized with val
func BuildCxlType(val string) CxlType {
	var field CxlType
	field.Value = val
	return field
}

//DKReason is a CHAR field
type DKReason struct{ message.CharValue }

//Tag returns tag.DKReason (127)
func (f DKReason) Tag() fix.Tag { return tag.DKReason }

//BuildDKReason returns a new DKReason initialized with val
func BuildDKReason(val string) DKReason {
	var field DKReason
	field.Value = val
	return field
}

//DateOfBirth is a LOCALMKTDATE field
type DateOfBirth struct{ message.LocalMktDateValue }

//Tag returns tag.DateOfBirth (486)
func (f DateOfBirth) Tag() fix.Tag { return tag.DateOfBirth }

//BuildDateOfBirth returns a new DateOfBirth initialized with val
func BuildDateOfBirth(val string) DateOfBirth {
	var field DateOfBirth
	field.Value = val
	return field
}

//DatedDate is a LOCALMKTDATE field
type DatedDate struct{ message.LocalMktDateValue }

//Tag returns tag.DatedDate (873)
func (f DatedDate) Tag() fix.Tag { return tag.DatedDate }

//BuildDatedDate returns a new DatedDate initialized with val
func BuildDatedDate(val string) DatedDate {
	var field DatedDate
	field.Value = val
	return field
}

//DayAvgPx is a PRICE field
type DayAvgPx struct{ message.PriceValue }

//Tag returns tag.DayAvgPx (426)
func (f DayAvgPx) Tag() fix.Tag { return tag.DayAvgPx }

//BuildDayAvgPx returns a new DayAvgPx initialized with val
func BuildDayAvgPx(val float64) DayAvgPx {
	var field DayAvgPx
	field.Value = val
	return field
}

//DayBookingInst is a CHAR field
type DayBookingInst struct{ message.CharValue }

//Tag returns tag.DayBookingInst (589)
func (f DayBookingInst) Tag() fix.Tag { return tag.DayBookingInst }

//BuildDayBookingInst returns a new DayBookingInst initialized with val
func BuildDayBookingInst(val string) DayBookingInst {
	var field DayBookingInst
	field.Value = val
	return field
}

//DayCumQty is a QTY field
type DayCumQty struct{ message.QtyValue }

//Tag returns tag.DayCumQty (425)
func (f DayCumQty) Tag() fix.Tag { return tag.DayCumQty }

//BuildDayCumQty returns a new DayCumQty initialized with val
func BuildDayCumQty(val float64) DayCumQty {
	var field DayCumQty
	field.Value = val
	return field
}

//DayOrderQty is a QTY field
type DayOrderQty struct{ message.QtyValue }

//Tag returns tag.DayOrderQty (424)
func (f DayOrderQty) Tag() fix.Tag { return tag.DayOrderQty }

//BuildDayOrderQty returns a new DayOrderQty initialized with val
func BuildDayOrderQty(val float64) DayOrderQty {
	var field DayOrderQty
	field.Value = val
	return field
}

//DealingCapacity is a CHAR field
type DealingCapacity struct{ message.CharValue }

//Tag returns tag.DealingCapacity (1048)
func (f DealingCapacity) Tag() fix.Tag { return tag.DealingCapacity }

//BuildDealingCapacity returns a new DealingCapacity initialized with val
func BuildDealingCapacity(val string) DealingCapacity {
	var field DealingCapacity
	field.Value = val
	return field
}

//DefBidSize is a QTY field
type DefBidSize struct{ message.QtyValue }

//Tag returns tag.DefBidSize (293)
func (f DefBidSize) Tag() fix.Tag { return tag.DefBidSize }

//BuildDefBidSize returns a new DefBidSize initialized with val
func BuildDefBidSize(val float64) DefBidSize {
	var field DefBidSize
	field.Value = val
	return field
}

//DefOfferSize is a QTY field
type DefOfferSize struct{ message.QtyValue }

//Tag returns tag.DefOfferSize (294)
func (f DefOfferSize) Tag() fix.Tag { return tag.DefOfferSize }

//BuildDefOfferSize returns a new DefOfferSize initialized with val
func BuildDefOfferSize(val float64) DefOfferSize {
	var field DefOfferSize
	field.Value = val
	return field
}

//DefaultApplExtID is a INT field
type DefaultApplExtID struct{ message.IntValue }

//Tag returns tag.DefaultApplExtID (1407)
func (f DefaultApplExtID) Tag() fix.Tag { return tag.DefaultApplExtID }

//BuildDefaultApplExtID returns a new DefaultApplExtID initialized with val
func BuildDefaultApplExtID(val int) DefaultApplExtID {
	var field DefaultApplExtID
	field.Value = val
	return field
}

//DefaultApplVerID is a STRING field
type DefaultApplVerID struct{ message.StringValue }

//Tag returns tag.DefaultApplVerID (1137)
func (f DefaultApplVerID) Tag() fix.Tag { return tag.DefaultApplVerID }

//BuildDefaultApplVerID returns a new DefaultApplVerID initialized with val
func BuildDefaultApplVerID(val string) DefaultApplVerID {
	var field DefaultApplVerID
	field.Value = val
	return field
}

//DefaultCstmApplVerID is a STRING field
type DefaultCstmApplVerID struct{ message.StringValue }

//Tag returns tag.DefaultCstmApplVerID (1408)
func (f DefaultCstmApplVerID) Tag() fix.Tag { return tag.DefaultCstmApplVerID }

//BuildDefaultCstmApplVerID returns a new DefaultCstmApplVerID initialized with val
func BuildDefaultCstmApplVerID(val string) DefaultCstmApplVerID {
	var field DefaultCstmApplVerID
	field.Value = val
	return field
}

//DefaultVerIndicator is a BOOLEAN field
type DefaultVerIndicator struct{ message.BooleanValue }

//Tag returns tag.DefaultVerIndicator (1410)
func (f DefaultVerIndicator) Tag() fix.Tag { return tag.DefaultVerIndicator }

//BuildDefaultVerIndicator returns a new DefaultVerIndicator initialized with val
func BuildDefaultVerIndicator(val bool) DefaultVerIndicator {
	var field DefaultVerIndicator
	field.Value = val
	return field
}

//DeleteReason is a CHAR field
type DeleteReason struct{ message.CharValue }

//Tag returns tag.DeleteReason (285)
func (f DeleteReason) Tag() fix.Tag { return tag.DeleteReason }

//BuildDeleteReason returns a new DeleteReason initialized with val
func BuildDeleteReason(val string) DeleteReason {
	var field DeleteReason
	field.Value = val
	return field
}

//DeliverToCompID is a STRING field
type DeliverToCompID struct{ message.StringValue }

//Tag returns tag.DeliverToCompID (128)
func (f DeliverToCompID) Tag() fix.Tag { return tag.DeliverToCompID }

//BuildDeliverToCompID returns a new DeliverToCompID initialized with val
func BuildDeliverToCompID(val string) DeliverToCompID {
	var field DeliverToCompID
	field.Value = val
	return field
}

//DeliverToLocationID is a STRING field
type DeliverToLocationID struct{ message.StringValue }

//Tag returns tag.DeliverToLocationID (145)
func (f DeliverToLocationID) Tag() fix.Tag { return tag.DeliverToLocationID }

//BuildDeliverToLocationID returns a new DeliverToLocationID initialized with val
func BuildDeliverToLocationID(val string) DeliverToLocationID {
	var field DeliverToLocationID
	field.Value = val
	return field
}

//DeliverToSubID is a STRING field
type DeliverToSubID struct{ message.StringValue }

//Tag returns tag.DeliverToSubID (129)
func (f DeliverToSubID) Tag() fix.Tag { return tag.DeliverToSubID }

//BuildDeliverToSubID returns a new DeliverToSubID initialized with val
func BuildDeliverToSubID(val string) DeliverToSubID {
	var field DeliverToSubID
	field.Value = val
	return field
}

//DeliveryDate is a LOCALMKTDATE field
type DeliveryDate struct{ message.LocalMktDateValue }

//Tag returns tag.DeliveryDate (743)
func (f DeliveryDate) Tag() fix.Tag { return tag.DeliveryDate }

//BuildDeliveryDate returns a new DeliveryDate initialized with val
func BuildDeliveryDate(val string) DeliveryDate {
	var field DeliveryDate
	field.Value = val
	return field
}

//DeliveryForm is a INT field
type DeliveryForm struct{ message.IntValue }

//Tag returns tag.DeliveryForm (668)
func (f DeliveryForm) Tag() fix.Tag { return tag.DeliveryForm }

//BuildDeliveryForm returns a new DeliveryForm initialized with val
func BuildDeliveryForm(val int) DeliveryForm {
	var field DeliveryForm
	field.Value = val
	return field
}

//DeliveryType is a INT field
type DeliveryType struct{ message.IntValue }

//Tag returns tag.DeliveryType (919)
func (f DeliveryType) Tag() fix.Tag { return tag.DeliveryType }

//BuildDeliveryType returns a new DeliveryType initialized with val
func BuildDeliveryType(val int) DeliveryType {
	var field DeliveryType
	field.Value = val
	return field
}

//DerivFlexProductEligibilityIndicator is a BOOLEAN field
type DerivFlexProductEligibilityIndicator struct{ message.BooleanValue }

//Tag returns tag.DerivFlexProductEligibilityIndicator (1243)
func (f DerivFlexProductEligibilityIndicator) Tag() fix.Tag {
	return tag.DerivFlexProductEligibilityIndicator
}

//BuildDerivFlexProductEligibilityIndicator returns a new DerivFlexProductEligibilityIndicator initialized with val
func BuildDerivFlexProductEligibilityIndicator(val bool) DerivFlexProductEligibilityIndicator {
	var field DerivFlexProductEligibilityIndicator
	field.Value = val
	return field
}

//DerivativeCFICode is a STRING field
type DerivativeCFICode struct{ message.StringValue }

//Tag returns tag.DerivativeCFICode (1248)
func (f DerivativeCFICode) Tag() fix.Tag { return tag.DerivativeCFICode }

//BuildDerivativeCFICode returns a new DerivativeCFICode initialized with val
func BuildDerivativeCFICode(val string) DerivativeCFICode {
	var field DerivativeCFICode
	field.Value = val
	return field
}

//DerivativeCapPrice is a PRICE field
type DerivativeCapPrice struct{ message.PriceValue }

//Tag returns tag.DerivativeCapPrice (1321)
func (f DerivativeCapPrice) Tag() fix.Tag { return tag.DerivativeCapPrice }

//BuildDerivativeCapPrice returns a new DerivativeCapPrice initialized with val
func BuildDerivativeCapPrice(val float64) DerivativeCapPrice {
	var field DerivativeCapPrice
	field.Value = val
	return field
}

//DerivativeContractMultiplier is a FLOAT field
type DerivativeContractMultiplier struct{ message.FloatValue }

//Tag returns tag.DerivativeContractMultiplier (1266)
func (f DerivativeContractMultiplier) Tag() fix.Tag { return tag.DerivativeContractMultiplier }

//BuildDerivativeContractMultiplier returns a new DerivativeContractMultiplier initialized with val
func BuildDerivativeContractMultiplier(val float64) DerivativeContractMultiplier {
	var field DerivativeContractMultiplier
	field.Value = val
	return field
}

//DerivativeContractMultiplierUnit is a INT field
type DerivativeContractMultiplierUnit struct{ message.IntValue }

//Tag returns tag.DerivativeContractMultiplierUnit (1438)
func (f DerivativeContractMultiplierUnit) Tag() fix.Tag { return tag.DerivativeContractMultiplierUnit }

//BuildDerivativeContractMultiplierUnit returns a new DerivativeContractMultiplierUnit initialized with val
func BuildDerivativeContractMultiplierUnit(val int) DerivativeContractMultiplierUnit {
	var field DerivativeContractMultiplierUnit
	field.Value = val
	return field
}

//DerivativeContractSettlMonth is a MONTHYEAR field
type DerivativeContractSettlMonth struct{ message.MonthYearValue }

//Tag returns tag.DerivativeContractSettlMonth (1285)
func (f DerivativeContractSettlMonth) Tag() fix.Tag { return tag.DerivativeContractSettlMonth }

//BuildDerivativeContractSettlMonth returns a new DerivativeContractSettlMonth initialized with val
func BuildDerivativeContractSettlMonth(val string) DerivativeContractSettlMonth {
	var field DerivativeContractSettlMonth
	field.Value = val
	return field
}

//DerivativeCountryOfIssue is a COUNTRY field
type DerivativeCountryOfIssue struct{ message.CountryValue }

//Tag returns tag.DerivativeCountryOfIssue (1258)
func (f DerivativeCountryOfIssue) Tag() fix.Tag { return tag.DerivativeCountryOfIssue }

//BuildDerivativeCountryOfIssue returns a new DerivativeCountryOfIssue initialized with val
func BuildDerivativeCountryOfIssue(val string) DerivativeCountryOfIssue {
	var field DerivativeCountryOfIssue
	field.Value = val
	return field
}

//DerivativeEncodedIssuer is a DATA field
type DerivativeEncodedIssuer struct{ message.DataValue }

//Tag returns tag.DerivativeEncodedIssuer (1278)
func (f DerivativeEncodedIssuer) Tag() fix.Tag { return tag.DerivativeEncodedIssuer }

//BuildDerivativeEncodedIssuer returns a new DerivativeEncodedIssuer initialized with val
func BuildDerivativeEncodedIssuer(val string) DerivativeEncodedIssuer {
	var field DerivativeEncodedIssuer
	field.Value = val
	return field
}

//DerivativeEncodedIssuerLen is a LENGTH field
type DerivativeEncodedIssuerLen struct{ message.LengthValue }

//Tag returns tag.DerivativeEncodedIssuerLen (1277)
func (f DerivativeEncodedIssuerLen) Tag() fix.Tag { return tag.DerivativeEncodedIssuerLen }

//BuildDerivativeEncodedIssuerLen returns a new DerivativeEncodedIssuerLen initialized with val
func BuildDerivativeEncodedIssuerLen(val int) DerivativeEncodedIssuerLen {
	var field DerivativeEncodedIssuerLen
	field.Value = val
	return field
}

//DerivativeEncodedSecurityDesc is a DATA field
type DerivativeEncodedSecurityDesc struct{ message.DataValue }

//Tag returns tag.DerivativeEncodedSecurityDesc (1281)
func (f DerivativeEncodedSecurityDesc) Tag() fix.Tag { return tag.DerivativeEncodedSecurityDesc }

//BuildDerivativeEncodedSecurityDesc returns a new DerivativeEncodedSecurityDesc initialized with val
func BuildDerivativeEncodedSecurityDesc(val string) DerivativeEncodedSecurityDesc {
	var field DerivativeEncodedSecurityDesc
	field.Value = val
	return field
}

//DerivativeEncodedSecurityDescLen is a LENGTH field
type DerivativeEncodedSecurityDescLen struct{ message.LengthValue }

//Tag returns tag.DerivativeEncodedSecurityDescLen (1280)
func (f DerivativeEncodedSecurityDescLen) Tag() fix.Tag { return tag.DerivativeEncodedSecurityDescLen }

//BuildDerivativeEncodedSecurityDescLen returns a new DerivativeEncodedSecurityDescLen initialized with val
func BuildDerivativeEncodedSecurityDescLen(val int) DerivativeEncodedSecurityDescLen {
	var field DerivativeEncodedSecurityDescLen
	field.Value = val
	return field
}

//DerivativeEventDate is a LOCALMKTDATE field
type DerivativeEventDate struct{ message.LocalMktDateValue }

//Tag returns tag.DerivativeEventDate (1288)
func (f DerivativeEventDate) Tag() fix.Tag { return tag.DerivativeEventDate }

//BuildDerivativeEventDate returns a new DerivativeEventDate initialized with val
func BuildDerivativeEventDate(val string) DerivativeEventDate {
	var field DerivativeEventDate
	field.Value = val
	return field
}

//DerivativeEventPx is a PRICE field
type DerivativeEventPx struct{ message.PriceValue }

//Tag returns tag.DerivativeEventPx (1290)
func (f DerivativeEventPx) Tag() fix.Tag { return tag.DerivativeEventPx }

//BuildDerivativeEventPx returns a new DerivativeEventPx initialized with val
func BuildDerivativeEventPx(val float64) DerivativeEventPx {
	var field DerivativeEventPx
	field.Value = val
	return field
}

//DerivativeEventText is a STRING field
type DerivativeEventText struct{ message.StringValue }

//Tag returns tag.DerivativeEventText (1291)
func (f DerivativeEventText) Tag() fix.Tag { return tag.DerivativeEventText }

//BuildDerivativeEventText returns a new DerivativeEventText initialized with val
func BuildDerivativeEventText(val string) DerivativeEventText {
	var field DerivativeEventText
	field.Value = val
	return field
}

//DerivativeEventTime is a UTCTIMESTAMP field
type DerivativeEventTime struct{ message.UTCTimestampValue }

//Tag returns tag.DerivativeEventTime (1289)
func (f DerivativeEventTime) Tag() fix.Tag { return tag.DerivativeEventTime }

//DerivativeEventType is a INT field
type DerivativeEventType struct{ message.IntValue }

//Tag returns tag.DerivativeEventType (1287)
func (f DerivativeEventType) Tag() fix.Tag { return tag.DerivativeEventType }

//BuildDerivativeEventType returns a new DerivativeEventType initialized with val
func BuildDerivativeEventType(val int) DerivativeEventType {
	var field DerivativeEventType
	field.Value = val
	return field
}

//DerivativeExerciseStyle is a CHAR field
type DerivativeExerciseStyle struct{ message.CharValue }

//Tag returns tag.DerivativeExerciseStyle (1299)
func (f DerivativeExerciseStyle) Tag() fix.Tag { return tag.DerivativeExerciseStyle }

//BuildDerivativeExerciseStyle returns a new DerivativeExerciseStyle initialized with val
func BuildDerivativeExerciseStyle(val string) DerivativeExerciseStyle {
	var field DerivativeExerciseStyle
	field.Value = val
	return field
}

//DerivativeFloorPrice is a PRICE field
type DerivativeFloorPrice struct{ message.PriceValue }

//Tag returns tag.DerivativeFloorPrice (1322)
func (f DerivativeFloorPrice) Tag() fix.Tag { return tag.DerivativeFloorPrice }

//BuildDerivativeFloorPrice returns a new DerivativeFloorPrice initialized with val
func BuildDerivativeFloorPrice(val float64) DerivativeFloorPrice {
	var field DerivativeFloorPrice
	field.Value = val
	return field
}

//DerivativeFlowScheduleType is a INT field
type DerivativeFlowScheduleType struct{ message.IntValue }

//Tag returns tag.DerivativeFlowScheduleType (1442)
func (f DerivativeFlowScheduleType) Tag() fix.Tag { return tag.DerivativeFlowScheduleType }

//BuildDerivativeFlowScheduleType returns a new DerivativeFlowScheduleType initialized with val
func BuildDerivativeFlowScheduleType(val int) DerivativeFlowScheduleType {
	var field DerivativeFlowScheduleType
	field.Value = val
	return field
}

//DerivativeFuturesValuationMethod is a STRING field
type DerivativeFuturesValuationMethod struct{ message.StringValue }

//Tag returns tag.DerivativeFuturesValuationMethod (1319)
func (f DerivativeFuturesValuationMethod) Tag() fix.Tag { return tag.DerivativeFuturesValuationMethod }

//BuildDerivativeFuturesValuationMethod returns a new DerivativeFuturesValuationMethod initialized with val
func BuildDerivativeFuturesValuationMethod(val string) DerivativeFuturesValuationMethod {
	var field DerivativeFuturesValuationMethod
	field.Value = val
	return field
}

//DerivativeInstrAttribType is a INT field
type DerivativeInstrAttribType struct{ message.IntValue }

//Tag returns tag.DerivativeInstrAttribType (1313)
func (f DerivativeInstrAttribType) Tag() fix.Tag { return tag.DerivativeInstrAttribType }

//BuildDerivativeInstrAttribType returns a new DerivativeInstrAttribType initialized with val
func BuildDerivativeInstrAttribType(val int) DerivativeInstrAttribType {
	var field DerivativeInstrAttribType
	field.Value = val
	return field
}

//DerivativeInstrAttribValue is a STRING field
type DerivativeInstrAttribValue struct{ message.StringValue }

//Tag returns tag.DerivativeInstrAttribValue (1314)
func (f DerivativeInstrAttribValue) Tag() fix.Tag { return tag.DerivativeInstrAttribValue }

//BuildDerivativeInstrAttribValue returns a new DerivativeInstrAttribValue initialized with val
func BuildDerivativeInstrAttribValue(val string) DerivativeInstrAttribValue {
	var field DerivativeInstrAttribValue
	field.Value = val
	return field
}

//DerivativeInstrRegistry is a STRING field
type DerivativeInstrRegistry struct{ message.StringValue }

//Tag returns tag.DerivativeInstrRegistry (1257)
func (f DerivativeInstrRegistry) Tag() fix.Tag { return tag.DerivativeInstrRegistry }

//BuildDerivativeInstrRegistry returns a new DerivativeInstrRegistry initialized with val
func BuildDerivativeInstrRegistry(val string) DerivativeInstrRegistry {
	var field DerivativeInstrRegistry
	field.Value = val
	return field
}

//DerivativeInstrmtAssignmentMethod is a CHAR field
type DerivativeInstrmtAssignmentMethod struct{ message.CharValue }

//Tag returns tag.DerivativeInstrmtAssignmentMethod (1255)
func (f DerivativeInstrmtAssignmentMethod) Tag() fix.Tag { return tag.DerivativeInstrmtAssignmentMethod }

//BuildDerivativeInstrmtAssignmentMethod returns a new DerivativeInstrmtAssignmentMethod initialized with val
func BuildDerivativeInstrmtAssignmentMethod(val string) DerivativeInstrmtAssignmentMethod {
	var field DerivativeInstrmtAssignmentMethod
	field.Value = val
	return field
}

//DerivativeInstrumentPartyID is a STRING field
type DerivativeInstrumentPartyID struct{ message.StringValue }

//Tag returns tag.DerivativeInstrumentPartyID (1293)
func (f DerivativeInstrumentPartyID) Tag() fix.Tag { return tag.DerivativeInstrumentPartyID }

//BuildDerivativeInstrumentPartyID returns a new DerivativeInstrumentPartyID initialized with val
func BuildDerivativeInstrumentPartyID(val string) DerivativeInstrumentPartyID {
	var field DerivativeInstrumentPartyID
	field.Value = val
	return field
}

//DerivativeInstrumentPartyIDSource is a STRING field
type DerivativeInstrumentPartyIDSource struct{ message.StringValue }

//Tag returns tag.DerivativeInstrumentPartyIDSource (1294)
func (f DerivativeInstrumentPartyIDSource) Tag() fix.Tag { return tag.DerivativeInstrumentPartyIDSource }

//BuildDerivativeInstrumentPartyIDSource returns a new DerivativeInstrumentPartyIDSource initialized with val
func BuildDerivativeInstrumentPartyIDSource(val string) DerivativeInstrumentPartyIDSource {
	var field DerivativeInstrumentPartyIDSource
	field.Value = val
	return field
}

//DerivativeInstrumentPartyRole is a INT field
type DerivativeInstrumentPartyRole struct{ message.IntValue }

//Tag returns tag.DerivativeInstrumentPartyRole (1295)
func (f DerivativeInstrumentPartyRole) Tag() fix.Tag { return tag.DerivativeInstrumentPartyRole }

//BuildDerivativeInstrumentPartyRole returns a new DerivativeInstrumentPartyRole initialized with val
func BuildDerivativeInstrumentPartyRole(val int) DerivativeInstrumentPartyRole {
	var field DerivativeInstrumentPartyRole
	field.Value = val
	return field
}

//DerivativeInstrumentPartySubID is a STRING field
type DerivativeInstrumentPartySubID struct{ message.StringValue }

//Tag returns tag.DerivativeInstrumentPartySubID (1297)
func (f DerivativeInstrumentPartySubID) Tag() fix.Tag { return tag.DerivativeInstrumentPartySubID }

//BuildDerivativeInstrumentPartySubID returns a new DerivativeInstrumentPartySubID initialized with val
func BuildDerivativeInstrumentPartySubID(val string) DerivativeInstrumentPartySubID {
	var field DerivativeInstrumentPartySubID
	field.Value = val
	return field
}

//DerivativeInstrumentPartySubIDType is a INT field
type DerivativeInstrumentPartySubIDType struct{ message.IntValue }

//Tag returns tag.DerivativeInstrumentPartySubIDType (1298)
func (f DerivativeInstrumentPartySubIDType) Tag() fix.Tag {
	return tag.DerivativeInstrumentPartySubIDType
}

//BuildDerivativeInstrumentPartySubIDType returns a new DerivativeInstrumentPartySubIDType initialized with val
func BuildDerivativeInstrumentPartySubIDType(val int) DerivativeInstrumentPartySubIDType {
	var field DerivativeInstrumentPartySubIDType
	field.Value = val
	return field
}

//DerivativeIssueDate is a LOCALMKTDATE field
type DerivativeIssueDate struct{ message.LocalMktDateValue }

//Tag returns tag.DerivativeIssueDate (1276)
func (f DerivativeIssueDate) Tag() fix.Tag { return tag.DerivativeIssueDate }

//BuildDerivativeIssueDate returns a new DerivativeIssueDate initialized with val
func BuildDerivativeIssueDate(val string) DerivativeIssueDate {
	var field DerivativeIssueDate
	field.Value = val
	return field
}

//DerivativeIssuer is a STRING field
type DerivativeIssuer struct{ message.StringValue }

//Tag returns tag.DerivativeIssuer (1275)
func (f DerivativeIssuer) Tag() fix.Tag { return tag.DerivativeIssuer }

//BuildDerivativeIssuer returns a new DerivativeIssuer initialized with val
func BuildDerivativeIssuer(val string) DerivativeIssuer {
	var field DerivativeIssuer
	field.Value = val
	return field
}

//DerivativeListMethod is a INT field
type DerivativeListMethod struct{ message.IntValue }

//Tag returns tag.DerivativeListMethod (1320)
func (f DerivativeListMethod) Tag() fix.Tag { return tag.DerivativeListMethod }

//BuildDerivativeListMethod returns a new DerivativeListMethod initialized with val
func BuildDerivativeListMethod(val int) DerivativeListMethod {
	var field DerivativeListMethod
	field.Value = val
	return field
}

//DerivativeLocaleOfIssue is a STRING field
type DerivativeLocaleOfIssue struct{ message.StringValue }

//Tag returns tag.DerivativeLocaleOfIssue (1260)
func (f DerivativeLocaleOfIssue) Tag() fix.Tag { return tag.DerivativeLocaleOfIssue }

//BuildDerivativeLocaleOfIssue returns a new DerivativeLocaleOfIssue initialized with val
func BuildDerivativeLocaleOfIssue(val string) DerivativeLocaleOfIssue {
	var field DerivativeLocaleOfIssue
	field.Value = val
	return field
}

//DerivativeMaturityDate is a LOCALMKTDATE field
type DerivativeMaturityDate struct{ message.LocalMktDateValue }

//Tag returns tag.DerivativeMaturityDate (1252)
func (f DerivativeMaturityDate) Tag() fix.Tag { return tag.DerivativeMaturityDate }

//BuildDerivativeMaturityDate returns a new DerivativeMaturityDate initialized with val
func BuildDerivativeMaturityDate(val string) DerivativeMaturityDate {
	var field DerivativeMaturityDate
	field.Value = val
	return field
}

//DerivativeMaturityMonthYear is a MONTHYEAR field
type DerivativeMaturityMonthYear struct{ message.MonthYearValue }

//Tag returns tag.DerivativeMaturityMonthYear (1251)
func (f DerivativeMaturityMonthYear) Tag() fix.Tag { return tag.DerivativeMaturityMonthYear }

//BuildDerivativeMaturityMonthYear returns a new DerivativeMaturityMonthYear initialized with val
func BuildDerivativeMaturityMonthYear(val string) DerivativeMaturityMonthYear {
	var field DerivativeMaturityMonthYear
	field.Value = val
	return field
}

//DerivativeMaturityTime is a TZTIMEONLY field
type DerivativeMaturityTime struct{ message.TZTimeOnlyValue }

//Tag returns tag.DerivativeMaturityTime (1253)
func (f DerivativeMaturityTime) Tag() fix.Tag { return tag.DerivativeMaturityTime }

//DerivativeMinPriceIncrement is a FLOAT field
type DerivativeMinPriceIncrement struct{ message.FloatValue }

//Tag returns tag.DerivativeMinPriceIncrement (1267)
func (f DerivativeMinPriceIncrement) Tag() fix.Tag { return tag.DerivativeMinPriceIncrement }

//BuildDerivativeMinPriceIncrement returns a new DerivativeMinPriceIncrement initialized with val
func BuildDerivativeMinPriceIncrement(val float64) DerivativeMinPriceIncrement {
	var field DerivativeMinPriceIncrement
	field.Value = val
	return field
}

//DerivativeMinPriceIncrementAmount is a AMT field
type DerivativeMinPriceIncrementAmount struct{ message.AmtValue }

//Tag returns tag.DerivativeMinPriceIncrementAmount (1268)
func (f DerivativeMinPriceIncrementAmount) Tag() fix.Tag { return tag.DerivativeMinPriceIncrementAmount }

//BuildDerivativeMinPriceIncrementAmount returns a new DerivativeMinPriceIncrementAmount initialized with val
func BuildDerivativeMinPriceIncrementAmount(val float64) DerivativeMinPriceIncrementAmount {
	var field DerivativeMinPriceIncrementAmount
	field.Value = val
	return field
}

//DerivativeNTPositionLimit is a INT field
type DerivativeNTPositionLimit struct{ message.IntValue }

//Tag returns tag.DerivativeNTPositionLimit (1274)
func (f DerivativeNTPositionLimit) Tag() fix.Tag { return tag.DerivativeNTPositionLimit }

//BuildDerivativeNTPositionLimit returns a new DerivativeNTPositionLimit initialized with val
func BuildDerivativeNTPositionLimit(val int) DerivativeNTPositionLimit {
	var field DerivativeNTPositionLimit
	field.Value = val
	return field
}

//DerivativeOptAttribute is a CHAR field
type DerivativeOptAttribute struct{ message.CharValue }

//Tag returns tag.DerivativeOptAttribute (1265)
func (f DerivativeOptAttribute) Tag() fix.Tag { return tag.DerivativeOptAttribute }

//BuildDerivativeOptAttribute returns a new DerivativeOptAttribute initialized with val
func BuildDerivativeOptAttribute(val string) DerivativeOptAttribute {
	var field DerivativeOptAttribute
	field.Value = val
	return field
}

//DerivativeOptPayAmount is a AMT field
type DerivativeOptPayAmount struct{ message.AmtValue }

//Tag returns tag.DerivativeOptPayAmount (1225)
func (f DerivativeOptPayAmount) Tag() fix.Tag { return tag.DerivativeOptPayAmount }

//BuildDerivativeOptPayAmount returns a new DerivativeOptPayAmount initialized with val
func BuildDerivativeOptPayAmount(val float64) DerivativeOptPayAmount {
	var field DerivativeOptPayAmount
	field.Value = val
	return field
}

//DerivativePositionLimit is a INT field
type DerivativePositionLimit struct{ message.IntValue }

//Tag returns tag.DerivativePositionLimit (1273)
func (f DerivativePositionLimit) Tag() fix.Tag { return tag.DerivativePositionLimit }

//BuildDerivativePositionLimit returns a new DerivativePositionLimit initialized with val
func BuildDerivativePositionLimit(val int) DerivativePositionLimit {
	var field DerivativePositionLimit
	field.Value = val
	return field
}

//DerivativePriceQuoteMethod is a STRING field
type DerivativePriceQuoteMethod struct{ message.StringValue }

//Tag returns tag.DerivativePriceQuoteMethod (1318)
func (f DerivativePriceQuoteMethod) Tag() fix.Tag { return tag.DerivativePriceQuoteMethod }

//BuildDerivativePriceQuoteMethod returns a new DerivativePriceQuoteMethod initialized with val
func BuildDerivativePriceQuoteMethod(val string) DerivativePriceQuoteMethod {
	var field DerivativePriceQuoteMethod
	field.Value = val
	return field
}

//DerivativePriceUnitOfMeasure is a STRING field
type DerivativePriceUnitOfMeasure struct{ message.StringValue }

//Tag returns tag.DerivativePriceUnitOfMeasure (1315)
func (f DerivativePriceUnitOfMeasure) Tag() fix.Tag { return tag.DerivativePriceUnitOfMeasure }

//BuildDerivativePriceUnitOfMeasure returns a new DerivativePriceUnitOfMeasure initialized with val
func BuildDerivativePriceUnitOfMeasure(val string) DerivativePriceUnitOfMeasure {
	var field DerivativePriceUnitOfMeasure
	field.Value = val
	return field
}

//DerivativePriceUnitOfMeasureQty is a QTY field
type DerivativePriceUnitOfMeasureQty struct{ message.QtyValue }

//Tag returns tag.DerivativePriceUnitOfMeasureQty (1316)
func (f DerivativePriceUnitOfMeasureQty) Tag() fix.Tag { return tag.DerivativePriceUnitOfMeasureQty }

//BuildDerivativePriceUnitOfMeasureQty returns a new DerivativePriceUnitOfMeasureQty initialized with val
func BuildDerivativePriceUnitOfMeasureQty(val float64) DerivativePriceUnitOfMeasureQty {
	var field DerivativePriceUnitOfMeasureQty
	field.Value = val
	return field
}

//DerivativeProduct is a INT field
type DerivativeProduct struct{ message.IntValue }

//Tag returns tag.DerivativeProduct (1246)
func (f DerivativeProduct) Tag() fix.Tag { return tag.DerivativeProduct }

//BuildDerivativeProduct returns a new DerivativeProduct initialized with val
func BuildDerivativeProduct(val int) DerivativeProduct {
	var field DerivativeProduct
	field.Value = val
	return field
}

//DerivativeProductComplex is a STRING field
type DerivativeProductComplex struct{ message.StringValue }

//Tag returns tag.DerivativeProductComplex (1228)
func (f DerivativeProductComplex) Tag() fix.Tag { return tag.DerivativeProductComplex }

//BuildDerivativeProductComplex returns a new DerivativeProductComplex initialized with val
func BuildDerivativeProductComplex(val string) DerivativeProductComplex {
	var field DerivativeProductComplex
	field.Value = val
	return field
}

//DerivativePutOrCall is a INT field
type DerivativePutOrCall struct{ message.IntValue }

//Tag returns tag.DerivativePutOrCall (1323)
func (f DerivativePutOrCall) Tag() fix.Tag { return tag.DerivativePutOrCall }

//BuildDerivativePutOrCall returns a new DerivativePutOrCall initialized with val
func BuildDerivativePutOrCall(val int) DerivativePutOrCall {
	var field DerivativePutOrCall
	field.Value = val
	return field
}

//DerivativeSecurityAltID is a STRING field
type DerivativeSecurityAltID struct{ message.StringValue }

//Tag returns tag.DerivativeSecurityAltID (1219)
func (f DerivativeSecurityAltID) Tag() fix.Tag { return tag.DerivativeSecurityAltID }

//BuildDerivativeSecurityAltID returns a new DerivativeSecurityAltID initialized with val
func BuildDerivativeSecurityAltID(val string) DerivativeSecurityAltID {
	var field DerivativeSecurityAltID
	field.Value = val
	return field
}

//DerivativeSecurityAltIDSource is a STRING field
type DerivativeSecurityAltIDSource struct{ message.StringValue }

//Tag returns tag.DerivativeSecurityAltIDSource (1220)
func (f DerivativeSecurityAltIDSource) Tag() fix.Tag { return tag.DerivativeSecurityAltIDSource }

//BuildDerivativeSecurityAltIDSource returns a new DerivativeSecurityAltIDSource initialized with val
func BuildDerivativeSecurityAltIDSource(val string) DerivativeSecurityAltIDSource {
	var field DerivativeSecurityAltIDSource
	field.Value = val
	return field
}

//DerivativeSecurityDesc is a STRING field
type DerivativeSecurityDesc struct{ message.StringValue }

//Tag returns tag.DerivativeSecurityDesc (1279)
func (f DerivativeSecurityDesc) Tag() fix.Tag { return tag.DerivativeSecurityDesc }

//BuildDerivativeSecurityDesc returns a new DerivativeSecurityDesc initialized with val
func BuildDerivativeSecurityDesc(val string) DerivativeSecurityDesc {
	var field DerivativeSecurityDesc
	field.Value = val
	return field
}

//DerivativeSecurityExchange is a EXCHANGE field
type DerivativeSecurityExchange struct{ message.ExchangeValue }

//Tag returns tag.DerivativeSecurityExchange (1272)
func (f DerivativeSecurityExchange) Tag() fix.Tag { return tag.DerivativeSecurityExchange }

//BuildDerivativeSecurityExchange returns a new DerivativeSecurityExchange initialized with val
func BuildDerivativeSecurityExchange(val string) DerivativeSecurityExchange {
	var field DerivativeSecurityExchange
	field.Value = val
	return field
}

//DerivativeSecurityGroup is a STRING field
type DerivativeSecurityGroup struct{ message.StringValue }

//Tag returns tag.DerivativeSecurityGroup (1247)
func (f DerivativeSecurityGroup) Tag() fix.Tag { return tag.DerivativeSecurityGroup }

//BuildDerivativeSecurityGroup returns a new DerivativeSecurityGroup initialized with val
func BuildDerivativeSecurityGroup(val string) DerivativeSecurityGroup {
	var field DerivativeSecurityGroup
	field.Value = val
	return field
}

//DerivativeSecurityID is a STRING field
type DerivativeSecurityID struct{ message.StringValue }

//Tag returns tag.DerivativeSecurityID (1216)
func (f DerivativeSecurityID) Tag() fix.Tag { return tag.DerivativeSecurityID }

//BuildDerivativeSecurityID returns a new DerivativeSecurityID initialized with val
func BuildDerivativeSecurityID(val string) DerivativeSecurityID {
	var field DerivativeSecurityID
	field.Value = val
	return field
}

//DerivativeSecurityIDSource is a STRING field
type DerivativeSecurityIDSource struct{ message.StringValue }

//Tag returns tag.DerivativeSecurityIDSource (1217)
func (f DerivativeSecurityIDSource) Tag() fix.Tag { return tag.DerivativeSecurityIDSource }

//BuildDerivativeSecurityIDSource returns a new DerivativeSecurityIDSource initialized with val
func BuildDerivativeSecurityIDSource(val string) DerivativeSecurityIDSource {
	var field DerivativeSecurityIDSource
	field.Value = val
	return field
}

//DerivativeSecurityListRequestType is a INT field
type DerivativeSecurityListRequestType struct{ message.IntValue }

//Tag returns tag.DerivativeSecurityListRequestType (1307)
func (f DerivativeSecurityListRequestType) Tag() fix.Tag { return tag.DerivativeSecurityListRequestType }

//BuildDerivativeSecurityListRequestType returns a new DerivativeSecurityListRequestType initialized with val
func BuildDerivativeSecurityListRequestType(val int) DerivativeSecurityListRequestType {
	var field DerivativeSecurityListRequestType
	field.Value = val
	return field
}

//DerivativeSecurityStatus is a STRING field
type DerivativeSecurityStatus struct{ message.StringValue }

//Tag returns tag.DerivativeSecurityStatus (1256)
func (f DerivativeSecurityStatus) Tag() fix.Tag { return tag.DerivativeSecurityStatus }

//BuildDerivativeSecurityStatus returns a new DerivativeSecurityStatus initialized with val
func BuildDerivativeSecurityStatus(val string) DerivativeSecurityStatus {
	var field DerivativeSecurityStatus
	field.Value = val
	return field
}

//DerivativeSecuritySubType is a STRING field
type DerivativeSecuritySubType struct{ message.StringValue }

//Tag returns tag.DerivativeSecuritySubType (1250)
func (f DerivativeSecuritySubType) Tag() fix.Tag { return tag.DerivativeSecuritySubType }

//BuildDerivativeSecuritySubType returns a new DerivativeSecuritySubType initialized with val
func BuildDerivativeSecuritySubType(val string) DerivativeSecuritySubType {
	var field DerivativeSecuritySubType
	field.Value = val
	return field
}

//DerivativeSecurityType is a STRING field
type DerivativeSecurityType struct{ message.StringValue }

//Tag returns tag.DerivativeSecurityType (1249)
func (f DerivativeSecurityType) Tag() fix.Tag { return tag.DerivativeSecurityType }

//BuildDerivativeSecurityType returns a new DerivativeSecurityType initialized with val
func BuildDerivativeSecurityType(val string) DerivativeSecurityType {
	var field DerivativeSecurityType
	field.Value = val
	return field
}

//DerivativeSecurityXML is a DATA field
type DerivativeSecurityXML struct{ message.DataValue }

//Tag returns tag.DerivativeSecurityXML (1283)
func (f DerivativeSecurityXML) Tag() fix.Tag { return tag.DerivativeSecurityXML }

//BuildDerivativeSecurityXML returns a new DerivativeSecurityXML initialized with val
func BuildDerivativeSecurityXML(val string) DerivativeSecurityXML {
	var field DerivativeSecurityXML
	field.Value = val
	return field
}

//DerivativeSecurityXMLLen is a LENGTH field
type DerivativeSecurityXMLLen struct{ message.LengthValue }

//Tag returns tag.DerivativeSecurityXMLLen (1282)
func (f DerivativeSecurityXMLLen) Tag() fix.Tag { return tag.DerivativeSecurityXMLLen }

//BuildDerivativeSecurityXMLLen returns a new DerivativeSecurityXMLLen initialized with val
func BuildDerivativeSecurityXMLLen(val int) DerivativeSecurityXMLLen {
	var field DerivativeSecurityXMLLen
	field.Value = val
	return field
}

//DerivativeSecurityXMLSchema is a STRING field
type DerivativeSecurityXMLSchema struct{ message.StringValue }

//Tag returns tag.DerivativeSecurityXMLSchema (1284)
func (f DerivativeSecurityXMLSchema) Tag() fix.Tag { return tag.DerivativeSecurityXMLSchema }

//BuildDerivativeSecurityXMLSchema returns a new DerivativeSecurityXMLSchema initialized with val
func BuildDerivativeSecurityXMLSchema(val string) DerivativeSecurityXMLSchema {
	var field DerivativeSecurityXMLSchema
	field.Value = val
	return field
}

//DerivativeSettlMethod is a CHAR field
type DerivativeSettlMethod struct{ message.CharValue }

//Tag returns tag.DerivativeSettlMethod (1317)
func (f DerivativeSettlMethod) Tag() fix.Tag { return tag.DerivativeSettlMethod }

//BuildDerivativeSettlMethod returns a new DerivativeSettlMethod initialized with val
func BuildDerivativeSettlMethod(val string) DerivativeSettlMethod {
	var field DerivativeSettlMethod
	field.Value = val
	return field
}

//DerivativeSettleOnOpenFlag is a STRING field
type DerivativeSettleOnOpenFlag struct{ message.StringValue }

//Tag returns tag.DerivativeSettleOnOpenFlag (1254)
func (f DerivativeSettleOnOpenFlag) Tag() fix.Tag { return tag.DerivativeSettleOnOpenFlag }

//BuildDerivativeSettleOnOpenFlag returns a new DerivativeSettleOnOpenFlag initialized with val
func BuildDerivativeSettleOnOpenFlag(val string) DerivativeSettleOnOpenFlag {
	var field DerivativeSettleOnOpenFlag
	field.Value = val
	return field
}

//DerivativeStateOrProvinceOfIssue is a STRING field
type DerivativeStateOrProvinceOfIssue struct{ message.StringValue }

//Tag returns tag.DerivativeStateOrProvinceOfIssue (1259)
func (f DerivativeStateOrProvinceOfIssue) Tag() fix.Tag { return tag.DerivativeStateOrProvinceOfIssue }

//BuildDerivativeStateOrProvinceOfIssue returns a new DerivativeStateOrProvinceOfIssue initialized with val
func BuildDerivativeStateOrProvinceOfIssue(val string) DerivativeStateOrProvinceOfIssue {
	var field DerivativeStateOrProvinceOfIssue
	field.Value = val
	return field
}

//DerivativeStrikeCurrency is a CURRENCY field
type DerivativeStrikeCurrency struct{ message.CurrencyValue }

//Tag returns tag.DerivativeStrikeCurrency (1262)
func (f DerivativeStrikeCurrency) Tag() fix.Tag { return tag.DerivativeStrikeCurrency }

//BuildDerivativeStrikeCurrency returns a new DerivativeStrikeCurrency initialized with val
func BuildDerivativeStrikeCurrency(val string) DerivativeStrikeCurrency {
	var field DerivativeStrikeCurrency
	field.Value = val
	return field
}

//DerivativeStrikeMultiplier is a FLOAT field
type DerivativeStrikeMultiplier struct{ message.FloatValue }

//Tag returns tag.DerivativeStrikeMultiplier (1263)
func (f DerivativeStrikeMultiplier) Tag() fix.Tag { return tag.DerivativeStrikeMultiplier }

//BuildDerivativeStrikeMultiplier returns a new DerivativeStrikeMultiplier initialized with val
func BuildDerivativeStrikeMultiplier(val float64) DerivativeStrikeMultiplier {
	var field DerivativeStrikeMultiplier
	field.Value = val
	return field
}

//DerivativeStrikePrice is a PRICE field
type DerivativeStrikePrice struct{ message.PriceValue }

//Tag returns tag.DerivativeStrikePrice (1261)
func (f DerivativeStrikePrice) Tag() fix.Tag { return tag.DerivativeStrikePrice }

//BuildDerivativeStrikePrice returns a new DerivativeStrikePrice initialized with val
func BuildDerivativeStrikePrice(val float64) DerivativeStrikePrice {
	var field DerivativeStrikePrice
	field.Value = val
	return field
}

//DerivativeStrikeValue is a FLOAT field
type DerivativeStrikeValue struct{ message.FloatValue }

//Tag returns tag.DerivativeStrikeValue (1264)
func (f DerivativeStrikeValue) Tag() fix.Tag { return tag.DerivativeStrikeValue }

//BuildDerivativeStrikeValue returns a new DerivativeStrikeValue initialized with val
func BuildDerivativeStrikeValue(val float64) DerivativeStrikeValue {
	var field DerivativeStrikeValue
	field.Value = val
	return field
}

//DerivativeSymbol is a STRING field
type DerivativeSymbol struct{ message.StringValue }

//Tag returns tag.DerivativeSymbol (1214)
func (f DerivativeSymbol) Tag() fix.Tag { return tag.DerivativeSymbol }

//BuildDerivativeSymbol returns a new DerivativeSymbol initialized with val
func BuildDerivativeSymbol(val string) DerivativeSymbol {
	var field DerivativeSymbol
	field.Value = val
	return field
}

//DerivativeSymbolSfx is a STRING field
type DerivativeSymbolSfx struct{ message.StringValue }

//Tag returns tag.DerivativeSymbolSfx (1215)
func (f DerivativeSymbolSfx) Tag() fix.Tag { return tag.DerivativeSymbolSfx }

//BuildDerivativeSymbolSfx returns a new DerivativeSymbolSfx initialized with val
func BuildDerivativeSymbolSfx(val string) DerivativeSymbolSfx {
	var field DerivativeSymbolSfx
	field.Value = val
	return field
}

//DerivativeTimeUnit is a STRING field
type DerivativeTimeUnit struct{ message.StringValue }

//Tag returns tag.DerivativeTimeUnit (1271)
func (f DerivativeTimeUnit) Tag() fix.Tag { return tag.DerivativeTimeUnit }

//BuildDerivativeTimeUnit returns a new DerivativeTimeUnit initialized with val
func BuildDerivativeTimeUnit(val string) DerivativeTimeUnit {
	var field DerivativeTimeUnit
	field.Value = val
	return field
}

//DerivativeUnitOfMeasure is a STRING field
type DerivativeUnitOfMeasure struct{ message.StringValue }

//Tag returns tag.DerivativeUnitOfMeasure (1269)
func (f DerivativeUnitOfMeasure) Tag() fix.Tag { return tag.DerivativeUnitOfMeasure }

//BuildDerivativeUnitOfMeasure returns a new DerivativeUnitOfMeasure initialized with val
func BuildDerivativeUnitOfMeasure(val string) DerivativeUnitOfMeasure {
	var field DerivativeUnitOfMeasure
	field.Value = val
	return field
}

//DerivativeUnitOfMeasureQty is a QTY field
type DerivativeUnitOfMeasureQty struct{ message.QtyValue }

//Tag returns tag.DerivativeUnitOfMeasureQty (1270)
func (f DerivativeUnitOfMeasureQty) Tag() fix.Tag { return tag.DerivativeUnitOfMeasureQty }

//BuildDerivativeUnitOfMeasureQty returns a new DerivativeUnitOfMeasureQty initialized with val
func BuildDerivativeUnitOfMeasureQty(val float64) DerivativeUnitOfMeasureQty {
	var field DerivativeUnitOfMeasureQty
	field.Value = val
	return field
}

//DerivativeValuationMethod is a STRING field
type DerivativeValuationMethod struct{ message.StringValue }

//Tag returns tag.DerivativeValuationMethod (1319)
func (f DerivativeValuationMethod) Tag() fix.Tag { return tag.DerivativeValuationMethod }

//BuildDerivativeValuationMethod returns a new DerivativeValuationMethod initialized with val
func BuildDerivativeValuationMethod(val string) DerivativeValuationMethod {
	var field DerivativeValuationMethod
	field.Value = val
	return field
}

//Designation is a STRING field
type Designation struct{ message.StringValue }

//Tag returns tag.Designation (494)
func (f Designation) Tag() fix.Tag { return tag.Designation }

//BuildDesignation returns a new Designation initialized with val
func BuildDesignation(val string) Designation {
	var field Designation
	field.Value = val
	return field
}

//DeskID is a STRING field
type DeskID struct{ message.StringValue }

//Tag returns tag.DeskID (284)
func (f DeskID) Tag() fix.Tag { return tag.DeskID }

//BuildDeskID returns a new DeskID initialized with val
func BuildDeskID(val string) DeskID {
	var field DeskID
	field.Value = val
	return field
}

//DeskOrderHandlingInst is a MULTIPLESTRINGVALUE field
type DeskOrderHandlingInst struct{ message.MultipleStringValue }

//Tag returns tag.DeskOrderHandlingInst (1035)
func (f DeskOrderHandlingInst) Tag() fix.Tag { return tag.DeskOrderHandlingInst }

//BuildDeskOrderHandlingInst returns a new DeskOrderHandlingInst initialized with val
func BuildDeskOrderHandlingInst(val string) DeskOrderHandlingInst {
	var field DeskOrderHandlingInst
	field.Value = val
	return field
}

//DeskType is a STRING field
type DeskType struct{ message.StringValue }

//Tag returns tag.DeskType (1033)
func (f DeskType) Tag() fix.Tag { return tag.DeskType }

//BuildDeskType returns a new DeskType initialized with val
func BuildDeskType(val string) DeskType {
	var field DeskType
	field.Value = val
	return field
}

//DeskTypeSource is a INT field
type DeskTypeSource struct{ message.IntValue }

//Tag returns tag.DeskTypeSource (1034)
func (f DeskTypeSource) Tag() fix.Tag { return tag.DeskTypeSource }

//BuildDeskTypeSource returns a new DeskTypeSource initialized with val
func BuildDeskTypeSource(val int) DeskTypeSource {
	var field DeskTypeSource
	field.Value = val
	return field
}

//DetachmentPoint is a PERCENTAGE field
type DetachmentPoint struct{ message.PercentageValue }

//Tag returns tag.DetachmentPoint (1458)
func (f DetachmentPoint) Tag() fix.Tag { return tag.DetachmentPoint }

//BuildDetachmentPoint returns a new DetachmentPoint initialized with val
func BuildDetachmentPoint(val float64) DetachmentPoint {
	var field DetachmentPoint
	field.Value = val
	return field
}

//DiscretionInst is a CHAR field
type DiscretionInst struct{ message.CharValue }

//Tag returns tag.DiscretionInst (388)
func (f DiscretionInst) Tag() fix.Tag { return tag.DiscretionInst }

//BuildDiscretionInst returns a new DiscretionInst initialized with val
func BuildDiscretionInst(val string) DiscretionInst {
	var field DiscretionInst
	field.Value = val
	return field
}

//DiscretionLimitType is a INT field
type DiscretionLimitType struct{ message.IntValue }

//Tag returns tag.DiscretionLimitType (843)
func (f DiscretionLimitType) Tag() fix.Tag { return tag.DiscretionLimitType }

//BuildDiscretionLimitType returns a new DiscretionLimitType initialized with val
func BuildDiscretionLimitType(val int) DiscretionLimitType {
	var field DiscretionLimitType
	field.Value = val
	return field
}

//DiscretionMoveType is a INT field
type DiscretionMoveType struct{ message.IntValue }

//Tag returns tag.DiscretionMoveType (841)
func (f DiscretionMoveType) Tag() fix.Tag { return tag.DiscretionMoveType }

//BuildDiscretionMoveType returns a new DiscretionMoveType initialized with val
func BuildDiscretionMoveType(val int) DiscretionMoveType {
	var field DiscretionMoveType
	field.Value = val
	return field
}

//DiscretionOffset is a PRICEOFFSET field
type DiscretionOffset struct{ message.PriceOffsetValue }

//Tag returns tag.DiscretionOffset (389)
func (f DiscretionOffset) Tag() fix.Tag { return tag.DiscretionOffset }

//BuildDiscretionOffset returns a new DiscretionOffset initialized with val
func BuildDiscretionOffset(val float64) DiscretionOffset {
	var field DiscretionOffset
	field.Value = val
	return field
}

//DiscretionOffsetType is a INT field
type DiscretionOffsetType struct{ message.IntValue }

//Tag returns tag.DiscretionOffsetType (842)
func (f DiscretionOffsetType) Tag() fix.Tag { return tag.DiscretionOffsetType }

//BuildDiscretionOffsetType returns a new DiscretionOffsetType initialized with val
func BuildDiscretionOffsetType(val int) DiscretionOffsetType {
	var field DiscretionOffsetType
	field.Value = val
	return field
}

//DiscretionOffsetValue is a FLOAT field
type DiscretionOffsetValue struct{ message.FloatValue }

//Tag returns tag.DiscretionOffsetValue (389)
func (f DiscretionOffsetValue) Tag() fix.Tag { return tag.DiscretionOffsetValue }

//BuildDiscretionOffsetValue returns a new DiscretionOffsetValue initialized with val
func BuildDiscretionOffsetValue(val float64) DiscretionOffsetValue {
	var field DiscretionOffsetValue
	field.Value = val
	return field
}

//DiscretionPrice is a PRICE field
type DiscretionPrice struct{ message.PriceValue }

//Tag returns tag.DiscretionPrice (845)
func (f DiscretionPrice) Tag() fix.Tag { return tag.DiscretionPrice }

//BuildDiscretionPrice returns a new DiscretionPrice initialized with val
func BuildDiscretionPrice(val float64) DiscretionPrice {
	var field DiscretionPrice
	field.Value = val
	return field
}

//DiscretionRoundDirection is a INT field
type DiscretionRoundDirection struct{ message.IntValue }

//Tag returns tag.DiscretionRoundDirection (844)
func (f DiscretionRoundDirection) Tag() fix.Tag { return tag.DiscretionRoundDirection }

//BuildDiscretionRoundDirection returns a new DiscretionRoundDirection initialized with val
func BuildDiscretionRoundDirection(val int) DiscretionRoundDirection {
	var field DiscretionRoundDirection
	field.Value = val
	return field
}

//DiscretionScope is a INT field
type DiscretionScope struct{ message.IntValue }

//Tag returns tag.DiscretionScope (846)
func (f DiscretionScope) Tag() fix.Tag { return tag.DiscretionScope }

//BuildDiscretionScope returns a new DiscretionScope initialized with val
func BuildDiscretionScope(val int) DiscretionScope {
	var field DiscretionScope
	field.Value = val
	return field
}

//DisplayHighQty is a QTY field
type DisplayHighQty struct{ message.QtyValue }

//Tag returns tag.DisplayHighQty (1086)
func (f DisplayHighQty) Tag() fix.Tag { return tag.DisplayHighQty }

//BuildDisplayHighQty returns a new DisplayHighQty initialized with val
func BuildDisplayHighQty(val float64) DisplayHighQty {
	var field DisplayHighQty
	field.Value = val
	return field
}

//DisplayLowQty is a QTY field
type DisplayLowQty struct{ message.QtyValue }

//Tag returns tag.DisplayLowQty (1085)
func (f DisplayLowQty) Tag() fix.Tag { return tag.DisplayLowQty }

//BuildDisplayLowQty returns a new DisplayLowQty initialized with val
func BuildDisplayLowQty(val float64) DisplayLowQty {
	var field DisplayLowQty
	field.Value = val
	return field
}

//DisplayMethod is a CHAR field
type DisplayMethod struct{ message.CharValue }

//Tag returns tag.DisplayMethod (1084)
func (f DisplayMethod) Tag() fix.Tag { return tag.DisplayMethod }

//BuildDisplayMethod returns a new DisplayMethod initialized with val
func BuildDisplayMethod(val string) DisplayMethod {
	var field DisplayMethod
	field.Value = val
	return field
}

//DisplayMinIncr is a QTY field
type DisplayMinIncr struct{ message.QtyValue }

//Tag returns tag.DisplayMinIncr (1087)
func (f DisplayMinIncr) Tag() fix.Tag { return tag.DisplayMinIncr }

//BuildDisplayMinIncr returns a new DisplayMinIncr initialized with val
func BuildDisplayMinIncr(val float64) DisplayMinIncr {
	var field DisplayMinIncr
	field.Value = val
	return field
}

//DisplayQty is a QTY field
type DisplayQty struct{ message.QtyValue }

//Tag returns tag.DisplayQty (1138)
func (f DisplayQty) Tag() fix.Tag { return tag.DisplayQty }

//BuildDisplayQty returns a new DisplayQty initialized with val
func BuildDisplayQty(val float64) DisplayQty {
	var field DisplayQty
	field.Value = val
	return field
}

//DisplayWhen is a CHAR field
type DisplayWhen struct{ message.CharValue }

//Tag returns tag.DisplayWhen (1083)
func (f DisplayWhen) Tag() fix.Tag { return tag.DisplayWhen }

//BuildDisplayWhen returns a new DisplayWhen initialized with val
func BuildDisplayWhen(val string) DisplayWhen {
	var field DisplayWhen
	field.Value = val
	return field
}

//DistribPaymentMethod is a INT field
type DistribPaymentMethod struct{ message.IntValue }

//Tag returns tag.DistribPaymentMethod (477)
func (f DistribPaymentMethod) Tag() fix.Tag { return tag.DistribPaymentMethod }

//BuildDistribPaymentMethod returns a new DistribPaymentMethod initialized with val
func BuildDistribPaymentMethod(val int) DistribPaymentMethod {
	var field DistribPaymentMethod
	field.Value = val
	return field
}

//DistribPercentage is a PERCENTAGE field
type DistribPercentage struct{ message.PercentageValue }

//Tag returns tag.DistribPercentage (512)
func (f DistribPercentage) Tag() fix.Tag { return tag.DistribPercentage }

//BuildDistribPercentage returns a new DistribPercentage initialized with val
func BuildDistribPercentage(val float64) DistribPercentage {
	var field DistribPercentage
	field.Value = val
	return field
}

//DividendYield is a PERCENTAGE field
type DividendYield struct{ message.PercentageValue }

//Tag returns tag.DividendYield (1380)
func (f DividendYield) Tag() fix.Tag { return tag.DividendYield }

//BuildDividendYield returns a new DividendYield initialized with val
func BuildDividendYield(val float64) DividendYield {
	var field DividendYield
	field.Value = val
	return field
}

//DlvyInst is a STRING field
type DlvyInst struct{ message.StringValue }

//Tag returns tag.DlvyInst (86)
func (f DlvyInst) Tag() fix.Tag { return tag.DlvyInst }

//BuildDlvyInst returns a new DlvyInst initialized with val
func BuildDlvyInst(val string) DlvyInst {
	var field DlvyInst
	field.Value = val
	return field
}

//DlvyInstType is a CHAR field
type DlvyInstType struct{ message.CharValue }

//Tag returns tag.DlvyInstType (787)
func (f DlvyInstType) Tag() fix.Tag { return tag.DlvyInstType }

//BuildDlvyInstType returns a new DlvyInstType initialized with val
func BuildDlvyInstType(val string) DlvyInstType {
	var field DlvyInstType
	field.Value = val
	return field
}

//DueToRelated is a BOOLEAN field
type DueToRelated struct{ message.BooleanValue }

//Tag returns tag.DueToRelated (329)
func (f DueToRelated) Tag() fix.Tag { return tag.DueToRelated }

//BuildDueToRelated returns a new DueToRelated initialized with val
func BuildDueToRelated(val bool) DueToRelated {
	var field DueToRelated
	field.Value = val
	return field
}

//EFPTrackingError is a PERCENTAGE field
type EFPTrackingError struct{ message.PercentageValue }

//Tag returns tag.EFPTrackingError (405)
func (f EFPTrackingError) Tag() fix.Tag { return tag.EFPTrackingError }

//BuildEFPTrackingError returns a new EFPTrackingError initialized with val
func BuildEFPTrackingError(val float64) EFPTrackingError {
	var field EFPTrackingError
	field.Value = val
	return field
}

//EffectiveTime is a UTCTIMESTAMP field
type EffectiveTime struct{ message.UTCTimestampValue }

//Tag returns tag.EffectiveTime (168)
func (f EffectiveTime) Tag() fix.Tag { return tag.EffectiveTime }

//EmailThreadID is a STRING field
type EmailThreadID struct{ message.StringValue }

//Tag returns tag.EmailThreadID (164)
func (f EmailThreadID) Tag() fix.Tag { return tag.EmailThreadID }

//BuildEmailThreadID returns a new EmailThreadID initialized with val
func BuildEmailThreadID(val string) EmailThreadID {
	var field EmailThreadID
	field.Value = val
	return field
}

//EmailType is a CHAR field
type EmailType struct{ message.CharValue }

//Tag returns tag.EmailType (94)
func (f EmailType) Tag() fix.Tag { return tag.EmailType }

//BuildEmailType returns a new EmailType initialized with val
func BuildEmailType(val string) EmailType {
	var field EmailType
	field.Value = val
	return field
}

//EncodedAllocText is a DATA field
type EncodedAllocText struct{ message.DataValue }

//Tag returns tag.EncodedAllocText (361)
func (f EncodedAllocText) Tag() fix.Tag { return tag.EncodedAllocText }

//BuildEncodedAllocText returns a new EncodedAllocText initialized with val
func BuildEncodedAllocText(val string) EncodedAllocText {
	var field EncodedAllocText
	field.Value = val
	return field
}

//EncodedAllocTextLen is a LENGTH field
type EncodedAllocTextLen struct{ message.LengthValue }

//Tag returns tag.EncodedAllocTextLen (360)
func (f EncodedAllocTextLen) Tag() fix.Tag { return tag.EncodedAllocTextLen }

//BuildEncodedAllocTextLen returns a new EncodedAllocTextLen initialized with val
func BuildEncodedAllocTextLen(val int) EncodedAllocTextLen {
	var field EncodedAllocTextLen
	field.Value = val
	return field
}

//EncodedHeadline is a DATA field
type EncodedHeadline struct{ message.DataValue }

//Tag returns tag.EncodedHeadline (359)
func (f EncodedHeadline) Tag() fix.Tag { return tag.EncodedHeadline }

//BuildEncodedHeadline returns a new EncodedHeadline initialized with val
func BuildEncodedHeadline(val string) EncodedHeadline {
	var field EncodedHeadline
	field.Value = val
	return field
}

//EncodedHeadlineLen is a LENGTH field
type EncodedHeadlineLen struct{ message.LengthValue }

//Tag returns tag.EncodedHeadlineLen (358)
func (f EncodedHeadlineLen) Tag() fix.Tag { return tag.EncodedHeadlineLen }

//BuildEncodedHeadlineLen returns a new EncodedHeadlineLen initialized with val
func BuildEncodedHeadlineLen(val int) EncodedHeadlineLen {
	var field EncodedHeadlineLen
	field.Value = val
	return field
}

//EncodedIssuer is a DATA field
type EncodedIssuer struct{ message.DataValue }

//Tag returns tag.EncodedIssuer (349)
func (f EncodedIssuer) Tag() fix.Tag { return tag.EncodedIssuer }

//BuildEncodedIssuer returns a new EncodedIssuer initialized with val
func BuildEncodedIssuer(val string) EncodedIssuer {
	var field EncodedIssuer
	field.Value = val
	return field
}

//EncodedIssuerLen is a LENGTH field
type EncodedIssuerLen struct{ message.LengthValue }

//Tag returns tag.EncodedIssuerLen (348)
func (f EncodedIssuerLen) Tag() fix.Tag { return tag.EncodedIssuerLen }

//BuildEncodedIssuerLen returns a new EncodedIssuerLen initialized with val
func BuildEncodedIssuerLen(val int) EncodedIssuerLen {
	var field EncodedIssuerLen
	field.Value = val
	return field
}

//EncodedLegIssuer is a DATA field
type EncodedLegIssuer struct{ message.DataValue }

//Tag returns tag.EncodedLegIssuer (619)
func (f EncodedLegIssuer) Tag() fix.Tag { return tag.EncodedLegIssuer }

//BuildEncodedLegIssuer returns a new EncodedLegIssuer initialized with val
func BuildEncodedLegIssuer(val string) EncodedLegIssuer {
	var field EncodedLegIssuer
	field.Value = val
	return field
}

//EncodedLegIssuerLen is a LENGTH field
type EncodedLegIssuerLen struct{ message.LengthValue }

//Tag returns tag.EncodedLegIssuerLen (618)
func (f EncodedLegIssuerLen) Tag() fix.Tag { return tag.EncodedLegIssuerLen }

//BuildEncodedLegIssuerLen returns a new EncodedLegIssuerLen initialized with val
func BuildEncodedLegIssuerLen(val int) EncodedLegIssuerLen {
	var field EncodedLegIssuerLen
	field.Value = val
	return field
}

//EncodedLegSecurityDesc is a DATA field
type EncodedLegSecurityDesc struct{ message.DataValue }

//Tag returns tag.EncodedLegSecurityDesc (622)
func (f EncodedLegSecurityDesc) Tag() fix.Tag { return tag.EncodedLegSecurityDesc }

//BuildEncodedLegSecurityDesc returns a new EncodedLegSecurityDesc initialized with val
func BuildEncodedLegSecurityDesc(val string) EncodedLegSecurityDesc {
	var field EncodedLegSecurityDesc
	field.Value = val
	return field
}

//EncodedLegSecurityDescLen is a LENGTH field
type EncodedLegSecurityDescLen struct{ message.LengthValue }

//Tag returns tag.EncodedLegSecurityDescLen (621)
func (f EncodedLegSecurityDescLen) Tag() fix.Tag { return tag.EncodedLegSecurityDescLen }

//BuildEncodedLegSecurityDescLen returns a new EncodedLegSecurityDescLen initialized with val
func BuildEncodedLegSecurityDescLen(val int) EncodedLegSecurityDescLen {
	var field EncodedLegSecurityDescLen
	field.Value = val
	return field
}

//EncodedListExecInst is a DATA field
type EncodedListExecInst struct{ message.DataValue }

//Tag returns tag.EncodedListExecInst (353)
func (f EncodedListExecInst) Tag() fix.Tag { return tag.EncodedListExecInst }

//BuildEncodedListExecInst returns a new EncodedListExecInst initialized with val
func BuildEncodedListExecInst(val string) EncodedListExecInst {
	var field EncodedListExecInst
	field.Value = val
	return field
}

//EncodedListExecInstLen is a LENGTH field
type EncodedListExecInstLen struct{ message.LengthValue }

//Tag returns tag.EncodedListExecInstLen (352)
func (f EncodedListExecInstLen) Tag() fix.Tag { return tag.EncodedListExecInstLen }

//BuildEncodedListExecInstLen returns a new EncodedListExecInstLen initialized with val
func BuildEncodedListExecInstLen(val int) EncodedListExecInstLen {
	var field EncodedListExecInstLen
	field.Value = val
	return field
}

//EncodedListStatusText is a DATA field
type EncodedListStatusText struct{ message.DataValue }

//Tag returns tag.EncodedListStatusText (446)
func (f EncodedListStatusText) Tag() fix.Tag { return tag.EncodedListStatusText }

//BuildEncodedListStatusText returns a new EncodedListStatusText initialized with val
func BuildEncodedListStatusText(val string) EncodedListStatusText {
	var field EncodedListStatusText
	field.Value = val
	return field
}

//EncodedListStatusTextLen is a LENGTH field
type EncodedListStatusTextLen struct{ message.LengthValue }

//Tag returns tag.EncodedListStatusTextLen (445)
func (f EncodedListStatusTextLen) Tag() fix.Tag { return tag.EncodedListStatusTextLen }

//BuildEncodedListStatusTextLen returns a new EncodedListStatusTextLen initialized with val
func BuildEncodedListStatusTextLen(val int) EncodedListStatusTextLen {
	var field EncodedListStatusTextLen
	field.Value = val
	return field
}

//EncodedMktSegmDesc is a DATA field
type EncodedMktSegmDesc struct{ message.DataValue }

//Tag returns tag.EncodedMktSegmDesc (1398)
func (f EncodedMktSegmDesc) Tag() fix.Tag { return tag.EncodedMktSegmDesc }

//BuildEncodedMktSegmDesc returns a new EncodedMktSegmDesc initialized with val
func BuildEncodedMktSegmDesc(val string) EncodedMktSegmDesc {
	var field EncodedMktSegmDesc
	field.Value = val
	return field
}

//EncodedMktSegmDescLen is a LENGTH field
type EncodedMktSegmDescLen struct{ message.LengthValue }

//Tag returns tag.EncodedMktSegmDescLen (1397)
func (f EncodedMktSegmDescLen) Tag() fix.Tag { return tag.EncodedMktSegmDescLen }

//BuildEncodedMktSegmDescLen returns a new EncodedMktSegmDescLen initialized with val
func BuildEncodedMktSegmDescLen(val int) EncodedMktSegmDescLen {
	var field EncodedMktSegmDescLen
	field.Value = val
	return field
}

//EncodedSecurityDesc is a DATA field
type EncodedSecurityDesc struct{ message.DataValue }

//Tag returns tag.EncodedSecurityDesc (351)
func (f EncodedSecurityDesc) Tag() fix.Tag { return tag.EncodedSecurityDesc }

//BuildEncodedSecurityDesc returns a new EncodedSecurityDesc initialized with val
func BuildEncodedSecurityDesc(val string) EncodedSecurityDesc {
	var field EncodedSecurityDesc
	field.Value = val
	return field
}

//EncodedSecurityDescLen is a LENGTH field
type EncodedSecurityDescLen struct{ message.LengthValue }

//Tag returns tag.EncodedSecurityDescLen (350)
func (f EncodedSecurityDescLen) Tag() fix.Tag { return tag.EncodedSecurityDescLen }

//BuildEncodedSecurityDescLen returns a new EncodedSecurityDescLen initialized with val
func BuildEncodedSecurityDescLen(val int) EncodedSecurityDescLen {
	var field EncodedSecurityDescLen
	field.Value = val
	return field
}

//EncodedSecurityListDesc is a DATA field
type EncodedSecurityListDesc struct{ message.DataValue }

//Tag returns tag.EncodedSecurityListDesc (1469)
func (f EncodedSecurityListDesc) Tag() fix.Tag { return tag.EncodedSecurityListDesc }

//BuildEncodedSecurityListDesc returns a new EncodedSecurityListDesc initialized with val
func BuildEncodedSecurityListDesc(val string) EncodedSecurityListDesc {
	var field EncodedSecurityListDesc
	field.Value = val
	return field
}

//EncodedSecurityListDescLen is a LENGTH field
type EncodedSecurityListDescLen struct{ message.LengthValue }

//Tag returns tag.EncodedSecurityListDescLen (1468)
func (f EncodedSecurityListDescLen) Tag() fix.Tag { return tag.EncodedSecurityListDescLen }

//BuildEncodedSecurityListDescLen returns a new EncodedSecurityListDescLen initialized with val
func BuildEncodedSecurityListDescLen(val int) EncodedSecurityListDescLen {
	var field EncodedSecurityListDescLen
	field.Value = val
	return field
}

//EncodedSubject is a DATA field
type EncodedSubject struct{ message.DataValue }

//Tag returns tag.EncodedSubject (357)
func (f EncodedSubject) Tag() fix.Tag { return tag.EncodedSubject }

//BuildEncodedSubject returns a new EncodedSubject initialized with val
func BuildEncodedSubject(val string) EncodedSubject {
	var field EncodedSubject
	field.Value = val
	return field
}

//EncodedSubjectLen is a LENGTH field
type EncodedSubjectLen struct{ message.LengthValue }

//Tag returns tag.EncodedSubjectLen (356)
func (f EncodedSubjectLen) Tag() fix.Tag { return tag.EncodedSubjectLen }

//BuildEncodedSubjectLen returns a new EncodedSubjectLen initialized with val
func BuildEncodedSubjectLen(val int) EncodedSubjectLen {
	var field EncodedSubjectLen
	field.Value = val
	return field
}

//EncodedSymbol is a DATA field
type EncodedSymbol struct{ message.DataValue }

//Tag returns tag.EncodedSymbol (1360)
func (f EncodedSymbol) Tag() fix.Tag { return tag.EncodedSymbol }

//BuildEncodedSymbol returns a new EncodedSymbol initialized with val
func BuildEncodedSymbol(val string) EncodedSymbol {
	var field EncodedSymbol
	field.Value = val
	return field
}

//EncodedSymbolLen is a LENGTH field
type EncodedSymbolLen struct{ message.LengthValue }

//Tag returns tag.EncodedSymbolLen (1359)
func (f EncodedSymbolLen) Tag() fix.Tag { return tag.EncodedSymbolLen }

//BuildEncodedSymbolLen returns a new EncodedSymbolLen initialized with val
func BuildEncodedSymbolLen(val int) EncodedSymbolLen {
	var field EncodedSymbolLen
	field.Value = val
	return field
}

//EncodedText is a DATA field
type EncodedText struct{ message.DataValue }

//Tag returns tag.EncodedText (355)
func (f EncodedText) Tag() fix.Tag { return tag.EncodedText }

//BuildEncodedText returns a new EncodedText initialized with val
func BuildEncodedText(val string) EncodedText {
	var field EncodedText
	field.Value = val
	return field
}

//EncodedTextLen is a LENGTH field
type EncodedTextLen struct{ message.LengthValue }

//Tag returns tag.EncodedTextLen (354)
func (f EncodedTextLen) Tag() fix.Tag { return tag.EncodedTextLen }

//BuildEncodedTextLen returns a new EncodedTextLen initialized with val
func BuildEncodedTextLen(val int) EncodedTextLen {
	var field EncodedTextLen
	field.Value = val
	return field
}

//EncodedUnderlyingIssuer is a DATA field
type EncodedUnderlyingIssuer struct{ message.DataValue }

//Tag returns tag.EncodedUnderlyingIssuer (363)
func (f EncodedUnderlyingIssuer) Tag() fix.Tag { return tag.EncodedUnderlyingIssuer }

//BuildEncodedUnderlyingIssuer returns a new EncodedUnderlyingIssuer initialized with val
func BuildEncodedUnderlyingIssuer(val string) EncodedUnderlyingIssuer {
	var field EncodedUnderlyingIssuer
	field.Value = val
	return field
}

//EncodedUnderlyingIssuerLen is a LENGTH field
type EncodedUnderlyingIssuerLen struct{ message.LengthValue }

//Tag returns tag.EncodedUnderlyingIssuerLen (362)
func (f EncodedUnderlyingIssuerLen) Tag() fix.Tag { return tag.EncodedUnderlyingIssuerLen }

//BuildEncodedUnderlyingIssuerLen returns a new EncodedUnderlyingIssuerLen initialized with val
func BuildEncodedUnderlyingIssuerLen(val int) EncodedUnderlyingIssuerLen {
	var field EncodedUnderlyingIssuerLen
	field.Value = val
	return field
}

//EncodedUnderlyingSecurityDesc is a DATA field
type EncodedUnderlyingSecurityDesc struct{ message.DataValue }

//Tag returns tag.EncodedUnderlyingSecurityDesc (365)
func (f EncodedUnderlyingSecurityDesc) Tag() fix.Tag { return tag.EncodedUnderlyingSecurityDesc }

//BuildEncodedUnderlyingSecurityDesc returns a new EncodedUnderlyingSecurityDesc initialized with val
func BuildEncodedUnderlyingSecurityDesc(val string) EncodedUnderlyingSecurityDesc {
	var field EncodedUnderlyingSecurityDesc
	field.Value = val
	return field
}

//EncodedUnderlyingSecurityDescLen is a LENGTH field
type EncodedUnderlyingSecurityDescLen struct{ message.LengthValue }

//Tag returns tag.EncodedUnderlyingSecurityDescLen (364)
func (f EncodedUnderlyingSecurityDescLen) Tag() fix.Tag { return tag.EncodedUnderlyingSecurityDescLen }

//BuildEncodedUnderlyingSecurityDescLen returns a new EncodedUnderlyingSecurityDescLen initialized with val
func BuildEncodedUnderlyingSecurityDescLen(val int) EncodedUnderlyingSecurityDescLen {
	var field EncodedUnderlyingSecurityDescLen
	field.Value = val
	return field
}

//EncryptMethod is a INT field
type EncryptMethod struct{ message.IntValue }

//Tag returns tag.EncryptMethod (98)
func (f EncryptMethod) Tag() fix.Tag { return tag.EncryptMethod }

//BuildEncryptMethod returns a new EncryptMethod initialized with val
func BuildEncryptMethod(val int) EncryptMethod {
	var field EncryptMethod
	field.Value = val
	return field
}

//EncryptedNewPassword is a DATA field
type EncryptedNewPassword struct{ message.DataValue }

//Tag returns tag.EncryptedNewPassword (1404)
func (f EncryptedNewPassword) Tag() fix.Tag { return tag.EncryptedNewPassword }

//BuildEncryptedNewPassword returns a new EncryptedNewPassword initialized with val
func BuildEncryptedNewPassword(val string) EncryptedNewPassword {
	var field EncryptedNewPassword
	field.Value = val
	return field
}

//EncryptedNewPasswordLen is a LENGTH field
type EncryptedNewPasswordLen struct{ message.LengthValue }

//Tag returns tag.EncryptedNewPasswordLen (1403)
func (f EncryptedNewPasswordLen) Tag() fix.Tag { return tag.EncryptedNewPasswordLen }

//BuildEncryptedNewPasswordLen returns a new EncryptedNewPasswordLen initialized with val
func BuildEncryptedNewPasswordLen(val int) EncryptedNewPasswordLen {
	var field EncryptedNewPasswordLen
	field.Value = val
	return field
}

//EncryptedPassword is a DATA field
type EncryptedPassword struct{ message.DataValue }

//Tag returns tag.EncryptedPassword (1402)
func (f EncryptedPassword) Tag() fix.Tag { return tag.EncryptedPassword }

//BuildEncryptedPassword returns a new EncryptedPassword initialized with val
func BuildEncryptedPassword(val string) EncryptedPassword {
	var field EncryptedPassword
	field.Value = val
	return field
}

//EncryptedPasswordLen is a LENGTH field
type EncryptedPasswordLen struct{ message.LengthValue }

//Tag returns tag.EncryptedPasswordLen (1401)
func (f EncryptedPasswordLen) Tag() fix.Tag { return tag.EncryptedPasswordLen }

//BuildEncryptedPasswordLen returns a new EncryptedPasswordLen initialized with val
func BuildEncryptedPasswordLen(val int) EncryptedPasswordLen {
	var field EncryptedPasswordLen
	field.Value = val
	return field
}

//EncryptedPasswordMethod is a INT field
type EncryptedPasswordMethod struct{ message.IntValue }

//Tag returns tag.EncryptedPasswordMethod (1400)
func (f EncryptedPasswordMethod) Tag() fix.Tag { return tag.EncryptedPasswordMethod }

//BuildEncryptedPasswordMethod returns a new EncryptedPasswordMethod initialized with val
func BuildEncryptedPasswordMethod(val int) EncryptedPasswordMethod {
	var field EncryptedPasswordMethod
	field.Value = val
	return field
}

//EndAccruedInterestAmt is a AMT field
type EndAccruedInterestAmt struct{ message.AmtValue }

//Tag returns tag.EndAccruedInterestAmt (920)
func (f EndAccruedInterestAmt) Tag() fix.Tag { return tag.EndAccruedInterestAmt }

//BuildEndAccruedInterestAmt returns a new EndAccruedInterestAmt initialized with val
func BuildEndAccruedInterestAmt(val float64) EndAccruedInterestAmt {
	var field EndAccruedInterestAmt
	field.Value = val
	return field
}

//EndCash is a AMT field
type EndCash struct{ message.AmtValue }

//Tag returns tag.EndCash (922)
func (f EndCash) Tag() fix.Tag { return tag.EndCash }

//BuildEndCash returns a new EndCash initialized with val
func BuildEndCash(val float64) EndCash {
	var field EndCash
	field.Value = val
	return field
}

//EndDate is a LOCALMKTDATE field
type EndDate struct{ message.LocalMktDateValue }

//Tag returns tag.EndDate (917)
func (f EndDate) Tag() fix.Tag { return tag.EndDate }

//BuildEndDate returns a new EndDate initialized with val
func BuildEndDate(val string) EndDate {
	var field EndDate
	field.Value = val
	return field
}

//EndMaturityMonthYear is a MONTHYEAR field
type EndMaturityMonthYear struct{ message.MonthYearValue }

//Tag returns tag.EndMaturityMonthYear (1226)
func (f EndMaturityMonthYear) Tag() fix.Tag { return tag.EndMaturityMonthYear }

//BuildEndMaturityMonthYear returns a new EndMaturityMonthYear initialized with val
func BuildEndMaturityMonthYear(val string) EndMaturityMonthYear {
	var field EndMaturityMonthYear
	field.Value = val
	return field
}

//EndSeqNo is a SEQNUM field
type EndSeqNo struct{ message.SeqNumValue }

//Tag returns tag.EndSeqNo (16)
func (f EndSeqNo) Tag() fix.Tag { return tag.EndSeqNo }

//BuildEndSeqNo returns a new EndSeqNo initialized with val
func BuildEndSeqNo(val int) EndSeqNo {
	var field EndSeqNo
	field.Value = val
	return field
}

//EndStrikePxRange is a PRICE field
type EndStrikePxRange struct{ message.PriceValue }

//Tag returns tag.EndStrikePxRange (1203)
func (f EndStrikePxRange) Tag() fix.Tag { return tag.EndStrikePxRange }

//BuildEndStrikePxRange returns a new EndStrikePxRange initialized with val
func BuildEndStrikePxRange(val float64) EndStrikePxRange {
	var field EndStrikePxRange
	field.Value = val
	return field
}

//EndTickPriceRange is a PRICE field
type EndTickPriceRange struct{ message.PriceValue }

//Tag returns tag.EndTickPriceRange (1207)
func (f EndTickPriceRange) Tag() fix.Tag { return tag.EndTickPriceRange }

//BuildEndTickPriceRange returns a new EndTickPriceRange initialized with val
func BuildEndTickPriceRange(val float64) EndTickPriceRange {
	var field EndTickPriceRange
	field.Value = val
	return field
}

//EventDate is a LOCALMKTDATE field
type EventDate struct{ message.LocalMktDateValue }

//Tag returns tag.EventDate (866)
func (f EventDate) Tag() fix.Tag { return tag.EventDate }

//BuildEventDate returns a new EventDate initialized with val
func BuildEventDate(val string) EventDate {
	var field EventDate
	field.Value = val
	return field
}

//EventPx is a PRICE field
type EventPx struct{ message.PriceValue }

//Tag returns tag.EventPx (867)
func (f EventPx) Tag() fix.Tag { return tag.EventPx }

//BuildEventPx returns a new EventPx initialized with val
func BuildEventPx(val float64) EventPx {
	var field EventPx
	field.Value = val
	return field
}

//EventText is a STRING field
type EventText struct{ message.StringValue }

//Tag returns tag.EventText (868)
func (f EventText) Tag() fix.Tag { return tag.EventText }

//BuildEventText returns a new EventText initialized with val
func BuildEventText(val string) EventText {
	var field EventText
	field.Value = val
	return field
}

//EventTime is a UTCTIMESTAMP field
type EventTime struct{ message.UTCTimestampValue }

//Tag returns tag.EventTime (1145)
func (f EventTime) Tag() fix.Tag { return tag.EventTime }

//EventType is a INT field
type EventType struct{ message.IntValue }

//Tag returns tag.EventType (865)
func (f EventType) Tag() fix.Tag { return tag.EventType }

//BuildEventType returns a new EventType initialized with val
func BuildEventType(val int) EventType {
	var field EventType
	field.Value = val
	return field
}

//ExDate is a LOCALMKTDATE field
type ExDate struct{ message.LocalMktDateValue }

//Tag returns tag.ExDate (230)
func (f ExDate) Tag() fix.Tag { return tag.ExDate }

//BuildExDate returns a new ExDate initialized with val
func BuildExDate(val string) ExDate {
	var field ExDate
	field.Value = val
	return field
}

//ExDestination is a EXCHANGE field
type ExDestination struct{ message.ExchangeValue }

//Tag returns tag.ExDestination (100)
func (f ExDestination) Tag() fix.Tag { return tag.ExDestination }

//BuildExDestination returns a new ExDestination initialized with val
func BuildExDestination(val string) ExDestination {
	var field ExDestination
	field.Value = val
	return field
}

//ExDestinationIDSource is a CHAR field
type ExDestinationIDSource struct{ message.CharValue }

//Tag returns tag.ExDestinationIDSource (1133)
func (f ExDestinationIDSource) Tag() fix.Tag { return tag.ExDestinationIDSource }

//BuildExDestinationIDSource returns a new ExDestinationIDSource initialized with val
func BuildExDestinationIDSource(val string) ExDestinationIDSource {
	var field ExDestinationIDSource
	field.Value = val
	return field
}

//ExchangeForPhysical is a BOOLEAN field
type ExchangeForPhysical struct{ message.BooleanValue }

//Tag returns tag.ExchangeForPhysical (411)
func (f ExchangeForPhysical) Tag() fix.Tag { return tag.ExchangeForPhysical }

//BuildExchangeForPhysical returns a new ExchangeForPhysical initialized with val
func BuildExchangeForPhysical(val bool) ExchangeForPhysical {
	var field ExchangeForPhysical
	field.Value = val
	return field
}

//ExchangeRule is a STRING field
type ExchangeRule struct{ message.StringValue }

//Tag returns tag.ExchangeRule (825)
func (f ExchangeRule) Tag() fix.Tag { return tag.ExchangeRule }

//BuildExchangeRule returns a new ExchangeRule initialized with val
func BuildExchangeRule(val string) ExchangeRule {
	var field ExchangeRule
	field.Value = val
	return field
}

//ExchangeSpecialInstructions is a STRING field
type ExchangeSpecialInstructions struct{ message.StringValue }

//Tag returns tag.ExchangeSpecialInstructions (1139)
func (f ExchangeSpecialInstructions) Tag() fix.Tag { return tag.ExchangeSpecialInstructions }

//BuildExchangeSpecialInstructions returns a new ExchangeSpecialInstructions initialized with val
func BuildExchangeSpecialInstructions(val string) ExchangeSpecialInstructions {
	var field ExchangeSpecialInstructions
	field.Value = val
	return field
}

//ExecAckStatus is a CHAR field
type ExecAckStatus struct{ message.CharValue }

//Tag returns tag.ExecAckStatus (1036)
func (f ExecAckStatus) Tag() fix.Tag { return tag.ExecAckStatus }

//BuildExecAckStatus returns a new ExecAckStatus initialized with val
func BuildExecAckStatus(val string) ExecAckStatus {
	var field ExecAckStatus
	field.Value = val
	return field
}

//ExecBroker is a STRING field
type ExecBroker struct{ message.StringValue }

//Tag returns tag.ExecBroker (76)
func (f ExecBroker) Tag() fix.Tag { return tag.ExecBroker }

//BuildExecBroker returns a new ExecBroker initialized with val
func BuildExecBroker(val string) ExecBroker {
	var field ExecBroker
	field.Value = val
	return field
}

//ExecID is a STRING field
type ExecID struct{ message.StringValue }

//Tag returns tag.ExecID (17)
func (f ExecID) Tag() fix.Tag { return tag.ExecID }

//BuildExecID returns a new ExecID initialized with val
func BuildExecID(val string) ExecID {
	var field ExecID
	field.Value = val
	return field
}

//ExecInst is a MULTIPLECHARVALUE field
type ExecInst struct{ message.MultipleCharValue }

//Tag returns tag.ExecInst (18)
func (f ExecInst) Tag() fix.Tag { return tag.ExecInst }

//BuildExecInst returns a new ExecInst initialized with val
func BuildExecInst(val string) ExecInst {
	var field ExecInst
	field.Value = val
	return field
}

//ExecInstValue is a CHAR field
type ExecInstValue struct{ message.CharValue }

//Tag returns tag.ExecInstValue (1308)
func (f ExecInstValue) Tag() fix.Tag { return tag.ExecInstValue }

//BuildExecInstValue returns a new ExecInstValue initialized with val
func BuildExecInstValue(val string) ExecInstValue {
	var field ExecInstValue
	field.Value = val
	return field
}

//ExecPriceAdjustment is a FLOAT field
type ExecPriceAdjustment struct{ message.FloatValue }

//Tag returns tag.ExecPriceAdjustment (485)
func (f ExecPriceAdjustment) Tag() fix.Tag { return tag.ExecPriceAdjustment }

//BuildExecPriceAdjustment returns a new ExecPriceAdjustment initialized with val
func BuildExecPriceAdjustment(val float64) ExecPriceAdjustment {
	var field ExecPriceAdjustment
	field.Value = val
	return field
}

//ExecPriceType is a CHAR field
type ExecPriceType struct{ message.CharValue }

//Tag returns tag.ExecPriceType (484)
func (f ExecPriceType) Tag() fix.Tag { return tag.ExecPriceType }

//BuildExecPriceType returns a new ExecPriceType initialized with val
func BuildExecPriceType(val string) ExecPriceType {
	var field ExecPriceType
	field.Value = val
	return field
}

//ExecRefID is a STRING field
type ExecRefID struct{ message.StringValue }

//Tag returns tag.ExecRefID (19)
func (f ExecRefID) Tag() fix.Tag { return tag.ExecRefID }

//BuildExecRefID returns a new ExecRefID initialized with val
func BuildExecRefID(val string) ExecRefID {
	var field ExecRefID
	field.Value = val
	return field
}

//ExecRestatementReason is a INT field
type ExecRestatementReason struct{ message.IntValue }

//Tag returns tag.ExecRestatementReason (378)
func (f ExecRestatementReason) Tag() fix.Tag { return tag.ExecRestatementReason }

//BuildExecRestatementReason returns a new ExecRestatementReason initialized with val
func BuildExecRestatementReason(val int) ExecRestatementReason {
	var field ExecRestatementReason
	field.Value = val
	return field
}

//ExecTransType is a CHAR field
type ExecTransType struct{ message.CharValue }

//Tag returns tag.ExecTransType (20)
func (f ExecTransType) Tag() fix.Tag { return tag.ExecTransType }

//BuildExecTransType returns a new ExecTransType initialized with val
func BuildExecTransType(val string) ExecTransType {
	var field ExecTransType
	field.Value = val
	return field
}

//ExecType is a CHAR field
type ExecType struct{ message.CharValue }

//Tag returns tag.ExecType (150)
func (f ExecType) Tag() fix.Tag { return tag.ExecType }

//BuildExecType returns a new ExecType initialized with val
func BuildExecType(val string) ExecType {
	var field ExecType
	field.Value = val
	return field
}

//ExecValuationPoint is a UTCTIMESTAMP field
type ExecValuationPoint struct{ message.UTCTimestampValue }

//Tag returns tag.ExecValuationPoint (515)
func (f ExecValuationPoint) Tag() fix.Tag { return tag.ExecValuationPoint }

//ExerciseMethod is a CHAR field
type ExerciseMethod struct{ message.CharValue }

//Tag returns tag.ExerciseMethod (747)
func (f ExerciseMethod) Tag() fix.Tag { return tag.ExerciseMethod }

//BuildExerciseMethod returns a new ExerciseMethod initialized with val
func BuildExerciseMethod(val string) ExerciseMethod {
	var field ExerciseMethod
	field.Value = val
	return field
}

//ExerciseStyle is a INT field
type ExerciseStyle struct{ message.IntValue }

//Tag returns tag.ExerciseStyle (1194)
func (f ExerciseStyle) Tag() fix.Tag { return tag.ExerciseStyle }

//BuildExerciseStyle returns a new ExerciseStyle initialized with val
func BuildExerciseStyle(val int) ExerciseStyle {
	var field ExerciseStyle
	field.Value = val
	return field
}

//ExpQty is a QTY field
type ExpQty struct{ message.QtyValue }

//Tag returns tag.ExpQty (983)
func (f ExpQty) Tag() fix.Tag { return tag.ExpQty }

//BuildExpQty returns a new ExpQty initialized with val
func BuildExpQty(val float64) ExpQty {
	var field ExpQty
	field.Value = val
	return field
}

//ExpType is a INT field
type ExpType struct{ message.IntValue }

//Tag returns tag.ExpType (982)
func (f ExpType) Tag() fix.Tag { return tag.ExpType }

//BuildExpType returns a new ExpType initialized with val
func BuildExpType(val int) ExpType {
	var field ExpType
	field.Value = val
	return field
}

//ExpirationCycle is a INT field
type ExpirationCycle struct{ message.IntValue }

//Tag returns tag.ExpirationCycle (827)
func (f ExpirationCycle) Tag() fix.Tag { return tag.ExpirationCycle }

//BuildExpirationCycle returns a new ExpirationCycle initialized with val
func BuildExpirationCycle(val int) ExpirationCycle {
	var field ExpirationCycle
	field.Value = val
	return field
}

//ExpirationQtyType is a INT field
type ExpirationQtyType struct{ message.IntValue }

//Tag returns tag.ExpirationQtyType (982)
func (f ExpirationQtyType) Tag() fix.Tag { return tag.ExpirationQtyType }

//BuildExpirationQtyType returns a new ExpirationQtyType initialized with val
func BuildExpirationQtyType(val int) ExpirationQtyType {
	var field ExpirationQtyType
	field.Value = val
	return field
}

//ExpireDate is a LOCALMKTDATE field
type ExpireDate struct{ message.LocalMktDateValue }

//Tag returns tag.ExpireDate (432)
func (f ExpireDate) Tag() fix.Tag { return tag.ExpireDate }

//BuildExpireDate returns a new ExpireDate initialized with val
func BuildExpireDate(val string) ExpireDate {
	var field ExpireDate
	field.Value = val
	return field
}

//ExpireTime is a UTCTIMESTAMP field
type ExpireTime struct{ message.UTCTimestampValue }

//Tag returns tag.ExpireTime (126)
func (f ExpireTime) Tag() fix.Tag { return tag.ExpireTime }

//Factor is a FLOAT field
type Factor struct{ message.FloatValue }

//Tag returns tag.Factor (228)
func (f Factor) Tag() fix.Tag { return tag.Factor }

//BuildFactor returns a new Factor initialized with val
func BuildFactor(val float64) Factor {
	var field Factor
	field.Value = val
	return field
}

//FairValue is a AMT field
type FairValue struct{ message.AmtValue }

//Tag returns tag.FairValue (406)
func (f FairValue) Tag() fix.Tag { return tag.FairValue }

//BuildFairValue returns a new FairValue initialized with val
func BuildFairValue(val float64) FairValue {
	var field FairValue
	field.Value = val
	return field
}

//FeeMultiplier is a FLOAT field
type FeeMultiplier struct{ message.FloatValue }

//Tag returns tag.FeeMultiplier (1329)
func (f FeeMultiplier) Tag() fix.Tag { return tag.FeeMultiplier }

//BuildFeeMultiplier returns a new FeeMultiplier initialized with val
func BuildFeeMultiplier(val float64) FeeMultiplier {
	var field FeeMultiplier
	field.Value = val
	return field
}

//FillExecID is a STRING field
type FillExecID struct{ message.StringValue }

//Tag returns tag.FillExecID (1363)
func (f FillExecID) Tag() fix.Tag { return tag.FillExecID }

//BuildFillExecID returns a new FillExecID initialized with val
func BuildFillExecID(val string) FillExecID {
	var field FillExecID
	field.Value = val
	return field
}

//FillLiquidityInd is a INT field
type FillLiquidityInd struct{ message.IntValue }

//Tag returns tag.FillLiquidityInd (1443)
func (f FillLiquidityInd) Tag() fix.Tag { return tag.FillLiquidityInd }

//BuildFillLiquidityInd returns a new FillLiquidityInd initialized with val
func BuildFillLiquidityInd(val int) FillLiquidityInd {
	var field FillLiquidityInd
	field.Value = val
	return field
}

//FillPx is a PRICE field
type FillPx struct{ message.PriceValue }

//Tag returns tag.FillPx (1364)
func (f FillPx) Tag() fix.Tag { return tag.FillPx }

//BuildFillPx returns a new FillPx initialized with val
func BuildFillPx(val float64) FillPx {
	var field FillPx
	field.Value = val
	return field
}

//FillQty is a QTY field
type FillQty struct{ message.QtyValue }

//Tag returns tag.FillQty (1365)
func (f FillQty) Tag() fix.Tag { return tag.FillQty }

//BuildFillQty returns a new FillQty initialized with val
func BuildFillQty(val float64) FillQty {
	var field FillQty
	field.Value = val
	return field
}

//FinancialStatus is a MULTIPLECHARVALUE field
type FinancialStatus struct{ message.MultipleCharValue }

//Tag returns tag.FinancialStatus (291)
func (f FinancialStatus) Tag() fix.Tag { return tag.FinancialStatus }

//BuildFinancialStatus returns a new FinancialStatus initialized with val
func BuildFinancialStatus(val string) FinancialStatus {
	var field FinancialStatus
	field.Value = val
	return field
}

//FirmTradeID is a STRING field
type FirmTradeID struct{ message.StringValue }

//Tag returns tag.FirmTradeID (1041)
func (f FirmTradeID) Tag() fix.Tag { return tag.FirmTradeID }

//BuildFirmTradeID returns a new FirmTradeID initialized with val
func BuildFirmTradeID(val string) FirmTradeID {
	var field FirmTradeID
	field.Value = val
	return field
}

//FirstPx is a PRICE field
type FirstPx struct{ message.PriceValue }

//Tag returns tag.FirstPx (1025)
func (f FirstPx) Tag() fix.Tag { return tag.FirstPx }

//BuildFirstPx returns a new FirstPx initialized with val
func BuildFirstPx(val float64) FirstPx {
	var field FirstPx
	field.Value = val
	return field
}

//FlexProductEligibilityIndicator is a BOOLEAN field
type FlexProductEligibilityIndicator struct{ message.BooleanValue }

//Tag returns tag.FlexProductEligibilityIndicator (1242)
func (f FlexProductEligibilityIndicator) Tag() fix.Tag { return tag.FlexProductEligibilityIndicator }

//BuildFlexProductEligibilityIndicator returns a new FlexProductEligibilityIndicator initialized with val
func BuildFlexProductEligibilityIndicator(val bool) FlexProductEligibilityIndicator {
	var field FlexProductEligibilityIndicator
	field.Value = val
	return field
}

//FlexibleIndicator is a BOOLEAN field
type FlexibleIndicator struct{ message.BooleanValue }

//Tag returns tag.FlexibleIndicator (1244)
func (f FlexibleIndicator) Tag() fix.Tag { return tag.FlexibleIndicator }

//BuildFlexibleIndicator returns a new FlexibleIndicator initialized with val
func BuildFlexibleIndicator(val bool) FlexibleIndicator {
	var field FlexibleIndicator
	field.Value = val
	return field
}

//FloorPrice is a PRICE field
type FloorPrice struct{ message.PriceValue }

//Tag returns tag.FloorPrice (1200)
func (f FloorPrice) Tag() fix.Tag { return tag.FloorPrice }

//BuildFloorPrice returns a new FloorPrice initialized with val
func BuildFloorPrice(val float64) FloorPrice {
	var field FloorPrice
	field.Value = val
	return field
}

//FlowScheduleType is a INT field
type FlowScheduleType struct{ message.IntValue }

//Tag returns tag.FlowScheduleType (1439)
func (f FlowScheduleType) Tag() fix.Tag { return tag.FlowScheduleType }

//BuildFlowScheduleType returns a new FlowScheduleType initialized with val
func BuildFlowScheduleType(val int) FlowScheduleType {
	var field FlowScheduleType
	field.Value = val
	return field
}

//ForexReq is a BOOLEAN field
type ForexReq struct{ message.BooleanValue }

//Tag returns tag.ForexReq (121)
func (f ForexReq) Tag() fix.Tag { return tag.ForexReq }

//BuildForexReq returns a new ForexReq initialized with val
func BuildForexReq(val bool) ForexReq {
	var field ForexReq
	field.Value = val
	return field
}

//FundRenewWaiv is a CHAR field
type FundRenewWaiv struct{ message.CharValue }

//Tag returns tag.FundRenewWaiv (497)
func (f FundRenewWaiv) Tag() fix.Tag { return tag.FundRenewWaiv }

//BuildFundRenewWaiv returns a new FundRenewWaiv initialized with val
func BuildFundRenewWaiv(val string) FundRenewWaiv {
	var field FundRenewWaiv
	field.Value = val
	return field
}

//FutSettDate is a LOCALMKTDATE field
type FutSettDate struct{ message.LocalMktDateValue }

//Tag returns tag.FutSettDate (64)
func (f FutSettDate) Tag() fix.Tag { return tag.FutSettDate }

//BuildFutSettDate returns a new FutSettDate initialized with val
func BuildFutSettDate(val string) FutSettDate {
	var field FutSettDate
	field.Value = val
	return field
}

//FutSettDate2 is a LOCALMKTDATE field
type FutSettDate2 struct{ message.LocalMktDateValue }

//Tag returns tag.FutSettDate2 (193)
func (f FutSettDate2) Tag() fix.Tag { return tag.FutSettDate2 }

//BuildFutSettDate2 returns a new FutSettDate2 initialized with val
func BuildFutSettDate2(val string) FutSettDate2 {
	var field FutSettDate2
	field.Value = val
	return field
}

//FuturesValuationMethod is a STRING field
type FuturesValuationMethod struct{ message.StringValue }

//Tag returns tag.FuturesValuationMethod (1197)
func (f FuturesValuationMethod) Tag() fix.Tag { return tag.FuturesValuationMethod }

//BuildFuturesValuationMethod returns a new FuturesValuationMethod initialized with val
func BuildFuturesValuationMethod(val string) FuturesValuationMethod {
	var field FuturesValuationMethod
	field.Value = val
	return field
}

//GTBookingInst is a INT field
type GTBookingInst struct{ message.IntValue }

//Tag returns tag.GTBookingInst (427)
func (f GTBookingInst) Tag() fix.Tag { return tag.GTBookingInst }

//BuildGTBookingInst returns a new GTBookingInst initialized with val
func BuildGTBookingInst(val int) GTBookingInst {
	var field GTBookingInst
	field.Value = val
	return field
}

//GapFillFlag is a BOOLEAN field
type GapFillFlag struct{ message.BooleanValue }

//Tag returns tag.GapFillFlag (123)
func (f GapFillFlag) Tag() fix.Tag { return tag.GapFillFlag }

//BuildGapFillFlag returns a new GapFillFlag initialized with val
func BuildGapFillFlag(val bool) GapFillFlag {
	var field GapFillFlag
	field.Value = val
	return field
}

//GrossTradeAmt is a AMT field
type GrossTradeAmt struct{ message.AmtValue }

//Tag returns tag.GrossTradeAmt (381)
func (f GrossTradeAmt) Tag() fix.Tag { return tag.GrossTradeAmt }

//BuildGrossTradeAmt returns a new GrossTradeAmt initialized with val
func BuildGrossTradeAmt(val float64) GrossTradeAmt {
	var field GrossTradeAmt
	field.Value = val
	return field
}

//HaltReasonChar is a CHAR field
type HaltReasonChar struct{ message.CharValue }

//Tag returns tag.HaltReasonChar (327)
func (f HaltReasonChar) Tag() fix.Tag { return tag.HaltReasonChar }

//BuildHaltReasonChar returns a new HaltReasonChar initialized with val
func BuildHaltReasonChar(val string) HaltReasonChar {
	var field HaltReasonChar
	field.Value = val
	return field
}

//HaltReasonInt is a INT field
type HaltReasonInt struct{ message.IntValue }

//Tag returns tag.HaltReasonInt (327)
func (f HaltReasonInt) Tag() fix.Tag { return tag.HaltReasonInt }

//BuildHaltReasonInt returns a new HaltReasonInt initialized with val
func BuildHaltReasonInt(val int) HaltReasonInt {
	var field HaltReasonInt
	field.Value = val
	return field
}

//HandlInst is a CHAR field
type HandlInst struct{ message.CharValue }

//Tag returns tag.HandlInst (21)
func (f HandlInst) Tag() fix.Tag { return tag.HandlInst }

//BuildHandlInst returns a new HandlInst initialized with val
func BuildHandlInst(val string) HandlInst {
	var field HandlInst
	field.Value = val
	return field
}

//Headline is a STRING field
type Headline struct{ message.StringValue }

//Tag returns tag.Headline (148)
func (f Headline) Tag() fix.Tag { return tag.Headline }

//BuildHeadline returns a new Headline initialized with val
func BuildHeadline(val string) Headline {
	var field Headline
	field.Value = val
	return field
}

//HeartBtInt is a INT field
type HeartBtInt struct{ message.IntValue }

//Tag returns tag.HeartBtInt (108)
func (f HeartBtInt) Tag() fix.Tag { return tag.HeartBtInt }

//BuildHeartBtInt returns a new HeartBtInt initialized with val
func BuildHeartBtInt(val int) HeartBtInt {
	var field HeartBtInt
	field.Value = val
	return field
}

//HighLimitPrice is a PRICE field
type HighLimitPrice struct{ message.PriceValue }

//Tag returns tag.HighLimitPrice (1149)
func (f HighLimitPrice) Tag() fix.Tag { return tag.HighLimitPrice }

//BuildHighLimitPrice returns a new HighLimitPrice initialized with val
func BuildHighLimitPrice(val float64) HighLimitPrice {
	var field HighLimitPrice
	field.Value = val
	return field
}

//HighPx is a PRICE field
type HighPx struct{ message.PriceValue }

//Tag returns tag.HighPx (332)
func (f HighPx) Tag() fix.Tag { return tag.HighPx }

//BuildHighPx returns a new HighPx initialized with val
func BuildHighPx(val float64) HighPx {
	var field HighPx
	field.Value = val
	return field
}

//HopCompID is a STRING field
type HopCompID struct{ message.StringValue }

//Tag returns tag.HopCompID (628)
func (f HopCompID) Tag() fix.Tag { return tag.HopCompID }

//BuildHopCompID returns a new HopCompID initialized with val
func BuildHopCompID(val string) HopCompID {
	var field HopCompID
	field.Value = val
	return field
}

//HopRefID is a SEQNUM field
type HopRefID struct{ message.SeqNumValue }

//Tag returns tag.HopRefID (630)
func (f HopRefID) Tag() fix.Tag { return tag.HopRefID }

//BuildHopRefID returns a new HopRefID initialized with val
func BuildHopRefID(val int) HopRefID {
	var field HopRefID
	field.Value = val
	return field
}

//HopSendingTime is a UTCTIMESTAMP field
type HopSendingTime struct{ message.UTCTimestampValue }

//Tag returns tag.HopSendingTime (629)
func (f HopSendingTime) Tag() fix.Tag { return tag.HopSendingTime }

//HostCrossID is a STRING field
type HostCrossID struct{ message.StringValue }

//Tag returns tag.HostCrossID (961)
func (f HostCrossID) Tag() fix.Tag { return tag.HostCrossID }

//BuildHostCrossID returns a new HostCrossID initialized with val
func BuildHostCrossID(val string) HostCrossID {
	var field HostCrossID
	field.Value = val
	return field
}

//IDSource is a STRING field
type IDSource struct{ message.StringValue }

//Tag returns tag.IDSource (22)
func (f IDSource) Tag() fix.Tag { return tag.IDSource }

//BuildIDSource returns a new IDSource initialized with val
func BuildIDSource(val string) IDSource {
	var field IDSource
	field.Value = val
	return field
}

//IOIID is a STRING field
type IOIID struct{ message.StringValue }

//Tag returns tag.IOIID (23)
func (f IOIID) Tag() fix.Tag { return tag.IOIID }

//BuildIOIID returns a new IOIID initialized with val
func BuildIOIID(val string) IOIID {
	var field IOIID
	field.Value = val
	return field
}

//IOINaturalFlag is a BOOLEAN field
type IOINaturalFlag struct{ message.BooleanValue }

//Tag returns tag.IOINaturalFlag (130)
func (f IOINaturalFlag) Tag() fix.Tag { return tag.IOINaturalFlag }

//BuildIOINaturalFlag returns a new IOINaturalFlag initialized with val
func BuildIOINaturalFlag(val bool) IOINaturalFlag {
	var field IOINaturalFlag
	field.Value = val
	return field
}

//IOIOthSvc is a CHAR field
type IOIOthSvc struct{ message.CharValue }

//Tag returns tag.IOIOthSvc (24)
func (f IOIOthSvc) Tag() fix.Tag { return tag.IOIOthSvc }

//BuildIOIOthSvc returns a new IOIOthSvc initialized with val
func BuildIOIOthSvc(val string) IOIOthSvc {
	var field IOIOthSvc
	field.Value = val
	return field
}

//IOIQltyInd is a CHAR field
type IOIQltyInd struct{ message.CharValue }

//Tag returns tag.IOIQltyInd (25)
func (f IOIQltyInd) Tag() fix.Tag { return tag.IOIQltyInd }

//BuildIOIQltyInd returns a new IOIQltyInd initialized with val
func BuildIOIQltyInd(val string) IOIQltyInd {
	var field IOIQltyInd
	field.Value = val
	return field
}

//IOIQty is a STRING field
type IOIQty struct{ message.StringValue }

//Tag returns tag.IOIQty (27)
func (f IOIQty) Tag() fix.Tag { return tag.IOIQty }

//BuildIOIQty returns a new IOIQty initialized with val
func BuildIOIQty(val string) IOIQty {
	var field IOIQty
	field.Value = val
	return field
}

//IOIQualifier is a CHAR field
type IOIQualifier struct{ message.CharValue }

//Tag returns tag.IOIQualifier (104)
func (f IOIQualifier) Tag() fix.Tag { return tag.IOIQualifier }

//BuildIOIQualifier returns a new IOIQualifier initialized with val
func BuildIOIQualifier(val string) IOIQualifier {
	var field IOIQualifier
	field.Value = val
	return field
}

//IOIRefID is a STRING field
type IOIRefID struct{ message.StringValue }

//Tag returns tag.IOIRefID (26)
func (f IOIRefID) Tag() fix.Tag { return tag.IOIRefID }

//BuildIOIRefID returns a new IOIRefID initialized with val
func BuildIOIRefID(val string) IOIRefID {
	var field IOIRefID
	field.Value = val
	return field
}

//IOIShares is a STRING field
type IOIShares struct{ message.StringValue }

//Tag returns tag.IOIShares (27)
func (f IOIShares) Tag() fix.Tag { return tag.IOIShares }

//BuildIOIShares returns a new IOIShares initialized with val
func BuildIOIShares(val string) IOIShares {
	var field IOIShares
	field.Value = val
	return field
}

//IOITransType is a CHAR field
type IOITransType struct{ message.CharValue }

//Tag returns tag.IOITransType (28)
func (f IOITransType) Tag() fix.Tag { return tag.IOITransType }

//BuildIOITransType returns a new IOITransType initialized with val
func BuildIOITransType(val string) IOITransType {
	var field IOITransType
	field.Value = val
	return field
}

//IOIid is a STRING field
type IOIid struct{ message.StringValue }

//Tag returns tag.IOIid (23)
func (f IOIid) Tag() fix.Tag { return tag.IOIid }

//BuildIOIid returns a new IOIid initialized with val
func BuildIOIid(val string) IOIid {
	var field IOIid
	field.Value = val
	return field
}

//ImpliedMarketIndicator is a INT field
type ImpliedMarketIndicator struct{ message.IntValue }

//Tag returns tag.ImpliedMarketIndicator (1144)
func (f ImpliedMarketIndicator) Tag() fix.Tag { return tag.ImpliedMarketIndicator }

//BuildImpliedMarketIndicator returns a new ImpliedMarketIndicator initialized with val
func BuildImpliedMarketIndicator(val int) ImpliedMarketIndicator {
	var field ImpliedMarketIndicator
	field.Value = val
	return field
}

//InViewOfCommon is a BOOLEAN field
type InViewOfCommon struct{ message.BooleanValue }

//Tag returns tag.InViewOfCommon (328)
func (f InViewOfCommon) Tag() fix.Tag { return tag.InViewOfCommon }

//BuildInViewOfCommon returns a new InViewOfCommon initialized with val
func BuildInViewOfCommon(val bool) InViewOfCommon {
	var field InViewOfCommon
	field.Value = val
	return field
}

//IncTaxInd is a INT field
type IncTaxInd struct{ message.IntValue }

//Tag returns tag.IncTaxInd (416)
func (f IncTaxInd) Tag() fix.Tag { return tag.IncTaxInd }

//BuildIncTaxInd returns a new IncTaxInd initialized with val
func BuildIncTaxInd(val int) IncTaxInd {
	var field IncTaxInd
	field.Value = val
	return field
}

//IndividualAllocID is a STRING field
type IndividualAllocID struct{ message.StringValue }

//Tag returns tag.IndividualAllocID (467)
func (f IndividualAllocID) Tag() fix.Tag { return tag.IndividualAllocID }

//BuildIndividualAllocID returns a new IndividualAllocID initialized with val
func BuildIndividualAllocID(val string) IndividualAllocID {
	var field IndividualAllocID
	field.Value = val
	return field
}

//IndividualAllocRejCode is a INT field
type IndividualAllocRejCode struct{ message.IntValue }

//Tag returns tag.IndividualAllocRejCode (776)
func (f IndividualAllocRejCode) Tag() fix.Tag { return tag.IndividualAllocRejCode }

//BuildIndividualAllocRejCode returns a new IndividualAllocRejCode initialized with val
func BuildIndividualAllocRejCode(val int) IndividualAllocRejCode {
	var field IndividualAllocRejCode
	field.Value = val
	return field
}

//IndividualAllocType is a INT field
type IndividualAllocType struct{ message.IntValue }

//Tag returns tag.IndividualAllocType (992)
func (f IndividualAllocType) Tag() fix.Tag { return tag.IndividualAllocType }

//BuildIndividualAllocType returns a new IndividualAllocType initialized with val
func BuildIndividualAllocType(val int) IndividualAllocType {
	var field IndividualAllocType
	field.Value = val
	return field
}

//InputSource is a STRING field
type InputSource struct{ message.StringValue }

//Tag returns tag.InputSource (979)
func (f InputSource) Tag() fix.Tag { return tag.InputSource }

//BuildInputSource returns a new InputSource initialized with val
func BuildInputSource(val string) InputSource {
	var field InputSource
	field.Value = val
	return field
}

//InstrAttribType is a INT field
type InstrAttribType struct{ message.IntValue }

//Tag returns tag.InstrAttribType (871)
func (f InstrAttribType) Tag() fix.Tag { return tag.InstrAttribType }

//BuildInstrAttribType returns a new InstrAttribType initialized with val
func BuildInstrAttribType(val int) InstrAttribType {
	var field InstrAttribType
	field.Value = val
	return field
}

//InstrAttribValue is a STRING field
type InstrAttribValue struct{ message.StringValue }

//Tag returns tag.InstrAttribValue (872)
func (f InstrAttribValue) Tag() fix.Tag { return tag.InstrAttribValue }

//BuildInstrAttribValue returns a new InstrAttribValue initialized with val
func BuildInstrAttribValue(val string) InstrAttribValue {
	var field InstrAttribValue
	field.Value = val
	return field
}

//InstrRegistry is a STRING field
type InstrRegistry struct{ message.StringValue }

//Tag returns tag.InstrRegistry (543)
func (f InstrRegistry) Tag() fix.Tag { return tag.InstrRegistry }

//BuildInstrRegistry returns a new InstrRegistry initialized with val
func BuildInstrRegistry(val string) InstrRegistry {
	var field InstrRegistry
	field.Value = val
	return field
}

//InstrmtAssignmentMethod is a CHAR field
type InstrmtAssignmentMethod struct{ message.CharValue }

//Tag returns tag.InstrmtAssignmentMethod (1049)
func (f InstrmtAssignmentMethod) Tag() fix.Tag { return tag.InstrmtAssignmentMethod }

//BuildInstrmtAssignmentMethod returns a new InstrmtAssignmentMethod initialized with val
func BuildInstrmtAssignmentMethod(val string) InstrmtAssignmentMethod {
	var field InstrmtAssignmentMethod
	field.Value = val
	return field
}

//InstrumentPartyID is a STRING field
type InstrumentPartyID struct{ message.StringValue }

//Tag returns tag.InstrumentPartyID (1019)
func (f InstrumentPartyID) Tag() fix.Tag { return tag.InstrumentPartyID }

//BuildInstrumentPartyID returns a new InstrumentPartyID initialized with val
func BuildInstrumentPartyID(val string) InstrumentPartyID {
	var field InstrumentPartyID
	field.Value = val
	return field
}

//InstrumentPartyIDSource is a CHAR field
type InstrumentPartyIDSource struct{ message.CharValue }

//Tag returns tag.InstrumentPartyIDSource (1050)
func (f InstrumentPartyIDSource) Tag() fix.Tag { return tag.InstrumentPartyIDSource }

//BuildInstrumentPartyIDSource returns a new InstrumentPartyIDSource initialized with val
func BuildInstrumentPartyIDSource(val string) InstrumentPartyIDSource {
	var field InstrumentPartyIDSource
	field.Value = val
	return field
}

//InstrumentPartyRole is a INT field
type InstrumentPartyRole struct{ message.IntValue }

//Tag returns tag.InstrumentPartyRole (1051)
func (f InstrumentPartyRole) Tag() fix.Tag { return tag.InstrumentPartyRole }

//BuildInstrumentPartyRole returns a new InstrumentPartyRole initialized with val
func BuildInstrumentPartyRole(val int) InstrumentPartyRole {
	var field InstrumentPartyRole
	field.Value = val
	return field
}

//InstrumentPartySubID is a STRING field
type InstrumentPartySubID struct{ message.StringValue }

//Tag returns tag.InstrumentPartySubID (1053)
func (f InstrumentPartySubID) Tag() fix.Tag { return tag.InstrumentPartySubID }

//BuildInstrumentPartySubID returns a new InstrumentPartySubID initialized with val
func BuildInstrumentPartySubID(val string) InstrumentPartySubID {
	var field InstrumentPartySubID
	field.Value = val
	return field
}

//InstrumentPartySubIDType is a INT field
type InstrumentPartySubIDType struct{ message.IntValue }

//Tag returns tag.InstrumentPartySubIDType (1054)
func (f InstrumentPartySubIDType) Tag() fix.Tag { return tag.InstrumentPartySubIDType }

//BuildInstrumentPartySubIDType returns a new InstrumentPartySubIDType initialized with val
func BuildInstrumentPartySubIDType(val int) InstrumentPartySubIDType {
	var field InstrumentPartySubIDType
	field.Value = val
	return field
}

//InterestAccrualDate is a LOCALMKTDATE field
type InterestAccrualDate struct{ message.LocalMktDateValue }

//Tag returns tag.InterestAccrualDate (874)
func (f InterestAccrualDate) Tag() fix.Tag { return tag.InterestAccrualDate }

//BuildInterestAccrualDate returns a new InterestAccrualDate initialized with val
func BuildInterestAccrualDate(val string) InterestAccrualDate {
	var field InterestAccrualDate
	field.Value = val
	return field
}

//InterestAtMaturity is a AMT field
type InterestAtMaturity struct{ message.AmtValue }

//Tag returns tag.InterestAtMaturity (738)
func (f InterestAtMaturity) Tag() fix.Tag { return tag.InterestAtMaturity }

//BuildInterestAtMaturity returns a new InterestAtMaturity initialized with val
func BuildInterestAtMaturity(val float64) InterestAtMaturity {
	var field InterestAtMaturity
	field.Value = val
	return field
}

//InvestorCountryOfResidence is a COUNTRY field
type InvestorCountryOfResidence struct{ message.CountryValue }

//Tag returns tag.InvestorCountryOfResidence (475)
func (f InvestorCountryOfResidence) Tag() fix.Tag { return tag.InvestorCountryOfResidence }

//BuildInvestorCountryOfResidence returns a new InvestorCountryOfResidence initialized with val
func BuildInvestorCountryOfResidence(val string) InvestorCountryOfResidence {
	var field InvestorCountryOfResidence
	field.Value = val
	return field
}

//IssueDate is a LOCALMKTDATE field
type IssueDate struct{ message.LocalMktDateValue }

//Tag returns tag.IssueDate (225)
func (f IssueDate) Tag() fix.Tag { return tag.IssueDate }

//BuildIssueDate returns a new IssueDate initialized with val
func BuildIssueDate(val string) IssueDate {
	var field IssueDate
	field.Value = val
	return field
}

//Issuer is a STRING field
type Issuer struct{ message.StringValue }

//Tag returns tag.Issuer (106)
func (f Issuer) Tag() fix.Tag { return tag.Issuer }

//BuildIssuer returns a new Issuer initialized with val
func BuildIssuer(val string) Issuer {
	var field Issuer
	field.Value = val
	return field
}

//LanguageCode is a LANGUAGE field
type LanguageCode struct{ message.LanguageValue }

//Tag returns tag.LanguageCode (1474)
func (f LanguageCode) Tag() fix.Tag { return tag.LanguageCode }

//BuildLanguageCode returns a new LanguageCode initialized with val
func BuildLanguageCode(val string) LanguageCode {
	var field LanguageCode
	field.Value = val
	return field
}

//LastCapacity is a CHAR field
type LastCapacity struct{ message.CharValue }

//Tag returns tag.LastCapacity (29)
func (f LastCapacity) Tag() fix.Tag { return tag.LastCapacity }

//BuildLastCapacity returns a new LastCapacity initialized with val
func BuildLastCapacity(val string) LastCapacity {
	var field LastCapacity
	field.Value = val
	return field
}

//LastForwardPoints is a PRICEOFFSET field
type LastForwardPoints struct{ message.PriceOffsetValue }

//Tag returns tag.LastForwardPoints (195)
func (f LastForwardPoints) Tag() fix.Tag { return tag.LastForwardPoints }

//BuildLastForwardPoints returns a new LastForwardPoints initialized with val
func BuildLastForwardPoints(val float64) LastForwardPoints {
	var field LastForwardPoints
	field.Value = val
	return field
}

//LastForwardPoints2 is a PRICEOFFSET field
type LastForwardPoints2 struct{ message.PriceOffsetValue }

//Tag returns tag.LastForwardPoints2 (641)
func (f LastForwardPoints2) Tag() fix.Tag { return tag.LastForwardPoints2 }

//BuildLastForwardPoints2 returns a new LastForwardPoints2 initialized with val
func BuildLastForwardPoints2(val float64) LastForwardPoints2 {
	var field LastForwardPoints2
	field.Value = val
	return field
}

//LastFragment is a BOOLEAN field
type LastFragment struct{ message.BooleanValue }

//Tag returns tag.LastFragment (893)
func (f LastFragment) Tag() fix.Tag { return tag.LastFragment }

//BuildLastFragment returns a new LastFragment initialized with val
func BuildLastFragment(val bool) LastFragment {
	var field LastFragment
	field.Value = val
	return field
}

//LastLiquidityInd is a INT field
type LastLiquidityInd struct{ message.IntValue }

//Tag returns tag.LastLiquidityInd (851)
func (f LastLiquidityInd) Tag() fix.Tag { return tag.LastLiquidityInd }

//BuildLastLiquidityInd returns a new LastLiquidityInd initialized with val
func BuildLastLiquidityInd(val int) LastLiquidityInd {
	var field LastLiquidityInd
	field.Value = val
	return field
}

//LastMkt is a EXCHANGE field
type LastMkt struct{ message.ExchangeValue }

//Tag returns tag.LastMkt (30)
func (f LastMkt) Tag() fix.Tag { return tag.LastMkt }

//BuildLastMkt returns a new LastMkt initialized with val
func BuildLastMkt(val string) LastMkt {
	var field LastMkt
	field.Value = val
	return field
}

//LastMsgSeqNumProcessed is a SEQNUM field
type LastMsgSeqNumProcessed struct{ message.SeqNumValue }

//Tag returns tag.LastMsgSeqNumProcessed (369)
func (f LastMsgSeqNumProcessed) Tag() fix.Tag { return tag.LastMsgSeqNumProcessed }

//BuildLastMsgSeqNumProcessed returns a new LastMsgSeqNumProcessed initialized with val
func BuildLastMsgSeqNumProcessed(val int) LastMsgSeqNumProcessed {
	var field LastMsgSeqNumProcessed
	field.Value = val
	return field
}

//LastNetworkResponseID is a STRING field
type LastNetworkResponseID struct{ message.StringValue }

//Tag returns tag.LastNetworkResponseID (934)
func (f LastNetworkResponseID) Tag() fix.Tag { return tag.LastNetworkResponseID }

//BuildLastNetworkResponseID returns a new LastNetworkResponseID initialized with val
func BuildLastNetworkResponseID(val string) LastNetworkResponseID {
	var field LastNetworkResponseID
	field.Value = val
	return field
}

//LastParPx is a PRICE field
type LastParPx struct{ message.PriceValue }

//Tag returns tag.LastParPx (669)
func (f LastParPx) Tag() fix.Tag { return tag.LastParPx }

//BuildLastParPx returns a new LastParPx initialized with val
func BuildLastParPx(val float64) LastParPx {
	var field LastParPx
	field.Value = val
	return field
}

//LastPx is a PRICE field
type LastPx struct{ message.PriceValue }

//Tag returns tag.LastPx (31)
func (f LastPx) Tag() fix.Tag { return tag.LastPx }

//BuildLastPx returns a new LastPx initialized with val
func BuildLastPx(val float64) LastPx {
	var field LastPx
	field.Value = val
	return field
}

//LastQty is a QTY field
type LastQty struct{ message.QtyValue }

//Tag returns tag.LastQty (32)
func (f LastQty) Tag() fix.Tag { return tag.LastQty }

//BuildLastQty returns a new LastQty initialized with val
func BuildLastQty(val float64) LastQty {
	var field LastQty
	field.Value = val
	return field
}

//LastRptRequested is a BOOLEAN field
type LastRptRequested struct{ message.BooleanValue }

//Tag returns tag.LastRptRequested (912)
func (f LastRptRequested) Tag() fix.Tag { return tag.LastRptRequested }

//BuildLastRptRequested returns a new LastRptRequested initialized with val
func BuildLastRptRequested(val bool) LastRptRequested {
	var field LastRptRequested
	field.Value = val
	return field
}

//LastShares is a QTY field
type LastShares struct{ message.QtyValue }

//Tag returns tag.LastShares (32)
func (f LastShares) Tag() fix.Tag { return tag.LastShares }

//BuildLastShares returns a new LastShares initialized with val
func BuildLastShares(val float64) LastShares {
	var field LastShares
	field.Value = val
	return field
}

//LastSpotRate is a PRICE field
type LastSpotRate struct{ message.PriceValue }

//Tag returns tag.LastSpotRate (194)
func (f LastSpotRate) Tag() fix.Tag { return tag.LastSpotRate }

//BuildLastSpotRate returns a new LastSpotRate initialized with val
func BuildLastSpotRate(val float64) LastSpotRate {
	var field LastSpotRate
	field.Value = val
	return field
}

//LastSwapPoints is a PRICEOFFSET field
type LastSwapPoints struct{ message.PriceOffsetValue }

//Tag returns tag.LastSwapPoints (1071)
func (f LastSwapPoints) Tag() fix.Tag { return tag.LastSwapPoints }

//BuildLastSwapPoints returns a new LastSwapPoints initialized with val
func BuildLastSwapPoints(val float64) LastSwapPoints {
	var field LastSwapPoints
	field.Value = val
	return field
}

//LastUpdateTime is a UTCTIMESTAMP field
type LastUpdateTime struct{ message.UTCTimestampValue }

//Tag returns tag.LastUpdateTime (779)
func (f LastUpdateTime) Tag() fix.Tag { return tag.LastUpdateTime }

//LateIndicator is a BOOLEAN field
type LateIndicator struct{ message.BooleanValue }

//Tag returns tag.LateIndicator (978)
func (f LateIndicator) Tag() fix.Tag { return tag.LateIndicator }

//BuildLateIndicator returns a new LateIndicator initialized with val
func BuildLateIndicator(val bool) LateIndicator {
	var field LateIndicator
	field.Value = val
	return field
}

//LeavesQty is a QTY field
type LeavesQty struct{ message.QtyValue }

//Tag returns tag.LeavesQty (151)
func (f LeavesQty) Tag() fix.Tag { return tag.LeavesQty }

//BuildLeavesQty returns a new LeavesQty initialized with val
func BuildLeavesQty(val float64) LeavesQty {
	var field LeavesQty
	field.Value = val
	return field
}

//LegAllocAccount is a STRING field
type LegAllocAccount struct{ message.StringValue }

//Tag returns tag.LegAllocAccount (671)
func (f LegAllocAccount) Tag() fix.Tag { return tag.LegAllocAccount }

//BuildLegAllocAccount returns a new LegAllocAccount initialized with val
func BuildLegAllocAccount(val string) LegAllocAccount {
	var field LegAllocAccount
	field.Value = val
	return field
}

//LegAllocAcctIDSource is a STRING field
type LegAllocAcctIDSource struct{ message.StringValue }

//Tag returns tag.LegAllocAcctIDSource (674)
func (f LegAllocAcctIDSource) Tag() fix.Tag { return tag.LegAllocAcctIDSource }

//BuildLegAllocAcctIDSource returns a new LegAllocAcctIDSource initialized with val
func BuildLegAllocAcctIDSource(val string) LegAllocAcctIDSource {
	var field LegAllocAcctIDSource
	field.Value = val
	return field
}

//LegAllocID is a STRING field
type LegAllocID struct{ message.StringValue }

//Tag returns tag.LegAllocID (1366)
func (f LegAllocID) Tag() fix.Tag { return tag.LegAllocID }

//BuildLegAllocID returns a new LegAllocID initialized with val
func BuildLegAllocID(val string) LegAllocID {
	var field LegAllocID
	field.Value = val
	return field
}

//LegAllocQty is a QTY field
type LegAllocQty struct{ message.QtyValue }

//Tag returns tag.LegAllocQty (673)
func (f LegAllocQty) Tag() fix.Tag { return tag.LegAllocQty }

//BuildLegAllocQty returns a new LegAllocQty initialized with val
func BuildLegAllocQty(val float64) LegAllocQty {
	var field LegAllocQty
	field.Value = val
	return field
}

//LegAllocSettlCurrency is a CURRENCY field
type LegAllocSettlCurrency struct{ message.CurrencyValue }

//Tag returns tag.LegAllocSettlCurrency (1367)
func (f LegAllocSettlCurrency) Tag() fix.Tag { return tag.LegAllocSettlCurrency }

//BuildLegAllocSettlCurrency returns a new LegAllocSettlCurrency initialized with val
func BuildLegAllocSettlCurrency(val string) LegAllocSettlCurrency {
	var field LegAllocSettlCurrency
	field.Value = val
	return field
}

//LegBenchmarkCurveCurrency is a CURRENCY field
type LegBenchmarkCurveCurrency struct{ message.CurrencyValue }

//Tag returns tag.LegBenchmarkCurveCurrency (676)
func (f LegBenchmarkCurveCurrency) Tag() fix.Tag { return tag.LegBenchmarkCurveCurrency }

//BuildLegBenchmarkCurveCurrency returns a new LegBenchmarkCurveCurrency initialized with val
func BuildLegBenchmarkCurveCurrency(val string) LegBenchmarkCurveCurrency {
	var field LegBenchmarkCurveCurrency
	field.Value = val
	return field
}

//LegBenchmarkCurveName is a STRING field
type LegBenchmarkCurveName struct{ message.StringValue }

//Tag returns tag.LegBenchmarkCurveName (677)
func (f LegBenchmarkCurveName) Tag() fix.Tag { return tag.LegBenchmarkCurveName }

//BuildLegBenchmarkCurveName returns a new LegBenchmarkCurveName initialized with val
func BuildLegBenchmarkCurveName(val string) LegBenchmarkCurveName {
	var field LegBenchmarkCurveName
	field.Value = val
	return field
}

//LegBenchmarkCurvePoint is a STRING field
type LegBenchmarkCurvePoint struct{ message.StringValue }

//Tag returns tag.LegBenchmarkCurvePoint (678)
func (f LegBenchmarkCurvePoint) Tag() fix.Tag { return tag.LegBenchmarkCurvePoint }

//BuildLegBenchmarkCurvePoint returns a new LegBenchmarkCurvePoint initialized with val
func BuildLegBenchmarkCurvePoint(val string) LegBenchmarkCurvePoint {
	var field LegBenchmarkCurvePoint
	field.Value = val
	return field
}

//LegBenchmarkPrice is a PRICE field
type LegBenchmarkPrice struct{ message.PriceValue }

//Tag returns tag.LegBenchmarkPrice (679)
func (f LegBenchmarkPrice) Tag() fix.Tag { return tag.LegBenchmarkPrice }

//BuildLegBenchmarkPrice returns a new LegBenchmarkPrice initialized with val
func BuildLegBenchmarkPrice(val float64) LegBenchmarkPrice {
	var field LegBenchmarkPrice
	field.Value = val
	return field
}

//LegBenchmarkPriceType is a INT field
type LegBenchmarkPriceType struct{ message.IntValue }

//Tag returns tag.LegBenchmarkPriceType (680)
func (f LegBenchmarkPriceType) Tag() fix.Tag { return tag.LegBenchmarkPriceType }

//BuildLegBenchmarkPriceType returns a new LegBenchmarkPriceType initialized with val
func BuildLegBenchmarkPriceType(val int) LegBenchmarkPriceType {
	var field LegBenchmarkPriceType
	field.Value = val
	return field
}

//LegBidForwardPoints is a PRICEOFFSET field
type LegBidForwardPoints struct{ message.PriceOffsetValue }

//Tag returns tag.LegBidForwardPoints (1067)
func (f LegBidForwardPoints) Tag() fix.Tag { return tag.LegBidForwardPoints }

//BuildLegBidForwardPoints returns a new LegBidForwardPoints initialized with val
func BuildLegBidForwardPoints(val float64) LegBidForwardPoints {
	var field LegBidForwardPoints
	field.Value = val
	return field
}

//LegBidPx is a PRICE field
type LegBidPx struct{ message.PriceValue }

//Tag returns tag.LegBidPx (681)
func (f LegBidPx) Tag() fix.Tag { return tag.LegBidPx }

//BuildLegBidPx returns a new LegBidPx initialized with val
func BuildLegBidPx(val float64) LegBidPx {
	var field LegBidPx
	field.Value = val
	return field
}

//LegCFICode is a STRING field
type LegCFICode struct{ message.StringValue }

//Tag returns tag.LegCFICode (608)
func (f LegCFICode) Tag() fix.Tag { return tag.LegCFICode }

//BuildLegCFICode returns a new LegCFICode initialized with val
func BuildLegCFICode(val string) LegCFICode {
	var field LegCFICode
	field.Value = val
	return field
}

//LegCalculatedCcyLastQty is a QTY field
type LegCalculatedCcyLastQty struct{ message.QtyValue }

//Tag returns tag.LegCalculatedCcyLastQty (1074)
func (f LegCalculatedCcyLastQty) Tag() fix.Tag { return tag.LegCalculatedCcyLastQty }

//BuildLegCalculatedCcyLastQty returns a new LegCalculatedCcyLastQty initialized with val
func BuildLegCalculatedCcyLastQty(val float64) LegCalculatedCcyLastQty {
	var field LegCalculatedCcyLastQty
	field.Value = val
	return field
}

//LegContractMultiplier is a FLOAT field
type LegContractMultiplier struct{ message.FloatValue }

//Tag returns tag.LegContractMultiplier (614)
func (f LegContractMultiplier) Tag() fix.Tag { return tag.LegContractMultiplier }

//BuildLegContractMultiplier returns a new LegContractMultiplier initialized with val
func BuildLegContractMultiplier(val float64) LegContractMultiplier {
	var field LegContractMultiplier
	field.Value = val
	return field
}

//LegContractMultiplierUnit is a INT field
type LegContractMultiplierUnit struct{ message.IntValue }

//Tag returns tag.LegContractMultiplierUnit (1436)
func (f LegContractMultiplierUnit) Tag() fix.Tag { return tag.LegContractMultiplierUnit }

//BuildLegContractMultiplierUnit returns a new LegContractMultiplierUnit initialized with val
func BuildLegContractMultiplierUnit(val int) LegContractMultiplierUnit {
	var field LegContractMultiplierUnit
	field.Value = val
	return field
}

//LegContractSettlMonth is a MONTHYEAR field
type LegContractSettlMonth struct{ message.MonthYearValue }

//Tag returns tag.LegContractSettlMonth (955)
func (f LegContractSettlMonth) Tag() fix.Tag { return tag.LegContractSettlMonth }

//BuildLegContractSettlMonth returns a new LegContractSettlMonth initialized with val
func BuildLegContractSettlMonth(val string) LegContractSettlMonth {
	var field LegContractSettlMonth
	field.Value = val
	return field
}

//LegCountryOfIssue is a COUNTRY field
type LegCountryOfIssue struct{ message.CountryValue }

//Tag returns tag.LegCountryOfIssue (596)
func (f LegCountryOfIssue) Tag() fix.Tag { return tag.LegCountryOfIssue }

//BuildLegCountryOfIssue returns a new LegCountryOfIssue initialized with val
func BuildLegCountryOfIssue(val string) LegCountryOfIssue {
	var field LegCountryOfIssue
	field.Value = val
	return field
}

//LegCouponPaymentDate is a LOCALMKTDATE field
type LegCouponPaymentDate struct{ message.LocalMktDateValue }

//Tag returns tag.LegCouponPaymentDate (248)
func (f LegCouponPaymentDate) Tag() fix.Tag { return tag.LegCouponPaymentDate }

//BuildLegCouponPaymentDate returns a new LegCouponPaymentDate initialized with val
func BuildLegCouponPaymentDate(val string) LegCouponPaymentDate {
	var field LegCouponPaymentDate
	field.Value = val
	return field
}

//LegCouponRate is a PERCENTAGE field
type LegCouponRate struct{ message.PercentageValue }

//Tag returns tag.LegCouponRate (615)
func (f LegCouponRate) Tag() fix.Tag { return tag.LegCouponRate }

//BuildLegCouponRate returns a new LegCouponRate initialized with val
func BuildLegCouponRate(val float64) LegCouponRate {
	var field LegCouponRate
	field.Value = val
	return field
}

//LegCoveredOrUncovered is a INT field
type LegCoveredOrUncovered struct{ message.IntValue }

//Tag returns tag.LegCoveredOrUncovered (565)
func (f LegCoveredOrUncovered) Tag() fix.Tag { return tag.LegCoveredOrUncovered }

//BuildLegCoveredOrUncovered returns a new LegCoveredOrUncovered initialized with val
func BuildLegCoveredOrUncovered(val int) LegCoveredOrUncovered {
	var field LegCoveredOrUncovered
	field.Value = val
	return field
}

//LegCreditRating is a STRING field
type LegCreditRating struct{ message.StringValue }

//Tag returns tag.LegCreditRating (257)
func (f LegCreditRating) Tag() fix.Tag { return tag.LegCreditRating }

//BuildLegCreditRating returns a new LegCreditRating initialized with val
func BuildLegCreditRating(val string) LegCreditRating {
	var field LegCreditRating
	field.Value = val
	return field
}

//LegCurrency is a CURRENCY field
type LegCurrency struct{ message.CurrencyValue }

//Tag returns tag.LegCurrency (556)
func (f LegCurrency) Tag() fix.Tag { return tag.LegCurrency }

//BuildLegCurrency returns a new LegCurrency initialized with val
func BuildLegCurrency(val string) LegCurrency {
	var field LegCurrency
	field.Value = val
	return field
}

//LegCurrencyRatio is a FLOAT field
type LegCurrencyRatio struct{ message.FloatValue }

//Tag returns tag.LegCurrencyRatio (1383)
func (f LegCurrencyRatio) Tag() fix.Tag { return tag.LegCurrencyRatio }

//BuildLegCurrencyRatio returns a new LegCurrencyRatio initialized with val
func BuildLegCurrencyRatio(val float64) LegCurrencyRatio {
	var field LegCurrencyRatio
	field.Value = val
	return field
}

//LegDatedDate is a LOCALMKTDATE field
type LegDatedDate struct{ message.LocalMktDateValue }

//Tag returns tag.LegDatedDate (739)
func (f LegDatedDate) Tag() fix.Tag { return tag.LegDatedDate }

//BuildLegDatedDate returns a new LegDatedDate initialized with val
func BuildLegDatedDate(val string) LegDatedDate {
	var field LegDatedDate
	field.Value = val
	return field
}

//LegDividendYield is a PERCENTAGE field
type LegDividendYield struct{ message.PercentageValue }

//Tag returns tag.LegDividendYield (1381)
func (f LegDividendYield) Tag() fix.Tag { return tag.LegDividendYield }

//BuildLegDividendYield returns a new LegDividendYield initialized with val
func BuildLegDividendYield(val float64) LegDividendYield {
	var field LegDividendYield
	field.Value = val
	return field
}

//LegExecInst is a MULTIPLECHARVALUE field
type LegExecInst struct{ message.MultipleCharValue }

//Tag returns tag.LegExecInst (1384)
func (f LegExecInst) Tag() fix.Tag { return tag.LegExecInst }

//BuildLegExecInst returns a new LegExecInst initialized with val
func BuildLegExecInst(val string) LegExecInst {
	var field LegExecInst
	field.Value = val
	return field
}

//LegExerciseStyle is a INT field
type LegExerciseStyle struct{ message.IntValue }

//Tag returns tag.LegExerciseStyle (1420)
func (f LegExerciseStyle) Tag() fix.Tag { return tag.LegExerciseStyle }

//BuildLegExerciseStyle returns a new LegExerciseStyle initialized with val
func BuildLegExerciseStyle(val int) LegExerciseStyle {
	var field LegExerciseStyle
	field.Value = val
	return field
}

//LegFactor is a FLOAT field
type LegFactor struct{ message.FloatValue }

//Tag returns tag.LegFactor (253)
func (f LegFactor) Tag() fix.Tag { return tag.LegFactor }

//BuildLegFactor returns a new LegFactor initialized with val
func BuildLegFactor(val float64) LegFactor {
	var field LegFactor
	field.Value = val
	return field
}

//LegFlowScheduleType is a INT field
type LegFlowScheduleType struct{ message.IntValue }

//Tag returns tag.LegFlowScheduleType (1440)
func (f LegFlowScheduleType) Tag() fix.Tag { return tag.LegFlowScheduleType }

//BuildLegFlowScheduleType returns a new LegFlowScheduleType initialized with val
func BuildLegFlowScheduleType(val int) LegFlowScheduleType {
	var field LegFlowScheduleType
	field.Value = val
	return field
}

//LegFutSettDate is a LOCALMKTDATE field
type LegFutSettDate struct{ message.LocalMktDateValue }

//Tag returns tag.LegFutSettDate (588)
func (f LegFutSettDate) Tag() fix.Tag { return tag.LegFutSettDate }

//BuildLegFutSettDate returns a new LegFutSettDate initialized with val
func BuildLegFutSettDate(val string) LegFutSettDate {
	var field LegFutSettDate
	field.Value = val
	return field
}

//LegGrossTradeAmt is a AMT field
type LegGrossTradeAmt struct{ message.AmtValue }

//Tag returns tag.LegGrossTradeAmt (1075)
func (f LegGrossTradeAmt) Tag() fix.Tag { return tag.LegGrossTradeAmt }

//BuildLegGrossTradeAmt returns a new LegGrossTradeAmt initialized with val
func BuildLegGrossTradeAmt(val float64) LegGrossTradeAmt {
	var field LegGrossTradeAmt
	field.Value = val
	return field
}

//LegIOIQty is a STRING field
type LegIOIQty struct{ message.StringValue }

//Tag returns tag.LegIOIQty (682)
func (f LegIOIQty) Tag() fix.Tag { return tag.LegIOIQty }

//BuildLegIOIQty returns a new LegIOIQty initialized with val
func BuildLegIOIQty(val string) LegIOIQty {
	var field LegIOIQty
	field.Value = val
	return field
}

//LegIndividualAllocID is a STRING field
type LegIndividualAllocID struct{ message.StringValue }

//Tag returns tag.LegIndividualAllocID (672)
func (f LegIndividualAllocID) Tag() fix.Tag { return tag.LegIndividualAllocID }

//BuildLegIndividualAllocID returns a new LegIndividualAllocID initialized with val
func BuildLegIndividualAllocID(val string) LegIndividualAllocID {
	var field LegIndividualAllocID
	field.Value = val
	return field
}

//LegInstrRegistry is a STRING field
type LegInstrRegistry struct{ message.StringValue }

//Tag returns tag.LegInstrRegistry (599)
func (f LegInstrRegistry) Tag() fix.Tag { return tag.LegInstrRegistry }

//BuildLegInstrRegistry returns a new LegInstrRegistry initialized with val
func BuildLegInstrRegistry(val string) LegInstrRegistry {
	var field LegInstrRegistry
	field.Value = val
	return field
}

//LegInterestAccrualDate is a LOCALMKTDATE field
type LegInterestAccrualDate struct{ message.LocalMktDateValue }

//Tag returns tag.LegInterestAccrualDate (956)
func (f LegInterestAccrualDate) Tag() fix.Tag { return tag.LegInterestAccrualDate }

//BuildLegInterestAccrualDate returns a new LegInterestAccrualDate initialized with val
func BuildLegInterestAccrualDate(val string) LegInterestAccrualDate {
	var field LegInterestAccrualDate
	field.Value = val
	return field
}

//LegIssueDate is a LOCALMKTDATE field
type LegIssueDate struct{ message.LocalMktDateValue }

//Tag returns tag.LegIssueDate (249)
func (f LegIssueDate) Tag() fix.Tag { return tag.LegIssueDate }

//BuildLegIssueDate returns a new LegIssueDate initialized with val
func BuildLegIssueDate(val string) LegIssueDate {
	var field LegIssueDate
	field.Value = val
	return field
}

//LegIssuer is a STRING field
type LegIssuer struct{ message.StringValue }

//Tag returns tag.LegIssuer (617)
func (f LegIssuer) Tag() fix.Tag { return tag.LegIssuer }

//BuildLegIssuer returns a new LegIssuer initialized with val
func BuildLegIssuer(val string) LegIssuer {
	var field LegIssuer
	field.Value = val
	return field
}

//LegLastForwardPoints is a PRICEOFFSET field
type LegLastForwardPoints struct{ message.PriceOffsetValue }

//Tag returns tag.LegLastForwardPoints (1073)
func (f LegLastForwardPoints) Tag() fix.Tag { return tag.LegLastForwardPoints }

//BuildLegLastForwardPoints returns a new LegLastForwardPoints initialized with val
func BuildLegLastForwardPoints(val float64) LegLastForwardPoints {
	var field LegLastForwardPoints
	field.Value = val
	return field
}

//LegLastPx is a PRICE field
type LegLastPx struct{ message.PriceValue }

//Tag returns tag.LegLastPx (637)
func (f LegLastPx) Tag() fix.Tag { return tag.LegLastPx }

//BuildLegLastPx returns a new LegLastPx initialized with val
func BuildLegLastPx(val float64) LegLastPx {
	var field LegLastPx
	field.Value = val
	return field
}

//LegLastQty is a QTY field
type LegLastQty struct{ message.QtyValue }

//Tag returns tag.LegLastQty (1418)
func (f LegLastQty) Tag() fix.Tag { return tag.LegLastQty }

//BuildLegLastQty returns a new LegLastQty initialized with val
func BuildLegLastQty(val float64) LegLastQty {
	var field LegLastQty
	field.Value = val
	return field
}

//LegLocaleOfIssue is a STRING field
type LegLocaleOfIssue struct{ message.StringValue }

//Tag returns tag.LegLocaleOfIssue (598)
func (f LegLocaleOfIssue) Tag() fix.Tag { return tag.LegLocaleOfIssue }

//BuildLegLocaleOfIssue returns a new LegLocaleOfIssue initialized with val
func BuildLegLocaleOfIssue(val string) LegLocaleOfIssue {
	var field LegLocaleOfIssue
	field.Value = val
	return field
}

//LegMaturityDate is a LOCALMKTDATE field
type LegMaturityDate struct{ message.LocalMktDateValue }

//Tag returns tag.LegMaturityDate (611)
func (f LegMaturityDate) Tag() fix.Tag { return tag.LegMaturityDate }

//BuildLegMaturityDate returns a new LegMaturityDate initialized with val
func BuildLegMaturityDate(val string) LegMaturityDate {
	var field LegMaturityDate
	field.Value = val
	return field
}

//LegMaturityMonthYear is a MONTHYEAR field
type LegMaturityMonthYear struct{ message.MonthYearValue }

//Tag returns tag.LegMaturityMonthYear (610)
func (f LegMaturityMonthYear) Tag() fix.Tag { return tag.LegMaturityMonthYear }

//BuildLegMaturityMonthYear returns a new LegMaturityMonthYear initialized with val
func BuildLegMaturityMonthYear(val string) LegMaturityMonthYear {
	var field LegMaturityMonthYear
	field.Value = val
	return field
}

//LegMaturityTime is a TZTIMEONLY field
type LegMaturityTime struct{ message.TZTimeOnlyValue }

//Tag returns tag.LegMaturityTime (1212)
func (f LegMaturityTime) Tag() fix.Tag { return tag.LegMaturityTime }

//LegNumber is a INT field
type LegNumber struct{ message.IntValue }

//Tag returns tag.LegNumber (1152)
func (f LegNumber) Tag() fix.Tag { return tag.LegNumber }

//BuildLegNumber returns a new LegNumber initialized with val
func BuildLegNumber(val int) LegNumber {
	var field LegNumber
	field.Value = val
	return field
}

//LegOfferForwardPoints is a PRICEOFFSET field
type LegOfferForwardPoints struct{ message.PriceOffsetValue }

//Tag returns tag.LegOfferForwardPoints (1068)
func (f LegOfferForwardPoints) Tag() fix.Tag { return tag.LegOfferForwardPoints }

//BuildLegOfferForwardPoints returns a new LegOfferForwardPoints initialized with val
func BuildLegOfferForwardPoints(val float64) LegOfferForwardPoints {
	var field LegOfferForwardPoints
	field.Value = val
	return field
}

//LegOfferPx is a PRICE field
type LegOfferPx struct{ message.PriceValue }

//Tag returns tag.LegOfferPx (684)
func (f LegOfferPx) Tag() fix.Tag { return tag.LegOfferPx }

//BuildLegOfferPx returns a new LegOfferPx initialized with val
func BuildLegOfferPx(val float64) LegOfferPx {
	var field LegOfferPx
	field.Value = val
	return field
}

//LegOptAttribute is a CHAR field
type LegOptAttribute struct{ message.CharValue }

//Tag returns tag.LegOptAttribute (613)
func (f LegOptAttribute) Tag() fix.Tag { return tag.LegOptAttribute }

//BuildLegOptAttribute returns a new LegOptAttribute initialized with val
func BuildLegOptAttribute(val string) LegOptAttribute {
	var field LegOptAttribute
	field.Value = val
	return field
}

//LegOptionRatio is a FLOAT field
type LegOptionRatio struct{ message.FloatValue }

//Tag returns tag.LegOptionRatio (1017)
func (f LegOptionRatio) Tag() fix.Tag { return tag.LegOptionRatio }

//BuildLegOptionRatio returns a new LegOptionRatio initialized with val
func BuildLegOptionRatio(val float64) LegOptionRatio {
	var field LegOptionRatio
	field.Value = val
	return field
}

//LegOrderQty is a QTY field
type LegOrderQty struct{ message.QtyValue }

//Tag returns tag.LegOrderQty (685)
func (f LegOrderQty) Tag() fix.Tag { return tag.LegOrderQty }

//BuildLegOrderQty returns a new LegOrderQty initialized with val
func BuildLegOrderQty(val float64) LegOrderQty {
	var field LegOrderQty
	field.Value = val
	return field
}

//LegPool is a STRING field
type LegPool struct{ message.StringValue }

//Tag returns tag.LegPool (740)
func (f LegPool) Tag() fix.Tag { return tag.LegPool }

//BuildLegPool returns a new LegPool initialized with val
func BuildLegPool(val string) LegPool {
	var field LegPool
	field.Value = val
	return field
}

//LegPositionEffect is a CHAR field
type LegPositionEffect struct{ message.CharValue }

//Tag returns tag.LegPositionEffect (564)
func (f LegPositionEffect) Tag() fix.Tag { return tag.LegPositionEffect }

//BuildLegPositionEffect returns a new LegPositionEffect initialized with val
func BuildLegPositionEffect(val string) LegPositionEffect {
	var field LegPositionEffect
	field.Value = val
	return field
}

//LegPrice is a PRICE field
type LegPrice struct{ message.PriceValue }

//Tag returns tag.LegPrice (566)
func (f LegPrice) Tag() fix.Tag { return tag.LegPrice }

//BuildLegPrice returns a new LegPrice initialized with val
func BuildLegPrice(val float64) LegPrice {
	var field LegPrice
	field.Value = val
	return field
}

//LegPriceType is a INT field
type LegPriceType struct{ message.IntValue }

//Tag returns tag.LegPriceType (686)
func (f LegPriceType) Tag() fix.Tag { return tag.LegPriceType }

//BuildLegPriceType returns a new LegPriceType initialized with val
func BuildLegPriceType(val int) LegPriceType {
	var field LegPriceType
	field.Value = val
	return field
}

//LegPriceUnitOfMeasure is a STRING field
type LegPriceUnitOfMeasure struct{ message.StringValue }

//Tag returns tag.LegPriceUnitOfMeasure (1421)
func (f LegPriceUnitOfMeasure) Tag() fix.Tag { return tag.LegPriceUnitOfMeasure }

//BuildLegPriceUnitOfMeasure returns a new LegPriceUnitOfMeasure initialized with val
func BuildLegPriceUnitOfMeasure(val string) LegPriceUnitOfMeasure {
	var field LegPriceUnitOfMeasure
	field.Value = val
	return field
}

//LegPriceUnitOfMeasureQty is a QTY field
type LegPriceUnitOfMeasureQty struct{ message.QtyValue }

//Tag returns tag.LegPriceUnitOfMeasureQty (1422)
func (f LegPriceUnitOfMeasureQty) Tag() fix.Tag { return tag.LegPriceUnitOfMeasureQty }

//BuildLegPriceUnitOfMeasureQty returns a new LegPriceUnitOfMeasureQty initialized with val
func BuildLegPriceUnitOfMeasureQty(val float64) LegPriceUnitOfMeasureQty {
	var field LegPriceUnitOfMeasureQty
	field.Value = val
	return field
}

//LegProduct is a INT field
type LegProduct struct{ message.IntValue }

//Tag returns tag.LegProduct (607)
func (f LegProduct) Tag() fix.Tag { return tag.LegProduct }

//BuildLegProduct returns a new LegProduct initialized with val
func BuildLegProduct(val int) LegProduct {
	var field LegProduct
	field.Value = val
	return field
}

//LegPutOrCall is a INT field
type LegPutOrCall struct{ message.IntValue }

//Tag returns tag.LegPutOrCall (1358)
func (f LegPutOrCall) Tag() fix.Tag { return tag.LegPutOrCall }

//BuildLegPutOrCall returns a new LegPutOrCall initialized with val
func BuildLegPutOrCall(val int) LegPutOrCall {
	var field LegPutOrCall
	field.Value = val
	return field
}

//LegQty is a QTY field
type LegQty struct{ message.QtyValue }

//Tag returns tag.LegQty (687)
func (f LegQty) Tag() fix.Tag { return tag.LegQty }

//BuildLegQty returns a new LegQty initialized with val
func BuildLegQty(val float64) LegQty {
	var field LegQty
	field.Value = val
	return field
}

//LegRatioQty is a FLOAT field
type LegRatioQty struct{ message.FloatValue }

//Tag returns tag.LegRatioQty (623)
func (f LegRatioQty) Tag() fix.Tag { return tag.LegRatioQty }

//BuildLegRatioQty returns a new LegRatioQty initialized with val
func BuildLegRatioQty(val float64) LegRatioQty {
	var field LegRatioQty
	field.Value = val
	return field
}

//LegRedemptionDate is a LOCALMKTDATE field
type LegRedemptionDate struct{ message.LocalMktDateValue }

//Tag returns tag.LegRedemptionDate (254)
func (f LegRedemptionDate) Tag() fix.Tag { return tag.LegRedemptionDate }

//BuildLegRedemptionDate returns a new LegRedemptionDate initialized with val
func BuildLegRedemptionDate(val string) LegRedemptionDate {
	var field LegRedemptionDate
	field.Value = val
	return field
}

//LegRefID is a STRING field
type LegRefID struct{ message.StringValue }

//Tag returns tag.LegRefID (654)
func (f LegRefID) Tag() fix.Tag { return tag.LegRefID }

//BuildLegRefID returns a new LegRefID initialized with val
func BuildLegRefID(val string) LegRefID {
	var field LegRefID
	field.Value = val
	return field
}

//LegRepoCollateralSecurityType is a INT field
type LegRepoCollateralSecurityType struct{ message.IntValue }

//Tag returns tag.LegRepoCollateralSecurityType (250)
func (f LegRepoCollateralSecurityType) Tag() fix.Tag { return tag.LegRepoCollateralSecurityType }

//BuildLegRepoCollateralSecurityType returns a new LegRepoCollateralSecurityType initialized with val
func BuildLegRepoCollateralSecurityType(val int) LegRepoCollateralSecurityType {
	var field LegRepoCollateralSecurityType
	field.Value = val
	return field
}

//LegReportID is a STRING field
type LegReportID struct{ message.StringValue }

//Tag returns tag.LegReportID (990)
func (f LegReportID) Tag() fix.Tag { return tag.LegReportID }

//BuildLegReportID returns a new LegReportID initialized with val
func BuildLegReportID(val string) LegReportID {
	var field LegReportID
	field.Value = val
	return field
}

//LegRepurchaseRate is a PERCENTAGE field
type LegRepurchaseRate struct{ message.PercentageValue }

//Tag returns tag.LegRepurchaseRate (252)
func (f LegRepurchaseRate) Tag() fix.Tag { return tag.LegRepurchaseRate }

//BuildLegRepurchaseRate returns a new LegRepurchaseRate initialized with val
func BuildLegRepurchaseRate(val float64) LegRepurchaseRate {
	var field LegRepurchaseRate
	field.Value = val
	return field
}

//LegRepurchaseTerm is a INT field
type LegRepurchaseTerm struct{ message.IntValue }

//Tag returns tag.LegRepurchaseTerm (251)
func (f LegRepurchaseTerm) Tag() fix.Tag { return tag.LegRepurchaseTerm }

//BuildLegRepurchaseTerm returns a new LegRepurchaseTerm initialized with val
func BuildLegRepurchaseTerm(val int) LegRepurchaseTerm {
	var field LegRepurchaseTerm
	field.Value = val
	return field
}

//LegSecurityAltID is a STRING field
type LegSecurityAltID struct{ message.StringValue }

//Tag returns tag.LegSecurityAltID (605)
func (f LegSecurityAltID) Tag() fix.Tag { return tag.LegSecurityAltID }

//BuildLegSecurityAltID returns a new LegSecurityAltID initialized with val
func BuildLegSecurityAltID(val string) LegSecurityAltID {
	var field LegSecurityAltID
	field.Value = val
	return field
}

//LegSecurityAltIDSource is a STRING field
type LegSecurityAltIDSource struct{ message.StringValue }

//Tag returns tag.LegSecurityAltIDSource (606)
func (f LegSecurityAltIDSource) Tag() fix.Tag { return tag.LegSecurityAltIDSource }

//BuildLegSecurityAltIDSource returns a new LegSecurityAltIDSource initialized with val
func BuildLegSecurityAltIDSource(val string) LegSecurityAltIDSource {
	var field LegSecurityAltIDSource
	field.Value = val
	return field
}

//LegSecurityDesc is a STRING field
type LegSecurityDesc struct{ message.StringValue }

//Tag returns tag.LegSecurityDesc (620)
func (f LegSecurityDesc) Tag() fix.Tag { return tag.LegSecurityDesc }

//BuildLegSecurityDesc returns a new LegSecurityDesc initialized with val
func BuildLegSecurityDesc(val string) LegSecurityDesc {
	var field LegSecurityDesc
	field.Value = val
	return field
}

//LegSecurityExchange is a EXCHANGE field
type LegSecurityExchange struct{ message.ExchangeValue }

//Tag returns tag.LegSecurityExchange (616)
func (f LegSecurityExchange) Tag() fix.Tag { return tag.LegSecurityExchange }

//BuildLegSecurityExchange returns a new LegSecurityExchange initialized with val
func BuildLegSecurityExchange(val string) LegSecurityExchange {
	var field LegSecurityExchange
	field.Value = val
	return field
}

//LegSecurityID is a STRING field
type LegSecurityID struct{ message.StringValue }

//Tag returns tag.LegSecurityID (602)
func (f LegSecurityID) Tag() fix.Tag { return tag.LegSecurityID }

//BuildLegSecurityID returns a new LegSecurityID initialized with val
func BuildLegSecurityID(val string) LegSecurityID {
	var field LegSecurityID
	field.Value = val
	return field
}

//LegSecurityIDSource is a STRING field
type LegSecurityIDSource struct{ message.StringValue }

//Tag returns tag.LegSecurityIDSource (603)
func (f LegSecurityIDSource) Tag() fix.Tag { return tag.LegSecurityIDSource }

//BuildLegSecurityIDSource returns a new LegSecurityIDSource initialized with val
func BuildLegSecurityIDSource(val string) LegSecurityIDSource {
	var field LegSecurityIDSource
	field.Value = val
	return field
}

//LegSecuritySubType is a STRING field
type LegSecuritySubType struct{ message.StringValue }

//Tag returns tag.LegSecuritySubType (764)
func (f LegSecuritySubType) Tag() fix.Tag { return tag.LegSecuritySubType }

//BuildLegSecuritySubType returns a new LegSecuritySubType initialized with val
func BuildLegSecuritySubType(val string) LegSecuritySubType {
	var field LegSecuritySubType
	field.Value = val
	return field
}

//LegSecurityType is a STRING field
type LegSecurityType struct{ message.StringValue }

//Tag returns tag.LegSecurityType (609)
func (f LegSecurityType) Tag() fix.Tag { return tag.LegSecurityType }

//BuildLegSecurityType returns a new LegSecurityType initialized with val
func BuildLegSecurityType(val string) LegSecurityType {
	var field LegSecurityType
	field.Value = val
	return field
}

//LegSettlCurrency is a CURRENCY field
type LegSettlCurrency struct{ message.CurrencyValue }

//Tag returns tag.LegSettlCurrency (675)
func (f LegSettlCurrency) Tag() fix.Tag { return tag.LegSettlCurrency }

//BuildLegSettlCurrency returns a new LegSettlCurrency initialized with val
func BuildLegSettlCurrency(val string) LegSettlCurrency {
	var field LegSettlCurrency
	field.Value = val
	return field
}

//LegSettlDate is a LOCALMKTDATE field
type LegSettlDate struct{ message.LocalMktDateValue }

//Tag returns tag.LegSettlDate (588)
func (f LegSettlDate) Tag() fix.Tag { return tag.LegSettlDate }

//BuildLegSettlDate returns a new LegSettlDate initialized with val
func BuildLegSettlDate(val string) LegSettlDate {
	var field LegSettlDate
	field.Value = val
	return field
}

//LegSettlType is a CHAR field
type LegSettlType struct{ message.CharValue }

//Tag returns tag.LegSettlType (587)
func (f LegSettlType) Tag() fix.Tag { return tag.LegSettlType }

//BuildLegSettlType returns a new LegSettlType initialized with val
func BuildLegSettlType(val string) LegSettlType {
	var field LegSettlType
	field.Value = val
	return field
}

//LegSettlmntTyp is a CHAR field
type LegSettlmntTyp struct{ message.CharValue }

//Tag returns tag.LegSettlmntTyp (587)
func (f LegSettlmntTyp) Tag() fix.Tag { return tag.LegSettlmntTyp }

//BuildLegSettlmntTyp returns a new LegSettlmntTyp initialized with val
func BuildLegSettlmntTyp(val string) LegSettlmntTyp {
	var field LegSettlmntTyp
	field.Value = val
	return field
}

//LegSide is a CHAR field
type LegSide struct{ message.CharValue }

//Tag returns tag.LegSide (624)
func (f LegSide) Tag() fix.Tag { return tag.LegSide }

//BuildLegSide returns a new LegSide initialized with val
func BuildLegSide(val string) LegSide {
	var field LegSide
	field.Value = val
	return field
}

//LegStateOrProvinceOfIssue is a STRING field
type LegStateOrProvinceOfIssue struct{ message.StringValue }

//Tag returns tag.LegStateOrProvinceOfIssue (597)
func (f LegStateOrProvinceOfIssue) Tag() fix.Tag { return tag.LegStateOrProvinceOfIssue }

//BuildLegStateOrProvinceOfIssue returns a new LegStateOrProvinceOfIssue initialized with val
func BuildLegStateOrProvinceOfIssue(val string) LegStateOrProvinceOfIssue {
	var field LegStateOrProvinceOfIssue
	field.Value = val
	return field
}

//LegStipulationType is a STRING field
type LegStipulationType struct{ message.StringValue }

//Tag returns tag.LegStipulationType (688)
func (f LegStipulationType) Tag() fix.Tag { return tag.LegStipulationType }

//BuildLegStipulationType returns a new LegStipulationType initialized with val
func BuildLegStipulationType(val string) LegStipulationType {
	var field LegStipulationType
	field.Value = val
	return field
}

//LegStipulationValue is a STRING field
type LegStipulationValue struct{ message.StringValue }

//Tag returns tag.LegStipulationValue (689)
func (f LegStipulationValue) Tag() fix.Tag { return tag.LegStipulationValue }

//BuildLegStipulationValue returns a new LegStipulationValue initialized with val
func BuildLegStipulationValue(val string) LegStipulationValue {
	var field LegStipulationValue
	field.Value = val
	return field
}

//LegStrikeCurrency is a CURRENCY field
type LegStrikeCurrency struct{ message.CurrencyValue }

//Tag returns tag.LegStrikeCurrency (942)
func (f LegStrikeCurrency) Tag() fix.Tag { return tag.LegStrikeCurrency }

//BuildLegStrikeCurrency returns a new LegStrikeCurrency initialized with val
func BuildLegStrikeCurrency(val string) LegStrikeCurrency {
	var field LegStrikeCurrency
	field.Value = val
	return field
}

//LegStrikePrice is a PRICE field
type LegStrikePrice struct{ message.PriceValue }

//Tag returns tag.LegStrikePrice (612)
func (f LegStrikePrice) Tag() fix.Tag { return tag.LegStrikePrice }

//BuildLegStrikePrice returns a new LegStrikePrice initialized with val
func BuildLegStrikePrice(val float64) LegStrikePrice {
	var field LegStrikePrice
	field.Value = val
	return field
}

//LegSwapType is a INT field
type LegSwapType struct{ message.IntValue }

//Tag returns tag.LegSwapType (690)
func (f LegSwapType) Tag() fix.Tag { return tag.LegSwapType }

//BuildLegSwapType returns a new LegSwapType initialized with val
func BuildLegSwapType(val int) LegSwapType {
	var field LegSwapType
	field.Value = val
	return field
}

//LegSymbol is a STRING field
type LegSymbol struct{ message.StringValue }

//Tag returns tag.LegSymbol (600)
func (f LegSymbol) Tag() fix.Tag { return tag.LegSymbol }

//BuildLegSymbol returns a new LegSymbol initialized with val
func BuildLegSymbol(val string) LegSymbol {
	var field LegSymbol
	field.Value = val
	return field
}

//LegSymbolSfx is a STRING field
type LegSymbolSfx struct{ message.StringValue }

//Tag returns tag.LegSymbolSfx (601)
func (f LegSymbolSfx) Tag() fix.Tag { return tag.LegSymbolSfx }

//BuildLegSymbolSfx returns a new LegSymbolSfx initialized with val
func BuildLegSymbolSfx(val string) LegSymbolSfx {
	var field LegSymbolSfx
	field.Value = val
	return field
}

//LegTimeUnit is a STRING field
type LegTimeUnit struct{ message.StringValue }

//Tag returns tag.LegTimeUnit (1001)
func (f LegTimeUnit) Tag() fix.Tag { return tag.LegTimeUnit }

//BuildLegTimeUnit returns a new LegTimeUnit initialized with val
func BuildLegTimeUnit(val string) LegTimeUnit {
	var field LegTimeUnit
	field.Value = val
	return field
}

//LegUnitOfMeasure is a STRING field
type LegUnitOfMeasure struct{ message.StringValue }

//Tag returns tag.LegUnitOfMeasure (999)
func (f LegUnitOfMeasure) Tag() fix.Tag { return tag.LegUnitOfMeasure }

//BuildLegUnitOfMeasure returns a new LegUnitOfMeasure initialized with val
func BuildLegUnitOfMeasure(val string) LegUnitOfMeasure {
	var field LegUnitOfMeasure
	field.Value = val
	return field
}

//LegUnitOfMeasureQty is a QTY field
type LegUnitOfMeasureQty struct{ message.QtyValue }

//Tag returns tag.LegUnitOfMeasureQty (1224)
func (f LegUnitOfMeasureQty) Tag() fix.Tag { return tag.LegUnitOfMeasureQty }

//BuildLegUnitOfMeasureQty returns a new LegUnitOfMeasureQty initialized with val
func BuildLegUnitOfMeasureQty(val float64) LegUnitOfMeasureQty {
	var field LegUnitOfMeasureQty
	field.Value = val
	return field
}

//LegVolatility is a FLOAT field
type LegVolatility struct{ message.FloatValue }

//Tag returns tag.LegVolatility (1379)
func (f LegVolatility) Tag() fix.Tag { return tag.LegVolatility }

//BuildLegVolatility returns a new LegVolatility initialized with val
func BuildLegVolatility(val float64) LegVolatility {
	var field LegVolatility
	field.Value = val
	return field
}

//LegalConfirm is a BOOLEAN field
type LegalConfirm struct{ message.BooleanValue }

//Tag returns tag.LegalConfirm (650)
func (f LegalConfirm) Tag() fix.Tag { return tag.LegalConfirm }

//BuildLegalConfirm returns a new LegalConfirm initialized with val
func BuildLegalConfirm(val bool) LegalConfirm {
	var field LegalConfirm
	field.Value = val
	return field
}

//LinesOfText is a NUMINGROUP field
type LinesOfText struct{ message.NumInGroupValue }

//Tag returns tag.LinesOfText (33)
func (f LinesOfText) Tag() fix.Tag { return tag.LinesOfText }

//BuildLinesOfText returns a new LinesOfText initialized with val
func BuildLinesOfText(val int) LinesOfText {
	var field LinesOfText
	field.Value = val
	return field
}

//LiquidityIndType is a INT field
type LiquidityIndType struct{ message.IntValue }

//Tag returns tag.LiquidityIndType (409)
func (f LiquidityIndType) Tag() fix.Tag { return tag.LiquidityIndType }

//BuildLiquidityIndType returns a new LiquidityIndType initialized with val
func BuildLiquidityIndType(val int) LiquidityIndType {
	var field LiquidityIndType
	field.Value = val
	return field
}

//LiquidityNumSecurities is a INT field
type LiquidityNumSecurities struct{ message.IntValue }

//Tag returns tag.LiquidityNumSecurities (441)
func (f LiquidityNumSecurities) Tag() fix.Tag { return tag.LiquidityNumSecurities }

//BuildLiquidityNumSecurities returns a new LiquidityNumSecurities initialized with val
func BuildLiquidityNumSecurities(val int) LiquidityNumSecurities {
	var field LiquidityNumSecurities
	field.Value = val
	return field
}

//LiquidityPctHigh is a PERCENTAGE field
type LiquidityPctHigh struct{ message.PercentageValue }

//Tag returns tag.LiquidityPctHigh (403)
func (f LiquidityPctHigh) Tag() fix.Tag { return tag.LiquidityPctHigh }

//BuildLiquidityPctHigh returns a new LiquidityPctHigh initialized with val
func BuildLiquidityPctHigh(val float64) LiquidityPctHigh {
	var field LiquidityPctHigh
	field.Value = val
	return field
}

//LiquidityPctLow is a PERCENTAGE field
type LiquidityPctLow struct{ message.PercentageValue }

//Tag returns tag.LiquidityPctLow (402)
func (f LiquidityPctLow) Tag() fix.Tag { return tag.LiquidityPctLow }

//BuildLiquidityPctLow returns a new LiquidityPctLow initialized with val
func BuildLiquidityPctLow(val float64) LiquidityPctLow {
	var field LiquidityPctLow
	field.Value = val
	return field
}

//LiquidityValue is a AMT field
type LiquidityValue struct{ message.AmtValue }

//Tag returns tag.LiquidityValue (404)
func (f LiquidityValue) Tag() fix.Tag { return tag.LiquidityValue }

//BuildLiquidityValue returns a new LiquidityValue initialized with val
func BuildLiquidityValue(val float64) LiquidityValue {
	var field LiquidityValue
	field.Value = val
	return field
}

//ListExecInst is a STRING field
type ListExecInst struct{ message.StringValue }

//Tag returns tag.ListExecInst (69)
func (f ListExecInst) Tag() fix.Tag { return tag.ListExecInst }

//BuildListExecInst returns a new ListExecInst initialized with val
func BuildListExecInst(val string) ListExecInst {
	var field ListExecInst
	field.Value = val
	return field
}

//ListExecInstType is a CHAR field
type ListExecInstType struct{ message.CharValue }

//Tag returns tag.ListExecInstType (433)
func (f ListExecInstType) Tag() fix.Tag { return tag.ListExecInstType }

//BuildListExecInstType returns a new ListExecInstType initialized with val
func BuildListExecInstType(val string) ListExecInstType {
	var field ListExecInstType
	field.Value = val
	return field
}

//ListID is a STRING field
type ListID struct{ message.StringValue }

//Tag returns tag.ListID (66)
func (f ListID) Tag() fix.Tag { return tag.ListID }

//BuildListID returns a new ListID initialized with val
func BuildListID(val string) ListID {
	var field ListID
	field.Value = val
	return field
}

//ListMethod is a INT field
type ListMethod struct{ message.IntValue }

//Tag returns tag.ListMethod (1198)
func (f ListMethod) Tag() fix.Tag { return tag.ListMethod }

//BuildListMethod returns a new ListMethod initialized with val
func BuildListMethod(val int) ListMethod {
	var field ListMethod
	field.Value = val
	return field
}

//ListName is a STRING field
type ListName struct{ message.StringValue }

//Tag returns tag.ListName (392)
func (f ListName) Tag() fix.Tag { return tag.ListName }

//BuildListName returns a new ListName initialized with val
func BuildListName(val string) ListName {
	var field ListName
	field.Value = val
	return field
}

//ListNoOrds is a INT field
type ListNoOrds struct{ message.IntValue }

//Tag returns tag.ListNoOrds (68)
func (f ListNoOrds) Tag() fix.Tag { return tag.ListNoOrds }

//BuildListNoOrds returns a new ListNoOrds initialized with val
func BuildListNoOrds(val int) ListNoOrds {
	var field ListNoOrds
	field.Value = val
	return field
}

//ListOrderStatus is a INT field
type ListOrderStatus struct{ message.IntValue }

//Tag returns tag.ListOrderStatus (431)
func (f ListOrderStatus) Tag() fix.Tag { return tag.ListOrderStatus }

//BuildListOrderStatus returns a new ListOrderStatus initialized with val
func BuildListOrderStatus(val int) ListOrderStatus {
	var field ListOrderStatus
	field.Value = val
	return field
}

//ListRejectReason is a INT field
type ListRejectReason struct{ message.IntValue }

//Tag returns tag.ListRejectReason (1386)
func (f ListRejectReason) Tag() fix.Tag { return tag.ListRejectReason }

//BuildListRejectReason returns a new ListRejectReason initialized with val
func BuildListRejectReason(val int) ListRejectReason {
	var field ListRejectReason
	field.Value = val
	return field
}

//ListSeqNo is a INT field
type ListSeqNo struct{ message.IntValue }

//Tag returns tag.ListSeqNo (67)
func (f ListSeqNo) Tag() fix.Tag { return tag.ListSeqNo }

//BuildListSeqNo returns a new ListSeqNo initialized with val
func BuildListSeqNo(val int) ListSeqNo {
	var field ListSeqNo
	field.Value = val
	return field
}

//ListStatusText is a STRING field
type ListStatusText struct{ message.StringValue }

//Tag returns tag.ListStatusText (444)
func (f ListStatusText) Tag() fix.Tag { return tag.ListStatusText }

//BuildListStatusText returns a new ListStatusText initialized with val
func BuildListStatusText(val string) ListStatusText {
	var field ListStatusText
	field.Value = val
	return field
}

//ListStatusType is a INT field
type ListStatusType struct{ message.IntValue }

//Tag returns tag.ListStatusType (429)
func (f ListStatusType) Tag() fix.Tag { return tag.ListStatusType }

//BuildListStatusType returns a new ListStatusType initialized with val
func BuildListStatusType(val int) ListStatusType {
	var field ListStatusType
	field.Value = val
	return field
}

//ListUpdateAction is a CHAR field
type ListUpdateAction struct{ message.CharValue }

//Tag returns tag.ListUpdateAction (1324)
func (f ListUpdateAction) Tag() fix.Tag { return tag.ListUpdateAction }

//BuildListUpdateAction returns a new ListUpdateAction initialized with val
func BuildListUpdateAction(val string) ListUpdateAction {
	var field ListUpdateAction
	field.Value = val
	return field
}

//LocaleOfIssue is a STRING field
type LocaleOfIssue struct{ message.StringValue }

//Tag returns tag.LocaleOfIssue (472)
func (f LocaleOfIssue) Tag() fix.Tag { return tag.LocaleOfIssue }

//BuildLocaleOfIssue returns a new LocaleOfIssue initialized with val
func BuildLocaleOfIssue(val string) LocaleOfIssue {
	var field LocaleOfIssue
	field.Value = val
	return field
}

//LocateReqd is a BOOLEAN field
type LocateReqd struct{ message.BooleanValue }

//Tag returns tag.LocateReqd (114)
func (f LocateReqd) Tag() fix.Tag { return tag.LocateReqd }

//BuildLocateReqd returns a new LocateReqd initialized with val
func BuildLocateReqd(val bool) LocateReqd {
	var field LocateReqd
	field.Value = val
	return field
}

//LocationID is a STRING field
type LocationID struct{ message.StringValue }

//Tag returns tag.LocationID (283)
func (f LocationID) Tag() fix.Tag { return tag.LocationID }

//BuildLocationID returns a new LocationID initialized with val
func BuildLocationID(val string) LocationID {
	var field LocationID
	field.Value = val
	return field
}

//LongQty is a QTY field
type LongQty struct{ message.QtyValue }

//Tag returns tag.LongQty (704)
func (f LongQty) Tag() fix.Tag { return tag.LongQty }

//BuildLongQty returns a new LongQty initialized with val
func BuildLongQty(val float64) LongQty {
	var field LongQty
	field.Value = val
	return field
}

//LotType is a CHAR field
type LotType struct{ message.CharValue }

//Tag returns tag.LotType (1093)
func (f LotType) Tag() fix.Tag { return tag.LotType }

//BuildLotType returns a new LotType initialized with val
func BuildLotType(val string) LotType {
	var field LotType
	field.Value = val
	return field
}

//LowLimitPrice is a PRICE field
type LowLimitPrice struct{ message.PriceValue }

//Tag returns tag.LowLimitPrice (1148)
func (f LowLimitPrice) Tag() fix.Tag { return tag.LowLimitPrice }

//BuildLowLimitPrice returns a new LowLimitPrice initialized with val
func BuildLowLimitPrice(val float64) LowLimitPrice {
	var field LowLimitPrice
	field.Value = val
	return field
}

//LowPx is a PRICE field
type LowPx struct{ message.PriceValue }

//Tag returns tag.LowPx (333)
func (f LowPx) Tag() fix.Tag { return tag.LowPx }

//BuildLowPx returns a new LowPx initialized with val
func BuildLowPx(val float64) LowPx {
	var field LowPx
	field.Value = val
	return field
}

//MDBookType is a INT field
type MDBookType struct{ message.IntValue }

//Tag returns tag.MDBookType (1021)
func (f MDBookType) Tag() fix.Tag { return tag.MDBookType }

//BuildMDBookType returns a new MDBookType initialized with val
func BuildMDBookType(val int) MDBookType {
	var field MDBookType
	field.Value = val
	return field
}

//MDEntryBuyer is a STRING field
type MDEntryBuyer struct{ message.StringValue }

//Tag returns tag.MDEntryBuyer (288)
func (f MDEntryBuyer) Tag() fix.Tag { return tag.MDEntryBuyer }

//BuildMDEntryBuyer returns a new MDEntryBuyer initialized with val
func BuildMDEntryBuyer(val string) MDEntryBuyer {
	var field MDEntryBuyer
	field.Value = val
	return field
}

//MDEntryDate is a UTCDATEONLY field
type MDEntryDate struct{ message.UTCDateOnlyValue }

//Tag returns tag.MDEntryDate (272)
func (f MDEntryDate) Tag() fix.Tag { return tag.MDEntryDate }

//MDEntryForwardPoints is a PRICEOFFSET field
type MDEntryForwardPoints struct{ message.PriceOffsetValue }

//Tag returns tag.MDEntryForwardPoints (1027)
func (f MDEntryForwardPoints) Tag() fix.Tag { return tag.MDEntryForwardPoints }

//BuildMDEntryForwardPoints returns a new MDEntryForwardPoints initialized with val
func BuildMDEntryForwardPoints(val float64) MDEntryForwardPoints {
	var field MDEntryForwardPoints
	field.Value = val
	return field
}

//MDEntryID is a STRING field
type MDEntryID struct{ message.StringValue }

//Tag returns tag.MDEntryID (278)
func (f MDEntryID) Tag() fix.Tag { return tag.MDEntryID }

//BuildMDEntryID returns a new MDEntryID initialized with val
func BuildMDEntryID(val string) MDEntryID {
	var field MDEntryID
	field.Value = val
	return field
}

//MDEntryOriginator is a STRING field
type MDEntryOriginator struct{ message.StringValue }

//Tag returns tag.MDEntryOriginator (282)
func (f MDEntryOriginator) Tag() fix.Tag { return tag.MDEntryOriginator }

//BuildMDEntryOriginator returns a new MDEntryOriginator initialized with val
func BuildMDEntryOriginator(val string) MDEntryOriginator {
	var field MDEntryOriginator
	field.Value = val
	return field
}

//MDEntryPositionNo is a INT field
type MDEntryPositionNo struct{ message.IntValue }

//Tag returns tag.MDEntryPositionNo (290)
func (f MDEntryPositionNo) Tag() fix.Tag { return tag.MDEntryPositionNo }

//BuildMDEntryPositionNo returns a new MDEntryPositionNo initialized with val
func BuildMDEntryPositionNo(val int) MDEntryPositionNo {
	var field MDEntryPositionNo
	field.Value = val
	return field
}

//MDEntryPx is a PRICE field
type MDEntryPx struct{ message.PriceValue }

//Tag returns tag.MDEntryPx (270)
func (f MDEntryPx) Tag() fix.Tag { return tag.MDEntryPx }

//BuildMDEntryPx returns a new MDEntryPx initialized with val
func BuildMDEntryPx(val float64) MDEntryPx {
	var field MDEntryPx
	field.Value = val
	return field
}

//MDEntryRefID is a STRING field
type MDEntryRefID struct{ message.StringValue }

//Tag returns tag.MDEntryRefID (280)
func (f MDEntryRefID) Tag() fix.Tag { return tag.MDEntryRefID }

//BuildMDEntryRefID returns a new MDEntryRefID initialized with val
func BuildMDEntryRefID(val string) MDEntryRefID {
	var field MDEntryRefID
	field.Value = val
	return field
}

//MDEntrySeller is a STRING field
type MDEntrySeller struct{ message.StringValue }

//Tag returns tag.MDEntrySeller (289)
func (f MDEntrySeller) Tag() fix.Tag { return tag.MDEntrySeller }

//BuildMDEntrySeller returns a new MDEntrySeller initialized with val
func BuildMDEntrySeller(val string) MDEntrySeller {
	var field MDEntrySeller
	field.Value = val
	return field
}

//MDEntrySize is a QTY field
type MDEntrySize struct{ message.QtyValue }

//Tag returns tag.MDEntrySize (271)
func (f MDEntrySize) Tag() fix.Tag { return tag.MDEntrySize }

//BuildMDEntrySize returns a new MDEntrySize initialized with val
func BuildMDEntrySize(val float64) MDEntrySize {
	var field MDEntrySize
	field.Value = val
	return field
}

//MDEntrySpotRate is a FLOAT field
type MDEntrySpotRate struct{ message.FloatValue }

//Tag returns tag.MDEntrySpotRate (1026)
func (f MDEntrySpotRate) Tag() fix.Tag { return tag.MDEntrySpotRate }

//BuildMDEntrySpotRate returns a new MDEntrySpotRate initialized with val
func BuildMDEntrySpotRate(val float64) MDEntrySpotRate {
	var field MDEntrySpotRate
	field.Value = val
	return field
}

//MDEntryTime is a UTCTIMEONLY field
type MDEntryTime struct{ message.UTCTimeOnlyValue }

//Tag returns tag.MDEntryTime (273)
func (f MDEntryTime) Tag() fix.Tag { return tag.MDEntryTime }

//MDEntryType is a CHAR field
type MDEntryType struct{ message.CharValue }

//Tag returns tag.MDEntryType (269)
func (f MDEntryType) Tag() fix.Tag { return tag.MDEntryType }

//BuildMDEntryType returns a new MDEntryType initialized with val
func BuildMDEntryType(val string) MDEntryType {
	var field MDEntryType
	field.Value = val
	return field
}

//MDFeedType is a STRING field
type MDFeedType struct{ message.StringValue }

//Tag returns tag.MDFeedType (1022)
func (f MDFeedType) Tag() fix.Tag { return tag.MDFeedType }

//BuildMDFeedType returns a new MDFeedType initialized with val
func BuildMDFeedType(val string) MDFeedType {
	var field MDFeedType
	field.Value = val
	return field
}

//MDImplicitDelete is a BOOLEAN field
type MDImplicitDelete struct{ message.BooleanValue }

//Tag returns tag.MDImplicitDelete (547)
func (f MDImplicitDelete) Tag() fix.Tag { return tag.MDImplicitDelete }

//BuildMDImplicitDelete returns a new MDImplicitDelete initialized with val
func BuildMDImplicitDelete(val bool) MDImplicitDelete {
	var field MDImplicitDelete
	field.Value = val
	return field
}

//MDMkt is a EXCHANGE field
type MDMkt struct{ message.ExchangeValue }

//Tag returns tag.MDMkt (275)
func (f MDMkt) Tag() fix.Tag { return tag.MDMkt }

//BuildMDMkt returns a new MDMkt initialized with val
func BuildMDMkt(val string) MDMkt {
	var field MDMkt
	field.Value = val
	return field
}

//MDOriginType is a INT field
type MDOriginType struct{ message.IntValue }

//Tag returns tag.MDOriginType (1024)
func (f MDOriginType) Tag() fix.Tag { return tag.MDOriginType }

//BuildMDOriginType returns a new MDOriginType initialized with val
func BuildMDOriginType(val int) MDOriginType {
	var field MDOriginType
	field.Value = val
	return field
}

//MDPriceLevel is a INT field
type MDPriceLevel struct{ message.IntValue }

//Tag returns tag.MDPriceLevel (1023)
func (f MDPriceLevel) Tag() fix.Tag { return tag.MDPriceLevel }

//BuildMDPriceLevel returns a new MDPriceLevel initialized with val
func BuildMDPriceLevel(val int) MDPriceLevel {
	var field MDPriceLevel
	field.Value = val
	return field
}

//MDQuoteType is a INT field
type MDQuoteType struct{ message.IntValue }

//Tag returns tag.MDQuoteType (1070)
func (f MDQuoteType) Tag() fix.Tag { return tag.MDQuoteType }

//BuildMDQuoteType returns a new MDQuoteType initialized with val
func BuildMDQuoteType(val int) MDQuoteType {
	var field MDQuoteType
	field.Value = val
	return field
}

//MDReportID is a INT field
type MDReportID struct{ message.IntValue }

//Tag returns tag.MDReportID (963)
func (f MDReportID) Tag() fix.Tag { return tag.MDReportID }

//BuildMDReportID returns a new MDReportID initialized with val
func BuildMDReportID(val int) MDReportID {
	var field MDReportID
	field.Value = val
	return field
}

//MDReqID is a STRING field
type MDReqID struct{ message.StringValue }

//Tag returns tag.MDReqID (262)
func (f MDReqID) Tag() fix.Tag { return tag.MDReqID }

//BuildMDReqID returns a new MDReqID initialized with val
func BuildMDReqID(val string) MDReqID {
	var field MDReqID
	field.Value = val
	return field
}

//MDReqRejReason is a CHAR field
type MDReqRejReason struct{ message.CharValue }

//Tag returns tag.MDReqRejReason (281)
func (f MDReqRejReason) Tag() fix.Tag { return tag.MDReqRejReason }

//BuildMDReqRejReason returns a new MDReqRejReason initialized with val
func BuildMDReqRejReason(val string) MDReqRejReason {
	var field MDReqRejReason
	field.Value = val
	return field
}

//MDSecSize is a QTY field
type MDSecSize struct{ message.QtyValue }

//Tag returns tag.MDSecSize (1179)
func (f MDSecSize) Tag() fix.Tag { return tag.MDSecSize }

//BuildMDSecSize returns a new MDSecSize initialized with val
func BuildMDSecSize(val float64) MDSecSize {
	var field MDSecSize
	field.Value = val
	return field
}

//MDSecSizeType is a INT field
type MDSecSizeType struct{ message.IntValue }

//Tag returns tag.MDSecSizeType (1178)
func (f MDSecSizeType) Tag() fix.Tag { return tag.MDSecSizeType }

//BuildMDSecSizeType returns a new MDSecSizeType initialized with val
func BuildMDSecSizeType(val int) MDSecSizeType {
	var field MDSecSizeType
	field.Value = val
	return field
}

//MDStreamID is a STRING field
type MDStreamID struct{ message.StringValue }

//Tag returns tag.MDStreamID (1500)
func (f MDStreamID) Tag() fix.Tag { return tag.MDStreamID }

//BuildMDStreamID returns a new MDStreamID initialized with val
func BuildMDStreamID(val string) MDStreamID {
	var field MDStreamID
	field.Value = val
	return field
}

//MDSubBookType is a INT field
type MDSubBookType struct{ message.IntValue }

//Tag returns tag.MDSubBookType (1173)
func (f MDSubBookType) Tag() fix.Tag { return tag.MDSubBookType }

//BuildMDSubBookType returns a new MDSubBookType initialized with val
func BuildMDSubBookType(val int) MDSubBookType {
	var field MDSubBookType
	field.Value = val
	return field
}

//MDUpdateAction is a CHAR field
type MDUpdateAction struct{ message.CharValue }

//Tag returns tag.MDUpdateAction (279)
func (f MDUpdateAction) Tag() fix.Tag { return tag.MDUpdateAction }

//BuildMDUpdateAction returns a new MDUpdateAction initialized with val
func BuildMDUpdateAction(val string) MDUpdateAction {
	var field MDUpdateAction
	field.Value = val
	return field
}

//MDUpdateType is a INT field
type MDUpdateType struct{ message.IntValue }

//Tag returns tag.MDUpdateType (265)
func (f MDUpdateType) Tag() fix.Tag { return tag.MDUpdateType }

//BuildMDUpdateType returns a new MDUpdateType initialized with val
func BuildMDUpdateType(val int) MDUpdateType {
	var field MDUpdateType
	field.Value = val
	return field
}

//MailingDtls is a STRING field
type MailingDtls struct{ message.StringValue }

//Tag returns tag.MailingDtls (474)
func (f MailingDtls) Tag() fix.Tag { return tag.MailingDtls }

//BuildMailingDtls returns a new MailingDtls initialized with val
func BuildMailingDtls(val string) MailingDtls {
	var field MailingDtls
	field.Value = val
	return field
}

//MailingInst is a STRING field
type MailingInst struct{ message.StringValue }

//Tag returns tag.MailingInst (482)
func (f MailingInst) Tag() fix.Tag { return tag.MailingInst }

//BuildMailingInst returns a new MailingInst initialized with val
func BuildMailingInst(val string) MailingInst {
	var field MailingInst
	field.Value = val
	return field
}

//ManualOrderIndicator is a BOOLEAN field
type ManualOrderIndicator struct{ message.BooleanValue }

//Tag returns tag.ManualOrderIndicator (1028)
func (f ManualOrderIndicator) Tag() fix.Tag { return tag.ManualOrderIndicator }

//BuildManualOrderIndicator returns a new ManualOrderIndicator initialized with val
func BuildManualOrderIndicator(val bool) ManualOrderIndicator {
	var field ManualOrderIndicator
	field.Value = val
	return field
}

//MarginExcess is a AMT field
type MarginExcess struct{ message.AmtValue }

//Tag returns tag.MarginExcess (899)
func (f MarginExcess) Tag() fix.Tag { return tag.MarginExcess }

//BuildMarginExcess returns a new MarginExcess initialized with val
func BuildMarginExcess(val float64) MarginExcess {
	var field MarginExcess
	field.Value = val
	return field
}

//MarginRatio is a PERCENTAGE field
type MarginRatio struct{ message.PercentageValue }

//Tag returns tag.MarginRatio (898)
func (f MarginRatio) Tag() fix.Tag { return tag.MarginRatio }

//BuildMarginRatio returns a new MarginRatio initialized with val
func BuildMarginRatio(val float64) MarginRatio {
	var field MarginRatio
	field.Value = val
	return field
}

//MarketDepth is a INT field
type MarketDepth struct{ message.IntValue }

//Tag returns tag.MarketDepth (264)
func (f MarketDepth) Tag() fix.Tag { return tag.MarketDepth }

//BuildMarketDepth returns a new MarketDepth initialized with val
func BuildMarketDepth(val int) MarketDepth {
	var field MarketDepth
	field.Value = val
	return field
}

//MarketID is a EXCHANGE field
type MarketID struct{ message.ExchangeValue }

//Tag returns tag.MarketID (1301)
func (f MarketID) Tag() fix.Tag { return tag.MarketID }

//BuildMarketID returns a new MarketID initialized with val
func BuildMarketID(val string) MarketID {
	var field MarketID
	field.Value = val
	return field
}

//MarketReportID is a STRING field
type MarketReportID struct{ message.StringValue }

//Tag returns tag.MarketReportID (1394)
func (f MarketReportID) Tag() fix.Tag { return tag.MarketReportID }

//BuildMarketReportID returns a new MarketReportID initialized with val
func BuildMarketReportID(val string) MarketReportID {
	var field MarketReportID
	field.Value = val
	return field
}

//MarketReqID is a STRING field
type MarketReqID struct{ message.StringValue }

//Tag returns tag.MarketReqID (1393)
func (f MarketReqID) Tag() fix.Tag { return tag.MarketReqID }

//BuildMarketReqID returns a new MarketReqID initialized with val
func BuildMarketReqID(val string) MarketReqID {
	var field MarketReqID
	field.Value = val
	return field
}

//MarketSegmentDesc is a STRING field
type MarketSegmentDesc struct{ message.StringValue }

//Tag returns tag.MarketSegmentDesc (1396)
func (f MarketSegmentDesc) Tag() fix.Tag { return tag.MarketSegmentDesc }

//BuildMarketSegmentDesc returns a new MarketSegmentDesc initialized with val
func BuildMarketSegmentDesc(val string) MarketSegmentDesc {
	var field MarketSegmentDesc
	field.Value = val
	return field
}

//MarketSegmentID is a STRING field
type MarketSegmentID struct{ message.StringValue }

//Tag returns tag.MarketSegmentID (1300)
func (f MarketSegmentID) Tag() fix.Tag { return tag.MarketSegmentID }

//BuildMarketSegmentID returns a new MarketSegmentID initialized with val
func BuildMarketSegmentID(val string) MarketSegmentID {
	var field MarketSegmentID
	field.Value = val
	return field
}

//MarketUpdateAction is a CHAR field
type MarketUpdateAction struct{ message.CharValue }

//Tag returns tag.MarketUpdateAction (1395)
func (f MarketUpdateAction) Tag() fix.Tag { return tag.MarketUpdateAction }

//BuildMarketUpdateAction returns a new MarketUpdateAction initialized with val
func BuildMarketUpdateAction(val string) MarketUpdateAction {
	var field MarketUpdateAction
	field.Value = val
	return field
}

//MassActionRejectReason is a INT field
type MassActionRejectReason struct{ message.IntValue }

//Tag returns tag.MassActionRejectReason (1376)
func (f MassActionRejectReason) Tag() fix.Tag { return tag.MassActionRejectReason }

//BuildMassActionRejectReason returns a new MassActionRejectReason initialized with val
func BuildMassActionRejectReason(val int) MassActionRejectReason {
	var field MassActionRejectReason
	field.Value = val
	return field
}

//MassActionReportID is a STRING field
type MassActionReportID struct{ message.StringValue }

//Tag returns tag.MassActionReportID (1369)
func (f MassActionReportID) Tag() fix.Tag { return tag.MassActionReportID }

//BuildMassActionReportID returns a new MassActionReportID initialized with val
func BuildMassActionReportID(val string) MassActionReportID {
	var field MassActionReportID
	field.Value = val
	return field
}

//MassActionResponse is a INT field
type MassActionResponse struct{ message.IntValue }

//Tag returns tag.MassActionResponse (1375)
func (f MassActionResponse) Tag() fix.Tag { return tag.MassActionResponse }

//BuildMassActionResponse returns a new MassActionResponse initialized with val
func BuildMassActionResponse(val int) MassActionResponse {
	var field MassActionResponse
	field.Value = val
	return field
}

//MassActionScope is a INT field
type MassActionScope struct{ message.IntValue }

//Tag returns tag.MassActionScope (1374)
func (f MassActionScope) Tag() fix.Tag { return tag.MassActionScope }

//BuildMassActionScope returns a new MassActionScope initialized with val
func BuildMassActionScope(val int) MassActionScope {
	var field MassActionScope
	field.Value = val
	return field
}

//MassActionType is a INT field
type MassActionType struct{ message.IntValue }

//Tag returns tag.MassActionType (1373)
func (f MassActionType) Tag() fix.Tag { return tag.MassActionType }

//BuildMassActionType returns a new MassActionType initialized with val
func BuildMassActionType(val int) MassActionType {
	var field MassActionType
	field.Value = val
	return field
}

//MassCancelRejectReason is a INT field
type MassCancelRejectReason struct{ message.IntValue }

//Tag returns tag.MassCancelRejectReason (532)
func (f MassCancelRejectReason) Tag() fix.Tag { return tag.MassCancelRejectReason }

//BuildMassCancelRejectReason returns a new MassCancelRejectReason initialized with val
func BuildMassCancelRejectReason(val int) MassCancelRejectReason {
	var field MassCancelRejectReason
	field.Value = val
	return field
}

//MassCancelRequestType is a CHAR field
type MassCancelRequestType struct{ message.CharValue }

//Tag returns tag.MassCancelRequestType (530)
func (f MassCancelRequestType) Tag() fix.Tag { return tag.MassCancelRequestType }

//BuildMassCancelRequestType returns a new MassCancelRequestType initialized with val
func BuildMassCancelRequestType(val string) MassCancelRequestType {
	var field MassCancelRequestType
	field.Value = val
	return field
}

//MassCancelResponse is a CHAR field
type MassCancelResponse struct{ message.CharValue }

//Tag returns tag.MassCancelResponse (531)
func (f MassCancelResponse) Tag() fix.Tag { return tag.MassCancelResponse }

//BuildMassCancelResponse returns a new MassCancelResponse initialized with val
func BuildMassCancelResponse(val string) MassCancelResponse {
	var field MassCancelResponse
	field.Value = val
	return field
}

//MassStatusReqID is a STRING field
type MassStatusReqID struct{ message.StringValue }

//Tag returns tag.MassStatusReqID (584)
func (f MassStatusReqID) Tag() fix.Tag { return tag.MassStatusReqID }

//BuildMassStatusReqID returns a new MassStatusReqID initialized with val
func BuildMassStatusReqID(val string) MassStatusReqID {
	var field MassStatusReqID
	field.Value = val
	return field
}

//MassStatusReqType is a INT field
type MassStatusReqType struct{ message.IntValue }

//Tag returns tag.MassStatusReqType (585)
func (f MassStatusReqType) Tag() fix.Tag { return tag.MassStatusReqType }

//BuildMassStatusReqType returns a new MassStatusReqType initialized with val
func BuildMassStatusReqType(val int) MassStatusReqType {
	var field MassStatusReqType
	field.Value = val
	return field
}

//MatchAlgorithm is a STRING field
type MatchAlgorithm struct{ message.StringValue }

//Tag returns tag.MatchAlgorithm (1142)
func (f MatchAlgorithm) Tag() fix.Tag { return tag.MatchAlgorithm }

//BuildMatchAlgorithm returns a new MatchAlgorithm initialized with val
func BuildMatchAlgorithm(val string) MatchAlgorithm {
	var field MatchAlgorithm
	field.Value = val
	return field
}

//MatchIncrement is a QTY field
type MatchIncrement struct{ message.QtyValue }

//Tag returns tag.MatchIncrement (1089)
func (f MatchIncrement) Tag() fix.Tag { return tag.MatchIncrement }

//BuildMatchIncrement returns a new MatchIncrement initialized with val
func BuildMatchIncrement(val float64) MatchIncrement {
	var field MatchIncrement
	field.Value = val
	return field
}

//MatchStatus is a CHAR field
type MatchStatus struct{ message.CharValue }

//Tag returns tag.MatchStatus (573)
func (f MatchStatus) Tag() fix.Tag { return tag.MatchStatus }

//BuildMatchStatus returns a new MatchStatus initialized with val
func BuildMatchStatus(val string) MatchStatus {
	var field MatchStatus
	field.Value = val
	return field
}

//MatchType is a STRING field
type MatchType struct{ message.StringValue }

//Tag returns tag.MatchType (574)
func (f MatchType) Tag() fix.Tag { return tag.MatchType }

//BuildMatchType returns a new MatchType initialized with val
func BuildMatchType(val string) MatchType {
	var field MatchType
	field.Value = val
	return field
}

//MaturityDate is a LOCALMKTDATE field
type MaturityDate struct{ message.LocalMktDateValue }

//Tag returns tag.MaturityDate (541)
func (f MaturityDate) Tag() fix.Tag { return tag.MaturityDate }

//BuildMaturityDate returns a new MaturityDate initialized with val
func BuildMaturityDate(val string) MaturityDate {
	var field MaturityDate
	field.Value = val
	return field
}

//MaturityDay is a DAYOFMONTH field
type MaturityDay struct{ message.DayOfMonthValue }

//Tag returns tag.MaturityDay (205)
func (f MaturityDay) Tag() fix.Tag { return tag.MaturityDay }

//BuildMaturityDay returns a new MaturityDay initialized with val
func BuildMaturityDay(val int) MaturityDay {
	var field MaturityDay
	field.Value = val
	return field
}

//MaturityMonthYear is a MONTHYEAR field
type MaturityMonthYear struct{ message.MonthYearValue }

//Tag returns tag.MaturityMonthYear (200)
func (f MaturityMonthYear) Tag() fix.Tag { return tag.MaturityMonthYear }

//BuildMaturityMonthYear returns a new MaturityMonthYear initialized with val
func BuildMaturityMonthYear(val string) MaturityMonthYear {
	var field MaturityMonthYear
	field.Value = val
	return field
}

//MaturityMonthYearFormat is a INT field
type MaturityMonthYearFormat struct{ message.IntValue }

//Tag returns tag.MaturityMonthYearFormat (1303)
func (f MaturityMonthYearFormat) Tag() fix.Tag { return tag.MaturityMonthYearFormat }

//BuildMaturityMonthYearFormat returns a new MaturityMonthYearFormat initialized with val
func BuildMaturityMonthYearFormat(val int) MaturityMonthYearFormat {
	var field MaturityMonthYearFormat
	field.Value = val
	return field
}

//MaturityMonthYearIncrement is a INT field
type MaturityMonthYearIncrement struct{ message.IntValue }

//Tag returns tag.MaturityMonthYearIncrement (1229)
func (f MaturityMonthYearIncrement) Tag() fix.Tag { return tag.MaturityMonthYearIncrement }

//BuildMaturityMonthYearIncrement returns a new MaturityMonthYearIncrement initialized with val
func BuildMaturityMonthYearIncrement(val int) MaturityMonthYearIncrement {
	var field MaturityMonthYearIncrement
	field.Value = val
	return field
}

//MaturityMonthYearIncrementUnits is a INT field
type MaturityMonthYearIncrementUnits struct{ message.IntValue }

//Tag returns tag.MaturityMonthYearIncrementUnits (1302)
func (f MaturityMonthYearIncrementUnits) Tag() fix.Tag { return tag.MaturityMonthYearIncrementUnits }

//BuildMaturityMonthYearIncrementUnits returns a new MaturityMonthYearIncrementUnits initialized with val
func BuildMaturityMonthYearIncrementUnits(val int) MaturityMonthYearIncrementUnits {
	var field MaturityMonthYearIncrementUnits
	field.Value = val
	return field
}

//MaturityNetMoney is a AMT field
type MaturityNetMoney struct{ message.AmtValue }

//Tag returns tag.MaturityNetMoney (890)
func (f MaturityNetMoney) Tag() fix.Tag { return tag.MaturityNetMoney }

//BuildMaturityNetMoney returns a new MaturityNetMoney initialized with val
func BuildMaturityNetMoney(val float64) MaturityNetMoney {
	var field MaturityNetMoney
	field.Value = val
	return field
}

//MaturityRuleID is a STRING field
type MaturityRuleID struct{ message.StringValue }

//Tag returns tag.MaturityRuleID (1222)
func (f MaturityRuleID) Tag() fix.Tag { return tag.MaturityRuleID }

//BuildMaturityRuleID returns a new MaturityRuleID initialized with val
func BuildMaturityRuleID(val string) MaturityRuleID {
	var field MaturityRuleID
	field.Value = val
	return field
}

//MaturityTime is a TZTIMEONLY field
type MaturityTime struct{ message.TZTimeOnlyValue }

//Tag returns tag.MaturityTime (1079)
func (f MaturityTime) Tag() fix.Tag { return tag.MaturityTime }

//MaxFloor is a QTY field
type MaxFloor struct{ message.QtyValue }

//Tag returns tag.MaxFloor (111)
func (f MaxFloor) Tag() fix.Tag { return tag.MaxFloor }

//BuildMaxFloor returns a new MaxFloor initialized with val
func BuildMaxFloor(val float64) MaxFloor {
	var field MaxFloor
	field.Value = val
	return field
}

//MaxMessageSize is a LENGTH field
type MaxMessageSize struct{ message.LengthValue }

//Tag returns tag.MaxMessageSize (383)
func (f MaxMessageSize) Tag() fix.Tag { return tag.MaxMessageSize }

//BuildMaxMessageSize returns a new MaxMessageSize initialized with val
func BuildMaxMessageSize(val int) MaxMessageSize {
	var field MaxMessageSize
	field.Value = val
	return field
}

//MaxPriceLevels is a INT field
type MaxPriceLevels struct{ message.IntValue }

//Tag returns tag.MaxPriceLevels (1090)
func (f MaxPriceLevels) Tag() fix.Tag { return tag.MaxPriceLevels }

//BuildMaxPriceLevels returns a new MaxPriceLevels initialized with val
func BuildMaxPriceLevels(val int) MaxPriceLevels {
	var field MaxPriceLevels
	field.Value = val
	return field
}

//MaxPriceVariation is a FLOAT field
type MaxPriceVariation struct{ message.FloatValue }

//Tag returns tag.MaxPriceVariation (1143)
func (f MaxPriceVariation) Tag() fix.Tag { return tag.MaxPriceVariation }

//BuildMaxPriceVariation returns a new MaxPriceVariation initialized with val
func BuildMaxPriceVariation(val float64) MaxPriceVariation {
	var field MaxPriceVariation
	field.Value = val
	return field
}

//MaxShow is a QTY field
type MaxShow struct{ message.QtyValue }

//Tag returns tag.MaxShow (210)
func (f MaxShow) Tag() fix.Tag { return tag.MaxShow }

//BuildMaxShow returns a new MaxShow initialized with val
func BuildMaxShow(val float64) MaxShow {
	var field MaxShow
	field.Value = val
	return field
}

//MaxTradeVol is a QTY field
type MaxTradeVol struct{ message.QtyValue }

//Tag returns tag.MaxTradeVol (1140)
func (f MaxTradeVol) Tag() fix.Tag { return tag.MaxTradeVol }

//BuildMaxTradeVol returns a new MaxTradeVol initialized with val
func BuildMaxTradeVol(val float64) MaxTradeVol {
	var field MaxTradeVol
	field.Value = val
	return field
}

//MessageEncoding is a STRING field
type MessageEncoding struct{ message.StringValue }

//Tag returns tag.MessageEncoding (347)
func (f MessageEncoding) Tag() fix.Tag { return tag.MessageEncoding }

//BuildMessageEncoding returns a new MessageEncoding initialized with val
func BuildMessageEncoding(val string) MessageEncoding {
	var field MessageEncoding
	field.Value = val
	return field
}

//MessageEventSource is a STRING field
type MessageEventSource struct{ message.StringValue }

//Tag returns tag.MessageEventSource (1011)
func (f MessageEventSource) Tag() fix.Tag { return tag.MessageEventSource }

//BuildMessageEventSource returns a new MessageEventSource initialized with val
func BuildMessageEventSource(val string) MessageEventSource {
	var field MessageEventSource
	field.Value = val
	return field
}

//MidPx is a PRICE field
type MidPx struct{ message.PriceValue }

//Tag returns tag.MidPx (631)
func (f MidPx) Tag() fix.Tag { return tag.MidPx }

//BuildMidPx returns a new MidPx initialized with val
func BuildMidPx(val float64) MidPx {
	var field MidPx
	field.Value = val
	return field
}

//MidYield is a PERCENTAGE field
type MidYield struct{ message.PercentageValue }

//Tag returns tag.MidYield (633)
func (f MidYield) Tag() fix.Tag { return tag.MidYield }

//BuildMidYield returns a new MidYield initialized with val
func BuildMidYield(val float64) MidYield {
	var field MidYield
	field.Value = val
	return field
}

//MinBidSize is a QTY field
type MinBidSize struct{ message.QtyValue }

//Tag returns tag.MinBidSize (647)
func (f MinBidSize) Tag() fix.Tag { return tag.MinBidSize }

//BuildMinBidSize returns a new MinBidSize initialized with val
func BuildMinBidSize(val float64) MinBidSize {
	var field MinBidSize
	field.Value = val
	return field
}

//MinLotSize is a QTY field
type MinLotSize struct{ message.QtyValue }

//Tag returns tag.MinLotSize (1231)
func (f MinLotSize) Tag() fix.Tag { return tag.MinLotSize }

//BuildMinLotSize returns a new MinLotSize initialized with val
func BuildMinLotSize(val float64) MinLotSize {
	var field MinLotSize
	field.Value = val
	return field
}

//MinOfferSize is a QTY field
type MinOfferSize struct{ message.QtyValue }

//Tag returns tag.MinOfferSize (648)
func (f MinOfferSize) Tag() fix.Tag { return tag.MinOfferSize }

//BuildMinOfferSize returns a new MinOfferSize initialized with val
func BuildMinOfferSize(val float64) MinOfferSize {
	var field MinOfferSize
	field.Value = val
	return field
}

//MinPriceIncrement is a FLOAT field
type MinPriceIncrement struct{ message.FloatValue }

//Tag returns tag.MinPriceIncrement (969)
func (f MinPriceIncrement) Tag() fix.Tag { return tag.MinPriceIncrement }

//BuildMinPriceIncrement returns a new MinPriceIncrement initialized with val
func BuildMinPriceIncrement(val float64) MinPriceIncrement {
	var field MinPriceIncrement
	field.Value = val
	return field
}

//MinPriceIncrementAmount is a AMT field
type MinPriceIncrementAmount struct{ message.AmtValue }

//Tag returns tag.MinPriceIncrementAmount (1146)
func (f MinPriceIncrementAmount) Tag() fix.Tag { return tag.MinPriceIncrementAmount }

//BuildMinPriceIncrementAmount returns a new MinPriceIncrementAmount initialized with val
func BuildMinPriceIncrementAmount(val float64) MinPriceIncrementAmount {
	var field MinPriceIncrementAmount
	field.Value = val
	return field
}

//MinQty is a QTY field
type MinQty struct{ message.QtyValue }

//Tag returns tag.MinQty (110)
func (f MinQty) Tag() fix.Tag { return tag.MinQty }

//BuildMinQty returns a new MinQty initialized with val
func BuildMinQty(val float64) MinQty {
	var field MinQty
	field.Value = val
	return field
}

//MinTradeVol is a QTY field
type MinTradeVol struct{ message.QtyValue }

//Tag returns tag.MinTradeVol (562)
func (f MinTradeVol) Tag() fix.Tag { return tag.MinTradeVol }

//BuildMinTradeVol returns a new MinTradeVol initialized with val
func BuildMinTradeVol(val float64) MinTradeVol {
	var field MinTradeVol
	field.Value = val
	return field
}

//MiscFeeAmt is a AMT field
type MiscFeeAmt struct{ message.AmtValue }

//Tag returns tag.MiscFeeAmt (137)
func (f MiscFeeAmt) Tag() fix.Tag { return tag.MiscFeeAmt }

//BuildMiscFeeAmt returns a new MiscFeeAmt initialized with val
func BuildMiscFeeAmt(val float64) MiscFeeAmt {
	var field MiscFeeAmt
	field.Value = val
	return field
}

//MiscFeeBasis is a INT field
type MiscFeeBasis struct{ message.IntValue }

//Tag returns tag.MiscFeeBasis (891)
func (f MiscFeeBasis) Tag() fix.Tag { return tag.MiscFeeBasis }

//BuildMiscFeeBasis returns a new MiscFeeBasis initialized with val
func BuildMiscFeeBasis(val int) MiscFeeBasis {
	var field MiscFeeBasis
	field.Value = val
	return field
}

//MiscFeeCurr is a CURRENCY field
type MiscFeeCurr struct{ message.CurrencyValue }

//Tag returns tag.MiscFeeCurr (138)
func (f MiscFeeCurr) Tag() fix.Tag { return tag.MiscFeeCurr }

//BuildMiscFeeCurr returns a new MiscFeeCurr initialized with val
func BuildMiscFeeCurr(val string) MiscFeeCurr {
	var field MiscFeeCurr
	field.Value = val
	return field
}

//MiscFeeType is a STRING field
type MiscFeeType struct{ message.StringValue }

//Tag returns tag.MiscFeeType (139)
func (f MiscFeeType) Tag() fix.Tag { return tag.MiscFeeType }

//BuildMiscFeeType returns a new MiscFeeType initialized with val
func BuildMiscFeeType(val string) MiscFeeType {
	var field MiscFeeType
	field.Value = val
	return field
}

//MktBidPx is a PRICE field
type MktBidPx struct{ message.PriceValue }

//Tag returns tag.MktBidPx (645)
func (f MktBidPx) Tag() fix.Tag { return tag.MktBidPx }

//BuildMktBidPx returns a new MktBidPx initialized with val
func BuildMktBidPx(val float64) MktBidPx {
	var field MktBidPx
	field.Value = val
	return field
}

//MktOfferPx is a PRICE field
type MktOfferPx struct{ message.PriceValue }

//Tag returns tag.MktOfferPx (646)
func (f MktOfferPx) Tag() fix.Tag { return tag.MktOfferPx }

//BuildMktOfferPx returns a new MktOfferPx initialized with val
func BuildMktOfferPx(val float64) MktOfferPx {
	var field MktOfferPx
	field.Value = val
	return field
}

//ModelType is a INT field
type ModelType struct{ message.IntValue }

//Tag returns tag.ModelType (1434)
func (f ModelType) Tag() fix.Tag { return tag.ModelType }

//BuildModelType returns a new ModelType initialized with val
func BuildModelType(val int) ModelType {
	var field ModelType
	field.Value = val
	return field
}

//MoneyLaunderingStatus is a CHAR field
type MoneyLaunderingStatus struct{ message.CharValue }

//Tag returns tag.MoneyLaunderingStatus (481)
func (f MoneyLaunderingStatus) Tag() fix.Tag { return tag.MoneyLaunderingStatus }

//BuildMoneyLaunderingStatus returns a new MoneyLaunderingStatus initialized with val
func BuildMoneyLaunderingStatus(val string) MoneyLaunderingStatus {
	var field MoneyLaunderingStatus
	field.Value = val
	return field
}

//MsgDirection is a CHAR field
type MsgDirection struct{ message.CharValue }

//Tag returns tag.MsgDirection (385)
func (f MsgDirection) Tag() fix.Tag { return tag.MsgDirection }

//BuildMsgDirection returns a new MsgDirection initialized with val
func BuildMsgDirection(val string) MsgDirection {
	var field MsgDirection
	field.Value = val
	return field
}

//MsgSeqNum is a SEQNUM field
type MsgSeqNum struct{ message.SeqNumValue }

//Tag returns tag.MsgSeqNum (34)
func (f MsgSeqNum) Tag() fix.Tag { return tag.MsgSeqNum }

//BuildMsgSeqNum returns a new MsgSeqNum initialized with val
func BuildMsgSeqNum(val int) MsgSeqNum {
	var field MsgSeqNum
	field.Value = val
	return field
}

//MsgType is a STRING field
type MsgType struct{ message.StringValue }

//Tag returns tag.MsgType (35)
func (f MsgType) Tag() fix.Tag { return tag.MsgType }

//BuildMsgType returns a new MsgType initialized with val
func BuildMsgType(val string) MsgType {
	var field MsgType
	field.Value = val
	return field
}

//MultiLegReportingType is a CHAR field
type MultiLegReportingType struct{ message.CharValue }

//Tag returns tag.MultiLegReportingType (442)
func (f MultiLegReportingType) Tag() fix.Tag { return tag.MultiLegReportingType }

//BuildMultiLegReportingType returns a new MultiLegReportingType initialized with val
func BuildMultiLegReportingType(val string) MultiLegReportingType {
	var field MultiLegReportingType
	field.Value = val
	return field
}

//MultiLegRptTypeReq is a INT field
type MultiLegRptTypeReq struct{ message.IntValue }

//Tag returns tag.MultiLegRptTypeReq (563)
func (f MultiLegRptTypeReq) Tag() fix.Tag { return tag.MultiLegRptTypeReq }

//BuildMultiLegRptTypeReq returns a new MultiLegRptTypeReq initialized with val
func BuildMultiLegRptTypeReq(val int) MultiLegRptTypeReq {
	var field MultiLegRptTypeReq
	field.Value = val
	return field
}

//MultilegModel is a INT field
type MultilegModel struct{ message.IntValue }

//Tag returns tag.MultilegModel (1377)
func (f MultilegModel) Tag() fix.Tag { return tag.MultilegModel }

//BuildMultilegModel returns a new MultilegModel initialized with val
func BuildMultilegModel(val int) MultilegModel {
	var field MultilegModel
	field.Value = val
	return field
}

//MultilegPriceMethod is a INT field
type MultilegPriceMethod struct{ message.IntValue }

//Tag returns tag.MultilegPriceMethod (1378)
func (f MultilegPriceMethod) Tag() fix.Tag { return tag.MultilegPriceMethod }

//BuildMultilegPriceMethod returns a new MultilegPriceMethod initialized with val
func BuildMultilegPriceMethod(val int) MultilegPriceMethod {
	var field MultilegPriceMethod
	field.Value = val
	return field
}

//NTPositionLimit is a INT field
type NTPositionLimit struct{ message.IntValue }

//Tag returns tag.NTPositionLimit (971)
func (f NTPositionLimit) Tag() fix.Tag { return tag.NTPositionLimit }

//BuildNTPositionLimit returns a new NTPositionLimit initialized with val
func BuildNTPositionLimit(val int) NTPositionLimit {
	var field NTPositionLimit
	field.Value = val
	return field
}

//Nested2PartyID is a STRING field
type Nested2PartyID struct{ message.StringValue }

//Tag returns tag.Nested2PartyID (757)
func (f Nested2PartyID) Tag() fix.Tag { return tag.Nested2PartyID }

//BuildNested2PartyID returns a new Nested2PartyID initialized with val
func BuildNested2PartyID(val string) Nested2PartyID {
	var field Nested2PartyID
	field.Value = val
	return field
}

//Nested2PartyIDSource is a CHAR field
type Nested2PartyIDSource struct{ message.CharValue }

//Tag returns tag.Nested2PartyIDSource (758)
func (f Nested2PartyIDSource) Tag() fix.Tag { return tag.Nested2PartyIDSource }

//BuildNested2PartyIDSource returns a new Nested2PartyIDSource initialized with val
func BuildNested2PartyIDSource(val string) Nested2PartyIDSource {
	var field Nested2PartyIDSource
	field.Value = val
	return field
}

//Nested2PartyRole is a INT field
type Nested2PartyRole struct{ message.IntValue }

//Tag returns tag.Nested2PartyRole (759)
func (f Nested2PartyRole) Tag() fix.Tag { return tag.Nested2PartyRole }

//BuildNested2PartyRole returns a new Nested2PartyRole initialized with val
func BuildNested2PartyRole(val int) Nested2PartyRole {
	var field Nested2PartyRole
	field.Value = val
	return field
}

//Nested2PartySubID is a STRING field
type Nested2PartySubID struct{ message.StringValue }

//Tag returns tag.Nested2PartySubID (760)
func (f Nested2PartySubID) Tag() fix.Tag { return tag.Nested2PartySubID }

//BuildNested2PartySubID returns a new Nested2PartySubID initialized with val
func BuildNested2PartySubID(val string) Nested2PartySubID {
	var field Nested2PartySubID
	field.Value = val
	return field
}

//Nested2PartySubIDType is a INT field
type Nested2PartySubIDType struct{ message.IntValue }

//Tag returns tag.Nested2PartySubIDType (807)
func (f Nested2PartySubIDType) Tag() fix.Tag { return tag.Nested2PartySubIDType }

//BuildNested2PartySubIDType returns a new Nested2PartySubIDType initialized with val
func BuildNested2PartySubIDType(val int) Nested2PartySubIDType {
	var field Nested2PartySubIDType
	field.Value = val
	return field
}

//Nested3PartyID is a STRING field
type Nested3PartyID struct{ message.StringValue }

//Tag returns tag.Nested3PartyID (949)
func (f Nested3PartyID) Tag() fix.Tag { return tag.Nested3PartyID }

//BuildNested3PartyID returns a new Nested3PartyID initialized with val
func BuildNested3PartyID(val string) Nested3PartyID {
	var field Nested3PartyID
	field.Value = val
	return field
}

//Nested3PartyIDSource is a CHAR field
type Nested3PartyIDSource struct{ message.CharValue }

//Tag returns tag.Nested3PartyIDSource (950)
func (f Nested3PartyIDSource) Tag() fix.Tag { return tag.Nested3PartyIDSource }

//BuildNested3PartyIDSource returns a new Nested3PartyIDSource initialized with val
func BuildNested3PartyIDSource(val string) Nested3PartyIDSource {
	var field Nested3PartyIDSource
	field.Value = val
	return field
}

//Nested3PartyRole is a INT field
type Nested3PartyRole struct{ message.IntValue }

//Tag returns tag.Nested3PartyRole (951)
func (f Nested3PartyRole) Tag() fix.Tag { return tag.Nested3PartyRole }

//BuildNested3PartyRole returns a new Nested3PartyRole initialized with val
func BuildNested3PartyRole(val int) Nested3PartyRole {
	var field Nested3PartyRole
	field.Value = val
	return field
}

//Nested3PartySubID is a STRING field
type Nested3PartySubID struct{ message.StringValue }

//Tag returns tag.Nested3PartySubID (953)
func (f Nested3PartySubID) Tag() fix.Tag { return tag.Nested3PartySubID }

//BuildNested3PartySubID returns a new Nested3PartySubID initialized with val
func BuildNested3PartySubID(val string) Nested3PartySubID {
	var field Nested3PartySubID
	field.Value = val
	return field
}

//Nested3PartySubIDType is a INT field
type Nested3PartySubIDType struct{ message.IntValue }

//Tag returns tag.Nested3PartySubIDType (954)
func (f Nested3PartySubIDType) Tag() fix.Tag { return tag.Nested3PartySubIDType }

//BuildNested3PartySubIDType returns a new Nested3PartySubIDType initialized with val
func BuildNested3PartySubIDType(val int) Nested3PartySubIDType {
	var field Nested3PartySubIDType
	field.Value = val
	return field
}

//Nested4PartyID is a STRING field
type Nested4PartyID struct{ message.StringValue }

//Tag returns tag.Nested4PartyID (1415)
func (f Nested4PartyID) Tag() fix.Tag { return tag.Nested4PartyID }

//BuildNested4PartyID returns a new Nested4PartyID initialized with val
func BuildNested4PartyID(val string) Nested4PartyID {
	var field Nested4PartyID
	field.Value = val
	return field
}

//Nested4PartyIDSource is a CHAR field
type Nested4PartyIDSource struct{ message.CharValue }

//Tag returns tag.Nested4PartyIDSource (1416)
func (f Nested4PartyIDSource) Tag() fix.Tag { return tag.Nested4PartyIDSource }

//BuildNested4PartyIDSource returns a new Nested4PartyIDSource initialized with val
func BuildNested4PartyIDSource(val string) Nested4PartyIDSource {
	var field Nested4PartyIDSource
	field.Value = val
	return field
}

//Nested4PartyRole is a INT field
type Nested4PartyRole struct{ message.IntValue }

//Tag returns tag.Nested4PartyRole (1417)
func (f Nested4PartyRole) Tag() fix.Tag { return tag.Nested4PartyRole }

//BuildNested4PartyRole returns a new Nested4PartyRole initialized with val
func BuildNested4PartyRole(val int) Nested4PartyRole {
	var field Nested4PartyRole
	field.Value = val
	return field
}

//Nested4PartySubID is a STRING field
type Nested4PartySubID struct{ message.StringValue }

//Tag returns tag.Nested4PartySubID (1412)
func (f Nested4PartySubID) Tag() fix.Tag { return tag.Nested4PartySubID }

//BuildNested4PartySubID returns a new Nested4PartySubID initialized with val
func BuildNested4PartySubID(val string) Nested4PartySubID {
	var field Nested4PartySubID
	field.Value = val
	return field
}

//Nested4PartySubIDType is a INT field
type Nested4PartySubIDType struct{ message.IntValue }

//Tag returns tag.Nested4PartySubIDType (1411)
func (f Nested4PartySubIDType) Tag() fix.Tag { return tag.Nested4PartySubIDType }

//BuildNested4PartySubIDType returns a new Nested4PartySubIDType initialized with val
func BuildNested4PartySubIDType(val int) Nested4PartySubIDType {
	var field Nested4PartySubIDType
	field.Value = val
	return field
}

//NestedInstrAttribType is a INT field
type NestedInstrAttribType struct{ message.IntValue }

//Tag returns tag.NestedInstrAttribType (1210)
func (f NestedInstrAttribType) Tag() fix.Tag { return tag.NestedInstrAttribType }

//BuildNestedInstrAttribType returns a new NestedInstrAttribType initialized with val
func BuildNestedInstrAttribType(val int) NestedInstrAttribType {
	var field NestedInstrAttribType
	field.Value = val
	return field
}

//NestedInstrAttribValue is a STRING field
type NestedInstrAttribValue struct{ message.StringValue }

//Tag returns tag.NestedInstrAttribValue (1211)
func (f NestedInstrAttribValue) Tag() fix.Tag { return tag.NestedInstrAttribValue }

//BuildNestedInstrAttribValue returns a new NestedInstrAttribValue initialized with val
func BuildNestedInstrAttribValue(val string) NestedInstrAttribValue {
	var field NestedInstrAttribValue
	field.Value = val
	return field
}

//NestedPartyID is a STRING field
type NestedPartyID struct{ message.StringValue }

//Tag returns tag.NestedPartyID (524)
func (f NestedPartyID) Tag() fix.Tag { return tag.NestedPartyID }

//BuildNestedPartyID returns a new NestedPartyID initialized with val
func BuildNestedPartyID(val string) NestedPartyID {
	var field NestedPartyID
	field.Value = val
	return field
}

//NestedPartyIDSource is a CHAR field
type NestedPartyIDSource struct{ message.CharValue }

//Tag returns tag.NestedPartyIDSource (525)
func (f NestedPartyIDSource) Tag() fix.Tag { return tag.NestedPartyIDSource }

//BuildNestedPartyIDSource returns a new NestedPartyIDSource initialized with val
func BuildNestedPartyIDSource(val string) NestedPartyIDSource {
	var field NestedPartyIDSource
	field.Value = val
	return field
}

//NestedPartyRole is a INT field
type NestedPartyRole struct{ message.IntValue }

//Tag returns tag.NestedPartyRole (538)
func (f NestedPartyRole) Tag() fix.Tag { return tag.NestedPartyRole }

//BuildNestedPartyRole returns a new NestedPartyRole initialized with val
func BuildNestedPartyRole(val int) NestedPartyRole {
	var field NestedPartyRole
	field.Value = val
	return field
}

//NestedPartySubID is a STRING field
type NestedPartySubID struct{ message.StringValue }

//Tag returns tag.NestedPartySubID (545)
func (f NestedPartySubID) Tag() fix.Tag { return tag.NestedPartySubID }

//BuildNestedPartySubID returns a new NestedPartySubID initialized with val
func BuildNestedPartySubID(val string) NestedPartySubID {
	var field NestedPartySubID
	field.Value = val
	return field
}

//NestedPartySubIDType is a INT field
type NestedPartySubIDType struct{ message.IntValue }

//Tag returns tag.NestedPartySubIDType (805)
func (f NestedPartySubIDType) Tag() fix.Tag { return tag.NestedPartySubIDType }

//BuildNestedPartySubIDType returns a new NestedPartySubIDType initialized with val
func BuildNestedPartySubIDType(val int) NestedPartySubIDType {
	var field NestedPartySubIDType
	field.Value = val
	return field
}

//NetChgPrevDay is a PRICEOFFSET field
type NetChgPrevDay struct{ message.PriceOffsetValue }

//Tag returns tag.NetChgPrevDay (451)
func (f NetChgPrevDay) Tag() fix.Tag { return tag.NetChgPrevDay }

//BuildNetChgPrevDay returns a new NetChgPrevDay initialized with val
func BuildNetChgPrevDay(val float64) NetChgPrevDay {
	var field NetChgPrevDay
	field.Value = val
	return field
}

//NetGrossInd is a INT field
type NetGrossInd struct{ message.IntValue }

//Tag returns tag.NetGrossInd (430)
func (f NetGrossInd) Tag() fix.Tag { return tag.NetGrossInd }

//BuildNetGrossInd returns a new NetGrossInd initialized with val
func BuildNetGrossInd(val int) NetGrossInd {
	var field NetGrossInd
	field.Value = val
	return field
}

//NetMoney is a AMT field
type NetMoney struct{ message.AmtValue }

//Tag returns tag.NetMoney (118)
func (f NetMoney) Tag() fix.Tag { return tag.NetMoney }

//BuildNetMoney returns a new NetMoney initialized with val
func BuildNetMoney(val float64) NetMoney {
	var field NetMoney
	field.Value = val
	return field
}

//NetworkRequestID is a STRING field
type NetworkRequestID struct{ message.StringValue }

//Tag returns tag.NetworkRequestID (933)
func (f NetworkRequestID) Tag() fix.Tag { return tag.NetworkRequestID }

//BuildNetworkRequestID returns a new NetworkRequestID initialized with val
func BuildNetworkRequestID(val string) NetworkRequestID {
	var field NetworkRequestID
	field.Value = val
	return field
}

//NetworkRequestType is a INT field
type NetworkRequestType struct{ message.IntValue }

//Tag returns tag.NetworkRequestType (935)
func (f NetworkRequestType) Tag() fix.Tag { return tag.NetworkRequestType }

//BuildNetworkRequestType returns a new NetworkRequestType initialized with val
func BuildNetworkRequestType(val int) NetworkRequestType {
	var field NetworkRequestType
	field.Value = val
	return field
}

//NetworkResponseID is a STRING field
type NetworkResponseID struct{ message.StringValue }

//Tag returns tag.NetworkResponseID (932)
func (f NetworkResponseID) Tag() fix.Tag { return tag.NetworkResponseID }

//BuildNetworkResponseID returns a new NetworkResponseID initialized with val
func BuildNetworkResponseID(val string) NetworkResponseID {
	var field NetworkResponseID
	field.Value = val
	return field
}

//NetworkStatusResponseType is a INT field
type NetworkStatusResponseType struct{ message.IntValue }

//Tag returns tag.NetworkStatusResponseType (937)
func (f NetworkStatusResponseType) Tag() fix.Tag { return tag.NetworkStatusResponseType }

//BuildNetworkStatusResponseType returns a new NetworkStatusResponseType initialized with val
func BuildNetworkStatusResponseType(val int) NetworkStatusResponseType {
	var field NetworkStatusResponseType
	field.Value = val
	return field
}

//NewPassword is a STRING field
type NewPassword struct{ message.StringValue }

//Tag returns tag.NewPassword (925)
func (f NewPassword) Tag() fix.Tag { return tag.NewPassword }

//BuildNewPassword returns a new NewPassword initialized with val
func BuildNewPassword(val string) NewPassword {
	var field NewPassword
	field.Value = val
	return field
}

//NewSeqNo is a SEQNUM field
type NewSeqNo struct{ message.SeqNumValue }

//Tag returns tag.NewSeqNo (36)
func (f NewSeqNo) Tag() fix.Tag { return tag.NewSeqNo }

//BuildNewSeqNo returns a new NewSeqNo initialized with val
func BuildNewSeqNo(val int) NewSeqNo {
	var field NewSeqNo
	field.Value = val
	return field
}

//NewsCategory is a INT field
type NewsCategory struct{ message.IntValue }

//Tag returns tag.NewsCategory (1473)
func (f NewsCategory) Tag() fix.Tag { return tag.NewsCategory }

//BuildNewsCategory returns a new NewsCategory initialized with val
func BuildNewsCategory(val int) NewsCategory {
	var field NewsCategory
	field.Value = val
	return field
}

//NewsID is a STRING field
type NewsID struct{ message.StringValue }

//Tag returns tag.NewsID (1472)
func (f NewsID) Tag() fix.Tag { return tag.NewsID }

//BuildNewsID returns a new NewsID initialized with val
func BuildNewsID(val string) NewsID {
	var field NewsID
	field.Value = val
	return field
}

//NewsRefID is a STRING field
type NewsRefID struct{ message.StringValue }

//Tag returns tag.NewsRefID (1476)
func (f NewsRefID) Tag() fix.Tag { return tag.NewsRefID }

//BuildNewsRefID returns a new NewsRefID initialized with val
func BuildNewsRefID(val string) NewsRefID {
	var field NewsRefID
	field.Value = val
	return field
}

//NewsRefType is a INT field
type NewsRefType struct{ message.IntValue }

//Tag returns tag.NewsRefType (1477)
func (f NewsRefType) Tag() fix.Tag { return tag.NewsRefType }

//BuildNewsRefType returns a new NewsRefType initialized with val
func BuildNewsRefType(val int) NewsRefType {
	var field NewsRefType
	field.Value = val
	return field
}

//NextExpectedMsgSeqNum is a SEQNUM field
type NextExpectedMsgSeqNum struct{ message.SeqNumValue }

//Tag returns tag.NextExpectedMsgSeqNum (789)
func (f NextExpectedMsgSeqNum) Tag() fix.Tag { return tag.NextExpectedMsgSeqNum }

//BuildNextExpectedMsgSeqNum returns a new NextExpectedMsgSeqNum initialized with val
func BuildNextExpectedMsgSeqNum(val int) NextExpectedMsgSeqNum {
	var field NextExpectedMsgSeqNum
	field.Value = val
	return field
}

//NoAffectedOrders is a NUMINGROUP field
type NoAffectedOrders struct{ message.NumInGroupValue }

//Tag returns tag.NoAffectedOrders (534)
func (f NoAffectedOrders) Tag() fix.Tag { return tag.NoAffectedOrders }

//BuildNoAffectedOrders returns a new NoAffectedOrders initialized with val
func BuildNoAffectedOrders(val int) NoAffectedOrders {
	var field NoAffectedOrders
	field.Value = val
	return field
}

//NoAllocs is a NUMINGROUP field
type NoAllocs struct{ message.NumInGroupValue }

//Tag returns tag.NoAllocs (78)
func (f NoAllocs) Tag() fix.Tag { return tag.NoAllocs }

//BuildNoAllocs returns a new NoAllocs initialized with val
func BuildNoAllocs(val int) NoAllocs {
	var field NoAllocs
	field.Value = val
	return field
}

//NoAltMDSource is a NUMINGROUP field
type NoAltMDSource struct{ message.NumInGroupValue }

//Tag returns tag.NoAltMDSource (816)
func (f NoAltMDSource) Tag() fix.Tag { return tag.NoAltMDSource }

//BuildNoAltMDSource returns a new NoAltMDSource initialized with val
func BuildNoAltMDSource(val int) NoAltMDSource {
	var field NoAltMDSource
	field.Value = val
	return field
}

//NoApplIDs is a NUMINGROUP field
type NoApplIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoApplIDs (1351)
func (f NoApplIDs) Tag() fix.Tag { return tag.NoApplIDs }

//BuildNoApplIDs returns a new NoApplIDs initialized with val
func BuildNoApplIDs(val int) NoApplIDs {
	var field NoApplIDs
	field.Value = val
	return field
}

//NoAsgnReqs is a NUMINGROUP field
type NoAsgnReqs struct{ message.NumInGroupValue }

//Tag returns tag.NoAsgnReqs (1499)
func (f NoAsgnReqs) Tag() fix.Tag { return tag.NoAsgnReqs }

//BuildNoAsgnReqs returns a new NoAsgnReqs initialized with val
func BuildNoAsgnReqs(val int) NoAsgnReqs {
	var field NoAsgnReqs
	field.Value = val
	return field
}

//NoBidComponents is a NUMINGROUP field
type NoBidComponents struct{ message.NumInGroupValue }

//Tag returns tag.NoBidComponents (420)
func (f NoBidComponents) Tag() fix.Tag { return tag.NoBidComponents }

//BuildNoBidComponents returns a new NoBidComponents initialized with val
func BuildNoBidComponents(val int) NoBidComponents {
	var field NoBidComponents
	field.Value = val
	return field
}

//NoBidDescriptors is a NUMINGROUP field
type NoBidDescriptors struct{ message.NumInGroupValue }

//Tag returns tag.NoBidDescriptors (398)
func (f NoBidDescriptors) Tag() fix.Tag { return tag.NoBidDescriptors }

//BuildNoBidDescriptors returns a new NoBidDescriptors initialized with val
func BuildNoBidDescriptors(val int) NoBidDescriptors {
	var field NoBidDescriptors
	field.Value = val
	return field
}

//NoCapacities is a NUMINGROUP field
type NoCapacities struct{ message.NumInGroupValue }

//Tag returns tag.NoCapacities (862)
func (f NoCapacities) Tag() fix.Tag { return tag.NoCapacities }

//BuildNoCapacities returns a new NoCapacities initialized with val
func BuildNoCapacities(val int) NoCapacities {
	var field NoCapacities
	field.Value = val
	return field
}

//NoClearingInstructions is a NUMINGROUP field
type NoClearingInstructions struct{ message.NumInGroupValue }

//Tag returns tag.NoClearingInstructions (576)
func (f NoClearingInstructions) Tag() fix.Tag { return tag.NoClearingInstructions }

//BuildNoClearingInstructions returns a new NoClearingInstructions initialized with val
func BuildNoClearingInstructions(val int) NoClearingInstructions {
	var field NoClearingInstructions
	field.Value = val
	return field
}

//NoCollInquiryQualifier is a NUMINGROUP field
type NoCollInquiryQualifier struct{ message.NumInGroupValue }

//Tag returns tag.NoCollInquiryQualifier (938)
func (f NoCollInquiryQualifier) Tag() fix.Tag { return tag.NoCollInquiryQualifier }

//BuildNoCollInquiryQualifier returns a new NoCollInquiryQualifier initialized with val
func BuildNoCollInquiryQualifier(val int) NoCollInquiryQualifier {
	var field NoCollInquiryQualifier
	field.Value = val
	return field
}

//NoCompIDs is a NUMINGROUP field
type NoCompIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoCompIDs (936)
func (f NoCompIDs) Tag() fix.Tag { return tag.NoCompIDs }

//BuildNoCompIDs returns a new NoCompIDs initialized with val
func BuildNoCompIDs(val int) NoCompIDs {
	var field NoCompIDs
	field.Value = val
	return field
}

//NoComplexEventDates is a NUMINGROUP field
type NoComplexEventDates struct{ message.NumInGroupValue }

//Tag returns tag.NoComplexEventDates (1491)
func (f NoComplexEventDates) Tag() fix.Tag { return tag.NoComplexEventDates }

//BuildNoComplexEventDates returns a new NoComplexEventDates initialized with val
func BuildNoComplexEventDates(val int) NoComplexEventDates {
	var field NoComplexEventDates
	field.Value = val
	return field
}

//NoComplexEventTimes is a NUMINGROUP field
type NoComplexEventTimes struct{ message.NumInGroupValue }

//Tag returns tag.NoComplexEventTimes (1494)
func (f NoComplexEventTimes) Tag() fix.Tag { return tag.NoComplexEventTimes }

//BuildNoComplexEventTimes returns a new NoComplexEventTimes initialized with val
func BuildNoComplexEventTimes(val int) NoComplexEventTimes {
	var field NoComplexEventTimes
	field.Value = val
	return field
}

//NoComplexEvents is a NUMINGROUP field
type NoComplexEvents struct{ message.NumInGroupValue }

//Tag returns tag.NoComplexEvents (1483)
func (f NoComplexEvents) Tag() fix.Tag { return tag.NoComplexEvents }

//BuildNoComplexEvents returns a new NoComplexEvents initialized with val
func BuildNoComplexEvents(val int) NoComplexEvents {
	var field NoComplexEvents
	field.Value = val
	return field
}

//NoContAmts is a NUMINGROUP field
type NoContAmts struct{ message.NumInGroupValue }

//Tag returns tag.NoContAmts (518)
func (f NoContAmts) Tag() fix.Tag { return tag.NoContAmts }

//BuildNoContAmts returns a new NoContAmts initialized with val
func BuildNoContAmts(val int) NoContAmts {
	var field NoContAmts
	field.Value = val
	return field
}

//NoContextPartyIDs is a NUMINGROUP field
type NoContextPartyIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoContextPartyIDs (1522)
func (f NoContextPartyIDs) Tag() fix.Tag { return tag.NoContextPartyIDs }

//BuildNoContextPartyIDs returns a new NoContextPartyIDs initialized with val
func BuildNoContextPartyIDs(val int) NoContextPartyIDs {
	var field NoContextPartyIDs
	field.Value = val
	return field
}

//NoContextPartySubIDs is a NUMINGROUP field
type NoContextPartySubIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoContextPartySubIDs (1526)
func (f NoContextPartySubIDs) Tag() fix.Tag { return tag.NoContextPartySubIDs }

//BuildNoContextPartySubIDs returns a new NoContextPartySubIDs initialized with val
func BuildNoContextPartySubIDs(val int) NoContextPartySubIDs {
	var field NoContextPartySubIDs
	field.Value = val
	return field
}

//NoContraBrokers is a NUMINGROUP field
type NoContraBrokers struct{ message.NumInGroupValue }

//Tag returns tag.NoContraBrokers (382)
func (f NoContraBrokers) Tag() fix.Tag { return tag.NoContraBrokers }

//BuildNoContraBrokers returns a new NoContraBrokers initialized with val
func BuildNoContraBrokers(val int) NoContraBrokers {
	var field NoContraBrokers
	field.Value = val
	return field
}

//NoDates is a INT field
type NoDates struct{ message.IntValue }

//Tag returns tag.NoDates (580)
func (f NoDates) Tag() fix.Tag { return tag.NoDates }

//BuildNoDates returns a new NoDates initialized with val
func BuildNoDates(val int) NoDates {
	var field NoDates
	field.Value = val
	return field
}

//NoDerivativeEvents is a NUMINGROUP field
type NoDerivativeEvents struct{ message.NumInGroupValue }

//Tag returns tag.NoDerivativeEvents (1286)
func (f NoDerivativeEvents) Tag() fix.Tag { return tag.NoDerivativeEvents }

//BuildNoDerivativeEvents returns a new NoDerivativeEvents initialized with val
func BuildNoDerivativeEvents(val int) NoDerivativeEvents {
	var field NoDerivativeEvents
	field.Value = val
	return field
}

//NoDerivativeInstrAttrib is a NUMINGROUP field
type NoDerivativeInstrAttrib struct{ message.NumInGroupValue }

//Tag returns tag.NoDerivativeInstrAttrib (1311)
func (f NoDerivativeInstrAttrib) Tag() fix.Tag { return tag.NoDerivativeInstrAttrib }

//BuildNoDerivativeInstrAttrib returns a new NoDerivativeInstrAttrib initialized with val
func BuildNoDerivativeInstrAttrib(val int) NoDerivativeInstrAttrib {
	var field NoDerivativeInstrAttrib
	field.Value = val
	return field
}

//NoDerivativeInstrumentParties is a NUMINGROUP field
type NoDerivativeInstrumentParties struct{ message.NumInGroupValue }

//Tag returns tag.NoDerivativeInstrumentParties (1292)
func (f NoDerivativeInstrumentParties) Tag() fix.Tag { return tag.NoDerivativeInstrumentParties }

//BuildNoDerivativeInstrumentParties returns a new NoDerivativeInstrumentParties initialized with val
func BuildNoDerivativeInstrumentParties(val int) NoDerivativeInstrumentParties {
	var field NoDerivativeInstrumentParties
	field.Value = val
	return field
}

//NoDerivativeInstrumentPartySubIDs is a NUMINGROUP field
type NoDerivativeInstrumentPartySubIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoDerivativeInstrumentPartySubIDs (1296)
func (f NoDerivativeInstrumentPartySubIDs) Tag() fix.Tag { return tag.NoDerivativeInstrumentPartySubIDs }

//BuildNoDerivativeInstrumentPartySubIDs returns a new NoDerivativeInstrumentPartySubIDs initialized with val
func BuildNoDerivativeInstrumentPartySubIDs(val int) NoDerivativeInstrumentPartySubIDs {
	var field NoDerivativeInstrumentPartySubIDs
	field.Value = val
	return field
}

//NoDerivativeSecurityAltID is a NUMINGROUP field
type NoDerivativeSecurityAltID struct{ message.NumInGroupValue }

//Tag returns tag.NoDerivativeSecurityAltID (1218)
func (f NoDerivativeSecurityAltID) Tag() fix.Tag { return tag.NoDerivativeSecurityAltID }

//BuildNoDerivativeSecurityAltID returns a new NoDerivativeSecurityAltID initialized with val
func BuildNoDerivativeSecurityAltID(val int) NoDerivativeSecurityAltID {
	var field NoDerivativeSecurityAltID
	field.Value = val
	return field
}

//NoDistribInsts is a NUMINGROUP field
type NoDistribInsts struct{ message.NumInGroupValue }

//Tag returns tag.NoDistribInsts (510)
func (f NoDistribInsts) Tag() fix.Tag { return tag.NoDistribInsts }

//BuildNoDistribInsts returns a new NoDistribInsts initialized with val
func BuildNoDistribInsts(val int) NoDistribInsts {
	var field NoDistribInsts
	field.Value = val
	return field
}

//NoDlvyInst is a NUMINGROUP field
type NoDlvyInst struct{ message.NumInGroupValue }

//Tag returns tag.NoDlvyInst (85)
func (f NoDlvyInst) Tag() fix.Tag { return tag.NoDlvyInst }

//BuildNoDlvyInst returns a new NoDlvyInst initialized with val
func BuildNoDlvyInst(val int) NoDlvyInst {
	var field NoDlvyInst
	field.Value = val
	return field
}

//NoEvents is a NUMINGROUP field
type NoEvents struct{ message.NumInGroupValue }

//Tag returns tag.NoEvents (864)
func (f NoEvents) Tag() fix.Tag { return tag.NoEvents }

//BuildNoEvents returns a new NoEvents initialized with val
func BuildNoEvents(val int) NoEvents {
	var field NoEvents
	field.Value = val
	return field
}

//NoExecInstRules is a NUMINGROUP field
type NoExecInstRules struct{ message.NumInGroupValue }

//Tag returns tag.NoExecInstRules (1232)
func (f NoExecInstRules) Tag() fix.Tag { return tag.NoExecInstRules }

//BuildNoExecInstRules returns a new NoExecInstRules initialized with val
func BuildNoExecInstRules(val int) NoExecInstRules {
	var field NoExecInstRules
	field.Value = val
	return field
}

//NoExecs is a NUMINGROUP field
type NoExecs struct{ message.NumInGroupValue }

//Tag returns tag.NoExecs (124)
func (f NoExecs) Tag() fix.Tag { return tag.NoExecs }

//BuildNoExecs returns a new NoExecs initialized with val
func BuildNoExecs(val int) NoExecs {
	var field NoExecs
	field.Value = val
	return field
}

//NoExpiration is a NUMINGROUP field
type NoExpiration struct{ message.NumInGroupValue }

//Tag returns tag.NoExpiration (981)
func (f NoExpiration) Tag() fix.Tag { return tag.NoExpiration }

//BuildNoExpiration returns a new NoExpiration initialized with val
func BuildNoExpiration(val int) NoExpiration {
	var field NoExpiration
	field.Value = val
	return field
}

//NoFills is a NUMINGROUP field
type NoFills struct{ message.NumInGroupValue }

//Tag returns tag.NoFills (1362)
func (f NoFills) Tag() fix.Tag { return tag.NoFills }

//BuildNoFills returns a new NoFills initialized with val
func BuildNoFills(val int) NoFills {
	var field NoFills
	field.Value = val
	return field
}

//NoHops is a NUMINGROUP field
type NoHops struct{ message.NumInGroupValue }

//Tag returns tag.NoHops (627)
func (f NoHops) Tag() fix.Tag { return tag.NoHops }

//BuildNoHops returns a new NoHops initialized with val
func BuildNoHops(val int) NoHops {
	var field NoHops
	field.Value = val
	return field
}

//NoIOIQualifiers is a NUMINGROUP field
type NoIOIQualifiers struct{ message.NumInGroupValue }

//Tag returns tag.NoIOIQualifiers (199)
func (f NoIOIQualifiers) Tag() fix.Tag { return tag.NoIOIQualifiers }

//BuildNoIOIQualifiers returns a new NoIOIQualifiers initialized with val
func BuildNoIOIQualifiers(val int) NoIOIQualifiers {
	var field NoIOIQualifiers
	field.Value = val
	return field
}

//NoInstrAttrib is a NUMINGROUP field
type NoInstrAttrib struct{ message.NumInGroupValue }

//Tag returns tag.NoInstrAttrib (870)
func (f NoInstrAttrib) Tag() fix.Tag { return tag.NoInstrAttrib }

//BuildNoInstrAttrib returns a new NoInstrAttrib initialized with val
func BuildNoInstrAttrib(val int) NoInstrAttrib {
	var field NoInstrAttrib
	field.Value = val
	return field
}

//NoInstrumentParties is a NUMINGROUP field
type NoInstrumentParties struct{ message.NumInGroupValue }

//Tag returns tag.NoInstrumentParties (1018)
func (f NoInstrumentParties) Tag() fix.Tag { return tag.NoInstrumentParties }

//BuildNoInstrumentParties returns a new NoInstrumentParties initialized with val
func BuildNoInstrumentParties(val int) NoInstrumentParties {
	var field NoInstrumentParties
	field.Value = val
	return field
}

//NoInstrumentPartySubIDs is a NUMINGROUP field
type NoInstrumentPartySubIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoInstrumentPartySubIDs (1052)
func (f NoInstrumentPartySubIDs) Tag() fix.Tag { return tag.NoInstrumentPartySubIDs }

//BuildNoInstrumentPartySubIDs returns a new NoInstrumentPartySubIDs initialized with val
func BuildNoInstrumentPartySubIDs(val int) NoInstrumentPartySubIDs {
	var field NoInstrumentPartySubIDs
	field.Value = val
	return field
}

//NoLegAllocs is a NUMINGROUP field
type NoLegAllocs struct{ message.NumInGroupValue }

//Tag returns tag.NoLegAllocs (670)
func (f NoLegAllocs) Tag() fix.Tag { return tag.NoLegAllocs }

//BuildNoLegAllocs returns a new NoLegAllocs initialized with val
func BuildNoLegAllocs(val int) NoLegAllocs {
	var field NoLegAllocs
	field.Value = val
	return field
}

//NoLegSecurityAltID is a STRING field
type NoLegSecurityAltID struct{ message.StringValue }

//Tag returns tag.NoLegSecurityAltID (604)
func (f NoLegSecurityAltID) Tag() fix.Tag { return tag.NoLegSecurityAltID }

//BuildNoLegSecurityAltID returns a new NoLegSecurityAltID initialized with val
func BuildNoLegSecurityAltID(val string) NoLegSecurityAltID {
	var field NoLegSecurityAltID
	field.Value = val
	return field
}

//NoLegStipulations is a NUMINGROUP field
type NoLegStipulations struct{ message.NumInGroupValue }

//Tag returns tag.NoLegStipulations (683)
func (f NoLegStipulations) Tag() fix.Tag { return tag.NoLegStipulations }

//BuildNoLegStipulations returns a new NoLegStipulations initialized with val
func BuildNoLegStipulations(val int) NoLegStipulations {
	var field NoLegStipulations
	field.Value = val
	return field
}

//NoLegs is a NUMINGROUP field
type NoLegs struct{ message.NumInGroupValue }

//Tag returns tag.NoLegs (555)
func (f NoLegs) Tag() fix.Tag { return tag.NoLegs }

//BuildNoLegs returns a new NoLegs initialized with val
func BuildNoLegs(val int) NoLegs {
	var field NoLegs
	field.Value = val
	return field
}

//NoLinesOfText is a NUMINGROUP field
type NoLinesOfText struct{ message.NumInGroupValue }

//Tag returns tag.NoLinesOfText (33)
func (f NoLinesOfText) Tag() fix.Tag { return tag.NoLinesOfText }

//BuildNoLinesOfText returns a new NoLinesOfText initialized with val
func BuildNoLinesOfText(val int) NoLinesOfText {
	var field NoLinesOfText
	field.Value = val
	return field
}

//NoLotTypeRules is a NUMINGROUP field
type NoLotTypeRules struct{ message.NumInGroupValue }

//Tag returns tag.NoLotTypeRules (1234)
func (f NoLotTypeRules) Tag() fix.Tag { return tag.NoLotTypeRules }

//BuildNoLotTypeRules returns a new NoLotTypeRules initialized with val
func BuildNoLotTypeRules(val int) NoLotTypeRules {
	var field NoLotTypeRules
	field.Value = val
	return field
}

//NoMDEntries is a NUMINGROUP field
type NoMDEntries struct{ message.NumInGroupValue }

//Tag returns tag.NoMDEntries (268)
func (f NoMDEntries) Tag() fix.Tag { return tag.NoMDEntries }

//BuildNoMDEntries returns a new NoMDEntries initialized with val
func BuildNoMDEntries(val int) NoMDEntries {
	var field NoMDEntries
	field.Value = val
	return field
}

//NoMDEntryTypes is a NUMINGROUP field
type NoMDEntryTypes struct{ message.NumInGroupValue }

//Tag returns tag.NoMDEntryTypes (267)
func (f NoMDEntryTypes) Tag() fix.Tag { return tag.NoMDEntryTypes }

//BuildNoMDEntryTypes returns a new NoMDEntryTypes initialized with val
func BuildNoMDEntryTypes(val int) NoMDEntryTypes {
	var field NoMDEntryTypes
	field.Value = val
	return field
}

//NoMDFeedTypes is a NUMINGROUP field
type NoMDFeedTypes struct{ message.NumInGroupValue }

//Tag returns tag.NoMDFeedTypes (1141)
func (f NoMDFeedTypes) Tag() fix.Tag { return tag.NoMDFeedTypes }

//BuildNoMDFeedTypes returns a new NoMDFeedTypes initialized with val
func BuildNoMDFeedTypes(val int) NoMDFeedTypes {
	var field NoMDFeedTypes
	field.Value = val
	return field
}

//NoMarketSegments is a NUMINGROUP field
type NoMarketSegments struct{ message.NumInGroupValue }

//Tag returns tag.NoMarketSegments (1310)
func (f NoMarketSegments) Tag() fix.Tag { return tag.NoMarketSegments }

//BuildNoMarketSegments returns a new NoMarketSegments initialized with val
func BuildNoMarketSegments(val int) NoMarketSegments {
	var field NoMarketSegments
	field.Value = val
	return field
}

//NoMatchRules is a NUMINGROUP field
type NoMatchRules struct{ message.NumInGroupValue }

//Tag returns tag.NoMatchRules (1235)
func (f NoMatchRules) Tag() fix.Tag { return tag.NoMatchRules }

//BuildNoMatchRules returns a new NoMatchRules initialized with val
func BuildNoMatchRules(val int) NoMatchRules {
	var field NoMatchRules
	field.Value = val
	return field
}

//NoMaturityRules is a NUMINGROUP field
type NoMaturityRules struct{ message.NumInGroupValue }

//Tag returns tag.NoMaturityRules (1236)
func (f NoMaturityRules) Tag() fix.Tag { return tag.NoMaturityRules }

//BuildNoMaturityRules returns a new NoMaturityRules initialized with val
func BuildNoMaturityRules(val int) NoMaturityRules {
	var field NoMaturityRules
	field.Value = val
	return field
}

//NoMiscFees is a NUMINGROUP field
type NoMiscFees struct{ message.NumInGroupValue }

//Tag returns tag.NoMiscFees (136)
func (f NoMiscFees) Tag() fix.Tag { return tag.NoMiscFees }

//BuildNoMiscFees returns a new NoMiscFees initialized with val
func BuildNoMiscFees(val int) NoMiscFees {
	var field NoMiscFees
	field.Value = val
	return field
}

//NoMsgTypes is a NUMINGROUP field
type NoMsgTypes struct{ message.NumInGroupValue }

//Tag returns tag.NoMsgTypes (384)
func (f NoMsgTypes) Tag() fix.Tag { return tag.NoMsgTypes }

//BuildNoMsgTypes returns a new NoMsgTypes initialized with val
func BuildNoMsgTypes(val int) NoMsgTypes {
	var field NoMsgTypes
	field.Value = val
	return field
}

//NoNested2PartyIDs is a NUMINGROUP field
type NoNested2PartyIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoNested2PartyIDs (756)
func (f NoNested2PartyIDs) Tag() fix.Tag { return tag.NoNested2PartyIDs }

//BuildNoNested2PartyIDs returns a new NoNested2PartyIDs initialized with val
func BuildNoNested2PartyIDs(val int) NoNested2PartyIDs {
	var field NoNested2PartyIDs
	field.Value = val
	return field
}

//NoNested2PartySubIDs is a NUMINGROUP field
type NoNested2PartySubIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoNested2PartySubIDs (806)
func (f NoNested2PartySubIDs) Tag() fix.Tag { return tag.NoNested2PartySubIDs }

//BuildNoNested2PartySubIDs returns a new NoNested2PartySubIDs initialized with val
func BuildNoNested2PartySubIDs(val int) NoNested2PartySubIDs {
	var field NoNested2PartySubIDs
	field.Value = val
	return field
}

//NoNested3PartyIDs is a NUMINGROUP field
type NoNested3PartyIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoNested3PartyIDs (948)
func (f NoNested3PartyIDs) Tag() fix.Tag { return tag.NoNested3PartyIDs }

//BuildNoNested3PartyIDs returns a new NoNested3PartyIDs initialized with val
func BuildNoNested3PartyIDs(val int) NoNested3PartyIDs {
	var field NoNested3PartyIDs
	field.Value = val
	return field
}

//NoNested3PartySubIDs is a NUMINGROUP field
type NoNested3PartySubIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoNested3PartySubIDs (952)
func (f NoNested3PartySubIDs) Tag() fix.Tag { return tag.NoNested3PartySubIDs }

//BuildNoNested3PartySubIDs returns a new NoNested3PartySubIDs initialized with val
func BuildNoNested3PartySubIDs(val int) NoNested3PartySubIDs {
	var field NoNested3PartySubIDs
	field.Value = val
	return field
}

//NoNested4PartyIDs is a NUMINGROUP field
type NoNested4PartyIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoNested4PartyIDs (1414)
func (f NoNested4PartyIDs) Tag() fix.Tag { return tag.NoNested4PartyIDs }

//BuildNoNested4PartyIDs returns a new NoNested4PartyIDs initialized with val
func BuildNoNested4PartyIDs(val int) NoNested4PartyIDs {
	var field NoNested4PartyIDs
	field.Value = val
	return field
}

//NoNested4PartySubIDs is a NUMINGROUP field
type NoNested4PartySubIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoNested4PartySubIDs (1413)
func (f NoNested4PartySubIDs) Tag() fix.Tag { return tag.NoNested4PartySubIDs }

//BuildNoNested4PartySubIDs returns a new NoNested4PartySubIDs initialized with val
func BuildNoNested4PartySubIDs(val int) NoNested4PartySubIDs {
	var field NoNested4PartySubIDs
	field.Value = val
	return field
}

//NoNestedInstrAttrib is a NUMINGROUP field
type NoNestedInstrAttrib struct{ message.NumInGroupValue }

//Tag returns tag.NoNestedInstrAttrib (1312)
func (f NoNestedInstrAttrib) Tag() fix.Tag { return tag.NoNestedInstrAttrib }

//BuildNoNestedInstrAttrib returns a new NoNestedInstrAttrib initialized with val
func BuildNoNestedInstrAttrib(val int) NoNestedInstrAttrib {
	var field NoNestedInstrAttrib
	field.Value = val
	return field
}

//NoNestedPartyIDs is a NUMINGROUP field
type NoNestedPartyIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoNestedPartyIDs (539)
func (f NoNestedPartyIDs) Tag() fix.Tag { return tag.NoNestedPartyIDs }

//BuildNoNestedPartyIDs returns a new NoNestedPartyIDs initialized with val
func BuildNoNestedPartyIDs(val int) NoNestedPartyIDs {
	var field NoNestedPartyIDs
	field.Value = val
	return field
}

//NoNestedPartySubIDs is a NUMINGROUP field
type NoNestedPartySubIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoNestedPartySubIDs (804)
func (f NoNestedPartySubIDs) Tag() fix.Tag { return tag.NoNestedPartySubIDs }

//BuildNoNestedPartySubIDs returns a new NoNestedPartySubIDs initialized with val
func BuildNoNestedPartySubIDs(val int) NoNestedPartySubIDs {
	var field NoNestedPartySubIDs
	field.Value = val
	return field
}

//NoNewsRefIDs is a NUMINGROUP field
type NoNewsRefIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoNewsRefIDs (1475)
func (f NoNewsRefIDs) Tag() fix.Tag { return tag.NoNewsRefIDs }

//BuildNoNewsRefIDs returns a new NoNewsRefIDs initialized with val
func BuildNoNewsRefIDs(val int) NoNewsRefIDs {
	var field NoNewsRefIDs
	field.Value = val
	return field
}

//NoNotAffectedOrders is a NUMINGROUP field
type NoNotAffectedOrders struct{ message.NumInGroupValue }

//Tag returns tag.NoNotAffectedOrders (1370)
func (f NoNotAffectedOrders) Tag() fix.Tag { return tag.NoNotAffectedOrders }

//BuildNoNotAffectedOrders returns a new NoNotAffectedOrders initialized with val
func BuildNoNotAffectedOrders(val int) NoNotAffectedOrders {
	var field NoNotAffectedOrders
	field.Value = val
	return field
}

//NoOfLegUnderlyings is a NUMINGROUP field
type NoOfLegUnderlyings struct{ message.NumInGroupValue }

//Tag returns tag.NoOfLegUnderlyings (1342)
func (f NoOfLegUnderlyings) Tag() fix.Tag { return tag.NoOfLegUnderlyings }

//BuildNoOfLegUnderlyings returns a new NoOfLegUnderlyings initialized with val
func BuildNoOfLegUnderlyings(val int) NoOfLegUnderlyings {
	var field NoOfLegUnderlyings
	field.Value = val
	return field
}

//NoOfSecSizes is a NUMINGROUP field
type NoOfSecSizes struct{ message.NumInGroupValue }

//Tag returns tag.NoOfSecSizes (1177)
func (f NoOfSecSizes) Tag() fix.Tag { return tag.NoOfSecSizes }

//BuildNoOfSecSizes returns a new NoOfSecSizes initialized with val
func BuildNoOfSecSizes(val int) NoOfSecSizes {
	var field NoOfSecSizes
	field.Value = val
	return field
}

//NoOrdTypeRules is a NUMINGROUP field
type NoOrdTypeRules struct{ message.NumInGroupValue }

//Tag returns tag.NoOrdTypeRules (1237)
func (f NoOrdTypeRules) Tag() fix.Tag { return tag.NoOrdTypeRules }

//BuildNoOrdTypeRules returns a new NoOrdTypeRules initialized with val
func BuildNoOrdTypeRules(val int) NoOrdTypeRules {
	var field NoOrdTypeRules
	field.Value = val
	return field
}

//NoOrders is a NUMINGROUP field
type NoOrders struct{ message.NumInGroupValue }

//Tag returns tag.NoOrders (73)
func (f NoOrders) Tag() fix.Tag { return tag.NoOrders }

//BuildNoOrders returns a new NoOrders initialized with val
func BuildNoOrders(val int) NoOrders {
	var field NoOrders
	field.Value = val
	return field
}

//NoPartyAltIDs is a NUMINGROUP field
type NoPartyAltIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoPartyAltIDs (1516)
func (f NoPartyAltIDs) Tag() fix.Tag { return tag.NoPartyAltIDs }

//BuildNoPartyAltIDs returns a new NoPartyAltIDs initialized with val
func BuildNoPartyAltIDs(val int) NoPartyAltIDs {
	var field NoPartyAltIDs
	field.Value = val
	return field
}

//NoPartyAltSubIDs is a NUMINGROUP field
type NoPartyAltSubIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoPartyAltSubIDs (1519)
func (f NoPartyAltSubIDs) Tag() fix.Tag { return tag.NoPartyAltSubIDs }

//BuildNoPartyAltSubIDs returns a new NoPartyAltSubIDs initialized with val
func BuildNoPartyAltSubIDs(val int) NoPartyAltSubIDs {
	var field NoPartyAltSubIDs
	field.Value = val
	return field
}

//NoPartyIDs is a NUMINGROUP field
type NoPartyIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoPartyIDs (453)
func (f NoPartyIDs) Tag() fix.Tag { return tag.NoPartyIDs }

//BuildNoPartyIDs returns a new NoPartyIDs initialized with val
func BuildNoPartyIDs(val int) NoPartyIDs {
	var field NoPartyIDs
	field.Value = val
	return field
}

//NoPartyList is a NUMINGROUP field
type NoPartyList struct{ message.NumInGroupValue }

//Tag returns tag.NoPartyList (1513)
func (f NoPartyList) Tag() fix.Tag { return tag.NoPartyList }

//BuildNoPartyList returns a new NoPartyList initialized with val
func BuildNoPartyList(val int) NoPartyList {
	var field NoPartyList
	field.Value = val
	return field
}

//NoPartyListResponseTypes is a NUMINGROUP field
type NoPartyListResponseTypes struct{ message.NumInGroupValue }

//Tag returns tag.NoPartyListResponseTypes (1506)
func (f NoPartyListResponseTypes) Tag() fix.Tag { return tag.NoPartyListResponseTypes }

//BuildNoPartyListResponseTypes returns a new NoPartyListResponseTypes initialized with val
func BuildNoPartyListResponseTypes(val int) NoPartyListResponseTypes {
	var field NoPartyListResponseTypes
	field.Value = val
	return field
}

//NoPartyRelationships is a NUMINGROUP field
type NoPartyRelationships struct{ message.NumInGroupValue }

//Tag returns tag.NoPartyRelationships (1514)
func (f NoPartyRelationships) Tag() fix.Tag { return tag.NoPartyRelationships }

//BuildNoPartyRelationships returns a new NoPartyRelationships initialized with val
func BuildNoPartyRelationships(val int) NoPartyRelationships {
	var field NoPartyRelationships
	field.Value = val
	return field
}

//NoPartySubIDs is a NUMINGROUP field
type NoPartySubIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoPartySubIDs (802)
func (f NoPartySubIDs) Tag() fix.Tag { return tag.NoPartySubIDs }

//BuildNoPartySubIDs returns a new NoPartySubIDs initialized with val
func BuildNoPartySubIDs(val int) NoPartySubIDs {
	var field NoPartySubIDs
	field.Value = val
	return field
}

//NoPosAmt is a NUMINGROUP field
type NoPosAmt struct{ message.NumInGroupValue }

//Tag returns tag.NoPosAmt (753)
func (f NoPosAmt) Tag() fix.Tag { return tag.NoPosAmt }

//BuildNoPosAmt returns a new NoPosAmt initialized with val
func BuildNoPosAmt(val int) NoPosAmt {
	var field NoPosAmt
	field.Value = val
	return field
}

//NoPositions is a NUMINGROUP field
type NoPositions struct{ message.NumInGroupValue }

//Tag returns tag.NoPositions (702)
func (f NoPositions) Tag() fix.Tag { return tag.NoPositions }

//BuildNoPositions returns a new NoPositions initialized with val
func BuildNoPositions(val int) NoPositions {
	var field NoPositions
	field.Value = val
	return field
}

//NoQuoteEntries is a NUMINGROUP field
type NoQuoteEntries struct{ message.NumInGroupValue }

//Tag returns tag.NoQuoteEntries (295)
func (f NoQuoteEntries) Tag() fix.Tag { return tag.NoQuoteEntries }

//BuildNoQuoteEntries returns a new NoQuoteEntries initialized with val
func BuildNoQuoteEntries(val int) NoQuoteEntries {
	var field NoQuoteEntries
	field.Value = val
	return field
}

//NoQuoteQualifiers is a NUMINGROUP field
type NoQuoteQualifiers struct{ message.NumInGroupValue }

//Tag returns tag.NoQuoteQualifiers (735)
func (f NoQuoteQualifiers) Tag() fix.Tag { return tag.NoQuoteQualifiers }

//BuildNoQuoteQualifiers returns a new NoQuoteQualifiers initialized with val
func BuildNoQuoteQualifiers(val int) NoQuoteQualifiers {
	var field NoQuoteQualifiers
	field.Value = val
	return field
}

//NoQuoteSets is a NUMINGROUP field
type NoQuoteSets struct{ message.NumInGroupValue }

//Tag returns tag.NoQuoteSets (296)
func (f NoQuoteSets) Tag() fix.Tag { return tag.NoQuoteSets }

//BuildNoQuoteSets returns a new NoQuoteSets initialized with val
func BuildNoQuoteSets(val int) NoQuoteSets {
	var field NoQuoteSets
	field.Value = val
	return field
}

//NoRateSources is a NUMINGROUP field
type NoRateSources struct{ message.NumInGroupValue }

//Tag returns tag.NoRateSources (1445)
func (f NoRateSources) Tag() fix.Tag { return tag.NoRateSources }

//BuildNoRateSources returns a new NoRateSources initialized with val
func BuildNoRateSources(val int) NoRateSources {
	var field NoRateSources
	field.Value = val
	return field
}

//NoRegistDtls is a NUMINGROUP field
type NoRegistDtls struct{ message.NumInGroupValue }

//Tag returns tag.NoRegistDtls (473)
func (f NoRegistDtls) Tag() fix.Tag { return tag.NoRegistDtls }

//BuildNoRegistDtls returns a new NoRegistDtls initialized with val
func BuildNoRegistDtls(val int) NoRegistDtls {
	var field NoRegistDtls
	field.Value = val
	return field
}

//NoRelatedContextPartyIDs is a NUMINGROUP field
type NoRelatedContextPartyIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoRelatedContextPartyIDs (1575)
func (f NoRelatedContextPartyIDs) Tag() fix.Tag { return tag.NoRelatedContextPartyIDs }

//BuildNoRelatedContextPartyIDs returns a new NoRelatedContextPartyIDs initialized with val
func BuildNoRelatedContextPartyIDs(val int) NoRelatedContextPartyIDs {
	var field NoRelatedContextPartyIDs
	field.Value = val
	return field
}

//NoRelatedContextPartySubIDs is a NUMINGROUP field
type NoRelatedContextPartySubIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoRelatedContextPartySubIDs (1579)
func (f NoRelatedContextPartySubIDs) Tag() fix.Tag { return tag.NoRelatedContextPartySubIDs }

//BuildNoRelatedContextPartySubIDs returns a new NoRelatedContextPartySubIDs initialized with val
func BuildNoRelatedContextPartySubIDs(val int) NoRelatedContextPartySubIDs {
	var field NoRelatedContextPartySubIDs
	field.Value = val
	return field
}

//NoRelatedPartyAltIDs is a NUMINGROUP field
type NoRelatedPartyAltIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoRelatedPartyAltIDs (1569)
func (f NoRelatedPartyAltIDs) Tag() fix.Tag { return tag.NoRelatedPartyAltIDs }

//BuildNoRelatedPartyAltIDs returns a new NoRelatedPartyAltIDs initialized with val
func BuildNoRelatedPartyAltIDs(val int) NoRelatedPartyAltIDs {
	var field NoRelatedPartyAltIDs
	field.Value = val
	return field
}

//NoRelatedPartyAltSubIDs is a NUMINGROUP field
type NoRelatedPartyAltSubIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoRelatedPartyAltSubIDs (1572)
func (f NoRelatedPartyAltSubIDs) Tag() fix.Tag { return tag.NoRelatedPartyAltSubIDs }

//BuildNoRelatedPartyAltSubIDs returns a new NoRelatedPartyAltSubIDs initialized with val
func BuildNoRelatedPartyAltSubIDs(val int) NoRelatedPartyAltSubIDs {
	var field NoRelatedPartyAltSubIDs
	field.Value = val
	return field
}

//NoRelatedPartyIDs is a NUMINGROUP field
type NoRelatedPartyIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoRelatedPartyIDs (1562)
func (f NoRelatedPartyIDs) Tag() fix.Tag { return tag.NoRelatedPartyIDs }

//BuildNoRelatedPartyIDs returns a new NoRelatedPartyIDs initialized with val
func BuildNoRelatedPartyIDs(val int) NoRelatedPartyIDs {
	var field NoRelatedPartyIDs
	field.Value = val
	return field
}

//NoRelatedPartySubIDs is a NUMINGROUP field
type NoRelatedPartySubIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoRelatedPartySubIDs (1566)
func (f NoRelatedPartySubIDs) Tag() fix.Tag { return tag.NoRelatedPartySubIDs }

//BuildNoRelatedPartySubIDs returns a new NoRelatedPartySubIDs initialized with val
func BuildNoRelatedPartySubIDs(val int) NoRelatedPartySubIDs {
	var field NoRelatedPartySubIDs
	field.Value = val
	return field
}

//NoRelatedSym is a NUMINGROUP field
type NoRelatedSym struct{ message.NumInGroupValue }

//Tag returns tag.NoRelatedSym (146)
func (f NoRelatedSym) Tag() fix.Tag { return tag.NoRelatedSym }

//BuildNoRelatedSym returns a new NoRelatedSym initialized with val
func BuildNoRelatedSym(val int) NoRelatedSym {
	var field NoRelatedSym
	field.Value = val
	return field
}

//NoRelationshipRiskInstruments is a NUMINGROUP field
type NoRelationshipRiskInstruments struct{ message.NumInGroupValue }

//Tag returns tag.NoRelationshipRiskInstruments (1587)
func (f NoRelationshipRiskInstruments) Tag() fix.Tag { return tag.NoRelationshipRiskInstruments }

//BuildNoRelationshipRiskInstruments returns a new NoRelationshipRiskInstruments initialized with val
func BuildNoRelationshipRiskInstruments(val int) NoRelationshipRiskInstruments {
	var field NoRelationshipRiskInstruments
	field.Value = val
	return field
}

//NoRelationshipRiskLimits is a NUMINGROUP field
type NoRelationshipRiskLimits struct{ message.NumInGroupValue }

//Tag returns tag.NoRelationshipRiskLimits (1582)
func (f NoRelationshipRiskLimits) Tag() fix.Tag { return tag.NoRelationshipRiskLimits }

//BuildNoRelationshipRiskLimits returns a new NoRelationshipRiskLimits initialized with val
func BuildNoRelationshipRiskLimits(val int) NoRelationshipRiskLimits {
	var field NoRelationshipRiskLimits
	field.Value = val
	return field
}

//NoRelationshipRiskSecurityAltID is a NUMINGROUP field
type NoRelationshipRiskSecurityAltID struct{ message.NumInGroupValue }

//Tag returns tag.NoRelationshipRiskSecurityAltID (1593)
func (f NoRelationshipRiskSecurityAltID) Tag() fix.Tag { return tag.NoRelationshipRiskSecurityAltID }

//BuildNoRelationshipRiskSecurityAltID returns a new NoRelationshipRiskSecurityAltID initialized with val
func BuildNoRelationshipRiskSecurityAltID(val int) NoRelationshipRiskSecurityAltID {
	var field NoRelationshipRiskSecurityAltID
	field.Value = val
	return field
}

//NoRelationshipRiskWarningLevels is a NUMINGROUP field
type NoRelationshipRiskWarningLevels struct{ message.NumInGroupValue }

//Tag returns tag.NoRelationshipRiskWarningLevels (1613)
func (f NoRelationshipRiskWarningLevels) Tag() fix.Tag { return tag.NoRelationshipRiskWarningLevels }

//BuildNoRelationshipRiskWarningLevels returns a new NoRelationshipRiskWarningLevels initialized with val
func BuildNoRelationshipRiskWarningLevels(val int) NoRelationshipRiskWarningLevels {
	var field NoRelationshipRiskWarningLevels
	field.Value = val
	return field
}

//NoRequestedPartyRoles is a NUMINGROUP field
type NoRequestedPartyRoles struct{ message.NumInGroupValue }

//Tag returns tag.NoRequestedPartyRoles (1508)
func (f NoRequestedPartyRoles) Tag() fix.Tag { return tag.NoRequestedPartyRoles }

//BuildNoRequestedPartyRoles returns a new NoRequestedPartyRoles initialized with val
func BuildNoRequestedPartyRoles(val int) NoRequestedPartyRoles {
	var field NoRequestedPartyRoles
	field.Value = val
	return field
}

//NoRiskInstruments is a NUMINGROUP field
type NoRiskInstruments struct{ message.NumInGroupValue }

//Tag returns tag.NoRiskInstruments (1534)
func (f NoRiskInstruments) Tag() fix.Tag { return tag.NoRiskInstruments }

//BuildNoRiskInstruments returns a new NoRiskInstruments initialized with val
func BuildNoRiskInstruments(val int) NoRiskInstruments {
	var field NoRiskInstruments
	field.Value = val
	return field
}

//NoRiskLimits is a NUMINGROUP field
type NoRiskLimits struct{ message.NumInGroupValue }

//Tag returns tag.NoRiskLimits (1529)
func (f NoRiskLimits) Tag() fix.Tag { return tag.NoRiskLimits }

//BuildNoRiskLimits returns a new NoRiskLimits initialized with val
func BuildNoRiskLimits(val int) NoRiskLimits {
	var field NoRiskLimits
	field.Value = val
	return field
}

//NoRiskSecurityAltID is a NUMINGROUP field
type NoRiskSecurityAltID struct{ message.NumInGroupValue }

//Tag returns tag.NoRiskSecurityAltID (1540)
func (f NoRiskSecurityAltID) Tag() fix.Tag { return tag.NoRiskSecurityAltID }

//BuildNoRiskSecurityAltID returns a new NoRiskSecurityAltID initialized with val
func BuildNoRiskSecurityAltID(val int) NoRiskSecurityAltID {
	var field NoRiskSecurityAltID
	field.Value = val
	return field
}

//NoRiskWarningLevels is a NUMINGROUP field
type NoRiskWarningLevels struct{ message.NumInGroupValue }

//Tag returns tag.NoRiskWarningLevels (1559)
func (f NoRiskWarningLevels) Tag() fix.Tag { return tag.NoRiskWarningLevels }

//BuildNoRiskWarningLevels returns a new NoRiskWarningLevels initialized with val
func BuildNoRiskWarningLevels(val int) NoRiskWarningLevels {
	var field NoRiskWarningLevels
	field.Value = val
	return field
}

//NoRootPartyIDs is a NUMINGROUP field
type NoRootPartyIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoRootPartyIDs (1116)
func (f NoRootPartyIDs) Tag() fix.Tag { return tag.NoRootPartyIDs }

//BuildNoRootPartyIDs returns a new NoRootPartyIDs initialized with val
func BuildNoRootPartyIDs(val int) NoRootPartyIDs {
	var field NoRootPartyIDs
	field.Value = val
	return field
}

//NoRootPartySubIDs is a NUMINGROUP field
type NoRootPartySubIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoRootPartySubIDs (1120)
func (f NoRootPartySubIDs) Tag() fix.Tag { return tag.NoRootPartySubIDs }

//BuildNoRootPartySubIDs returns a new NoRootPartySubIDs initialized with val
func BuildNoRootPartySubIDs(val int) NoRootPartySubIDs {
	var field NoRootPartySubIDs
	field.Value = val
	return field
}

//NoRoutingIDs is a NUMINGROUP field
type NoRoutingIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoRoutingIDs (215)
func (f NoRoutingIDs) Tag() fix.Tag { return tag.NoRoutingIDs }

//BuildNoRoutingIDs returns a new NoRoutingIDs initialized with val
func BuildNoRoutingIDs(val int) NoRoutingIDs {
	var field NoRoutingIDs
	field.Value = val
	return field
}

//NoRpts is a INT field
type NoRpts struct{ message.IntValue }

//Tag returns tag.NoRpts (82)
func (f NoRpts) Tag() fix.Tag { return tag.NoRpts }

//BuildNoRpts returns a new NoRpts initialized with val
func BuildNoRpts(val int) NoRpts {
	var field NoRpts
	field.Value = val
	return field
}

//NoSecurityAltID is a NUMINGROUP field
type NoSecurityAltID struct{ message.NumInGroupValue }

//Tag returns tag.NoSecurityAltID (454)
func (f NoSecurityAltID) Tag() fix.Tag { return tag.NoSecurityAltID }

//BuildNoSecurityAltID returns a new NoSecurityAltID initialized with val
func BuildNoSecurityAltID(val int) NoSecurityAltID {
	var field NoSecurityAltID
	field.Value = val
	return field
}

//NoSecurityTypes is a NUMINGROUP field
type NoSecurityTypes struct{ message.NumInGroupValue }

//Tag returns tag.NoSecurityTypes (558)
func (f NoSecurityTypes) Tag() fix.Tag { return tag.NoSecurityTypes }

//BuildNoSecurityTypes returns a new NoSecurityTypes initialized with val
func BuildNoSecurityTypes(val int) NoSecurityTypes {
	var field NoSecurityTypes
	field.Value = val
	return field
}

//NoSettlDetails is a NUMINGROUP field
type NoSettlDetails struct{ message.NumInGroupValue }

//Tag returns tag.NoSettlDetails (1158)
func (f NoSettlDetails) Tag() fix.Tag { return tag.NoSettlDetails }

//BuildNoSettlDetails returns a new NoSettlDetails initialized with val
func BuildNoSettlDetails(val int) NoSettlDetails {
	var field NoSettlDetails
	field.Value = val
	return field
}

//NoSettlInst is a NUMINGROUP field
type NoSettlInst struct{ message.NumInGroupValue }

//Tag returns tag.NoSettlInst (778)
func (f NoSettlInst) Tag() fix.Tag { return tag.NoSettlInst }

//BuildNoSettlInst returns a new NoSettlInst initialized with val
func BuildNoSettlInst(val int) NoSettlInst {
	var field NoSettlInst
	field.Value = val
	return field
}

//NoSettlOblig is a NUMINGROUP field
type NoSettlOblig struct{ message.NumInGroupValue }

//Tag returns tag.NoSettlOblig (1165)
func (f NoSettlOblig) Tag() fix.Tag { return tag.NoSettlOblig }

//BuildNoSettlOblig returns a new NoSettlOblig initialized with val
func BuildNoSettlOblig(val int) NoSettlOblig {
	var field NoSettlOblig
	field.Value = val
	return field
}

//NoSettlPartyIDs is a NUMINGROUP field
type NoSettlPartyIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoSettlPartyIDs (781)
func (f NoSettlPartyIDs) Tag() fix.Tag { return tag.NoSettlPartyIDs }

//BuildNoSettlPartyIDs returns a new NoSettlPartyIDs initialized with val
func BuildNoSettlPartyIDs(val int) NoSettlPartyIDs {
	var field NoSettlPartyIDs
	field.Value = val
	return field
}

//NoSettlPartySubIDs is a NUMINGROUP field
type NoSettlPartySubIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoSettlPartySubIDs (801)
func (f NoSettlPartySubIDs) Tag() fix.Tag { return tag.NoSettlPartySubIDs }

//BuildNoSettlPartySubIDs returns a new NoSettlPartySubIDs initialized with val
func BuildNoSettlPartySubIDs(val int) NoSettlPartySubIDs {
	var field NoSettlPartySubIDs
	field.Value = val
	return field
}

//NoSideTrdRegTS is a NUMINGROUP field
type NoSideTrdRegTS struct{ message.NumInGroupValue }

//Tag returns tag.NoSideTrdRegTS (1016)
func (f NoSideTrdRegTS) Tag() fix.Tag { return tag.NoSideTrdRegTS }

//BuildNoSideTrdRegTS returns a new NoSideTrdRegTS initialized with val
func BuildNoSideTrdRegTS(val int) NoSideTrdRegTS {
	var field NoSideTrdRegTS
	field.Value = val
	return field
}

//NoSides is a NUMINGROUP field
type NoSides struct{ message.NumInGroupValue }

//Tag returns tag.NoSides (552)
func (f NoSides) Tag() fix.Tag { return tag.NoSides }

//BuildNoSides returns a new NoSides initialized with val
func BuildNoSides(val int) NoSides {
	var field NoSides
	field.Value = val
	return field
}

//NoStatsIndicators is a NUMINGROUP field
type NoStatsIndicators struct{ message.NumInGroupValue }

//Tag returns tag.NoStatsIndicators (1175)
func (f NoStatsIndicators) Tag() fix.Tag { return tag.NoStatsIndicators }

//BuildNoStatsIndicators returns a new NoStatsIndicators initialized with val
func BuildNoStatsIndicators(val int) NoStatsIndicators {
	var field NoStatsIndicators
	field.Value = val
	return field
}

//NoStipulations is a NUMINGROUP field
type NoStipulations struct{ message.NumInGroupValue }

//Tag returns tag.NoStipulations (232)
func (f NoStipulations) Tag() fix.Tag { return tag.NoStipulations }

//BuildNoStipulations returns a new NoStipulations initialized with val
func BuildNoStipulations(val int) NoStipulations {
	var field NoStipulations
	field.Value = val
	return field
}

//NoStrategyParameters is a NUMINGROUP field
type NoStrategyParameters struct{ message.NumInGroupValue }

//Tag returns tag.NoStrategyParameters (957)
func (f NoStrategyParameters) Tag() fix.Tag { return tag.NoStrategyParameters }

//BuildNoStrategyParameters returns a new NoStrategyParameters initialized with val
func BuildNoStrategyParameters(val int) NoStrategyParameters {
	var field NoStrategyParameters
	field.Value = val
	return field
}

//NoStrikeRules is a NUMINGROUP field
type NoStrikeRules struct{ message.NumInGroupValue }

//Tag returns tag.NoStrikeRules (1201)
func (f NoStrikeRules) Tag() fix.Tag { return tag.NoStrikeRules }

//BuildNoStrikeRules returns a new NoStrikeRules initialized with val
func BuildNoStrikeRules(val int) NoStrikeRules {
	var field NoStrikeRules
	field.Value = val
	return field
}

//NoStrikes is a NUMINGROUP field
type NoStrikes struct{ message.NumInGroupValue }

//Tag returns tag.NoStrikes (428)
func (f NoStrikes) Tag() fix.Tag { return tag.NoStrikes }

//BuildNoStrikes returns a new NoStrikes initialized with val
func BuildNoStrikes(val int) NoStrikes {
	var field NoStrikes
	field.Value = val
	return field
}

//NoTargetPartyIDs is a NUMINGROUP field
type NoTargetPartyIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoTargetPartyIDs (1461)
func (f NoTargetPartyIDs) Tag() fix.Tag { return tag.NoTargetPartyIDs }

//BuildNoTargetPartyIDs returns a new NoTargetPartyIDs initialized with val
func BuildNoTargetPartyIDs(val int) NoTargetPartyIDs {
	var field NoTargetPartyIDs
	field.Value = val
	return field
}

//NoTickRules is a NUMINGROUP field
type NoTickRules struct{ message.NumInGroupValue }

//Tag returns tag.NoTickRules (1205)
func (f NoTickRules) Tag() fix.Tag { return tag.NoTickRules }

//BuildNoTickRules returns a new NoTickRules initialized with val
func BuildNoTickRules(val int) NoTickRules {
	var field NoTickRules
	field.Value = val
	return field
}

//NoTimeInForceRules is a NUMINGROUP field
type NoTimeInForceRules struct{ message.NumInGroupValue }

//Tag returns tag.NoTimeInForceRules (1239)
func (f NoTimeInForceRules) Tag() fix.Tag { return tag.NoTimeInForceRules }

//BuildNoTimeInForceRules returns a new NoTimeInForceRules initialized with val
func BuildNoTimeInForceRules(val int) NoTimeInForceRules {
	var field NoTimeInForceRules
	field.Value = val
	return field
}

//NoTrades is a NUMINGROUP field
type NoTrades struct{ message.NumInGroupValue }

//Tag returns tag.NoTrades (897)
func (f NoTrades) Tag() fix.Tag { return tag.NoTrades }

//BuildNoTrades returns a new NoTrades initialized with val
func BuildNoTrades(val int) NoTrades {
	var field NoTrades
	field.Value = val
	return field
}

//NoTradingSessionRules is a NUMINGROUP field
type NoTradingSessionRules struct{ message.NumInGroupValue }

//Tag returns tag.NoTradingSessionRules (1309)
func (f NoTradingSessionRules) Tag() fix.Tag { return tag.NoTradingSessionRules }

//BuildNoTradingSessionRules returns a new NoTradingSessionRules initialized with val
func BuildNoTradingSessionRules(val int) NoTradingSessionRules {
	var field NoTradingSessionRules
	field.Value = val
	return field
}

//NoTradingSessions is a NUMINGROUP field
type NoTradingSessions struct{ message.NumInGroupValue }

//Tag returns tag.NoTradingSessions (386)
func (f NoTradingSessions) Tag() fix.Tag { return tag.NoTradingSessions }

//BuildNoTradingSessions returns a new NoTradingSessions initialized with val
func BuildNoTradingSessions(val int) NoTradingSessions {
	var field NoTradingSessions
	field.Value = val
	return field
}

//NoTrdRegTimestamps is a NUMINGROUP field
type NoTrdRegTimestamps struct{ message.NumInGroupValue }

//Tag returns tag.NoTrdRegTimestamps (768)
func (f NoTrdRegTimestamps) Tag() fix.Tag { return tag.NoTrdRegTimestamps }

//BuildNoTrdRegTimestamps returns a new NoTrdRegTimestamps initialized with val
func BuildNoTrdRegTimestamps(val int) NoTrdRegTimestamps {
	var field NoTrdRegTimestamps
	field.Value = val
	return field
}

//NoTrdRepIndicators is a NUMINGROUP field
type NoTrdRepIndicators struct{ message.NumInGroupValue }

//Tag returns tag.NoTrdRepIndicators (1387)
func (f NoTrdRepIndicators) Tag() fix.Tag { return tag.NoTrdRepIndicators }

//BuildNoTrdRepIndicators returns a new NoTrdRepIndicators initialized with val
func BuildNoTrdRepIndicators(val int) NoTrdRepIndicators {
	var field NoTrdRepIndicators
	field.Value = val
	return field
}

//NoUnderlyingAmounts is a NUMINGROUP field
type NoUnderlyingAmounts struct{ message.NumInGroupValue }

//Tag returns tag.NoUnderlyingAmounts (984)
func (f NoUnderlyingAmounts) Tag() fix.Tag { return tag.NoUnderlyingAmounts }

//BuildNoUnderlyingAmounts returns a new NoUnderlyingAmounts initialized with val
func BuildNoUnderlyingAmounts(val int) NoUnderlyingAmounts {
	var field NoUnderlyingAmounts
	field.Value = val
	return field
}

//NoUnderlyingLegSecurityAltID is a NUMINGROUP field
type NoUnderlyingLegSecurityAltID struct{ message.NumInGroupValue }

//Tag returns tag.NoUnderlyingLegSecurityAltID (1334)
func (f NoUnderlyingLegSecurityAltID) Tag() fix.Tag { return tag.NoUnderlyingLegSecurityAltID }

//BuildNoUnderlyingLegSecurityAltID returns a new NoUnderlyingLegSecurityAltID initialized with val
func BuildNoUnderlyingLegSecurityAltID(val int) NoUnderlyingLegSecurityAltID {
	var field NoUnderlyingLegSecurityAltID
	field.Value = val
	return field
}

//NoUnderlyingSecurityAltID is a NUMINGROUP field
type NoUnderlyingSecurityAltID struct{ message.NumInGroupValue }

//Tag returns tag.NoUnderlyingSecurityAltID (457)
func (f NoUnderlyingSecurityAltID) Tag() fix.Tag { return tag.NoUnderlyingSecurityAltID }

//BuildNoUnderlyingSecurityAltID returns a new NoUnderlyingSecurityAltID initialized with val
func BuildNoUnderlyingSecurityAltID(val int) NoUnderlyingSecurityAltID {
	var field NoUnderlyingSecurityAltID
	field.Value = val
	return field
}

//NoUnderlyingStips is a NUMINGROUP field
type NoUnderlyingStips struct{ message.NumInGroupValue }

//Tag returns tag.NoUnderlyingStips (887)
func (f NoUnderlyingStips) Tag() fix.Tag { return tag.NoUnderlyingStips }

//BuildNoUnderlyingStips returns a new NoUnderlyingStips initialized with val
func BuildNoUnderlyingStips(val int) NoUnderlyingStips {
	var field NoUnderlyingStips
	field.Value = val
	return field
}

//NoUnderlyings is a NUMINGROUP field
type NoUnderlyings struct{ message.NumInGroupValue }

//Tag returns tag.NoUnderlyings (711)
func (f NoUnderlyings) Tag() fix.Tag { return tag.NoUnderlyings }

//BuildNoUnderlyings returns a new NoUnderlyings initialized with val
func BuildNoUnderlyings(val int) NoUnderlyings {
	var field NoUnderlyings
	field.Value = val
	return field
}

//NoUndlyInstrumentParties is a NUMINGROUP field
type NoUndlyInstrumentParties struct{ message.NumInGroupValue }

//Tag returns tag.NoUndlyInstrumentParties (1058)
func (f NoUndlyInstrumentParties) Tag() fix.Tag { return tag.NoUndlyInstrumentParties }

//BuildNoUndlyInstrumentParties returns a new NoUndlyInstrumentParties initialized with val
func BuildNoUndlyInstrumentParties(val int) NoUndlyInstrumentParties {
	var field NoUndlyInstrumentParties
	field.Value = val
	return field
}

//NoUndlyInstrumentPartySubIDs is a NUMINGROUP field
type NoUndlyInstrumentPartySubIDs struct{ message.NumInGroupValue }

//Tag returns tag.NoUndlyInstrumentPartySubIDs (1062)
func (f NoUndlyInstrumentPartySubIDs) Tag() fix.Tag { return tag.NoUndlyInstrumentPartySubIDs }

//BuildNoUndlyInstrumentPartySubIDs returns a new NoUndlyInstrumentPartySubIDs initialized with val
func BuildNoUndlyInstrumentPartySubIDs(val int) NoUndlyInstrumentPartySubIDs {
	var field NoUndlyInstrumentPartySubIDs
	field.Value = val
	return field
}

//NotAffOrigClOrdID is a STRING field
type NotAffOrigClOrdID struct{ message.StringValue }

//Tag returns tag.NotAffOrigClOrdID (1372)
func (f NotAffOrigClOrdID) Tag() fix.Tag { return tag.NotAffOrigClOrdID }

//BuildNotAffOrigClOrdID returns a new NotAffOrigClOrdID initialized with val
func BuildNotAffOrigClOrdID(val string) NotAffOrigClOrdID {
	var field NotAffOrigClOrdID
	field.Value = val
	return field
}

//NotAffectedOrderID is a STRING field
type NotAffectedOrderID struct{ message.StringValue }

//Tag returns tag.NotAffectedOrderID (1371)
func (f NotAffectedOrderID) Tag() fix.Tag { return tag.NotAffectedOrderID }

//BuildNotAffectedOrderID returns a new NotAffectedOrderID initialized with val
func BuildNotAffectedOrderID(val string) NotAffectedOrderID {
	var field NotAffectedOrderID
	field.Value = val
	return field
}

//NotifyBrokerOfCredit is a BOOLEAN field
type NotifyBrokerOfCredit struct{ message.BooleanValue }

//Tag returns tag.NotifyBrokerOfCredit (208)
func (f NotifyBrokerOfCredit) Tag() fix.Tag { return tag.NotifyBrokerOfCredit }

//BuildNotifyBrokerOfCredit returns a new NotifyBrokerOfCredit initialized with val
func BuildNotifyBrokerOfCredit(val bool) NotifyBrokerOfCredit {
	var field NotifyBrokerOfCredit
	field.Value = val
	return field
}

//NotionalPercentageOutstanding is a PERCENTAGE field
type NotionalPercentageOutstanding struct{ message.PercentageValue }

//Tag returns tag.NotionalPercentageOutstanding (1451)
func (f NotionalPercentageOutstanding) Tag() fix.Tag { return tag.NotionalPercentageOutstanding }

//BuildNotionalPercentageOutstanding returns a new NotionalPercentageOutstanding initialized with val
func BuildNotionalPercentageOutstanding(val float64) NotionalPercentageOutstanding {
	var field NotionalPercentageOutstanding
	field.Value = val
	return field
}

//NumBidders is a INT field
type NumBidders struct{ message.IntValue }

//Tag returns tag.NumBidders (417)
func (f NumBidders) Tag() fix.Tag { return tag.NumBidders }

//BuildNumBidders returns a new NumBidders initialized with val
func BuildNumBidders(val int) NumBidders {
	var field NumBidders
	field.Value = val
	return field
}

//NumDaysInterest is a INT field
type NumDaysInterest struct{ message.IntValue }

//Tag returns tag.NumDaysInterest (157)
func (f NumDaysInterest) Tag() fix.Tag { return tag.NumDaysInterest }

//BuildNumDaysInterest returns a new NumDaysInterest initialized with val
func BuildNumDaysInterest(val int) NumDaysInterest {
	var field NumDaysInterest
	field.Value = val
	return field
}

//NumTickets is a INT field
type NumTickets struct{ message.IntValue }

//Tag returns tag.NumTickets (395)
func (f NumTickets) Tag() fix.Tag { return tag.NumTickets }

//BuildNumTickets returns a new NumTickets initialized with val
func BuildNumTickets(val int) NumTickets {
	var field NumTickets
	field.Value = val
	return field
}

//NumberOfOrders is a INT field
type NumberOfOrders struct{ message.IntValue }

//Tag returns tag.NumberOfOrders (346)
func (f NumberOfOrders) Tag() fix.Tag { return tag.NumberOfOrders }

//BuildNumberOfOrders returns a new NumberOfOrders initialized with val
func BuildNumberOfOrders(val int) NumberOfOrders {
	var field NumberOfOrders
	field.Value = val
	return field
}

//OddLot is a BOOLEAN field
type OddLot struct{ message.BooleanValue }

//Tag returns tag.OddLot (575)
func (f OddLot) Tag() fix.Tag { return tag.OddLot }

//BuildOddLot returns a new OddLot initialized with val
func BuildOddLot(val bool) OddLot {
	var field OddLot
	field.Value = val
	return field
}

//OfferForwardPoints is a PRICEOFFSET field
type OfferForwardPoints struct{ message.PriceOffsetValue }

//Tag returns tag.OfferForwardPoints (191)
func (f OfferForwardPoints) Tag() fix.Tag { return tag.OfferForwardPoints }

//BuildOfferForwardPoints returns a new OfferForwardPoints initialized with val
func BuildOfferForwardPoints(val float64) OfferForwardPoints {
	var field OfferForwardPoints
	field.Value = val
	return field
}

//OfferForwardPoints2 is a PRICEOFFSET field
type OfferForwardPoints2 struct{ message.PriceOffsetValue }

//Tag returns tag.OfferForwardPoints2 (643)
func (f OfferForwardPoints2) Tag() fix.Tag { return tag.OfferForwardPoints2 }

//BuildOfferForwardPoints2 returns a new OfferForwardPoints2 initialized with val
func BuildOfferForwardPoints2(val float64) OfferForwardPoints2 {
	var field OfferForwardPoints2
	field.Value = val
	return field
}

//OfferPx is a PRICE field
type OfferPx struct{ message.PriceValue }

//Tag returns tag.OfferPx (133)
func (f OfferPx) Tag() fix.Tag { return tag.OfferPx }

//BuildOfferPx returns a new OfferPx initialized with val
func BuildOfferPx(val float64) OfferPx {
	var field OfferPx
	field.Value = val
	return field
}

//OfferSize is a QTY field
type OfferSize struct{ message.QtyValue }

//Tag returns tag.OfferSize (135)
func (f OfferSize) Tag() fix.Tag { return tag.OfferSize }

//BuildOfferSize returns a new OfferSize initialized with val
func BuildOfferSize(val float64) OfferSize {
	var field OfferSize
	field.Value = val
	return field
}

//OfferSpotRate is a PRICE field
type OfferSpotRate struct{ message.PriceValue }

//Tag returns tag.OfferSpotRate (190)
func (f OfferSpotRate) Tag() fix.Tag { return tag.OfferSpotRate }

//BuildOfferSpotRate returns a new OfferSpotRate initialized with val
func BuildOfferSpotRate(val float64) OfferSpotRate {
	var field OfferSpotRate
	field.Value = val
	return field
}

//OfferSwapPoints is a PRICEOFFSET field
type OfferSwapPoints struct{ message.PriceOffsetValue }

//Tag returns tag.OfferSwapPoints (1066)
func (f OfferSwapPoints) Tag() fix.Tag { return tag.OfferSwapPoints }

//BuildOfferSwapPoints returns a new OfferSwapPoints initialized with val
func BuildOfferSwapPoints(val float64) OfferSwapPoints {
	var field OfferSwapPoints
	field.Value = val
	return field
}

//OfferYield is a PERCENTAGE field
type OfferYield struct{ message.PercentageValue }

//Tag returns tag.OfferYield (634)
func (f OfferYield) Tag() fix.Tag { return tag.OfferYield }

//BuildOfferYield returns a new OfferYield initialized with val
func BuildOfferYield(val float64) OfferYield {
	var field OfferYield
	field.Value = val
	return field
}

//OnBehalfOfCompID is a STRING field
type OnBehalfOfCompID struct{ message.StringValue }

//Tag returns tag.OnBehalfOfCompID (115)
func (f OnBehalfOfCompID) Tag() fix.Tag { return tag.OnBehalfOfCompID }

//BuildOnBehalfOfCompID returns a new OnBehalfOfCompID initialized with val
func BuildOnBehalfOfCompID(val string) OnBehalfOfCompID {
	var field OnBehalfOfCompID
	field.Value = val
	return field
}

//OnBehalfOfLocationID is a STRING field
type OnBehalfOfLocationID struct{ message.StringValue }

//Tag returns tag.OnBehalfOfLocationID (144)
func (f OnBehalfOfLocationID) Tag() fix.Tag { return tag.OnBehalfOfLocationID }

//BuildOnBehalfOfLocationID returns a new OnBehalfOfLocationID initialized with val
func BuildOnBehalfOfLocationID(val string) OnBehalfOfLocationID {
	var field OnBehalfOfLocationID
	field.Value = val
	return field
}

//OnBehalfOfSendingTime is a UTCTIMESTAMP field
type OnBehalfOfSendingTime struct{ message.UTCTimestampValue }

//Tag returns tag.OnBehalfOfSendingTime (370)
func (f OnBehalfOfSendingTime) Tag() fix.Tag { return tag.OnBehalfOfSendingTime }

//OnBehalfOfSubID is a STRING field
type OnBehalfOfSubID struct{ message.StringValue }

//Tag returns tag.OnBehalfOfSubID (116)
func (f OnBehalfOfSubID) Tag() fix.Tag { return tag.OnBehalfOfSubID }

//BuildOnBehalfOfSubID returns a new OnBehalfOfSubID initialized with val
func BuildOnBehalfOfSubID(val string) OnBehalfOfSubID {
	var field OnBehalfOfSubID
	field.Value = val
	return field
}

//OpenClose is a CHAR field
type OpenClose struct{ message.CharValue }

//Tag returns tag.OpenClose (77)
func (f OpenClose) Tag() fix.Tag { return tag.OpenClose }

//BuildOpenClose returns a new OpenClose initialized with val
func BuildOpenClose(val string) OpenClose {
	var field OpenClose
	field.Value = val
	return field
}

//OpenCloseSettlFlag is a MULTIPLECHARVALUE field
type OpenCloseSettlFlag struct{ message.MultipleCharValue }

//Tag returns tag.OpenCloseSettlFlag (286)
func (f OpenCloseSettlFlag) Tag() fix.Tag { return tag.OpenCloseSettlFlag }

//BuildOpenCloseSettlFlag returns a new OpenCloseSettlFlag initialized with val
func BuildOpenCloseSettlFlag(val string) OpenCloseSettlFlag {
	var field OpenCloseSettlFlag
	field.Value = val
	return field
}

//OpenCloseSettleFlag is a MULTIPLEVALUESTRING field
type OpenCloseSettleFlag struct{ message.MultipleStringValue }

//Tag returns tag.OpenCloseSettleFlag (286)
func (f OpenCloseSettleFlag) Tag() fix.Tag { return tag.OpenCloseSettleFlag }

//BuildOpenCloseSettleFlag returns a new OpenCloseSettleFlag initialized with val
func BuildOpenCloseSettleFlag(val string) OpenCloseSettleFlag {
	var field OpenCloseSettleFlag
	field.Value = val
	return field
}

//OpenInterest is a AMT field
type OpenInterest struct{ message.AmtValue }

//Tag returns tag.OpenInterest (746)
func (f OpenInterest) Tag() fix.Tag { return tag.OpenInterest }

//BuildOpenInterest returns a new OpenInterest initialized with val
func BuildOpenInterest(val float64) OpenInterest {
	var field OpenInterest
	field.Value = val
	return field
}

//OptAttribute is a CHAR field
type OptAttribute struct{ message.CharValue }

//Tag returns tag.OptAttribute (206)
func (f OptAttribute) Tag() fix.Tag { return tag.OptAttribute }

//BuildOptAttribute returns a new OptAttribute initialized with val
func BuildOptAttribute(val string) OptAttribute {
	var field OptAttribute
	field.Value = val
	return field
}

//OptPayAmount is a AMT field
type OptPayAmount struct{ message.AmtValue }

//Tag returns tag.OptPayAmount (1195)
func (f OptPayAmount) Tag() fix.Tag { return tag.OptPayAmount }

//BuildOptPayAmount returns a new OptPayAmount initialized with val
func BuildOptPayAmount(val float64) OptPayAmount {
	var field OptPayAmount
	field.Value = val
	return field
}

//OptPayoutAmount is a AMT field
type OptPayoutAmount struct{ message.AmtValue }

//Tag returns tag.OptPayoutAmount (1195)
func (f OptPayoutAmount) Tag() fix.Tag { return tag.OptPayoutAmount }

//BuildOptPayoutAmount returns a new OptPayoutAmount initialized with val
func BuildOptPayoutAmount(val float64) OptPayoutAmount {
	var field OptPayoutAmount
	field.Value = val
	return field
}

//OptPayoutType is a INT field
type OptPayoutType struct{ message.IntValue }

//Tag returns tag.OptPayoutType (1482)
func (f OptPayoutType) Tag() fix.Tag { return tag.OptPayoutType }

//BuildOptPayoutType returns a new OptPayoutType initialized with val
func BuildOptPayoutType(val int) OptPayoutType {
	var field OptPayoutType
	field.Value = val
	return field
}

//OrdRejReason is a INT field
type OrdRejReason struct{ message.IntValue }

//Tag returns tag.OrdRejReason (103)
func (f OrdRejReason) Tag() fix.Tag { return tag.OrdRejReason }

//BuildOrdRejReason returns a new OrdRejReason initialized with val
func BuildOrdRejReason(val int) OrdRejReason {
	var field OrdRejReason
	field.Value = val
	return field
}

//OrdStatus is a CHAR field
type OrdStatus struct{ message.CharValue }

//Tag returns tag.OrdStatus (39)
func (f OrdStatus) Tag() fix.Tag { return tag.OrdStatus }

//BuildOrdStatus returns a new OrdStatus initialized with val
func BuildOrdStatus(val string) OrdStatus {
	var field OrdStatus
	field.Value = val
	return field
}

//OrdStatusReqID is a STRING field
type OrdStatusReqID struct{ message.StringValue }

//Tag returns tag.OrdStatusReqID (790)
func (f OrdStatusReqID) Tag() fix.Tag { return tag.OrdStatusReqID }

//BuildOrdStatusReqID returns a new OrdStatusReqID initialized with val
func BuildOrdStatusReqID(val string) OrdStatusReqID {
	var field OrdStatusReqID
	field.Value = val
	return field
}

//OrdType is a CHAR field
type OrdType struct{ message.CharValue }

//Tag returns tag.OrdType (40)
func (f OrdType) Tag() fix.Tag { return tag.OrdType }

//BuildOrdType returns a new OrdType initialized with val
func BuildOrdType(val string) OrdType {
	var field OrdType
	field.Value = val
	return field
}

//OrderAvgPx is a PRICE field
type OrderAvgPx struct{ message.PriceValue }

//Tag returns tag.OrderAvgPx (799)
func (f OrderAvgPx) Tag() fix.Tag { return tag.OrderAvgPx }

//BuildOrderAvgPx returns a new OrderAvgPx initialized with val
func BuildOrderAvgPx(val float64) OrderAvgPx {
	var field OrderAvgPx
	field.Value = val
	return field
}

//OrderBookingQty is a QTY field
type OrderBookingQty struct{ message.QtyValue }

//Tag returns tag.OrderBookingQty (800)
func (f OrderBookingQty) Tag() fix.Tag { return tag.OrderBookingQty }

//BuildOrderBookingQty returns a new OrderBookingQty initialized with val
func BuildOrderBookingQty(val float64) OrderBookingQty {
	var field OrderBookingQty
	field.Value = val
	return field
}

//OrderCapacity is a CHAR field
type OrderCapacity struct{ message.CharValue }

//Tag returns tag.OrderCapacity (528)
func (f OrderCapacity) Tag() fix.Tag { return tag.OrderCapacity }

//BuildOrderCapacity returns a new OrderCapacity initialized with val
func BuildOrderCapacity(val string) OrderCapacity {
	var field OrderCapacity
	field.Value = val
	return field
}

//OrderCapacityQty is a QTY field
type OrderCapacityQty struct{ message.QtyValue }

//Tag returns tag.OrderCapacityQty (863)
func (f OrderCapacityQty) Tag() fix.Tag { return tag.OrderCapacityQty }

//BuildOrderCapacityQty returns a new OrderCapacityQty initialized with val
func BuildOrderCapacityQty(val float64) OrderCapacityQty {
	var field OrderCapacityQty
	field.Value = val
	return field
}

//OrderCategory is a CHAR field
type OrderCategory struct{ message.CharValue }

//Tag returns tag.OrderCategory (1115)
func (f OrderCategory) Tag() fix.Tag { return tag.OrderCategory }

//BuildOrderCategory returns a new OrderCategory initialized with val
func BuildOrderCategory(val string) OrderCategory {
	var field OrderCategory
	field.Value = val
	return field
}

//OrderDelay is a INT field
type OrderDelay struct{ message.IntValue }

//Tag returns tag.OrderDelay (1428)
func (f OrderDelay) Tag() fix.Tag { return tag.OrderDelay }

//BuildOrderDelay returns a new OrderDelay initialized with val
func BuildOrderDelay(val int) OrderDelay {
	var field OrderDelay
	field.Value = val
	return field
}

//OrderDelayUnit is a INT field
type OrderDelayUnit struct{ message.IntValue }

//Tag returns tag.OrderDelayUnit (1429)
func (f OrderDelayUnit) Tag() fix.Tag { return tag.OrderDelayUnit }

//BuildOrderDelayUnit returns a new OrderDelayUnit initialized with val
func BuildOrderDelayUnit(val int) OrderDelayUnit {
	var field OrderDelayUnit
	field.Value = val
	return field
}

//OrderHandlingInstSource is a INT field
type OrderHandlingInstSource struct{ message.IntValue }

//Tag returns tag.OrderHandlingInstSource (1032)
func (f OrderHandlingInstSource) Tag() fix.Tag { return tag.OrderHandlingInstSource }

//BuildOrderHandlingInstSource returns a new OrderHandlingInstSource initialized with val
func BuildOrderHandlingInstSource(val int) OrderHandlingInstSource {
	var field OrderHandlingInstSource
	field.Value = val
	return field
}

//OrderID is a STRING field
type OrderID struct{ message.StringValue }

//Tag returns tag.OrderID (37)
func (f OrderID) Tag() fix.Tag { return tag.OrderID }

//BuildOrderID returns a new OrderID initialized with val
func BuildOrderID(val string) OrderID {
	var field OrderID
	field.Value = val
	return field
}

//OrderInputDevice is a STRING field
type OrderInputDevice struct{ message.StringValue }

//Tag returns tag.OrderInputDevice (821)
func (f OrderInputDevice) Tag() fix.Tag { return tag.OrderInputDevice }

//BuildOrderInputDevice returns a new OrderInputDevice initialized with val
func BuildOrderInputDevice(val string) OrderInputDevice {
	var field OrderInputDevice
	field.Value = val
	return field
}

//OrderPercent is a PERCENTAGE field
type OrderPercent struct{ message.PercentageValue }

//Tag returns tag.OrderPercent (516)
func (f OrderPercent) Tag() fix.Tag { return tag.OrderPercent }

//BuildOrderPercent returns a new OrderPercent initialized with val
func BuildOrderPercent(val float64) OrderPercent {
	var field OrderPercent
	field.Value = val
	return field
}

//OrderQty is a QTY field
type OrderQty struct{ message.QtyValue }

//Tag returns tag.OrderQty (38)
func (f OrderQty) Tag() fix.Tag { return tag.OrderQty }

//BuildOrderQty returns a new OrderQty initialized with val
func BuildOrderQty(val float64) OrderQty {
	var field OrderQty
	field.Value = val
	return field
}

//OrderQty2 is a QTY field
type OrderQty2 struct{ message.QtyValue }

//Tag returns tag.OrderQty2 (192)
func (f OrderQty2) Tag() fix.Tag { return tag.OrderQty2 }

//BuildOrderQty2 returns a new OrderQty2 initialized with val
func BuildOrderQty2(val float64) OrderQty2 {
	var field OrderQty2
	field.Value = val
	return field
}

//OrderRestrictions is a MULTIPLECHARVALUE field
type OrderRestrictions struct{ message.MultipleCharValue }

//Tag returns tag.OrderRestrictions (529)
func (f OrderRestrictions) Tag() fix.Tag { return tag.OrderRestrictions }

//BuildOrderRestrictions returns a new OrderRestrictions initialized with val
func BuildOrderRestrictions(val string) OrderRestrictions {
	var field OrderRestrictions
	field.Value = val
	return field
}

//OrigClOrdID is a STRING field
type OrigClOrdID struct{ message.StringValue }

//Tag returns tag.OrigClOrdID (41)
func (f OrigClOrdID) Tag() fix.Tag { return tag.OrigClOrdID }

//BuildOrigClOrdID returns a new OrigClOrdID initialized with val
func BuildOrigClOrdID(val string) OrigClOrdID {
	var field OrigClOrdID
	field.Value = val
	return field
}

//OrigCrossID is a STRING field
type OrigCrossID struct{ message.StringValue }

//Tag returns tag.OrigCrossID (551)
func (f OrigCrossID) Tag() fix.Tag { return tag.OrigCrossID }

//BuildOrigCrossID returns a new OrigCrossID initialized with val
func BuildOrigCrossID(val string) OrigCrossID {
	var field OrigCrossID
	field.Value = val
	return field
}

//OrigCustOrderCapacity is a INT field
type OrigCustOrderCapacity struct{ message.IntValue }

//Tag returns tag.OrigCustOrderCapacity (1432)
func (f OrigCustOrderCapacity) Tag() fix.Tag { return tag.OrigCustOrderCapacity }

//BuildOrigCustOrderCapacity returns a new OrigCustOrderCapacity initialized with val
func BuildOrigCustOrderCapacity(val int) OrigCustOrderCapacity {
	var field OrigCustOrderCapacity
	field.Value = val
	return field
}

//OrigOrdModTime is a UTCTIMESTAMP field
type OrigOrdModTime struct{ message.UTCTimestampValue }

//Tag returns tag.OrigOrdModTime (586)
func (f OrigOrdModTime) Tag() fix.Tag { return tag.OrigOrdModTime }

//OrigPosReqRefID is a STRING field
type OrigPosReqRefID struct{ message.StringValue }

//Tag returns tag.OrigPosReqRefID (713)
func (f OrigPosReqRefID) Tag() fix.Tag { return tag.OrigPosReqRefID }

//BuildOrigPosReqRefID returns a new OrigPosReqRefID initialized with val
func BuildOrigPosReqRefID(val string) OrigPosReqRefID {
	var field OrigPosReqRefID
	field.Value = val
	return field
}

//OrigSecondaryTradeID is a STRING field
type OrigSecondaryTradeID struct{ message.StringValue }

//Tag returns tag.OrigSecondaryTradeID (1127)
func (f OrigSecondaryTradeID) Tag() fix.Tag { return tag.OrigSecondaryTradeID }

//BuildOrigSecondaryTradeID returns a new OrigSecondaryTradeID initialized with val
func BuildOrigSecondaryTradeID(val string) OrigSecondaryTradeID {
	var field OrigSecondaryTradeID
	field.Value = val
	return field
}

//OrigSendingTime is a UTCTIMESTAMP field
type OrigSendingTime struct{ message.UTCTimestampValue }

//Tag returns tag.OrigSendingTime (122)
func (f OrigSendingTime) Tag() fix.Tag { return tag.OrigSendingTime }

//OrigTime is a UTCTIMESTAMP field
type OrigTime struct{ message.UTCTimestampValue }

//Tag returns tag.OrigTime (42)
func (f OrigTime) Tag() fix.Tag { return tag.OrigTime }

//OrigTradeDate is a LOCALMKTDATE field
type OrigTradeDate struct{ message.LocalMktDateValue }

//Tag returns tag.OrigTradeDate (1125)
func (f OrigTradeDate) Tag() fix.Tag { return tag.OrigTradeDate }

//BuildOrigTradeDate returns a new OrigTradeDate initialized with val
func BuildOrigTradeDate(val string) OrigTradeDate {
	var field OrigTradeDate
	field.Value = val
	return field
}

//OrigTradeHandlingInstr is a CHAR field
type OrigTradeHandlingInstr struct{ message.CharValue }

//Tag returns tag.OrigTradeHandlingInstr (1124)
func (f OrigTradeHandlingInstr) Tag() fix.Tag { return tag.OrigTradeHandlingInstr }

//BuildOrigTradeHandlingInstr returns a new OrigTradeHandlingInstr initialized with val
func BuildOrigTradeHandlingInstr(val string) OrigTradeHandlingInstr {
	var field OrigTradeHandlingInstr
	field.Value = val
	return field
}

//OrigTradeID is a STRING field
type OrigTradeID struct{ message.StringValue }

//Tag returns tag.OrigTradeID (1126)
func (f OrigTradeID) Tag() fix.Tag { return tag.OrigTradeID }

//BuildOrigTradeID returns a new OrigTradeID initialized with val
func BuildOrigTradeID(val string) OrigTradeID {
	var field OrigTradeID
	field.Value = val
	return field
}

//OriginalNotionalPercentageOutstanding is a PERCENTAGE field
type OriginalNotionalPercentageOutstanding struct{ message.PercentageValue }

//Tag returns tag.OriginalNotionalPercentageOutstanding (1452)
func (f OriginalNotionalPercentageOutstanding) Tag() fix.Tag {
	return tag.OriginalNotionalPercentageOutstanding
}

//BuildOriginalNotionalPercentageOutstanding returns a new OriginalNotionalPercentageOutstanding initialized with val
func BuildOriginalNotionalPercentageOutstanding(val float64) OriginalNotionalPercentageOutstanding {
	var field OriginalNotionalPercentageOutstanding
	field.Value = val
	return field
}

//OutMainCntryUIndex is a AMT field
type OutMainCntryUIndex struct{ message.AmtValue }

//Tag returns tag.OutMainCntryUIndex (412)
func (f OutMainCntryUIndex) Tag() fix.Tag { return tag.OutMainCntryUIndex }

//BuildOutMainCntryUIndex returns a new OutMainCntryUIndex initialized with val
func BuildOutMainCntryUIndex(val float64) OutMainCntryUIndex {
	var field OutMainCntryUIndex
	field.Value = val
	return field
}

//OutsideIndexPct is a PERCENTAGE field
type OutsideIndexPct struct{ message.PercentageValue }

//Tag returns tag.OutsideIndexPct (407)
func (f OutsideIndexPct) Tag() fix.Tag { return tag.OutsideIndexPct }

//BuildOutsideIndexPct returns a new OutsideIndexPct initialized with val
func BuildOutsideIndexPct(val float64) OutsideIndexPct {
	var field OutsideIndexPct
	field.Value = val
	return field
}

//OwnerType is a INT field
type OwnerType struct{ message.IntValue }

//Tag returns tag.OwnerType (522)
func (f OwnerType) Tag() fix.Tag { return tag.OwnerType }

//BuildOwnerType returns a new OwnerType initialized with val
func BuildOwnerType(val int) OwnerType {
	var field OwnerType
	field.Value = val
	return field
}

//OwnershipType is a CHAR field
type OwnershipType struct{ message.CharValue }

//Tag returns tag.OwnershipType (517)
func (f OwnershipType) Tag() fix.Tag { return tag.OwnershipType }

//BuildOwnershipType returns a new OwnershipType initialized with val
func BuildOwnershipType(val string) OwnershipType {
	var field OwnershipType
	field.Value = val
	return field
}

//ParentMktSegmID is a STRING field
type ParentMktSegmID struct{ message.StringValue }

//Tag returns tag.ParentMktSegmID (1325)
func (f ParentMktSegmID) Tag() fix.Tag { return tag.ParentMktSegmID }

//BuildParentMktSegmID returns a new ParentMktSegmID initialized with val
func BuildParentMktSegmID(val string) ParentMktSegmID {
	var field ParentMktSegmID
	field.Value = val
	return field
}

//ParticipationRate is a PERCENTAGE field
type ParticipationRate struct{ message.PercentageValue }

//Tag returns tag.ParticipationRate (849)
func (f ParticipationRate) Tag() fix.Tag { return tag.ParticipationRate }

//BuildParticipationRate returns a new ParticipationRate initialized with val
func BuildParticipationRate(val float64) ParticipationRate {
	var field ParticipationRate
	field.Value = val
	return field
}

//PartyAltID is a STRING field
type PartyAltID struct{ message.StringValue }

//Tag returns tag.PartyAltID (1517)
func (f PartyAltID) Tag() fix.Tag { return tag.PartyAltID }

//BuildPartyAltID returns a new PartyAltID initialized with val
func BuildPartyAltID(val string) PartyAltID {
	var field PartyAltID
	field.Value = val
	return field
}

//PartyAltIDSource is a CHAR field
type PartyAltIDSource struct{ message.CharValue }

//Tag returns tag.PartyAltIDSource (1518)
func (f PartyAltIDSource) Tag() fix.Tag { return tag.PartyAltIDSource }

//BuildPartyAltIDSource returns a new PartyAltIDSource initialized with val
func BuildPartyAltIDSource(val string) PartyAltIDSource {
	var field PartyAltIDSource
	field.Value = val
	return field
}

//PartyAltSubID is a STRING field
type PartyAltSubID struct{ message.StringValue }

//Tag returns tag.PartyAltSubID (1520)
func (f PartyAltSubID) Tag() fix.Tag { return tag.PartyAltSubID }

//BuildPartyAltSubID returns a new PartyAltSubID initialized with val
func BuildPartyAltSubID(val string) PartyAltSubID {
	var field PartyAltSubID
	field.Value = val
	return field
}

//PartyAltSubIDType is a INT field
type PartyAltSubIDType struct{ message.IntValue }

//Tag returns tag.PartyAltSubIDType (1521)
func (f PartyAltSubIDType) Tag() fix.Tag { return tag.PartyAltSubIDType }

//BuildPartyAltSubIDType returns a new PartyAltSubIDType initialized with val
func BuildPartyAltSubIDType(val int) PartyAltSubIDType {
	var field PartyAltSubIDType
	field.Value = val
	return field
}

//PartyDetailsListReportID is a STRING field
type PartyDetailsListReportID struct{ message.StringValue }

//Tag returns tag.PartyDetailsListReportID (1510)
func (f PartyDetailsListReportID) Tag() fix.Tag { return tag.PartyDetailsListReportID }

//BuildPartyDetailsListReportID returns a new PartyDetailsListReportID initialized with val
func BuildPartyDetailsListReportID(val string) PartyDetailsListReportID {
	var field PartyDetailsListReportID
	field.Value = val
	return field
}

//PartyDetailsListRequestID is a STRING field
type PartyDetailsListRequestID struct{ message.StringValue }

//Tag returns tag.PartyDetailsListRequestID (1505)
func (f PartyDetailsListRequestID) Tag() fix.Tag { return tag.PartyDetailsListRequestID }

//BuildPartyDetailsListRequestID returns a new PartyDetailsListRequestID initialized with val
func BuildPartyDetailsListRequestID(val string) PartyDetailsListRequestID {
	var field PartyDetailsListRequestID
	field.Value = val
	return field
}

//PartyDetailsRequestResult is a INT field
type PartyDetailsRequestResult struct{ message.IntValue }

//Tag returns tag.PartyDetailsRequestResult (1511)
func (f PartyDetailsRequestResult) Tag() fix.Tag { return tag.PartyDetailsRequestResult }

//BuildPartyDetailsRequestResult returns a new PartyDetailsRequestResult initialized with val
func BuildPartyDetailsRequestResult(val int) PartyDetailsRequestResult {
	var field PartyDetailsRequestResult
	field.Value = val
	return field
}

//PartyID is a STRING field
type PartyID struct{ message.StringValue }

//Tag returns tag.PartyID (448)
func (f PartyID) Tag() fix.Tag { return tag.PartyID }

//BuildPartyID returns a new PartyID initialized with val
func BuildPartyID(val string) PartyID {
	var field PartyID
	field.Value = val
	return field
}

//PartyIDSource is a CHAR field
type PartyIDSource struct{ message.CharValue }

//Tag returns tag.PartyIDSource (447)
func (f PartyIDSource) Tag() fix.Tag { return tag.PartyIDSource }

//BuildPartyIDSource returns a new PartyIDSource initialized with val
func BuildPartyIDSource(val string) PartyIDSource {
	var field PartyIDSource
	field.Value = val
	return field
}

//PartyListResponseType is a INT field
type PartyListResponseType struct{ message.IntValue }

//Tag returns tag.PartyListResponseType (1507)
func (f PartyListResponseType) Tag() fix.Tag { return tag.PartyListResponseType }

//BuildPartyListResponseType returns a new PartyListResponseType initialized with val
func BuildPartyListResponseType(val int) PartyListResponseType {
	var field PartyListResponseType
	field.Value = val
	return field
}

//PartyRelationship is a INT field
type PartyRelationship struct{ message.IntValue }

//Tag returns tag.PartyRelationship (1515)
func (f PartyRelationship) Tag() fix.Tag { return tag.PartyRelationship }

//BuildPartyRelationship returns a new PartyRelationship initialized with val
func BuildPartyRelationship(val int) PartyRelationship {
	var field PartyRelationship
	field.Value = val
	return field
}

//PartyRole is a INT field
type PartyRole struct{ message.IntValue }

//Tag returns tag.PartyRole (452)
func (f PartyRole) Tag() fix.Tag { return tag.PartyRole }

//BuildPartyRole returns a new PartyRole initialized with val
func BuildPartyRole(val int) PartyRole {
	var field PartyRole
	field.Value = val
	return field
}

//PartySubID is a STRING field
type PartySubID struct{ message.StringValue }

//Tag returns tag.PartySubID (523)
func (f PartySubID) Tag() fix.Tag { return tag.PartySubID }

//BuildPartySubID returns a new PartySubID initialized with val
func BuildPartySubID(val string) PartySubID {
	var field PartySubID
	field.Value = val
	return field
}

//PartySubIDType is a INT field
type PartySubIDType struct{ message.IntValue }

//Tag returns tag.PartySubIDType (803)
func (f PartySubIDType) Tag() fix.Tag { return tag.PartySubIDType }

//BuildPartySubIDType returns a new PartySubIDType initialized with val
func BuildPartySubIDType(val int) PartySubIDType {
	var field PartySubIDType
	field.Value = val
	return field
}

//Password is a STRING field
type Password struct{ message.StringValue }

//Tag returns tag.Password (554)
func (f Password) Tag() fix.Tag { return tag.Password }

//BuildPassword returns a new Password initialized with val
func BuildPassword(val string) Password {
	var field Password
	field.Value = val
	return field
}

//PaymentDate is a LOCALMKTDATE field
type PaymentDate struct{ message.LocalMktDateValue }

//Tag returns tag.PaymentDate (504)
func (f PaymentDate) Tag() fix.Tag { return tag.PaymentDate }

//BuildPaymentDate returns a new PaymentDate initialized with val
func BuildPaymentDate(val string) PaymentDate {
	var field PaymentDate
	field.Value = val
	return field
}

//PaymentMethod is a INT field
type PaymentMethod struct{ message.IntValue }

//Tag returns tag.PaymentMethod (492)
func (f PaymentMethod) Tag() fix.Tag { return tag.PaymentMethod }

//BuildPaymentMethod returns a new PaymentMethod initialized with val
func BuildPaymentMethod(val int) PaymentMethod {
	var field PaymentMethod
	field.Value = val
	return field
}

//PaymentRef is a STRING field
type PaymentRef struct{ message.StringValue }

//Tag returns tag.PaymentRef (476)
func (f PaymentRef) Tag() fix.Tag { return tag.PaymentRef }

//BuildPaymentRef returns a new PaymentRef initialized with val
func BuildPaymentRef(val string) PaymentRef {
	var field PaymentRef
	field.Value = val
	return field
}

//PaymentRemitterID is a STRING field
type PaymentRemitterID struct{ message.StringValue }

//Tag returns tag.PaymentRemitterID (505)
func (f PaymentRemitterID) Tag() fix.Tag { return tag.PaymentRemitterID }

//BuildPaymentRemitterID returns a new PaymentRemitterID initialized with val
func BuildPaymentRemitterID(val string) PaymentRemitterID {
	var field PaymentRemitterID
	field.Value = val
	return field
}

//PctAtRisk is a PERCENTAGE field
type PctAtRisk struct{ message.PercentageValue }

//Tag returns tag.PctAtRisk (869)
func (f PctAtRisk) Tag() fix.Tag { return tag.PctAtRisk }

//BuildPctAtRisk returns a new PctAtRisk initialized with val
func BuildPctAtRisk(val float64) PctAtRisk {
	var field PctAtRisk
	field.Value = val
	return field
}

//PegDifference is a PRICEOFFSET field
type PegDifference struct{ message.PriceOffsetValue }

//Tag returns tag.PegDifference (211)
func (f PegDifference) Tag() fix.Tag { return tag.PegDifference }

//BuildPegDifference returns a new PegDifference initialized with val
func BuildPegDifference(val float64) PegDifference {
	var field PegDifference
	field.Value = val
	return field
}

//PegLimitType is a INT field
type PegLimitType struct{ message.IntValue }

//Tag returns tag.PegLimitType (837)
func (f PegLimitType) Tag() fix.Tag { return tag.PegLimitType }

//BuildPegLimitType returns a new PegLimitType initialized with val
func BuildPegLimitType(val int) PegLimitType {
	var field PegLimitType
	field.Value = val
	return field
}

//PegMoveType is a INT field
type PegMoveType struct{ message.IntValue }

//Tag returns tag.PegMoveType (835)
func (f PegMoveType) Tag() fix.Tag { return tag.PegMoveType }

//BuildPegMoveType returns a new PegMoveType initialized with val
func BuildPegMoveType(val int) PegMoveType {
	var field PegMoveType
	field.Value = val
	return field
}

//PegOffsetType is a INT field
type PegOffsetType struct{ message.IntValue }

//Tag returns tag.PegOffsetType (836)
func (f PegOffsetType) Tag() fix.Tag { return tag.PegOffsetType }

//BuildPegOffsetType returns a new PegOffsetType initialized with val
func BuildPegOffsetType(val int) PegOffsetType {
	var field PegOffsetType
	field.Value = val
	return field
}

//PegOffsetValue is a FLOAT field
type PegOffsetValue struct{ message.FloatValue }

//Tag returns tag.PegOffsetValue (211)
func (f PegOffsetValue) Tag() fix.Tag { return tag.PegOffsetValue }

//BuildPegOffsetValue returns a new PegOffsetValue initialized with val
func BuildPegOffsetValue(val float64) PegOffsetValue {
	var field PegOffsetValue
	field.Value = val
	return field
}

//PegPriceType is a INT field
type PegPriceType struct{ message.IntValue }

//Tag returns tag.PegPriceType (1094)
func (f PegPriceType) Tag() fix.Tag { return tag.PegPriceType }

//BuildPegPriceType returns a new PegPriceType initialized with val
func BuildPegPriceType(val int) PegPriceType {
	var field PegPriceType
	field.Value = val
	return field
}

//PegRoundDirection is a INT field
type PegRoundDirection struct{ message.IntValue }

//Tag returns tag.PegRoundDirection (838)
func (f PegRoundDirection) Tag() fix.Tag { return tag.PegRoundDirection }

//BuildPegRoundDirection returns a new PegRoundDirection initialized with val
func BuildPegRoundDirection(val int) PegRoundDirection {
	var field PegRoundDirection
	field.Value = val
	return field
}

//PegScope is a INT field
type PegScope struct{ message.IntValue }

//Tag returns tag.PegScope (840)
func (f PegScope) Tag() fix.Tag { return tag.PegScope }

//BuildPegScope returns a new PegScope initialized with val
func BuildPegScope(val int) PegScope {
	var field PegScope
	field.Value = val
	return field
}

//PegSecurityDesc is a STRING field
type PegSecurityDesc struct{ message.StringValue }

//Tag returns tag.PegSecurityDesc (1099)
func (f PegSecurityDesc) Tag() fix.Tag { return tag.PegSecurityDesc }

//BuildPegSecurityDesc returns a new PegSecurityDesc initialized with val
func BuildPegSecurityDesc(val string) PegSecurityDesc {
	var field PegSecurityDesc
	field.Value = val
	return field
}

//PegSecurityID is a STRING field
type PegSecurityID struct{ message.StringValue }

//Tag returns tag.PegSecurityID (1097)
func (f PegSecurityID) Tag() fix.Tag { return tag.PegSecurityID }

//BuildPegSecurityID returns a new PegSecurityID initialized with val
func BuildPegSecurityID(val string) PegSecurityID {
	var field PegSecurityID
	field.Value = val
	return field
}

//PegSecurityIDSource is a STRING field
type PegSecurityIDSource struct{ message.StringValue }

//Tag returns tag.PegSecurityIDSource (1096)
func (f PegSecurityIDSource) Tag() fix.Tag { return tag.PegSecurityIDSource }

//BuildPegSecurityIDSource returns a new PegSecurityIDSource initialized with val
func BuildPegSecurityIDSource(val string) PegSecurityIDSource {
	var field PegSecurityIDSource
	field.Value = val
	return field
}

//PegSymbol is a STRING field
type PegSymbol struct{ message.StringValue }

//Tag returns tag.PegSymbol (1098)
func (f PegSymbol) Tag() fix.Tag { return tag.PegSymbol }

//BuildPegSymbol returns a new PegSymbol initialized with val
func BuildPegSymbol(val string) PegSymbol {
	var field PegSymbol
	field.Value = val
	return field
}

//PeggedPrice is a PRICE field
type PeggedPrice struct{ message.PriceValue }

//Tag returns tag.PeggedPrice (839)
func (f PeggedPrice) Tag() fix.Tag { return tag.PeggedPrice }

//BuildPeggedPrice returns a new PeggedPrice initialized with val
func BuildPeggedPrice(val float64) PeggedPrice {
	var field PeggedPrice
	field.Value = val
	return field
}

//PeggedRefPrice is a PRICE field
type PeggedRefPrice struct{ message.PriceValue }

//Tag returns tag.PeggedRefPrice (1095)
func (f PeggedRefPrice) Tag() fix.Tag { return tag.PeggedRefPrice }

//BuildPeggedRefPrice returns a new PeggedRefPrice initialized with val
func BuildPeggedRefPrice(val float64) PeggedRefPrice {
	var field PeggedRefPrice
	field.Value = val
	return field
}

//Pool is a STRING field
type Pool struct{ message.StringValue }

//Tag returns tag.Pool (691)
func (f Pool) Tag() fix.Tag { return tag.Pool }

//BuildPool returns a new Pool initialized with val
func BuildPool(val string) Pool {
	var field Pool
	field.Value = val
	return field
}

//PosAmt is a AMT field
type PosAmt struct{ message.AmtValue }

//Tag returns tag.PosAmt (708)
func (f PosAmt) Tag() fix.Tag { return tag.PosAmt }

//BuildPosAmt returns a new PosAmt initialized with val
func BuildPosAmt(val float64) PosAmt {
	var field PosAmt
	field.Value = val
	return field
}

//PosAmtType is a STRING field
type PosAmtType struct{ message.StringValue }

//Tag returns tag.PosAmtType (707)
func (f PosAmtType) Tag() fix.Tag { return tag.PosAmtType }

//BuildPosAmtType returns a new PosAmtType initialized with val
func BuildPosAmtType(val string) PosAmtType {
	var field PosAmtType
	field.Value = val
	return field
}

//PosMaintAction is a INT field
type PosMaintAction struct{ message.IntValue }

//Tag returns tag.PosMaintAction (712)
func (f PosMaintAction) Tag() fix.Tag { return tag.PosMaintAction }

//BuildPosMaintAction returns a new PosMaintAction initialized with val
func BuildPosMaintAction(val int) PosMaintAction {
	var field PosMaintAction
	field.Value = val
	return field
}

//PosMaintResult is a INT field
type PosMaintResult struct{ message.IntValue }

//Tag returns tag.PosMaintResult (723)
func (f PosMaintResult) Tag() fix.Tag { return tag.PosMaintResult }

//BuildPosMaintResult returns a new PosMaintResult initialized with val
func BuildPosMaintResult(val int) PosMaintResult {
	var field PosMaintResult
	field.Value = val
	return field
}

//PosMaintRptID is a STRING field
type PosMaintRptID struct{ message.StringValue }

//Tag returns tag.PosMaintRptID (721)
func (f PosMaintRptID) Tag() fix.Tag { return tag.PosMaintRptID }

//BuildPosMaintRptID returns a new PosMaintRptID initialized with val
func BuildPosMaintRptID(val string) PosMaintRptID {
	var field PosMaintRptID
	field.Value = val
	return field
}

//PosMaintRptRefID is a STRING field
type PosMaintRptRefID struct{ message.StringValue }

//Tag returns tag.PosMaintRptRefID (714)
func (f PosMaintRptRefID) Tag() fix.Tag { return tag.PosMaintRptRefID }

//BuildPosMaintRptRefID returns a new PosMaintRptRefID initialized with val
func BuildPosMaintRptRefID(val string) PosMaintRptRefID {
	var field PosMaintRptRefID
	field.Value = val
	return field
}

//PosMaintStatus is a INT field
type PosMaintStatus struct{ message.IntValue }

//Tag returns tag.PosMaintStatus (722)
func (f PosMaintStatus) Tag() fix.Tag { return tag.PosMaintStatus }

//BuildPosMaintStatus returns a new PosMaintStatus initialized with val
func BuildPosMaintStatus(val int) PosMaintStatus {
	var field PosMaintStatus
	field.Value = val
	return field
}

//PosQtyStatus is a INT field
type PosQtyStatus struct{ message.IntValue }

//Tag returns tag.PosQtyStatus (706)
func (f PosQtyStatus) Tag() fix.Tag { return tag.PosQtyStatus }

//BuildPosQtyStatus returns a new PosQtyStatus initialized with val
func BuildPosQtyStatus(val int) PosQtyStatus {
	var field PosQtyStatus
	field.Value = val
	return field
}

//PosReqID is a STRING field
type PosReqID struct{ message.StringValue }

//Tag returns tag.PosReqID (710)
func (f PosReqID) Tag() fix.Tag { return tag.PosReqID }

//BuildPosReqID returns a new PosReqID initialized with val
func BuildPosReqID(val string) PosReqID {
	var field PosReqID
	field.Value = val
	return field
}

//PosReqResult is a INT field
type PosReqResult struct{ message.IntValue }

//Tag returns tag.PosReqResult (728)
func (f PosReqResult) Tag() fix.Tag { return tag.PosReqResult }

//BuildPosReqResult returns a new PosReqResult initialized with val
func BuildPosReqResult(val int) PosReqResult {
	var field PosReqResult
	field.Value = val
	return field
}

//PosReqStatus is a INT field
type PosReqStatus struct{ message.IntValue }

//Tag returns tag.PosReqStatus (729)
func (f PosReqStatus) Tag() fix.Tag { return tag.PosReqStatus }

//BuildPosReqStatus returns a new PosReqStatus initialized with val
func BuildPosReqStatus(val int) PosReqStatus {
	var field PosReqStatus
	field.Value = val
	return field
}

//PosReqType is a INT field
type PosReqType struct{ message.IntValue }

//Tag returns tag.PosReqType (724)
func (f PosReqType) Tag() fix.Tag { return tag.PosReqType }

//BuildPosReqType returns a new PosReqType initialized with val
func BuildPosReqType(val int) PosReqType {
	var field PosReqType
	field.Value = val
	return field
}

//PosTransType is a INT field
type PosTransType struct{ message.IntValue }

//Tag returns tag.PosTransType (709)
func (f PosTransType) Tag() fix.Tag { return tag.PosTransType }

//BuildPosTransType returns a new PosTransType initialized with val
func BuildPosTransType(val int) PosTransType {
	var field PosTransType
	field.Value = val
	return field
}

//PosType is a STRING field
type PosType struct{ message.StringValue }

//Tag returns tag.PosType (703)
func (f PosType) Tag() fix.Tag { return tag.PosType }

//BuildPosType returns a new PosType initialized with val
func BuildPosType(val string) PosType {
	var field PosType
	field.Value = val
	return field
}

//PositionCurrency is a STRING field
type PositionCurrency struct{ message.StringValue }

//Tag returns tag.PositionCurrency (1055)
func (f PositionCurrency) Tag() fix.Tag { return tag.PositionCurrency }

//BuildPositionCurrency returns a new PositionCurrency initialized with val
func BuildPositionCurrency(val string) PositionCurrency {
	var field PositionCurrency
	field.Value = val
	return field
}

//PositionEffect is a CHAR field
type PositionEffect struct{ message.CharValue }

//Tag returns tag.PositionEffect (77)
func (f PositionEffect) Tag() fix.Tag { return tag.PositionEffect }

//BuildPositionEffect returns a new PositionEffect initialized with val
func BuildPositionEffect(val string) PositionEffect {
	var field PositionEffect
	field.Value = val
	return field
}

//PositionLimit is a INT field
type PositionLimit struct{ message.IntValue }

//Tag returns tag.PositionLimit (970)
func (f PositionLimit) Tag() fix.Tag { return tag.PositionLimit }

//BuildPositionLimit returns a new PositionLimit initialized with val
func BuildPositionLimit(val int) PositionLimit {
	var field PositionLimit
	field.Value = val
	return field
}

//PossDupFlag is a BOOLEAN field
type PossDupFlag struct{ message.BooleanValue }

//Tag returns tag.PossDupFlag (43)
func (f PossDupFlag) Tag() fix.Tag { return tag.PossDupFlag }

//BuildPossDupFlag returns a new PossDupFlag initialized with val
func BuildPossDupFlag(val bool) PossDupFlag {
	var field PossDupFlag
	field.Value = val
	return field
}

//PossResend is a BOOLEAN field
type PossResend struct{ message.BooleanValue }

//Tag returns tag.PossResend (97)
func (f PossResend) Tag() fix.Tag { return tag.PossResend }

//BuildPossResend returns a new PossResend initialized with val
func BuildPossResend(val bool) PossResend {
	var field PossResend
	field.Value = val
	return field
}

//PreTradeAnonymity is a BOOLEAN field
type PreTradeAnonymity struct{ message.BooleanValue }

//Tag returns tag.PreTradeAnonymity (1091)
func (f PreTradeAnonymity) Tag() fix.Tag { return tag.PreTradeAnonymity }

//BuildPreTradeAnonymity returns a new PreTradeAnonymity initialized with val
func BuildPreTradeAnonymity(val bool) PreTradeAnonymity {
	var field PreTradeAnonymity
	field.Value = val
	return field
}

//PreallocMethod is a CHAR field
type PreallocMethod struct{ message.CharValue }

//Tag returns tag.PreallocMethod (591)
func (f PreallocMethod) Tag() fix.Tag { return tag.PreallocMethod }

//BuildPreallocMethod returns a new PreallocMethod initialized with val
func BuildPreallocMethod(val string) PreallocMethod {
	var field PreallocMethod
	field.Value = val
	return field
}

//PrevClosePx is a PRICE field
type PrevClosePx struct{ message.PriceValue }

//Tag returns tag.PrevClosePx (140)
func (f PrevClosePx) Tag() fix.Tag { return tag.PrevClosePx }

//BuildPrevClosePx returns a new PrevClosePx initialized with val
func BuildPrevClosePx(val float64) PrevClosePx {
	var field PrevClosePx
	field.Value = val
	return field
}

//PreviouslyReported is a BOOLEAN field
type PreviouslyReported struct{ message.BooleanValue }

//Tag returns tag.PreviouslyReported (570)
func (f PreviouslyReported) Tag() fix.Tag { return tag.PreviouslyReported }

//BuildPreviouslyReported returns a new PreviouslyReported initialized with val
func BuildPreviouslyReported(val bool) PreviouslyReported {
	var field PreviouslyReported
	field.Value = val
	return field
}

//Price is a PRICE field
type Price struct{ message.PriceValue }

//Tag returns tag.Price (44)
func (f Price) Tag() fix.Tag { return tag.Price }

//BuildPrice returns a new Price initialized with val
func BuildPrice(val float64) Price {
	var field Price
	field.Value = val
	return field
}

//Price2 is a PRICE field
type Price2 struct{ message.PriceValue }

//Tag returns tag.Price2 (640)
func (f Price2) Tag() fix.Tag { return tag.Price2 }

//BuildPrice2 returns a new Price2 initialized with val
func BuildPrice2(val float64) Price2 {
	var field Price2
	field.Value = val
	return field
}

//PriceDelta is a FLOAT field
type PriceDelta struct{ message.FloatValue }

//Tag returns tag.PriceDelta (811)
func (f PriceDelta) Tag() fix.Tag { return tag.PriceDelta }

//BuildPriceDelta returns a new PriceDelta initialized with val
func BuildPriceDelta(val float64) PriceDelta {
	var field PriceDelta
	field.Value = val
	return field
}

//PriceImprovement is a PRICEOFFSET field
type PriceImprovement struct{ message.PriceOffsetValue }

//Tag returns tag.PriceImprovement (639)
func (f PriceImprovement) Tag() fix.Tag { return tag.PriceImprovement }

//BuildPriceImprovement returns a new PriceImprovement initialized with val
func BuildPriceImprovement(val float64) PriceImprovement {
	var field PriceImprovement
	field.Value = val
	return field
}

//PriceLimitType is a INT field
type PriceLimitType struct{ message.IntValue }

//Tag returns tag.PriceLimitType (1306)
func (f PriceLimitType) Tag() fix.Tag { return tag.PriceLimitType }

//BuildPriceLimitType returns a new PriceLimitType initialized with val
func BuildPriceLimitType(val int) PriceLimitType {
	var field PriceLimitType
	field.Value = val
	return field
}

//PriceProtectionScope is a CHAR field
type PriceProtectionScope struct{ message.CharValue }

//Tag returns tag.PriceProtectionScope (1092)
func (f PriceProtectionScope) Tag() fix.Tag { return tag.PriceProtectionScope }

//BuildPriceProtectionScope returns a new PriceProtectionScope initialized with val
func BuildPriceProtectionScope(val string) PriceProtectionScope {
	var field PriceProtectionScope
	field.Value = val
	return field
}

//PriceQuoteMethod is a STRING field
type PriceQuoteMethod struct{ message.StringValue }

//Tag returns tag.PriceQuoteMethod (1196)
func (f PriceQuoteMethod) Tag() fix.Tag { return tag.PriceQuoteMethod }

//BuildPriceQuoteMethod returns a new PriceQuoteMethod initialized with val
func BuildPriceQuoteMethod(val string) PriceQuoteMethod {
	var field PriceQuoteMethod
	field.Value = val
	return field
}

//PriceType is a INT field
type PriceType struct{ message.IntValue }

//Tag returns tag.PriceType (423)
func (f PriceType) Tag() fix.Tag { return tag.PriceType }

//BuildPriceType returns a new PriceType initialized with val
func BuildPriceType(val int) PriceType {
	var field PriceType
	field.Value = val
	return field
}

//PriceUnitOfMeasure is a STRING field
type PriceUnitOfMeasure struct{ message.StringValue }

//Tag returns tag.PriceUnitOfMeasure (1191)
func (f PriceUnitOfMeasure) Tag() fix.Tag { return tag.PriceUnitOfMeasure }

//BuildPriceUnitOfMeasure returns a new PriceUnitOfMeasure initialized with val
func BuildPriceUnitOfMeasure(val string) PriceUnitOfMeasure {
	var field PriceUnitOfMeasure
	field.Value = val
	return field
}

//PriceUnitOfMeasureQty is a QTY field
type PriceUnitOfMeasureQty struct{ message.QtyValue }

//Tag returns tag.PriceUnitOfMeasureQty (1192)
func (f PriceUnitOfMeasureQty) Tag() fix.Tag { return tag.PriceUnitOfMeasureQty }

//BuildPriceUnitOfMeasureQty returns a new PriceUnitOfMeasureQty initialized with val
func BuildPriceUnitOfMeasureQty(val float64) PriceUnitOfMeasureQty {
	var field PriceUnitOfMeasureQty
	field.Value = val
	return field
}

//PriorSettlPrice is a PRICE field
type PriorSettlPrice struct{ message.PriceValue }

//Tag returns tag.PriorSettlPrice (734)
func (f PriorSettlPrice) Tag() fix.Tag { return tag.PriorSettlPrice }

//BuildPriorSettlPrice returns a new PriorSettlPrice initialized with val
func BuildPriorSettlPrice(val float64) PriorSettlPrice {
	var field PriorSettlPrice
	field.Value = val
	return field
}

//PriorSpreadIndicator is a BOOLEAN field
type PriorSpreadIndicator struct{ message.BooleanValue }

//Tag returns tag.PriorSpreadIndicator (720)
func (f PriorSpreadIndicator) Tag() fix.Tag { return tag.PriorSpreadIndicator }

//BuildPriorSpreadIndicator returns a new PriorSpreadIndicator initialized with val
func BuildPriorSpreadIndicator(val bool) PriorSpreadIndicator {
	var field PriorSpreadIndicator
	field.Value = val
	return field
}

//PriorityIndicator is a INT field
type PriorityIndicator struct{ message.IntValue }

//Tag returns tag.PriorityIndicator (638)
func (f PriorityIndicator) Tag() fix.Tag { return tag.PriorityIndicator }

//BuildPriorityIndicator returns a new PriorityIndicator initialized with val
func BuildPriorityIndicator(val int) PriorityIndicator {
	var field PriorityIndicator
	field.Value = val
	return field
}

//PrivateQuote is a BOOLEAN field
type PrivateQuote struct{ message.BooleanValue }

//Tag returns tag.PrivateQuote (1171)
func (f PrivateQuote) Tag() fix.Tag { return tag.PrivateQuote }

//BuildPrivateQuote returns a new PrivateQuote initialized with val
func BuildPrivateQuote(val bool) PrivateQuote {
	var field PrivateQuote
	field.Value = val
	return field
}

//ProcessCode is a CHAR field
type ProcessCode struct{ message.CharValue }

//Tag returns tag.ProcessCode (81)
func (f ProcessCode) Tag() fix.Tag { return tag.ProcessCode }

//BuildProcessCode returns a new ProcessCode initialized with val
func BuildProcessCode(val string) ProcessCode {
	var field ProcessCode
	field.Value = val
	return field
}

//Product is a INT field
type Product struct{ message.IntValue }

//Tag returns tag.Product (460)
func (f Product) Tag() fix.Tag { return tag.Product }

//BuildProduct returns a new Product initialized with val
func BuildProduct(val int) Product {
	var field Product
	field.Value = val
	return field
}

//ProductComplex is a STRING field
type ProductComplex struct{ message.StringValue }

//Tag returns tag.ProductComplex (1227)
func (f ProductComplex) Tag() fix.Tag { return tag.ProductComplex }

//BuildProductComplex returns a new ProductComplex initialized with val
func BuildProductComplex(val string) ProductComplex {
	var field ProductComplex
	field.Value = val
	return field
}

//ProgPeriodInterval is a INT field
type ProgPeriodInterval struct{ message.IntValue }

//Tag returns tag.ProgPeriodInterval (415)
func (f ProgPeriodInterval) Tag() fix.Tag { return tag.ProgPeriodInterval }

//BuildProgPeriodInterval returns a new ProgPeriodInterval initialized with val
func BuildProgPeriodInterval(val int) ProgPeriodInterval {
	var field ProgPeriodInterval
	field.Value = val
	return field
}

//ProgRptReqs is a INT field
type ProgRptReqs struct{ message.IntValue }

//Tag returns tag.ProgRptReqs (414)
func (f ProgRptReqs) Tag() fix.Tag { return tag.ProgRptReqs }

//BuildProgRptReqs returns a new ProgRptReqs initialized with val
func BuildProgRptReqs(val int) ProgRptReqs {
	var field ProgRptReqs
	field.Value = val
	return field
}

//PublishTrdIndicator is a BOOLEAN field
type PublishTrdIndicator struct{ message.BooleanValue }

//Tag returns tag.PublishTrdIndicator (852)
func (f PublishTrdIndicator) Tag() fix.Tag { return tag.PublishTrdIndicator }

//BuildPublishTrdIndicator returns a new PublishTrdIndicator initialized with val
func BuildPublishTrdIndicator(val bool) PublishTrdIndicator {
	var field PublishTrdIndicator
	field.Value = val
	return field
}

//PutOrCall is a INT field
type PutOrCall struct{ message.IntValue }

//Tag returns tag.PutOrCall (201)
func (f PutOrCall) Tag() fix.Tag { return tag.PutOrCall }

//BuildPutOrCall returns a new PutOrCall initialized with val
func BuildPutOrCall(val int) PutOrCall {
	var field PutOrCall
	field.Value = val
	return field
}

//QtyType is a INT field
type QtyType struct{ message.IntValue }

//Tag returns tag.QtyType (854)
func (f QtyType) Tag() fix.Tag { return tag.QtyType }

//BuildQtyType returns a new QtyType initialized with val
func BuildQtyType(val int) QtyType {
	var field QtyType
	field.Value = val
	return field
}

//Quantity is a QTY field
type Quantity struct{ message.QtyValue }

//Tag returns tag.Quantity (53)
func (f Quantity) Tag() fix.Tag { return tag.Quantity }

//BuildQuantity returns a new Quantity initialized with val
func BuildQuantity(val float64) Quantity {
	var field Quantity
	field.Value = val
	return field
}

//QuantityDate is a LOCALMKTDATE field
type QuantityDate struct{ message.LocalMktDateValue }

//Tag returns tag.QuantityDate (976)
func (f QuantityDate) Tag() fix.Tag { return tag.QuantityDate }

//BuildQuantityDate returns a new QuantityDate initialized with val
func BuildQuantityDate(val string) QuantityDate {
	var field QuantityDate
	field.Value = val
	return field
}

//QuantityType is a INT field
type QuantityType struct{ message.IntValue }

//Tag returns tag.QuantityType (465)
func (f QuantityType) Tag() fix.Tag { return tag.QuantityType }

//BuildQuantityType returns a new QuantityType initialized with val
func BuildQuantityType(val int) QuantityType {
	var field QuantityType
	field.Value = val
	return field
}

//QuoteAckStatus is a INT field
type QuoteAckStatus struct{ message.IntValue }

//Tag returns tag.QuoteAckStatus (297)
func (f QuoteAckStatus) Tag() fix.Tag { return tag.QuoteAckStatus }

//BuildQuoteAckStatus returns a new QuoteAckStatus initialized with val
func BuildQuoteAckStatus(val int) QuoteAckStatus {
	var field QuoteAckStatus
	field.Value = val
	return field
}

//QuoteCancelType is a INT field
type QuoteCancelType struct{ message.IntValue }

//Tag returns tag.QuoteCancelType (298)
func (f QuoteCancelType) Tag() fix.Tag { return tag.QuoteCancelType }

//BuildQuoteCancelType returns a new QuoteCancelType initialized with val
func BuildQuoteCancelType(val int) QuoteCancelType {
	var field QuoteCancelType
	field.Value = val
	return field
}

//QuoteCondition is a MULTIPLESTRINGVALUE field
type QuoteCondition struct{ message.MultipleStringValue }

//Tag returns tag.QuoteCondition (276)
func (f QuoteCondition) Tag() fix.Tag { return tag.QuoteCondition }

//BuildQuoteCondition returns a new QuoteCondition initialized with val
func BuildQuoteCondition(val string) QuoteCondition {
	var field QuoteCondition
	field.Value = val
	return field
}

//QuoteEntryID is a STRING field
type QuoteEntryID struct{ message.StringValue }

//Tag returns tag.QuoteEntryID (299)
func (f QuoteEntryID) Tag() fix.Tag { return tag.QuoteEntryID }

//BuildQuoteEntryID returns a new QuoteEntryID initialized with val
func BuildQuoteEntryID(val string) QuoteEntryID {
	var field QuoteEntryID
	field.Value = val
	return field
}

//QuoteEntryRejectReason is a INT field
type QuoteEntryRejectReason struct{ message.IntValue }

//Tag returns tag.QuoteEntryRejectReason (368)
func (f QuoteEntryRejectReason) Tag() fix.Tag { return tag.QuoteEntryRejectReason }

//BuildQuoteEntryRejectReason returns a new QuoteEntryRejectReason initialized with val
func BuildQuoteEntryRejectReason(val int) QuoteEntryRejectReason {
	var field QuoteEntryRejectReason
	field.Value = val
	return field
}

//QuoteEntryStatus is a INT field
type QuoteEntryStatus struct{ message.IntValue }

//Tag returns tag.QuoteEntryStatus (1167)
func (f QuoteEntryStatus) Tag() fix.Tag { return tag.QuoteEntryStatus }

//BuildQuoteEntryStatus returns a new QuoteEntryStatus initialized with val
func BuildQuoteEntryStatus(val int) QuoteEntryStatus {
	var field QuoteEntryStatus
	field.Value = val
	return field
}

//QuoteID is a STRING field
type QuoteID struct{ message.StringValue }

//Tag returns tag.QuoteID (117)
func (f QuoteID) Tag() fix.Tag { return tag.QuoteID }

//BuildQuoteID returns a new QuoteID initialized with val
func BuildQuoteID(val string) QuoteID {
	var field QuoteID
	field.Value = val
	return field
}

//QuoteMsgID is a STRING field
type QuoteMsgID struct{ message.StringValue }

//Tag returns tag.QuoteMsgID (1166)
func (f QuoteMsgID) Tag() fix.Tag { return tag.QuoteMsgID }

//BuildQuoteMsgID returns a new QuoteMsgID initialized with val
func BuildQuoteMsgID(val string) QuoteMsgID {
	var field QuoteMsgID
	field.Value = val
	return field
}

//QuotePriceType is a INT field
type QuotePriceType struct{ message.IntValue }

//Tag returns tag.QuotePriceType (692)
func (f QuotePriceType) Tag() fix.Tag { return tag.QuotePriceType }

//BuildQuotePriceType returns a new QuotePriceType initialized with val
func BuildQuotePriceType(val int) QuotePriceType {
	var field QuotePriceType
	field.Value = val
	return field
}

//QuoteQualifier is a CHAR field
type QuoteQualifier struct{ message.CharValue }

//Tag returns tag.QuoteQualifier (695)
func (f QuoteQualifier) Tag() fix.Tag { return tag.QuoteQualifier }

//BuildQuoteQualifier returns a new QuoteQualifier initialized with val
func BuildQuoteQualifier(val string) QuoteQualifier {
	var field QuoteQualifier
	field.Value = val
	return field
}

//QuoteRejectReason is a INT field
type QuoteRejectReason struct{ message.IntValue }

//Tag returns tag.QuoteRejectReason (300)
func (f QuoteRejectReason) Tag() fix.Tag { return tag.QuoteRejectReason }

//BuildQuoteRejectReason returns a new QuoteRejectReason initialized with val
func BuildQuoteRejectReason(val int) QuoteRejectReason {
	var field QuoteRejectReason
	field.Value = val
	return field
}

//QuoteReqID is a STRING field
type QuoteReqID struct{ message.StringValue }

//Tag returns tag.QuoteReqID (131)
func (f QuoteReqID) Tag() fix.Tag { return tag.QuoteReqID }

//BuildQuoteReqID returns a new QuoteReqID initialized with val
func BuildQuoteReqID(val string) QuoteReqID {
	var field QuoteReqID
	field.Value = val
	return field
}

//QuoteRequestRejectReason is a INT field
type QuoteRequestRejectReason struct{ message.IntValue }

//Tag returns tag.QuoteRequestRejectReason (658)
func (f QuoteRequestRejectReason) Tag() fix.Tag { return tag.QuoteRequestRejectReason }

//BuildQuoteRequestRejectReason returns a new QuoteRequestRejectReason initialized with val
func BuildQuoteRequestRejectReason(val int) QuoteRequestRejectReason {
	var field QuoteRequestRejectReason
	field.Value = val
	return field
}

//QuoteRequestType is a INT field
type QuoteRequestType struct{ message.IntValue }

//Tag returns tag.QuoteRequestType (303)
func (f QuoteRequestType) Tag() fix.Tag { return tag.QuoteRequestType }

//BuildQuoteRequestType returns a new QuoteRequestType initialized with val
func BuildQuoteRequestType(val int) QuoteRequestType {
	var field QuoteRequestType
	field.Value = val
	return field
}

//QuoteRespID is a STRING field
type QuoteRespID struct{ message.StringValue }

//Tag returns tag.QuoteRespID (693)
func (f QuoteRespID) Tag() fix.Tag { return tag.QuoteRespID }

//BuildQuoteRespID returns a new QuoteRespID initialized with val
func BuildQuoteRespID(val string) QuoteRespID {
	var field QuoteRespID
	field.Value = val
	return field
}

//QuoteRespType is a INT field
type QuoteRespType struct{ message.IntValue }

//Tag returns tag.QuoteRespType (694)
func (f QuoteRespType) Tag() fix.Tag { return tag.QuoteRespType }

//BuildQuoteRespType returns a new QuoteRespType initialized with val
func BuildQuoteRespType(val int) QuoteRespType {
	var field QuoteRespType
	field.Value = val
	return field
}

//QuoteResponseLevel is a INT field
type QuoteResponseLevel struct{ message.IntValue }

//Tag returns tag.QuoteResponseLevel (301)
func (f QuoteResponseLevel) Tag() fix.Tag { return tag.QuoteResponseLevel }

//BuildQuoteResponseLevel returns a new QuoteResponseLevel initialized with val
func BuildQuoteResponseLevel(val int) QuoteResponseLevel {
	var field QuoteResponseLevel
	field.Value = val
	return field
}

//QuoteSetID is a STRING field
type QuoteSetID struct{ message.StringValue }

//Tag returns tag.QuoteSetID (302)
func (f QuoteSetID) Tag() fix.Tag { return tag.QuoteSetID }

//BuildQuoteSetID returns a new QuoteSetID initialized with val
func BuildQuoteSetID(val string) QuoteSetID {
	var field QuoteSetID
	field.Value = val
	return field
}

//QuoteSetValidUntilTime is a UTCTIMESTAMP field
type QuoteSetValidUntilTime struct{ message.UTCTimestampValue }

//Tag returns tag.QuoteSetValidUntilTime (367)
func (f QuoteSetValidUntilTime) Tag() fix.Tag { return tag.QuoteSetValidUntilTime }

//QuoteStatus is a INT field
type QuoteStatus struct{ message.IntValue }

//Tag returns tag.QuoteStatus (297)
func (f QuoteStatus) Tag() fix.Tag { return tag.QuoteStatus }

//BuildQuoteStatus returns a new QuoteStatus initialized with val
func BuildQuoteStatus(val int) QuoteStatus {
	var field QuoteStatus
	field.Value = val
	return field
}

//QuoteStatusReqID is a STRING field
type QuoteStatusReqID struct{ message.StringValue }

//Tag returns tag.QuoteStatusReqID (649)
func (f QuoteStatusReqID) Tag() fix.Tag { return tag.QuoteStatusReqID }

//BuildQuoteStatusReqID returns a new QuoteStatusReqID initialized with val
func BuildQuoteStatusReqID(val string) QuoteStatusReqID {
	var field QuoteStatusReqID
	field.Value = val
	return field
}

//QuoteType is a INT field
type QuoteType struct{ message.IntValue }

//Tag returns tag.QuoteType (537)
func (f QuoteType) Tag() fix.Tag { return tag.QuoteType }

//BuildQuoteType returns a new QuoteType initialized with val
func BuildQuoteType(val int) QuoteType {
	var field QuoteType
	field.Value = val
	return field
}

//RFQReqID is a STRING field
type RFQReqID struct{ message.StringValue }

//Tag returns tag.RFQReqID (644)
func (f RFQReqID) Tag() fix.Tag { return tag.RFQReqID }

//BuildRFQReqID returns a new RFQReqID initialized with val
func BuildRFQReqID(val string) RFQReqID {
	var field RFQReqID
	field.Value = val
	return field
}

//RateSource is a INT field
type RateSource struct{ message.IntValue }

//Tag returns tag.RateSource (1446)
func (f RateSource) Tag() fix.Tag { return tag.RateSource }

//BuildRateSource returns a new RateSource initialized with val
func BuildRateSource(val int) RateSource {
	var field RateSource
	field.Value = val
	return field
}

//RateSourceType is a INT field
type RateSourceType struct{ message.IntValue }

//Tag returns tag.RateSourceType (1447)
func (f RateSourceType) Tag() fix.Tag { return tag.RateSourceType }

//BuildRateSourceType returns a new RateSourceType initialized with val
func BuildRateSourceType(val int) RateSourceType {
	var field RateSourceType
	field.Value = val
	return field
}

//RatioQty is a QTY field
type RatioQty struct{ message.QtyValue }

//Tag returns tag.RatioQty (319)
func (f RatioQty) Tag() fix.Tag { return tag.RatioQty }

//BuildRatioQty returns a new RatioQty initialized with val
func BuildRatioQty(val float64) RatioQty {
	var field RatioQty
	field.Value = val
	return field
}

//RawData is a DATA field
type RawData struct{ message.DataValue }

//Tag returns tag.RawData (96)
func (f RawData) Tag() fix.Tag { return tag.RawData }

//BuildRawData returns a new RawData initialized with val
func BuildRawData(val string) RawData {
	var field RawData
	field.Value = val
	return field
}

//RawDataLength is a LENGTH field
type RawDataLength struct{ message.LengthValue }

//Tag returns tag.RawDataLength (95)
func (f RawDataLength) Tag() fix.Tag { return tag.RawDataLength }

//BuildRawDataLength returns a new RawDataLength initialized with val
func BuildRawDataLength(val int) RawDataLength {
	var field RawDataLength
	field.Value = val
	return field
}

//ReceivedDeptID is a STRING field
type ReceivedDeptID struct{ message.StringValue }

//Tag returns tag.ReceivedDeptID (1030)
func (f ReceivedDeptID) Tag() fix.Tag { return tag.ReceivedDeptID }

//BuildReceivedDeptID returns a new ReceivedDeptID initialized with val
func BuildReceivedDeptID(val string) ReceivedDeptID {
	var field ReceivedDeptID
	field.Value = val
	return field
}

//RedemptionDate is a LOCALMKTDATE field
type RedemptionDate struct{ message.LocalMktDateValue }

//Tag returns tag.RedemptionDate (240)
func (f RedemptionDate) Tag() fix.Tag { return tag.RedemptionDate }

//BuildRedemptionDate returns a new RedemptionDate initialized with val
func BuildRedemptionDate(val string) RedemptionDate {
	var field RedemptionDate
	field.Value = val
	return field
}

//RefAllocID is a STRING field
type RefAllocID struct{ message.StringValue }

//Tag returns tag.RefAllocID (72)
func (f RefAllocID) Tag() fix.Tag { return tag.RefAllocID }

//BuildRefAllocID returns a new RefAllocID initialized with val
func BuildRefAllocID(val string) RefAllocID {
	var field RefAllocID
	field.Value = val
	return field
}

//RefApplExtID is a INT field
type RefApplExtID struct{ message.IntValue }

//Tag returns tag.RefApplExtID (1406)
func (f RefApplExtID) Tag() fix.Tag { return tag.RefApplExtID }

//BuildRefApplExtID returns a new RefApplExtID initialized with val
func BuildRefApplExtID(val int) RefApplExtID {
	var field RefApplExtID
	field.Value = val
	return field
}

//RefApplID is a STRING field
type RefApplID struct{ message.StringValue }

//Tag returns tag.RefApplID (1355)
func (f RefApplID) Tag() fix.Tag { return tag.RefApplID }

//BuildRefApplID returns a new RefApplID initialized with val
func BuildRefApplID(val string) RefApplID {
	var field RefApplID
	field.Value = val
	return field
}

//RefApplLastSeqNum is a SEQNUM field
type RefApplLastSeqNum struct{ message.SeqNumValue }

//Tag returns tag.RefApplLastSeqNum (1357)
func (f RefApplLastSeqNum) Tag() fix.Tag { return tag.RefApplLastSeqNum }

//BuildRefApplLastSeqNum returns a new RefApplLastSeqNum initialized with val
func BuildRefApplLastSeqNum(val int) RefApplLastSeqNum {
	var field RefApplLastSeqNum
	field.Value = val
	return field
}

//RefApplReqID is a STRING field
type RefApplReqID struct{ message.StringValue }

//Tag returns tag.RefApplReqID (1433)
func (f RefApplReqID) Tag() fix.Tag { return tag.RefApplReqID }

//BuildRefApplReqID returns a new RefApplReqID initialized with val
func BuildRefApplReqID(val string) RefApplReqID {
	var field RefApplReqID
	field.Value = val
	return field
}

//RefApplVerID is a STRING field
type RefApplVerID struct{ message.StringValue }

//Tag returns tag.RefApplVerID (1130)
func (f RefApplVerID) Tag() fix.Tag { return tag.RefApplVerID }

//BuildRefApplVerID returns a new RefApplVerID initialized with val
func BuildRefApplVerID(val string) RefApplVerID {
	var field RefApplVerID
	field.Value = val
	return field
}

//RefCompID is a STRING field
type RefCompID struct{ message.StringValue }

//Tag returns tag.RefCompID (930)
func (f RefCompID) Tag() fix.Tag { return tag.RefCompID }

//BuildRefCompID returns a new RefCompID initialized with val
func BuildRefCompID(val string) RefCompID {
	var field RefCompID
	field.Value = val
	return field
}

//RefCstmApplVerID is a STRING field
type RefCstmApplVerID struct{ message.StringValue }

//Tag returns tag.RefCstmApplVerID (1131)
func (f RefCstmApplVerID) Tag() fix.Tag { return tag.RefCstmApplVerID }

//BuildRefCstmApplVerID returns a new RefCstmApplVerID initialized with val
func BuildRefCstmApplVerID(val string) RefCstmApplVerID {
	var field RefCstmApplVerID
	field.Value = val
	return field
}

//RefMsgType is a STRING field
type RefMsgType struct{ message.StringValue }

//Tag returns tag.RefMsgType (372)
func (f RefMsgType) Tag() fix.Tag { return tag.RefMsgType }

//BuildRefMsgType returns a new RefMsgType initialized with val
func BuildRefMsgType(val string) RefMsgType {
	var field RefMsgType
	field.Value = val
	return field
}

//RefOrdIDReason is a INT field
type RefOrdIDReason struct{ message.IntValue }

//Tag returns tag.RefOrdIDReason (1431)
func (f RefOrdIDReason) Tag() fix.Tag { return tag.RefOrdIDReason }

//BuildRefOrdIDReason returns a new RefOrdIDReason initialized with val
func BuildRefOrdIDReason(val int) RefOrdIDReason {
	var field RefOrdIDReason
	field.Value = val
	return field
}

//RefOrderID is a STRING field
type RefOrderID struct{ message.StringValue }

//Tag returns tag.RefOrderID (1080)
func (f RefOrderID) Tag() fix.Tag { return tag.RefOrderID }

//BuildRefOrderID returns a new RefOrderID initialized with val
func BuildRefOrderID(val string) RefOrderID {
	var field RefOrderID
	field.Value = val
	return field
}

//RefOrderIDSource is a CHAR field
type RefOrderIDSource struct{ message.CharValue }

//Tag returns tag.RefOrderIDSource (1081)
func (f RefOrderIDSource) Tag() fix.Tag { return tag.RefOrderIDSource }

//BuildRefOrderIDSource returns a new RefOrderIDSource initialized with val
func BuildRefOrderIDSource(val string) RefOrderIDSource {
	var field RefOrderIDSource
	field.Value = val
	return field
}

//RefSeqNum is a SEQNUM field
type RefSeqNum struct{ message.SeqNumValue }

//Tag returns tag.RefSeqNum (45)
func (f RefSeqNum) Tag() fix.Tag { return tag.RefSeqNum }

//BuildRefSeqNum returns a new RefSeqNum initialized with val
func BuildRefSeqNum(val int) RefSeqNum {
	var field RefSeqNum
	field.Value = val
	return field
}

//RefSubID is a STRING field
type RefSubID struct{ message.StringValue }

//Tag returns tag.RefSubID (931)
func (f RefSubID) Tag() fix.Tag { return tag.RefSubID }

//BuildRefSubID returns a new RefSubID initialized with val
func BuildRefSubID(val string) RefSubID {
	var field RefSubID
	field.Value = val
	return field
}

//RefTagID is a INT field
type RefTagID struct{ message.IntValue }

//Tag returns tag.RefTagID (371)
func (f RefTagID) Tag() fix.Tag { return tag.RefTagID }

//BuildRefTagID returns a new RefTagID initialized with val
func BuildRefTagID(val int) RefTagID {
	var field RefTagID
	field.Value = val
	return field
}

//ReferencePage is a STRING field
type ReferencePage struct{ message.StringValue }

//Tag returns tag.ReferencePage (1448)
func (f ReferencePage) Tag() fix.Tag { return tag.ReferencePage }

//BuildReferencePage returns a new ReferencePage initialized with val
func BuildReferencePage(val string) ReferencePage {
	var field ReferencePage
	field.Value = val
	return field
}

//RefreshIndicator is a BOOLEAN field
type RefreshIndicator struct{ message.BooleanValue }

//Tag returns tag.RefreshIndicator (1187)
func (f RefreshIndicator) Tag() fix.Tag { return tag.RefreshIndicator }

//BuildRefreshIndicator returns a new RefreshIndicator initialized with val
func BuildRefreshIndicator(val bool) RefreshIndicator {
	var field RefreshIndicator
	field.Value = val
	return field
}

//RefreshQty is a QTY field
type RefreshQty struct{ message.QtyValue }

//Tag returns tag.RefreshQty (1088)
func (f RefreshQty) Tag() fix.Tag { return tag.RefreshQty }

//BuildRefreshQty returns a new RefreshQty initialized with val
func BuildRefreshQty(val float64) RefreshQty {
	var field RefreshQty
	field.Value = val
	return field
}

//RegistAcctType is a STRING field
type RegistAcctType struct{ message.StringValue }

//Tag returns tag.RegistAcctType (493)
func (f RegistAcctType) Tag() fix.Tag { return tag.RegistAcctType }

//BuildRegistAcctType returns a new RegistAcctType initialized with val
func BuildRegistAcctType(val string) RegistAcctType {
	var field RegistAcctType
	field.Value = val
	return field
}

//RegistDetls is a STRING field
type RegistDetls struct{ message.StringValue }

//Tag returns tag.RegistDetls (509)
func (f RegistDetls) Tag() fix.Tag { return tag.RegistDetls }

//BuildRegistDetls returns a new RegistDetls initialized with val
func BuildRegistDetls(val string) RegistDetls {
	var field RegistDetls
	field.Value = val
	return field
}

//RegistDtls is a STRING field
type RegistDtls struct{ message.StringValue }

//Tag returns tag.RegistDtls (509)
func (f RegistDtls) Tag() fix.Tag { return tag.RegistDtls }

//BuildRegistDtls returns a new RegistDtls initialized with val
func BuildRegistDtls(val string) RegistDtls {
	var field RegistDtls
	field.Value = val
	return field
}

//RegistEmail is a STRING field
type RegistEmail struct{ message.StringValue }

//Tag returns tag.RegistEmail (511)
func (f RegistEmail) Tag() fix.Tag { return tag.RegistEmail }

//BuildRegistEmail returns a new RegistEmail initialized with val
func BuildRegistEmail(val string) RegistEmail {
	var field RegistEmail
	field.Value = val
	return field
}

//RegistID is a STRING field
type RegistID struct{ message.StringValue }

//Tag returns tag.RegistID (513)
func (f RegistID) Tag() fix.Tag { return tag.RegistID }

//BuildRegistID returns a new RegistID initialized with val
func BuildRegistID(val string) RegistID {
	var field RegistID
	field.Value = val
	return field
}

//RegistRefID is a STRING field
type RegistRefID struct{ message.StringValue }

//Tag returns tag.RegistRefID (508)
func (f RegistRefID) Tag() fix.Tag { return tag.RegistRefID }

//BuildRegistRefID returns a new RegistRefID initialized with val
func BuildRegistRefID(val string) RegistRefID {
	var field RegistRefID
	field.Value = val
	return field
}

//RegistRejReasonCode is a INT field
type RegistRejReasonCode struct{ message.IntValue }

//Tag returns tag.RegistRejReasonCode (507)
func (f RegistRejReasonCode) Tag() fix.Tag { return tag.RegistRejReasonCode }

//BuildRegistRejReasonCode returns a new RegistRejReasonCode initialized with val
func BuildRegistRejReasonCode(val int) RegistRejReasonCode {
	var field RegistRejReasonCode
	field.Value = val
	return field
}

//RegistRejReasonText is a STRING field
type RegistRejReasonText struct{ message.StringValue }

//Tag returns tag.RegistRejReasonText (496)
func (f RegistRejReasonText) Tag() fix.Tag { return tag.RegistRejReasonText }

//BuildRegistRejReasonText returns a new RegistRejReasonText initialized with val
func BuildRegistRejReasonText(val string) RegistRejReasonText {
	var field RegistRejReasonText
	field.Value = val
	return field
}

//RegistStatus is a CHAR field
type RegistStatus struct{ message.CharValue }

//Tag returns tag.RegistStatus (506)
func (f RegistStatus) Tag() fix.Tag { return tag.RegistStatus }

//BuildRegistStatus returns a new RegistStatus initialized with val
func BuildRegistStatus(val string) RegistStatus {
	var field RegistStatus
	field.Value = val
	return field
}

//RegistTransType is a CHAR field
type RegistTransType struct{ message.CharValue }

//Tag returns tag.RegistTransType (514)
func (f RegistTransType) Tag() fix.Tag { return tag.RegistTransType }

//BuildRegistTransType returns a new RegistTransType initialized with val
func BuildRegistTransType(val string) RegistTransType {
	var field RegistTransType
	field.Value = val
	return field
}

//RejectText is a STRING field
type RejectText struct{ message.StringValue }

//Tag returns tag.RejectText (1328)
func (f RejectText) Tag() fix.Tag { return tag.RejectText }

//BuildRejectText returns a new RejectText initialized with val
func BuildRejectText(val string) RejectText {
	var field RejectText
	field.Value = val
	return field
}

//RelSymTransactTime is a UTCTIMESTAMP field
type RelSymTransactTime struct{ message.UTCTimestampValue }

//Tag returns tag.RelSymTransactTime (1504)
func (f RelSymTransactTime) Tag() fix.Tag { return tag.RelSymTransactTime }

//RelatdSym is a STRING field
type RelatdSym struct{ message.StringValue }

//Tag returns tag.RelatdSym (46)
func (f RelatdSym) Tag() fix.Tag { return tag.RelatdSym }

//BuildRelatdSym returns a new RelatdSym initialized with val
func BuildRelatdSym(val string) RelatdSym {
	var field RelatdSym
	field.Value = val
	return field
}

//RelatedContextPartyID is a STRING field
type RelatedContextPartyID struct{ message.StringValue }

//Tag returns tag.RelatedContextPartyID (1576)
func (f RelatedContextPartyID) Tag() fix.Tag { return tag.RelatedContextPartyID }

//BuildRelatedContextPartyID returns a new RelatedContextPartyID initialized with val
func BuildRelatedContextPartyID(val string) RelatedContextPartyID {
	var field RelatedContextPartyID
	field.Value = val
	return field
}

//RelatedContextPartyIDSource is a CHAR field
type RelatedContextPartyIDSource struct{ message.CharValue }

//Tag returns tag.RelatedContextPartyIDSource (1577)
func (f RelatedContextPartyIDSource) Tag() fix.Tag { return tag.RelatedContextPartyIDSource }

//BuildRelatedContextPartyIDSource returns a new RelatedContextPartyIDSource initialized with val
func BuildRelatedContextPartyIDSource(val string) RelatedContextPartyIDSource {
	var field RelatedContextPartyIDSource
	field.Value = val
	return field
}

//RelatedContextPartyRole is a INT field
type RelatedContextPartyRole struct{ message.IntValue }

//Tag returns tag.RelatedContextPartyRole (1578)
func (f RelatedContextPartyRole) Tag() fix.Tag { return tag.RelatedContextPartyRole }

//BuildRelatedContextPartyRole returns a new RelatedContextPartyRole initialized with val
func BuildRelatedContextPartyRole(val int) RelatedContextPartyRole {
	var field RelatedContextPartyRole
	field.Value = val
	return field
}

//RelatedContextPartySubID is a STRING field
type RelatedContextPartySubID struct{ message.StringValue }

//Tag returns tag.RelatedContextPartySubID (1580)
func (f RelatedContextPartySubID) Tag() fix.Tag { return tag.RelatedContextPartySubID }

//BuildRelatedContextPartySubID returns a new RelatedContextPartySubID initialized with val
func BuildRelatedContextPartySubID(val string) RelatedContextPartySubID {
	var field RelatedContextPartySubID
	field.Value = val
	return field
}

//RelatedContextPartySubIDType is a INT field
type RelatedContextPartySubIDType struct{ message.IntValue }

//Tag returns tag.RelatedContextPartySubIDType (1581)
func (f RelatedContextPartySubIDType) Tag() fix.Tag { return tag.RelatedContextPartySubIDType }

//BuildRelatedContextPartySubIDType returns a new RelatedContextPartySubIDType initialized with val
func BuildRelatedContextPartySubIDType(val int) RelatedContextPartySubIDType {
	var field RelatedContextPartySubIDType
	field.Value = val
	return field
}

//RelatedPartyAltID is a STRING field
type RelatedPartyAltID struct{ message.StringValue }

//Tag returns tag.RelatedPartyAltID (1570)
func (f RelatedPartyAltID) Tag() fix.Tag { return tag.RelatedPartyAltID }

//BuildRelatedPartyAltID returns a new RelatedPartyAltID initialized with val
func BuildRelatedPartyAltID(val string) RelatedPartyAltID {
	var field RelatedPartyAltID
	field.Value = val
	return field
}

//RelatedPartyAltIDSource is a CHAR field
type RelatedPartyAltIDSource struct{ message.CharValue }

//Tag returns tag.RelatedPartyAltIDSource (1571)
func (f RelatedPartyAltIDSource) Tag() fix.Tag { return tag.RelatedPartyAltIDSource }

//BuildRelatedPartyAltIDSource returns a new RelatedPartyAltIDSource initialized with val
func BuildRelatedPartyAltIDSource(val string) RelatedPartyAltIDSource {
	var field RelatedPartyAltIDSource
	field.Value = val
	return field
}

//RelatedPartyAltSubID is a STRING field
type RelatedPartyAltSubID struct{ message.StringValue }

//Tag returns tag.RelatedPartyAltSubID (1573)
func (f RelatedPartyAltSubID) Tag() fix.Tag { return tag.RelatedPartyAltSubID }

//BuildRelatedPartyAltSubID returns a new RelatedPartyAltSubID initialized with val
func BuildRelatedPartyAltSubID(val string) RelatedPartyAltSubID {
	var field RelatedPartyAltSubID
	field.Value = val
	return field
}

//RelatedPartyAltSubIDType is a INT field
type RelatedPartyAltSubIDType struct{ message.IntValue }

//Tag returns tag.RelatedPartyAltSubIDType (1574)
func (f RelatedPartyAltSubIDType) Tag() fix.Tag { return tag.RelatedPartyAltSubIDType }

//BuildRelatedPartyAltSubIDType returns a new RelatedPartyAltSubIDType initialized with val
func BuildRelatedPartyAltSubIDType(val int) RelatedPartyAltSubIDType {
	var field RelatedPartyAltSubIDType
	field.Value = val
	return field
}

//RelatedPartyID is a STRING field
type RelatedPartyID struct{ message.StringValue }

//Tag returns tag.RelatedPartyID (1563)
func (f RelatedPartyID) Tag() fix.Tag { return tag.RelatedPartyID }

//BuildRelatedPartyID returns a new RelatedPartyID initialized with val
func BuildRelatedPartyID(val string) RelatedPartyID {
	var field RelatedPartyID
	field.Value = val
	return field
}

//RelatedPartyIDSource is a CHAR field
type RelatedPartyIDSource struct{ message.CharValue }

//Tag returns tag.RelatedPartyIDSource (1564)
func (f RelatedPartyIDSource) Tag() fix.Tag { return tag.RelatedPartyIDSource }

//BuildRelatedPartyIDSource returns a new RelatedPartyIDSource initialized with val
func BuildRelatedPartyIDSource(val string) RelatedPartyIDSource {
	var field RelatedPartyIDSource
	field.Value = val
	return field
}

//RelatedPartyRole is a INT field
type RelatedPartyRole struct{ message.IntValue }

//Tag returns tag.RelatedPartyRole (1565)
func (f RelatedPartyRole) Tag() fix.Tag { return tag.RelatedPartyRole }

//BuildRelatedPartyRole returns a new RelatedPartyRole initialized with val
func BuildRelatedPartyRole(val int) RelatedPartyRole {
	var field RelatedPartyRole
	field.Value = val
	return field
}

//RelatedPartySubID is a STRING field
type RelatedPartySubID struct{ message.StringValue }

//Tag returns tag.RelatedPartySubID (1567)
func (f RelatedPartySubID) Tag() fix.Tag { return tag.RelatedPartySubID }

//BuildRelatedPartySubID returns a new RelatedPartySubID initialized with val
func BuildRelatedPartySubID(val string) RelatedPartySubID {
	var field RelatedPartySubID
	field.Value = val
	return field
}

//RelatedPartySubIDType is a INT field
type RelatedPartySubIDType struct{ message.IntValue }

//Tag returns tag.RelatedPartySubIDType (1568)
func (f RelatedPartySubIDType) Tag() fix.Tag { return tag.RelatedPartySubIDType }

//BuildRelatedPartySubIDType returns a new RelatedPartySubIDType initialized with val
func BuildRelatedPartySubIDType(val int) RelatedPartySubIDType {
	var field RelatedPartySubIDType
	field.Value = val
	return field
}

//RelationshipRiskCFICode is a STRING field
type RelationshipRiskCFICode struct{ message.StringValue }

//Tag returns tag.RelationshipRiskCFICode (1599)
func (f RelationshipRiskCFICode) Tag() fix.Tag { return tag.RelationshipRiskCFICode }

//BuildRelationshipRiskCFICode returns a new RelationshipRiskCFICode initialized with val
func BuildRelationshipRiskCFICode(val string) RelationshipRiskCFICode {
	var field RelationshipRiskCFICode
	field.Value = val
	return field
}

//RelationshipRiskCouponRate is a PERCENTAGE field
type RelationshipRiskCouponRate struct{ message.PercentageValue }

//Tag returns tag.RelationshipRiskCouponRate (1608)
func (f RelationshipRiskCouponRate) Tag() fix.Tag { return tag.RelationshipRiskCouponRate }

//BuildRelationshipRiskCouponRate returns a new RelationshipRiskCouponRate initialized with val
func BuildRelationshipRiskCouponRate(val float64) RelationshipRiskCouponRate {
	var field RelationshipRiskCouponRate
	field.Value = val
	return field
}

//RelationshipRiskEncodedSecurityDesc is a DATA field
type RelationshipRiskEncodedSecurityDesc struct{ message.DataValue }

//Tag returns tag.RelationshipRiskEncodedSecurityDesc (1619)
func (f RelationshipRiskEncodedSecurityDesc) Tag() fix.Tag {
	return tag.RelationshipRiskEncodedSecurityDesc
}

//BuildRelationshipRiskEncodedSecurityDesc returns a new RelationshipRiskEncodedSecurityDesc initialized with val
func BuildRelationshipRiskEncodedSecurityDesc(val string) RelationshipRiskEncodedSecurityDesc {
	var field RelationshipRiskEncodedSecurityDesc
	field.Value = val
	return field
}

//RelationshipRiskEncodedSecurityDescLen is a LENGTH field
type RelationshipRiskEncodedSecurityDescLen struct{ message.LengthValue }

//Tag returns tag.RelationshipRiskEncodedSecurityDescLen (1618)
func (f RelationshipRiskEncodedSecurityDescLen) Tag() fix.Tag {
	return tag.RelationshipRiskEncodedSecurityDescLen
}

//BuildRelationshipRiskEncodedSecurityDescLen returns a new RelationshipRiskEncodedSecurityDescLen initialized with val
func BuildRelationshipRiskEncodedSecurityDescLen(val int) RelationshipRiskEncodedSecurityDescLen {
	var field RelationshipRiskEncodedSecurityDescLen
	field.Value = val
	return field
}

//RelationshipRiskFlexibleIndicator is a BOOLEAN field
type RelationshipRiskFlexibleIndicator struct{ message.BooleanValue }

//Tag returns tag.RelationshipRiskFlexibleIndicator (1607)
func (f RelationshipRiskFlexibleIndicator) Tag() fix.Tag { return tag.RelationshipRiskFlexibleIndicator }

//BuildRelationshipRiskFlexibleIndicator returns a new RelationshipRiskFlexibleIndicator initialized with val
func BuildRelationshipRiskFlexibleIndicator(val bool) RelationshipRiskFlexibleIndicator {
	var field RelationshipRiskFlexibleIndicator
	field.Value = val
	return field
}

//RelationshipRiskInstrumentMultiplier is a FLOAT field
type RelationshipRiskInstrumentMultiplier struct{ message.FloatValue }

//Tag returns tag.RelationshipRiskInstrumentMultiplier (1612)
func (f RelationshipRiskInstrumentMultiplier) Tag() fix.Tag {
	return tag.RelationshipRiskInstrumentMultiplier
}

//BuildRelationshipRiskInstrumentMultiplier returns a new RelationshipRiskInstrumentMultiplier initialized with val
func BuildRelationshipRiskInstrumentMultiplier(val float64) RelationshipRiskInstrumentMultiplier {
	var field RelationshipRiskInstrumentMultiplier
	field.Value = val
	return field
}

//RelationshipRiskInstrumentOperator is a INT field
type RelationshipRiskInstrumentOperator struct{ message.IntValue }

//Tag returns tag.RelationshipRiskInstrumentOperator (1588)
func (f RelationshipRiskInstrumentOperator) Tag() fix.Tag {
	return tag.RelationshipRiskInstrumentOperator
}

//BuildRelationshipRiskInstrumentOperator returns a new RelationshipRiskInstrumentOperator initialized with val
func BuildRelationshipRiskInstrumentOperator(val int) RelationshipRiskInstrumentOperator {
	var field RelationshipRiskInstrumentOperator
	field.Value = val
	return field
}

//RelationshipRiskInstrumentSettlType is a STRING field
type RelationshipRiskInstrumentSettlType struct{ message.StringValue }

//Tag returns tag.RelationshipRiskInstrumentSettlType (1611)
func (f RelationshipRiskInstrumentSettlType) Tag() fix.Tag {
	return tag.RelationshipRiskInstrumentSettlType
}

//BuildRelationshipRiskInstrumentSettlType returns a new RelationshipRiskInstrumentSettlType initialized with val
func BuildRelationshipRiskInstrumentSettlType(val string) RelationshipRiskInstrumentSettlType {
	var field RelationshipRiskInstrumentSettlType
	field.Value = val
	return field
}

//RelationshipRiskLimitAmount is a AMT field
type RelationshipRiskLimitAmount struct{ message.AmtValue }

//Tag returns tag.RelationshipRiskLimitAmount (1584)
func (f RelationshipRiskLimitAmount) Tag() fix.Tag { return tag.RelationshipRiskLimitAmount }

//BuildRelationshipRiskLimitAmount returns a new RelationshipRiskLimitAmount initialized with val
func BuildRelationshipRiskLimitAmount(val float64) RelationshipRiskLimitAmount {
	var field RelationshipRiskLimitAmount
	field.Value = val
	return field
}

//RelationshipRiskLimitCurrency is a CURRENCY field
type RelationshipRiskLimitCurrency struct{ message.CurrencyValue }

//Tag returns tag.RelationshipRiskLimitCurrency (1585)
func (f RelationshipRiskLimitCurrency) Tag() fix.Tag { return tag.RelationshipRiskLimitCurrency }

//BuildRelationshipRiskLimitCurrency returns a new RelationshipRiskLimitCurrency initialized with val
func BuildRelationshipRiskLimitCurrency(val string) RelationshipRiskLimitCurrency {
	var field RelationshipRiskLimitCurrency
	field.Value = val
	return field
}

//RelationshipRiskLimitPlatform is a STRING field
type RelationshipRiskLimitPlatform struct{ message.StringValue }

//Tag returns tag.RelationshipRiskLimitPlatform (1586)
func (f RelationshipRiskLimitPlatform) Tag() fix.Tag { return tag.RelationshipRiskLimitPlatform }

//BuildRelationshipRiskLimitPlatform returns a new RelationshipRiskLimitPlatform initialized with val
func BuildRelationshipRiskLimitPlatform(val string) RelationshipRiskLimitPlatform {
	var field RelationshipRiskLimitPlatform
	field.Value = val
	return field
}

//RelationshipRiskLimitType is a INT field
type RelationshipRiskLimitType struct{ message.IntValue }

//Tag returns tag.RelationshipRiskLimitType (1583)
func (f RelationshipRiskLimitType) Tag() fix.Tag { return tag.RelationshipRiskLimitType }

//BuildRelationshipRiskLimitType returns a new RelationshipRiskLimitType initialized with val
func BuildRelationshipRiskLimitType(val int) RelationshipRiskLimitType {
	var field RelationshipRiskLimitType
	field.Value = val
	return field
}

//RelationshipRiskMaturityMonthYear is a MONTHYEAR field
type RelationshipRiskMaturityMonthYear struct{ message.MonthYearValue }

//Tag returns tag.RelationshipRiskMaturityMonthYear (1602)
func (f RelationshipRiskMaturityMonthYear) Tag() fix.Tag { return tag.RelationshipRiskMaturityMonthYear }

//BuildRelationshipRiskMaturityMonthYear returns a new RelationshipRiskMaturityMonthYear initialized with val
func BuildRelationshipRiskMaturityMonthYear(val string) RelationshipRiskMaturityMonthYear {
	var field RelationshipRiskMaturityMonthYear
	field.Value = val
	return field
}

//RelationshipRiskMaturityTime is a TZTIMEONLY field
type RelationshipRiskMaturityTime struct{ message.TZTimeOnlyValue }

//Tag returns tag.RelationshipRiskMaturityTime (1603)
func (f RelationshipRiskMaturityTime) Tag() fix.Tag { return tag.RelationshipRiskMaturityTime }

//RelationshipRiskProduct is a INT field
type RelationshipRiskProduct struct{ message.IntValue }

//Tag returns tag.RelationshipRiskProduct (1596)
func (f RelationshipRiskProduct) Tag() fix.Tag { return tag.RelationshipRiskProduct }

//BuildRelationshipRiskProduct returns a new RelationshipRiskProduct initialized with val
func BuildRelationshipRiskProduct(val int) RelationshipRiskProduct {
	var field RelationshipRiskProduct
	field.Value = val
	return field
}

//RelationshipRiskProductComplex is a STRING field
type RelationshipRiskProductComplex struct{ message.StringValue }

//Tag returns tag.RelationshipRiskProductComplex (1597)
func (f RelationshipRiskProductComplex) Tag() fix.Tag { return tag.RelationshipRiskProductComplex }

//BuildRelationshipRiskProductComplex returns a new RelationshipRiskProductComplex initialized with val
func BuildRelationshipRiskProductComplex(val string) RelationshipRiskProductComplex {
	var field RelationshipRiskProductComplex
	field.Value = val
	return field
}

//RelationshipRiskPutOrCall is a INT field
type RelationshipRiskPutOrCall struct{ message.IntValue }

//Tag returns tag.RelationshipRiskPutOrCall (1606)
func (f RelationshipRiskPutOrCall) Tag() fix.Tag { return tag.RelationshipRiskPutOrCall }

//BuildRelationshipRiskPutOrCall returns a new RelationshipRiskPutOrCall initialized with val
func BuildRelationshipRiskPutOrCall(val int) RelationshipRiskPutOrCall {
	var field RelationshipRiskPutOrCall
	field.Value = val
	return field
}

//RelationshipRiskRestructuringType is a STRING field
type RelationshipRiskRestructuringType struct{ message.StringValue }

//Tag returns tag.RelationshipRiskRestructuringType (1604)
func (f RelationshipRiskRestructuringType) Tag() fix.Tag { return tag.RelationshipRiskRestructuringType }

//BuildRelationshipRiskRestructuringType returns a new RelationshipRiskRestructuringType initialized with val
func BuildRelationshipRiskRestructuringType(val string) RelationshipRiskRestructuringType {
	var field RelationshipRiskRestructuringType
	field.Value = val
	return field
}

//RelationshipRiskSecurityAltID is a STRING field
type RelationshipRiskSecurityAltID struct{ message.StringValue }

//Tag returns tag.RelationshipRiskSecurityAltID (1594)
func (f RelationshipRiskSecurityAltID) Tag() fix.Tag { return tag.RelationshipRiskSecurityAltID }

//BuildRelationshipRiskSecurityAltID returns a new RelationshipRiskSecurityAltID initialized with val
func BuildRelationshipRiskSecurityAltID(val string) RelationshipRiskSecurityAltID {
	var field RelationshipRiskSecurityAltID
	field.Value = val
	return field
}

//RelationshipRiskSecurityAltIDSource is a STRING field
type RelationshipRiskSecurityAltIDSource struct{ message.StringValue }

//Tag returns tag.RelationshipRiskSecurityAltIDSource (1595)
func (f RelationshipRiskSecurityAltIDSource) Tag() fix.Tag {
	return tag.RelationshipRiskSecurityAltIDSource
}

//BuildRelationshipRiskSecurityAltIDSource returns a new RelationshipRiskSecurityAltIDSource initialized with val
func BuildRelationshipRiskSecurityAltIDSource(val string) RelationshipRiskSecurityAltIDSource {
	var field RelationshipRiskSecurityAltIDSource
	field.Value = val
	return field
}

//RelationshipRiskSecurityDesc is a STRING field
type RelationshipRiskSecurityDesc struct{ message.StringValue }

//Tag returns tag.RelationshipRiskSecurityDesc (1610)
func (f RelationshipRiskSecurityDesc) Tag() fix.Tag { return tag.RelationshipRiskSecurityDesc }

//BuildRelationshipRiskSecurityDesc returns a new RelationshipRiskSecurityDesc initialized with val
func BuildRelationshipRiskSecurityDesc(val string) RelationshipRiskSecurityDesc {
	var field RelationshipRiskSecurityDesc
	field.Value = val
	return field
}

//RelationshipRiskSecurityExchange is a EXCHANGE field
type RelationshipRiskSecurityExchange struct{ message.ExchangeValue }

//Tag returns tag.RelationshipRiskSecurityExchange (1609)
func (f RelationshipRiskSecurityExchange) Tag() fix.Tag { return tag.RelationshipRiskSecurityExchange }

//BuildRelationshipRiskSecurityExchange returns a new RelationshipRiskSecurityExchange initialized with val
func BuildRelationshipRiskSecurityExchange(val string) RelationshipRiskSecurityExchange {
	var field RelationshipRiskSecurityExchange
	field.Value = val
	return field
}

//RelationshipRiskSecurityGroup is a STRING field
type RelationshipRiskSecurityGroup struct{ message.StringValue }

//Tag returns tag.RelationshipRiskSecurityGroup (1598)
func (f RelationshipRiskSecurityGroup) Tag() fix.Tag { return tag.RelationshipRiskSecurityGroup }

//BuildRelationshipRiskSecurityGroup returns a new RelationshipRiskSecurityGroup initialized with val
func BuildRelationshipRiskSecurityGroup(val string) RelationshipRiskSecurityGroup {
	var field RelationshipRiskSecurityGroup
	field.Value = val
	return field
}

//RelationshipRiskSecurityID is a STRING field
type RelationshipRiskSecurityID struct{ message.StringValue }

//Tag returns tag.RelationshipRiskSecurityID (1591)
func (f RelationshipRiskSecurityID) Tag() fix.Tag { return tag.RelationshipRiskSecurityID }

//BuildRelationshipRiskSecurityID returns a new RelationshipRiskSecurityID initialized with val
func BuildRelationshipRiskSecurityID(val string) RelationshipRiskSecurityID {
	var field RelationshipRiskSecurityID
	field.Value = val
	return field
}

//RelationshipRiskSecurityIDSource is a STRING field
type RelationshipRiskSecurityIDSource struct{ message.StringValue }

//Tag returns tag.RelationshipRiskSecurityIDSource (1592)
func (f RelationshipRiskSecurityIDSource) Tag() fix.Tag { return tag.RelationshipRiskSecurityIDSource }

//BuildRelationshipRiskSecurityIDSource returns a new RelationshipRiskSecurityIDSource initialized with val
func BuildRelationshipRiskSecurityIDSource(val string) RelationshipRiskSecurityIDSource {
	var field RelationshipRiskSecurityIDSource
	field.Value = val
	return field
}

//RelationshipRiskSecuritySubType is a STRING field
type RelationshipRiskSecuritySubType struct{ message.StringValue }

//Tag returns tag.RelationshipRiskSecuritySubType (1601)
func (f RelationshipRiskSecuritySubType) Tag() fix.Tag { return tag.RelationshipRiskSecuritySubType }

//BuildRelationshipRiskSecuritySubType returns a new RelationshipRiskSecuritySubType initialized with val
func BuildRelationshipRiskSecuritySubType(val string) RelationshipRiskSecuritySubType {
	var field RelationshipRiskSecuritySubType
	field.Value = val
	return field
}

//RelationshipRiskSecurityType is a STRING field
type RelationshipRiskSecurityType struct{ message.StringValue }

//Tag returns tag.RelationshipRiskSecurityType (1600)
func (f RelationshipRiskSecurityType) Tag() fix.Tag { return tag.RelationshipRiskSecurityType }

//BuildRelationshipRiskSecurityType returns a new RelationshipRiskSecurityType initialized with val
func BuildRelationshipRiskSecurityType(val string) RelationshipRiskSecurityType {
	var field RelationshipRiskSecurityType
	field.Value = val
	return field
}

//RelationshipRiskSeniority is a STRING field
type RelationshipRiskSeniority struct{ message.StringValue }

//Tag returns tag.RelationshipRiskSeniority (1605)
func (f RelationshipRiskSeniority) Tag() fix.Tag { return tag.RelationshipRiskSeniority }

//BuildRelationshipRiskSeniority returns a new RelationshipRiskSeniority initialized with val
func BuildRelationshipRiskSeniority(val string) RelationshipRiskSeniority {
	var field RelationshipRiskSeniority
	field.Value = val
	return field
}

//RelationshipRiskSymbol is a STRING field
type RelationshipRiskSymbol struct{ message.StringValue }

//Tag returns tag.RelationshipRiskSymbol (1589)
func (f RelationshipRiskSymbol) Tag() fix.Tag { return tag.RelationshipRiskSymbol }

//BuildRelationshipRiskSymbol returns a new RelationshipRiskSymbol initialized with val
func BuildRelationshipRiskSymbol(val string) RelationshipRiskSymbol {
	var field RelationshipRiskSymbol
	field.Value = val
	return field
}

//RelationshipRiskSymbolSfx is a STRING field
type RelationshipRiskSymbolSfx struct{ message.StringValue }

//Tag returns tag.RelationshipRiskSymbolSfx (1590)
func (f RelationshipRiskSymbolSfx) Tag() fix.Tag { return tag.RelationshipRiskSymbolSfx }

//BuildRelationshipRiskSymbolSfx returns a new RelationshipRiskSymbolSfx initialized with val
func BuildRelationshipRiskSymbolSfx(val string) RelationshipRiskSymbolSfx {
	var field RelationshipRiskSymbolSfx
	field.Value = val
	return field
}

//RelationshipRiskWarningLevelName is a STRING field
type RelationshipRiskWarningLevelName struct{ message.StringValue }

//Tag returns tag.RelationshipRiskWarningLevelName (1615)
func (f RelationshipRiskWarningLevelName) Tag() fix.Tag { return tag.RelationshipRiskWarningLevelName }

//BuildRelationshipRiskWarningLevelName returns a new RelationshipRiskWarningLevelName initialized with val
func BuildRelationshipRiskWarningLevelName(val string) RelationshipRiskWarningLevelName {
	var field RelationshipRiskWarningLevelName
	field.Value = val
	return field
}

//RelationshipRiskWarningLevelPercent is a PERCENTAGE field
type RelationshipRiskWarningLevelPercent struct{ message.PercentageValue }

//Tag returns tag.RelationshipRiskWarningLevelPercent (1614)
func (f RelationshipRiskWarningLevelPercent) Tag() fix.Tag {
	return tag.RelationshipRiskWarningLevelPercent
}

//BuildRelationshipRiskWarningLevelPercent returns a new RelationshipRiskWarningLevelPercent initialized with val
func BuildRelationshipRiskWarningLevelPercent(val float64) RelationshipRiskWarningLevelPercent {
	var field RelationshipRiskWarningLevelPercent
	field.Value = val
	return field
}

//RepoCollateralSecurityType is a INT field
type RepoCollateralSecurityType struct{ message.IntValue }

//Tag returns tag.RepoCollateralSecurityType (239)
func (f RepoCollateralSecurityType) Tag() fix.Tag { return tag.RepoCollateralSecurityType }

//BuildRepoCollateralSecurityType returns a new RepoCollateralSecurityType initialized with val
func BuildRepoCollateralSecurityType(val int) RepoCollateralSecurityType {
	var field RepoCollateralSecurityType
	field.Value = val
	return field
}

//ReportToExch is a BOOLEAN field
type ReportToExch struct{ message.BooleanValue }

//Tag returns tag.ReportToExch (113)
func (f ReportToExch) Tag() fix.Tag { return tag.ReportToExch }

//BuildReportToExch returns a new ReportToExch initialized with val
func BuildReportToExch(val bool) ReportToExch {
	var field ReportToExch
	field.Value = val
	return field
}

//ReportedPx is a PRICE field
type ReportedPx struct{ message.PriceValue }

//Tag returns tag.ReportedPx (861)
func (f ReportedPx) Tag() fix.Tag { return tag.ReportedPx }

//BuildReportedPx returns a new ReportedPx initialized with val
func BuildReportedPx(val float64) ReportedPx {
	var field ReportedPx
	field.Value = val
	return field
}

//ReportedPxDiff is a BOOLEAN field
type ReportedPxDiff struct{ message.BooleanValue }

//Tag returns tag.ReportedPxDiff (1134)
func (f ReportedPxDiff) Tag() fix.Tag { return tag.ReportedPxDiff }

//BuildReportedPxDiff returns a new ReportedPxDiff initialized with val
func BuildReportedPxDiff(val bool) ReportedPxDiff {
	var field ReportedPxDiff
	field.Value = val
	return field
}

//RepurchaseRate is a PERCENTAGE field
type RepurchaseRate struct{ message.PercentageValue }

//Tag returns tag.RepurchaseRate (227)
func (f RepurchaseRate) Tag() fix.Tag { return tag.RepurchaseRate }

//BuildRepurchaseRate returns a new RepurchaseRate initialized with val
func BuildRepurchaseRate(val float64) RepurchaseRate {
	var field RepurchaseRate
	field.Value = val
	return field
}

//RepurchaseTerm is a INT field
type RepurchaseTerm struct{ message.IntValue }

//Tag returns tag.RepurchaseTerm (226)
func (f RepurchaseTerm) Tag() fix.Tag { return tag.RepurchaseTerm }

//BuildRepurchaseTerm returns a new RepurchaseTerm initialized with val
func BuildRepurchaseTerm(val int) RepurchaseTerm {
	var field RepurchaseTerm
	field.Value = val
	return field
}

//RequestedPartyRole is a INT field
type RequestedPartyRole struct{ message.IntValue }

//Tag returns tag.RequestedPartyRole (1509)
func (f RequestedPartyRole) Tag() fix.Tag { return tag.RequestedPartyRole }

//BuildRequestedPartyRole returns a new RequestedPartyRole initialized with val
func BuildRequestedPartyRole(val int) RequestedPartyRole {
	var field RequestedPartyRole
	field.Value = val
	return field
}

//ResetSeqNumFlag is a BOOLEAN field
type ResetSeqNumFlag struct{ message.BooleanValue }

//Tag returns tag.ResetSeqNumFlag (141)
func (f ResetSeqNumFlag) Tag() fix.Tag { return tag.ResetSeqNumFlag }

//BuildResetSeqNumFlag returns a new ResetSeqNumFlag initialized with val
func BuildResetSeqNumFlag(val bool) ResetSeqNumFlag {
	var field ResetSeqNumFlag
	field.Value = val
	return field
}

//RespondentType is a INT field
type RespondentType struct{ message.IntValue }

//Tag returns tag.RespondentType (1172)
func (f RespondentType) Tag() fix.Tag { return tag.RespondentType }

//BuildRespondentType returns a new RespondentType initialized with val
func BuildRespondentType(val int) RespondentType {
	var field RespondentType
	field.Value = val
	return field
}

//ResponseDestination is a STRING field
type ResponseDestination struct{ message.StringValue }

//Tag returns tag.ResponseDestination (726)
func (f ResponseDestination) Tag() fix.Tag { return tag.ResponseDestination }

//BuildResponseDestination returns a new ResponseDestination initialized with val
func BuildResponseDestination(val string) ResponseDestination {
	var field ResponseDestination
	field.Value = val
	return field
}

//ResponseTransportType is a INT field
type ResponseTransportType struct{ message.IntValue }

//Tag returns tag.ResponseTransportType (725)
func (f ResponseTransportType) Tag() fix.Tag { return tag.ResponseTransportType }

//BuildResponseTransportType returns a new ResponseTransportType initialized with val
func BuildResponseTransportType(val int) ResponseTransportType {
	var field ResponseTransportType
	field.Value = val
	return field
}

//RestructuringType is a STRING field
type RestructuringType struct{ message.StringValue }

//Tag returns tag.RestructuringType (1449)
func (f RestructuringType) Tag() fix.Tag { return tag.RestructuringType }

//BuildRestructuringType returns a new RestructuringType initialized with val
func BuildRestructuringType(val string) RestructuringType {
	var field RestructuringType
	field.Value = val
	return field
}

//ReversalIndicator is a BOOLEAN field
type ReversalIndicator struct{ message.BooleanValue }

//Tag returns tag.ReversalIndicator (700)
func (f ReversalIndicator) Tag() fix.Tag { return tag.ReversalIndicator }

//BuildReversalIndicator returns a new ReversalIndicator initialized with val
func BuildReversalIndicator(val bool) ReversalIndicator {
	var field ReversalIndicator
	field.Value = val
	return field
}

//RiskCFICode is a STRING field
type RiskCFICode struct{ message.StringValue }

//Tag returns tag.RiskCFICode (1546)
func (f RiskCFICode) Tag() fix.Tag { return tag.RiskCFICode }

//BuildRiskCFICode returns a new RiskCFICode initialized with val
func BuildRiskCFICode(val string) RiskCFICode {
	var field RiskCFICode
	field.Value = val
	return field
}

//RiskCouponRate is a PERCENTAGE field
type RiskCouponRate struct{ message.PercentageValue }

//Tag returns tag.RiskCouponRate (1555)
func (f RiskCouponRate) Tag() fix.Tag { return tag.RiskCouponRate }

//BuildRiskCouponRate returns a new RiskCouponRate initialized with val
func BuildRiskCouponRate(val float64) RiskCouponRate {
	var field RiskCouponRate
	field.Value = val
	return field
}

//RiskEncodedSecurityDesc is a DATA field
type RiskEncodedSecurityDesc struct{ message.DataValue }

//Tag returns tag.RiskEncodedSecurityDesc (1621)
func (f RiskEncodedSecurityDesc) Tag() fix.Tag { return tag.RiskEncodedSecurityDesc }

//BuildRiskEncodedSecurityDesc returns a new RiskEncodedSecurityDesc initialized with val
func BuildRiskEncodedSecurityDesc(val string) RiskEncodedSecurityDesc {
	var field RiskEncodedSecurityDesc
	field.Value = val
	return field
}

//RiskEncodedSecurityDescLen is a LENGTH field
type RiskEncodedSecurityDescLen struct{ message.LengthValue }

//Tag returns tag.RiskEncodedSecurityDescLen (1620)
func (f RiskEncodedSecurityDescLen) Tag() fix.Tag { return tag.RiskEncodedSecurityDescLen }

//BuildRiskEncodedSecurityDescLen returns a new RiskEncodedSecurityDescLen initialized with val
func BuildRiskEncodedSecurityDescLen(val int) RiskEncodedSecurityDescLen {
	var field RiskEncodedSecurityDescLen
	field.Value = val
	return field
}

//RiskFlexibleIndicator is a BOOLEAN field
type RiskFlexibleIndicator struct{ message.BooleanValue }

//Tag returns tag.RiskFlexibleIndicator (1554)
func (f RiskFlexibleIndicator) Tag() fix.Tag { return tag.RiskFlexibleIndicator }

//BuildRiskFlexibleIndicator returns a new RiskFlexibleIndicator initialized with val
func BuildRiskFlexibleIndicator(val bool) RiskFlexibleIndicator {
	var field RiskFlexibleIndicator
	field.Value = val
	return field
}

//RiskFreeRate is a FLOAT field
type RiskFreeRate struct{ message.FloatValue }

//Tag returns tag.RiskFreeRate (1190)
func (f RiskFreeRate) Tag() fix.Tag { return tag.RiskFreeRate }

//BuildRiskFreeRate returns a new RiskFreeRate initialized with val
func BuildRiskFreeRate(val float64) RiskFreeRate {
	var field RiskFreeRate
	field.Value = val
	return field
}

//RiskInstrumentMultiplier is a FLOAT field
type RiskInstrumentMultiplier struct{ message.FloatValue }

//Tag returns tag.RiskInstrumentMultiplier (1558)
func (f RiskInstrumentMultiplier) Tag() fix.Tag { return tag.RiskInstrumentMultiplier }

//BuildRiskInstrumentMultiplier returns a new RiskInstrumentMultiplier initialized with val
func BuildRiskInstrumentMultiplier(val float64) RiskInstrumentMultiplier {
	var field RiskInstrumentMultiplier
	field.Value = val
	return field
}

//RiskInstrumentOperator is a INT field
type RiskInstrumentOperator struct{ message.IntValue }

//Tag returns tag.RiskInstrumentOperator (1535)
func (f RiskInstrumentOperator) Tag() fix.Tag { return tag.RiskInstrumentOperator }

//BuildRiskInstrumentOperator returns a new RiskInstrumentOperator initialized with val
func BuildRiskInstrumentOperator(val int) RiskInstrumentOperator {
	var field RiskInstrumentOperator
	field.Value = val
	return field
}

//RiskInstrumentSettlType is a STRING field
type RiskInstrumentSettlType struct{ message.StringValue }

//Tag returns tag.RiskInstrumentSettlType (1557)
func (f RiskInstrumentSettlType) Tag() fix.Tag { return tag.RiskInstrumentSettlType }

//BuildRiskInstrumentSettlType returns a new RiskInstrumentSettlType initialized with val
func BuildRiskInstrumentSettlType(val string) RiskInstrumentSettlType {
	var field RiskInstrumentSettlType
	field.Value = val
	return field
}

//RiskLimitAmount is a AMT field
type RiskLimitAmount struct{ message.AmtValue }

//Tag returns tag.RiskLimitAmount (1531)
func (f RiskLimitAmount) Tag() fix.Tag { return tag.RiskLimitAmount }

//BuildRiskLimitAmount returns a new RiskLimitAmount initialized with val
func BuildRiskLimitAmount(val float64) RiskLimitAmount {
	var field RiskLimitAmount
	field.Value = val
	return field
}

//RiskLimitCurrency is a CURRENCY field
type RiskLimitCurrency struct{ message.CurrencyValue }

//Tag returns tag.RiskLimitCurrency (1532)
func (f RiskLimitCurrency) Tag() fix.Tag { return tag.RiskLimitCurrency }

//BuildRiskLimitCurrency returns a new RiskLimitCurrency initialized with val
func BuildRiskLimitCurrency(val string) RiskLimitCurrency {
	var field RiskLimitCurrency
	field.Value = val
	return field
}

//RiskLimitPlatform is a STRING field
type RiskLimitPlatform struct{ message.StringValue }

//Tag returns tag.RiskLimitPlatform (1533)
func (f RiskLimitPlatform) Tag() fix.Tag { return tag.RiskLimitPlatform }

//BuildRiskLimitPlatform returns a new RiskLimitPlatform initialized with val
func BuildRiskLimitPlatform(val string) RiskLimitPlatform {
	var field RiskLimitPlatform
	field.Value = val
	return field
}

//RiskLimitType is a INT field
type RiskLimitType struct{ message.IntValue }

//Tag returns tag.RiskLimitType (1530)
func (f RiskLimitType) Tag() fix.Tag { return tag.RiskLimitType }

//BuildRiskLimitType returns a new RiskLimitType initialized with val
func BuildRiskLimitType(val int) RiskLimitType {
	var field RiskLimitType
	field.Value = val
	return field
}

//RiskMaturityMonthYear is a MONTHYEAR field
type RiskMaturityMonthYear struct{ message.MonthYearValue }

//Tag returns tag.RiskMaturityMonthYear (1549)
func (f RiskMaturityMonthYear) Tag() fix.Tag { return tag.RiskMaturityMonthYear }

//BuildRiskMaturityMonthYear returns a new RiskMaturityMonthYear initialized with val
func BuildRiskMaturityMonthYear(val string) RiskMaturityMonthYear {
	var field RiskMaturityMonthYear
	field.Value = val
	return field
}

//RiskMaturityTime is a TZTIMEONLY field
type RiskMaturityTime struct{ message.TZTimeOnlyValue }

//Tag returns tag.RiskMaturityTime (1550)
func (f RiskMaturityTime) Tag() fix.Tag { return tag.RiskMaturityTime }

//RiskProduct is a INT field
type RiskProduct struct{ message.IntValue }

//Tag returns tag.RiskProduct (1543)
func (f RiskProduct) Tag() fix.Tag { return tag.RiskProduct }

//BuildRiskProduct returns a new RiskProduct initialized with val
func BuildRiskProduct(val int) RiskProduct {
	var field RiskProduct
	field.Value = val
	return field
}

//RiskProductComplex is a STRING field
type RiskProductComplex struct{ message.StringValue }

//Tag returns tag.RiskProductComplex (1544)
func (f RiskProductComplex) Tag() fix.Tag { return tag.RiskProductComplex }

//BuildRiskProductComplex returns a new RiskProductComplex initialized with val
func BuildRiskProductComplex(val string) RiskProductComplex {
	var field RiskProductComplex
	field.Value = val
	return field
}

//RiskPutOrCall is a INT field
type RiskPutOrCall struct{ message.IntValue }

//Tag returns tag.RiskPutOrCall (1553)
func (f RiskPutOrCall) Tag() fix.Tag { return tag.RiskPutOrCall }

//BuildRiskPutOrCall returns a new RiskPutOrCall initialized with val
func BuildRiskPutOrCall(val int) RiskPutOrCall {
	var field RiskPutOrCall
	field.Value = val
	return field
}

//RiskRestructuringType is a STRING field
type RiskRestructuringType struct{ message.StringValue }

//Tag returns tag.RiskRestructuringType (1551)
func (f RiskRestructuringType) Tag() fix.Tag { return tag.RiskRestructuringType }

//BuildRiskRestructuringType returns a new RiskRestructuringType initialized with val
func BuildRiskRestructuringType(val string) RiskRestructuringType {
	var field RiskRestructuringType
	field.Value = val
	return field
}

//RiskSecurityAltID is a STRING field
type RiskSecurityAltID struct{ message.StringValue }

//Tag returns tag.RiskSecurityAltID (1541)
func (f RiskSecurityAltID) Tag() fix.Tag { return tag.RiskSecurityAltID }

//BuildRiskSecurityAltID returns a new RiskSecurityAltID initialized with val
func BuildRiskSecurityAltID(val string) RiskSecurityAltID {
	var field RiskSecurityAltID
	field.Value = val
	return field
}

//RiskSecurityAltIDSource is a STRING field
type RiskSecurityAltIDSource struct{ message.StringValue }

//Tag returns tag.RiskSecurityAltIDSource (1542)
func (f RiskSecurityAltIDSource) Tag() fix.Tag { return tag.RiskSecurityAltIDSource }

//BuildRiskSecurityAltIDSource returns a new RiskSecurityAltIDSource initialized with val
func BuildRiskSecurityAltIDSource(val string) RiskSecurityAltIDSource {
	var field RiskSecurityAltIDSource
	field.Value = val
	return field
}

//RiskSecurityDesc is a STRING field
type RiskSecurityDesc struct{ message.StringValue }

//Tag returns tag.RiskSecurityDesc (1556)
func (f RiskSecurityDesc) Tag() fix.Tag { return tag.RiskSecurityDesc }

//BuildRiskSecurityDesc returns a new RiskSecurityDesc initialized with val
func BuildRiskSecurityDesc(val string) RiskSecurityDesc {
	var field RiskSecurityDesc
	field.Value = val
	return field
}

//RiskSecurityExchange is a EXCHANGE field
type RiskSecurityExchange struct{ message.ExchangeValue }

//Tag returns tag.RiskSecurityExchange (1616)
func (f RiskSecurityExchange) Tag() fix.Tag { return tag.RiskSecurityExchange }

//BuildRiskSecurityExchange returns a new RiskSecurityExchange initialized with val
func BuildRiskSecurityExchange(val string) RiskSecurityExchange {
	var field RiskSecurityExchange
	field.Value = val
	return field
}

//RiskSecurityGroup is a STRING field
type RiskSecurityGroup struct{ message.StringValue }

//Tag returns tag.RiskSecurityGroup (1545)
func (f RiskSecurityGroup) Tag() fix.Tag { return tag.RiskSecurityGroup }

//BuildRiskSecurityGroup returns a new RiskSecurityGroup initialized with val
func BuildRiskSecurityGroup(val string) RiskSecurityGroup {
	var field RiskSecurityGroup
	field.Value = val
	return field
}

//RiskSecurityID is a STRING field
type RiskSecurityID struct{ message.StringValue }

//Tag returns tag.RiskSecurityID (1538)
func (f RiskSecurityID) Tag() fix.Tag { return tag.RiskSecurityID }

//BuildRiskSecurityID returns a new RiskSecurityID initialized with val
func BuildRiskSecurityID(val string) RiskSecurityID {
	var field RiskSecurityID
	field.Value = val
	return field
}

//RiskSecurityIDSource is a STRING field
type RiskSecurityIDSource struct{ message.StringValue }

//Tag returns tag.RiskSecurityIDSource (1539)
func (f RiskSecurityIDSource) Tag() fix.Tag { return tag.RiskSecurityIDSource }

//BuildRiskSecurityIDSource returns a new RiskSecurityIDSource initialized with val
func BuildRiskSecurityIDSource(val string) RiskSecurityIDSource {
	var field RiskSecurityIDSource
	field.Value = val
	return field
}

//RiskSecuritySubType is a STRING field
type RiskSecuritySubType struct{ message.StringValue }

//Tag returns tag.RiskSecuritySubType (1548)
func (f RiskSecuritySubType) Tag() fix.Tag { return tag.RiskSecuritySubType }

//BuildRiskSecuritySubType returns a new RiskSecuritySubType initialized with val
func BuildRiskSecuritySubType(val string) RiskSecuritySubType {
	var field RiskSecuritySubType
	field.Value = val
	return field
}

//RiskSecurityType is a STRING field
type RiskSecurityType struct{ message.StringValue }

//Tag returns tag.RiskSecurityType (1547)
func (f RiskSecurityType) Tag() fix.Tag { return tag.RiskSecurityType }

//BuildRiskSecurityType returns a new RiskSecurityType initialized with val
func BuildRiskSecurityType(val string) RiskSecurityType {
	var field RiskSecurityType
	field.Value = val
	return field
}

//RiskSeniority is a STRING field
type RiskSeniority struct{ message.StringValue }

//Tag returns tag.RiskSeniority (1552)
func (f RiskSeniority) Tag() fix.Tag { return tag.RiskSeniority }

//BuildRiskSeniority returns a new RiskSeniority initialized with val
func BuildRiskSeniority(val string) RiskSeniority {
	var field RiskSeniority
	field.Value = val
	return field
}

//RiskSymbol is a STRING field
type RiskSymbol struct{ message.StringValue }

//Tag returns tag.RiskSymbol (1536)
func (f RiskSymbol) Tag() fix.Tag { return tag.RiskSymbol }

//BuildRiskSymbol returns a new RiskSymbol initialized with val
func BuildRiskSymbol(val string) RiskSymbol {
	var field RiskSymbol
	field.Value = val
	return field
}

//RiskSymbolSfx is a STRING field
type RiskSymbolSfx struct{ message.StringValue }

//Tag returns tag.RiskSymbolSfx (1537)
func (f RiskSymbolSfx) Tag() fix.Tag { return tag.RiskSymbolSfx }

//BuildRiskSymbolSfx returns a new RiskSymbolSfx initialized with val
func BuildRiskSymbolSfx(val string) RiskSymbolSfx {
	var field RiskSymbolSfx
	field.Value = val
	return field
}

//RiskWarningLevelName is a STRING field
type RiskWarningLevelName struct{ message.StringValue }

//Tag returns tag.RiskWarningLevelName (1561)
func (f RiskWarningLevelName) Tag() fix.Tag { return tag.RiskWarningLevelName }

//BuildRiskWarningLevelName returns a new RiskWarningLevelName initialized with val
func BuildRiskWarningLevelName(val string) RiskWarningLevelName {
	var field RiskWarningLevelName
	field.Value = val
	return field
}

//RiskWarningLevelPercent is a PERCENTAGE field
type RiskWarningLevelPercent struct{ message.PercentageValue }

//Tag returns tag.RiskWarningLevelPercent (1560)
func (f RiskWarningLevelPercent) Tag() fix.Tag { return tag.RiskWarningLevelPercent }

//BuildRiskWarningLevelPercent returns a new RiskWarningLevelPercent initialized with val
func BuildRiskWarningLevelPercent(val float64) RiskWarningLevelPercent {
	var field RiskWarningLevelPercent
	field.Value = val
	return field
}

//RndPx is a PRICE field
type RndPx struct{ message.PriceValue }

//Tag returns tag.RndPx (991)
func (f RndPx) Tag() fix.Tag { return tag.RndPx }

//BuildRndPx returns a new RndPx initialized with val
func BuildRndPx(val float64) RndPx {
	var field RndPx
	field.Value = val
	return field
}

//RootPartyID is a STRING field
type RootPartyID struct{ message.StringValue }

//Tag returns tag.RootPartyID (1117)
func (f RootPartyID) Tag() fix.Tag { return tag.RootPartyID }

//BuildRootPartyID returns a new RootPartyID initialized with val
func BuildRootPartyID(val string) RootPartyID {
	var field RootPartyID
	field.Value = val
	return field
}

//RootPartyIDSource is a CHAR field
type RootPartyIDSource struct{ message.CharValue }

//Tag returns tag.RootPartyIDSource (1118)
func (f RootPartyIDSource) Tag() fix.Tag { return tag.RootPartyIDSource }

//BuildRootPartyIDSource returns a new RootPartyIDSource initialized with val
func BuildRootPartyIDSource(val string) RootPartyIDSource {
	var field RootPartyIDSource
	field.Value = val
	return field
}

//RootPartyRole is a INT field
type RootPartyRole struct{ message.IntValue }

//Tag returns tag.RootPartyRole (1119)
func (f RootPartyRole) Tag() fix.Tag { return tag.RootPartyRole }

//BuildRootPartyRole returns a new RootPartyRole initialized with val
func BuildRootPartyRole(val int) RootPartyRole {
	var field RootPartyRole
	field.Value = val
	return field
}

//RootPartySubID is a STRING field
type RootPartySubID struct{ message.StringValue }

//Tag returns tag.RootPartySubID (1121)
func (f RootPartySubID) Tag() fix.Tag { return tag.RootPartySubID }

//BuildRootPartySubID returns a new RootPartySubID initialized with val
func BuildRootPartySubID(val string) RootPartySubID {
	var field RootPartySubID
	field.Value = val
	return field
}

//RootPartySubIDType is a INT field
type RootPartySubIDType struct{ message.IntValue }

//Tag returns tag.RootPartySubIDType (1122)
func (f RootPartySubIDType) Tag() fix.Tag { return tag.RootPartySubIDType }

//BuildRootPartySubIDType returns a new RootPartySubIDType initialized with val
func BuildRootPartySubIDType(val int) RootPartySubIDType {
	var field RootPartySubIDType
	field.Value = val
	return field
}

//RoundLot is a QTY field
type RoundLot struct{ message.QtyValue }

//Tag returns tag.RoundLot (561)
func (f RoundLot) Tag() fix.Tag { return tag.RoundLot }

//BuildRoundLot returns a new RoundLot initialized with val
func BuildRoundLot(val float64) RoundLot {
	var field RoundLot
	field.Value = val
	return field
}

//RoundingDirection is a CHAR field
type RoundingDirection struct{ message.CharValue }

//Tag returns tag.RoundingDirection (468)
func (f RoundingDirection) Tag() fix.Tag { return tag.RoundingDirection }

//BuildRoundingDirection returns a new RoundingDirection initialized with val
func BuildRoundingDirection(val string) RoundingDirection {
	var field RoundingDirection
	field.Value = val
	return field
}

//RoundingModulus is a FLOAT field
type RoundingModulus struct{ message.FloatValue }

//Tag returns tag.RoundingModulus (469)
func (f RoundingModulus) Tag() fix.Tag { return tag.RoundingModulus }

//BuildRoundingModulus returns a new RoundingModulus initialized with val
func BuildRoundingModulus(val float64) RoundingModulus {
	var field RoundingModulus
	field.Value = val
	return field
}

//RoutingID is a STRING field
type RoutingID struct{ message.StringValue }

//Tag returns tag.RoutingID (217)
func (f RoutingID) Tag() fix.Tag { return tag.RoutingID }

//BuildRoutingID returns a new RoutingID initialized with val
func BuildRoutingID(val string) RoutingID {
	var field RoutingID
	field.Value = val
	return field
}

//RoutingType is a INT field
type RoutingType struct{ message.IntValue }

//Tag returns tag.RoutingType (216)
func (f RoutingType) Tag() fix.Tag { return tag.RoutingType }

//BuildRoutingType returns a new RoutingType initialized with val
func BuildRoutingType(val int) RoutingType {
	var field RoutingType
	field.Value = val
	return field
}

//RptSeq is a INT field
type RptSeq struct{ message.IntValue }

//Tag returns tag.RptSeq (83)
func (f RptSeq) Tag() fix.Tag { return tag.RptSeq }

//BuildRptSeq returns a new RptSeq initialized with val
func BuildRptSeq(val int) RptSeq {
	var field RptSeq
	field.Value = val
	return field
}

//RptSys is a STRING field
type RptSys struct{ message.StringValue }

//Tag returns tag.RptSys (1135)
func (f RptSys) Tag() fix.Tag { return tag.RptSys }

//BuildRptSys returns a new RptSys initialized with val
func BuildRptSys(val string) RptSys {
	var field RptSys
	field.Value = val
	return field
}

//Rule80A is a CHAR field
type Rule80A struct{ message.CharValue }

//Tag returns tag.Rule80A (47)
func (f Rule80A) Tag() fix.Tag { return tag.Rule80A }

//BuildRule80A returns a new Rule80A initialized with val
func BuildRule80A(val string) Rule80A {
	var field Rule80A
	field.Value = val
	return field
}

//Scope is a MULTIPLECHARVALUE field
type Scope struct{ message.MultipleCharValue }

//Tag returns tag.Scope (546)
func (f Scope) Tag() fix.Tag { return tag.Scope }

//BuildScope returns a new Scope initialized with val
func BuildScope(val string) Scope {
	var field Scope
	field.Value = val
	return field
}

//SecDefStatus is a INT field
type SecDefStatus struct{ message.IntValue }

//Tag returns tag.SecDefStatus (653)
func (f SecDefStatus) Tag() fix.Tag { return tag.SecDefStatus }

//BuildSecDefStatus returns a new SecDefStatus initialized with val
func BuildSecDefStatus(val int) SecDefStatus {
	var field SecDefStatus
	field.Value = val
	return field
}

//SecondaryAllocID is a STRING field
type SecondaryAllocID struct{ message.StringValue }

//Tag returns tag.SecondaryAllocID (793)
func (f SecondaryAllocID) Tag() fix.Tag { return tag.SecondaryAllocID }

//BuildSecondaryAllocID returns a new SecondaryAllocID initialized with val
func BuildSecondaryAllocID(val string) SecondaryAllocID {
	var field SecondaryAllocID
	field.Value = val
	return field
}

//SecondaryClOrdID is a STRING field
type SecondaryClOrdID struct{ message.StringValue }

//Tag returns tag.SecondaryClOrdID (526)
func (f SecondaryClOrdID) Tag() fix.Tag { return tag.SecondaryClOrdID }

//BuildSecondaryClOrdID returns a new SecondaryClOrdID initialized with val
func BuildSecondaryClOrdID(val string) SecondaryClOrdID {
	var field SecondaryClOrdID
	field.Value = val
	return field
}

//SecondaryDisplayQty is a QTY field
type SecondaryDisplayQty struct{ message.QtyValue }

//Tag returns tag.SecondaryDisplayQty (1082)
func (f SecondaryDisplayQty) Tag() fix.Tag { return tag.SecondaryDisplayQty }

//BuildSecondaryDisplayQty returns a new SecondaryDisplayQty initialized with val
func BuildSecondaryDisplayQty(val float64) SecondaryDisplayQty {
	var field SecondaryDisplayQty
	field.Value = val
	return field
}

//SecondaryExecID is a STRING field
type SecondaryExecID struct{ message.StringValue }

//Tag returns tag.SecondaryExecID (527)
func (f SecondaryExecID) Tag() fix.Tag { return tag.SecondaryExecID }

//BuildSecondaryExecID returns a new SecondaryExecID initialized with val
func BuildSecondaryExecID(val string) SecondaryExecID {
	var field SecondaryExecID
	field.Value = val
	return field
}

//SecondaryFirmTradeID is a STRING field
type SecondaryFirmTradeID struct{ message.StringValue }

//Tag returns tag.SecondaryFirmTradeID (1042)
func (f SecondaryFirmTradeID) Tag() fix.Tag { return tag.SecondaryFirmTradeID }

//BuildSecondaryFirmTradeID returns a new SecondaryFirmTradeID initialized with val
func BuildSecondaryFirmTradeID(val string) SecondaryFirmTradeID {
	var field SecondaryFirmTradeID
	field.Value = val
	return field
}

//SecondaryHighLimitPrice is a PRICE field
type SecondaryHighLimitPrice struct{ message.PriceValue }

//Tag returns tag.SecondaryHighLimitPrice (1230)
func (f SecondaryHighLimitPrice) Tag() fix.Tag { return tag.SecondaryHighLimitPrice }

//BuildSecondaryHighLimitPrice returns a new SecondaryHighLimitPrice initialized with val
func BuildSecondaryHighLimitPrice(val float64) SecondaryHighLimitPrice {
	var field SecondaryHighLimitPrice
	field.Value = val
	return field
}

//SecondaryIndividualAllocID is a STRING field
type SecondaryIndividualAllocID struct{ message.StringValue }

//Tag returns tag.SecondaryIndividualAllocID (989)
func (f SecondaryIndividualAllocID) Tag() fix.Tag { return tag.SecondaryIndividualAllocID }

//BuildSecondaryIndividualAllocID returns a new SecondaryIndividualAllocID initialized with val
func BuildSecondaryIndividualAllocID(val string) SecondaryIndividualAllocID {
	var field SecondaryIndividualAllocID
	field.Value = val
	return field
}

//SecondaryLowLimitPrice is a PRICE field
type SecondaryLowLimitPrice struct{ message.PriceValue }

//Tag returns tag.SecondaryLowLimitPrice (1221)
func (f SecondaryLowLimitPrice) Tag() fix.Tag { return tag.SecondaryLowLimitPrice }

//BuildSecondaryLowLimitPrice returns a new SecondaryLowLimitPrice initialized with val
func BuildSecondaryLowLimitPrice(val float64) SecondaryLowLimitPrice {
	var field SecondaryLowLimitPrice
	field.Value = val
	return field
}

//SecondaryOrderID is a STRING field
type SecondaryOrderID struct{ message.StringValue }

//Tag returns tag.SecondaryOrderID (198)
func (f SecondaryOrderID) Tag() fix.Tag { return tag.SecondaryOrderID }

//BuildSecondaryOrderID returns a new SecondaryOrderID initialized with val
func BuildSecondaryOrderID(val string) SecondaryOrderID {
	var field SecondaryOrderID
	field.Value = val
	return field
}

//SecondaryPriceLimitType is a INT field
type SecondaryPriceLimitType struct{ message.IntValue }

//Tag returns tag.SecondaryPriceLimitType (1305)
func (f SecondaryPriceLimitType) Tag() fix.Tag { return tag.SecondaryPriceLimitType }

//BuildSecondaryPriceLimitType returns a new SecondaryPriceLimitType initialized with val
func BuildSecondaryPriceLimitType(val int) SecondaryPriceLimitType {
	var field SecondaryPriceLimitType
	field.Value = val
	return field
}

//SecondaryTradeID is a STRING field
type SecondaryTradeID struct{ message.StringValue }

//Tag returns tag.SecondaryTradeID (1040)
func (f SecondaryTradeID) Tag() fix.Tag { return tag.SecondaryTradeID }

//BuildSecondaryTradeID returns a new SecondaryTradeID initialized with val
func BuildSecondaryTradeID(val string) SecondaryTradeID {
	var field SecondaryTradeID
	field.Value = val
	return field
}

//SecondaryTradeReportID is a STRING field
type SecondaryTradeReportID struct{ message.StringValue }

//Tag returns tag.SecondaryTradeReportID (818)
func (f SecondaryTradeReportID) Tag() fix.Tag { return tag.SecondaryTradeReportID }

//BuildSecondaryTradeReportID returns a new SecondaryTradeReportID initialized with val
func BuildSecondaryTradeReportID(val string) SecondaryTradeReportID {
	var field SecondaryTradeReportID
	field.Value = val
	return field
}

//SecondaryTradeReportRefID is a STRING field
type SecondaryTradeReportRefID struct{ message.StringValue }

//Tag returns tag.SecondaryTradeReportRefID (881)
func (f SecondaryTradeReportRefID) Tag() fix.Tag { return tag.SecondaryTradeReportRefID }

//BuildSecondaryTradeReportRefID returns a new SecondaryTradeReportRefID initialized with val
func BuildSecondaryTradeReportRefID(val string) SecondaryTradeReportRefID {
	var field SecondaryTradeReportRefID
	field.Value = val
	return field
}

//SecondaryTradingReferencePrice is a PRICE field
type SecondaryTradingReferencePrice struct{ message.PriceValue }

//Tag returns tag.SecondaryTradingReferencePrice (1240)
func (f SecondaryTradingReferencePrice) Tag() fix.Tag { return tag.SecondaryTradingReferencePrice }

//BuildSecondaryTradingReferencePrice returns a new SecondaryTradingReferencePrice initialized with val
func BuildSecondaryTradingReferencePrice(val float64) SecondaryTradingReferencePrice {
	var field SecondaryTradingReferencePrice
	field.Value = val
	return field
}

//SecondaryTrdType is a INT field
type SecondaryTrdType struct{ message.IntValue }

//Tag returns tag.SecondaryTrdType (855)
func (f SecondaryTrdType) Tag() fix.Tag { return tag.SecondaryTrdType }

//BuildSecondaryTrdType returns a new SecondaryTrdType initialized with val
func BuildSecondaryTrdType(val int) SecondaryTrdType {
	var field SecondaryTrdType
	field.Value = val
	return field
}

//SecureData is a DATA field
type SecureData struct{ message.DataValue }

//Tag returns tag.SecureData (91)
func (f SecureData) Tag() fix.Tag { return tag.SecureData }

//BuildSecureData returns a new SecureData initialized with val
func BuildSecureData(val string) SecureData {
	var field SecureData
	field.Value = val
	return field
}

//SecureDataLen is a LENGTH field
type SecureDataLen struct{ message.LengthValue }

//Tag returns tag.SecureDataLen (90)
func (f SecureDataLen) Tag() fix.Tag { return tag.SecureDataLen }

//BuildSecureDataLen returns a new SecureDataLen initialized with val
func BuildSecureDataLen(val int) SecureDataLen {
	var field SecureDataLen
	field.Value = val
	return field
}

//SecurityAltID is a STRING field
type SecurityAltID struct{ message.StringValue }

//Tag returns tag.SecurityAltID (455)
func (f SecurityAltID) Tag() fix.Tag { return tag.SecurityAltID }

//BuildSecurityAltID returns a new SecurityAltID initialized with val
func BuildSecurityAltID(val string) SecurityAltID {
	var field SecurityAltID
	field.Value = val
	return field
}

//SecurityAltIDSource is a STRING field
type SecurityAltIDSource struct{ message.StringValue }

//Tag returns tag.SecurityAltIDSource (456)
func (f SecurityAltIDSource) Tag() fix.Tag { return tag.SecurityAltIDSource }

//BuildSecurityAltIDSource returns a new SecurityAltIDSource initialized with val
func BuildSecurityAltIDSource(val string) SecurityAltIDSource {
	var field SecurityAltIDSource
	field.Value = val
	return field
}

//SecurityDesc is a STRING field
type SecurityDesc struct{ message.StringValue }

//Tag returns tag.SecurityDesc (107)
func (f SecurityDesc) Tag() fix.Tag { return tag.SecurityDesc }

//BuildSecurityDesc returns a new SecurityDesc initialized with val
func BuildSecurityDesc(val string) SecurityDesc {
	var field SecurityDesc
	field.Value = val
	return field
}

//SecurityExchange is a EXCHANGE field
type SecurityExchange struct{ message.ExchangeValue }

//Tag returns tag.SecurityExchange (207)
func (f SecurityExchange) Tag() fix.Tag { return tag.SecurityExchange }

//BuildSecurityExchange returns a new SecurityExchange initialized with val
func BuildSecurityExchange(val string) SecurityExchange {
	var field SecurityExchange
	field.Value = val
	return field
}

//SecurityGroup is a STRING field
type SecurityGroup struct{ message.StringValue }

//Tag returns tag.SecurityGroup (1151)
func (f SecurityGroup) Tag() fix.Tag { return tag.SecurityGroup }

//BuildSecurityGroup returns a new SecurityGroup initialized with val
func BuildSecurityGroup(val string) SecurityGroup {
	var field SecurityGroup
	field.Value = val
	return field
}

//SecurityID is a STRING field
type SecurityID struct{ message.StringValue }

//Tag returns tag.SecurityID (48)
func (f SecurityID) Tag() fix.Tag { return tag.SecurityID }

//BuildSecurityID returns a new SecurityID initialized with val
func BuildSecurityID(val string) SecurityID {
	var field SecurityID
	field.Value = val
	return field
}

//SecurityIDSource is a STRING field
type SecurityIDSource struct{ message.StringValue }

//Tag returns tag.SecurityIDSource (22)
func (f SecurityIDSource) Tag() fix.Tag { return tag.SecurityIDSource }

//BuildSecurityIDSource returns a new SecurityIDSource initialized with val
func BuildSecurityIDSource(val string) SecurityIDSource {
	var field SecurityIDSource
	field.Value = val
	return field
}

//SecurityListDesc is a STRING field
type SecurityListDesc struct{ message.StringValue }

//Tag returns tag.SecurityListDesc (1467)
func (f SecurityListDesc) Tag() fix.Tag { return tag.SecurityListDesc }

//BuildSecurityListDesc returns a new SecurityListDesc initialized with val
func BuildSecurityListDesc(val string) SecurityListDesc {
	var field SecurityListDesc
	field.Value = val
	return field
}

//SecurityListID is a STRING field
type SecurityListID struct{ message.StringValue }

//Tag returns tag.SecurityListID (1465)
func (f SecurityListID) Tag() fix.Tag { return tag.SecurityListID }

//BuildSecurityListID returns a new SecurityListID initialized with val
func BuildSecurityListID(val string) SecurityListID {
	var field SecurityListID
	field.Value = val
	return field
}

//SecurityListRefID is a STRING field
type SecurityListRefID struct{ message.StringValue }

//Tag returns tag.SecurityListRefID (1466)
func (f SecurityListRefID) Tag() fix.Tag { return tag.SecurityListRefID }

//BuildSecurityListRefID returns a new SecurityListRefID initialized with val
func BuildSecurityListRefID(val string) SecurityListRefID {
	var field SecurityListRefID
	field.Value = val
	return field
}

//SecurityListRequestType is a INT field
type SecurityListRequestType struct{ message.IntValue }

//Tag returns tag.SecurityListRequestType (559)
func (f SecurityListRequestType) Tag() fix.Tag { return tag.SecurityListRequestType }

//BuildSecurityListRequestType returns a new SecurityListRequestType initialized with val
func BuildSecurityListRequestType(val int) SecurityListRequestType {
	var field SecurityListRequestType
	field.Value = val
	return field
}

//SecurityListType is a INT field
type SecurityListType struct{ message.IntValue }

//Tag returns tag.SecurityListType (1470)
func (f SecurityListType) Tag() fix.Tag { return tag.SecurityListType }

//BuildSecurityListType returns a new SecurityListType initialized with val
func BuildSecurityListType(val int) SecurityListType {
	var field SecurityListType
	field.Value = val
	return field
}

//SecurityListTypeSource is a INT field
type SecurityListTypeSource struct{ message.IntValue }

//Tag returns tag.SecurityListTypeSource (1471)
func (f SecurityListTypeSource) Tag() fix.Tag { return tag.SecurityListTypeSource }

//BuildSecurityListTypeSource returns a new SecurityListTypeSource initialized with val
func BuildSecurityListTypeSource(val int) SecurityListTypeSource {
	var field SecurityListTypeSource
	field.Value = val
	return field
}

//SecurityReportID is a INT field
type SecurityReportID struct{ message.IntValue }

//Tag returns tag.SecurityReportID (964)
func (f SecurityReportID) Tag() fix.Tag { return tag.SecurityReportID }

//BuildSecurityReportID returns a new SecurityReportID initialized with val
func BuildSecurityReportID(val int) SecurityReportID {
	var field SecurityReportID
	field.Value = val
	return field
}

//SecurityReqID is a STRING field
type SecurityReqID struct{ message.StringValue }

//Tag returns tag.SecurityReqID (320)
func (f SecurityReqID) Tag() fix.Tag { return tag.SecurityReqID }

//BuildSecurityReqID returns a new SecurityReqID initialized with val
func BuildSecurityReqID(val string) SecurityReqID {
	var field SecurityReqID
	field.Value = val
	return field
}

//SecurityRequestResult is a INT field
type SecurityRequestResult struct{ message.IntValue }

//Tag returns tag.SecurityRequestResult (560)
func (f SecurityRequestResult) Tag() fix.Tag { return tag.SecurityRequestResult }

//BuildSecurityRequestResult returns a new SecurityRequestResult initialized with val
func BuildSecurityRequestResult(val int) SecurityRequestResult {
	var field SecurityRequestResult
	field.Value = val
	return field
}

//SecurityRequestType is a INT field
type SecurityRequestType struct{ message.IntValue }

//Tag returns tag.SecurityRequestType (321)
func (f SecurityRequestType) Tag() fix.Tag { return tag.SecurityRequestType }

//BuildSecurityRequestType returns a new SecurityRequestType initialized with val
func BuildSecurityRequestType(val int) SecurityRequestType {
	var field SecurityRequestType
	field.Value = val
	return field
}

//SecurityResponseID is a STRING field
type SecurityResponseID struct{ message.StringValue }

//Tag returns tag.SecurityResponseID (322)
func (f SecurityResponseID) Tag() fix.Tag { return tag.SecurityResponseID }

//BuildSecurityResponseID returns a new SecurityResponseID initialized with val
func BuildSecurityResponseID(val string) SecurityResponseID {
	var field SecurityResponseID
	field.Value = val
	return field
}

//SecurityResponseType is a INT field
type SecurityResponseType struct{ message.IntValue }

//Tag returns tag.SecurityResponseType (323)
func (f SecurityResponseType) Tag() fix.Tag { return tag.SecurityResponseType }

//BuildSecurityResponseType returns a new SecurityResponseType initialized with val
func BuildSecurityResponseType(val int) SecurityResponseType {
	var field SecurityResponseType
	field.Value = val
	return field
}

//SecuritySettlAgentAcctName is a STRING field
type SecuritySettlAgentAcctName struct{ message.StringValue }

//Tag returns tag.SecuritySettlAgentAcctName (179)
func (f SecuritySettlAgentAcctName) Tag() fix.Tag { return tag.SecuritySettlAgentAcctName }

//BuildSecuritySettlAgentAcctName returns a new SecuritySettlAgentAcctName initialized with val
func BuildSecuritySettlAgentAcctName(val string) SecuritySettlAgentAcctName {
	var field SecuritySettlAgentAcctName
	field.Value = val
	return field
}

//SecuritySettlAgentAcctNum is a STRING field
type SecuritySettlAgentAcctNum struct{ message.StringValue }

//Tag returns tag.SecuritySettlAgentAcctNum (178)
func (f SecuritySettlAgentAcctNum) Tag() fix.Tag { return tag.SecuritySettlAgentAcctNum }

//BuildSecuritySettlAgentAcctNum returns a new SecuritySettlAgentAcctNum initialized with val
func BuildSecuritySettlAgentAcctNum(val string) SecuritySettlAgentAcctNum {
	var field SecuritySettlAgentAcctNum
	field.Value = val
	return field
}

//SecuritySettlAgentCode is a STRING field
type SecuritySettlAgentCode struct{ message.StringValue }

//Tag returns tag.SecuritySettlAgentCode (177)
func (f SecuritySettlAgentCode) Tag() fix.Tag { return tag.SecuritySettlAgentCode }

//BuildSecuritySettlAgentCode returns a new SecuritySettlAgentCode initialized with val
func BuildSecuritySettlAgentCode(val string) SecuritySettlAgentCode {
	var field SecuritySettlAgentCode
	field.Value = val
	return field
}

//SecuritySettlAgentContactName is a STRING field
type SecuritySettlAgentContactName struct{ message.StringValue }

//Tag returns tag.SecuritySettlAgentContactName (180)
func (f SecuritySettlAgentContactName) Tag() fix.Tag { return tag.SecuritySettlAgentContactName }

//BuildSecuritySettlAgentContactName returns a new SecuritySettlAgentContactName initialized with val
func BuildSecuritySettlAgentContactName(val string) SecuritySettlAgentContactName {
	var field SecuritySettlAgentContactName
	field.Value = val
	return field
}

//SecuritySettlAgentContactPhone is a STRING field
type SecuritySettlAgentContactPhone struct{ message.StringValue }

//Tag returns tag.SecuritySettlAgentContactPhone (181)
func (f SecuritySettlAgentContactPhone) Tag() fix.Tag { return tag.SecuritySettlAgentContactPhone }

//BuildSecuritySettlAgentContactPhone returns a new SecuritySettlAgentContactPhone initialized with val
func BuildSecuritySettlAgentContactPhone(val string) SecuritySettlAgentContactPhone {
	var field SecuritySettlAgentContactPhone
	field.Value = val
	return field
}

//SecuritySettlAgentName is a STRING field
type SecuritySettlAgentName struct{ message.StringValue }

//Tag returns tag.SecuritySettlAgentName (176)
func (f SecuritySettlAgentName) Tag() fix.Tag { return tag.SecuritySettlAgentName }

//BuildSecuritySettlAgentName returns a new SecuritySettlAgentName initialized with val
func BuildSecuritySettlAgentName(val string) SecuritySettlAgentName {
	var field SecuritySettlAgentName
	field.Value = val
	return field
}

//SecurityStatus is a STRING field
type SecurityStatus struct{ message.StringValue }

//Tag returns tag.SecurityStatus (965)
func (f SecurityStatus) Tag() fix.Tag { return tag.SecurityStatus }

//BuildSecurityStatus returns a new SecurityStatus initialized with val
func BuildSecurityStatus(val string) SecurityStatus {
	var field SecurityStatus
	field.Value = val
	return field
}

//SecurityStatusReqID is a STRING field
type SecurityStatusReqID struct{ message.StringValue }

//Tag returns tag.SecurityStatusReqID (324)
func (f SecurityStatusReqID) Tag() fix.Tag { return tag.SecurityStatusReqID }

//BuildSecurityStatusReqID returns a new SecurityStatusReqID initialized with val
func BuildSecurityStatusReqID(val string) SecurityStatusReqID {
	var field SecurityStatusReqID
	field.Value = val
	return field
}

//SecuritySubType is a STRING field
type SecuritySubType struct{ message.StringValue }

//Tag returns tag.SecuritySubType (762)
func (f SecuritySubType) Tag() fix.Tag { return tag.SecuritySubType }

//BuildSecuritySubType returns a new SecuritySubType initialized with val
func BuildSecuritySubType(val string) SecuritySubType {
	var field SecuritySubType
	field.Value = val
	return field
}

//SecurityTradingEvent is a INT field
type SecurityTradingEvent struct{ message.IntValue }

//Tag returns tag.SecurityTradingEvent (1174)
func (f SecurityTradingEvent) Tag() fix.Tag { return tag.SecurityTradingEvent }

//BuildSecurityTradingEvent returns a new SecurityTradingEvent initialized with val
func BuildSecurityTradingEvent(val int) SecurityTradingEvent {
	var field SecurityTradingEvent
	field.Value = val
	return field
}

//SecurityTradingStatus is a INT field
type SecurityTradingStatus struct{ message.IntValue }

//Tag returns tag.SecurityTradingStatus (326)
func (f SecurityTradingStatus) Tag() fix.Tag { return tag.SecurityTradingStatus }

//BuildSecurityTradingStatus returns a new SecurityTradingStatus initialized with val
func BuildSecurityTradingStatus(val int) SecurityTradingStatus {
	var field SecurityTradingStatus
	field.Value = val
	return field
}

//SecurityType is a STRING field
type SecurityType struct{ message.StringValue }

//Tag returns tag.SecurityType (167)
func (f SecurityType) Tag() fix.Tag { return tag.SecurityType }

//BuildSecurityType returns a new SecurityType initialized with val
func BuildSecurityType(val string) SecurityType {
	var field SecurityType
	field.Value = val
	return field
}

//SecurityUpdateAction is a CHAR field
type SecurityUpdateAction struct{ message.CharValue }

//Tag returns tag.SecurityUpdateAction (980)
func (f SecurityUpdateAction) Tag() fix.Tag { return tag.SecurityUpdateAction }

//BuildSecurityUpdateAction returns a new SecurityUpdateAction initialized with val
func BuildSecurityUpdateAction(val string) SecurityUpdateAction {
	var field SecurityUpdateAction
	field.Value = val
	return field
}

//SecurityXML is a XMLDATA field
type SecurityXML struct{ message.XMLDataValue }

//Tag returns tag.SecurityXML (1185)
func (f SecurityXML) Tag() fix.Tag { return tag.SecurityXML }

//BuildSecurityXML returns a new SecurityXML initialized with val
func BuildSecurityXML(val string) SecurityXML {
	var field SecurityXML
	field.Value = val
	return field
}

//SecurityXMLLen is a LENGTH field
type SecurityXMLLen struct{ message.LengthValue }

//Tag returns tag.SecurityXMLLen (1184)
func (f SecurityXMLLen) Tag() fix.Tag { return tag.SecurityXMLLen }

//BuildSecurityXMLLen returns a new SecurityXMLLen initialized with val
func BuildSecurityXMLLen(val int) SecurityXMLLen {
	var field SecurityXMLLen
	field.Value = val
	return field
}

//SecurityXMLSchema is a STRING field
type SecurityXMLSchema struct{ message.StringValue }

//Tag returns tag.SecurityXMLSchema (1186)
func (f SecurityXMLSchema) Tag() fix.Tag { return tag.SecurityXMLSchema }

//BuildSecurityXMLSchema returns a new SecurityXMLSchema initialized with val
func BuildSecurityXMLSchema(val string) SecurityXMLSchema {
	var field SecurityXMLSchema
	field.Value = val
	return field
}

//SellVolume is a QTY field
type SellVolume struct{ message.QtyValue }

//Tag returns tag.SellVolume (331)
func (f SellVolume) Tag() fix.Tag { return tag.SellVolume }

//BuildSellVolume returns a new SellVolume initialized with val
func BuildSellVolume(val float64) SellVolume {
	var field SellVolume
	field.Value = val
	return field
}

//SellerDays is a INT field
type SellerDays struct{ message.IntValue }

//Tag returns tag.SellerDays (287)
func (f SellerDays) Tag() fix.Tag { return tag.SellerDays }

//BuildSellerDays returns a new SellerDays initialized with val
func BuildSellerDays(val int) SellerDays {
	var field SellerDays
	field.Value = val
	return field
}

//SenderCompID is a STRING field
type SenderCompID struct{ message.StringValue }

//Tag returns tag.SenderCompID (49)
func (f SenderCompID) Tag() fix.Tag { return tag.SenderCompID }

//BuildSenderCompID returns a new SenderCompID initialized with val
func BuildSenderCompID(val string) SenderCompID {
	var field SenderCompID
	field.Value = val
	return field
}

//SenderLocationID is a STRING field
type SenderLocationID struct{ message.StringValue }

//Tag returns tag.SenderLocationID (142)
func (f SenderLocationID) Tag() fix.Tag { return tag.SenderLocationID }

//BuildSenderLocationID returns a new SenderLocationID initialized with val
func BuildSenderLocationID(val string) SenderLocationID {
	var field SenderLocationID
	field.Value = val
	return field
}

//SenderSubID is a STRING field
type SenderSubID struct{ message.StringValue }

//Tag returns tag.SenderSubID (50)
func (f SenderSubID) Tag() fix.Tag { return tag.SenderSubID }

//BuildSenderSubID returns a new SenderSubID initialized with val
func BuildSenderSubID(val string) SenderSubID {
	var field SenderSubID
	field.Value = val
	return field
}

//SendingDate is a LOCALMKTDATE field
type SendingDate struct{ message.LocalMktDateValue }

//Tag returns tag.SendingDate (51)
func (f SendingDate) Tag() fix.Tag { return tag.SendingDate }

//BuildSendingDate returns a new SendingDate initialized with val
func BuildSendingDate(val string) SendingDate {
	var field SendingDate
	field.Value = val
	return field
}

//SendingTime is a UTCTIMESTAMP field
type SendingTime struct{ message.UTCTimestampValue }

//Tag returns tag.SendingTime (52)
func (f SendingTime) Tag() fix.Tag { return tag.SendingTime }

//Seniority is a STRING field
type Seniority struct{ message.StringValue }

//Tag returns tag.Seniority (1450)
func (f Seniority) Tag() fix.Tag { return tag.Seniority }

//BuildSeniority returns a new Seniority initialized with val
func BuildSeniority(val string) Seniority {
	var field Seniority
	field.Value = val
	return field
}

//SessionRejectReason is a INT field
type SessionRejectReason struct{ message.IntValue }

//Tag returns tag.SessionRejectReason (373)
func (f SessionRejectReason) Tag() fix.Tag { return tag.SessionRejectReason }

//BuildSessionRejectReason returns a new SessionRejectReason initialized with val
func BuildSessionRejectReason(val int) SessionRejectReason {
	var field SessionRejectReason
	field.Value = val
	return field
}

//SessionStatus is a INT field
type SessionStatus struct{ message.IntValue }

//Tag returns tag.SessionStatus (1409)
func (f SessionStatus) Tag() fix.Tag { return tag.SessionStatus }

//BuildSessionStatus returns a new SessionStatus initialized with val
func BuildSessionStatus(val int) SessionStatus {
	var field SessionStatus
	field.Value = val
	return field
}

//SettlBrkrCode is a STRING field
type SettlBrkrCode struct{ message.StringValue }

//Tag returns tag.SettlBrkrCode (174)
func (f SettlBrkrCode) Tag() fix.Tag { return tag.SettlBrkrCode }

//BuildSettlBrkrCode returns a new SettlBrkrCode initialized with val
func BuildSettlBrkrCode(val string) SettlBrkrCode {
	var field SettlBrkrCode
	field.Value = val
	return field
}

//SettlCurrAmt is a AMT field
type SettlCurrAmt struct{ message.AmtValue }

//Tag returns tag.SettlCurrAmt (119)
func (f SettlCurrAmt) Tag() fix.Tag { return tag.SettlCurrAmt }

//BuildSettlCurrAmt returns a new SettlCurrAmt initialized with val
func BuildSettlCurrAmt(val float64) SettlCurrAmt {
	var field SettlCurrAmt
	field.Value = val
	return field
}

//SettlCurrBidFxRate is a FLOAT field
type SettlCurrBidFxRate struct{ message.FloatValue }

//Tag returns tag.SettlCurrBidFxRate (656)
func (f SettlCurrBidFxRate) Tag() fix.Tag { return tag.SettlCurrBidFxRate }

//BuildSettlCurrBidFxRate returns a new SettlCurrBidFxRate initialized with val
func BuildSettlCurrBidFxRate(val float64) SettlCurrBidFxRate {
	var field SettlCurrBidFxRate
	field.Value = val
	return field
}

//SettlCurrFxRate is a FLOAT field
type SettlCurrFxRate struct{ message.FloatValue }

//Tag returns tag.SettlCurrFxRate (155)
func (f SettlCurrFxRate) Tag() fix.Tag { return tag.SettlCurrFxRate }

//BuildSettlCurrFxRate returns a new SettlCurrFxRate initialized with val
func BuildSettlCurrFxRate(val float64) SettlCurrFxRate {
	var field SettlCurrFxRate
	field.Value = val
	return field
}

//SettlCurrFxRateCalc is a CHAR field
type SettlCurrFxRateCalc struct{ message.CharValue }

//Tag returns tag.SettlCurrFxRateCalc (156)
func (f SettlCurrFxRateCalc) Tag() fix.Tag { return tag.SettlCurrFxRateCalc }

//BuildSettlCurrFxRateCalc returns a new SettlCurrFxRateCalc initialized with val
func BuildSettlCurrFxRateCalc(val string) SettlCurrFxRateCalc {
	var field SettlCurrFxRateCalc
	field.Value = val
	return field
}

//SettlCurrOfferFxRate is a FLOAT field
type SettlCurrOfferFxRate struct{ message.FloatValue }

//Tag returns tag.SettlCurrOfferFxRate (657)
func (f SettlCurrOfferFxRate) Tag() fix.Tag { return tag.SettlCurrOfferFxRate }

//BuildSettlCurrOfferFxRate returns a new SettlCurrOfferFxRate initialized with val
func BuildSettlCurrOfferFxRate(val float64) SettlCurrOfferFxRate {
	var field SettlCurrOfferFxRate
	field.Value = val
	return field
}

//SettlCurrency is a CURRENCY field
type SettlCurrency struct{ message.CurrencyValue }

//Tag returns tag.SettlCurrency (120)
func (f SettlCurrency) Tag() fix.Tag { return tag.SettlCurrency }

//BuildSettlCurrency returns a new SettlCurrency initialized with val
func BuildSettlCurrency(val string) SettlCurrency {
	var field SettlCurrency
	field.Value = val
	return field
}

//SettlDate is a LOCALMKTDATE field
type SettlDate struct{ message.LocalMktDateValue }

//Tag returns tag.SettlDate (64)
func (f SettlDate) Tag() fix.Tag { return tag.SettlDate }

//BuildSettlDate returns a new SettlDate initialized with val
func BuildSettlDate(val string) SettlDate {
	var field SettlDate
	field.Value = val
	return field
}

//SettlDate2 is a LOCALMKTDATE field
type SettlDate2 struct{ message.LocalMktDateValue }

//Tag returns tag.SettlDate2 (193)
func (f SettlDate2) Tag() fix.Tag { return tag.SettlDate2 }

//BuildSettlDate2 returns a new SettlDate2 initialized with val
func BuildSettlDate2(val string) SettlDate2 {
	var field SettlDate2
	field.Value = val
	return field
}

//SettlDeliveryType is a INT field
type SettlDeliveryType struct{ message.IntValue }

//Tag returns tag.SettlDeliveryType (172)
func (f SettlDeliveryType) Tag() fix.Tag { return tag.SettlDeliveryType }

//BuildSettlDeliveryType returns a new SettlDeliveryType initialized with val
func BuildSettlDeliveryType(val int) SettlDeliveryType {
	var field SettlDeliveryType
	field.Value = val
	return field
}

//SettlDepositoryCode is a STRING field
type SettlDepositoryCode struct{ message.StringValue }

//Tag returns tag.SettlDepositoryCode (173)
func (f SettlDepositoryCode) Tag() fix.Tag { return tag.SettlDepositoryCode }

//BuildSettlDepositoryCode returns a new SettlDepositoryCode initialized with val
func BuildSettlDepositoryCode(val string) SettlDepositoryCode {
	var field SettlDepositoryCode
	field.Value = val
	return field
}

//SettlInstCode is a STRING field
type SettlInstCode struct{ message.StringValue }

//Tag returns tag.SettlInstCode (175)
func (f SettlInstCode) Tag() fix.Tag { return tag.SettlInstCode }

//BuildSettlInstCode returns a new SettlInstCode initialized with val
func BuildSettlInstCode(val string) SettlInstCode {
	var field SettlInstCode
	field.Value = val
	return field
}

//SettlInstID is a STRING field
type SettlInstID struct{ message.StringValue }

//Tag returns tag.SettlInstID (162)
func (f SettlInstID) Tag() fix.Tag { return tag.SettlInstID }

//BuildSettlInstID returns a new SettlInstID initialized with val
func BuildSettlInstID(val string) SettlInstID {
	var field SettlInstID
	field.Value = val
	return field
}

//SettlInstMode is a CHAR field
type SettlInstMode struct{ message.CharValue }

//Tag returns tag.SettlInstMode (160)
func (f SettlInstMode) Tag() fix.Tag { return tag.SettlInstMode }

//BuildSettlInstMode returns a new SettlInstMode initialized with val
func BuildSettlInstMode(val string) SettlInstMode {
	var field SettlInstMode
	field.Value = val
	return field
}

//SettlInstMsgID is a STRING field
type SettlInstMsgID struct{ message.StringValue }

//Tag returns tag.SettlInstMsgID (777)
func (f SettlInstMsgID) Tag() fix.Tag { return tag.SettlInstMsgID }

//BuildSettlInstMsgID returns a new SettlInstMsgID initialized with val
func BuildSettlInstMsgID(val string) SettlInstMsgID {
	var field SettlInstMsgID
	field.Value = val
	return field
}

//SettlInstRefID is a STRING field
type SettlInstRefID struct{ message.StringValue }

//Tag returns tag.SettlInstRefID (214)
func (f SettlInstRefID) Tag() fix.Tag { return tag.SettlInstRefID }

//BuildSettlInstRefID returns a new SettlInstRefID initialized with val
func BuildSettlInstRefID(val string) SettlInstRefID {
	var field SettlInstRefID
	field.Value = val
	return field
}

//SettlInstReqID is a STRING field
type SettlInstReqID struct{ message.StringValue }

//Tag returns tag.SettlInstReqID (791)
func (f SettlInstReqID) Tag() fix.Tag { return tag.SettlInstReqID }

//BuildSettlInstReqID returns a new SettlInstReqID initialized with val
func BuildSettlInstReqID(val string) SettlInstReqID {
	var field SettlInstReqID
	field.Value = val
	return field
}

//SettlInstReqRejCode is a INT field
type SettlInstReqRejCode struct{ message.IntValue }

//Tag returns tag.SettlInstReqRejCode (792)
func (f SettlInstReqRejCode) Tag() fix.Tag { return tag.SettlInstReqRejCode }

//BuildSettlInstReqRejCode returns a new SettlInstReqRejCode initialized with val
func BuildSettlInstReqRejCode(val int) SettlInstReqRejCode {
	var field SettlInstReqRejCode
	field.Value = val
	return field
}

//SettlInstSource is a CHAR field
type SettlInstSource struct{ message.CharValue }

//Tag returns tag.SettlInstSource (165)
func (f SettlInstSource) Tag() fix.Tag { return tag.SettlInstSource }

//BuildSettlInstSource returns a new SettlInstSource initialized with val
func BuildSettlInstSource(val string) SettlInstSource {
	var field SettlInstSource
	field.Value = val
	return field
}

//SettlInstTransType is a CHAR field
type SettlInstTransType struct{ message.CharValue }

//Tag returns tag.SettlInstTransType (163)
func (f SettlInstTransType) Tag() fix.Tag { return tag.SettlInstTransType }

//BuildSettlInstTransType returns a new SettlInstTransType initialized with val
func BuildSettlInstTransType(val string) SettlInstTransType {
	var field SettlInstTransType
	field.Value = val
	return field
}

//SettlLocation is a STRING field
type SettlLocation struct{ message.StringValue }

//Tag returns tag.SettlLocation (166)
func (f SettlLocation) Tag() fix.Tag { return tag.SettlLocation }

//BuildSettlLocation returns a new SettlLocation initialized with val
func BuildSettlLocation(val string) SettlLocation {
	var field SettlLocation
	field.Value = val
	return field
}

//SettlMethod is a CHAR field
type SettlMethod struct{ message.CharValue }

//Tag returns tag.SettlMethod (1193)
func (f SettlMethod) Tag() fix.Tag { return tag.SettlMethod }

//BuildSettlMethod returns a new SettlMethod initialized with val
func BuildSettlMethod(val string) SettlMethod {
	var field SettlMethod
	field.Value = val
	return field
}

//SettlObligID is a STRING field
type SettlObligID struct{ message.StringValue }

//Tag returns tag.SettlObligID (1161)
func (f SettlObligID) Tag() fix.Tag { return tag.SettlObligID }

//BuildSettlObligID returns a new SettlObligID initialized with val
func BuildSettlObligID(val string) SettlObligID {
	var field SettlObligID
	field.Value = val
	return field
}

//SettlObligMode is a INT field
type SettlObligMode struct{ message.IntValue }

//Tag returns tag.SettlObligMode (1159)
func (f SettlObligMode) Tag() fix.Tag { return tag.SettlObligMode }

//BuildSettlObligMode returns a new SettlObligMode initialized with val
func BuildSettlObligMode(val int) SettlObligMode {
	var field SettlObligMode
	field.Value = val
	return field
}

//SettlObligMsgID is a STRING field
type SettlObligMsgID struct{ message.StringValue }

//Tag returns tag.SettlObligMsgID (1160)
func (f SettlObligMsgID) Tag() fix.Tag { return tag.SettlObligMsgID }

//BuildSettlObligMsgID returns a new SettlObligMsgID initialized with val
func BuildSettlObligMsgID(val string) SettlObligMsgID {
	var field SettlObligMsgID
	field.Value = val
	return field
}

//SettlObligRefID is a STRING field
type SettlObligRefID struct{ message.StringValue }

//Tag returns tag.SettlObligRefID (1163)
func (f SettlObligRefID) Tag() fix.Tag { return tag.SettlObligRefID }

//BuildSettlObligRefID returns a new SettlObligRefID initialized with val
func BuildSettlObligRefID(val string) SettlObligRefID {
	var field SettlObligRefID
	field.Value = val
	return field
}

//SettlObligSource is a CHAR field
type SettlObligSource struct{ message.CharValue }

//Tag returns tag.SettlObligSource (1164)
func (f SettlObligSource) Tag() fix.Tag { return tag.SettlObligSource }

//BuildSettlObligSource returns a new SettlObligSource initialized with val
func BuildSettlObligSource(val string) SettlObligSource {
	var field SettlObligSource
	field.Value = val
	return field
}

//SettlObligTransType is a CHAR field
type SettlObligTransType struct{ message.CharValue }

//Tag returns tag.SettlObligTransType (1162)
func (f SettlObligTransType) Tag() fix.Tag { return tag.SettlObligTransType }

//BuildSettlObligTransType returns a new SettlObligTransType initialized with val
func BuildSettlObligTransType(val string) SettlObligTransType {
	var field SettlObligTransType
	field.Value = val
	return field
}

//SettlPartyID is a STRING field
type SettlPartyID struct{ message.StringValue }

//Tag returns tag.SettlPartyID (782)
func (f SettlPartyID) Tag() fix.Tag { return tag.SettlPartyID }

//BuildSettlPartyID returns a new SettlPartyID initialized with val
func BuildSettlPartyID(val string) SettlPartyID {
	var field SettlPartyID
	field.Value = val
	return field
}

//SettlPartyIDSource is a CHAR field
type SettlPartyIDSource struct{ message.CharValue }

//Tag returns tag.SettlPartyIDSource (783)
func (f SettlPartyIDSource) Tag() fix.Tag { return tag.SettlPartyIDSource }

//BuildSettlPartyIDSource returns a new SettlPartyIDSource initialized with val
func BuildSettlPartyIDSource(val string) SettlPartyIDSource {
	var field SettlPartyIDSource
	field.Value = val
	return field
}

//SettlPartyRole is a INT field
type SettlPartyRole struct{ message.IntValue }

//Tag returns tag.SettlPartyRole (784)
func (f SettlPartyRole) Tag() fix.Tag { return tag.SettlPartyRole }

//BuildSettlPartyRole returns a new SettlPartyRole initialized with val
func BuildSettlPartyRole(val int) SettlPartyRole {
	var field SettlPartyRole
	field.Value = val
	return field
}

//SettlPartySubID is a STRING field
type SettlPartySubID struct{ message.StringValue }

//Tag returns tag.SettlPartySubID (785)
func (f SettlPartySubID) Tag() fix.Tag { return tag.SettlPartySubID }

//BuildSettlPartySubID returns a new SettlPartySubID initialized with val
func BuildSettlPartySubID(val string) SettlPartySubID {
	var field SettlPartySubID
	field.Value = val
	return field
}

//SettlPartySubIDType is a INT field
type SettlPartySubIDType struct{ message.IntValue }

//Tag returns tag.SettlPartySubIDType (786)
func (f SettlPartySubIDType) Tag() fix.Tag { return tag.SettlPartySubIDType }

//BuildSettlPartySubIDType returns a new SettlPartySubIDType initialized with val
func BuildSettlPartySubIDType(val int) SettlPartySubIDType {
	var field SettlPartySubIDType
	field.Value = val
	return field
}

//SettlPrice is a PRICE field
type SettlPrice struct{ message.PriceValue }

//Tag returns tag.SettlPrice (730)
func (f SettlPrice) Tag() fix.Tag { return tag.SettlPrice }

//BuildSettlPrice returns a new SettlPrice initialized with val
func BuildSettlPrice(val float64) SettlPrice {
	var field SettlPrice
	field.Value = val
	return field
}

//SettlPriceType is a INT field
type SettlPriceType struct{ message.IntValue }

//Tag returns tag.SettlPriceType (731)
func (f SettlPriceType) Tag() fix.Tag { return tag.SettlPriceType }

//BuildSettlPriceType returns a new SettlPriceType initialized with val
func BuildSettlPriceType(val int) SettlPriceType {
	var field SettlPriceType
	field.Value = val
	return field
}

//SettlSessID is a STRING field
type SettlSessID struct{ message.StringValue }

//Tag returns tag.SettlSessID (716)
func (f SettlSessID) Tag() fix.Tag { return tag.SettlSessID }

//BuildSettlSessID returns a new SettlSessID initialized with val
func BuildSettlSessID(val string) SettlSessID {
	var field SettlSessID
	field.Value = val
	return field
}

//SettlSessSubID is a STRING field
type SettlSessSubID struct{ message.StringValue }

//Tag returns tag.SettlSessSubID (717)
func (f SettlSessSubID) Tag() fix.Tag { return tag.SettlSessSubID }

//BuildSettlSessSubID returns a new SettlSessSubID initialized with val
func BuildSettlSessSubID(val string) SettlSessSubID {
	var field SettlSessSubID
	field.Value = val
	return field
}

//SettlType is a STRING field
type SettlType struct{ message.StringValue }

//Tag returns tag.SettlType (63)
func (f SettlType) Tag() fix.Tag { return tag.SettlType }

//BuildSettlType returns a new SettlType initialized with val
func BuildSettlType(val string) SettlType {
	var field SettlType
	field.Value = val
	return field
}

//SettleOnOpenFlag is a STRING field
type SettleOnOpenFlag struct{ message.StringValue }

//Tag returns tag.SettleOnOpenFlag (966)
func (f SettleOnOpenFlag) Tag() fix.Tag { return tag.SettleOnOpenFlag }

//BuildSettleOnOpenFlag returns a new SettleOnOpenFlag initialized with val
func BuildSettleOnOpenFlag(val string) SettleOnOpenFlag {
	var field SettleOnOpenFlag
	field.Value = val
	return field
}

//SettlementCycleNo is a INT field
type SettlementCycleNo struct{ message.IntValue }

//Tag returns tag.SettlementCycleNo (1153)
func (f SettlementCycleNo) Tag() fix.Tag { return tag.SettlementCycleNo }

//BuildSettlementCycleNo returns a new SettlementCycleNo initialized with val
func BuildSettlementCycleNo(val int) SettlementCycleNo {
	var field SettlementCycleNo
	field.Value = val
	return field
}

//SettlmntTyp is a CHAR field
type SettlmntTyp struct{ message.CharValue }

//Tag returns tag.SettlmntTyp (63)
func (f SettlmntTyp) Tag() fix.Tag { return tag.SettlmntTyp }

//BuildSettlmntTyp returns a new SettlmntTyp initialized with val
func BuildSettlmntTyp(val string) SettlmntTyp {
	var field SettlmntTyp
	field.Value = val
	return field
}

//SharedCommission is a AMT field
type SharedCommission struct{ message.AmtValue }

//Tag returns tag.SharedCommission (858)
func (f SharedCommission) Tag() fix.Tag { return tag.SharedCommission }

//BuildSharedCommission returns a new SharedCommission initialized with val
func BuildSharedCommission(val float64) SharedCommission {
	var field SharedCommission
	field.Value = val
	return field
}

//Shares is a QTY field
type Shares struct{ message.QtyValue }

//Tag returns tag.Shares (53)
func (f Shares) Tag() fix.Tag { return tag.Shares }

//BuildShares returns a new Shares initialized with val
func BuildShares(val float64) Shares {
	var field Shares
	field.Value = val
	return field
}

//ShortQty is a QTY field
type ShortQty struct{ message.QtyValue }

//Tag returns tag.ShortQty (705)
func (f ShortQty) Tag() fix.Tag { return tag.ShortQty }

//BuildShortQty returns a new ShortQty initialized with val
func BuildShortQty(val float64) ShortQty {
	var field ShortQty
	field.Value = val
	return field
}

//ShortSaleReason is a INT field
type ShortSaleReason struct{ message.IntValue }

//Tag returns tag.ShortSaleReason (853)
func (f ShortSaleReason) Tag() fix.Tag { return tag.ShortSaleReason }

//BuildShortSaleReason returns a new ShortSaleReason initialized with val
func BuildShortSaleReason(val int) ShortSaleReason {
	var field ShortSaleReason
	field.Value = val
	return field
}

//Side is a CHAR field
type Side struct{ message.CharValue }

//Tag returns tag.Side (54)
func (f Side) Tag() fix.Tag { return tag.Side }

//BuildSide returns a new Side initialized with val
func BuildSide(val string) Side {
	var field Side
	field.Value = val
	return field
}

//SideComplianceID is a STRING field
type SideComplianceID struct{ message.StringValue }

//Tag returns tag.SideComplianceID (659)
func (f SideComplianceID) Tag() fix.Tag { return tag.SideComplianceID }

//BuildSideComplianceID returns a new SideComplianceID initialized with val
func BuildSideComplianceID(val string) SideComplianceID {
	var field SideComplianceID
	field.Value = val
	return field
}

//SideCurrency is a CURRENCY field
type SideCurrency struct{ message.CurrencyValue }

//Tag returns tag.SideCurrency (1154)
func (f SideCurrency) Tag() fix.Tag { return tag.SideCurrency }

//BuildSideCurrency returns a new SideCurrency initialized with val
func BuildSideCurrency(val string) SideCurrency {
	var field SideCurrency
	field.Value = val
	return field
}

//SideExecID is a STRING field
type SideExecID struct{ message.StringValue }

//Tag returns tag.SideExecID (1427)
func (f SideExecID) Tag() fix.Tag { return tag.SideExecID }

//BuildSideExecID returns a new SideExecID initialized with val
func BuildSideExecID(val string) SideExecID {
	var field SideExecID
	field.Value = val
	return field
}

//SideFillStationCd is a STRING field
type SideFillStationCd struct{ message.StringValue }

//Tag returns tag.SideFillStationCd (1006)
func (f SideFillStationCd) Tag() fix.Tag { return tag.SideFillStationCd }

//BuildSideFillStationCd returns a new SideFillStationCd initialized with val
func BuildSideFillStationCd(val string) SideFillStationCd {
	var field SideFillStationCd
	field.Value = val
	return field
}

//SideGrossTradeAmt is a AMT field
type SideGrossTradeAmt struct{ message.AmtValue }

//Tag returns tag.SideGrossTradeAmt (1072)
func (f SideGrossTradeAmt) Tag() fix.Tag { return tag.SideGrossTradeAmt }

//BuildSideGrossTradeAmt returns a new SideGrossTradeAmt initialized with val
func BuildSideGrossTradeAmt(val float64) SideGrossTradeAmt {
	var field SideGrossTradeAmt
	field.Value = val
	return field
}

//SideLastQty is a INT field
type SideLastQty struct{ message.IntValue }

//Tag returns tag.SideLastQty (1009)
func (f SideLastQty) Tag() fix.Tag { return tag.SideLastQty }

//BuildSideLastQty returns a new SideLastQty initialized with val
func BuildSideLastQty(val int) SideLastQty {
	var field SideLastQty
	field.Value = val
	return field
}

//SideLiquidityInd is a INT field
type SideLiquidityInd struct{ message.IntValue }

//Tag returns tag.SideLiquidityInd (1444)
func (f SideLiquidityInd) Tag() fix.Tag { return tag.SideLiquidityInd }

//BuildSideLiquidityInd returns a new SideLiquidityInd initialized with val
func BuildSideLiquidityInd(val int) SideLiquidityInd {
	var field SideLiquidityInd
	field.Value = val
	return field
}

//SideMultiLegReportingType is a INT field
type SideMultiLegReportingType struct{ message.IntValue }

//Tag returns tag.SideMultiLegReportingType (752)
func (f SideMultiLegReportingType) Tag() fix.Tag { return tag.SideMultiLegReportingType }

//BuildSideMultiLegReportingType returns a new SideMultiLegReportingType initialized with val
func BuildSideMultiLegReportingType(val int) SideMultiLegReportingType {
	var field SideMultiLegReportingType
	field.Value = val
	return field
}

//SideQty is a INT field
type SideQty struct{ message.IntValue }

//Tag returns tag.SideQty (1009)
func (f SideQty) Tag() fix.Tag { return tag.SideQty }

//BuildSideQty returns a new SideQty initialized with val
func BuildSideQty(val int) SideQty {
	var field SideQty
	field.Value = val
	return field
}

//SideReasonCd is a STRING field
type SideReasonCd struct{ message.StringValue }

//Tag returns tag.SideReasonCd (1007)
func (f SideReasonCd) Tag() fix.Tag { return tag.SideReasonCd }

//BuildSideReasonCd returns a new SideReasonCd initialized with val
func BuildSideReasonCd(val string) SideReasonCd {
	var field SideReasonCd
	field.Value = val
	return field
}

//SideSettlCurrency is a CURRENCY field
type SideSettlCurrency struct{ message.CurrencyValue }

//Tag returns tag.SideSettlCurrency (1155)
func (f SideSettlCurrency) Tag() fix.Tag { return tag.SideSettlCurrency }

//BuildSideSettlCurrency returns a new SideSettlCurrency initialized with val
func BuildSideSettlCurrency(val string) SideSettlCurrency {
	var field SideSettlCurrency
	field.Value = val
	return field
}

//SideTimeInForce is a UTCTIMESTAMP field
type SideTimeInForce struct{ message.UTCTimestampValue }

//Tag returns tag.SideTimeInForce (962)
func (f SideTimeInForce) Tag() fix.Tag { return tag.SideTimeInForce }

//SideTradeReportID is a STRING field
type SideTradeReportID struct{ message.StringValue }

//Tag returns tag.SideTradeReportID (1005)
func (f SideTradeReportID) Tag() fix.Tag { return tag.SideTradeReportID }

//BuildSideTradeReportID returns a new SideTradeReportID initialized with val
func BuildSideTradeReportID(val string) SideTradeReportID {
	var field SideTradeReportID
	field.Value = val
	return field
}

//SideTrdRegTimestamp is a UTCTIMESTAMP field
type SideTrdRegTimestamp struct{ message.UTCTimestampValue }

//Tag returns tag.SideTrdRegTimestamp (1012)
func (f SideTrdRegTimestamp) Tag() fix.Tag { return tag.SideTrdRegTimestamp }

//SideTrdRegTimestampSrc is a STRING field
type SideTrdRegTimestampSrc struct{ message.StringValue }

//Tag returns tag.SideTrdRegTimestampSrc (1014)
func (f SideTrdRegTimestampSrc) Tag() fix.Tag { return tag.SideTrdRegTimestampSrc }

//BuildSideTrdRegTimestampSrc returns a new SideTrdRegTimestampSrc initialized with val
func BuildSideTrdRegTimestampSrc(val string) SideTrdRegTimestampSrc {
	var field SideTrdRegTimestampSrc
	field.Value = val
	return field
}

//SideTrdRegTimestampType is a INT field
type SideTrdRegTimestampType struct{ message.IntValue }

//Tag returns tag.SideTrdRegTimestampType (1013)
func (f SideTrdRegTimestampType) Tag() fix.Tag { return tag.SideTrdRegTimestampType }

//BuildSideTrdRegTimestampType returns a new SideTrdRegTimestampType initialized with val
func BuildSideTrdRegTimestampType(val int) SideTrdRegTimestampType {
	var field SideTrdRegTimestampType
	field.Value = val
	return field
}

//SideTrdSubTyp is a INT field
type SideTrdSubTyp struct{ message.IntValue }

//Tag returns tag.SideTrdSubTyp (1008)
func (f SideTrdSubTyp) Tag() fix.Tag { return tag.SideTrdSubTyp }

//BuildSideTrdSubTyp returns a new SideTrdSubTyp initialized with val
func BuildSideTrdSubTyp(val int) SideTrdSubTyp {
	var field SideTrdSubTyp
	field.Value = val
	return field
}

//SideValue1 is a AMT field
type SideValue1 struct{ message.AmtValue }

//Tag returns tag.SideValue1 (396)
func (f SideValue1) Tag() fix.Tag { return tag.SideValue1 }

//BuildSideValue1 returns a new SideValue1 initialized with val
func BuildSideValue1(val float64) SideValue1 {
	var field SideValue1
	field.Value = val
	return field
}

//SideValue2 is a AMT field
type SideValue2 struct{ message.AmtValue }

//Tag returns tag.SideValue2 (397)
func (f SideValue2) Tag() fix.Tag { return tag.SideValue2 }

//BuildSideValue2 returns a new SideValue2 initialized with val
func BuildSideValue2(val float64) SideValue2 {
	var field SideValue2
	field.Value = val
	return field
}

//SideValueInd is a INT field
type SideValueInd struct{ message.IntValue }

//Tag returns tag.SideValueInd (401)
func (f SideValueInd) Tag() fix.Tag { return tag.SideValueInd }

//BuildSideValueInd returns a new SideValueInd initialized with val
func BuildSideValueInd(val int) SideValueInd {
	var field SideValueInd
	field.Value = val
	return field
}

//Signature is a DATA field
type Signature struct{ message.DataValue }

//Tag returns tag.Signature (89)
func (f Signature) Tag() fix.Tag { return tag.Signature }

//BuildSignature returns a new Signature initialized with val
func BuildSignature(val string) Signature {
	var field Signature
	field.Value = val
	return field
}

//SignatureLength is a LENGTH field
type SignatureLength struct{ message.LengthValue }

//Tag returns tag.SignatureLength (93)
func (f SignatureLength) Tag() fix.Tag { return tag.SignatureLength }

//BuildSignatureLength returns a new SignatureLength initialized with val
func BuildSignatureLength(val int) SignatureLength {
	var field SignatureLength
	field.Value = val
	return field
}

//SolicitedFlag is a BOOLEAN field
type SolicitedFlag struct{ message.BooleanValue }

//Tag returns tag.SolicitedFlag (377)
func (f SolicitedFlag) Tag() fix.Tag { return tag.SolicitedFlag }

//BuildSolicitedFlag returns a new SolicitedFlag initialized with val
func BuildSolicitedFlag(val bool) SolicitedFlag {
	var field SolicitedFlag
	field.Value = val
	return field
}

//Spread is a PRICEOFFSET field
type Spread struct{ message.PriceOffsetValue }

//Tag returns tag.Spread (218)
func (f Spread) Tag() fix.Tag { return tag.Spread }

//BuildSpread returns a new Spread initialized with val
func BuildSpread(val float64) Spread {
	var field Spread
	field.Value = val
	return field
}

//SpreadToBenchmark is a PRICEOFFSET field
type SpreadToBenchmark struct{ message.PriceOffsetValue }

//Tag returns tag.SpreadToBenchmark (218)
func (f SpreadToBenchmark) Tag() fix.Tag { return tag.SpreadToBenchmark }

//BuildSpreadToBenchmark returns a new SpreadToBenchmark initialized with val
func BuildSpreadToBenchmark(val float64) SpreadToBenchmark {
	var field SpreadToBenchmark
	field.Value = val
	return field
}

//StandInstDbID is a STRING field
type StandInstDbID struct{ message.StringValue }

//Tag returns tag.StandInstDbID (171)
func (f StandInstDbID) Tag() fix.Tag { return tag.StandInstDbID }

//BuildStandInstDbID returns a new StandInstDbID initialized with val
func BuildStandInstDbID(val string) StandInstDbID {
	var field StandInstDbID
	field.Value = val
	return field
}

//StandInstDbName is a STRING field
type StandInstDbName struct{ message.StringValue }

//Tag returns tag.StandInstDbName (170)
func (f StandInstDbName) Tag() fix.Tag { return tag.StandInstDbName }

//BuildStandInstDbName returns a new StandInstDbName initialized with val
func BuildStandInstDbName(val string) StandInstDbName {
	var field StandInstDbName
	field.Value = val
	return field
}

//StandInstDbType is a INT field
type StandInstDbType struct{ message.IntValue }

//Tag returns tag.StandInstDbType (169)
func (f StandInstDbType) Tag() fix.Tag { return tag.StandInstDbType }

//BuildStandInstDbType returns a new StandInstDbType initialized with val
func BuildStandInstDbType(val int) StandInstDbType {
	var field StandInstDbType
	field.Value = val
	return field
}

//StartCash is a AMT field
type StartCash struct{ message.AmtValue }

//Tag returns tag.StartCash (921)
func (f StartCash) Tag() fix.Tag { return tag.StartCash }

//BuildStartCash returns a new StartCash initialized with val
func BuildStartCash(val float64) StartCash {
	var field StartCash
	field.Value = val
	return field
}

//StartDate is a LOCALMKTDATE field
type StartDate struct{ message.LocalMktDateValue }

//Tag returns tag.StartDate (916)
func (f StartDate) Tag() fix.Tag { return tag.StartDate }

//BuildStartDate returns a new StartDate initialized with val
func BuildStartDate(val string) StartDate {
	var field StartDate
	field.Value = val
	return field
}

//StartMaturityMonthYear is a MONTHYEAR field
type StartMaturityMonthYear struct{ message.MonthYearValue }

//Tag returns tag.StartMaturityMonthYear (1241)
func (f StartMaturityMonthYear) Tag() fix.Tag { return tag.StartMaturityMonthYear }

//BuildStartMaturityMonthYear returns a new StartMaturityMonthYear initialized with val
func BuildStartMaturityMonthYear(val string) StartMaturityMonthYear {
	var field StartMaturityMonthYear
	field.Value = val
	return field
}

//StartStrikePxRange is a PRICE field
type StartStrikePxRange struct{ message.PriceValue }

//Tag returns tag.StartStrikePxRange (1202)
func (f StartStrikePxRange) Tag() fix.Tag { return tag.StartStrikePxRange }

//BuildStartStrikePxRange returns a new StartStrikePxRange initialized with val
func BuildStartStrikePxRange(val float64) StartStrikePxRange {
	var field StartStrikePxRange
	field.Value = val
	return field
}

//StartTickPriceRange is a PRICE field
type StartTickPriceRange struct{ message.PriceValue }

//Tag returns tag.StartTickPriceRange (1206)
func (f StartTickPriceRange) Tag() fix.Tag { return tag.StartTickPriceRange }

//BuildStartTickPriceRange returns a new StartTickPriceRange initialized with val
func BuildStartTickPriceRange(val float64) StartTickPriceRange {
	var field StartTickPriceRange
	field.Value = val
	return field
}

//StateOrProvinceOfIssue is a STRING field
type StateOrProvinceOfIssue struct{ message.StringValue }

//Tag returns tag.StateOrProvinceOfIssue (471)
func (f StateOrProvinceOfIssue) Tag() fix.Tag { return tag.StateOrProvinceOfIssue }

//BuildStateOrProvinceOfIssue returns a new StateOrProvinceOfIssue initialized with val
func BuildStateOrProvinceOfIssue(val string) StateOrProvinceOfIssue {
	var field StateOrProvinceOfIssue
	field.Value = val
	return field
}

//StatsType is a INT field
type StatsType struct{ message.IntValue }

//Tag returns tag.StatsType (1176)
func (f StatsType) Tag() fix.Tag { return tag.StatsType }

//BuildStatsType returns a new StatsType initialized with val
func BuildStatsType(val int) StatsType {
	var field StatsType
	field.Value = val
	return field
}

//StatusText is a STRING field
type StatusText struct{ message.StringValue }

//Tag returns tag.StatusText (929)
func (f StatusText) Tag() fix.Tag { return tag.StatusText }

//BuildStatusText returns a new StatusText initialized with val
func BuildStatusText(val string) StatusText {
	var field StatusText
	field.Value = val
	return field
}

//StatusValue is a INT field
type StatusValue struct{ message.IntValue }

//Tag returns tag.StatusValue (928)
func (f StatusValue) Tag() fix.Tag { return tag.StatusValue }

//BuildStatusValue returns a new StatusValue initialized with val
func BuildStatusValue(val int) StatusValue {
	var field StatusValue
	field.Value = val
	return field
}

//StipulationType is a STRING field
type StipulationType struct{ message.StringValue }

//Tag returns tag.StipulationType (233)
func (f StipulationType) Tag() fix.Tag { return tag.StipulationType }

//BuildStipulationType returns a new StipulationType initialized with val
func BuildStipulationType(val string) StipulationType {
	var field StipulationType
	field.Value = val
	return field
}

//StipulationValue is a STRING field
type StipulationValue struct{ message.StringValue }

//Tag returns tag.StipulationValue (234)
func (f StipulationValue) Tag() fix.Tag { return tag.StipulationValue }

//BuildStipulationValue returns a new StipulationValue initialized with val
func BuildStipulationValue(val string) StipulationValue {
	var field StipulationValue
	field.Value = val
	return field
}

//StopPx is a PRICE field
type StopPx struct{ message.PriceValue }

//Tag returns tag.StopPx (99)
func (f StopPx) Tag() fix.Tag { return tag.StopPx }

//BuildStopPx returns a new StopPx initialized with val
func BuildStopPx(val float64) StopPx {
	var field StopPx
	field.Value = val
	return field
}

//StrategyParameterName is a STRING field
type StrategyParameterName struct{ message.StringValue }

//Tag returns tag.StrategyParameterName (958)
func (f StrategyParameterName) Tag() fix.Tag { return tag.StrategyParameterName }

//BuildStrategyParameterName returns a new StrategyParameterName initialized with val
func BuildStrategyParameterName(val string) StrategyParameterName {
	var field StrategyParameterName
	field.Value = val
	return field
}

//StrategyParameterType is a INT field
type StrategyParameterType struct{ message.IntValue }

//Tag returns tag.StrategyParameterType (959)
func (f StrategyParameterType) Tag() fix.Tag { return tag.StrategyParameterType }

//BuildStrategyParameterType returns a new StrategyParameterType initialized with val
func BuildStrategyParameterType(val int) StrategyParameterType {
	var field StrategyParameterType
	field.Value = val
	return field
}

//StrategyParameterValue is a STRING field
type StrategyParameterValue struct{ message.StringValue }

//Tag returns tag.StrategyParameterValue (960)
func (f StrategyParameterValue) Tag() fix.Tag { return tag.StrategyParameterValue }

//BuildStrategyParameterValue returns a new StrategyParameterValue initialized with val
func BuildStrategyParameterValue(val string) StrategyParameterValue {
	var field StrategyParameterValue
	field.Value = val
	return field
}

//StreamAsgnAckType is a INT field
type StreamAsgnAckType struct{ message.IntValue }

//Tag returns tag.StreamAsgnAckType (1503)
func (f StreamAsgnAckType) Tag() fix.Tag { return tag.StreamAsgnAckType }

//BuildStreamAsgnAckType returns a new StreamAsgnAckType initialized with val
func BuildStreamAsgnAckType(val int) StreamAsgnAckType {
	var field StreamAsgnAckType
	field.Value = val
	return field
}

//StreamAsgnRejReason is a INT field
type StreamAsgnRejReason struct{ message.IntValue }

//Tag returns tag.StreamAsgnRejReason (1502)
func (f StreamAsgnRejReason) Tag() fix.Tag { return tag.StreamAsgnRejReason }

//BuildStreamAsgnRejReason returns a new StreamAsgnRejReason initialized with val
func BuildStreamAsgnRejReason(val int) StreamAsgnRejReason {
	var field StreamAsgnRejReason
	field.Value = val
	return field
}

//StreamAsgnReqID is a STRING field
type StreamAsgnReqID struct{ message.StringValue }

//Tag returns tag.StreamAsgnReqID (1497)
func (f StreamAsgnReqID) Tag() fix.Tag { return tag.StreamAsgnReqID }

//BuildStreamAsgnReqID returns a new StreamAsgnReqID initialized with val
func BuildStreamAsgnReqID(val string) StreamAsgnReqID {
	var field StreamAsgnReqID
	field.Value = val
	return field
}

//StreamAsgnReqType is a INT field
type StreamAsgnReqType struct{ message.IntValue }

//Tag returns tag.StreamAsgnReqType (1498)
func (f StreamAsgnReqType) Tag() fix.Tag { return tag.StreamAsgnReqType }

//BuildStreamAsgnReqType returns a new StreamAsgnReqType initialized with val
func BuildStreamAsgnReqType(val int) StreamAsgnReqType {
	var field StreamAsgnReqType
	field.Value = val
	return field
}

//StreamAsgnRptID is a STRING field
type StreamAsgnRptID struct{ message.StringValue }

//Tag returns tag.StreamAsgnRptID (1501)
func (f StreamAsgnRptID) Tag() fix.Tag { return tag.StreamAsgnRptID }

//BuildStreamAsgnRptID returns a new StreamAsgnRptID initialized with val
func BuildStreamAsgnRptID(val string) StreamAsgnRptID {
	var field StreamAsgnRptID
	field.Value = val
	return field
}

//StreamAsgnType is a INT field
type StreamAsgnType struct{ message.IntValue }

//Tag returns tag.StreamAsgnType (1617)
func (f StreamAsgnType) Tag() fix.Tag { return tag.StreamAsgnType }

//BuildStreamAsgnType returns a new StreamAsgnType initialized with val
func BuildStreamAsgnType(val int) StreamAsgnType {
	var field StreamAsgnType
	field.Value = val
	return field
}

//StrikeCurrency is a CURRENCY field
type StrikeCurrency struct{ message.CurrencyValue }

//Tag returns tag.StrikeCurrency (947)
func (f StrikeCurrency) Tag() fix.Tag { return tag.StrikeCurrency }

//BuildStrikeCurrency returns a new StrikeCurrency initialized with val
func BuildStrikeCurrency(val string) StrikeCurrency {
	var field StrikeCurrency
	field.Value = val
	return field
}

//StrikeExerciseStyle is a INT field
type StrikeExerciseStyle struct{ message.IntValue }

//Tag returns tag.StrikeExerciseStyle (1304)
func (f StrikeExerciseStyle) Tag() fix.Tag { return tag.StrikeExerciseStyle }

//BuildStrikeExerciseStyle returns a new StrikeExerciseStyle initialized with val
func BuildStrikeExerciseStyle(val int) StrikeExerciseStyle {
	var field StrikeExerciseStyle
	field.Value = val
	return field
}

//StrikeIncrement is a FLOAT field
type StrikeIncrement struct{ message.FloatValue }

//Tag returns tag.StrikeIncrement (1204)
func (f StrikeIncrement) Tag() fix.Tag { return tag.StrikeIncrement }

//BuildStrikeIncrement returns a new StrikeIncrement initialized with val
func BuildStrikeIncrement(val float64) StrikeIncrement {
	var field StrikeIncrement
	field.Value = val
	return field
}

//StrikeMultiplier is a FLOAT field
type StrikeMultiplier struct{ message.FloatValue }

//Tag returns tag.StrikeMultiplier (967)
func (f StrikeMultiplier) Tag() fix.Tag { return tag.StrikeMultiplier }

//BuildStrikeMultiplier returns a new StrikeMultiplier initialized with val
func BuildStrikeMultiplier(val float64) StrikeMultiplier {
	var field StrikeMultiplier
	field.Value = val
	return field
}

//StrikePrice is a PRICE field
type StrikePrice struct{ message.PriceValue }

//Tag returns tag.StrikePrice (202)
func (f StrikePrice) Tag() fix.Tag { return tag.StrikePrice }

//BuildStrikePrice returns a new StrikePrice initialized with val
func BuildStrikePrice(val float64) StrikePrice {
	var field StrikePrice
	field.Value = val
	return field
}

//StrikePriceBoundaryMethod is a INT field
type StrikePriceBoundaryMethod struct{ message.IntValue }

//Tag returns tag.StrikePriceBoundaryMethod (1479)
func (f StrikePriceBoundaryMethod) Tag() fix.Tag { return tag.StrikePriceBoundaryMethod }

//BuildStrikePriceBoundaryMethod returns a new StrikePriceBoundaryMethod initialized with val
func BuildStrikePriceBoundaryMethod(val int) StrikePriceBoundaryMethod {
	var field StrikePriceBoundaryMethod
	field.Value = val
	return field
}

//StrikePriceBoundaryPrecision is a PERCENTAGE field
type StrikePriceBoundaryPrecision struct{ message.PercentageValue }

//Tag returns tag.StrikePriceBoundaryPrecision (1480)
func (f StrikePriceBoundaryPrecision) Tag() fix.Tag { return tag.StrikePriceBoundaryPrecision }

//BuildStrikePriceBoundaryPrecision returns a new StrikePriceBoundaryPrecision initialized with val
func BuildStrikePriceBoundaryPrecision(val float64) StrikePriceBoundaryPrecision {
	var field StrikePriceBoundaryPrecision
	field.Value = val
	return field
}

//StrikePriceDeterminationMethod is a INT field
type StrikePriceDeterminationMethod struct{ message.IntValue }

//Tag returns tag.StrikePriceDeterminationMethod (1478)
func (f StrikePriceDeterminationMethod) Tag() fix.Tag { return tag.StrikePriceDeterminationMethod }

//BuildStrikePriceDeterminationMethod returns a new StrikePriceDeterminationMethod initialized with val
func BuildStrikePriceDeterminationMethod(val int) StrikePriceDeterminationMethod {
	var field StrikePriceDeterminationMethod
	field.Value = val
	return field
}

//StrikeRuleID is a STRING field
type StrikeRuleID struct{ message.StringValue }

//Tag returns tag.StrikeRuleID (1223)
func (f StrikeRuleID) Tag() fix.Tag { return tag.StrikeRuleID }

//BuildStrikeRuleID returns a new StrikeRuleID initialized with val
func BuildStrikeRuleID(val string) StrikeRuleID {
	var field StrikeRuleID
	field.Value = val
	return field
}

//StrikeTime is a UTCTIMESTAMP field
type StrikeTime struct{ message.UTCTimestampValue }

//Tag returns tag.StrikeTime (443)
func (f StrikeTime) Tag() fix.Tag { return tag.StrikeTime }

//StrikeValue is a FLOAT field
type StrikeValue struct{ message.FloatValue }

//Tag returns tag.StrikeValue (968)
func (f StrikeValue) Tag() fix.Tag { return tag.StrikeValue }

//BuildStrikeValue returns a new StrikeValue initialized with val
func BuildStrikeValue(val float64) StrikeValue {
	var field StrikeValue
	field.Value = val
	return field
}

//Subject is a STRING field
type Subject struct{ message.StringValue }

//Tag returns tag.Subject (147)
func (f Subject) Tag() fix.Tag { return tag.Subject }

//BuildSubject returns a new Subject initialized with val
func BuildSubject(val string) Subject {
	var field Subject
	field.Value = val
	return field
}

//SubscriptionRequestType is a CHAR field
type SubscriptionRequestType struct{ message.CharValue }

//Tag returns tag.SubscriptionRequestType (263)
func (f SubscriptionRequestType) Tag() fix.Tag { return tag.SubscriptionRequestType }

//BuildSubscriptionRequestType returns a new SubscriptionRequestType initialized with val
func BuildSubscriptionRequestType(val string) SubscriptionRequestType {
	var field SubscriptionRequestType
	field.Value = val
	return field
}

//SwapPoints is a PRICEOFFSET field
type SwapPoints struct{ message.PriceOffsetValue }

//Tag returns tag.SwapPoints (1069)
func (f SwapPoints) Tag() fix.Tag { return tag.SwapPoints }

//BuildSwapPoints returns a new SwapPoints initialized with val
func BuildSwapPoints(val float64) SwapPoints {
	var field SwapPoints
	field.Value = val
	return field
}

//Symbol is a STRING field
type Symbol struct{ message.StringValue }

//Tag returns tag.Symbol (55)
func (f Symbol) Tag() fix.Tag { return tag.Symbol }

//BuildSymbol returns a new Symbol initialized with val
func BuildSymbol(val string) Symbol {
	var field Symbol
	field.Value = val
	return field
}

//SymbolSfx is a STRING field
type SymbolSfx struct{ message.StringValue }

//Tag returns tag.SymbolSfx (65)
func (f SymbolSfx) Tag() fix.Tag { return tag.SymbolSfx }

//BuildSymbolSfx returns a new SymbolSfx initialized with val
func BuildSymbolSfx(val string) SymbolSfx {
	var field SymbolSfx
	field.Value = val
	return field
}

//TZTransactTime is a TZTIMESTAMP field
type TZTransactTime struct{ message.TZTimestampValue }

//Tag returns tag.TZTransactTime (1132)
func (f TZTransactTime) Tag() fix.Tag { return tag.TZTransactTime }

//TargetCompID is a STRING field
type TargetCompID struct{ message.StringValue }

//Tag returns tag.TargetCompID (56)
func (f TargetCompID) Tag() fix.Tag { return tag.TargetCompID }

//BuildTargetCompID returns a new TargetCompID initialized with val
func BuildTargetCompID(val string) TargetCompID {
	var field TargetCompID
	field.Value = val
	return field
}

//TargetLocationID is a STRING field
type TargetLocationID struct{ message.StringValue }

//Tag returns tag.TargetLocationID (143)
func (f TargetLocationID) Tag() fix.Tag { return tag.TargetLocationID }

//BuildTargetLocationID returns a new TargetLocationID initialized with val
func BuildTargetLocationID(val string) TargetLocationID {
	var field TargetLocationID
	field.Value = val
	return field
}

//TargetPartyID is a STRING field
type TargetPartyID struct{ message.StringValue }

//Tag returns tag.TargetPartyID (1462)
func (f TargetPartyID) Tag() fix.Tag { return tag.TargetPartyID }

//BuildTargetPartyID returns a new TargetPartyID initialized with val
func BuildTargetPartyID(val string) TargetPartyID {
	var field TargetPartyID
	field.Value = val
	return field
}

//TargetPartyIDSource is a CHAR field
type TargetPartyIDSource struct{ message.CharValue }

//Tag returns tag.TargetPartyIDSource (1463)
func (f TargetPartyIDSource) Tag() fix.Tag { return tag.TargetPartyIDSource }

//BuildTargetPartyIDSource returns a new TargetPartyIDSource initialized with val
func BuildTargetPartyIDSource(val string) TargetPartyIDSource {
	var field TargetPartyIDSource
	field.Value = val
	return field
}

//TargetPartyRole is a INT field
type TargetPartyRole struct{ message.IntValue }

//Tag returns tag.TargetPartyRole (1464)
func (f TargetPartyRole) Tag() fix.Tag { return tag.TargetPartyRole }

//BuildTargetPartyRole returns a new TargetPartyRole initialized with val
func BuildTargetPartyRole(val int) TargetPartyRole {
	var field TargetPartyRole
	field.Value = val
	return field
}

//TargetStrategy is a INT field
type TargetStrategy struct{ message.IntValue }

//Tag returns tag.TargetStrategy (847)
func (f TargetStrategy) Tag() fix.Tag { return tag.TargetStrategy }

//BuildTargetStrategy returns a new TargetStrategy initialized with val
func BuildTargetStrategy(val int) TargetStrategy {
	var field TargetStrategy
	field.Value = val
	return field
}

//TargetStrategyParameters is a STRING field
type TargetStrategyParameters struct{ message.StringValue }

//Tag returns tag.TargetStrategyParameters (848)
func (f TargetStrategyParameters) Tag() fix.Tag { return tag.TargetStrategyParameters }

//BuildTargetStrategyParameters returns a new TargetStrategyParameters initialized with val
func BuildTargetStrategyParameters(val string) TargetStrategyParameters {
	var field TargetStrategyParameters
	field.Value = val
	return field
}

//TargetStrategyPerformance is a FLOAT field
type TargetStrategyPerformance struct{ message.FloatValue }

//Tag returns tag.TargetStrategyPerformance (850)
func (f TargetStrategyPerformance) Tag() fix.Tag { return tag.TargetStrategyPerformance }

//BuildTargetStrategyPerformance returns a new TargetStrategyPerformance initialized with val
func BuildTargetStrategyPerformance(val float64) TargetStrategyPerformance {
	var field TargetStrategyPerformance
	field.Value = val
	return field
}

//TargetSubID is a STRING field
type TargetSubID struct{ message.StringValue }

//Tag returns tag.TargetSubID (57)
func (f TargetSubID) Tag() fix.Tag { return tag.TargetSubID }

//BuildTargetSubID returns a new TargetSubID initialized with val
func BuildTargetSubID(val string) TargetSubID {
	var field TargetSubID
	field.Value = val
	return field
}

//TaxAdvantageType is a INT field
type TaxAdvantageType struct{ message.IntValue }

//Tag returns tag.TaxAdvantageType (495)
func (f TaxAdvantageType) Tag() fix.Tag { return tag.TaxAdvantageType }

//BuildTaxAdvantageType returns a new TaxAdvantageType initialized with val
func BuildTaxAdvantageType(val int) TaxAdvantageType {
	var field TaxAdvantageType
	field.Value = val
	return field
}

//TerminationType is a INT field
type TerminationType struct{ message.IntValue }

//Tag returns tag.TerminationType (788)
func (f TerminationType) Tag() fix.Tag { return tag.TerminationType }

//BuildTerminationType returns a new TerminationType initialized with val
func BuildTerminationType(val int) TerminationType {
	var field TerminationType
	field.Value = val
	return field
}

//TestMessageIndicator is a BOOLEAN field
type TestMessageIndicator struct{ message.BooleanValue }

//Tag returns tag.TestMessageIndicator (464)
func (f TestMessageIndicator) Tag() fix.Tag { return tag.TestMessageIndicator }

//BuildTestMessageIndicator returns a new TestMessageIndicator initialized with val
func BuildTestMessageIndicator(val bool) TestMessageIndicator {
	var field TestMessageIndicator
	field.Value = val
	return field
}

//TestReqID is a STRING field
type TestReqID struct{ message.StringValue }

//Tag returns tag.TestReqID (112)
func (f TestReqID) Tag() fix.Tag { return tag.TestReqID }

//BuildTestReqID returns a new TestReqID initialized with val
func BuildTestReqID(val string) TestReqID {
	var field TestReqID
	field.Value = val
	return field
}

//Text is a STRING field
type Text struct{ message.StringValue }

//Tag returns tag.Text (58)
func (f Text) Tag() fix.Tag { return tag.Text }

//BuildText returns a new Text initialized with val
func BuildText(val string) Text {
	var field Text
	field.Value = val
	return field
}

//ThresholdAmount is a PRICEOFFSET field
type ThresholdAmount struct{ message.PriceOffsetValue }

//Tag returns tag.ThresholdAmount (834)
func (f ThresholdAmount) Tag() fix.Tag { return tag.ThresholdAmount }

//BuildThresholdAmount returns a new ThresholdAmount initialized with val
func BuildThresholdAmount(val float64) ThresholdAmount {
	var field ThresholdAmount
	field.Value = val
	return field
}

//TickDirection is a CHAR field
type TickDirection struct{ message.CharValue }

//Tag returns tag.TickDirection (274)
func (f TickDirection) Tag() fix.Tag { return tag.TickDirection }

//BuildTickDirection returns a new TickDirection initialized with val
func BuildTickDirection(val string) TickDirection {
	var field TickDirection
	field.Value = val
	return field
}

//TickIncrement is a PRICE field
type TickIncrement struct{ message.PriceValue }

//Tag returns tag.TickIncrement (1208)
func (f TickIncrement) Tag() fix.Tag { return tag.TickIncrement }

//BuildTickIncrement returns a new TickIncrement initialized with val
func BuildTickIncrement(val float64) TickIncrement {
	var field TickIncrement
	field.Value = val
	return field
}

//TickRuleType is a INT field
type TickRuleType struct{ message.IntValue }

//Tag returns tag.TickRuleType (1209)
func (f TickRuleType) Tag() fix.Tag { return tag.TickRuleType }

//BuildTickRuleType returns a new TickRuleType initialized with val
func BuildTickRuleType(val int) TickRuleType {
	var field TickRuleType
	field.Value = val
	return field
}

//TierCode is a STRING field
type TierCode struct{ message.StringValue }

//Tag returns tag.TierCode (994)
func (f TierCode) Tag() fix.Tag { return tag.TierCode }

//BuildTierCode returns a new TierCode initialized with val
func BuildTierCode(val string) TierCode {
	var field TierCode
	field.Value = val
	return field
}

//TimeBracket is a STRING field
type TimeBracket struct{ message.StringValue }

//Tag returns tag.TimeBracket (943)
func (f TimeBracket) Tag() fix.Tag { return tag.TimeBracket }

//BuildTimeBracket returns a new TimeBracket initialized with val
func BuildTimeBracket(val string) TimeBracket {
	var field TimeBracket
	field.Value = val
	return field
}

//TimeInForce is a CHAR field
type TimeInForce struct{ message.CharValue }

//Tag returns tag.TimeInForce (59)
func (f TimeInForce) Tag() fix.Tag { return tag.TimeInForce }

//BuildTimeInForce returns a new TimeInForce initialized with val
func BuildTimeInForce(val string) TimeInForce {
	var field TimeInForce
	field.Value = val
	return field
}

//TimeToExpiration is a FLOAT field
type TimeToExpiration struct{ message.FloatValue }

//Tag returns tag.TimeToExpiration (1189)
func (f TimeToExpiration) Tag() fix.Tag { return tag.TimeToExpiration }

//BuildTimeToExpiration returns a new TimeToExpiration initialized with val
func BuildTimeToExpiration(val float64) TimeToExpiration {
	var field TimeToExpiration
	field.Value = val
	return field
}

//TimeUnit is a STRING field
type TimeUnit struct{ message.StringValue }

//Tag returns tag.TimeUnit (997)
func (f TimeUnit) Tag() fix.Tag { return tag.TimeUnit }

//BuildTimeUnit returns a new TimeUnit initialized with val
func BuildTimeUnit(val string) TimeUnit {
	var field TimeUnit
	field.Value = val
	return field
}

//TotNoAccQuotes is a INT field
type TotNoAccQuotes struct{ message.IntValue }

//Tag returns tag.TotNoAccQuotes (1169)
func (f TotNoAccQuotes) Tag() fix.Tag { return tag.TotNoAccQuotes }

//BuildTotNoAccQuotes returns a new TotNoAccQuotes initialized with val
func BuildTotNoAccQuotes(val int) TotNoAccQuotes {
	var field TotNoAccQuotes
	field.Value = val
	return field
}

//TotNoAllocs is a INT field
type TotNoAllocs struct{ message.IntValue }

//Tag returns tag.TotNoAllocs (892)
func (f TotNoAllocs) Tag() fix.Tag { return tag.TotNoAllocs }

//BuildTotNoAllocs returns a new TotNoAllocs initialized with val
func BuildTotNoAllocs(val int) TotNoAllocs {
	var field TotNoAllocs
	field.Value = val
	return field
}

//TotNoCxldQuotes is a INT field
type TotNoCxldQuotes struct{ message.IntValue }

//Tag returns tag.TotNoCxldQuotes (1168)
func (f TotNoCxldQuotes) Tag() fix.Tag { return tag.TotNoCxldQuotes }

//BuildTotNoCxldQuotes returns a new TotNoCxldQuotes initialized with val
func BuildTotNoCxldQuotes(val int) TotNoCxldQuotes {
	var field TotNoCxldQuotes
	field.Value = val
	return field
}

//TotNoFills is a INT field
type TotNoFills struct{ message.IntValue }

//Tag returns tag.TotNoFills (1361)
func (f TotNoFills) Tag() fix.Tag { return tag.TotNoFills }

//BuildTotNoFills returns a new TotNoFills initialized with val
func BuildTotNoFills(val int) TotNoFills {
	var field TotNoFills
	field.Value = val
	return field
}

//TotNoOrders is a INT field
type TotNoOrders struct{ message.IntValue }

//Tag returns tag.TotNoOrders (68)
func (f TotNoOrders) Tag() fix.Tag { return tag.TotNoOrders }

//BuildTotNoOrders returns a new TotNoOrders initialized with val
func BuildTotNoOrders(val int) TotNoOrders {
	var field TotNoOrders
	field.Value = val
	return field
}

//TotNoPartyList is a INT field
type TotNoPartyList struct{ message.IntValue }

//Tag returns tag.TotNoPartyList (1512)
func (f TotNoPartyList) Tag() fix.Tag { return tag.TotNoPartyList }

//BuildTotNoPartyList returns a new TotNoPartyList initialized with val
func BuildTotNoPartyList(val int) TotNoPartyList {
	var field TotNoPartyList
	field.Value = val
	return field
}

//TotNoQuoteEntries is a INT field
type TotNoQuoteEntries struct{ message.IntValue }

//Tag returns tag.TotNoQuoteEntries (304)
func (f TotNoQuoteEntries) Tag() fix.Tag { return tag.TotNoQuoteEntries }

//BuildTotNoQuoteEntries returns a new TotNoQuoteEntries initialized with val
func BuildTotNoQuoteEntries(val int) TotNoQuoteEntries {
	var field TotNoQuoteEntries
	field.Value = val
	return field
}

//TotNoRejQuotes is a INT field
type TotNoRejQuotes struct{ message.IntValue }

//Tag returns tag.TotNoRejQuotes (1170)
func (f TotNoRejQuotes) Tag() fix.Tag { return tag.TotNoRejQuotes }

//BuildTotNoRejQuotes returns a new TotNoRejQuotes initialized with val
func BuildTotNoRejQuotes(val int) TotNoRejQuotes {
	var field TotNoRejQuotes
	field.Value = val
	return field
}

//TotNoRelatedSym is a INT field
type TotNoRelatedSym struct{ message.IntValue }

//Tag returns tag.TotNoRelatedSym (393)
func (f TotNoRelatedSym) Tag() fix.Tag { return tag.TotNoRelatedSym }

//BuildTotNoRelatedSym returns a new TotNoRelatedSym initialized with val
func BuildTotNoRelatedSym(val int) TotNoRelatedSym {
	var field TotNoRelatedSym
	field.Value = val
	return field
}

//TotNoSecurityTypes is a INT field
type TotNoSecurityTypes struct{ message.IntValue }

//Tag returns tag.TotNoSecurityTypes (557)
func (f TotNoSecurityTypes) Tag() fix.Tag { return tag.TotNoSecurityTypes }

//BuildTotNoSecurityTypes returns a new TotNoSecurityTypes initialized with val
func BuildTotNoSecurityTypes(val int) TotNoSecurityTypes {
	var field TotNoSecurityTypes
	field.Value = val
	return field
}

//TotNoStrikes is a INT field
type TotNoStrikes struct{ message.IntValue }

//Tag returns tag.TotNoStrikes (422)
func (f TotNoStrikes) Tag() fix.Tag { return tag.TotNoStrikes }

//BuildTotNoStrikes returns a new TotNoStrikes initialized with val
func BuildTotNoStrikes(val int) TotNoStrikes {
	var field TotNoStrikes
	field.Value = val
	return field
}

//TotNumAssignmentReports is a INT field
type TotNumAssignmentReports struct{ message.IntValue }

//Tag returns tag.TotNumAssignmentReports (832)
func (f TotNumAssignmentReports) Tag() fix.Tag { return tag.TotNumAssignmentReports }

//BuildTotNumAssignmentReports returns a new TotNumAssignmentReports initialized with val
func BuildTotNumAssignmentReports(val int) TotNumAssignmentReports {
	var field TotNumAssignmentReports
	field.Value = val
	return field
}

//TotNumReports is a INT field
type TotNumReports struct{ message.IntValue }

//Tag returns tag.TotNumReports (911)
func (f TotNumReports) Tag() fix.Tag { return tag.TotNumReports }

//BuildTotNumReports returns a new TotNumReports initialized with val
func BuildTotNumReports(val int) TotNumReports {
	var field TotNumReports
	field.Value = val
	return field
}

//TotNumTradeReports is a INT field
type TotNumTradeReports struct{ message.IntValue }

//Tag returns tag.TotNumTradeReports (748)
func (f TotNumTradeReports) Tag() fix.Tag { return tag.TotNumTradeReports }

//BuildTotNumTradeReports returns a new TotNumTradeReports initialized with val
func BuildTotNumTradeReports(val int) TotNumTradeReports {
	var field TotNumTradeReports
	field.Value = val
	return field
}

//TotQuoteEntries is a INT field
type TotQuoteEntries struct{ message.IntValue }

//Tag returns tag.TotQuoteEntries (304)
func (f TotQuoteEntries) Tag() fix.Tag { return tag.TotQuoteEntries }

//BuildTotQuoteEntries returns a new TotQuoteEntries initialized with val
func BuildTotQuoteEntries(val int) TotQuoteEntries {
	var field TotQuoteEntries
	field.Value = val
	return field
}

//TotalAccruedInterestAmt is a AMT field
type TotalAccruedInterestAmt struct{ message.AmtValue }

//Tag returns tag.TotalAccruedInterestAmt (540)
func (f TotalAccruedInterestAmt) Tag() fix.Tag { return tag.TotalAccruedInterestAmt }

//BuildTotalAccruedInterestAmt returns a new TotalAccruedInterestAmt initialized with val
func BuildTotalAccruedInterestAmt(val float64) TotalAccruedInterestAmt {
	var field TotalAccruedInterestAmt
	field.Value = val
	return field
}

//TotalAffectedOrders is a INT field
type TotalAffectedOrders struct{ message.IntValue }

//Tag returns tag.TotalAffectedOrders (533)
func (f TotalAffectedOrders) Tag() fix.Tag { return tag.TotalAffectedOrders }

//BuildTotalAffectedOrders returns a new TotalAffectedOrders initialized with val
func BuildTotalAffectedOrders(val int) TotalAffectedOrders {
	var field TotalAffectedOrders
	field.Value = val
	return field
}

//TotalNetValue is a AMT field
type TotalNetValue struct{ message.AmtValue }

//Tag returns tag.TotalNetValue (900)
func (f TotalNetValue) Tag() fix.Tag { return tag.TotalNetValue }

//BuildTotalNetValue returns a new TotalNetValue initialized with val
func BuildTotalNetValue(val float64) TotalNetValue {
	var field TotalNetValue
	field.Value = val
	return field
}

//TotalNumPosReports is a INT field
type TotalNumPosReports struct{ message.IntValue }

//Tag returns tag.TotalNumPosReports (727)
func (f TotalNumPosReports) Tag() fix.Tag { return tag.TotalNumPosReports }

//BuildTotalNumPosReports returns a new TotalNumPosReports initialized with val
func BuildTotalNumPosReports(val int) TotalNumPosReports {
	var field TotalNumPosReports
	field.Value = val
	return field
}

//TotalNumSecurities is a INT field
type TotalNumSecurities struct{ message.IntValue }

//Tag returns tag.TotalNumSecurities (393)
func (f TotalNumSecurities) Tag() fix.Tag { return tag.TotalNumSecurities }

//BuildTotalNumSecurities returns a new TotalNumSecurities initialized with val
func BuildTotalNumSecurities(val int) TotalNumSecurities {
	var field TotalNumSecurities
	field.Value = val
	return field
}

//TotalNumSecurityTypes is a INT field
type TotalNumSecurityTypes struct{ message.IntValue }

//Tag returns tag.TotalNumSecurityTypes (557)
func (f TotalNumSecurityTypes) Tag() fix.Tag { return tag.TotalNumSecurityTypes }

//BuildTotalNumSecurityTypes returns a new TotalNumSecurityTypes initialized with val
func BuildTotalNumSecurityTypes(val int) TotalNumSecurityTypes {
	var field TotalNumSecurityTypes
	field.Value = val
	return field
}

//TotalTakedown is a AMT field
type TotalTakedown struct{ message.AmtValue }

//Tag returns tag.TotalTakedown (237)
func (f TotalTakedown) Tag() fix.Tag { return tag.TotalTakedown }

//BuildTotalTakedown returns a new TotalTakedown initialized with val
func BuildTotalTakedown(val float64) TotalTakedown {
	var field TotalTakedown
	field.Value = val
	return field
}

//TotalVolumeTraded is a QTY field
type TotalVolumeTraded struct{ message.QtyValue }

//Tag returns tag.TotalVolumeTraded (387)
func (f TotalVolumeTraded) Tag() fix.Tag { return tag.TotalVolumeTraded }

//BuildTotalVolumeTraded returns a new TotalVolumeTraded initialized with val
func BuildTotalVolumeTraded(val float64) TotalVolumeTraded {
	var field TotalVolumeTraded
	field.Value = val
	return field
}

//TotalVolumeTradedDate is a UTCDATEONLY field
type TotalVolumeTradedDate struct{ message.UTCDateOnlyValue }

//Tag returns tag.TotalVolumeTradedDate (449)
func (f TotalVolumeTradedDate) Tag() fix.Tag { return tag.TotalVolumeTradedDate }

//TotalVolumeTradedTime is a UTCTIMEONLY field
type TotalVolumeTradedTime struct{ message.UTCTimeOnlyValue }

//Tag returns tag.TotalVolumeTradedTime (450)
func (f TotalVolumeTradedTime) Tag() fix.Tag { return tag.TotalVolumeTradedTime }

//TradSesCloseTime is a UTCTIMESTAMP field
type TradSesCloseTime struct{ message.UTCTimestampValue }

//Tag returns tag.TradSesCloseTime (344)
func (f TradSesCloseTime) Tag() fix.Tag { return tag.TradSesCloseTime }

//TradSesEndTime is a UTCTIMESTAMP field
type TradSesEndTime struct{ message.UTCTimestampValue }

//Tag returns tag.TradSesEndTime (345)
func (f TradSesEndTime) Tag() fix.Tag { return tag.TradSesEndTime }

//TradSesEvent is a INT field
type TradSesEvent struct{ message.IntValue }

//Tag returns tag.TradSesEvent (1368)
func (f TradSesEvent) Tag() fix.Tag { return tag.TradSesEvent }

//BuildTradSesEvent returns a new TradSesEvent initialized with val
func BuildTradSesEvent(val int) TradSesEvent {
	var field TradSesEvent
	field.Value = val
	return field
}

//TradSesMethod is a INT field
type TradSesMethod struct{ message.IntValue }

//Tag returns tag.TradSesMethod (338)
func (f TradSesMethod) Tag() fix.Tag { return tag.TradSesMethod }

//BuildTradSesMethod returns a new TradSesMethod initialized with val
func BuildTradSesMethod(val int) TradSesMethod {
	var field TradSesMethod
	field.Value = val
	return field
}

//TradSesMode is a INT field
type TradSesMode struct{ message.IntValue }

//Tag returns tag.TradSesMode (339)
func (f TradSesMode) Tag() fix.Tag { return tag.TradSesMode }

//BuildTradSesMode returns a new TradSesMode initialized with val
func BuildTradSesMode(val int) TradSesMode {
	var field TradSesMode
	field.Value = val
	return field
}

//TradSesOpenTime is a UTCTIMESTAMP field
type TradSesOpenTime struct{ message.UTCTimestampValue }

//Tag returns tag.TradSesOpenTime (342)
func (f TradSesOpenTime) Tag() fix.Tag { return tag.TradSesOpenTime }

//TradSesPreCloseTime is a UTCTIMESTAMP field
type TradSesPreCloseTime struct{ message.UTCTimestampValue }

//Tag returns tag.TradSesPreCloseTime (343)
func (f TradSesPreCloseTime) Tag() fix.Tag { return tag.TradSesPreCloseTime }

//TradSesReqID is a STRING field
type TradSesReqID struct{ message.StringValue }

//Tag returns tag.TradSesReqID (335)
func (f TradSesReqID) Tag() fix.Tag { return tag.TradSesReqID }

//BuildTradSesReqID returns a new TradSesReqID initialized with val
func BuildTradSesReqID(val string) TradSesReqID {
	var field TradSesReqID
	field.Value = val
	return field
}

//TradSesStartTime is a UTCTIMESTAMP field
type TradSesStartTime struct{ message.UTCTimestampValue }

//Tag returns tag.TradSesStartTime (341)
func (f TradSesStartTime) Tag() fix.Tag { return tag.TradSesStartTime }

//TradSesStatus is a INT field
type TradSesStatus struct{ message.IntValue }

//Tag returns tag.TradSesStatus (340)
func (f TradSesStatus) Tag() fix.Tag { return tag.TradSesStatus }

//BuildTradSesStatus returns a new TradSesStatus initialized with val
func BuildTradSesStatus(val int) TradSesStatus {
	var field TradSesStatus
	field.Value = val
	return field
}

//TradSesStatusRejReason is a INT field
type TradSesStatusRejReason struct{ message.IntValue }

//Tag returns tag.TradSesStatusRejReason (567)
func (f TradSesStatusRejReason) Tag() fix.Tag { return tag.TradSesStatusRejReason }

//BuildTradSesStatusRejReason returns a new TradSesStatusRejReason initialized with val
func BuildTradSesStatusRejReason(val int) TradSesStatusRejReason {
	var field TradSesStatusRejReason
	field.Value = val
	return field
}

//TradSesUpdateAction is a CHAR field
type TradSesUpdateAction struct{ message.CharValue }

//Tag returns tag.TradSesUpdateAction (1327)
func (f TradSesUpdateAction) Tag() fix.Tag { return tag.TradSesUpdateAction }

//BuildTradSesUpdateAction returns a new TradSesUpdateAction initialized with val
func BuildTradSesUpdateAction(val string) TradSesUpdateAction {
	var field TradSesUpdateAction
	field.Value = val
	return field
}

//TradeAllocIndicator is a INT field
type TradeAllocIndicator struct{ message.IntValue }

//Tag returns tag.TradeAllocIndicator (826)
func (f TradeAllocIndicator) Tag() fix.Tag { return tag.TradeAllocIndicator }

//BuildTradeAllocIndicator returns a new TradeAllocIndicator initialized with val
func BuildTradeAllocIndicator(val int) TradeAllocIndicator {
	var field TradeAllocIndicator
	field.Value = val
	return field
}

//TradeCondition is a MULTIPLESTRINGVALUE field
type TradeCondition struct{ message.MultipleStringValue }

//Tag returns tag.TradeCondition (277)
func (f TradeCondition) Tag() fix.Tag { return tag.TradeCondition }

//BuildTradeCondition returns a new TradeCondition initialized with val
func BuildTradeCondition(val string) TradeCondition {
	var field TradeCondition
	field.Value = val
	return field
}

//TradeDate is a LOCALMKTDATE field
type TradeDate struct{ message.LocalMktDateValue }

//Tag returns tag.TradeDate (75)
func (f TradeDate) Tag() fix.Tag { return tag.TradeDate }

//BuildTradeDate returns a new TradeDate initialized with val
func BuildTradeDate(val string) TradeDate {
	var field TradeDate
	field.Value = val
	return field
}

//TradeHandlingInstr is a CHAR field
type TradeHandlingInstr struct{ message.CharValue }

//Tag returns tag.TradeHandlingInstr (1123)
func (f TradeHandlingInstr) Tag() fix.Tag { return tag.TradeHandlingInstr }

//BuildTradeHandlingInstr returns a new TradeHandlingInstr initialized with val
func BuildTradeHandlingInstr(val string) TradeHandlingInstr {
	var field TradeHandlingInstr
	field.Value = val
	return field
}

//TradeID is a STRING field
type TradeID struct{ message.StringValue }

//Tag returns tag.TradeID (1003)
func (f TradeID) Tag() fix.Tag { return tag.TradeID }

//BuildTradeID returns a new TradeID initialized with val
func BuildTradeID(val string) TradeID {
	var field TradeID
	field.Value = val
	return field
}

//TradeInputDevice is a STRING field
type TradeInputDevice struct{ message.StringValue }

//Tag returns tag.TradeInputDevice (579)
func (f TradeInputDevice) Tag() fix.Tag { return tag.TradeInputDevice }

//BuildTradeInputDevice returns a new TradeInputDevice initialized with val
func BuildTradeInputDevice(val string) TradeInputDevice {
	var field TradeInputDevice
	field.Value = val
	return field
}

//TradeInputSource is a STRING field
type TradeInputSource struct{ message.StringValue }

//Tag returns tag.TradeInputSource (578)
func (f TradeInputSource) Tag() fix.Tag { return tag.TradeInputSource }

//BuildTradeInputSource returns a new TradeInputSource initialized with val
func BuildTradeInputSource(val string) TradeInputSource {
	var field TradeInputSource
	field.Value = val
	return field
}

//TradeLegRefID is a STRING field
type TradeLegRefID struct{ message.StringValue }

//Tag returns tag.TradeLegRefID (824)
func (f TradeLegRefID) Tag() fix.Tag { return tag.TradeLegRefID }

//BuildTradeLegRefID returns a new TradeLegRefID initialized with val
func BuildTradeLegRefID(val string) TradeLegRefID {
	var field TradeLegRefID
	field.Value = val
	return field
}

//TradeLinkID is a STRING field
type TradeLinkID struct{ message.StringValue }

//Tag returns tag.TradeLinkID (820)
func (f TradeLinkID) Tag() fix.Tag { return tag.TradeLinkID }

//BuildTradeLinkID returns a new TradeLinkID initialized with val
func BuildTradeLinkID(val string) TradeLinkID {
	var field TradeLinkID
	field.Value = val
	return field
}

//TradeOriginationDate is a LOCALMKTDATE field
type TradeOriginationDate struct{ message.LocalMktDateValue }

//Tag returns tag.TradeOriginationDate (229)
func (f TradeOriginationDate) Tag() fix.Tag { return tag.TradeOriginationDate }

//BuildTradeOriginationDate returns a new TradeOriginationDate initialized with val
func BuildTradeOriginationDate(val string) TradeOriginationDate {
	var field TradeOriginationDate
	field.Value = val
	return field
}

//TradePublishIndicator is a INT field
type TradePublishIndicator struct{ message.IntValue }

//Tag returns tag.TradePublishIndicator (1390)
func (f TradePublishIndicator) Tag() fix.Tag { return tag.TradePublishIndicator }

//BuildTradePublishIndicator returns a new TradePublishIndicator initialized with val
func BuildTradePublishIndicator(val int) TradePublishIndicator {
	var field TradePublishIndicator
	field.Value = val
	return field
}

//TradeReportID is a STRING field
type TradeReportID struct{ message.StringValue }

//Tag returns tag.TradeReportID (571)
func (f TradeReportID) Tag() fix.Tag { return tag.TradeReportID }

//BuildTradeReportID returns a new TradeReportID initialized with val
func BuildTradeReportID(val string) TradeReportID {
	var field TradeReportID
	field.Value = val
	return field
}

//TradeReportRefID is a STRING field
type TradeReportRefID struct{ message.StringValue }

//Tag returns tag.TradeReportRefID (572)
func (f TradeReportRefID) Tag() fix.Tag { return tag.TradeReportRefID }

//BuildTradeReportRefID returns a new TradeReportRefID initialized with val
func BuildTradeReportRefID(val string) TradeReportRefID {
	var field TradeReportRefID
	field.Value = val
	return field
}

//TradeReportRejectReason is a INT field
type TradeReportRejectReason struct{ message.IntValue }

//Tag returns tag.TradeReportRejectReason (751)
func (f TradeReportRejectReason) Tag() fix.Tag { return tag.TradeReportRejectReason }

//BuildTradeReportRejectReason returns a new TradeReportRejectReason initialized with val
func BuildTradeReportRejectReason(val int) TradeReportRejectReason {
	var field TradeReportRejectReason
	field.Value = val
	return field
}

//TradeReportTransType is a INT field
type TradeReportTransType struct{ message.IntValue }

//Tag returns tag.TradeReportTransType (487)
func (f TradeReportTransType) Tag() fix.Tag { return tag.TradeReportTransType }

//BuildTradeReportTransType returns a new TradeReportTransType initialized with val
func BuildTradeReportTransType(val int) TradeReportTransType {
	var field TradeReportTransType
	field.Value = val
	return field
}

//TradeReportType is a INT field
type TradeReportType struct{ message.IntValue }

//Tag returns tag.TradeReportType (856)
func (f TradeReportType) Tag() fix.Tag { return tag.TradeReportType }

//BuildTradeReportType returns a new TradeReportType initialized with val
func BuildTradeReportType(val int) TradeReportType {
	var field TradeReportType
	field.Value = val
	return field
}

//TradeRequestID is a STRING field
type TradeRequestID struct{ message.StringValue }

//Tag returns tag.TradeRequestID (568)
func (f TradeRequestID) Tag() fix.Tag { return tag.TradeRequestID }

//BuildTradeRequestID returns a new TradeRequestID initialized with val
func BuildTradeRequestID(val string) TradeRequestID {
	var field TradeRequestID
	field.Value = val
	return field
}

//TradeRequestResult is a INT field
type TradeRequestResult struct{ message.IntValue }

//Tag returns tag.TradeRequestResult (749)
func (f TradeRequestResult) Tag() fix.Tag { return tag.TradeRequestResult }

//BuildTradeRequestResult returns a new TradeRequestResult initialized with val
func BuildTradeRequestResult(val int) TradeRequestResult {
	var field TradeRequestResult
	field.Value = val
	return field
}

//TradeRequestStatus is a INT field
type TradeRequestStatus struct{ message.IntValue }

//Tag returns tag.TradeRequestStatus (750)
func (f TradeRequestStatus) Tag() fix.Tag { return tag.TradeRequestStatus }

//BuildTradeRequestStatus returns a new TradeRequestStatus initialized with val
func BuildTradeRequestStatus(val int) TradeRequestStatus {
	var field TradeRequestStatus
	field.Value = val
	return field
}

//TradeRequestType is a INT field
type TradeRequestType struct{ message.IntValue }

//Tag returns tag.TradeRequestType (569)
func (f TradeRequestType) Tag() fix.Tag { return tag.TradeRequestType }

//BuildTradeRequestType returns a new TradeRequestType initialized with val
func BuildTradeRequestType(val int) TradeRequestType {
	var field TradeRequestType
	field.Value = val
	return field
}

//TradeType is a CHAR field
type TradeType struct{ message.CharValue }

//Tag returns tag.TradeType (418)
func (f TradeType) Tag() fix.Tag { return tag.TradeType }

//BuildTradeType returns a new TradeType initialized with val
func BuildTradeType(val string) TradeType {
	var field TradeType
	field.Value = val
	return field
}

//TradeVolume is a QTY field
type TradeVolume struct{ message.QtyValue }

//Tag returns tag.TradeVolume (1020)
func (f TradeVolume) Tag() fix.Tag { return tag.TradeVolume }

//BuildTradeVolume returns a new TradeVolume initialized with val
func BuildTradeVolume(val float64) TradeVolume {
	var field TradeVolume
	field.Value = val
	return field
}

//TradedFlatSwitch is a BOOLEAN field
type TradedFlatSwitch struct{ message.BooleanValue }

//Tag returns tag.TradedFlatSwitch (258)
func (f TradedFlatSwitch) Tag() fix.Tag { return tag.TradedFlatSwitch }

//BuildTradedFlatSwitch returns a new TradedFlatSwitch initialized with val
func BuildTradedFlatSwitch(val bool) TradedFlatSwitch {
	var field TradedFlatSwitch
	field.Value = val
	return field
}

//TradingCurrency is a CURRENCY field
type TradingCurrency struct{ message.CurrencyValue }

//Tag returns tag.TradingCurrency (1245)
func (f TradingCurrency) Tag() fix.Tag { return tag.TradingCurrency }

//BuildTradingCurrency returns a new TradingCurrency initialized with val
func BuildTradingCurrency(val string) TradingCurrency {
	var field TradingCurrency
	field.Value = val
	return field
}

//TradingReferencePrice is a PRICE field
type TradingReferencePrice struct{ message.PriceValue }

//Tag returns tag.TradingReferencePrice (1150)
func (f TradingReferencePrice) Tag() fix.Tag { return tag.TradingReferencePrice }

//BuildTradingReferencePrice returns a new TradingReferencePrice initialized with val
func BuildTradingReferencePrice(val float64) TradingReferencePrice {
	var field TradingReferencePrice
	field.Value = val
	return field
}

//TradingSessionDesc is a STRING field
type TradingSessionDesc struct{ message.StringValue }

//Tag returns tag.TradingSessionDesc (1326)
func (f TradingSessionDesc) Tag() fix.Tag { return tag.TradingSessionDesc }

//BuildTradingSessionDesc returns a new TradingSessionDesc initialized with val
func BuildTradingSessionDesc(val string) TradingSessionDesc {
	var field TradingSessionDesc
	field.Value = val
	return field
}

//TradingSessionID is a STRING field
type TradingSessionID struct{ message.StringValue }

//Tag returns tag.TradingSessionID (336)
func (f TradingSessionID) Tag() fix.Tag { return tag.TradingSessionID }

//BuildTradingSessionID returns a new TradingSessionID initialized with val
func BuildTradingSessionID(val string) TradingSessionID {
	var field TradingSessionID
	field.Value = val
	return field
}

//TradingSessionSubID is a STRING field
type TradingSessionSubID struct{ message.StringValue }

//Tag returns tag.TradingSessionSubID (625)
func (f TradingSessionSubID) Tag() fix.Tag { return tag.TradingSessionSubID }

//BuildTradingSessionSubID returns a new TradingSessionSubID initialized with val
func BuildTradingSessionSubID(val string) TradingSessionSubID {
	var field TradingSessionSubID
	field.Value = val
	return field
}

//TransBkdTime is a UTCTIMESTAMP field
type TransBkdTime struct{ message.UTCTimestampValue }

//Tag returns tag.TransBkdTime (483)
func (f TransBkdTime) Tag() fix.Tag { return tag.TransBkdTime }

//TransactTime is a UTCTIMESTAMP field
type TransactTime struct{ message.UTCTimestampValue }

//Tag returns tag.TransactTime (60)
func (f TransactTime) Tag() fix.Tag { return tag.TransactTime }

//TransferReason is a STRING field
type TransferReason struct{ message.StringValue }

//Tag returns tag.TransferReason (830)
func (f TransferReason) Tag() fix.Tag { return tag.TransferReason }

//BuildTransferReason returns a new TransferReason initialized with val
func BuildTransferReason(val string) TransferReason {
	var field TransferReason
	field.Value = val
	return field
}

//TrdMatchID is a STRING field
type TrdMatchID struct{ message.StringValue }

//Tag returns tag.TrdMatchID (880)
func (f TrdMatchID) Tag() fix.Tag { return tag.TrdMatchID }

//BuildTrdMatchID returns a new TrdMatchID initialized with val
func BuildTrdMatchID(val string) TrdMatchID {
	var field TrdMatchID
	field.Value = val
	return field
}

//TrdRegTimestamp is a UTCTIMESTAMP field
type TrdRegTimestamp struct{ message.UTCTimestampValue }

//Tag returns tag.TrdRegTimestamp (769)
func (f TrdRegTimestamp) Tag() fix.Tag { return tag.TrdRegTimestamp }

//TrdRegTimestampOrigin is a STRING field
type TrdRegTimestampOrigin struct{ message.StringValue }

//Tag returns tag.TrdRegTimestampOrigin (771)
func (f TrdRegTimestampOrigin) Tag() fix.Tag { return tag.TrdRegTimestampOrigin }

//BuildTrdRegTimestampOrigin returns a new TrdRegTimestampOrigin initialized with val
func BuildTrdRegTimestampOrigin(val string) TrdRegTimestampOrigin {
	var field TrdRegTimestampOrigin
	field.Value = val
	return field
}

//TrdRegTimestampType is a INT field
type TrdRegTimestampType struct{ message.IntValue }

//Tag returns tag.TrdRegTimestampType (770)
func (f TrdRegTimestampType) Tag() fix.Tag { return tag.TrdRegTimestampType }

//BuildTrdRegTimestampType returns a new TrdRegTimestampType initialized with val
func BuildTrdRegTimestampType(val int) TrdRegTimestampType {
	var field TrdRegTimestampType
	field.Value = val
	return field
}

//TrdRepIndicator is a BOOLEAN field
type TrdRepIndicator struct{ message.BooleanValue }

//Tag returns tag.TrdRepIndicator (1389)
func (f TrdRepIndicator) Tag() fix.Tag { return tag.TrdRepIndicator }

//BuildTrdRepIndicator returns a new TrdRepIndicator initialized with val
func BuildTrdRepIndicator(val bool) TrdRepIndicator {
	var field TrdRepIndicator
	field.Value = val
	return field
}

//TrdRepPartyRole is a INT field
type TrdRepPartyRole struct{ message.IntValue }

//Tag returns tag.TrdRepPartyRole (1388)
func (f TrdRepPartyRole) Tag() fix.Tag { return tag.TrdRepPartyRole }

//BuildTrdRepPartyRole returns a new TrdRepPartyRole initialized with val
func BuildTrdRepPartyRole(val int) TrdRepPartyRole {
	var field TrdRepPartyRole
	field.Value = val
	return field
}

//TrdRptStatus is a INT field
type TrdRptStatus struct{ message.IntValue }

//Tag returns tag.TrdRptStatus (939)
func (f TrdRptStatus) Tag() fix.Tag { return tag.TrdRptStatus }

//BuildTrdRptStatus returns a new TrdRptStatus initialized with val
func BuildTrdRptStatus(val int) TrdRptStatus {
	var field TrdRptStatus
	field.Value = val
	return field
}

//TrdSubType is a INT field
type TrdSubType struct{ message.IntValue }

//Tag returns tag.TrdSubType (829)
func (f TrdSubType) Tag() fix.Tag { return tag.TrdSubType }

//BuildTrdSubType returns a new TrdSubType initialized with val
func BuildTrdSubType(val int) TrdSubType {
	var field TrdSubType
	field.Value = val
	return field
}

//TrdType is a INT field
type TrdType struct{ message.IntValue }

//Tag returns tag.TrdType (828)
func (f TrdType) Tag() fix.Tag { return tag.TrdType }

//BuildTrdType returns a new TrdType initialized with val
func BuildTrdType(val int) TrdType {
	var field TrdType
	field.Value = val
	return field
}

//TriggerAction is a CHAR field
type TriggerAction struct{ message.CharValue }

//Tag returns tag.TriggerAction (1101)
func (f TriggerAction) Tag() fix.Tag { return tag.TriggerAction }

//BuildTriggerAction returns a new TriggerAction initialized with val
func BuildTriggerAction(val string) TriggerAction {
	var field TriggerAction
	field.Value = val
	return field
}

//TriggerNewPrice is a PRICE field
type TriggerNewPrice struct{ message.PriceValue }

//Tag returns tag.TriggerNewPrice (1110)
func (f TriggerNewPrice) Tag() fix.Tag { return tag.TriggerNewPrice }

//BuildTriggerNewPrice returns a new TriggerNewPrice initialized with val
func BuildTriggerNewPrice(val float64) TriggerNewPrice {
	var field TriggerNewPrice
	field.Value = val
	return field
}

//TriggerNewQty is a QTY field
type TriggerNewQty struct{ message.QtyValue }

//Tag returns tag.TriggerNewQty (1112)
func (f TriggerNewQty) Tag() fix.Tag { return tag.TriggerNewQty }

//BuildTriggerNewQty returns a new TriggerNewQty initialized with val
func BuildTriggerNewQty(val float64) TriggerNewQty {
	var field TriggerNewQty
	field.Value = val
	return field
}

//TriggerOrderType is a CHAR field
type TriggerOrderType struct{ message.CharValue }

//Tag returns tag.TriggerOrderType (1111)
func (f TriggerOrderType) Tag() fix.Tag { return tag.TriggerOrderType }

//BuildTriggerOrderType returns a new TriggerOrderType initialized with val
func BuildTriggerOrderType(val string) TriggerOrderType {
	var field TriggerOrderType
	field.Value = val
	return field
}

//TriggerPrice is a PRICE field
type TriggerPrice struct{ message.PriceValue }

//Tag returns tag.TriggerPrice (1102)
func (f TriggerPrice) Tag() fix.Tag { return tag.TriggerPrice }

//BuildTriggerPrice returns a new TriggerPrice initialized with val
func BuildTriggerPrice(val float64) TriggerPrice {
	var field TriggerPrice
	field.Value = val
	return field
}

//TriggerPriceDirection is a CHAR field
type TriggerPriceDirection struct{ message.CharValue }

//Tag returns tag.TriggerPriceDirection (1109)
func (f TriggerPriceDirection) Tag() fix.Tag { return tag.TriggerPriceDirection }

//BuildTriggerPriceDirection returns a new TriggerPriceDirection initialized with val
func BuildTriggerPriceDirection(val string) TriggerPriceDirection {
	var field TriggerPriceDirection
	field.Value = val
	return field
}

//TriggerPriceType is a CHAR field
type TriggerPriceType struct{ message.CharValue }

//Tag returns tag.TriggerPriceType (1107)
func (f TriggerPriceType) Tag() fix.Tag { return tag.TriggerPriceType }

//BuildTriggerPriceType returns a new TriggerPriceType initialized with val
func BuildTriggerPriceType(val string) TriggerPriceType {
	var field TriggerPriceType
	field.Value = val
	return field
}

//TriggerPriceTypeScope is a CHAR field
type TriggerPriceTypeScope struct{ message.CharValue }

//Tag returns tag.TriggerPriceTypeScope (1108)
func (f TriggerPriceTypeScope) Tag() fix.Tag { return tag.TriggerPriceTypeScope }

//BuildTriggerPriceTypeScope returns a new TriggerPriceTypeScope initialized with val
func BuildTriggerPriceTypeScope(val string) TriggerPriceTypeScope {
	var field TriggerPriceTypeScope
	field.Value = val
	return field
}

//TriggerSecurityDesc is a STRING field
type TriggerSecurityDesc struct{ message.StringValue }

//Tag returns tag.TriggerSecurityDesc (1106)
func (f TriggerSecurityDesc) Tag() fix.Tag { return tag.TriggerSecurityDesc }

//BuildTriggerSecurityDesc returns a new TriggerSecurityDesc initialized with val
func BuildTriggerSecurityDesc(val string) TriggerSecurityDesc {
	var field TriggerSecurityDesc
	field.Value = val
	return field
}

//TriggerSecurityID is a STRING field
type TriggerSecurityID struct{ message.StringValue }

//Tag returns tag.TriggerSecurityID (1104)
func (f TriggerSecurityID) Tag() fix.Tag { return tag.TriggerSecurityID }

//BuildTriggerSecurityID returns a new TriggerSecurityID initialized with val
func BuildTriggerSecurityID(val string) TriggerSecurityID {
	var field TriggerSecurityID
	field.Value = val
	return field
}

//TriggerSecurityIDSource is a STRING field
type TriggerSecurityIDSource struct{ message.StringValue }

//Tag returns tag.TriggerSecurityIDSource (1105)
func (f TriggerSecurityIDSource) Tag() fix.Tag { return tag.TriggerSecurityIDSource }

//BuildTriggerSecurityIDSource returns a new TriggerSecurityIDSource initialized with val
func BuildTriggerSecurityIDSource(val string) TriggerSecurityIDSource {
	var field TriggerSecurityIDSource
	field.Value = val
	return field
}

//TriggerSymbol is a STRING field
type TriggerSymbol struct{ message.StringValue }

//Tag returns tag.TriggerSymbol (1103)
func (f TriggerSymbol) Tag() fix.Tag { return tag.TriggerSymbol }

//BuildTriggerSymbol returns a new TriggerSymbol initialized with val
func BuildTriggerSymbol(val string) TriggerSymbol {
	var field TriggerSymbol
	field.Value = val
	return field
}

//TriggerTradingSessionID is a STRING field
type TriggerTradingSessionID struct{ message.StringValue }

//Tag returns tag.TriggerTradingSessionID (1113)
func (f TriggerTradingSessionID) Tag() fix.Tag { return tag.TriggerTradingSessionID }

//BuildTriggerTradingSessionID returns a new TriggerTradingSessionID initialized with val
func BuildTriggerTradingSessionID(val string) TriggerTradingSessionID {
	var field TriggerTradingSessionID
	field.Value = val
	return field
}

//TriggerTradingSessionSubID is a STRING field
type TriggerTradingSessionSubID struct{ message.StringValue }

//Tag returns tag.TriggerTradingSessionSubID (1114)
func (f TriggerTradingSessionSubID) Tag() fix.Tag { return tag.TriggerTradingSessionSubID }

//BuildTriggerTradingSessionSubID returns a new TriggerTradingSessionSubID initialized with val
func BuildTriggerTradingSessionSubID(val string) TriggerTradingSessionSubID {
	var field TriggerTradingSessionSubID
	field.Value = val
	return field
}

//TriggerType is a CHAR field
type TriggerType struct{ message.CharValue }

//Tag returns tag.TriggerType (1100)
func (f TriggerType) Tag() fix.Tag { return tag.TriggerType }

//BuildTriggerType returns a new TriggerType initialized with val
func BuildTriggerType(val string) TriggerType {
	var field TriggerType
	field.Value = val
	return field
}

//URLLink is a STRING field
type URLLink struct{ message.StringValue }

//Tag returns tag.URLLink (149)
func (f URLLink) Tag() fix.Tag { return tag.URLLink }

//BuildURLLink returns a new URLLink initialized with val
func BuildURLLink(val string) URLLink {
	var field URLLink
	field.Value = val
	return field
}

//UnderlyingAdjustedQuantity is a QTY field
type UnderlyingAdjustedQuantity struct{ message.QtyValue }

//Tag returns tag.UnderlyingAdjustedQuantity (1044)
func (f UnderlyingAdjustedQuantity) Tag() fix.Tag { return tag.UnderlyingAdjustedQuantity }

//BuildUnderlyingAdjustedQuantity returns a new UnderlyingAdjustedQuantity initialized with val
func BuildUnderlyingAdjustedQuantity(val float64) UnderlyingAdjustedQuantity {
	var field UnderlyingAdjustedQuantity
	field.Value = val
	return field
}

//UnderlyingAllocationPercent is a PERCENTAGE field
type UnderlyingAllocationPercent struct{ message.PercentageValue }

//Tag returns tag.UnderlyingAllocationPercent (972)
func (f UnderlyingAllocationPercent) Tag() fix.Tag { return tag.UnderlyingAllocationPercent }

//BuildUnderlyingAllocationPercent returns a new UnderlyingAllocationPercent initialized with val
func BuildUnderlyingAllocationPercent(val float64) UnderlyingAllocationPercent {
	var field UnderlyingAllocationPercent
	field.Value = val
	return field
}

//UnderlyingAttachmentPoint is a PERCENTAGE field
type UnderlyingAttachmentPoint struct{ message.PercentageValue }

//Tag returns tag.UnderlyingAttachmentPoint (1459)
func (f UnderlyingAttachmentPoint) Tag() fix.Tag { return tag.UnderlyingAttachmentPoint }

//BuildUnderlyingAttachmentPoint returns a new UnderlyingAttachmentPoint initialized with val
func BuildUnderlyingAttachmentPoint(val float64) UnderlyingAttachmentPoint {
	var field UnderlyingAttachmentPoint
	field.Value = val
	return field
}

//UnderlyingCFICode is a STRING field
type UnderlyingCFICode struct{ message.StringValue }

//Tag returns tag.UnderlyingCFICode (463)
func (f UnderlyingCFICode) Tag() fix.Tag { return tag.UnderlyingCFICode }

//BuildUnderlyingCFICode returns a new UnderlyingCFICode initialized with val
func BuildUnderlyingCFICode(val string) UnderlyingCFICode {
	var field UnderlyingCFICode
	field.Value = val
	return field
}

//UnderlyingCPProgram is a STRING field
type UnderlyingCPProgram struct{ message.StringValue }

//Tag returns tag.UnderlyingCPProgram (877)
func (f UnderlyingCPProgram) Tag() fix.Tag { return tag.UnderlyingCPProgram }

//BuildUnderlyingCPProgram returns a new UnderlyingCPProgram initialized with val
func BuildUnderlyingCPProgram(val string) UnderlyingCPProgram {
	var field UnderlyingCPProgram
	field.Value = val
	return field
}

//UnderlyingCPRegType is a STRING field
type UnderlyingCPRegType struct{ message.StringValue }

//Tag returns tag.UnderlyingCPRegType (878)
func (f UnderlyingCPRegType) Tag() fix.Tag { return tag.UnderlyingCPRegType }

//BuildUnderlyingCPRegType returns a new UnderlyingCPRegType initialized with val
func BuildUnderlyingCPRegType(val string) UnderlyingCPRegType {
	var field UnderlyingCPRegType
	field.Value = val
	return field
}

//UnderlyingCapValue is a AMT field
type UnderlyingCapValue struct{ message.AmtValue }

//Tag returns tag.UnderlyingCapValue (1038)
func (f UnderlyingCapValue) Tag() fix.Tag { return tag.UnderlyingCapValue }

//BuildUnderlyingCapValue returns a new UnderlyingCapValue initialized with val
func BuildUnderlyingCapValue(val float64) UnderlyingCapValue {
	var field UnderlyingCapValue
	field.Value = val
	return field
}

//UnderlyingCashAmount is a AMT field
type UnderlyingCashAmount struct{ message.AmtValue }

//Tag returns tag.UnderlyingCashAmount (973)
func (f UnderlyingCashAmount) Tag() fix.Tag { return tag.UnderlyingCashAmount }

//BuildUnderlyingCashAmount returns a new UnderlyingCashAmount initialized with val
func BuildUnderlyingCashAmount(val float64) UnderlyingCashAmount {
	var field UnderlyingCashAmount
	field.Value = val
	return field
}

//UnderlyingCashType is a STRING field
type UnderlyingCashType struct{ message.StringValue }

//Tag returns tag.UnderlyingCashType (974)
func (f UnderlyingCashType) Tag() fix.Tag { return tag.UnderlyingCashType }

//BuildUnderlyingCashType returns a new UnderlyingCashType initialized with val
func BuildUnderlyingCashType(val string) UnderlyingCashType {
	var field UnderlyingCashType
	field.Value = val
	return field
}

//UnderlyingCollectAmount is a AMT field
type UnderlyingCollectAmount struct{ message.AmtValue }

//Tag returns tag.UnderlyingCollectAmount (986)
func (f UnderlyingCollectAmount) Tag() fix.Tag { return tag.UnderlyingCollectAmount }

//BuildUnderlyingCollectAmount returns a new UnderlyingCollectAmount initialized with val
func BuildUnderlyingCollectAmount(val float64) UnderlyingCollectAmount {
	var field UnderlyingCollectAmount
	field.Value = val
	return field
}

//UnderlyingContractMultiplier is a FLOAT field
type UnderlyingContractMultiplier struct{ message.FloatValue }

//Tag returns tag.UnderlyingContractMultiplier (436)
func (f UnderlyingContractMultiplier) Tag() fix.Tag { return tag.UnderlyingContractMultiplier }

//BuildUnderlyingContractMultiplier returns a new UnderlyingContractMultiplier initialized with val
func BuildUnderlyingContractMultiplier(val float64) UnderlyingContractMultiplier {
	var field UnderlyingContractMultiplier
	field.Value = val
	return field
}

//UnderlyingContractMultiplierUnit is a INT field
type UnderlyingContractMultiplierUnit struct{ message.IntValue }

//Tag returns tag.UnderlyingContractMultiplierUnit (1437)
func (f UnderlyingContractMultiplierUnit) Tag() fix.Tag { return tag.UnderlyingContractMultiplierUnit }

//BuildUnderlyingContractMultiplierUnit returns a new UnderlyingContractMultiplierUnit initialized with val
func BuildUnderlyingContractMultiplierUnit(val int) UnderlyingContractMultiplierUnit {
	var field UnderlyingContractMultiplierUnit
	field.Value = val
	return field
}

//UnderlyingCountryOfIssue is a COUNTRY field
type UnderlyingCountryOfIssue struct{ message.CountryValue }

//Tag returns tag.UnderlyingCountryOfIssue (592)
func (f UnderlyingCountryOfIssue) Tag() fix.Tag { return tag.UnderlyingCountryOfIssue }

//BuildUnderlyingCountryOfIssue returns a new UnderlyingCountryOfIssue initialized with val
func BuildUnderlyingCountryOfIssue(val string) UnderlyingCountryOfIssue {
	var field UnderlyingCountryOfIssue
	field.Value = val
	return field
}

//UnderlyingCouponPaymentDate is a LOCALMKTDATE field
type UnderlyingCouponPaymentDate struct{ message.LocalMktDateValue }

//Tag returns tag.UnderlyingCouponPaymentDate (241)
func (f UnderlyingCouponPaymentDate) Tag() fix.Tag { return tag.UnderlyingCouponPaymentDate }

//BuildUnderlyingCouponPaymentDate returns a new UnderlyingCouponPaymentDate initialized with val
func BuildUnderlyingCouponPaymentDate(val string) UnderlyingCouponPaymentDate {
	var field UnderlyingCouponPaymentDate
	field.Value = val
	return field
}

//UnderlyingCouponRate is a PERCENTAGE field
type UnderlyingCouponRate struct{ message.PercentageValue }

//Tag returns tag.UnderlyingCouponRate (435)
func (f UnderlyingCouponRate) Tag() fix.Tag { return tag.UnderlyingCouponRate }

//BuildUnderlyingCouponRate returns a new UnderlyingCouponRate initialized with val
func BuildUnderlyingCouponRate(val float64) UnderlyingCouponRate {
	var field UnderlyingCouponRate
	field.Value = val
	return field
}

//UnderlyingCreditRating is a STRING field
type UnderlyingCreditRating struct{ message.StringValue }

//Tag returns tag.UnderlyingCreditRating (256)
func (f UnderlyingCreditRating) Tag() fix.Tag { return tag.UnderlyingCreditRating }

//BuildUnderlyingCreditRating returns a new UnderlyingCreditRating initialized with val
func BuildUnderlyingCreditRating(val string) UnderlyingCreditRating {
	var field UnderlyingCreditRating
	field.Value = val
	return field
}

//UnderlyingCurrency is a CURRENCY field
type UnderlyingCurrency struct{ message.CurrencyValue }

//Tag returns tag.UnderlyingCurrency (318)
func (f UnderlyingCurrency) Tag() fix.Tag { return tag.UnderlyingCurrency }

//BuildUnderlyingCurrency returns a new UnderlyingCurrency initialized with val
func BuildUnderlyingCurrency(val string) UnderlyingCurrency {
	var field UnderlyingCurrency
	field.Value = val
	return field
}

//UnderlyingCurrentValue is a AMT field
type UnderlyingCurrentValue struct{ message.AmtValue }

//Tag returns tag.UnderlyingCurrentValue (885)
func (f UnderlyingCurrentValue) Tag() fix.Tag { return tag.UnderlyingCurrentValue }

//BuildUnderlyingCurrentValue returns a new UnderlyingCurrentValue initialized with val
func BuildUnderlyingCurrentValue(val float64) UnderlyingCurrentValue {
	var field UnderlyingCurrentValue
	field.Value = val
	return field
}

//UnderlyingDeliveryAmount is a AMT field
type UnderlyingDeliveryAmount struct{ message.AmtValue }

//Tag returns tag.UnderlyingDeliveryAmount (1037)
func (f UnderlyingDeliveryAmount) Tag() fix.Tag { return tag.UnderlyingDeliveryAmount }

//BuildUnderlyingDeliveryAmount returns a new UnderlyingDeliveryAmount initialized with val
func BuildUnderlyingDeliveryAmount(val float64) UnderlyingDeliveryAmount {
	var field UnderlyingDeliveryAmount
	field.Value = val
	return field
}

//UnderlyingDetachmentPoint is a PERCENTAGE field
type UnderlyingDetachmentPoint struct{ message.PercentageValue }

//Tag returns tag.UnderlyingDetachmentPoint (1460)
func (f UnderlyingDetachmentPoint) Tag() fix.Tag { return tag.UnderlyingDetachmentPoint }

//BuildUnderlyingDetachmentPoint returns a new UnderlyingDetachmentPoint initialized with val
func BuildUnderlyingDetachmentPoint(val float64) UnderlyingDetachmentPoint {
	var field UnderlyingDetachmentPoint
	field.Value = val
	return field
}

//UnderlyingDirtyPrice is a PRICE field
type UnderlyingDirtyPrice struct{ message.PriceValue }

//Tag returns tag.UnderlyingDirtyPrice (882)
func (f UnderlyingDirtyPrice) Tag() fix.Tag { return tag.UnderlyingDirtyPrice }

//BuildUnderlyingDirtyPrice returns a new UnderlyingDirtyPrice initialized with val
func BuildUnderlyingDirtyPrice(val float64) UnderlyingDirtyPrice {
	var field UnderlyingDirtyPrice
	field.Value = val
	return field
}

//UnderlyingEndPrice is a PRICE field
type UnderlyingEndPrice struct{ message.PriceValue }

//Tag returns tag.UnderlyingEndPrice (883)
func (f UnderlyingEndPrice) Tag() fix.Tag { return tag.UnderlyingEndPrice }

//BuildUnderlyingEndPrice returns a new UnderlyingEndPrice initialized with val
func BuildUnderlyingEndPrice(val float64) UnderlyingEndPrice {
	var field UnderlyingEndPrice
	field.Value = val
	return field
}

//UnderlyingEndValue is a AMT field
type UnderlyingEndValue struct{ message.AmtValue }

//Tag returns tag.UnderlyingEndValue (886)
func (f UnderlyingEndValue) Tag() fix.Tag { return tag.UnderlyingEndValue }

//BuildUnderlyingEndValue returns a new UnderlyingEndValue initialized with val
func BuildUnderlyingEndValue(val float64) UnderlyingEndValue {
	var field UnderlyingEndValue
	field.Value = val
	return field
}

//UnderlyingExerciseStyle is a INT field
type UnderlyingExerciseStyle struct{ message.IntValue }

//Tag returns tag.UnderlyingExerciseStyle (1419)
func (f UnderlyingExerciseStyle) Tag() fix.Tag { return tag.UnderlyingExerciseStyle }

//BuildUnderlyingExerciseStyle returns a new UnderlyingExerciseStyle initialized with val
func BuildUnderlyingExerciseStyle(val int) UnderlyingExerciseStyle {
	var field UnderlyingExerciseStyle
	field.Value = val
	return field
}

//UnderlyingFXRate is a FLOAT field
type UnderlyingFXRate struct{ message.FloatValue }

//Tag returns tag.UnderlyingFXRate (1045)
func (f UnderlyingFXRate) Tag() fix.Tag { return tag.UnderlyingFXRate }

//BuildUnderlyingFXRate returns a new UnderlyingFXRate initialized with val
func BuildUnderlyingFXRate(val float64) UnderlyingFXRate {
	var field UnderlyingFXRate
	field.Value = val
	return field
}

//UnderlyingFXRateCalc is a CHAR field
type UnderlyingFXRateCalc struct{ message.CharValue }

//Tag returns tag.UnderlyingFXRateCalc (1046)
func (f UnderlyingFXRateCalc) Tag() fix.Tag { return tag.UnderlyingFXRateCalc }

//BuildUnderlyingFXRateCalc returns a new UnderlyingFXRateCalc initialized with val
func BuildUnderlyingFXRateCalc(val string) UnderlyingFXRateCalc {
	var field UnderlyingFXRateCalc
	field.Value = val
	return field
}

//UnderlyingFactor is a FLOAT field
type UnderlyingFactor struct{ message.FloatValue }

//Tag returns tag.UnderlyingFactor (246)
func (f UnderlyingFactor) Tag() fix.Tag { return tag.UnderlyingFactor }

//BuildUnderlyingFactor returns a new UnderlyingFactor initialized with val
func BuildUnderlyingFactor(val float64) UnderlyingFactor {
	var field UnderlyingFactor
	field.Value = val
	return field
}

//UnderlyingFlowScheduleType is a INT field
type UnderlyingFlowScheduleType struct{ message.IntValue }

//Tag returns tag.UnderlyingFlowScheduleType (1441)
func (f UnderlyingFlowScheduleType) Tag() fix.Tag { return tag.UnderlyingFlowScheduleType }

//BuildUnderlyingFlowScheduleType returns a new UnderlyingFlowScheduleType initialized with val
func BuildUnderlyingFlowScheduleType(val int) UnderlyingFlowScheduleType {
	var field UnderlyingFlowScheduleType
	field.Value = val
	return field
}

//UnderlyingIDSource is a STRING field
type UnderlyingIDSource struct{ message.StringValue }

//Tag returns tag.UnderlyingIDSource (305)
func (f UnderlyingIDSource) Tag() fix.Tag { return tag.UnderlyingIDSource }

//BuildUnderlyingIDSource returns a new UnderlyingIDSource initialized with val
func BuildUnderlyingIDSource(val string) UnderlyingIDSource {
	var field UnderlyingIDSource
	field.Value = val
	return field
}

//UnderlyingInstrRegistry is a STRING field
type UnderlyingInstrRegistry struct{ message.StringValue }

//Tag returns tag.UnderlyingInstrRegistry (595)
func (f UnderlyingInstrRegistry) Tag() fix.Tag { return tag.UnderlyingInstrRegistry }

//BuildUnderlyingInstrRegistry returns a new UnderlyingInstrRegistry initialized with val
func BuildUnderlyingInstrRegistry(val string) UnderlyingInstrRegistry {
	var field UnderlyingInstrRegistry
	field.Value = val
	return field
}

//UnderlyingInstrumentPartyID is a STRING field
type UnderlyingInstrumentPartyID struct{ message.StringValue }

//Tag returns tag.UnderlyingInstrumentPartyID (1059)
func (f UnderlyingInstrumentPartyID) Tag() fix.Tag { return tag.UnderlyingInstrumentPartyID }

//BuildUnderlyingInstrumentPartyID returns a new UnderlyingInstrumentPartyID initialized with val
func BuildUnderlyingInstrumentPartyID(val string) UnderlyingInstrumentPartyID {
	var field UnderlyingInstrumentPartyID
	field.Value = val
	return field
}

//UnderlyingInstrumentPartyIDSource is a CHAR field
type UnderlyingInstrumentPartyIDSource struct{ message.CharValue }

//Tag returns tag.UnderlyingInstrumentPartyIDSource (1060)
func (f UnderlyingInstrumentPartyIDSource) Tag() fix.Tag { return tag.UnderlyingInstrumentPartyIDSource }

//BuildUnderlyingInstrumentPartyIDSource returns a new UnderlyingInstrumentPartyIDSource initialized with val
func BuildUnderlyingInstrumentPartyIDSource(val string) UnderlyingInstrumentPartyIDSource {
	var field UnderlyingInstrumentPartyIDSource
	field.Value = val
	return field
}

//UnderlyingInstrumentPartyRole is a INT field
type UnderlyingInstrumentPartyRole struct{ message.IntValue }

//Tag returns tag.UnderlyingInstrumentPartyRole (1061)
func (f UnderlyingInstrumentPartyRole) Tag() fix.Tag { return tag.UnderlyingInstrumentPartyRole }

//BuildUnderlyingInstrumentPartyRole returns a new UnderlyingInstrumentPartyRole initialized with val
func BuildUnderlyingInstrumentPartyRole(val int) UnderlyingInstrumentPartyRole {
	var field UnderlyingInstrumentPartyRole
	field.Value = val
	return field
}

//UnderlyingInstrumentPartySubID is a STRING field
type UnderlyingInstrumentPartySubID struct{ message.StringValue }

//Tag returns tag.UnderlyingInstrumentPartySubID (1063)
func (f UnderlyingInstrumentPartySubID) Tag() fix.Tag { return tag.UnderlyingInstrumentPartySubID }

//BuildUnderlyingInstrumentPartySubID returns a new UnderlyingInstrumentPartySubID initialized with val
func BuildUnderlyingInstrumentPartySubID(val string) UnderlyingInstrumentPartySubID {
	var field UnderlyingInstrumentPartySubID
	field.Value = val
	return field
}

//UnderlyingInstrumentPartySubIDType is a INT field
type UnderlyingInstrumentPartySubIDType struct{ message.IntValue }

//Tag returns tag.UnderlyingInstrumentPartySubIDType (1064)
func (f UnderlyingInstrumentPartySubIDType) Tag() fix.Tag {
	return tag.UnderlyingInstrumentPartySubIDType
}

//BuildUnderlyingInstrumentPartySubIDType returns a new UnderlyingInstrumentPartySubIDType initialized with val
func BuildUnderlyingInstrumentPartySubIDType(val int) UnderlyingInstrumentPartySubIDType {
	var field UnderlyingInstrumentPartySubIDType
	field.Value = val
	return field
}

//UnderlyingIssueDate is a LOCALMKTDATE field
type UnderlyingIssueDate struct{ message.LocalMktDateValue }

//Tag returns tag.UnderlyingIssueDate (242)
func (f UnderlyingIssueDate) Tag() fix.Tag { return tag.UnderlyingIssueDate }

//BuildUnderlyingIssueDate returns a new UnderlyingIssueDate initialized with val
func BuildUnderlyingIssueDate(val string) UnderlyingIssueDate {
	var field UnderlyingIssueDate
	field.Value = val
	return field
}

//UnderlyingIssuer is a STRING field
type UnderlyingIssuer struct{ message.StringValue }

//Tag returns tag.UnderlyingIssuer (306)
func (f UnderlyingIssuer) Tag() fix.Tag { return tag.UnderlyingIssuer }

//BuildUnderlyingIssuer returns a new UnderlyingIssuer initialized with val
func BuildUnderlyingIssuer(val string) UnderlyingIssuer {
	var field UnderlyingIssuer
	field.Value = val
	return field
}

//UnderlyingLastPx is a PRICE field
type UnderlyingLastPx struct{ message.PriceValue }

//Tag returns tag.UnderlyingLastPx (651)
func (f UnderlyingLastPx) Tag() fix.Tag { return tag.UnderlyingLastPx }

//BuildUnderlyingLastPx returns a new UnderlyingLastPx initialized with val
func BuildUnderlyingLastPx(val float64) UnderlyingLastPx {
	var field UnderlyingLastPx
	field.Value = val
	return field
}

//UnderlyingLastQty is a QTY field
type UnderlyingLastQty struct{ message.QtyValue }

//Tag returns tag.UnderlyingLastQty (652)
func (f UnderlyingLastQty) Tag() fix.Tag { return tag.UnderlyingLastQty }

//BuildUnderlyingLastQty returns a new UnderlyingLastQty initialized with val
func BuildUnderlyingLastQty(val float64) UnderlyingLastQty {
	var field UnderlyingLastQty
	field.Value = val
	return field
}

//UnderlyingLegCFICode is a STRING field
type UnderlyingLegCFICode struct{ message.StringValue }

//Tag returns tag.UnderlyingLegCFICode (1344)
func (f UnderlyingLegCFICode) Tag() fix.Tag { return tag.UnderlyingLegCFICode }

//BuildUnderlyingLegCFICode returns a new UnderlyingLegCFICode initialized with val
func BuildUnderlyingLegCFICode(val string) UnderlyingLegCFICode {
	var field UnderlyingLegCFICode
	field.Value = val
	return field
}

//UnderlyingLegMaturityDate is a LOCALMKTDATE field
type UnderlyingLegMaturityDate struct{ message.LocalMktDateValue }

//Tag returns tag.UnderlyingLegMaturityDate (1345)
func (f UnderlyingLegMaturityDate) Tag() fix.Tag { return tag.UnderlyingLegMaturityDate }

//BuildUnderlyingLegMaturityDate returns a new UnderlyingLegMaturityDate initialized with val
func BuildUnderlyingLegMaturityDate(val string) UnderlyingLegMaturityDate {
	var field UnderlyingLegMaturityDate
	field.Value = val
	return field
}

//UnderlyingLegMaturityMonthYear is a MONTHYEAR field
type UnderlyingLegMaturityMonthYear struct{ message.MonthYearValue }

//Tag returns tag.UnderlyingLegMaturityMonthYear (1339)
func (f UnderlyingLegMaturityMonthYear) Tag() fix.Tag { return tag.UnderlyingLegMaturityMonthYear }

//BuildUnderlyingLegMaturityMonthYear returns a new UnderlyingLegMaturityMonthYear initialized with val
func BuildUnderlyingLegMaturityMonthYear(val string) UnderlyingLegMaturityMonthYear {
	var field UnderlyingLegMaturityMonthYear
	field.Value = val
	return field
}

//UnderlyingLegMaturityTime is a TZTIMEONLY field
type UnderlyingLegMaturityTime struct{ message.TZTimeOnlyValue }

//Tag returns tag.UnderlyingLegMaturityTime (1405)
func (f UnderlyingLegMaturityTime) Tag() fix.Tag { return tag.UnderlyingLegMaturityTime }

//UnderlyingLegOptAttribute is a CHAR field
type UnderlyingLegOptAttribute struct{ message.CharValue }

//Tag returns tag.UnderlyingLegOptAttribute (1391)
func (f UnderlyingLegOptAttribute) Tag() fix.Tag { return tag.UnderlyingLegOptAttribute }

//BuildUnderlyingLegOptAttribute returns a new UnderlyingLegOptAttribute initialized with val
func BuildUnderlyingLegOptAttribute(val string) UnderlyingLegOptAttribute {
	var field UnderlyingLegOptAttribute
	field.Value = val
	return field
}

//UnderlyingLegPutOrCall is a INT field
type UnderlyingLegPutOrCall struct{ message.IntValue }

//Tag returns tag.UnderlyingLegPutOrCall (1343)
func (f UnderlyingLegPutOrCall) Tag() fix.Tag { return tag.UnderlyingLegPutOrCall }

//BuildUnderlyingLegPutOrCall returns a new UnderlyingLegPutOrCall initialized with val
func BuildUnderlyingLegPutOrCall(val int) UnderlyingLegPutOrCall {
	var field UnderlyingLegPutOrCall
	field.Value = val
	return field
}

//UnderlyingLegSecurityAltID is a STRING field
type UnderlyingLegSecurityAltID struct{ message.StringValue }

//Tag returns tag.UnderlyingLegSecurityAltID (1335)
func (f UnderlyingLegSecurityAltID) Tag() fix.Tag { return tag.UnderlyingLegSecurityAltID }

//BuildUnderlyingLegSecurityAltID returns a new UnderlyingLegSecurityAltID initialized with val
func BuildUnderlyingLegSecurityAltID(val string) UnderlyingLegSecurityAltID {
	var field UnderlyingLegSecurityAltID
	field.Value = val
	return field
}

//UnderlyingLegSecurityAltIDSource is a STRING field
type UnderlyingLegSecurityAltIDSource struct{ message.StringValue }

//Tag returns tag.UnderlyingLegSecurityAltIDSource (1336)
func (f UnderlyingLegSecurityAltIDSource) Tag() fix.Tag { return tag.UnderlyingLegSecurityAltIDSource }

//BuildUnderlyingLegSecurityAltIDSource returns a new UnderlyingLegSecurityAltIDSource initialized with val
func BuildUnderlyingLegSecurityAltIDSource(val string) UnderlyingLegSecurityAltIDSource {
	var field UnderlyingLegSecurityAltIDSource
	field.Value = val
	return field
}

//UnderlyingLegSecurityDesc is a STRING field
type UnderlyingLegSecurityDesc struct{ message.StringValue }

//Tag returns tag.UnderlyingLegSecurityDesc (1392)
func (f UnderlyingLegSecurityDesc) Tag() fix.Tag { return tag.UnderlyingLegSecurityDesc }

//BuildUnderlyingLegSecurityDesc returns a new UnderlyingLegSecurityDesc initialized with val
func BuildUnderlyingLegSecurityDesc(val string) UnderlyingLegSecurityDesc {
	var field UnderlyingLegSecurityDesc
	field.Value = val
	return field
}

//UnderlyingLegSecurityExchange is a STRING field
type UnderlyingLegSecurityExchange struct{ message.StringValue }

//Tag returns tag.UnderlyingLegSecurityExchange (1341)
func (f UnderlyingLegSecurityExchange) Tag() fix.Tag { return tag.UnderlyingLegSecurityExchange }

//BuildUnderlyingLegSecurityExchange returns a new UnderlyingLegSecurityExchange initialized with val
func BuildUnderlyingLegSecurityExchange(val string) UnderlyingLegSecurityExchange {
	var field UnderlyingLegSecurityExchange
	field.Value = val
	return field
}

//UnderlyingLegSecurityID is a STRING field
type UnderlyingLegSecurityID struct{ message.StringValue }

//Tag returns tag.UnderlyingLegSecurityID (1332)
func (f UnderlyingLegSecurityID) Tag() fix.Tag { return tag.UnderlyingLegSecurityID }

//BuildUnderlyingLegSecurityID returns a new UnderlyingLegSecurityID initialized with val
func BuildUnderlyingLegSecurityID(val string) UnderlyingLegSecurityID {
	var field UnderlyingLegSecurityID
	field.Value = val
	return field
}

//UnderlyingLegSecurityIDSource is a STRING field
type UnderlyingLegSecurityIDSource struct{ message.StringValue }

//Tag returns tag.UnderlyingLegSecurityIDSource (1333)
func (f UnderlyingLegSecurityIDSource) Tag() fix.Tag { return tag.UnderlyingLegSecurityIDSource }

//BuildUnderlyingLegSecurityIDSource returns a new UnderlyingLegSecurityIDSource initialized with val
func BuildUnderlyingLegSecurityIDSource(val string) UnderlyingLegSecurityIDSource {
	var field UnderlyingLegSecurityIDSource
	field.Value = val
	return field
}

//UnderlyingLegSecuritySubType is a STRING field
type UnderlyingLegSecuritySubType struct{ message.StringValue }

//Tag returns tag.UnderlyingLegSecuritySubType (1338)
func (f UnderlyingLegSecuritySubType) Tag() fix.Tag { return tag.UnderlyingLegSecuritySubType }

//BuildUnderlyingLegSecuritySubType returns a new UnderlyingLegSecuritySubType initialized with val
func BuildUnderlyingLegSecuritySubType(val string) UnderlyingLegSecuritySubType {
	var field UnderlyingLegSecuritySubType
	field.Value = val
	return field
}

//UnderlyingLegSecurityType is a STRING field
type UnderlyingLegSecurityType struct{ message.StringValue }

//Tag returns tag.UnderlyingLegSecurityType (1337)
func (f UnderlyingLegSecurityType) Tag() fix.Tag { return tag.UnderlyingLegSecurityType }

//BuildUnderlyingLegSecurityType returns a new UnderlyingLegSecurityType initialized with val
func BuildUnderlyingLegSecurityType(val string) UnderlyingLegSecurityType {
	var field UnderlyingLegSecurityType
	field.Value = val
	return field
}

//UnderlyingLegStrikePrice is a PRICE field
type UnderlyingLegStrikePrice struct{ message.PriceValue }

//Tag returns tag.UnderlyingLegStrikePrice (1340)
func (f UnderlyingLegStrikePrice) Tag() fix.Tag { return tag.UnderlyingLegStrikePrice }

//BuildUnderlyingLegStrikePrice returns a new UnderlyingLegStrikePrice initialized with val
func BuildUnderlyingLegStrikePrice(val float64) UnderlyingLegStrikePrice {
	var field UnderlyingLegStrikePrice
	field.Value = val
	return field
}

//UnderlyingLegSymbol is a STRING field
type UnderlyingLegSymbol struct{ message.StringValue }

//Tag returns tag.UnderlyingLegSymbol (1330)
func (f UnderlyingLegSymbol) Tag() fix.Tag { return tag.UnderlyingLegSymbol }

//BuildUnderlyingLegSymbol returns a new UnderlyingLegSymbol initialized with val
func BuildUnderlyingLegSymbol(val string) UnderlyingLegSymbol {
	var field UnderlyingLegSymbol
	field.Value = val
	return field
}

//UnderlyingLegSymbolSfx is a STRING field
type UnderlyingLegSymbolSfx struct{ message.StringValue }

//Tag returns tag.UnderlyingLegSymbolSfx (1331)
func (f UnderlyingLegSymbolSfx) Tag() fix.Tag { return tag.UnderlyingLegSymbolSfx }

//BuildUnderlyingLegSymbolSfx returns a new UnderlyingLegSymbolSfx initialized with val
func BuildUnderlyingLegSymbolSfx(val string) UnderlyingLegSymbolSfx {
	var field UnderlyingLegSymbolSfx
	field.Value = val
	return field
}

//UnderlyingLocaleOfIssue is a STRING field
type UnderlyingLocaleOfIssue struct{ message.StringValue }

//Tag returns tag.UnderlyingLocaleOfIssue (594)
func (f UnderlyingLocaleOfIssue) Tag() fix.Tag { return tag.UnderlyingLocaleOfIssue }

//BuildUnderlyingLocaleOfIssue returns a new UnderlyingLocaleOfIssue initialized with val
func BuildUnderlyingLocaleOfIssue(val string) UnderlyingLocaleOfIssue {
	var field UnderlyingLocaleOfIssue
	field.Value = val
	return field
}

//UnderlyingMaturityDate is a LOCALMKTDATE field
type UnderlyingMaturityDate struct{ message.LocalMktDateValue }

//Tag returns tag.UnderlyingMaturityDate (542)
func (f UnderlyingMaturityDate) Tag() fix.Tag { return tag.UnderlyingMaturityDate }

//BuildUnderlyingMaturityDate returns a new UnderlyingMaturityDate initialized with val
func BuildUnderlyingMaturityDate(val string) UnderlyingMaturityDate {
	var field UnderlyingMaturityDate
	field.Value = val
	return field
}

//UnderlyingMaturityDay is a DAYOFMONTH field
type UnderlyingMaturityDay struct{ message.DayOfMonthValue }

//Tag returns tag.UnderlyingMaturityDay (314)
func (f UnderlyingMaturityDay) Tag() fix.Tag { return tag.UnderlyingMaturityDay }

//BuildUnderlyingMaturityDay returns a new UnderlyingMaturityDay initialized with val
func BuildUnderlyingMaturityDay(val int) UnderlyingMaturityDay {
	var field UnderlyingMaturityDay
	field.Value = val
	return field
}

//UnderlyingMaturityMonthYear is a MONTHYEAR field
type UnderlyingMaturityMonthYear struct{ message.MonthYearValue }

//Tag returns tag.UnderlyingMaturityMonthYear (313)
func (f UnderlyingMaturityMonthYear) Tag() fix.Tag { return tag.UnderlyingMaturityMonthYear }

//BuildUnderlyingMaturityMonthYear returns a new UnderlyingMaturityMonthYear initialized with val
func BuildUnderlyingMaturityMonthYear(val string) UnderlyingMaturityMonthYear {
	var field UnderlyingMaturityMonthYear
	field.Value = val
	return field
}

//UnderlyingMaturityTime is a TZTIMEONLY field
type UnderlyingMaturityTime struct{ message.TZTimeOnlyValue }

//Tag returns tag.UnderlyingMaturityTime (1213)
func (f UnderlyingMaturityTime) Tag() fix.Tag { return tag.UnderlyingMaturityTime }

//UnderlyingNotionalPercentageOutstanding is a PERCENTAGE field
type UnderlyingNotionalPercentageOutstanding struct{ message.PercentageValue }

//Tag returns tag.UnderlyingNotionalPercentageOutstanding (1455)
func (f UnderlyingNotionalPercentageOutstanding) Tag() fix.Tag {
	return tag.UnderlyingNotionalPercentageOutstanding
}

//BuildUnderlyingNotionalPercentageOutstanding returns a new UnderlyingNotionalPercentageOutstanding initialized with val
func BuildUnderlyingNotionalPercentageOutstanding(val float64) UnderlyingNotionalPercentageOutstanding {
	var field UnderlyingNotionalPercentageOutstanding
	field.Value = val
	return field
}

//UnderlyingOptAttribute is a CHAR field
type UnderlyingOptAttribute struct{ message.CharValue }

//Tag returns tag.UnderlyingOptAttribute (317)
func (f UnderlyingOptAttribute) Tag() fix.Tag { return tag.UnderlyingOptAttribute }

//BuildUnderlyingOptAttribute returns a new UnderlyingOptAttribute initialized with val
func BuildUnderlyingOptAttribute(val string) UnderlyingOptAttribute {
	var field UnderlyingOptAttribute
	field.Value = val
	return field
}

//UnderlyingOriginalNotionalPercentageOutstanding is a PERCENTAGE field
type UnderlyingOriginalNotionalPercentageOutstanding struct{ message.PercentageValue }

//Tag returns tag.UnderlyingOriginalNotionalPercentageOutstanding (1456)
func (f UnderlyingOriginalNotionalPercentageOutstanding) Tag() fix.Tag {
	return tag.UnderlyingOriginalNotionalPercentageOutstanding
}

//BuildUnderlyingOriginalNotionalPercentageOutstanding returns a new UnderlyingOriginalNotionalPercentageOutstanding initialized with val
func BuildUnderlyingOriginalNotionalPercentageOutstanding(val float64) UnderlyingOriginalNotionalPercentageOutstanding {
	var field UnderlyingOriginalNotionalPercentageOutstanding
	field.Value = val
	return field
}

//UnderlyingPayAmount is a AMT field
type UnderlyingPayAmount struct{ message.AmtValue }

//Tag returns tag.UnderlyingPayAmount (985)
func (f UnderlyingPayAmount) Tag() fix.Tag { return tag.UnderlyingPayAmount }

//BuildUnderlyingPayAmount returns a new UnderlyingPayAmount initialized with val
func BuildUnderlyingPayAmount(val float64) UnderlyingPayAmount {
	var field UnderlyingPayAmount
	field.Value = val
	return field
}

//UnderlyingPriceDeterminationMethod is a INT field
type UnderlyingPriceDeterminationMethod struct{ message.IntValue }

//Tag returns tag.UnderlyingPriceDeterminationMethod (1481)
func (f UnderlyingPriceDeterminationMethod) Tag() fix.Tag {
	return tag.UnderlyingPriceDeterminationMethod
}

//BuildUnderlyingPriceDeterminationMethod returns a new UnderlyingPriceDeterminationMethod initialized with val
func BuildUnderlyingPriceDeterminationMethod(val int) UnderlyingPriceDeterminationMethod {
	var field UnderlyingPriceDeterminationMethod
	field.Value = val
	return field
}

//UnderlyingPriceUnitOfMeasure is a STRING field
type UnderlyingPriceUnitOfMeasure struct{ message.StringValue }

//Tag returns tag.UnderlyingPriceUnitOfMeasure (1424)
func (f UnderlyingPriceUnitOfMeasure) Tag() fix.Tag { return tag.UnderlyingPriceUnitOfMeasure }

//BuildUnderlyingPriceUnitOfMeasure returns a new UnderlyingPriceUnitOfMeasure initialized with val
func BuildUnderlyingPriceUnitOfMeasure(val string) UnderlyingPriceUnitOfMeasure {
	var field UnderlyingPriceUnitOfMeasure
	field.Value = val
	return field
}

//UnderlyingPriceUnitOfMeasureQty is a QTY field
type UnderlyingPriceUnitOfMeasureQty struct{ message.QtyValue }

//Tag returns tag.UnderlyingPriceUnitOfMeasureQty (1425)
func (f UnderlyingPriceUnitOfMeasureQty) Tag() fix.Tag { return tag.UnderlyingPriceUnitOfMeasureQty }

//BuildUnderlyingPriceUnitOfMeasureQty returns a new UnderlyingPriceUnitOfMeasureQty initialized with val
func BuildUnderlyingPriceUnitOfMeasureQty(val float64) UnderlyingPriceUnitOfMeasureQty {
	var field UnderlyingPriceUnitOfMeasureQty
	field.Value = val
	return field
}

//UnderlyingProduct is a INT field
type UnderlyingProduct struct{ message.IntValue }

//Tag returns tag.UnderlyingProduct (462)
func (f UnderlyingProduct) Tag() fix.Tag { return tag.UnderlyingProduct }

//BuildUnderlyingProduct returns a new UnderlyingProduct initialized with val
func BuildUnderlyingProduct(val int) UnderlyingProduct {
	var field UnderlyingProduct
	field.Value = val
	return field
}

//UnderlyingPutOrCall is a INT field
type UnderlyingPutOrCall struct{ message.IntValue }

//Tag returns tag.UnderlyingPutOrCall (315)
func (f UnderlyingPutOrCall) Tag() fix.Tag { return tag.UnderlyingPutOrCall }

//BuildUnderlyingPutOrCall returns a new UnderlyingPutOrCall initialized with val
func BuildUnderlyingPutOrCall(val int) UnderlyingPutOrCall {
	var field UnderlyingPutOrCall
	field.Value = val
	return field
}

//UnderlyingPx is a PRICE field
type UnderlyingPx struct{ message.PriceValue }

//Tag returns tag.UnderlyingPx (810)
func (f UnderlyingPx) Tag() fix.Tag { return tag.UnderlyingPx }

//BuildUnderlyingPx returns a new UnderlyingPx initialized with val
func BuildUnderlyingPx(val float64) UnderlyingPx {
	var field UnderlyingPx
	field.Value = val
	return field
}

//UnderlyingQty is a QTY field
type UnderlyingQty struct{ message.QtyValue }

//Tag returns tag.UnderlyingQty (879)
func (f UnderlyingQty) Tag() fix.Tag { return tag.UnderlyingQty }

//BuildUnderlyingQty returns a new UnderlyingQty initialized with val
func BuildUnderlyingQty(val float64) UnderlyingQty {
	var field UnderlyingQty
	field.Value = val
	return field
}

//UnderlyingRedemptionDate is a LOCALMKTDATE field
type UnderlyingRedemptionDate struct{ message.LocalMktDateValue }

//Tag returns tag.UnderlyingRedemptionDate (247)
func (f UnderlyingRedemptionDate) Tag() fix.Tag { return tag.UnderlyingRedemptionDate }

//BuildUnderlyingRedemptionDate returns a new UnderlyingRedemptionDate initialized with val
func BuildUnderlyingRedemptionDate(val string) UnderlyingRedemptionDate {
	var field UnderlyingRedemptionDate
	field.Value = val
	return field
}

//UnderlyingRepoCollateralSecurityType is a INT field
type UnderlyingRepoCollateralSecurityType struct{ message.IntValue }

//Tag returns tag.UnderlyingRepoCollateralSecurityType (243)
func (f UnderlyingRepoCollateralSecurityType) Tag() fix.Tag {
	return tag.UnderlyingRepoCollateralSecurityType
}

//BuildUnderlyingRepoCollateralSecurityType returns a new UnderlyingRepoCollateralSecurityType initialized with val
func BuildUnderlyingRepoCollateralSecurityType(val int) UnderlyingRepoCollateralSecurityType {
	var field UnderlyingRepoCollateralSecurityType
	field.Value = val
	return field
}

//UnderlyingRepurchaseRate is a PERCENTAGE field
type UnderlyingRepurchaseRate struct{ message.PercentageValue }

//Tag returns tag.UnderlyingRepurchaseRate (245)
func (f UnderlyingRepurchaseRate) Tag() fix.Tag { return tag.UnderlyingRepurchaseRate }

//BuildUnderlyingRepurchaseRate returns a new UnderlyingRepurchaseRate initialized with val
func BuildUnderlyingRepurchaseRate(val float64) UnderlyingRepurchaseRate {
	var field UnderlyingRepurchaseRate
	field.Value = val
	return field
}

//UnderlyingRepurchaseTerm is a INT field
type UnderlyingRepurchaseTerm struct{ message.IntValue }

//Tag returns tag.UnderlyingRepurchaseTerm (244)
func (f UnderlyingRepurchaseTerm) Tag() fix.Tag { return tag.UnderlyingRepurchaseTerm }

//BuildUnderlyingRepurchaseTerm returns a new UnderlyingRepurchaseTerm initialized with val
func BuildUnderlyingRepurchaseTerm(val int) UnderlyingRepurchaseTerm {
	var field UnderlyingRepurchaseTerm
	field.Value = val
	return field
}

//UnderlyingRestructuringType is a STRING field
type UnderlyingRestructuringType struct{ message.StringValue }

//Tag returns tag.UnderlyingRestructuringType (1453)
func (f UnderlyingRestructuringType) Tag() fix.Tag { return tag.UnderlyingRestructuringType }

//BuildUnderlyingRestructuringType returns a new UnderlyingRestructuringType initialized with val
func BuildUnderlyingRestructuringType(val string) UnderlyingRestructuringType {
	var field UnderlyingRestructuringType
	field.Value = val
	return field
}

//UnderlyingSecurityAltID is a STRING field
type UnderlyingSecurityAltID struct{ message.StringValue }

//Tag returns tag.UnderlyingSecurityAltID (458)
func (f UnderlyingSecurityAltID) Tag() fix.Tag { return tag.UnderlyingSecurityAltID }

//BuildUnderlyingSecurityAltID returns a new UnderlyingSecurityAltID initialized with val
func BuildUnderlyingSecurityAltID(val string) UnderlyingSecurityAltID {
	var field UnderlyingSecurityAltID
	field.Value = val
	return field
}

//UnderlyingSecurityAltIDSource is a STRING field
type UnderlyingSecurityAltIDSource struct{ message.StringValue }

//Tag returns tag.UnderlyingSecurityAltIDSource (459)
func (f UnderlyingSecurityAltIDSource) Tag() fix.Tag { return tag.UnderlyingSecurityAltIDSource }

//BuildUnderlyingSecurityAltIDSource returns a new UnderlyingSecurityAltIDSource initialized with val
func BuildUnderlyingSecurityAltIDSource(val string) UnderlyingSecurityAltIDSource {
	var field UnderlyingSecurityAltIDSource
	field.Value = val
	return field
}

//UnderlyingSecurityDesc is a STRING field
type UnderlyingSecurityDesc struct{ message.StringValue }

//Tag returns tag.UnderlyingSecurityDesc (307)
func (f UnderlyingSecurityDesc) Tag() fix.Tag { return tag.UnderlyingSecurityDesc }

//BuildUnderlyingSecurityDesc returns a new UnderlyingSecurityDesc initialized with val
func BuildUnderlyingSecurityDesc(val string) UnderlyingSecurityDesc {
	var field UnderlyingSecurityDesc
	field.Value = val
	return field
}

//UnderlyingSecurityExchange is a EXCHANGE field
type UnderlyingSecurityExchange struct{ message.ExchangeValue }

//Tag returns tag.UnderlyingSecurityExchange (308)
func (f UnderlyingSecurityExchange) Tag() fix.Tag { return tag.UnderlyingSecurityExchange }

//BuildUnderlyingSecurityExchange returns a new UnderlyingSecurityExchange initialized with val
func BuildUnderlyingSecurityExchange(val string) UnderlyingSecurityExchange {
	var field UnderlyingSecurityExchange
	field.Value = val
	return field
}

//UnderlyingSecurityID is a STRING field
type UnderlyingSecurityID struct{ message.StringValue }

//Tag returns tag.UnderlyingSecurityID (309)
func (f UnderlyingSecurityID) Tag() fix.Tag { return tag.UnderlyingSecurityID }

//BuildUnderlyingSecurityID returns a new UnderlyingSecurityID initialized with val
func BuildUnderlyingSecurityID(val string) UnderlyingSecurityID {
	var field UnderlyingSecurityID
	field.Value = val
	return field
}

//UnderlyingSecurityIDSource is a STRING field
type UnderlyingSecurityIDSource struct{ message.StringValue }

//Tag returns tag.UnderlyingSecurityIDSource (305)
func (f UnderlyingSecurityIDSource) Tag() fix.Tag { return tag.UnderlyingSecurityIDSource }

//BuildUnderlyingSecurityIDSource returns a new UnderlyingSecurityIDSource initialized with val
func BuildUnderlyingSecurityIDSource(val string) UnderlyingSecurityIDSource {
	var field UnderlyingSecurityIDSource
	field.Value = val
	return field
}

//UnderlyingSecuritySubType is a STRING field
type UnderlyingSecuritySubType struct{ message.StringValue }

//Tag returns tag.UnderlyingSecuritySubType (763)
func (f UnderlyingSecuritySubType) Tag() fix.Tag { return tag.UnderlyingSecuritySubType }

//BuildUnderlyingSecuritySubType returns a new UnderlyingSecuritySubType initialized with val
func BuildUnderlyingSecuritySubType(val string) UnderlyingSecuritySubType {
	var field UnderlyingSecuritySubType
	field.Value = val
	return field
}

//UnderlyingSecurityType is a STRING field
type UnderlyingSecurityType struct{ message.StringValue }

//Tag returns tag.UnderlyingSecurityType (310)
func (f UnderlyingSecurityType) Tag() fix.Tag { return tag.UnderlyingSecurityType }

//BuildUnderlyingSecurityType returns a new UnderlyingSecurityType initialized with val
func BuildUnderlyingSecurityType(val string) UnderlyingSecurityType {
	var field UnderlyingSecurityType
	field.Value = val
	return field
}

//UnderlyingSeniority is a STRING field
type UnderlyingSeniority struct{ message.StringValue }

//Tag returns tag.UnderlyingSeniority (1454)
func (f UnderlyingSeniority) Tag() fix.Tag { return tag.UnderlyingSeniority }

//BuildUnderlyingSeniority returns a new UnderlyingSeniority initialized with val
func BuildUnderlyingSeniority(val string) UnderlyingSeniority {
	var field UnderlyingSeniority
	field.Value = val
	return field
}

//UnderlyingSettlMethod is a STRING field
type UnderlyingSettlMethod struct{ message.StringValue }

//Tag returns tag.UnderlyingSettlMethod (1039)
func (f UnderlyingSettlMethod) Tag() fix.Tag { return tag.UnderlyingSettlMethod }

//BuildUnderlyingSettlMethod returns a new UnderlyingSettlMethod initialized with val
func BuildUnderlyingSettlMethod(val string) UnderlyingSettlMethod {
	var field UnderlyingSettlMethod
	field.Value = val
	return field
}

//UnderlyingSettlPrice is a PRICE field
type UnderlyingSettlPrice struct{ message.PriceValue }

//Tag returns tag.UnderlyingSettlPrice (732)
func (f UnderlyingSettlPrice) Tag() fix.Tag { return tag.UnderlyingSettlPrice }

//BuildUnderlyingSettlPrice returns a new UnderlyingSettlPrice initialized with val
func BuildUnderlyingSettlPrice(val float64) UnderlyingSettlPrice {
	var field UnderlyingSettlPrice
	field.Value = val
	return field
}

//UnderlyingSettlPriceType is a INT field
type UnderlyingSettlPriceType struct{ message.IntValue }

//Tag returns tag.UnderlyingSettlPriceType (733)
func (f UnderlyingSettlPriceType) Tag() fix.Tag { return tag.UnderlyingSettlPriceType }

//BuildUnderlyingSettlPriceType returns a new UnderlyingSettlPriceType initialized with val
func BuildUnderlyingSettlPriceType(val int) UnderlyingSettlPriceType {
	var field UnderlyingSettlPriceType
	field.Value = val
	return field
}

//UnderlyingSettlementDate is a LOCALMKTDATE field
type UnderlyingSettlementDate struct{ message.LocalMktDateValue }

//Tag returns tag.UnderlyingSettlementDate (987)
func (f UnderlyingSettlementDate) Tag() fix.Tag { return tag.UnderlyingSettlementDate }

//BuildUnderlyingSettlementDate returns a new UnderlyingSettlementDate initialized with val
func BuildUnderlyingSettlementDate(val string) UnderlyingSettlementDate {
	var field UnderlyingSettlementDate
	field.Value = val
	return field
}

//UnderlyingSettlementStatus is a STRING field
type UnderlyingSettlementStatus struct{ message.StringValue }

//Tag returns tag.UnderlyingSettlementStatus (988)
func (f UnderlyingSettlementStatus) Tag() fix.Tag { return tag.UnderlyingSettlementStatus }

//BuildUnderlyingSettlementStatus returns a new UnderlyingSettlementStatus initialized with val
func BuildUnderlyingSettlementStatus(val string) UnderlyingSettlementStatus {
	var field UnderlyingSettlementStatus
	field.Value = val
	return field
}

//UnderlyingSettlementType is a INT field
type UnderlyingSettlementType struct{ message.IntValue }

//Tag returns tag.UnderlyingSettlementType (975)
func (f UnderlyingSettlementType) Tag() fix.Tag { return tag.UnderlyingSettlementType }

//BuildUnderlyingSettlementType returns a new UnderlyingSettlementType initialized with val
func BuildUnderlyingSettlementType(val int) UnderlyingSettlementType {
	var field UnderlyingSettlementType
	field.Value = val
	return field
}

//UnderlyingStartValue is a AMT field
type UnderlyingStartValue struct{ message.AmtValue }

//Tag returns tag.UnderlyingStartValue (884)
func (f UnderlyingStartValue) Tag() fix.Tag { return tag.UnderlyingStartValue }

//BuildUnderlyingStartValue returns a new UnderlyingStartValue initialized with val
func BuildUnderlyingStartValue(val float64) UnderlyingStartValue {
	var field UnderlyingStartValue
	field.Value = val
	return field
}

//UnderlyingStateOrProvinceOfIssue is a STRING field
type UnderlyingStateOrProvinceOfIssue struct{ message.StringValue }

//Tag returns tag.UnderlyingStateOrProvinceOfIssue (593)
func (f UnderlyingStateOrProvinceOfIssue) Tag() fix.Tag { return tag.UnderlyingStateOrProvinceOfIssue }

//BuildUnderlyingStateOrProvinceOfIssue returns a new UnderlyingStateOrProvinceOfIssue initialized with val
func BuildUnderlyingStateOrProvinceOfIssue(val string) UnderlyingStateOrProvinceOfIssue {
	var field UnderlyingStateOrProvinceOfIssue
	field.Value = val
	return field
}

//UnderlyingStipType is a STRING field
type UnderlyingStipType struct{ message.StringValue }

//Tag returns tag.UnderlyingStipType (888)
func (f UnderlyingStipType) Tag() fix.Tag { return tag.UnderlyingStipType }

//BuildUnderlyingStipType returns a new UnderlyingStipType initialized with val
func BuildUnderlyingStipType(val string) UnderlyingStipType {
	var field UnderlyingStipType
	field.Value = val
	return field
}

//UnderlyingStipValue is a STRING field
type UnderlyingStipValue struct{ message.StringValue }

//Tag returns tag.UnderlyingStipValue (889)
func (f UnderlyingStipValue) Tag() fix.Tag { return tag.UnderlyingStipValue }

//BuildUnderlyingStipValue returns a new UnderlyingStipValue initialized with val
func BuildUnderlyingStipValue(val string) UnderlyingStipValue {
	var field UnderlyingStipValue
	field.Value = val
	return field
}

//UnderlyingStrikeCurrency is a CURRENCY field
type UnderlyingStrikeCurrency struct{ message.CurrencyValue }

//Tag returns tag.UnderlyingStrikeCurrency (941)
func (f UnderlyingStrikeCurrency) Tag() fix.Tag { return tag.UnderlyingStrikeCurrency }

//BuildUnderlyingStrikeCurrency returns a new UnderlyingStrikeCurrency initialized with val
func BuildUnderlyingStrikeCurrency(val string) UnderlyingStrikeCurrency {
	var field UnderlyingStrikeCurrency
	field.Value = val
	return field
}

//UnderlyingStrikePrice is a PRICE field
type UnderlyingStrikePrice struct{ message.PriceValue }

//Tag returns tag.UnderlyingStrikePrice (316)
func (f UnderlyingStrikePrice) Tag() fix.Tag { return tag.UnderlyingStrikePrice }

//BuildUnderlyingStrikePrice returns a new UnderlyingStrikePrice initialized with val
func BuildUnderlyingStrikePrice(val float64) UnderlyingStrikePrice {
	var field UnderlyingStrikePrice
	field.Value = val
	return field
}

//UnderlyingSymbol is a STRING field
type UnderlyingSymbol struct{ message.StringValue }

//Tag returns tag.UnderlyingSymbol (311)
func (f UnderlyingSymbol) Tag() fix.Tag { return tag.UnderlyingSymbol }

//BuildUnderlyingSymbol returns a new UnderlyingSymbol initialized with val
func BuildUnderlyingSymbol(val string) UnderlyingSymbol {
	var field UnderlyingSymbol
	field.Value = val
	return field
}

//UnderlyingSymbolSfx is a STRING field
type UnderlyingSymbolSfx struct{ message.StringValue }

//Tag returns tag.UnderlyingSymbolSfx (312)
func (f UnderlyingSymbolSfx) Tag() fix.Tag { return tag.UnderlyingSymbolSfx }

//BuildUnderlyingSymbolSfx returns a new UnderlyingSymbolSfx initialized with val
func BuildUnderlyingSymbolSfx(val string) UnderlyingSymbolSfx {
	var field UnderlyingSymbolSfx
	field.Value = val
	return field
}

//UnderlyingTimeUnit is a STRING field
type UnderlyingTimeUnit struct{ message.StringValue }

//Tag returns tag.UnderlyingTimeUnit (1000)
func (f UnderlyingTimeUnit) Tag() fix.Tag { return tag.UnderlyingTimeUnit }

//BuildUnderlyingTimeUnit returns a new UnderlyingTimeUnit initialized with val
func BuildUnderlyingTimeUnit(val string) UnderlyingTimeUnit {
	var field UnderlyingTimeUnit
	field.Value = val
	return field
}

//UnderlyingTradingSessionID is a STRING field
type UnderlyingTradingSessionID struct{ message.StringValue }

//Tag returns tag.UnderlyingTradingSessionID (822)
func (f UnderlyingTradingSessionID) Tag() fix.Tag { return tag.UnderlyingTradingSessionID }

//BuildUnderlyingTradingSessionID returns a new UnderlyingTradingSessionID initialized with val
func BuildUnderlyingTradingSessionID(val string) UnderlyingTradingSessionID {
	var field UnderlyingTradingSessionID
	field.Value = val
	return field
}

//UnderlyingTradingSessionSubID is a STRING field
type UnderlyingTradingSessionSubID struct{ message.StringValue }

//Tag returns tag.UnderlyingTradingSessionSubID (823)
func (f UnderlyingTradingSessionSubID) Tag() fix.Tag { return tag.UnderlyingTradingSessionSubID }

//BuildUnderlyingTradingSessionSubID returns a new UnderlyingTradingSessionSubID initialized with val
func BuildUnderlyingTradingSessionSubID(val string) UnderlyingTradingSessionSubID {
	var field UnderlyingTradingSessionSubID
	field.Value = val
	return field
}

//UnderlyingUnitOfMeasure is a STRING field
type UnderlyingUnitOfMeasure struct{ message.StringValue }

//Tag returns tag.UnderlyingUnitOfMeasure (998)
func (f UnderlyingUnitOfMeasure) Tag() fix.Tag { return tag.UnderlyingUnitOfMeasure }

//BuildUnderlyingUnitOfMeasure returns a new UnderlyingUnitOfMeasure initialized with val
func BuildUnderlyingUnitOfMeasure(val string) UnderlyingUnitOfMeasure {
	var field UnderlyingUnitOfMeasure
	field.Value = val
	return field
}

//UnderlyingUnitOfMeasureQty is a QTY field
type UnderlyingUnitOfMeasureQty struct{ message.QtyValue }

//Tag returns tag.UnderlyingUnitOfMeasureQty (1423)
func (f UnderlyingUnitOfMeasureQty) Tag() fix.Tag { return tag.UnderlyingUnitOfMeasureQty }

//BuildUnderlyingUnitOfMeasureQty returns a new UnderlyingUnitOfMeasureQty initialized with val
func BuildUnderlyingUnitOfMeasureQty(val float64) UnderlyingUnitOfMeasureQty {
	var field UnderlyingUnitOfMeasureQty
	field.Value = val
	return field
}

//UndlyInstrumentPartyID is a STRING field
type UndlyInstrumentPartyID struct{ message.StringValue }

//Tag returns tag.UndlyInstrumentPartyID (1059)
func (f UndlyInstrumentPartyID) Tag() fix.Tag { return tag.UndlyInstrumentPartyID }

//BuildUndlyInstrumentPartyID returns a new UndlyInstrumentPartyID initialized with val
func BuildUndlyInstrumentPartyID(val string) UndlyInstrumentPartyID {
	var field UndlyInstrumentPartyID
	field.Value = val
	return field
}

//UndlyInstrumentPartyIDSource is a CHAR field
type UndlyInstrumentPartyIDSource struct{ message.CharValue }

//Tag returns tag.UndlyInstrumentPartyIDSource (1060)
func (f UndlyInstrumentPartyIDSource) Tag() fix.Tag { return tag.UndlyInstrumentPartyIDSource }

//BuildUndlyInstrumentPartyIDSource returns a new UndlyInstrumentPartyIDSource initialized with val
func BuildUndlyInstrumentPartyIDSource(val string) UndlyInstrumentPartyIDSource {
	var field UndlyInstrumentPartyIDSource
	field.Value = val
	return field
}

//UndlyInstrumentPartyRole is a INT field
type UndlyInstrumentPartyRole struct{ message.IntValue }

//Tag returns tag.UndlyInstrumentPartyRole (1061)
func (f UndlyInstrumentPartyRole) Tag() fix.Tag { return tag.UndlyInstrumentPartyRole }

//BuildUndlyInstrumentPartyRole returns a new UndlyInstrumentPartyRole initialized with val
func BuildUndlyInstrumentPartyRole(val int) UndlyInstrumentPartyRole {
	var field UndlyInstrumentPartyRole
	field.Value = val
	return field
}

//UndlyInstrumentPartySubID is a STRING field
type UndlyInstrumentPartySubID struct{ message.StringValue }

//Tag returns tag.UndlyInstrumentPartySubID (1063)
func (f UndlyInstrumentPartySubID) Tag() fix.Tag { return tag.UndlyInstrumentPartySubID }

//BuildUndlyInstrumentPartySubID returns a new UndlyInstrumentPartySubID initialized with val
func BuildUndlyInstrumentPartySubID(val string) UndlyInstrumentPartySubID {
	var field UndlyInstrumentPartySubID
	field.Value = val
	return field
}

//UndlyInstrumentPartySubIDType is a INT field
type UndlyInstrumentPartySubIDType struct{ message.IntValue }

//Tag returns tag.UndlyInstrumentPartySubIDType (1064)
func (f UndlyInstrumentPartySubIDType) Tag() fix.Tag { return tag.UndlyInstrumentPartySubIDType }

//BuildUndlyInstrumentPartySubIDType returns a new UndlyInstrumentPartySubIDType initialized with val
func BuildUndlyInstrumentPartySubIDType(val int) UndlyInstrumentPartySubIDType {
	var field UndlyInstrumentPartySubIDType
	field.Value = val
	return field
}

//UnitOfMeasure is a STRING field
type UnitOfMeasure struct{ message.StringValue }

//Tag returns tag.UnitOfMeasure (996)
func (f UnitOfMeasure) Tag() fix.Tag { return tag.UnitOfMeasure }

//BuildUnitOfMeasure returns a new UnitOfMeasure initialized with val
func BuildUnitOfMeasure(val string) UnitOfMeasure {
	var field UnitOfMeasure
	field.Value = val
	return field
}

//UnitOfMeasureQty is a QTY field
type UnitOfMeasureQty struct{ message.QtyValue }

//Tag returns tag.UnitOfMeasureQty (1147)
func (f UnitOfMeasureQty) Tag() fix.Tag { return tag.UnitOfMeasureQty }

//BuildUnitOfMeasureQty returns a new UnitOfMeasureQty initialized with val
func BuildUnitOfMeasureQty(val float64) UnitOfMeasureQty {
	var field UnitOfMeasureQty
	field.Value = val
	return field
}

//UnsolicitedIndicator is a BOOLEAN field
type UnsolicitedIndicator struct{ message.BooleanValue }

//Tag returns tag.UnsolicitedIndicator (325)
func (f UnsolicitedIndicator) Tag() fix.Tag { return tag.UnsolicitedIndicator }

//BuildUnsolicitedIndicator returns a new UnsolicitedIndicator initialized with val
func BuildUnsolicitedIndicator(val bool) UnsolicitedIndicator {
	var field UnsolicitedIndicator
	field.Value = val
	return field
}

//Urgency is a CHAR field
type Urgency struct{ message.CharValue }

//Tag returns tag.Urgency (61)
func (f Urgency) Tag() fix.Tag { return tag.Urgency }

//BuildUrgency returns a new Urgency initialized with val
func BuildUrgency(val string) Urgency {
	var field Urgency
	field.Value = val
	return field
}

//UserRequestID is a STRING field
type UserRequestID struct{ message.StringValue }

//Tag returns tag.UserRequestID (923)
func (f UserRequestID) Tag() fix.Tag { return tag.UserRequestID }

//BuildUserRequestID returns a new UserRequestID initialized with val
func BuildUserRequestID(val string) UserRequestID {
	var field UserRequestID
	field.Value = val
	return field
}

//UserRequestType is a INT field
type UserRequestType struct{ message.IntValue }

//Tag returns tag.UserRequestType (924)
func (f UserRequestType) Tag() fix.Tag { return tag.UserRequestType }

//BuildUserRequestType returns a new UserRequestType initialized with val
func BuildUserRequestType(val int) UserRequestType {
	var field UserRequestType
	field.Value = val
	return field
}

//UserStatus is a INT field
type UserStatus struct{ message.IntValue }

//Tag returns tag.UserStatus (926)
func (f UserStatus) Tag() fix.Tag { return tag.UserStatus }

//BuildUserStatus returns a new UserStatus initialized with val
func BuildUserStatus(val int) UserStatus {
	var field UserStatus
	field.Value = val
	return field
}

//UserStatusText is a STRING field
type UserStatusText struct{ message.StringValue }

//Tag returns tag.UserStatusText (927)
func (f UserStatusText) Tag() fix.Tag { return tag.UserStatusText }

//BuildUserStatusText returns a new UserStatusText initialized with val
func BuildUserStatusText(val string) UserStatusText {
	var field UserStatusText
	field.Value = val
	return field
}

//Username is a STRING field
type Username struct{ message.StringValue }

//Tag returns tag.Username (553)
func (f Username) Tag() fix.Tag { return tag.Username }

//BuildUsername returns a new Username initialized with val
func BuildUsername(val string) Username {
	var field Username
	field.Value = val
	return field
}

//ValidUntilTime is a UTCTIMESTAMP field
type ValidUntilTime struct{ message.UTCTimestampValue }

//Tag returns tag.ValidUntilTime (62)
func (f ValidUntilTime) Tag() fix.Tag { return tag.ValidUntilTime }

//ValuationMethod is a STRING field
type ValuationMethod struct{ message.StringValue }

//Tag returns tag.ValuationMethod (1197)
func (f ValuationMethod) Tag() fix.Tag { return tag.ValuationMethod }

//BuildValuationMethod returns a new ValuationMethod initialized with val
func BuildValuationMethod(val string) ValuationMethod {
	var field ValuationMethod
	field.Value = val
	return field
}

//ValueOfFutures is a AMT field
type ValueOfFutures struct{ message.AmtValue }

//Tag returns tag.ValueOfFutures (408)
func (f ValueOfFutures) Tag() fix.Tag { return tag.ValueOfFutures }

//BuildValueOfFutures returns a new ValueOfFutures initialized with val
func BuildValueOfFutures(val float64) ValueOfFutures {
	var field ValueOfFutures
	field.Value = val
	return field
}

//VenueType is a CHAR field
type VenueType struct{ message.CharValue }

//Tag returns tag.VenueType (1430)
func (f VenueType) Tag() fix.Tag { return tag.VenueType }

//BuildVenueType returns a new VenueType initialized with val
func BuildVenueType(val string) VenueType {
	var field VenueType
	field.Value = val
	return field
}

//Volatility is a FLOAT field
type Volatility struct{ message.FloatValue }

//Tag returns tag.Volatility (1188)
func (f Volatility) Tag() fix.Tag { return tag.Volatility }

//BuildVolatility returns a new Volatility initialized with val
func BuildVolatility(val float64) Volatility {
	var field Volatility
	field.Value = val
	return field
}

//WaveNo is a STRING field
type WaveNo struct{ message.StringValue }

//Tag returns tag.WaveNo (105)
func (f WaveNo) Tag() fix.Tag { return tag.WaveNo }

//BuildWaveNo returns a new WaveNo initialized with val
func BuildWaveNo(val string) WaveNo {
	var field WaveNo
	field.Value = val
	return field
}

//WorkingIndicator is a BOOLEAN field
type WorkingIndicator struct{ message.BooleanValue }

//Tag returns tag.WorkingIndicator (636)
func (f WorkingIndicator) Tag() fix.Tag { return tag.WorkingIndicator }

//BuildWorkingIndicator returns a new WorkingIndicator initialized with val
func BuildWorkingIndicator(val bool) WorkingIndicator {
	var field WorkingIndicator
	field.Value = val
	return field
}

//WtAverageLiquidity is a PERCENTAGE field
type WtAverageLiquidity struct{ message.PercentageValue }

//Tag returns tag.WtAverageLiquidity (410)
func (f WtAverageLiquidity) Tag() fix.Tag { return tag.WtAverageLiquidity }

//BuildWtAverageLiquidity returns a new WtAverageLiquidity initialized with val
func BuildWtAverageLiquidity(val float64) WtAverageLiquidity {
	var field WtAverageLiquidity
	field.Value = val
	return field
}

//XmlData is a DATA field
type XmlData struct{ message.DataValue }

//Tag returns tag.XmlData (213)
func (f XmlData) Tag() fix.Tag { return tag.XmlData }

//BuildXmlData returns a new XmlData initialized with val
func BuildXmlData(val string) XmlData {
	var field XmlData
	field.Value = val
	return field
}

//XmlDataLen is a LENGTH field
type XmlDataLen struct{ message.LengthValue }

//Tag returns tag.XmlDataLen (212)
func (f XmlDataLen) Tag() fix.Tag { return tag.XmlDataLen }

//BuildXmlDataLen returns a new XmlDataLen initialized with val
func BuildXmlDataLen(val int) XmlDataLen {
	var field XmlDataLen
	field.Value = val
	return field
}

//Yield is a PERCENTAGE field
type Yield struct{ message.PercentageValue }

//Tag returns tag.Yield (236)
func (f Yield) Tag() fix.Tag { return tag.Yield }

//BuildYield returns a new Yield initialized with val
func BuildYield(val float64) Yield {
	var field Yield
	field.Value = val
	return field
}

//YieldCalcDate is a LOCALMKTDATE field
type YieldCalcDate struct{ message.LocalMktDateValue }

//Tag returns tag.YieldCalcDate (701)
func (f YieldCalcDate) Tag() fix.Tag { return tag.YieldCalcDate }

//BuildYieldCalcDate returns a new YieldCalcDate initialized with val
func BuildYieldCalcDate(val string) YieldCalcDate {
	var field YieldCalcDate
	field.Value = val
	return field
}

//YieldRedemptionDate is a LOCALMKTDATE field
type YieldRedemptionDate struct{ message.LocalMktDateValue }

//Tag returns tag.YieldRedemptionDate (696)
func (f YieldRedemptionDate) Tag() fix.Tag { return tag.YieldRedemptionDate }

//BuildYieldRedemptionDate returns a new YieldRedemptionDate initialized with val
func BuildYieldRedemptionDate(val string) YieldRedemptionDate {
	var field YieldRedemptionDate
	field.Value = val
	return field
}

//YieldRedemptionPrice is a PRICE field
type YieldRedemptionPrice struct{ message.PriceValue }

//Tag returns tag.YieldRedemptionPrice (697)
func (f YieldRedemptionPrice) Tag() fix.Tag { return tag.YieldRedemptionPrice }

//BuildYieldRedemptionPrice returns a new YieldRedemptionPrice initialized with val
func BuildYieldRedemptionPrice(val float64) YieldRedemptionPrice {
	var field YieldRedemptionPrice
	field.Value = val
	return field
}

//YieldRedemptionPriceType is a INT field
type YieldRedemptionPriceType struct{ message.IntValue }

//Tag returns tag.YieldRedemptionPriceType (698)
func (f YieldRedemptionPriceType) Tag() fix.Tag { return tag.YieldRedemptionPriceType }

//BuildYieldRedemptionPriceType returns a new YieldRedemptionPriceType initialized with val
func BuildYieldRedemptionPriceType(val int) YieldRedemptionPriceType {
	var field YieldRedemptionPriceType
	field.Value = val
	return field
}

//YieldType is a STRING field
type YieldType struct{ message.StringValue }

//Tag returns tag.YieldType (235)
func (f YieldType) Tag() fix.Tag { return tag.YieldType }

//BuildYieldType returns a new YieldType initialized with val
func BuildYieldType(val string) YieldType {
	var field YieldType
	field.Value = val
	return field
}
