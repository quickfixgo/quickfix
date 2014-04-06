package field

import (
	"github.com/quickfixgo/quickfix"
)
import (
	"github.com/quickfixgo/quickfix/tag"
)

type Account struct{ quickfix.StringValue }

func (f Account) Tag() tag.Tag { return tag.Account }

type AccountType struct{ quickfix.IntValue }

func (f AccountType) Tag() tag.Tag { return tag.AccountType }

type AccruedInterestAmt struct{ quickfix.AmtValue }

func (f AccruedInterestAmt) Tag() tag.Tag { return tag.AccruedInterestAmt }

type AccruedInterestRate struct{ quickfix.PercentageValue }

func (f AccruedInterestRate) Tag() tag.Tag { return tag.AccruedInterestRate }

type AcctIDSource struct{ quickfix.IntValue }

func (f AcctIDSource) Tag() tag.Tag { return tag.AcctIDSource }

type Adjustment struct{ quickfix.IntValue }

func (f Adjustment) Tag() tag.Tag { return tag.Adjustment }

type AdjustmentType struct{ quickfix.IntValue }

func (f AdjustmentType) Tag() tag.Tag { return tag.AdjustmentType }

type AdvId struct{ quickfix.StringValue }

func (f AdvId) Tag() tag.Tag { return tag.AdvId }

type AdvRefID struct{ quickfix.StringValue }

func (f AdvRefID) Tag() tag.Tag { return tag.AdvRefID }

type AdvSide struct{ quickfix.CharValue }

func (f AdvSide) Tag() tag.Tag { return tag.AdvSide }

type AdvTransType struct{ quickfix.StringValue }

func (f AdvTransType) Tag() tag.Tag { return tag.AdvTransType }

type AffectedOrderID struct{ quickfix.StringValue }

func (f AffectedOrderID) Tag() tag.Tag { return tag.AffectedOrderID }

type AffectedSecondaryOrderID struct{ quickfix.StringValue }

func (f AffectedSecondaryOrderID) Tag() tag.Tag { return tag.AffectedSecondaryOrderID }

type AffirmStatus struct{ quickfix.IntValue }

func (f AffirmStatus) Tag() tag.Tag { return tag.AffirmStatus }

type AggregatedBook struct{ quickfix.BooleanValue }

func (f AggregatedBook) Tag() tag.Tag { return tag.AggregatedBook }

type AggressorIndicator struct{ quickfix.BooleanValue }

func (f AggressorIndicator) Tag() tag.Tag { return tag.AggressorIndicator }

type AgreementCurrency struct{ quickfix.CurrencyValue }

func (f AgreementCurrency) Tag() tag.Tag { return tag.AgreementCurrency }

type AgreementDate struct{ quickfix.LocalMktDateValue }

func (f AgreementDate) Tag() tag.Tag { return tag.AgreementDate }

type AgreementDesc struct{ quickfix.StringValue }

func (f AgreementDesc) Tag() tag.Tag { return tag.AgreementDesc }

type AgreementID struct{ quickfix.StringValue }

func (f AgreementID) Tag() tag.Tag { return tag.AgreementID }

type AllocAccount struct{ quickfix.StringValue }

func (f AllocAccount) Tag() tag.Tag { return tag.AllocAccount }

type AllocAccountType struct{ quickfix.IntValue }

func (f AllocAccountType) Tag() tag.Tag { return tag.AllocAccountType }

type AllocAccruedInterestAmt struct{ quickfix.AmtValue }

func (f AllocAccruedInterestAmt) Tag() tag.Tag { return tag.AllocAccruedInterestAmt }

type AllocAcctIDSource struct{ quickfix.IntValue }

func (f AllocAcctIDSource) Tag() tag.Tag { return tag.AllocAcctIDSource }

type AllocAvgPx struct{ quickfix.PriceValue }

func (f AllocAvgPx) Tag() tag.Tag { return tag.AllocAvgPx }

type AllocCancReplaceReason struct{ quickfix.IntValue }

func (f AllocCancReplaceReason) Tag() tag.Tag { return tag.AllocCancReplaceReason }

type AllocClearingFeeIndicator struct{ quickfix.StringValue }

func (f AllocClearingFeeIndicator) Tag() tag.Tag { return tag.AllocClearingFeeIndicator }

type AllocCustomerCapacity struct{ quickfix.StringValue }

func (f AllocCustomerCapacity) Tag() tag.Tag { return tag.AllocCustomerCapacity }

type AllocHandlInst struct{ quickfix.IntValue }

func (f AllocHandlInst) Tag() tag.Tag { return tag.AllocHandlInst }

type AllocID struct{ quickfix.StringValue }

func (f AllocID) Tag() tag.Tag { return tag.AllocID }

type AllocInterestAtMaturity struct{ quickfix.AmtValue }

func (f AllocInterestAtMaturity) Tag() tag.Tag { return tag.AllocInterestAtMaturity }

type AllocIntermedReqType struct{ quickfix.IntValue }

func (f AllocIntermedReqType) Tag() tag.Tag { return tag.AllocIntermedReqType }

type AllocLinkID struct{ quickfix.StringValue }

func (f AllocLinkID) Tag() tag.Tag { return tag.AllocLinkID }

type AllocLinkType struct{ quickfix.IntValue }

func (f AllocLinkType) Tag() tag.Tag { return tag.AllocLinkType }

type AllocMethod struct{ quickfix.IntValue }

func (f AllocMethod) Tag() tag.Tag { return tag.AllocMethod }

type AllocNetMoney struct{ quickfix.AmtValue }

func (f AllocNetMoney) Tag() tag.Tag { return tag.AllocNetMoney }

type AllocNoOrdersType struct{ quickfix.IntValue }

func (f AllocNoOrdersType) Tag() tag.Tag { return tag.AllocNoOrdersType }

type AllocPositionEffect struct{ quickfix.CharValue }

func (f AllocPositionEffect) Tag() tag.Tag { return tag.AllocPositionEffect }

type AllocPrice struct{ quickfix.PriceValue }

func (f AllocPrice) Tag() tag.Tag { return tag.AllocPrice }

type AllocQty struct{ quickfix.QtyValue }

func (f AllocQty) Tag() tag.Tag { return tag.AllocQty }

type AllocRejCode struct{ quickfix.IntValue }

func (f AllocRejCode) Tag() tag.Tag { return tag.AllocRejCode }

type AllocReportID struct{ quickfix.StringValue }

func (f AllocReportID) Tag() tag.Tag { return tag.AllocReportID }

type AllocReportRefID struct{ quickfix.StringValue }

func (f AllocReportRefID) Tag() tag.Tag { return tag.AllocReportRefID }

type AllocReportType struct{ quickfix.IntValue }

func (f AllocReportType) Tag() tag.Tag { return tag.AllocReportType }

type AllocSettlCurrAmt struct{ quickfix.AmtValue }

func (f AllocSettlCurrAmt) Tag() tag.Tag { return tag.AllocSettlCurrAmt }

type AllocSettlCurrency struct{ quickfix.CurrencyValue }

func (f AllocSettlCurrency) Tag() tag.Tag { return tag.AllocSettlCurrency }

type AllocSettlInstType struct{ quickfix.IntValue }

func (f AllocSettlInstType) Tag() tag.Tag { return tag.AllocSettlInstType }

type AllocShares struct{ quickfix.QtyValue }

func (f AllocShares) Tag() tag.Tag { return tag.AllocShares }

type AllocStatus struct{ quickfix.IntValue }

func (f AllocStatus) Tag() tag.Tag { return tag.AllocStatus }

type AllocText struct{ quickfix.StringValue }

func (f AllocText) Tag() tag.Tag { return tag.AllocText }

type AllocTransType struct{ quickfix.CharValue }

func (f AllocTransType) Tag() tag.Tag { return tag.AllocTransType }

type AllocType struct{ quickfix.IntValue }

func (f AllocType) Tag() tag.Tag { return tag.AllocType }

type AllowableOneSidednessCurr struct{ quickfix.CurrencyValue }

func (f AllowableOneSidednessCurr) Tag() tag.Tag { return tag.AllowableOneSidednessCurr }

type AllowableOneSidednessPct struct{ quickfix.PercentageValue }

func (f AllowableOneSidednessPct) Tag() tag.Tag { return tag.AllowableOneSidednessPct }

type AllowableOneSidednessValue struct{ quickfix.AmtValue }

func (f AllowableOneSidednessValue) Tag() tag.Tag { return tag.AllowableOneSidednessValue }

type AltMDSourceID struct{ quickfix.StringValue }

func (f AltMDSourceID) Tag() tag.Tag { return tag.AltMDSourceID }

type ApplBegSeqNum struct{ quickfix.SeqNumValue }

func (f ApplBegSeqNum) Tag() tag.Tag { return tag.ApplBegSeqNum }

type ApplEndSeqNum struct{ quickfix.SeqNumValue }

func (f ApplEndSeqNum) Tag() tag.Tag { return tag.ApplEndSeqNum }

type ApplExtID struct{ quickfix.IntValue }

func (f ApplExtID) Tag() tag.Tag { return tag.ApplExtID }

type ApplID struct{ quickfix.StringValue }

func (f ApplID) Tag() tag.Tag { return tag.ApplID }

type ApplLastSeqNum struct{ quickfix.SeqNumValue }

func (f ApplLastSeqNum) Tag() tag.Tag { return tag.ApplLastSeqNum }

type ApplNewSeqNum struct{ quickfix.SeqNumValue }

func (f ApplNewSeqNum) Tag() tag.Tag { return tag.ApplNewSeqNum }

type ApplQueueAction struct{ quickfix.IntValue }

func (f ApplQueueAction) Tag() tag.Tag { return tag.ApplQueueAction }

type ApplQueueDepth struct{ quickfix.IntValue }

func (f ApplQueueDepth) Tag() tag.Tag { return tag.ApplQueueDepth }

type ApplQueueMax struct{ quickfix.IntValue }

func (f ApplQueueMax) Tag() tag.Tag { return tag.ApplQueueMax }

type ApplQueueResolution struct{ quickfix.IntValue }

func (f ApplQueueResolution) Tag() tag.Tag { return tag.ApplQueueResolution }

type ApplReportID struct{ quickfix.StringValue }

func (f ApplReportID) Tag() tag.Tag { return tag.ApplReportID }

type ApplReportType struct{ quickfix.IntValue }

func (f ApplReportType) Tag() tag.Tag { return tag.ApplReportType }

type ApplReqID struct{ quickfix.StringValue }

func (f ApplReqID) Tag() tag.Tag { return tag.ApplReqID }

type ApplReqType struct{ quickfix.IntValue }

func (f ApplReqType) Tag() tag.Tag { return tag.ApplReqType }

type ApplResendFlag struct{ quickfix.BooleanValue }

func (f ApplResendFlag) Tag() tag.Tag { return tag.ApplResendFlag }

type ApplResponseError struct{ quickfix.IntValue }

func (f ApplResponseError) Tag() tag.Tag { return tag.ApplResponseError }

type ApplResponseID struct{ quickfix.StringValue }

func (f ApplResponseID) Tag() tag.Tag { return tag.ApplResponseID }

type ApplResponseType struct{ quickfix.IntValue }

func (f ApplResponseType) Tag() tag.Tag { return tag.ApplResponseType }

type ApplSeqNum struct{ quickfix.SeqNumValue }

func (f ApplSeqNum) Tag() tag.Tag { return tag.ApplSeqNum }

type ApplTotalMessageCount struct{ quickfix.IntValue }

func (f ApplTotalMessageCount) Tag() tag.Tag { return tag.ApplTotalMessageCount }

type ApplVerID struct{ quickfix.StringValue }

func (f ApplVerID) Tag() tag.Tag { return tag.ApplVerID }

type AsOfIndicator struct{ quickfix.CharValue }

func (f AsOfIndicator) Tag() tag.Tag { return tag.AsOfIndicator }

type AsgnReqID struct{ quickfix.StringValue }

func (f AsgnReqID) Tag() tag.Tag { return tag.AsgnReqID }

type AsgnRptID struct{ quickfix.StringValue }

func (f AsgnRptID) Tag() tag.Tag { return tag.AsgnRptID }

type AssignmentMethod struct{ quickfix.CharValue }

func (f AssignmentMethod) Tag() tag.Tag { return tag.AssignmentMethod }

type AssignmentUnit struct{ quickfix.QtyValue }

func (f AssignmentUnit) Tag() tag.Tag { return tag.AssignmentUnit }

type AttachmentPoint struct{ quickfix.PercentageValue }

func (f AttachmentPoint) Tag() tag.Tag { return tag.AttachmentPoint }

type AutoAcceptIndicator struct{ quickfix.BooleanValue }

func (f AutoAcceptIndicator) Tag() tag.Tag { return tag.AutoAcceptIndicator }

type AvgParPx struct{ quickfix.PriceValue }

func (f AvgParPx) Tag() tag.Tag { return tag.AvgParPx }

type AvgPrxPrecision struct{ quickfix.IntValue }

func (f AvgPrxPrecision) Tag() tag.Tag { return tag.AvgPrxPrecision }

type AvgPx struct{ quickfix.PriceValue }

func (f AvgPx) Tag() tag.Tag { return tag.AvgPx }

type AvgPxIndicator struct{ quickfix.IntValue }

func (f AvgPxIndicator) Tag() tag.Tag { return tag.AvgPxIndicator }

type AvgPxPrecision struct{ quickfix.IntValue }

func (f AvgPxPrecision) Tag() tag.Tag { return tag.AvgPxPrecision }

type BasisFeatureDate struct{ quickfix.LocalMktDateValue }

func (f BasisFeatureDate) Tag() tag.Tag { return tag.BasisFeatureDate }

type BasisFeaturePrice struct{ quickfix.PriceValue }

func (f BasisFeaturePrice) Tag() tag.Tag { return tag.BasisFeaturePrice }

type BasisPxType struct{ quickfix.CharValue }

func (f BasisPxType) Tag() tag.Tag { return tag.BasisPxType }

type BeginSeqNo struct{ quickfix.SeqNumValue }

func (f BeginSeqNo) Tag() tag.Tag { return tag.BeginSeqNo }

type BeginString struct{ quickfix.StringValue }

func (f BeginString) Tag() tag.Tag { return tag.BeginString }

type Benchmark struct{ quickfix.CharValue }

func (f Benchmark) Tag() tag.Tag { return tag.Benchmark }

type BenchmarkCurveCurrency struct{ quickfix.CurrencyValue }

func (f BenchmarkCurveCurrency) Tag() tag.Tag { return tag.BenchmarkCurveCurrency }

type BenchmarkCurveName struct{ quickfix.StringValue }

func (f BenchmarkCurveName) Tag() tag.Tag { return tag.BenchmarkCurveName }

type BenchmarkCurvePoint struct{ quickfix.StringValue }

func (f BenchmarkCurvePoint) Tag() tag.Tag { return tag.BenchmarkCurvePoint }

type BenchmarkPrice struct{ quickfix.PriceValue }

func (f BenchmarkPrice) Tag() tag.Tag { return tag.BenchmarkPrice }

type BenchmarkPriceType struct{ quickfix.IntValue }

func (f BenchmarkPriceType) Tag() tag.Tag { return tag.BenchmarkPriceType }

type BenchmarkSecurityID struct{ quickfix.StringValue }

func (f BenchmarkSecurityID) Tag() tag.Tag { return tag.BenchmarkSecurityID }

type BenchmarkSecurityIDSource struct{ quickfix.StringValue }

func (f BenchmarkSecurityIDSource) Tag() tag.Tag { return tag.BenchmarkSecurityIDSource }

type BidDescriptor struct{ quickfix.StringValue }

func (f BidDescriptor) Tag() tag.Tag { return tag.BidDescriptor }

type BidDescriptorType struct{ quickfix.IntValue }

func (f BidDescriptorType) Tag() tag.Tag { return tag.BidDescriptorType }

type BidForwardPoints struct{ quickfix.PriceOffsetValue }

func (f BidForwardPoints) Tag() tag.Tag { return tag.BidForwardPoints }

type BidForwardPoints2 struct{ quickfix.PriceOffsetValue }

func (f BidForwardPoints2) Tag() tag.Tag { return tag.BidForwardPoints2 }

type BidID struct{ quickfix.StringValue }

func (f BidID) Tag() tag.Tag { return tag.BidID }

type BidPx struct{ quickfix.PriceValue }

func (f BidPx) Tag() tag.Tag { return tag.BidPx }

type BidRequestTransType struct{ quickfix.CharValue }

func (f BidRequestTransType) Tag() tag.Tag { return tag.BidRequestTransType }

type BidSize struct{ quickfix.QtyValue }

func (f BidSize) Tag() tag.Tag { return tag.BidSize }

type BidSpotRate struct{ quickfix.PriceValue }

func (f BidSpotRate) Tag() tag.Tag { return tag.BidSpotRate }

type BidSwapPoints struct{ quickfix.PriceOffsetValue }

func (f BidSwapPoints) Tag() tag.Tag { return tag.BidSwapPoints }

type BidTradeType struct{ quickfix.CharValue }

func (f BidTradeType) Tag() tag.Tag { return tag.BidTradeType }

type BidType struct{ quickfix.IntValue }

func (f BidType) Tag() tag.Tag { return tag.BidType }

type BidYield struct{ quickfix.PercentageValue }

func (f BidYield) Tag() tag.Tag { return tag.BidYield }

type BodyLength struct{ quickfix.LengthValue }

func (f BodyLength) Tag() tag.Tag { return tag.BodyLength }

type BookingRefID struct{ quickfix.StringValue }

func (f BookingRefID) Tag() tag.Tag { return tag.BookingRefID }

type BookingType struct{ quickfix.IntValue }

func (f BookingType) Tag() tag.Tag { return tag.BookingType }

type BookingUnit struct{ quickfix.CharValue }

func (f BookingUnit) Tag() tag.Tag { return tag.BookingUnit }

type BrokerOfCredit struct{ quickfix.StringValue }

func (f BrokerOfCredit) Tag() tag.Tag { return tag.BrokerOfCredit }

type BusinessRejectReason struct{ quickfix.IntValue }

func (f BusinessRejectReason) Tag() tag.Tag { return tag.BusinessRejectReason }

type BusinessRejectRefID struct{ quickfix.StringValue }

func (f BusinessRejectRefID) Tag() tag.Tag { return tag.BusinessRejectRefID }

type BuyVolume struct{ quickfix.QtyValue }

func (f BuyVolume) Tag() tag.Tag { return tag.BuyVolume }

type CFICode struct{ quickfix.StringValue }

func (f CFICode) Tag() tag.Tag { return tag.CFICode }

type CPProgram struct{ quickfix.IntValue }

func (f CPProgram) Tag() tag.Tag { return tag.CPProgram }

type CPRegType struct{ quickfix.StringValue }

func (f CPRegType) Tag() tag.Tag { return tag.CPRegType }

type CalculatedCcyLastQty struct{ quickfix.QtyValue }

func (f CalculatedCcyLastQty) Tag() tag.Tag { return tag.CalculatedCcyLastQty }

type CancellationRights struct{ quickfix.CharValue }

func (f CancellationRights) Tag() tag.Tag { return tag.CancellationRights }

type CapPrice struct{ quickfix.PriceValue }

func (f CapPrice) Tag() tag.Tag { return tag.CapPrice }

type CardExpDate struct{ quickfix.LocalMktDateValue }

func (f CardExpDate) Tag() tag.Tag { return tag.CardExpDate }

type CardHolderName struct{ quickfix.StringValue }

func (f CardHolderName) Tag() tag.Tag { return tag.CardHolderName }

type CardIssNo struct{ quickfix.StringValue }

func (f CardIssNo) Tag() tag.Tag { return tag.CardIssNo }

type CardIssNum struct{ quickfix.StringValue }

func (f CardIssNum) Tag() tag.Tag { return tag.CardIssNum }

type CardNumber struct{ quickfix.StringValue }

func (f CardNumber) Tag() tag.Tag { return tag.CardNumber }

type CardStartDate struct{ quickfix.LocalMktDateValue }

func (f CardStartDate) Tag() tag.Tag { return tag.CardStartDate }

type CashDistribAgentAcctName struct{ quickfix.StringValue }

func (f CashDistribAgentAcctName) Tag() tag.Tag { return tag.CashDistribAgentAcctName }

type CashDistribAgentAcctNumber struct{ quickfix.StringValue }

func (f CashDistribAgentAcctNumber) Tag() tag.Tag { return tag.CashDistribAgentAcctNumber }

type CashDistribAgentCode struct{ quickfix.StringValue }

func (f CashDistribAgentCode) Tag() tag.Tag { return tag.CashDistribAgentCode }

type CashDistribAgentName struct{ quickfix.StringValue }

func (f CashDistribAgentName) Tag() tag.Tag { return tag.CashDistribAgentName }

type CashDistribCurr struct{ quickfix.CurrencyValue }

func (f CashDistribCurr) Tag() tag.Tag { return tag.CashDistribCurr }

type CashDistribPayRef struct{ quickfix.StringValue }

func (f CashDistribPayRef) Tag() tag.Tag { return tag.CashDistribPayRef }

type CashMargin struct{ quickfix.CharValue }

func (f CashMargin) Tag() tag.Tag { return tag.CashMargin }

type CashOrderQty struct{ quickfix.QtyValue }

func (f CashOrderQty) Tag() tag.Tag { return tag.CashOrderQty }

type CashOutstanding struct{ quickfix.AmtValue }

func (f CashOutstanding) Tag() tag.Tag { return tag.CashOutstanding }

type CashSettlAgentAcctName struct{ quickfix.StringValue }

func (f CashSettlAgentAcctName) Tag() tag.Tag { return tag.CashSettlAgentAcctName }

type CashSettlAgentAcctNum struct{ quickfix.StringValue }

func (f CashSettlAgentAcctNum) Tag() tag.Tag { return tag.CashSettlAgentAcctNum }

type CashSettlAgentCode struct{ quickfix.StringValue }

func (f CashSettlAgentCode) Tag() tag.Tag { return tag.CashSettlAgentCode }

type CashSettlAgentContactName struct{ quickfix.StringValue }

func (f CashSettlAgentContactName) Tag() tag.Tag { return tag.CashSettlAgentContactName }

type CashSettlAgentContactPhone struct{ quickfix.StringValue }

func (f CashSettlAgentContactPhone) Tag() tag.Tag { return tag.CashSettlAgentContactPhone }

type CashSettlAgentName struct{ quickfix.StringValue }

func (f CashSettlAgentName) Tag() tag.Tag { return tag.CashSettlAgentName }

type CcyAmt struct{ quickfix.AmtValue }

func (f CcyAmt) Tag() tag.Tag { return tag.CcyAmt }

type CheckSum struct{ quickfix.StringValue }

func (f CheckSum) Tag() tag.Tag { return tag.CheckSum }

type ClOrdID struct{ quickfix.StringValue }

func (f ClOrdID) Tag() tag.Tag { return tag.ClOrdID }

type ClOrdLinkID struct{ quickfix.StringValue }

func (f ClOrdLinkID) Tag() tag.Tag { return tag.ClOrdLinkID }

type ClearingAccount struct{ quickfix.StringValue }

func (f ClearingAccount) Tag() tag.Tag { return tag.ClearingAccount }

type ClearingBusinessDate struct{ quickfix.LocalMktDateValue }

func (f ClearingBusinessDate) Tag() tag.Tag { return tag.ClearingBusinessDate }

type ClearingFeeIndicator struct{ quickfix.StringValue }

func (f ClearingFeeIndicator) Tag() tag.Tag { return tag.ClearingFeeIndicator }

type ClearingFirm struct{ quickfix.StringValue }

func (f ClearingFirm) Tag() tag.Tag { return tag.ClearingFirm }

type ClearingInstruction struct{ quickfix.IntValue }

func (f ClearingInstruction) Tag() tag.Tag { return tag.ClearingInstruction }

type ClientBidID struct{ quickfix.StringValue }

func (f ClientBidID) Tag() tag.Tag { return tag.ClientBidID }

type ClientID struct{ quickfix.StringValue }

func (f ClientID) Tag() tag.Tag { return tag.ClientID }

type CollAction struct{ quickfix.IntValue }

func (f CollAction) Tag() tag.Tag { return tag.CollAction }

type CollApplType struct{ quickfix.IntValue }

func (f CollApplType) Tag() tag.Tag { return tag.CollApplType }

type CollAsgnID struct{ quickfix.StringValue }

func (f CollAsgnID) Tag() tag.Tag { return tag.CollAsgnID }

type CollAsgnReason struct{ quickfix.IntValue }

func (f CollAsgnReason) Tag() tag.Tag { return tag.CollAsgnReason }

type CollAsgnRefID struct{ quickfix.StringValue }

func (f CollAsgnRefID) Tag() tag.Tag { return tag.CollAsgnRefID }

type CollAsgnRejectReason struct{ quickfix.IntValue }

func (f CollAsgnRejectReason) Tag() tag.Tag { return tag.CollAsgnRejectReason }

type CollAsgnRespType struct{ quickfix.IntValue }

func (f CollAsgnRespType) Tag() tag.Tag { return tag.CollAsgnRespType }

type CollAsgnTransType struct{ quickfix.IntValue }

func (f CollAsgnTransType) Tag() tag.Tag { return tag.CollAsgnTransType }

type CollInquiryID struct{ quickfix.StringValue }

func (f CollInquiryID) Tag() tag.Tag { return tag.CollInquiryID }

type CollInquiryQualifier struct{ quickfix.IntValue }

func (f CollInquiryQualifier) Tag() tag.Tag { return tag.CollInquiryQualifier }

type CollInquiryResult struct{ quickfix.IntValue }

func (f CollInquiryResult) Tag() tag.Tag { return tag.CollInquiryResult }

type CollInquiryStatus struct{ quickfix.IntValue }

func (f CollInquiryStatus) Tag() tag.Tag { return tag.CollInquiryStatus }

type CollReqID struct{ quickfix.StringValue }

func (f CollReqID) Tag() tag.Tag { return tag.CollReqID }

type CollRespID struct{ quickfix.StringValue }

func (f CollRespID) Tag() tag.Tag { return tag.CollRespID }

type CollRptID struct{ quickfix.StringValue }

func (f CollRptID) Tag() tag.Tag { return tag.CollRptID }

type CollStatus struct{ quickfix.IntValue }

func (f CollStatus) Tag() tag.Tag { return tag.CollStatus }

type CommCurrency struct{ quickfix.CurrencyValue }

func (f CommCurrency) Tag() tag.Tag { return tag.CommCurrency }

type CommType struct{ quickfix.CharValue }

func (f CommType) Tag() tag.Tag { return tag.CommType }

type Commission struct{ quickfix.AmtValue }

func (f Commission) Tag() tag.Tag { return tag.Commission }

type ComplexEventCondition struct{ quickfix.IntValue }

func (f ComplexEventCondition) Tag() tag.Tag { return tag.ComplexEventCondition }

type ComplexEventEndDate struct{ quickfix.UTCTimestampValue }

func (f ComplexEventEndDate) Tag() tag.Tag { return tag.ComplexEventEndDate }

type ComplexEventEndTime struct{ quickfix.UTCTimeOnlyValue }

func (f ComplexEventEndTime) Tag() tag.Tag { return tag.ComplexEventEndTime }

type ComplexEventPrice struct{ quickfix.PriceValue }

func (f ComplexEventPrice) Tag() tag.Tag { return tag.ComplexEventPrice }

type ComplexEventPriceBoundaryMethod struct{ quickfix.IntValue }

func (f ComplexEventPriceBoundaryMethod) Tag() tag.Tag { return tag.ComplexEventPriceBoundaryMethod }

type ComplexEventPriceBoundaryPrecision struct{ quickfix.PercentageValue }

func (f ComplexEventPriceBoundaryPrecision) Tag() tag.Tag {
	return tag.ComplexEventPriceBoundaryPrecision
}

type ComplexEventPriceTimeType struct{ quickfix.IntValue }

func (f ComplexEventPriceTimeType) Tag() tag.Tag { return tag.ComplexEventPriceTimeType }

type ComplexEventStartDate struct{ quickfix.UTCTimestampValue }

func (f ComplexEventStartDate) Tag() tag.Tag { return tag.ComplexEventStartDate }

type ComplexEventStartTime struct{ quickfix.UTCTimeOnlyValue }

func (f ComplexEventStartTime) Tag() tag.Tag { return tag.ComplexEventStartTime }

type ComplexEventType struct{ quickfix.IntValue }

func (f ComplexEventType) Tag() tag.Tag { return tag.ComplexEventType }

type ComplexOptPayoutAmount struct{ quickfix.AmtValue }

func (f ComplexOptPayoutAmount) Tag() tag.Tag { return tag.ComplexOptPayoutAmount }

type ComplianceID struct{ quickfix.StringValue }

func (f ComplianceID) Tag() tag.Tag { return tag.ComplianceID }

type Concession struct{ quickfix.AmtValue }

func (f Concession) Tag() tag.Tag { return tag.Concession }

type ConfirmID struct{ quickfix.StringValue }

func (f ConfirmID) Tag() tag.Tag { return tag.ConfirmID }

type ConfirmRefID struct{ quickfix.StringValue }

func (f ConfirmRefID) Tag() tag.Tag { return tag.ConfirmRefID }

type ConfirmRejReason struct{ quickfix.IntValue }

func (f ConfirmRejReason) Tag() tag.Tag { return tag.ConfirmRejReason }

type ConfirmReqID struct{ quickfix.StringValue }

func (f ConfirmReqID) Tag() tag.Tag { return tag.ConfirmReqID }

type ConfirmStatus struct{ quickfix.IntValue }

func (f ConfirmStatus) Tag() tag.Tag { return tag.ConfirmStatus }

type ConfirmTransType struct{ quickfix.IntValue }

func (f ConfirmTransType) Tag() tag.Tag { return tag.ConfirmTransType }

type ConfirmType struct{ quickfix.IntValue }

func (f ConfirmType) Tag() tag.Tag { return tag.ConfirmType }

type ContAmtCurr struct{ quickfix.CurrencyValue }

func (f ContAmtCurr) Tag() tag.Tag { return tag.ContAmtCurr }

type ContAmtType struct{ quickfix.IntValue }

func (f ContAmtType) Tag() tag.Tag { return tag.ContAmtType }

type ContAmtValue struct{ quickfix.FloatValue }

func (f ContAmtValue) Tag() tag.Tag { return tag.ContAmtValue }

type ContIntRptID struct{ quickfix.StringValue }

func (f ContIntRptID) Tag() tag.Tag { return tag.ContIntRptID }

type ContextPartyID struct{ quickfix.StringValue }

func (f ContextPartyID) Tag() tag.Tag { return tag.ContextPartyID }

type ContextPartyIDSource struct{ quickfix.CharValue }

func (f ContextPartyIDSource) Tag() tag.Tag { return tag.ContextPartyIDSource }

type ContextPartyRole struct{ quickfix.IntValue }

func (f ContextPartyRole) Tag() tag.Tag { return tag.ContextPartyRole }

type ContextPartySubID struct{ quickfix.StringValue }

func (f ContextPartySubID) Tag() tag.Tag { return tag.ContextPartySubID }

type ContextPartySubIDType struct{ quickfix.IntValue }

func (f ContextPartySubIDType) Tag() tag.Tag { return tag.ContextPartySubIDType }

type ContingencyType struct{ quickfix.IntValue }

func (f ContingencyType) Tag() tag.Tag { return tag.ContingencyType }

type ContraBroker struct{ quickfix.StringValue }

func (f ContraBroker) Tag() tag.Tag { return tag.ContraBroker }

type ContraLegRefID struct{ quickfix.StringValue }

func (f ContraLegRefID) Tag() tag.Tag { return tag.ContraLegRefID }

type ContraTradeQty struct{ quickfix.QtyValue }

func (f ContraTradeQty) Tag() tag.Tag { return tag.ContraTradeQty }

type ContraTradeTime struct{ quickfix.UTCTimestampValue }

func (f ContraTradeTime) Tag() tag.Tag { return tag.ContraTradeTime }

type ContraTrader struct{ quickfix.StringValue }

func (f ContraTrader) Tag() tag.Tag { return tag.ContraTrader }

type ContractMultiplier struct{ quickfix.FloatValue }

func (f ContractMultiplier) Tag() tag.Tag { return tag.ContractMultiplier }

type ContractMultiplierUnit struct{ quickfix.IntValue }

func (f ContractMultiplierUnit) Tag() tag.Tag { return tag.ContractMultiplierUnit }

type ContractSettlMonth struct{ quickfix.MonthYearValue }

func (f ContractSettlMonth) Tag() tag.Tag { return tag.ContractSettlMonth }

type ContraryInstructionIndicator struct{ quickfix.BooleanValue }

func (f ContraryInstructionIndicator) Tag() tag.Tag { return tag.ContraryInstructionIndicator }

type CopyMsgIndicator struct{ quickfix.BooleanValue }

func (f CopyMsgIndicator) Tag() tag.Tag { return tag.CopyMsgIndicator }

type CorporateAction struct{ quickfix.MultipleCharValue }

func (f CorporateAction) Tag() tag.Tag { return tag.CorporateAction }

type Country struct{ quickfix.CountryValue }

func (f Country) Tag() tag.Tag { return tag.Country }

type CountryOfIssue struct{ quickfix.CountryValue }

func (f CountryOfIssue) Tag() tag.Tag { return tag.CountryOfIssue }

type CouponPaymentDate struct{ quickfix.LocalMktDateValue }

func (f CouponPaymentDate) Tag() tag.Tag { return tag.CouponPaymentDate }

type CouponRate struct{ quickfix.PercentageValue }

func (f CouponRate) Tag() tag.Tag { return tag.CouponRate }

type CoveredOrUncovered struct{ quickfix.IntValue }

func (f CoveredOrUncovered) Tag() tag.Tag { return tag.CoveredOrUncovered }

type CreditRating struct{ quickfix.StringValue }

func (f CreditRating) Tag() tag.Tag { return tag.CreditRating }

type CrossID struct{ quickfix.StringValue }

func (f CrossID) Tag() tag.Tag { return tag.CrossID }

type CrossPercent struct{ quickfix.PercentageValue }

func (f CrossPercent) Tag() tag.Tag { return tag.CrossPercent }

type CrossPrioritization struct{ quickfix.IntValue }

func (f CrossPrioritization) Tag() tag.Tag { return tag.CrossPrioritization }

type CrossType struct{ quickfix.IntValue }

func (f CrossType) Tag() tag.Tag { return tag.CrossType }

type CstmApplVerID struct{ quickfix.StringValue }

func (f CstmApplVerID) Tag() tag.Tag { return tag.CstmApplVerID }

type CumQty struct{ quickfix.QtyValue }

func (f CumQty) Tag() tag.Tag { return tag.CumQty }

type Currency struct{ quickfix.CurrencyValue }

func (f Currency) Tag() tag.Tag { return tag.Currency }

type CurrencyRatio struct{ quickfix.FloatValue }

func (f CurrencyRatio) Tag() tag.Tag { return tag.CurrencyRatio }

type CustDirectedOrder struct{ quickfix.BooleanValue }

func (f CustDirectedOrder) Tag() tag.Tag { return tag.CustDirectedOrder }

type CustOrderCapacity struct{ quickfix.IntValue }

func (f CustOrderCapacity) Tag() tag.Tag { return tag.CustOrderCapacity }

type CustOrderHandlingInst struct{ quickfix.MultipleStringValue }

func (f CustOrderHandlingInst) Tag() tag.Tag { return tag.CustOrderHandlingInst }

type CustomerOrFirm struct{ quickfix.IntValue }

func (f CustomerOrFirm) Tag() tag.Tag { return tag.CustomerOrFirm }

type CxlQty struct{ quickfix.QtyValue }

func (f CxlQty) Tag() tag.Tag { return tag.CxlQty }

type CxlRejReason struct{ quickfix.IntValue }

func (f CxlRejReason) Tag() tag.Tag { return tag.CxlRejReason }

type CxlRejResponseTo struct{ quickfix.CharValue }

func (f CxlRejResponseTo) Tag() tag.Tag { return tag.CxlRejResponseTo }

type CxlType struct{ quickfix.CharValue }

func (f CxlType) Tag() tag.Tag { return tag.CxlType }

type DKReason struct{ quickfix.CharValue }

func (f DKReason) Tag() tag.Tag { return tag.DKReason }

type DateOfBirth struct{ quickfix.LocalMktDateValue }

func (f DateOfBirth) Tag() tag.Tag { return tag.DateOfBirth }

type DatedDate struct{ quickfix.LocalMktDateValue }

func (f DatedDate) Tag() tag.Tag { return tag.DatedDate }

type DayAvgPx struct{ quickfix.PriceValue }

func (f DayAvgPx) Tag() tag.Tag { return tag.DayAvgPx }

type DayBookingInst struct{ quickfix.CharValue }

func (f DayBookingInst) Tag() tag.Tag { return tag.DayBookingInst }

type DayCumQty struct{ quickfix.QtyValue }

func (f DayCumQty) Tag() tag.Tag { return tag.DayCumQty }

type DayOrderQty struct{ quickfix.QtyValue }

func (f DayOrderQty) Tag() tag.Tag { return tag.DayOrderQty }

type DealingCapacity struct{ quickfix.CharValue }

func (f DealingCapacity) Tag() tag.Tag { return tag.DealingCapacity }

type DefBidSize struct{ quickfix.QtyValue }

func (f DefBidSize) Tag() tag.Tag { return tag.DefBidSize }

type DefOfferSize struct{ quickfix.QtyValue }

func (f DefOfferSize) Tag() tag.Tag { return tag.DefOfferSize }

type DefaultApplExtID struct{ quickfix.IntValue }

func (f DefaultApplExtID) Tag() tag.Tag { return tag.DefaultApplExtID }

type DefaultApplVerID struct{ quickfix.StringValue }

func (f DefaultApplVerID) Tag() tag.Tag { return tag.DefaultApplVerID }

type DefaultCstmApplVerID struct{ quickfix.StringValue }

func (f DefaultCstmApplVerID) Tag() tag.Tag { return tag.DefaultCstmApplVerID }

type DefaultVerIndicator struct{ quickfix.BooleanValue }

func (f DefaultVerIndicator) Tag() tag.Tag { return tag.DefaultVerIndicator }

type DeleteReason struct{ quickfix.CharValue }

func (f DeleteReason) Tag() tag.Tag { return tag.DeleteReason }

type DeliverToCompID struct{ quickfix.StringValue }

func (f DeliverToCompID) Tag() tag.Tag { return tag.DeliverToCompID }

type DeliverToLocationID struct{ quickfix.StringValue }

func (f DeliverToLocationID) Tag() tag.Tag { return tag.DeliverToLocationID }

type DeliverToSubID struct{ quickfix.StringValue }

func (f DeliverToSubID) Tag() tag.Tag { return tag.DeliverToSubID }

type DeliveryDate struct{ quickfix.LocalMktDateValue }

func (f DeliveryDate) Tag() tag.Tag { return tag.DeliveryDate }

type DeliveryForm struct{ quickfix.IntValue }

func (f DeliveryForm) Tag() tag.Tag { return tag.DeliveryForm }

type DeliveryType struct{ quickfix.IntValue }

func (f DeliveryType) Tag() tag.Tag { return tag.DeliveryType }

type DerivFlexProductEligibilityIndicator struct{ quickfix.BooleanValue }

func (f DerivFlexProductEligibilityIndicator) Tag() tag.Tag {
	return tag.DerivFlexProductEligibilityIndicator
}

type DerivativeCFICode struct{ quickfix.StringValue }

func (f DerivativeCFICode) Tag() tag.Tag { return tag.DerivativeCFICode }

type DerivativeCapPrice struct{ quickfix.PriceValue }

func (f DerivativeCapPrice) Tag() tag.Tag { return tag.DerivativeCapPrice }

type DerivativeContractMultiplier struct{ quickfix.FloatValue }

func (f DerivativeContractMultiplier) Tag() tag.Tag { return tag.DerivativeContractMultiplier }

type DerivativeContractMultiplierUnit struct{ quickfix.IntValue }

func (f DerivativeContractMultiplierUnit) Tag() tag.Tag { return tag.DerivativeContractMultiplierUnit }

type DerivativeContractSettlMonth struct{ quickfix.MonthYearValue }

func (f DerivativeContractSettlMonth) Tag() tag.Tag { return tag.DerivativeContractSettlMonth }

type DerivativeCountryOfIssue struct{ quickfix.CountryValue }

func (f DerivativeCountryOfIssue) Tag() tag.Tag { return tag.DerivativeCountryOfIssue }

type DerivativeEncodedIssuer struct{ quickfix.DataValue }

func (f DerivativeEncodedIssuer) Tag() tag.Tag { return tag.DerivativeEncodedIssuer }

type DerivativeEncodedIssuerLen struct{ quickfix.LengthValue }

func (f DerivativeEncodedIssuerLen) Tag() tag.Tag { return tag.DerivativeEncodedIssuerLen }

type DerivativeEncodedSecurityDesc struct{ quickfix.DataValue }

func (f DerivativeEncodedSecurityDesc) Tag() tag.Tag { return tag.DerivativeEncodedSecurityDesc }

type DerivativeEncodedSecurityDescLen struct{ quickfix.LengthValue }

func (f DerivativeEncodedSecurityDescLen) Tag() tag.Tag { return tag.DerivativeEncodedSecurityDescLen }

type DerivativeEventDate struct{ quickfix.LocalMktDateValue }

func (f DerivativeEventDate) Tag() tag.Tag { return tag.DerivativeEventDate }

type DerivativeEventPx struct{ quickfix.PriceValue }

func (f DerivativeEventPx) Tag() tag.Tag { return tag.DerivativeEventPx }

type DerivativeEventText struct{ quickfix.StringValue }

func (f DerivativeEventText) Tag() tag.Tag { return tag.DerivativeEventText }

type DerivativeEventTime struct{ quickfix.UTCTimestampValue }

func (f DerivativeEventTime) Tag() tag.Tag { return tag.DerivativeEventTime }

type DerivativeEventType struct{ quickfix.IntValue }

func (f DerivativeEventType) Tag() tag.Tag { return tag.DerivativeEventType }

type DerivativeExerciseStyle struct{ quickfix.CharValue }

func (f DerivativeExerciseStyle) Tag() tag.Tag { return tag.DerivativeExerciseStyle }

type DerivativeFloorPrice struct{ quickfix.PriceValue }

func (f DerivativeFloorPrice) Tag() tag.Tag { return tag.DerivativeFloorPrice }

type DerivativeFlowScheduleType struct{ quickfix.IntValue }

func (f DerivativeFlowScheduleType) Tag() tag.Tag { return tag.DerivativeFlowScheduleType }

type DerivativeFuturesValuationMethod struct{ quickfix.StringValue }

func (f DerivativeFuturesValuationMethod) Tag() tag.Tag { return tag.DerivativeFuturesValuationMethod }

type DerivativeInstrAttribType struct{ quickfix.IntValue }

func (f DerivativeInstrAttribType) Tag() tag.Tag { return tag.DerivativeInstrAttribType }

type DerivativeInstrAttribValue struct{ quickfix.StringValue }

func (f DerivativeInstrAttribValue) Tag() tag.Tag { return tag.DerivativeInstrAttribValue }

type DerivativeInstrRegistry struct{ quickfix.StringValue }

func (f DerivativeInstrRegistry) Tag() tag.Tag { return tag.DerivativeInstrRegistry }

type DerivativeInstrmtAssignmentMethod struct{ quickfix.CharValue }

func (f DerivativeInstrmtAssignmentMethod) Tag() tag.Tag { return tag.DerivativeInstrmtAssignmentMethod }

type DerivativeInstrumentPartyID struct{ quickfix.StringValue }

func (f DerivativeInstrumentPartyID) Tag() tag.Tag { return tag.DerivativeInstrumentPartyID }

type DerivativeInstrumentPartyIDSource struct{ quickfix.StringValue }

func (f DerivativeInstrumentPartyIDSource) Tag() tag.Tag { return tag.DerivativeInstrumentPartyIDSource }

type DerivativeInstrumentPartyRole struct{ quickfix.IntValue }

func (f DerivativeInstrumentPartyRole) Tag() tag.Tag { return tag.DerivativeInstrumentPartyRole }

type DerivativeInstrumentPartySubID struct{ quickfix.StringValue }

func (f DerivativeInstrumentPartySubID) Tag() tag.Tag { return tag.DerivativeInstrumentPartySubID }

type DerivativeInstrumentPartySubIDType struct{ quickfix.IntValue }

func (f DerivativeInstrumentPartySubIDType) Tag() tag.Tag {
	return tag.DerivativeInstrumentPartySubIDType
}

type DerivativeIssueDate struct{ quickfix.LocalMktDateValue }

func (f DerivativeIssueDate) Tag() tag.Tag { return tag.DerivativeIssueDate }

type DerivativeIssuer struct{ quickfix.StringValue }

func (f DerivativeIssuer) Tag() tag.Tag { return tag.DerivativeIssuer }

type DerivativeListMethod struct{ quickfix.IntValue }

func (f DerivativeListMethod) Tag() tag.Tag { return tag.DerivativeListMethod }

type DerivativeLocaleOfIssue struct{ quickfix.StringValue }

func (f DerivativeLocaleOfIssue) Tag() tag.Tag { return tag.DerivativeLocaleOfIssue }

type DerivativeMaturityDate struct{ quickfix.LocalMktDateValue }

func (f DerivativeMaturityDate) Tag() tag.Tag { return tag.DerivativeMaturityDate }

type DerivativeMaturityMonthYear struct{ quickfix.MonthYearValue }

func (f DerivativeMaturityMonthYear) Tag() tag.Tag { return tag.DerivativeMaturityMonthYear }

type DerivativeMaturityTime struct{ quickfix.TZTimeOnlyValue }

func (f DerivativeMaturityTime) Tag() tag.Tag { return tag.DerivativeMaturityTime }

type DerivativeMinPriceIncrement struct{ quickfix.FloatValue }

func (f DerivativeMinPriceIncrement) Tag() tag.Tag { return tag.DerivativeMinPriceIncrement }

type DerivativeMinPriceIncrementAmount struct{ quickfix.AmtValue }

func (f DerivativeMinPriceIncrementAmount) Tag() tag.Tag { return tag.DerivativeMinPriceIncrementAmount }

type DerivativeNTPositionLimit struct{ quickfix.IntValue }

func (f DerivativeNTPositionLimit) Tag() tag.Tag { return tag.DerivativeNTPositionLimit }

type DerivativeOptAttribute struct{ quickfix.CharValue }

func (f DerivativeOptAttribute) Tag() tag.Tag { return tag.DerivativeOptAttribute }

type DerivativeOptPayAmount struct{ quickfix.AmtValue }

func (f DerivativeOptPayAmount) Tag() tag.Tag { return tag.DerivativeOptPayAmount }

type DerivativePositionLimit struct{ quickfix.IntValue }

func (f DerivativePositionLimit) Tag() tag.Tag { return tag.DerivativePositionLimit }

type DerivativePriceQuoteMethod struct{ quickfix.StringValue }

func (f DerivativePriceQuoteMethod) Tag() tag.Tag { return tag.DerivativePriceQuoteMethod }

type DerivativePriceUnitOfMeasure struct{ quickfix.StringValue }

func (f DerivativePriceUnitOfMeasure) Tag() tag.Tag { return tag.DerivativePriceUnitOfMeasure }

type DerivativePriceUnitOfMeasureQty struct{ quickfix.QtyValue }

func (f DerivativePriceUnitOfMeasureQty) Tag() tag.Tag { return tag.DerivativePriceUnitOfMeasureQty }

type DerivativeProduct struct{ quickfix.IntValue }

func (f DerivativeProduct) Tag() tag.Tag { return tag.DerivativeProduct }

type DerivativeProductComplex struct{ quickfix.StringValue }

func (f DerivativeProductComplex) Tag() tag.Tag { return tag.DerivativeProductComplex }

type DerivativePutOrCall struct{ quickfix.IntValue }

func (f DerivativePutOrCall) Tag() tag.Tag { return tag.DerivativePutOrCall }

type DerivativeSecurityAltID struct{ quickfix.StringValue }

func (f DerivativeSecurityAltID) Tag() tag.Tag { return tag.DerivativeSecurityAltID }

type DerivativeSecurityAltIDSource struct{ quickfix.StringValue }

func (f DerivativeSecurityAltIDSource) Tag() tag.Tag { return tag.DerivativeSecurityAltIDSource }

type DerivativeSecurityDesc struct{ quickfix.StringValue }

func (f DerivativeSecurityDesc) Tag() tag.Tag { return tag.DerivativeSecurityDesc }

type DerivativeSecurityExchange struct{ quickfix.ExchangeValue }

func (f DerivativeSecurityExchange) Tag() tag.Tag { return tag.DerivativeSecurityExchange }

type DerivativeSecurityGroup struct{ quickfix.StringValue }

func (f DerivativeSecurityGroup) Tag() tag.Tag { return tag.DerivativeSecurityGroup }

type DerivativeSecurityID struct{ quickfix.StringValue }

func (f DerivativeSecurityID) Tag() tag.Tag { return tag.DerivativeSecurityID }

type DerivativeSecurityIDSource struct{ quickfix.StringValue }

func (f DerivativeSecurityIDSource) Tag() tag.Tag { return tag.DerivativeSecurityIDSource }

type DerivativeSecurityListRequestType struct{ quickfix.IntValue }

func (f DerivativeSecurityListRequestType) Tag() tag.Tag { return tag.DerivativeSecurityListRequestType }

type DerivativeSecurityStatus struct{ quickfix.StringValue }

func (f DerivativeSecurityStatus) Tag() tag.Tag { return tag.DerivativeSecurityStatus }

type DerivativeSecuritySubType struct{ quickfix.StringValue }

func (f DerivativeSecuritySubType) Tag() tag.Tag { return tag.DerivativeSecuritySubType }

type DerivativeSecurityType struct{ quickfix.StringValue }

func (f DerivativeSecurityType) Tag() tag.Tag { return tag.DerivativeSecurityType }

type DerivativeSecurityXML struct{ quickfix.DataValue }

func (f DerivativeSecurityXML) Tag() tag.Tag { return tag.DerivativeSecurityXML }

type DerivativeSecurityXMLLen struct{ quickfix.LengthValue }

func (f DerivativeSecurityXMLLen) Tag() tag.Tag { return tag.DerivativeSecurityXMLLen }

type DerivativeSecurityXMLSchema struct{ quickfix.StringValue }

func (f DerivativeSecurityXMLSchema) Tag() tag.Tag { return tag.DerivativeSecurityXMLSchema }

type DerivativeSettlMethod struct{ quickfix.CharValue }

func (f DerivativeSettlMethod) Tag() tag.Tag { return tag.DerivativeSettlMethod }

type DerivativeSettleOnOpenFlag struct{ quickfix.StringValue }

func (f DerivativeSettleOnOpenFlag) Tag() tag.Tag { return tag.DerivativeSettleOnOpenFlag }

type DerivativeStateOrProvinceOfIssue struct{ quickfix.StringValue }

func (f DerivativeStateOrProvinceOfIssue) Tag() tag.Tag { return tag.DerivativeStateOrProvinceOfIssue }

type DerivativeStrikeCurrency struct{ quickfix.CurrencyValue }

func (f DerivativeStrikeCurrency) Tag() tag.Tag { return tag.DerivativeStrikeCurrency }

type DerivativeStrikeMultiplier struct{ quickfix.FloatValue }

func (f DerivativeStrikeMultiplier) Tag() tag.Tag { return tag.DerivativeStrikeMultiplier }

type DerivativeStrikePrice struct{ quickfix.PriceValue }

func (f DerivativeStrikePrice) Tag() tag.Tag { return tag.DerivativeStrikePrice }

type DerivativeStrikeValue struct{ quickfix.FloatValue }

func (f DerivativeStrikeValue) Tag() tag.Tag { return tag.DerivativeStrikeValue }

type DerivativeSymbol struct{ quickfix.StringValue }

func (f DerivativeSymbol) Tag() tag.Tag { return tag.DerivativeSymbol }

type DerivativeSymbolSfx struct{ quickfix.StringValue }

func (f DerivativeSymbolSfx) Tag() tag.Tag { return tag.DerivativeSymbolSfx }

type DerivativeTimeUnit struct{ quickfix.StringValue }

func (f DerivativeTimeUnit) Tag() tag.Tag { return tag.DerivativeTimeUnit }

type DerivativeUnitOfMeasure struct{ quickfix.StringValue }

func (f DerivativeUnitOfMeasure) Tag() tag.Tag { return tag.DerivativeUnitOfMeasure }

type DerivativeUnitOfMeasureQty struct{ quickfix.QtyValue }

func (f DerivativeUnitOfMeasureQty) Tag() tag.Tag { return tag.DerivativeUnitOfMeasureQty }

type DerivativeValuationMethod struct{ quickfix.StringValue }

func (f DerivativeValuationMethod) Tag() tag.Tag { return tag.DerivativeValuationMethod }

type Designation struct{ quickfix.StringValue }

func (f Designation) Tag() tag.Tag { return tag.Designation }

type DeskID struct{ quickfix.StringValue }

func (f DeskID) Tag() tag.Tag { return tag.DeskID }

type DeskOrderHandlingInst struct{ quickfix.MultipleStringValue }

func (f DeskOrderHandlingInst) Tag() tag.Tag { return tag.DeskOrderHandlingInst }

type DeskType struct{ quickfix.StringValue }

func (f DeskType) Tag() tag.Tag { return tag.DeskType }

type DeskTypeSource struct{ quickfix.IntValue }

func (f DeskTypeSource) Tag() tag.Tag { return tag.DeskTypeSource }

type DetachmentPoint struct{ quickfix.PercentageValue }

func (f DetachmentPoint) Tag() tag.Tag { return tag.DetachmentPoint }

type DiscretionInst struct{ quickfix.CharValue }

func (f DiscretionInst) Tag() tag.Tag { return tag.DiscretionInst }

type DiscretionLimitType struct{ quickfix.IntValue }

func (f DiscretionLimitType) Tag() tag.Tag { return tag.DiscretionLimitType }

type DiscretionMoveType struct{ quickfix.IntValue }

func (f DiscretionMoveType) Tag() tag.Tag { return tag.DiscretionMoveType }

type DiscretionOffset struct{ quickfix.PriceOffsetValue }

func (f DiscretionOffset) Tag() tag.Tag { return tag.DiscretionOffset }

type DiscretionOffsetType struct{ quickfix.IntValue }

func (f DiscretionOffsetType) Tag() tag.Tag { return tag.DiscretionOffsetType }

type DiscretionOffsetValue struct{ quickfix.FloatValue }

func (f DiscretionOffsetValue) Tag() tag.Tag { return tag.DiscretionOffsetValue }

type DiscretionPrice struct{ quickfix.PriceValue }

func (f DiscretionPrice) Tag() tag.Tag { return tag.DiscretionPrice }

type DiscretionRoundDirection struct{ quickfix.IntValue }

func (f DiscretionRoundDirection) Tag() tag.Tag { return tag.DiscretionRoundDirection }

type DiscretionScope struct{ quickfix.IntValue }

func (f DiscretionScope) Tag() tag.Tag { return tag.DiscretionScope }

type DisplayHighQty struct{ quickfix.QtyValue }

func (f DisplayHighQty) Tag() tag.Tag { return tag.DisplayHighQty }

type DisplayLowQty struct{ quickfix.QtyValue }

func (f DisplayLowQty) Tag() tag.Tag { return tag.DisplayLowQty }

type DisplayMethod struct{ quickfix.CharValue }

func (f DisplayMethod) Tag() tag.Tag { return tag.DisplayMethod }

type DisplayMinIncr struct{ quickfix.QtyValue }

func (f DisplayMinIncr) Tag() tag.Tag { return tag.DisplayMinIncr }

type DisplayQty struct{ quickfix.QtyValue }

func (f DisplayQty) Tag() tag.Tag { return tag.DisplayQty }

type DisplayWhen struct{ quickfix.CharValue }

func (f DisplayWhen) Tag() tag.Tag { return tag.DisplayWhen }

type DistribPaymentMethod struct{ quickfix.IntValue }

func (f DistribPaymentMethod) Tag() tag.Tag { return tag.DistribPaymentMethod }

type DistribPercentage struct{ quickfix.PercentageValue }

func (f DistribPercentage) Tag() tag.Tag { return tag.DistribPercentage }

type DividendYield struct{ quickfix.PercentageValue }

func (f DividendYield) Tag() tag.Tag { return tag.DividendYield }

type DlvyInst struct{ quickfix.StringValue }

func (f DlvyInst) Tag() tag.Tag { return tag.DlvyInst }

type DlvyInstType struct{ quickfix.CharValue }

func (f DlvyInstType) Tag() tag.Tag { return tag.DlvyInstType }

type DueToRelated struct{ quickfix.BooleanValue }

func (f DueToRelated) Tag() tag.Tag { return tag.DueToRelated }

type EFPTrackingError struct{ quickfix.PercentageValue }

func (f EFPTrackingError) Tag() tag.Tag { return tag.EFPTrackingError }

type EffectiveTime struct{ quickfix.UTCTimestampValue }

func (f EffectiveTime) Tag() tag.Tag { return tag.EffectiveTime }

type EmailThreadID struct{ quickfix.StringValue }

func (f EmailThreadID) Tag() tag.Tag { return tag.EmailThreadID }

type EmailType struct{ quickfix.CharValue }

func (f EmailType) Tag() tag.Tag { return tag.EmailType }

type EncodedAllocText struct{ quickfix.DataValue }

func (f EncodedAllocText) Tag() tag.Tag { return tag.EncodedAllocText }

type EncodedAllocTextLen struct{ quickfix.LengthValue }

func (f EncodedAllocTextLen) Tag() tag.Tag { return tag.EncodedAllocTextLen }

type EncodedHeadline struct{ quickfix.DataValue }

func (f EncodedHeadline) Tag() tag.Tag { return tag.EncodedHeadline }

type EncodedHeadlineLen struct{ quickfix.LengthValue }

func (f EncodedHeadlineLen) Tag() tag.Tag { return tag.EncodedHeadlineLen }

type EncodedIssuer struct{ quickfix.DataValue }

func (f EncodedIssuer) Tag() tag.Tag { return tag.EncodedIssuer }

type EncodedIssuerLen struct{ quickfix.LengthValue }

func (f EncodedIssuerLen) Tag() tag.Tag { return tag.EncodedIssuerLen }

type EncodedLegIssuer struct{ quickfix.DataValue }

func (f EncodedLegIssuer) Tag() tag.Tag { return tag.EncodedLegIssuer }

type EncodedLegIssuerLen struct{ quickfix.LengthValue }

func (f EncodedLegIssuerLen) Tag() tag.Tag { return tag.EncodedLegIssuerLen }

type EncodedLegSecurityDesc struct{ quickfix.DataValue }

func (f EncodedLegSecurityDesc) Tag() tag.Tag { return tag.EncodedLegSecurityDesc }

type EncodedLegSecurityDescLen struct{ quickfix.LengthValue }

func (f EncodedLegSecurityDescLen) Tag() tag.Tag { return tag.EncodedLegSecurityDescLen }

type EncodedListExecInst struct{ quickfix.DataValue }

func (f EncodedListExecInst) Tag() tag.Tag { return tag.EncodedListExecInst }

type EncodedListExecInstLen struct{ quickfix.LengthValue }

func (f EncodedListExecInstLen) Tag() tag.Tag { return tag.EncodedListExecInstLen }

type EncodedListStatusText struct{ quickfix.DataValue }

func (f EncodedListStatusText) Tag() tag.Tag { return tag.EncodedListStatusText }

type EncodedListStatusTextLen struct{ quickfix.LengthValue }

func (f EncodedListStatusTextLen) Tag() tag.Tag { return tag.EncodedListStatusTextLen }

type EncodedMktSegmDesc struct{ quickfix.DataValue }

func (f EncodedMktSegmDesc) Tag() tag.Tag { return tag.EncodedMktSegmDesc }

type EncodedMktSegmDescLen struct{ quickfix.LengthValue }

func (f EncodedMktSegmDescLen) Tag() tag.Tag { return tag.EncodedMktSegmDescLen }

type EncodedSecurityDesc struct{ quickfix.DataValue }

func (f EncodedSecurityDesc) Tag() tag.Tag { return tag.EncodedSecurityDesc }

type EncodedSecurityDescLen struct{ quickfix.LengthValue }

func (f EncodedSecurityDescLen) Tag() tag.Tag { return tag.EncodedSecurityDescLen }

type EncodedSecurityListDesc struct{ quickfix.DataValue }

func (f EncodedSecurityListDesc) Tag() tag.Tag { return tag.EncodedSecurityListDesc }

type EncodedSecurityListDescLen struct{ quickfix.LengthValue }

func (f EncodedSecurityListDescLen) Tag() tag.Tag { return tag.EncodedSecurityListDescLen }

type EncodedSubject struct{ quickfix.DataValue }

func (f EncodedSubject) Tag() tag.Tag { return tag.EncodedSubject }

type EncodedSubjectLen struct{ quickfix.LengthValue }

func (f EncodedSubjectLen) Tag() tag.Tag { return tag.EncodedSubjectLen }

type EncodedSymbol struct{ quickfix.DataValue }

func (f EncodedSymbol) Tag() tag.Tag { return tag.EncodedSymbol }

type EncodedSymbolLen struct{ quickfix.LengthValue }

func (f EncodedSymbolLen) Tag() tag.Tag { return tag.EncodedSymbolLen }

type EncodedText struct{ quickfix.DataValue }

func (f EncodedText) Tag() tag.Tag { return tag.EncodedText }

type EncodedTextLen struct{ quickfix.LengthValue }

func (f EncodedTextLen) Tag() tag.Tag { return tag.EncodedTextLen }

type EncodedUnderlyingIssuer struct{ quickfix.DataValue }

func (f EncodedUnderlyingIssuer) Tag() tag.Tag { return tag.EncodedUnderlyingIssuer }

type EncodedUnderlyingIssuerLen struct{ quickfix.LengthValue }

func (f EncodedUnderlyingIssuerLen) Tag() tag.Tag { return tag.EncodedUnderlyingIssuerLen }

type EncodedUnderlyingSecurityDesc struct{ quickfix.DataValue }

func (f EncodedUnderlyingSecurityDesc) Tag() tag.Tag { return tag.EncodedUnderlyingSecurityDesc }

type EncodedUnderlyingSecurityDescLen struct{ quickfix.LengthValue }

func (f EncodedUnderlyingSecurityDescLen) Tag() tag.Tag { return tag.EncodedUnderlyingSecurityDescLen }

type EncryptMethod struct{ quickfix.IntValue }

func (f EncryptMethod) Tag() tag.Tag { return tag.EncryptMethod }

type EncryptedNewPassword struct{ quickfix.DataValue }

func (f EncryptedNewPassword) Tag() tag.Tag { return tag.EncryptedNewPassword }

type EncryptedNewPasswordLen struct{ quickfix.LengthValue }

func (f EncryptedNewPasswordLen) Tag() tag.Tag { return tag.EncryptedNewPasswordLen }

type EncryptedPassword struct{ quickfix.DataValue }

func (f EncryptedPassword) Tag() tag.Tag { return tag.EncryptedPassword }

type EncryptedPasswordLen struct{ quickfix.LengthValue }

func (f EncryptedPasswordLen) Tag() tag.Tag { return tag.EncryptedPasswordLen }

type EncryptedPasswordMethod struct{ quickfix.IntValue }

func (f EncryptedPasswordMethod) Tag() tag.Tag { return tag.EncryptedPasswordMethod }

type EndAccruedInterestAmt struct{ quickfix.AmtValue }

func (f EndAccruedInterestAmt) Tag() tag.Tag { return tag.EndAccruedInterestAmt }

type EndCash struct{ quickfix.AmtValue }

func (f EndCash) Tag() tag.Tag { return tag.EndCash }

type EndDate struct{ quickfix.LocalMktDateValue }

func (f EndDate) Tag() tag.Tag { return tag.EndDate }

type EndMaturityMonthYear struct{ quickfix.MonthYearValue }

func (f EndMaturityMonthYear) Tag() tag.Tag { return tag.EndMaturityMonthYear }

type EndSeqNo struct{ quickfix.SeqNumValue }

func (f EndSeqNo) Tag() tag.Tag { return tag.EndSeqNo }

type EndStrikePxRange struct{ quickfix.PriceValue }

func (f EndStrikePxRange) Tag() tag.Tag { return tag.EndStrikePxRange }

type EndTickPriceRange struct{ quickfix.PriceValue }

func (f EndTickPriceRange) Tag() tag.Tag { return tag.EndTickPriceRange }

type EventDate struct{ quickfix.LocalMktDateValue }

func (f EventDate) Tag() tag.Tag { return tag.EventDate }

type EventPx struct{ quickfix.PriceValue }

func (f EventPx) Tag() tag.Tag { return tag.EventPx }

type EventText struct{ quickfix.StringValue }

func (f EventText) Tag() tag.Tag { return tag.EventText }

type EventTime struct{ quickfix.UTCTimestampValue }

func (f EventTime) Tag() tag.Tag { return tag.EventTime }

type EventType struct{ quickfix.IntValue }

func (f EventType) Tag() tag.Tag { return tag.EventType }

type ExDate struct{ quickfix.LocalMktDateValue }

func (f ExDate) Tag() tag.Tag { return tag.ExDate }

type ExDestination struct{ quickfix.ExchangeValue }

func (f ExDestination) Tag() tag.Tag { return tag.ExDestination }

type ExDestinationIDSource struct{ quickfix.CharValue }

func (f ExDestinationIDSource) Tag() tag.Tag { return tag.ExDestinationIDSource }

type ExchangeForPhysical struct{ quickfix.BooleanValue }

func (f ExchangeForPhysical) Tag() tag.Tag { return tag.ExchangeForPhysical }

type ExchangeRule struct{ quickfix.StringValue }

func (f ExchangeRule) Tag() tag.Tag { return tag.ExchangeRule }

type ExchangeSpecialInstructions struct{ quickfix.StringValue }

func (f ExchangeSpecialInstructions) Tag() tag.Tag { return tag.ExchangeSpecialInstructions }

type ExecAckStatus struct{ quickfix.CharValue }

func (f ExecAckStatus) Tag() tag.Tag { return tag.ExecAckStatus }

type ExecBroker struct{ quickfix.StringValue }

func (f ExecBroker) Tag() tag.Tag { return tag.ExecBroker }

type ExecID struct{ quickfix.StringValue }

func (f ExecID) Tag() tag.Tag { return tag.ExecID }

type ExecInst struct{ quickfix.MultipleCharValue }

func (f ExecInst) Tag() tag.Tag { return tag.ExecInst }

type ExecInstValue struct{ quickfix.CharValue }

func (f ExecInstValue) Tag() tag.Tag { return tag.ExecInstValue }

type ExecPriceAdjustment struct{ quickfix.FloatValue }

func (f ExecPriceAdjustment) Tag() tag.Tag { return tag.ExecPriceAdjustment }

type ExecPriceType struct{ quickfix.CharValue }

func (f ExecPriceType) Tag() tag.Tag { return tag.ExecPriceType }

type ExecRefID struct{ quickfix.StringValue }

func (f ExecRefID) Tag() tag.Tag { return tag.ExecRefID }

type ExecRestatementReason struct{ quickfix.IntValue }

func (f ExecRestatementReason) Tag() tag.Tag { return tag.ExecRestatementReason }

type ExecTransType struct{ quickfix.CharValue }

func (f ExecTransType) Tag() tag.Tag { return tag.ExecTransType }

type ExecType struct{ quickfix.CharValue }

func (f ExecType) Tag() tag.Tag { return tag.ExecType }

type ExecValuationPoint struct{ quickfix.UTCTimestampValue }

func (f ExecValuationPoint) Tag() tag.Tag { return tag.ExecValuationPoint }

type ExerciseMethod struct{ quickfix.CharValue }

func (f ExerciseMethod) Tag() tag.Tag { return tag.ExerciseMethod }

type ExerciseStyle struct{ quickfix.IntValue }

func (f ExerciseStyle) Tag() tag.Tag { return tag.ExerciseStyle }

type ExpQty struct{ quickfix.QtyValue }

func (f ExpQty) Tag() tag.Tag { return tag.ExpQty }

type ExpType struct{ quickfix.IntValue }

func (f ExpType) Tag() tag.Tag { return tag.ExpType }

type ExpirationCycle struct{ quickfix.IntValue }

func (f ExpirationCycle) Tag() tag.Tag { return tag.ExpirationCycle }

type ExpirationQtyType struct{ quickfix.IntValue }

func (f ExpirationQtyType) Tag() tag.Tag { return tag.ExpirationQtyType }

type ExpireDate struct{ quickfix.LocalMktDateValue }

func (f ExpireDate) Tag() tag.Tag { return tag.ExpireDate }

type ExpireTime struct{ quickfix.UTCTimestampValue }

func (f ExpireTime) Tag() tag.Tag { return tag.ExpireTime }

type Factor struct{ quickfix.FloatValue }

func (f Factor) Tag() tag.Tag { return tag.Factor }

type FairValue struct{ quickfix.AmtValue }

func (f FairValue) Tag() tag.Tag { return tag.FairValue }

type FeeMultiplier struct{ quickfix.FloatValue }

func (f FeeMultiplier) Tag() tag.Tag { return tag.FeeMultiplier }

type FillExecID struct{ quickfix.StringValue }

func (f FillExecID) Tag() tag.Tag { return tag.FillExecID }

type FillLiquidityInd struct{ quickfix.IntValue }

func (f FillLiquidityInd) Tag() tag.Tag { return tag.FillLiquidityInd }

type FillPx struct{ quickfix.PriceValue }

func (f FillPx) Tag() tag.Tag { return tag.FillPx }

type FillQty struct{ quickfix.QtyValue }

func (f FillQty) Tag() tag.Tag { return tag.FillQty }

type FinancialStatus struct{ quickfix.MultipleCharValue }

func (f FinancialStatus) Tag() tag.Tag { return tag.FinancialStatus }

type FirmTradeID struct{ quickfix.StringValue }

func (f FirmTradeID) Tag() tag.Tag { return tag.FirmTradeID }

type FirstPx struct{ quickfix.PriceValue }

func (f FirstPx) Tag() tag.Tag { return tag.FirstPx }

type FlexProductEligibilityIndicator struct{ quickfix.BooleanValue }

func (f FlexProductEligibilityIndicator) Tag() tag.Tag { return tag.FlexProductEligibilityIndicator }

type FlexibleIndicator struct{ quickfix.BooleanValue }

func (f FlexibleIndicator) Tag() tag.Tag { return tag.FlexibleIndicator }

type FloorPrice struct{ quickfix.PriceValue }

func (f FloorPrice) Tag() tag.Tag { return tag.FloorPrice }

type FlowScheduleType struct{ quickfix.IntValue }

func (f FlowScheduleType) Tag() tag.Tag { return tag.FlowScheduleType }

type ForexReq struct{ quickfix.BooleanValue }

func (f ForexReq) Tag() tag.Tag { return tag.ForexReq }

type FundRenewWaiv struct{ quickfix.CharValue }

func (f FundRenewWaiv) Tag() tag.Tag { return tag.FundRenewWaiv }

type FutSettDate struct{ quickfix.LocalMktDateValue }

func (f FutSettDate) Tag() tag.Tag { return tag.FutSettDate }

type FutSettDate2 struct{ quickfix.LocalMktDateValue }

func (f FutSettDate2) Tag() tag.Tag { return tag.FutSettDate2 }

type FuturesValuationMethod struct{ quickfix.StringValue }

func (f FuturesValuationMethod) Tag() tag.Tag { return tag.FuturesValuationMethod }

type GTBookingInst struct{ quickfix.IntValue }

func (f GTBookingInst) Tag() tag.Tag { return tag.GTBookingInst }

type GapFillFlag struct{ quickfix.BooleanValue }

func (f GapFillFlag) Tag() tag.Tag { return tag.GapFillFlag }

type GrossTradeAmt struct{ quickfix.AmtValue }

func (f GrossTradeAmt) Tag() tag.Tag { return tag.GrossTradeAmt }

type HaltReasonChar struct{ quickfix.CharValue }

func (f HaltReasonChar) Tag() tag.Tag { return tag.HaltReasonChar }

type HaltReasonInt struct{ quickfix.IntValue }

func (f HaltReasonInt) Tag() tag.Tag { return tag.HaltReasonInt }

type HandlInst struct{ quickfix.CharValue }

func (f HandlInst) Tag() tag.Tag { return tag.HandlInst }

type Headline struct{ quickfix.StringValue }

func (f Headline) Tag() tag.Tag { return tag.Headline }

type HeartBtInt struct{ quickfix.IntValue }

func (f HeartBtInt) Tag() tag.Tag { return tag.HeartBtInt }

type HighLimitPrice struct{ quickfix.PriceValue }

func (f HighLimitPrice) Tag() tag.Tag { return tag.HighLimitPrice }

type HighPx struct{ quickfix.PriceValue }

func (f HighPx) Tag() tag.Tag { return tag.HighPx }

type HopCompID struct{ quickfix.StringValue }

func (f HopCompID) Tag() tag.Tag { return tag.HopCompID }

type HopRefID struct{ quickfix.SeqNumValue }

func (f HopRefID) Tag() tag.Tag { return tag.HopRefID }

type HopSendingTime struct{ quickfix.UTCTimestampValue }

func (f HopSendingTime) Tag() tag.Tag { return tag.HopSendingTime }

type HostCrossID struct{ quickfix.StringValue }

func (f HostCrossID) Tag() tag.Tag { return tag.HostCrossID }

type IDSource struct{ quickfix.StringValue }

func (f IDSource) Tag() tag.Tag { return tag.IDSource }

type IOIID struct{ quickfix.StringValue }

func (f IOIID) Tag() tag.Tag { return tag.IOIID }

type IOINaturalFlag struct{ quickfix.BooleanValue }

func (f IOINaturalFlag) Tag() tag.Tag { return tag.IOINaturalFlag }

type IOIOthSvc struct{ quickfix.CharValue }

func (f IOIOthSvc) Tag() tag.Tag { return tag.IOIOthSvc }

type IOIQltyInd struct{ quickfix.CharValue }

func (f IOIQltyInd) Tag() tag.Tag { return tag.IOIQltyInd }

type IOIQty struct{ quickfix.StringValue }

func (f IOIQty) Tag() tag.Tag { return tag.IOIQty }

type IOIQualifier struct{ quickfix.CharValue }

func (f IOIQualifier) Tag() tag.Tag { return tag.IOIQualifier }

type IOIRefID struct{ quickfix.StringValue }

func (f IOIRefID) Tag() tag.Tag { return tag.IOIRefID }

type IOIShares struct{ quickfix.StringValue }

func (f IOIShares) Tag() tag.Tag { return tag.IOIShares }

type IOITransType struct{ quickfix.CharValue }

func (f IOITransType) Tag() tag.Tag { return tag.IOITransType }

type IOIid struct{ quickfix.StringValue }

func (f IOIid) Tag() tag.Tag { return tag.IOIid }

type ImpliedMarketIndicator struct{ quickfix.IntValue }

func (f ImpliedMarketIndicator) Tag() tag.Tag { return tag.ImpliedMarketIndicator }

type InViewOfCommon struct{ quickfix.BooleanValue }

func (f InViewOfCommon) Tag() tag.Tag { return tag.InViewOfCommon }

type IncTaxInd struct{ quickfix.IntValue }

func (f IncTaxInd) Tag() tag.Tag { return tag.IncTaxInd }

type IndividualAllocID struct{ quickfix.StringValue }

func (f IndividualAllocID) Tag() tag.Tag { return tag.IndividualAllocID }

type IndividualAllocRejCode struct{ quickfix.IntValue }

func (f IndividualAllocRejCode) Tag() tag.Tag { return tag.IndividualAllocRejCode }

type IndividualAllocType struct{ quickfix.IntValue }

func (f IndividualAllocType) Tag() tag.Tag { return tag.IndividualAllocType }

type InputSource struct{ quickfix.StringValue }

func (f InputSource) Tag() tag.Tag { return tag.InputSource }

type InstrAttribType struct{ quickfix.IntValue }

func (f InstrAttribType) Tag() tag.Tag { return tag.InstrAttribType }

type InstrAttribValue struct{ quickfix.StringValue }

func (f InstrAttribValue) Tag() tag.Tag { return tag.InstrAttribValue }

type InstrRegistry struct{ quickfix.StringValue }

func (f InstrRegistry) Tag() tag.Tag { return tag.InstrRegistry }

type InstrmtAssignmentMethod struct{ quickfix.CharValue }

func (f InstrmtAssignmentMethod) Tag() tag.Tag { return tag.InstrmtAssignmentMethod }

type InstrumentPartyID struct{ quickfix.StringValue }

func (f InstrumentPartyID) Tag() tag.Tag { return tag.InstrumentPartyID }

type InstrumentPartyIDSource struct{ quickfix.CharValue }

func (f InstrumentPartyIDSource) Tag() tag.Tag { return tag.InstrumentPartyIDSource }

type InstrumentPartyRole struct{ quickfix.IntValue }

func (f InstrumentPartyRole) Tag() tag.Tag { return tag.InstrumentPartyRole }

type InstrumentPartySubID struct{ quickfix.StringValue }

func (f InstrumentPartySubID) Tag() tag.Tag { return tag.InstrumentPartySubID }

type InstrumentPartySubIDType struct{ quickfix.IntValue }

func (f InstrumentPartySubIDType) Tag() tag.Tag { return tag.InstrumentPartySubIDType }

type InterestAccrualDate struct{ quickfix.LocalMktDateValue }

func (f InterestAccrualDate) Tag() tag.Tag { return tag.InterestAccrualDate }

type InterestAtMaturity struct{ quickfix.AmtValue }

func (f InterestAtMaturity) Tag() tag.Tag { return tag.InterestAtMaturity }

type InvestorCountryOfResidence struct{ quickfix.CountryValue }

func (f InvestorCountryOfResidence) Tag() tag.Tag { return tag.InvestorCountryOfResidence }

type IssueDate struct{ quickfix.LocalMktDateValue }

func (f IssueDate) Tag() tag.Tag { return tag.IssueDate }

type Issuer struct{ quickfix.StringValue }

func (f Issuer) Tag() tag.Tag { return tag.Issuer }

type LanguageCode struct{ quickfix.LanguageValue }

func (f LanguageCode) Tag() tag.Tag { return tag.LanguageCode }

type LastCapacity struct{ quickfix.CharValue }

func (f LastCapacity) Tag() tag.Tag { return tag.LastCapacity }

type LastForwardPoints struct{ quickfix.PriceOffsetValue }

func (f LastForwardPoints) Tag() tag.Tag { return tag.LastForwardPoints }

type LastForwardPoints2 struct{ quickfix.PriceOffsetValue }

func (f LastForwardPoints2) Tag() tag.Tag { return tag.LastForwardPoints2 }

type LastFragment struct{ quickfix.BooleanValue }

func (f LastFragment) Tag() tag.Tag { return tag.LastFragment }

type LastLiquidityInd struct{ quickfix.IntValue }

func (f LastLiquidityInd) Tag() tag.Tag { return tag.LastLiquidityInd }

type LastMkt struct{ quickfix.ExchangeValue }

func (f LastMkt) Tag() tag.Tag { return tag.LastMkt }

type LastMsgSeqNumProcessed struct{ quickfix.SeqNumValue }

func (f LastMsgSeqNumProcessed) Tag() tag.Tag { return tag.LastMsgSeqNumProcessed }

type LastNetworkResponseID struct{ quickfix.StringValue }

func (f LastNetworkResponseID) Tag() tag.Tag { return tag.LastNetworkResponseID }

type LastParPx struct{ quickfix.PriceValue }

func (f LastParPx) Tag() tag.Tag { return tag.LastParPx }

type LastPx struct{ quickfix.PriceValue }

func (f LastPx) Tag() tag.Tag { return tag.LastPx }

type LastQty struct{ quickfix.QtyValue }

func (f LastQty) Tag() tag.Tag { return tag.LastQty }

type LastRptRequested struct{ quickfix.BooleanValue }

func (f LastRptRequested) Tag() tag.Tag { return tag.LastRptRequested }

type LastShares struct{ quickfix.QtyValue }

func (f LastShares) Tag() tag.Tag { return tag.LastShares }

type LastSpotRate struct{ quickfix.PriceValue }

func (f LastSpotRate) Tag() tag.Tag { return tag.LastSpotRate }

type LastSwapPoints struct{ quickfix.PriceOffsetValue }

func (f LastSwapPoints) Tag() tag.Tag { return tag.LastSwapPoints }

type LastUpdateTime struct{ quickfix.UTCTimestampValue }

func (f LastUpdateTime) Tag() tag.Tag { return tag.LastUpdateTime }

type LateIndicator struct{ quickfix.BooleanValue }

func (f LateIndicator) Tag() tag.Tag { return tag.LateIndicator }

type LeavesQty struct{ quickfix.QtyValue }

func (f LeavesQty) Tag() tag.Tag { return tag.LeavesQty }

type LegAllocAccount struct{ quickfix.StringValue }

func (f LegAllocAccount) Tag() tag.Tag { return tag.LegAllocAccount }

type LegAllocAcctIDSource struct{ quickfix.StringValue }

func (f LegAllocAcctIDSource) Tag() tag.Tag { return tag.LegAllocAcctIDSource }

type LegAllocID struct{ quickfix.StringValue }

func (f LegAllocID) Tag() tag.Tag { return tag.LegAllocID }

type LegAllocQty struct{ quickfix.QtyValue }

func (f LegAllocQty) Tag() tag.Tag { return tag.LegAllocQty }

type LegAllocSettlCurrency struct{ quickfix.CurrencyValue }

func (f LegAllocSettlCurrency) Tag() tag.Tag { return tag.LegAllocSettlCurrency }

type LegBenchmarkCurveCurrency struct{ quickfix.CurrencyValue }

func (f LegBenchmarkCurveCurrency) Tag() tag.Tag { return tag.LegBenchmarkCurveCurrency }

type LegBenchmarkCurveName struct{ quickfix.StringValue }

func (f LegBenchmarkCurveName) Tag() tag.Tag { return tag.LegBenchmarkCurveName }

type LegBenchmarkCurvePoint struct{ quickfix.StringValue }

func (f LegBenchmarkCurvePoint) Tag() tag.Tag { return tag.LegBenchmarkCurvePoint }

type LegBenchmarkPrice struct{ quickfix.PriceValue }

func (f LegBenchmarkPrice) Tag() tag.Tag { return tag.LegBenchmarkPrice }

type LegBenchmarkPriceType struct{ quickfix.IntValue }

func (f LegBenchmarkPriceType) Tag() tag.Tag { return tag.LegBenchmarkPriceType }

type LegBidForwardPoints struct{ quickfix.PriceOffsetValue }

func (f LegBidForwardPoints) Tag() tag.Tag { return tag.LegBidForwardPoints }

type LegBidPx struct{ quickfix.PriceValue }

func (f LegBidPx) Tag() tag.Tag { return tag.LegBidPx }

type LegCFICode struct{ quickfix.StringValue }

func (f LegCFICode) Tag() tag.Tag { return tag.LegCFICode }

type LegCalculatedCcyLastQty struct{ quickfix.QtyValue }

func (f LegCalculatedCcyLastQty) Tag() tag.Tag { return tag.LegCalculatedCcyLastQty }

type LegContractMultiplier struct{ quickfix.FloatValue }

func (f LegContractMultiplier) Tag() tag.Tag { return tag.LegContractMultiplier }

type LegContractMultiplierUnit struct{ quickfix.IntValue }

func (f LegContractMultiplierUnit) Tag() tag.Tag { return tag.LegContractMultiplierUnit }

type LegContractSettlMonth struct{ quickfix.MonthYearValue }

func (f LegContractSettlMonth) Tag() tag.Tag { return tag.LegContractSettlMonth }

type LegCountryOfIssue struct{ quickfix.CountryValue }

func (f LegCountryOfIssue) Tag() tag.Tag { return tag.LegCountryOfIssue }

type LegCouponPaymentDate struct{ quickfix.LocalMktDateValue }

func (f LegCouponPaymentDate) Tag() tag.Tag { return tag.LegCouponPaymentDate }

type LegCouponRate struct{ quickfix.PercentageValue }

func (f LegCouponRate) Tag() tag.Tag { return tag.LegCouponRate }

type LegCoveredOrUncovered struct{ quickfix.IntValue }

func (f LegCoveredOrUncovered) Tag() tag.Tag { return tag.LegCoveredOrUncovered }

type LegCreditRating struct{ quickfix.StringValue }

func (f LegCreditRating) Tag() tag.Tag { return tag.LegCreditRating }

type LegCurrency struct{ quickfix.CurrencyValue }

func (f LegCurrency) Tag() tag.Tag { return tag.LegCurrency }

type LegCurrencyRatio struct{ quickfix.FloatValue }

func (f LegCurrencyRatio) Tag() tag.Tag { return tag.LegCurrencyRatio }

type LegDatedDate struct{ quickfix.LocalMktDateValue }

func (f LegDatedDate) Tag() tag.Tag { return tag.LegDatedDate }

type LegDividendYield struct{ quickfix.PercentageValue }

func (f LegDividendYield) Tag() tag.Tag { return tag.LegDividendYield }

type LegExecInst struct{ quickfix.MultipleCharValue }

func (f LegExecInst) Tag() tag.Tag { return tag.LegExecInst }

type LegExerciseStyle struct{ quickfix.IntValue }

func (f LegExerciseStyle) Tag() tag.Tag { return tag.LegExerciseStyle }

type LegFactor struct{ quickfix.FloatValue }

func (f LegFactor) Tag() tag.Tag { return tag.LegFactor }

type LegFlowScheduleType struct{ quickfix.IntValue }

func (f LegFlowScheduleType) Tag() tag.Tag { return tag.LegFlowScheduleType }

type LegFutSettDate struct{ quickfix.LocalMktDateValue }

func (f LegFutSettDate) Tag() tag.Tag { return tag.LegFutSettDate }

type LegGrossTradeAmt struct{ quickfix.AmtValue }

func (f LegGrossTradeAmt) Tag() tag.Tag { return tag.LegGrossTradeAmt }

type LegIOIQty struct{ quickfix.StringValue }

func (f LegIOIQty) Tag() tag.Tag { return tag.LegIOIQty }

type LegIndividualAllocID struct{ quickfix.StringValue }

func (f LegIndividualAllocID) Tag() tag.Tag { return tag.LegIndividualAllocID }

type LegInstrRegistry struct{ quickfix.StringValue }

func (f LegInstrRegistry) Tag() tag.Tag { return tag.LegInstrRegistry }

type LegInterestAccrualDate struct{ quickfix.LocalMktDateValue }

func (f LegInterestAccrualDate) Tag() tag.Tag { return tag.LegInterestAccrualDate }

type LegIssueDate struct{ quickfix.LocalMktDateValue }

func (f LegIssueDate) Tag() tag.Tag { return tag.LegIssueDate }

type LegIssuer struct{ quickfix.StringValue }

func (f LegIssuer) Tag() tag.Tag { return tag.LegIssuer }

type LegLastForwardPoints struct{ quickfix.PriceOffsetValue }

func (f LegLastForwardPoints) Tag() tag.Tag { return tag.LegLastForwardPoints }

type LegLastPx struct{ quickfix.PriceValue }

func (f LegLastPx) Tag() tag.Tag { return tag.LegLastPx }

type LegLastQty struct{ quickfix.QtyValue }

func (f LegLastQty) Tag() tag.Tag { return tag.LegLastQty }

type LegLocaleOfIssue struct{ quickfix.StringValue }

func (f LegLocaleOfIssue) Tag() tag.Tag { return tag.LegLocaleOfIssue }

type LegMaturityDate struct{ quickfix.LocalMktDateValue }

func (f LegMaturityDate) Tag() tag.Tag { return tag.LegMaturityDate }

type LegMaturityMonthYear struct{ quickfix.MonthYearValue }

func (f LegMaturityMonthYear) Tag() tag.Tag { return tag.LegMaturityMonthYear }

type LegMaturityTime struct{ quickfix.TZTimeOnlyValue }

func (f LegMaturityTime) Tag() tag.Tag { return tag.LegMaturityTime }

type LegNumber struct{ quickfix.IntValue }

func (f LegNumber) Tag() tag.Tag { return tag.LegNumber }

type LegOfferForwardPoints struct{ quickfix.PriceOffsetValue }

func (f LegOfferForwardPoints) Tag() tag.Tag { return tag.LegOfferForwardPoints }

type LegOfferPx struct{ quickfix.PriceValue }

func (f LegOfferPx) Tag() tag.Tag { return tag.LegOfferPx }

type LegOptAttribute struct{ quickfix.CharValue }

func (f LegOptAttribute) Tag() tag.Tag { return tag.LegOptAttribute }

type LegOptionRatio struct{ quickfix.FloatValue }

func (f LegOptionRatio) Tag() tag.Tag { return tag.LegOptionRatio }

type LegOrderQty struct{ quickfix.QtyValue }

func (f LegOrderQty) Tag() tag.Tag { return tag.LegOrderQty }

type LegPool struct{ quickfix.StringValue }

func (f LegPool) Tag() tag.Tag { return tag.LegPool }

type LegPositionEffect struct{ quickfix.CharValue }

func (f LegPositionEffect) Tag() tag.Tag { return tag.LegPositionEffect }

type LegPrice struct{ quickfix.PriceValue }

func (f LegPrice) Tag() tag.Tag { return tag.LegPrice }

type LegPriceType struct{ quickfix.IntValue }

func (f LegPriceType) Tag() tag.Tag { return tag.LegPriceType }

type LegPriceUnitOfMeasure struct{ quickfix.StringValue }

func (f LegPriceUnitOfMeasure) Tag() tag.Tag { return tag.LegPriceUnitOfMeasure }

type LegPriceUnitOfMeasureQty struct{ quickfix.QtyValue }

func (f LegPriceUnitOfMeasureQty) Tag() tag.Tag { return tag.LegPriceUnitOfMeasureQty }

type LegProduct struct{ quickfix.IntValue }

func (f LegProduct) Tag() tag.Tag { return tag.LegProduct }

type LegPutOrCall struct{ quickfix.IntValue }

func (f LegPutOrCall) Tag() tag.Tag { return tag.LegPutOrCall }

type LegQty struct{ quickfix.QtyValue }

func (f LegQty) Tag() tag.Tag { return tag.LegQty }

type LegRatioQty struct{ quickfix.FloatValue }

func (f LegRatioQty) Tag() tag.Tag { return tag.LegRatioQty }

type LegRedemptionDate struct{ quickfix.LocalMktDateValue }

func (f LegRedemptionDate) Tag() tag.Tag { return tag.LegRedemptionDate }

type LegRefID struct{ quickfix.StringValue }

func (f LegRefID) Tag() tag.Tag { return tag.LegRefID }

type LegRepoCollateralSecurityType struct{ quickfix.IntValue }

func (f LegRepoCollateralSecurityType) Tag() tag.Tag { return tag.LegRepoCollateralSecurityType }

type LegReportID struct{ quickfix.StringValue }

func (f LegReportID) Tag() tag.Tag { return tag.LegReportID }

type LegRepurchaseRate struct{ quickfix.PercentageValue }

func (f LegRepurchaseRate) Tag() tag.Tag { return tag.LegRepurchaseRate }

type LegRepurchaseTerm struct{ quickfix.IntValue }

func (f LegRepurchaseTerm) Tag() tag.Tag { return tag.LegRepurchaseTerm }

type LegSecurityAltID struct{ quickfix.StringValue }

func (f LegSecurityAltID) Tag() tag.Tag { return tag.LegSecurityAltID }

type LegSecurityAltIDSource struct{ quickfix.StringValue }

func (f LegSecurityAltIDSource) Tag() tag.Tag { return tag.LegSecurityAltIDSource }

type LegSecurityDesc struct{ quickfix.StringValue }

func (f LegSecurityDesc) Tag() tag.Tag { return tag.LegSecurityDesc }

type LegSecurityExchange struct{ quickfix.ExchangeValue }

func (f LegSecurityExchange) Tag() tag.Tag { return tag.LegSecurityExchange }

type LegSecurityID struct{ quickfix.StringValue }

func (f LegSecurityID) Tag() tag.Tag { return tag.LegSecurityID }

type LegSecurityIDSource struct{ quickfix.StringValue }

func (f LegSecurityIDSource) Tag() tag.Tag { return tag.LegSecurityIDSource }

type LegSecuritySubType struct{ quickfix.StringValue }

func (f LegSecuritySubType) Tag() tag.Tag { return tag.LegSecuritySubType }

type LegSecurityType struct{ quickfix.StringValue }

func (f LegSecurityType) Tag() tag.Tag { return tag.LegSecurityType }

type LegSettlCurrency struct{ quickfix.CurrencyValue }

func (f LegSettlCurrency) Tag() tag.Tag { return tag.LegSettlCurrency }

type LegSettlDate struct{ quickfix.LocalMktDateValue }

func (f LegSettlDate) Tag() tag.Tag { return tag.LegSettlDate }

type LegSettlType struct{ quickfix.CharValue }

func (f LegSettlType) Tag() tag.Tag { return tag.LegSettlType }

type LegSettlmntTyp struct{ quickfix.CharValue }

func (f LegSettlmntTyp) Tag() tag.Tag { return tag.LegSettlmntTyp }

type LegSide struct{ quickfix.CharValue }

func (f LegSide) Tag() tag.Tag { return tag.LegSide }

type LegStateOrProvinceOfIssue struct{ quickfix.StringValue }

func (f LegStateOrProvinceOfIssue) Tag() tag.Tag { return tag.LegStateOrProvinceOfIssue }

type LegStipulationType struct{ quickfix.StringValue }

func (f LegStipulationType) Tag() tag.Tag { return tag.LegStipulationType }

type LegStipulationValue struct{ quickfix.StringValue }

func (f LegStipulationValue) Tag() tag.Tag { return tag.LegStipulationValue }

type LegStrikeCurrency struct{ quickfix.CurrencyValue }

func (f LegStrikeCurrency) Tag() tag.Tag { return tag.LegStrikeCurrency }

type LegStrikePrice struct{ quickfix.PriceValue }

func (f LegStrikePrice) Tag() tag.Tag { return tag.LegStrikePrice }

type LegSwapType struct{ quickfix.IntValue }

func (f LegSwapType) Tag() tag.Tag { return tag.LegSwapType }

type LegSymbol struct{ quickfix.StringValue }

func (f LegSymbol) Tag() tag.Tag { return tag.LegSymbol }

type LegSymbolSfx struct{ quickfix.StringValue }

func (f LegSymbolSfx) Tag() tag.Tag { return tag.LegSymbolSfx }

type LegTimeUnit struct{ quickfix.StringValue }

func (f LegTimeUnit) Tag() tag.Tag { return tag.LegTimeUnit }

type LegUnitOfMeasure struct{ quickfix.StringValue }

func (f LegUnitOfMeasure) Tag() tag.Tag { return tag.LegUnitOfMeasure }

type LegUnitOfMeasureQty struct{ quickfix.QtyValue }

func (f LegUnitOfMeasureQty) Tag() tag.Tag { return tag.LegUnitOfMeasureQty }

type LegVolatility struct{ quickfix.FloatValue }

func (f LegVolatility) Tag() tag.Tag { return tag.LegVolatility }

type LegalConfirm struct{ quickfix.BooleanValue }

func (f LegalConfirm) Tag() tag.Tag { return tag.LegalConfirm }

type LinesOfText struct{ quickfix.NumInGroupValue }

func (f LinesOfText) Tag() tag.Tag { return tag.LinesOfText }

type LiquidityIndType struct{ quickfix.IntValue }

func (f LiquidityIndType) Tag() tag.Tag { return tag.LiquidityIndType }

type LiquidityNumSecurities struct{ quickfix.IntValue }

func (f LiquidityNumSecurities) Tag() tag.Tag { return tag.LiquidityNumSecurities }

type LiquidityPctHigh struct{ quickfix.PercentageValue }

func (f LiquidityPctHigh) Tag() tag.Tag { return tag.LiquidityPctHigh }

type LiquidityPctLow struct{ quickfix.PercentageValue }

func (f LiquidityPctLow) Tag() tag.Tag { return tag.LiquidityPctLow }

type LiquidityValue struct{ quickfix.AmtValue }

func (f LiquidityValue) Tag() tag.Tag { return tag.LiquidityValue }

type ListExecInst struct{ quickfix.StringValue }

func (f ListExecInst) Tag() tag.Tag { return tag.ListExecInst }

type ListExecInstType struct{ quickfix.CharValue }

func (f ListExecInstType) Tag() tag.Tag { return tag.ListExecInstType }

type ListID struct{ quickfix.StringValue }

func (f ListID) Tag() tag.Tag { return tag.ListID }

type ListMethod struct{ quickfix.IntValue }

func (f ListMethod) Tag() tag.Tag { return tag.ListMethod }

type ListName struct{ quickfix.StringValue }

func (f ListName) Tag() tag.Tag { return tag.ListName }

type ListNoOrds struct{ quickfix.IntValue }

func (f ListNoOrds) Tag() tag.Tag { return tag.ListNoOrds }

type ListOrderStatus struct{ quickfix.IntValue }

func (f ListOrderStatus) Tag() tag.Tag { return tag.ListOrderStatus }

type ListRejectReason struct{ quickfix.IntValue }

func (f ListRejectReason) Tag() tag.Tag { return tag.ListRejectReason }

type ListSeqNo struct{ quickfix.IntValue }

func (f ListSeqNo) Tag() tag.Tag { return tag.ListSeqNo }

type ListStatusText struct{ quickfix.StringValue }

func (f ListStatusText) Tag() tag.Tag { return tag.ListStatusText }

type ListStatusType struct{ quickfix.IntValue }

func (f ListStatusType) Tag() tag.Tag { return tag.ListStatusType }

type ListUpdateAction struct{ quickfix.CharValue }

func (f ListUpdateAction) Tag() tag.Tag { return tag.ListUpdateAction }

type LocaleOfIssue struct{ quickfix.StringValue }

func (f LocaleOfIssue) Tag() tag.Tag { return tag.LocaleOfIssue }

type LocateReqd struct{ quickfix.BooleanValue }

func (f LocateReqd) Tag() tag.Tag { return tag.LocateReqd }

type LocationID struct{ quickfix.StringValue }

func (f LocationID) Tag() tag.Tag { return tag.LocationID }

type LongQty struct{ quickfix.QtyValue }

func (f LongQty) Tag() tag.Tag { return tag.LongQty }

type LotType struct{ quickfix.CharValue }

func (f LotType) Tag() tag.Tag { return tag.LotType }

type LowLimitPrice struct{ quickfix.PriceValue }

func (f LowLimitPrice) Tag() tag.Tag { return tag.LowLimitPrice }

type LowPx struct{ quickfix.PriceValue }

func (f LowPx) Tag() tag.Tag { return tag.LowPx }

type MDBookType struct{ quickfix.IntValue }

func (f MDBookType) Tag() tag.Tag { return tag.MDBookType }

type MDEntryBuyer struct{ quickfix.StringValue }

func (f MDEntryBuyer) Tag() tag.Tag { return tag.MDEntryBuyer }

type MDEntryDate struct{ quickfix.UTCDateOnlyValue }

func (f MDEntryDate) Tag() tag.Tag { return tag.MDEntryDate }

type MDEntryForwardPoints struct{ quickfix.PriceOffsetValue }

func (f MDEntryForwardPoints) Tag() tag.Tag { return tag.MDEntryForwardPoints }

type MDEntryID struct{ quickfix.StringValue }

func (f MDEntryID) Tag() tag.Tag { return tag.MDEntryID }

type MDEntryOriginator struct{ quickfix.StringValue }

func (f MDEntryOriginator) Tag() tag.Tag { return tag.MDEntryOriginator }

type MDEntryPositionNo struct{ quickfix.IntValue }

func (f MDEntryPositionNo) Tag() tag.Tag { return tag.MDEntryPositionNo }

type MDEntryPx struct{ quickfix.PriceValue }

func (f MDEntryPx) Tag() tag.Tag { return tag.MDEntryPx }

type MDEntryRefID struct{ quickfix.StringValue }

func (f MDEntryRefID) Tag() tag.Tag { return tag.MDEntryRefID }

type MDEntrySeller struct{ quickfix.StringValue }

func (f MDEntrySeller) Tag() tag.Tag { return tag.MDEntrySeller }

type MDEntrySize struct{ quickfix.QtyValue }

func (f MDEntrySize) Tag() tag.Tag { return tag.MDEntrySize }

type MDEntrySpotRate struct{ quickfix.FloatValue }

func (f MDEntrySpotRate) Tag() tag.Tag { return tag.MDEntrySpotRate }

type MDEntryTime struct{ quickfix.UTCTimeOnlyValue }

func (f MDEntryTime) Tag() tag.Tag { return tag.MDEntryTime }

type MDEntryType struct{ quickfix.CharValue }

func (f MDEntryType) Tag() tag.Tag { return tag.MDEntryType }

type MDFeedType struct{ quickfix.StringValue }

func (f MDFeedType) Tag() tag.Tag { return tag.MDFeedType }

type MDImplicitDelete struct{ quickfix.BooleanValue }

func (f MDImplicitDelete) Tag() tag.Tag { return tag.MDImplicitDelete }

type MDMkt struct{ quickfix.ExchangeValue }

func (f MDMkt) Tag() tag.Tag { return tag.MDMkt }

type MDOriginType struct{ quickfix.IntValue }

func (f MDOriginType) Tag() tag.Tag { return tag.MDOriginType }

type MDPriceLevel struct{ quickfix.IntValue }

func (f MDPriceLevel) Tag() tag.Tag { return tag.MDPriceLevel }

type MDQuoteType struct{ quickfix.IntValue }

func (f MDQuoteType) Tag() tag.Tag { return tag.MDQuoteType }

type MDReportID struct{ quickfix.IntValue }

func (f MDReportID) Tag() tag.Tag { return tag.MDReportID }

type MDReqID struct{ quickfix.StringValue }

func (f MDReqID) Tag() tag.Tag { return tag.MDReqID }

type MDReqRejReason struct{ quickfix.CharValue }

func (f MDReqRejReason) Tag() tag.Tag { return tag.MDReqRejReason }

type MDSecSize struct{ quickfix.QtyValue }

func (f MDSecSize) Tag() tag.Tag { return tag.MDSecSize }

type MDSecSizeType struct{ quickfix.IntValue }

func (f MDSecSizeType) Tag() tag.Tag { return tag.MDSecSizeType }

type MDStreamID struct{ quickfix.StringValue }

func (f MDStreamID) Tag() tag.Tag { return tag.MDStreamID }

type MDSubBookType struct{ quickfix.IntValue }

func (f MDSubBookType) Tag() tag.Tag { return tag.MDSubBookType }

type MDUpdateAction struct{ quickfix.CharValue }

func (f MDUpdateAction) Tag() tag.Tag { return tag.MDUpdateAction }

type MDUpdateType struct{ quickfix.IntValue }

func (f MDUpdateType) Tag() tag.Tag { return tag.MDUpdateType }

type MailingDtls struct{ quickfix.StringValue }

func (f MailingDtls) Tag() tag.Tag { return tag.MailingDtls }

type MailingInst struct{ quickfix.StringValue }

func (f MailingInst) Tag() tag.Tag { return tag.MailingInst }

type ManualOrderIndicator struct{ quickfix.BooleanValue }

func (f ManualOrderIndicator) Tag() tag.Tag { return tag.ManualOrderIndicator }

type MarginExcess struct{ quickfix.AmtValue }

func (f MarginExcess) Tag() tag.Tag { return tag.MarginExcess }

type MarginRatio struct{ quickfix.PercentageValue }

func (f MarginRatio) Tag() tag.Tag { return tag.MarginRatio }

type MarketDepth struct{ quickfix.IntValue }

func (f MarketDepth) Tag() tag.Tag { return tag.MarketDepth }

type MarketID struct{ quickfix.ExchangeValue }

func (f MarketID) Tag() tag.Tag { return tag.MarketID }

type MarketReportID struct{ quickfix.StringValue }

func (f MarketReportID) Tag() tag.Tag { return tag.MarketReportID }

type MarketReqID struct{ quickfix.StringValue }

func (f MarketReqID) Tag() tag.Tag { return tag.MarketReqID }

type MarketSegmentDesc struct{ quickfix.StringValue }

func (f MarketSegmentDesc) Tag() tag.Tag { return tag.MarketSegmentDesc }

type MarketSegmentID struct{ quickfix.StringValue }

func (f MarketSegmentID) Tag() tag.Tag { return tag.MarketSegmentID }

type MarketUpdateAction struct{ quickfix.CharValue }

func (f MarketUpdateAction) Tag() tag.Tag { return tag.MarketUpdateAction }

type MassActionRejectReason struct{ quickfix.IntValue }

func (f MassActionRejectReason) Tag() tag.Tag { return tag.MassActionRejectReason }

type MassActionReportID struct{ quickfix.StringValue }

func (f MassActionReportID) Tag() tag.Tag { return tag.MassActionReportID }

type MassActionResponse struct{ quickfix.IntValue }

func (f MassActionResponse) Tag() tag.Tag { return tag.MassActionResponse }

type MassActionScope struct{ quickfix.IntValue }

func (f MassActionScope) Tag() tag.Tag { return tag.MassActionScope }

type MassActionType struct{ quickfix.IntValue }

func (f MassActionType) Tag() tag.Tag { return tag.MassActionType }

type MassCancelRejectReason struct{ quickfix.IntValue }

func (f MassCancelRejectReason) Tag() tag.Tag { return tag.MassCancelRejectReason }

type MassCancelRequestType struct{ quickfix.CharValue }

func (f MassCancelRequestType) Tag() tag.Tag { return tag.MassCancelRequestType }

type MassCancelResponse struct{ quickfix.CharValue }

func (f MassCancelResponse) Tag() tag.Tag { return tag.MassCancelResponse }

type MassStatusReqID struct{ quickfix.StringValue }

func (f MassStatusReqID) Tag() tag.Tag { return tag.MassStatusReqID }

type MassStatusReqType struct{ quickfix.IntValue }

func (f MassStatusReqType) Tag() tag.Tag { return tag.MassStatusReqType }

type MatchAlgorithm struct{ quickfix.StringValue }

func (f MatchAlgorithm) Tag() tag.Tag { return tag.MatchAlgorithm }

type MatchIncrement struct{ quickfix.QtyValue }

func (f MatchIncrement) Tag() tag.Tag { return tag.MatchIncrement }

type MatchStatus struct{ quickfix.CharValue }

func (f MatchStatus) Tag() tag.Tag { return tag.MatchStatus }

type MatchType struct{ quickfix.StringValue }

func (f MatchType) Tag() tag.Tag { return tag.MatchType }

type MaturityDate struct{ quickfix.LocalMktDateValue }

func (f MaturityDate) Tag() tag.Tag { return tag.MaturityDate }

type MaturityDay struct{ quickfix.DayOfMonthValue }

func (f MaturityDay) Tag() tag.Tag { return tag.MaturityDay }

type MaturityMonthYear struct{ quickfix.MonthYearValue }

func (f MaturityMonthYear) Tag() tag.Tag { return tag.MaturityMonthYear }

type MaturityMonthYearFormat struct{ quickfix.IntValue }

func (f MaturityMonthYearFormat) Tag() tag.Tag { return tag.MaturityMonthYearFormat }

type MaturityMonthYearIncrement struct{ quickfix.IntValue }

func (f MaturityMonthYearIncrement) Tag() tag.Tag { return tag.MaturityMonthYearIncrement }

type MaturityMonthYearIncrementUnits struct{ quickfix.IntValue }

func (f MaturityMonthYearIncrementUnits) Tag() tag.Tag { return tag.MaturityMonthYearIncrementUnits }

type MaturityNetMoney struct{ quickfix.AmtValue }

func (f MaturityNetMoney) Tag() tag.Tag { return tag.MaturityNetMoney }

type MaturityRuleID struct{ quickfix.StringValue }

func (f MaturityRuleID) Tag() tag.Tag { return tag.MaturityRuleID }

type MaturityTime struct{ quickfix.TZTimeOnlyValue }

func (f MaturityTime) Tag() tag.Tag { return tag.MaturityTime }

type MaxFloor struct{ quickfix.QtyValue }

func (f MaxFloor) Tag() tag.Tag { return tag.MaxFloor }

type MaxMessageSize struct{ quickfix.LengthValue }

func (f MaxMessageSize) Tag() tag.Tag { return tag.MaxMessageSize }

type MaxPriceLevels struct{ quickfix.IntValue }

func (f MaxPriceLevels) Tag() tag.Tag { return tag.MaxPriceLevels }

type MaxPriceVariation struct{ quickfix.FloatValue }

func (f MaxPriceVariation) Tag() tag.Tag { return tag.MaxPriceVariation }

type MaxShow struct{ quickfix.QtyValue }

func (f MaxShow) Tag() tag.Tag { return tag.MaxShow }

type MaxTradeVol struct{ quickfix.QtyValue }

func (f MaxTradeVol) Tag() tag.Tag { return tag.MaxTradeVol }

type MessageEncoding struct{ quickfix.StringValue }

func (f MessageEncoding) Tag() tag.Tag { return tag.MessageEncoding }

type MessageEventSource struct{ quickfix.StringValue }

func (f MessageEventSource) Tag() tag.Tag { return tag.MessageEventSource }

type MidPx struct{ quickfix.PriceValue }

func (f MidPx) Tag() tag.Tag { return tag.MidPx }

type MidYield struct{ quickfix.PercentageValue }

func (f MidYield) Tag() tag.Tag { return tag.MidYield }

type MinBidSize struct{ quickfix.QtyValue }

func (f MinBidSize) Tag() tag.Tag { return tag.MinBidSize }

type MinLotSize struct{ quickfix.QtyValue }

func (f MinLotSize) Tag() tag.Tag { return tag.MinLotSize }

type MinOfferSize struct{ quickfix.QtyValue }

func (f MinOfferSize) Tag() tag.Tag { return tag.MinOfferSize }

type MinPriceIncrement struct{ quickfix.FloatValue }

func (f MinPriceIncrement) Tag() tag.Tag { return tag.MinPriceIncrement }

type MinPriceIncrementAmount struct{ quickfix.AmtValue }

func (f MinPriceIncrementAmount) Tag() tag.Tag { return tag.MinPriceIncrementAmount }

type MinQty struct{ quickfix.QtyValue }

func (f MinQty) Tag() tag.Tag { return tag.MinQty }

type MinTradeVol struct{ quickfix.QtyValue }

func (f MinTradeVol) Tag() tag.Tag { return tag.MinTradeVol }

type MiscFeeAmt struct{ quickfix.AmtValue }

func (f MiscFeeAmt) Tag() tag.Tag { return tag.MiscFeeAmt }

type MiscFeeBasis struct{ quickfix.IntValue }

func (f MiscFeeBasis) Tag() tag.Tag { return tag.MiscFeeBasis }

type MiscFeeCurr struct{ quickfix.CurrencyValue }

func (f MiscFeeCurr) Tag() tag.Tag { return tag.MiscFeeCurr }

type MiscFeeType struct{ quickfix.StringValue }

func (f MiscFeeType) Tag() tag.Tag { return tag.MiscFeeType }

type MktBidPx struct{ quickfix.PriceValue }

func (f MktBidPx) Tag() tag.Tag { return tag.MktBidPx }

type MktOfferPx struct{ quickfix.PriceValue }

func (f MktOfferPx) Tag() tag.Tag { return tag.MktOfferPx }

type ModelType struct{ quickfix.IntValue }

func (f ModelType) Tag() tag.Tag { return tag.ModelType }

type MoneyLaunderingStatus struct{ quickfix.CharValue }

func (f MoneyLaunderingStatus) Tag() tag.Tag { return tag.MoneyLaunderingStatus }

type MsgDirection struct{ quickfix.CharValue }

func (f MsgDirection) Tag() tag.Tag { return tag.MsgDirection }

type MsgSeqNum struct{ quickfix.SeqNumValue }

func (f MsgSeqNum) Tag() tag.Tag { return tag.MsgSeqNum }

type MsgType struct{ quickfix.StringValue }

func (f MsgType) Tag() tag.Tag { return tag.MsgType }

type MultiLegReportingType struct{ quickfix.CharValue }

func (f MultiLegReportingType) Tag() tag.Tag { return tag.MultiLegReportingType }

type MultiLegRptTypeReq struct{ quickfix.IntValue }

func (f MultiLegRptTypeReq) Tag() tag.Tag { return tag.MultiLegRptTypeReq }

type MultilegModel struct{ quickfix.IntValue }

func (f MultilegModel) Tag() tag.Tag { return tag.MultilegModel }

type MultilegPriceMethod struct{ quickfix.IntValue }

func (f MultilegPriceMethod) Tag() tag.Tag { return tag.MultilegPriceMethod }

type NTPositionLimit struct{ quickfix.IntValue }

func (f NTPositionLimit) Tag() tag.Tag { return tag.NTPositionLimit }

type Nested2PartyID struct{ quickfix.StringValue }

func (f Nested2PartyID) Tag() tag.Tag { return tag.Nested2PartyID }

type Nested2PartyIDSource struct{ quickfix.CharValue }

func (f Nested2PartyIDSource) Tag() tag.Tag { return tag.Nested2PartyIDSource }

type Nested2PartyRole struct{ quickfix.IntValue }

func (f Nested2PartyRole) Tag() tag.Tag { return tag.Nested2PartyRole }

type Nested2PartySubID struct{ quickfix.StringValue }

func (f Nested2PartySubID) Tag() tag.Tag { return tag.Nested2PartySubID }

type Nested2PartySubIDType struct{ quickfix.IntValue }

func (f Nested2PartySubIDType) Tag() tag.Tag { return tag.Nested2PartySubIDType }

type Nested3PartyID struct{ quickfix.StringValue }

func (f Nested3PartyID) Tag() tag.Tag { return tag.Nested3PartyID }

type Nested3PartyIDSource struct{ quickfix.CharValue }

func (f Nested3PartyIDSource) Tag() tag.Tag { return tag.Nested3PartyIDSource }

type Nested3PartyRole struct{ quickfix.IntValue }

func (f Nested3PartyRole) Tag() tag.Tag { return tag.Nested3PartyRole }

type Nested3PartySubID struct{ quickfix.StringValue }

func (f Nested3PartySubID) Tag() tag.Tag { return tag.Nested3PartySubID }

type Nested3PartySubIDType struct{ quickfix.IntValue }

func (f Nested3PartySubIDType) Tag() tag.Tag { return tag.Nested3PartySubIDType }

type Nested4PartyID struct{ quickfix.StringValue }

func (f Nested4PartyID) Tag() tag.Tag { return tag.Nested4PartyID }

type Nested4PartyIDSource struct{ quickfix.CharValue }

func (f Nested4PartyIDSource) Tag() tag.Tag { return tag.Nested4PartyIDSource }

type Nested4PartyRole struct{ quickfix.IntValue }

func (f Nested4PartyRole) Tag() tag.Tag { return tag.Nested4PartyRole }

type Nested4PartySubID struct{ quickfix.StringValue }

func (f Nested4PartySubID) Tag() tag.Tag { return tag.Nested4PartySubID }

type Nested4PartySubIDType struct{ quickfix.IntValue }

func (f Nested4PartySubIDType) Tag() tag.Tag { return tag.Nested4PartySubIDType }

type NestedInstrAttribType struct{ quickfix.IntValue }

func (f NestedInstrAttribType) Tag() tag.Tag { return tag.NestedInstrAttribType }

type NestedInstrAttribValue struct{ quickfix.StringValue }

func (f NestedInstrAttribValue) Tag() tag.Tag { return tag.NestedInstrAttribValue }

type NestedPartyID struct{ quickfix.StringValue }

func (f NestedPartyID) Tag() tag.Tag { return tag.NestedPartyID }

type NestedPartyIDSource struct{ quickfix.CharValue }

func (f NestedPartyIDSource) Tag() tag.Tag { return tag.NestedPartyIDSource }

type NestedPartyRole struct{ quickfix.IntValue }

func (f NestedPartyRole) Tag() tag.Tag { return tag.NestedPartyRole }

type NestedPartySubID struct{ quickfix.StringValue }

func (f NestedPartySubID) Tag() tag.Tag { return tag.NestedPartySubID }

type NestedPartySubIDType struct{ quickfix.IntValue }

func (f NestedPartySubIDType) Tag() tag.Tag { return tag.NestedPartySubIDType }

type NetChgPrevDay struct{ quickfix.PriceOffsetValue }

func (f NetChgPrevDay) Tag() tag.Tag { return tag.NetChgPrevDay }

type NetGrossInd struct{ quickfix.IntValue }

func (f NetGrossInd) Tag() tag.Tag { return tag.NetGrossInd }

type NetMoney struct{ quickfix.AmtValue }

func (f NetMoney) Tag() tag.Tag { return tag.NetMoney }

type NetworkRequestID struct{ quickfix.StringValue }

func (f NetworkRequestID) Tag() tag.Tag { return tag.NetworkRequestID }

type NetworkRequestType struct{ quickfix.IntValue }

func (f NetworkRequestType) Tag() tag.Tag { return tag.NetworkRequestType }

type NetworkResponseID struct{ quickfix.StringValue }

func (f NetworkResponseID) Tag() tag.Tag { return tag.NetworkResponseID }

type NetworkStatusResponseType struct{ quickfix.IntValue }

func (f NetworkStatusResponseType) Tag() tag.Tag { return tag.NetworkStatusResponseType }

type NewPassword struct{ quickfix.StringValue }

func (f NewPassword) Tag() tag.Tag { return tag.NewPassword }

type NewSeqNo struct{ quickfix.SeqNumValue }

func (f NewSeqNo) Tag() tag.Tag { return tag.NewSeqNo }

type NewsCategory struct{ quickfix.IntValue }

func (f NewsCategory) Tag() tag.Tag { return tag.NewsCategory }

type NewsID struct{ quickfix.StringValue }

func (f NewsID) Tag() tag.Tag { return tag.NewsID }

type NewsRefID struct{ quickfix.StringValue }

func (f NewsRefID) Tag() tag.Tag { return tag.NewsRefID }

type NewsRefType struct{ quickfix.IntValue }

func (f NewsRefType) Tag() tag.Tag { return tag.NewsRefType }

type NextExpectedMsgSeqNum struct{ quickfix.SeqNumValue }

func (f NextExpectedMsgSeqNum) Tag() tag.Tag { return tag.NextExpectedMsgSeqNum }

type NoAffectedOrders struct{ quickfix.NumInGroupValue }

func (f NoAffectedOrders) Tag() tag.Tag { return tag.NoAffectedOrders }

type NoAllocs struct{ quickfix.NumInGroupValue }

func (f NoAllocs) Tag() tag.Tag { return tag.NoAllocs }

type NoAltMDSource struct{ quickfix.NumInGroupValue }

func (f NoAltMDSource) Tag() tag.Tag { return tag.NoAltMDSource }

type NoApplIDs struct{ quickfix.NumInGroupValue }

func (f NoApplIDs) Tag() tag.Tag { return tag.NoApplIDs }

type NoAsgnReqs struct{ quickfix.NumInGroupValue }

func (f NoAsgnReqs) Tag() tag.Tag { return tag.NoAsgnReqs }

type NoBidComponents struct{ quickfix.NumInGroupValue }

func (f NoBidComponents) Tag() tag.Tag { return tag.NoBidComponents }

type NoBidDescriptors struct{ quickfix.NumInGroupValue }

func (f NoBidDescriptors) Tag() tag.Tag { return tag.NoBidDescriptors }

type NoCapacities struct{ quickfix.NumInGroupValue }

func (f NoCapacities) Tag() tag.Tag { return tag.NoCapacities }

type NoClearingInstructions struct{ quickfix.NumInGroupValue }

func (f NoClearingInstructions) Tag() tag.Tag { return tag.NoClearingInstructions }

type NoCollInquiryQualifier struct{ quickfix.NumInGroupValue }

func (f NoCollInquiryQualifier) Tag() tag.Tag { return tag.NoCollInquiryQualifier }

type NoCompIDs struct{ quickfix.NumInGroupValue }

func (f NoCompIDs) Tag() tag.Tag { return tag.NoCompIDs }

type NoComplexEventDates struct{ quickfix.NumInGroupValue }

func (f NoComplexEventDates) Tag() tag.Tag { return tag.NoComplexEventDates }

type NoComplexEventTimes struct{ quickfix.NumInGroupValue }

func (f NoComplexEventTimes) Tag() tag.Tag { return tag.NoComplexEventTimes }

type NoComplexEvents struct{ quickfix.NumInGroupValue }

func (f NoComplexEvents) Tag() tag.Tag { return tag.NoComplexEvents }

type NoContAmts struct{ quickfix.NumInGroupValue }

func (f NoContAmts) Tag() tag.Tag { return tag.NoContAmts }

type NoContextPartyIDs struct{ quickfix.NumInGroupValue }

func (f NoContextPartyIDs) Tag() tag.Tag { return tag.NoContextPartyIDs }

type NoContextPartySubIDs struct{ quickfix.NumInGroupValue }

func (f NoContextPartySubIDs) Tag() tag.Tag { return tag.NoContextPartySubIDs }

type NoContraBrokers struct{ quickfix.NumInGroupValue }

func (f NoContraBrokers) Tag() tag.Tag { return tag.NoContraBrokers }

type NoDates struct{ quickfix.IntValue }

func (f NoDates) Tag() tag.Tag { return tag.NoDates }

type NoDerivativeEvents struct{ quickfix.NumInGroupValue }

func (f NoDerivativeEvents) Tag() tag.Tag { return tag.NoDerivativeEvents }

type NoDerivativeInstrAttrib struct{ quickfix.NumInGroupValue }

func (f NoDerivativeInstrAttrib) Tag() tag.Tag { return tag.NoDerivativeInstrAttrib }

type NoDerivativeInstrumentParties struct{ quickfix.NumInGroupValue }

func (f NoDerivativeInstrumentParties) Tag() tag.Tag { return tag.NoDerivativeInstrumentParties }

type NoDerivativeInstrumentPartySubIDs struct{ quickfix.NumInGroupValue }

func (f NoDerivativeInstrumentPartySubIDs) Tag() tag.Tag { return tag.NoDerivativeInstrumentPartySubIDs }

type NoDerivativeSecurityAltID struct{ quickfix.NumInGroupValue }

func (f NoDerivativeSecurityAltID) Tag() tag.Tag { return tag.NoDerivativeSecurityAltID }

type NoDistribInsts struct{ quickfix.NumInGroupValue }

func (f NoDistribInsts) Tag() tag.Tag { return tag.NoDistribInsts }

type NoDlvyInst struct{ quickfix.NumInGroupValue }

func (f NoDlvyInst) Tag() tag.Tag { return tag.NoDlvyInst }

type NoEvents struct{ quickfix.NumInGroupValue }

func (f NoEvents) Tag() tag.Tag { return tag.NoEvents }

type NoExecInstRules struct{ quickfix.NumInGroupValue }

func (f NoExecInstRules) Tag() tag.Tag { return tag.NoExecInstRules }

type NoExecs struct{ quickfix.NumInGroupValue }

func (f NoExecs) Tag() tag.Tag { return tag.NoExecs }

type NoExpiration struct{ quickfix.NumInGroupValue }

func (f NoExpiration) Tag() tag.Tag { return tag.NoExpiration }

type NoFills struct{ quickfix.NumInGroupValue }

func (f NoFills) Tag() tag.Tag { return tag.NoFills }

type NoHops struct{ quickfix.NumInGroupValue }

func (f NoHops) Tag() tag.Tag { return tag.NoHops }

type NoIOIQualifiers struct{ quickfix.NumInGroupValue }

func (f NoIOIQualifiers) Tag() tag.Tag { return tag.NoIOIQualifiers }

type NoInstrAttrib struct{ quickfix.NumInGroupValue }

func (f NoInstrAttrib) Tag() tag.Tag { return tag.NoInstrAttrib }

type NoInstrumentParties struct{ quickfix.NumInGroupValue }

func (f NoInstrumentParties) Tag() tag.Tag { return tag.NoInstrumentParties }

type NoInstrumentPartySubIDs struct{ quickfix.NumInGroupValue }

func (f NoInstrumentPartySubIDs) Tag() tag.Tag { return tag.NoInstrumentPartySubIDs }

type NoLegAllocs struct{ quickfix.NumInGroupValue }

func (f NoLegAllocs) Tag() tag.Tag { return tag.NoLegAllocs }

type NoLegSecurityAltID struct{ quickfix.StringValue }

func (f NoLegSecurityAltID) Tag() tag.Tag { return tag.NoLegSecurityAltID }

type NoLegStipulations struct{ quickfix.NumInGroupValue }

func (f NoLegStipulations) Tag() tag.Tag { return tag.NoLegStipulations }

type NoLegs struct{ quickfix.NumInGroupValue }

func (f NoLegs) Tag() tag.Tag { return tag.NoLegs }

type NoLinesOfText struct{ quickfix.NumInGroupValue }

func (f NoLinesOfText) Tag() tag.Tag { return tag.NoLinesOfText }

type NoLotTypeRules struct{ quickfix.NumInGroupValue }

func (f NoLotTypeRules) Tag() tag.Tag { return tag.NoLotTypeRules }

type NoMDEntries struct{ quickfix.NumInGroupValue }

func (f NoMDEntries) Tag() tag.Tag { return tag.NoMDEntries }

type NoMDEntryTypes struct{ quickfix.NumInGroupValue }

func (f NoMDEntryTypes) Tag() tag.Tag { return tag.NoMDEntryTypes }

type NoMDFeedTypes struct{ quickfix.NumInGroupValue }

func (f NoMDFeedTypes) Tag() tag.Tag { return tag.NoMDFeedTypes }

type NoMarketSegments struct{ quickfix.NumInGroupValue }

func (f NoMarketSegments) Tag() tag.Tag { return tag.NoMarketSegments }

type NoMatchRules struct{ quickfix.NumInGroupValue }

func (f NoMatchRules) Tag() tag.Tag { return tag.NoMatchRules }

type NoMaturityRules struct{ quickfix.NumInGroupValue }

func (f NoMaturityRules) Tag() tag.Tag { return tag.NoMaturityRules }

type NoMiscFees struct{ quickfix.NumInGroupValue }

func (f NoMiscFees) Tag() tag.Tag { return tag.NoMiscFees }

type NoMsgTypes struct{ quickfix.NumInGroupValue }

func (f NoMsgTypes) Tag() tag.Tag { return tag.NoMsgTypes }

type NoNested2PartyIDs struct{ quickfix.NumInGroupValue }

func (f NoNested2PartyIDs) Tag() tag.Tag { return tag.NoNested2PartyIDs }

type NoNested2PartySubIDs struct{ quickfix.NumInGroupValue }

func (f NoNested2PartySubIDs) Tag() tag.Tag { return tag.NoNested2PartySubIDs }

type NoNested3PartyIDs struct{ quickfix.NumInGroupValue }

func (f NoNested3PartyIDs) Tag() tag.Tag { return tag.NoNested3PartyIDs }

type NoNested3PartySubIDs struct{ quickfix.NumInGroupValue }

func (f NoNested3PartySubIDs) Tag() tag.Tag { return tag.NoNested3PartySubIDs }

type NoNested4PartyIDs struct{ quickfix.NumInGroupValue }

func (f NoNested4PartyIDs) Tag() tag.Tag { return tag.NoNested4PartyIDs }

type NoNested4PartySubIDs struct{ quickfix.NumInGroupValue }

func (f NoNested4PartySubIDs) Tag() tag.Tag { return tag.NoNested4PartySubIDs }

type NoNestedInstrAttrib struct{ quickfix.NumInGroupValue }

func (f NoNestedInstrAttrib) Tag() tag.Tag { return tag.NoNestedInstrAttrib }

type NoNestedPartyIDs struct{ quickfix.NumInGroupValue }

func (f NoNestedPartyIDs) Tag() tag.Tag { return tag.NoNestedPartyIDs }

type NoNestedPartySubIDs struct{ quickfix.NumInGroupValue }

func (f NoNestedPartySubIDs) Tag() tag.Tag { return tag.NoNestedPartySubIDs }

type NoNewsRefIDs struct{ quickfix.NumInGroupValue }

func (f NoNewsRefIDs) Tag() tag.Tag { return tag.NoNewsRefIDs }

type NoNotAffectedOrders struct{ quickfix.NumInGroupValue }

func (f NoNotAffectedOrders) Tag() tag.Tag { return tag.NoNotAffectedOrders }

type NoOfLegUnderlyings struct{ quickfix.NumInGroupValue }

func (f NoOfLegUnderlyings) Tag() tag.Tag { return tag.NoOfLegUnderlyings }

type NoOfSecSizes struct{ quickfix.NumInGroupValue }

func (f NoOfSecSizes) Tag() tag.Tag { return tag.NoOfSecSizes }

type NoOrdTypeRules struct{ quickfix.NumInGroupValue }

func (f NoOrdTypeRules) Tag() tag.Tag { return tag.NoOrdTypeRules }

type NoOrders struct{ quickfix.NumInGroupValue }

func (f NoOrders) Tag() tag.Tag { return tag.NoOrders }

type NoPartyAltIDs struct{ quickfix.NumInGroupValue }

func (f NoPartyAltIDs) Tag() tag.Tag { return tag.NoPartyAltIDs }

type NoPartyAltSubIDs struct{ quickfix.NumInGroupValue }

func (f NoPartyAltSubIDs) Tag() tag.Tag { return tag.NoPartyAltSubIDs }

type NoPartyIDs struct{ quickfix.NumInGroupValue }

func (f NoPartyIDs) Tag() tag.Tag { return tag.NoPartyIDs }

type NoPartyList struct{ quickfix.NumInGroupValue }

func (f NoPartyList) Tag() tag.Tag { return tag.NoPartyList }

type NoPartyListResponseTypes struct{ quickfix.NumInGroupValue }

func (f NoPartyListResponseTypes) Tag() tag.Tag { return tag.NoPartyListResponseTypes }

type NoPartyRelationships struct{ quickfix.NumInGroupValue }

func (f NoPartyRelationships) Tag() tag.Tag { return tag.NoPartyRelationships }

type NoPartySubIDs struct{ quickfix.NumInGroupValue }

func (f NoPartySubIDs) Tag() tag.Tag { return tag.NoPartySubIDs }

type NoPosAmt struct{ quickfix.NumInGroupValue }

func (f NoPosAmt) Tag() tag.Tag { return tag.NoPosAmt }

type NoPositions struct{ quickfix.NumInGroupValue }

func (f NoPositions) Tag() tag.Tag { return tag.NoPositions }

type NoQuoteEntries struct{ quickfix.NumInGroupValue }

func (f NoQuoteEntries) Tag() tag.Tag { return tag.NoQuoteEntries }

type NoQuoteQualifiers struct{ quickfix.NumInGroupValue }

func (f NoQuoteQualifiers) Tag() tag.Tag { return tag.NoQuoteQualifiers }

type NoQuoteSets struct{ quickfix.NumInGroupValue }

func (f NoQuoteSets) Tag() tag.Tag { return tag.NoQuoteSets }

type NoRateSources struct{ quickfix.NumInGroupValue }

func (f NoRateSources) Tag() tag.Tag { return tag.NoRateSources }

type NoRegistDtls struct{ quickfix.NumInGroupValue }

func (f NoRegistDtls) Tag() tag.Tag { return tag.NoRegistDtls }

type NoRelatedContextPartyIDs struct{ quickfix.NumInGroupValue }

func (f NoRelatedContextPartyIDs) Tag() tag.Tag { return tag.NoRelatedContextPartyIDs }

type NoRelatedContextPartySubIDs struct{ quickfix.NumInGroupValue }

func (f NoRelatedContextPartySubIDs) Tag() tag.Tag { return tag.NoRelatedContextPartySubIDs }

type NoRelatedPartyAltIDs struct{ quickfix.NumInGroupValue }

func (f NoRelatedPartyAltIDs) Tag() tag.Tag { return tag.NoRelatedPartyAltIDs }

type NoRelatedPartyAltSubIDs struct{ quickfix.NumInGroupValue }

func (f NoRelatedPartyAltSubIDs) Tag() tag.Tag { return tag.NoRelatedPartyAltSubIDs }

type NoRelatedPartyIDs struct{ quickfix.NumInGroupValue }

func (f NoRelatedPartyIDs) Tag() tag.Tag { return tag.NoRelatedPartyIDs }

type NoRelatedPartySubIDs struct{ quickfix.NumInGroupValue }

func (f NoRelatedPartySubIDs) Tag() tag.Tag { return tag.NoRelatedPartySubIDs }

type NoRelatedSym struct{ quickfix.NumInGroupValue }

func (f NoRelatedSym) Tag() tag.Tag { return tag.NoRelatedSym }

type NoRelationshipRiskInstruments struct{ quickfix.NumInGroupValue }

func (f NoRelationshipRiskInstruments) Tag() tag.Tag { return tag.NoRelationshipRiskInstruments }

type NoRelationshipRiskLimits struct{ quickfix.NumInGroupValue }

func (f NoRelationshipRiskLimits) Tag() tag.Tag { return tag.NoRelationshipRiskLimits }

type NoRelationshipRiskSecurityAltID struct{ quickfix.NumInGroupValue }

func (f NoRelationshipRiskSecurityAltID) Tag() tag.Tag { return tag.NoRelationshipRiskSecurityAltID }

type NoRelationshipRiskWarningLevels struct{ quickfix.NumInGroupValue }

func (f NoRelationshipRiskWarningLevels) Tag() tag.Tag { return tag.NoRelationshipRiskWarningLevels }

type NoRequestedPartyRoles struct{ quickfix.NumInGroupValue }

func (f NoRequestedPartyRoles) Tag() tag.Tag { return tag.NoRequestedPartyRoles }

type NoRiskInstruments struct{ quickfix.NumInGroupValue }

func (f NoRiskInstruments) Tag() tag.Tag { return tag.NoRiskInstruments }

type NoRiskLimits struct{ quickfix.NumInGroupValue }

func (f NoRiskLimits) Tag() tag.Tag { return tag.NoRiskLimits }

type NoRiskSecurityAltID struct{ quickfix.NumInGroupValue }

func (f NoRiskSecurityAltID) Tag() tag.Tag { return tag.NoRiskSecurityAltID }

type NoRiskWarningLevels struct{ quickfix.NumInGroupValue }

func (f NoRiskWarningLevels) Tag() tag.Tag { return tag.NoRiskWarningLevels }

type NoRootPartyIDs struct{ quickfix.NumInGroupValue }

func (f NoRootPartyIDs) Tag() tag.Tag { return tag.NoRootPartyIDs }

type NoRootPartySubIDs struct{ quickfix.NumInGroupValue }

func (f NoRootPartySubIDs) Tag() tag.Tag { return tag.NoRootPartySubIDs }

type NoRoutingIDs struct{ quickfix.NumInGroupValue }

func (f NoRoutingIDs) Tag() tag.Tag { return tag.NoRoutingIDs }

type NoRpts struct{ quickfix.IntValue }

func (f NoRpts) Tag() tag.Tag { return tag.NoRpts }

type NoSecurityAltID struct{ quickfix.NumInGroupValue }

func (f NoSecurityAltID) Tag() tag.Tag { return tag.NoSecurityAltID }

type NoSecurityTypes struct{ quickfix.NumInGroupValue }

func (f NoSecurityTypes) Tag() tag.Tag { return tag.NoSecurityTypes }

type NoSettlDetails struct{ quickfix.NumInGroupValue }

func (f NoSettlDetails) Tag() tag.Tag { return tag.NoSettlDetails }

type NoSettlInst struct{ quickfix.NumInGroupValue }

func (f NoSettlInst) Tag() tag.Tag { return tag.NoSettlInst }

type NoSettlOblig struct{ quickfix.NumInGroupValue }

func (f NoSettlOblig) Tag() tag.Tag { return tag.NoSettlOblig }

type NoSettlPartyIDs struct{ quickfix.NumInGroupValue }

func (f NoSettlPartyIDs) Tag() tag.Tag { return tag.NoSettlPartyIDs }

type NoSettlPartySubIDs struct{ quickfix.NumInGroupValue }

func (f NoSettlPartySubIDs) Tag() tag.Tag { return tag.NoSettlPartySubIDs }

type NoSideTrdRegTS struct{ quickfix.NumInGroupValue }

func (f NoSideTrdRegTS) Tag() tag.Tag { return tag.NoSideTrdRegTS }

type NoSides struct{ quickfix.NumInGroupValue }

func (f NoSides) Tag() tag.Tag { return tag.NoSides }

type NoStatsIndicators struct{ quickfix.NumInGroupValue }

func (f NoStatsIndicators) Tag() tag.Tag { return tag.NoStatsIndicators }

type NoStipulations struct{ quickfix.NumInGroupValue }

func (f NoStipulations) Tag() tag.Tag { return tag.NoStipulations }

type NoStrategyParameters struct{ quickfix.NumInGroupValue }

func (f NoStrategyParameters) Tag() tag.Tag { return tag.NoStrategyParameters }

type NoStrikeRules struct{ quickfix.NumInGroupValue }

func (f NoStrikeRules) Tag() tag.Tag { return tag.NoStrikeRules }

type NoStrikes struct{ quickfix.NumInGroupValue }

func (f NoStrikes) Tag() tag.Tag { return tag.NoStrikes }

type NoTargetPartyIDs struct{ quickfix.NumInGroupValue }

func (f NoTargetPartyIDs) Tag() tag.Tag { return tag.NoTargetPartyIDs }

type NoTickRules struct{ quickfix.NumInGroupValue }

func (f NoTickRules) Tag() tag.Tag { return tag.NoTickRules }

type NoTimeInForceRules struct{ quickfix.NumInGroupValue }

func (f NoTimeInForceRules) Tag() tag.Tag { return tag.NoTimeInForceRules }

type NoTrades struct{ quickfix.NumInGroupValue }

func (f NoTrades) Tag() tag.Tag { return tag.NoTrades }

type NoTradingSessionRules struct{ quickfix.NumInGroupValue }

func (f NoTradingSessionRules) Tag() tag.Tag { return tag.NoTradingSessionRules }

type NoTradingSessions struct{ quickfix.NumInGroupValue }

func (f NoTradingSessions) Tag() tag.Tag { return tag.NoTradingSessions }

type NoTrdRegTimestamps struct{ quickfix.NumInGroupValue }

func (f NoTrdRegTimestamps) Tag() tag.Tag { return tag.NoTrdRegTimestamps }

type NoTrdRepIndicators struct{ quickfix.NumInGroupValue }

func (f NoTrdRepIndicators) Tag() tag.Tag { return tag.NoTrdRepIndicators }

type NoUnderlyingAmounts struct{ quickfix.NumInGroupValue }

func (f NoUnderlyingAmounts) Tag() tag.Tag { return tag.NoUnderlyingAmounts }

type NoUnderlyingLegSecurityAltID struct{ quickfix.NumInGroupValue }

func (f NoUnderlyingLegSecurityAltID) Tag() tag.Tag { return tag.NoUnderlyingLegSecurityAltID }

type NoUnderlyingSecurityAltID struct{ quickfix.NumInGroupValue }

func (f NoUnderlyingSecurityAltID) Tag() tag.Tag { return tag.NoUnderlyingSecurityAltID }

type NoUnderlyingStips struct{ quickfix.NumInGroupValue }

func (f NoUnderlyingStips) Tag() tag.Tag { return tag.NoUnderlyingStips }

type NoUnderlyings struct{ quickfix.NumInGroupValue }

func (f NoUnderlyings) Tag() tag.Tag { return tag.NoUnderlyings }

type NoUndlyInstrumentParties struct{ quickfix.NumInGroupValue }

func (f NoUndlyInstrumentParties) Tag() tag.Tag { return tag.NoUndlyInstrumentParties }

type NoUndlyInstrumentPartySubIDs struct{ quickfix.NumInGroupValue }

func (f NoUndlyInstrumentPartySubIDs) Tag() tag.Tag { return tag.NoUndlyInstrumentPartySubIDs }

type NotAffOrigClOrdID struct{ quickfix.StringValue }

func (f NotAffOrigClOrdID) Tag() tag.Tag { return tag.NotAffOrigClOrdID }

type NotAffectedOrderID struct{ quickfix.StringValue }

func (f NotAffectedOrderID) Tag() tag.Tag { return tag.NotAffectedOrderID }

type NotifyBrokerOfCredit struct{ quickfix.BooleanValue }

func (f NotifyBrokerOfCredit) Tag() tag.Tag { return tag.NotifyBrokerOfCredit }

type NotionalPercentageOutstanding struct{ quickfix.PercentageValue }

func (f NotionalPercentageOutstanding) Tag() tag.Tag { return tag.NotionalPercentageOutstanding }

type NumBidders struct{ quickfix.IntValue }

func (f NumBidders) Tag() tag.Tag { return tag.NumBidders }

type NumDaysInterest struct{ quickfix.IntValue }

func (f NumDaysInterest) Tag() tag.Tag { return tag.NumDaysInterest }

type NumTickets struct{ quickfix.IntValue }

func (f NumTickets) Tag() tag.Tag { return tag.NumTickets }

type NumberOfOrders struct{ quickfix.IntValue }

func (f NumberOfOrders) Tag() tag.Tag { return tag.NumberOfOrders }

type OddLot struct{ quickfix.BooleanValue }

func (f OddLot) Tag() tag.Tag { return tag.OddLot }

type OfferForwardPoints struct{ quickfix.PriceOffsetValue }

func (f OfferForwardPoints) Tag() tag.Tag { return tag.OfferForwardPoints }

type OfferForwardPoints2 struct{ quickfix.PriceOffsetValue }

func (f OfferForwardPoints2) Tag() tag.Tag { return tag.OfferForwardPoints2 }

type OfferPx struct{ quickfix.PriceValue }

func (f OfferPx) Tag() tag.Tag { return tag.OfferPx }

type OfferSize struct{ quickfix.QtyValue }

func (f OfferSize) Tag() tag.Tag { return tag.OfferSize }

type OfferSpotRate struct{ quickfix.PriceValue }

func (f OfferSpotRate) Tag() tag.Tag { return tag.OfferSpotRate }

type OfferSwapPoints struct{ quickfix.PriceOffsetValue }

func (f OfferSwapPoints) Tag() tag.Tag { return tag.OfferSwapPoints }

type OfferYield struct{ quickfix.PercentageValue }

func (f OfferYield) Tag() tag.Tag { return tag.OfferYield }

type OnBehalfOfCompID struct{ quickfix.StringValue }

func (f OnBehalfOfCompID) Tag() tag.Tag { return tag.OnBehalfOfCompID }

type OnBehalfOfLocationID struct{ quickfix.StringValue }

func (f OnBehalfOfLocationID) Tag() tag.Tag { return tag.OnBehalfOfLocationID }

type OnBehalfOfSendingTime struct{ quickfix.UTCTimestampValue }

func (f OnBehalfOfSendingTime) Tag() tag.Tag { return tag.OnBehalfOfSendingTime }

type OnBehalfOfSubID struct{ quickfix.StringValue }

func (f OnBehalfOfSubID) Tag() tag.Tag { return tag.OnBehalfOfSubID }

type OpenClose struct{ quickfix.CharValue }

func (f OpenClose) Tag() tag.Tag { return tag.OpenClose }

type OpenCloseSettlFlag struct{ quickfix.MultipleCharValue }

func (f OpenCloseSettlFlag) Tag() tag.Tag { return tag.OpenCloseSettlFlag }

type OpenCloseSettleFlag struct{ quickfix.MultipleStringValue }

func (f OpenCloseSettleFlag) Tag() tag.Tag { return tag.OpenCloseSettleFlag }

type OpenInterest struct{ quickfix.AmtValue }

func (f OpenInterest) Tag() tag.Tag { return tag.OpenInterest }

type OptAttribute struct{ quickfix.CharValue }

func (f OptAttribute) Tag() tag.Tag { return tag.OptAttribute }

type OptPayAmount struct{ quickfix.AmtValue }

func (f OptPayAmount) Tag() tag.Tag { return tag.OptPayAmount }

type OptPayoutAmount struct{ quickfix.AmtValue }

func (f OptPayoutAmount) Tag() tag.Tag { return tag.OptPayoutAmount }

type OptPayoutType struct{ quickfix.IntValue }

func (f OptPayoutType) Tag() tag.Tag { return tag.OptPayoutType }

type OrdRejReason struct{ quickfix.IntValue }

func (f OrdRejReason) Tag() tag.Tag { return tag.OrdRejReason }

type OrdStatus struct{ quickfix.CharValue }

func (f OrdStatus) Tag() tag.Tag { return tag.OrdStatus }

type OrdStatusReqID struct{ quickfix.StringValue }

func (f OrdStatusReqID) Tag() tag.Tag { return tag.OrdStatusReqID }

type OrdType struct{ quickfix.CharValue }

func (f OrdType) Tag() tag.Tag { return tag.OrdType }

type OrderAvgPx struct{ quickfix.PriceValue }

func (f OrderAvgPx) Tag() tag.Tag { return tag.OrderAvgPx }

type OrderBookingQty struct{ quickfix.QtyValue }

func (f OrderBookingQty) Tag() tag.Tag { return tag.OrderBookingQty }

type OrderCapacity struct{ quickfix.CharValue }

func (f OrderCapacity) Tag() tag.Tag { return tag.OrderCapacity }

type OrderCapacityQty struct{ quickfix.QtyValue }

func (f OrderCapacityQty) Tag() tag.Tag { return tag.OrderCapacityQty }

type OrderCategory struct{ quickfix.CharValue }

func (f OrderCategory) Tag() tag.Tag { return tag.OrderCategory }

type OrderDelay struct{ quickfix.IntValue }

func (f OrderDelay) Tag() tag.Tag { return tag.OrderDelay }

type OrderDelayUnit struct{ quickfix.IntValue }

func (f OrderDelayUnit) Tag() tag.Tag { return tag.OrderDelayUnit }

type OrderHandlingInstSource struct{ quickfix.IntValue }

func (f OrderHandlingInstSource) Tag() tag.Tag { return tag.OrderHandlingInstSource }

type OrderID struct{ quickfix.StringValue }

func (f OrderID) Tag() tag.Tag { return tag.OrderID }

type OrderInputDevice struct{ quickfix.StringValue }

func (f OrderInputDevice) Tag() tag.Tag { return tag.OrderInputDevice }

type OrderPercent struct{ quickfix.PercentageValue }

func (f OrderPercent) Tag() tag.Tag { return tag.OrderPercent }

type OrderQty struct{ quickfix.QtyValue }

func (f OrderQty) Tag() tag.Tag { return tag.OrderQty }

type OrderQty2 struct{ quickfix.QtyValue }

func (f OrderQty2) Tag() tag.Tag { return tag.OrderQty2 }

type OrderRestrictions struct{ quickfix.MultipleCharValue }

func (f OrderRestrictions) Tag() tag.Tag { return tag.OrderRestrictions }

type OrigClOrdID struct{ quickfix.StringValue }

func (f OrigClOrdID) Tag() tag.Tag { return tag.OrigClOrdID }

type OrigCrossID struct{ quickfix.StringValue }

func (f OrigCrossID) Tag() tag.Tag { return tag.OrigCrossID }

type OrigCustOrderCapacity struct{ quickfix.IntValue }

func (f OrigCustOrderCapacity) Tag() tag.Tag { return tag.OrigCustOrderCapacity }

type OrigOrdModTime struct{ quickfix.UTCTimestampValue }

func (f OrigOrdModTime) Tag() tag.Tag { return tag.OrigOrdModTime }

type OrigPosReqRefID struct{ quickfix.StringValue }

func (f OrigPosReqRefID) Tag() tag.Tag { return tag.OrigPosReqRefID }

type OrigSecondaryTradeID struct{ quickfix.StringValue }

func (f OrigSecondaryTradeID) Tag() tag.Tag { return tag.OrigSecondaryTradeID }

type OrigSendingTime struct{ quickfix.UTCTimestampValue }

func (f OrigSendingTime) Tag() tag.Tag { return tag.OrigSendingTime }

type OrigTime struct{ quickfix.UTCTimestampValue }

func (f OrigTime) Tag() tag.Tag { return tag.OrigTime }

type OrigTradeDate struct{ quickfix.LocalMktDateValue }

func (f OrigTradeDate) Tag() tag.Tag { return tag.OrigTradeDate }

type OrigTradeHandlingInstr struct{ quickfix.CharValue }

func (f OrigTradeHandlingInstr) Tag() tag.Tag { return tag.OrigTradeHandlingInstr }

type OrigTradeID struct{ quickfix.StringValue }

func (f OrigTradeID) Tag() tag.Tag { return tag.OrigTradeID }

type OriginalNotionalPercentageOutstanding struct{ quickfix.PercentageValue }

func (f OriginalNotionalPercentageOutstanding) Tag() tag.Tag {
	return tag.OriginalNotionalPercentageOutstanding
}

type OutMainCntryUIndex struct{ quickfix.AmtValue }

func (f OutMainCntryUIndex) Tag() tag.Tag { return tag.OutMainCntryUIndex }

type OutsideIndexPct struct{ quickfix.PercentageValue }

func (f OutsideIndexPct) Tag() tag.Tag { return tag.OutsideIndexPct }

type OwnerType struct{ quickfix.IntValue }

func (f OwnerType) Tag() tag.Tag { return tag.OwnerType }

type OwnershipType struct{ quickfix.CharValue }

func (f OwnershipType) Tag() tag.Tag { return tag.OwnershipType }

type ParentMktSegmID struct{ quickfix.StringValue }

func (f ParentMktSegmID) Tag() tag.Tag { return tag.ParentMktSegmID }

type ParticipationRate struct{ quickfix.PercentageValue }

func (f ParticipationRate) Tag() tag.Tag { return tag.ParticipationRate }

type PartyAltID struct{ quickfix.StringValue }

func (f PartyAltID) Tag() tag.Tag { return tag.PartyAltID }

type PartyAltIDSource struct{ quickfix.CharValue }

func (f PartyAltIDSource) Tag() tag.Tag { return tag.PartyAltIDSource }

type PartyAltSubID struct{ quickfix.StringValue }

func (f PartyAltSubID) Tag() tag.Tag { return tag.PartyAltSubID }

type PartyAltSubIDType struct{ quickfix.IntValue }

func (f PartyAltSubIDType) Tag() tag.Tag { return tag.PartyAltSubIDType }

type PartyDetailsListReportID struct{ quickfix.StringValue }

func (f PartyDetailsListReportID) Tag() tag.Tag { return tag.PartyDetailsListReportID }

type PartyDetailsListRequestID struct{ quickfix.StringValue }

func (f PartyDetailsListRequestID) Tag() tag.Tag { return tag.PartyDetailsListRequestID }

type PartyDetailsRequestResult struct{ quickfix.IntValue }

func (f PartyDetailsRequestResult) Tag() tag.Tag { return tag.PartyDetailsRequestResult }

type PartyID struct{ quickfix.StringValue }

func (f PartyID) Tag() tag.Tag { return tag.PartyID }

type PartyIDSource struct{ quickfix.CharValue }

func (f PartyIDSource) Tag() tag.Tag { return tag.PartyIDSource }

type PartyListResponseType struct{ quickfix.IntValue }

func (f PartyListResponseType) Tag() tag.Tag { return tag.PartyListResponseType }

type PartyRelationship struct{ quickfix.IntValue }

func (f PartyRelationship) Tag() tag.Tag { return tag.PartyRelationship }

type PartyRole struct{ quickfix.IntValue }

func (f PartyRole) Tag() tag.Tag { return tag.PartyRole }

type PartySubID struct{ quickfix.StringValue }

func (f PartySubID) Tag() tag.Tag { return tag.PartySubID }

type PartySubIDType struct{ quickfix.IntValue }

func (f PartySubIDType) Tag() tag.Tag { return tag.PartySubIDType }

type Password struct{ quickfix.StringValue }

func (f Password) Tag() tag.Tag { return tag.Password }

type PaymentDate struct{ quickfix.LocalMktDateValue }

func (f PaymentDate) Tag() tag.Tag { return tag.PaymentDate }

type PaymentMethod struct{ quickfix.IntValue }

func (f PaymentMethod) Tag() tag.Tag { return tag.PaymentMethod }

type PaymentRef struct{ quickfix.StringValue }

func (f PaymentRef) Tag() tag.Tag { return tag.PaymentRef }

type PaymentRemitterID struct{ quickfix.StringValue }

func (f PaymentRemitterID) Tag() tag.Tag { return tag.PaymentRemitterID }

type PctAtRisk struct{ quickfix.PercentageValue }

func (f PctAtRisk) Tag() tag.Tag { return tag.PctAtRisk }

type PegDifference struct{ quickfix.PriceOffsetValue }

func (f PegDifference) Tag() tag.Tag { return tag.PegDifference }

type PegLimitType struct{ quickfix.IntValue }

func (f PegLimitType) Tag() tag.Tag { return tag.PegLimitType }

type PegMoveType struct{ quickfix.IntValue }

func (f PegMoveType) Tag() tag.Tag { return tag.PegMoveType }

type PegOffsetType struct{ quickfix.IntValue }

func (f PegOffsetType) Tag() tag.Tag { return tag.PegOffsetType }

type PegOffsetValue struct{ quickfix.FloatValue }

func (f PegOffsetValue) Tag() tag.Tag { return tag.PegOffsetValue }

type PegPriceType struct{ quickfix.IntValue }

func (f PegPriceType) Tag() tag.Tag { return tag.PegPriceType }

type PegRoundDirection struct{ quickfix.IntValue }

func (f PegRoundDirection) Tag() tag.Tag { return tag.PegRoundDirection }

type PegScope struct{ quickfix.IntValue }

func (f PegScope) Tag() tag.Tag { return tag.PegScope }

type PegSecurityDesc struct{ quickfix.StringValue }

func (f PegSecurityDesc) Tag() tag.Tag { return tag.PegSecurityDesc }

type PegSecurityID struct{ quickfix.StringValue }

func (f PegSecurityID) Tag() tag.Tag { return tag.PegSecurityID }

type PegSecurityIDSource struct{ quickfix.StringValue }

func (f PegSecurityIDSource) Tag() tag.Tag { return tag.PegSecurityIDSource }

type PegSymbol struct{ quickfix.StringValue }

func (f PegSymbol) Tag() tag.Tag { return tag.PegSymbol }

type PeggedPrice struct{ quickfix.PriceValue }

func (f PeggedPrice) Tag() tag.Tag { return tag.PeggedPrice }

type PeggedRefPrice struct{ quickfix.PriceValue }

func (f PeggedRefPrice) Tag() tag.Tag { return tag.PeggedRefPrice }

type Pool struct{ quickfix.StringValue }

func (f Pool) Tag() tag.Tag { return tag.Pool }

type PosAmt struct{ quickfix.AmtValue }

func (f PosAmt) Tag() tag.Tag { return tag.PosAmt }

type PosAmtType struct{ quickfix.StringValue }

func (f PosAmtType) Tag() tag.Tag { return tag.PosAmtType }

type PosMaintAction struct{ quickfix.IntValue }

func (f PosMaintAction) Tag() tag.Tag { return tag.PosMaintAction }

type PosMaintResult struct{ quickfix.IntValue }

func (f PosMaintResult) Tag() tag.Tag { return tag.PosMaintResult }

type PosMaintRptID struct{ quickfix.StringValue }

func (f PosMaintRptID) Tag() tag.Tag { return tag.PosMaintRptID }

type PosMaintRptRefID struct{ quickfix.StringValue }

func (f PosMaintRptRefID) Tag() tag.Tag { return tag.PosMaintRptRefID }

type PosMaintStatus struct{ quickfix.IntValue }

func (f PosMaintStatus) Tag() tag.Tag { return tag.PosMaintStatus }

type PosQtyStatus struct{ quickfix.IntValue }

func (f PosQtyStatus) Tag() tag.Tag { return tag.PosQtyStatus }

type PosReqID struct{ quickfix.StringValue }

func (f PosReqID) Tag() tag.Tag { return tag.PosReqID }

type PosReqResult struct{ quickfix.IntValue }

func (f PosReqResult) Tag() tag.Tag { return tag.PosReqResult }

type PosReqStatus struct{ quickfix.IntValue }

func (f PosReqStatus) Tag() tag.Tag { return tag.PosReqStatus }

type PosReqType struct{ quickfix.IntValue }

func (f PosReqType) Tag() tag.Tag { return tag.PosReqType }

type PosTransType struct{ quickfix.IntValue }

func (f PosTransType) Tag() tag.Tag { return tag.PosTransType }

type PosType struct{ quickfix.StringValue }

func (f PosType) Tag() tag.Tag { return tag.PosType }

type PositionCurrency struct{ quickfix.StringValue }

func (f PositionCurrency) Tag() tag.Tag { return tag.PositionCurrency }

type PositionEffect struct{ quickfix.CharValue }

func (f PositionEffect) Tag() tag.Tag { return tag.PositionEffect }

type PositionLimit struct{ quickfix.IntValue }

func (f PositionLimit) Tag() tag.Tag { return tag.PositionLimit }

type PossDupFlag struct{ quickfix.BooleanValue }

func (f PossDupFlag) Tag() tag.Tag { return tag.PossDupFlag }

type PossResend struct{ quickfix.BooleanValue }

func (f PossResend) Tag() tag.Tag { return tag.PossResend }

type PreTradeAnonymity struct{ quickfix.BooleanValue }

func (f PreTradeAnonymity) Tag() tag.Tag { return tag.PreTradeAnonymity }

type PreallocMethod struct{ quickfix.CharValue }

func (f PreallocMethod) Tag() tag.Tag { return tag.PreallocMethod }

type PrevClosePx struct{ quickfix.PriceValue }

func (f PrevClosePx) Tag() tag.Tag { return tag.PrevClosePx }

type PreviouslyReported struct{ quickfix.BooleanValue }

func (f PreviouslyReported) Tag() tag.Tag { return tag.PreviouslyReported }

type Price struct{ quickfix.PriceValue }

func (f Price) Tag() tag.Tag { return tag.Price }

type Price2 struct{ quickfix.PriceValue }

func (f Price2) Tag() tag.Tag { return tag.Price2 }

type PriceDelta struct{ quickfix.FloatValue }

func (f PriceDelta) Tag() tag.Tag { return tag.PriceDelta }

type PriceImprovement struct{ quickfix.PriceOffsetValue }

func (f PriceImprovement) Tag() tag.Tag { return tag.PriceImprovement }

type PriceLimitType struct{ quickfix.IntValue }

func (f PriceLimitType) Tag() tag.Tag { return tag.PriceLimitType }

type PriceProtectionScope struct{ quickfix.CharValue }

func (f PriceProtectionScope) Tag() tag.Tag { return tag.PriceProtectionScope }

type PriceQuoteMethod struct{ quickfix.StringValue }

func (f PriceQuoteMethod) Tag() tag.Tag { return tag.PriceQuoteMethod }

type PriceType struct{ quickfix.IntValue }

func (f PriceType) Tag() tag.Tag { return tag.PriceType }

type PriceUnitOfMeasure struct{ quickfix.StringValue }

func (f PriceUnitOfMeasure) Tag() tag.Tag { return tag.PriceUnitOfMeasure }

type PriceUnitOfMeasureQty struct{ quickfix.QtyValue }

func (f PriceUnitOfMeasureQty) Tag() tag.Tag { return tag.PriceUnitOfMeasureQty }

type PriorSettlPrice struct{ quickfix.PriceValue }

func (f PriorSettlPrice) Tag() tag.Tag { return tag.PriorSettlPrice }

type PriorSpreadIndicator struct{ quickfix.BooleanValue }

func (f PriorSpreadIndicator) Tag() tag.Tag { return tag.PriorSpreadIndicator }

type PriorityIndicator struct{ quickfix.IntValue }

func (f PriorityIndicator) Tag() tag.Tag { return tag.PriorityIndicator }

type PrivateQuote struct{ quickfix.BooleanValue }

func (f PrivateQuote) Tag() tag.Tag { return tag.PrivateQuote }

type ProcessCode struct{ quickfix.CharValue }

func (f ProcessCode) Tag() tag.Tag { return tag.ProcessCode }

type Product struct{ quickfix.IntValue }

func (f Product) Tag() tag.Tag { return tag.Product }

type ProductComplex struct{ quickfix.StringValue }

func (f ProductComplex) Tag() tag.Tag { return tag.ProductComplex }

type ProgPeriodInterval struct{ quickfix.IntValue }

func (f ProgPeriodInterval) Tag() tag.Tag { return tag.ProgPeriodInterval }

type ProgRptReqs struct{ quickfix.IntValue }

func (f ProgRptReqs) Tag() tag.Tag { return tag.ProgRptReqs }

type PublishTrdIndicator struct{ quickfix.BooleanValue }

func (f PublishTrdIndicator) Tag() tag.Tag { return tag.PublishTrdIndicator }

type PutOrCall struct{ quickfix.IntValue }

func (f PutOrCall) Tag() tag.Tag { return tag.PutOrCall }

type QtyType struct{ quickfix.IntValue }

func (f QtyType) Tag() tag.Tag { return tag.QtyType }

type Quantity struct{ quickfix.QtyValue }

func (f Quantity) Tag() tag.Tag { return tag.Quantity }

type QuantityDate struct{ quickfix.LocalMktDateValue }

func (f QuantityDate) Tag() tag.Tag { return tag.QuantityDate }

type QuantityType struct{ quickfix.IntValue }

func (f QuantityType) Tag() tag.Tag { return tag.QuantityType }

type QuoteAckStatus struct{ quickfix.IntValue }

func (f QuoteAckStatus) Tag() tag.Tag { return tag.QuoteAckStatus }

type QuoteCancelType struct{ quickfix.IntValue }

func (f QuoteCancelType) Tag() tag.Tag { return tag.QuoteCancelType }

type QuoteCondition struct{ quickfix.MultipleStringValue }

func (f QuoteCondition) Tag() tag.Tag { return tag.QuoteCondition }

type QuoteEntryID struct{ quickfix.StringValue }

func (f QuoteEntryID) Tag() tag.Tag { return tag.QuoteEntryID }

type QuoteEntryRejectReason struct{ quickfix.IntValue }

func (f QuoteEntryRejectReason) Tag() tag.Tag { return tag.QuoteEntryRejectReason }

type QuoteEntryStatus struct{ quickfix.IntValue }

func (f QuoteEntryStatus) Tag() tag.Tag { return tag.QuoteEntryStatus }

type QuoteID struct{ quickfix.StringValue }

func (f QuoteID) Tag() tag.Tag { return tag.QuoteID }

type QuoteMsgID struct{ quickfix.StringValue }

func (f QuoteMsgID) Tag() tag.Tag { return tag.QuoteMsgID }

type QuotePriceType struct{ quickfix.IntValue }

func (f QuotePriceType) Tag() tag.Tag { return tag.QuotePriceType }

type QuoteQualifier struct{ quickfix.CharValue }

func (f QuoteQualifier) Tag() tag.Tag { return tag.QuoteQualifier }

type QuoteRejectReason struct{ quickfix.IntValue }

func (f QuoteRejectReason) Tag() tag.Tag { return tag.QuoteRejectReason }

type QuoteReqID struct{ quickfix.StringValue }

func (f QuoteReqID) Tag() tag.Tag { return tag.QuoteReqID }

type QuoteRequestRejectReason struct{ quickfix.IntValue }

func (f QuoteRequestRejectReason) Tag() tag.Tag { return tag.QuoteRequestRejectReason }

type QuoteRequestType struct{ quickfix.IntValue }

func (f QuoteRequestType) Tag() tag.Tag { return tag.QuoteRequestType }

type QuoteRespID struct{ quickfix.StringValue }

func (f QuoteRespID) Tag() tag.Tag { return tag.QuoteRespID }

type QuoteRespType struct{ quickfix.IntValue }

func (f QuoteRespType) Tag() tag.Tag { return tag.QuoteRespType }

type QuoteResponseLevel struct{ quickfix.IntValue }

func (f QuoteResponseLevel) Tag() tag.Tag { return tag.QuoteResponseLevel }

type QuoteSetID struct{ quickfix.StringValue }

func (f QuoteSetID) Tag() tag.Tag { return tag.QuoteSetID }

type QuoteSetValidUntilTime struct{ quickfix.UTCTimestampValue }

func (f QuoteSetValidUntilTime) Tag() tag.Tag { return tag.QuoteSetValidUntilTime }

type QuoteStatus struct{ quickfix.IntValue }

func (f QuoteStatus) Tag() tag.Tag { return tag.QuoteStatus }

type QuoteStatusReqID struct{ quickfix.StringValue }

func (f QuoteStatusReqID) Tag() tag.Tag { return tag.QuoteStatusReqID }

type QuoteType struct{ quickfix.IntValue }

func (f QuoteType) Tag() tag.Tag { return tag.QuoteType }

type RFQReqID struct{ quickfix.StringValue }

func (f RFQReqID) Tag() tag.Tag { return tag.RFQReqID }

type RateSource struct{ quickfix.IntValue }

func (f RateSource) Tag() tag.Tag { return tag.RateSource }

type RateSourceType struct{ quickfix.IntValue }

func (f RateSourceType) Tag() tag.Tag { return tag.RateSourceType }

type RatioQty struct{ quickfix.QtyValue }

func (f RatioQty) Tag() tag.Tag { return tag.RatioQty }

type RawData struct{ quickfix.DataValue }

func (f RawData) Tag() tag.Tag { return tag.RawData }

type RawDataLength struct{ quickfix.LengthValue }

func (f RawDataLength) Tag() tag.Tag { return tag.RawDataLength }

type ReceivedDeptID struct{ quickfix.StringValue }

func (f ReceivedDeptID) Tag() tag.Tag { return tag.ReceivedDeptID }

type RedemptionDate struct{ quickfix.LocalMktDateValue }

func (f RedemptionDate) Tag() tag.Tag { return tag.RedemptionDate }

type RefAllocID struct{ quickfix.StringValue }

func (f RefAllocID) Tag() tag.Tag { return tag.RefAllocID }

type RefApplExtID struct{ quickfix.IntValue }

func (f RefApplExtID) Tag() tag.Tag { return tag.RefApplExtID }

type RefApplID struct{ quickfix.StringValue }

func (f RefApplID) Tag() tag.Tag { return tag.RefApplID }

type RefApplLastSeqNum struct{ quickfix.SeqNumValue }

func (f RefApplLastSeqNum) Tag() tag.Tag { return tag.RefApplLastSeqNum }

type RefApplReqID struct{ quickfix.StringValue }

func (f RefApplReqID) Tag() tag.Tag { return tag.RefApplReqID }

type RefApplVerID struct{ quickfix.StringValue }

func (f RefApplVerID) Tag() tag.Tag { return tag.RefApplVerID }

type RefCompID struct{ quickfix.StringValue }

func (f RefCompID) Tag() tag.Tag { return tag.RefCompID }

type RefCstmApplVerID struct{ quickfix.StringValue }

func (f RefCstmApplVerID) Tag() tag.Tag { return tag.RefCstmApplVerID }

type RefMsgType struct{ quickfix.StringValue }

func (f RefMsgType) Tag() tag.Tag { return tag.RefMsgType }

type RefOrdIDReason struct{ quickfix.IntValue }

func (f RefOrdIDReason) Tag() tag.Tag { return tag.RefOrdIDReason }

type RefOrderID struct{ quickfix.StringValue }

func (f RefOrderID) Tag() tag.Tag { return tag.RefOrderID }

type RefOrderIDSource struct{ quickfix.CharValue }

func (f RefOrderIDSource) Tag() tag.Tag { return tag.RefOrderIDSource }

type RefSeqNum struct{ quickfix.SeqNumValue }

func (f RefSeqNum) Tag() tag.Tag { return tag.RefSeqNum }

type RefSubID struct{ quickfix.StringValue }

func (f RefSubID) Tag() tag.Tag { return tag.RefSubID }

type RefTagID struct{ quickfix.IntValue }

func (f RefTagID) Tag() tag.Tag { return tag.RefTagID }

type ReferencePage struct{ quickfix.StringValue }

func (f ReferencePage) Tag() tag.Tag { return tag.ReferencePage }

type RefreshIndicator struct{ quickfix.BooleanValue }

func (f RefreshIndicator) Tag() tag.Tag { return tag.RefreshIndicator }

type RefreshQty struct{ quickfix.QtyValue }

func (f RefreshQty) Tag() tag.Tag { return tag.RefreshQty }

type RegistAcctType struct{ quickfix.StringValue }

func (f RegistAcctType) Tag() tag.Tag { return tag.RegistAcctType }

type RegistDetls struct{ quickfix.StringValue }

func (f RegistDetls) Tag() tag.Tag { return tag.RegistDetls }

type RegistDtls struct{ quickfix.StringValue }

func (f RegistDtls) Tag() tag.Tag { return tag.RegistDtls }

type RegistEmail struct{ quickfix.StringValue }

func (f RegistEmail) Tag() tag.Tag { return tag.RegistEmail }

type RegistID struct{ quickfix.StringValue }

func (f RegistID) Tag() tag.Tag { return tag.RegistID }

type RegistRefID struct{ quickfix.StringValue }

func (f RegistRefID) Tag() tag.Tag { return tag.RegistRefID }

type RegistRejReasonCode struct{ quickfix.IntValue }

func (f RegistRejReasonCode) Tag() tag.Tag { return tag.RegistRejReasonCode }

type RegistRejReasonText struct{ quickfix.StringValue }

func (f RegistRejReasonText) Tag() tag.Tag { return tag.RegistRejReasonText }

type RegistStatus struct{ quickfix.CharValue }

func (f RegistStatus) Tag() tag.Tag { return tag.RegistStatus }

type RegistTransType struct{ quickfix.CharValue }

func (f RegistTransType) Tag() tag.Tag { return tag.RegistTransType }

type RejectText struct{ quickfix.StringValue }

func (f RejectText) Tag() tag.Tag { return tag.RejectText }

type RelSymTransactTime struct{ quickfix.UTCTimestampValue }

func (f RelSymTransactTime) Tag() tag.Tag { return tag.RelSymTransactTime }

type RelatdSym struct{ quickfix.StringValue }

func (f RelatdSym) Tag() tag.Tag { return tag.RelatdSym }

type RelatedContextPartyID struct{ quickfix.StringValue }

func (f RelatedContextPartyID) Tag() tag.Tag { return tag.RelatedContextPartyID }

type RelatedContextPartyIDSource struct{ quickfix.CharValue }

func (f RelatedContextPartyIDSource) Tag() tag.Tag { return tag.RelatedContextPartyIDSource }

type RelatedContextPartyRole struct{ quickfix.IntValue }

func (f RelatedContextPartyRole) Tag() tag.Tag { return tag.RelatedContextPartyRole }

type RelatedContextPartySubID struct{ quickfix.StringValue }

func (f RelatedContextPartySubID) Tag() tag.Tag { return tag.RelatedContextPartySubID }

type RelatedContextPartySubIDType struct{ quickfix.IntValue }

func (f RelatedContextPartySubIDType) Tag() tag.Tag { return tag.RelatedContextPartySubIDType }

type RelatedPartyAltID struct{ quickfix.StringValue }

func (f RelatedPartyAltID) Tag() tag.Tag { return tag.RelatedPartyAltID }

type RelatedPartyAltIDSource struct{ quickfix.CharValue }

func (f RelatedPartyAltIDSource) Tag() tag.Tag { return tag.RelatedPartyAltIDSource }

type RelatedPartyAltSubID struct{ quickfix.StringValue }

func (f RelatedPartyAltSubID) Tag() tag.Tag { return tag.RelatedPartyAltSubID }

type RelatedPartyAltSubIDType struct{ quickfix.IntValue }

func (f RelatedPartyAltSubIDType) Tag() tag.Tag { return tag.RelatedPartyAltSubIDType }

type RelatedPartyID struct{ quickfix.StringValue }

func (f RelatedPartyID) Tag() tag.Tag { return tag.RelatedPartyID }

type RelatedPartyIDSource struct{ quickfix.CharValue }

func (f RelatedPartyIDSource) Tag() tag.Tag { return tag.RelatedPartyIDSource }

type RelatedPartyRole struct{ quickfix.IntValue }

func (f RelatedPartyRole) Tag() tag.Tag { return tag.RelatedPartyRole }

type RelatedPartySubID struct{ quickfix.StringValue }

func (f RelatedPartySubID) Tag() tag.Tag { return tag.RelatedPartySubID }

type RelatedPartySubIDType struct{ quickfix.IntValue }

func (f RelatedPartySubIDType) Tag() tag.Tag { return tag.RelatedPartySubIDType }

type RelationshipRiskCFICode struct{ quickfix.StringValue }

func (f RelationshipRiskCFICode) Tag() tag.Tag { return tag.RelationshipRiskCFICode }

type RelationshipRiskCouponRate struct{ quickfix.PercentageValue }

func (f RelationshipRiskCouponRate) Tag() tag.Tag { return tag.RelationshipRiskCouponRate }

type RelationshipRiskEncodedSecurityDesc struct{ quickfix.DataValue }

func (f RelationshipRiskEncodedSecurityDesc) Tag() tag.Tag {
	return tag.RelationshipRiskEncodedSecurityDesc
}

type RelationshipRiskEncodedSecurityDescLen struct{ quickfix.LengthValue }

func (f RelationshipRiskEncodedSecurityDescLen) Tag() tag.Tag {
	return tag.RelationshipRiskEncodedSecurityDescLen
}

type RelationshipRiskFlexibleIndicator struct{ quickfix.BooleanValue }

func (f RelationshipRiskFlexibleIndicator) Tag() tag.Tag { return tag.RelationshipRiskFlexibleIndicator }

type RelationshipRiskInstrumentMultiplier struct{ quickfix.FloatValue }

func (f RelationshipRiskInstrumentMultiplier) Tag() tag.Tag {
	return tag.RelationshipRiskInstrumentMultiplier
}

type RelationshipRiskInstrumentOperator struct{ quickfix.IntValue }

func (f RelationshipRiskInstrumentOperator) Tag() tag.Tag {
	return tag.RelationshipRiskInstrumentOperator
}

type RelationshipRiskInstrumentSettlType struct{ quickfix.StringValue }

func (f RelationshipRiskInstrumentSettlType) Tag() tag.Tag {
	return tag.RelationshipRiskInstrumentSettlType
}

type RelationshipRiskLimitAmount struct{ quickfix.AmtValue }

func (f RelationshipRiskLimitAmount) Tag() tag.Tag { return tag.RelationshipRiskLimitAmount }

type RelationshipRiskLimitCurrency struct{ quickfix.CurrencyValue }

func (f RelationshipRiskLimitCurrency) Tag() tag.Tag { return tag.RelationshipRiskLimitCurrency }

type RelationshipRiskLimitPlatform struct{ quickfix.StringValue }

func (f RelationshipRiskLimitPlatform) Tag() tag.Tag { return tag.RelationshipRiskLimitPlatform }

type RelationshipRiskLimitType struct{ quickfix.IntValue }

func (f RelationshipRiskLimitType) Tag() tag.Tag { return tag.RelationshipRiskLimitType }

type RelationshipRiskMaturityMonthYear struct{ quickfix.MonthYearValue }

func (f RelationshipRiskMaturityMonthYear) Tag() tag.Tag { return tag.RelationshipRiskMaturityMonthYear }

type RelationshipRiskMaturityTime struct{ quickfix.TZTimeOnlyValue }

func (f RelationshipRiskMaturityTime) Tag() tag.Tag { return tag.RelationshipRiskMaturityTime }

type RelationshipRiskProduct struct{ quickfix.IntValue }

func (f RelationshipRiskProduct) Tag() tag.Tag { return tag.RelationshipRiskProduct }

type RelationshipRiskProductComplex struct{ quickfix.StringValue }

func (f RelationshipRiskProductComplex) Tag() tag.Tag { return tag.RelationshipRiskProductComplex }

type RelationshipRiskPutOrCall struct{ quickfix.IntValue }

func (f RelationshipRiskPutOrCall) Tag() tag.Tag { return tag.RelationshipRiskPutOrCall }

type RelationshipRiskRestructuringType struct{ quickfix.StringValue }

func (f RelationshipRiskRestructuringType) Tag() tag.Tag { return tag.RelationshipRiskRestructuringType }

type RelationshipRiskSecurityAltID struct{ quickfix.StringValue }

func (f RelationshipRiskSecurityAltID) Tag() tag.Tag { return tag.RelationshipRiskSecurityAltID }

type RelationshipRiskSecurityAltIDSource struct{ quickfix.StringValue }

func (f RelationshipRiskSecurityAltIDSource) Tag() tag.Tag {
	return tag.RelationshipRiskSecurityAltIDSource
}

type RelationshipRiskSecurityDesc struct{ quickfix.StringValue }

func (f RelationshipRiskSecurityDesc) Tag() tag.Tag { return tag.RelationshipRiskSecurityDesc }

type RelationshipRiskSecurityExchange struct{ quickfix.ExchangeValue }

func (f RelationshipRiskSecurityExchange) Tag() tag.Tag { return tag.RelationshipRiskSecurityExchange }

type RelationshipRiskSecurityGroup struct{ quickfix.StringValue }

func (f RelationshipRiskSecurityGroup) Tag() tag.Tag { return tag.RelationshipRiskSecurityGroup }

type RelationshipRiskSecurityID struct{ quickfix.StringValue }

func (f RelationshipRiskSecurityID) Tag() tag.Tag { return tag.RelationshipRiskSecurityID }

type RelationshipRiskSecurityIDSource struct{ quickfix.StringValue }

func (f RelationshipRiskSecurityIDSource) Tag() tag.Tag { return tag.RelationshipRiskSecurityIDSource }

type RelationshipRiskSecuritySubType struct{ quickfix.StringValue }

func (f RelationshipRiskSecuritySubType) Tag() tag.Tag { return tag.RelationshipRiskSecuritySubType }

type RelationshipRiskSecurityType struct{ quickfix.StringValue }

func (f RelationshipRiskSecurityType) Tag() tag.Tag { return tag.RelationshipRiskSecurityType }

type RelationshipRiskSeniority struct{ quickfix.StringValue }

func (f RelationshipRiskSeniority) Tag() tag.Tag { return tag.RelationshipRiskSeniority }

type RelationshipRiskSymbol struct{ quickfix.StringValue }

func (f RelationshipRiskSymbol) Tag() tag.Tag { return tag.RelationshipRiskSymbol }

type RelationshipRiskSymbolSfx struct{ quickfix.StringValue }

func (f RelationshipRiskSymbolSfx) Tag() tag.Tag { return tag.RelationshipRiskSymbolSfx }

type RelationshipRiskWarningLevelName struct{ quickfix.StringValue }

func (f RelationshipRiskWarningLevelName) Tag() tag.Tag { return tag.RelationshipRiskWarningLevelName }

type RelationshipRiskWarningLevelPercent struct{ quickfix.PercentageValue }

func (f RelationshipRiskWarningLevelPercent) Tag() tag.Tag {
	return tag.RelationshipRiskWarningLevelPercent
}

type RepoCollateralSecurityType struct{ quickfix.IntValue }

func (f RepoCollateralSecurityType) Tag() tag.Tag { return tag.RepoCollateralSecurityType }

type ReportToExch struct{ quickfix.BooleanValue }

func (f ReportToExch) Tag() tag.Tag { return tag.ReportToExch }

type ReportedPx struct{ quickfix.PriceValue }

func (f ReportedPx) Tag() tag.Tag { return tag.ReportedPx }

type ReportedPxDiff struct{ quickfix.BooleanValue }

func (f ReportedPxDiff) Tag() tag.Tag { return tag.ReportedPxDiff }

type RepurchaseRate struct{ quickfix.PercentageValue }

func (f RepurchaseRate) Tag() tag.Tag { return tag.RepurchaseRate }

type RepurchaseTerm struct{ quickfix.IntValue }

func (f RepurchaseTerm) Tag() tag.Tag { return tag.RepurchaseTerm }

type RequestedPartyRole struct{ quickfix.IntValue }

func (f RequestedPartyRole) Tag() tag.Tag { return tag.RequestedPartyRole }

type ResetSeqNumFlag struct{ quickfix.BooleanValue }

func (f ResetSeqNumFlag) Tag() tag.Tag { return tag.ResetSeqNumFlag }

type RespondentType struct{ quickfix.IntValue }

func (f RespondentType) Tag() tag.Tag { return tag.RespondentType }

type ResponseDestination struct{ quickfix.StringValue }

func (f ResponseDestination) Tag() tag.Tag { return tag.ResponseDestination }

type ResponseTransportType struct{ quickfix.IntValue }

func (f ResponseTransportType) Tag() tag.Tag { return tag.ResponseTransportType }

type RestructuringType struct{ quickfix.StringValue }

func (f RestructuringType) Tag() tag.Tag { return tag.RestructuringType }

type ReversalIndicator struct{ quickfix.BooleanValue }

func (f ReversalIndicator) Tag() tag.Tag { return tag.ReversalIndicator }

type RiskCFICode struct{ quickfix.StringValue }

func (f RiskCFICode) Tag() tag.Tag { return tag.RiskCFICode }

type RiskCouponRate struct{ quickfix.PercentageValue }

func (f RiskCouponRate) Tag() tag.Tag { return tag.RiskCouponRate }

type RiskEncodedSecurityDesc struct{ quickfix.DataValue }

func (f RiskEncodedSecurityDesc) Tag() tag.Tag { return tag.RiskEncodedSecurityDesc }

type RiskEncodedSecurityDescLen struct{ quickfix.LengthValue }

func (f RiskEncodedSecurityDescLen) Tag() tag.Tag { return tag.RiskEncodedSecurityDescLen }

type RiskFlexibleIndicator struct{ quickfix.BooleanValue }

func (f RiskFlexibleIndicator) Tag() tag.Tag { return tag.RiskFlexibleIndicator }

type RiskFreeRate struct{ quickfix.FloatValue }

func (f RiskFreeRate) Tag() tag.Tag { return tag.RiskFreeRate }

type RiskInstrumentMultiplier struct{ quickfix.FloatValue }

func (f RiskInstrumentMultiplier) Tag() tag.Tag { return tag.RiskInstrumentMultiplier }

type RiskInstrumentOperator struct{ quickfix.IntValue }

func (f RiskInstrumentOperator) Tag() tag.Tag { return tag.RiskInstrumentOperator }

type RiskInstrumentSettlType struct{ quickfix.StringValue }

func (f RiskInstrumentSettlType) Tag() tag.Tag { return tag.RiskInstrumentSettlType }

type RiskLimitAmount struct{ quickfix.AmtValue }

func (f RiskLimitAmount) Tag() tag.Tag { return tag.RiskLimitAmount }

type RiskLimitCurrency struct{ quickfix.CurrencyValue }

func (f RiskLimitCurrency) Tag() tag.Tag { return tag.RiskLimitCurrency }

type RiskLimitPlatform struct{ quickfix.StringValue }

func (f RiskLimitPlatform) Tag() tag.Tag { return tag.RiskLimitPlatform }

type RiskLimitType struct{ quickfix.IntValue }

func (f RiskLimitType) Tag() tag.Tag { return tag.RiskLimitType }

type RiskMaturityMonthYear struct{ quickfix.MonthYearValue }

func (f RiskMaturityMonthYear) Tag() tag.Tag { return tag.RiskMaturityMonthYear }

type RiskMaturityTime struct{ quickfix.TZTimeOnlyValue }

func (f RiskMaturityTime) Tag() tag.Tag { return tag.RiskMaturityTime }

type RiskProduct struct{ quickfix.IntValue }

func (f RiskProduct) Tag() tag.Tag { return tag.RiskProduct }

type RiskProductComplex struct{ quickfix.StringValue }

func (f RiskProductComplex) Tag() tag.Tag { return tag.RiskProductComplex }

type RiskPutOrCall struct{ quickfix.IntValue }

func (f RiskPutOrCall) Tag() tag.Tag { return tag.RiskPutOrCall }

type RiskRestructuringType struct{ quickfix.StringValue }

func (f RiskRestructuringType) Tag() tag.Tag { return tag.RiskRestructuringType }

type RiskSecurityAltID struct{ quickfix.StringValue }

func (f RiskSecurityAltID) Tag() tag.Tag { return tag.RiskSecurityAltID }

type RiskSecurityAltIDSource struct{ quickfix.StringValue }

func (f RiskSecurityAltIDSource) Tag() tag.Tag { return tag.RiskSecurityAltIDSource }

type RiskSecurityDesc struct{ quickfix.StringValue }

func (f RiskSecurityDesc) Tag() tag.Tag { return tag.RiskSecurityDesc }

type RiskSecurityExchange struct{ quickfix.ExchangeValue }

func (f RiskSecurityExchange) Tag() tag.Tag { return tag.RiskSecurityExchange }

type RiskSecurityGroup struct{ quickfix.StringValue }

func (f RiskSecurityGroup) Tag() tag.Tag { return tag.RiskSecurityGroup }

type RiskSecurityID struct{ quickfix.StringValue }

func (f RiskSecurityID) Tag() tag.Tag { return tag.RiskSecurityID }

type RiskSecurityIDSource struct{ quickfix.StringValue }

func (f RiskSecurityIDSource) Tag() tag.Tag { return tag.RiskSecurityIDSource }

type RiskSecuritySubType struct{ quickfix.StringValue }

func (f RiskSecuritySubType) Tag() tag.Tag { return tag.RiskSecuritySubType }

type RiskSecurityType struct{ quickfix.StringValue }

func (f RiskSecurityType) Tag() tag.Tag { return tag.RiskSecurityType }

type RiskSeniority struct{ quickfix.StringValue }

func (f RiskSeniority) Tag() tag.Tag { return tag.RiskSeniority }

type RiskSymbol struct{ quickfix.StringValue }

func (f RiskSymbol) Tag() tag.Tag { return tag.RiskSymbol }

type RiskSymbolSfx struct{ quickfix.StringValue }

func (f RiskSymbolSfx) Tag() tag.Tag { return tag.RiskSymbolSfx }

type RiskWarningLevelName struct{ quickfix.StringValue }

func (f RiskWarningLevelName) Tag() tag.Tag { return tag.RiskWarningLevelName }

type RiskWarningLevelPercent struct{ quickfix.PercentageValue }

func (f RiskWarningLevelPercent) Tag() tag.Tag { return tag.RiskWarningLevelPercent }

type RndPx struct{ quickfix.PriceValue }

func (f RndPx) Tag() tag.Tag { return tag.RndPx }

type RootPartyID struct{ quickfix.StringValue }

func (f RootPartyID) Tag() tag.Tag { return tag.RootPartyID }

type RootPartyIDSource struct{ quickfix.CharValue }

func (f RootPartyIDSource) Tag() tag.Tag { return tag.RootPartyIDSource }

type RootPartyRole struct{ quickfix.IntValue }

func (f RootPartyRole) Tag() tag.Tag { return tag.RootPartyRole }

type RootPartySubID struct{ quickfix.StringValue }

func (f RootPartySubID) Tag() tag.Tag { return tag.RootPartySubID }

type RootPartySubIDType struct{ quickfix.IntValue }

func (f RootPartySubIDType) Tag() tag.Tag { return tag.RootPartySubIDType }

type RoundLot struct{ quickfix.QtyValue }

func (f RoundLot) Tag() tag.Tag { return tag.RoundLot }

type RoundingDirection struct{ quickfix.CharValue }

func (f RoundingDirection) Tag() tag.Tag { return tag.RoundingDirection }

type RoundingModulus struct{ quickfix.FloatValue }

func (f RoundingModulus) Tag() tag.Tag { return tag.RoundingModulus }

type RoutingID struct{ quickfix.StringValue }

func (f RoutingID) Tag() tag.Tag { return tag.RoutingID }

type RoutingType struct{ quickfix.IntValue }

func (f RoutingType) Tag() tag.Tag { return tag.RoutingType }

type RptSeq struct{ quickfix.IntValue }

func (f RptSeq) Tag() tag.Tag { return tag.RptSeq }

type RptSys struct{ quickfix.StringValue }

func (f RptSys) Tag() tag.Tag { return tag.RptSys }

type Rule80A struct{ quickfix.CharValue }

func (f Rule80A) Tag() tag.Tag { return tag.Rule80A }

type Scope struct{ quickfix.MultipleCharValue }

func (f Scope) Tag() tag.Tag { return tag.Scope }

type SecDefStatus struct{ quickfix.IntValue }

func (f SecDefStatus) Tag() tag.Tag { return tag.SecDefStatus }

type SecondaryAllocID struct{ quickfix.StringValue }

func (f SecondaryAllocID) Tag() tag.Tag { return tag.SecondaryAllocID }

type SecondaryClOrdID struct{ quickfix.StringValue }

func (f SecondaryClOrdID) Tag() tag.Tag { return tag.SecondaryClOrdID }

type SecondaryDisplayQty struct{ quickfix.QtyValue }

func (f SecondaryDisplayQty) Tag() tag.Tag { return tag.SecondaryDisplayQty }

type SecondaryExecID struct{ quickfix.StringValue }

func (f SecondaryExecID) Tag() tag.Tag { return tag.SecondaryExecID }

type SecondaryFirmTradeID struct{ quickfix.StringValue }

func (f SecondaryFirmTradeID) Tag() tag.Tag { return tag.SecondaryFirmTradeID }

type SecondaryHighLimitPrice struct{ quickfix.PriceValue }

func (f SecondaryHighLimitPrice) Tag() tag.Tag { return tag.SecondaryHighLimitPrice }

type SecondaryIndividualAllocID struct{ quickfix.StringValue }

func (f SecondaryIndividualAllocID) Tag() tag.Tag { return tag.SecondaryIndividualAllocID }

type SecondaryLowLimitPrice struct{ quickfix.PriceValue }

func (f SecondaryLowLimitPrice) Tag() tag.Tag { return tag.SecondaryLowLimitPrice }

type SecondaryOrderID struct{ quickfix.StringValue }

func (f SecondaryOrderID) Tag() tag.Tag { return tag.SecondaryOrderID }

type SecondaryPriceLimitType struct{ quickfix.IntValue }

func (f SecondaryPriceLimitType) Tag() tag.Tag { return tag.SecondaryPriceLimitType }

type SecondaryTradeID struct{ quickfix.StringValue }

func (f SecondaryTradeID) Tag() tag.Tag { return tag.SecondaryTradeID }

type SecondaryTradeReportID struct{ quickfix.StringValue }

func (f SecondaryTradeReportID) Tag() tag.Tag { return tag.SecondaryTradeReportID }

type SecondaryTradeReportRefID struct{ quickfix.StringValue }

func (f SecondaryTradeReportRefID) Tag() tag.Tag { return tag.SecondaryTradeReportRefID }

type SecondaryTradingReferencePrice struct{ quickfix.PriceValue }

func (f SecondaryTradingReferencePrice) Tag() tag.Tag { return tag.SecondaryTradingReferencePrice }

type SecondaryTrdType struct{ quickfix.IntValue }

func (f SecondaryTrdType) Tag() tag.Tag { return tag.SecondaryTrdType }

type SecureData struct{ quickfix.DataValue }

func (f SecureData) Tag() tag.Tag { return tag.SecureData }

type SecureDataLen struct{ quickfix.LengthValue }

func (f SecureDataLen) Tag() tag.Tag { return tag.SecureDataLen }

type SecurityAltID struct{ quickfix.StringValue }

func (f SecurityAltID) Tag() tag.Tag { return tag.SecurityAltID }

type SecurityAltIDSource struct{ quickfix.StringValue }

func (f SecurityAltIDSource) Tag() tag.Tag { return tag.SecurityAltIDSource }

type SecurityDesc struct{ quickfix.StringValue }

func (f SecurityDesc) Tag() tag.Tag { return tag.SecurityDesc }

type SecurityExchange struct{ quickfix.ExchangeValue }

func (f SecurityExchange) Tag() tag.Tag { return tag.SecurityExchange }

type SecurityGroup struct{ quickfix.StringValue }

func (f SecurityGroup) Tag() tag.Tag { return tag.SecurityGroup }

type SecurityID struct{ quickfix.StringValue }

func (f SecurityID) Tag() tag.Tag { return tag.SecurityID }

type SecurityIDSource struct{ quickfix.StringValue }

func (f SecurityIDSource) Tag() tag.Tag { return tag.SecurityIDSource }

type SecurityListDesc struct{ quickfix.StringValue }

func (f SecurityListDesc) Tag() tag.Tag { return tag.SecurityListDesc }

type SecurityListID struct{ quickfix.StringValue }

func (f SecurityListID) Tag() tag.Tag { return tag.SecurityListID }

type SecurityListRefID struct{ quickfix.StringValue }

func (f SecurityListRefID) Tag() tag.Tag { return tag.SecurityListRefID }

type SecurityListRequestType struct{ quickfix.IntValue }

func (f SecurityListRequestType) Tag() tag.Tag { return tag.SecurityListRequestType }

type SecurityListType struct{ quickfix.IntValue }

func (f SecurityListType) Tag() tag.Tag { return tag.SecurityListType }

type SecurityListTypeSource struct{ quickfix.IntValue }

func (f SecurityListTypeSource) Tag() tag.Tag { return tag.SecurityListTypeSource }

type SecurityReportID struct{ quickfix.IntValue }

func (f SecurityReportID) Tag() tag.Tag { return tag.SecurityReportID }

type SecurityReqID struct{ quickfix.StringValue }

func (f SecurityReqID) Tag() tag.Tag { return tag.SecurityReqID }

type SecurityRequestResult struct{ quickfix.IntValue }

func (f SecurityRequestResult) Tag() tag.Tag { return tag.SecurityRequestResult }

type SecurityRequestType struct{ quickfix.IntValue }

func (f SecurityRequestType) Tag() tag.Tag { return tag.SecurityRequestType }

type SecurityResponseID struct{ quickfix.StringValue }

func (f SecurityResponseID) Tag() tag.Tag { return tag.SecurityResponseID }

type SecurityResponseType struct{ quickfix.IntValue }

func (f SecurityResponseType) Tag() tag.Tag { return tag.SecurityResponseType }

type SecuritySettlAgentAcctName struct{ quickfix.StringValue }

func (f SecuritySettlAgentAcctName) Tag() tag.Tag { return tag.SecuritySettlAgentAcctName }

type SecuritySettlAgentAcctNum struct{ quickfix.StringValue }

func (f SecuritySettlAgentAcctNum) Tag() tag.Tag { return tag.SecuritySettlAgentAcctNum }

type SecuritySettlAgentCode struct{ quickfix.StringValue }

func (f SecuritySettlAgentCode) Tag() tag.Tag { return tag.SecuritySettlAgentCode }

type SecuritySettlAgentContactName struct{ quickfix.StringValue }

func (f SecuritySettlAgentContactName) Tag() tag.Tag { return tag.SecuritySettlAgentContactName }

type SecuritySettlAgentContactPhone struct{ quickfix.StringValue }

func (f SecuritySettlAgentContactPhone) Tag() tag.Tag { return tag.SecuritySettlAgentContactPhone }

type SecuritySettlAgentName struct{ quickfix.StringValue }

func (f SecuritySettlAgentName) Tag() tag.Tag { return tag.SecuritySettlAgentName }

type SecurityStatus struct{ quickfix.StringValue }

func (f SecurityStatus) Tag() tag.Tag { return tag.SecurityStatus }

type SecurityStatusReqID struct{ quickfix.StringValue }

func (f SecurityStatusReqID) Tag() tag.Tag { return tag.SecurityStatusReqID }

type SecuritySubType struct{ quickfix.StringValue }

func (f SecuritySubType) Tag() tag.Tag { return tag.SecuritySubType }

type SecurityTradingEvent struct{ quickfix.IntValue }

func (f SecurityTradingEvent) Tag() tag.Tag { return tag.SecurityTradingEvent }

type SecurityTradingStatus struct{ quickfix.IntValue }

func (f SecurityTradingStatus) Tag() tag.Tag { return tag.SecurityTradingStatus }

type SecurityType struct{ quickfix.StringValue }

func (f SecurityType) Tag() tag.Tag { return tag.SecurityType }

type SecurityUpdateAction struct{ quickfix.CharValue }

func (f SecurityUpdateAction) Tag() tag.Tag { return tag.SecurityUpdateAction }

type SecurityXML struct{ quickfix.XMLDataValue }

func (f SecurityXML) Tag() tag.Tag { return tag.SecurityXML }

type SecurityXMLLen struct{ quickfix.LengthValue }

func (f SecurityXMLLen) Tag() tag.Tag { return tag.SecurityXMLLen }

type SecurityXMLSchema struct{ quickfix.StringValue }

func (f SecurityXMLSchema) Tag() tag.Tag { return tag.SecurityXMLSchema }

type SellVolume struct{ quickfix.QtyValue }

func (f SellVolume) Tag() tag.Tag { return tag.SellVolume }

type SellerDays struct{ quickfix.IntValue }

func (f SellerDays) Tag() tag.Tag { return tag.SellerDays }

type SenderCompID struct{ quickfix.StringValue }

func (f SenderCompID) Tag() tag.Tag { return tag.SenderCompID }

type SenderLocationID struct{ quickfix.StringValue }

func (f SenderLocationID) Tag() tag.Tag { return tag.SenderLocationID }

type SenderSubID struct{ quickfix.StringValue }

func (f SenderSubID) Tag() tag.Tag { return tag.SenderSubID }

type SendingDate struct{ quickfix.LocalMktDateValue }

func (f SendingDate) Tag() tag.Tag { return tag.SendingDate }

type SendingTime struct{ quickfix.UTCTimestampValue }

func (f SendingTime) Tag() tag.Tag { return tag.SendingTime }

type Seniority struct{ quickfix.StringValue }

func (f Seniority) Tag() tag.Tag { return tag.Seniority }

type SessionRejectReason struct{ quickfix.IntValue }

func (f SessionRejectReason) Tag() tag.Tag { return tag.SessionRejectReason }

type SessionStatus struct{ quickfix.IntValue }

func (f SessionStatus) Tag() tag.Tag { return tag.SessionStatus }

type SettlBrkrCode struct{ quickfix.StringValue }

func (f SettlBrkrCode) Tag() tag.Tag { return tag.SettlBrkrCode }

type SettlCurrAmt struct{ quickfix.AmtValue }

func (f SettlCurrAmt) Tag() tag.Tag { return tag.SettlCurrAmt }

type SettlCurrBidFxRate struct{ quickfix.FloatValue }

func (f SettlCurrBidFxRate) Tag() tag.Tag { return tag.SettlCurrBidFxRate }

type SettlCurrFxRate struct{ quickfix.FloatValue }

func (f SettlCurrFxRate) Tag() tag.Tag { return tag.SettlCurrFxRate }

type SettlCurrFxRateCalc struct{ quickfix.CharValue }

func (f SettlCurrFxRateCalc) Tag() tag.Tag { return tag.SettlCurrFxRateCalc }

type SettlCurrOfferFxRate struct{ quickfix.FloatValue }

func (f SettlCurrOfferFxRate) Tag() tag.Tag { return tag.SettlCurrOfferFxRate }

type SettlCurrency struct{ quickfix.CurrencyValue }

func (f SettlCurrency) Tag() tag.Tag { return tag.SettlCurrency }

type SettlDate struct{ quickfix.LocalMktDateValue }

func (f SettlDate) Tag() tag.Tag { return tag.SettlDate }

type SettlDate2 struct{ quickfix.LocalMktDateValue }

func (f SettlDate2) Tag() tag.Tag { return tag.SettlDate2 }

type SettlDeliveryType struct{ quickfix.IntValue }

func (f SettlDeliveryType) Tag() tag.Tag { return tag.SettlDeliveryType }

type SettlDepositoryCode struct{ quickfix.StringValue }

func (f SettlDepositoryCode) Tag() tag.Tag { return tag.SettlDepositoryCode }

type SettlInstCode struct{ quickfix.StringValue }

func (f SettlInstCode) Tag() tag.Tag { return tag.SettlInstCode }

type SettlInstID struct{ quickfix.StringValue }

func (f SettlInstID) Tag() tag.Tag { return tag.SettlInstID }

type SettlInstMode struct{ quickfix.CharValue }

func (f SettlInstMode) Tag() tag.Tag { return tag.SettlInstMode }

type SettlInstMsgID struct{ quickfix.StringValue }

func (f SettlInstMsgID) Tag() tag.Tag { return tag.SettlInstMsgID }

type SettlInstRefID struct{ quickfix.StringValue }

func (f SettlInstRefID) Tag() tag.Tag { return tag.SettlInstRefID }

type SettlInstReqID struct{ quickfix.StringValue }

func (f SettlInstReqID) Tag() tag.Tag { return tag.SettlInstReqID }

type SettlInstReqRejCode struct{ quickfix.IntValue }

func (f SettlInstReqRejCode) Tag() tag.Tag { return tag.SettlInstReqRejCode }

type SettlInstSource struct{ quickfix.CharValue }

func (f SettlInstSource) Tag() tag.Tag { return tag.SettlInstSource }

type SettlInstTransType struct{ quickfix.CharValue }

func (f SettlInstTransType) Tag() tag.Tag { return tag.SettlInstTransType }

type SettlLocation struct{ quickfix.StringValue }

func (f SettlLocation) Tag() tag.Tag { return tag.SettlLocation }

type SettlMethod struct{ quickfix.CharValue }

func (f SettlMethod) Tag() tag.Tag { return tag.SettlMethod }

type SettlObligID struct{ quickfix.StringValue }

func (f SettlObligID) Tag() tag.Tag { return tag.SettlObligID }

type SettlObligMode struct{ quickfix.IntValue }

func (f SettlObligMode) Tag() tag.Tag { return tag.SettlObligMode }

type SettlObligMsgID struct{ quickfix.StringValue }

func (f SettlObligMsgID) Tag() tag.Tag { return tag.SettlObligMsgID }

type SettlObligRefID struct{ quickfix.StringValue }

func (f SettlObligRefID) Tag() tag.Tag { return tag.SettlObligRefID }

type SettlObligSource struct{ quickfix.CharValue }

func (f SettlObligSource) Tag() tag.Tag { return tag.SettlObligSource }

type SettlObligTransType struct{ quickfix.CharValue }

func (f SettlObligTransType) Tag() tag.Tag { return tag.SettlObligTransType }

type SettlPartyID struct{ quickfix.StringValue }

func (f SettlPartyID) Tag() tag.Tag { return tag.SettlPartyID }

type SettlPartyIDSource struct{ quickfix.CharValue }

func (f SettlPartyIDSource) Tag() tag.Tag { return tag.SettlPartyIDSource }

type SettlPartyRole struct{ quickfix.IntValue }

func (f SettlPartyRole) Tag() tag.Tag { return tag.SettlPartyRole }

type SettlPartySubID struct{ quickfix.StringValue }

func (f SettlPartySubID) Tag() tag.Tag { return tag.SettlPartySubID }

type SettlPartySubIDType struct{ quickfix.IntValue }

func (f SettlPartySubIDType) Tag() tag.Tag { return tag.SettlPartySubIDType }

type SettlPrice struct{ quickfix.PriceValue }

func (f SettlPrice) Tag() tag.Tag { return tag.SettlPrice }

type SettlPriceType struct{ quickfix.IntValue }

func (f SettlPriceType) Tag() tag.Tag { return tag.SettlPriceType }

type SettlSessID struct{ quickfix.StringValue }

func (f SettlSessID) Tag() tag.Tag { return tag.SettlSessID }

type SettlSessSubID struct{ quickfix.StringValue }

func (f SettlSessSubID) Tag() tag.Tag { return tag.SettlSessSubID }

type SettlType struct{ quickfix.StringValue }

func (f SettlType) Tag() tag.Tag { return tag.SettlType }

type SettleOnOpenFlag struct{ quickfix.StringValue }

func (f SettleOnOpenFlag) Tag() tag.Tag { return tag.SettleOnOpenFlag }

type SettlementCycleNo struct{ quickfix.IntValue }

func (f SettlementCycleNo) Tag() tag.Tag { return tag.SettlementCycleNo }

type SettlmntTyp struct{ quickfix.CharValue }

func (f SettlmntTyp) Tag() tag.Tag { return tag.SettlmntTyp }

type SharedCommission struct{ quickfix.AmtValue }

func (f SharedCommission) Tag() tag.Tag { return tag.SharedCommission }

type Shares struct{ quickfix.QtyValue }

func (f Shares) Tag() tag.Tag { return tag.Shares }

type ShortQty struct{ quickfix.QtyValue }

func (f ShortQty) Tag() tag.Tag { return tag.ShortQty }

type ShortSaleReason struct{ quickfix.IntValue }

func (f ShortSaleReason) Tag() tag.Tag { return tag.ShortSaleReason }

type Side struct{ quickfix.CharValue }

func (f Side) Tag() tag.Tag { return tag.Side }

type SideComplianceID struct{ quickfix.StringValue }

func (f SideComplianceID) Tag() tag.Tag { return tag.SideComplianceID }

type SideCurrency struct{ quickfix.CurrencyValue }

func (f SideCurrency) Tag() tag.Tag { return tag.SideCurrency }

type SideExecID struct{ quickfix.StringValue }

func (f SideExecID) Tag() tag.Tag { return tag.SideExecID }

type SideFillStationCd struct{ quickfix.StringValue }

func (f SideFillStationCd) Tag() tag.Tag { return tag.SideFillStationCd }

type SideGrossTradeAmt struct{ quickfix.AmtValue }

func (f SideGrossTradeAmt) Tag() tag.Tag { return tag.SideGrossTradeAmt }

type SideLastQty struct{ quickfix.IntValue }

func (f SideLastQty) Tag() tag.Tag { return tag.SideLastQty }

type SideLiquidityInd struct{ quickfix.IntValue }

func (f SideLiquidityInd) Tag() tag.Tag { return tag.SideLiquidityInd }

type SideMultiLegReportingType struct{ quickfix.IntValue }

func (f SideMultiLegReportingType) Tag() tag.Tag { return tag.SideMultiLegReportingType }

type SideQty struct{ quickfix.IntValue }

func (f SideQty) Tag() tag.Tag { return tag.SideQty }

type SideReasonCd struct{ quickfix.StringValue }

func (f SideReasonCd) Tag() tag.Tag { return tag.SideReasonCd }

type SideSettlCurrency struct{ quickfix.CurrencyValue }

func (f SideSettlCurrency) Tag() tag.Tag { return tag.SideSettlCurrency }

type SideTimeInForce struct{ quickfix.UTCTimestampValue }

func (f SideTimeInForce) Tag() tag.Tag { return tag.SideTimeInForce }

type SideTradeReportID struct{ quickfix.StringValue }

func (f SideTradeReportID) Tag() tag.Tag { return tag.SideTradeReportID }

type SideTrdRegTimestamp struct{ quickfix.UTCTimestampValue }

func (f SideTrdRegTimestamp) Tag() tag.Tag { return tag.SideTrdRegTimestamp }

type SideTrdRegTimestampSrc struct{ quickfix.StringValue }

func (f SideTrdRegTimestampSrc) Tag() tag.Tag { return tag.SideTrdRegTimestampSrc }

type SideTrdRegTimestampType struct{ quickfix.IntValue }

func (f SideTrdRegTimestampType) Tag() tag.Tag { return tag.SideTrdRegTimestampType }

type SideTrdSubTyp struct{ quickfix.IntValue }

func (f SideTrdSubTyp) Tag() tag.Tag { return tag.SideTrdSubTyp }

type SideValue1 struct{ quickfix.AmtValue }

func (f SideValue1) Tag() tag.Tag { return tag.SideValue1 }

type SideValue2 struct{ quickfix.AmtValue }

func (f SideValue2) Tag() tag.Tag { return tag.SideValue2 }

type SideValueInd struct{ quickfix.IntValue }

func (f SideValueInd) Tag() tag.Tag { return tag.SideValueInd }

type Signature struct{ quickfix.DataValue }

func (f Signature) Tag() tag.Tag { return tag.Signature }

type SignatureLength struct{ quickfix.LengthValue }

func (f SignatureLength) Tag() tag.Tag { return tag.SignatureLength }

type SolicitedFlag struct{ quickfix.BooleanValue }

func (f SolicitedFlag) Tag() tag.Tag { return tag.SolicitedFlag }

type Spread struct{ quickfix.PriceOffsetValue }

func (f Spread) Tag() tag.Tag { return tag.Spread }

type SpreadToBenchmark struct{ quickfix.PriceOffsetValue }

func (f SpreadToBenchmark) Tag() tag.Tag { return tag.SpreadToBenchmark }

type StandInstDbID struct{ quickfix.StringValue }

func (f StandInstDbID) Tag() tag.Tag { return tag.StandInstDbID }

type StandInstDbName struct{ quickfix.StringValue }

func (f StandInstDbName) Tag() tag.Tag { return tag.StandInstDbName }

type StandInstDbType struct{ quickfix.IntValue }

func (f StandInstDbType) Tag() tag.Tag { return tag.StandInstDbType }

type StartCash struct{ quickfix.AmtValue }

func (f StartCash) Tag() tag.Tag { return tag.StartCash }

type StartDate struct{ quickfix.LocalMktDateValue }

func (f StartDate) Tag() tag.Tag { return tag.StartDate }

type StartMaturityMonthYear struct{ quickfix.MonthYearValue }

func (f StartMaturityMonthYear) Tag() tag.Tag { return tag.StartMaturityMonthYear }

type StartStrikePxRange struct{ quickfix.PriceValue }

func (f StartStrikePxRange) Tag() tag.Tag { return tag.StartStrikePxRange }

type StartTickPriceRange struct{ quickfix.PriceValue }

func (f StartTickPriceRange) Tag() tag.Tag { return tag.StartTickPriceRange }

type StateOrProvinceOfIssue struct{ quickfix.StringValue }

func (f StateOrProvinceOfIssue) Tag() tag.Tag { return tag.StateOrProvinceOfIssue }

type StatsType struct{ quickfix.IntValue }

func (f StatsType) Tag() tag.Tag { return tag.StatsType }

type StatusText struct{ quickfix.StringValue }

func (f StatusText) Tag() tag.Tag { return tag.StatusText }

type StatusValue struct{ quickfix.IntValue }

func (f StatusValue) Tag() tag.Tag { return tag.StatusValue }

type StipulationType struct{ quickfix.StringValue }

func (f StipulationType) Tag() tag.Tag { return tag.StipulationType }

type StipulationValue struct{ quickfix.StringValue }

func (f StipulationValue) Tag() tag.Tag { return tag.StipulationValue }

type StopPx struct{ quickfix.PriceValue }

func (f StopPx) Tag() tag.Tag { return tag.StopPx }

type StrategyParameterName struct{ quickfix.StringValue }

func (f StrategyParameterName) Tag() tag.Tag { return tag.StrategyParameterName }

type StrategyParameterType struct{ quickfix.IntValue }

func (f StrategyParameterType) Tag() tag.Tag { return tag.StrategyParameterType }

type StrategyParameterValue struct{ quickfix.StringValue }

func (f StrategyParameterValue) Tag() tag.Tag { return tag.StrategyParameterValue }

type StreamAsgnAckType struct{ quickfix.IntValue }

func (f StreamAsgnAckType) Tag() tag.Tag { return tag.StreamAsgnAckType }

type StreamAsgnRejReason struct{ quickfix.IntValue }

func (f StreamAsgnRejReason) Tag() tag.Tag { return tag.StreamAsgnRejReason }

type StreamAsgnReqID struct{ quickfix.StringValue }

func (f StreamAsgnReqID) Tag() tag.Tag { return tag.StreamAsgnReqID }

type StreamAsgnReqType struct{ quickfix.IntValue }

func (f StreamAsgnReqType) Tag() tag.Tag { return tag.StreamAsgnReqType }

type StreamAsgnRptID struct{ quickfix.StringValue }

func (f StreamAsgnRptID) Tag() tag.Tag { return tag.StreamAsgnRptID }

type StreamAsgnType struct{ quickfix.IntValue }

func (f StreamAsgnType) Tag() tag.Tag { return tag.StreamAsgnType }

type StrikeCurrency struct{ quickfix.CurrencyValue }

func (f StrikeCurrency) Tag() tag.Tag { return tag.StrikeCurrency }

type StrikeExerciseStyle struct{ quickfix.IntValue }

func (f StrikeExerciseStyle) Tag() tag.Tag { return tag.StrikeExerciseStyle }

type StrikeIncrement struct{ quickfix.FloatValue }

func (f StrikeIncrement) Tag() tag.Tag { return tag.StrikeIncrement }

type StrikeMultiplier struct{ quickfix.FloatValue }

func (f StrikeMultiplier) Tag() tag.Tag { return tag.StrikeMultiplier }

type StrikePrice struct{ quickfix.PriceValue }

func (f StrikePrice) Tag() tag.Tag { return tag.StrikePrice }

type StrikePriceBoundaryMethod struct{ quickfix.IntValue }

func (f StrikePriceBoundaryMethod) Tag() tag.Tag { return tag.StrikePriceBoundaryMethod }

type StrikePriceBoundaryPrecision struct{ quickfix.PercentageValue }

func (f StrikePriceBoundaryPrecision) Tag() tag.Tag { return tag.StrikePriceBoundaryPrecision }

type StrikePriceDeterminationMethod struct{ quickfix.IntValue }

func (f StrikePriceDeterminationMethod) Tag() tag.Tag { return tag.StrikePriceDeterminationMethod }

type StrikeRuleID struct{ quickfix.StringValue }

func (f StrikeRuleID) Tag() tag.Tag { return tag.StrikeRuleID }

type StrikeTime struct{ quickfix.UTCTimestampValue }

func (f StrikeTime) Tag() tag.Tag { return tag.StrikeTime }

type StrikeValue struct{ quickfix.FloatValue }

func (f StrikeValue) Tag() tag.Tag { return tag.StrikeValue }

type Subject struct{ quickfix.StringValue }

func (f Subject) Tag() tag.Tag { return tag.Subject }

type SubscriptionRequestType struct{ quickfix.CharValue }

func (f SubscriptionRequestType) Tag() tag.Tag { return tag.SubscriptionRequestType }

type SwapPoints struct{ quickfix.PriceOffsetValue }

func (f SwapPoints) Tag() tag.Tag { return tag.SwapPoints }

type Symbol struct{ quickfix.StringValue }

func (f Symbol) Tag() tag.Tag { return tag.Symbol }

type SymbolSfx struct{ quickfix.StringValue }

func (f SymbolSfx) Tag() tag.Tag { return tag.SymbolSfx }

type TZTransactTime struct{ quickfix.TZTimestampValue }

func (f TZTransactTime) Tag() tag.Tag { return tag.TZTransactTime }

type TargetCompID struct{ quickfix.StringValue }

func (f TargetCompID) Tag() tag.Tag { return tag.TargetCompID }

type TargetLocationID struct{ quickfix.StringValue }

func (f TargetLocationID) Tag() tag.Tag { return tag.TargetLocationID }

type TargetPartyID struct{ quickfix.StringValue }

func (f TargetPartyID) Tag() tag.Tag { return tag.TargetPartyID }

type TargetPartyIDSource struct{ quickfix.CharValue }

func (f TargetPartyIDSource) Tag() tag.Tag { return tag.TargetPartyIDSource }

type TargetPartyRole struct{ quickfix.IntValue }

func (f TargetPartyRole) Tag() tag.Tag { return tag.TargetPartyRole }

type TargetStrategy struct{ quickfix.IntValue }

func (f TargetStrategy) Tag() tag.Tag { return tag.TargetStrategy }

type TargetStrategyParameters struct{ quickfix.StringValue }

func (f TargetStrategyParameters) Tag() tag.Tag { return tag.TargetStrategyParameters }

type TargetStrategyPerformance struct{ quickfix.FloatValue }

func (f TargetStrategyPerformance) Tag() tag.Tag { return tag.TargetStrategyPerformance }

type TargetSubID struct{ quickfix.StringValue }

func (f TargetSubID) Tag() tag.Tag { return tag.TargetSubID }

type TaxAdvantageType struct{ quickfix.IntValue }

func (f TaxAdvantageType) Tag() tag.Tag { return tag.TaxAdvantageType }

type TerminationType struct{ quickfix.IntValue }

func (f TerminationType) Tag() tag.Tag { return tag.TerminationType }

type TestMessageIndicator struct{ quickfix.BooleanValue }

func (f TestMessageIndicator) Tag() tag.Tag { return tag.TestMessageIndicator }

type TestReqID struct{ quickfix.StringValue }

func (f TestReqID) Tag() tag.Tag { return tag.TestReqID }

type Text struct{ quickfix.StringValue }

func (f Text) Tag() tag.Tag { return tag.Text }

type ThresholdAmount struct{ quickfix.PriceOffsetValue }

func (f ThresholdAmount) Tag() tag.Tag { return tag.ThresholdAmount }

type TickDirection struct{ quickfix.CharValue }

func (f TickDirection) Tag() tag.Tag { return tag.TickDirection }

type TickIncrement struct{ quickfix.PriceValue }

func (f TickIncrement) Tag() tag.Tag { return tag.TickIncrement }

type TickRuleType struct{ quickfix.IntValue }

func (f TickRuleType) Tag() tag.Tag { return tag.TickRuleType }

type TierCode struct{ quickfix.StringValue }

func (f TierCode) Tag() tag.Tag { return tag.TierCode }

type TimeBracket struct{ quickfix.StringValue }

func (f TimeBracket) Tag() tag.Tag { return tag.TimeBracket }

type TimeInForce struct{ quickfix.CharValue }

func (f TimeInForce) Tag() tag.Tag { return tag.TimeInForce }

type TimeToExpiration struct{ quickfix.FloatValue }

func (f TimeToExpiration) Tag() tag.Tag { return tag.TimeToExpiration }

type TimeUnit struct{ quickfix.StringValue }

func (f TimeUnit) Tag() tag.Tag { return tag.TimeUnit }

type TotNoAccQuotes struct{ quickfix.IntValue }

func (f TotNoAccQuotes) Tag() tag.Tag { return tag.TotNoAccQuotes }

type TotNoAllocs struct{ quickfix.IntValue }

func (f TotNoAllocs) Tag() tag.Tag { return tag.TotNoAllocs }

type TotNoCxldQuotes struct{ quickfix.IntValue }

func (f TotNoCxldQuotes) Tag() tag.Tag { return tag.TotNoCxldQuotes }

type TotNoFills struct{ quickfix.IntValue }

func (f TotNoFills) Tag() tag.Tag { return tag.TotNoFills }

type TotNoOrders struct{ quickfix.IntValue }

func (f TotNoOrders) Tag() tag.Tag { return tag.TotNoOrders }

type TotNoPartyList struct{ quickfix.IntValue }

func (f TotNoPartyList) Tag() tag.Tag { return tag.TotNoPartyList }

type TotNoQuoteEntries struct{ quickfix.IntValue }

func (f TotNoQuoteEntries) Tag() tag.Tag { return tag.TotNoQuoteEntries }

type TotNoRejQuotes struct{ quickfix.IntValue }

func (f TotNoRejQuotes) Tag() tag.Tag { return tag.TotNoRejQuotes }

type TotNoRelatedSym struct{ quickfix.IntValue }

func (f TotNoRelatedSym) Tag() tag.Tag { return tag.TotNoRelatedSym }

type TotNoSecurityTypes struct{ quickfix.IntValue }

func (f TotNoSecurityTypes) Tag() tag.Tag { return tag.TotNoSecurityTypes }

type TotNoStrikes struct{ quickfix.IntValue }

func (f TotNoStrikes) Tag() tag.Tag { return tag.TotNoStrikes }

type TotNumAssignmentReports struct{ quickfix.IntValue }

func (f TotNumAssignmentReports) Tag() tag.Tag { return tag.TotNumAssignmentReports }

type TotNumReports struct{ quickfix.IntValue }

func (f TotNumReports) Tag() tag.Tag { return tag.TotNumReports }

type TotNumTradeReports struct{ quickfix.IntValue }

func (f TotNumTradeReports) Tag() tag.Tag { return tag.TotNumTradeReports }

type TotQuoteEntries struct{ quickfix.IntValue }

func (f TotQuoteEntries) Tag() tag.Tag { return tag.TotQuoteEntries }

type TotalAccruedInterestAmt struct{ quickfix.AmtValue }

func (f TotalAccruedInterestAmt) Tag() tag.Tag { return tag.TotalAccruedInterestAmt }

type TotalAffectedOrders struct{ quickfix.IntValue }

func (f TotalAffectedOrders) Tag() tag.Tag { return tag.TotalAffectedOrders }

type TotalNetValue struct{ quickfix.AmtValue }

func (f TotalNetValue) Tag() tag.Tag { return tag.TotalNetValue }

type TotalNumPosReports struct{ quickfix.IntValue }

func (f TotalNumPosReports) Tag() tag.Tag { return tag.TotalNumPosReports }

type TotalNumSecurities struct{ quickfix.IntValue }

func (f TotalNumSecurities) Tag() tag.Tag { return tag.TotalNumSecurities }

type TotalNumSecurityTypes struct{ quickfix.IntValue }

func (f TotalNumSecurityTypes) Tag() tag.Tag { return tag.TotalNumSecurityTypes }

type TotalTakedown struct{ quickfix.AmtValue }

func (f TotalTakedown) Tag() tag.Tag { return tag.TotalTakedown }

type TotalVolumeTraded struct{ quickfix.QtyValue }

func (f TotalVolumeTraded) Tag() tag.Tag { return tag.TotalVolumeTraded }

type TotalVolumeTradedDate struct{ quickfix.UTCDateOnlyValue }

func (f TotalVolumeTradedDate) Tag() tag.Tag { return tag.TotalVolumeTradedDate }

type TotalVolumeTradedTime struct{ quickfix.UTCTimeOnlyValue }

func (f TotalVolumeTradedTime) Tag() tag.Tag { return tag.TotalVolumeTradedTime }

type TradSesCloseTime struct{ quickfix.UTCTimestampValue }

func (f TradSesCloseTime) Tag() tag.Tag { return tag.TradSesCloseTime }

type TradSesEndTime struct{ quickfix.UTCTimestampValue }

func (f TradSesEndTime) Tag() tag.Tag { return tag.TradSesEndTime }

type TradSesEvent struct{ quickfix.IntValue }

func (f TradSesEvent) Tag() tag.Tag { return tag.TradSesEvent }

type TradSesMethod struct{ quickfix.IntValue }

func (f TradSesMethod) Tag() tag.Tag { return tag.TradSesMethod }

type TradSesMode struct{ quickfix.IntValue }

func (f TradSesMode) Tag() tag.Tag { return tag.TradSesMode }

type TradSesOpenTime struct{ quickfix.UTCTimestampValue }

func (f TradSesOpenTime) Tag() tag.Tag { return tag.TradSesOpenTime }

type TradSesPreCloseTime struct{ quickfix.UTCTimestampValue }

func (f TradSesPreCloseTime) Tag() tag.Tag { return tag.TradSesPreCloseTime }

type TradSesReqID struct{ quickfix.StringValue }

func (f TradSesReqID) Tag() tag.Tag { return tag.TradSesReqID }

type TradSesStartTime struct{ quickfix.UTCTimestampValue }

func (f TradSesStartTime) Tag() tag.Tag { return tag.TradSesStartTime }

type TradSesStatus struct{ quickfix.IntValue }

func (f TradSesStatus) Tag() tag.Tag { return tag.TradSesStatus }

type TradSesStatusRejReason struct{ quickfix.IntValue }

func (f TradSesStatusRejReason) Tag() tag.Tag { return tag.TradSesStatusRejReason }

type TradSesUpdateAction struct{ quickfix.CharValue }

func (f TradSesUpdateAction) Tag() tag.Tag { return tag.TradSesUpdateAction }

type TradeAllocIndicator struct{ quickfix.IntValue }

func (f TradeAllocIndicator) Tag() tag.Tag { return tag.TradeAllocIndicator }

type TradeCondition struct{ quickfix.MultipleStringValue }

func (f TradeCondition) Tag() tag.Tag { return tag.TradeCondition }

type TradeDate struct{ quickfix.LocalMktDateValue }

func (f TradeDate) Tag() tag.Tag { return tag.TradeDate }

type TradeHandlingInstr struct{ quickfix.CharValue }

func (f TradeHandlingInstr) Tag() tag.Tag { return tag.TradeHandlingInstr }

type TradeID struct{ quickfix.StringValue }

func (f TradeID) Tag() tag.Tag { return tag.TradeID }

type TradeInputDevice struct{ quickfix.StringValue }

func (f TradeInputDevice) Tag() tag.Tag { return tag.TradeInputDevice }

type TradeInputSource struct{ quickfix.StringValue }

func (f TradeInputSource) Tag() tag.Tag { return tag.TradeInputSource }

type TradeLegRefID struct{ quickfix.StringValue }

func (f TradeLegRefID) Tag() tag.Tag { return tag.TradeLegRefID }

type TradeLinkID struct{ quickfix.StringValue }

func (f TradeLinkID) Tag() tag.Tag { return tag.TradeLinkID }

type TradeOriginationDate struct{ quickfix.LocalMktDateValue }

func (f TradeOriginationDate) Tag() tag.Tag { return tag.TradeOriginationDate }

type TradePublishIndicator struct{ quickfix.IntValue }

func (f TradePublishIndicator) Tag() tag.Tag { return tag.TradePublishIndicator }

type TradeReportID struct{ quickfix.StringValue }

func (f TradeReportID) Tag() tag.Tag { return tag.TradeReportID }

type TradeReportRefID struct{ quickfix.StringValue }

func (f TradeReportRefID) Tag() tag.Tag { return tag.TradeReportRefID }

type TradeReportRejectReason struct{ quickfix.IntValue }

func (f TradeReportRejectReason) Tag() tag.Tag { return tag.TradeReportRejectReason }

type TradeReportTransType struct{ quickfix.IntValue }

func (f TradeReportTransType) Tag() tag.Tag { return tag.TradeReportTransType }

type TradeReportType struct{ quickfix.IntValue }

func (f TradeReportType) Tag() tag.Tag { return tag.TradeReportType }

type TradeRequestID struct{ quickfix.StringValue }

func (f TradeRequestID) Tag() tag.Tag { return tag.TradeRequestID }

type TradeRequestResult struct{ quickfix.IntValue }

func (f TradeRequestResult) Tag() tag.Tag { return tag.TradeRequestResult }

type TradeRequestStatus struct{ quickfix.IntValue }

func (f TradeRequestStatus) Tag() tag.Tag { return tag.TradeRequestStatus }

type TradeRequestType struct{ quickfix.IntValue }

func (f TradeRequestType) Tag() tag.Tag { return tag.TradeRequestType }

type TradeType struct{ quickfix.CharValue }

func (f TradeType) Tag() tag.Tag { return tag.TradeType }

type TradeVolume struct{ quickfix.QtyValue }

func (f TradeVolume) Tag() tag.Tag { return tag.TradeVolume }

type TradedFlatSwitch struct{ quickfix.BooleanValue }

func (f TradedFlatSwitch) Tag() tag.Tag { return tag.TradedFlatSwitch }

type TradingCurrency struct{ quickfix.CurrencyValue }

func (f TradingCurrency) Tag() tag.Tag { return tag.TradingCurrency }

type TradingReferencePrice struct{ quickfix.PriceValue }

func (f TradingReferencePrice) Tag() tag.Tag { return tag.TradingReferencePrice }

type TradingSessionDesc struct{ quickfix.StringValue }

func (f TradingSessionDesc) Tag() tag.Tag { return tag.TradingSessionDesc }

type TradingSessionID struct{ quickfix.StringValue }

func (f TradingSessionID) Tag() tag.Tag { return tag.TradingSessionID }

type TradingSessionSubID struct{ quickfix.StringValue }

func (f TradingSessionSubID) Tag() tag.Tag { return tag.TradingSessionSubID }

type TransBkdTime struct{ quickfix.UTCTimestampValue }

func (f TransBkdTime) Tag() tag.Tag { return tag.TransBkdTime }

type TransactTime struct{ quickfix.UTCTimestampValue }

func (f TransactTime) Tag() tag.Tag { return tag.TransactTime }

type TransferReason struct{ quickfix.StringValue }

func (f TransferReason) Tag() tag.Tag { return tag.TransferReason }

type TrdMatchID struct{ quickfix.StringValue }

func (f TrdMatchID) Tag() tag.Tag { return tag.TrdMatchID }

type TrdRegTimestamp struct{ quickfix.UTCTimestampValue }

func (f TrdRegTimestamp) Tag() tag.Tag { return tag.TrdRegTimestamp }

type TrdRegTimestampOrigin struct{ quickfix.StringValue }

func (f TrdRegTimestampOrigin) Tag() tag.Tag { return tag.TrdRegTimestampOrigin }

type TrdRegTimestampType struct{ quickfix.IntValue }

func (f TrdRegTimestampType) Tag() tag.Tag { return tag.TrdRegTimestampType }

type TrdRepIndicator struct{ quickfix.BooleanValue }

func (f TrdRepIndicator) Tag() tag.Tag { return tag.TrdRepIndicator }

type TrdRepPartyRole struct{ quickfix.IntValue }

func (f TrdRepPartyRole) Tag() tag.Tag { return tag.TrdRepPartyRole }

type TrdRptStatus struct{ quickfix.IntValue }

func (f TrdRptStatus) Tag() tag.Tag { return tag.TrdRptStatus }

type TrdSubType struct{ quickfix.IntValue }

func (f TrdSubType) Tag() tag.Tag { return tag.TrdSubType }

type TrdType struct{ quickfix.IntValue }

func (f TrdType) Tag() tag.Tag { return tag.TrdType }

type TriggerAction struct{ quickfix.CharValue }

func (f TriggerAction) Tag() tag.Tag { return tag.TriggerAction }

type TriggerNewPrice struct{ quickfix.PriceValue }

func (f TriggerNewPrice) Tag() tag.Tag { return tag.TriggerNewPrice }

type TriggerNewQty struct{ quickfix.QtyValue }

func (f TriggerNewQty) Tag() tag.Tag { return tag.TriggerNewQty }

type TriggerOrderType struct{ quickfix.CharValue }

func (f TriggerOrderType) Tag() tag.Tag { return tag.TriggerOrderType }

type TriggerPrice struct{ quickfix.PriceValue }

func (f TriggerPrice) Tag() tag.Tag { return tag.TriggerPrice }

type TriggerPriceDirection struct{ quickfix.CharValue }

func (f TriggerPriceDirection) Tag() tag.Tag { return tag.TriggerPriceDirection }

type TriggerPriceType struct{ quickfix.CharValue }

func (f TriggerPriceType) Tag() tag.Tag { return tag.TriggerPriceType }

type TriggerPriceTypeScope struct{ quickfix.CharValue }

func (f TriggerPriceTypeScope) Tag() tag.Tag { return tag.TriggerPriceTypeScope }

type TriggerSecurityDesc struct{ quickfix.StringValue }

func (f TriggerSecurityDesc) Tag() tag.Tag { return tag.TriggerSecurityDesc }

type TriggerSecurityID struct{ quickfix.StringValue }

func (f TriggerSecurityID) Tag() tag.Tag { return tag.TriggerSecurityID }

type TriggerSecurityIDSource struct{ quickfix.StringValue }

func (f TriggerSecurityIDSource) Tag() tag.Tag { return tag.TriggerSecurityIDSource }

type TriggerSymbol struct{ quickfix.StringValue }

func (f TriggerSymbol) Tag() tag.Tag { return tag.TriggerSymbol }

type TriggerTradingSessionID struct{ quickfix.StringValue }

func (f TriggerTradingSessionID) Tag() tag.Tag { return tag.TriggerTradingSessionID }

type TriggerTradingSessionSubID struct{ quickfix.StringValue }

func (f TriggerTradingSessionSubID) Tag() tag.Tag { return tag.TriggerTradingSessionSubID }

type TriggerType struct{ quickfix.CharValue }

func (f TriggerType) Tag() tag.Tag { return tag.TriggerType }

type URLLink struct{ quickfix.StringValue }

func (f URLLink) Tag() tag.Tag { return tag.URLLink }

type UnderlyingAdjustedQuantity struct{ quickfix.QtyValue }

func (f UnderlyingAdjustedQuantity) Tag() tag.Tag { return tag.UnderlyingAdjustedQuantity }

type UnderlyingAllocationPercent struct{ quickfix.PercentageValue }

func (f UnderlyingAllocationPercent) Tag() tag.Tag { return tag.UnderlyingAllocationPercent }

type UnderlyingAttachmentPoint struct{ quickfix.PercentageValue }

func (f UnderlyingAttachmentPoint) Tag() tag.Tag { return tag.UnderlyingAttachmentPoint }

type UnderlyingCFICode struct{ quickfix.StringValue }

func (f UnderlyingCFICode) Tag() tag.Tag { return tag.UnderlyingCFICode }

type UnderlyingCPProgram struct{ quickfix.StringValue }

func (f UnderlyingCPProgram) Tag() tag.Tag { return tag.UnderlyingCPProgram }

type UnderlyingCPRegType struct{ quickfix.StringValue }

func (f UnderlyingCPRegType) Tag() tag.Tag { return tag.UnderlyingCPRegType }

type UnderlyingCapValue struct{ quickfix.AmtValue }

func (f UnderlyingCapValue) Tag() tag.Tag { return tag.UnderlyingCapValue }

type UnderlyingCashAmount struct{ quickfix.AmtValue }

func (f UnderlyingCashAmount) Tag() tag.Tag { return tag.UnderlyingCashAmount }

type UnderlyingCashType struct{ quickfix.StringValue }

func (f UnderlyingCashType) Tag() tag.Tag { return tag.UnderlyingCashType }

type UnderlyingCollectAmount struct{ quickfix.AmtValue }

func (f UnderlyingCollectAmount) Tag() tag.Tag { return tag.UnderlyingCollectAmount }

type UnderlyingContractMultiplier struct{ quickfix.FloatValue }

func (f UnderlyingContractMultiplier) Tag() tag.Tag { return tag.UnderlyingContractMultiplier }

type UnderlyingContractMultiplierUnit struct{ quickfix.IntValue }

func (f UnderlyingContractMultiplierUnit) Tag() tag.Tag { return tag.UnderlyingContractMultiplierUnit }

type UnderlyingCountryOfIssue struct{ quickfix.CountryValue }

func (f UnderlyingCountryOfIssue) Tag() tag.Tag { return tag.UnderlyingCountryOfIssue }

type UnderlyingCouponPaymentDate struct{ quickfix.LocalMktDateValue }

func (f UnderlyingCouponPaymentDate) Tag() tag.Tag { return tag.UnderlyingCouponPaymentDate }

type UnderlyingCouponRate struct{ quickfix.PercentageValue }

func (f UnderlyingCouponRate) Tag() tag.Tag { return tag.UnderlyingCouponRate }

type UnderlyingCreditRating struct{ quickfix.StringValue }

func (f UnderlyingCreditRating) Tag() tag.Tag { return tag.UnderlyingCreditRating }

type UnderlyingCurrency struct{ quickfix.CurrencyValue }

func (f UnderlyingCurrency) Tag() tag.Tag { return tag.UnderlyingCurrency }

type UnderlyingCurrentValue struct{ quickfix.AmtValue }

func (f UnderlyingCurrentValue) Tag() tag.Tag { return tag.UnderlyingCurrentValue }

type UnderlyingDeliveryAmount struct{ quickfix.AmtValue }

func (f UnderlyingDeliveryAmount) Tag() tag.Tag { return tag.UnderlyingDeliveryAmount }

type UnderlyingDetachmentPoint struct{ quickfix.PercentageValue }

func (f UnderlyingDetachmentPoint) Tag() tag.Tag { return tag.UnderlyingDetachmentPoint }

type UnderlyingDirtyPrice struct{ quickfix.PriceValue }

func (f UnderlyingDirtyPrice) Tag() tag.Tag { return tag.UnderlyingDirtyPrice }

type UnderlyingEndPrice struct{ quickfix.PriceValue }

func (f UnderlyingEndPrice) Tag() tag.Tag { return tag.UnderlyingEndPrice }

type UnderlyingEndValue struct{ quickfix.AmtValue }

func (f UnderlyingEndValue) Tag() tag.Tag { return tag.UnderlyingEndValue }

type UnderlyingExerciseStyle struct{ quickfix.IntValue }

func (f UnderlyingExerciseStyle) Tag() tag.Tag { return tag.UnderlyingExerciseStyle }

type UnderlyingFXRate struct{ quickfix.FloatValue }

func (f UnderlyingFXRate) Tag() tag.Tag { return tag.UnderlyingFXRate }

type UnderlyingFXRateCalc struct{ quickfix.CharValue }

func (f UnderlyingFXRateCalc) Tag() tag.Tag { return tag.UnderlyingFXRateCalc }

type UnderlyingFactor struct{ quickfix.FloatValue }

func (f UnderlyingFactor) Tag() tag.Tag { return tag.UnderlyingFactor }

type UnderlyingFlowScheduleType struct{ quickfix.IntValue }

func (f UnderlyingFlowScheduleType) Tag() tag.Tag { return tag.UnderlyingFlowScheduleType }

type UnderlyingIDSource struct{ quickfix.StringValue }

func (f UnderlyingIDSource) Tag() tag.Tag { return tag.UnderlyingIDSource }

type UnderlyingInstrRegistry struct{ quickfix.StringValue }

func (f UnderlyingInstrRegistry) Tag() tag.Tag { return tag.UnderlyingInstrRegistry }

type UnderlyingInstrumentPartyID struct{ quickfix.StringValue }

func (f UnderlyingInstrumentPartyID) Tag() tag.Tag { return tag.UnderlyingInstrumentPartyID }

type UnderlyingInstrumentPartyIDSource struct{ quickfix.CharValue }

func (f UnderlyingInstrumentPartyIDSource) Tag() tag.Tag { return tag.UnderlyingInstrumentPartyIDSource }

type UnderlyingInstrumentPartyRole struct{ quickfix.IntValue }

func (f UnderlyingInstrumentPartyRole) Tag() tag.Tag { return tag.UnderlyingInstrumentPartyRole }

type UnderlyingInstrumentPartySubID struct{ quickfix.StringValue }

func (f UnderlyingInstrumentPartySubID) Tag() tag.Tag { return tag.UnderlyingInstrumentPartySubID }

type UnderlyingInstrumentPartySubIDType struct{ quickfix.IntValue }

func (f UnderlyingInstrumentPartySubIDType) Tag() tag.Tag {
	return tag.UnderlyingInstrumentPartySubIDType
}

type UnderlyingIssueDate struct{ quickfix.LocalMktDateValue }

func (f UnderlyingIssueDate) Tag() tag.Tag { return tag.UnderlyingIssueDate }

type UnderlyingIssuer struct{ quickfix.StringValue }

func (f UnderlyingIssuer) Tag() tag.Tag { return tag.UnderlyingIssuer }

type UnderlyingLastPx struct{ quickfix.PriceValue }

func (f UnderlyingLastPx) Tag() tag.Tag { return tag.UnderlyingLastPx }

type UnderlyingLastQty struct{ quickfix.QtyValue }

func (f UnderlyingLastQty) Tag() tag.Tag { return tag.UnderlyingLastQty }

type UnderlyingLegCFICode struct{ quickfix.StringValue }

func (f UnderlyingLegCFICode) Tag() tag.Tag { return tag.UnderlyingLegCFICode }

type UnderlyingLegMaturityDate struct{ quickfix.LocalMktDateValue }

func (f UnderlyingLegMaturityDate) Tag() tag.Tag { return tag.UnderlyingLegMaturityDate }

type UnderlyingLegMaturityMonthYear struct{ quickfix.MonthYearValue }

func (f UnderlyingLegMaturityMonthYear) Tag() tag.Tag { return tag.UnderlyingLegMaturityMonthYear }

type UnderlyingLegMaturityTime struct{ quickfix.TZTimeOnlyValue }

func (f UnderlyingLegMaturityTime) Tag() tag.Tag { return tag.UnderlyingLegMaturityTime }

type UnderlyingLegOptAttribute struct{ quickfix.CharValue }

func (f UnderlyingLegOptAttribute) Tag() tag.Tag { return tag.UnderlyingLegOptAttribute }

type UnderlyingLegPutOrCall struct{ quickfix.IntValue }

func (f UnderlyingLegPutOrCall) Tag() tag.Tag { return tag.UnderlyingLegPutOrCall }

type UnderlyingLegSecurityAltID struct{ quickfix.StringValue }

func (f UnderlyingLegSecurityAltID) Tag() tag.Tag { return tag.UnderlyingLegSecurityAltID }

type UnderlyingLegSecurityAltIDSource struct{ quickfix.StringValue }

func (f UnderlyingLegSecurityAltIDSource) Tag() tag.Tag { return tag.UnderlyingLegSecurityAltIDSource }

type UnderlyingLegSecurityDesc struct{ quickfix.StringValue }

func (f UnderlyingLegSecurityDesc) Tag() tag.Tag { return tag.UnderlyingLegSecurityDesc }

type UnderlyingLegSecurityExchange struct{ quickfix.StringValue }

func (f UnderlyingLegSecurityExchange) Tag() tag.Tag { return tag.UnderlyingLegSecurityExchange }

type UnderlyingLegSecurityID struct{ quickfix.StringValue }

func (f UnderlyingLegSecurityID) Tag() tag.Tag { return tag.UnderlyingLegSecurityID }

type UnderlyingLegSecurityIDSource struct{ quickfix.StringValue }

func (f UnderlyingLegSecurityIDSource) Tag() tag.Tag { return tag.UnderlyingLegSecurityIDSource }

type UnderlyingLegSecuritySubType struct{ quickfix.StringValue }

func (f UnderlyingLegSecuritySubType) Tag() tag.Tag { return tag.UnderlyingLegSecuritySubType }

type UnderlyingLegSecurityType struct{ quickfix.StringValue }

func (f UnderlyingLegSecurityType) Tag() tag.Tag { return tag.UnderlyingLegSecurityType }

type UnderlyingLegStrikePrice struct{ quickfix.PriceValue }

func (f UnderlyingLegStrikePrice) Tag() tag.Tag { return tag.UnderlyingLegStrikePrice }

type UnderlyingLegSymbol struct{ quickfix.StringValue }

func (f UnderlyingLegSymbol) Tag() tag.Tag { return tag.UnderlyingLegSymbol }

type UnderlyingLegSymbolSfx struct{ quickfix.StringValue }

func (f UnderlyingLegSymbolSfx) Tag() tag.Tag { return tag.UnderlyingLegSymbolSfx }

type UnderlyingLocaleOfIssue struct{ quickfix.StringValue }

func (f UnderlyingLocaleOfIssue) Tag() tag.Tag { return tag.UnderlyingLocaleOfIssue }

type UnderlyingMaturityDate struct{ quickfix.LocalMktDateValue }

func (f UnderlyingMaturityDate) Tag() tag.Tag { return tag.UnderlyingMaturityDate }

type UnderlyingMaturityDay struct{ quickfix.DayOfMonthValue }

func (f UnderlyingMaturityDay) Tag() tag.Tag { return tag.UnderlyingMaturityDay }

type UnderlyingMaturityMonthYear struct{ quickfix.MonthYearValue }

func (f UnderlyingMaturityMonthYear) Tag() tag.Tag { return tag.UnderlyingMaturityMonthYear }

type UnderlyingMaturityTime struct{ quickfix.TZTimeOnlyValue }

func (f UnderlyingMaturityTime) Tag() tag.Tag { return tag.UnderlyingMaturityTime }

type UnderlyingNotionalPercentageOutstanding struct{ quickfix.PercentageValue }

func (f UnderlyingNotionalPercentageOutstanding) Tag() tag.Tag {
	return tag.UnderlyingNotionalPercentageOutstanding
}

type UnderlyingOptAttribute struct{ quickfix.CharValue }

func (f UnderlyingOptAttribute) Tag() tag.Tag { return tag.UnderlyingOptAttribute }

type UnderlyingOriginalNotionalPercentageOutstanding struct{ quickfix.PercentageValue }

func (f UnderlyingOriginalNotionalPercentageOutstanding) Tag() tag.Tag {
	return tag.UnderlyingOriginalNotionalPercentageOutstanding
}

type UnderlyingPayAmount struct{ quickfix.AmtValue }

func (f UnderlyingPayAmount) Tag() tag.Tag { return tag.UnderlyingPayAmount }

type UnderlyingPriceDeterminationMethod struct{ quickfix.IntValue }

func (f UnderlyingPriceDeterminationMethod) Tag() tag.Tag {
	return tag.UnderlyingPriceDeterminationMethod
}

type UnderlyingPriceUnitOfMeasure struct{ quickfix.StringValue }

func (f UnderlyingPriceUnitOfMeasure) Tag() tag.Tag { return tag.UnderlyingPriceUnitOfMeasure }

type UnderlyingPriceUnitOfMeasureQty struct{ quickfix.QtyValue }

func (f UnderlyingPriceUnitOfMeasureQty) Tag() tag.Tag { return tag.UnderlyingPriceUnitOfMeasureQty }

type UnderlyingProduct struct{ quickfix.IntValue }

func (f UnderlyingProduct) Tag() tag.Tag { return tag.UnderlyingProduct }

type UnderlyingPutOrCall struct{ quickfix.IntValue }

func (f UnderlyingPutOrCall) Tag() tag.Tag { return tag.UnderlyingPutOrCall }

type UnderlyingPx struct{ quickfix.PriceValue }

func (f UnderlyingPx) Tag() tag.Tag { return tag.UnderlyingPx }

type UnderlyingQty struct{ quickfix.QtyValue }

func (f UnderlyingQty) Tag() tag.Tag { return tag.UnderlyingQty }

type UnderlyingRedemptionDate struct{ quickfix.LocalMktDateValue }

func (f UnderlyingRedemptionDate) Tag() tag.Tag { return tag.UnderlyingRedemptionDate }

type UnderlyingRepoCollateralSecurityType struct{ quickfix.IntValue }

func (f UnderlyingRepoCollateralSecurityType) Tag() tag.Tag {
	return tag.UnderlyingRepoCollateralSecurityType
}

type UnderlyingRepurchaseRate struct{ quickfix.PercentageValue }

func (f UnderlyingRepurchaseRate) Tag() tag.Tag { return tag.UnderlyingRepurchaseRate }

type UnderlyingRepurchaseTerm struct{ quickfix.IntValue }

func (f UnderlyingRepurchaseTerm) Tag() tag.Tag { return tag.UnderlyingRepurchaseTerm }

type UnderlyingRestructuringType struct{ quickfix.StringValue }

func (f UnderlyingRestructuringType) Tag() tag.Tag { return tag.UnderlyingRestructuringType }

type UnderlyingSecurityAltID struct{ quickfix.StringValue }

func (f UnderlyingSecurityAltID) Tag() tag.Tag { return tag.UnderlyingSecurityAltID }

type UnderlyingSecurityAltIDSource struct{ quickfix.StringValue }

func (f UnderlyingSecurityAltIDSource) Tag() tag.Tag { return tag.UnderlyingSecurityAltIDSource }

type UnderlyingSecurityDesc struct{ quickfix.StringValue }

func (f UnderlyingSecurityDesc) Tag() tag.Tag { return tag.UnderlyingSecurityDesc }

type UnderlyingSecurityExchange struct{ quickfix.ExchangeValue }

func (f UnderlyingSecurityExchange) Tag() tag.Tag { return tag.UnderlyingSecurityExchange }

type UnderlyingSecurityID struct{ quickfix.StringValue }

func (f UnderlyingSecurityID) Tag() tag.Tag { return tag.UnderlyingSecurityID }

type UnderlyingSecurityIDSource struct{ quickfix.StringValue }

func (f UnderlyingSecurityIDSource) Tag() tag.Tag { return tag.UnderlyingSecurityIDSource }

type UnderlyingSecuritySubType struct{ quickfix.StringValue }

func (f UnderlyingSecuritySubType) Tag() tag.Tag { return tag.UnderlyingSecuritySubType }

type UnderlyingSecurityType struct{ quickfix.StringValue }

func (f UnderlyingSecurityType) Tag() tag.Tag { return tag.UnderlyingSecurityType }

type UnderlyingSeniority struct{ quickfix.StringValue }

func (f UnderlyingSeniority) Tag() tag.Tag { return tag.UnderlyingSeniority }

type UnderlyingSettlMethod struct{ quickfix.StringValue }

func (f UnderlyingSettlMethod) Tag() tag.Tag { return tag.UnderlyingSettlMethod }

type UnderlyingSettlPrice struct{ quickfix.PriceValue }

func (f UnderlyingSettlPrice) Tag() tag.Tag { return tag.UnderlyingSettlPrice }

type UnderlyingSettlPriceType struct{ quickfix.IntValue }

func (f UnderlyingSettlPriceType) Tag() tag.Tag { return tag.UnderlyingSettlPriceType }

type UnderlyingSettlementDate struct{ quickfix.LocalMktDateValue }

func (f UnderlyingSettlementDate) Tag() tag.Tag { return tag.UnderlyingSettlementDate }

type UnderlyingSettlementStatus struct{ quickfix.StringValue }

func (f UnderlyingSettlementStatus) Tag() tag.Tag { return tag.UnderlyingSettlementStatus }

type UnderlyingSettlementType struct{ quickfix.IntValue }

func (f UnderlyingSettlementType) Tag() tag.Tag { return tag.UnderlyingSettlementType }

type UnderlyingStartValue struct{ quickfix.AmtValue }

func (f UnderlyingStartValue) Tag() tag.Tag { return tag.UnderlyingStartValue }

type UnderlyingStateOrProvinceOfIssue struct{ quickfix.StringValue }

func (f UnderlyingStateOrProvinceOfIssue) Tag() tag.Tag { return tag.UnderlyingStateOrProvinceOfIssue }

type UnderlyingStipType struct{ quickfix.StringValue }

func (f UnderlyingStipType) Tag() tag.Tag { return tag.UnderlyingStipType }

type UnderlyingStipValue struct{ quickfix.StringValue }

func (f UnderlyingStipValue) Tag() tag.Tag { return tag.UnderlyingStipValue }

type UnderlyingStrikeCurrency struct{ quickfix.CurrencyValue }

func (f UnderlyingStrikeCurrency) Tag() tag.Tag { return tag.UnderlyingStrikeCurrency }

type UnderlyingStrikePrice struct{ quickfix.PriceValue }

func (f UnderlyingStrikePrice) Tag() tag.Tag { return tag.UnderlyingStrikePrice }

type UnderlyingSymbol struct{ quickfix.StringValue }

func (f UnderlyingSymbol) Tag() tag.Tag { return tag.UnderlyingSymbol }

type UnderlyingSymbolSfx struct{ quickfix.StringValue }

func (f UnderlyingSymbolSfx) Tag() tag.Tag { return tag.UnderlyingSymbolSfx }

type UnderlyingTimeUnit struct{ quickfix.StringValue }

func (f UnderlyingTimeUnit) Tag() tag.Tag { return tag.UnderlyingTimeUnit }

type UnderlyingTradingSessionID struct{ quickfix.StringValue }

func (f UnderlyingTradingSessionID) Tag() tag.Tag { return tag.UnderlyingTradingSessionID }

type UnderlyingTradingSessionSubID struct{ quickfix.StringValue }

func (f UnderlyingTradingSessionSubID) Tag() tag.Tag { return tag.UnderlyingTradingSessionSubID }

type UnderlyingUnitOfMeasure struct{ quickfix.StringValue }

func (f UnderlyingUnitOfMeasure) Tag() tag.Tag { return tag.UnderlyingUnitOfMeasure }

type UnderlyingUnitOfMeasureQty struct{ quickfix.QtyValue }

func (f UnderlyingUnitOfMeasureQty) Tag() tag.Tag { return tag.UnderlyingUnitOfMeasureQty }

type UndlyInstrumentPartyID struct{ quickfix.StringValue }

func (f UndlyInstrumentPartyID) Tag() tag.Tag { return tag.UndlyInstrumentPartyID }

type UndlyInstrumentPartyIDSource struct{ quickfix.CharValue }

func (f UndlyInstrumentPartyIDSource) Tag() tag.Tag { return tag.UndlyInstrumentPartyIDSource }

type UndlyInstrumentPartyRole struct{ quickfix.IntValue }

func (f UndlyInstrumentPartyRole) Tag() tag.Tag { return tag.UndlyInstrumentPartyRole }

type UndlyInstrumentPartySubID struct{ quickfix.StringValue }

func (f UndlyInstrumentPartySubID) Tag() tag.Tag { return tag.UndlyInstrumentPartySubID }

type UndlyInstrumentPartySubIDType struct{ quickfix.IntValue }

func (f UndlyInstrumentPartySubIDType) Tag() tag.Tag { return tag.UndlyInstrumentPartySubIDType }

type UnitOfMeasure struct{ quickfix.StringValue }

func (f UnitOfMeasure) Tag() tag.Tag { return tag.UnitOfMeasure }

type UnitOfMeasureQty struct{ quickfix.QtyValue }

func (f UnitOfMeasureQty) Tag() tag.Tag { return tag.UnitOfMeasureQty }

type UnsolicitedIndicator struct{ quickfix.BooleanValue }

func (f UnsolicitedIndicator) Tag() tag.Tag { return tag.UnsolicitedIndicator }

type Urgency struct{ quickfix.CharValue }

func (f Urgency) Tag() tag.Tag { return tag.Urgency }

type UserRequestID struct{ quickfix.StringValue }

func (f UserRequestID) Tag() tag.Tag { return tag.UserRequestID }

type UserRequestType struct{ quickfix.IntValue }

func (f UserRequestType) Tag() tag.Tag { return tag.UserRequestType }

type UserStatus struct{ quickfix.IntValue }

func (f UserStatus) Tag() tag.Tag { return tag.UserStatus }

type UserStatusText struct{ quickfix.StringValue }

func (f UserStatusText) Tag() tag.Tag { return tag.UserStatusText }

type Username struct{ quickfix.StringValue }

func (f Username) Tag() tag.Tag { return tag.Username }

type ValidUntilTime struct{ quickfix.UTCTimestampValue }

func (f ValidUntilTime) Tag() tag.Tag { return tag.ValidUntilTime }

type ValuationMethod struct{ quickfix.StringValue }

func (f ValuationMethod) Tag() tag.Tag { return tag.ValuationMethod }

type ValueOfFutures struct{ quickfix.AmtValue }

func (f ValueOfFutures) Tag() tag.Tag { return tag.ValueOfFutures }

type VenueType struct{ quickfix.CharValue }

func (f VenueType) Tag() tag.Tag { return tag.VenueType }

type Volatility struct{ quickfix.FloatValue }

func (f Volatility) Tag() tag.Tag { return tag.Volatility }

type WaveNo struct{ quickfix.StringValue }

func (f WaveNo) Tag() tag.Tag { return tag.WaveNo }

type WorkingIndicator struct{ quickfix.BooleanValue }

func (f WorkingIndicator) Tag() tag.Tag { return tag.WorkingIndicator }

type WtAverageLiquidity struct{ quickfix.PercentageValue }

func (f WtAverageLiquidity) Tag() tag.Tag { return tag.WtAverageLiquidity }

type XmlData struct{ quickfix.DataValue }

func (f XmlData) Tag() tag.Tag { return tag.XmlData }

type XmlDataLen struct{ quickfix.LengthValue }

func (f XmlDataLen) Tag() tag.Tag { return tag.XmlDataLen }

type Yield struct{ quickfix.PercentageValue }

func (f Yield) Tag() tag.Tag { return tag.Yield }

type YieldCalcDate struct{ quickfix.LocalMktDateValue }

func (f YieldCalcDate) Tag() tag.Tag { return tag.YieldCalcDate }

type YieldRedemptionDate struct{ quickfix.LocalMktDateValue }

func (f YieldRedemptionDate) Tag() tag.Tag { return tag.YieldRedemptionDate }

type YieldRedemptionPrice struct{ quickfix.PriceValue }

func (f YieldRedemptionPrice) Tag() tag.Tag { return tag.YieldRedemptionPrice }

type YieldRedemptionPriceType struct{ quickfix.IntValue }

func (f YieldRedemptionPriceType) Tag() tag.Tag { return tag.YieldRedemptionPriceType }

type YieldType struct{ quickfix.StringValue }

func (f YieldType) Tag() tag.Tag { return tag.YieldType }
