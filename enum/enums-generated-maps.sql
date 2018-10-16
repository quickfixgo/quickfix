CREATE SCHEMA IF NOT EXISTS enum_maps;

CREATE TABLE IF NOT EXISTS enum_maps.AccountType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AccountType VALUES 
('1', 'AccountType_ACCOUNT_IS_CARRIED_ON_CUSTOMER_SIDE_OF_THE_BOOKS'),('2', 'AccountType_ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS'),('3', 'AccountType_HOUSE_TRADER'),('4', 'AccountType_FLOOR_TRADER'),('6', 'AccountType_ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS_AND_IS_CROSS_MARGINED'),('7', 'AccountType_ACCOUNT_IS_HOUSE_TRADER_AND_IS_CROSS_MARGINED'),('8', 'AccountType_JOINT_BACK_OFFICE_ACCOUNT');

CREATE TABLE IF NOT EXISTS enum_maps.AcctIDSource (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AcctIDSource VALUES 
('1', 'AcctIDSource_BIC'),('2', 'AcctIDSource_SID_CODE'),('3', 'AcctIDSource_TFM'),('4', 'AcctIDSource_OMGEO'),('5', 'AcctIDSource_DTCC_CODE'),('99', 'AcctIDSource_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.Adjustment (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.Adjustment VALUES 
('1', 'Adjustment_CANCEL'),('2', 'Adjustment_ERROR'),('3', 'Adjustment_CORRECTION');

CREATE TABLE IF NOT EXISTS enum_maps.AdjustmentType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AdjustmentType VALUES 
('0', 'AdjustmentType_PROCESS_REQUEST_AS_MARGIN_DISPOSITION'),('1', 'AdjustmentType_DELTA_PLUS'),('2', 'AdjustmentType_DELTA_MINUS'),('3', 'AdjustmentType_FINAL');

CREATE TABLE IF NOT EXISTS enum_maps.AdvSide (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AdvSide VALUES 
('B', 'AdvSide_BUY'),('S', 'AdvSide_SELL'),('T', 'AdvSide_TRADE'),('X', 'AdvSide_CROSS');

CREATE TABLE IF NOT EXISTS enum_maps.AdvTransType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AdvTransType VALUES 
('C', 'AdvTransType_CANCEL'),('N', 'AdvTransType_NEW'),('R', 'AdvTransType_REPLACE');

CREATE TABLE IF NOT EXISTS enum_maps.AffirmStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AffirmStatus VALUES 
('1', 'AffirmStatus_RECEIVED'),('2', 'AffirmStatus_CONFIRM_REJECTED_IE_NOT_AFFIRMED'),('3', 'AffirmStatus_AFFIRMED');

CREATE TABLE IF NOT EXISTS enum_maps.AggregatedBook (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AggregatedBook VALUES 
('N', 'AggregatedBook_NO'),('Y', 'AggregatedBook_YES');

CREATE TABLE IF NOT EXISTS enum_maps.AggressorIndicator (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AggressorIndicator VALUES 
('N', 'AggressorIndicator_NO'),('Y', 'AggressorIndicator_YES');

CREATE TABLE IF NOT EXISTS enum_maps.AllocAccountType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AllocAccountType VALUES 
('1', 'AllocAccountType_ACCOUNT_IS_CARRIED_PN_CUSTOMER_SIDE_OF_BOOKS'),('2', 'AllocAccountType_ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS'),('3', 'AllocAccountType_HOUSE_TRADER'),('4', 'AllocAccountType_FLOOR_TRADER'),('6', 'AllocAccountType_ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS_AND_IS_CROSS_MARGINED'),('7', 'AllocAccountType_ACCOUNT_IS_HOUSE_TRADER_AND_IS_CROSS_MARGINED'),('8', 'AllocAccountType_JOINT_BACK_OFFICE_ACCOUNT');

CREATE TABLE IF NOT EXISTS enum_maps.AllocCancReplaceReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AllocCancReplaceReason VALUES 
('1', 'AllocCancReplaceReason_ORIGINAL_DETAILS_INCOMPLETE_INCORRECT'),('2', 'AllocCancReplaceReason_CHANGE_IN_UNDERLYING_ORDER_DETAILS'),('99', 'AllocCancReplaceReason_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.AllocHandlInst (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AllocHandlInst VALUES 
('1', 'AllocHandlInst_MATCH'),('2', 'AllocHandlInst_FORWARD'),('3', 'AllocHandlInst_FORWARD_AND_MATCH');

CREATE TABLE IF NOT EXISTS enum_maps.AllocIntermedReqType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AllocIntermedReqType VALUES 
('1', 'AllocIntermedReqType_PENDING_ACCEPT'),('2', 'AllocIntermedReqType_PENDING_RELEASE'),('3', 'AllocIntermedReqType_PENDING_REVERSAL'),('4', 'AllocIntermedReqType_ACCEPT'),('5', 'AllocIntermedReqType_BLOCK_LEVEL_REJECT'),('6', 'AllocIntermedReqType_ACCOUNT_LEVEL_REJECT');

CREATE TABLE IF NOT EXISTS enum_maps.AllocLinkType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AllocLinkType VALUES 
('0', 'AllocLinkType_FX_NETTING'),('1', 'AllocLinkType_FX_SWAP');

CREATE TABLE IF NOT EXISTS enum_maps.AllocMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AllocMethod VALUES 
('1', 'AllocMethod_AUTOMATIC'),('2', 'AllocMethod_GUARANTOR'),('3', 'AllocMethod_MANUAL');

CREATE TABLE IF NOT EXISTS enum_maps.AllocNoOrdersType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AllocNoOrdersType VALUES 
('0', 'AllocNoOrdersType_NOT_SPECIFIED'),('1', 'AllocNoOrdersType_EXPLICIT_LIST_PROVIDED');

CREATE TABLE IF NOT EXISTS enum_maps.AllocPositionEffect (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AllocPositionEffect VALUES 
('C', 'AllocPositionEffect_CLOSE'),('F', 'AllocPositionEffect_FIFO'),('O', 'AllocPositionEffect_OPEN'),('R', 'AllocPositionEffect_ROLLED');

CREATE TABLE IF NOT EXISTS enum_maps.AllocRejCode (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AllocRejCode VALUES 
('0', 'AllocRejCode_UNKNOWN_ACCOUNT'),('1', 'AllocRejCode_INCORRECT_QUANTITY'),('10', 'AllocRejCode_UNKNOWN_OR_STALE_EXECID'),('11', 'AllocRejCode_MISMATCHED_DATA'),('12', 'AllocRejCode_UNKNOWN_CLORDID'),('13', 'AllocRejCode_WAREHOUSE_REQUEST_REJECTED'),('2', 'AllocRejCode_INCORRECT_AVERAGEG_PRICE'),('3', 'AllocRejCode_UNKNOWN_EXECUTING_BROKER_MNEMONIC'),('4', 'AllocRejCode_COMMISSION_DIFFERENCE'),('5', 'AllocRejCode_UNKNOWN_ORDERID'),('6', 'AllocRejCode_UNKNOWN_LISTID'),('7', 'AllocRejCode_OTHER_7'),('8', 'AllocRejCode_INCORRECT_ALLOCATED_QUANTITY'),('9', 'AllocRejCode_CALCULATION_DIFFERENCE'),('99', 'AllocRejCode_OTHER_99');

CREATE TABLE IF NOT EXISTS enum_maps.AllocReportType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AllocReportType VALUES 
('10', 'AllocReportType_REJECT'),('11', 'AllocReportType_ACCEPT_PENDING'),('12', 'AllocReportType_COMPLETE'),('14', 'AllocReportType_REVERSE_PENDING'),('2', 'AllocReportType_PRELIMINARY_REQUEST_TO_INTERMEDIARY'),('3', 'AllocReportType_SELLSIDE_CALCULATED_USING_PRELIMINARY'),('4', 'AllocReportType_SELLSIDE_CALCULATED_WITHOUT_PRELIMINARY'),('5', 'AllocReportType_WAREHOUSE_RECAP'),('8', 'AllocReportType_REQUEST_TO_INTERMEDIARY'),('9', 'AllocReportType_ACCEPT');

CREATE TABLE IF NOT EXISTS enum_maps.AllocSettlInstType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AllocSettlInstType VALUES 
('0', 'AllocSettlInstType_USE_DEFAULT_INSTRUCTIONS'),('1', 'AllocSettlInstType_DERIVE_FROM_PARAMETERS_PROVIDED'),('2', 'AllocSettlInstType_FULL_DETAILS_PROVIDED'),('3', 'AllocSettlInstType_SSI_DB_IDS_PROVIDED'),('4', 'AllocSettlInstType_PHONE_FOR_INSTRUCTIONS');

CREATE TABLE IF NOT EXISTS enum_maps.AllocStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AllocStatus VALUES 
('0', 'AllocStatus_ACCEPTED'),('1', 'AllocStatus_BLOCK_LEVEL_REJECT'),('2', 'AllocStatus_ACCOUNT_LEVEL_REJECT'),('3', 'AllocStatus_RECEIVED'),('4', 'AllocStatus_INCOMPLETE'),('5', 'AllocStatus_REJECTED_BY_INTERMEDIARY'),('6', 'AllocStatus_ALLOCATION_PENDING'),('7', 'AllocStatus_REVERSED');

CREATE TABLE IF NOT EXISTS enum_maps.AllocTransType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AllocTransType VALUES 
('0', 'AllocTransType_NEW'),('1', 'AllocTransType_REPLACE'),('2', 'AllocTransType_CANCEL'),('3', 'AllocTransType_PRELIMINARY'),('4', 'AllocTransType_CALCULATED'),('5', 'AllocTransType_CALCULATED_WITHOUT_PRELIMINARY'),('6', 'AllocTransType_REVERSAL');

CREATE TABLE IF NOT EXISTS enum_maps.AllocType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AllocType VALUES 
('1', 'AllocType_CALCULATED'),('10', 'AllocType_REJECT'),('11', 'AllocType_ACCEPT_PENDING'),('12', 'AllocType_INCOMPLETE_GROUP'),('13', 'AllocType_COMPLETE_GROUP'),('14', 'AllocType_REVERSAL_PENDING'),('2', 'AllocType_PRELIMINARY'),('3', 'AllocType_SELLSIDE_CALCULATED_USING_PRELIMINARY'),('4', 'AllocType_SELLSIDE_CALCULATED_WITHOUT_PRELIMINARY'),('5', 'AllocType_READY_TO_BOOK'),('6', 'AllocType_BUYSIDE_READY_TO_BOOK'),('7', 'AllocType_WAREHOUSE_INSTRUCTION'),('8', 'AllocType_REQUEST_TO_INTERMEDIARY'),('9', 'AllocType_ACCEPT');

CREATE TABLE IF NOT EXISTS enum_maps.ApplQueueAction (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ApplQueueAction VALUES 
('0', 'ApplQueueAction_NO_ACTION_TAKEN'),('1', 'ApplQueueAction_QUEUE_FLUSHED'),('2', 'ApplQueueAction_OVERLAY_LAST'),('3', 'ApplQueueAction_END_SESSION');

CREATE TABLE IF NOT EXISTS enum_maps.ApplQueueResolution (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ApplQueueResolution VALUES 
('0', 'ApplQueueResolution_NO_ACTION_TAKEN'),('1', 'ApplQueueResolution_QUEUE_FLUSHED'),('2', 'ApplQueueResolution_OVERLAY_LAST'),('3', 'ApplQueueResolution_END_SESSION');

CREATE TABLE IF NOT EXISTS enum_maps.ApplReportType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ApplReportType VALUES 
('0', 'ApplReportType_RESET_APPLSEQNUM_TO_NEW_VALUE_SPECIFIED_IN_APPLNEWSEQNUM'),('1', 'ApplReportType_REPORTS_THAT_THE_LAST_MESSAGE_HAS_BEEN_SENT_FOR_THE_APPLIDS_REFER_TO_REFAPPLLASTSEQNUM'),('2', 'ApplReportType_HEARTBEAT_MESSAGE_INDICATING_THAT_APPLICATION_IDENTIFIED_BY_REFAPPLID'),('3', 'ApplReportType_APPLICATION_MESSAGE_RE_SEND_COMPLETED');

CREATE TABLE IF NOT EXISTS enum_maps.ApplReqType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ApplReqType VALUES 
('0', 'ApplReqType_RETRANSMISSION_OF_APPLICATION_MESSAGES_FOR_THE_SPECIFIED_APPLICATIONS'),('1', 'ApplReqType_SUBSCRIPTION_TO_THE_SPECIFIED_APPLICATIONS'),('2', 'ApplReqType_REQUEST_FOR_THE_LAST_APPLLASTSEQNUM_PUBLISHED_FOR_THE_SPECIFIED_APPLICATIONS'),('3', 'ApplReqType_REQUEST_VALID_SET_OF_APPLICATIONS'),('4', 'ApplReqType_UNSUBSCRIBE_TO_THE_SPECIFIED_APPLICATIONS'),('5', 'ApplReqType_CANCEL_RETRANSMISSION'),('6', 'ApplReqType_CANCEL_RETRANSMISSION_AND_UNSUBSCRIBE_TO_THE_SPECIFIED_APPLICATIONS');

CREATE TABLE IF NOT EXISTS enum_maps.ApplResponseError (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ApplResponseError VALUES 
('0', 'ApplResponseError_APPLICATION_DOES_NOT_EXIST'),('1', 'ApplResponseError_MESSAGES_REQUESTED_ARE_NOT_AVAILABLE'),('2', 'ApplResponseError_USER_NOT_AUTHORIZED_FOR_APPLICATION');

CREATE TABLE IF NOT EXISTS enum_maps.ApplResponseType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ApplResponseType VALUES 
('0', 'ApplResponseType_REQUEST_SUCCESSFULLY_PROCESSED'),('1', 'ApplResponseType_APPLICATION_DOES_NOT_EXIST'),('2', 'ApplResponseType_MESSAGES_NOT_AVAILABLE');

CREATE TABLE IF NOT EXISTS enum_maps.ApplVerID (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ApplVerID VALUES 
('0', 'ApplVerID_FIX27'),('1', 'ApplVerID_FIX30'),('2', 'ApplVerID_FIX40'),('3', 'ApplVerID_FIX41'),('4', 'ApplVerID_FIX42'),('5', 'ApplVerID_FIX43'),('6', 'ApplVerID_FIX44'),('7', 'ApplVerID_FIX50'),('8', 'ApplVerID_FIX50SP1'),('9', 'ApplVerID_FIX50SP2');

CREATE TABLE IF NOT EXISTS enum_maps.AsOfIndicator (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AsOfIndicator VALUES 
('0', 'AsOfIndicator_FALSE'),('1', 'AsOfIndicator_TRUE');

CREATE TABLE IF NOT EXISTS enum_maps.AssignmentMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AssignmentMethod VALUES 
('P', 'AssignmentMethod_PRO_RATA'),('R', 'AssignmentMethod_RANDOM');

CREATE TABLE IF NOT EXISTS enum_maps.AvgPxIndicator (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.AvgPxIndicator VALUES 
('0', 'AvgPxIndicator_NO_AVERAGE_PRICING'),('1', 'AvgPxIndicator_TRADE_IS_PART_OF_AN_AVERAGE_PRICE_GROUP_IDENTIFIED_BY_THE_TRADELINKID'),('2', 'AvgPxIndicator_LAST_TRADE_IS_THE_AVERAGE_PRICE_GROUP_IDENTIFIED_BY_THE_TRADELINKID');

CREATE TABLE IF NOT EXISTS enum_maps.BasisPxType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.BasisPxType VALUES 
('2', 'BasisPxType_CLOSING_PRICE_AT_MORNINGN_SESSION'),('3', 'BasisPxType_CLOSING_PRICE'),('4', 'BasisPxType_CURRENT_PRICE'),('5', 'BasisPxType_SQ'),('6', 'BasisPxType_VWAP_THROUGH_A_DAY'),('7', 'BasisPxType_VWAP_THROUGH_A_MORNING_SESSION'),('8', 'BasisPxType_VWAP_THROUGH_AN_AFTERNOON_SESSION'),('9', 'BasisPxType_VWAP_THROUGH_A_DAY_EXCEPT_YORI'),('A', 'BasisPxType_VWAP_THROUGH_A_MORNING_SESSION_EXCEPT_YORI'),('B', 'BasisPxType_VWAP_THROUGH_AN_AFTERNOON_SESSION_EXCEPT_YORI'),('C', 'BasisPxType_STRIKE'),('D', 'BasisPxType_OPEN'),('Z', 'BasisPxType_OTHERS');

CREATE TABLE IF NOT EXISTS enum_maps.Benchmark (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.Benchmark VALUES 
('1', 'Benchmark_CURVE'),('2', 'Benchmark_5YR'),('3', 'Benchmark_OLD5'),('4', 'Benchmark_10YR'),('5', 'Benchmark_OLD10'),('6', 'Benchmark_30YR'),('7', 'Benchmark_OLD30'),('8', 'Benchmark_3MOLIBOR'),('9', 'Benchmark_6MOLIBOR');

CREATE TABLE IF NOT EXISTS enum_maps.BenchmarkCurveName (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.BenchmarkCurveName VALUES 
('EONIA', 'BenchmarkCurveName_EONIA'),('EUREPO', 'BenchmarkCurveName_EUREPO'),('Euribor', 'BenchmarkCurveName_EURIBOR'),('FutureSWAP', 'BenchmarkCurveName_FUTURESWAP'),('LIBID', 'BenchmarkCurveName_LIBID'),('LIBOR', 'BenchmarkCurveName_LIBOR'),('MuniAAA', 'BenchmarkCurveName_MUNIAAA'),('OTHER', 'BenchmarkCurveName_OTHER'),('Pfandbriefe', 'BenchmarkCurveName_PFANDBRIEFE'),('SONIA', 'BenchmarkCurveName_SONIA'),('SWAP', 'BenchmarkCurveName_SWAP'),('Treasury', 'BenchmarkCurveName_TREASURY');

CREATE TABLE IF NOT EXISTS enum_maps.BidDescriptorType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.BidDescriptorType VALUES 
('1', 'BidDescriptorType_SECTOR'),('2', 'BidDescriptorType_COUNTRY'),('3', 'BidDescriptorType_INDEX');

CREATE TABLE IF NOT EXISTS enum_maps.BidRequestTransType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.BidRequestTransType VALUES 
('C', 'BidRequestTransType_CANCEL'),('N', 'BidRequestTransType_NO');

CREATE TABLE IF NOT EXISTS enum_maps.BidTradeType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.BidTradeType VALUES 
('A', 'BidTradeType_AGENCY'),('G', 'BidTradeType_VWAP_GUARANTEE'),('J', 'BidTradeType_GUARANTEED_CLOSE'),('R', 'BidTradeType_RISK_TRADE');

CREATE TABLE IF NOT EXISTS enum_maps.BidType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.BidType VALUES 
('1', 'BidType_NON_DISCLOSED_STYLE'),('2', 'BidType_DISCLOSED_SYTLE'),('3', 'BidType_NO_BIDDING_PROCESS');

CREATE TABLE IF NOT EXISTS enum_maps.BookingType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.BookingType VALUES 
('0', 'BookingType_REGULAR_BOOKING'),('1', 'BookingType_CFD'),('2', 'BookingType_TOTAL_RETURN_SWAP');

CREATE TABLE IF NOT EXISTS enum_maps.BookingUnit (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.BookingUnit VALUES 
('0', 'BookingUnit_EACH_PARTIAL_EXECUTION_IS_A_BOOKABLE_UNIT'),('1', 'BookingUnit_AGGREGATE_PARTIAL_EXECUTIONS_ON_THIS_ORDER_AND_BOOK_ONE_TRADE_PER_ORDER'),('2', 'BookingUnit_AGGREGATE_EXECUTIONS_FOR_THIS_SYMBOL_SIDE_AND_SETTLEMENT_DATE');

CREATE TABLE IF NOT EXISTS enum_maps.BusinessRejectReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.BusinessRejectReason VALUES 
('0', 'BusinessRejectReason_OTHER'),('1', 'BusinessRejectReason_UNKNOWN_ID'),('18', 'BusinessRejectReason_INVALID_PRICE_INCREMENT'),('2', 'BusinessRejectReason_UNKNOWN_SECURITY'),('3', 'BusinessRejectReason_UNSUPPORTED_MESSAGE_TYPE'),('4', 'BusinessRejectReason_APPLICATION_NOT_AVAILABLE'),('5', 'BusinessRejectReason_CONDITIONALLY_REQUIRED_FIELD_MISSING'),('6', 'BusinessRejectReason_NOT_AUTHORIZED'),('7', 'BusinessRejectReason_DELIVERTO_FIRM_NOT_AVAILABLE_AT_THIS_TIME');

CREATE TABLE IF NOT EXISTS enum_maps.CPProgram (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CPProgram VALUES 
('1', 'CPProgram_3'),('2', 'CPProgram_4'),('99', 'CPProgram_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.CancellationRights (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CancellationRights VALUES 
('M', 'CancellationRights_NO_M'),('N', 'CancellationRights_NO_N'),('O', 'CancellationRights_NO_O'),('Y', 'CancellationRights_YES');

CREATE TABLE IF NOT EXISTS enum_maps.CashMargin (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CashMargin VALUES 
('1', 'CashMargin_CASH'),('2', 'CashMargin_MARGIN_OPEN'),('3', 'CashMargin_MARGIN_CLOSE');

CREATE TABLE IF NOT EXISTS enum_maps.ClearingFeeIndicator (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ClearingFeeIndicator VALUES 
('1', 'ClearingFeeIndicator_1ST_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT'),('2', 'ClearingFeeIndicator_2ND_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT'),('3', 'ClearingFeeIndicator_3RD_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT'),('4', 'ClearingFeeIndicator_4TH_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT'),('5', 'ClearingFeeIndicator_5TH_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT'),('9', 'ClearingFeeIndicator_6TH_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT'),('B', 'ClearingFeeIndicator_CBOE_MEMBER'),('C', 'ClearingFeeIndicator_NON_MEMBER_AND_CUSTOMER'),('E', 'ClearingFeeIndicator_EQUITY_MEMBER_AND_CLEARING_MEMBER'),('F', 'ClearingFeeIndicator_FULL_AND_ASSOCIATE_MEMBER_TRADING_FOR_OWN_ACCOUNT_AND_AS_FLOOR_BROKERS'),('H', 'ClearingFeeIndicator_106H_AND_106J_FIRMS'),('I', 'ClearingFeeIndicator_GIM_IDEM_AND_COM_MEMBERSHIP_INTEREST_HOLDERS'),('L', 'ClearingFeeIndicator_LESSEE_106F_EMPLOYEES'),('M', 'ClearingFeeIndicator_ALL_OTHER_OWNERSHIP_TYPES');

CREATE TABLE IF NOT EXISTS enum_maps.ClearingInstruction (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ClearingInstruction VALUES 
('0', 'ClearingInstruction_PROCESS_NORMALLY'),('1', 'ClearingInstruction_EXCLUDE_FROM_ALL_NETTING'),('10', 'ClearingInstruction_AUTOMATIC_GIVE_UP_MODE'),('11', 'ClearingInstruction_QUALIFIED_SERVICE_REPRESENTATIVE_QSR'),('12', 'ClearingInstruction_CUSTOMER_TRADE'),('13', 'ClearingInstruction_SELF_CLEARING'),('2', 'ClearingInstruction_BILATERAL_NETTING_ONLY'),('3', 'ClearingInstruction_EX_CLEARING'),('4', 'ClearingInstruction_SPECIAL_TRADE'),('5', 'ClearingInstruction_MULTILATERAL_NETTING'),('6', 'ClearingInstruction_CLEAR_AGAINST_CENTRAL_COUNTERPARTY'),('7', 'ClearingInstruction_EXCLUDE_FROM_CENTRAL_COUNTERPARTY'),('8', 'ClearingInstruction_MANUAL_MODE'),('9', 'ClearingInstruction_AUTOMATIC_POSTING_MODE');

CREATE TABLE IF NOT EXISTS enum_maps.CollAction (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CollAction VALUES 
('0', 'CollAction_RETAIN'),('1', 'CollAction_ADD'),('2', 'CollAction_REMOVE');

CREATE TABLE IF NOT EXISTS enum_maps.CollApplType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CollApplType VALUES 
('0', 'CollApplType_SPECIFIC_DEPOSIT'),('1', 'CollApplType_GENERAL');

CREATE TABLE IF NOT EXISTS enum_maps.CollAsgnReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CollAsgnReason VALUES 
('0', 'CollAsgnReason_INITIAL'),('1', 'CollAsgnReason_SCHEDULED'),('2', 'CollAsgnReason_TIME_WARNING'),('3', 'CollAsgnReason_MARGIN_DEFICIENCY'),('4', 'CollAsgnReason_MARGIN_EXCESS'),('5', 'CollAsgnReason_FORWARD_COLLATERAL_DEMAND'),('6', 'CollAsgnReason_EVENT_OF_DEFAULT'),('7', 'CollAsgnReason_ADVERSE_TAX_EVENT');

CREATE TABLE IF NOT EXISTS enum_maps.CollAsgnRejectReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CollAsgnRejectReason VALUES 
('0', 'CollAsgnRejectReason_UNKNOWN_DEAL'),('1', 'CollAsgnRejectReason_UNKNOWN_OR_INVALID_INSTRUMENT'),('2', 'CollAsgnRejectReason_UNAUTHORIZED_TRANSACTION'),('3', 'CollAsgnRejectReason_INSUFFICIENT_COLLATERAL'),('4', 'CollAsgnRejectReason_INVALID_TYPE_OF_COLLATERAL'),('5', 'CollAsgnRejectReason_EXCESSIVE_SUBSTITUTION'),('99', 'CollAsgnRejectReason_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.CollAsgnRespType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CollAsgnRespType VALUES 
('0', 'CollAsgnRespType_RECEIVED'),('1', 'CollAsgnRespType_ACCEPTED'),('2', 'CollAsgnRespType_DECLINED'),('3', 'CollAsgnRespType_REJECTED');

CREATE TABLE IF NOT EXISTS enum_maps.CollAsgnTransType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CollAsgnTransType VALUES 
('0', 'CollAsgnTransType_NEW'),('1', 'CollAsgnTransType_REPLACE'),('2', 'CollAsgnTransType_CANCEL'),('3', 'CollAsgnTransType_RELEASE'),('4', 'CollAsgnTransType_REVERSE');

CREATE TABLE IF NOT EXISTS enum_maps.CollInquiryQualifier (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CollInquiryQualifier VALUES 
('0', 'CollInquiryQualifier_TRADE_DATE'),('1', 'CollInquiryQualifier_GC_INSTRUMENT'),('2', 'CollInquiryQualifier_COLLATERAL_INSTRUMENT'),('3', 'CollInquiryQualifier_SUBSTITUTION_ELIGIBLE'),('4', 'CollInquiryQualifier_NOT_ASSIGNED'),('5', 'CollInquiryQualifier_PARTIALLY_ASSIGNED'),('6', 'CollInquiryQualifier_FULLY_ASSIGNED'),('7', 'CollInquiryQualifier_OUTSTANDING_TRADES');

CREATE TABLE IF NOT EXISTS enum_maps.CollInquiryResult (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CollInquiryResult VALUES 
('0', 'CollInquiryResult_SUCCESSFUL'),('1', 'CollInquiryResult_INVALID_OR_UNKNOWN_INSTRUMENT'),('2', 'CollInquiryResult_INVALID_OR_UNKNOWN_COLLATERAL_TYPE'),('3', 'CollInquiryResult_INVALID_PARTIES'),('4', 'CollInquiryResult_INVALID_TRANSPORT_TYPE_REQUESTED'),('5', 'CollInquiryResult_INVALID_DESTINATION_REQUESTED'),('6', 'CollInquiryResult_NO_COLLATERAL_FOUND_FOR_THE_TRADE_SPECIFIED'),('7', 'CollInquiryResult_NO_COLLATERAL_FOUND_FOR_THE_ORDER_SPECIFIED'),('8', 'CollInquiryResult_COLLATERAL_INQUIRY_TYPE_NOT_SUPPORTED'),('9', 'CollInquiryResult_UNAUTHORIZED_FOR_COLLATERAL_INQUIRY'),('99', 'CollInquiryResult_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.CollInquiryStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CollInquiryStatus VALUES 
('0', 'CollInquiryStatus_ACCEPTED'),('1', 'CollInquiryStatus_ACCEPTED_WITH_WARNINGS'),('2', 'CollInquiryStatus_COMPLETED'),('3', 'CollInquiryStatus_COMPLETED_WITH_WARNINGS'),('4', 'CollInquiryStatus_REJECTED');

CREATE TABLE IF NOT EXISTS enum_maps.CollStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CollStatus VALUES 
('0', 'CollStatus_UNASSIGNED'),('1', 'CollStatus_PARTIALLY_ASSIGNED'),('2', 'CollStatus_ASSIGNMENT_PROPOSED'),('3', 'CollStatus_ASSIGNED'),('4', 'CollStatus_CHALLENGED');

CREATE TABLE IF NOT EXISTS enum_maps.CommType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CommType VALUES 
('1', 'CommType_PER_UNIT'),('2', 'CommType_PERCENT'),('3', 'CommType_ABSOLUTE'),('4', 'CommType_PERCENTAGE_WAIVED_4'),('5', 'CommType_PERCENTAGE_WAIVED_5'),('6', 'CommType_POINTS_PER_BOND_OR_CONTRACT');

CREATE TABLE IF NOT EXISTS enum_maps.ComplexEventCondition (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ComplexEventCondition VALUES 
('1', 'ComplexEventCondition_AND'),('2', 'ComplexEventCondition_OR');

CREATE TABLE IF NOT EXISTS enum_maps.ComplexEventPriceBoundaryMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ComplexEventPriceBoundaryMethod VALUES 
('1', 'ComplexEventPriceBoundaryMethod_LESS_THAN_COMPLEXEVENTPRICE'),('2', 'ComplexEventPriceBoundaryMethod_LESS_THAN_OR_EQUAL_TO_COMPLEXEVENTPRICE'),('3', 'ComplexEventPriceBoundaryMethod_EQUAL_TO_COMPLEXEVENTPRICE'),('4', 'ComplexEventPriceBoundaryMethod_GREATER_THAN_OR_EQUAL_TO_COMPLEXEVENTPRICE'),('5', 'ComplexEventPriceBoundaryMethod_GREATER_THAN_COMPLEXEVENTPRICE');

CREATE TABLE IF NOT EXISTS enum_maps.ComplexEventPriceTimeType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ComplexEventPriceTimeType VALUES 
('1', 'ComplexEventPriceTimeType_EXPIRATION'),('2', 'ComplexEventPriceTimeType_IMMEDIATE'),('3', 'ComplexEventPriceTimeType_SPECIFIED_DATE_TIME');

CREATE TABLE IF NOT EXISTS enum_maps.ComplexEventType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ComplexEventType VALUES 
('1', 'ComplexEventType_CAPPED'),('2', 'ComplexEventType_TRIGGER'),('3', 'ComplexEventType_KNOCK_IN_UP'),('4', 'ComplexEventType_KOCK_IN_DOWN'),('5', 'ComplexEventType_KNOCK_OUT_UP'),('6', 'ComplexEventType_KNOCK_OUT_DOWN'),('7', 'ComplexEventType_UNDERLYING'),('8', 'ComplexEventType_RESET_BARRIER'),('9', 'ComplexEventType_ROLLING_BARRIER');

CREATE TABLE IF NOT EXISTS enum_maps.ConfirmRejReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ConfirmRejReason VALUES 
('1', 'ConfirmRejReason_MISMATCHED_ACCOUNT'),('2', 'ConfirmRejReason_MISSING_SETTLEMENT_INSTRUCTIONS'),('99', 'ConfirmRejReason_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.ConfirmStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ConfirmStatus VALUES 
('1', 'ConfirmStatus_RECEIVED'),('2', 'ConfirmStatus_MISMATCHED_ACCOUNT'),('3', 'ConfirmStatus_MISSING_SETTLEMENT_INSTRUCTIONS'),('4', 'ConfirmStatus_CONFIRMED'),('5', 'ConfirmStatus_REQUEST_REJECTED');

CREATE TABLE IF NOT EXISTS enum_maps.ConfirmTransType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ConfirmTransType VALUES 
('0', 'ConfirmTransType_NEW'),('1', 'ConfirmTransType_REPLACE'),('2', 'ConfirmTransType_CANCEL');

CREATE TABLE IF NOT EXISTS enum_maps.ConfirmType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ConfirmType VALUES 
('1', 'ConfirmType_STATUS'),('2', 'ConfirmType_CONFIRMATION'),('3', 'ConfirmType_CONFIRMATION_REQUEST_REJECTED');

CREATE TABLE IF NOT EXISTS enum_maps.ContAmtType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ContAmtType VALUES 
('1', 'ContAmtType_COMMISSION_AMOUNT'),('10', 'ContAmtType_EXIT_CHARGE_PERCENT'),('11', 'ContAmtType_FUND_BASED_RENEWAL_COMMISSION_PERCENT'),('12', 'ContAmtType_PROJECTED_FUND_VALUE'),('13', 'ContAmtType_FUND_BASED_RENEWAL_COMMISSION_AMOUNT_13'),('14', 'ContAmtType_FUND_BASED_RENEWAL_COMMISSION_AMOUNT_14'),('15', 'ContAmtType_NET_SETTLEMENT_AMOUNT'),('2', 'ContAmtType_COMMISSION_PERCENT'),('3', 'ContAmtType_INITIAL_CHARGE_AMOUNT'),('4', 'ContAmtType_INITIAL_CHARGE_PERCENT'),('5', 'ContAmtType_DISCOUNT_AMOUNT'),('6', 'ContAmtType_DISCOUNT_PERCENT'),('7', 'ContAmtType_DILUTION_LEVY_AMOUNT'),('8', 'ContAmtType_DILUTION_LEVY_PERCENT'),('9', 'ContAmtType_EXIT_CHARGE_AMOUNT');

CREATE TABLE IF NOT EXISTS enum_maps.ContingencyType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ContingencyType VALUES 
('1', 'ContingencyType_ONE_CANCELS_THE_OTHER'),('2', 'ContingencyType_ONE_TRIGGERS_THE_OTHER'),('3', 'ContingencyType_ONE_UPDATES_THE_OTHER_3'),('4', 'ContingencyType_ONE_UPDATES_THE_OTHER_4');

CREATE TABLE IF NOT EXISTS enum_maps.ContractMultiplierUnit (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ContractMultiplierUnit VALUES 
('0', 'ContractMultiplierUnit_SHARES'),('1', 'ContractMultiplierUnit_HOURS'),('2', 'ContractMultiplierUnit_DAYS');

CREATE TABLE IF NOT EXISTS enum_maps.CorporateAction (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CorporateAction VALUES 
('A', 'CorporateAction_EX_DIVIDEND'),('B', 'CorporateAction_EX_DISTRIBUTION'),('C', 'CorporateAction_EX_RIGHTS'),('D', 'CorporateAction_NEW'),('E', 'CorporateAction_EX_INTEREST'),('F', 'CorporateAction_CASH_DIVIDEND'),('G', 'CorporateAction_STOCK_DIVIDEND'),('H', 'CorporateAction_NON_INTEGER_STOCK_SPLIT'),('I', 'CorporateAction_REVERSE_STOCK_SPLIT'),('J', 'CorporateAction_STANDARD_INTEGER_STOCK_SPLIT'),('K', 'CorporateAction_POSITION_CONSOLIDATION'),('L', 'CorporateAction_LIQUIDATION_REORGANIZATION'),('M', 'CorporateAction_MERGER_REORGANIZATION'),('N', 'CorporateAction_RIGHTS_OFFERING'),('O', 'CorporateAction_SHAREHOLDER_MEETING'),('P', 'CorporateAction_SPINOFF'),('Q', 'CorporateAction_TENDER_OFFER'),('R', 'CorporateAction_WARRANT'),('S', 'CorporateAction_SPECIAL_ACTION'),('T', 'CorporateAction_SYMBOL_CONVERSION'),('U', 'CorporateAction_CUSIP'),('V', 'CorporateAction_LEAP_ROLLOVER'),('W', 'CorporateAction_SUCCESSION_EVENT');

CREATE TABLE IF NOT EXISTS enum_maps.CoveredOrUncovered (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CoveredOrUncovered VALUES 
('0', 'CoveredOrUncovered_COVERED'),('1', 'CoveredOrUncovered_UNCOVERED');

CREATE TABLE IF NOT EXISTS enum_maps.CrossPrioritization (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CrossPrioritization VALUES 
('0', 'CrossPrioritization_NONE'),('1', 'CrossPrioritization_BUY_SIDE_IS_PRIORITIZED'),('2', 'CrossPrioritization_SELL_SIDE_IS_PRIORITIZED');

CREATE TABLE IF NOT EXISTS enum_maps.CrossType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CrossType VALUES 
('1', 'CrossType_CROSS_AON'),('2', 'CrossType_CROSS_IOC'),('3', 'CrossType_CROSS_ONE_SIDE'),('4', 'CrossType_CROSS_SAME_PRICE');

CREATE TABLE IF NOT EXISTS enum_maps.CustOrderCapacity (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CustOrderCapacity VALUES 
('1', 'CustOrderCapacity_MEMBER_TRADING_FOR_THEIR_OWN_ACCOUNT'),('2', 'CustOrderCapacity_CLEARING_FIRM_TRADING_FOR_ITS_PROPRIETARY_ACCOUNT'),('3', 'CustOrderCapacity_MEMBER_TRADING_FOR_ANOTHER_MEMBER'),('4', 'CustOrderCapacity_ALL_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.CustOrderHandlingInst (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CustOrderHandlingInst VALUES 
('ADD', 'CustOrderHandlingInst_ADD_ON_ORDER'),('AON', 'CustOrderHandlingInst_ALL_OR_NONE'),('CNH', 'CustOrderHandlingInst_CASH_NOT_HELD'),('DIR', 'CustOrderHandlingInst_DIRECTED_ORDER'),('E.W', 'CustOrderHandlingInst_EXCHANGE_FOR_PHYSICAL_TRANSACTION'),('FOK', 'CustOrderHandlingInst_FILL_OR_KILL'),('IO', 'CustOrderHandlingInst_IMBALANCE_ONLY'),('IOC', 'CustOrderHandlingInst_IMMEDIATE_OR_CANCEL'),('LOC', 'CustOrderHandlingInst_LIMIT_ON_CLOSE'),('LOO', 'CustOrderHandlingInst_LIMIT_ON_OPEN'),('MAC', 'CustOrderHandlingInst_MARKET_AT_CLOSE'),('MAO', 'CustOrderHandlingInst_MARKET_AT_OPEN'),('MOC', 'CustOrderHandlingInst_MARKET_ON_CLOSE'),('MOO', 'CustOrderHandlingInst_MARKET_ON_OPEN'),('MQT', 'CustOrderHandlingInst_MINIMUM_QUANTITY'),('NH', 'CustOrderHandlingInst_NOT_HELD'),('OVD', 'CustOrderHandlingInst_OVER_THE_DAY'),('PEG', 'CustOrderHandlingInst_PEGGED'),('RSV', 'CustOrderHandlingInst_RESERVE_SIZE_ORDER'),('S.W', 'CustOrderHandlingInst_STOP_STOCK_TRANSACTION'),('SCL', 'CustOrderHandlingInst_SCALE'),('TMO', 'CustOrderHandlingInst_TIME_ORDER'),('TS', 'CustOrderHandlingInst_TRAILING_STOP'),('WRK', 'CustOrderHandlingInst_WORK');

CREATE TABLE IF NOT EXISTS enum_maps.CustomerOrFirm (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CustomerOrFirm VALUES 
('0', 'CustomerOrFirm_CUSTOMER'),('1', 'CustomerOrFirm_FIRM');

CREATE TABLE IF NOT EXISTS enum_maps.CxlRejReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CxlRejReason VALUES 
('0', 'CxlRejReason_TOO_LATE_TO_CANCEL'),('1', 'CxlRejReason_UNKNOWN_ORDER'),('18', 'CxlRejReason_INVALID_PRICE_INCREMENT'),('2', 'CxlRejReason_BROKER'),('3', 'CxlRejReason_ORDER_ALREADY_IN_PENDING_CANCEL_OR_PENDING_REPLACE_STATUS'),('4', 'CxlRejReason_UNABLE_TO_PROCESS_ORDER_MASS_CANCEL_REQUEST'),('5', 'CxlRejReason_ORIGORDMODTIME'),('6', 'CxlRejReason_DUPLICATE_CLORDID'),('7', 'CxlRejReason_PRICE_EXCEEDS_CURRENT_PRICE'),('8', 'CxlRejReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND'),('99', 'CxlRejReason_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.CxlRejResponseTo (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CxlRejResponseTo VALUES 
('1', 'CxlRejResponseTo_ORDER_CANCEL_REQUEST'),('2', 'CxlRejResponseTo_ORDER_CANCEL_REPLACE_REQUEST');

CREATE TABLE IF NOT EXISTS enum_maps.CxlType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.CxlType VALUES 
('F', 'CxlType_FULL_REMAINING_QUANTITY'),('P', 'CxlType_PARTIAL_CANCEL');

CREATE TABLE IF NOT EXISTS enum_maps.DKReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DKReason VALUES 
('A', 'DKReason_UNKNOWN_SYMBOL'),('B', 'DKReason_WRONG_SIDE'),('C', 'DKReason_QUANTITY_EXCEEDS_ORDER'),('D', 'DKReason_NO_MATCHING_ORDER'),('E', 'DKReason_PRICE_EXCEEDS_LIMIT'),('F', 'DKReason_CALCULATION_DIFFERENCE'),('Z', 'DKReason_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.DayBookingInst (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DayBookingInst VALUES 
('0', 'DayBookingInst_CAN_TRIGGER_BOOKING_WITHOUT_REFERENCE_TO_THE_ORDER_INITIATOR'),('1', 'DayBookingInst_SPEAK_WITH_ORDER_INITIATOR_BEFORE_BOOKING'),('2', 'DayBookingInst_ACCUMULATE');

CREATE TABLE IF NOT EXISTS enum_maps.DealingCapacity (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DealingCapacity VALUES 
('A', 'DealingCapacity_AGENT'),('P', 'DealingCapacity_PRINCIPAL'),('R', 'DealingCapacity_RISKLESS_PRINCIPAL');

CREATE TABLE IF NOT EXISTS enum_maps.DeleteReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DeleteReason VALUES 
('0', 'DeleteReason_CANCELLATION'),('1', 'DeleteReason_ERROR');

CREATE TABLE IF NOT EXISTS enum_maps.DeliveryForm (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DeliveryForm VALUES 
('1', 'DeliveryForm_BOOK_ENTRY'),('2', 'DeliveryForm_BEARER');

CREATE TABLE IF NOT EXISTS enum_maps.DeliveryType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DeliveryType VALUES 
('0', 'DeliveryType_VERSUS_PAYMENT_DELIVER'),('1', 'DeliveryType_FREE_DELIVER'),('2', 'DeliveryType_TRI_PARTY'),('3', 'DeliveryType_HOLD_IN_CUSTODY');

CREATE TABLE IF NOT EXISTS enum_maps.DerivativeSecurityListRequestType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DerivativeSecurityListRequestType VALUES 
('0', 'DerivativeSecurityListRequestType_SYMBOL'),('1', 'DerivativeSecurityListRequestType_SECURITYTYPE_AND_OR_CFICODE'),('2', 'DerivativeSecurityListRequestType_PRODUCT'),('3', 'DerivativeSecurityListRequestType_TRADINGSESSIONID'),('4', 'DerivativeSecurityListRequestType_ALL_SECURITIES'),('5', 'DerivativeSecurityListRequestType_UNDELYINGSYMBOL'),('6', 'DerivativeSecurityListRequestType_UNDERLYING_SECURITYTYPE_AND_OR_CFICODE'),('7', 'DerivativeSecurityListRequestType_UNDERLYING_PRODUCT'),('8', 'DerivativeSecurityListRequestType_MARKETID_OR_MARKETID_PLUS_MARKETSEGMENTID');

CREATE TABLE IF NOT EXISTS enum_maps.DeskOrderHandlingInst (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DeskOrderHandlingInst VALUES 
('ADD', 'DeskOrderHandlingInst_ADD_ON_ORDER'),('AON', 'DeskOrderHandlingInst_ALL_OR_NONE'),('CNH', 'DeskOrderHandlingInst_CASH_NOT_HELD'),('DIR', 'DeskOrderHandlingInst_DIRECTED_ORDER'),('E.W', 'DeskOrderHandlingInst_EXCHANGE_FOR_PHYSICAL_TRANSACTION'),('FOK', 'DeskOrderHandlingInst_FILL_OR_KILL'),('IO', 'DeskOrderHandlingInst_IMBALANCE_ONLY'),('IOC', 'DeskOrderHandlingInst_IMMEDIATE_OR_CANCEL'),('LOC', 'DeskOrderHandlingInst_LIMIT_ON_CLOSE'),('LOO', 'DeskOrderHandlingInst_LIMIT_ON_OPEN'),('MAC', 'DeskOrderHandlingInst_MARKET_AT_CLOSE'),('MAO', 'DeskOrderHandlingInst_MARKET_AT_OPEN'),('MOC', 'DeskOrderHandlingInst_MARKET_ON_CLOSE'),('MOO', 'DeskOrderHandlingInst_MARKET_ON_OPEN'),('MQT', 'DeskOrderHandlingInst_MINIMUM_QUANTITY'),('NH', 'DeskOrderHandlingInst_NOT_HELD'),('OVD', 'DeskOrderHandlingInst_OVER_THE_DAY'),('PEG', 'DeskOrderHandlingInst_PEGGED'),('RSV', 'DeskOrderHandlingInst_RESERVE_SIZE_ORDER'),('S.W', 'DeskOrderHandlingInst_STOP_STOCK_TRANSACTION'),('SCL', 'DeskOrderHandlingInst_SCALE'),('TMO', 'DeskOrderHandlingInst_TIME_ORDER'),('TS', 'DeskOrderHandlingInst_TRAILING_STOP'),('WRK', 'DeskOrderHandlingInst_WORK');

CREATE TABLE IF NOT EXISTS enum_maps.DeskType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DeskType VALUES 
('A', 'DeskType_AGENCY'),('AR', 'DeskType_ARBITRAGE'),('D', 'DeskType_DERIVATIVES'),('IN', 'DeskType_INTERNATIONAL'),('IS', 'DeskType_INSTITUTIONAL'),('O', 'DeskType_OTHER'),('PF', 'DeskType_PREFERRED_TRADING'),('PR', 'DeskType_PROPRIETARY'),('PT', 'DeskType_PROGRAM_TRADING'),('S', 'DeskType_SALES'),('T', 'DeskType_TRADING');

CREATE TABLE IF NOT EXISTS enum_maps.DeskTypeSource (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DeskTypeSource VALUES 
('1', 'DeskTypeSource_NASD_OATS');

CREATE TABLE IF NOT EXISTS enum_maps.DiscretionInst (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DiscretionInst VALUES 
('0', 'DiscretionInst_RELATED_TO_DISPLAYED_PRICE'),('1', 'DiscretionInst_RELATED_TO_MARKET_PRICE'),('2', 'DiscretionInst_RELATED_TO_PRIMARY_PRICE'),('3', 'DiscretionInst_RELATED_TO_LOCAL_PRIMARY_PRICE'),('4', 'DiscretionInst_RELATED_TO_MIDPOINT_PRICE'),('5', 'DiscretionInst_RELATED_TO_LAST_TRADE_PRICE'),('6', 'DiscretionInst_RELATED_TO_VWAP'),('7', 'DiscretionInst_AVERAGE_PRICE_GUARANTEE');

CREATE TABLE IF NOT EXISTS enum_maps.DiscretionLimitType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DiscretionLimitType VALUES 
('0', 'DiscretionLimitType_OR_BETTER'),('1', 'DiscretionLimitType_STRICT'),('2', 'DiscretionLimitType_OR_WORSE');

CREATE TABLE IF NOT EXISTS enum_maps.DiscretionMoveType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DiscretionMoveType VALUES 
('0', 'DiscretionMoveType_FLOATING'),('1', 'DiscretionMoveType_FIXED');

CREATE TABLE IF NOT EXISTS enum_maps.DiscretionOffsetType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DiscretionOffsetType VALUES 
('0', 'DiscretionOffsetType_PRICE'),('1', 'DiscretionOffsetType_BASIS_POINTS'),('2', 'DiscretionOffsetType_TICKS'),('3', 'DiscretionOffsetType_PRICE_TIER');

CREATE TABLE IF NOT EXISTS enum_maps.DiscretionRoundDirection (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DiscretionRoundDirection VALUES 
('1', 'DiscretionRoundDirection_MORE_AGGRESSIVE'),('2', 'DiscretionRoundDirection_MORE_PASSIVE');

CREATE TABLE IF NOT EXISTS enum_maps.DiscretionScope (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DiscretionScope VALUES 
('1', 'DiscretionScope_LOCAL'),('2', 'DiscretionScope_NATIONAL'),('3', 'DiscretionScope_GLOBAL'),('4', 'DiscretionScope_NATIONAL_EXCLUDING_LOCAL');

CREATE TABLE IF NOT EXISTS enum_maps.DisplayMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DisplayMethod VALUES 
('1', 'DisplayMethod_INITIAL'),('2', 'DisplayMethod_NEW'),('3', 'DisplayMethod_RANDOM'),('4', 'DisplayMethod_UNDISCLOSED');

CREATE TABLE IF NOT EXISTS enum_maps.DisplayWhen (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DisplayWhen VALUES 
('1', 'DisplayWhen_IMMEDIATE'),('2', 'DisplayWhen_EXHAUST');

CREATE TABLE IF NOT EXISTS enum_maps.DistribPaymentMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DistribPaymentMethod VALUES 
('1', 'DistribPaymentMethod_CREST'),('10', 'DistribPaymentMethod_BPAY'),('11', 'DistribPaymentMethod_HIGH_VALUE_CLEARING_SYSTEM_HVACS'),('12', 'DistribPaymentMethod_REINVEST_IN_FUND'),('2', 'DistribPaymentMethod_NSCC'),('3', 'DistribPaymentMethod_EUROCLEAR'),('4', 'DistribPaymentMethod_CLEARSTREAM'),('5', 'DistribPaymentMethod_CHEQUE'),('6', 'DistribPaymentMethod_TELEGRAPHIC_TRANSFER'),('7', 'DistribPaymentMethod_FED_WIRE'),('8', 'DistribPaymentMethod_DIRECT_CREDIT'),('9', 'DistribPaymentMethod_ACH_CREDIT');

CREATE TABLE IF NOT EXISTS enum_maps.DlvyInstType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DlvyInstType VALUES 
('C', 'DlvyInstType_CASH'),('S', 'DlvyInstType_SECURITIES');

CREATE TABLE IF NOT EXISTS enum_maps.DueToRelated (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.DueToRelated VALUES 
('N', 'DueToRelated_NO'),('Y', 'DueToRelated_YES');

CREATE TABLE IF NOT EXISTS enum_maps.EmailType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.EmailType VALUES 
('0', 'EmailType_NEW'),('1', 'EmailType_REPLY'),('2', 'EmailType_ADMIN_REPLY');

CREATE TABLE IF NOT EXISTS enum_maps.EncryptMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.EncryptMethod VALUES 
('0', 'EncryptMethod_NONE_OTHER'),('1', 'EncryptMethod_PKCS'),('2', 'EncryptMethod_DES'),('3', 'EncryptMethod_PKCS_DES'),('4', 'EncryptMethod_PGP_DES'),('5', 'EncryptMethod_PGP_DES_MD5'),('6', 'EncryptMethod_PEM_DES_MD5');

CREATE TABLE IF NOT EXISTS enum_maps.EventType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.EventType VALUES 
('1', 'EventType_PUT'),('10', 'EventType_SWAP_ROLL_DATE'),('11', 'EventType_SWAP_NEXT_START_DATE'),('12', 'EventType_SWAP_NEXT_ROLL_DATE'),('13', 'EventType_FIRST_DELIVERY_DATE'),('14', 'EventType_LAST_DELIVERY_DATE'),('15', 'EventType_INITIAL_INVENTORY_DUE_DATE'),('16', 'EventType_FINAL_INVENTORY_DUE_DATE'),('17', 'EventType_FIRST_INTENT_DATE'),('18', 'EventType_LAST_INTENT_DATE'),('19', 'EventType_POSITION_REMOVAL_DATE'),('2', 'EventType_CALL'),('3', 'EventType_TENDER'),('4', 'EventType_SINKING_FUND_CALL'),('5', 'EventType_ACTIVATION'),('6', 'EventType_INACTIVIATION'),('7', 'EventType_LAST_ELIGIBLE_TRADE_DATE'),('8', 'EventType_SWAP_START_DATE'),('9', 'EventType_SWAP_END_DATE'),('99', 'EventType_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.ExDestination (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ExDestination VALUES 
('0', 'ExDestination_NONE'),('4', 'ExDestination_POSIT');

CREATE TABLE IF NOT EXISTS enum_maps.ExDestinationIDSource (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ExDestinationIDSource VALUES 
('B', 'ExDestinationIDSource_BIC'),('C', 'ExDestinationIDSource_GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER'),('D', 'ExDestinationIDSource_PROPRIETARY'),('E', 'ExDestinationIDSource_ISO_COUNTRY_CODE'),('G', 'ExDestinationIDSource_MIC');

CREATE TABLE IF NOT EXISTS enum_maps.ExchangeForPhysical (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ExchangeForPhysical VALUES 
('N', 'ExchangeForPhysical_NO'),('Y', 'ExchangeForPhysical_YES');

CREATE TABLE IF NOT EXISTS enum_maps.ExecAckStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ExecAckStatus VALUES 
('0', 'ExecAckStatus_RECEIVED_NOT_YET_PROCESSED'),('1', 'ExecAckStatus_ACCEPTED'),('2', 'ExecAckStatus_DONT_KNOW');

CREATE TABLE IF NOT EXISTS enum_maps.ExecInst (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ExecInst VALUES 
('0', 'ExecInst_STAY_ON_OFFER_SIDE'),('1', 'ExecInst_NOT_HELD'),('2', 'ExecInst_WORK'),('3', 'ExecInst_GO_ALONG'),('4', 'ExecInst_OVER_THE_DAY'),('5', 'ExecInst_HELD'),('6', 'ExecInst_PARTICIPANT_DONT_INITIATE'),('7', 'ExecInst_STRICT_SCALE'),('8', 'ExecInst_TRY_TO_SCALE'),('9', 'ExecInst_STAY_ON_BID_SIDE'),('A', 'ExecInst_NO_CROSS'),('B', 'ExecInst_OK_TO_CROSS'),('C', 'ExecInst_CALL_FIRST'),('D', 'ExecInst_PERCENT_OF_VOLUME'),('E', 'ExecInst_DO_NOT_INCREASE'),('F', 'ExecInst_DO_NOT_REDUCE'),('G', 'ExecInst_ALL_OR_NONE'),('H', 'ExecInst_REINSTATE_ON_SYSTEM_FAILURE'),('I', 'ExecInst_INSTITUTIONS_ONLY'),('J', 'ExecInst_REINSTATE_ON_TRADING_HALT'),('K', 'ExecInst_CANCEL_ON_TRADING_HALT'),('L', 'ExecInst_LAST_PEG'),('M', 'ExecInst_MID_PRICE_PEG'),('N', 'ExecInst_NON_NEGOTIABLE'),('O', 'ExecInst_OPENING_PEG'),('P', 'ExecInst_MARKET_PEG'),('Q', 'ExecInst_CANCEL_ON_SYSTEM_FAILURE'),('R', 'ExecInst_PRIMARY_PEG'),('S', 'ExecInst_SUSPEND'),('T', 'ExecInst_FIXED_PEG_TO_LOCAL_BEST_BID_OR_OFFER_AT_TIME_OF_ORDER'),('U', 'ExecInst_CUSTOMER_DISPLAY_INSTRUCTION'),('V', 'ExecInst_NETTING'),('W', 'ExecInst_PEG_TO_VWAP'),('X', 'ExecInst_TRADE_ALONG'),('Y', 'ExecInst_TRY_TO_STOP'),('Z', 'ExecInst_CANCEL_IF_NOT_BEST'),('a', 'ExecInst_TRAILING_STOP_PEG'),('b', 'ExecInst_STRICT_LIMIT'),('c', 'ExecInst_IGNORE_PRICE_VALIDITY_CHECKS'),('d', 'ExecInst_PEG_TO_LIMIT_PRICE'),('e', 'ExecInst_WORK_TO_TARGET_STRATEGY'),('f', 'ExecInst_INTERMARKET_SWEEP'),('g', 'ExecInst_EXTERNAL_ROUTING_ALLOWED'),('h', 'ExecInst_EXTERNAL_ROUTING_NOT_ALLOWED'),('i', 'ExecInst_IMBALANCE_ONLY'),('j', 'ExecInst_SINGLE_EXECUTION_REQUESTED_FOR_BLOCK_TRADE'),('k', 'ExecInst_BEST_EXECUTION'),('l', 'ExecInst_SUSPEND_ON_SYSTEM_FAILURE'),('m', 'ExecInst_SUSPEND_ON_TRADING_HALT'),('n', 'ExecInst_REINSTATE_ON_CONNECTION_LOSS'),('o', 'ExecInst_CANCEL_ON_CONNECTION_LOSS'),('p', 'ExecInst_SUSPEND_ON_CONNECTION_LOSS'),('q', 'ExecInst_RELEASE_FROM_SUSPENSION'),('r', 'ExecInst_EXECUTE_AS_DELTA_NEUTRAL_USING_VOLATILITY_PROVIDED'),('s', 'ExecInst_EXECUTE_AS_DURATION_NEUTRAL'),('t', 'ExecInst_EXECUTE_AS_FX_NEUTRAL');

CREATE TABLE IF NOT EXISTS enum_maps.ExecPriceType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ExecPriceType VALUES 
('B', 'ExecPriceType_BID_PRICE'),('C', 'ExecPriceType_CREATION_PRICE'),('D', 'ExecPriceType_CREATION_PRICE_PLUS_ADJUSTMENT_PERCENT'),('E', 'ExecPriceType_CREATION_PRICE_PLUS_ADJUSTMENT_AMOUNT'),('O', 'ExecPriceType_OFFER_PRICE'),('P', 'ExecPriceType_OFFER_PRICE_MINUS_ADJUSTMENT_PERCENT'),('Q', 'ExecPriceType_OFFER_PRICE_MINUS_ADJUSTMENT_AMOUNT'),('S', 'ExecPriceType_SINGLE_PRICE');

CREATE TABLE IF NOT EXISTS enum_maps.ExecRestatementReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ExecRestatementReason VALUES 
('0', 'ExecRestatementReason_GT_CORPORATE_ACTION'),('1', 'ExecRestatementReason_GT_RENEWAL'),('10', 'ExecRestatementReason_WAREHOUSE_RECAP'),('11', 'ExecRestatementReason_PEG_REFRESH'),('2', 'ExecRestatementReason_VERBAL_CHANGE'),('3', 'ExecRestatementReason_REPRICING_OF_ORDER'),('4', 'ExecRestatementReason_BROKER_OPTION'),('5', 'ExecRestatementReason_PARTIAL_DECLINE_OF_ORDERQTY'),('6', 'ExecRestatementReason_CANCEL_ON_TRADING_HALT'),('7', 'ExecRestatementReason_CANCEL_ON_SYSTEM_FAILURE'),('8', 'ExecRestatementReason_MARKET'),('9', 'ExecRestatementReason_CANCELED_NOT_BEST'),('99', 'ExecRestatementReason_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.ExecTransType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ExecTransType VALUES 
('0', 'ExecTransType_NEW'),('1', 'ExecTransType_CANCEL'),('2', 'ExecTransType_CORRECT'),('3', 'ExecTransType_STATUS');

CREATE TABLE IF NOT EXISTS enum_maps.ExecType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ExecType VALUES 
('0', 'ExecType_NEW'),('1', 'ExecType_PARTIAL_FILL'),('2', 'ExecType_FILL'),('3', 'ExecType_DONE_FOR_DAY'),('4', 'ExecType_CANCELED'),('5', 'ExecType_REPLACED'),('6', 'ExecType_PENDING_CANCEL'),('7', 'ExecType_STOPPED'),('8', 'ExecType_REJECTED'),('9', 'ExecType_SUSPENDED'),('A', 'ExecType_PENDING_NEW'),('B', 'ExecType_CALCULATED'),('C', 'ExecType_EXPIRED'),('D', 'ExecType_RESTATED'),('E', 'ExecType_PENDING_REPLACE'),('F', 'ExecType_TRADE'),('G', 'ExecType_TRADE_CORRECT'),('H', 'ExecType_TRADE_CANCEL'),('I', 'ExecType_ORDER_STATUS'),('J', 'ExecType_TRADE_IN_A_CLEARING_HOLD'),('K', 'ExecType_TRADE_HAS_BEEN_RELEASED_TO_CLEARING'),('L', 'ExecType_TRIGGERED_OR_ACTIVATED_BY_SYSTEM');

CREATE TABLE IF NOT EXISTS enum_maps.ExerciseMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ExerciseMethod VALUES 
('A', 'ExerciseMethod_AUTOMATIC'),('M', 'ExerciseMethod_MANUAL');

CREATE TABLE IF NOT EXISTS enum_maps.ExerciseStyle (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ExerciseStyle VALUES 
('0', 'ExerciseStyle_EUROPEAN'),('1', 'ExerciseStyle_AMERICAN'),('2', 'ExerciseStyle_BERMUDA');

CREATE TABLE IF NOT EXISTS enum_maps.ExpType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ExpType VALUES 
('1', 'ExpType_AUTO_EXERCISE'),('2', 'ExpType_NON_AUTO_EXERCISE'),('3', 'ExpType_FINAL_WILL_BE_EXERCISED'),('4', 'ExpType_CONTRARY_INTENTION'),('5', 'ExpType_DIFFERENCE');

CREATE TABLE IF NOT EXISTS enum_maps.ExpirationCycle (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ExpirationCycle VALUES 
('0', 'ExpirationCycle_EXPIRE_ON_TRADING_SESSION_CLOSE'),('1', 'ExpirationCycle_EXPIRE_ON_TRADING_SESSION_OPEN'),('2', 'ExpirationCycle_TRADING_ELIGIBILITY_EXPIRATION_SPECIFIED_IN_THE_DATE_AND_TIME_FIELDS_EVENTDATE');

CREATE TABLE IF NOT EXISTS enum_maps.ExpirationQtyType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ExpirationQtyType VALUES 
('1', 'ExpirationQtyType_AUTO_EXERCISE'),('2', 'ExpirationQtyType_NON_AUTO_EXERCISE'),('3', 'ExpirationQtyType_FINAL_WILL_BE_EXERCISED'),('4', 'ExpirationQtyType_CONTRARY_INTENTION'),('5', 'ExpirationQtyType_DIFFERENCE');

CREATE TABLE IF NOT EXISTS enum_maps.FinancialStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.FinancialStatus VALUES 
('1', 'FinancialStatus_BANKRUPT'),('2', 'FinancialStatus_PENDING_DELISTING'),('3', 'FinancialStatus_RESTRICTED');

CREATE TABLE IF NOT EXISTS enum_maps.FlowScheduleType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.FlowScheduleType VALUES 
('0', 'FlowScheduleType_NERC_EASTERN_OFF_PEAK'),('1', 'FlowScheduleType_NERC_WESTERN_OFF_PEAK'),('2', 'FlowScheduleType_NERC_CALENDAR_ALL_DAYS_IN_MONTH'),('3', 'FlowScheduleType_NERC_EASTERN_PEAK'),('4', 'FlowScheduleType_NERC_WESTERN_PEAK');

CREATE TABLE IF NOT EXISTS enum_maps.ForexReq (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ForexReq VALUES 
('N', 'ForexReq_NO'),('Y', 'ForexReq_YES');

CREATE TABLE IF NOT EXISTS enum_maps.FundRenewWaiv (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.FundRenewWaiv VALUES 
('N', 'FundRenewWaiv_NO'),('Y', 'FundRenewWaiv_YES');

CREATE TABLE IF NOT EXISTS enum_maps.FuturesValuationMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.FuturesValuationMethod VALUES 
('EQTY', 'FuturesValuationMethod_PREMIUM_STYLE'),('FUT', 'FuturesValuationMethod_FUTURES_STYLE_MARK_TO_MARKET'),('FUTDA', 'FuturesValuationMethod_FUTURES_STYLE_WITH_AN_ATTACHED_CASH_ADJUSTMENT');

CREATE TABLE IF NOT EXISTS enum_maps.GTBookingInst (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.GTBookingInst VALUES 
('0', 'GTBookingInst_BOOK_OUT_ALL_TRADES_ON_DAY_OF_EXECUTION'),('1', 'GTBookingInst_ACCUMULATE_EXECTUIONS_UNTIL_FORDER_IS_FILLED_OR_EXPIRES'),('2', 'GTBookingInst_ACCUMULATE_UNTIL_VERBALLLY_NOTIFIED_OTHERWISE');

CREATE TABLE IF NOT EXISTS enum_maps.GapFillFlag (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.GapFillFlag VALUES 
('N', 'GapFillFlag_NO'),('Y', 'GapFillFlag_YES');

CREATE TABLE IF NOT EXISTS enum_maps.HaltReasonChar (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.HaltReasonChar VALUES 
('D', 'HaltReasonChar_NEWS_DISSEMINATION'),('E', 'HaltReasonChar_ORDER_INFLUX'),('I', 'HaltReasonChar_ORDER_IMBALANCE'),('M', 'HaltReasonChar_ADDITIONAL_INFORMATION'),('P', 'HaltReasonChar_NEW_PENDING'),('X', 'HaltReasonChar_EQUIPMENT_CHANGEOVER');

CREATE TABLE IF NOT EXISTS enum_maps.HaltReasonInt (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.HaltReasonInt VALUES 
('0', 'HaltReasonInt_NEWS_DISSEMINATION'),('1', 'HaltReasonInt_ORDER_INFLUX'),('2', 'HaltReasonInt_ORDER_IMBALANCE'),('3', 'HaltReasonInt_ADDITIONAL_INFORMATION'),('4', 'HaltReasonInt_NEWS_PENDING'),('5', 'HaltReasonInt_EQUIPMENT_CHANGEOVER');

CREATE TABLE IF NOT EXISTS enum_maps.HandlInst (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.HandlInst VALUES 
('1', 'HandlInst_AUTOMATED_EXECUTION_ORDER_PRIVATE_NO_BROKER_INTERVENTION'),('2', 'HandlInst_AUTOMATED_EXECUTION_ORDER_PUBLIC_BROKER_INTERVENTION_OK'),('3', 'HandlInst_MANUAL_ORDER_BEST_EXECUTION');

CREATE TABLE IF NOT EXISTS enum_maps.IDSource (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.IDSource VALUES 
('1', 'IDSource_CUSIP'),('2', 'IDSource_SEDOL'),('3', 'IDSource_QUIK'),('4', 'IDSource_ISIN_NUMBER'),('5', 'IDSource_RIC_CODE'),('6', 'IDSource_ISO_CURRENCY_CODE'),('7', 'IDSource_ISO_COUNTRY_CODE'),('8', 'IDSource_EXCHANGE_SYMBOL'),('9', 'IDSource_CONSOLIDATED_TAPE_ASSOCIATION');

CREATE TABLE IF NOT EXISTS enum_maps.IOINaturalFlag (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.IOINaturalFlag VALUES 
('N', 'IOINaturalFlag_NO'),('Y', 'IOINaturalFlag_YES');

CREATE TABLE IF NOT EXISTS enum_maps.IOIOthSvc (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.IOIOthSvc VALUES 
('A', 'IOIOthSvc_AUTEX'),('B', 'IOIOthSvc_BRIDGE');

CREATE TABLE IF NOT EXISTS enum_maps.IOIQltyInd (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.IOIQltyInd VALUES 
('H', 'IOIQltyInd_HIGH'),('L', 'IOIQltyInd_LOW'),('M', 'IOIQltyInd_MEDIUM');

CREATE TABLE IF NOT EXISTS enum_maps.IOIQty (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.IOIQty VALUES 
('0', 'IOIQty_1000000000'),('L', 'IOIQty_LARGE'),('M', 'IOIQty_MEDIUM'),('S', 'IOIQty_SMALL'),('U', 'IOIQty_UNDISCLOSED_QUANTITY');

CREATE TABLE IF NOT EXISTS enum_maps.IOIQualifier (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.IOIQualifier VALUES 
('A', 'IOIQualifier_ALL_OR_NONE'),('B', 'IOIQualifier_MARKET_ON_CLOSE'),('C', 'IOIQualifier_AT_THE_CLOSE'),('D', 'IOIQualifier_VWAP'),('I', 'IOIQualifier_IN_TOUCH_WITH'),('L', 'IOIQualifier_LIMIT'),('M', 'IOIQualifier_MORE_BEHIND'),('O', 'IOIQualifier_AT_THE_OPEN'),('P', 'IOIQualifier_TAKING_A_POSITION'),('Q', 'IOIQualifier_AT_THE_MARKET'),('R', 'IOIQualifier_READY_TO_TRADE'),('S', 'IOIQualifier_PORTFOLIO_SHOWN'),('T', 'IOIQualifier_THROUGH_THE_DAY'),('V', 'IOIQualifier_VERSUS'),('W', 'IOIQualifier_INDICATION'),('X', 'IOIQualifier_CROSSING_OPPORTUNITY'),('Y', 'IOIQualifier_AT_THE_MIDPOINT'),('Z', 'IOIQualifier_PRE_OPEN');

CREATE TABLE IF NOT EXISTS enum_maps.IOIShares (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.IOIShares VALUES 
('L', 'IOIShares_LARGE'),('M', 'IOIShares_MEDIUM'),('S', 'IOIShares_SMALL');

CREATE TABLE IF NOT EXISTS enum_maps.IOITransType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.IOITransType VALUES 
('C', 'IOITransType_CANCEL'),('N', 'IOITransType_NEW'),('R', 'IOITransType_REPLACE');

CREATE TABLE IF NOT EXISTS enum_maps.ImpliedMarketIndicator (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ImpliedMarketIndicator VALUES 
('0', 'ImpliedMarketIndicator_NOT_IMPLIED'),('1', 'ImpliedMarketIndicator_IMPLIED_IN'),('2', 'ImpliedMarketIndicator_IMPLIED_OUT'),('3', 'ImpliedMarketIndicator_BOTH_IMPLIED_IN_AND_IMPLIED_OUT');

CREATE TABLE IF NOT EXISTS enum_maps.InViewOfCommon (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.InViewOfCommon VALUES 
('N', 'InViewOfCommon_NO'),('Y', 'InViewOfCommon_YES');

CREATE TABLE IF NOT EXISTS enum_maps.IncTaxInd (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.IncTaxInd VALUES 
('1', 'IncTaxInd_NET'),('2', 'IncTaxInd_GROSS');

CREATE TABLE IF NOT EXISTS enum_maps.IndividualAllocType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.IndividualAllocType VALUES 
('1', 'IndividualAllocType_SUB_ALLOCATE'),('2', 'IndividualAllocType_THIRD_PARTY_ALLOCATION');

CREATE TABLE IF NOT EXISTS enum_maps.InstrAttribType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.InstrAttribType VALUES 
('1', 'InstrAttribType_FLAT'),('10', 'InstrAttribType_ORIGINAL_ISSUE_DISCOUNT'),('11', 'InstrAttribType_CALLABLE_PUTTABLE'),('12', 'InstrAttribType_ESCROWED_TO_MATURITY'),('13', 'InstrAttribType_ESCROWED_TO_REDEMPTION_DATE'),('14', 'InstrAttribType_PRE_REFUNDED'),('15', 'InstrAttribType_IN_DEFAULT'),('16', 'InstrAttribType_UNRATED'),('17', 'InstrAttribType_TAXABLE'),('18', 'InstrAttribType_INDEXED'),('19', 'InstrAttribType_SUBJECT_TO_ALTERNATIVE_MINIMUM_TAX'),('2', 'InstrAttribType_ZERO_COUPON'),('20', 'InstrAttribType_ORIGINAL_ISSUE_DISCOUNT_PRICE_SUPPLY_PRICE_IN_THE_INSTRATTRIBVALUE'),('21', 'InstrAttribType_CALLABLE_BELOW_MATURITY_VALUE'),('22', 'InstrAttribType_CALLABLE_WITHOUT_NOTICE_BY_MAIL_TO_HOLDER_UNLESS_REGISTERED'),('23', 'InstrAttribType_PRICE_TICK_RULES_FOR_SECURITY'),('24', 'InstrAttribType_TRADE_TYPE_ELIGIBILITY_DETAILS_FOR_SECURITY'),('25', 'InstrAttribType_INSTRUMENT_DENOMINATOR'),('26', 'InstrAttribType_INSTRUMENT_NUMERATOR'),('27', 'InstrAttribType_INSTRUMENT_PRICE_PRECISION'),('28', 'InstrAttribType_INSTRUMENT_STRIKE_PRICE'),('29', 'InstrAttribType_TRADEABLE_INDICATOR'),('3', 'InstrAttribType_INTEREST_BEARING'),('4', 'InstrAttribType_NO_PERIODIC_PAYMENTS'),('5', 'InstrAttribType_VARIABLE_RATE'),('6', 'InstrAttribType_LESS_FEE_FOR_PUT'),('7', 'InstrAttribType_STEPPED_COUPON'),('8', 'InstrAttribType_COUPON_PERIOD'),('9', 'InstrAttribType_WHEN_AND_IF_ISSUED'),('99', 'InstrAttribType_TEXT_SUPPLY_THE_TEXT_OF_THE_ATTRIBUTE_OR_DISCLAIMER_IN_THE_INSTRATTRIBVALUE');

CREATE TABLE IF NOT EXISTS enum_maps.InstrRegistry (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.InstrRegistry VALUES 
('BIC', 'InstrRegistry_CUSTODIAN'),('ISO', 'InstrRegistry_COUNTRY'),('ZZ', 'InstrRegistry_PHYSICAL');

CREATE TABLE IF NOT EXISTS enum_maps.LastCapacity (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.LastCapacity VALUES 
('1', 'LastCapacity_AGENT'),('2', 'LastCapacity_CROSS_AS_AGENT'),('3', 'LastCapacity_CROSS_AS_PRINCIPAL'),('4', 'LastCapacity_PRINCIPAL');

CREATE TABLE IF NOT EXISTS enum_maps.LastFragment (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.LastFragment VALUES 
('N', 'LastFragment_NO'),('Y', 'LastFragment_YES');

CREATE TABLE IF NOT EXISTS enum_maps.LastLiquidityInd (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.LastLiquidityInd VALUES 
('1', 'LastLiquidityInd_ADDED_LIQUIDITY'),('2', 'LastLiquidityInd_REMOVED_LIQUIDITY'),('3', 'LastLiquidityInd_LIQUIDITY_ROUTED_OUT'),('4', 'LastLiquidityInd_AUCTION');

CREATE TABLE IF NOT EXISTS enum_maps.LastRptRequested (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.LastRptRequested VALUES 
('N', 'LastRptRequested_NO'),('Y', 'LastRptRequested_YES');

CREATE TABLE IF NOT EXISTS enum_maps.LegSwapType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.LegSwapType VALUES 
('1', 'LegSwapType_PAR_FOR_PAR'),('2', 'LegSwapType_MODIFIED_DURATION'),('4', 'LegSwapType_RISK'),('5', 'LegSwapType_PROCEEDS');

CREATE TABLE IF NOT EXISTS enum_maps.LegalConfirm (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.LegalConfirm VALUES 
('N', 'LegalConfirm_NO'),('Y', 'LegalConfirm_YES');

CREATE TABLE IF NOT EXISTS enum_maps.LiquidityIndType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.LiquidityIndType VALUES 
('1', 'LiquidityIndType_5_DAY_MOVING_AVERAGE'),('2', 'LiquidityIndType_20_DAY_MOVING_AVERAGE'),('3', 'LiquidityIndType_NORMAL_MARKET_SIZE'),('4', 'LiquidityIndType_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.ListExecInstType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ListExecInstType VALUES 
('1', 'ListExecInstType_IMMEDIATE'),('2', 'ListExecInstType_WAIT_FOR_EXECUT_INSTRUCTION'),('3', 'ListExecInstType_EXCHANGE_SWITCH_CIV_ORDER_3'),('4', 'ListExecInstType_EXCHANGE_SWITCH_CIV_ORDER_4'),('5', 'ListExecInstType_EXCHANGE_SWITCH_CIV_ORDER_5');

CREATE TABLE IF NOT EXISTS enum_maps.ListMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ListMethod VALUES 
('0', 'ListMethod_PRE_LISTED_ONLY'),('1', 'ListMethod_USER_REQUESTED');

CREATE TABLE IF NOT EXISTS enum_maps.ListOrderStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ListOrderStatus VALUES 
('1', 'ListOrderStatus_IN_BIDDING_PROCESS'),('2', 'ListOrderStatus_RECEIVED_FOR_EXECUTION'),('3', 'ListOrderStatus_EXECUTING'),('4', 'ListOrderStatus_CANCELLING'),('5', 'ListOrderStatus_ALERT'),('6', 'ListOrderStatus_ALL_DONE'),('7', 'ListOrderStatus_REJECT');

CREATE TABLE IF NOT EXISTS enum_maps.ListRejectReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ListRejectReason VALUES 
('0', 'ListRejectReason_BROKER'),('11', 'ListRejectReason_UNSUPPORTED_ORDER_CHARACTERISTIC'),('2', 'ListRejectReason_EXCHANGE_CLOSED'),('4', 'ListRejectReason_TOO_LATE_TO_ENTER'),('5', 'ListRejectReason_UNKNOWN_ORDER'),('6', 'ListRejectReason_DUPLICATE_ORDER'),('99', 'ListRejectReason_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.ListStatusType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ListStatusType VALUES 
('1', 'ListStatusType_ACK'),('2', 'ListStatusType_RESPONSE'),('3', 'ListStatusType_TIMED'),('4', 'ListStatusType_EXEC_STARTED'),('5', 'ListStatusType_ALL_DONE'),('6', 'ListStatusType_ALERT');

CREATE TABLE IF NOT EXISTS enum_maps.LocateReqd (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.LocateReqd VALUES 
('N', 'LocateReqd_NO'),('Y', 'LocateReqd_YES');

CREATE TABLE IF NOT EXISTS enum_maps.LotType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.LotType VALUES 
('1', 'LotType_ODD_LOT'),('2', 'LotType_ROUND_LOT'),('3', 'LotType_BLOCK_LOT'),('4', 'LotType_ROUND_LOT_BASED_UPON_UNITOFMEASURE');

CREATE TABLE IF NOT EXISTS enum_maps.MDBookType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MDBookType VALUES 
('1', 'MDBookType_TOP_OF_BOOK'),('2', 'MDBookType_PRICE_DEPTH'),('3', 'MDBookType_ORDER_DEPTH');

CREATE TABLE IF NOT EXISTS enum_maps.MDEntryType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MDEntryType VALUES 
('0', 'MDEntryType_BID'),('1', 'MDEntryType_OFFER'),('2', 'MDEntryType_TRADE'),('3', 'MDEntryType_INDEX_VALUE'),('4', 'MDEntryType_OPENING_PRICE'),('5', 'MDEntryType_CLOSING_PRICE'),('6', 'MDEntryType_SETTLEMENT_PRICE'),('7', 'MDEntryType_TRADING_SESSION_HIGH_PRICE'),('8', 'MDEntryType_TRADING_SESSION_LOW_PRICE'),('9', 'MDEntryType_TRADING_SESSION_VWAP_PRICE'),('A', 'MDEntryType_IMBALANCE'),('B', 'MDEntryType_TRADE_VOLUME'),('C', 'MDEntryType_OPEN_INTEREST'),('D', 'MDEntryType_COMPOSITE_UNDERLYING_PRICE'),('E', 'MDEntryType_SIMULATED_SELL_PRICE'),('F', 'MDEntryType_SIMULATED_BUY_PRICE'),('G', 'MDEntryType_MARGIN_RATE'),('H', 'MDEntryType_MID_PRICE'),('J', 'MDEntryType_EMPTY_BOOK'),('K', 'MDEntryType_SETTLE_HIGH_PRICE'),('L', 'MDEntryType_SETTLE_LOW_PRICE'),('M', 'MDEntryType_PRIOR_SETTLE_PRICE'),('N', 'MDEntryType_SESSION_HIGH_BID'),('O', 'MDEntryType_SESSION_LOW_OFFER'),('P', 'MDEntryType_EARLY_PRICES'),('Q', 'MDEntryType_AUCTION_CLEARING_PRICE'),('R', 'MDEntryType_DAILY_VALUE_ADJUSTMENT_FOR_LONG_POSITIONS'),('S', 'MDEntryType_SWAP_VALUE_FACTOR'),('T', 'MDEntryType_CUMULATIVE_VALUE_ADJUSTMENT_FOR_LONG_POSITIONS'),('U', 'MDEntryType_DAILY_VALUE_ADJUSTMENT_FOR_SHORT_POSITIONS'),('V', 'MDEntryType_CUMULATIVE_VALUE_ADJUSTMENT_FOR_SHORT_POSITIONS'),('W', 'MDEntryType_FIXING_PRICE'),('X', 'MDEntryType_CASH_RATE'),('Y', 'MDEntryType_RECOVERY_RATE'),('Z', 'MDEntryType_RECOVERY_RATE_FOR_LONG'),('a', 'MDEntryType_RECOVERY_RATE_FOR_SHORT');

CREATE TABLE IF NOT EXISTS enum_maps.MDImplicitDelete (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MDImplicitDelete VALUES 
('N', 'MDImplicitDelete_NO'),('Y', 'MDImplicitDelete_YES');

CREATE TABLE IF NOT EXISTS enum_maps.MDOriginType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MDOriginType VALUES 
('0', 'MDOriginType_BOOK'),('1', 'MDOriginType_OFF_BOOK'),('2', 'MDOriginType_CROSS');

CREATE TABLE IF NOT EXISTS enum_maps.MDQuoteType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MDQuoteType VALUES 
('0', 'MDQuoteType_INDICATIVE'),('1', 'MDQuoteType_TRADEABLE'),('2', 'MDQuoteType_RESTRICTED_TRADEABLE'),('3', 'MDQuoteType_COUNTER'),('4', 'MDQuoteType_INDICATIVE_AND_TRADEABLE');

CREATE TABLE IF NOT EXISTS enum_maps.MDReqRejReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MDReqRejReason VALUES 
('0', 'MDReqRejReason_UNKNOWN_SYMBOL'),('1', 'MDReqRejReason_DUPLICATE_MDREQID'),('2', 'MDReqRejReason_INSUFFICIENT_BANDWIDTH'),('3', 'MDReqRejReason_INSUFFICIENT_PERMISSIONS'),('4', 'MDReqRejReason_UNSUPPORTED_SUBSCRIPTIONREQUESTTYPE'),('5', 'MDReqRejReason_UNSUPPORTED_MARKETDEPTH'),('6', 'MDReqRejReason_UNSUPPORTED_MDUPDATETYPE'),('7', 'MDReqRejReason_UNSUPPORTED_AGGREGATEDBOOK'),('8', 'MDReqRejReason_UNSUPPORTED_MDENTRYTYPE'),('9', 'MDReqRejReason_UNSUPPORTED_TRADINGSESSIONID'),('A', 'MDReqRejReason_UNSUPPORTED_SCOPE'),('B', 'MDReqRejReason_UNSUPPORTED_OPENCLOSESETTLEFLAG'),('C', 'MDReqRejReason_UNSUPPORTED_MDIMPLICITDELETE'),('D', 'MDReqRejReason_INSUFFICIENT_CREDIT');

CREATE TABLE IF NOT EXISTS enum_maps.MDSecSizeType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MDSecSizeType VALUES 
('1', 'MDSecSizeType_CUSTOMER');

CREATE TABLE IF NOT EXISTS enum_maps.MDUpdateAction (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MDUpdateAction VALUES 
('0', 'MDUpdateAction_NEW'),('1', 'MDUpdateAction_CHANGE'),('2', 'MDUpdateAction_DELETE'),('3', 'MDUpdateAction_DELETE_THRU'),('4', 'MDUpdateAction_DELETE_FROM'),('5', 'MDUpdateAction_OVERLAY');

CREATE TABLE IF NOT EXISTS enum_maps.MDUpdateType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MDUpdateType VALUES 
('0', 'MDUpdateType_FULL_REFRESH'),('1', 'MDUpdateType_INCREMENTAL_REFRESH');

CREATE TABLE IF NOT EXISTS enum_maps.MarketUpdateAction (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MarketUpdateAction VALUES 
('A', 'MarketUpdateAction_ADD'),('D', 'MarketUpdateAction_DELETE'),('M', 'MarketUpdateAction_MODIFY');

CREATE TABLE IF NOT EXISTS enum_maps.MassActionRejectReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MassActionRejectReason VALUES 
('0', 'MassActionRejectReason_MASS_ACTION_NOT_SUPPORTED'),('1', 'MassActionRejectReason_INVALID_OR_UNKNOWN_SECURITY'),('10', 'MassActionRejectReason_INVALID_OR_UNKNOWN_SECURITY_ISSUER'),('11', 'MassActionRejectReason_INVALID_OR_UNKNOWN_ISSUER_OF_UNDERLYING_SECURITY'),('2', 'MassActionRejectReason_INVALID_OR_UNKNOWN_UNDERLYING_SECURITY'),('3', 'MassActionRejectReason_INVALID_OR_UNKNOWN_PRODUCT'),('4', 'MassActionRejectReason_INVALID_OR_UNKNOWN_CFICODE'),('5', 'MassActionRejectReason_INVALID_OR_UNKNOWN_SECURITYTYPE'),('6', 'MassActionRejectReason_INVALID_OR_UNKNOWN_TRADING_SESSION'),('7', 'MassActionRejectReason_INVALID_OR_UNKNOWN_MARKET'),('8', 'MassActionRejectReason_INVALID_OR_UNKNOWN_MARKET_SEGMENT'),('9', 'MassActionRejectReason_INVALID_OR_UNKNOWN_SECURITY_GROUP'),('99', 'MassActionRejectReason_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.MassActionResponse (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MassActionResponse VALUES 
('0', 'MassActionResponse_REJECTED'),('1', 'MassActionResponse_ACCEPTED');

CREATE TABLE IF NOT EXISTS enum_maps.MassActionScope (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MassActionScope VALUES 
('1', 'MassActionScope_ALL_ORDERS_FOR_A_SECURITY'),('10', 'MassActionScope_ALL_ORDERS_FOR_A_SECURITY_GROUP'),('11', 'MassActionScope_CANCEL_FOR_SECURITY_ISSUER'),('12', 'MassActionScope_CANCEL_FOR_ISSUER_OF_UNDERLYING_SECURITY'),('2', 'MassActionScope_ALL_ORDERS_FOR_AN_UNDERLYING_SECURITY'),('3', 'MassActionScope_ALL_ORDERS_FOR_A_PRODUCT'),('4', 'MassActionScope_ALL_ORDERS_FOR_A_CFICODE'),('5', 'MassActionScope_ALL_ORDERS_FOR_A_SECURITYTYPE'),('6', 'MassActionScope_ALL_ORDERS_FOR_A_TRADING_SESSION'),('7', 'MassActionScope_ALL_ORDERS'),('8', 'MassActionScope_ALL_ORDERS_FOR_A_MARKET'),('9', 'MassActionScope_ALL_ORDERS_FOR_A_MARKET_SEGMENT');

CREATE TABLE IF NOT EXISTS enum_maps.MassActionType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MassActionType VALUES 
('1', 'MassActionType_SUSPEND_ORDERS'),('2', 'MassActionType_RELEASE_ORDERS_FROM_SUSPENSION'),('3', 'MassActionType_CANCEL_ORDERS');

CREATE TABLE IF NOT EXISTS enum_maps.MassCancelRejectReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MassCancelRejectReason VALUES 
('0', 'MassCancelRejectReason_MASS_CANCEL_NOT_SUPPORTED'),('1', 'MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITY'),('10', 'MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITY_ISSUER'),('11', 'MassCancelRejectReason_INVALID_OR_UNKNOWN_ISSUER_OF_UNDERLYING_SECURITY'),('2', 'MassCancelRejectReason_INVALID_OR_UNKOWN_UNDERLYING_SECURITY'),('3', 'MassCancelRejectReason_INVALID_OR_UNKNOWN_PRODUCT'),('4', 'MassCancelRejectReason_INVALID_OR_UNKNOWN_CFICODE'),('5', 'MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITYTYPE'),('6', 'MassCancelRejectReason_INVALID_OR_UNKNOWN_TRADING_SESSION'),('7', 'MassCancelRejectReason_INVALID_OR_UNKNOWN_MARKET'),('8', 'MassCancelRejectReason_INVALID_OR_UNKOWN_MARKET_SEGMENT'),('9', 'MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITY_GROUP'),('99', 'MassCancelRejectReason_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.MassCancelRequestType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MassCancelRequestType VALUES 
('1', 'MassCancelRequestType_CANCEL_ORDERS_FOR_A_SECURITY'),('2', 'MassCancelRequestType_CANCEL_ORDERS_FOR_AN_UNDERLYING_SECURITY'),('3', 'MassCancelRequestType_CANCEL_ORDERS_FOR_A_PRODUCT'),('4', 'MassCancelRequestType_CANCEL_ORDERS_FOR_A_CFICODE'),('5', 'MassCancelRequestType_CANCEL_ORDERS_FOR_A_SECURITYTYPE'),('6', 'MassCancelRequestType_CANCEL_ORDERS_FOR_A_TRADING_SESSION'),('7', 'MassCancelRequestType_CANCEL_ALL_ORDERS'),('8', 'MassCancelRequestType_CANCEL_ORDERS_FOR_A_MARKET'),('9', 'MassCancelRequestType_CANCEL_ORDERS_FOR_A_MARKET_SEGMENT'),('A', 'MassCancelRequestType_CANCEL_ORDERS_FOR_A_SECURITY_GROUP'),('B', 'MassCancelRequestType_CANCEL_FOR_SECURITY_ISSUER'),('C', 'MassCancelRequestType_CANCEL_FOR_ISSUER_OF_UNDERLYING_SECURITY');

CREATE TABLE IF NOT EXISTS enum_maps.MassCancelResponse (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MassCancelResponse VALUES 
('0', 'MassCancelResponse_CANCEL_REQUEST_REJECTED'),('1', 'MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITY'),('2', 'MassCancelResponse_CANCEL_ORDERS_FOR_AN_UNDERLYING_SECURITY'),('3', 'MassCancelResponse_CANCEL_ORDERS_FOR_A_PRODUCT'),('4', 'MassCancelResponse_CANCEL_ORDERS_FOR_A_CFICODE'),('5', 'MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITYTYPE'),('6', 'MassCancelResponse_CANCEL_ORDERS_FOR_A_TRADING_SESSION'),('7', 'MassCancelResponse_CANCEL_ALL_ORDERS'),('8', 'MassCancelResponse_CANCEL_ORDERS_FOR_A_MARKET'),('9', 'MassCancelResponse_CANCEL_ORDERS_FOR_A_MARKET_SEGMENT'),('A', 'MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITY_GROUP'),('B', 'MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITIES_ISSUER'),('C', 'MassCancelResponse_CANCEL_ORDERS_FOR_ISSUER_OF_UNDERLYING_SECURITY');

CREATE TABLE IF NOT EXISTS enum_maps.MassStatusReqType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MassStatusReqType VALUES 
('1', 'MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_SECURITY'),('10', 'MassStatusReqType_STATUS_FOR_ISSUER_OF_UNDERLYING_SECURITY'),('2', 'MassStatusReqType_STATUS_FOR_ORDERS_FOR_AN_UNDERLYING_SECURITY'),('3', 'MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_PRODUCT'),('4', 'MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_CFICODE'),('5', 'MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_SECURITYTYPE'),('6', 'MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_TRADING_SESSION'),('7', 'MassStatusReqType_STATUS_FOR_ALL_ORDERS'),('8', 'MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_PARTYID'),('9', 'MassStatusReqType_STATUS_FOR_SECURITY_ISSUER');

CREATE TABLE IF NOT EXISTS enum_maps.MatchStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MatchStatus VALUES 
('0', 'MatchStatus_COMPARED_MATCHED_OR_AFFIRMED'),('1', 'MatchStatus_UNCOMPARED_UNMATCHED_OR_UNAFFIRMED'),('2', 'MatchStatus_ADVISORY_OR_ALERT');

CREATE TABLE IF NOT EXISTS enum_maps.MatchType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MatchType VALUES 
('1', 'MatchType_ONE_PARTY_TRADE_REPORT'),('2', 'MatchType_TWO_PARTY_TRADE_REPORT'),('3', 'MatchType_CONFIRMED_TRADE_REPORT'),('4', 'MatchType_AUTO_MATCH'),('5', 'MatchType_CROSS_AUCTION'),('6', 'MatchType_COUNTER_ORDER_SELECTION'),('60', 'MatchType_ONE_PARTY_PRIVATELY_NEGOTIATED_TRADE_REPORT'),('61', 'MatchType_TWO_PARTY_PRIVATELY_NEGOTIATED_TRADE_REPORT'),('62', 'MatchType_CONTINUOUS_AUTO_MATCH'),('63', 'MatchType_CROSS_AUCTION_63'),('64', 'MatchType_COUNTER_ORDER_SELECTION_64'),('65', 'MatchType_CALL_AUCTION_65'),('7', 'MatchType_CALL_AUCTION'),('8', 'MatchType_ISSUING_BUY_BACK_AUCTION'),('A1', 'MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_FOUR_BADGES_AND_EXECUTION_TIME'),('A2', 'MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_FOUR_BADGES'),('A3', 'MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_TWO_BADGES_AND_EXECUTION_TIME'),('A4', 'MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_TWO_BADGES'),('A5', 'MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADETYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_EXECUTION_TIME'),('ACTM1', 'MatchType_NASDAQACTM1MATCH'),('ACTM2', 'MatchType_NASDAQACTM2MATCH'),('ACTM3', 'MatchType_NASDAQACTACCEPTEDTRADE'),('ACTM4', 'MatchType_NASDAQACTDEFAULTTRADE'),('ACTM5', 'MatchType_NASDAQACTDEFAULTAFTERM2'),('ACTM6', 'MatchType_NASDAQACTM6MATCH'),('ACTMT', 'MatchType_NASDAQNONACT'),('AQ', 'MatchType_COMPARED_RECORDS_RESULTING_FROM_STAMPED_ADVISORIES_OR_SPECIALIST_ACCEPTS_PAIR_OFFS'),('M1', 'MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_MINUS_BADGES_AND_TIMES_ACT_M1_MATCH'),('M2', 'MatchType_SUMMARIZED_MATCH_MINUS_BADGES_AND_TIMES_ACT_M2_MATCH'),('M3', 'MatchType_ACT_ACCEPTED_TRADE'),('M4', 'MatchType_ACT_DEFAULT_TRADE'),('M5', 'MatchType_ACT_DEFAULT_AFTER_M2'),('M6', 'MatchType_ACT_M6_MATCH'),('MT', 'MatchType_OCS_LOCKED_IN_NON_ACT'),('S1', 'MatchType_SUMMARIZED_MATCH_USING_A1_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIED'),('S2', 'MatchType_SUMMARIZED_MATCH_USING_A2_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIZED'),('S3', 'MatchType_SUMMARIZED_MATCH_USING_A3_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIZED'),('S4', 'MatchType_SUMMARIZED_MATCH_USING_A4_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIZED'),('S5', 'MatchType_SUMMARIZED_MATCH_USING_A5_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIZED');

CREATE TABLE IF NOT EXISTS enum_maps.MaturityMonthYearFormat (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MaturityMonthYearFormat VALUES 
('0', 'MaturityMonthYearFormat_YEARMONTH_ONLY'),('1', 'MaturityMonthYearFormat_YEARMONTHDAY'),('2', 'MaturityMonthYearFormat_YEARMONTHWEEK');

CREATE TABLE IF NOT EXISTS enum_maps.MaturityMonthYearIncrementUnits (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MaturityMonthYearIncrementUnits VALUES 
('0', 'MaturityMonthYearIncrementUnits_MONTHS'),('1', 'MaturityMonthYearIncrementUnits_DAYS'),('2', 'MaturityMonthYearIncrementUnits_WEEKS'),('3', 'MaturityMonthYearIncrementUnits_YEARS');

CREATE TABLE IF NOT EXISTS enum_maps.MessageEncoding (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MessageEncoding VALUES 
('EUC-JP', 'MessageEncoding_EUC_JP'),('ISO-2022-JP', 'MessageEncoding_ISO_2022_JP'),('SHIFT_JIS', 'MessageEncoding_SHIFT_JIS'),('UTF-8', 'MessageEncoding_UTF_8');

CREATE TABLE IF NOT EXISTS enum_maps.MiscFeeBasis (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MiscFeeBasis VALUES 
('0', 'MiscFeeBasis_ABSOLUTE'),('1', 'MiscFeeBasis_PER_UNIT'),('2', 'MiscFeeBasis_PERCENTAGE');

CREATE TABLE IF NOT EXISTS enum_maps.MiscFeeType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MiscFeeType VALUES 
('1', 'MiscFeeType_REGULATORY'),('10', 'MiscFeeType_PER_TRANSACTION'),('11', 'MiscFeeType_CONVERSION'),('12', 'MiscFeeType_AGENT'),('13', 'MiscFeeType_TRANSFER_FEE'),('14', 'MiscFeeType_SECURITY_LENDING'),('2', 'MiscFeeType_TAX'),('3', 'MiscFeeType_LOCAL_COMMISSION'),('4', 'MiscFeeType_EXCHANGE_FEES'),('5', 'MiscFeeType_STAMP'),('6', 'MiscFeeType_LEVY'),('7', 'MiscFeeType_OTHER'),('8', 'MiscFeeType_MARKUP'),('9', 'MiscFeeType_CONSUMPTION_TAX');

CREATE TABLE IF NOT EXISTS enum_maps.ModelType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ModelType VALUES 
('0', 'ModelType_UTILITY_PROVIDED_STANDARD_MODEL'),('1', 'ModelType_PROPRIETARY');

CREATE TABLE IF NOT EXISTS enum_maps.MoneyLaunderingStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MoneyLaunderingStatus VALUES 
('1', 'MoneyLaunderingStatus_EXEMPT_1'),('2', 'MoneyLaunderingStatus_EXEMPT_2'),('3', 'MoneyLaunderingStatus_EXEMPT_3'),('N', 'MoneyLaunderingStatus_NOT_CHECKED'),('Y', 'MoneyLaunderingStatus_PASSED');

CREATE TABLE IF NOT EXISTS enum_maps.MsgDirection (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MsgDirection VALUES 
('R', 'MsgDirection_RECEIVE'),('S', 'MsgDirection_SEND');

CREATE TABLE IF NOT EXISTS enum_maps.MsgType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MsgType VALUES 
('0', 'MsgType_HEARTBEAT'),('1', 'MsgType_TEST_REQUEST'),('2', 'MsgType_RESEND_REQUEST'),('3', 'MsgType_REJECT'),('4', 'MsgType_SEQUENCE_RESET'),('5', 'MsgType_LOGOUT'),('6', 'MsgType_INDICATION_OF_INTEREST'),('7', 'MsgType_ADVERTISEMENT'),('8', 'MsgType_EXECUTION_REPORT'),('9', 'MsgType_ORDER_CANCEL_REJECT'),('A', 'MsgType_LOGON'),('AA', 'MsgType_DERIVATIVE_SECURITY_LIST'),('AB', 'MsgType_NEW_ORDER_MULTILEG'),('AC', 'MsgType_MULTILEG_ORDER_CANCEL_REPLACE'),('AD', 'MsgType_TRADE_CAPTURE_REPORT_REQUEST'),('AE', 'MsgType_TRADE_CAPTURE_REPORT'),('AF', 'MsgType_ORDER_MASS_STATUS_REQUEST'),('AG', 'MsgType_QUOTE_REQUEST_REJECT'),('AH', 'MsgType_RFQ_REQUEST'),('AI', 'MsgType_QUOTE_STATUS_REPORT'),('AJ', 'MsgType_QUOTE_RESPONSE'),('AK', 'MsgType_CONFIRMATION'),('AL', 'MsgType_POSITION_MAINTENANCE_REQUEST'),('AM', 'MsgType_POSITION_MAINTENANCE_REPORT'),('AN', 'MsgType_REQUEST_FOR_POSITIONS'),('AO', 'MsgType_REQUEST_FOR_POSITIONS_ACK'),('AP', 'MsgType_POSITION_REPORT'),('AQ', 'MsgType_TRADE_CAPTURE_REPORT_REQUEST_ACK'),('AR', 'MsgType_TRADE_CAPTURE_REPORT_ACK'),('AS', 'MsgType_ALLOCATION_REPORT'),('AT', 'MsgType_ALLOCATION_REPORT_ACK'),('AU', 'MsgType_CONFIRMATION_ACK'),('AV', 'MsgType_SETTLEMENT_INSTRUCTION_REQUEST'),('AW', 'MsgType_ASSIGNMENT_REPORT'),('AX', 'MsgType_COLLATERAL_REQUEST'),('AY', 'MsgType_COLLATERAL_ASSIGNMENT'),('AZ', 'MsgType_COLLATERAL_RESPONSE'),('B', 'MsgType_NEWS'),('BA', 'MsgType_COLLATERAL_REPORT'),('BB', 'MsgType_COLLATERAL_INQUIRY'),('BC', 'MsgType_NETWORK_STATUS_REQUEST'),('BD', 'MsgType_NETWORK_STATUS_RESPONSE'),('BE', 'MsgType_USER_REQUEST'),('BF', 'MsgType_USER_RESPONSE'),('BG', 'MsgType_COLLATERAL_INQUIRY_ACK'),('BH', 'MsgType_CONFIRMATION_REQUEST'),('BI', 'MsgType_TRADING_SESSION_LIST_REQUEST'),('BJ', 'MsgType_TRADING_SESSION_LIST'),('BK', 'MsgType_SECURITY_LIST_UPDATE_REPORT'),('BL', 'MsgType_ADJUSTED_POSITION_REPORT'),('BM', 'MsgType_ALLOCATION_INSTRUCTION_ALERT'),('BN', 'MsgType_EXECUTION_ACKNOWLEDGEMENT'),('BO', 'MsgType_CONTRARY_INTENTION_REPORT'),('BP', 'MsgType_SECURITY_DEFINITION_UPDATE_REPORT'),('BQ', 'MsgType_SETTLEMENTOBLIGATIONREPORT'),('BR', 'MsgType_DERIVATIVESECURITYLISTUPDATEREPORT'),('BS', 'MsgType_TRADINGSESSIONLISTUPDATEREPORT'),('BT', 'MsgType_MARKETDEFINITIONREQUEST'),('BU', 'MsgType_MARKETDEFINITION'),('BV', 'MsgType_MARKETDEFINITIONUPDATEREPORT'),('BW', 'MsgType_APPLICATIONMESSAGEREQUEST'),('BX', 'MsgType_APPLICATIONMESSAGEREQUESTACK'),('BY', 'MsgType_APPLICATIONMESSAGEREPORT'),('BZ', 'MsgType_ORDERMASSACTIONREPORT'),('C', 'MsgType_EMAIL'),('CA', 'MsgType_ORDERMASSACTIONREQUEST'),('CB', 'MsgType_USERNOTIFICATION'),('CC', 'MsgType_STREAMASSIGNMENTREQUEST'),('CD', 'MsgType_STREAMASSIGNMENTREPORT'),('CE', 'MsgType_STREAMASSIGNMENTREPORTACK'),('CF', 'MsgType_PARTYDETAILSLISTREQUEST'),('CG', 'MsgType_PARTYDETAILSLISTREPORT'),('D', 'MsgType_ORDER_SINGLE'),('E', 'MsgType_ORDER_LIST'),('F', 'MsgType_ORDER_CANCEL_REQUEST'),('G', 'MsgType_ORDER_CANCEL_REPLACE_REQUEST'),('H', 'MsgType_ORDER_STATUS_REQUEST'),('J', 'MsgType_ALLOCATION_INSTRUCTION'),('K', 'MsgType_LIST_CANCEL_REQUEST'),('L', 'MsgType_LIST_EXECUTE'),('M', 'MsgType_LIST_STATUS_REQUEST'),('N', 'MsgType_LIST_STATUS'),('P', 'MsgType_ALLOCATION_INSTRUCTION_ACK'),('Q', 'MsgType_DONT_KNOW_TRADE'),('R', 'MsgType_QUOTE_REQUEST'),('S', 'MsgType_QUOTE'),('T', 'MsgType_SETTLEMENT_INSTRUCTIONS'),('V', 'MsgType_MARKET_DATA_REQUEST'),('W', 'MsgType_MARKET_DATA_SNAPSHOT_FULL_REFRESH'),('X', 'MsgType_MARKET_DATA_INCREMENTAL_REFRESH'),('Y', 'MsgType_MARKET_DATA_REQUEST_REJECT'),('Z', 'MsgType_QUOTE_CANCEL'),('a', 'MsgType_QUOTE_STATUS_REQUEST'),('b', 'MsgType_MASS_QUOTE_ACKNOWLEDGEMENT'),('c', 'MsgType_SECURITY_DEFINITION_REQUEST'),('d', 'MsgType_SECURITY_DEFINITION'),('e', 'MsgType_SECURITY_STATUS_REQUEST'),('f', 'MsgType_SECURITY_STATUS'),('g', 'MsgType_TRADING_SESSION_STATUS_REQUEST'),('h', 'MsgType_TRADING_SESSION_STATUS'),('i', 'MsgType_MASS_QUOTE'),('j', 'MsgType_BUSINESS_MESSAGE_REJECT'),('k', 'MsgType_BID_REQUEST'),('l', 'MsgType_BID_RESPONSE'),('m', 'MsgType_LIST_STRIKE_PRICE'),('n', 'MsgType_XML_MESSAGE'),('o', 'MsgType_REGISTRATION_INSTRUCTIONS'),('p', 'MsgType_REGISTRATION_INSTRUCTIONS_RESPONSE'),('q', 'MsgType_ORDER_MASS_CANCEL_REQUEST'),('r', 'MsgType_ORDER_MASS_CANCEL_REPORT'),('s', 'MsgType_NEW_ORDER_CROSS'),('t', 'MsgType_CROSS_ORDER_CANCEL_REPLACE_REQUEST'),('u', 'MsgType_CROSS_ORDER_CANCEL_REQUEST'),('v', 'MsgType_SECURITY_TYPE_REQUEST'),('w', 'MsgType_SECURITY_TYPES'),('x', 'MsgType_SECURITY_LIST_REQUEST'),('y', 'MsgType_SECURITY_LIST'),('z', 'MsgType_DERIVATIVE_SECURITY_LIST_REQUEST');

CREATE TABLE IF NOT EXISTS enum_maps.MultiLegReportingType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MultiLegReportingType VALUES 
('1', 'MultiLegReportingType_SINGLE_SECURITY'),('2', 'MultiLegReportingType_INDIVIDUAL_LEG_OF_A_MULTI_LEG_SECURITY'),('3', 'MultiLegReportingType_MULTI_LEG_SECURITY');

CREATE TABLE IF NOT EXISTS enum_maps.MultiLegRptTypeReq (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MultiLegRptTypeReq VALUES 
('0', 'MultiLegRptTypeReq_REPORT_BY_MULITLEG_SECURITY_ONLY'),('1', 'MultiLegRptTypeReq_REPORT_BY_MULTILEG_SECURITY_AND_BY_INSTRUMENT_LEGS_BELONGING_TO_THE_MULTILEG_SECURITY'),('2', 'MultiLegRptTypeReq_REPORT_BY_INSTRUMENT_LEGS_BELONGING_TO_THE_MULTILEG_SECURITY_ONLY');

CREATE TABLE IF NOT EXISTS enum_maps.MultilegModel (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MultilegModel VALUES 
('0', 'MultilegModel_PREDEFINED_MULTILEG_SECURITY'),('1', 'MultilegModel_USER_DEFINED_MULTLEG_SECURITY'),('2', 'MultilegModel_USER_DEFINED_NON_SECURITIZED_MULTILEG');

CREATE TABLE IF NOT EXISTS enum_maps.MultilegPriceMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.MultilegPriceMethod VALUES 
('0', 'MultilegPriceMethod_NET_PRICE'),('1', 'MultilegPriceMethod_REVERSED_NET_PRICE'),('2', 'MultilegPriceMethod_YIELD_DIFFERENCE'),('3', 'MultilegPriceMethod_INDIVIDUAL'),('4', 'MultilegPriceMethod_CONTRACT_WEIGHTED_AVERAGE_PRICE'),('5', 'MultilegPriceMethod_MULTIPLIED_PRICE');

CREATE TABLE IF NOT EXISTS enum_maps.NetGrossInd (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.NetGrossInd VALUES 
('1', 'NetGrossInd_NET'),('2', 'NetGrossInd_GROSS');

CREATE TABLE IF NOT EXISTS enum_maps.NetworkRequestType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.NetworkRequestType VALUES 
('1', 'NetworkRequestType_SNAPSHOT'),('2', 'NetworkRequestType_SUBSCRIBE'),('4', 'NetworkRequestType_STOP_SUBSCRIBING'),('8', 'NetworkRequestType_LEVEL_OF_DETAIL_THEN_NOCOMPIDS_BECOMES_REQUIRED');

CREATE TABLE IF NOT EXISTS enum_maps.NetworkStatusResponseType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.NetworkStatusResponseType VALUES 
('1', 'NetworkStatusResponseType_FULL'),('2', 'NetworkStatusResponseType_INCREMENTAL_UPDATE');

CREATE TABLE IF NOT EXISTS enum_maps.NewsCategory (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.NewsCategory VALUES 
('0', 'NewsCategory_COMPANY_NEWS'),('1', 'NewsCategory_MARKETPLACE_NEWS'),('2', 'NewsCategory_FINANCIAL_MARKET_NEWS'),('3', 'NewsCategory_TECHNICAL_NEWS'),('99', 'NewsCategory_OTHER_NEWS');

CREATE TABLE IF NOT EXISTS enum_maps.NewsRefType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.NewsRefType VALUES 
('0', 'NewsRefType_REPLACEMENT'),('1', 'NewsRefType_OTHER_LANGUAGE'),('2', 'NewsRefType_COMPLIMENTARY');

CREATE TABLE IF NOT EXISTS enum_maps.NoSides (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.NoSides VALUES 
('1', 'NoSides_ONE_SIDE'),('2', 'NoSides_BOTH_SIDES');

CREATE TABLE IF NOT EXISTS enum_maps.NotifyBrokerOfCredit (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.NotifyBrokerOfCredit VALUES 
('N', 'NotifyBrokerOfCredit_NO'),('Y', 'NotifyBrokerOfCredit_YES');

CREATE TABLE IF NOT EXISTS enum_maps.OddLot (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.OddLot VALUES 
('N', 'OddLot_NO'),('Y', 'OddLot_YES');

CREATE TABLE IF NOT EXISTS enum_maps.OpenClose (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.OpenClose VALUES 
('C', 'OpenClose_CLOSE'),('O', 'OpenClose_OPEN');

CREATE TABLE IF NOT EXISTS enum_maps.OpenCloseSettlFlag (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.OpenCloseSettlFlag VALUES 
('0', 'OpenCloseSettlFlag_DAILY_OPEN'),('1', 'OpenCloseSettlFlag_SESSION_OPEN'),('2', 'OpenCloseSettlFlag_DELIVERY_SETTLEMENT_ENTRY'),('3', 'OpenCloseSettlFlag_EXPECTED_ENTRY'),('4', 'OpenCloseSettlFlag_ENTRY_FROM_PREVIOUS_BUSINESS_DAY'),('5', 'OpenCloseSettlFlag_THEORETICAL_PRICE_VALUE');

CREATE TABLE IF NOT EXISTS enum_maps.OpenCloseSettleFlag (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.OpenCloseSettleFlag VALUES 
('0', 'OpenCloseSettleFlag_DAILY_OPEN'),('1', 'OpenCloseSettleFlag_SESSION_OPEN'),('2', 'OpenCloseSettleFlag_DELIVERY_SETTLEMENT_PRICE'),('3', 'OpenCloseSettleFlag_EXPECTED_PRICE'),('4', 'OpenCloseSettleFlag_PRICE_FROM_PREVIOUS_BUSINESS_DAY');

CREATE TABLE IF NOT EXISTS enum_maps.OptPayoutType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.OptPayoutType VALUES 
('1', 'OptPayoutType_VANILLA'),('2', 'OptPayoutType_CAPPED'),('3', 'OptPayoutType_BINARY');

CREATE TABLE IF NOT EXISTS enum_maps.OrdRejReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.OrdRejReason VALUES 
('0', 'OrdRejReason_BROKER'),('1', 'OrdRejReason_UNKNOWN_SYMBOL'),('10', 'OrdRejReason_INVALID_INVESTOR_ID'),('11', 'OrdRejReason_UNSUPPORTED_ORDER_CHARACTERISTIC'),('12', 'OrdRejReason_SURVEILLENCE_OPTION'),('13', 'OrdRejReason_INCORRECT_QUANTITY'),('14', 'OrdRejReason_INCORRECT_ALLOCATED_QUANTITY'),('15', 'OrdRejReason_UNKNOWN_ACCOUNT'),('16', 'OrdRejReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND'),('18', 'OrdRejReason_INVALID_PRICE_INCREMENT'),('2', 'OrdRejReason_EXCHANGE_CLOSED'),('3', 'OrdRejReason_ORDER_EXCEEDS_LIMIT'),('4', 'OrdRejReason_TOO_LATE_TO_ENTER'),('5', 'OrdRejReason_UNKNOWN_ORDER'),('6', 'OrdRejReason_DUPLICATE_ORDER'),('7', 'OrdRejReason_DUPLICATE_OF_A_VERBALLY_COMMUNICATED_ORDER'),('8', 'OrdRejReason_STALE_ORDER'),('9', 'OrdRejReason_TRADE_ALONG_REQUIRED'),('99', 'OrdRejReason_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.OrdStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.OrdStatus VALUES 
('0', 'OrdStatus_NEW'),('1', 'OrdStatus_PARTIALLY_FILLED'),('2', 'OrdStatus_FILLED'),('3', 'OrdStatus_DONE_FOR_DAY'),('4', 'OrdStatus_CANCELED'),('5', 'OrdStatus_REPLACED'),('6', 'OrdStatus_PENDING_CANCEL'),('7', 'OrdStatus_STOPPED'),('8', 'OrdStatus_REJECTED'),('9', 'OrdStatus_SUSPENDED'),('A', 'OrdStatus_PENDING_NEW'),('B', 'OrdStatus_CALCULATED'),('C', 'OrdStatus_EXPIRED'),('D', 'OrdStatus_ACCEPTED_FOR_BIDDING'),('E', 'OrdStatus_PENDING_REPLACE');

CREATE TABLE IF NOT EXISTS enum_maps.OrdType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.OrdType VALUES 
('1', 'OrdType_MARKET'),('2', 'OrdType_LIMIT'),('3', 'OrdType_STOP'),('4', 'OrdType_STOP_LIMIT'),('5', 'OrdType_MARKET_ON_CLOSE'),('6', 'OrdType_WITH_OR_WITHOUT'),('7', 'OrdType_LIMIT_OR_BETTER'),('8', 'OrdType_LIMIT_WITH_OR_WITHOUT'),('9', 'OrdType_ON_BASIS'),('A', 'OrdType_ON_CLOSE'),('B', 'OrdType_LIMIT_ON_CLOSE'),('C', 'OrdType_FOREX_MARKET'),('D', 'OrdType_PREVIOUSLY_QUOTED'),('E', 'OrdType_PREVIOUSLY_INDICATED'),('F', 'OrdType_FOREX_LIMIT'),('G', 'OrdType_FOREX_SWAP'),('H', 'OrdType_FOREX_PREVIOUSLY_QUOTED'),('I', 'OrdType_FUNARI'),('J', 'OrdType_MARKET_IF_TOUCHED'),('K', 'OrdType_MARKET_WITH_LEFT_OVER_AS_LIMIT'),('L', 'OrdType_PREVIOUS_FUND_VALUATION_POINT'),('M', 'OrdType_NEXT_FUND_VALUATION_POINT'),('P', 'OrdType_PEGGED'),('Q', 'OrdType_COUNTER_ORDER_SELECTION');

CREATE TABLE IF NOT EXISTS enum_maps.OrderCapacity (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.OrderCapacity VALUES 
('A', 'OrderCapacity_AGENCY'),('G', 'OrderCapacity_PROPRIETARY'),('I', 'OrderCapacity_INDIVIDUAL'),('P', 'OrderCapacity_PRINCIPAL'),('R', 'OrderCapacity_RISKLESS_PRINCIPAL'),('W', 'OrderCapacity_AGENT_FOR_OTHER_MEMBER');

CREATE TABLE IF NOT EXISTS enum_maps.OrderCategory (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.OrderCategory VALUES 
('1', 'OrderCategory_ORDER'),('2', 'OrderCategory_QUOTE'),('3', 'OrderCategory_PRIVATELY_NEGOTIATED_TRADE'),('4', 'OrderCategory_MULTILEG_ORDER'),('5', 'OrderCategory_LINKED_ORDER'),('6', 'OrderCategory_QUOTE_REQUEST'),('7', 'OrderCategory_IMPLIED_ORDER'),('8', 'OrderCategory_CROSS_ORDER'),('9', 'OrderCategory_STREAMING_PRICE');

CREATE TABLE IF NOT EXISTS enum_maps.OrderDelayUnit (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.OrderDelayUnit VALUES 
('0', 'OrderDelayUnit_SECONDS'),('1', 'OrderDelayUnit_TENTHS_OF_A_SECOND'),('10', 'OrderDelayUnit_MINUTES'),('11', 'OrderDelayUnit_HOURS'),('12', 'OrderDelayUnit_DAYS'),('13', 'OrderDelayUnit_WEEKS'),('14', 'OrderDelayUnit_MONTHS'),('15', 'OrderDelayUnit_YEARS'),('2', 'OrderDelayUnit_HUNDREDTHS_OF_A_SECOND'),('3', 'OrderDelayUnit_MILLISECONDS'),('4', 'OrderDelayUnit_MICROSECONDS'),('5', 'OrderDelayUnit_NANOSECONDS');

CREATE TABLE IF NOT EXISTS enum_maps.OrderHandlingInstSource (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.OrderHandlingInstSource VALUES 
('1', 'OrderHandlingInstSource_NASD_OATS');

CREATE TABLE IF NOT EXISTS enum_maps.OrderRestrictions (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.OrderRestrictions VALUES 
('1', 'OrderRestrictions_PROGRAM_TRADE'),('2', 'OrderRestrictions_INDEX_ARBITRAGE'),('3', 'OrderRestrictions_NON_INDEX_ARBITRAGE'),('4', 'OrderRestrictions_COMPETING_MARKET_MAKER'),('5', 'OrderRestrictions_ACTING_AS_MARKET_MAKER_OR_SPECIALIST_IN_THE_SECURITY'),('6', 'OrderRestrictions_ACTING_AS_MARKET_MAKER_OR_SPECIALIST_IN_THE_UNDERLYING_SECURITY_OF_A_DERIVATIVE_SECURITY'),('7', 'OrderRestrictions_FOREIGN_ENTITY'),('8', 'OrderRestrictions_EXTERNAL_MARKET_PARTICIPANT'),('9', 'OrderRestrictions_EXTERNAL_INTER_CONNECTED_MARKET_LINKAGE'),('A', 'OrderRestrictions_RISKLESS_ARBITRAGE'),('B', 'OrderRestrictions_ISSUER_HOLDING'),('C', 'OrderRestrictions_ISSUE_PRICE_STABILIZATION'),('D', 'OrderRestrictions_NON_ALGORITHMIC'),('E', 'OrderRestrictions_ALGORITHMIC'),('F', 'OrderRestrictions_CROSS');

CREATE TABLE IF NOT EXISTS enum_maps.OrigCustOrderCapacity (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.OrigCustOrderCapacity VALUES 
('1', 'OrigCustOrderCapacity_MEMBER_TRADING_FOR_THEIR_OWN_ACCOUNT'),('2', 'OrigCustOrderCapacity_CLEARING_FIRM_TRADING_FOR_ITS_PROPRIETARY_ACCOUNT'),('3', 'OrigCustOrderCapacity_MEMBER_TRADING_FOR_ANOTHER_MEMBER'),('4', 'OrigCustOrderCapacity_ALL_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.OwnerType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.OwnerType VALUES 
('1', 'OwnerType_INDIVIDUAL_INVESTOR'),('10', 'OwnerType_NETWORKING_SUB_ACCOUNT'),('11', 'OwnerType_NON_PROFIT_ORGANIZATION'),('12', 'OwnerType_CORPORATE_BODY'),('13', 'OwnerType_NOMINEE'),('2', 'OwnerType_PUBLIC_COMPANY'),('3', 'OwnerType_PRIVATE_COMPANY'),('4', 'OwnerType_INDIVIDUAL_TRUSTEE'),('5', 'OwnerType_COMPANY_TRUSTEE'),('6', 'OwnerType_PENSION_PLAN'),('7', 'OwnerType_CUSTODIAN_UNDER_GIFTS_TO_MINORS_ACT'),('8', 'OwnerType_TRUSTS'),('9', 'OwnerType_FIDUCIARIES');

CREATE TABLE IF NOT EXISTS enum_maps.OwnershipType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.OwnershipType VALUES 
('2', 'OwnershipType_JOINT_TRUSTEES'),('J', 'OwnershipType_JOINT_INVESTORS'),('T', 'OwnershipType_TENANTS_IN_COMMON');

CREATE TABLE IF NOT EXISTS enum_maps.PartyDetailsRequestResult (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PartyDetailsRequestResult VALUES 
('0', 'PartyDetailsRequestResult_VALID_REQUEST'),('1', 'PartyDetailsRequestResult_INVALID_OR_UNSUPPORTED_REQUEST'),('2', 'PartyDetailsRequestResult_NO_PARTIES_OR_PARTY_DETAILS_FOUND_THAT_MATCH_SELECTION_CRITERIA'),('3', 'PartyDetailsRequestResult_UNSUPPORTED_PARTYLISTRESPONSETYPE'),('4', 'PartyDetailsRequestResult_NOT_AUTHORIZED_TO_RETRIEVE_PARTIES_OR_PARTY_DETAILS_DATA'),('5', 'PartyDetailsRequestResult_PARTIES_OR_PARTY_DETAILS_DATA_TEMPORARILY_UNAVAILABLE'),('6', 'PartyDetailsRequestResult_REQUEST_FOR_PARTIES_DATA_NOT_SUPPORTED'),('99', 'PartyDetailsRequestResult_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.PartyIDSource (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PartyIDSource VALUES 
('1', 'PartyIDSource_KOREAN_INVESTOR_ID'),('2', 'PartyIDSource_TAIWANESE_QUALIFIED_FOREIGN_INVESTOR_ID_QFII_FID'),('3', 'PartyIDSource_TAIWANESE_TRADING_ACCT'),('4', 'PartyIDSource_MALAYSIAN_CENTRAL_DEPOSITORY'),('5', 'PartyIDSource_CHINESE_INVESTOR_ID'),('6', 'PartyIDSource_UK_NATIONAL_INSURANCE_OR_PENSION_NUMBER'),('7', 'PartyIDSource_US_SOCIAL_SECURITY_NUMBER'),('8', 'PartyIDSource_US_EMPLOYER_OR_TAX_ID_NUMBER'),('9', 'PartyIDSource_AUSTRALIAN_BUSINESS_NUMBER'),('A', 'PartyIDSource_AUSTRALIAN_TAX_FILE_NUMBER'),('B', 'PartyIDSource_BIC'),('C', 'PartyIDSource_GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER'),('D', 'PartyIDSource_PROPRIETARY'),('E', 'PartyIDSource_ISO_COUNTRY_CODE'),('F', 'PartyIDSource_SETTLEMENT_ENTITY_LOCATION'),('G', 'PartyIDSource_MIC'),('H', 'PartyIDSource_CSD_PARTICIPANT_MEMBER_CODE'),('I', 'PartyIDSource_DIRECTED_BROKER_THREE_CHARACTER_ACRONYM_AS_DEFINED_IN_ISITC_ETC_BEST_PRACTICE_GUIDELINES_DOCUMENT');

CREATE TABLE IF NOT EXISTS enum_maps.PartyListResponseType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PartyListResponseType VALUES 
('0', 'PartyListResponseType_RETURN_ALL_AVAILABLE_INFORMATION_ON_PARTIES_AND_RELATED_PARTIES'),('1', 'PartyListResponseType_RETURN_ONLY_PARTY_INFORMATION'),('2', 'PartyListResponseType_INCLUDE_INFORMATION_ON_RELATED_PARTIES'),('3', 'PartyListResponseType_INCLUDE_RISK_LIMIT_INFORMATION');

CREATE TABLE IF NOT EXISTS enum_maps.PartyRelationship (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PartyRelationship VALUES 
('0', 'PartyRelationship_IS_ALSO'),('1', 'PartyRelationship_CLEARS_FOR'),('10', 'PartyRelationship_HAS_MEMBERS'),('11', 'PartyRelationship_PROVIDES_MARKETPLACE_FOR'),('12', 'PartyRelationship_PARTICIPANT_OF_MARKETPLACE'),('13', 'PartyRelationship_CARRIES_POSITIONS_FOR'),('14', 'PartyRelationship_POSTS_TRADES_TO'),('15', 'PartyRelationship_ENTERS_TRADES_FOR'),('16', 'PartyRelationship_ENTERS_TRADES_THROUGH'),('17', 'PartyRelationship_PROVIDES_QUOTES_TO'),('18', 'PartyRelationship_REQUESTS_QUOTES_FROM'),('19', 'PartyRelationship_INVESTS_FOR'),('2', 'PartyRelationship_CLEARS_THROUGH'),('20', 'PartyRelationship_INVESTS_THROUGH'),('21', 'PartyRelationship_BROKERS_TRADES_FOR'),('22', 'PartyRelationship_BROKERS_TRADES_THROUGH'),('23', 'PartyRelationship_PROVIDES_TRADING_SERVICES_FOR'),('24', 'PartyRelationship_USES_TRADING_SERVICES_OF'),('25', 'PartyRelationship_APPROVES_OF'),('26', 'PartyRelationship_APPROVED_BY'),('27', 'PartyRelationship_PARENT_FIRM_FOR'),('28', 'PartyRelationship_SUBSIDIARY_OF'),('29', 'PartyRelationship_REGULATORY_OWNER_OF'),('3', 'PartyRelationship_TRADES_FOR'),('30', 'PartyRelationship_OWNED_BY_30'),('31', 'PartyRelationship_CONTROLS'),('32', 'PartyRelationship_IS_CONTROLLED_BY'),('33', 'PartyRelationship_LEGAL'),('34', 'PartyRelationship_OWNED_BY_34'),('35', 'PartyRelationship_BENEFICIAL_OWNER_OF'),('36', 'PartyRelationship_OWNED_BY_36'),('4', 'PartyRelationship_TRADES_THROUGH'),('5', 'PartyRelationship_SPONSORS'),('6', 'PartyRelationship_SPONSORED_THROUGH'),('7', 'PartyRelationship_PROVIDES_GUARANTEE_FOR'),('8', 'PartyRelationship_IS_GUARANTEED_BY'),('9', 'PartyRelationship_MEMBER_OF');

CREATE TABLE IF NOT EXISTS enum_maps.PartyRole (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PartyRole VALUES 
('1', 'PartyRole_EXECUTING_FIRM'),('10', 'PartyRole_SETTLEMENT_LOCATION'),('11', 'PartyRole_ORDER_ORIGINATION_TRADER'),('12', 'PartyRole_EXECUTING_TRADER'),('13', 'PartyRole_ORDER_ORIGINATION_FIRM'),('14', 'PartyRole_GIVEUP_CLEARING_FIRM'),('15', 'PartyRole_CORRESPONDANT_CLEARING_FIRM'),('16', 'PartyRole_EXECUTING_SYSTEM'),('17', 'PartyRole_CONTRA_FIRM'),('18', 'PartyRole_CONTRA_CLEARING_FIRM'),('19', 'PartyRole_SPONSORING_FIRM'),('2', 'PartyRole_BROKER_OF_CREDIT'),('20', 'PartyRole_UNDERLYING_CONTRA_FIRM'),('21', 'PartyRole_CLEARING_ORGANIZATION'),('22', 'PartyRole_EXCHANGE'),('24', 'PartyRole_CUSTOMER_ACCOUNT'),('25', 'PartyRole_CORRESPONDENT_CLEARING_ORGANIZATION'),('26', 'PartyRole_CORRESPONDENT_BROKER'),('27', 'PartyRole_BUYER_SELLER'),('28', 'PartyRole_CUSTODIAN'),('29', 'PartyRole_INTERMEDIARY'),('3', 'PartyRole_CLIENT_ID'),('30', 'PartyRole_AGENT'),('31', 'PartyRole_SUB_CUSTODIAN'),('32', 'PartyRole_BENEFICIARY'),('33', 'PartyRole_INTERESTED_PARTY'),('34', 'PartyRole_REGULATORY_BODY'),('35', 'PartyRole_LIQUIDITY_PROVIDER'),('36', 'PartyRole_ENTERING_TRADER'),('37', 'PartyRole_CONTRA_TRADER'),('38', 'PartyRole_POSITION_ACCOUNT'),('39', 'PartyRole_CONTRA_INVESTOR_ID'),('4', 'PartyRole_CLEARING_FIRM'),('40', 'PartyRole_TRANSFER_TO_FIRM'),('41', 'PartyRole_CONTRA_POSITION_ACCOUNT'),('42', 'PartyRole_CONTRA_EXCHANGE'),('43', 'PartyRole_INTERNAL_CARRY_ACCOUNT'),('44', 'PartyRole_ORDER_ENTRY_OPERATOR_ID'),('45', 'PartyRole_SECONDARY_ACCOUNT_NUMBER'),('46', 'PartyRole_FOREIGN_FIRM'),('47', 'PartyRole_THIRD_PARTY_ALLOCATION_FIRM'),('48', 'PartyRole_CLAIMING_ACCOUNT'),('49', 'PartyRole_ASSET_MANAGER'),('5', 'PartyRole_INVESTOR_ID'),('50', 'PartyRole_PLEDGOR_ACCOUNT'),('51', 'PartyRole_PLEDGEE_ACCOUNT'),('52', 'PartyRole_LARGE_TRADER_REPORTABLE_ACCOUNT'),('53', 'PartyRole_TRADER_MNEMONIC'),('54', 'PartyRole_SENDER_LOCATION'),('55', 'PartyRole_SESSION_ID'),('56', 'PartyRole_ACCEPTABLE_COUNTERPARTY'),('57', 'PartyRole_UNACCEPTABLE_COUNTERPARTY'),('58', 'PartyRole_ENTERING_UNIT'),('59', 'PartyRole_EXECUTING_UNIT'),('6', 'PartyRole_INTRODUCING_FIRM'),('60', 'PartyRole_INTRODUCING_BROKER'),('61', 'PartyRole_QUOTE_ORIGINATOR'),('62', 'PartyRole_REPORT_ORIGINATOR'),('63', 'PartyRole_SYSTEMATIC_INTERNALISER'),('64', 'PartyRole_MULTILATERAL_TRADING_FACILITY'),('65', 'PartyRole_REGULATED_MARKET'),('66', 'PartyRole_MARKET_MAKER'),('67', 'PartyRole_INVESTMENT_FIRM'),('68', 'PartyRole_HOST_COMPETENT_AUTHORITY'),('69', 'PartyRole_HOME_COMPETENT_AUTHORITY'),('7', 'PartyRole_ENTERING_FIRM'),('70', 'PartyRole_COMPETENT_AUTHORITY_OF_THE_MOST_RELEVANT_MARKET_IN_TERMS_OF_LIQUIDITY'),('71', 'PartyRole_COMPETENT_AUTHORITY_OF_THE_TRANSACTION'),('72', 'PartyRole_REPORTING_INTERMEDIARY'),('73', 'PartyRole_EXECUTION_VENUE'),('74', 'PartyRole_MARKET_DATA_ENTRY_ORIGINATOR'),('75', 'PartyRole_LOCATION_ID'),('76', 'PartyRole_DESK_ID'),('77', 'PartyRole_MARKET_DATA_MARKET'),('78', 'PartyRole_ALLOCATION_ENTITY'),('79', 'PartyRole_PRIME_BROKER_PROVIDING_GENERAL_TRADE_SERVICES'),('8', 'PartyRole_LOCATE'),('80', 'PartyRole_STEP_OUT_FIRM'),('81', 'PartyRole_BROKERCLEARINGID'),('82', 'PartyRole_CENTRAL_REGISTRATION_DEPOSITORY'),('83', 'PartyRole_CLEARING_ACCOUNT'),('84', 'PartyRole_ACCEPTABLE_SETTLING_COUNTERPARTY'),('85', 'PartyRole_UNACCEPTABLE_SETTLING_COUNTERPARTY'),('9', 'PartyRole_FUND_MANAGER_CLIENT_ID');

CREATE TABLE IF NOT EXISTS enum_maps.PartySubIDType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PartySubIDType VALUES 
('1', 'PartySubIDType_FIRM'),('10', 'PartySubIDType_SECURITIES_ACCOUNT_NUMBER'),('11', 'PartySubIDType_REGISTRATION_NUMBER'),('12', 'PartySubIDType_REGISTERED_ADDRESS_12'),('13', 'PartySubIDType_REGULATORY_STATUS'),('14', 'PartySubIDType_REGISTRATION_NAME'),('15', 'PartySubIDType_CASH_ACCOUNT_NUMBER'),('16', 'PartySubIDType_BIC'),('17', 'PartySubIDType_CSD_PARTICIPANT_MEMBER_CODE'),('18', 'PartySubIDType_REGISTERED_ADDRESS_18'),('19', 'PartySubIDType_FUND_ACCOUNT_NAME'),('2', 'PartySubIDType_PERSON'),('20', 'PartySubIDType_TELEX_NUMBER'),('21', 'PartySubIDType_FAX_NUMBER'),('22', 'PartySubIDType_SECURITIES_ACCOUNT_NAME'),('23', 'PartySubIDType_CASH_ACCOUNT_NAME'),('24', 'PartySubIDType_DEPARTMENT'),('25', 'PartySubIDType_LOCATION_DESK'),('26', 'PartySubIDType_POSITION_ACCOUNT_TYPE'),('27', 'PartySubIDType_SECURITY_LOCATE_ID'),('28', 'PartySubIDType_MARKET_MAKER'),('29', 'PartySubIDType_ELIGIBLE_COUNTERPARTY'),('3', 'PartySubIDType_SYSTEM'),('30', 'PartySubIDType_PROFESSIONAL_CLIENT'),('31', 'PartySubIDType_LOCATION'),('32', 'PartySubIDType_EXECUTION_VENUE'),('33', 'PartySubIDType_CURRENCY_DELIVERY_IDENTIFIER'),('4', 'PartySubIDType_APPLICATION'),('4000', 'PartySubIDType_RESERVEDANDAVAILABLEFORBILATERALLYAGREEDUPONUSERDEFINEDVALUES'),('5', 'PartySubIDType_FULL_LEGAL_NAME_OF_FIRM'),('6', 'PartySubIDType_POSTAL_ADDRESS'),('7', 'PartySubIDType_PHONE_NUMBER'),('8', 'PartySubIDType_EMAIL_ADDRESS'),('9', 'PartySubIDType_CONTACT_NAME');

CREATE TABLE IF NOT EXISTS enum_maps.PaymentMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PaymentMethod VALUES 
('1', 'PaymentMethod_CREST'),('10', 'PaymentMethod_DIRECT_CREDIT'),('11', 'PaymentMethod_CREDIT_CARD'),('12', 'PaymentMethod_ACH_DEBIT'),('13', 'PaymentMethod_ACH_CREDIT'),('14', 'PaymentMethod_BPAY'),('15', 'PaymentMethod_HIGH_VALUE_CLEARING_SYSTEM'),('2', 'PaymentMethod_NSCC'),('3', 'PaymentMethod_EUROCLEAR'),('4', 'PaymentMethod_CLEARSTREAM'),('5', 'PaymentMethod_CHEQUE'),('6', 'PaymentMethod_TELEGRAPHIC_TRANSFER'),('7', 'PaymentMethod_FED_WIRE'),('8', 'PaymentMethod_DEBIT_CARD'),('9', 'PaymentMethod_DIRECT_DEBIT');

CREATE TABLE IF NOT EXISTS enum_maps.PegLimitType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PegLimitType VALUES 
('0', 'PegLimitType_OR_BETTER'),('1', 'PegLimitType_STRICT'),('2', 'PegLimitType_OR_WORSE');

CREATE TABLE IF NOT EXISTS enum_maps.PegMoveType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PegMoveType VALUES 
('0', 'PegMoveType_FLOATING'),('1', 'PegMoveType_FIXED');

CREATE TABLE IF NOT EXISTS enum_maps.PegOffsetType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PegOffsetType VALUES 
('0', 'PegOffsetType_PRICE'),('1', 'PegOffsetType_BASIS_POINTS'),('2', 'PegOffsetType_TICKS'),('3', 'PegOffsetType_PRICE_TIER');

CREATE TABLE IF NOT EXISTS enum_maps.PegPriceType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PegPriceType VALUES 
('1', 'PegPriceType_LAST_PEG'),('2', 'PegPriceType_MID_PRICE_PEG'),('3', 'PegPriceType_OPENING_PEG'),('4', 'PegPriceType_MARKET_PEG'),('5', 'PegPriceType_PRIMARY_PEG'),('6', 'PegPriceType_FIXED_PEG_TO_LOCAL_BEST_BID_OR_OFFER_AT_TIME_OF_ORDER'),('7', 'PegPriceType_PEG_TO_VWAP'),('8', 'PegPriceType_TRAILING_STOP_PEG'),('9', 'PegPriceType_PEG_TO_LIMIT_PRICE');

CREATE TABLE IF NOT EXISTS enum_maps.PegRoundDirection (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PegRoundDirection VALUES 
('1', 'PegRoundDirection_MORE_AGGRESSIVE'),('2', 'PegRoundDirection_MORE_PASSIVE');

CREATE TABLE IF NOT EXISTS enum_maps.PegScope (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PegScope VALUES 
('1', 'PegScope_LOCAL'),('2', 'PegScope_NATIONAL'),('3', 'PegScope_GLOBAL'),('4', 'PegScope_NATIONAL_EXCLUDING_LOCAL');

CREATE TABLE IF NOT EXISTS enum_maps.PosAmtType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PosAmtType VALUES 
('ACPN', 'PosAmtType_ACCRUED_COUPON_AMOUNT'),('BANK', 'PosAmtType_TOTAL_BANKED_AMOUNT'),('CASH', 'PosAmtType_CASH_AMOUNT'),('CMTM', 'PosAmtType_COLLATERALIZED_MARK_TO_MARKET'),('COLAT', 'PosAmtType_TOTAL_COLLATERALIZED_AMOUNT'),('CPN', 'PosAmtType_COUPON_AMOUNT'),('CRES', 'PosAmtType_CASH_RESIDUAL_AMOUNT'),('DLV', 'PosAmtType_COMPENSATION_AMOUNT'),('FMTM', 'PosAmtType_FINAL_MARK_TO_MARKET_AMOUNT'),('IACPN', 'PosAmtType_INCREMENTAL_ACCRUED_COUPON'),('ICMTM', 'PosAmtType_INCREMENTAL_COLLATERALIZED_MARK_TO_MARKET'),('ICPN', 'PosAmtType_INITIAL_TRADE_COUPON_AMOUNT'),('IMTM', 'PosAmtType_INCREMENTAL_MARK_TO_MARKET_AMOUNT'),('PREM', 'PosAmtType_PREMIUM_AMOUNT'),('SETL', 'PosAmtType_SETTLEMENT_VALUE'),('SMTM', 'PosAmtType_START_OF_DAY_MARK_TO_MARKET_AMOUNT'),('TVAR', 'PosAmtType_TRADE_VARIATION_AMOUNT'),('VADJ', 'PosAmtType_VALUE_ADJUSTED_AMOUNT');

CREATE TABLE IF NOT EXISTS enum_maps.PosMaintAction (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PosMaintAction VALUES 
('1', 'PosMaintAction_NEW'),('2', 'PosMaintAction_REPLACE'),('3', 'PosMaintAction_CANCEL'),('4', 'PosMaintAction_REVERSE');

CREATE TABLE IF NOT EXISTS enum_maps.PosMaintResult (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PosMaintResult VALUES 
('0', 'PosMaintResult_SUCCESSFUL_COMPLETION'),('1', 'PosMaintResult_REJECTED'),('99', 'PosMaintResult_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.PosMaintStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PosMaintStatus VALUES 
('0', 'PosMaintStatus_ACCEPTED'),('1', 'PosMaintStatus_ACCEPTED_WITH_WARNINGS'),('2', 'PosMaintStatus_REJECTED'),('3', 'PosMaintStatus_COMPLETED'),('4', 'PosMaintStatus_COMPLETED_WITH_WARNINGS');

CREATE TABLE IF NOT EXISTS enum_maps.PosQtyStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PosQtyStatus VALUES 
('0', 'PosQtyStatus_SUBMITTED'),('1', 'PosQtyStatus_ACCEPTED'),('2', 'PosQtyStatus_REJECTED');

CREATE TABLE IF NOT EXISTS enum_maps.PosReqResult (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PosReqResult VALUES 
('0', 'PosReqResult_VALID_REQUEST'),('1', 'PosReqResult_INVALID_OR_UNSUPPORTED_REQUEST'),('2', 'PosReqResult_NO_POSITIONS_FOUND_THAT_MATCH_CRITERIA'),('3', 'PosReqResult_NOT_AUTHORIZED_TO_REQUEST_POSITIONS'),('4', 'PosReqResult_REQUEST_FOR_POSITION_NOT_SUPPORTED'),('99', 'PosReqResult_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.PosReqStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PosReqStatus VALUES 
('0', 'PosReqStatus_COMPLETED'),('1', 'PosReqStatus_COMPLETED_WITH_WARNINGS'),('2', 'PosReqStatus_REJECTED');

CREATE TABLE IF NOT EXISTS enum_maps.PosReqType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PosReqType VALUES 
('0', 'PosReqType_POSITIONS'),('1', 'PosReqType_TRADES'),('2', 'PosReqType_EXERCISES'),('3', 'PosReqType_ASSIGNMENTS'),('4', 'PosReqType_SETTLEMENT_ACTIVITY'),('5', 'PosReqType_BACKOUT_MESSAGE'),('6', 'PosReqType_DELTA_POSITIONS');

CREATE TABLE IF NOT EXISTS enum_maps.PosTransType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PosTransType VALUES 
('1', 'PosTransType_EXERCISE'),('2', 'PosTransType_DO_NOT_EXERCISE'),('3', 'PosTransType_POSITION_ADJUSTMENT'),('4', 'PosTransType_POSITION_CHANGE_SUBMISSION_MARGIN_DISPOSITION'),('5', 'PosTransType_PLEDGE'),('6', 'PosTransType_LARGE_TRADER_SUBMISSION');

CREATE TABLE IF NOT EXISTS enum_maps.PosType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PosType VALUES 
('ALC', 'PosType_ALLOCATION_TRADE_QTY'),('AS', 'PosType_OPTION_ASSIGNMENT'),('ASF', 'PosType_AS_OF_TRADE_QTY'),('CAA', 'PosType_CORPORATE_ACTION_ADJUSTMENT'),('CEA', 'PosType_CREDIT_EVENT_ADJUSTMENT'),('DLT', 'PosType_NET_DELTA_QTY'),('DLV', 'PosType_DELIVERY_QTY'),('DN', 'PosType_DELIVERY_NOTICE_QTY'),('EP', 'PosType_EXCHANGE_FOR_PHYSICAL_QTY'),('ETR', 'PosType_ELECTRONIC_TRADE_QTY'),('EX', 'PosType_OPTION_EXERCISE_QTY'),('FIN', 'PosType_END_OF_DAY_QTY'),('IAS', 'PosType_INTRA_SPREAD_QTY'),('IES', 'PosType_INTER_SPREAD_QTY'),('PA', 'PosType_ADJUSTMENT_QTY'),('PIT', 'PosType_PIT_TRADE_QTY'),('PNTN', 'PosType_PRIVATELY_NEGOTIATED_TRADE_QTY'),('RCV', 'PosType_RECEIVE_QUANTITY'),('SEA', 'PosType_SUCCESSION_EVENT_ADJUSTMENT'),('SOD', 'PosType_START_OF_DAY_QTY'),('SPL', 'PosType_INTEGRAL_SPLIT'),('TA', 'PosType_TRANSACTION_FROM_ASSIGNMENT'),('TOT', 'PosType_TOTAL_TRANSACTION_QTY'),('TQ', 'PosType_TRANSACTION_QUANTITY'),('TRF', 'PosType_TRANSFER_TRADE_QTY'),('TX', 'PosType_TRANSACTION_FROM_EXERCISE'),('XM', 'PosType_CROSS_MARGIN_QTY');

CREATE TABLE IF NOT EXISTS enum_maps.PositionEffect (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PositionEffect VALUES 
('C', 'PositionEffect_CLOSE'),('D', 'PositionEffect_DEFAULT'),('F', 'PositionEffect_FIFO'),('N', 'PositionEffect_CLOSE_BUT_NOTIFY_ON_OPEN'),('O', 'PositionEffect_OPEN'),('R', 'PositionEffect_ROLLED');

CREATE TABLE IF NOT EXISTS enum_maps.PossDupFlag (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PossDupFlag VALUES 
('N', 'PossDupFlag_NO'),('Y', 'PossDupFlag_YES');

CREATE TABLE IF NOT EXISTS enum_maps.PossResend (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PossResend VALUES 
('N', 'PossResend_NO'),('Y', 'PossResend_YES');

CREATE TABLE IF NOT EXISTS enum_maps.PreallocMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PreallocMethod VALUES 
('0', 'PreallocMethod_PRO_RATA'),('1', 'PreallocMethod_DO_NOT_PRO_RATA');

CREATE TABLE IF NOT EXISTS enum_maps.PreviouslyReported (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PreviouslyReported VALUES 
('N', 'PreviouslyReported_NO'),('Y', 'PreviouslyReported_YES');

CREATE TABLE IF NOT EXISTS enum_maps.PriceLimitType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PriceLimitType VALUES 
('0', 'PriceLimitType_PRICE'),('1', 'PriceLimitType_TICKS'),('2', 'PriceLimitType_PERCENTAGE');

CREATE TABLE IF NOT EXISTS enum_maps.PriceProtectionScope (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PriceProtectionScope VALUES 
('0', 'PriceProtectionScope_NONE'),('1', 'PriceProtectionScope_LOCAL'),('2', 'PriceProtectionScope_NATIONAL'),('3', 'PriceProtectionScope_GLOBAL');

CREATE TABLE IF NOT EXISTS enum_maps.PriceQuoteMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PriceQuoteMethod VALUES 
('INT', 'PriceQuoteMethod_INTEREST_RATE_INDEX'),('INX', 'PriceQuoteMethod_INDEX'),('PCTPAR', 'PriceQuoteMethod_PERCENT_OF_PAR'),('STD', 'PriceQuoteMethod_STANDARD');

CREATE TABLE IF NOT EXISTS enum_maps.PriceType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PriceType VALUES 
('1', 'PriceType_PERCENTAGE'),('10', 'PriceType_FIXED_CABINET_TRADE_PRICE'),('11', 'PriceType_VARIABLE_CABINET_TRADE_PRICE'),('13', 'PriceType_PRODUCT_TICKS_IN_HALFS'),('14', 'PriceType_PRODUCT_TICKS_IN_FOURTHS'),('15', 'PriceType_PRODUCT_TICKS_IN_EIGHTS'),('16', 'PriceType_PRODUCT_TICKS_IN_SIXTEENTHS'),('17', 'PriceType_PRODUCT_TICKS_IN_THIRTY_SECONDS'),('18', 'PriceType_PRODUCT_TICKS_IN_SIXTY_FORTHS'),('19', 'PriceType_PRODUCT_TICKS_IN_ONE_TWENTY_EIGHTS'),('2', 'PriceType_PER_UNIT'),('3', 'PriceType_FIXED_AMOUNT'),('4', 'PriceType_DISCOUNT'),('5', 'PriceType_PREMIUM'),('6', 'PriceType_SPREAD'),('7', 'PriceType_TED_PRICE'),('8', 'PriceType_TED_YIELD'),('9', 'PriceType_YIELD');

CREATE TABLE IF NOT EXISTS enum_maps.PriorityIndicator (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PriorityIndicator VALUES 
('0', 'PriorityIndicator_PRIORITY_UNCHANGED'),('1', 'PriorityIndicator_LOST_PRIORITY_AS_RESULT_OF_ORDER_CHANGE');

CREATE TABLE IF NOT EXISTS enum_maps.ProcessCode (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ProcessCode VALUES 
('0', 'ProcessCode_REGULAR'),('1', 'ProcessCode_SOFT_DOLLAR'),('2', 'ProcessCode_STEP_IN'),('3', 'ProcessCode_STEP_OUT'),('4', 'ProcessCode_SOFT_DOLLAR_STEP_IN'),('5', 'ProcessCode_SOFT_DOLLAR_STEP_OUT'),('6', 'ProcessCode_PLAN_SPONSOR');

CREATE TABLE IF NOT EXISTS enum_maps.Product (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.Product VALUES 
('1', 'Product_AGENCY'),('10', 'Product_MORTGAGE'),('11', 'Product_MUNICIPAL'),('12', 'Product_OTHER'),('13', 'Product_FINANCING'),('2', 'Product_COMMODITY'),('3', 'Product_CORPORATE'),('4', 'Product_CURRENCY'),('5', 'Product_EQUITY'),('6', 'Product_GOVERNMENT'),('7', 'Product_INDEX'),('8', 'Product_LOAN'),('9', 'Product_MONEYMARKET');

CREATE TABLE IF NOT EXISTS enum_maps.ProgRptReqs (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ProgRptReqs VALUES 
('1', 'ProgRptReqs_BUY_SIDE_EXPLICITLY_REQUESTS_STATUS_USING_STATUE_REQUEST'),('2', 'ProgRptReqs_SELL_SIDE_PERIODICALLY_SENDS_STATUS_USING_LIST_STATUS_PERIOD_OPTIONALLY_SPECIFIED_IN_PROGRESSPERIOD'),('3', 'ProgRptReqs_REAL_TIME_EXECUTION_REPORTS');

CREATE TABLE IF NOT EXISTS enum_maps.PublishTrdIndicator (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PublishTrdIndicator VALUES 
('N', 'PublishTrdIndicator_NO'),('Y', 'PublishTrdIndicator_YES');

CREATE TABLE IF NOT EXISTS enum_maps.PutOrCall (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.PutOrCall VALUES 
('0', 'PutOrCall_PUT'),('1', 'PutOrCall_CALL');

CREATE TABLE IF NOT EXISTS enum_maps.QtyType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.QtyType VALUES 
('0', 'QtyType_UNITS'),('1', 'QtyType_CONTRACTS'),('2', 'QtyType_UNITS_OF_MEASURE_PER_TIME_UNIT');

CREATE TABLE IF NOT EXISTS enum_maps.QuantityType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.QuantityType VALUES 
('1', 'QuantityType_SHARES'),('2', 'QuantityType_BONDS'),('3', 'QuantityType_CURRENTFACE'),('4', 'QuantityType_ORIGINALFACE'),('5', 'QuantityType_CURRENCY'),('6', 'QuantityType_CONTRACTS'),('7', 'QuantityType_OTHER'),('8', 'QuantityType_PAR');

CREATE TABLE IF NOT EXISTS enum_maps.QuoteAckStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.QuoteAckStatus VALUES 
('0', 'QuoteAckStatus_ACCEPTED'),('1', 'QuoteAckStatus_CANCELED_FOR_SYMBOL'),('2', 'QuoteAckStatus_CANCELED_FOR_SECURITY_TYPE'),('3', 'QuoteAckStatus_CANCELED_FOR_UNDERLYING'),('4', 'QuoteAckStatus_CANCELED_ALL'),('5', 'QuoteAckStatus_REJECTED');

CREATE TABLE IF NOT EXISTS enum_maps.QuoteCancelType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.QuoteCancelType VALUES 
('1', 'QuoteCancelType_CANCEL_FOR_ONE_OR_MORE_SECURITIES'),('2', 'QuoteCancelType_CANCEL_FOR_SECURITY_TYPE'),('3', 'QuoteCancelType_CANCEL_FOR_UNDERLYING_SECURITY'),('4', 'QuoteCancelType_CANCEL_ALL_QUOTES'),('5', 'QuoteCancelType_CANCEL_QUOTE_SPECIFIED_IN_QUOTEID'),('6', 'QuoteCancelType_CANCEL_BY_QUOTETYPE'),('7', 'QuoteCancelType_CANCEL_FOR_SECURITY_ISSUER'),('8', 'QuoteCancelType_CANCEL_FOR_ISSUER_OF_UNDERLYING_SECURITY');

CREATE TABLE IF NOT EXISTS enum_maps.QuoteCondition (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.QuoteCondition VALUES 
('0', 'QuoteCondition_RESERVED_SAM'),('1', 'QuoteCondition_NO_ACTIVE_SAM'),('2', 'QuoteCondition_RESTRICTED'),('3', 'QuoteCondition_REST_OF_BOOK_VWAP'),('4', 'QuoteCondition_BETTER_PRICES_IN_CONDITIONAL_ORDERS'),('5', 'QuoteCondition_MEDIAN_PRICE'),('6', 'QuoteCondition_FULL_CURVE'),('7', 'QuoteCondition_FLAT_CURVE'),('A', 'QuoteCondition_OPEN_ACTIVE'),('B', 'QuoteCondition_CLOSED_INACTIVE'),('C', 'QuoteCondition_EXCHANGE_BEST'),('D', 'QuoteCondition_CONSOLIDATED_BEST'),('E', 'QuoteCondition_LOCKED'),('F', 'QuoteCondition_CROSSED'),('G', 'QuoteCondition_DEPTH'),('H', 'QuoteCondition_FAST_TRADING'),('I', 'QuoteCondition_NON_FIRM'),('J', 'QuoteCondition_OUTRIGHT_PRICE'),('K', 'QuoteCondition_IMPLIED_PRICE'),('L', 'QuoteCondition_MANUAL_SLOW_QUOTE'),('M', 'QuoteCondition_DEPTH_ON_OFFER'),('N', 'QuoteCondition_DEPTH_ON_BID'),('O', 'QuoteCondition_CLOSING'),('P', 'QuoteCondition_NEWS_DISSEMINATION'),('Q', 'QuoteCondition_TRADING_RANGE'),('R', 'QuoteCondition_ORDER_INFLUX'),('S', 'QuoteCondition_DUE_TO_RELATED'),('T', 'QuoteCondition_NEWS_PENDING'),('U', 'QuoteCondition_ADDITIONAL_INFO'),('V', 'QuoteCondition_ADDITIONAL_INFO_DUE_TO_RELATED'),('W', 'QuoteCondition_RESUME'),('X', 'QuoteCondition_VIEW_OF_COMMON'),('Y', 'QuoteCondition_VOLUME_ALERT'),('Z', 'QuoteCondition_ORDER_IMBALANCE'),('a', 'QuoteCondition_EQUIPMENT_CHANGEOVER'),('b', 'QuoteCondition_NO_OPEN'),('c', 'QuoteCondition_REGULAR_ETH'),('d', 'QuoteCondition_AUTOMATIC_EXECUTION'),('e', 'QuoteCondition_AUTOMATIC_EXECUTION_ETH'),('f', 'QuoteCondition_FAST_MARKET_ETH'),('g', 'QuoteCondition_INACTIVE_ETH'),('h', 'QuoteCondition_ROTATION'),('i', 'QuoteCondition_ROTATION_ETH'),('j', 'QuoteCondition_HALT'),('k', 'QuoteCondition_HALT_ETH'),('l', 'QuoteCondition_DUE_TO_NEWS_DISSEMINATION'),('m', 'QuoteCondition_DUE_TO_NEWS_PENDING'),('n', 'QuoteCondition_TRADING_RESUME'),('o', 'QuoteCondition_OUT_OF_SEQUENCE'),('p', 'QuoteCondition_BID_SPECIALIST'),('q', 'QuoteCondition_OFFER_SPECIALIST'),('r', 'QuoteCondition_BID_OFFER_SPECIALIST'),('s', 'QuoteCondition_END_OF_DAY_SAM'),('t', 'QuoteCondition_FORBIDDEN_SAM'),('u', 'QuoteCondition_FROZEN_SAM'),('v', 'QuoteCondition_PREOPENING_SAM'),('w', 'QuoteCondition_OPENING_SAM'),('x', 'QuoteCondition_OPEN_SAM'),('y', 'QuoteCondition_SURVEILLANCE_SAM'),('z', 'QuoteCondition_SUSPENDED_SAM');

CREATE TABLE IF NOT EXISTS enum_maps.QuoteEntryRejectReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.QuoteEntryRejectReason VALUES 
('1', 'QuoteEntryRejectReason_UNKNOWN_SYMBOL'),('2', 'QuoteEntryRejectReason_EXHCNAGE'),('3', 'QuoteEntryRejectReason_QUOTE_EXCEEDS_LIMIT'),('4', 'QuoteEntryRejectReason_TOO_LATE_TO_ENTER'),('5', 'QuoteEntryRejectReason_UNKNOWN_QUOTE'),('6', 'QuoteEntryRejectReason_DUPLICATE_QUOTE'),('7', 'QuoteEntryRejectReason_INVALID_BID_ASK_SPREAD'),('8', 'QuoteEntryRejectReason_INVALID_PRICE'),('9', 'QuoteEntryRejectReason_NOT_AUTHORIZED_TO_QUOTE_SECURITY'),('99', 'QuoteEntryRejectReason_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.QuoteEntryStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.QuoteEntryStatus VALUES 
('0', 'QuoteEntryStatus_ACCEPTED'),('12', 'QuoteEntryStatus_LOCKED_MARKET_WARNING'),('13', 'QuoteEntryStatus_CROSS_MARKET_WARNING'),('14', 'QuoteEntryStatus_CANCELED_DUE_TO_LOCK_MARKET'),('15', 'QuoteEntryStatus_CANCELED_DUE_TO_CROSS_MARKET'),('16', 'QuoteEntryStatus_ACTIVE'),('5', 'QuoteEntryStatus_REJECTED'),('6', 'QuoteEntryStatus_REMOVED_FROM_MARKET'),('7', 'QuoteEntryStatus_EXPIRED');

CREATE TABLE IF NOT EXISTS enum_maps.QuotePriceType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.QuotePriceType VALUES 
('1', 'QuotePriceType_PERCENT'),('10', 'QuotePriceType_YIELD'),('2', 'QuotePriceType_PER_SHARE'),('3', 'QuotePriceType_FIXED_AMOUNT'),('4', 'QuotePriceType_DISCOUNT'),('5', 'QuotePriceType_PREMIUM'),('6', 'QuotePriceType_SPREAD'),('7', 'QuotePriceType_TED_PRICE'),('8', 'QuotePriceType_TED_YIELD'),('9', 'QuotePriceType_YIELD_SPREAD');

CREATE TABLE IF NOT EXISTS enum_maps.QuoteRejectReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.QuoteRejectReason VALUES 
('1', 'QuoteRejectReason_UNKNOWN_SYMBOL'),('10', 'QuoteRejectReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND'),('11', 'QuoteRejectReason_QUOTE_LOCKED'),('12', 'QuoteRejectReason_INVALID_OR_UNKNOWN_SECURITY_ISSUER'),('13', 'QuoteRejectReason_INVALID_OR_UNKNOWN_ISSUER_OF_UNDERLYING_SECURITY'),('2', 'QuoteRejectReason_EXCHANGE'),('3', 'QuoteRejectReason_QUOTE_REQUEST_EXCEEDS_LIMIT'),('4', 'QuoteRejectReason_TOO_LATE_TO_ENTER'),('5', 'QuoteRejectReason_UNKNOWN_QUOTE'),('6', 'QuoteRejectReason_DUPLICATE_QUOTE'),('7', 'QuoteRejectReason_INVALID_BID_ASK_SPREAD'),('8', 'QuoteRejectReason_INVALID_PRICE'),('9', 'QuoteRejectReason_NOT_AUTHORIZED_TO_QUOTE_SECURITY'),('99', 'QuoteRejectReason_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.QuoteRequestRejectReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.QuoteRequestRejectReason VALUES 
('1', 'QuoteRequestRejectReason_UNKNOWN_SYMBOL'),('10', 'QuoteRequestRejectReason_PASS'),('11', 'QuoteRequestRejectReason_INSUFFICIENT_CREDIT'),('2', 'QuoteRequestRejectReason_EXCHANGE'),('3', 'QuoteRequestRejectReason_QUOTE_REQUEST_EXCEEDS_LIMIT'),('4', 'QuoteRequestRejectReason_TOO_LATE_TO_ENTER'),('5', 'QuoteRequestRejectReason_INVALID_PRICE'),('6', 'QuoteRequestRejectReason_NOT_AUTHORIZED_TO_REQUEST_QUOTE'),('7', 'QuoteRequestRejectReason_NO_MATCH_FOR_INQUIRY'),('8', 'QuoteRequestRejectReason_NO_MARKET_FOR_INSTRUMENT'),('9', 'QuoteRequestRejectReason_NO_INVENTORY'),('99', 'QuoteRequestRejectReason_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.QuoteRequestType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.QuoteRequestType VALUES 
('1', 'QuoteRequestType_MANUAL'),('2', 'QuoteRequestType_AUTOMATIC');

CREATE TABLE IF NOT EXISTS enum_maps.QuoteRespType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.QuoteRespType VALUES 
('1', 'QuoteRespType_HIT_LIFT'),('2', 'QuoteRespType_COUNTER'),('3', 'QuoteRespType_EXPIRED'),('4', 'QuoteRespType_COVER'),('5', 'QuoteRespType_DONE_AWAY'),('6', 'QuoteRespType_PASS'),('7', 'QuoteRespType_END_TRADE'),('8', 'QuoteRespType_TIMED_OUT');

CREATE TABLE IF NOT EXISTS enum_maps.QuoteResponseLevel (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.QuoteResponseLevel VALUES 
('0', 'QuoteResponseLevel_NO_ACKNOWLEDGEMENT'),('1', 'QuoteResponseLevel_ACKNOWLEDGE_ONLY_NEGATIVE_OR_ERRONEOUS_QUOTES'),('2', 'QuoteResponseLevel_ACKNOWLEDGE_EACH_QUOTE_MESSAGE'),('3', 'QuoteResponseLevel_SUMMARY_ACKNOWLEDGEMENT');

CREATE TABLE IF NOT EXISTS enum_maps.QuoteStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.QuoteStatus VALUES 
('0', 'QuoteStatus_ACCEPTED'),('1', 'QuoteStatus_CANCEL_FOR_SYMBOL'),('10', 'QuoteStatus_PENDING'),('11', 'QuoteStatus_PASS'),('12', 'QuoteStatus_LOCKED_MARKET_WARNING'),('13', 'QuoteStatus_CROSS_MARKET_WARNING'),('14', 'QuoteStatus_CANCELED_DUE_TO_LOCK_MARKET'),('15', 'QuoteStatus_CANCELED_DUE_TO_CROSS_MARKET'),('16', 'QuoteStatus_ACTIVE'),('17', 'QuoteStatus_CANCELED'),('18', 'QuoteStatus_UNSOLICITED_QUOTE_REPLENISHMENT'),('19', 'QuoteStatus_PENDING_END_TRADE'),('2', 'QuoteStatus_CANCELED_FOR_SECURITY_TYPE'),('20', 'QuoteStatus_TOO_LATE_TO_END'),('3', 'QuoteStatus_CANCELED_FOR_UNDERLYING'),('4', 'QuoteStatus_CANCELED_ALL'),('5', 'QuoteStatus_REJECTED'),('6', 'QuoteStatus_REMOVED_FROM_MARKET'),('7', 'QuoteStatus_EXPIRED'),('8', 'QuoteStatus_QUERY'),('9', 'QuoteStatus_QUOTE_NOT_FOUND');

CREATE TABLE IF NOT EXISTS enum_maps.QuoteType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.QuoteType VALUES 
('0', 'QuoteType_INDICATIVE'),('1', 'QuoteType_TRADEABLE'),('2', 'QuoteType_RESTRICTED_TRADEABLE'),('3', 'QuoteType_COUNTER');

CREATE TABLE IF NOT EXISTS enum_maps.RateSource (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.RateSource VALUES 
('0', 'RateSource_BLOOMBERG'),('1', 'RateSource_REUTERS'),('2', 'RateSource_TELERATE'),('99', 'RateSource_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.RateSourceType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.RateSourceType VALUES 
('0', 'RateSourceType_PRIMARY'),('1', 'RateSourceType_SECONDARY');

CREATE TABLE IF NOT EXISTS enum_maps.RefOrdIDReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.RefOrdIDReason VALUES 
('0', 'RefOrdIDReason_GTC_FROM_PREVIOUS_DAY'),('1', 'RefOrdIDReason_PARTIAL_FILL_REMAINING'),('2', 'RefOrdIDReason_ORDER_CHANGED');

CREATE TABLE IF NOT EXISTS enum_maps.RefOrderIDSource (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.RefOrderIDSource VALUES 
('0', 'RefOrderIDSource_SECONDARYORDERID'),('1', 'RefOrderIDSource_ORDERID'),('2', 'RefOrderIDSource_MDENTRYID'),('3', 'RefOrderIDSource_QUOTEENTRYID'),('4', 'RefOrderIDSource_ORIGINAL_ORDER_ID');

CREATE TABLE IF NOT EXISTS enum_maps.RegistRejReasonCode (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.RegistRejReasonCode VALUES 
('1', 'RegistRejReasonCode_INVALID_UNACCEPTABLE_ACCOUNT_TYPE'),('10', 'RegistRejReasonCode_INVALID_UNACEEPTABLE_INVESTOR_ID_SOURCE'),('11', 'RegistRejReasonCode_INVALID_UNACCEPTABLE_DATE_OF_BIRTH'),('12', 'RegistRejReasonCode_INVALID_UNACCEPTABLE_INVESTOR_COUNTRY_OF_RESIDENCE'),('13', 'RegistRejReasonCode_INVALID_UNACCEPTABLE_NO_DISTRIB_INSTNS'),('14', 'RegistRejReasonCode_INVALID_UNACCEPTABLE_DISTRIB_PERCENTAGE'),('15', 'RegistRejReasonCode_INVALID_UNACCEPTABLE_DISTRIB_PAYMENT_METHOD'),('16', 'RegistRejReasonCode_INVALID_UNACCEPTABLE_CASH_DISTRIB_AGENT_ACCT_NAME'),('17', 'RegistRejReasonCode_INVALID_UNACCEPTABLE_CASH_DISTRIB_AGENT_CODE'),('18', 'RegistRejReasonCode_INVALID_UNACCEPTABLE_CASH_DISTRIB_AGENT_ACCT_NUM'),('2', 'RegistRejReasonCode_INVALID_UNACCEPTABLE_TAX_EXEMPT_TYPE'),('3', 'RegistRejReasonCode_INVALID_UNACCEPTABLE_OWNERSHIP_TYPE'),('4', 'RegistRejReasonCode_INVALID_UNACCEPTABLE_NO_REG_DETAILS'),('5', 'RegistRejReasonCode_INVALID_UNACCEPTABLE_REG_SEQ_NO'),('6', 'RegistRejReasonCode_INVALID_UNACCEPTABLE_REG_DETAILS'),('7', 'RegistRejReasonCode_INVALID_UNACCEPTABLE_MAILING_DETAILS'),('8', 'RegistRejReasonCode_INVALID_UNACCEPTABLE_MAILING_INSTRUCTIONS'),('9', 'RegistRejReasonCode_INVALID_UNACCEPTABLE_INVESTOR_ID'),('99', 'RegistRejReasonCode_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.RegistStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.RegistStatus VALUES 
('A', 'RegistStatus_ACCEPTED'),('H', 'RegistStatus_HELD'),('N', 'RegistStatus_REMINDER'),('R', 'RegistStatus_REJECTED');

CREATE TABLE IF NOT EXISTS enum_maps.RegistTransType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.RegistTransType VALUES 
('0', 'RegistTransType_NEW'),('1', 'RegistTransType_REPLACE'),('2', 'RegistTransType_CANCEL');

CREATE TABLE IF NOT EXISTS enum_maps.ReportToExch (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ReportToExch VALUES 
('N', 'ReportToExch_NO'),('Y', 'ReportToExch_YES');

CREATE TABLE IF NOT EXISTS enum_maps.ResetSeqNumFlag (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ResetSeqNumFlag VALUES 
('N', 'ResetSeqNumFlag_NO'),('Y', 'ResetSeqNumFlag_YES');

CREATE TABLE IF NOT EXISTS enum_maps.RespondentType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.RespondentType VALUES 
('1', 'RespondentType_ALL_MARKET_PARTICIPANTS'),('2', 'RespondentType_SPECIFIED_MARKET_PARTICIPANTS'),('3', 'RespondentType_ALL_MARKET_MAKERS'),('4', 'RespondentType_PRIMARY_MARKET_MAKER');

CREATE TABLE IF NOT EXISTS enum_maps.ResponseTransportType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ResponseTransportType VALUES 
('0', 'ResponseTransportType_INBAND'),('1', 'ResponseTransportType_OUT_OF_BAND');

CREATE TABLE IF NOT EXISTS enum_maps.RestructuringType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.RestructuringType VALUES 
('FR', 'RestructuringType_FULL_RESTRUCTURING'),('MM', 'RestructuringType_MODIFIED_MOD_RESTRUCTURING'),('MR', 'RestructuringType_MODIFIED_RESTRUCTURING'),('XR', 'RestructuringType_NO_RESTRUCTURING_SPECIFIED');

CREATE TABLE IF NOT EXISTS enum_maps.RiskInstrumentOperator (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.RiskInstrumentOperator VALUES 
('1', 'RiskInstrumentOperator_INCLUDE'),('2', 'RiskInstrumentOperator_EXCLUDE');

CREATE TABLE IF NOT EXISTS enum_maps.RiskLimitType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.RiskLimitType VALUES 
('1', 'RiskLimitType_GROSS_LIMIT'),('2', 'RiskLimitType_NET_LIMIT'),('3', 'RiskLimitType_EXPOSURE'),('4', 'RiskLimitType_LONG_LIMIT'),('5', 'RiskLimitType_SHORT_LIMIT');

CREATE TABLE IF NOT EXISTS enum_maps.RoundingDirection (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.RoundingDirection VALUES 
('0', 'RoundingDirection_ROUND_TO_NEAREST'),('1', 'RoundingDirection_ROUND_DOWN'),('2', 'RoundingDirection_ROUND_UP');

CREATE TABLE IF NOT EXISTS enum_maps.RoutingType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.RoutingType VALUES 
('1', 'RoutingType_TARGET_FIRM'),('2', 'RoutingType_TARGET_LIST'),('3', 'RoutingType_BLOCK_FIRM'),('4', 'RoutingType_BLOCK_LIST');

CREATE TABLE IF NOT EXISTS enum_maps.Rule80A (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.Rule80A VALUES 
('A', 'Rule80A_AGENCY_SINGLE_ORDER'),('B', 'Rule80A_SHORT_EXEMPT_TRANSACTION_B'),('C', 'Rule80A_PROPRIETARY_NON_ALGORITHMIC_PROGRAM_TRADE'),('D', 'Rule80A_PROGRAM_ORDER_INDEX_ARB_FOR_MEMBER_FIRM_ORG'),('E', 'Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_PRINCIPAL'),('F', 'Rule80A_SHORT_EXEMPT_TRANSACTION_F'),('H', 'Rule80A_SHORT_EXEMPT_TRANSACTION_H'),('I', 'Rule80A_INDIVIDUAL_INVESTOR_SINGLE_ORDER'),('J', 'Rule80A_PROPRIETARY_ALGORITHMIC_PROGRAM_TRADING'),('K', 'Rule80A_AGENCY_ALGORITHMIC_PROGRAM_TRADING'),('L', 'Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_MEMBER_COMPETING_MARKET_MAKER_AFFLIATED_WITH_THE_FIRM_CLEARING_THE_TRADE'),('M', 'Rule80A_PROGRAM_ORDER_INDEX_ARB_FOR_OTHER_MEMBER'),('N', 'Rule80A_AGENT_FOR_OTHER_MEMBER_NON_ALGORITHMIC_PROGRAM_TRADE'),('O', 'Rule80A_PROPRIETARY_TRANSACTIONS_FOR_COMPETING_MARKET_MAKER_THAT_IS_AFFILIATED_WITH_THE_CLEARING_MEMBER'),('P', 'Rule80A_PRINCIPAL'),('R', 'Rule80A_TRANSACTIONS_FOR_THE_ACCOUNT_OF_A_NON_MEMBER_COMPTING_MARKET_MAKER'),('S', 'Rule80A_SPECIALIST_TRADES'),('T', 'Rule80A_TRANSACTIONS_FOR_THE_ACCOUNT_OF_AN_UNAFFILIATED_MEMBERS_COMPETING_MARKET_MAKER'),('U', 'Rule80A_AGENCY_INDEX_ARBITRAGE'),('W', 'Rule80A_ALL_OTHER_ORDERS_AS_AGENT_FOR_OTHER_MEMBER'),('X', 'Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_MEMBER_COMPETING_MARKET_MAKER_NOT_AFFILIATED_WITH_THE_FIRM_CLEARING_THE_TRADE'),('Y', 'Rule80A_AGENCY_NON_ALGORITHMIC_PROGRAM_TRADE'),('Z', 'Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_NON_MEMBER_COMPETING_MARKET_MAKER');

CREATE TABLE IF NOT EXISTS enum_maps.Scope (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.Scope VALUES 
('1', 'Scope_LOCAL_MARKET'),('2', 'Scope_NATIONAL'),('3', 'Scope_GLOBAL');

CREATE TABLE IF NOT EXISTS enum_maps.SecDefStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SecDefStatus VALUES 
('0', 'SecDefStatus_PENDING_APPROVAL'),('1', 'SecDefStatus_APPROVED'),('2', 'SecDefStatus_REJECTED'),('3', 'SecDefStatus_UNAUTHORIZED_REQUEST'),('4', 'SecDefStatus_INVALID_DEFINITION_REQUEST');

CREATE TABLE IF NOT EXISTS enum_maps.SecurityIDSource (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SecurityIDSource VALUES 
('1', 'SecurityIDSource_CUSIP'),('2', 'SecurityIDSource_SEDOL'),('3', 'SecurityIDSource_QUIK'),('4', 'SecurityIDSource_ISIN_NUMBER'),('5', 'SecurityIDSource_RIC_CODE'),('6', 'SecurityIDSource_ISO_CURRENCY_CODE'),('7', 'SecurityIDSource_ISO_COUNTRY_CODE'),('8', 'SecurityIDSource_EXCHANGE_SYMBOL'),('9', 'SecurityIDSource_CONSOLIDATED_TAPE_ASSOCIATION'),('A', 'SecurityIDSource_BLOOMBERG_SYMBOL'),('B', 'SecurityIDSource_WERTPAPIER'),('C', 'SecurityIDSource_DUTCH'),('D', 'SecurityIDSource_VALOREN'),('E', 'SecurityIDSource_SICOVAM'),('F', 'SecurityIDSource_BELGIAN'),('G', 'SecurityIDSource_COMMON'),('H', 'SecurityIDSource_CLEARING_HOUSE'),('I', 'SecurityIDSource_ISDA_FPML_PRODUCT_SPECIFICATION'),('J', 'SecurityIDSource_OPTION_PRICE_REPORTING_AUTHORITY'),('K', 'SecurityIDSource_ISDA_FPML_PRODUCT_URL'),('L', 'SecurityIDSource_LETTER_OF_CREDIT'),('M', 'SecurityIDSource_MARKETPLACE_ASSIGNED_IDENTIFIER');

CREATE TABLE IF NOT EXISTS enum_maps.SecurityListRequestType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SecurityListRequestType VALUES 
('0', 'SecurityListRequestType_SYMBOL'),('1', 'SecurityListRequestType_SECURITYTYPE_AND_OR_CFICODE'),('2', 'SecurityListRequestType_PRODUCT'),('3', 'SecurityListRequestType_TRADINGSESSIONID'),('4', 'SecurityListRequestType_ALL_SECURITIES'),('5', 'SecurityListRequestType_MARKETID_OR_MARKETID_PLUS_MARKETSEGMENTID');

CREATE TABLE IF NOT EXISTS enum_maps.SecurityListType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SecurityListType VALUES 
('1', 'SecurityListType_INDUSTRY_CLASSIFICATION'),('2', 'SecurityListType_TRADING_LIST'),('3', 'SecurityListType_MARKET'),('4', 'SecurityListType_NEWSPAPER_LIST');

CREATE TABLE IF NOT EXISTS enum_maps.SecurityListTypeSource (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SecurityListTypeSource VALUES 
('1', 'SecurityListTypeSource_ICB'),('2', 'SecurityListTypeSource_NAICS'),('3', 'SecurityListTypeSource_GICS');

CREATE TABLE IF NOT EXISTS enum_maps.SecurityRequestResult (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SecurityRequestResult VALUES 
('0', 'SecurityRequestResult_VALID_REQUEST'),('1', 'SecurityRequestResult_INVALID_OR_UNSUPPORTED_REQUEST'),('2', 'SecurityRequestResult_NO_INSTRUMENTS_FOUND_THAT_MATCH_SELECTION_CRITERIA'),('3', 'SecurityRequestResult_NOT_AUTHORIZED_TO_RETRIEVE_INSTRUMENT_DATA'),('4', 'SecurityRequestResult_INSTRUMENT_DATA_TEMPORARILY_UNAVAILABLE'),('5', 'SecurityRequestResult_REQUEST_FOR_INSTRUMENT_DATA_NOT_SUPPORTED');

CREATE TABLE IF NOT EXISTS enum_maps.SecurityRequestType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SecurityRequestType VALUES 
('0', 'SecurityRequestType_REQUEST_SECURITY_IDENTITY_AND_SPECIFICATIONS'),('1', 'SecurityRequestType_REQUEST_SECURITY_IDENTITY_FOR_THE_SPECIFICATIONS_PROVIDED'),('2', 'SecurityRequestType_REQUEST_LIST_SECURITY_TYPES'),('3', 'SecurityRequestType_REQUEST_LIST_SECURITIES'),('4', 'SecurityRequestType_SYMBOL'),('5', 'SecurityRequestType_SECURITYTYPE_AND_OR_CFICODE'),('6', 'SecurityRequestType_PRODUCT'),('7', 'SecurityRequestType_TRADINGSESSIONID'),('8', 'SecurityRequestType_ALL_SECURITIES'),('9', 'SecurityRequestType_MARKETID_OR_MARKETID_PLUS_MARKETSEGMENTID');

CREATE TABLE IF NOT EXISTS enum_maps.SecurityResponseType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SecurityResponseType VALUES 
('1', 'SecurityResponseType_ACCEPT_SECURITY_PROPOSAL_AS_IS'),('2', 'SecurityResponseType_ACCEPT_SECURITY_PROPOSAL_WITH_REVISIONS_AS_INDICATED_IN_THE_MESSAGE'),('3', 'SecurityResponseType_LIST_OF_SECURITY_TYPES_RETURNED_PER_REQUEST'),('4', 'SecurityResponseType_LIST_OF_SECURITIES_RETURNED_PER_REQUEST'),('5', 'SecurityResponseType_REJECT_SECURITY_PROPOSAL'),('6', 'SecurityResponseType_CANNOT_MATCH_SELECTION_CRITERIA');

CREATE TABLE IF NOT EXISTS enum_maps.SecurityStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SecurityStatus VALUES 
('1', 'SecurityStatus_ACTIVE'),('2', 'SecurityStatus_INACTIVE');

CREATE TABLE IF NOT EXISTS enum_maps.SecurityTradingEvent (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SecurityTradingEvent VALUES 
('1', 'SecurityTradingEvent_ORDER_IMBALANCE_AUCTION_IS_EXTENDED'),('2', 'SecurityTradingEvent_TRADING_RESUMES'),('3', 'SecurityTradingEvent_PRICE_VOLATILITY_INTERRUPTION'),('4', 'SecurityTradingEvent_CHANGE_OF_TRADING_SESSION'),('5', 'SecurityTradingEvent_CHANGE_OF_TRADING_SUBSESSION'),('6', 'SecurityTradingEvent_CHANGE_OF_SECURITY_TRADING_STATUS'),('7', 'SecurityTradingEvent_CHANGE_OF_BOOK_TYPE'),('8', 'SecurityTradingEvent_CHANGE_OF_MARKET_DEPTH');

CREATE TABLE IF NOT EXISTS enum_maps.SecurityTradingStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SecurityTradingStatus VALUES 
('1', 'SecurityTradingStatus_OPENING_DELAY'),('10', 'SecurityTradingStatus_MARKET_ON_CLOSE_IMBALANCE_SELL'),('11', 'SecurityTradingStatus_11'),('12', 'SecurityTradingStatus_NO_MARKET_IMBALANCE'),('13', 'SecurityTradingStatus_NO_MARKET_ON_CLOSE_IMBALANCE'),('14', 'SecurityTradingStatus_ITS_PRE_OPENING'),('15', 'SecurityTradingStatus_NEW_PRICE_INDICATION'),('16', 'SecurityTradingStatus_TRADE_DISSEMINATION_TIME'),('17', 'SecurityTradingStatus_READY_TO_TRADE'),('18', 'SecurityTradingStatus_NOT_AVAILABLE_FOR_TRADING'),('19', 'SecurityTradingStatus_NOT_TRADED_ON_THIS_MARKET'),('2', 'SecurityTradingStatus_TRADING_HALT'),('20', 'SecurityTradingStatus_UNKNOWN_OR_INVALID'),('21', 'SecurityTradingStatus_PRE_OPEN'),('22', 'SecurityTradingStatus_OPENING_ROTATION'),('23', 'SecurityTradingStatus_FAST_MARKET'),('24', 'SecurityTradingStatus_PRE_CROSS'),('25', 'SecurityTradingStatus_CROSS'),('26', 'SecurityTradingStatus_POST_CLOSE'),('3', 'SecurityTradingStatus_RESUME'),('4', 'SecurityTradingStatus_NO_OPEN'),('5', 'SecurityTradingStatus_PRICE_INDICATION'),('6', 'SecurityTradingStatus_TRADING_RANGE_INDICATION'),('7', 'SecurityTradingStatus_MARKET_IMBALANCE_BUY'),('8', 'SecurityTradingStatus_MARKET_IMBALANCE_SELL'),('9', 'SecurityTradingStatus_MARKET_ON_CLOSE_IMBALANCE_BUY');

CREATE TABLE IF NOT EXISTS enum_maps.SecurityType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SecurityType VALUES 
('?', 'SecurityType_WILDCARD_ENTRY_FOR_USE_ON_SECURITY_DEFINITION_REQUEST'),('ABS', 'SecurityType_ASSET_BACKED_SECURITIES'),('AMENDED', 'SecurityType_AMENDED_RESTATED'),('AN', 'SecurityType_OTHER_ANTICIPATION_NOTES'),('BA', 'SecurityType_BANKERS_ACCEPTANCE'),('BDN', 'SecurityType_BANK_DEPOSITORY_NOTE'),('BN', 'SecurityType_BANK_NOTES'),('BOX', 'SecurityType_BILL_OF_EXCHANGES'),('BRADY', 'SecurityType_BRADY_BOND'),('BRIDGE', 'SecurityType_BRIDGE_LOAN'),('BUYSELL', 'SecurityType_BUY_SELLBACK'),('CAMM', 'SecurityType_CANADIAN_MONEY_MARKETS'),('CAN', 'SecurityType_CANADIAN_TREASURY_NOTES'),('CASH', 'SecurityType_CASH'),('CB', 'SecurityType_CONVERTIBLE_BOND'),('CD', 'SecurityType_CERTIFICATE_OF_DEPOSIT'),('CDS', 'SecurityType_CREDIT_DEFAULT_SWAP'),('CL', 'SecurityType_CALL_LOANS'),('CMB', 'SecurityType_CANADIAN_MORTGAGE_BONDS'),('CMBS', 'SecurityType_CORP_MORTGAGE_BACKED_SECURITIES'),('CMO', 'SecurityType_COLLATERALIZED_MORTGAGE_OBLIGATION'),('COFO', 'SecurityType_CERTIFICATE_OF_OBLIGATION'),('COFP', 'SecurityType_CERTIFICATE_OF_PARTICIPATION'),('CORP', 'SecurityType_CORPORATE_BOND'),('CP', 'SecurityType_COMMERCIAL_PAPER'),('CPP', 'SecurityType_CORPORATE_PRIVATE_PLACEMENT'),('CS', 'SecurityType_COMMON_STOCK'),('CTB', 'SecurityType_CANADIAN_TREASURY_BILLS'),('DEFLTED', 'SecurityType_DEFAULTED'),('DINP', 'SecurityType_DEBTOR_IN_POSSESSION'),('DN', 'SecurityType_DEPOSIT_NOTES'),('DUAL', 'SecurityType_DUAL_CURRENCY'),('EUCD', 'SecurityType_EURO_CERTIFICATE_OF_DEPOSIT'),('EUCORP', 'SecurityType_EURO_CORPORATE_BOND'),('EUCP', 'SecurityType_EURO_COMMERCIAL_PAPER'),('EUFRN', 'SecurityType_EURO_CORPORATE_FLOATING_RATE_NOTES'),('EUSOV', 'SecurityType_EURO_SOVEREIGNS'),('EUSUPRA', 'SecurityType_EURO_SUPRANATIONAL_COUPONS'),('FAC', 'SecurityType_FEDERAL_AGENCY_COUPON'),('FADN', 'SecurityType_FEDERAL_AGENCY_DISCOUNT_NOTE'),('FHA', 'SecurityType_FEDERAL_HOUSING_AUTHORITY'),('FHL', 'SecurityType_FEDERAL_HOME_LOAN'),('FN', 'SecurityType_FEDERAL_NATIONAL_MORTGAGE_ASSOCIATION'),('FOR', 'SecurityType_FOREIGN_EXCHANGE_CONTRACT'),('FORWARD', 'SecurityType_FORWARD'),('FRN', 'SecurityType_US_CORPORATE_FLOATING_RATE_NOTES'),('FUT', 'SecurityType_FUTURE'),('FXFWD', 'SecurityType_FX_FORWARD'),('FXNDF', 'SecurityType_NON_DELIVERABLE_FORWARD'),('FXSPOT', 'SecurityType_FX_SPOT'),('FXSWAP', 'SecurityType_FX_SWAP'),('GN', 'SecurityType_GOVERNMENT_NATIONAL_MORTGAGE_ASSOCIATION'),('GO', 'SecurityType_GENERAL_OBLIGATION_BONDS'),('GOVT', 'SecurityType_TREASURIES_PLUS_AGENCY_DEBENTURE'),('IET', 'SecurityType_IOETTE_MORTGAGE'),('IRS', 'SecurityType_INTEREST_RATE_SWAP'),('LOFC', 'SecurityType_LETTER_OF_CREDIT'),('LQN', 'SecurityType_LIQUIDITY_NOTE'),('MATURED', 'SecurityType_MATURED'),('MBS', 'SecurityType_MORTGAGE_BACKED_SECURITIES'),('MF', 'SecurityType_MUTUAL_FUND'),('MIO', 'SecurityType_MORTGAGE_INTEREST_ONLY'),('MLEG', 'SecurityType_MULTILEG_INSTRUMENT'),('MPO', 'SecurityType_MORTGAGE_PRINCIPAL_ONLY'),('MPP', 'SecurityType_MORTGAGE_PRIVATE_PLACEMENT'),('MPT', 'SecurityType_MISCELLANEOUS_PASS_THROUGH'),('MT', 'SecurityType_MANDATORY_TENDER'),('MTN', 'SecurityType_MEDIUM_TERM_NOTES'),('MUNI', 'SecurityType_MUNICIPAL_BOND'),('NONE', 'SecurityType_NO_SECURITY_TYPE'),('ONITE', 'SecurityType_OVERNIGHT'),('OOC', 'SecurityType_OPTIONS_ON_COMBO'),('OOF', 'SecurityType_OPTIONS_ON_FUTURES'),('OOP', 'SecurityType_OPTIONS_ON_PHYSICAL'),('OPT', 'SecurityType_OPTION'),('PEF', 'SecurityType_PRIVATE_EXPORT_FUNDING'),('PFAND', 'SecurityType_PFANDBRIEFE'),('PN', 'SecurityType_PROMISSORY_NOTE'),('POOL', 'SecurityType_AGENCY_POOLS'),('PROV', 'SecurityType_CANADIAN_PROVINCIAL_BONDS'),('PS', 'SecurityType_PREFERRED_STOCK'),('PZFJ', 'SecurityType_PLAZOS_FIJOS'),('RAN', 'SecurityType_REVENUE_ANTICIPATION_NOTE'),('REPLACD', 'SecurityType_REPLACED'),('REPO', 'SecurityType_REPURCHASE'),('RETIRED', 'SecurityType_RETIRED'),('REV', 'SecurityType_REVENUE_BONDS'),('RP', 'SecurityType_REPURCHASE_AGREEMENT'),('RVLV', 'SecurityType_REVOLVER_LOAN'),('RVLVTRM', 'SecurityType_REVOLVER_TERM_LOAN'),('RVRP', 'SecurityType_REVERSE_REPURCHASE_AGREEMENT'),('SECLOAN', 'SecurityType_SECURITIES_LOAN'),('SECPLEDGE', 'SecurityType_SECURITIES_PLEDGE'),('SL', 'SecurityType_STUDENT_LOAN_MARKETING_ASSOCIATION'),('SLQN', 'SecurityType_SECURED_LIQUIDITY_NOTE'),('SPCLA', 'SecurityType_SPECIAL_ASSESSMENT'),('SPCLO', 'SecurityType_SPECIAL_OBLIGATION'),('SPCLT', 'SecurityType_SPECIAL_TAX'),('STN', 'SecurityType_SHORT_TERM_LOAN_NOTE'),('STRUCT', 'SecurityType_STRUCTURED_NOTES'),('SUPRA', 'SecurityType_USD_SUPRANATIONAL_COUPONS'),('SWING', 'SecurityType_SWING_LINE_FACILITY'),('TAN', 'SecurityType_TAX_ANTICIPATION_NOTE'),('TAXA', 'SecurityType_TAX_ALLOCATION'),('TB', 'SecurityType_TREASURY_BILL'),('TBA', 'SecurityType_TO_BE_ANNOUNCED'),('TBILL', 'SecurityType_US_TREASURY_BILL_TBILL'),('TBOND', 'SecurityType_US_TREASURY_BOND'),('TCAL', 'SecurityType_PRINCIPAL_STRIP_OF_A_CALLABLE_BOND_OR_NOTE'),('TD', 'SecurityType_TIME_DEPOSIT'),('TECP', 'SecurityType_TAX_EXEMPT_COMMERCIAL_PAPER'),('TERM', 'SecurityType_TERM_LOAN'),('TINT', 'SecurityType_INTEREST_STRIP_FROM_ANY_BOND_OR_NOTE'),('TIPS', 'SecurityType_TREASURY_INFLATION_PROTECTED_SECURITIES'),('TLQN', 'SecurityType_TERM_LIQUIDITY_NOTE'),('TMCP', 'SecurityType_TAXABLE_MUNICIPAL_CP'),('TNOTE', 'SecurityType_US_TREASURY_NOTE_TNOTE'),('TPRN', 'SecurityType_PRINCIPAL_STRIP_FROM_A_NON_CALLABLE_BOND_OR_NOTE'),('TRAN', 'SecurityType_TAX_REVENUE_ANTICIPATION_NOTE'),('UST', 'SecurityType_US_TREASURY_NOTE_UST'),('USTB', 'SecurityType_US_TREASURY_BILL_USTB'),('VRDN', 'SecurityType_VARIABLE_RATE_DEMAND_NOTE'),('WAR', 'SecurityType_WARRANT'),('WITHDRN', 'SecurityType_WITHDRAWN'),('WLD', 'SecurityType_WILDCARD_ENTRY'),('XCN', 'SecurityType_EXTENDED_COMM_NOTE'),('XLINKD', 'SecurityType_INDEXED_LINKED'),('YANK', 'SecurityType_YANKEE_CORPORATE_BOND'),('YCD', 'SecurityType_YANKEE_CERTIFICATE_OF_DEPOSIT'),('ZOO', 'SecurityType_CATS_TIGERS_LIONS');

CREATE TABLE IF NOT EXISTS enum_maps.SecurityUpdateAction (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SecurityUpdateAction VALUES 
('A', 'SecurityUpdateAction_ADD'),('D', 'SecurityUpdateAction_DELETE'),('M', 'SecurityUpdateAction_MODIFY');

CREATE TABLE IF NOT EXISTS enum_maps.Seniority (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.Seniority VALUES 
('SB', 'Seniority_SUBORDINATED'),('SD', 'Seniority_SENIOR_SECURED'),('SR', 'Seniority_SENIOR');

CREATE TABLE IF NOT EXISTS enum_maps.SessionRejectReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SessionRejectReason VALUES 
('0', 'SessionRejectReason_INVALID_TAG_NUMBER'),('1', 'SessionRejectReason_REQUIRED_TAG_MISSING'),('10', 'SessionRejectReason_SENDINGTIME_ACCURACY_PROBLEM'),('11', 'SessionRejectReason_INVALID_MSGTYPE'),('12', 'SessionRejectReason_XML_VALIDATION_ERROR'),('13', 'SessionRejectReason_TAG_APPEARS_MORE_THAN_ONCE'),('14', 'SessionRejectReason_TAG_SPECIFIED_OUT_OF_REQUIRED_ORDER'),('15', 'SessionRejectReason_REPEATING_GROUP_FIELDS_OUT_OF_ORDER'),('16', 'SessionRejectReason_INCORRECT_NUMINGROUP_COUNT_FOR_REPEATING_GROUP'),('17', 'SessionRejectReason_NON_DATA_VALUE_INCLUDES_FIELD_DELIMITER'),('18', 'SessionRejectReason_INVALID_UNSUPPORTED_APPLICATION_VERSION'),('2', 'SessionRejectReason_TAG_NOT_DEFINED_FOR_THIS_MESSAGE_TYPE'),('3', 'SessionRejectReason_UNDEFINED_TAG'),('4', 'SessionRejectReason_TAG_SPECIFIED_WITHOUT_A_VALUE'),('5', 'SessionRejectReason_VALUE_IS_INCORRECT'),('6', 'SessionRejectReason_INCORRECT_DATA_FORMAT_FOR_VALUE'),('7', 'SessionRejectReason_DECRYPTION_PROBLEM'),('8', 'SessionRejectReason_SIGNATURE_PROBLEM'),('9', 'SessionRejectReason_COMPID_PROBLEM'),('99', 'SessionRejectReason_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.SessionStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SessionStatus VALUES 
('0', 'SessionStatus_SESSION_ACTIVE'),('1', 'SessionStatus_SESSION_PASSWORD_CHANGED'),('2', 'SessionStatus_SESSION_PASSWORD_DUE_TO_EXPIRE'),('3', 'SessionStatus_NEW_SESSION_PASSWORD_DOES_NOT_COMPLY_WITH_POLICY'),('4', 'SessionStatus_SESSION_LOGOUT_COMPLETE'),('5', 'SessionStatus_INVALID_USERNAME_OR_PASSWORD'),('6', 'SessionStatus_ACCOUNT_LOCKED'),('7', 'SessionStatus_LOGONS_ARE_NOT_ALLOWED_AT_THIS_TIME'),('8', 'SessionStatus_PASSWORD_EXPIRED');

CREATE TABLE IF NOT EXISTS enum_maps.SettlCurrFxRateCalc (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SettlCurrFxRateCalc VALUES 
('D', 'SettlCurrFxRateCalc_DIVIDE'),('M', 'SettlCurrFxRateCalc_MULTIPLY');

CREATE TABLE IF NOT EXISTS enum_maps.SettlDeliveryType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SettlDeliveryType VALUES 
('0', 'SettlDeliveryType_VERSUS_PAYMENT_DELIVER'),('1', 'SettlDeliveryType_FREE_DELIVER'),('2', 'SettlDeliveryType_TRI_PARTY'),('3', 'SettlDeliveryType_HOLD_IN_CUSTODY');

CREATE TABLE IF NOT EXISTS enum_maps.SettlInstMode (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SettlInstMode VALUES 
('0', 'SettlInstMode_DEFAULT'),('1', 'SettlInstMode_STANDING_INSTRUCTIONS_PROVIDED'),('2', 'SettlInstMode_SPECIFIC_ALLOCATION_ACCOUNT_OVERRIDING'),('3', 'SettlInstMode_SPECIFIC_ALLOCATION_ACCOUNT_STANDING'),('4', 'SettlInstMode_SPECIFIC_ORDER_FOR_A_SINGLE_ACCOUNT'),('5', 'SettlInstMode_REQUEST_REJECT');

CREATE TABLE IF NOT EXISTS enum_maps.SettlInstReqRejCode (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SettlInstReqRejCode VALUES 
('0', 'SettlInstReqRejCode_UNABLE_TO_PROCESS_REQUEST'),('1', 'SettlInstReqRejCode_UNKNOWN_ACCOUNT'),('2', 'SettlInstReqRejCode_NO_MATCHING_SETTLEMENT_INSTRUCTIONS_FOUND'),('99', 'SettlInstReqRejCode_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.SettlInstSource (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SettlInstSource VALUES 
('1', 'SettlInstSource_BROKERS_INSTRUCTIONS'),('2', 'SettlInstSource_INSTITUTIONS_INSTRUCTIONS'),('3', 'SettlInstSource_INVESTOR');

CREATE TABLE IF NOT EXISTS enum_maps.SettlInstTransType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SettlInstTransType VALUES 
('C', 'SettlInstTransType_CANCEL'),('N', 'SettlInstTransType_NEW'),('R', 'SettlInstTransType_REPLACE'),('T', 'SettlInstTransType_RESTATE');

CREATE TABLE IF NOT EXISTS enum_maps.SettlLocation (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SettlLocation VALUES 
('CED', 'SettlLocation_CEDEL'),('DTC', 'SettlLocation_DEPOSITORY_TRUST_COMPANY'),('EUR', 'SettlLocation_EURO_CLEAR'),('FED', 'SettlLocation_FEDERAL_BOOK_ENTRY'),('ISO_Country_Code', 'SettlLocation_LOCAL_MARKET_SETTLE_LOCATION'),('PNY', 'SettlLocation_PHYSICAL'),('PTC', 'SettlLocation_PARTICIPANT_TRUST_COMPANY');

CREATE TABLE IF NOT EXISTS enum_maps.SettlMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SettlMethod VALUES 
('C', 'SettlMethod_CASH_SETTLEMENT_REQUIRED'),('P', 'SettlMethod_PHYSICAL_SETTLEMENT_REQUIRED');

CREATE TABLE IF NOT EXISTS enum_maps.SettlObligMode (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SettlObligMode VALUES 
('1', 'SettlObligMode_PRELIMINARY'),('2', 'SettlObligMode_FINAL');

CREATE TABLE IF NOT EXISTS enum_maps.SettlObligSource (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SettlObligSource VALUES 
('1', 'SettlObligSource_INSTRUCTIONS_OF_BROKER'),('2', 'SettlObligSource_INSTRUCTIONS_FOR_INSTITUTION'),('3', 'SettlObligSource_INVESTOR');

CREATE TABLE IF NOT EXISTS enum_maps.SettlObligTransType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SettlObligTransType VALUES 
('C', 'SettlObligTransType_CANCEL'),('N', 'SettlObligTransType_NEW'),('R', 'SettlObligTransType_REPLACE'),('T', 'SettlObligTransType_RESTATE');

CREATE TABLE IF NOT EXISTS enum_maps.SettlPriceType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SettlPriceType VALUES 
('1', 'SettlPriceType_FINAL'),('2', 'SettlPriceType_THEORETICAL');

CREATE TABLE IF NOT EXISTS enum_maps.SettlSessID (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SettlSessID VALUES 
('EOD', 'SettlSessID_END_OF_DAY'),('ETH', 'SettlSessID_ELECTRONIC_TRADING_HOURS'),('ITD', 'SettlSessID_INTRADAY'),('RTH', 'SettlSessID_REGULAR_TRADING_HOURS');

CREATE TABLE IF NOT EXISTS enum_maps.SettlType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SettlType VALUES 
('0', 'SettlType_REGULAR'),('1', 'SettlType_CASH'),('2', 'SettlType_NEXT_DAY'),('3', 'SettlType_T_PLUS_2'),('4', 'SettlType_T_PLUS_3'),('5', 'SettlType_T_PLUS_4'),('6', 'SettlType_FUTURE'),('7', 'SettlType_WHEN_AND_IF_ISSUED'),('8', 'SettlType_SELLERS_OPTION'),('9', 'SettlType_T_PLUS_5'),('B', 'SettlType_BROKEN_DATE'),('C', 'SettlType_FX_SPOT_NEXT_SETTLEMENT');

CREATE TABLE IF NOT EXISTS enum_maps.SettlmntTyp (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SettlmntTyp VALUES 
('0', 'SettlmntTyp_REGULAR'),('1', 'SettlmntTyp_CASH'),('2', 'SettlmntTyp_NEXT_DAY'),('3', 'SettlmntTyp_T_PLUS_2'),('4', 'SettlmntTyp_T_PLUS_3'),('5', 'SettlmntTyp_T_PLUS_4'),('6', 'SettlmntTyp_FUTURE'),('7', 'SettlmntTyp_WHEN_AND_IF_ISSUED'),('8', 'SettlmntTyp_SELLERS_OPTION'),('9', 'SettlmntTyp_T_PLUS_5'),('A', 'SettlmntTyp_T_PLUS_1');

CREATE TABLE IF NOT EXISTS enum_maps.ShortSaleReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ShortSaleReason VALUES 
('0', 'ShortSaleReason_DEALER_SOLD_SHORT'),('1', 'ShortSaleReason_DEALER_SOLD_SHORT_EXEMPT'),('2', 'ShortSaleReason_SELLING_CUSTOMER_SOLD_SHORT'),('3', 'ShortSaleReason_SELLING_CUSTOMER_SOLD_SHORT_EXEMPT'),('4', 'ShortSaleReason_QUALIFIED_SERVICE_REPRESENTATIVE'),('5', 'ShortSaleReason_QSR_OR_AGU_CONTRA_SIDE_SOLD_SHORT_EXEMPT');

CREATE TABLE IF NOT EXISTS enum_maps.Side (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.Side VALUES 
('1', 'Side_BUY'),('2', 'Side_SELL'),('3', 'Side_BUY_MINUS'),('4', 'Side_SELL_PLUS'),('5', 'Side_SELL_SHORT'),('6', 'Side_SELL_SHORT_EXEMPT'),('7', 'Side_UNDISCLOSED'),('8', 'Side_CROSS'),('9', 'Side_CROSS_SHORT'),('A', 'Side_CROSS_SHORT_EXEMPT'),('B', 'Side_AS_DEFINED'),('C', 'Side_OPPOSITE'),('D', 'Side_SUBSCRIBE'),('E', 'Side_REDEEM'),('F', 'Side_LEND'),('G', 'Side_BORROW');

CREATE TABLE IF NOT EXISTS enum_maps.SideMultiLegReportingType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SideMultiLegReportingType VALUES 
('1', 'SideMultiLegReportingType_SINGLE_SECURITY'),('2', 'SideMultiLegReportingType_INDIVIDUAL_LEG_OF_A_MULTILEG_SECURITY'),('3', 'SideMultiLegReportingType_MULTILEG_SECURITY');

CREATE TABLE IF NOT EXISTS enum_maps.SideTrdSubTyp (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SideTrdSubTyp VALUES 
('0', 'SideTrdSubTyp_CMTA'),('1', 'SideTrdSubTyp_INTERNAL_TRANSFER'),('10', 'SideTrdSubTyp_TRANSACTION_FROM_ASSIGNMENT'),('2', 'SideTrdSubTyp_EXTERNAL_TRANSFER'),('3', 'SideTrdSubTyp_REJECT_FOR_SUBMITTING_TRADE'),('4', 'SideTrdSubTyp_ADVISORY_FOR_CONTRA_SIDE'),('5', 'SideTrdSubTyp_OFFSET_DUE_TO_AN_ALLOCATION'),('6', 'SideTrdSubTyp_ONSET_DUE_TO_AN_ALLOCATION'),('7', 'SideTrdSubTyp_DIFFERENTIAL_SPREAD'),('8', 'SideTrdSubTyp_IMPLIED_SPREAD_LEG_EXECUTED_AGAINST_AN_OUTRIGHT'),('9', 'SideTrdSubTyp_TRANSACTION_FROM_EXERCISE');

CREATE TABLE IF NOT EXISTS enum_maps.SideValueInd (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SideValueInd VALUES 
('1', 'SideValueInd_SIDE_VALUE_1'),('2', 'SideValueInd_SIDE_VALUE_2');

CREATE TABLE IF NOT EXISTS enum_maps.SolicitedFlag (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SolicitedFlag VALUES 
('N', 'SolicitedFlag_NO'),('Y', 'SolicitedFlag_YES');

CREATE TABLE IF NOT EXISTS enum_maps.StandInstDbType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.StandInstDbType VALUES 
('0', 'StandInstDbType_OTHER'),('1', 'StandInstDbType_DTC_SID'),('2', 'StandInstDbType_THOMSON_ALERT'),('3', 'StandInstDbType_A_GLOBAL_CUSTODIAN'),('4', 'StandInstDbType_ACCOUNTNET');

CREATE TABLE IF NOT EXISTS enum_maps.StatsType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.StatsType VALUES 
('1', 'StatsType_EXCHANGE_LAST'),('2', 'StatsType_HIGH'),('3', 'StatsType_AVERAGE_PRICE'),('4', 'StatsType_TURNOVER');

CREATE TABLE IF NOT EXISTS enum_maps.StatusValue (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.StatusValue VALUES 
('1', 'StatusValue_CONNECTED'),('2', 'StatusValue_NOT_CONNECTED_2'),('3', 'StatusValue_NOT_CONNECTED_3'),('4', 'StatusValue_IN_PROCESS');

CREATE TABLE IF NOT EXISTS enum_maps.StipulationType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.StipulationType VALUES 
('ABS', 'StipulationType_ABSOLUTE_PREPAYMENT_SPEED'),('AMT', 'StipulationType_ALTERNATIVE_MINIMUM_TAX'),('AUTOREINV', 'StipulationType_AUTO_REINVESTMENT_AT_RATE_OR_BETTER'),('AVAILQTY', 'StipulationType_AVAILABLE_OFFER_QUANTITY_TO_BE_SHOWN_TO_THE_STREET'),('AVFICO', 'StipulationType_AVERAGE_FICO_SCORE'),('AVSIZE', 'StipulationType_AVERAGE_LOAN_SIZE'),('BANKQUAL', 'StipulationType_BANK_QUALIFIED'),('BGNCON', 'StipulationType_BARGAIN_CONDITIONS'),('BROKERCREDIT', 'StipulationType_BROKERS_SALES_CREDIT'),('COUPON', 'StipulationType_COUPON_RANGE'),('CPP', 'StipulationType_CONSTANT_PREPAYMENT_PENALTY'),('CPR', 'StipulationType_CONSTANT_PREPAYMENT_RATE'),('CPY', 'StipulationType_CONSTANT_PREPAYMENT_YIELD'),('CURRENCY', 'StipulationType_ISO_CURRENCY_CODE'),('CUSTOMDATE', 'StipulationType_CUSTOM_START_END_DATE'),('DISCOUNT', 'StipulationType_DISCOUNT_RATE'),('GEOG', 'StipulationType_GEOGRAPHICS_AND_RANGE'),('HAIRCUT', 'StipulationType_VALUATION_DISCOUNT'),('HEP', 'StipulationType_FINAL_CPR_OF_HOME_EQUITY_PREPAYMENT_CURVE'),('INSURED', 'StipulationType_INSURED'),('INTERNALPX', 'StipulationType_OFFER_PRICE_TO_BE_SHOWN_TO_INTERNAL_BROKERS'),('INTERNALQTY', 'StipulationType_OFFER_QUANTITY_TO_BE_SHOWN_TO_INTERNAL_BROKERS'),('ISSUE', 'StipulationType_YEAR_OR_YEAR_MONTH_OF_ISSUE'),('ISSUER', 'StipulationType_ISSUERS_TICKER'),('ISSUESIZE', 'StipulationType_ISSUE_SIZE_RANGE'),('LEAVEQTY', 'StipulationType_THE_MINIMUM_RESIDUAL_OFFER_QUANTITY'),('LOOKBACK', 'StipulationType_LOOKBACK_DAYS'),('LOT', 'StipulationType_EXPLICIT_LOT_IDENTIFIER'),('LOTVAR', 'StipulationType_LOT_VARIANCE'),('MAT', 'StipulationType_MATURITY_YEAR_AND_MONTH'),('MATURITY', 'StipulationType_MATURITY_RANGE'),('MAXBAL', 'StipulationType_MAXIMUM_LOAN_BALANCE'),('MAXDNOM', 'StipulationType_MAXIMUMDENOMINATION'),('MAXORDQTY', 'StipulationType_MAXIMUM_ORDER_SIZE'),('MAXSUBS', 'StipulationType_MAXIMUM_SUBSTITUTIONS'),('MHP', 'StipulationType_PERCENT_OF_MANUFACTURED_HOUSING_PREPAYMENT_CURVE'),('MINDNOM', 'StipulationType_MINIMUM_DENOMINATION'),('MININCR', 'StipulationType_MINIMUM_INCREMENT'),('MINQTY', 'StipulationType_MINIMUM_QUANTITY'),('MPR', 'StipulationType_MONTHLY_PREPAYMENT_RATE'),('ORDRINCR', 'StipulationType_ORDER_QUANTITY_INCREMENT'),('PAYFREQ', 'StipulationType_PAYMENT_FREQUENCY_CALENDAR'),('PIECES', 'StipulationType_NUMBER_OF_PIECES'),('PMAX', 'StipulationType_POOLS_MAXIMUM'),('PMIN', 'StipulationType_POOLSMINIMUM'),('POOL', 'StipulationType_POOL_IDENTIFIER'),('PPC', 'StipulationType_PERCENT_OF_PROSPECTUS_PREPAYMENT_CURVE'),('PPL', 'StipulationType_POOLS_PER_LOT'),('PPM', 'StipulationType_POOLS_PER_MILLION'),('PPT', 'StipulationType_POOLS_PER_TRADE'),('PRICE', 'StipulationType_PRICE_RANGE'),('PRICEFREQ', 'StipulationType_PRICING_FREQUENCY'),('PRIMARY', 'StipulationType_PRIMARY_OR_SECONDARY_MARKET_INDICATOR'),('PROD', 'StipulationType_PRODUCTION_YEAR'),('PROTECT', 'StipulationType_CALL_PROTECTION'),('PSA', 'StipulationType_PERCENT_OF_BMA_PREPAYMENT_CURVE'),('PURPOSE', 'StipulationType_PURPOSE'),('PXSOURCE', 'StipulationType_BENCHMARK_PRICE_SOURCE'),('RATING', 'StipulationType_RATING_SOURCE_AND_RANGE'),('REDEMPTION', 'StipulationType_TYPE_OF_REDEMPTION'),('REFINT', 'StipulationType_INTEREST_OF_ROLLING_OR_CLOSING_TRADE'),('REFPRIN', 'StipulationType_PRINCIPAL_OF_ROLLING_OR_CLOSING_TRADE'),('REFTRADE', 'StipulationType_REFERENCE_TO_ROLLING_OR_CLOSING_TRADE'),('RESTRICTED', 'StipulationType_RESTRICTED'),('ROLLTYPE', 'StipulationType_TYPE_OF_ROLL_TRADE'),('SALESCREDITOVR', 'StipulationType_BROKER_SALES_CREDIT_OVERRIDE'),('SECTOR', 'StipulationType_MARKET_SECTOR'),('SECTYPE', 'StipulationType_SECURITY_TYPE_INCLUDED_OR_EXCLUDED'),('SMM', 'StipulationType_SINGLE_MONTHLY_MORTALITY'),('STRUCT', 'StipulationType_STRUCTURE'),('SUBSFREQ', 'StipulationType_SUBSTITUTIONS_FREQUENCY'),('SUBSLEFT', 'StipulationType_SUBSTITUTIONS_LEFT'),('TEXT', 'StipulationType_FREEFORM_TEXT'),('TRADERCREDIT', 'StipulationType_TRADERS_CREDIT'),('TRDVAR', 'StipulationType_TRADE_VARIANCE'),('WAC', 'StipulationType_WEIGHTED_AVERAGE_COUPON'),('WAL', 'StipulationType_WEIGHTED_AVERAGE_LIFE_COUPON'),('WALA', 'StipulationType_WEIGHTED_AVERAGE_LOAN_AGE'),('WAM', 'StipulationType_WEIGHTED_AVERAGE_MATURITY'),('WHOLE', 'StipulationType_WHOLE_POOL'),('YIELD', 'StipulationType_YIELD_RANGE'),('YTM', 'StipulationType_YIELD_TO_MATURITY');

CREATE TABLE IF NOT EXISTS enum_maps.StrategyParameterType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.StrategyParameterType VALUES 
('1', 'StrategyParameterType_INT'),('10', 'StrategyParameterType_AMT'),('11', 'StrategyParameterType_PERCENTAGE'),('12', 'StrategyParameterType_CHAR'),('13', 'StrategyParameterType_BOOLEAN'),('14', 'StrategyParameterType_STRING'),('15', 'StrategyParameterType_MULTIPLECHARVALUE'),('16', 'StrategyParameterType_CURRENCY'),('17', 'StrategyParameterType_EXCHANGE'),('18', 'StrategyParameterType_MONTHYEAR'),('19', 'StrategyParameterType_UTCTIMESTAMP'),('2', 'StrategyParameterType_LENGTH'),('20', 'StrategyParameterType_UTCTIMEONLY'),('21', 'StrategyParameterType_LOCALMKTDATE'),('22', 'StrategyParameterType_UTCDATEONLY'),('23', 'StrategyParameterType_DATA'),('24', 'StrategyParameterType_MULTIPLESTRINGVALUE'),('25', 'StrategyParameterType_COUNTRY'),('26', 'StrategyParameterType_LANGUAGE'),('27', 'StrategyParameterType_TZTIMEONLY'),('28', 'StrategyParameterType_TZTIMESTAMP'),('29', 'StrategyParameterType_TENOR'),('3', 'StrategyParameterType_NUMINGROUP'),('4', 'StrategyParameterType_SEQNUM'),('5', 'StrategyParameterType_TAGNUM'),('6', 'StrategyParameterType_FLOAT'),('7', 'StrategyParameterType_QTY'),('8', 'StrategyParameterType_PRICE'),('9', 'StrategyParameterType_PRICEOFFSET');

CREATE TABLE IF NOT EXISTS enum_maps.StreamAsgnAckType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.StreamAsgnAckType VALUES 
('0', 'StreamAsgnAckType_ASSIGNMENT_ACCEPTED'),('1', 'StreamAsgnAckType_ASSIGNMENT_REJECTED');

CREATE TABLE IF NOT EXISTS enum_maps.StreamAsgnRejReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.StreamAsgnRejReason VALUES 
('0', 'StreamAsgnRejReason_UNKNOWN_CLIENT'),('1', 'StreamAsgnRejReason_EXCEEDS_MAXIMUM_SIZE'),('2', 'StreamAsgnRejReason_UNKNOWN_OR_INVALID_CURRENCY_PAIR'),('3', 'StreamAsgnRejReason_NO_AVAILABLE_STREAM'),('99', 'StreamAsgnRejReason_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.StreamAsgnReqType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.StreamAsgnReqType VALUES 
('1', 'StreamAsgnReqType_STREAM_ASSIGNMENT_FOR_NEW_CUSTOMER'),('2', 'StreamAsgnReqType_STREAM_ASSIGNMENT_FOR_EXISTING_CUSTOMER');

CREATE TABLE IF NOT EXISTS enum_maps.StreamAsgnType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.StreamAsgnType VALUES 
('1', 'StreamAsgnType_ASSIGNMENT'),('2', 'StreamAsgnType_REJECTED'),('3', 'StreamAsgnType_TERMINATE_UNASSIGN');

CREATE TABLE IF NOT EXISTS enum_maps.StrikePriceBoundaryMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.StrikePriceBoundaryMethod VALUES 
('1', 'StrikePriceBoundaryMethod_LESS_THAN_UNDERLYING_PRICE_IS_IN_THE_MONEY'),('2', 'StrikePriceBoundaryMethod_LESS_THAN_OR_EQUAL_TO_THE_UNDERLYING_PRICE_IS_IN_THE_MONEY'),('3', 'StrikePriceBoundaryMethod_EQUAL_TO_THE_UNDERLYING_PRICE_IS_IN_THE_MONEY'),('4', 'StrikePriceBoundaryMethod_GREATER_THAN_OR_EQUAL_TO_UNDERLYING_PRICE_IS_IN_THE_MONEY'),('5', 'StrikePriceBoundaryMethod_GREATER_THAN_UNDERLYING_IS_IN_THE_MONEY');

CREATE TABLE IF NOT EXISTS enum_maps.StrikePriceDeterminationMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.StrikePriceDeterminationMethod VALUES 
('1', 'StrikePriceDeterminationMethod_FIXED_STRIKE'),('2', 'StrikePriceDeterminationMethod_STRIKE_SET_AT_EXPIRATION_TO_UNDERLYING_OR_OTHER_VALUE'),('3', 'StrikePriceDeterminationMethod_STRIKE_SET_TO_AVERAGE_OF_UNDERLYING_SETTLEMENT_PRICE_ACROSS_THE_LIFE_OF_THE_OPTION'),('4', 'StrikePriceDeterminationMethod_STRIKE_SET_TO_OPTIMAL_VALUE');

CREATE TABLE IF NOT EXISTS enum_maps.SubscriptionRequestType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SubscriptionRequestType VALUES 
('0', 'SubscriptionRequestType_SNAPSHOT'),('1', 'SubscriptionRequestType_SNAPSHOT_PLUS_UPDATES'),('2', 'SubscriptionRequestType_DISABLE_PREVIOUS_SNAPSHOT_PLUS_UPDATE_REQUEST');

CREATE TABLE IF NOT EXISTS enum_maps.SymbolSfx (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.SymbolSfx VALUES 
('CD', 'SymbolSfx_EUCP_WITH_LUMP_SUM_INTEREST_RATHER_THAN_DISCOUNT_PRICE'),('WI', 'SymbolSfx_WHEN_ISSUED_FOR_A_SECURITY_TO_BE_REISSUED_UNDER_AN_OLD_CUSIP_OR_ISIN');

CREATE TABLE IF NOT EXISTS enum_maps.TargetStrategy (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TargetStrategy VALUES 
('1', 'TargetStrategy_VWAP'),('1000', 'TargetStrategy_RESERVEDANDAVAILABLEFORBILATERALLYAGREEDUPONUSERDEFINEDVALUES'),('2', 'TargetStrategy_PARTICIPATE'),('3', 'TargetStrategy_MININIZE_MARKET_IMPACT');

CREATE TABLE IF NOT EXISTS enum_maps.TaxAdvantageType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TaxAdvantageType VALUES 
('0', 'TaxAdvantageType_NONE_NOT_APPLICABLE'),('1', 'TaxAdvantageType_MAXI_ISA'),('10', 'TaxAdvantageType_EMPLOYEE_10'),('11', 'TaxAdvantageType_EMPLOYER_11'),('12', 'TaxAdvantageType_EMPLOYER_12'),('13', 'TaxAdvantageType_NON_FUND_PROTOTYPE_IRA'),('14', 'TaxAdvantageType_NON_FUND_QUALIFIED_PLAN'),('15', 'TaxAdvantageType_DEFINED_CONTRIBUTION_PLAN'),('16', 'TaxAdvantageType_INDIVIDUAL_RETIREMENT_ACCOUNT_16'),('17', 'TaxAdvantageType_INDIVIDUAL_RETIREMENT_ACCOUNT_17'),('18', 'TaxAdvantageType_KEOGH'),('19', 'TaxAdvantageType_PROFIT_SHARING_PLAN'),('2', 'TaxAdvantageType_TESSA'),('20', 'TaxAdvantageType_401'),('21', 'TaxAdvantageType_SELF_DIRECTED_IRA'),('22', 'TaxAdvantageType_403'),('23', 'TaxAdvantageType_457'),('24', 'TaxAdvantageType_ROTH_IRA_24'),('25', 'TaxAdvantageType_ROTH_IRA_25'),('26', 'TaxAdvantageType_ROTH_CONVERSION_IRA_26'),('27', 'TaxAdvantageType_ROTH_CONVERSION_IRA_27'),('28', 'TaxAdvantageType_EDUCATION_IRA_28'),('29', 'TaxAdvantageType_EDUCATION_IRA_29'),('3', 'TaxAdvantageType_MINI_CASH_ISA'),('4', 'TaxAdvantageType_MINI_STOCKS_AND_SHARES_ISA'),('5', 'TaxAdvantageType_MINI_INSURANCE_ISA'),('6', 'TaxAdvantageType_CURRENT_YEAR_PAYMENT'),('7', 'TaxAdvantageType_PRIOR_YEAR_PAYMENT'),('8', 'TaxAdvantageType_ASSET_TRANSFER'),('9', 'TaxAdvantageType_EMPLOYEE_9'),('999', 'TaxAdvantageType_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.TerminationType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TerminationType VALUES 
('1', 'TerminationType_OVERNIGHT'),('2', 'TerminationType_TERM'),('3', 'TerminationType_FLEXIBLE'),('4', 'TerminationType_OPEN');

CREATE TABLE IF NOT EXISTS enum_maps.TestMessageIndicator (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TestMessageIndicator VALUES 
('N', 'TestMessageIndicator_NO'),('Y', 'TestMessageIndicator_YES');

CREATE TABLE IF NOT EXISTS enum_maps.TickDirection (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TickDirection VALUES 
('0', 'TickDirection_PLUS_TICK'),('1', 'TickDirection_ZERO_PLUS_TICK'),('2', 'TickDirection_MINUS_TICK'),('3', 'TickDirection_ZERO_MINUS_TICK');

CREATE TABLE IF NOT EXISTS enum_maps.TickRuleType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TickRuleType VALUES 
('0', 'TickRuleType_REGULAR'),('1', 'TickRuleType_VARIABLE'),('2', 'TickRuleType_FIXED'),('3', 'TickRuleType_TRADED_AS_A_SPREAD_LEG'),('4', 'TickRuleType_SETTLED_AS_A_SPREAD_LEG');

CREATE TABLE IF NOT EXISTS enum_maps.TimeInForce (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TimeInForce VALUES 
('0', 'TimeInForce_DAY'),('1', 'TimeInForce_GOOD_TILL_CANCEL'),('2', 'TimeInForce_AT_THE_OPENING'),('3', 'TimeInForce_IMMEDIATE_OR_CANCEL'),('4', 'TimeInForce_FILL_OR_KILL'),('5', 'TimeInForce_GOOD_TILL_CROSSING'),('6', 'TimeInForce_GOOD_TILL_DATE'),('7', 'TimeInForce_AT_THE_CLOSE'),('8', 'TimeInForce_GOOD_THROUGH_CROSSING'),('9', 'TimeInForce_AT_CROSSING');

CREATE TABLE IF NOT EXISTS enum_maps.TimeUnit (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TimeUnit VALUES 
('D', 'TimeUnit_DAY'),('H', 'TimeUnit_HOUR'),('Min', 'TimeUnit_MINUTE'),('Mo', 'TimeUnit_MONTH'),('S', 'TimeUnit_SECOND'),('Wk', 'TimeUnit_WEEK'),('Yr', 'TimeUnit_YEAR');

CREATE TABLE IF NOT EXISTS enum_maps.TradSesEvent (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradSesEvent VALUES 
('0', 'TradSesEvent_TRADING_RESUMES'),('1', 'TradSesEvent_CHANGE_OF_TRADING_SESSION'),('2', 'TradSesEvent_CHANGE_OF_TRADING_SUBSESSION'),('3', 'TradSesEvent_CHANGE_OF_TRADING_STATUS');

CREATE TABLE IF NOT EXISTS enum_maps.TradSesMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradSesMethod VALUES 
('1', 'TradSesMethod_ELECTRONIC'),('2', 'TradSesMethod_OPEN_OUTCRY'),('3', 'TradSesMethod_TWO_PARTY');

CREATE TABLE IF NOT EXISTS enum_maps.TradSesMode (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradSesMode VALUES 
('1', 'TradSesMode_TESTING'),('2', 'TradSesMode_SIMULATED'),('3', 'TradSesMode_PRODUCTION');

CREATE TABLE IF NOT EXISTS enum_maps.TradSesStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradSesStatus VALUES 
('0', 'TradSesStatus_UNKNOWN'),('1', 'TradSesStatus_HALTED'),('2', 'TradSesStatus_OPEN'),('3', 'TradSesStatus_CLOSED'),('4', 'TradSesStatus_PRE_OPEN'),('5', 'TradSesStatus_PRE_CLOSE'),('6', 'TradSesStatus_REQUEST_REJECTED');

CREATE TABLE IF NOT EXISTS enum_maps.TradSesStatusRejReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradSesStatusRejReason VALUES 
('1', 'TradSesStatusRejReason_UNKNOWN_OR_INVALID_TRADINGSESSIONID'),('99', 'TradSesStatusRejReason_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.TradeAllocIndicator (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradeAllocIndicator VALUES 
('0', 'TradeAllocIndicator_ALLOCATION_NOT_REQUIRED'),('1', 'TradeAllocIndicator_ALLOCATION_REQUIRED'),('2', 'TradeAllocIndicator_USE_ALLOCATION_PROVIDED_WITH_THE_TRADE'),('3', 'TradeAllocIndicator_ALLOCATION_GIVE_UP_EXECUTOR'),('4', 'TradeAllocIndicator_ALLOCATION_FROM_EXECUTOR'),('5', 'TradeAllocIndicator_ALLOCATION_TO_CLAIM_ACCOUNT');

CREATE TABLE IF NOT EXISTS enum_maps.TradeCondition (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradeCondition VALUES 
('0', 'TradeCondition_CANCEL'),('1', 'TradeCondition_IMPLIED_TRADE'),('2', 'TradeCondition_MARKETPLACE_ENTERED_TRADE'),('3', 'TradeCondition_MULT_ASSET_CLASS_MULTILEG_TRADE'),('4', 'TradeCondition_MULTILEG_TO_MULTILEG_TRADE'),('A', 'TradeCondition_CASH'),('AA', 'TradeCondition_SPREAD'),('AB', 'TradeCondition_SPREAD_ETH'),('AC', 'TradeCondition_STRADDLE'),('AD', 'TradeCondition_STRADDLE_ETH'),('AE', 'TradeCondition_STOPPED'),('AF', 'TradeCondition_STOPPED_ETH'),('AG', 'TradeCondition_REGULAR_ETH'),('AH', 'TradeCondition_COMBO'),('AI', 'TradeCondition_COMBO_ETH'),('AJ', 'TradeCondition_OFFICIAL_CLOSING_PRICE'),('AK', 'TradeCondition_PRIOR_REFERENCE_PRICE'),('AL', 'TradeCondition_STOPPED_SOLD_LAST'),('AM', 'TradeCondition_STOPPED_OUT_OF_SEQUENCE'),('AN', 'TradeCondition_OFFICAL_CLOSING_PRICE'),('AO', 'TradeCondition_CROSSED_AO'),('AP', 'TradeCondition_FAST_MARKET'),('AQ', 'TradeCondition_AUTOMATIC_EXECUTION'),('AR', 'TradeCondition_FORM_T'),('AS', 'TradeCondition_BASKET_INDEX'),('AT', 'TradeCondition_BURST_BASKET'),('AV', 'TradeCondition_OUTSIDE_SPREAD'),('B', 'TradeCondition_AVERAGE_PRICE_TRADE'),('C', 'TradeCondition_CASH_TRADE'),('D', 'TradeCondition_NEXT_DAY'),('E', 'TradeCondition_OPENING_REOPENING_TRADE_DETAIL'),('F', 'TradeCondition_INTRADAY_TRADE_DETAIL'),('G', 'TradeCondition_RULE_127_TRADE'),('H', 'TradeCondition_RULE_155_TRADE'),('I', 'TradeCondition_SOLD_LAST'),('J', 'TradeCondition_NEXT_DAY_TRADE'),('K', 'TradeCondition_OPENED'),('L', 'TradeCondition_SELLER'),('M', 'TradeCondition_SOLD'),('N', 'TradeCondition_STOPPED_STOCK'),('P', 'TradeCondition_IMBALANCE_MORE_BUYERS'),('Q', 'TradeCondition_IMBALANCE_MORE_SELLERS'),('R', 'TradeCondition_OPENING_PRICE'),('S', 'TradeCondition_BARGAIN_CONDITION'),('T', 'TradeCondition_CONVERTED_PRICE_INDICATOR'),('U', 'TradeCondition_EXCHANGE_LAST'),('V', 'TradeCondition_FINAL_PRICE_OF_SESSION'),('W', 'TradeCondition_EX_PIT'),('X', 'TradeCondition_CROSSED_X'),('Y', 'TradeCondition_TRADES_RESULTING_FROM_MANUAL_SLOW_QUOTE'),('Z', 'TradeCondition_TRADES_RESULTING_FROM_INTERMARKET_SWEEP'),('a', 'TradeCondition_VOLUME_ONLY'),('b', 'TradeCondition_DIRECT_PLUS'),('c', 'TradeCondition_ACQUISITION'),('d', 'TradeCondition_BUNCHED'),('e', 'TradeCondition_DISTRIBUTION'),('f', 'TradeCondition_BUNCHED_SALE'),('g', 'TradeCondition_SPLIT_TRADE'),('h', 'TradeCondition_CANCEL_STOPPED'),('i', 'TradeCondition_CANCEL_ETH'),('j', 'TradeCondition_CANCEL_STOPPED_ETH'),('k', 'TradeCondition_OUT_OF_SEQUENCE_ETH'),('l', 'TradeCondition_CANCEL_LAST_ETH'),('m', 'TradeCondition_SOLD_LAST_SALE_ETH'),('n', 'TradeCondition_CANCEL_LAST'),('o', 'TradeCondition_SOLD_LAST_SALE'),('p', 'TradeCondition_CANCEL_OPEN'),('q', 'TradeCondition_CANCEL_OPEN_ETH'),('r', 'TradeCondition_OPENED_SALE_ETH'),('s', 'TradeCondition_CANCEL_ONLY'),('t', 'TradeCondition_CANCEL_ONLY_ETH'),('u', 'TradeCondition_LATE_OPEN_ETH'),('v', 'TradeCondition_AUTO_EXECUTION_ETH'),('w', 'TradeCondition_REOPEN'),('x', 'TradeCondition_REOPEN_ETH'),('y', 'TradeCondition_ADJUSTED'),('z', 'TradeCondition_ADJUSTED_ETH');

CREATE TABLE IF NOT EXISTS enum_maps.TradeHandlingInstr (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradeHandlingInstr VALUES 
('0', 'TradeHandlingInstr_TRADE_CONFIRMATION'),('1', 'TradeHandlingInstr_TWO_PARTY_REPORT'),('2', 'TradeHandlingInstr_ONE_PARTY_REPORT_FOR_MATCHING'),('3', 'TradeHandlingInstr_ONE_PARTY_REPORT_FOR_PASS_THROUGH'),('4', 'TradeHandlingInstr_AUTOMATED_FLOOR_ORDER_ROUTING'),('5', 'TradeHandlingInstr_TWO_PARTY_REPORT_FOR_CLAIM');

CREATE TABLE IF NOT EXISTS enum_maps.TradePublishIndicator (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradePublishIndicator VALUES 
('0', 'TradePublishIndicator_DO_NOT_PUBLISH_TRADE'),('1', 'TradePublishIndicator_PUBLISH_TRADE'),('2', 'TradePublishIndicator_DEFERRED_PUBLICATION');

CREATE TABLE IF NOT EXISTS enum_maps.TradeReportRejectReason (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradeReportRejectReason VALUES 
('0', 'TradeReportRejectReason_SUCCESSFUL'),('1', 'TradeReportRejectReason_INVALID_PARTY_ONFORMATION'),('2', 'TradeReportRejectReason_UNKNOWN_INSTRUMENT'),('3', 'TradeReportRejectReason_UNAUTHORIZED_TO_REPORT_TRADES'),('4', 'TradeReportRejectReason_INVALID_TRADE_TYPE'),('99', 'TradeReportRejectReason_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.TradeReportTransType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradeReportTransType VALUES 
('0', 'TradeReportTransType_NEW'),('1', 'TradeReportTransType_CANCEL'),('2', 'TradeReportTransType_REPLACE'),('3', 'TradeReportTransType_RELEASE'),('4', 'TradeReportTransType_REVERSE'),('5', 'TradeReportTransType_CANCEL_DUE_TO_BACK_OUT_OF_TRADE');

CREATE TABLE IF NOT EXISTS enum_maps.TradeReportType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradeReportType VALUES 
('0', 'TradeReportType_SUBMIT'),('1', 'TradeReportType_ALLEGED_1'),('10', 'TradeReportType_PENDED'),('11', 'TradeReportType_ALLEGED_NEW'),('12', 'TradeReportType_ALLEGED_ADDENDUM'),('13', 'TradeReportType_ALLEGED_NO_WAS'),('14', 'TradeReportType_ALLEGED_TRADE_REPORT_CANCEL'),('15', 'TradeReportType_ALLEGED_15'),('2', 'TradeReportType_ACCEPT'),('3', 'TradeReportType_DECLINE'),('4', 'TradeReportType_ADDENDUM'),('5', 'TradeReportType_NO_WAS'),('6', 'TradeReportType_TRADE_REPORT_CANCEL'),('7', 'TradeReportType_7'),('8', 'TradeReportType_DEFAULTED'),('9', 'TradeReportType_INVALID_CMTA');

CREATE TABLE IF NOT EXISTS enum_maps.TradeRequestResult (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradeRequestResult VALUES 
('0', 'TradeRequestResult_SUCCESSFUL'),('1', 'TradeRequestResult_INVALID_OR_UNKNOWN_INSTRUMENT'),('2', 'TradeRequestResult_INVALID_TYPE_OF_TRADE_REQUESTED'),('3', 'TradeRequestResult_INVALID_PARTIES'),('4', 'TradeRequestResult_INVALID_TRANSPORT_TYPE_REQUESTED'),('5', 'TradeRequestResult_INVALID_DESTINATION_REQUESTED'),('8', 'TradeRequestResult_TRADEREQUESTTYPE_NOT_SUPPORTED'),('9', 'TradeRequestResult_NOT_AUTHORIZED'),('99', 'TradeRequestResult_OTHER');

CREATE TABLE IF NOT EXISTS enum_maps.TradeRequestStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradeRequestStatus VALUES 
('0', 'TradeRequestStatus_ACCEPTED'),('1', 'TradeRequestStatus_COMPLETED'),('2', 'TradeRequestStatus_REJECTED');

CREATE TABLE IF NOT EXISTS enum_maps.TradeRequestType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradeRequestType VALUES 
('0', 'TradeRequestType_ALL_TRADES'),('1', 'TradeRequestType_MATCHED_TRADES_MATCHING_CRITERIA_PROVIDED_ON_REQUEST'),('2', 'TradeRequestType_UNMATCHED_TRADES_THAT_MATCH_CRITERIA'),('3', 'TradeRequestType_UNREPORTED_TRADES_THAT_MATCH_CRITERIA'),('4', 'TradeRequestType_ADVISORIES_THAT_MATCH_CRITERIA');

CREATE TABLE IF NOT EXISTS enum_maps.TradeType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradeType VALUES 
('A', 'TradeType_AGENCY'),('G', 'TradeType_VWAP_GUARANTEE'),('J', 'TradeType_GUARANTEED_CLOSE'),('R', 'TradeType_RISK_TRADE');

CREATE TABLE IF NOT EXISTS enum_maps.TradedFlatSwitch (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradedFlatSwitch VALUES 
('N', 'TradedFlatSwitch_NO'),('Y', 'TradedFlatSwitch_YES');

CREATE TABLE IF NOT EXISTS enum_maps.TradingSessionID (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradingSessionID VALUES 
('1', 'TradingSessionID_DAY'),('2', 'TradingSessionID_HALFDAY'),('3', 'TradingSessionID_MORNING'),('4', 'TradingSessionID_AFTERNOON'),('5', 'TradingSessionID_EVENING'),('6', 'TradingSessionID_AFTER_HOURS');

CREATE TABLE IF NOT EXISTS enum_maps.TradingSessionSubID (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TradingSessionSubID VALUES 
('1', 'TradingSessionSubID_PRE_TRADING'),('2', 'TradingSessionSubID_OPENING_OR_OPENING_AUCTION'),('3', 'TradingSessionSubID_3'),('4', 'TradingSessionSubID_CLOSING_OR_CLOSING_AUCTION'),('5', 'TradingSessionSubID_POST_TRADING'),('6', 'TradingSessionSubID_INTRADAY_AUCTION'),('7', 'TradingSessionSubID_QUIESCENT');

CREATE TABLE IF NOT EXISTS enum_maps.TrdRegTimestampType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TrdRegTimestampType VALUES 
('1', 'TrdRegTimestampType_EXECUTION_TIME'),('2', 'TrdRegTimestampType_TIME_IN'),('3', 'TrdRegTimestampType_TIME_OUT'),('4', 'TrdRegTimestampType_BROKER_RECEIPT'),('5', 'TrdRegTimestampType_BROKER_EXECUTION'),('6', 'TrdRegTimestampType_DESK_RECEIPT'),('7', 'TrdRegTimestampType_SUBMISSION_TO_CLEARING');

CREATE TABLE IF NOT EXISTS enum_maps.TrdRptStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TrdRptStatus VALUES 
('0', 'TrdRptStatus_ACCEPTED'),('1', 'TrdRptStatus_REJECTED'),('3', 'TrdRptStatus_ACCEPTED_WITH_ERRORS');

CREATE TABLE IF NOT EXISTS enum_maps.TrdSubType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TrdSubType VALUES 
('0', 'TrdSubType_CMTA'),('1', 'TrdSubType_INTERNAL_TRANSFER_OR_ADJUSTMENT'),('10', 'TrdSubType_TRANSACTION_FROM_ASSIGNMENT'),('11', 'TrdSubType_ACATS'),('14', 'TrdSubType_AI'),('15', 'TrdSubType_B'),('16', 'TrdSubType_K'),('17', 'TrdSubType_LC'),('18', 'TrdSubType_M'),('19', 'TrdSubType_N'),('2', 'TrdSubType_EXTERNAL_TRANSFER_OR_TRANSFER_OF_ACCOUNT'),('20', 'TrdSubType_NM'),('21', 'TrdSubType_NR'),('22', 'TrdSubType_P'),('23', 'TrdSubType_PA'),('24', 'TrdSubType_PC'),('25', 'TrdSubType_PN'),('26', 'TrdSubType_R'),('27', 'TrdSubType_RO'),('28', 'TrdSubType_RT'),('29', 'TrdSubType_SW'),('3', 'TrdSubType_REJECT_FOR_SUBMITTING_SIDE'),('30', 'TrdSubType_T'),('31', 'TrdSubType_WN'),('32', 'TrdSubType_WT'),('33', 'TrdSubType_OFF_HOURS_TRADE'),('34', 'TrdSubType_ON_HOURS_TRADE'),('35', 'TrdSubType_OTC_QUOTE'),('36', 'TrdSubType_CONVERTED_SWAP'),('37', 'TrdSubType_CROSSED_TRADE'),('38', 'TrdSubType_INTERIM_PROTECTED_TRADE'),('39', 'TrdSubType_LARGE_IN_SCALE'),('4', 'TrdSubType_ADVISORY_FOR_CONTRA_SIDE'),('5', 'TrdSubType_OFFSET_DUE_TO_AN_ALLOCATION'),('6', 'TrdSubType_ONSET_DUE_TO_AN_ALLOCATION'),('7', 'TrdSubType_DIFFERENTIAL_SPREAD'),('8', 'TrdSubType_IMPLIED_SPREAD_LEG_EXECUTED_AGAINST_AN_OUTRIGHT'),('9', 'TrdSubType_TRANSACTION_FROM_EXERCISE');

CREATE TABLE IF NOT EXISTS enum_maps.TrdType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TrdType VALUES 
('0', 'TrdType_REGULAR_TRADE'),('1', 'TrdType_BLOCK_TRADE_1'),('10', 'TrdType_AFTER_HOURS_TRADE'),('11', 'TrdType_EXCHANGE_FOR_RISK'),('12', 'TrdType_EXCHANGE_FOR_SWAP'),('13', 'TrdType_EXCHANGE_OF_FUTURES_FOR'),('14', 'TrdType_EXCHANGE_OF_OPTIONS_FOR_OPTIONS'),('15', 'TrdType_TRADING_AT_SETTLEMENT'),('16', 'TrdType_ALL_OR_NONE'),('17', 'TrdType_FUTURES_LARGE_ORDER_EXECUTION'),('18', 'TrdType_EXCHANGE_OF_FUTURES_FOR_FUTURES'),('19', 'TrdType_OPTION_INTERIM_TRADE'),('2', 'TrdType_EFP'),('20', 'TrdType_OPTION_CABINET_TRADE'),('22', 'TrdType_PRIVATELY_NEGOTIATED_TRADES'),('23', 'TrdType_SUBSTITUTION_OF_FUTURES_FOR_FORWARDS'),('24', 'TrdType_ERROR_TRADE'),('25', 'TrdType_SPECIAL_CUM_DIVIDEND'),('26', 'TrdType_SPECIAL_EX_DIVIDEND'),('27', 'TrdType_SPECIAL_CUM_COUPON'),('28', 'TrdType_SPECIAL_EX_COUPON'),('29', 'TrdType_CASH_SETTLEMENT'),('3', 'TrdType_TRANSFER'),('30', 'TrdType_SPECIAL_PRICE'),('31', 'TrdType_GUARANTEED_DELIVERY'),('32', 'TrdType_SPECIAL_CUM_RIGHTS'),('33', 'TrdType_SPECIAL_EX_RIGHTS'),('34', 'TrdType_SPECIAL_CUM_CAPITAL_REPAYMENTS'),('35', 'TrdType_SPECIAL_EX_CAPITAL_REPAYMENTS'),('36', 'TrdType_SPECIAL_CUM_BONUS'),('37', 'TrdType_SPECIAL_EX_BONUS'),('38', 'TrdType_BLOCK_TRADE_38'),('39', 'TrdType_WORKED_PRINCIPAL_TRADE'),('4', 'TrdType_LATE_TRADE'),('40', 'TrdType_BLOCK_TRADES'),('41', 'TrdType_NAME_CHANGE'),('42', 'TrdType_PORTFOLIO_TRANSFER'),('43', 'TrdType_PROROGATION_BUY'),('44', 'TrdType_PROROGATION_SELL'),('45', 'TrdType_OPTION_EXERCISE'),('46', 'TrdType_DELTA_NEUTRAL_TRANSACTION'),('47', 'TrdType_FINANCING_TRANSACTION'),('48', 'TrdType_NON_STANDARD_SETTLEMENT'),('49', 'TrdType_DERIVATIVE_RELATED_TRANSACTION'),('5', 'TrdType_T_TRADE'),('50', 'TrdType_PORTFOLIO_TRADE'),('51', 'TrdType_VOLUME_WEIGHTED_AVERAGE_TRADE'),('52', 'TrdType_EXCHANGE_GRANTED_TRADE'),('53', 'TrdType_REPURCHASE_AGREEMENT'),('54', 'TrdType_OTC'),('55', 'TrdType_EXCHANGE_BASIS_FACILITY'),('6', 'TrdType_WEIGHTED_AVERAGE_PRICE_TRADE'),('7', 'TrdType_BUNCHED_TRADE'),('8', 'TrdType_LATE_BUNCHED_TRADE'),('9', 'TrdType_PRIOR_REFERENCE_PRICE_TRADE');

CREATE TABLE IF NOT EXISTS enum_maps.TriggerAction (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TriggerAction VALUES 
('1', 'TriggerAction_ACTIVATE'),('2', 'TriggerAction_MODIFY'),('3', 'TriggerAction_CANCEL');

CREATE TABLE IF NOT EXISTS enum_maps.TriggerOrderType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TriggerOrderType VALUES 
('1', 'TriggerOrderType_MARKET'),('2', 'TriggerOrderType_LIMIT');

CREATE TABLE IF NOT EXISTS enum_maps.TriggerPriceDirection (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TriggerPriceDirection VALUES 
('D', 'TriggerPriceDirection_TRIGGER_IF_THE_PRICE_OF_THE_SPECIFIED_TYPE_GOES_DOWN_TO_OR_THROUGH_THE_SPECIFIED_TRIGGER_PRICE'),('U', 'TriggerPriceDirection_TRIGGER_IF_THE_PRICE_OF_THE_SPECIFIED_TYPE_GOES_UP_TO_OR_THROUGH_THE_SPECIFIED_TRIGGER_PRICE');

CREATE TABLE IF NOT EXISTS enum_maps.TriggerPriceType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TriggerPriceType VALUES 
('1', 'TriggerPriceType_BEST_OFFER'),('2', 'TriggerPriceType_LAST_TRADE'),('3', 'TriggerPriceType_BEST_BID'),('4', 'TriggerPriceType_BEST_BID_OR_LAST_TRADE'),('5', 'TriggerPriceType_BEST_OFFER_OR_LAST_TRADE'),('6', 'TriggerPriceType_BEST_MID');

CREATE TABLE IF NOT EXISTS enum_maps.TriggerPriceTypeScope (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TriggerPriceTypeScope VALUES 
('0', 'TriggerPriceTypeScope_NONE'),('1', 'TriggerPriceTypeScope_LOCAL'),('2', 'TriggerPriceTypeScope_NATIONAL'),('3', 'TriggerPriceTypeScope_GLOBAL');

CREATE TABLE IF NOT EXISTS enum_maps.TriggerType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.TriggerType VALUES 
('1', 'TriggerType_PARTIAL_EXECUTION'),('2', 'TriggerType_SPECIFIED_TRADING_SESSION'),('3', 'TriggerType_NEXT_AUCTION'),('4', 'TriggerType_PRICE_MOVEMENT');

CREATE TABLE IF NOT EXISTS enum_maps.UnderlyingCashType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.UnderlyingCashType VALUES 
('DIFF', 'UnderlyingCashType_DIFF'),('FIXED', 'UnderlyingCashType_FIXED');

CREATE TABLE IF NOT EXISTS enum_maps.UnderlyingFXRateCalc (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.UnderlyingFXRateCalc VALUES 
('D', 'UnderlyingFXRateCalc_DIVIDE'),('M', 'UnderlyingFXRateCalc_MULTIPLY');

CREATE TABLE IF NOT EXISTS enum_maps.UnderlyingPriceDeterminationMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.UnderlyingPriceDeterminationMethod VALUES 
('1', 'UnderlyingPriceDeterminationMethod_REGULAR'),('2', 'UnderlyingPriceDeterminationMethod_SPECIAL_REFERENCE'),('3', 'UnderlyingPriceDeterminationMethod_OPTIMAL_VALUE'),('4', 'UnderlyingPriceDeterminationMethod_AVERAGE_VALUE');

CREATE TABLE IF NOT EXISTS enum_maps.UnderlyingSettlementType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.UnderlyingSettlementType VALUES 
('2', 'UnderlyingSettlementType_T_PLUS_1'),('4', 'UnderlyingSettlementType_T_PLUS_3'),('5', 'UnderlyingSettlementType_T_PLUS_4');

CREATE TABLE IF NOT EXISTS enum_maps.UnitOfMeasure (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.UnitOfMeasure VALUES 
('Alw', 'UnitOfMeasure_ALLOWANCES'),('Bbl', 'UnitOfMeasure_BARRELS'),('Bcf', 'UnitOfMeasure_BILLION_CUBIC_FEET'),('Bu', 'UnitOfMeasure_BUSHELS'),('Gal', 'UnitOfMeasure_GALLONS'),('MMBtu', 'UnitOfMeasure_ONE_MILLION_BTU'),('MMbbl', 'UnitOfMeasure_MILLION_BARRELS'),('MWh', 'UnitOfMeasure_MEGAWATT_HOURS'),('USD', 'UnitOfMeasure_US_DOLLARS'),('lbs', 'UnitOfMeasure_POUNDS'),('oz_tr', 'UnitOfMeasure_TROY_OUNCES'),('t', 'UnitOfMeasure_METRIC_TONS'),('tn', 'UnitOfMeasure_TONS');

CREATE TABLE IF NOT EXISTS enum_maps.UnsolicitedIndicator (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.UnsolicitedIndicator VALUES 
('N', 'UnsolicitedIndicator_NO'),('Y', 'UnsolicitedIndicator_YES');

CREATE TABLE IF NOT EXISTS enum_maps.Urgency (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.Urgency VALUES 
('0', 'Urgency_NORMAL'),('1', 'Urgency_FLASH'),('2', 'Urgency_BACKGROUND');

CREATE TABLE IF NOT EXISTS enum_maps.UserRequestType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.UserRequestType VALUES 
('1', 'UserRequestType_LOG_ON_USER'),('2', 'UserRequestType_LOG_OFF_USER'),('3', 'UserRequestType_CHANGE_PASSWORD_FOR_USER'),('4', 'UserRequestType_REQUEST_INDIVIDUAL_USER_STATUS');

CREATE TABLE IF NOT EXISTS enum_maps.UserStatus (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.UserStatus VALUES 
('1', 'UserStatus_LOGGED_IN'),('2', 'UserStatus_NOT_LOGGED_IN'),('3', 'UserStatus_USER_NOT_RECOGNISED'),('4', 'UserStatus_PASSWORD_INCORRECT'),('5', 'UserStatus_PASSWORD_CHANGED'),('6', 'UserStatus_OTHER'),('7', 'UserStatus_FORCED_USER_LOGOUT_BY_EXCHANGE'),('8', 'UserStatus_SESSION_SHUTDOWN_WARNING');

CREATE TABLE IF NOT EXISTS enum_maps.ValuationMethod (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.ValuationMethod VALUES 
('CDS', 'ValuationMethod_CDS_STYLE_COLLATERALIZATION_OF_MARKET_TO_MARKET_AND_COUPON'),('CDSD', 'ValuationMethod_CDS_IN_DELIVERY'),('EQTY', 'ValuationMethod_PREMIUM_STYLE'),('FUT', 'ValuationMethod_FUTURES_STYLE_MARK_TO_MARKET'),('FUTDA', 'ValuationMethod_FUTURES_STYLE_WITH_AN_ATTACHED_CASH_ADJUSTMENT');

CREATE TABLE IF NOT EXISTS enum_maps.VenueType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.VenueType VALUES 
('E', 'VenueType_ELECTRONIC'),('P', 'VenueType_PIT'),('X', 'VenueType_EX_PIT');

CREATE TABLE IF NOT EXISTS enum_maps.WorkingIndicator (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.WorkingIndicator VALUES 
('N', 'WorkingIndicator_NO'),('Y', 'WorkingIndicator_YES');

CREATE TABLE IF NOT EXISTS enum_maps.YieldType (value VARCHAR, name VARCHAR);
INSERT INTO enum_maps.YieldType VALUES 
('AFTERTAX', 'YieldType_AFTER_TAX_YIELD'),('ANNUAL', 'YieldType_ANNUAL_YIELD'),('ATISSUE', 'YieldType_YIELD_AT_ISSUE'),('AVGLIFE', 'YieldType_YIELD_TO_AVERAGE_LIFE_THE_YIELD_ASSUMING_THAT_ALL_SINKS'),('AVGMATURITY', 'YieldType_YIELD_TO_AVG_MATURITY'),('BOOK', 'YieldType_BOOK_YIELD'),('CALL', 'YieldType_YIELD_TO_NEXT_CALL'),('CHANGE', 'YieldType_YIELD_CHANGE_SINCE_CLOSE'),('CLOSE', 'YieldType_CLOSING_YIELD'),('COMPOUND', 'YieldType_COMPOUND_YIELD'),('CURRENT', 'YieldType_CURRENT_YIELD'),('GOVTEQUIV', 'YieldType_GVNT_EQUIVALENT_YIELD'),('GROSS', 'YieldType_TRUE_GROSS_YIELD'),('INFLATION', 'YieldType_YIELD_WITH_INFLATION_ASSUMPTION'),('INVERSEFLOATER', 'YieldType_INVERSE_FLOATER_BOND_YIELD'),('LASTCLOSE', 'YieldType_MOST_RECENT_CLOSING_YIELD'),('LASTMONTH', 'YieldType_CLOSING_YIELD_MOST_RECENT_MONTH'),('LASTQUARTER', 'YieldType_CLOSING_YIELD_MOST_RECENT_QUARTER'),('LASTYEAR', 'YieldType_CLOSING_YIELD_MOST_RECENT_YEAR'),('LONGAVGLIFE', 'YieldType_YIELD_TO_LONGEST_AVERAGE_LIFE'),('LONGEST', 'YieldType_YIELD_TO_LONGEST_AVERAGE'),('MARK', 'YieldType_MARK_TO_MARKET_YIELD'),('MATURITY', 'YieldType_YIELD_TO_MATURITY'),('NEXTREFUND', 'YieldType_YIELD_TO_NEXT_REFUND'),('OPENAVG', 'YieldType_OPEN_AVERAGE_YIELD'),('PREVCLOSE', 'YieldType_PREVIOUS_CLOSE_YIELD'),('PROCEEDS', 'YieldType_PROCEEDS_YIELD'),('PUT', 'YieldType_YIELD_TO_NEXT_PUT'),('SEMIANNUAL', 'YieldType_SEMI_ANNUAL_YIELD'),('SHORTAVGLIFE', 'YieldType_YIELD_TO_SHORTEST_AVERAGE_LIFE'),('SHORTEST', 'YieldType_YIELD_TO_SHORTEST_AVERAGE'),('SIMPLE', 'YieldType_SIMPLE_YIELD'),('TAXEQUIV', 'YieldType_TAX_EQUIVALENT_YIELD'),('TENDER', 'YieldType_YIELD_TO_TENDER_DATE'),('TRUE', 'YieldType_TRUE_YIELD'),('VALUE1/32', 'YieldType_YIELD_VALUE_OF_1_32_THE_AMOUNT_THAT_THE_YIELD_WILL_CHANGE_FOR_A_1_32ND_CHANGE_IN_PRICE'),('VALUE1_32', 'YieldType_YIELD_VALUE_OF_1_32'),('WORST', 'YieldType_YIELD_TO_WORST');

