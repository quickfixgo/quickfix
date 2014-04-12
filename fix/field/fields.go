package field

import (
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/tag"
	"github.com/quickfixgo/quickfix/message"
)

type Account struct{ message.StringValue }

func (f Account) Tag() fix.Tag { return tag.Account }

type AccountType struct{ message.IntValue }

func (f AccountType) Tag() fix.Tag { return tag.AccountType }

type AccruedInterestAmt struct{ message.AmtValue }

func (f AccruedInterestAmt) Tag() fix.Tag { return tag.AccruedInterestAmt }

type AccruedInterestRate struct{ message.PercentageValue }

func (f AccruedInterestRate) Tag() fix.Tag { return tag.AccruedInterestRate }

type AcctIDSource struct{ message.IntValue }

func (f AcctIDSource) Tag() fix.Tag { return tag.AcctIDSource }

type Adjustment struct{ message.IntValue }

func (f Adjustment) Tag() fix.Tag { return tag.Adjustment }

type AdjustmentType struct{ message.IntValue }

func (f AdjustmentType) Tag() fix.Tag { return tag.AdjustmentType }

type AdvId struct{ message.StringValue }

func (f AdvId) Tag() fix.Tag { return tag.AdvId }

type AdvRefID struct{ message.StringValue }

func (f AdvRefID) Tag() fix.Tag { return tag.AdvRefID }

type AdvSide struct{ message.CharValue }

func (f AdvSide) Tag() fix.Tag { return tag.AdvSide }

type AdvTransType struct{ message.StringValue }

func (f AdvTransType) Tag() fix.Tag { return tag.AdvTransType }

type AffectedOrderID struct{ message.StringValue }

func (f AffectedOrderID) Tag() fix.Tag { return tag.AffectedOrderID }

type AffectedSecondaryOrderID struct{ message.StringValue }

func (f AffectedSecondaryOrderID) Tag() fix.Tag { return tag.AffectedSecondaryOrderID }

type AffirmStatus struct{ message.IntValue }

func (f AffirmStatus) Tag() fix.Tag { return tag.AffirmStatus }

type AggregatedBook struct{ message.BooleanValue }

func (f AggregatedBook) Tag() fix.Tag { return tag.AggregatedBook }

type AggressorIndicator struct{ message.BooleanValue }

func (f AggressorIndicator) Tag() fix.Tag { return tag.AggressorIndicator }

type AgreementCurrency struct{ message.CurrencyValue }

func (f AgreementCurrency) Tag() fix.Tag { return tag.AgreementCurrency }

type AgreementDate struct{ message.LocalMktDateValue }

func (f AgreementDate) Tag() fix.Tag { return tag.AgreementDate }

type AgreementDesc struct{ message.StringValue }

func (f AgreementDesc) Tag() fix.Tag { return tag.AgreementDesc }

type AgreementID struct{ message.StringValue }

func (f AgreementID) Tag() fix.Tag { return tag.AgreementID }

type AllocAccount struct{ message.StringValue }

func (f AllocAccount) Tag() fix.Tag { return tag.AllocAccount }

type AllocAccountType struct{ message.IntValue }

func (f AllocAccountType) Tag() fix.Tag { return tag.AllocAccountType }

type AllocAccruedInterestAmt struct{ message.AmtValue }

func (f AllocAccruedInterestAmt) Tag() fix.Tag { return tag.AllocAccruedInterestAmt }

type AllocAcctIDSource struct{ message.IntValue }

func (f AllocAcctIDSource) Tag() fix.Tag { return tag.AllocAcctIDSource }

type AllocAvgPx struct{ message.PriceValue }

func (f AllocAvgPx) Tag() fix.Tag { return tag.AllocAvgPx }

type AllocCancReplaceReason struct{ message.IntValue }

func (f AllocCancReplaceReason) Tag() fix.Tag { return tag.AllocCancReplaceReason }

type AllocClearingFeeIndicator struct{ message.StringValue }

func (f AllocClearingFeeIndicator) Tag() fix.Tag { return tag.AllocClearingFeeIndicator }

type AllocCustomerCapacity struct{ message.StringValue }

func (f AllocCustomerCapacity) Tag() fix.Tag { return tag.AllocCustomerCapacity }

type AllocHandlInst struct{ message.IntValue }

func (f AllocHandlInst) Tag() fix.Tag { return tag.AllocHandlInst }

type AllocID struct{ message.StringValue }

func (f AllocID) Tag() fix.Tag { return tag.AllocID }

type AllocInterestAtMaturity struct{ message.AmtValue }

func (f AllocInterestAtMaturity) Tag() fix.Tag { return tag.AllocInterestAtMaturity }

type AllocIntermedReqType struct{ message.IntValue }

func (f AllocIntermedReqType) Tag() fix.Tag { return tag.AllocIntermedReqType }

type AllocLinkID struct{ message.StringValue }

func (f AllocLinkID) Tag() fix.Tag { return tag.AllocLinkID }

type AllocLinkType struct{ message.IntValue }

func (f AllocLinkType) Tag() fix.Tag { return tag.AllocLinkType }

type AllocMethod struct{ message.IntValue }

func (f AllocMethod) Tag() fix.Tag { return tag.AllocMethod }

type AllocNetMoney struct{ message.AmtValue }

func (f AllocNetMoney) Tag() fix.Tag { return tag.AllocNetMoney }

type AllocNoOrdersType struct{ message.IntValue }

func (f AllocNoOrdersType) Tag() fix.Tag { return tag.AllocNoOrdersType }

type AllocPositionEffect struct{ message.CharValue }

func (f AllocPositionEffect) Tag() fix.Tag { return tag.AllocPositionEffect }

type AllocPrice struct{ message.PriceValue }

func (f AllocPrice) Tag() fix.Tag { return tag.AllocPrice }

type AllocQty struct{ message.QtyValue }

func (f AllocQty) Tag() fix.Tag { return tag.AllocQty }

type AllocRejCode struct{ message.IntValue }

func (f AllocRejCode) Tag() fix.Tag { return tag.AllocRejCode }

type AllocReportID struct{ message.StringValue }

func (f AllocReportID) Tag() fix.Tag { return tag.AllocReportID }

type AllocReportRefID struct{ message.StringValue }

func (f AllocReportRefID) Tag() fix.Tag { return tag.AllocReportRefID }

type AllocReportType struct{ message.IntValue }

func (f AllocReportType) Tag() fix.Tag { return tag.AllocReportType }

type AllocSettlCurrAmt struct{ message.AmtValue }

func (f AllocSettlCurrAmt) Tag() fix.Tag { return tag.AllocSettlCurrAmt }

type AllocSettlCurrency struct{ message.CurrencyValue }

func (f AllocSettlCurrency) Tag() fix.Tag { return tag.AllocSettlCurrency }

type AllocSettlInstType struct{ message.IntValue }

func (f AllocSettlInstType) Tag() fix.Tag { return tag.AllocSettlInstType }

type AllocShares struct{ message.QtyValue }

func (f AllocShares) Tag() fix.Tag { return tag.AllocShares }

type AllocStatus struct{ message.IntValue }

func (f AllocStatus) Tag() fix.Tag { return tag.AllocStatus }

type AllocText struct{ message.StringValue }

func (f AllocText) Tag() fix.Tag { return tag.AllocText }

type AllocTransType struct{ message.CharValue }

func (f AllocTransType) Tag() fix.Tag { return tag.AllocTransType }

type AllocType struct{ message.IntValue }

func (f AllocType) Tag() fix.Tag { return tag.AllocType }

type AllowableOneSidednessCurr struct{ message.CurrencyValue }

func (f AllowableOneSidednessCurr) Tag() fix.Tag { return tag.AllowableOneSidednessCurr }

type AllowableOneSidednessPct struct{ message.PercentageValue }

func (f AllowableOneSidednessPct) Tag() fix.Tag { return tag.AllowableOneSidednessPct }

type AllowableOneSidednessValue struct{ message.AmtValue }

func (f AllowableOneSidednessValue) Tag() fix.Tag { return tag.AllowableOneSidednessValue }

type AltMDSourceID struct{ message.StringValue }

func (f AltMDSourceID) Tag() fix.Tag { return tag.AltMDSourceID }

type ApplBegSeqNum struct{ message.SeqNumValue }

func (f ApplBegSeqNum) Tag() fix.Tag { return tag.ApplBegSeqNum }

type ApplEndSeqNum struct{ message.SeqNumValue }

func (f ApplEndSeqNum) Tag() fix.Tag { return tag.ApplEndSeqNum }

type ApplExtID struct{ message.IntValue }

func (f ApplExtID) Tag() fix.Tag { return tag.ApplExtID }

type ApplID struct{ message.StringValue }

func (f ApplID) Tag() fix.Tag { return tag.ApplID }

type ApplLastSeqNum struct{ message.SeqNumValue }

func (f ApplLastSeqNum) Tag() fix.Tag { return tag.ApplLastSeqNum }

type ApplNewSeqNum struct{ message.SeqNumValue }

func (f ApplNewSeqNum) Tag() fix.Tag { return tag.ApplNewSeqNum }

type ApplQueueAction struct{ message.IntValue }

func (f ApplQueueAction) Tag() fix.Tag { return tag.ApplQueueAction }

type ApplQueueDepth struct{ message.IntValue }

func (f ApplQueueDepth) Tag() fix.Tag { return tag.ApplQueueDepth }

type ApplQueueMax struct{ message.IntValue }

func (f ApplQueueMax) Tag() fix.Tag { return tag.ApplQueueMax }

type ApplQueueResolution struct{ message.IntValue }

func (f ApplQueueResolution) Tag() fix.Tag { return tag.ApplQueueResolution }

type ApplReportID struct{ message.StringValue }

func (f ApplReportID) Tag() fix.Tag { return tag.ApplReportID }

type ApplReportType struct{ message.IntValue }

func (f ApplReportType) Tag() fix.Tag { return tag.ApplReportType }

type ApplReqID struct{ message.StringValue }

func (f ApplReqID) Tag() fix.Tag { return tag.ApplReqID }

type ApplReqType struct{ message.IntValue }

func (f ApplReqType) Tag() fix.Tag { return tag.ApplReqType }

type ApplResendFlag struct{ message.BooleanValue }

func (f ApplResendFlag) Tag() fix.Tag { return tag.ApplResendFlag }

type ApplResponseError struct{ message.IntValue }

func (f ApplResponseError) Tag() fix.Tag { return tag.ApplResponseError }

type ApplResponseID struct{ message.StringValue }

func (f ApplResponseID) Tag() fix.Tag { return tag.ApplResponseID }

type ApplResponseType struct{ message.IntValue }

func (f ApplResponseType) Tag() fix.Tag { return tag.ApplResponseType }

type ApplSeqNum struct{ message.SeqNumValue }

func (f ApplSeqNum) Tag() fix.Tag { return tag.ApplSeqNum }

type ApplTotalMessageCount struct{ message.IntValue }

func (f ApplTotalMessageCount) Tag() fix.Tag { return tag.ApplTotalMessageCount }

type ApplVerID struct{ message.StringValue }

func (f ApplVerID) Tag() fix.Tag { return tag.ApplVerID }

type AsOfIndicator struct{ message.CharValue }

func (f AsOfIndicator) Tag() fix.Tag { return tag.AsOfIndicator }

type AsgnReqID struct{ message.StringValue }

func (f AsgnReqID) Tag() fix.Tag { return tag.AsgnReqID }

type AsgnRptID struct{ message.StringValue }

func (f AsgnRptID) Tag() fix.Tag { return tag.AsgnRptID }

type AssignmentMethod struct{ message.CharValue }

func (f AssignmentMethod) Tag() fix.Tag { return tag.AssignmentMethod }

type AssignmentUnit struct{ message.QtyValue }

func (f AssignmentUnit) Tag() fix.Tag { return tag.AssignmentUnit }

type AttachmentPoint struct{ message.PercentageValue }

func (f AttachmentPoint) Tag() fix.Tag { return tag.AttachmentPoint }

type AutoAcceptIndicator struct{ message.BooleanValue }

func (f AutoAcceptIndicator) Tag() fix.Tag { return tag.AutoAcceptIndicator }

type AvgParPx struct{ message.PriceValue }

func (f AvgParPx) Tag() fix.Tag { return tag.AvgParPx }

type AvgPrxPrecision struct{ message.IntValue }

func (f AvgPrxPrecision) Tag() fix.Tag { return tag.AvgPrxPrecision }

type AvgPx struct{ message.PriceValue }

func (f AvgPx) Tag() fix.Tag { return tag.AvgPx }

type AvgPxIndicator struct{ message.IntValue }

func (f AvgPxIndicator) Tag() fix.Tag { return tag.AvgPxIndicator }

type AvgPxPrecision struct{ message.IntValue }

func (f AvgPxPrecision) Tag() fix.Tag { return tag.AvgPxPrecision }

type BasisFeatureDate struct{ message.LocalMktDateValue }

func (f BasisFeatureDate) Tag() fix.Tag { return tag.BasisFeatureDate }

type BasisFeaturePrice struct{ message.PriceValue }

func (f BasisFeaturePrice) Tag() fix.Tag { return tag.BasisFeaturePrice }

type BasisPxType struct{ message.CharValue }

func (f BasisPxType) Tag() fix.Tag { return tag.BasisPxType }

type BeginSeqNo struct{ message.SeqNumValue }

func (f BeginSeqNo) Tag() fix.Tag { return tag.BeginSeqNo }

type BeginString struct{ message.StringValue }

func (f BeginString) Tag() fix.Tag { return tag.BeginString }

type Benchmark struct{ message.CharValue }

func (f Benchmark) Tag() fix.Tag { return tag.Benchmark }

type BenchmarkCurveCurrency struct{ message.CurrencyValue }

func (f BenchmarkCurveCurrency) Tag() fix.Tag { return tag.BenchmarkCurveCurrency }

type BenchmarkCurveName struct{ message.StringValue }

func (f BenchmarkCurveName) Tag() fix.Tag { return tag.BenchmarkCurveName }

type BenchmarkCurvePoint struct{ message.StringValue }

func (f BenchmarkCurvePoint) Tag() fix.Tag { return tag.BenchmarkCurvePoint }

type BenchmarkPrice struct{ message.PriceValue }

func (f BenchmarkPrice) Tag() fix.Tag { return tag.BenchmarkPrice }

type BenchmarkPriceType struct{ message.IntValue }

func (f BenchmarkPriceType) Tag() fix.Tag { return tag.BenchmarkPriceType }

type BenchmarkSecurityID struct{ message.StringValue }

func (f BenchmarkSecurityID) Tag() fix.Tag { return tag.BenchmarkSecurityID }

type BenchmarkSecurityIDSource struct{ message.StringValue }

func (f BenchmarkSecurityIDSource) Tag() fix.Tag { return tag.BenchmarkSecurityIDSource }

type BidDescriptor struct{ message.StringValue }

func (f BidDescriptor) Tag() fix.Tag { return tag.BidDescriptor }

type BidDescriptorType struct{ message.IntValue }

func (f BidDescriptorType) Tag() fix.Tag { return tag.BidDescriptorType }

type BidForwardPoints struct{ message.PriceOffsetValue }

func (f BidForwardPoints) Tag() fix.Tag { return tag.BidForwardPoints }

type BidForwardPoints2 struct{ message.PriceOffsetValue }

func (f BidForwardPoints2) Tag() fix.Tag { return tag.BidForwardPoints2 }

type BidID struct{ message.StringValue }

func (f BidID) Tag() fix.Tag { return tag.BidID }

type BidPx struct{ message.PriceValue }

func (f BidPx) Tag() fix.Tag { return tag.BidPx }

type BidRequestTransType struct{ message.CharValue }

func (f BidRequestTransType) Tag() fix.Tag { return tag.BidRequestTransType }

type BidSize struct{ message.QtyValue }

func (f BidSize) Tag() fix.Tag { return tag.BidSize }

type BidSpotRate struct{ message.PriceValue }

func (f BidSpotRate) Tag() fix.Tag { return tag.BidSpotRate }

type BidSwapPoints struct{ message.PriceOffsetValue }

func (f BidSwapPoints) Tag() fix.Tag { return tag.BidSwapPoints }

type BidTradeType struct{ message.CharValue }

func (f BidTradeType) Tag() fix.Tag { return tag.BidTradeType }

type BidType struct{ message.IntValue }

func (f BidType) Tag() fix.Tag { return tag.BidType }

type BidYield struct{ message.PercentageValue }

func (f BidYield) Tag() fix.Tag { return tag.BidYield }

type BodyLength struct{ message.LengthValue }

func (f BodyLength) Tag() fix.Tag { return tag.BodyLength }

type BookingRefID struct{ message.StringValue }

func (f BookingRefID) Tag() fix.Tag { return tag.BookingRefID }

type BookingType struct{ message.IntValue }

func (f BookingType) Tag() fix.Tag { return tag.BookingType }

type BookingUnit struct{ message.CharValue }

func (f BookingUnit) Tag() fix.Tag { return tag.BookingUnit }

type BrokerOfCredit struct{ message.StringValue }

func (f BrokerOfCredit) Tag() fix.Tag { return tag.BrokerOfCredit }

type BusinessRejectReason struct{ message.IntValue }

func (f BusinessRejectReason) Tag() fix.Tag { return tag.BusinessRejectReason }

type BusinessRejectRefID struct{ message.StringValue }

func (f BusinessRejectRefID) Tag() fix.Tag { return tag.BusinessRejectRefID }

type BuyVolume struct{ message.QtyValue }

func (f BuyVolume) Tag() fix.Tag { return tag.BuyVolume }

type CFICode struct{ message.StringValue }

func (f CFICode) Tag() fix.Tag { return tag.CFICode }

type CPProgram struct{ message.IntValue }

func (f CPProgram) Tag() fix.Tag { return tag.CPProgram }

type CPRegType struct{ message.StringValue }

func (f CPRegType) Tag() fix.Tag { return tag.CPRegType }

type CalculatedCcyLastQty struct{ message.QtyValue }

func (f CalculatedCcyLastQty) Tag() fix.Tag { return tag.CalculatedCcyLastQty }

type CancellationRights struct{ message.CharValue }

func (f CancellationRights) Tag() fix.Tag { return tag.CancellationRights }

type CapPrice struct{ message.PriceValue }

func (f CapPrice) Tag() fix.Tag { return tag.CapPrice }

type CardExpDate struct{ message.LocalMktDateValue }

func (f CardExpDate) Tag() fix.Tag { return tag.CardExpDate }

type CardHolderName struct{ message.StringValue }

func (f CardHolderName) Tag() fix.Tag { return tag.CardHolderName }

type CardIssNo struct{ message.StringValue }

func (f CardIssNo) Tag() fix.Tag { return tag.CardIssNo }

type CardIssNum struct{ message.StringValue }

func (f CardIssNum) Tag() fix.Tag { return tag.CardIssNum }

type CardNumber struct{ message.StringValue }

func (f CardNumber) Tag() fix.Tag { return tag.CardNumber }

type CardStartDate struct{ message.LocalMktDateValue }

func (f CardStartDate) Tag() fix.Tag { return tag.CardStartDate }

type CashDistribAgentAcctName struct{ message.StringValue }

func (f CashDistribAgentAcctName) Tag() fix.Tag { return tag.CashDistribAgentAcctName }

type CashDistribAgentAcctNumber struct{ message.StringValue }

func (f CashDistribAgentAcctNumber) Tag() fix.Tag { return tag.CashDistribAgentAcctNumber }

type CashDistribAgentCode struct{ message.StringValue }

func (f CashDistribAgentCode) Tag() fix.Tag { return tag.CashDistribAgentCode }

type CashDistribAgentName struct{ message.StringValue }

func (f CashDistribAgentName) Tag() fix.Tag { return tag.CashDistribAgentName }

type CashDistribCurr struct{ message.CurrencyValue }

func (f CashDistribCurr) Tag() fix.Tag { return tag.CashDistribCurr }

type CashDistribPayRef struct{ message.StringValue }

func (f CashDistribPayRef) Tag() fix.Tag { return tag.CashDistribPayRef }

type CashMargin struct{ message.CharValue }

func (f CashMargin) Tag() fix.Tag { return tag.CashMargin }

type CashOrderQty struct{ message.QtyValue }

func (f CashOrderQty) Tag() fix.Tag { return tag.CashOrderQty }

type CashOutstanding struct{ message.AmtValue }

func (f CashOutstanding) Tag() fix.Tag { return tag.CashOutstanding }

type CashSettlAgentAcctName struct{ message.StringValue }

func (f CashSettlAgentAcctName) Tag() fix.Tag { return tag.CashSettlAgentAcctName }

type CashSettlAgentAcctNum struct{ message.StringValue }

func (f CashSettlAgentAcctNum) Tag() fix.Tag { return tag.CashSettlAgentAcctNum }

type CashSettlAgentCode struct{ message.StringValue }

func (f CashSettlAgentCode) Tag() fix.Tag { return tag.CashSettlAgentCode }

type CashSettlAgentContactName struct{ message.StringValue }

func (f CashSettlAgentContactName) Tag() fix.Tag { return tag.CashSettlAgentContactName }

type CashSettlAgentContactPhone struct{ message.StringValue }

func (f CashSettlAgentContactPhone) Tag() fix.Tag { return tag.CashSettlAgentContactPhone }

type CashSettlAgentName struct{ message.StringValue }

func (f CashSettlAgentName) Tag() fix.Tag { return tag.CashSettlAgentName }

type CcyAmt struct{ message.AmtValue }

func (f CcyAmt) Tag() fix.Tag { return tag.CcyAmt }

type CheckSum struct{ message.StringValue }

func (f CheckSum) Tag() fix.Tag { return tag.CheckSum }

type ClOrdID struct{ message.StringValue }

func (f ClOrdID) Tag() fix.Tag { return tag.ClOrdID }

type ClOrdLinkID struct{ message.StringValue }

func (f ClOrdLinkID) Tag() fix.Tag { return tag.ClOrdLinkID }

type ClearingAccount struct{ message.StringValue }

func (f ClearingAccount) Tag() fix.Tag { return tag.ClearingAccount }

type ClearingBusinessDate struct{ message.LocalMktDateValue }

func (f ClearingBusinessDate) Tag() fix.Tag { return tag.ClearingBusinessDate }

type ClearingFeeIndicator struct{ message.StringValue }

func (f ClearingFeeIndicator) Tag() fix.Tag { return tag.ClearingFeeIndicator }

type ClearingFirm struct{ message.StringValue }

func (f ClearingFirm) Tag() fix.Tag { return tag.ClearingFirm }

type ClearingInstruction struct{ message.IntValue }

func (f ClearingInstruction) Tag() fix.Tag { return tag.ClearingInstruction }

type ClientBidID struct{ message.StringValue }

func (f ClientBidID) Tag() fix.Tag { return tag.ClientBidID }

type ClientID struct{ message.StringValue }

func (f ClientID) Tag() fix.Tag { return tag.ClientID }

type CollAction struct{ message.IntValue }

func (f CollAction) Tag() fix.Tag { return tag.CollAction }

type CollApplType struct{ message.IntValue }

func (f CollApplType) Tag() fix.Tag { return tag.CollApplType }

type CollAsgnID struct{ message.StringValue }

func (f CollAsgnID) Tag() fix.Tag { return tag.CollAsgnID }

type CollAsgnReason struct{ message.IntValue }

func (f CollAsgnReason) Tag() fix.Tag { return tag.CollAsgnReason }

type CollAsgnRefID struct{ message.StringValue }

func (f CollAsgnRefID) Tag() fix.Tag { return tag.CollAsgnRefID }

type CollAsgnRejectReason struct{ message.IntValue }

func (f CollAsgnRejectReason) Tag() fix.Tag { return tag.CollAsgnRejectReason }

type CollAsgnRespType struct{ message.IntValue }

func (f CollAsgnRespType) Tag() fix.Tag { return tag.CollAsgnRespType }

type CollAsgnTransType struct{ message.IntValue }

func (f CollAsgnTransType) Tag() fix.Tag { return tag.CollAsgnTransType }

type CollInquiryID struct{ message.StringValue }

func (f CollInquiryID) Tag() fix.Tag { return tag.CollInquiryID }

type CollInquiryQualifier struct{ message.IntValue }

func (f CollInquiryQualifier) Tag() fix.Tag { return tag.CollInquiryQualifier }

type CollInquiryResult struct{ message.IntValue }

func (f CollInquiryResult) Tag() fix.Tag { return tag.CollInquiryResult }

type CollInquiryStatus struct{ message.IntValue }

func (f CollInquiryStatus) Tag() fix.Tag { return tag.CollInquiryStatus }

type CollReqID struct{ message.StringValue }

func (f CollReqID) Tag() fix.Tag { return tag.CollReqID }

type CollRespID struct{ message.StringValue }

func (f CollRespID) Tag() fix.Tag { return tag.CollRespID }

type CollRptID struct{ message.StringValue }

func (f CollRptID) Tag() fix.Tag { return tag.CollRptID }

type CollStatus struct{ message.IntValue }

func (f CollStatus) Tag() fix.Tag { return tag.CollStatus }

type CommCurrency struct{ message.CurrencyValue }

func (f CommCurrency) Tag() fix.Tag { return tag.CommCurrency }

type CommType struct{ message.CharValue }

func (f CommType) Tag() fix.Tag { return tag.CommType }

type Commission struct{ message.AmtValue }

func (f Commission) Tag() fix.Tag { return tag.Commission }

type ComplexEventCondition struct{ message.IntValue }

func (f ComplexEventCondition) Tag() fix.Tag { return tag.ComplexEventCondition }

type ComplexEventEndDate struct{ message.UTCTimestampValue }

func (f ComplexEventEndDate) Tag() fix.Tag { return tag.ComplexEventEndDate }

type ComplexEventEndTime struct{ message.UTCTimeOnlyValue }

func (f ComplexEventEndTime) Tag() fix.Tag { return tag.ComplexEventEndTime }

type ComplexEventPrice struct{ message.PriceValue }

func (f ComplexEventPrice) Tag() fix.Tag { return tag.ComplexEventPrice }

type ComplexEventPriceBoundaryMethod struct{ message.IntValue }

func (f ComplexEventPriceBoundaryMethod) Tag() fix.Tag { return tag.ComplexEventPriceBoundaryMethod }

type ComplexEventPriceBoundaryPrecision struct{ message.PercentageValue }

func (f ComplexEventPriceBoundaryPrecision) Tag() fix.Tag {
	return tag.ComplexEventPriceBoundaryPrecision
}

type ComplexEventPriceTimeType struct{ message.IntValue }

func (f ComplexEventPriceTimeType) Tag() fix.Tag { return tag.ComplexEventPriceTimeType }

type ComplexEventStartDate struct{ message.UTCTimestampValue }

func (f ComplexEventStartDate) Tag() fix.Tag { return tag.ComplexEventStartDate }

type ComplexEventStartTime struct{ message.UTCTimeOnlyValue }

func (f ComplexEventStartTime) Tag() fix.Tag { return tag.ComplexEventStartTime }

type ComplexEventType struct{ message.IntValue }

func (f ComplexEventType) Tag() fix.Tag { return tag.ComplexEventType }

type ComplexOptPayoutAmount struct{ message.AmtValue }

func (f ComplexOptPayoutAmount) Tag() fix.Tag { return tag.ComplexOptPayoutAmount }

type ComplianceID struct{ message.StringValue }

func (f ComplianceID) Tag() fix.Tag { return tag.ComplianceID }

type Concession struct{ message.AmtValue }

func (f Concession) Tag() fix.Tag { return tag.Concession }

type ConfirmID struct{ message.StringValue }

func (f ConfirmID) Tag() fix.Tag { return tag.ConfirmID }

type ConfirmRefID struct{ message.StringValue }

func (f ConfirmRefID) Tag() fix.Tag { return tag.ConfirmRefID }

type ConfirmRejReason struct{ message.IntValue }

func (f ConfirmRejReason) Tag() fix.Tag { return tag.ConfirmRejReason }

type ConfirmReqID struct{ message.StringValue }

func (f ConfirmReqID) Tag() fix.Tag { return tag.ConfirmReqID }

type ConfirmStatus struct{ message.IntValue }

func (f ConfirmStatus) Tag() fix.Tag { return tag.ConfirmStatus }

type ConfirmTransType struct{ message.IntValue }

func (f ConfirmTransType) Tag() fix.Tag { return tag.ConfirmTransType }

type ConfirmType struct{ message.IntValue }

func (f ConfirmType) Tag() fix.Tag { return tag.ConfirmType }

type ContAmtCurr struct{ message.CurrencyValue }

func (f ContAmtCurr) Tag() fix.Tag { return tag.ContAmtCurr }

type ContAmtType struct{ message.IntValue }

func (f ContAmtType) Tag() fix.Tag { return tag.ContAmtType }

type ContAmtValue struct{ message.FloatValue }

func (f ContAmtValue) Tag() fix.Tag { return tag.ContAmtValue }

type ContIntRptID struct{ message.StringValue }

func (f ContIntRptID) Tag() fix.Tag { return tag.ContIntRptID }

type ContextPartyID struct{ message.StringValue }

func (f ContextPartyID) Tag() fix.Tag { return tag.ContextPartyID }

type ContextPartyIDSource struct{ message.CharValue }

func (f ContextPartyIDSource) Tag() fix.Tag { return tag.ContextPartyIDSource }

type ContextPartyRole struct{ message.IntValue }

func (f ContextPartyRole) Tag() fix.Tag { return tag.ContextPartyRole }

type ContextPartySubID struct{ message.StringValue }

func (f ContextPartySubID) Tag() fix.Tag { return tag.ContextPartySubID }

type ContextPartySubIDType struct{ message.IntValue }

func (f ContextPartySubIDType) Tag() fix.Tag { return tag.ContextPartySubIDType }

type ContingencyType struct{ message.IntValue }

func (f ContingencyType) Tag() fix.Tag { return tag.ContingencyType }

type ContraBroker struct{ message.StringValue }

func (f ContraBroker) Tag() fix.Tag { return tag.ContraBroker }

type ContraLegRefID struct{ message.StringValue }

func (f ContraLegRefID) Tag() fix.Tag { return tag.ContraLegRefID }

type ContraTradeQty struct{ message.QtyValue }

func (f ContraTradeQty) Tag() fix.Tag { return tag.ContraTradeQty }

type ContraTradeTime struct{ message.UTCTimestampValue }

func (f ContraTradeTime) Tag() fix.Tag { return tag.ContraTradeTime }

type ContraTrader struct{ message.StringValue }

func (f ContraTrader) Tag() fix.Tag { return tag.ContraTrader }

type ContractMultiplier struct{ message.FloatValue }

func (f ContractMultiplier) Tag() fix.Tag { return tag.ContractMultiplier }

type ContractMultiplierUnit struct{ message.IntValue }

func (f ContractMultiplierUnit) Tag() fix.Tag { return tag.ContractMultiplierUnit }

type ContractSettlMonth struct{ message.MonthYearValue }

func (f ContractSettlMonth) Tag() fix.Tag { return tag.ContractSettlMonth }

type ContraryInstructionIndicator struct{ message.BooleanValue }

func (f ContraryInstructionIndicator) Tag() fix.Tag { return tag.ContraryInstructionIndicator }

type CopyMsgIndicator struct{ message.BooleanValue }

func (f CopyMsgIndicator) Tag() fix.Tag { return tag.CopyMsgIndicator }

type CorporateAction struct{ message.MultipleCharValue }

func (f CorporateAction) Tag() fix.Tag { return tag.CorporateAction }

type Country struct{ message.CountryValue }

func (f Country) Tag() fix.Tag { return tag.Country }

type CountryOfIssue struct{ message.CountryValue }

func (f CountryOfIssue) Tag() fix.Tag { return tag.CountryOfIssue }

type CouponPaymentDate struct{ message.LocalMktDateValue }

func (f CouponPaymentDate) Tag() fix.Tag { return tag.CouponPaymentDate }

type CouponRate struct{ message.PercentageValue }

func (f CouponRate) Tag() fix.Tag { return tag.CouponRate }

type CoveredOrUncovered struct{ message.IntValue }

func (f CoveredOrUncovered) Tag() fix.Tag { return tag.CoveredOrUncovered }

type CreditRating struct{ message.StringValue }

func (f CreditRating) Tag() fix.Tag { return tag.CreditRating }

type CrossID struct{ message.StringValue }

func (f CrossID) Tag() fix.Tag { return tag.CrossID }

type CrossPercent struct{ message.PercentageValue }

func (f CrossPercent) Tag() fix.Tag { return tag.CrossPercent }

type CrossPrioritization struct{ message.IntValue }

func (f CrossPrioritization) Tag() fix.Tag { return tag.CrossPrioritization }

type CrossType struct{ message.IntValue }

func (f CrossType) Tag() fix.Tag { return tag.CrossType }

type CstmApplVerID struct{ message.StringValue }

func (f CstmApplVerID) Tag() fix.Tag { return tag.CstmApplVerID }

type CumQty struct{ message.QtyValue }

func (f CumQty) Tag() fix.Tag { return tag.CumQty }

type Currency struct{ message.CurrencyValue }

func (f Currency) Tag() fix.Tag { return tag.Currency }

type CurrencyRatio struct{ message.FloatValue }

func (f CurrencyRatio) Tag() fix.Tag { return tag.CurrencyRatio }

type CustDirectedOrder struct{ message.BooleanValue }

func (f CustDirectedOrder) Tag() fix.Tag { return tag.CustDirectedOrder }

type CustOrderCapacity struct{ message.IntValue }

func (f CustOrderCapacity) Tag() fix.Tag { return tag.CustOrderCapacity }

type CustOrderHandlingInst struct{ message.MultipleStringValue }

func (f CustOrderHandlingInst) Tag() fix.Tag { return tag.CustOrderHandlingInst }

type CustomerOrFirm struct{ message.IntValue }

func (f CustomerOrFirm) Tag() fix.Tag { return tag.CustomerOrFirm }

type CxlQty struct{ message.QtyValue }

func (f CxlQty) Tag() fix.Tag { return tag.CxlQty }

type CxlRejReason struct{ message.IntValue }

func (f CxlRejReason) Tag() fix.Tag { return tag.CxlRejReason }

type CxlRejResponseTo struct{ message.CharValue }

func (f CxlRejResponseTo) Tag() fix.Tag { return tag.CxlRejResponseTo }

type CxlType struct{ message.CharValue }

func (f CxlType) Tag() fix.Tag { return tag.CxlType }

type DKReason struct{ message.CharValue }

func (f DKReason) Tag() fix.Tag { return tag.DKReason }

type DateOfBirth struct{ message.LocalMktDateValue }

func (f DateOfBirth) Tag() fix.Tag { return tag.DateOfBirth }

type DatedDate struct{ message.LocalMktDateValue }

func (f DatedDate) Tag() fix.Tag { return tag.DatedDate }

type DayAvgPx struct{ message.PriceValue }

func (f DayAvgPx) Tag() fix.Tag { return tag.DayAvgPx }

type DayBookingInst struct{ message.CharValue }

func (f DayBookingInst) Tag() fix.Tag { return tag.DayBookingInst }

type DayCumQty struct{ message.QtyValue }

func (f DayCumQty) Tag() fix.Tag { return tag.DayCumQty }

type DayOrderQty struct{ message.QtyValue }

func (f DayOrderQty) Tag() fix.Tag { return tag.DayOrderQty }

type DealingCapacity struct{ message.CharValue }

func (f DealingCapacity) Tag() fix.Tag { return tag.DealingCapacity }

type DefBidSize struct{ message.QtyValue }

func (f DefBidSize) Tag() fix.Tag { return tag.DefBidSize }

type DefOfferSize struct{ message.QtyValue }

func (f DefOfferSize) Tag() fix.Tag { return tag.DefOfferSize }

type DefaultApplExtID struct{ message.IntValue }

func (f DefaultApplExtID) Tag() fix.Tag { return tag.DefaultApplExtID }

type DefaultApplVerID struct{ message.StringValue }

func (f DefaultApplVerID) Tag() fix.Tag { return tag.DefaultApplVerID }

type DefaultCstmApplVerID struct{ message.StringValue }

func (f DefaultCstmApplVerID) Tag() fix.Tag { return tag.DefaultCstmApplVerID }

type DefaultVerIndicator struct{ message.BooleanValue }

func (f DefaultVerIndicator) Tag() fix.Tag { return tag.DefaultVerIndicator }

type DeleteReason struct{ message.CharValue }

func (f DeleteReason) Tag() fix.Tag { return tag.DeleteReason }

type DeliverToCompID struct{ message.StringValue }

func (f DeliverToCompID) Tag() fix.Tag { return tag.DeliverToCompID }

type DeliverToLocationID struct{ message.StringValue }

func (f DeliverToLocationID) Tag() fix.Tag { return tag.DeliverToLocationID }

type DeliverToSubID struct{ message.StringValue }

func (f DeliverToSubID) Tag() fix.Tag { return tag.DeliverToSubID }

type DeliveryDate struct{ message.LocalMktDateValue }

func (f DeliveryDate) Tag() fix.Tag { return tag.DeliveryDate }

type DeliveryForm struct{ message.IntValue }

func (f DeliveryForm) Tag() fix.Tag { return tag.DeliveryForm }

type DeliveryType struct{ message.IntValue }

func (f DeliveryType) Tag() fix.Tag { return tag.DeliveryType }

type DerivFlexProductEligibilityIndicator struct{ message.BooleanValue }

func (f DerivFlexProductEligibilityIndicator) Tag() fix.Tag {
	return tag.DerivFlexProductEligibilityIndicator
}

type DerivativeCFICode struct{ message.StringValue }

func (f DerivativeCFICode) Tag() fix.Tag { return tag.DerivativeCFICode }

type DerivativeCapPrice struct{ message.PriceValue }

func (f DerivativeCapPrice) Tag() fix.Tag { return tag.DerivativeCapPrice }

type DerivativeContractMultiplier struct{ message.FloatValue }

func (f DerivativeContractMultiplier) Tag() fix.Tag { return tag.DerivativeContractMultiplier }

type DerivativeContractMultiplierUnit struct{ message.IntValue }

func (f DerivativeContractMultiplierUnit) Tag() fix.Tag { return tag.DerivativeContractMultiplierUnit }

type DerivativeContractSettlMonth struct{ message.MonthYearValue }

func (f DerivativeContractSettlMonth) Tag() fix.Tag { return tag.DerivativeContractSettlMonth }

type DerivativeCountryOfIssue struct{ message.CountryValue }

func (f DerivativeCountryOfIssue) Tag() fix.Tag { return tag.DerivativeCountryOfIssue }

type DerivativeEncodedIssuer struct{ message.DataValue }

func (f DerivativeEncodedIssuer) Tag() fix.Tag { return tag.DerivativeEncodedIssuer }

type DerivativeEncodedIssuerLen struct{ message.LengthValue }

func (f DerivativeEncodedIssuerLen) Tag() fix.Tag { return tag.DerivativeEncodedIssuerLen }

type DerivativeEncodedSecurityDesc struct{ message.DataValue }

func (f DerivativeEncodedSecurityDesc) Tag() fix.Tag { return tag.DerivativeEncodedSecurityDesc }

type DerivativeEncodedSecurityDescLen struct{ message.LengthValue }

func (f DerivativeEncodedSecurityDescLen) Tag() fix.Tag { return tag.DerivativeEncodedSecurityDescLen }

type DerivativeEventDate struct{ message.LocalMktDateValue }

func (f DerivativeEventDate) Tag() fix.Tag { return tag.DerivativeEventDate }

type DerivativeEventPx struct{ message.PriceValue }

func (f DerivativeEventPx) Tag() fix.Tag { return tag.DerivativeEventPx }

type DerivativeEventText struct{ message.StringValue }

func (f DerivativeEventText) Tag() fix.Tag { return tag.DerivativeEventText }

type DerivativeEventTime struct{ message.UTCTimestampValue }

func (f DerivativeEventTime) Tag() fix.Tag { return tag.DerivativeEventTime }

type DerivativeEventType struct{ message.IntValue }

func (f DerivativeEventType) Tag() fix.Tag { return tag.DerivativeEventType }

type DerivativeExerciseStyle struct{ message.CharValue }

func (f DerivativeExerciseStyle) Tag() fix.Tag { return tag.DerivativeExerciseStyle }

type DerivativeFloorPrice struct{ message.PriceValue }

func (f DerivativeFloorPrice) Tag() fix.Tag { return tag.DerivativeFloorPrice }

type DerivativeFlowScheduleType struct{ message.IntValue }

func (f DerivativeFlowScheduleType) Tag() fix.Tag { return tag.DerivativeFlowScheduleType }

type DerivativeFuturesValuationMethod struct{ message.StringValue }

func (f DerivativeFuturesValuationMethod) Tag() fix.Tag { return tag.DerivativeFuturesValuationMethod }

type DerivativeInstrAttribType struct{ message.IntValue }

func (f DerivativeInstrAttribType) Tag() fix.Tag { return tag.DerivativeInstrAttribType }

type DerivativeInstrAttribValue struct{ message.StringValue }

func (f DerivativeInstrAttribValue) Tag() fix.Tag { return tag.DerivativeInstrAttribValue }

type DerivativeInstrRegistry struct{ message.StringValue }

func (f DerivativeInstrRegistry) Tag() fix.Tag { return tag.DerivativeInstrRegistry }

type DerivativeInstrmtAssignmentMethod struct{ message.CharValue }

func (f DerivativeInstrmtAssignmentMethod) Tag() fix.Tag { return tag.DerivativeInstrmtAssignmentMethod }

type DerivativeInstrumentPartyID struct{ message.StringValue }

func (f DerivativeInstrumentPartyID) Tag() fix.Tag { return tag.DerivativeInstrumentPartyID }

type DerivativeInstrumentPartyIDSource struct{ message.StringValue }

func (f DerivativeInstrumentPartyIDSource) Tag() fix.Tag { return tag.DerivativeInstrumentPartyIDSource }

type DerivativeInstrumentPartyRole struct{ message.IntValue }

func (f DerivativeInstrumentPartyRole) Tag() fix.Tag { return tag.DerivativeInstrumentPartyRole }

type DerivativeInstrumentPartySubID struct{ message.StringValue }

func (f DerivativeInstrumentPartySubID) Tag() fix.Tag { return tag.DerivativeInstrumentPartySubID }

type DerivativeInstrumentPartySubIDType struct{ message.IntValue }

func (f DerivativeInstrumentPartySubIDType) Tag() fix.Tag {
	return tag.DerivativeInstrumentPartySubIDType
}

type DerivativeIssueDate struct{ message.LocalMktDateValue }

func (f DerivativeIssueDate) Tag() fix.Tag { return tag.DerivativeIssueDate }

type DerivativeIssuer struct{ message.StringValue }

func (f DerivativeIssuer) Tag() fix.Tag { return tag.DerivativeIssuer }

type DerivativeListMethod struct{ message.IntValue }

func (f DerivativeListMethod) Tag() fix.Tag { return tag.DerivativeListMethod }

type DerivativeLocaleOfIssue struct{ message.StringValue }

func (f DerivativeLocaleOfIssue) Tag() fix.Tag { return tag.DerivativeLocaleOfIssue }

type DerivativeMaturityDate struct{ message.LocalMktDateValue }

func (f DerivativeMaturityDate) Tag() fix.Tag { return tag.DerivativeMaturityDate }

type DerivativeMaturityMonthYear struct{ message.MonthYearValue }

func (f DerivativeMaturityMonthYear) Tag() fix.Tag { return tag.DerivativeMaturityMonthYear }

type DerivativeMaturityTime struct{ message.TZTimeOnlyValue }

func (f DerivativeMaturityTime) Tag() fix.Tag { return tag.DerivativeMaturityTime }

type DerivativeMinPriceIncrement struct{ message.FloatValue }

func (f DerivativeMinPriceIncrement) Tag() fix.Tag { return tag.DerivativeMinPriceIncrement }

type DerivativeMinPriceIncrementAmount struct{ message.AmtValue }

func (f DerivativeMinPriceIncrementAmount) Tag() fix.Tag { return tag.DerivativeMinPriceIncrementAmount }

type DerivativeNTPositionLimit struct{ message.IntValue }

func (f DerivativeNTPositionLimit) Tag() fix.Tag { return tag.DerivativeNTPositionLimit }

type DerivativeOptAttribute struct{ message.CharValue }

func (f DerivativeOptAttribute) Tag() fix.Tag { return tag.DerivativeOptAttribute }

type DerivativeOptPayAmount struct{ message.AmtValue }

func (f DerivativeOptPayAmount) Tag() fix.Tag { return tag.DerivativeOptPayAmount }

type DerivativePositionLimit struct{ message.IntValue }

func (f DerivativePositionLimit) Tag() fix.Tag { return tag.DerivativePositionLimit }

type DerivativePriceQuoteMethod struct{ message.StringValue }

func (f DerivativePriceQuoteMethod) Tag() fix.Tag { return tag.DerivativePriceQuoteMethod }

type DerivativePriceUnitOfMeasure struct{ message.StringValue }

func (f DerivativePriceUnitOfMeasure) Tag() fix.Tag { return tag.DerivativePriceUnitOfMeasure }

type DerivativePriceUnitOfMeasureQty struct{ message.QtyValue }

func (f DerivativePriceUnitOfMeasureQty) Tag() fix.Tag { return tag.DerivativePriceUnitOfMeasureQty }

type DerivativeProduct struct{ message.IntValue }

func (f DerivativeProduct) Tag() fix.Tag { return tag.DerivativeProduct }

type DerivativeProductComplex struct{ message.StringValue }

func (f DerivativeProductComplex) Tag() fix.Tag { return tag.DerivativeProductComplex }

type DerivativePutOrCall struct{ message.IntValue }

func (f DerivativePutOrCall) Tag() fix.Tag { return tag.DerivativePutOrCall }

type DerivativeSecurityAltID struct{ message.StringValue }

func (f DerivativeSecurityAltID) Tag() fix.Tag { return tag.DerivativeSecurityAltID }

type DerivativeSecurityAltIDSource struct{ message.StringValue }

func (f DerivativeSecurityAltIDSource) Tag() fix.Tag { return tag.DerivativeSecurityAltIDSource }

type DerivativeSecurityDesc struct{ message.StringValue }

func (f DerivativeSecurityDesc) Tag() fix.Tag { return tag.DerivativeSecurityDesc }

type DerivativeSecurityExchange struct{ message.ExchangeValue }

func (f DerivativeSecurityExchange) Tag() fix.Tag { return tag.DerivativeSecurityExchange }

type DerivativeSecurityGroup struct{ message.StringValue }

func (f DerivativeSecurityGroup) Tag() fix.Tag { return tag.DerivativeSecurityGroup }

type DerivativeSecurityID struct{ message.StringValue }

func (f DerivativeSecurityID) Tag() fix.Tag { return tag.DerivativeSecurityID }

type DerivativeSecurityIDSource struct{ message.StringValue }

func (f DerivativeSecurityIDSource) Tag() fix.Tag { return tag.DerivativeSecurityIDSource }

type DerivativeSecurityListRequestType struct{ message.IntValue }

func (f DerivativeSecurityListRequestType) Tag() fix.Tag { return tag.DerivativeSecurityListRequestType }

type DerivativeSecurityStatus struct{ message.StringValue }

func (f DerivativeSecurityStatus) Tag() fix.Tag { return tag.DerivativeSecurityStatus }

type DerivativeSecuritySubType struct{ message.StringValue }

func (f DerivativeSecuritySubType) Tag() fix.Tag { return tag.DerivativeSecuritySubType }

type DerivativeSecurityType struct{ message.StringValue }

func (f DerivativeSecurityType) Tag() fix.Tag { return tag.DerivativeSecurityType }

type DerivativeSecurityXML struct{ message.DataValue }

func (f DerivativeSecurityXML) Tag() fix.Tag { return tag.DerivativeSecurityXML }

type DerivativeSecurityXMLLen struct{ message.LengthValue }

func (f DerivativeSecurityXMLLen) Tag() fix.Tag { return tag.DerivativeSecurityXMLLen }

type DerivativeSecurityXMLSchema struct{ message.StringValue }

func (f DerivativeSecurityXMLSchema) Tag() fix.Tag { return tag.DerivativeSecurityXMLSchema }

type DerivativeSettlMethod struct{ message.CharValue }

func (f DerivativeSettlMethod) Tag() fix.Tag { return tag.DerivativeSettlMethod }

type DerivativeSettleOnOpenFlag struct{ message.StringValue }

func (f DerivativeSettleOnOpenFlag) Tag() fix.Tag { return tag.DerivativeSettleOnOpenFlag }

type DerivativeStateOrProvinceOfIssue struct{ message.StringValue }

func (f DerivativeStateOrProvinceOfIssue) Tag() fix.Tag { return tag.DerivativeStateOrProvinceOfIssue }

type DerivativeStrikeCurrency struct{ message.CurrencyValue }

func (f DerivativeStrikeCurrency) Tag() fix.Tag { return tag.DerivativeStrikeCurrency }

type DerivativeStrikeMultiplier struct{ message.FloatValue }

func (f DerivativeStrikeMultiplier) Tag() fix.Tag { return tag.DerivativeStrikeMultiplier }

type DerivativeStrikePrice struct{ message.PriceValue }

func (f DerivativeStrikePrice) Tag() fix.Tag { return tag.DerivativeStrikePrice }

type DerivativeStrikeValue struct{ message.FloatValue }

func (f DerivativeStrikeValue) Tag() fix.Tag { return tag.DerivativeStrikeValue }

type DerivativeSymbol struct{ message.StringValue }

func (f DerivativeSymbol) Tag() fix.Tag { return tag.DerivativeSymbol }

type DerivativeSymbolSfx struct{ message.StringValue }

func (f DerivativeSymbolSfx) Tag() fix.Tag { return tag.DerivativeSymbolSfx }

type DerivativeTimeUnit struct{ message.StringValue }

func (f DerivativeTimeUnit) Tag() fix.Tag { return tag.DerivativeTimeUnit }

type DerivativeUnitOfMeasure struct{ message.StringValue }

func (f DerivativeUnitOfMeasure) Tag() fix.Tag { return tag.DerivativeUnitOfMeasure }

type DerivativeUnitOfMeasureQty struct{ message.QtyValue }

func (f DerivativeUnitOfMeasureQty) Tag() fix.Tag { return tag.DerivativeUnitOfMeasureQty }

type DerivativeValuationMethod struct{ message.StringValue }

func (f DerivativeValuationMethod) Tag() fix.Tag { return tag.DerivativeValuationMethod }

type Designation struct{ message.StringValue }

func (f Designation) Tag() fix.Tag { return tag.Designation }

type DeskID struct{ message.StringValue }

func (f DeskID) Tag() fix.Tag { return tag.DeskID }

type DeskOrderHandlingInst struct{ message.MultipleStringValue }

func (f DeskOrderHandlingInst) Tag() fix.Tag { return tag.DeskOrderHandlingInst }

type DeskType struct{ message.StringValue }

func (f DeskType) Tag() fix.Tag { return tag.DeskType }

type DeskTypeSource struct{ message.IntValue }

func (f DeskTypeSource) Tag() fix.Tag { return tag.DeskTypeSource }

type DetachmentPoint struct{ message.PercentageValue }

func (f DetachmentPoint) Tag() fix.Tag { return tag.DetachmentPoint }

type DiscretionInst struct{ message.CharValue }

func (f DiscretionInst) Tag() fix.Tag { return tag.DiscretionInst }

type DiscretionLimitType struct{ message.IntValue }

func (f DiscretionLimitType) Tag() fix.Tag { return tag.DiscretionLimitType }

type DiscretionMoveType struct{ message.IntValue }

func (f DiscretionMoveType) Tag() fix.Tag { return tag.DiscretionMoveType }

type DiscretionOffset struct{ message.PriceOffsetValue }

func (f DiscretionOffset) Tag() fix.Tag { return tag.DiscretionOffset }

type DiscretionOffsetType struct{ message.IntValue }

func (f DiscretionOffsetType) Tag() fix.Tag { return tag.DiscretionOffsetType }

type DiscretionOffsetValue struct{ message.FloatValue }

func (f DiscretionOffsetValue) Tag() fix.Tag { return tag.DiscretionOffsetValue }

type DiscretionPrice struct{ message.PriceValue }

func (f DiscretionPrice) Tag() fix.Tag { return tag.DiscretionPrice }

type DiscretionRoundDirection struct{ message.IntValue }

func (f DiscretionRoundDirection) Tag() fix.Tag { return tag.DiscretionRoundDirection }

type DiscretionScope struct{ message.IntValue }

func (f DiscretionScope) Tag() fix.Tag { return tag.DiscretionScope }

type DisplayHighQty struct{ message.QtyValue }

func (f DisplayHighQty) Tag() fix.Tag { return tag.DisplayHighQty }

type DisplayLowQty struct{ message.QtyValue }

func (f DisplayLowQty) Tag() fix.Tag { return tag.DisplayLowQty }

type DisplayMethod struct{ message.CharValue }

func (f DisplayMethod) Tag() fix.Tag { return tag.DisplayMethod }

type DisplayMinIncr struct{ message.QtyValue }

func (f DisplayMinIncr) Tag() fix.Tag { return tag.DisplayMinIncr }

type DisplayQty struct{ message.QtyValue }

func (f DisplayQty) Tag() fix.Tag { return tag.DisplayQty }

type DisplayWhen struct{ message.CharValue }

func (f DisplayWhen) Tag() fix.Tag { return tag.DisplayWhen }

type DistribPaymentMethod struct{ message.IntValue }

func (f DistribPaymentMethod) Tag() fix.Tag { return tag.DistribPaymentMethod }

type DistribPercentage struct{ message.PercentageValue }

func (f DistribPercentage) Tag() fix.Tag { return tag.DistribPercentage }

type DividendYield struct{ message.PercentageValue }

func (f DividendYield) Tag() fix.Tag { return tag.DividendYield }

type DlvyInst struct{ message.StringValue }

func (f DlvyInst) Tag() fix.Tag { return tag.DlvyInst }

type DlvyInstType struct{ message.CharValue }

func (f DlvyInstType) Tag() fix.Tag { return tag.DlvyInstType }

type DueToRelated struct{ message.BooleanValue }

func (f DueToRelated) Tag() fix.Tag { return tag.DueToRelated }

type EFPTrackingError struct{ message.PercentageValue }

func (f EFPTrackingError) Tag() fix.Tag { return tag.EFPTrackingError }

type EffectiveTime struct{ message.UTCTimestampValue }

func (f EffectiveTime) Tag() fix.Tag { return tag.EffectiveTime }

type EmailThreadID struct{ message.StringValue }

func (f EmailThreadID) Tag() fix.Tag { return tag.EmailThreadID }

type EmailType struct{ message.CharValue }

func (f EmailType) Tag() fix.Tag { return tag.EmailType }

type EncodedAllocText struct{ message.DataValue }

func (f EncodedAllocText) Tag() fix.Tag { return tag.EncodedAllocText }

type EncodedAllocTextLen struct{ message.LengthValue }

func (f EncodedAllocTextLen) Tag() fix.Tag { return tag.EncodedAllocTextLen }

type EncodedHeadline struct{ message.DataValue }

func (f EncodedHeadline) Tag() fix.Tag { return tag.EncodedHeadline }

type EncodedHeadlineLen struct{ message.LengthValue }

func (f EncodedHeadlineLen) Tag() fix.Tag { return tag.EncodedHeadlineLen }

type EncodedIssuer struct{ message.DataValue }

func (f EncodedIssuer) Tag() fix.Tag { return tag.EncodedIssuer }

type EncodedIssuerLen struct{ message.LengthValue }

func (f EncodedIssuerLen) Tag() fix.Tag { return tag.EncodedIssuerLen }

type EncodedLegIssuer struct{ message.DataValue }

func (f EncodedLegIssuer) Tag() fix.Tag { return tag.EncodedLegIssuer }

type EncodedLegIssuerLen struct{ message.LengthValue }

func (f EncodedLegIssuerLen) Tag() fix.Tag { return tag.EncodedLegIssuerLen }

type EncodedLegSecurityDesc struct{ message.DataValue }

func (f EncodedLegSecurityDesc) Tag() fix.Tag { return tag.EncodedLegSecurityDesc }

type EncodedLegSecurityDescLen struct{ message.LengthValue }

func (f EncodedLegSecurityDescLen) Tag() fix.Tag { return tag.EncodedLegSecurityDescLen }

type EncodedListExecInst struct{ message.DataValue }

func (f EncodedListExecInst) Tag() fix.Tag { return tag.EncodedListExecInst }

type EncodedListExecInstLen struct{ message.LengthValue }

func (f EncodedListExecInstLen) Tag() fix.Tag { return tag.EncodedListExecInstLen }

type EncodedListStatusText struct{ message.DataValue }

func (f EncodedListStatusText) Tag() fix.Tag { return tag.EncodedListStatusText }

type EncodedListStatusTextLen struct{ message.LengthValue }

func (f EncodedListStatusTextLen) Tag() fix.Tag { return tag.EncodedListStatusTextLen }

type EncodedMktSegmDesc struct{ message.DataValue }

func (f EncodedMktSegmDesc) Tag() fix.Tag { return tag.EncodedMktSegmDesc }

type EncodedMktSegmDescLen struct{ message.LengthValue }

func (f EncodedMktSegmDescLen) Tag() fix.Tag { return tag.EncodedMktSegmDescLen }

type EncodedSecurityDesc struct{ message.DataValue }

func (f EncodedSecurityDesc) Tag() fix.Tag { return tag.EncodedSecurityDesc }

type EncodedSecurityDescLen struct{ message.LengthValue }

func (f EncodedSecurityDescLen) Tag() fix.Tag { return tag.EncodedSecurityDescLen }

type EncodedSecurityListDesc struct{ message.DataValue }

func (f EncodedSecurityListDesc) Tag() fix.Tag { return tag.EncodedSecurityListDesc }

type EncodedSecurityListDescLen struct{ message.LengthValue }

func (f EncodedSecurityListDescLen) Tag() fix.Tag { return tag.EncodedSecurityListDescLen }

type EncodedSubject struct{ message.DataValue }

func (f EncodedSubject) Tag() fix.Tag { return tag.EncodedSubject }

type EncodedSubjectLen struct{ message.LengthValue }

func (f EncodedSubjectLen) Tag() fix.Tag { return tag.EncodedSubjectLen }

type EncodedSymbol struct{ message.DataValue }

func (f EncodedSymbol) Tag() fix.Tag { return tag.EncodedSymbol }

type EncodedSymbolLen struct{ message.LengthValue }

func (f EncodedSymbolLen) Tag() fix.Tag { return tag.EncodedSymbolLen }

type EncodedText struct{ message.DataValue }

func (f EncodedText) Tag() fix.Tag { return tag.EncodedText }

type EncodedTextLen struct{ message.LengthValue }

func (f EncodedTextLen) Tag() fix.Tag { return tag.EncodedTextLen }

type EncodedUnderlyingIssuer struct{ message.DataValue }

func (f EncodedUnderlyingIssuer) Tag() fix.Tag { return tag.EncodedUnderlyingIssuer }

type EncodedUnderlyingIssuerLen struct{ message.LengthValue }

func (f EncodedUnderlyingIssuerLen) Tag() fix.Tag { return tag.EncodedUnderlyingIssuerLen }

type EncodedUnderlyingSecurityDesc struct{ message.DataValue }

func (f EncodedUnderlyingSecurityDesc) Tag() fix.Tag { return tag.EncodedUnderlyingSecurityDesc }

type EncodedUnderlyingSecurityDescLen struct{ message.LengthValue }

func (f EncodedUnderlyingSecurityDescLen) Tag() fix.Tag { return tag.EncodedUnderlyingSecurityDescLen }

type EncryptMethod struct{ message.IntValue }

func (f EncryptMethod) Tag() fix.Tag { return tag.EncryptMethod }

type EncryptedNewPassword struct{ message.DataValue }

func (f EncryptedNewPassword) Tag() fix.Tag { return tag.EncryptedNewPassword }

type EncryptedNewPasswordLen struct{ message.LengthValue }

func (f EncryptedNewPasswordLen) Tag() fix.Tag { return tag.EncryptedNewPasswordLen }

type EncryptedPassword struct{ message.DataValue }

func (f EncryptedPassword) Tag() fix.Tag { return tag.EncryptedPassword }

type EncryptedPasswordLen struct{ message.LengthValue }

func (f EncryptedPasswordLen) Tag() fix.Tag { return tag.EncryptedPasswordLen }

type EncryptedPasswordMethod struct{ message.IntValue }

func (f EncryptedPasswordMethod) Tag() fix.Tag { return tag.EncryptedPasswordMethod }

type EndAccruedInterestAmt struct{ message.AmtValue }

func (f EndAccruedInterestAmt) Tag() fix.Tag { return tag.EndAccruedInterestAmt }

type EndCash struct{ message.AmtValue }

func (f EndCash) Tag() fix.Tag { return tag.EndCash }

type EndDate struct{ message.LocalMktDateValue }

func (f EndDate) Tag() fix.Tag { return tag.EndDate }

type EndMaturityMonthYear struct{ message.MonthYearValue }

func (f EndMaturityMonthYear) Tag() fix.Tag { return tag.EndMaturityMonthYear }

type EndSeqNo struct{ message.SeqNumValue }

func (f EndSeqNo) Tag() fix.Tag { return tag.EndSeqNo }

type EndStrikePxRange struct{ message.PriceValue }

func (f EndStrikePxRange) Tag() fix.Tag { return tag.EndStrikePxRange }

type EndTickPriceRange struct{ message.PriceValue }

func (f EndTickPriceRange) Tag() fix.Tag { return tag.EndTickPriceRange }

type EventDate struct{ message.LocalMktDateValue }

func (f EventDate) Tag() fix.Tag { return tag.EventDate }

type EventPx struct{ message.PriceValue }

func (f EventPx) Tag() fix.Tag { return tag.EventPx }

type EventText struct{ message.StringValue }

func (f EventText) Tag() fix.Tag { return tag.EventText }

type EventTime struct{ message.UTCTimestampValue }

func (f EventTime) Tag() fix.Tag { return tag.EventTime }

type EventType struct{ message.IntValue }

func (f EventType) Tag() fix.Tag { return tag.EventType }

type ExDate struct{ message.LocalMktDateValue }

func (f ExDate) Tag() fix.Tag { return tag.ExDate }

type ExDestination struct{ message.ExchangeValue }

func (f ExDestination) Tag() fix.Tag { return tag.ExDestination }

type ExDestinationIDSource struct{ message.CharValue }

func (f ExDestinationIDSource) Tag() fix.Tag { return tag.ExDestinationIDSource }

type ExchangeForPhysical struct{ message.BooleanValue }

func (f ExchangeForPhysical) Tag() fix.Tag { return tag.ExchangeForPhysical }

type ExchangeRule struct{ message.StringValue }

func (f ExchangeRule) Tag() fix.Tag { return tag.ExchangeRule }

type ExchangeSpecialInstructions struct{ message.StringValue }

func (f ExchangeSpecialInstructions) Tag() fix.Tag { return tag.ExchangeSpecialInstructions }

type ExecAckStatus struct{ message.CharValue }

func (f ExecAckStatus) Tag() fix.Tag { return tag.ExecAckStatus }

type ExecBroker struct{ message.StringValue }

func (f ExecBroker) Tag() fix.Tag { return tag.ExecBroker }

type ExecID struct{ message.StringValue }

func (f ExecID) Tag() fix.Tag { return tag.ExecID }

type ExecInst struct{ message.MultipleCharValue }

func (f ExecInst) Tag() fix.Tag { return tag.ExecInst }

type ExecInstValue struct{ message.CharValue }

func (f ExecInstValue) Tag() fix.Tag { return tag.ExecInstValue }

type ExecPriceAdjustment struct{ message.FloatValue }

func (f ExecPriceAdjustment) Tag() fix.Tag { return tag.ExecPriceAdjustment }

type ExecPriceType struct{ message.CharValue }

func (f ExecPriceType) Tag() fix.Tag { return tag.ExecPriceType }

type ExecRefID struct{ message.StringValue }

func (f ExecRefID) Tag() fix.Tag { return tag.ExecRefID }

type ExecRestatementReason struct{ message.IntValue }

func (f ExecRestatementReason) Tag() fix.Tag { return tag.ExecRestatementReason }

type ExecTransType struct{ message.CharValue }

func (f ExecTransType) Tag() fix.Tag { return tag.ExecTransType }

type ExecType struct{ message.CharValue }

func (f ExecType) Tag() fix.Tag { return tag.ExecType }

type ExecValuationPoint struct{ message.UTCTimestampValue }

func (f ExecValuationPoint) Tag() fix.Tag { return tag.ExecValuationPoint }

type ExerciseMethod struct{ message.CharValue }

func (f ExerciseMethod) Tag() fix.Tag { return tag.ExerciseMethod }

type ExerciseStyle struct{ message.IntValue }

func (f ExerciseStyle) Tag() fix.Tag { return tag.ExerciseStyle }

type ExpQty struct{ message.QtyValue }

func (f ExpQty) Tag() fix.Tag { return tag.ExpQty }

type ExpType struct{ message.IntValue }

func (f ExpType) Tag() fix.Tag { return tag.ExpType }

type ExpirationCycle struct{ message.IntValue }

func (f ExpirationCycle) Tag() fix.Tag { return tag.ExpirationCycle }

type ExpirationQtyType struct{ message.IntValue }

func (f ExpirationQtyType) Tag() fix.Tag { return tag.ExpirationQtyType }

type ExpireDate struct{ message.LocalMktDateValue }

func (f ExpireDate) Tag() fix.Tag { return tag.ExpireDate }

type ExpireTime struct{ message.UTCTimestampValue }

func (f ExpireTime) Tag() fix.Tag { return tag.ExpireTime }

type Factor struct{ message.FloatValue }

func (f Factor) Tag() fix.Tag { return tag.Factor }

type FairValue struct{ message.AmtValue }

func (f FairValue) Tag() fix.Tag { return tag.FairValue }

type FeeMultiplier struct{ message.FloatValue }

func (f FeeMultiplier) Tag() fix.Tag { return tag.FeeMultiplier }

type FillExecID struct{ message.StringValue }

func (f FillExecID) Tag() fix.Tag { return tag.FillExecID }

type FillLiquidityInd struct{ message.IntValue }

func (f FillLiquidityInd) Tag() fix.Tag { return tag.FillLiquidityInd }

type FillPx struct{ message.PriceValue }

func (f FillPx) Tag() fix.Tag { return tag.FillPx }

type FillQty struct{ message.QtyValue }

func (f FillQty) Tag() fix.Tag { return tag.FillQty }

type FinancialStatus struct{ message.MultipleCharValue }

func (f FinancialStatus) Tag() fix.Tag { return tag.FinancialStatus }

type FirmTradeID struct{ message.StringValue }

func (f FirmTradeID) Tag() fix.Tag { return tag.FirmTradeID }

type FirstPx struct{ message.PriceValue }

func (f FirstPx) Tag() fix.Tag { return tag.FirstPx }

type FlexProductEligibilityIndicator struct{ message.BooleanValue }

func (f FlexProductEligibilityIndicator) Tag() fix.Tag { return tag.FlexProductEligibilityIndicator }

type FlexibleIndicator struct{ message.BooleanValue }

func (f FlexibleIndicator) Tag() fix.Tag { return tag.FlexibleIndicator }

type FloorPrice struct{ message.PriceValue }

func (f FloorPrice) Tag() fix.Tag { return tag.FloorPrice }

type FlowScheduleType struct{ message.IntValue }

func (f FlowScheduleType) Tag() fix.Tag { return tag.FlowScheduleType }

type ForexReq struct{ message.BooleanValue }

func (f ForexReq) Tag() fix.Tag { return tag.ForexReq }

type FundRenewWaiv struct{ message.CharValue }

func (f FundRenewWaiv) Tag() fix.Tag { return tag.FundRenewWaiv }

type FutSettDate struct{ message.LocalMktDateValue }

func (f FutSettDate) Tag() fix.Tag { return tag.FutSettDate }

type FutSettDate2 struct{ message.LocalMktDateValue }

func (f FutSettDate2) Tag() fix.Tag { return tag.FutSettDate2 }

type FuturesValuationMethod struct{ message.StringValue }

func (f FuturesValuationMethod) Tag() fix.Tag { return tag.FuturesValuationMethod }

type GTBookingInst struct{ message.IntValue }

func (f GTBookingInst) Tag() fix.Tag { return tag.GTBookingInst }

type GapFillFlag struct{ message.BooleanValue }

func (f GapFillFlag) Tag() fix.Tag { return tag.GapFillFlag }

type GrossTradeAmt struct{ message.AmtValue }

func (f GrossTradeAmt) Tag() fix.Tag { return tag.GrossTradeAmt }

type HaltReasonChar struct{ message.CharValue }

func (f HaltReasonChar) Tag() fix.Tag { return tag.HaltReasonChar }

type HaltReasonInt struct{ message.IntValue }

func (f HaltReasonInt) Tag() fix.Tag { return tag.HaltReasonInt }

type HandlInst struct{ message.CharValue }

func (f HandlInst) Tag() fix.Tag { return tag.HandlInst }

type Headline struct{ message.StringValue }

func (f Headline) Tag() fix.Tag { return tag.Headline }

type HeartBtInt struct{ message.IntValue }

func (f HeartBtInt) Tag() fix.Tag { return tag.HeartBtInt }

type HighLimitPrice struct{ message.PriceValue }

func (f HighLimitPrice) Tag() fix.Tag { return tag.HighLimitPrice }

type HighPx struct{ message.PriceValue }

func (f HighPx) Tag() fix.Tag { return tag.HighPx }

type HopCompID struct{ message.StringValue }

func (f HopCompID) Tag() fix.Tag { return tag.HopCompID }

type HopRefID struct{ message.SeqNumValue }

func (f HopRefID) Tag() fix.Tag { return tag.HopRefID }

type HopSendingTime struct{ message.UTCTimestampValue }

func (f HopSendingTime) Tag() fix.Tag { return tag.HopSendingTime }

type HostCrossID struct{ message.StringValue }

func (f HostCrossID) Tag() fix.Tag { return tag.HostCrossID }

type IDSource struct{ message.StringValue }

func (f IDSource) Tag() fix.Tag { return tag.IDSource }

type IOIID struct{ message.StringValue }

func (f IOIID) Tag() fix.Tag { return tag.IOIID }

type IOINaturalFlag struct{ message.BooleanValue }

func (f IOINaturalFlag) Tag() fix.Tag { return tag.IOINaturalFlag }

type IOIOthSvc struct{ message.CharValue }

func (f IOIOthSvc) Tag() fix.Tag { return tag.IOIOthSvc }

type IOIQltyInd struct{ message.CharValue }

func (f IOIQltyInd) Tag() fix.Tag { return tag.IOIQltyInd }

type IOIQty struct{ message.StringValue }

func (f IOIQty) Tag() fix.Tag { return tag.IOIQty }

type IOIQualifier struct{ message.CharValue }

func (f IOIQualifier) Tag() fix.Tag { return tag.IOIQualifier }

type IOIRefID struct{ message.StringValue }

func (f IOIRefID) Tag() fix.Tag { return tag.IOIRefID }

type IOIShares struct{ message.StringValue }

func (f IOIShares) Tag() fix.Tag { return tag.IOIShares }

type IOITransType struct{ message.CharValue }

func (f IOITransType) Tag() fix.Tag { return tag.IOITransType }

type IOIid struct{ message.StringValue }

func (f IOIid) Tag() fix.Tag { return tag.IOIid }

type ImpliedMarketIndicator struct{ message.IntValue }

func (f ImpliedMarketIndicator) Tag() fix.Tag { return tag.ImpliedMarketIndicator }

type InViewOfCommon struct{ message.BooleanValue }

func (f InViewOfCommon) Tag() fix.Tag { return tag.InViewOfCommon }

type IncTaxInd struct{ message.IntValue }

func (f IncTaxInd) Tag() fix.Tag { return tag.IncTaxInd }

type IndividualAllocID struct{ message.StringValue }

func (f IndividualAllocID) Tag() fix.Tag { return tag.IndividualAllocID }

type IndividualAllocRejCode struct{ message.IntValue }

func (f IndividualAllocRejCode) Tag() fix.Tag { return tag.IndividualAllocRejCode }

type IndividualAllocType struct{ message.IntValue }

func (f IndividualAllocType) Tag() fix.Tag { return tag.IndividualAllocType }

type InputSource struct{ message.StringValue }

func (f InputSource) Tag() fix.Tag { return tag.InputSource }

type InstrAttribType struct{ message.IntValue }

func (f InstrAttribType) Tag() fix.Tag { return tag.InstrAttribType }

type InstrAttribValue struct{ message.StringValue }

func (f InstrAttribValue) Tag() fix.Tag { return tag.InstrAttribValue }

type InstrRegistry struct{ message.StringValue }

func (f InstrRegistry) Tag() fix.Tag { return tag.InstrRegistry }

type InstrmtAssignmentMethod struct{ message.CharValue }

func (f InstrmtAssignmentMethod) Tag() fix.Tag { return tag.InstrmtAssignmentMethod }

type InstrumentPartyID struct{ message.StringValue }

func (f InstrumentPartyID) Tag() fix.Tag { return tag.InstrumentPartyID }

type InstrumentPartyIDSource struct{ message.CharValue }

func (f InstrumentPartyIDSource) Tag() fix.Tag { return tag.InstrumentPartyIDSource }

type InstrumentPartyRole struct{ message.IntValue }

func (f InstrumentPartyRole) Tag() fix.Tag { return tag.InstrumentPartyRole }

type InstrumentPartySubID struct{ message.StringValue }

func (f InstrumentPartySubID) Tag() fix.Tag { return tag.InstrumentPartySubID }

type InstrumentPartySubIDType struct{ message.IntValue }

func (f InstrumentPartySubIDType) Tag() fix.Tag { return tag.InstrumentPartySubIDType }

type InterestAccrualDate struct{ message.LocalMktDateValue }

func (f InterestAccrualDate) Tag() fix.Tag { return tag.InterestAccrualDate }

type InterestAtMaturity struct{ message.AmtValue }

func (f InterestAtMaturity) Tag() fix.Tag { return tag.InterestAtMaturity }

type InvestorCountryOfResidence struct{ message.CountryValue }

func (f InvestorCountryOfResidence) Tag() fix.Tag { return tag.InvestorCountryOfResidence }

type IssueDate struct{ message.LocalMktDateValue }

func (f IssueDate) Tag() fix.Tag { return tag.IssueDate }

type Issuer struct{ message.StringValue }

func (f Issuer) Tag() fix.Tag { return tag.Issuer }

type LanguageCode struct{ message.LanguageValue }

func (f LanguageCode) Tag() fix.Tag { return tag.LanguageCode }

type LastCapacity struct{ message.CharValue }

func (f LastCapacity) Tag() fix.Tag { return tag.LastCapacity }

type LastForwardPoints struct{ message.PriceOffsetValue }

func (f LastForwardPoints) Tag() fix.Tag { return tag.LastForwardPoints }

type LastForwardPoints2 struct{ message.PriceOffsetValue }

func (f LastForwardPoints2) Tag() fix.Tag { return tag.LastForwardPoints2 }

type LastFragment struct{ message.BooleanValue }

func (f LastFragment) Tag() fix.Tag { return tag.LastFragment }

type LastLiquidityInd struct{ message.IntValue }

func (f LastLiquidityInd) Tag() fix.Tag { return tag.LastLiquidityInd }

type LastMkt struct{ message.ExchangeValue }

func (f LastMkt) Tag() fix.Tag { return tag.LastMkt }

type LastMsgSeqNumProcessed struct{ message.SeqNumValue }

func (f LastMsgSeqNumProcessed) Tag() fix.Tag { return tag.LastMsgSeqNumProcessed }

type LastNetworkResponseID struct{ message.StringValue }

func (f LastNetworkResponseID) Tag() fix.Tag { return tag.LastNetworkResponseID }

type LastParPx struct{ message.PriceValue }

func (f LastParPx) Tag() fix.Tag { return tag.LastParPx }

type LastPx struct{ message.PriceValue }

func (f LastPx) Tag() fix.Tag { return tag.LastPx }

type LastQty struct{ message.QtyValue }

func (f LastQty) Tag() fix.Tag { return tag.LastQty }

type LastRptRequested struct{ message.BooleanValue }

func (f LastRptRequested) Tag() fix.Tag { return tag.LastRptRequested }

type LastShares struct{ message.QtyValue }

func (f LastShares) Tag() fix.Tag { return tag.LastShares }

type LastSpotRate struct{ message.PriceValue }

func (f LastSpotRate) Tag() fix.Tag { return tag.LastSpotRate }

type LastSwapPoints struct{ message.PriceOffsetValue }

func (f LastSwapPoints) Tag() fix.Tag { return tag.LastSwapPoints }

type LastUpdateTime struct{ message.UTCTimestampValue }

func (f LastUpdateTime) Tag() fix.Tag { return tag.LastUpdateTime }

type LateIndicator struct{ message.BooleanValue }

func (f LateIndicator) Tag() fix.Tag { return tag.LateIndicator }

type LeavesQty struct{ message.QtyValue }

func (f LeavesQty) Tag() fix.Tag { return tag.LeavesQty }

type LegAllocAccount struct{ message.StringValue }

func (f LegAllocAccount) Tag() fix.Tag { return tag.LegAllocAccount }

type LegAllocAcctIDSource struct{ message.StringValue }

func (f LegAllocAcctIDSource) Tag() fix.Tag { return tag.LegAllocAcctIDSource }

type LegAllocID struct{ message.StringValue }

func (f LegAllocID) Tag() fix.Tag { return tag.LegAllocID }

type LegAllocQty struct{ message.QtyValue }

func (f LegAllocQty) Tag() fix.Tag { return tag.LegAllocQty }

type LegAllocSettlCurrency struct{ message.CurrencyValue }

func (f LegAllocSettlCurrency) Tag() fix.Tag { return tag.LegAllocSettlCurrency }

type LegBenchmarkCurveCurrency struct{ message.CurrencyValue }

func (f LegBenchmarkCurveCurrency) Tag() fix.Tag { return tag.LegBenchmarkCurveCurrency }

type LegBenchmarkCurveName struct{ message.StringValue }

func (f LegBenchmarkCurveName) Tag() fix.Tag { return tag.LegBenchmarkCurveName }

type LegBenchmarkCurvePoint struct{ message.StringValue }

func (f LegBenchmarkCurvePoint) Tag() fix.Tag { return tag.LegBenchmarkCurvePoint }

type LegBenchmarkPrice struct{ message.PriceValue }

func (f LegBenchmarkPrice) Tag() fix.Tag { return tag.LegBenchmarkPrice }

type LegBenchmarkPriceType struct{ message.IntValue }

func (f LegBenchmarkPriceType) Tag() fix.Tag { return tag.LegBenchmarkPriceType }

type LegBidForwardPoints struct{ message.PriceOffsetValue }

func (f LegBidForwardPoints) Tag() fix.Tag { return tag.LegBidForwardPoints }

type LegBidPx struct{ message.PriceValue }

func (f LegBidPx) Tag() fix.Tag { return tag.LegBidPx }

type LegCFICode struct{ message.StringValue }

func (f LegCFICode) Tag() fix.Tag { return tag.LegCFICode }

type LegCalculatedCcyLastQty struct{ message.QtyValue }

func (f LegCalculatedCcyLastQty) Tag() fix.Tag { return tag.LegCalculatedCcyLastQty }

type LegContractMultiplier struct{ message.FloatValue }

func (f LegContractMultiplier) Tag() fix.Tag { return tag.LegContractMultiplier }

type LegContractMultiplierUnit struct{ message.IntValue }

func (f LegContractMultiplierUnit) Tag() fix.Tag { return tag.LegContractMultiplierUnit }

type LegContractSettlMonth struct{ message.MonthYearValue }

func (f LegContractSettlMonth) Tag() fix.Tag { return tag.LegContractSettlMonth }

type LegCountryOfIssue struct{ message.CountryValue }

func (f LegCountryOfIssue) Tag() fix.Tag { return tag.LegCountryOfIssue }

type LegCouponPaymentDate struct{ message.LocalMktDateValue }

func (f LegCouponPaymentDate) Tag() fix.Tag { return tag.LegCouponPaymentDate }

type LegCouponRate struct{ message.PercentageValue }

func (f LegCouponRate) Tag() fix.Tag { return tag.LegCouponRate }

type LegCoveredOrUncovered struct{ message.IntValue }

func (f LegCoveredOrUncovered) Tag() fix.Tag { return tag.LegCoveredOrUncovered }

type LegCreditRating struct{ message.StringValue }

func (f LegCreditRating) Tag() fix.Tag { return tag.LegCreditRating }

type LegCurrency struct{ message.CurrencyValue }

func (f LegCurrency) Tag() fix.Tag { return tag.LegCurrency }

type LegCurrencyRatio struct{ message.FloatValue }

func (f LegCurrencyRatio) Tag() fix.Tag { return tag.LegCurrencyRatio }

type LegDatedDate struct{ message.LocalMktDateValue }

func (f LegDatedDate) Tag() fix.Tag { return tag.LegDatedDate }

type LegDividendYield struct{ message.PercentageValue }

func (f LegDividendYield) Tag() fix.Tag { return tag.LegDividendYield }

type LegExecInst struct{ message.MultipleCharValue }

func (f LegExecInst) Tag() fix.Tag { return tag.LegExecInst }

type LegExerciseStyle struct{ message.IntValue }

func (f LegExerciseStyle) Tag() fix.Tag { return tag.LegExerciseStyle }

type LegFactor struct{ message.FloatValue }

func (f LegFactor) Tag() fix.Tag { return tag.LegFactor }

type LegFlowScheduleType struct{ message.IntValue }

func (f LegFlowScheduleType) Tag() fix.Tag { return tag.LegFlowScheduleType }

type LegFutSettDate struct{ message.LocalMktDateValue }

func (f LegFutSettDate) Tag() fix.Tag { return tag.LegFutSettDate }

type LegGrossTradeAmt struct{ message.AmtValue }

func (f LegGrossTradeAmt) Tag() fix.Tag { return tag.LegGrossTradeAmt }

type LegIOIQty struct{ message.StringValue }

func (f LegIOIQty) Tag() fix.Tag { return tag.LegIOIQty }

type LegIndividualAllocID struct{ message.StringValue }

func (f LegIndividualAllocID) Tag() fix.Tag { return tag.LegIndividualAllocID }

type LegInstrRegistry struct{ message.StringValue }

func (f LegInstrRegistry) Tag() fix.Tag { return tag.LegInstrRegistry }

type LegInterestAccrualDate struct{ message.LocalMktDateValue }

func (f LegInterestAccrualDate) Tag() fix.Tag { return tag.LegInterestAccrualDate }

type LegIssueDate struct{ message.LocalMktDateValue }

func (f LegIssueDate) Tag() fix.Tag { return tag.LegIssueDate }

type LegIssuer struct{ message.StringValue }

func (f LegIssuer) Tag() fix.Tag { return tag.LegIssuer }

type LegLastForwardPoints struct{ message.PriceOffsetValue }

func (f LegLastForwardPoints) Tag() fix.Tag { return tag.LegLastForwardPoints }

type LegLastPx struct{ message.PriceValue }

func (f LegLastPx) Tag() fix.Tag { return tag.LegLastPx }

type LegLastQty struct{ message.QtyValue }

func (f LegLastQty) Tag() fix.Tag { return tag.LegLastQty }

type LegLocaleOfIssue struct{ message.StringValue }

func (f LegLocaleOfIssue) Tag() fix.Tag { return tag.LegLocaleOfIssue }

type LegMaturityDate struct{ message.LocalMktDateValue }

func (f LegMaturityDate) Tag() fix.Tag { return tag.LegMaturityDate }

type LegMaturityMonthYear struct{ message.MonthYearValue }

func (f LegMaturityMonthYear) Tag() fix.Tag { return tag.LegMaturityMonthYear }

type LegMaturityTime struct{ message.TZTimeOnlyValue }

func (f LegMaturityTime) Tag() fix.Tag { return tag.LegMaturityTime }

type LegNumber struct{ message.IntValue }

func (f LegNumber) Tag() fix.Tag { return tag.LegNumber }

type LegOfferForwardPoints struct{ message.PriceOffsetValue }

func (f LegOfferForwardPoints) Tag() fix.Tag { return tag.LegOfferForwardPoints }

type LegOfferPx struct{ message.PriceValue }

func (f LegOfferPx) Tag() fix.Tag { return tag.LegOfferPx }

type LegOptAttribute struct{ message.CharValue }

func (f LegOptAttribute) Tag() fix.Tag { return tag.LegOptAttribute }

type LegOptionRatio struct{ message.FloatValue }

func (f LegOptionRatio) Tag() fix.Tag { return tag.LegOptionRatio }

type LegOrderQty struct{ message.QtyValue }

func (f LegOrderQty) Tag() fix.Tag { return tag.LegOrderQty }

type LegPool struct{ message.StringValue }

func (f LegPool) Tag() fix.Tag { return tag.LegPool }

type LegPositionEffect struct{ message.CharValue }

func (f LegPositionEffect) Tag() fix.Tag { return tag.LegPositionEffect }

type LegPrice struct{ message.PriceValue }

func (f LegPrice) Tag() fix.Tag { return tag.LegPrice }

type LegPriceType struct{ message.IntValue }

func (f LegPriceType) Tag() fix.Tag { return tag.LegPriceType }

type LegPriceUnitOfMeasure struct{ message.StringValue }

func (f LegPriceUnitOfMeasure) Tag() fix.Tag { return tag.LegPriceUnitOfMeasure }

type LegPriceUnitOfMeasureQty struct{ message.QtyValue }

func (f LegPriceUnitOfMeasureQty) Tag() fix.Tag { return tag.LegPriceUnitOfMeasureQty }

type LegProduct struct{ message.IntValue }

func (f LegProduct) Tag() fix.Tag { return tag.LegProduct }

type LegPutOrCall struct{ message.IntValue }

func (f LegPutOrCall) Tag() fix.Tag { return tag.LegPutOrCall }

type LegQty struct{ message.QtyValue }

func (f LegQty) Tag() fix.Tag { return tag.LegQty }

type LegRatioQty struct{ message.FloatValue }

func (f LegRatioQty) Tag() fix.Tag { return tag.LegRatioQty }

type LegRedemptionDate struct{ message.LocalMktDateValue }

func (f LegRedemptionDate) Tag() fix.Tag { return tag.LegRedemptionDate }

type LegRefID struct{ message.StringValue }

func (f LegRefID) Tag() fix.Tag { return tag.LegRefID }

type LegRepoCollateralSecurityType struct{ message.IntValue }

func (f LegRepoCollateralSecurityType) Tag() fix.Tag { return tag.LegRepoCollateralSecurityType }

type LegReportID struct{ message.StringValue }

func (f LegReportID) Tag() fix.Tag { return tag.LegReportID }

type LegRepurchaseRate struct{ message.PercentageValue }

func (f LegRepurchaseRate) Tag() fix.Tag { return tag.LegRepurchaseRate }

type LegRepurchaseTerm struct{ message.IntValue }

func (f LegRepurchaseTerm) Tag() fix.Tag { return tag.LegRepurchaseTerm }

type LegSecurityAltID struct{ message.StringValue }

func (f LegSecurityAltID) Tag() fix.Tag { return tag.LegSecurityAltID }

type LegSecurityAltIDSource struct{ message.StringValue }

func (f LegSecurityAltIDSource) Tag() fix.Tag { return tag.LegSecurityAltIDSource }

type LegSecurityDesc struct{ message.StringValue }

func (f LegSecurityDesc) Tag() fix.Tag { return tag.LegSecurityDesc }

type LegSecurityExchange struct{ message.ExchangeValue }

func (f LegSecurityExchange) Tag() fix.Tag { return tag.LegSecurityExchange }

type LegSecurityID struct{ message.StringValue }

func (f LegSecurityID) Tag() fix.Tag { return tag.LegSecurityID }

type LegSecurityIDSource struct{ message.StringValue }

func (f LegSecurityIDSource) Tag() fix.Tag { return tag.LegSecurityIDSource }

type LegSecuritySubType struct{ message.StringValue }

func (f LegSecuritySubType) Tag() fix.Tag { return tag.LegSecuritySubType }

type LegSecurityType struct{ message.StringValue }

func (f LegSecurityType) Tag() fix.Tag { return tag.LegSecurityType }

type LegSettlCurrency struct{ message.CurrencyValue }

func (f LegSettlCurrency) Tag() fix.Tag { return tag.LegSettlCurrency }

type LegSettlDate struct{ message.LocalMktDateValue }

func (f LegSettlDate) Tag() fix.Tag { return tag.LegSettlDate }

type LegSettlType struct{ message.CharValue }

func (f LegSettlType) Tag() fix.Tag { return tag.LegSettlType }

type LegSettlmntTyp struct{ message.CharValue }

func (f LegSettlmntTyp) Tag() fix.Tag { return tag.LegSettlmntTyp }

type LegSide struct{ message.CharValue }

func (f LegSide) Tag() fix.Tag { return tag.LegSide }

type LegStateOrProvinceOfIssue struct{ message.StringValue }

func (f LegStateOrProvinceOfIssue) Tag() fix.Tag { return tag.LegStateOrProvinceOfIssue }

type LegStipulationType struct{ message.StringValue }

func (f LegStipulationType) Tag() fix.Tag { return tag.LegStipulationType }

type LegStipulationValue struct{ message.StringValue }

func (f LegStipulationValue) Tag() fix.Tag { return tag.LegStipulationValue }

type LegStrikeCurrency struct{ message.CurrencyValue }

func (f LegStrikeCurrency) Tag() fix.Tag { return tag.LegStrikeCurrency }

type LegStrikePrice struct{ message.PriceValue }

func (f LegStrikePrice) Tag() fix.Tag { return tag.LegStrikePrice }

type LegSwapType struct{ message.IntValue }

func (f LegSwapType) Tag() fix.Tag { return tag.LegSwapType }

type LegSymbol struct{ message.StringValue }

func (f LegSymbol) Tag() fix.Tag { return tag.LegSymbol }

type LegSymbolSfx struct{ message.StringValue }

func (f LegSymbolSfx) Tag() fix.Tag { return tag.LegSymbolSfx }

type LegTimeUnit struct{ message.StringValue }

func (f LegTimeUnit) Tag() fix.Tag { return tag.LegTimeUnit }

type LegUnitOfMeasure struct{ message.StringValue }

func (f LegUnitOfMeasure) Tag() fix.Tag { return tag.LegUnitOfMeasure }

type LegUnitOfMeasureQty struct{ message.QtyValue }

func (f LegUnitOfMeasureQty) Tag() fix.Tag { return tag.LegUnitOfMeasureQty }

type LegVolatility struct{ message.FloatValue }

func (f LegVolatility) Tag() fix.Tag { return tag.LegVolatility }

type LegalConfirm struct{ message.BooleanValue }

func (f LegalConfirm) Tag() fix.Tag { return tag.LegalConfirm }

type LinesOfText struct{ message.NumInGroupValue }

func (f LinesOfText) Tag() fix.Tag { return tag.LinesOfText }

type LiquidityIndType struct{ message.IntValue }

func (f LiquidityIndType) Tag() fix.Tag { return tag.LiquidityIndType }

type LiquidityNumSecurities struct{ message.IntValue }

func (f LiquidityNumSecurities) Tag() fix.Tag { return tag.LiquidityNumSecurities }

type LiquidityPctHigh struct{ message.PercentageValue }

func (f LiquidityPctHigh) Tag() fix.Tag { return tag.LiquidityPctHigh }

type LiquidityPctLow struct{ message.PercentageValue }

func (f LiquidityPctLow) Tag() fix.Tag { return tag.LiquidityPctLow }

type LiquidityValue struct{ message.AmtValue }

func (f LiquidityValue) Tag() fix.Tag { return tag.LiquidityValue }

type ListExecInst struct{ message.StringValue }

func (f ListExecInst) Tag() fix.Tag { return tag.ListExecInst }

type ListExecInstType struct{ message.CharValue }

func (f ListExecInstType) Tag() fix.Tag { return tag.ListExecInstType }

type ListID struct{ message.StringValue }

func (f ListID) Tag() fix.Tag { return tag.ListID }

type ListMethod struct{ message.IntValue }

func (f ListMethod) Tag() fix.Tag { return tag.ListMethod }

type ListName struct{ message.StringValue }

func (f ListName) Tag() fix.Tag { return tag.ListName }

type ListNoOrds struct{ message.IntValue }

func (f ListNoOrds) Tag() fix.Tag { return tag.ListNoOrds }

type ListOrderStatus struct{ message.IntValue }

func (f ListOrderStatus) Tag() fix.Tag { return tag.ListOrderStatus }

type ListRejectReason struct{ message.IntValue }

func (f ListRejectReason) Tag() fix.Tag { return tag.ListRejectReason }

type ListSeqNo struct{ message.IntValue }

func (f ListSeqNo) Tag() fix.Tag { return tag.ListSeqNo }

type ListStatusText struct{ message.StringValue }

func (f ListStatusText) Tag() fix.Tag { return tag.ListStatusText }

type ListStatusType struct{ message.IntValue }

func (f ListStatusType) Tag() fix.Tag { return tag.ListStatusType }

type ListUpdateAction struct{ message.CharValue }

func (f ListUpdateAction) Tag() fix.Tag { return tag.ListUpdateAction }

type LocaleOfIssue struct{ message.StringValue }

func (f LocaleOfIssue) Tag() fix.Tag { return tag.LocaleOfIssue }

type LocateReqd struct{ message.BooleanValue }

func (f LocateReqd) Tag() fix.Tag { return tag.LocateReqd }

type LocationID struct{ message.StringValue }

func (f LocationID) Tag() fix.Tag { return tag.LocationID }

type LongQty struct{ message.QtyValue }

func (f LongQty) Tag() fix.Tag { return tag.LongQty }

type LotType struct{ message.CharValue }

func (f LotType) Tag() fix.Tag { return tag.LotType }

type LowLimitPrice struct{ message.PriceValue }

func (f LowLimitPrice) Tag() fix.Tag { return tag.LowLimitPrice }

type LowPx struct{ message.PriceValue }

func (f LowPx) Tag() fix.Tag { return tag.LowPx }

type MDBookType struct{ message.IntValue }

func (f MDBookType) Tag() fix.Tag { return tag.MDBookType }

type MDEntryBuyer struct{ message.StringValue }

func (f MDEntryBuyer) Tag() fix.Tag { return tag.MDEntryBuyer }

type MDEntryDate struct{ message.UTCDateOnlyValue }

func (f MDEntryDate) Tag() fix.Tag { return tag.MDEntryDate }

type MDEntryForwardPoints struct{ message.PriceOffsetValue }

func (f MDEntryForwardPoints) Tag() fix.Tag { return tag.MDEntryForwardPoints }

type MDEntryID struct{ message.StringValue }

func (f MDEntryID) Tag() fix.Tag { return tag.MDEntryID }

type MDEntryOriginator struct{ message.StringValue }

func (f MDEntryOriginator) Tag() fix.Tag { return tag.MDEntryOriginator }

type MDEntryPositionNo struct{ message.IntValue }

func (f MDEntryPositionNo) Tag() fix.Tag { return tag.MDEntryPositionNo }

type MDEntryPx struct{ message.PriceValue }

func (f MDEntryPx) Tag() fix.Tag { return tag.MDEntryPx }

type MDEntryRefID struct{ message.StringValue }

func (f MDEntryRefID) Tag() fix.Tag { return tag.MDEntryRefID }

type MDEntrySeller struct{ message.StringValue }

func (f MDEntrySeller) Tag() fix.Tag { return tag.MDEntrySeller }

type MDEntrySize struct{ message.QtyValue }

func (f MDEntrySize) Tag() fix.Tag { return tag.MDEntrySize }

type MDEntrySpotRate struct{ message.FloatValue }

func (f MDEntrySpotRate) Tag() fix.Tag { return tag.MDEntrySpotRate }

type MDEntryTime struct{ message.UTCTimeOnlyValue }

func (f MDEntryTime) Tag() fix.Tag { return tag.MDEntryTime }

type MDEntryType struct{ message.CharValue }

func (f MDEntryType) Tag() fix.Tag { return tag.MDEntryType }

type MDFeedType struct{ message.StringValue }

func (f MDFeedType) Tag() fix.Tag { return tag.MDFeedType }

type MDImplicitDelete struct{ message.BooleanValue }

func (f MDImplicitDelete) Tag() fix.Tag { return tag.MDImplicitDelete }

type MDMkt struct{ message.ExchangeValue }

func (f MDMkt) Tag() fix.Tag { return tag.MDMkt }

type MDOriginType struct{ message.IntValue }

func (f MDOriginType) Tag() fix.Tag { return tag.MDOriginType }

type MDPriceLevel struct{ message.IntValue }

func (f MDPriceLevel) Tag() fix.Tag { return tag.MDPriceLevel }

type MDQuoteType struct{ message.IntValue }

func (f MDQuoteType) Tag() fix.Tag { return tag.MDQuoteType }

type MDReportID struct{ message.IntValue }

func (f MDReportID) Tag() fix.Tag { return tag.MDReportID }

type MDReqID struct{ message.StringValue }

func (f MDReqID) Tag() fix.Tag { return tag.MDReqID }

type MDReqRejReason struct{ message.CharValue }

func (f MDReqRejReason) Tag() fix.Tag { return tag.MDReqRejReason }

type MDSecSize struct{ message.QtyValue }

func (f MDSecSize) Tag() fix.Tag { return tag.MDSecSize }

type MDSecSizeType struct{ message.IntValue }

func (f MDSecSizeType) Tag() fix.Tag { return tag.MDSecSizeType }

type MDStreamID struct{ message.StringValue }

func (f MDStreamID) Tag() fix.Tag { return tag.MDStreamID }

type MDSubBookType struct{ message.IntValue }

func (f MDSubBookType) Tag() fix.Tag { return tag.MDSubBookType }

type MDUpdateAction struct{ message.CharValue }

func (f MDUpdateAction) Tag() fix.Tag { return tag.MDUpdateAction }

type MDUpdateType struct{ message.IntValue }

func (f MDUpdateType) Tag() fix.Tag { return tag.MDUpdateType }

type MailingDtls struct{ message.StringValue }

func (f MailingDtls) Tag() fix.Tag { return tag.MailingDtls }

type MailingInst struct{ message.StringValue }

func (f MailingInst) Tag() fix.Tag { return tag.MailingInst }

type ManualOrderIndicator struct{ message.BooleanValue }

func (f ManualOrderIndicator) Tag() fix.Tag { return tag.ManualOrderIndicator }

type MarginExcess struct{ message.AmtValue }

func (f MarginExcess) Tag() fix.Tag { return tag.MarginExcess }

type MarginRatio struct{ message.PercentageValue }

func (f MarginRatio) Tag() fix.Tag { return tag.MarginRatio }

type MarketDepth struct{ message.IntValue }

func (f MarketDepth) Tag() fix.Tag { return tag.MarketDepth }

type MarketID struct{ message.ExchangeValue }

func (f MarketID) Tag() fix.Tag { return tag.MarketID }

type MarketReportID struct{ message.StringValue }

func (f MarketReportID) Tag() fix.Tag { return tag.MarketReportID }

type MarketReqID struct{ message.StringValue }

func (f MarketReqID) Tag() fix.Tag { return tag.MarketReqID }

type MarketSegmentDesc struct{ message.StringValue }

func (f MarketSegmentDesc) Tag() fix.Tag { return tag.MarketSegmentDesc }

type MarketSegmentID struct{ message.StringValue }

func (f MarketSegmentID) Tag() fix.Tag { return tag.MarketSegmentID }

type MarketUpdateAction struct{ message.CharValue }

func (f MarketUpdateAction) Tag() fix.Tag { return tag.MarketUpdateAction }

type MassActionRejectReason struct{ message.IntValue }

func (f MassActionRejectReason) Tag() fix.Tag { return tag.MassActionRejectReason }

type MassActionReportID struct{ message.StringValue }

func (f MassActionReportID) Tag() fix.Tag { return tag.MassActionReportID }

type MassActionResponse struct{ message.IntValue }

func (f MassActionResponse) Tag() fix.Tag { return tag.MassActionResponse }

type MassActionScope struct{ message.IntValue }

func (f MassActionScope) Tag() fix.Tag { return tag.MassActionScope }

type MassActionType struct{ message.IntValue }

func (f MassActionType) Tag() fix.Tag { return tag.MassActionType }

type MassCancelRejectReason struct{ message.IntValue }

func (f MassCancelRejectReason) Tag() fix.Tag { return tag.MassCancelRejectReason }

type MassCancelRequestType struct{ message.CharValue }

func (f MassCancelRequestType) Tag() fix.Tag { return tag.MassCancelRequestType }

type MassCancelResponse struct{ message.CharValue }

func (f MassCancelResponse) Tag() fix.Tag { return tag.MassCancelResponse }

type MassStatusReqID struct{ message.StringValue }

func (f MassStatusReqID) Tag() fix.Tag { return tag.MassStatusReqID }

type MassStatusReqType struct{ message.IntValue }

func (f MassStatusReqType) Tag() fix.Tag { return tag.MassStatusReqType }

type MatchAlgorithm struct{ message.StringValue }

func (f MatchAlgorithm) Tag() fix.Tag { return tag.MatchAlgorithm }

type MatchIncrement struct{ message.QtyValue }

func (f MatchIncrement) Tag() fix.Tag { return tag.MatchIncrement }

type MatchStatus struct{ message.CharValue }

func (f MatchStatus) Tag() fix.Tag { return tag.MatchStatus }

type MatchType struct{ message.StringValue }

func (f MatchType) Tag() fix.Tag { return tag.MatchType }

type MaturityDate struct{ message.LocalMktDateValue }

func (f MaturityDate) Tag() fix.Tag { return tag.MaturityDate }

type MaturityDay struct{ message.DayOfMonthValue }

func (f MaturityDay) Tag() fix.Tag { return tag.MaturityDay }

type MaturityMonthYear struct{ message.MonthYearValue }

func (f MaturityMonthYear) Tag() fix.Tag { return tag.MaturityMonthYear }

type MaturityMonthYearFormat struct{ message.IntValue }

func (f MaturityMonthYearFormat) Tag() fix.Tag { return tag.MaturityMonthYearFormat }

type MaturityMonthYearIncrement struct{ message.IntValue }

func (f MaturityMonthYearIncrement) Tag() fix.Tag { return tag.MaturityMonthYearIncrement }

type MaturityMonthYearIncrementUnits struct{ message.IntValue }

func (f MaturityMonthYearIncrementUnits) Tag() fix.Tag { return tag.MaturityMonthYearIncrementUnits }

type MaturityNetMoney struct{ message.AmtValue }

func (f MaturityNetMoney) Tag() fix.Tag { return tag.MaturityNetMoney }

type MaturityRuleID struct{ message.StringValue }

func (f MaturityRuleID) Tag() fix.Tag { return tag.MaturityRuleID }

type MaturityTime struct{ message.TZTimeOnlyValue }

func (f MaturityTime) Tag() fix.Tag { return tag.MaturityTime }

type MaxFloor struct{ message.QtyValue }

func (f MaxFloor) Tag() fix.Tag { return tag.MaxFloor }

type MaxMessageSize struct{ message.LengthValue }

func (f MaxMessageSize) Tag() fix.Tag { return tag.MaxMessageSize }

type MaxPriceLevels struct{ message.IntValue }

func (f MaxPriceLevels) Tag() fix.Tag { return tag.MaxPriceLevels }

type MaxPriceVariation struct{ message.FloatValue }

func (f MaxPriceVariation) Tag() fix.Tag { return tag.MaxPriceVariation }

type MaxShow struct{ message.QtyValue }

func (f MaxShow) Tag() fix.Tag { return tag.MaxShow }

type MaxTradeVol struct{ message.QtyValue }

func (f MaxTradeVol) Tag() fix.Tag { return tag.MaxTradeVol }

type MessageEncoding struct{ message.StringValue }

func (f MessageEncoding) Tag() fix.Tag { return tag.MessageEncoding }

type MessageEventSource struct{ message.StringValue }

func (f MessageEventSource) Tag() fix.Tag { return tag.MessageEventSource }

type MidPx struct{ message.PriceValue }

func (f MidPx) Tag() fix.Tag { return tag.MidPx }

type MidYield struct{ message.PercentageValue }

func (f MidYield) Tag() fix.Tag { return tag.MidYield }

type MinBidSize struct{ message.QtyValue }

func (f MinBidSize) Tag() fix.Tag { return tag.MinBidSize }

type MinLotSize struct{ message.QtyValue }

func (f MinLotSize) Tag() fix.Tag { return tag.MinLotSize }

type MinOfferSize struct{ message.QtyValue }

func (f MinOfferSize) Tag() fix.Tag { return tag.MinOfferSize }

type MinPriceIncrement struct{ message.FloatValue }

func (f MinPriceIncrement) Tag() fix.Tag { return tag.MinPriceIncrement }

type MinPriceIncrementAmount struct{ message.AmtValue }

func (f MinPriceIncrementAmount) Tag() fix.Tag { return tag.MinPriceIncrementAmount }

type MinQty struct{ message.QtyValue }

func (f MinQty) Tag() fix.Tag { return tag.MinQty }

type MinTradeVol struct{ message.QtyValue }

func (f MinTradeVol) Tag() fix.Tag { return tag.MinTradeVol }

type MiscFeeAmt struct{ message.AmtValue }

func (f MiscFeeAmt) Tag() fix.Tag { return tag.MiscFeeAmt }

type MiscFeeBasis struct{ message.IntValue }

func (f MiscFeeBasis) Tag() fix.Tag { return tag.MiscFeeBasis }

type MiscFeeCurr struct{ message.CurrencyValue }

func (f MiscFeeCurr) Tag() fix.Tag { return tag.MiscFeeCurr }

type MiscFeeType struct{ message.StringValue }

func (f MiscFeeType) Tag() fix.Tag { return tag.MiscFeeType }

type MktBidPx struct{ message.PriceValue }

func (f MktBidPx) Tag() fix.Tag { return tag.MktBidPx }

type MktOfferPx struct{ message.PriceValue }

func (f MktOfferPx) Tag() fix.Tag { return tag.MktOfferPx }

type ModelType struct{ message.IntValue }

func (f ModelType) Tag() fix.Tag { return tag.ModelType }

type MoneyLaunderingStatus struct{ message.CharValue }

func (f MoneyLaunderingStatus) Tag() fix.Tag { return tag.MoneyLaunderingStatus }

type MsgDirection struct{ message.CharValue }

func (f MsgDirection) Tag() fix.Tag { return tag.MsgDirection }

type MsgSeqNum struct{ message.SeqNumValue }

func (f MsgSeqNum) Tag() fix.Tag { return tag.MsgSeqNum }

type MsgType struct{ message.StringValue }

func (f MsgType) Tag() fix.Tag { return tag.MsgType }

type MultiLegReportingType struct{ message.CharValue }

func (f MultiLegReportingType) Tag() fix.Tag { return tag.MultiLegReportingType }

type MultiLegRptTypeReq struct{ message.IntValue }

func (f MultiLegRptTypeReq) Tag() fix.Tag { return tag.MultiLegRptTypeReq }

type MultilegModel struct{ message.IntValue }

func (f MultilegModel) Tag() fix.Tag { return tag.MultilegModel }

type MultilegPriceMethod struct{ message.IntValue }

func (f MultilegPriceMethod) Tag() fix.Tag { return tag.MultilegPriceMethod }

type NTPositionLimit struct{ message.IntValue }

func (f NTPositionLimit) Tag() fix.Tag { return tag.NTPositionLimit }

type Nested2PartyID struct{ message.StringValue }

func (f Nested2PartyID) Tag() fix.Tag { return tag.Nested2PartyID }

type Nested2PartyIDSource struct{ message.CharValue }

func (f Nested2PartyIDSource) Tag() fix.Tag { return tag.Nested2PartyIDSource }

type Nested2PartyRole struct{ message.IntValue }

func (f Nested2PartyRole) Tag() fix.Tag { return tag.Nested2PartyRole }

type Nested2PartySubID struct{ message.StringValue }

func (f Nested2PartySubID) Tag() fix.Tag { return tag.Nested2PartySubID }

type Nested2PartySubIDType struct{ message.IntValue }

func (f Nested2PartySubIDType) Tag() fix.Tag { return tag.Nested2PartySubIDType }

type Nested3PartyID struct{ message.StringValue }

func (f Nested3PartyID) Tag() fix.Tag { return tag.Nested3PartyID }

type Nested3PartyIDSource struct{ message.CharValue }

func (f Nested3PartyIDSource) Tag() fix.Tag { return tag.Nested3PartyIDSource }

type Nested3PartyRole struct{ message.IntValue }

func (f Nested3PartyRole) Tag() fix.Tag { return tag.Nested3PartyRole }

type Nested3PartySubID struct{ message.StringValue }

func (f Nested3PartySubID) Tag() fix.Tag { return tag.Nested3PartySubID }

type Nested3PartySubIDType struct{ message.IntValue }

func (f Nested3PartySubIDType) Tag() fix.Tag { return tag.Nested3PartySubIDType }

type Nested4PartyID struct{ message.StringValue }

func (f Nested4PartyID) Tag() fix.Tag { return tag.Nested4PartyID }

type Nested4PartyIDSource struct{ message.CharValue }

func (f Nested4PartyIDSource) Tag() fix.Tag { return tag.Nested4PartyIDSource }

type Nested4PartyRole struct{ message.IntValue }

func (f Nested4PartyRole) Tag() fix.Tag { return tag.Nested4PartyRole }

type Nested4PartySubID struct{ message.StringValue }

func (f Nested4PartySubID) Tag() fix.Tag { return tag.Nested4PartySubID }

type Nested4PartySubIDType struct{ message.IntValue }

func (f Nested4PartySubIDType) Tag() fix.Tag { return tag.Nested4PartySubIDType }

type NestedInstrAttribType struct{ message.IntValue }

func (f NestedInstrAttribType) Tag() fix.Tag { return tag.NestedInstrAttribType }

type NestedInstrAttribValue struct{ message.StringValue }

func (f NestedInstrAttribValue) Tag() fix.Tag { return tag.NestedInstrAttribValue }

type NestedPartyID struct{ message.StringValue }

func (f NestedPartyID) Tag() fix.Tag { return tag.NestedPartyID }

type NestedPartyIDSource struct{ message.CharValue }

func (f NestedPartyIDSource) Tag() fix.Tag { return tag.NestedPartyIDSource }

type NestedPartyRole struct{ message.IntValue }

func (f NestedPartyRole) Tag() fix.Tag { return tag.NestedPartyRole }

type NestedPartySubID struct{ message.StringValue }

func (f NestedPartySubID) Tag() fix.Tag { return tag.NestedPartySubID }

type NestedPartySubIDType struct{ message.IntValue }

func (f NestedPartySubIDType) Tag() fix.Tag { return tag.NestedPartySubIDType }

type NetChgPrevDay struct{ message.PriceOffsetValue }

func (f NetChgPrevDay) Tag() fix.Tag { return tag.NetChgPrevDay }

type NetGrossInd struct{ message.IntValue }

func (f NetGrossInd) Tag() fix.Tag { return tag.NetGrossInd }

type NetMoney struct{ message.AmtValue }

func (f NetMoney) Tag() fix.Tag { return tag.NetMoney }

type NetworkRequestID struct{ message.StringValue }

func (f NetworkRequestID) Tag() fix.Tag { return tag.NetworkRequestID }

type NetworkRequestType struct{ message.IntValue }

func (f NetworkRequestType) Tag() fix.Tag { return tag.NetworkRequestType }

type NetworkResponseID struct{ message.StringValue }

func (f NetworkResponseID) Tag() fix.Tag { return tag.NetworkResponseID }

type NetworkStatusResponseType struct{ message.IntValue }

func (f NetworkStatusResponseType) Tag() fix.Tag { return tag.NetworkStatusResponseType }

type NewPassword struct{ message.StringValue }

func (f NewPassword) Tag() fix.Tag { return tag.NewPassword }

type NewSeqNo struct{ message.SeqNumValue }

func (f NewSeqNo) Tag() fix.Tag { return tag.NewSeqNo }

type NewsCategory struct{ message.IntValue }

func (f NewsCategory) Tag() fix.Tag { return tag.NewsCategory }

type NewsID struct{ message.StringValue }

func (f NewsID) Tag() fix.Tag { return tag.NewsID }

type NewsRefID struct{ message.StringValue }

func (f NewsRefID) Tag() fix.Tag { return tag.NewsRefID }

type NewsRefType struct{ message.IntValue }

func (f NewsRefType) Tag() fix.Tag { return tag.NewsRefType }

type NextExpectedMsgSeqNum struct{ message.SeqNumValue }

func (f NextExpectedMsgSeqNum) Tag() fix.Tag { return tag.NextExpectedMsgSeqNum }

type NoAffectedOrders struct{ message.NumInGroupValue }

func (f NoAffectedOrders) Tag() fix.Tag { return tag.NoAffectedOrders }

type NoAllocs struct{ message.NumInGroupValue }

func (f NoAllocs) Tag() fix.Tag { return tag.NoAllocs }

type NoAltMDSource struct{ message.NumInGroupValue }

func (f NoAltMDSource) Tag() fix.Tag { return tag.NoAltMDSource }

type NoApplIDs struct{ message.NumInGroupValue }

func (f NoApplIDs) Tag() fix.Tag { return tag.NoApplIDs }

type NoAsgnReqs struct{ message.NumInGroupValue }

func (f NoAsgnReqs) Tag() fix.Tag { return tag.NoAsgnReqs }

type NoBidComponents struct{ message.NumInGroupValue }

func (f NoBidComponents) Tag() fix.Tag { return tag.NoBidComponents }

type NoBidDescriptors struct{ message.NumInGroupValue }

func (f NoBidDescriptors) Tag() fix.Tag { return tag.NoBidDescriptors }

type NoCapacities struct{ message.NumInGroupValue }

func (f NoCapacities) Tag() fix.Tag { return tag.NoCapacities }

type NoClearingInstructions struct{ message.NumInGroupValue }

func (f NoClearingInstructions) Tag() fix.Tag { return tag.NoClearingInstructions }

type NoCollInquiryQualifier struct{ message.NumInGroupValue }

func (f NoCollInquiryQualifier) Tag() fix.Tag { return tag.NoCollInquiryQualifier }

type NoCompIDs struct{ message.NumInGroupValue }

func (f NoCompIDs) Tag() fix.Tag { return tag.NoCompIDs }

type NoComplexEventDates struct{ message.NumInGroupValue }

func (f NoComplexEventDates) Tag() fix.Tag { return tag.NoComplexEventDates }

type NoComplexEventTimes struct{ message.NumInGroupValue }

func (f NoComplexEventTimes) Tag() fix.Tag { return tag.NoComplexEventTimes }

type NoComplexEvents struct{ message.NumInGroupValue }

func (f NoComplexEvents) Tag() fix.Tag { return tag.NoComplexEvents }

type NoContAmts struct{ message.NumInGroupValue }

func (f NoContAmts) Tag() fix.Tag { return tag.NoContAmts }

type NoContextPartyIDs struct{ message.NumInGroupValue }

func (f NoContextPartyIDs) Tag() fix.Tag { return tag.NoContextPartyIDs }

type NoContextPartySubIDs struct{ message.NumInGroupValue }

func (f NoContextPartySubIDs) Tag() fix.Tag { return tag.NoContextPartySubIDs }

type NoContraBrokers struct{ message.NumInGroupValue }

func (f NoContraBrokers) Tag() fix.Tag { return tag.NoContraBrokers }

type NoDates struct{ message.IntValue }

func (f NoDates) Tag() fix.Tag { return tag.NoDates }

type NoDerivativeEvents struct{ message.NumInGroupValue }

func (f NoDerivativeEvents) Tag() fix.Tag { return tag.NoDerivativeEvents }

type NoDerivativeInstrAttrib struct{ message.NumInGroupValue }

func (f NoDerivativeInstrAttrib) Tag() fix.Tag { return tag.NoDerivativeInstrAttrib }

type NoDerivativeInstrumentParties struct{ message.NumInGroupValue }

func (f NoDerivativeInstrumentParties) Tag() fix.Tag { return tag.NoDerivativeInstrumentParties }

type NoDerivativeInstrumentPartySubIDs struct{ message.NumInGroupValue }

func (f NoDerivativeInstrumentPartySubIDs) Tag() fix.Tag { return tag.NoDerivativeInstrumentPartySubIDs }

type NoDerivativeSecurityAltID struct{ message.NumInGroupValue }

func (f NoDerivativeSecurityAltID) Tag() fix.Tag { return tag.NoDerivativeSecurityAltID }

type NoDistribInsts struct{ message.NumInGroupValue }

func (f NoDistribInsts) Tag() fix.Tag { return tag.NoDistribInsts }

type NoDlvyInst struct{ message.NumInGroupValue }

func (f NoDlvyInst) Tag() fix.Tag { return tag.NoDlvyInst }

type NoEvents struct{ message.NumInGroupValue }

func (f NoEvents) Tag() fix.Tag { return tag.NoEvents }

type NoExecInstRules struct{ message.NumInGroupValue }

func (f NoExecInstRules) Tag() fix.Tag { return tag.NoExecInstRules }

type NoExecs struct{ message.NumInGroupValue }

func (f NoExecs) Tag() fix.Tag { return tag.NoExecs }

type NoExpiration struct{ message.NumInGroupValue }

func (f NoExpiration) Tag() fix.Tag { return tag.NoExpiration }

type NoFills struct{ message.NumInGroupValue }

func (f NoFills) Tag() fix.Tag { return tag.NoFills }

type NoHops struct{ message.NumInGroupValue }

func (f NoHops) Tag() fix.Tag { return tag.NoHops }

type NoIOIQualifiers struct{ message.NumInGroupValue }

func (f NoIOIQualifiers) Tag() fix.Tag { return tag.NoIOIQualifiers }

type NoInstrAttrib struct{ message.NumInGroupValue }

func (f NoInstrAttrib) Tag() fix.Tag { return tag.NoInstrAttrib }

type NoInstrumentParties struct{ message.NumInGroupValue }

func (f NoInstrumentParties) Tag() fix.Tag { return tag.NoInstrumentParties }

type NoInstrumentPartySubIDs struct{ message.NumInGroupValue }

func (f NoInstrumentPartySubIDs) Tag() fix.Tag { return tag.NoInstrumentPartySubIDs }

type NoLegAllocs struct{ message.NumInGroupValue }

func (f NoLegAllocs) Tag() fix.Tag { return tag.NoLegAllocs }

type NoLegSecurityAltID struct{ message.StringValue }

func (f NoLegSecurityAltID) Tag() fix.Tag { return tag.NoLegSecurityAltID }

type NoLegStipulations struct{ message.NumInGroupValue }

func (f NoLegStipulations) Tag() fix.Tag { return tag.NoLegStipulations }

type NoLegs struct{ message.NumInGroupValue }

func (f NoLegs) Tag() fix.Tag { return tag.NoLegs }

type NoLinesOfText struct{ message.NumInGroupValue }

func (f NoLinesOfText) Tag() fix.Tag { return tag.NoLinesOfText }

type NoLotTypeRules struct{ message.NumInGroupValue }

func (f NoLotTypeRules) Tag() fix.Tag { return tag.NoLotTypeRules }

type NoMDEntries struct{ message.NumInGroupValue }

func (f NoMDEntries) Tag() fix.Tag { return tag.NoMDEntries }

type NoMDEntryTypes struct{ message.NumInGroupValue }

func (f NoMDEntryTypes) Tag() fix.Tag { return tag.NoMDEntryTypes }

type NoMDFeedTypes struct{ message.NumInGroupValue }

func (f NoMDFeedTypes) Tag() fix.Tag { return tag.NoMDFeedTypes }

type NoMarketSegments struct{ message.NumInGroupValue }

func (f NoMarketSegments) Tag() fix.Tag { return tag.NoMarketSegments }

type NoMatchRules struct{ message.NumInGroupValue }

func (f NoMatchRules) Tag() fix.Tag { return tag.NoMatchRules }

type NoMaturityRules struct{ message.NumInGroupValue }

func (f NoMaturityRules) Tag() fix.Tag { return tag.NoMaturityRules }

type NoMiscFees struct{ message.NumInGroupValue }

func (f NoMiscFees) Tag() fix.Tag { return tag.NoMiscFees }

type NoMsgTypes struct{ message.NumInGroupValue }

func (f NoMsgTypes) Tag() fix.Tag { return tag.NoMsgTypes }

type NoNested2PartyIDs struct{ message.NumInGroupValue }

func (f NoNested2PartyIDs) Tag() fix.Tag { return tag.NoNested2PartyIDs }

type NoNested2PartySubIDs struct{ message.NumInGroupValue }

func (f NoNested2PartySubIDs) Tag() fix.Tag { return tag.NoNested2PartySubIDs }

type NoNested3PartyIDs struct{ message.NumInGroupValue }

func (f NoNested3PartyIDs) Tag() fix.Tag { return tag.NoNested3PartyIDs }

type NoNested3PartySubIDs struct{ message.NumInGroupValue }

func (f NoNested3PartySubIDs) Tag() fix.Tag { return tag.NoNested3PartySubIDs }

type NoNested4PartyIDs struct{ message.NumInGroupValue }

func (f NoNested4PartyIDs) Tag() fix.Tag { return tag.NoNested4PartyIDs }

type NoNested4PartySubIDs struct{ message.NumInGroupValue }

func (f NoNested4PartySubIDs) Tag() fix.Tag { return tag.NoNested4PartySubIDs }

type NoNestedInstrAttrib struct{ message.NumInGroupValue }

func (f NoNestedInstrAttrib) Tag() fix.Tag { return tag.NoNestedInstrAttrib }

type NoNestedPartyIDs struct{ message.NumInGroupValue }

func (f NoNestedPartyIDs) Tag() fix.Tag { return tag.NoNestedPartyIDs }

type NoNestedPartySubIDs struct{ message.NumInGroupValue }

func (f NoNestedPartySubIDs) Tag() fix.Tag { return tag.NoNestedPartySubIDs }

type NoNewsRefIDs struct{ message.NumInGroupValue }

func (f NoNewsRefIDs) Tag() fix.Tag { return tag.NoNewsRefIDs }

type NoNotAffectedOrders struct{ message.NumInGroupValue }

func (f NoNotAffectedOrders) Tag() fix.Tag { return tag.NoNotAffectedOrders }

type NoOfLegUnderlyings struct{ message.NumInGroupValue }

func (f NoOfLegUnderlyings) Tag() fix.Tag { return tag.NoOfLegUnderlyings }

type NoOfSecSizes struct{ message.NumInGroupValue }

func (f NoOfSecSizes) Tag() fix.Tag { return tag.NoOfSecSizes }

type NoOrdTypeRules struct{ message.NumInGroupValue }

func (f NoOrdTypeRules) Tag() fix.Tag { return tag.NoOrdTypeRules }

type NoOrders struct{ message.NumInGroupValue }

func (f NoOrders) Tag() fix.Tag { return tag.NoOrders }

type NoPartyAltIDs struct{ message.NumInGroupValue }

func (f NoPartyAltIDs) Tag() fix.Tag { return tag.NoPartyAltIDs }

type NoPartyAltSubIDs struct{ message.NumInGroupValue }

func (f NoPartyAltSubIDs) Tag() fix.Tag { return tag.NoPartyAltSubIDs }

type NoPartyIDs struct{ message.NumInGroupValue }

func (f NoPartyIDs) Tag() fix.Tag { return tag.NoPartyIDs }

type NoPartyList struct{ message.NumInGroupValue }

func (f NoPartyList) Tag() fix.Tag { return tag.NoPartyList }

type NoPartyListResponseTypes struct{ message.NumInGroupValue }

func (f NoPartyListResponseTypes) Tag() fix.Tag { return tag.NoPartyListResponseTypes }

type NoPartyRelationships struct{ message.NumInGroupValue }

func (f NoPartyRelationships) Tag() fix.Tag { return tag.NoPartyRelationships }

type NoPartySubIDs struct{ message.NumInGroupValue }

func (f NoPartySubIDs) Tag() fix.Tag { return tag.NoPartySubIDs }

type NoPosAmt struct{ message.NumInGroupValue }

func (f NoPosAmt) Tag() fix.Tag { return tag.NoPosAmt }

type NoPositions struct{ message.NumInGroupValue }

func (f NoPositions) Tag() fix.Tag { return tag.NoPositions }

type NoQuoteEntries struct{ message.NumInGroupValue }

func (f NoQuoteEntries) Tag() fix.Tag { return tag.NoQuoteEntries }

type NoQuoteQualifiers struct{ message.NumInGroupValue }

func (f NoQuoteQualifiers) Tag() fix.Tag { return tag.NoQuoteQualifiers }

type NoQuoteSets struct{ message.NumInGroupValue }

func (f NoQuoteSets) Tag() fix.Tag { return tag.NoQuoteSets }

type NoRateSources struct{ message.NumInGroupValue }

func (f NoRateSources) Tag() fix.Tag { return tag.NoRateSources }

type NoRegistDtls struct{ message.NumInGroupValue }

func (f NoRegistDtls) Tag() fix.Tag { return tag.NoRegistDtls }

type NoRelatedContextPartyIDs struct{ message.NumInGroupValue }

func (f NoRelatedContextPartyIDs) Tag() fix.Tag { return tag.NoRelatedContextPartyIDs }

type NoRelatedContextPartySubIDs struct{ message.NumInGroupValue }

func (f NoRelatedContextPartySubIDs) Tag() fix.Tag { return tag.NoRelatedContextPartySubIDs }

type NoRelatedPartyAltIDs struct{ message.NumInGroupValue }

func (f NoRelatedPartyAltIDs) Tag() fix.Tag { return tag.NoRelatedPartyAltIDs }

type NoRelatedPartyAltSubIDs struct{ message.NumInGroupValue }

func (f NoRelatedPartyAltSubIDs) Tag() fix.Tag { return tag.NoRelatedPartyAltSubIDs }

type NoRelatedPartyIDs struct{ message.NumInGroupValue }

func (f NoRelatedPartyIDs) Tag() fix.Tag { return tag.NoRelatedPartyIDs }

type NoRelatedPartySubIDs struct{ message.NumInGroupValue }

func (f NoRelatedPartySubIDs) Tag() fix.Tag { return tag.NoRelatedPartySubIDs }

type NoRelatedSym struct{ message.NumInGroupValue }

func (f NoRelatedSym) Tag() fix.Tag { return tag.NoRelatedSym }

type NoRelationshipRiskInstruments struct{ message.NumInGroupValue }

func (f NoRelationshipRiskInstruments) Tag() fix.Tag { return tag.NoRelationshipRiskInstruments }

type NoRelationshipRiskLimits struct{ message.NumInGroupValue }

func (f NoRelationshipRiskLimits) Tag() fix.Tag { return tag.NoRelationshipRiskLimits }

type NoRelationshipRiskSecurityAltID struct{ message.NumInGroupValue }

func (f NoRelationshipRiskSecurityAltID) Tag() fix.Tag { return tag.NoRelationshipRiskSecurityAltID }

type NoRelationshipRiskWarningLevels struct{ message.NumInGroupValue }

func (f NoRelationshipRiskWarningLevels) Tag() fix.Tag { return tag.NoRelationshipRiskWarningLevels }

type NoRequestedPartyRoles struct{ message.NumInGroupValue }

func (f NoRequestedPartyRoles) Tag() fix.Tag { return tag.NoRequestedPartyRoles }

type NoRiskInstruments struct{ message.NumInGroupValue }

func (f NoRiskInstruments) Tag() fix.Tag { return tag.NoRiskInstruments }

type NoRiskLimits struct{ message.NumInGroupValue }

func (f NoRiskLimits) Tag() fix.Tag { return tag.NoRiskLimits }

type NoRiskSecurityAltID struct{ message.NumInGroupValue }

func (f NoRiskSecurityAltID) Tag() fix.Tag { return tag.NoRiskSecurityAltID }

type NoRiskWarningLevels struct{ message.NumInGroupValue }

func (f NoRiskWarningLevels) Tag() fix.Tag { return tag.NoRiskWarningLevels }

type NoRootPartyIDs struct{ message.NumInGroupValue }

func (f NoRootPartyIDs) Tag() fix.Tag { return tag.NoRootPartyIDs }

type NoRootPartySubIDs struct{ message.NumInGroupValue }

func (f NoRootPartySubIDs) Tag() fix.Tag { return tag.NoRootPartySubIDs }

type NoRoutingIDs struct{ message.NumInGroupValue }

func (f NoRoutingIDs) Tag() fix.Tag { return tag.NoRoutingIDs }

type NoRpts struct{ message.IntValue }

func (f NoRpts) Tag() fix.Tag { return tag.NoRpts }

type NoSecurityAltID struct{ message.NumInGroupValue }

func (f NoSecurityAltID) Tag() fix.Tag { return tag.NoSecurityAltID }

type NoSecurityTypes struct{ message.NumInGroupValue }

func (f NoSecurityTypes) Tag() fix.Tag { return tag.NoSecurityTypes }

type NoSettlDetails struct{ message.NumInGroupValue }

func (f NoSettlDetails) Tag() fix.Tag { return tag.NoSettlDetails }

type NoSettlInst struct{ message.NumInGroupValue }

func (f NoSettlInst) Tag() fix.Tag { return tag.NoSettlInst }

type NoSettlOblig struct{ message.NumInGroupValue }

func (f NoSettlOblig) Tag() fix.Tag { return tag.NoSettlOblig }

type NoSettlPartyIDs struct{ message.NumInGroupValue }

func (f NoSettlPartyIDs) Tag() fix.Tag { return tag.NoSettlPartyIDs }

type NoSettlPartySubIDs struct{ message.NumInGroupValue }

func (f NoSettlPartySubIDs) Tag() fix.Tag { return tag.NoSettlPartySubIDs }

type NoSideTrdRegTS struct{ message.NumInGroupValue }

func (f NoSideTrdRegTS) Tag() fix.Tag { return tag.NoSideTrdRegTS }

type NoSides struct{ message.NumInGroupValue }

func (f NoSides) Tag() fix.Tag { return tag.NoSides }

type NoStatsIndicators struct{ message.NumInGroupValue }

func (f NoStatsIndicators) Tag() fix.Tag { return tag.NoStatsIndicators }

type NoStipulations struct{ message.NumInGroupValue }

func (f NoStipulations) Tag() fix.Tag { return tag.NoStipulations }

type NoStrategyParameters struct{ message.NumInGroupValue }

func (f NoStrategyParameters) Tag() fix.Tag { return tag.NoStrategyParameters }

type NoStrikeRules struct{ message.NumInGroupValue }

func (f NoStrikeRules) Tag() fix.Tag { return tag.NoStrikeRules }

type NoStrikes struct{ message.NumInGroupValue }

func (f NoStrikes) Tag() fix.Tag { return tag.NoStrikes }

type NoTargetPartyIDs struct{ message.NumInGroupValue }

func (f NoTargetPartyIDs) Tag() fix.Tag { return tag.NoTargetPartyIDs }

type NoTickRules struct{ message.NumInGroupValue }

func (f NoTickRules) Tag() fix.Tag { return tag.NoTickRules }

type NoTimeInForceRules struct{ message.NumInGroupValue }

func (f NoTimeInForceRules) Tag() fix.Tag { return tag.NoTimeInForceRules }

type NoTrades struct{ message.NumInGroupValue }

func (f NoTrades) Tag() fix.Tag { return tag.NoTrades }

type NoTradingSessionRules struct{ message.NumInGroupValue }

func (f NoTradingSessionRules) Tag() fix.Tag { return tag.NoTradingSessionRules }

type NoTradingSessions struct{ message.NumInGroupValue }

func (f NoTradingSessions) Tag() fix.Tag { return tag.NoTradingSessions }

type NoTrdRegTimestamps struct{ message.NumInGroupValue }

func (f NoTrdRegTimestamps) Tag() fix.Tag { return tag.NoTrdRegTimestamps }

type NoTrdRepIndicators struct{ message.NumInGroupValue }

func (f NoTrdRepIndicators) Tag() fix.Tag { return tag.NoTrdRepIndicators }

type NoUnderlyingAmounts struct{ message.NumInGroupValue }

func (f NoUnderlyingAmounts) Tag() fix.Tag { return tag.NoUnderlyingAmounts }

type NoUnderlyingLegSecurityAltID struct{ message.NumInGroupValue }

func (f NoUnderlyingLegSecurityAltID) Tag() fix.Tag { return tag.NoUnderlyingLegSecurityAltID }

type NoUnderlyingSecurityAltID struct{ message.NumInGroupValue }

func (f NoUnderlyingSecurityAltID) Tag() fix.Tag { return tag.NoUnderlyingSecurityAltID }

type NoUnderlyingStips struct{ message.NumInGroupValue }

func (f NoUnderlyingStips) Tag() fix.Tag { return tag.NoUnderlyingStips }

type NoUnderlyings struct{ message.NumInGroupValue }

func (f NoUnderlyings) Tag() fix.Tag { return tag.NoUnderlyings }

type NoUndlyInstrumentParties struct{ message.NumInGroupValue }

func (f NoUndlyInstrumentParties) Tag() fix.Tag { return tag.NoUndlyInstrumentParties }

type NoUndlyInstrumentPartySubIDs struct{ message.NumInGroupValue }

func (f NoUndlyInstrumentPartySubIDs) Tag() fix.Tag { return tag.NoUndlyInstrumentPartySubIDs }

type NotAffOrigClOrdID struct{ message.StringValue }

func (f NotAffOrigClOrdID) Tag() fix.Tag { return tag.NotAffOrigClOrdID }

type NotAffectedOrderID struct{ message.StringValue }

func (f NotAffectedOrderID) Tag() fix.Tag { return tag.NotAffectedOrderID }

type NotifyBrokerOfCredit struct{ message.BooleanValue }

func (f NotifyBrokerOfCredit) Tag() fix.Tag { return tag.NotifyBrokerOfCredit }

type NotionalPercentageOutstanding struct{ message.PercentageValue }

func (f NotionalPercentageOutstanding) Tag() fix.Tag { return tag.NotionalPercentageOutstanding }

type NumBidders struct{ message.IntValue }

func (f NumBidders) Tag() fix.Tag { return tag.NumBidders }

type NumDaysInterest struct{ message.IntValue }

func (f NumDaysInterest) Tag() fix.Tag { return tag.NumDaysInterest }

type NumTickets struct{ message.IntValue }

func (f NumTickets) Tag() fix.Tag { return tag.NumTickets }

type NumberOfOrders struct{ message.IntValue }

func (f NumberOfOrders) Tag() fix.Tag { return tag.NumberOfOrders }

type OddLot struct{ message.BooleanValue }

func (f OddLot) Tag() fix.Tag { return tag.OddLot }

type OfferForwardPoints struct{ message.PriceOffsetValue }

func (f OfferForwardPoints) Tag() fix.Tag { return tag.OfferForwardPoints }

type OfferForwardPoints2 struct{ message.PriceOffsetValue }

func (f OfferForwardPoints2) Tag() fix.Tag { return tag.OfferForwardPoints2 }

type OfferPx struct{ message.PriceValue }

func (f OfferPx) Tag() fix.Tag { return tag.OfferPx }

type OfferSize struct{ message.QtyValue }

func (f OfferSize) Tag() fix.Tag { return tag.OfferSize }

type OfferSpotRate struct{ message.PriceValue }

func (f OfferSpotRate) Tag() fix.Tag { return tag.OfferSpotRate }

type OfferSwapPoints struct{ message.PriceOffsetValue }

func (f OfferSwapPoints) Tag() fix.Tag { return tag.OfferSwapPoints }

type OfferYield struct{ message.PercentageValue }

func (f OfferYield) Tag() fix.Tag { return tag.OfferYield }

type OnBehalfOfCompID struct{ message.StringValue }

func (f OnBehalfOfCompID) Tag() fix.Tag { return tag.OnBehalfOfCompID }

type OnBehalfOfLocationID struct{ message.StringValue }

func (f OnBehalfOfLocationID) Tag() fix.Tag { return tag.OnBehalfOfLocationID }

type OnBehalfOfSendingTime struct{ message.UTCTimestampValue }

func (f OnBehalfOfSendingTime) Tag() fix.Tag { return tag.OnBehalfOfSendingTime }

type OnBehalfOfSubID struct{ message.StringValue }

func (f OnBehalfOfSubID) Tag() fix.Tag { return tag.OnBehalfOfSubID }

type OpenClose struct{ message.CharValue }

func (f OpenClose) Tag() fix.Tag { return tag.OpenClose }

type OpenCloseSettlFlag struct{ message.MultipleCharValue }

func (f OpenCloseSettlFlag) Tag() fix.Tag { return tag.OpenCloseSettlFlag }

type OpenCloseSettleFlag struct{ message.MultipleStringValue }

func (f OpenCloseSettleFlag) Tag() fix.Tag { return tag.OpenCloseSettleFlag }

type OpenInterest struct{ message.AmtValue }

func (f OpenInterest) Tag() fix.Tag { return tag.OpenInterest }

type OptAttribute struct{ message.CharValue }

func (f OptAttribute) Tag() fix.Tag { return tag.OptAttribute }

type OptPayAmount struct{ message.AmtValue }

func (f OptPayAmount) Tag() fix.Tag { return tag.OptPayAmount }

type OptPayoutAmount struct{ message.AmtValue }

func (f OptPayoutAmount) Tag() fix.Tag { return tag.OptPayoutAmount }

type OptPayoutType struct{ message.IntValue }

func (f OptPayoutType) Tag() fix.Tag { return tag.OptPayoutType }

type OrdRejReason struct{ message.IntValue }

func (f OrdRejReason) Tag() fix.Tag { return tag.OrdRejReason }

type OrdStatus struct{ message.CharValue }

func (f OrdStatus) Tag() fix.Tag { return tag.OrdStatus }

type OrdStatusReqID struct{ message.StringValue }

func (f OrdStatusReqID) Tag() fix.Tag { return tag.OrdStatusReqID }

type OrdType struct{ message.CharValue }

func (f OrdType) Tag() fix.Tag { return tag.OrdType }

type OrderAvgPx struct{ message.PriceValue }

func (f OrderAvgPx) Tag() fix.Tag { return tag.OrderAvgPx }

type OrderBookingQty struct{ message.QtyValue }

func (f OrderBookingQty) Tag() fix.Tag { return tag.OrderBookingQty }

type OrderCapacity struct{ message.CharValue }

func (f OrderCapacity) Tag() fix.Tag { return tag.OrderCapacity }

type OrderCapacityQty struct{ message.QtyValue }

func (f OrderCapacityQty) Tag() fix.Tag { return tag.OrderCapacityQty }

type OrderCategory struct{ message.CharValue }

func (f OrderCategory) Tag() fix.Tag { return tag.OrderCategory }

type OrderDelay struct{ message.IntValue }

func (f OrderDelay) Tag() fix.Tag { return tag.OrderDelay }

type OrderDelayUnit struct{ message.IntValue }

func (f OrderDelayUnit) Tag() fix.Tag { return tag.OrderDelayUnit }

type OrderHandlingInstSource struct{ message.IntValue }

func (f OrderHandlingInstSource) Tag() fix.Tag { return tag.OrderHandlingInstSource }

type OrderID struct{ message.StringValue }

func (f OrderID) Tag() fix.Tag { return tag.OrderID }

type OrderInputDevice struct{ message.StringValue }

func (f OrderInputDevice) Tag() fix.Tag { return tag.OrderInputDevice }

type OrderPercent struct{ message.PercentageValue }

func (f OrderPercent) Tag() fix.Tag { return tag.OrderPercent }

type OrderQty struct{ message.QtyValue }

func (f OrderQty) Tag() fix.Tag { return tag.OrderQty }

type OrderQty2 struct{ message.QtyValue }

func (f OrderQty2) Tag() fix.Tag { return tag.OrderQty2 }

type OrderRestrictions struct{ message.MultipleCharValue }

func (f OrderRestrictions) Tag() fix.Tag { return tag.OrderRestrictions }

type OrigClOrdID struct{ message.StringValue }

func (f OrigClOrdID) Tag() fix.Tag { return tag.OrigClOrdID }

type OrigCrossID struct{ message.StringValue }

func (f OrigCrossID) Tag() fix.Tag { return tag.OrigCrossID }

type OrigCustOrderCapacity struct{ message.IntValue }

func (f OrigCustOrderCapacity) Tag() fix.Tag { return tag.OrigCustOrderCapacity }

type OrigOrdModTime struct{ message.UTCTimestampValue }

func (f OrigOrdModTime) Tag() fix.Tag { return tag.OrigOrdModTime }

type OrigPosReqRefID struct{ message.StringValue }

func (f OrigPosReqRefID) Tag() fix.Tag { return tag.OrigPosReqRefID }

type OrigSecondaryTradeID struct{ message.StringValue }

func (f OrigSecondaryTradeID) Tag() fix.Tag { return tag.OrigSecondaryTradeID }

type OrigSendingTime struct{ message.UTCTimestampValue }

func (f OrigSendingTime) Tag() fix.Tag { return tag.OrigSendingTime }

type OrigTime struct{ message.UTCTimestampValue }

func (f OrigTime) Tag() fix.Tag { return tag.OrigTime }

type OrigTradeDate struct{ message.LocalMktDateValue }

func (f OrigTradeDate) Tag() fix.Tag { return tag.OrigTradeDate }

type OrigTradeHandlingInstr struct{ message.CharValue }

func (f OrigTradeHandlingInstr) Tag() fix.Tag { return tag.OrigTradeHandlingInstr }

type OrigTradeID struct{ message.StringValue }

func (f OrigTradeID) Tag() fix.Tag { return tag.OrigTradeID }

type OriginalNotionalPercentageOutstanding struct{ message.PercentageValue }

func (f OriginalNotionalPercentageOutstanding) Tag() fix.Tag {
	return tag.OriginalNotionalPercentageOutstanding
}

type OutMainCntryUIndex struct{ message.AmtValue }

func (f OutMainCntryUIndex) Tag() fix.Tag { return tag.OutMainCntryUIndex }

type OutsideIndexPct struct{ message.PercentageValue }

func (f OutsideIndexPct) Tag() fix.Tag { return tag.OutsideIndexPct }

type OwnerType struct{ message.IntValue }

func (f OwnerType) Tag() fix.Tag { return tag.OwnerType }

type OwnershipType struct{ message.CharValue }

func (f OwnershipType) Tag() fix.Tag { return tag.OwnershipType }

type ParentMktSegmID struct{ message.StringValue }

func (f ParentMktSegmID) Tag() fix.Tag { return tag.ParentMktSegmID }

type ParticipationRate struct{ message.PercentageValue }

func (f ParticipationRate) Tag() fix.Tag { return tag.ParticipationRate }

type PartyAltID struct{ message.StringValue }

func (f PartyAltID) Tag() fix.Tag { return tag.PartyAltID }

type PartyAltIDSource struct{ message.CharValue }

func (f PartyAltIDSource) Tag() fix.Tag { return tag.PartyAltIDSource }

type PartyAltSubID struct{ message.StringValue }

func (f PartyAltSubID) Tag() fix.Tag { return tag.PartyAltSubID }

type PartyAltSubIDType struct{ message.IntValue }

func (f PartyAltSubIDType) Tag() fix.Tag { return tag.PartyAltSubIDType }

type PartyDetailsListReportID struct{ message.StringValue }

func (f PartyDetailsListReportID) Tag() fix.Tag { return tag.PartyDetailsListReportID }

type PartyDetailsListRequestID struct{ message.StringValue }

func (f PartyDetailsListRequestID) Tag() fix.Tag { return tag.PartyDetailsListRequestID }

type PartyDetailsRequestResult struct{ message.IntValue }

func (f PartyDetailsRequestResult) Tag() fix.Tag { return tag.PartyDetailsRequestResult }

type PartyID struct{ message.StringValue }

func (f PartyID) Tag() fix.Tag { return tag.PartyID }

type PartyIDSource struct{ message.CharValue }

func (f PartyIDSource) Tag() fix.Tag { return tag.PartyIDSource }

type PartyListResponseType struct{ message.IntValue }

func (f PartyListResponseType) Tag() fix.Tag { return tag.PartyListResponseType }

type PartyRelationship struct{ message.IntValue }

func (f PartyRelationship) Tag() fix.Tag { return tag.PartyRelationship }

type PartyRole struct{ message.IntValue }

func (f PartyRole) Tag() fix.Tag { return tag.PartyRole }

type PartySubID struct{ message.StringValue }

func (f PartySubID) Tag() fix.Tag { return tag.PartySubID }

type PartySubIDType struct{ message.IntValue }

func (f PartySubIDType) Tag() fix.Tag { return tag.PartySubIDType }

type Password struct{ message.StringValue }

func (f Password) Tag() fix.Tag { return tag.Password }

type PaymentDate struct{ message.LocalMktDateValue }

func (f PaymentDate) Tag() fix.Tag { return tag.PaymentDate }

type PaymentMethod struct{ message.IntValue }

func (f PaymentMethod) Tag() fix.Tag { return tag.PaymentMethod }

type PaymentRef struct{ message.StringValue }

func (f PaymentRef) Tag() fix.Tag { return tag.PaymentRef }

type PaymentRemitterID struct{ message.StringValue }

func (f PaymentRemitterID) Tag() fix.Tag { return tag.PaymentRemitterID }

type PctAtRisk struct{ message.PercentageValue }

func (f PctAtRisk) Tag() fix.Tag { return tag.PctAtRisk }

type PegDifference struct{ message.PriceOffsetValue }

func (f PegDifference) Tag() fix.Tag { return tag.PegDifference }

type PegLimitType struct{ message.IntValue }

func (f PegLimitType) Tag() fix.Tag { return tag.PegLimitType }

type PegMoveType struct{ message.IntValue }

func (f PegMoveType) Tag() fix.Tag { return tag.PegMoveType }

type PegOffsetType struct{ message.IntValue }

func (f PegOffsetType) Tag() fix.Tag { return tag.PegOffsetType }

type PegOffsetValue struct{ message.FloatValue }

func (f PegOffsetValue) Tag() fix.Tag { return tag.PegOffsetValue }

type PegPriceType struct{ message.IntValue }

func (f PegPriceType) Tag() fix.Tag { return tag.PegPriceType }

type PegRoundDirection struct{ message.IntValue }

func (f PegRoundDirection) Tag() fix.Tag { return tag.PegRoundDirection }

type PegScope struct{ message.IntValue }

func (f PegScope) Tag() fix.Tag { return tag.PegScope }

type PegSecurityDesc struct{ message.StringValue }

func (f PegSecurityDesc) Tag() fix.Tag { return tag.PegSecurityDesc }

type PegSecurityID struct{ message.StringValue }

func (f PegSecurityID) Tag() fix.Tag { return tag.PegSecurityID }

type PegSecurityIDSource struct{ message.StringValue }

func (f PegSecurityIDSource) Tag() fix.Tag { return tag.PegSecurityIDSource }

type PegSymbol struct{ message.StringValue }

func (f PegSymbol) Tag() fix.Tag { return tag.PegSymbol }

type PeggedPrice struct{ message.PriceValue }

func (f PeggedPrice) Tag() fix.Tag { return tag.PeggedPrice }

type PeggedRefPrice struct{ message.PriceValue }

func (f PeggedRefPrice) Tag() fix.Tag { return tag.PeggedRefPrice }

type Pool struct{ message.StringValue }

func (f Pool) Tag() fix.Tag { return tag.Pool }

type PosAmt struct{ message.AmtValue }

func (f PosAmt) Tag() fix.Tag { return tag.PosAmt }

type PosAmtType struct{ message.StringValue }

func (f PosAmtType) Tag() fix.Tag { return tag.PosAmtType }

type PosMaintAction struct{ message.IntValue }

func (f PosMaintAction) Tag() fix.Tag { return tag.PosMaintAction }

type PosMaintResult struct{ message.IntValue }

func (f PosMaintResult) Tag() fix.Tag { return tag.PosMaintResult }

type PosMaintRptID struct{ message.StringValue }

func (f PosMaintRptID) Tag() fix.Tag { return tag.PosMaintRptID }

type PosMaintRptRefID struct{ message.StringValue }

func (f PosMaintRptRefID) Tag() fix.Tag { return tag.PosMaintRptRefID }

type PosMaintStatus struct{ message.IntValue }

func (f PosMaintStatus) Tag() fix.Tag { return tag.PosMaintStatus }

type PosQtyStatus struct{ message.IntValue }

func (f PosQtyStatus) Tag() fix.Tag { return tag.PosQtyStatus }

type PosReqID struct{ message.StringValue }

func (f PosReqID) Tag() fix.Tag { return tag.PosReqID }

type PosReqResult struct{ message.IntValue }

func (f PosReqResult) Tag() fix.Tag { return tag.PosReqResult }

type PosReqStatus struct{ message.IntValue }

func (f PosReqStatus) Tag() fix.Tag { return tag.PosReqStatus }

type PosReqType struct{ message.IntValue }

func (f PosReqType) Tag() fix.Tag { return tag.PosReqType }

type PosTransType struct{ message.IntValue }

func (f PosTransType) Tag() fix.Tag { return tag.PosTransType }

type PosType struct{ message.StringValue }

func (f PosType) Tag() fix.Tag { return tag.PosType }

type PositionCurrency struct{ message.StringValue }

func (f PositionCurrency) Tag() fix.Tag { return tag.PositionCurrency }

type PositionEffect struct{ message.CharValue }

func (f PositionEffect) Tag() fix.Tag { return tag.PositionEffect }

type PositionLimit struct{ message.IntValue }

func (f PositionLimit) Tag() fix.Tag { return tag.PositionLimit }

type PossDupFlag struct{ message.BooleanValue }

func (f PossDupFlag) Tag() fix.Tag { return tag.PossDupFlag }

type PossResend struct{ message.BooleanValue }

func (f PossResend) Tag() fix.Tag { return tag.PossResend }

type PreTradeAnonymity struct{ message.BooleanValue }

func (f PreTradeAnonymity) Tag() fix.Tag { return tag.PreTradeAnonymity }

type PreallocMethod struct{ message.CharValue }

func (f PreallocMethod) Tag() fix.Tag { return tag.PreallocMethod }

type PrevClosePx struct{ message.PriceValue }

func (f PrevClosePx) Tag() fix.Tag { return tag.PrevClosePx }

type PreviouslyReported struct{ message.BooleanValue }

func (f PreviouslyReported) Tag() fix.Tag { return tag.PreviouslyReported }

type Price struct{ message.PriceValue }

func (f Price) Tag() fix.Tag { return tag.Price }

type Price2 struct{ message.PriceValue }

func (f Price2) Tag() fix.Tag { return tag.Price2 }

type PriceDelta struct{ message.FloatValue }

func (f PriceDelta) Tag() fix.Tag { return tag.PriceDelta }

type PriceImprovement struct{ message.PriceOffsetValue }

func (f PriceImprovement) Tag() fix.Tag { return tag.PriceImprovement }

type PriceLimitType struct{ message.IntValue }

func (f PriceLimitType) Tag() fix.Tag { return tag.PriceLimitType }

type PriceProtectionScope struct{ message.CharValue }

func (f PriceProtectionScope) Tag() fix.Tag { return tag.PriceProtectionScope }

type PriceQuoteMethod struct{ message.StringValue }

func (f PriceQuoteMethod) Tag() fix.Tag { return tag.PriceQuoteMethod }

type PriceType struct{ message.IntValue }

func (f PriceType) Tag() fix.Tag { return tag.PriceType }

type PriceUnitOfMeasure struct{ message.StringValue }

func (f PriceUnitOfMeasure) Tag() fix.Tag { return tag.PriceUnitOfMeasure }

type PriceUnitOfMeasureQty struct{ message.QtyValue }

func (f PriceUnitOfMeasureQty) Tag() fix.Tag { return tag.PriceUnitOfMeasureQty }

type PriorSettlPrice struct{ message.PriceValue }

func (f PriorSettlPrice) Tag() fix.Tag { return tag.PriorSettlPrice }

type PriorSpreadIndicator struct{ message.BooleanValue }

func (f PriorSpreadIndicator) Tag() fix.Tag { return tag.PriorSpreadIndicator }

type PriorityIndicator struct{ message.IntValue }

func (f PriorityIndicator) Tag() fix.Tag { return tag.PriorityIndicator }

type PrivateQuote struct{ message.BooleanValue }

func (f PrivateQuote) Tag() fix.Tag { return tag.PrivateQuote }

type ProcessCode struct{ message.CharValue }

func (f ProcessCode) Tag() fix.Tag { return tag.ProcessCode }

type Product struct{ message.IntValue }

func (f Product) Tag() fix.Tag { return tag.Product }

type ProductComplex struct{ message.StringValue }

func (f ProductComplex) Tag() fix.Tag { return tag.ProductComplex }

type ProgPeriodInterval struct{ message.IntValue }

func (f ProgPeriodInterval) Tag() fix.Tag { return tag.ProgPeriodInterval }

type ProgRptReqs struct{ message.IntValue }

func (f ProgRptReqs) Tag() fix.Tag { return tag.ProgRptReqs }

type PublishTrdIndicator struct{ message.BooleanValue }

func (f PublishTrdIndicator) Tag() fix.Tag { return tag.PublishTrdIndicator }

type PutOrCall struct{ message.IntValue }

func (f PutOrCall) Tag() fix.Tag { return tag.PutOrCall }

type QtyType struct{ message.IntValue }

func (f QtyType) Tag() fix.Tag { return tag.QtyType }

type Quantity struct{ message.QtyValue }

func (f Quantity) Tag() fix.Tag { return tag.Quantity }

type QuantityDate struct{ message.LocalMktDateValue }

func (f QuantityDate) Tag() fix.Tag { return tag.QuantityDate }

type QuantityType struct{ message.IntValue }

func (f QuantityType) Tag() fix.Tag { return tag.QuantityType }

type QuoteAckStatus struct{ message.IntValue }

func (f QuoteAckStatus) Tag() fix.Tag { return tag.QuoteAckStatus }

type QuoteCancelType struct{ message.IntValue }

func (f QuoteCancelType) Tag() fix.Tag { return tag.QuoteCancelType }

type QuoteCondition struct{ message.MultipleStringValue }

func (f QuoteCondition) Tag() fix.Tag { return tag.QuoteCondition }

type QuoteEntryID struct{ message.StringValue }

func (f QuoteEntryID) Tag() fix.Tag { return tag.QuoteEntryID }

type QuoteEntryRejectReason struct{ message.IntValue }

func (f QuoteEntryRejectReason) Tag() fix.Tag { return tag.QuoteEntryRejectReason }

type QuoteEntryStatus struct{ message.IntValue }

func (f QuoteEntryStatus) Tag() fix.Tag { return tag.QuoteEntryStatus }

type QuoteID struct{ message.StringValue }

func (f QuoteID) Tag() fix.Tag { return tag.QuoteID }

type QuoteMsgID struct{ message.StringValue }

func (f QuoteMsgID) Tag() fix.Tag { return tag.QuoteMsgID }

type QuotePriceType struct{ message.IntValue }

func (f QuotePriceType) Tag() fix.Tag { return tag.QuotePriceType }

type QuoteQualifier struct{ message.CharValue }

func (f QuoteQualifier) Tag() fix.Tag { return tag.QuoteQualifier }

type QuoteRejectReason struct{ message.IntValue }

func (f QuoteRejectReason) Tag() fix.Tag { return tag.QuoteRejectReason }

type QuoteReqID struct{ message.StringValue }

func (f QuoteReqID) Tag() fix.Tag { return tag.QuoteReqID }

type QuoteRequestRejectReason struct{ message.IntValue }

func (f QuoteRequestRejectReason) Tag() fix.Tag { return tag.QuoteRequestRejectReason }

type QuoteRequestType struct{ message.IntValue }

func (f QuoteRequestType) Tag() fix.Tag { return tag.QuoteRequestType }

type QuoteRespID struct{ message.StringValue }

func (f QuoteRespID) Tag() fix.Tag { return tag.QuoteRespID }

type QuoteRespType struct{ message.IntValue }

func (f QuoteRespType) Tag() fix.Tag { return tag.QuoteRespType }

type QuoteResponseLevel struct{ message.IntValue }

func (f QuoteResponseLevel) Tag() fix.Tag { return tag.QuoteResponseLevel }

type QuoteSetID struct{ message.StringValue }

func (f QuoteSetID) Tag() fix.Tag { return tag.QuoteSetID }

type QuoteSetValidUntilTime struct{ message.UTCTimestampValue }

func (f QuoteSetValidUntilTime) Tag() fix.Tag { return tag.QuoteSetValidUntilTime }

type QuoteStatus struct{ message.IntValue }

func (f QuoteStatus) Tag() fix.Tag { return tag.QuoteStatus }

type QuoteStatusReqID struct{ message.StringValue }

func (f QuoteStatusReqID) Tag() fix.Tag { return tag.QuoteStatusReqID }

type QuoteType struct{ message.IntValue }

func (f QuoteType) Tag() fix.Tag { return tag.QuoteType }

type RFQReqID struct{ message.StringValue }

func (f RFQReqID) Tag() fix.Tag { return tag.RFQReqID }

type RateSource struct{ message.IntValue }

func (f RateSource) Tag() fix.Tag { return tag.RateSource }

type RateSourceType struct{ message.IntValue }

func (f RateSourceType) Tag() fix.Tag { return tag.RateSourceType }

type RatioQty struct{ message.QtyValue }

func (f RatioQty) Tag() fix.Tag { return tag.RatioQty }

type RawData struct{ message.DataValue }

func (f RawData) Tag() fix.Tag { return tag.RawData }

type RawDataLength struct{ message.LengthValue }

func (f RawDataLength) Tag() fix.Tag { return tag.RawDataLength }

type ReceivedDeptID struct{ message.StringValue }

func (f ReceivedDeptID) Tag() fix.Tag { return tag.ReceivedDeptID }

type RedemptionDate struct{ message.LocalMktDateValue }

func (f RedemptionDate) Tag() fix.Tag { return tag.RedemptionDate }

type RefAllocID struct{ message.StringValue }

func (f RefAllocID) Tag() fix.Tag { return tag.RefAllocID }

type RefApplExtID struct{ message.IntValue }

func (f RefApplExtID) Tag() fix.Tag { return tag.RefApplExtID }

type RefApplID struct{ message.StringValue }

func (f RefApplID) Tag() fix.Tag { return tag.RefApplID }

type RefApplLastSeqNum struct{ message.SeqNumValue }

func (f RefApplLastSeqNum) Tag() fix.Tag { return tag.RefApplLastSeqNum }

type RefApplReqID struct{ message.StringValue }

func (f RefApplReqID) Tag() fix.Tag { return tag.RefApplReqID }

type RefApplVerID struct{ message.StringValue }

func (f RefApplVerID) Tag() fix.Tag { return tag.RefApplVerID }

type RefCompID struct{ message.StringValue }

func (f RefCompID) Tag() fix.Tag { return tag.RefCompID }

type RefCstmApplVerID struct{ message.StringValue }

func (f RefCstmApplVerID) Tag() fix.Tag { return tag.RefCstmApplVerID }

type RefMsgType struct{ message.StringValue }

func (f RefMsgType) Tag() fix.Tag { return tag.RefMsgType }

type RefOrdIDReason struct{ message.IntValue }

func (f RefOrdIDReason) Tag() fix.Tag { return tag.RefOrdIDReason }

type RefOrderID struct{ message.StringValue }

func (f RefOrderID) Tag() fix.Tag { return tag.RefOrderID }

type RefOrderIDSource struct{ message.CharValue }

func (f RefOrderIDSource) Tag() fix.Tag { return tag.RefOrderIDSource }

type RefSeqNum struct{ message.SeqNumValue }

func (f RefSeqNum) Tag() fix.Tag { return tag.RefSeqNum }

type RefSubID struct{ message.StringValue }

func (f RefSubID) Tag() fix.Tag { return tag.RefSubID }

type RefTagID struct{ message.IntValue }

func (f RefTagID) Tag() fix.Tag { return tag.RefTagID }

type ReferencePage struct{ message.StringValue }

func (f ReferencePage) Tag() fix.Tag { return tag.ReferencePage }

type RefreshIndicator struct{ message.BooleanValue }

func (f RefreshIndicator) Tag() fix.Tag { return tag.RefreshIndicator }

type RefreshQty struct{ message.QtyValue }

func (f RefreshQty) Tag() fix.Tag { return tag.RefreshQty }

type RegistAcctType struct{ message.StringValue }

func (f RegistAcctType) Tag() fix.Tag { return tag.RegistAcctType }

type RegistDetls struct{ message.StringValue }

func (f RegistDetls) Tag() fix.Tag { return tag.RegistDetls }

type RegistDtls struct{ message.StringValue }

func (f RegistDtls) Tag() fix.Tag { return tag.RegistDtls }

type RegistEmail struct{ message.StringValue }

func (f RegistEmail) Tag() fix.Tag { return tag.RegistEmail }

type RegistID struct{ message.StringValue }

func (f RegistID) Tag() fix.Tag { return tag.RegistID }

type RegistRefID struct{ message.StringValue }

func (f RegistRefID) Tag() fix.Tag { return tag.RegistRefID }

type RegistRejReasonCode struct{ message.IntValue }

func (f RegistRejReasonCode) Tag() fix.Tag { return tag.RegistRejReasonCode }

type RegistRejReasonText struct{ message.StringValue }

func (f RegistRejReasonText) Tag() fix.Tag { return tag.RegistRejReasonText }

type RegistStatus struct{ message.CharValue }

func (f RegistStatus) Tag() fix.Tag { return tag.RegistStatus }

type RegistTransType struct{ message.CharValue }

func (f RegistTransType) Tag() fix.Tag { return tag.RegistTransType }

type RejectText struct{ message.StringValue }

func (f RejectText) Tag() fix.Tag { return tag.RejectText }

type RelSymTransactTime struct{ message.UTCTimestampValue }

func (f RelSymTransactTime) Tag() fix.Tag { return tag.RelSymTransactTime }

type RelatdSym struct{ message.StringValue }

func (f RelatdSym) Tag() fix.Tag { return tag.RelatdSym }

type RelatedContextPartyID struct{ message.StringValue }

func (f RelatedContextPartyID) Tag() fix.Tag { return tag.RelatedContextPartyID }

type RelatedContextPartyIDSource struct{ message.CharValue }

func (f RelatedContextPartyIDSource) Tag() fix.Tag { return tag.RelatedContextPartyIDSource }

type RelatedContextPartyRole struct{ message.IntValue }

func (f RelatedContextPartyRole) Tag() fix.Tag { return tag.RelatedContextPartyRole }

type RelatedContextPartySubID struct{ message.StringValue }

func (f RelatedContextPartySubID) Tag() fix.Tag { return tag.RelatedContextPartySubID }

type RelatedContextPartySubIDType struct{ message.IntValue }

func (f RelatedContextPartySubIDType) Tag() fix.Tag { return tag.RelatedContextPartySubIDType }

type RelatedPartyAltID struct{ message.StringValue }

func (f RelatedPartyAltID) Tag() fix.Tag { return tag.RelatedPartyAltID }

type RelatedPartyAltIDSource struct{ message.CharValue }

func (f RelatedPartyAltIDSource) Tag() fix.Tag { return tag.RelatedPartyAltIDSource }

type RelatedPartyAltSubID struct{ message.StringValue }

func (f RelatedPartyAltSubID) Tag() fix.Tag { return tag.RelatedPartyAltSubID }

type RelatedPartyAltSubIDType struct{ message.IntValue }

func (f RelatedPartyAltSubIDType) Tag() fix.Tag { return tag.RelatedPartyAltSubIDType }

type RelatedPartyID struct{ message.StringValue }

func (f RelatedPartyID) Tag() fix.Tag { return tag.RelatedPartyID }

type RelatedPartyIDSource struct{ message.CharValue }

func (f RelatedPartyIDSource) Tag() fix.Tag { return tag.RelatedPartyIDSource }

type RelatedPartyRole struct{ message.IntValue }

func (f RelatedPartyRole) Tag() fix.Tag { return tag.RelatedPartyRole }

type RelatedPartySubID struct{ message.StringValue }

func (f RelatedPartySubID) Tag() fix.Tag { return tag.RelatedPartySubID }

type RelatedPartySubIDType struct{ message.IntValue }

func (f RelatedPartySubIDType) Tag() fix.Tag { return tag.RelatedPartySubIDType }

type RelationshipRiskCFICode struct{ message.StringValue }

func (f RelationshipRiskCFICode) Tag() fix.Tag { return tag.RelationshipRiskCFICode }

type RelationshipRiskCouponRate struct{ message.PercentageValue }

func (f RelationshipRiskCouponRate) Tag() fix.Tag { return tag.RelationshipRiskCouponRate }

type RelationshipRiskEncodedSecurityDesc struct{ message.DataValue }

func (f RelationshipRiskEncodedSecurityDesc) Tag() fix.Tag {
	return tag.RelationshipRiskEncodedSecurityDesc
}

type RelationshipRiskEncodedSecurityDescLen struct{ message.LengthValue }

func (f RelationshipRiskEncodedSecurityDescLen) Tag() fix.Tag {
	return tag.RelationshipRiskEncodedSecurityDescLen
}

type RelationshipRiskFlexibleIndicator struct{ message.BooleanValue }

func (f RelationshipRiskFlexibleIndicator) Tag() fix.Tag { return tag.RelationshipRiskFlexibleIndicator }

type RelationshipRiskInstrumentMultiplier struct{ message.FloatValue }

func (f RelationshipRiskInstrumentMultiplier) Tag() fix.Tag {
	return tag.RelationshipRiskInstrumentMultiplier
}

type RelationshipRiskInstrumentOperator struct{ message.IntValue }

func (f RelationshipRiskInstrumentOperator) Tag() fix.Tag {
	return tag.RelationshipRiskInstrumentOperator
}

type RelationshipRiskInstrumentSettlType struct{ message.StringValue }

func (f RelationshipRiskInstrumentSettlType) Tag() fix.Tag {
	return tag.RelationshipRiskInstrumentSettlType
}

type RelationshipRiskLimitAmount struct{ message.AmtValue }

func (f RelationshipRiskLimitAmount) Tag() fix.Tag { return tag.RelationshipRiskLimitAmount }

type RelationshipRiskLimitCurrency struct{ message.CurrencyValue }

func (f RelationshipRiskLimitCurrency) Tag() fix.Tag { return tag.RelationshipRiskLimitCurrency }

type RelationshipRiskLimitPlatform struct{ message.StringValue }

func (f RelationshipRiskLimitPlatform) Tag() fix.Tag { return tag.RelationshipRiskLimitPlatform }

type RelationshipRiskLimitType struct{ message.IntValue }

func (f RelationshipRiskLimitType) Tag() fix.Tag { return tag.RelationshipRiskLimitType }

type RelationshipRiskMaturityMonthYear struct{ message.MonthYearValue }

func (f RelationshipRiskMaturityMonthYear) Tag() fix.Tag { return tag.RelationshipRiskMaturityMonthYear }

type RelationshipRiskMaturityTime struct{ message.TZTimeOnlyValue }

func (f RelationshipRiskMaturityTime) Tag() fix.Tag { return tag.RelationshipRiskMaturityTime }

type RelationshipRiskProduct struct{ message.IntValue }

func (f RelationshipRiskProduct) Tag() fix.Tag { return tag.RelationshipRiskProduct }

type RelationshipRiskProductComplex struct{ message.StringValue }

func (f RelationshipRiskProductComplex) Tag() fix.Tag { return tag.RelationshipRiskProductComplex }

type RelationshipRiskPutOrCall struct{ message.IntValue }

func (f RelationshipRiskPutOrCall) Tag() fix.Tag { return tag.RelationshipRiskPutOrCall }

type RelationshipRiskRestructuringType struct{ message.StringValue }

func (f RelationshipRiskRestructuringType) Tag() fix.Tag { return tag.RelationshipRiskRestructuringType }

type RelationshipRiskSecurityAltID struct{ message.StringValue }

func (f RelationshipRiskSecurityAltID) Tag() fix.Tag { return tag.RelationshipRiskSecurityAltID }

type RelationshipRiskSecurityAltIDSource struct{ message.StringValue }

func (f RelationshipRiskSecurityAltIDSource) Tag() fix.Tag {
	return tag.RelationshipRiskSecurityAltIDSource
}

type RelationshipRiskSecurityDesc struct{ message.StringValue }

func (f RelationshipRiskSecurityDesc) Tag() fix.Tag { return tag.RelationshipRiskSecurityDesc }

type RelationshipRiskSecurityExchange struct{ message.ExchangeValue }

func (f RelationshipRiskSecurityExchange) Tag() fix.Tag { return tag.RelationshipRiskSecurityExchange }

type RelationshipRiskSecurityGroup struct{ message.StringValue }

func (f RelationshipRiskSecurityGroup) Tag() fix.Tag { return tag.RelationshipRiskSecurityGroup }

type RelationshipRiskSecurityID struct{ message.StringValue }

func (f RelationshipRiskSecurityID) Tag() fix.Tag { return tag.RelationshipRiskSecurityID }

type RelationshipRiskSecurityIDSource struct{ message.StringValue }

func (f RelationshipRiskSecurityIDSource) Tag() fix.Tag { return tag.RelationshipRiskSecurityIDSource }

type RelationshipRiskSecuritySubType struct{ message.StringValue }

func (f RelationshipRiskSecuritySubType) Tag() fix.Tag { return tag.RelationshipRiskSecuritySubType }

type RelationshipRiskSecurityType struct{ message.StringValue }

func (f RelationshipRiskSecurityType) Tag() fix.Tag { return tag.RelationshipRiskSecurityType }

type RelationshipRiskSeniority struct{ message.StringValue }

func (f RelationshipRiskSeniority) Tag() fix.Tag { return tag.RelationshipRiskSeniority }

type RelationshipRiskSymbol struct{ message.StringValue }

func (f RelationshipRiskSymbol) Tag() fix.Tag { return tag.RelationshipRiskSymbol }

type RelationshipRiskSymbolSfx struct{ message.StringValue }

func (f RelationshipRiskSymbolSfx) Tag() fix.Tag { return tag.RelationshipRiskSymbolSfx }

type RelationshipRiskWarningLevelName struct{ message.StringValue }

func (f RelationshipRiskWarningLevelName) Tag() fix.Tag { return tag.RelationshipRiskWarningLevelName }

type RelationshipRiskWarningLevelPercent struct{ message.PercentageValue }

func (f RelationshipRiskWarningLevelPercent) Tag() fix.Tag {
	return tag.RelationshipRiskWarningLevelPercent
}

type RepoCollateralSecurityType struct{ message.IntValue }

func (f RepoCollateralSecurityType) Tag() fix.Tag { return tag.RepoCollateralSecurityType }

type ReportToExch struct{ message.BooleanValue }

func (f ReportToExch) Tag() fix.Tag { return tag.ReportToExch }

type ReportedPx struct{ message.PriceValue }

func (f ReportedPx) Tag() fix.Tag { return tag.ReportedPx }

type ReportedPxDiff struct{ message.BooleanValue }

func (f ReportedPxDiff) Tag() fix.Tag { return tag.ReportedPxDiff }

type RepurchaseRate struct{ message.PercentageValue }

func (f RepurchaseRate) Tag() fix.Tag { return tag.RepurchaseRate }

type RepurchaseTerm struct{ message.IntValue }

func (f RepurchaseTerm) Tag() fix.Tag { return tag.RepurchaseTerm }

type RequestedPartyRole struct{ message.IntValue }

func (f RequestedPartyRole) Tag() fix.Tag { return tag.RequestedPartyRole }

type ResetSeqNumFlag struct{ message.BooleanValue }

func (f ResetSeqNumFlag) Tag() fix.Tag { return tag.ResetSeqNumFlag }

type RespondentType struct{ message.IntValue }

func (f RespondentType) Tag() fix.Tag { return tag.RespondentType }

type ResponseDestination struct{ message.StringValue }

func (f ResponseDestination) Tag() fix.Tag { return tag.ResponseDestination }

type ResponseTransportType struct{ message.IntValue }

func (f ResponseTransportType) Tag() fix.Tag { return tag.ResponseTransportType }

type RestructuringType struct{ message.StringValue }

func (f RestructuringType) Tag() fix.Tag { return tag.RestructuringType }

type ReversalIndicator struct{ message.BooleanValue }

func (f ReversalIndicator) Tag() fix.Tag { return tag.ReversalIndicator }

type RiskCFICode struct{ message.StringValue }

func (f RiskCFICode) Tag() fix.Tag { return tag.RiskCFICode }

type RiskCouponRate struct{ message.PercentageValue }

func (f RiskCouponRate) Tag() fix.Tag { return tag.RiskCouponRate }

type RiskEncodedSecurityDesc struct{ message.DataValue }

func (f RiskEncodedSecurityDesc) Tag() fix.Tag { return tag.RiskEncodedSecurityDesc }

type RiskEncodedSecurityDescLen struct{ message.LengthValue }

func (f RiskEncodedSecurityDescLen) Tag() fix.Tag { return tag.RiskEncodedSecurityDescLen }

type RiskFlexibleIndicator struct{ message.BooleanValue }

func (f RiskFlexibleIndicator) Tag() fix.Tag { return tag.RiskFlexibleIndicator }

type RiskFreeRate struct{ message.FloatValue }

func (f RiskFreeRate) Tag() fix.Tag { return tag.RiskFreeRate }

type RiskInstrumentMultiplier struct{ message.FloatValue }

func (f RiskInstrumentMultiplier) Tag() fix.Tag { return tag.RiskInstrumentMultiplier }

type RiskInstrumentOperator struct{ message.IntValue }

func (f RiskInstrumentOperator) Tag() fix.Tag { return tag.RiskInstrumentOperator }

type RiskInstrumentSettlType struct{ message.StringValue }

func (f RiskInstrumentSettlType) Tag() fix.Tag { return tag.RiskInstrumentSettlType }

type RiskLimitAmount struct{ message.AmtValue }

func (f RiskLimitAmount) Tag() fix.Tag { return tag.RiskLimitAmount }

type RiskLimitCurrency struct{ message.CurrencyValue }

func (f RiskLimitCurrency) Tag() fix.Tag { return tag.RiskLimitCurrency }

type RiskLimitPlatform struct{ message.StringValue }

func (f RiskLimitPlatform) Tag() fix.Tag { return tag.RiskLimitPlatform }

type RiskLimitType struct{ message.IntValue }

func (f RiskLimitType) Tag() fix.Tag { return tag.RiskLimitType }

type RiskMaturityMonthYear struct{ message.MonthYearValue }

func (f RiskMaturityMonthYear) Tag() fix.Tag { return tag.RiskMaturityMonthYear }

type RiskMaturityTime struct{ message.TZTimeOnlyValue }

func (f RiskMaturityTime) Tag() fix.Tag { return tag.RiskMaturityTime }

type RiskProduct struct{ message.IntValue }

func (f RiskProduct) Tag() fix.Tag { return tag.RiskProduct }

type RiskProductComplex struct{ message.StringValue }

func (f RiskProductComplex) Tag() fix.Tag { return tag.RiskProductComplex }

type RiskPutOrCall struct{ message.IntValue }

func (f RiskPutOrCall) Tag() fix.Tag { return tag.RiskPutOrCall }

type RiskRestructuringType struct{ message.StringValue }

func (f RiskRestructuringType) Tag() fix.Tag { return tag.RiskRestructuringType }

type RiskSecurityAltID struct{ message.StringValue }

func (f RiskSecurityAltID) Tag() fix.Tag { return tag.RiskSecurityAltID }

type RiskSecurityAltIDSource struct{ message.StringValue }

func (f RiskSecurityAltIDSource) Tag() fix.Tag { return tag.RiskSecurityAltIDSource }

type RiskSecurityDesc struct{ message.StringValue }

func (f RiskSecurityDesc) Tag() fix.Tag { return tag.RiskSecurityDesc }

type RiskSecurityExchange struct{ message.ExchangeValue }

func (f RiskSecurityExchange) Tag() fix.Tag { return tag.RiskSecurityExchange }

type RiskSecurityGroup struct{ message.StringValue }

func (f RiskSecurityGroup) Tag() fix.Tag { return tag.RiskSecurityGroup }

type RiskSecurityID struct{ message.StringValue }

func (f RiskSecurityID) Tag() fix.Tag { return tag.RiskSecurityID }

type RiskSecurityIDSource struct{ message.StringValue }

func (f RiskSecurityIDSource) Tag() fix.Tag { return tag.RiskSecurityIDSource }

type RiskSecuritySubType struct{ message.StringValue }

func (f RiskSecuritySubType) Tag() fix.Tag { return tag.RiskSecuritySubType }

type RiskSecurityType struct{ message.StringValue }

func (f RiskSecurityType) Tag() fix.Tag { return tag.RiskSecurityType }

type RiskSeniority struct{ message.StringValue }

func (f RiskSeniority) Tag() fix.Tag { return tag.RiskSeniority }

type RiskSymbol struct{ message.StringValue }

func (f RiskSymbol) Tag() fix.Tag { return tag.RiskSymbol }

type RiskSymbolSfx struct{ message.StringValue }

func (f RiskSymbolSfx) Tag() fix.Tag { return tag.RiskSymbolSfx }

type RiskWarningLevelName struct{ message.StringValue }

func (f RiskWarningLevelName) Tag() fix.Tag { return tag.RiskWarningLevelName }

type RiskWarningLevelPercent struct{ message.PercentageValue }

func (f RiskWarningLevelPercent) Tag() fix.Tag { return tag.RiskWarningLevelPercent }

type RndPx struct{ message.PriceValue }

func (f RndPx) Tag() fix.Tag { return tag.RndPx }

type RootPartyID struct{ message.StringValue }

func (f RootPartyID) Tag() fix.Tag { return tag.RootPartyID }

type RootPartyIDSource struct{ message.CharValue }

func (f RootPartyIDSource) Tag() fix.Tag { return tag.RootPartyIDSource }

type RootPartyRole struct{ message.IntValue }

func (f RootPartyRole) Tag() fix.Tag { return tag.RootPartyRole }

type RootPartySubID struct{ message.StringValue }

func (f RootPartySubID) Tag() fix.Tag { return tag.RootPartySubID }

type RootPartySubIDType struct{ message.IntValue }

func (f RootPartySubIDType) Tag() fix.Tag { return tag.RootPartySubIDType }

type RoundLot struct{ message.QtyValue }

func (f RoundLot) Tag() fix.Tag { return tag.RoundLot }

type RoundingDirection struct{ message.CharValue }

func (f RoundingDirection) Tag() fix.Tag { return tag.RoundingDirection }

type RoundingModulus struct{ message.FloatValue }

func (f RoundingModulus) Tag() fix.Tag { return tag.RoundingModulus }

type RoutingID struct{ message.StringValue }

func (f RoutingID) Tag() fix.Tag { return tag.RoutingID }

type RoutingType struct{ message.IntValue }

func (f RoutingType) Tag() fix.Tag { return tag.RoutingType }

type RptSeq struct{ message.IntValue }

func (f RptSeq) Tag() fix.Tag { return tag.RptSeq }

type RptSys struct{ message.StringValue }

func (f RptSys) Tag() fix.Tag { return tag.RptSys }

type Rule80A struct{ message.CharValue }

func (f Rule80A) Tag() fix.Tag { return tag.Rule80A }

type Scope struct{ message.MultipleCharValue }

func (f Scope) Tag() fix.Tag { return tag.Scope }

type SecDefStatus struct{ message.IntValue }

func (f SecDefStatus) Tag() fix.Tag { return tag.SecDefStatus }

type SecondaryAllocID struct{ message.StringValue }

func (f SecondaryAllocID) Tag() fix.Tag { return tag.SecondaryAllocID }

type SecondaryClOrdID struct{ message.StringValue }

func (f SecondaryClOrdID) Tag() fix.Tag { return tag.SecondaryClOrdID }

type SecondaryDisplayQty struct{ message.QtyValue }

func (f SecondaryDisplayQty) Tag() fix.Tag { return tag.SecondaryDisplayQty }

type SecondaryExecID struct{ message.StringValue }

func (f SecondaryExecID) Tag() fix.Tag { return tag.SecondaryExecID }

type SecondaryFirmTradeID struct{ message.StringValue }

func (f SecondaryFirmTradeID) Tag() fix.Tag { return tag.SecondaryFirmTradeID }

type SecondaryHighLimitPrice struct{ message.PriceValue }

func (f SecondaryHighLimitPrice) Tag() fix.Tag { return tag.SecondaryHighLimitPrice }

type SecondaryIndividualAllocID struct{ message.StringValue }

func (f SecondaryIndividualAllocID) Tag() fix.Tag { return tag.SecondaryIndividualAllocID }

type SecondaryLowLimitPrice struct{ message.PriceValue }

func (f SecondaryLowLimitPrice) Tag() fix.Tag { return tag.SecondaryLowLimitPrice }

type SecondaryOrderID struct{ message.StringValue }

func (f SecondaryOrderID) Tag() fix.Tag { return tag.SecondaryOrderID }

type SecondaryPriceLimitType struct{ message.IntValue }

func (f SecondaryPriceLimitType) Tag() fix.Tag { return tag.SecondaryPriceLimitType }

type SecondaryTradeID struct{ message.StringValue }

func (f SecondaryTradeID) Tag() fix.Tag { return tag.SecondaryTradeID }

type SecondaryTradeReportID struct{ message.StringValue }

func (f SecondaryTradeReportID) Tag() fix.Tag { return tag.SecondaryTradeReportID }

type SecondaryTradeReportRefID struct{ message.StringValue }

func (f SecondaryTradeReportRefID) Tag() fix.Tag { return tag.SecondaryTradeReportRefID }

type SecondaryTradingReferencePrice struct{ message.PriceValue }

func (f SecondaryTradingReferencePrice) Tag() fix.Tag { return tag.SecondaryTradingReferencePrice }

type SecondaryTrdType struct{ message.IntValue }

func (f SecondaryTrdType) Tag() fix.Tag { return tag.SecondaryTrdType }

type SecureData struct{ message.DataValue }

func (f SecureData) Tag() fix.Tag { return tag.SecureData }

type SecureDataLen struct{ message.LengthValue }

func (f SecureDataLen) Tag() fix.Tag { return tag.SecureDataLen }

type SecurityAltID struct{ message.StringValue }

func (f SecurityAltID) Tag() fix.Tag { return tag.SecurityAltID }

type SecurityAltIDSource struct{ message.StringValue }

func (f SecurityAltIDSource) Tag() fix.Tag { return tag.SecurityAltIDSource }

type SecurityDesc struct{ message.StringValue }

func (f SecurityDesc) Tag() fix.Tag { return tag.SecurityDesc }

type SecurityExchange struct{ message.ExchangeValue }

func (f SecurityExchange) Tag() fix.Tag { return tag.SecurityExchange }

type SecurityGroup struct{ message.StringValue }

func (f SecurityGroup) Tag() fix.Tag { return tag.SecurityGroup }

type SecurityID struct{ message.StringValue }

func (f SecurityID) Tag() fix.Tag { return tag.SecurityID }

type SecurityIDSource struct{ message.StringValue }

func (f SecurityIDSource) Tag() fix.Tag { return tag.SecurityIDSource }

type SecurityListDesc struct{ message.StringValue }

func (f SecurityListDesc) Tag() fix.Tag { return tag.SecurityListDesc }

type SecurityListID struct{ message.StringValue }

func (f SecurityListID) Tag() fix.Tag { return tag.SecurityListID }

type SecurityListRefID struct{ message.StringValue }

func (f SecurityListRefID) Tag() fix.Tag { return tag.SecurityListRefID }

type SecurityListRequestType struct{ message.IntValue }

func (f SecurityListRequestType) Tag() fix.Tag { return tag.SecurityListRequestType }

type SecurityListType struct{ message.IntValue }

func (f SecurityListType) Tag() fix.Tag { return tag.SecurityListType }

type SecurityListTypeSource struct{ message.IntValue }

func (f SecurityListTypeSource) Tag() fix.Tag { return tag.SecurityListTypeSource }

type SecurityReportID struct{ message.IntValue }

func (f SecurityReportID) Tag() fix.Tag { return tag.SecurityReportID }

type SecurityReqID struct{ message.StringValue }

func (f SecurityReqID) Tag() fix.Tag { return tag.SecurityReqID }

type SecurityRequestResult struct{ message.IntValue }

func (f SecurityRequestResult) Tag() fix.Tag { return tag.SecurityRequestResult }

type SecurityRequestType struct{ message.IntValue }

func (f SecurityRequestType) Tag() fix.Tag { return tag.SecurityRequestType }

type SecurityResponseID struct{ message.StringValue }

func (f SecurityResponseID) Tag() fix.Tag { return tag.SecurityResponseID }

type SecurityResponseType struct{ message.IntValue }

func (f SecurityResponseType) Tag() fix.Tag { return tag.SecurityResponseType }

type SecuritySettlAgentAcctName struct{ message.StringValue }

func (f SecuritySettlAgentAcctName) Tag() fix.Tag { return tag.SecuritySettlAgentAcctName }

type SecuritySettlAgentAcctNum struct{ message.StringValue }

func (f SecuritySettlAgentAcctNum) Tag() fix.Tag { return tag.SecuritySettlAgentAcctNum }

type SecuritySettlAgentCode struct{ message.StringValue }

func (f SecuritySettlAgentCode) Tag() fix.Tag { return tag.SecuritySettlAgentCode }

type SecuritySettlAgentContactName struct{ message.StringValue }

func (f SecuritySettlAgentContactName) Tag() fix.Tag { return tag.SecuritySettlAgentContactName }

type SecuritySettlAgentContactPhone struct{ message.StringValue }

func (f SecuritySettlAgentContactPhone) Tag() fix.Tag { return tag.SecuritySettlAgentContactPhone }

type SecuritySettlAgentName struct{ message.StringValue }

func (f SecuritySettlAgentName) Tag() fix.Tag { return tag.SecuritySettlAgentName }

type SecurityStatus struct{ message.StringValue }

func (f SecurityStatus) Tag() fix.Tag { return tag.SecurityStatus }

type SecurityStatusReqID struct{ message.StringValue }

func (f SecurityStatusReqID) Tag() fix.Tag { return tag.SecurityStatusReqID }

type SecuritySubType struct{ message.StringValue }

func (f SecuritySubType) Tag() fix.Tag { return tag.SecuritySubType }

type SecurityTradingEvent struct{ message.IntValue }

func (f SecurityTradingEvent) Tag() fix.Tag { return tag.SecurityTradingEvent }

type SecurityTradingStatus struct{ message.IntValue }

func (f SecurityTradingStatus) Tag() fix.Tag { return tag.SecurityTradingStatus }

type SecurityType struct{ message.StringValue }

func (f SecurityType) Tag() fix.Tag { return tag.SecurityType }

type SecurityUpdateAction struct{ message.CharValue }

func (f SecurityUpdateAction) Tag() fix.Tag { return tag.SecurityUpdateAction }

type SecurityXML struct{ message.XMLDataValue }

func (f SecurityXML) Tag() fix.Tag { return tag.SecurityXML }

type SecurityXMLLen struct{ message.LengthValue }

func (f SecurityXMLLen) Tag() fix.Tag { return tag.SecurityXMLLen }

type SecurityXMLSchema struct{ message.StringValue }

func (f SecurityXMLSchema) Tag() fix.Tag { return tag.SecurityXMLSchema }

type SellVolume struct{ message.QtyValue }

func (f SellVolume) Tag() fix.Tag { return tag.SellVolume }

type SellerDays struct{ message.IntValue }

func (f SellerDays) Tag() fix.Tag { return tag.SellerDays }

type SenderCompID struct{ message.StringValue }

func (f SenderCompID) Tag() fix.Tag { return tag.SenderCompID }

type SenderLocationID struct{ message.StringValue }

func (f SenderLocationID) Tag() fix.Tag { return tag.SenderLocationID }

type SenderSubID struct{ message.StringValue }

func (f SenderSubID) Tag() fix.Tag { return tag.SenderSubID }

type SendingDate struct{ message.LocalMktDateValue }

func (f SendingDate) Tag() fix.Tag { return tag.SendingDate }

type SendingTime struct{ message.UTCTimestampValue }

func (f SendingTime) Tag() fix.Tag { return tag.SendingTime }

type Seniority struct{ message.StringValue }

func (f Seniority) Tag() fix.Tag { return tag.Seniority }

type SessionRejectReason struct{ message.IntValue }

func (f SessionRejectReason) Tag() fix.Tag { return tag.SessionRejectReason }

type SessionStatus struct{ message.IntValue }

func (f SessionStatus) Tag() fix.Tag { return tag.SessionStatus }

type SettlBrkrCode struct{ message.StringValue }

func (f SettlBrkrCode) Tag() fix.Tag { return tag.SettlBrkrCode }

type SettlCurrAmt struct{ message.AmtValue }

func (f SettlCurrAmt) Tag() fix.Tag { return tag.SettlCurrAmt }

type SettlCurrBidFxRate struct{ message.FloatValue }

func (f SettlCurrBidFxRate) Tag() fix.Tag { return tag.SettlCurrBidFxRate }

type SettlCurrFxRate struct{ message.FloatValue }

func (f SettlCurrFxRate) Tag() fix.Tag { return tag.SettlCurrFxRate }

type SettlCurrFxRateCalc struct{ message.CharValue }

func (f SettlCurrFxRateCalc) Tag() fix.Tag { return tag.SettlCurrFxRateCalc }

type SettlCurrOfferFxRate struct{ message.FloatValue }

func (f SettlCurrOfferFxRate) Tag() fix.Tag { return tag.SettlCurrOfferFxRate }

type SettlCurrency struct{ message.CurrencyValue }

func (f SettlCurrency) Tag() fix.Tag { return tag.SettlCurrency }

type SettlDate struct{ message.LocalMktDateValue }

func (f SettlDate) Tag() fix.Tag { return tag.SettlDate }

type SettlDate2 struct{ message.LocalMktDateValue }

func (f SettlDate2) Tag() fix.Tag { return tag.SettlDate2 }

type SettlDeliveryType struct{ message.IntValue }

func (f SettlDeliveryType) Tag() fix.Tag { return tag.SettlDeliveryType }

type SettlDepositoryCode struct{ message.StringValue }

func (f SettlDepositoryCode) Tag() fix.Tag { return tag.SettlDepositoryCode }

type SettlInstCode struct{ message.StringValue }

func (f SettlInstCode) Tag() fix.Tag { return tag.SettlInstCode }

type SettlInstID struct{ message.StringValue }

func (f SettlInstID) Tag() fix.Tag { return tag.SettlInstID }

type SettlInstMode struct{ message.CharValue }

func (f SettlInstMode) Tag() fix.Tag { return tag.SettlInstMode }

type SettlInstMsgID struct{ message.StringValue }

func (f SettlInstMsgID) Tag() fix.Tag { return tag.SettlInstMsgID }

type SettlInstRefID struct{ message.StringValue }

func (f SettlInstRefID) Tag() fix.Tag { return tag.SettlInstRefID }

type SettlInstReqID struct{ message.StringValue }

func (f SettlInstReqID) Tag() fix.Tag { return tag.SettlInstReqID }

type SettlInstReqRejCode struct{ message.IntValue }

func (f SettlInstReqRejCode) Tag() fix.Tag { return tag.SettlInstReqRejCode }

type SettlInstSource struct{ message.CharValue }

func (f SettlInstSource) Tag() fix.Tag { return tag.SettlInstSource }

type SettlInstTransType struct{ message.CharValue }

func (f SettlInstTransType) Tag() fix.Tag { return tag.SettlInstTransType }

type SettlLocation struct{ message.StringValue }

func (f SettlLocation) Tag() fix.Tag { return tag.SettlLocation }

type SettlMethod struct{ message.CharValue }

func (f SettlMethod) Tag() fix.Tag { return tag.SettlMethod }

type SettlObligID struct{ message.StringValue }

func (f SettlObligID) Tag() fix.Tag { return tag.SettlObligID }

type SettlObligMode struct{ message.IntValue }

func (f SettlObligMode) Tag() fix.Tag { return tag.SettlObligMode }

type SettlObligMsgID struct{ message.StringValue }

func (f SettlObligMsgID) Tag() fix.Tag { return tag.SettlObligMsgID }

type SettlObligRefID struct{ message.StringValue }

func (f SettlObligRefID) Tag() fix.Tag { return tag.SettlObligRefID }

type SettlObligSource struct{ message.CharValue }

func (f SettlObligSource) Tag() fix.Tag { return tag.SettlObligSource }

type SettlObligTransType struct{ message.CharValue }

func (f SettlObligTransType) Tag() fix.Tag { return tag.SettlObligTransType }

type SettlPartyID struct{ message.StringValue }

func (f SettlPartyID) Tag() fix.Tag { return tag.SettlPartyID }

type SettlPartyIDSource struct{ message.CharValue }

func (f SettlPartyIDSource) Tag() fix.Tag { return tag.SettlPartyIDSource }

type SettlPartyRole struct{ message.IntValue }

func (f SettlPartyRole) Tag() fix.Tag { return tag.SettlPartyRole }

type SettlPartySubID struct{ message.StringValue }

func (f SettlPartySubID) Tag() fix.Tag { return tag.SettlPartySubID }

type SettlPartySubIDType struct{ message.IntValue }

func (f SettlPartySubIDType) Tag() fix.Tag { return tag.SettlPartySubIDType }

type SettlPrice struct{ message.PriceValue }

func (f SettlPrice) Tag() fix.Tag { return tag.SettlPrice }

type SettlPriceType struct{ message.IntValue }

func (f SettlPriceType) Tag() fix.Tag { return tag.SettlPriceType }

type SettlSessID struct{ message.StringValue }

func (f SettlSessID) Tag() fix.Tag { return tag.SettlSessID }

type SettlSessSubID struct{ message.StringValue }

func (f SettlSessSubID) Tag() fix.Tag { return tag.SettlSessSubID }

type SettlType struct{ message.StringValue }

func (f SettlType) Tag() fix.Tag { return tag.SettlType }

type SettleOnOpenFlag struct{ message.StringValue }

func (f SettleOnOpenFlag) Tag() fix.Tag { return tag.SettleOnOpenFlag }

type SettlementCycleNo struct{ message.IntValue }

func (f SettlementCycleNo) Tag() fix.Tag { return tag.SettlementCycleNo }

type SettlmntTyp struct{ message.CharValue }

func (f SettlmntTyp) Tag() fix.Tag { return tag.SettlmntTyp }

type SharedCommission struct{ message.AmtValue }

func (f SharedCommission) Tag() fix.Tag { return tag.SharedCommission }

type Shares struct{ message.QtyValue }

func (f Shares) Tag() fix.Tag { return tag.Shares }

type ShortQty struct{ message.QtyValue }

func (f ShortQty) Tag() fix.Tag { return tag.ShortQty }

type ShortSaleReason struct{ message.IntValue }

func (f ShortSaleReason) Tag() fix.Tag { return tag.ShortSaleReason }

type Side struct{ message.CharValue }

func (f Side) Tag() fix.Tag { return tag.Side }

type SideComplianceID struct{ message.StringValue }

func (f SideComplianceID) Tag() fix.Tag { return tag.SideComplianceID }

type SideCurrency struct{ message.CurrencyValue }

func (f SideCurrency) Tag() fix.Tag { return tag.SideCurrency }

type SideExecID struct{ message.StringValue }

func (f SideExecID) Tag() fix.Tag { return tag.SideExecID }

type SideFillStationCd struct{ message.StringValue }

func (f SideFillStationCd) Tag() fix.Tag { return tag.SideFillStationCd }

type SideGrossTradeAmt struct{ message.AmtValue }

func (f SideGrossTradeAmt) Tag() fix.Tag { return tag.SideGrossTradeAmt }

type SideLastQty struct{ message.IntValue }

func (f SideLastQty) Tag() fix.Tag { return tag.SideLastQty }

type SideLiquidityInd struct{ message.IntValue }

func (f SideLiquidityInd) Tag() fix.Tag { return tag.SideLiquidityInd }

type SideMultiLegReportingType struct{ message.IntValue }

func (f SideMultiLegReportingType) Tag() fix.Tag { return tag.SideMultiLegReportingType }

type SideQty struct{ message.IntValue }

func (f SideQty) Tag() fix.Tag { return tag.SideQty }

type SideReasonCd struct{ message.StringValue }

func (f SideReasonCd) Tag() fix.Tag { return tag.SideReasonCd }

type SideSettlCurrency struct{ message.CurrencyValue }

func (f SideSettlCurrency) Tag() fix.Tag { return tag.SideSettlCurrency }

type SideTimeInForce struct{ message.UTCTimestampValue }

func (f SideTimeInForce) Tag() fix.Tag { return tag.SideTimeInForce }

type SideTradeReportID struct{ message.StringValue }

func (f SideTradeReportID) Tag() fix.Tag { return tag.SideTradeReportID }

type SideTrdRegTimestamp struct{ message.UTCTimestampValue }

func (f SideTrdRegTimestamp) Tag() fix.Tag { return tag.SideTrdRegTimestamp }

type SideTrdRegTimestampSrc struct{ message.StringValue }

func (f SideTrdRegTimestampSrc) Tag() fix.Tag { return tag.SideTrdRegTimestampSrc }

type SideTrdRegTimestampType struct{ message.IntValue }

func (f SideTrdRegTimestampType) Tag() fix.Tag { return tag.SideTrdRegTimestampType }

type SideTrdSubTyp struct{ message.IntValue }

func (f SideTrdSubTyp) Tag() fix.Tag { return tag.SideTrdSubTyp }

type SideValue1 struct{ message.AmtValue }

func (f SideValue1) Tag() fix.Tag { return tag.SideValue1 }

type SideValue2 struct{ message.AmtValue }

func (f SideValue2) Tag() fix.Tag { return tag.SideValue2 }

type SideValueInd struct{ message.IntValue }

func (f SideValueInd) Tag() fix.Tag { return tag.SideValueInd }

type Signature struct{ message.DataValue }

func (f Signature) Tag() fix.Tag { return tag.Signature }

type SignatureLength struct{ message.LengthValue }

func (f SignatureLength) Tag() fix.Tag { return tag.SignatureLength }

type SolicitedFlag struct{ message.BooleanValue }

func (f SolicitedFlag) Tag() fix.Tag { return tag.SolicitedFlag }

type Spread struct{ message.PriceOffsetValue }

func (f Spread) Tag() fix.Tag { return tag.Spread }

type SpreadToBenchmark struct{ message.PriceOffsetValue }

func (f SpreadToBenchmark) Tag() fix.Tag { return tag.SpreadToBenchmark }

type StandInstDbID struct{ message.StringValue }

func (f StandInstDbID) Tag() fix.Tag { return tag.StandInstDbID }

type StandInstDbName struct{ message.StringValue }

func (f StandInstDbName) Tag() fix.Tag { return tag.StandInstDbName }

type StandInstDbType struct{ message.IntValue }

func (f StandInstDbType) Tag() fix.Tag { return tag.StandInstDbType }

type StartCash struct{ message.AmtValue }

func (f StartCash) Tag() fix.Tag { return tag.StartCash }

type StartDate struct{ message.LocalMktDateValue }

func (f StartDate) Tag() fix.Tag { return tag.StartDate }

type StartMaturityMonthYear struct{ message.MonthYearValue }

func (f StartMaturityMonthYear) Tag() fix.Tag { return tag.StartMaturityMonthYear }

type StartStrikePxRange struct{ message.PriceValue }

func (f StartStrikePxRange) Tag() fix.Tag { return tag.StartStrikePxRange }

type StartTickPriceRange struct{ message.PriceValue }

func (f StartTickPriceRange) Tag() fix.Tag { return tag.StartTickPriceRange }

type StateOrProvinceOfIssue struct{ message.StringValue }

func (f StateOrProvinceOfIssue) Tag() fix.Tag { return tag.StateOrProvinceOfIssue }

type StatsType struct{ message.IntValue }

func (f StatsType) Tag() fix.Tag { return tag.StatsType }

type StatusText struct{ message.StringValue }

func (f StatusText) Tag() fix.Tag { return tag.StatusText }

type StatusValue struct{ message.IntValue }

func (f StatusValue) Tag() fix.Tag { return tag.StatusValue }

type StipulationType struct{ message.StringValue }

func (f StipulationType) Tag() fix.Tag { return tag.StipulationType }

type StipulationValue struct{ message.StringValue }

func (f StipulationValue) Tag() fix.Tag { return tag.StipulationValue }

type StopPx struct{ message.PriceValue }

func (f StopPx) Tag() fix.Tag { return tag.StopPx }

type StrategyParameterName struct{ message.StringValue }

func (f StrategyParameterName) Tag() fix.Tag { return tag.StrategyParameterName }

type StrategyParameterType struct{ message.IntValue }

func (f StrategyParameterType) Tag() fix.Tag { return tag.StrategyParameterType }

type StrategyParameterValue struct{ message.StringValue }

func (f StrategyParameterValue) Tag() fix.Tag { return tag.StrategyParameterValue }

type StreamAsgnAckType struct{ message.IntValue }

func (f StreamAsgnAckType) Tag() fix.Tag { return tag.StreamAsgnAckType }

type StreamAsgnRejReason struct{ message.IntValue }

func (f StreamAsgnRejReason) Tag() fix.Tag { return tag.StreamAsgnRejReason }

type StreamAsgnReqID struct{ message.StringValue }

func (f StreamAsgnReqID) Tag() fix.Tag { return tag.StreamAsgnReqID }

type StreamAsgnReqType struct{ message.IntValue }

func (f StreamAsgnReqType) Tag() fix.Tag { return tag.StreamAsgnReqType }

type StreamAsgnRptID struct{ message.StringValue }

func (f StreamAsgnRptID) Tag() fix.Tag { return tag.StreamAsgnRptID }

type StreamAsgnType struct{ message.IntValue }

func (f StreamAsgnType) Tag() fix.Tag { return tag.StreamAsgnType }

type StrikeCurrency struct{ message.CurrencyValue }

func (f StrikeCurrency) Tag() fix.Tag { return tag.StrikeCurrency }

type StrikeExerciseStyle struct{ message.IntValue }

func (f StrikeExerciseStyle) Tag() fix.Tag { return tag.StrikeExerciseStyle }

type StrikeIncrement struct{ message.FloatValue }

func (f StrikeIncrement) Tag() fix.Tag { return tag.StrikeIncrement }

type StrikeMultiplier struct{ message.FloatValue }

func (f StrikeMultiplier) Tag() fix.Tag { return tag.StrikeMultiplier }

type StrikePrice struct{ message.PriceValue }

func (f StrikePrice) Tag() fix.Tag { return tag.StrikePrice }

type StrikePriceBoundaryMethod struct{ message.IntValue }

func (f StrikePriceBoundaryMethod) Tag() fix.Tag { return tag.StrikePriceBoundaryMethod }

type StrikePriceBoundaryPrecision struct{ message.PercentageValue }

func (f StrikePriceBoundaryPrecision) Tag() fix.Tag { return tag.StrikePriceBoundaryPrecision }

type StrikePriceDeterminationMethod struct{ message.IntValue }

func (f StrikePriceDeterminationMethod) Tag() fix.Tag { return tag.StrikePriceDeterminationMethod }

type StrikeRuleID struct{ message.StringValue }

func (f StrikeRuleID) Tag() fix.Tag { return tag.StrikeRuleID }

type StrikeTime struct{ message.UTCTimestampValue }

func (f StrikeTime) Tag() fix.Tag { return tag.StrikeTime }

type StrikeValue struct{ message.FloatValue }

func (f StrikeValue) Tag() fix.Tag { return tag.StrikeValue }

type Subject struct{ message.StringValue }

func (f Subject) Tag() fix.Tag { return tag.Subject }

type SubscriptionRequestType struct{ message.CharValue }

func (f SubscriptionRequestType) Tag() fix.Tag { return tag.SubscriptionRequestType }

type SwapPoints struct{ message.PriceOffsetValue }

func (f SwapPoints) Tag() fix.Tag { return tag.SwapPoints }

type Symbol struct{ message.StringValue }

func (f Symbol) Tag() fix.Tag { return tag.Symbol }

type SymbolSfx struct{ message.StringValue }

func (f SymbolSfx) Tag() fix.Tag { return tag.SymbolSfx }

type TZTransactTime struct{ message.TZTimestampValue }

func (f TZTransactTime) Tag() fix.Tag { return tag.TZTransactTime }

type TargetCompID struct{ message.StringValue }

func (f TargetCompID) Tag() fix.Tag { return tag.TargetCompID }

type TargetLocationID struct{ message.StringValue }

func (f TargetLocationID) Tag() fix.Tag { return tag.TargetLocationID }

type TargetPartyID struct{ message.StringValue }

func (f TargetPartyID) Tag() fix.Tag { return tag.TargetPartyID }

type TargetPartyIDSource struct{ message.CharValue }

func (f TargetPartyIDSource) Tag() fix.Tag { return tag.TargetPartyIDSource }

type TargetPartyRole struct{ message.IntValue }

func (f TargetPartyRole) Tag() fix.Tag { return tag.TargetPartyRole }

type TargetStrategy struct{ message.IntValue }

func (f TargetStrategy) Tag() fix.Tag { return tag.TargetStrategy }

type TargetStrategyParameters struct{ message.StringValue }

func (f TargetStrategyParameters) Tag() fix.Tag { return tag.TargetStrategyParameters }

type TargetStrategyPerformance struct{ message.FloatValue }

func (f TargetStrategyPerformance) Tag() fix.Tag { return tag.TargetStrategyPerformance }

type TargetSubID struct{ message.StringValue }

func (f TargetSubID) Tag() fix.Tag { return tag.TargetSubID }

type TaxAdvantageType struct{ message.IntValue }

func (f TaxAdvantageType) Tag() fix.Tag { return tag.TaxAdvantageType }

type TerminationType struct{ message.IntValue }

func (f TerminationType) Tag() fix.Tag { return tag.TerminationType }

type TestMessageIndicator struct{ message.BooleanValue }

func (f TestMessageIndicator) Tag() fix.Tag { return tag.TestMessageIndicator }

type TestReqID struct{ message.StringValue }

func (f TestReqID) Tag() fix.Tag { return tag.TestReqID }

type Text struct{ message.StringValue }

func (f Text) Tag() fix.Tag { return tag.Text }

type ThresholdAmount struct{ message.PriceOffsetValue }

func (f ThresholdAmount) Tag() fix.Tag { return tag.ThresholdAmount }

type TickDirection struct{ message.CharValue }

func (f TickDirection) Tag() fix.Tag { return tag.TickDirection }

type TickIncrement struct{ message.PriceValue }

func (f TickIncrement) Tag() fix.Tag { return tag.TickIncrement }

type TickRuleType struct{ message.IntValue }

func (f TickRuleType) Tag() fix.Tag { return tag.TickRuleType }

type TierCode struct{ message.StringValue }

func (f TierCode) Tag() fix.Tag { return tag.TierCode }

type TimeBracket struct{ message.StringValue }

func (f TimeBracket) Tag() fix.Tag { return tag.TimeBracket }

type TimeInForce struct{ message.CharValue }

func (f TimeInForce) Tag() fix.Tag { return tag.TimeInForce }

type TimeToExpiration struct{ message.FloatValue }

func (f TimeToExpiration) Tag() fix.Tag { return tag.TimeToExpiration }

type TimeUnit struct{ message.StringValue }

func (f TimeUnit) Tag() fix.Tag { return tag.TimeUnit }

type TotNoAccQuotes struct{ message.IntValue }

func (f TotNoAccQuotes) Tag() fix.Tag { return tag.TotNoAccQuotes }

type TotNoAllocs struct{ message.IntValue }

func (f TotNoAllocs) Tag() fix.Tag { return tag.TotNoAllocs }

type TotNoCxldQuotes struct{ message.IntValue }

func (f TotNoCxldQuotes) Tag() fix.Tag { return tag.TotNoCxldQuotes }

type TotNoFills struct{ message.IntValue }

func (f TotNoFills) Tag() fix.Tag { return tag.TotNoFills }

type TotNoOrders struct{ message.IntValue }

func (f TotNoOrders) Tag() fix.Tag { return tag.TotNoOrders }

type TotNoPartyList struct{ message.IntValue }

func (f TotNoPartyList) Tag() fix.Tag { return tag.TotNoPartyList }

type TotNoQuoteEntries struct{ message.IntValue }

func (f TotNoQuoteEntries) Tag() fix.Tag { return tag.TotNoQuoteEntries }

type TotNoRejQuotes struct{ message.IntValue }

func (f TotNoRejQuotes) Tag() fix.Tag { return tag.TotNoRejQuotes }

type TotNoRelatedSym struct{ message.IntValue }

func (f TotNoRelatedSym) Tag() fix.Tag { return tag.TotNoRelatedSym }

type TotNoSecurityTypes struct{ message.IntValue }

func (f TotNoSecurityTypes) Tag() fix.Tag { return tag.TotNoSecurityTypes }

type TotNoStrikes struct{ message.IntValue }

func (f TotNoStrikes) Tag() fix.Tag { return tag.TotNoStrikes }

type TotNumAssignmentReports struct{ message.IntValue }

func (f TotNumAssignmentReports) Tag() fix.Tag { return tag.TotNumAssignmentReports }

type TotNumReports struct{ message.IntValue }

func (f TotNumReports) Tag() fix.Tag { return tag.TotNumReports }

type TotNumTradeReports struct{ message.IntValue }

func (f TotNumTradeReports) Tag() fix.Tag { return tag.TotNumTradeReports }

type TotQuoteEntries struct{ message.IntValue }

func (f TotQuoteEntries) Tag() fix.Tag { return tag.TotQuoteEntries }

type TotalAccruedInterestAmt struct{ message.AmtValue }

func (f TotalAccruedInterestAmt) Tag() fix.Tag { return tag.TotalAccruedInterestAmt }

type TotalAffectedOrders struct{ message.IntValue }

func (f TotalAffectedOrders) Tag() fix.Tag { return tag.TotalAffectedOrders }

type TotalNetValue struct{ message.AmtValue }

func (f TotalNetValue) Tag() fix.Tag { return tag.TotalNetValue }

type TotalNumPosReports struct{ message.IntValue }

func (f TotalNumPosReports) Tag() fix.Tag { return tag.TotalNumPosReports }

type TotalNumSecurities struct{ message.IntValue }

func (f TotalNumSecurities) Tag() fix.Tag { return tag.TotalNumSecurities }

type TotalNumSecurityTypes struct{ message.IntValue }

func (f TotalNumSecurityTypes) Tag() fix.Tag { return tag.TotalNumSecurityTypes }

type TotalTakedown struct{ message.AmtValue }

func (f TotalTakedown) Tag() fix.Tag { return tag.TotalTakedown }

type TotalVolumeTraded struct{ message.QtyValue }

func (f TotalVolumeTraded) Tag() fix.Tag { return tag.TotalVolumeTraded }

type TotalVolumeTradedDate struct{ message.UTCDateOnlyValue }

func (f TotalVolumeTradedDate) Tag() fix.Tag { return tag.TotalVolumeTradedDate }

type TotalVolumeTradedTime struct{ message.UTCTimeOnlyValue }

func (f TotalVolumeTradedTime) Tag() fix.Tag { return tag.TotalVolumeTradedTime }

type TradSesCloseTime struct{ message.UTCTimestampValue }

func (f TradSesCloseTime) Tag() fix.Tag { return tag.TradSesCloseTime }

type TradSesEndTime struct{ message.UTCTimestampValue }

func (f TradSesEndTime) Tag() fix.Tag { return tag.TradSesEndTime }

type TradSesEvent struct{ message.IntValue }

func (f TradSesEvent) Tag() fix.Tag { return tag.TradSesEvent }

type TradSesMethod struct{ message.IntValue }

func (f TradSesMethod) Tag() fix.Tag { return tag.TradSesMethod }

type TradSesMode struct{ message.IntValue }

func (f TradSesMode) Tag() fix.Tag { return tag.TradSesMode }

type TradSesOpenTime struct{ message.UTCTimestampValue }

func (f TradSesOpenTime) Tag() fix.Tag { return tag.TradSesOpenTime }

type TradSesPreCloseTime struct{ message.UTCTimestampValue }

func (f TradSesPreCloseTime) Tag() fix.Tag { return tag.TradSesPreCloseTime }

type TradSesReqID struct{ message.StringValue }

func (f TradSesReqID) Tag() fix.Tag { return tag.TradSesReqID }

type TradSesStartTime struct{ message.UTCTimestampValue }

func (f TradSesStartTime) Tag() fix.Tag { return tag.TradSesStartTime }

type TradSesStatus struct{ message.IntValue }

func (f TradSesStatus) Tag() fix.Tag { return tag.TradSesStatus }

type TradSesStatusRejReason struct{ message.IntValue }

func (f TradSesStatusRejReason) Tag() fix.Tag { return tag.TradSesStatusRejReason }

type TradSesUpdateAction struct{ message.CharValue }

func (f TradSesUpdateAction) Tag() fix.Tag { return tag.TradSesUpdateAction }

type TradeAllocIndicator struct{ message.IntValue }

func (f TradeAllocIndicator) Tag() fix.Tag { return tag.TradeAllocIndicator }

type TradeCondition struct{ message.MultipleStringValue }

func (f TradeCondition) Tag() fix.Tag { return tag.TradeCondition }

type TradeDate struct{ message.LocalMktDateValue }

func (f TradeDate) Tag() fix.Tag { return tag.TradeDate }

type TradeHandlingInstr struct{ message.CharValue }

func (f TradeHandlingInstr) Tag() fix.Tag { return tag.TradeHandlingInstr }

type TradeID struct{ message.StringValue }

func (f TradeID) Tag() fix.Tag { return tag.TradeID }

type TradeInputDevice struct{ message.StringValue }

func (f TradeInputDevice) Tag() fix.Tag { return tag.TradeInputDevice }

type TradeInputSource struct{ message.StringValue }

func (f TradeInputSource) Tag() fix.Tag { return tag.TradeInputSource }

type TradeLegRefID struct{ message.StringValue }

func (f TradeLegRefID) Tag() fix.Tag { return tag.TradeLegRefID }

type TradeLinkID struct{ message.StringValue }

func (f TradeLinkID) Tag() fix.Tag { return tag.TradeLinkID }

type TradeOriginationDate struct{ message.LocalMktDateValue }

func (f TradeOriginationDate) Tag() fix.Tag { return tag.TradeOriginationDate }

type TradePublishIndicator struct{ message.IntValue }

func (f TradePublishIndicator) Tag() fix.Tag { return tag.TradePublishIndicator }

type TradeReportID struct{ message.StringValue }

func (f TradeReportID) Tag() fix.Tag { return tag.TradeReportID }

type TradeReportRefID struct{ message.StringValue }

func (f TradeReportRefID) Tag() fix.Tag { return tag.TradeReportRefID }

type TradeReportRejectReason struct{ message.IntValue }

func (f TradeReportRejectReason) Tag() fix.Tag { return tag.TradeReportRejectReason }

type TradeReportTransType struct{ message.IntValue }

func (f TradeReportTransType) Tag() fix.Tag { return tag.TradeReportTransType }

type TradeReportType struct{ message.IntValue }

func (f TradeReportType) Tag() fix.Tag { return tag.TradeReportType }

type TradeRequestID struct{ message.StringValue }

func (f TradeRequestID) Tag() fix.Tag { return tag.TradeRequestID }

type TradeRequestResult struct{ message.IntValue }

func (f TradeRequestResult) Tag() fix.Tag { return tag.TradeRequestResult }

type TradeRequestStatus struct{ message.IntValue }

func (f TradeRequestStatus) Tag() fix.Tag { return tag.TradeRequestStatus }

type TradeRequestType struct{ message.IntValue }

func (f TradeRequestType) Tag() fix.Tag { return tag.TradeRequestType }

type TradeType struct{ message.CharValue }

func (f TradeType) Tag() fix.Tag { return tag.TradeType }

type TradeVolume struct{ message.QtyValue }

func (f TradeVolume) Tag() fix.Tag { return tag.TradeVolume }

type TradedFlatSwitch struct{ message.BooleanValue }

func (f TradedFlatSwitch) Tag() fix.Tag { return tag.TradedFlatSwitch }

type TradingCurrency struct{ message.CurrencyValue }

func (f TradingCurrency) Tag() fix.Tag { return tag.TradingCurrency }

type TradingReferencePrice struct{ message.PriceValue }

func (f TradingReferencePrice) Tag() fix.Tag { return tag.TradingReferencePrice }

type TradingSessionDesc struct{ message.StringValue }

func (f TradingSessionDesc) Tag() fix.Tag { return tag.TradingSessionDesc }

type TradingSessionID struct{ message.StringValue }

func (f TradingSessionID) Tag() fix.Tag { return tag.TradingSessionID }

type TradingSessionSubID struct{ message.StringValue }

func (f TradingSessionSubID) Tag() fix.Tag { return tag.TradingSessionSubID }

type TransBkdTime struct{ message.UTCTimestampValue }

func (f TransBkdTime) Tag() fix.Tag { return tag.TransBkdTime }

type TransactTime struct{ message.UTCTimestampValue }

func (f TransactTime) Tag() fix.Tag { return tag.TransactTime }

type TransferReason struct{ message.StringValue }

func (f TransferReason) Tag() fix.Tag { return tag.TransferReason }

type TrdMatchID struct{ message.StringValue }

func (f TrdMatchID) Tag() fix.Tag { return tag.TrdMatchID }

type TrdRegTimestamp struct{ message.UTCTimestampValue }

func (f TrdRegTimestamp) Tag() fix.Tag { return tag.TrdRegTimestamp }

type TrdRegTimestampOrigin struct{ message.StringValue }

func (f TrdRegTimestampOrigin) Tag() fix.Tag { return tag.TrdRegTimestampOrigin }

type TrdRegTimestampType struct{ message.IntValue }

func (f TrdRegTimestampType) Tag() fix.Tag { return tag.TrdRegTimestampType }

type TrdRepIndicator struct{ message.BooleanValue }

func (f TrdRepIndicator) Tag() fix.Tag { return tag.TrdRepIndicator }

type TrdRepPartyRole struct{ message.IntValue }

func (f TrdRepPartyRole) Tag() fix.Tag { return tag.TrdRepPartyRole }

type TrdRptStatus struct{ message.IntValue }

func (f TrdRptStatus) Tag() fix.Tag { return tag.TrdRptStatus }

type TrdSubType struct{ message.IntValue }

func (f TrdSubType) Tag() fix.Tag { return tag.TrdSubType }

type TrdType struct{ message.IntValue }

func (f TrdType) Tag() fix.Tag { return tag.TrdType }

type TriggerAction struct{ message.CharValue }

func (f TriggerAction) Tag() fix.Tag { return tag.TriggerAction }

type TriggerNewPrice struct{ message.PriceValue }

func (f TriggerNewPrice) Tag() fix.Tag { return tag.TriggerNewPrice }

type TriggerNewQty struct{ message.QtyValue }

func (f TriggerNewQty) Tag() fix.Tag { return tag.TriggerNewQty }

type TriggerOrderType struct{ message.CharValue }

func (f TriggerOrderType) Tag() fix.Tag { return tag.TriggerOrderType }

type TriggerPrice struct{ message.PriceValue }

func (f TriggerPrice) Tag() fix.Tag { return tag.TriggerPrice }

type TriggerPriceDirection struct{ message.CharValue }

func (f TriggerPriceDirection) Tag() fix.Tag { return tag.TriggerPriceDirection }

type TriggerPriceType struct{ message.CharValue }

func (f TriggerPriceType) Tag() fix.Tag { return tag.TriggerPriceType }

type TriggerPriceTypeScope struct{ message.CharValue }

func (f TriggerPriceTypeScope) Tag() fix.Tag { return tag.TriggerPriceTypeScope }

type TriggerSecurityDesc struct{ message.StringValue }

func (f TriggerSecurityDesc) Tag() fix.Tag { return tag.TriggerSecurityDesc }

type TriggerSecurityID struct{ message.StringValue }

func (f TriggerSecurityID) Tag() fix.Tag { return tag.TriggerSecurityID }

type TriggerSecurityIDSource struct{ message.StringValue }

func (f TriggerSecurityIDSource) Tag() fix.Tag { return tag.TriggerSecurityIDSource }

type TriggerSymbol struct{ message.StringValue }

func (f TriggerSymbol) Tag() fix.Tag { return tag.TriggerSymbol }

type TriggerTradingSessionID struct{ message.StringValue }

func (f TriggerTradingSessionID) Tag() fix.Tag { return tag.TriggerTradingSessionID }

type TriggerTradingSessionSubID struct{ message.StringValue }

func (f TriggerTradingSessionSubID) Tag() fix.Tag { return tag.TriggerTradingSessionSubID }

type TriggerType struct{ message.CharValue }

func (f TriggerType) Tag() fix.Tag { return tag.TriggerType }

type URLLink struct{ message.StringValue }

func (f URLLink) Tag() fix.Tag { return tag.URLLink }

type UnderlyingAdjustedQuantity struct{ message.QtyValue }

func (f UnderlyingAdjustedQuantity) Tag() fix.Tag { return tag.UnderlyingAdjustedQuantity }

type UnderlyingAllocationPercent struct{ message.PercentageValue }

func (f UnderlyingAllocationPercent) Tag() fix.Tag { return tag.UnderlyingAllocationPercent }

type UnderlyingAttachmentPoint struct{ message.PercentageValue }

func (f UnderlyingAttachmentPoint) Tag() fix.Tag { return tag.UnderlyingAttachmentPoint }

type UnderlyingCFICode struct{ message.StringValue }

func (f UnderlyingCFICode) Tag() fix.Tag { return tag.UnderlyingCFICode }

type UnderlyingCPProgram struct{ message.StringValue }

func (f UnderlyingCPProgram) Tag() fix.Tag { return tag.UnderlyingCPProgram }

type UnderlyingCPRegType struct{ message.StringValue }

func (f UnderlyingCPRegType) Tag() fix.Tag { return tag.UnderlyingCPRegType }

type UnderlyingCapValue struct{ message.AmtValue }

func (f UnderlyingCapValue) Tag() fix.Tag { return tag.UnderlyingCapValue }

type UnderlyingCashAmount struct{ message.AmtValue }

func (f UnderlyingCashAmount) Tag() fix.Tag { return tag.UnderlyingCashAmount }

type UnderlyingCashType struct{ message.StringValue }

func (f UnderlyingCashType) Tag() fix.Tag { return tag.UnderlyingCashType }

type UnderlyingCollectAmount struct{ message.AmtValue }

func (f UnderlyingCollectAmount) Tag() fix.Tag { return tag.UnderlyingCollectAmount }

type UnderlyingContractMultiplier struct{ message.FloatValue }

func (f UnderlyingContractMultiplier) Tag() fix.Tag { return tag.UnderlyingContractMultiplier }

type UnderlyingContractMultiplierUnit struct{ message.IntValue }

func (f UnderlyingContractMultiplierUnit) Tag() fix.Tag { return tag.UnderlyingContractMultiplierUnit }

type UnderlyingCountryOfIssue struct{ message.CountryValue }

func (f UnderlyingCountryOfIssue) Tag() fix.Tag { return tag.UnderlyingCountryOfIssue }

type UnderlyingCouponPaymentDate struct{ message.LocalMktDateValue }

func (f UnderlyingCouponPaymentDate) Tag() fix.Tag { return tag.UnderlyingCouponPaymentDate }

type UnderlyingCouponRate struct{ message.PercentageValue }

func (f UnderlyingCouponRate) Tag() fix.Tag { return tag.UnderlyingCouponRate }

type UnderlyingCreditRating struct{ message.StringValue }

func (f UnderlyingCreditRating) Tag() fix.Tag { return tag.UnderlyingCreditRating }

type UnderlyingCurrency struct{ message.CurrencyValue }

func (f UnderlyingCurrency) Tag() fix.Tag { return tag.UnderlyingCurrency }

type UnderlyingCurrentValue struct{ message.AmtValue }

func (f UnderlyingCurrentValue) Tag() fix.Tag { return tag.UnderlyingCurrentValue }

type UnderlyingDeliveryAmount struct{ message.AmtValue }

func (f UnderlyingDeliveryAmount) Tag() fix.Tag { return tag.UnderlyingDeliveryAmount }

type UnderlyingDetachmentPoint struct{ message.PercentageValue }

func (f UnderlyingDetachmentPoint) Tag() fix.Tag { return tag.UnderlyingDetachmentPoint }

type UnderlyingDirtyPrice struct{ message.PriceValue }

func (f UnderlyingDirtyPrice) Tag() fix.Tag { return tag.UnderlyingDirtyPrice }

type UnderlyingEndPrice struct{ message.PriceValue }

func (f UnderlyingEndPrice) Tag() fix.Tag { return tag.UnderlyingEndPrice }

type UnderlyingEndValue struct{ message.AmtValue }

func (f UnderlyingEndValue) Tag() fix.Tag { return tag.UnderlyingEndValue }

type UnderlyingExerciseStyle struct{ message.IntValue }

func (f UnderlyingExerciseStyle) Tag() fix.Tag { return tag.UnderlyingExerciseStyle }

type UnderlyingFXRate struct{ message.FloatValue }

func (f UnderlyingFXRate) Tag() fix.Tag { return tag.UnderlyingFXRate }

type UnderlyingFXRateCalc struct{ message.CharValue }

func (f UnderlyingFXRateCalc) Tag() fix.Tag { return tag.UnderlyingFXRateCalc }

type UnderlyingFactor struct{ message.FloatValue }

func (f UnderlyingFactor) Tag() fix.Tag { return tag.UnderlyingFactor }

type UnderlyingFlowScheduleType struct{ message.IntValue }

func (f UnderlyingFlowScheduleType) Tag() fix.Tag { return tag.UnderlyingFlowScheduleType }

type UnderlyingIDSource struct{ message.StringValue }

func (f UnderlyingIDSource) Tag() fix.Tag { return tag.UnderlyingIDSource }

type UnderlyingInstrRegistry struct{ message.StringValue }

func (f UnderlyingInstrRegistry) Tag() fix.Tag { return tag.UnderlyingInstrRegistry }

type UnderlyingInstrumentPartyID struct{ message.StringValue }

func (f UnderlyingInstrumentPartyID) Tag() fix.Tag { return tag.UnderlyingInstrumentPartyID }

type UnderlyingInstrumentPartyIDSource struct{ message.CharValue }

func (f UnderlyingInstrumentPartyIDSource) Tag() fix.Tag { return tag.UnderlyingInstrumentPartyIDSource }

type UnderlyingInstrumentPartyRole struct{ message.IntValue }

func (f UnderlyingInstrumentPartyRole) Tag() fix.Tag { return tag.UnderlyingInstrumentPartyRole }

type UnderlyingInstrumentPartySubID struct{ message.StringValue }

func (f UnderlyingInstrumentPartySubID) Tag() fix.Tag { return tag.UnderlyingInstrumentPartySubID }

type UnderlyingInstrumentPartySubIDType struct{ message.IntValue }

func (f UnderlyingInstrumentPartySubIDType) Tag() fix.Tag {
	return tag.UnderlyingInstrumentPartySubIDType
}

type UnderlyingIssueDate struct{ message.LocalMktDateValue }

func (f UnderlyingIssueDate) Tag() fix.Tag { return tag.UnderlyingIssueDate }

type UnderlyingIssuer struct{ message.StringValue }

func (f UnderlyingIssuer) Tag() fix.Tag { return tag.UnderlyingIssuer }

type UnderlyingLastPx struct{ message.PriceValue }

func (f UnderlyingLastPx) Tag() fix.Tag { return tag.UnderlyingLastPx }

type UnderlyingLastQty struct{ message.QtyValue }

func (f UnderlyingLastQty) Tag() fix.Tag { return tag.UnderlyingLastQty }

type UnderlyingLegCFICode struct{ message.StringValue }

func (f UnderlyingLegCFICode) Tag() fix.Tag { return tag.UnderlyingLegCFICode }

type UnderlyingLegMaturityDate struct{ message.LocalMktDateValue }

func (f UnderlyingLegMaturityDate) Tag() fix.Tag { return tag.UnderlyingLegMaturityDate }

type UnderlyingLegMaturityMonthYear struct{ message.MonthYearValue }

func (f UnderlyingLegMaturityMonthYear) Tag() fix.Tag { return tag.UnderlyingLegMaturityMonthYear }

type UnderlyingLegMaturityTime struct{ message.TZTimeOnlyValue }

func (f UnderlyingLegMaturityTime) Tag() fix.Tag { return tag.UnderlyingLegMaturityTime }

type UnderlyingLegOptAttribute struct{ message.CharValue }

func (f UnderlyingLegOptAttribute) Tag() fix.Tag { return tag.UnderlyingLegOptAttribute }

type UnderlyingLegPutOrCall struct{ message.IntValue }

func (f UnderlyingLegPutOrCall) Tag() fix.Tag { return tag.UnderlyingLegPutOrCall }

type UnderlyingLegSecurityAltID struct{ message.StringValue }

func (f UnderlyingLegSecurityAltID) Tag() fix.Tag { return tag.UnderlyingLegSecurityAltID }

type UnderlyingLegSecurityAltIDSource struct{ message.StringValue }

func (f UnderlyingLegSecurityAltIDSource) Tag() fix.Tag { return tag.UnderlyingLegSecurityAltIDSource }

type UnderlyingLegSecurityDesc struct{ message.StringValue }

func (f UnderlyingLegSecurityDesc) Tag() fix.Tag { return tag.UnderlyingLegSecurityDesc }

type UnderlyingLegSecurityExchange struct{ message.StringValue }

func (f UnderlyingLegSecurityExchange) Tag() fix.Tag { return tag.UnderlyingLegSecurityExchange }

type UnderlyingLegSecurityID struct{ message.StringValue }

func (f UnderlyingLegSecurityID) Tag() fix.Tag { return tag.UnderlyingLegSecurityID }

type UnderlyingLegSecurityIDSource struct{ message.StringValue }

func (f UnderlyingLegSecurityIDSource) Tag() fix.Tag { return tag.UnderlyingLegSecurityIDSource }

type UnderlyingLegSecuritySubType struct{ message.StringValue }

func (f UnderlyingLegSecuritySubType) Tag() fix.Tag { return tag.UnderlyingLegSecuritySubType }

type UnderlyingLegSecurityType struct{ message.StringValue }

func (f UnderlyingLegSecurityType) Tag() fix.Tag { return tag.UnderlyingLegSecurityType }

type UnderlyingLegStrikePrice struct{ message.PriceValue }

func (f UnderlyingLegStrikePrice) Tag() fix.Tag { return tag.UnderlyingLegStrikePrice }

type UnderlyingLegSymbol struct{ message.StringValue }

func (f UnderlyingLegSymbol) Tag() fix.Tag { return tag.UnderlyingLegSymbol }

type UnderlyingLegSymbolSfx struct{ message.StringValue }

func (f UnderlyingLegSymbolSfx) Tag() fix.Tag { return tag.UnderlyingLegSymbolSfx }

type UnderlyingLocaleOfIssue struct{ message.StringValue }

func (f UnderlyingLocaleOfIssue) Tag() fix.Tag { return tag.UnderlyingLocaleOfIssue }

type UnderlyingMaturityDate struct{ message.LocalMktDateValue }

func (f UnderlyingMaturityDate) Tag() fix.Tag { return tag.UnderlyingMaturityDate }

type UnderlyingMaturityDay struct{ message.DayOfMonthValue }

func (f UnderlyingMaturityDay) Tag() fix.Tag { return tag.UnderlyingMaturityDay }

type UnderlyingMaturityMonthYear struct{ message.MonthYearValue }

func (f UnderlyingMaturityMonthYear) Tag() fix.Tag { return tag.UnderlyingMaturityMonthYear }

type UnderlyingMaturityTime struct{ message.TZTimeOnlyValue }

func (f UnderlyingMaturityTime) Tag() fix.Tag { return tag.UnderlyingMaturityTime }

type UnderlyingNotionalPercentageOutstanding struct{ message.PercentageValue }

func (f UnderlyingNotionalPercentageOutstanding) Tag() fix.Tag {
	return tag.UnderlyingNotionalPercentageOutstanding
}

type UnderlyingOptAttribute struct{ message.CharValue }

func (f UnderlyingOptAttribute) Tag() fix.Tag { return tag.UnderlyingOptAttribute }

type UnderlyingOriginalNotionalPercentageOutstanding struct{ message.PercentageValue }

func (f UnderlyingOriginalNotionalPercentageOutstanding) Tag() fix.Tag {
	return tag.UnderlyingOriginalNotionalPercentageOutstanding
}

type UnderlyingPayAmount struct{ message.AmtValue }

func (f UnderlyingPayAmount) Tag() fix.Tag { return tag.UnderlyingPayAmount }

type UnderlyingPriceDeterminationMethod struct{ message.IntValue }

func (f UnderlyingPriceDeterminationMethod) Tag() fix.Tag {
	return tag.UnderlyingPriceDeterminationMethod
}

type UnderlyingPriceUnitOfMeasure struct{ message.StringValue }

func (f UnderlyingPriceUnitOfMeasure) Tag() fix.Tag { return tag.UnderlyingPriceUnitOfMeasure }

type UnderlyingPriceUnitOfMeasureQty struct{ message.QtyValue }

func (f UnderlyingPriceUnitOfMeasureQty) Tag() fix.Tag { return tag.UnderlyingPriceUnitOfMeasureQty }

type UnderlyingProduct struct{ message.IntValue }

func (f UnderlyingProduct) Tag() fix.Tag { return tag.UnderlyingProduct }

type UnderlyingPutOrCall struct{ message.IntValue }

func (f UnderlyingPutOrCall) Tag() fix.Tag { return tag.UnderlyingPutOrCall }

type UnderlyingPx struct{ message.PriceValue }

func (f UnderlyingPx) Tag() fix.Tag { return tag.UnderlyingPx }

type UnderlyingQty struct{ message.QtyValue }

func (f UnderlyingQty) Tag() fix.Tag { return tag.UnderlyingQty }

type UnderlyingRedemptionDate struct{ message.LocalMktDateValue }

func (f UnderlyingRedemptionDate) Tag() fix.Tag { return tag.UnderlyingRedemptionDate }

type UnderlyingRepoCollateralSecurityType struct{ message.IntValue }

func (f UnderlyingRepoCollateralSecurityType) Tag() fix.Tag {
	return tag.UnderlyingRepoCollateralSecurityType
}

type UnderlyingRepurchaseRate struct{ message.PercentageValue }

func (f UnderlyingRepurchaseRate) Tag() fix.Tag { return tag.UnderlyingRepurchaseRate }

type UnderlyingRepurchaseTerm struct{ message.IntValue }

func (f UnderlyingRepurchaseTerm) Tag() fix.Tag { return tag.UnderlyingRepurchaseTerm }

type UnderlyingRestructuringType struct{ message.StringValue }

func (f UnderlyingRestructuringType) Tag() fix.Tag { return tag.UnderlyingRestructuringType }

type UnderlyingSecurityAltID struct{ message.StringValue }

func (f UnderlyingSecurityAltID) Tag() fix.Tag { return tag.UnderlyingSecurityAltID }

type UnderlyingSecurityAltIDSource struct{ message.StringValue }

func (f UnderlyingSecurityAltIDSource) Tag() fix.Tag { return tag.UnderlyingSecurityAltIDSource }

type UnderlyingSecurityDesc struct{ message.StringValue }

func (f UnderlyingSecurityDesc) Tag() fix.Tag { return tag.UnderlyingSecurityDesc }

type UnderlyingSecurityExchange struct{ message.ExchangeValue }

func (f UnderlyingSecurityExchange) Tag() fix.Tag { return tag.UnderlyingSecurityExchange }

type UnderlyingSecurityID struct{ message.StringValue }

func (f UnderlyingSecurityID) Tag() fix.Tag { return tag.UnderlyingSecurityID }

type UnderlyingSecurityIDSource struct{ message.StringValue }

func (f UnderlyingSecurityIDSource) Tag() fix.Tag { return tag.UnderlyingSecurityIDSource }

type UnderlyingSecuritySubType struct{ message.StringValue }

func (f UnderlyingSecuritySubType) Tag() fix.Tag { return tag.UnderlyingSecuritySubType }

type UnderlyingSecurityType struct{ message.StringValue }

func (f UnderlyingSecurityType) Tag() fix.Tag { return tag.UnderlyingSecurityType }

type UnderlyingSeniority struct{ message.StringValue }

func (f UnderlyingSeniority) Tag() fix.Tag { return tag.UnderlyingSeniority }

type UnderlyingSettlMethod struct{ message.StringValue }

func (f UnderlyingSettlMethod) Tag() fix.Tag { return tag.UnderlyingSettlMethod }

type UnderlyingSettlPrice struct{ message.PriceValue }

func (f UnderlyingSettlPrice) Tag() fix.Tag { return tag.UnderlyingSettlPrice }

type UnderlyingSettlPriceType struct{ message.IntValue }

func (f UnderlyingSettlPriceType) Tag() fix.Tag { return tag.UnderlyingSettlPriceType }

type UnderlyingSettlementDate struct{ message.LocalMktDateValue }

func (f UnderlyingSettlementDate) Tag() fix.Tag { return tag.UnderlyingSettlementDate }

type UnderlyingSettlementStatus struct{ message.StringValue }

func (f UnderlyingSettlementStatus) Tag() fix.Tag { return tag.UnderlyingSettlementStatus }

type UnderlyingSettlementType struct{ message.IntValue }

func (f UnderlyingSettlementType) Tag() fix.Tag { return tag.UnderlyingSettlementType }

type UnderlyingStartValue struct{ message.AmtValue }

func (f UnderlyingStartValue) Tag() fix.Tag { return tag.UnderlyingStartValue }

type UnderlyingStateOrProvinceOfIssue struct{ message.StringValue }

func (f UnderlyingStateOrProvinceOfIssue) Tag() fix.Tag { return tag.UnderlyingStateOrProvinceOfIssue }

type UnderlyingStipType struct{ message.StringValue }

func (f UnderlyingStipType) Tag() fix.Tag { return tag.UnderlyingStipType }

type UnderlyingStipValue struct{ message.StringValue }

func (f UnderlyingStipValue) Tag() fix.Tag { return tag.UnderlyingStipValue }

type UnderlyingStrikeCurrency struct{ message.CurrencyValue }

func (f UnderlyingStrikeCurrency) Tag() fix.Tag { return tag.UnderlyingStrikeCurrency }

type UnderlyingStrikePrice struct{ message.PriceValue }

func (f UnderlyingStrikePrice) Tag() fix.Tag { return tag.UnderlyingStrikePrice }

type UnderlyingSymbol struct{ message.StringValue }

func (f UnderlyingSymbol) Tag() fix.Tag { return tag.UnderlyingSymbol }

type UnderlyingSymbolSfx struct{ message.StringValue }

func (f UnderlyingSymbolSfx) Tag() fix.Tag { return tag.UnderlyingSymbolSfx }

type UnderlyingTimeUnit struct{ message.StringValue }

func (f UnderlyingTimeUnit) Tag() fix.Tag { return tag.UnderlyingTimeUnit }

type UnderlyingTradingSessionID struct{ message.StringValue }

func (f UnderlyingTradingSessionID) Tag() fix.Tag { return tag.UnderlyingTradingSessionID }

type UnderlyingTradingSessionSubID struct{ message.StringValue }

func (f UnderlyingTradingSessionSubID) Tag() fix.Tag { return tag.UnderlyingTradingSessionSubID }

type UnderlyingUnitOfMeasure struct{ message.StringValue }

func (f UnderlyingUnitOfMeasure) Tag() fix.Tag { return tag.UnderlyingUnitOfMeasure }

type UnderlyingUnitOfMeasureQty struct{ message.QtyValue }

func (f UnderlyingUnitOfMeasureQty) Tag() fix.Tag { return tag.UnderlyingUnitOfMeasureQty }

type UndlyInstrumentPartyID struct{ message.StringValue }

func (f UndlyInstrumentPartyID) Tag() fix.Tag { return tag.UndlyInstrumentPartyID }

type UndlyInstrumentPartyIDSource struct{ message.CharValue }

func (f UndlyInstrumentPartyIDSource) Tag() fix.Tag { return tag.UndlyInstrumentPartyIDSource }

type UndlyInstrumentPartyRole struct{ message.IntValue }

func (f UndlyInstrumentPartyRole) Tag() fix.Tag { return tag.UndlyInstrumentPartyRole }

type UndlyInstrumentPartySubID struct{ message.StringValue }

func (f UndlyInstrumentPartySubID) Tag() fix.Tag { return tag.UndlyInstrumentPartySubID }

type UndlyInstrumentPartySubIDType struct{ message.IntValue }

func (f UndlyInstrumentPartySubIDType) Tag() fix.Tag { return tag.UndlyInstrumentPartySubIDType }

type UnitOfMeasure struct{ message.StringValue }

func (f UnitOfMeasure) Tag() fix.Tag { return tag.UnitOfMeasure }

type UnitOfMeasureQty struct{ message.QtyValue }

func (f UnitOfMeasureQty) Tag() fix.Tag { return tag.UnitOfMeasureQty }

type UnsolicitedIndicator struct{ message.BooleanValue }

func (f UnsolicitedIndicator) Tag() fix.Tag { return tag.UnsolicitedIndicator }

type Urgency struct{ message.CharValue }

func (f Urgency) Tag() fix.Tag { return tag.Urgency }

type UserRequestID struct{ message.StringValue }

func (f UserRequestID) Tag() fix.Tag { return tag.UserRequestID }

type UserRequestType struct{ message.IntValue }

func (f UserRequestType) Tag() fix.Tag { return tag.UserRequestType }

type UserStatus struct{ message.IntValue }

func (f UserStatus) Tag() fix.Tag { return tag.UserStatus }

type UserStatusText struct{ message.StringValue }

func (f UserStatusText) Tag() fix.Tag { return tag.UserStatusText }

type Username struct{ message.StringValue }

func (f Username) Tag() fix.Tag { return tag.Username }

type ValidUntilTime struct{ message.UTCTimestampValue }

func (f ValidUntilTime) Tag() fix.Tag { return tag.ValidUntilTime }

type ValuationMethod struct{ message.StringValue }

func (f ValuationMethod) Tag() fix.Tag { return tag.ValuationMethod }

type ValueOfFutures struct{ message.AmtValue }

func (f ValueOfFutures) Tag() fix.Tag { return tag.ValueOfFutures }

type VenueType struct{ message.CharValue }

func (f VenueType) Tag() fix.Tag { return tag.VenueType }

type Volatility struct{ message.FloatValue }

func (f Volatility) Tag() fix.Tag { return tag.Volatility }

type WaveNo struct{ message.StringValue }

func (f WaveNo) Tag() fix.Tag { return tag.WaveNo }

type WorkingIndicator struct{ message.BooleanValue }

func (f WorkingIndicator) Tag() fix.Tag { return tag.WorkingIndicator }

type WtAverageLiquidity struct{ message.PercentageValue }

func (f WtAverageLiquidity) Tag() fix.Tag { return tag.WtAverageLiquidity }

type XmlData struct{ message.DataValue }

func (f XmlData) Tag() fix.Tag { return tag.XmlData }

type XmlDataLen struct{ message.LengthValue }

func (f XmlDataLen) Tag() fix.Tag { return tag.XmlDataLen }

type Yield struct{ message.PercentageValue }

func (f Yield) Tag() fix.Tag { return tag.Yield }

type YieldCalcDate struct{ message.LocalMktDateValue }

func (f YieldCalcDate) Tag() fix.Tag { return tag.YieldCalcDate }

type YieldRedemptionDate struct{ message.LocalMktDateValue }

func (f YieldRedemptionDate) Tag() fix.Tag { return tag.YieldRedemptionDate }

type YieldRedemptionPrice struct{ message.PriceValue }

func (f YieldRedemptionPrice) Tag() fix.Tag { return tag.YieldRedemptionPrice }

type YieldRedemptionPriceType struct{ message.IntValue }

func (f YieldRedemptionPriceType) Tag() fix.Tag { return tag.YieldRedemptionPriceType }

type YieldType struct{ message.StringValue }

func (f YieldType) Tag() fix.Tag { return tag.YieldType }
