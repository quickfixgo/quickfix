package enum

//Enum values for AccountType
const (
	AccountType_ACCOUNT_IS_CARRIED_ON_CUSTOMER_SIDE_OF_THE_BOOKS                       = "1"
	AccountType_ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS                       = "2"
	AccountType_HOUSE_TRADER                                                           = "3"
	AccountType_FLOOR_TRADER                                                           = "4"
	AccountType_ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS_AND_IS_CROSS_MARGINED = "6"
	AccountType_ACCOUNT_IS_HOUSE_TRADER_AND_IS_CROSS_MARGINED                          = "7"
	AccountType_JOINT_BACK_OFFICE_ACCOUNT                                              = "8"
)

//Enum values for AcctIDSource
const (
	AcctIDSource_BIC       = "1"
	AcctIDSource_SID_CODE  = "2"
	AcctIDSource_TFM       = "3"
	AcctIDSource_OMGEO     = "4"
	AcctIDSource_DTCC_CODE = "5"
	AcctIDSource_OTHER     = "99"
)

//Enum values for Adjustment
const (
	Adjustment_CANCEL     = "1"
	Adjustment_ERROR      = "2"
	Adjustment_CORRECTION = "3"
)

//Enum values for AdjustmentType
const (
	AdjustmentType_PROCESS_REQUEST_AS_MARGIN_DISPOSITION = "0"
	AdjustmentType_DELTA_PLUS                            = "1"
	AdjustmentType_DELTA_MINUS                           = "2"
	AdjustmentType_FINAL                                 = "3"
)

//Enum values for AdvSide
const (
	AdvSide_BUY   = "B"
	AdvSide_SELL  = "S"
	AdvSide_TRADE = "T"
	AdvSide_CROSS = "X"
)

//Enum values for AdvTransType
const (
	AdvTransType_CANCEL  = "C"
	AdvTransType_NEW     = "N"
	AdvTransType_REPLACE = "R"
)

//Enum values for AffirmStatus
const (
	AffirmStatus_RECEIVED                         = "1"
	AffirmStatus_CONFIRM_REJECTED_IE_NOT_AFFIRMED = "2"
	AffirmStatus_AFFIRMED                         = "3"
)

//Enum values for AggregatedBook
const (
	AggregatedBook_NO  = "N"
	AggregatedBook_YES = "Y"
)

//Enum values for AggressorIndicator
const (
	AggressorIndicator_NO  = "N"
	AggressorIndicator_YES = "Y"
)

//Enum values for AllocAccountType
const (
	AllocAccountType_ACCOUNT_IS_CARRIED_PN_CUSTOMER_SIDE_OF_BOOKS                           = "1"
	AllocAccountType_ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS                       = "2"
	AllocAccountType_HOUSE_TRADER                                                           = "3"
	AllocAccountType_FLOOR_TRADER                                                           = "4"
	AllocAccountType_ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS_AND_IS_CROSS_MARGINED = "6"
	AllocAccountType_ACCOUNT_IS_HOUSE_TRADER_AND_IS_CROSS_MARGINED                          = "7"
	AllocAccountType_JOINT_BACK_OFFICE_ACCOUNT                                              = "8"
)

//Enum values for AllocCancReplaceReason
const (
	AllocCancReplaceReason_ORIGINAL_DETAILS_INCOMPLETE_INCORRECT = "1"
	AllocCancReplaceReason_CHANGE_IN_UNDERLYING_ORDER_DETAILS    = "2"
	AllocCancReplaceReason_OTHER                                 = "99"
)

//Enum values for AllocHandlInst
const (
	AllocHandlInst_MATCH             = "1"
	AllocHandlInst_FORWARD           = "2"
	AllocHandlInst_FORWARD_AND_MATCH = "3"
)

//Enum values for AllocIntermedReqType
const (
	AllocIntermedReqType_PENDING_ACCEPT       = "1"
	AllocIntermedReqType_PENDING_RELEASE      = "2"
	AllocIntermedReqType_PENDING_REVERSAL     = "3"
	AllocIntermedReqType_ACCEPT               = "4"
	AllocIntermedReqType_BLOCK_LEVEL_REJECT   = "5"
	AllocIntermedReqType_ACCOUNT_LEVEL_REJECT = "6"
)

//Enum values for AllocLinkType
const (
	AllocLinkType_FX_NETTING = "0"
	AllocLinkType_FX_SWAP    = "1"
)

//Enum values for AllocMethod
const (
	AllocMethod_AUTOMATIC = "1"
	AllocMethod_GUARANTOR = "2"
	AllocMethod_MANUAL    = "3"
)

//Enum values for AllocNoOrdersType
const (
	AllocNoOrdersType_NOT_SPECIFIED          = "0"
	AllocNoOrdersType_EXPLICIT_LIST_PROVIDED = "1"
)

//Enum values for AllocPositionEffect
const (
	AllocPositionEffect_CLOSE  = "C"
	AllocPositionEffect_FIFO   = "F"
	AllocPositionEffect_OPEN   = "O"
	AllocPositionEffect_ROLLED = "R"
)

//Enum values for AllocRejCode
const (
	AllocRejCode_UNKNOWN_ACCOUNT                   = "0"
	AllocRejCode_INCORRECT_QUANTITY                = "1"
	AllocRejCode_UNKNOWN_OR_STALE_EXECID           = "10"
	AllocRejCode_MISMATCHED_DATA                   = "11"
	AllocRejCode_UNKNOWN_CLORDID                   = "12"
	AllocRejCode_WAREHOUSE_REQUEST_REJECTED        = "13"
	AllocRejCode_INCORRECT_AVERAGEG_PRICE          = "2"
	AllocRejCode_UNKNOWN_EXECUTING_BROKER_MNEMONIC = "3"
	AllocRejCode_COMMISSION_DIFFERENCE             = "4"
	AllocRejCode_UNKNOWN_ORDERID                   = "5"
	AllocRejCode_UNKNOWN_LISTID                    = "6"
	AllocRejCode_OTHER_7                           = "7"
	AllocRejCode_INCORRECT_ALLOCATED_QUANTITY      = "8"
	AllocRejCode_CALCULATION_DIFFERENCE            = "9"
	AllocRejCode_OTHER_99                          = "99"
)

//Enum values for AllocReportType
const (
	AllocReportType_REJECT                                  = "10"
	AllocReportType_ACCEPT_PENDING                          = "11"
	AllocReportType_COMPLETE                                = "12"
	AllocReportType_REVERSE_PENDING                         = "14"
	AllocReportType_PRELIMINARY_REQUEST_TO_INTERMEDIARY     = "2"
	AllocReportType_SELLSIDE_CALCULATED_USING_PRELIMINARY   = "3"
	AllocReportType_SELLSIDE_CALCULATED_WITHOUT_PRELIMINARY = "4"
	AllocReportType_WAREHOUSE_RECAP                         = "5"
	AllocReportType_REQUEST_TO_INTERMEDIARY                 = "8"
	AllocReportType_ACCEPT                                  = "9"
)

//Enum values for AllocSettlInstType
const (
	AllocSettlInstType_USE_DEFAULT_INSTRUCTIONS        = "0"
	AllocSettlInstType_DERIVE_FROM_PARAMETERS_PROVIDED = "1"
	AllocSettlInstType_FULL_DETAILS_PROVIDED           = "2"
	AllocSettlInstType_SSI_DB_IDS_PROVIDED             = "3"
	AllocSettlInstType_PHONE_FOR_INSTRUCTIONS          = "4"
)

//Enum values for AllocStatus
const (
	AllocStatus_ACCEPTED                 = "0"
	AllocStatus_BLOCK_LEVEL_REJECT       = "1"
	AllocStatus_ACCOUNT_LEVEL_REJECT     = "2"
	AllocStatus_RECEIVED                 = "3"
	AllocStatus_INCOMPLETE               = "4"
	AllocStatus_REJECTED_BY_INTERMEDIARY = "5"
	AllocStatus_ALLOCATION_PENDING       = "6"
	AllocStatus_REVERSED                 = "7"
)

//Enum values for AllocTransType
const (
	AllocTransType_NEW                            = "0"
	AllocTransType_REPLACE                        = "1"
	AllocTransType_CANCEL                         = "2"
	AllocTransType_PRELIMINARY                    = "3"
	AllocTransType_CALCULATED                     = "4"
	AllocTransType_CALCULATED_WITHOUT_PRELIMINARY = "5"
	AllocTransType_REVERSAL                       = "6"
)

//Enum values for AllocType
const (
	AllocType_CALCULATED                              = "1"
	AllocType_REJECT                                  = "10"
	AllocType_ACCEPT_PENDING                          = "11"
	AllocType_INCOMPLETE_GROUP                        = "12"
	AllocType_COMPLETE_GROUP                          = "13"
	AllocType_REVERSAL_PENDING                        = "14"
	AllocType_PRELIMINARY                             = "2"
	AllocType_SELLSIDE_CALCULATED_USING_PRELIMINARY   = "3"
	AllocType_SELLSIDE_CALCULATED_WITHOUT_PRELIMINARY = "4"
	AllocType_READY_TO_BOOK                           = "5"
	AllocType_BUYSIDE_READY_TO_BOOK                   = "6"
	AllocType_WAREHOUSE_INSTRUCTION                   = "7"
	AllocType_REQUEST_TO_INTERMEDIARY                 = "8"
	AllocType_ACCEPT                                  = "9"
)

//Enum values for ApplQueueAction
const (
	ApplQueueAction_NO_ACTION_TAKEN = "0"
	ApplQueueAction_QUEUE_FLUSHED   = "1"
	ApplQueueAction_OVERLAY_LAST    = "2"
	ApplQueueAction_END_SESSION     = "3"
)

//Enum values for ApplQueueResolution
const (
	ApplQueueResolution_NO_ACTION_TAKEN = "0"
	ApplQueueResolution_QUEUE_FLUSHED   = "1"
	ApplQueueResolution_OVERLAY_LAST    = "2"
	ApplQueueResolution_END_SESSION     = "3"
)

//Enum values for ApplReportType
const (
	ApplReportType_RESET_APPLSEQNUM_TO_NEW_VALUE_SPECIFIED_IN_APPLNEWSEQNUM                               = "0"
	ApplReportType_REPORTS_THAT_THE_LAST_MESSAGE_HAS_BEEN_SENT_FOR_THE_APPLIDS_REFER_TO_REFAPPLLASTSEQNUM = "1"
	ApplReportType_HEARTBEAT_MESSAGE_INDICATING_THAT_APPLICATION_IDENTIFIED_BY_REFAPPLID                  = "2"
	ApplReportType_APPLICATION_MESSAGE_RE_SEND_COMPLETED                                                  = "3"
)

//Enum values for ApplReqType
const (
	ApplReqType_RETRANSMISSION_OF_APPLICATION_MESSAGES_FOR_THE_SPECIFIED_APPLICATIONS        = "0"
	ApplReqType_SUBSCRIPTION_TO_THE_SPECIFIED_APPLICATIONS                                   = "1"
	ApplReqType_REQUEST_FOR_THE_LAST_APPLLASTSEQNUM_PUBLISHED_FOR_THE_SPECIFIED_APPLICATIONS = "2"
	ApplReqType_REQUEST_VALID_SET_OF_APPLICATIONS                                            = "3"
	ApplReqType_UNSUBSCRIBE_TO_THE_SPECIFIED_APPLICATIONS                                    = "4"
	ApplReqType_CANCEL_RETRANSMISSION                                                        = "5"
	ApplReqType_CANCEL_RETRANSMISSION_AND_UNSUBSCRIBE_TO_THE_SPECIFIED_APPLICATIONS          = "6"
)

//Enum values for ApplResponseError
const (
	ApplResponseError_APPLICATION_DOES_NOT_EXIST           = "0"
	ApplResponseError_MESSAGES_REQUESTED_ARE_NOT_AVAILABLE = "1"
	ApplResponseError_USER_NOT_AUTHORIZED_FOR_APPLICATION  = "2"
)

//Enum values for ApplResponseType
const (
	ApplResponseType_REQUEST_SUCCESSFULLY_PROCESSED = "0"
	ApplResponseType_APPLICATION_DOES_NOT_EXIST     = "1"
	ApplResponseType_MESSAGES_NOT_AVAILABLE         = "2"
)

//Enum values for ApplVerID
const (
	ApplVerID_FIX27    = "0"
	ApplVerID_FIX30    = "1"
	ApplVerID_FIX40    = "2"
	ApplVerID_FIX41    = "3"
	ApplVerID_FIX42    = "4"
	ApplVerID_FIX43    = "5"
	ApplVerID_FIX44    = "6"
	ApplVerID_FIX50    = "7"
	ApplVerID_FIX50SP1 = "8"
	ApplVerID_FIX50SP2 = "9"
)

//Enum values for AsOfIndicator
const (
	AsOfIndicator_FALSE = "0"
	AsOfIndicator_TRUE  = "1"
)

//Enum values for AssignmentMethod
const (
	AssignmentMethod_PRO_RATA = "P"
	AssignmentMethod_RANDOM   = "R"
)

//Enum values for AvgPxIndicator
const (
	AvgPxIndicator_NO_AVERAGE_PRICING                                                    = "0"
	AvgPxIndicator_TRADE_IS_PART_OF_AN_AVERAGE_PRICE_GROUP_IDENTIFIED_BY_THE_TRADELINKID = "1"
	AvgPxIndicator_LAST_TRADE_IS_THE_AVERAGE_PRICE_GROUP_IDENTIFIED_BY_THE_TRADELINKID   = "2"
)

//Enum values for BasisPxType
const (
	BasisPxType_CLOSING_PRICE_AT_MORNINGN_SESSION             = "2"
	BasisPxType_CLOSING_PRICE                                 = "3"
	BasisPxType_CURRENT_PRICE                                 = "4"
	BasisPxType_SQ                                            = "5"
	BasisPxType_VWAP_THROUGH_A_DAY                            = "6"
	BasisPxType_VWAP_THROUGH_A_MORNING_SESSION                = "7"
	BasisPxType_VWAP_THROUGH_AN_AFTERNOON_SESSION             = "8"
	BasisPxType_VWAP_THROUGH_A_DAY_EXCEPT_YORI                = "9"
	BasisPxType_VWAP_THROUGH_A_MORNING_SESSION_EXCEPT_YORI    = "A"
	BasisPxType_VWAP_THROUGH_AN_AFTERNOON_SESSION_EXCEPT_YORI = "B"
	BasisPxType_STRIKE                                        = "C"
	BasisPxType_OPEN                                          = "D"
	BasisPxType_OTHERS                                        = "Z"
)

//Enum values for Benchmark
const (
	Benchmark_CURVE    = "1"
	Benchmark_5YR      = "2"
	Benchmark_OLD5     = "3"
	Benchmark_10YR     = "4"
	Benchmark_OLD10    = "5"
	Benchmark_30YR     = "6"
	Benchmark_OLD30    = "7"
	Benchmark_3MOLIBOR = "8"
	Benchmark_6MOLIBOR = "9"
)

//Enum values for BenchmarkCurveName
const (
	BenchmarkCurveName_EONIA       = "EONIA"
	BenchmarkCurveName_EUREPO      = "EUREPO"
	BenchmarkCurveName_EURIBOR     = "Euribor"
	BenchmarkCurveName_FUTURESWAP  = "FutureSWAP"
	BenchmarkCurveName_LIBID       = "LIBID"
	BenchmarkCurveName_LIBOR       = "LIBOR"
	BenchmarkCurveName_MUNIAAA     = "MuniAAA"
	BenchmarkCurveName_OTHER       = "OTHER"
	BenchmarkCurveName_PFANDBRIEFE = "Pfandbriefe"
	BenchmarkCurveName_SONIA       = "SONIA"
	BenchmarkCurveName_SWAP        = "SWAP"
	BenchmarkCurveName_TREASURY    = "Treasury"
)

//Enum values for BidDescriptorType
const (
	BidDescriptorType_SECTOR  = "1"
	BidDescriptorType_COUNTRY = "2"
	BidDescriptorType_INDEX   = "3"
)

//Enum values for BidRequestTransType
const (
	BidRequestTransType_CANCEL = "C"
	BidRequestTransType_NO     = "N"
)

//Enum values for BidTradeType
const (
	BidTradeType_AGENCY           = "A"
	BidTradeType_VWAP_GUARANTEE   = "G"
	BidTradeType_GUARANTEED_CLOSE = "J"
	BidTradeType_RISK_TRADE       = "R"
)

//Enum values for BidType
const (
	BidType_NON_DISCLOSED_STYLE = "1"
	BidType_DISCLOSED_SYTLE     = "2"
	BidType_NO_BIDDING_PROCESS  = "3"
)

//Enum values for BookingType
const (
	BookingType_REGULAR_BOOKING   = "0"
	BookingType_CFD               = "1"
	BookingType_TOTAL_RETURN_SWAP = "2"
)

//Enum values for BookingUnit
const (
	BookingUnit_EACH_PARTIAL_EXECUTION_IS_A_BOOKABLE_UNIT                               = "0"
	BookingUnit_AGGREGATE_PARTIAL_EXECUTIONS_ON_THIS_ORDER_AND_BOOK_ONE_TRADE_PER_ORDER = "1"
	BookingUnit_AGGREGATE_EXECUTIONS_FOR_THIS_SYMBOL_SIDE_AND_SETTLEMENT_DATE           = "2"
)

//Enum values for BusinessRejectReason
const (
	BusinessRejectReason_OTHER                                     = "0"
	BusinessRejectReason_UNKNOWN_ID                                = "1"
	BusinessRejectReason_INVALID_PRICE_INCREMENT                   = "18"
	BusinessRejectReason_UNKNOWN_SECURITY                          = "2"
	BusinessRejectReason_UNSUPPORTED_MESSAGE_TYPE                  = "3"
	BusinessRejectReason_APPLICATION_NOT_AVAILABLE                 = "4"
	BusinessRejectReason_CONDITIONALLY_REQUIRED_FIELD_MISSING      = "5"
	BusinessRejectReason_NOT_AUTHORIZED                            = "6"
	BusinessRejectReason_DELIVERTO_FIRM_NOT_AVAILABLE_AT_THIS_TIME = "7"
)

//Enum values for CPProgram
const (
	CPProgram_3     = "1"
	CPProgram_4     = "2"
	CPProgram_OTHER = "99"
)

//Enum values for CancellationRights
const (
	CancellationRights_NO_M = "M"
	CancellationRights_NO_N = "N"
	CancellationRights_NO_O = "O"
	CancellationRights_YES  = "Y"
)

//Enum values for CashMargin
const (
	CashMargin_CASH         = "1"
	CashMargin_MARGIN_OPEN  = "2"
	CashMargin_MARGIN_CLOSE = "3"
)

//Enum values for ClearingFeeIndicator
const (
	ClearingFeeIndicator_1ST_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              = "1"
	ClearingFeeIndicator_2ND_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              = "2"
	ClearingFeeIndicator_3RD_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              = "3"
	ClearingFeeIndicator_4TH_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              = "4"
	ClearingFeeIndicator_5TH_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              = "5"
	ClearingFeeIndicator_6TH_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              = "9"
	ClearingFeeIndicator_CBOE_MEMBER                                                            = "B"
	ClearingFeeIndicator_NON_MEMBER_AND_CUSTOMER                                                = "C"
	ClearingFeeIndicator_EQUITY_MEMBER_AND_CLEARING_MEMBER                                      = "E"
	ClearingFeeIndicator_FULL_AND_ASSOCIATE_MEMBER_TRADING_FOR_OWN_ACCOUNT_AND_AS_FLOOR_BROKERS = "F"
	ClearingFeeIndicator_106H_AND_106J_FIRMS                                                    = "H"
	ClearingFeeIndicator_GIM_IDEM_AND_COM_MEMBERSHIP_INTEREST_HOLDERS                           = "I"
	ClearingFeeIndicator_LESSEE_106F_EMPLOYEES                                                  = "L"
	ClearingFeeIndicator_ALL_OTHER_OWNERSHIP_TYPES                                              = "M"
)

//Enum values for ClearingInstruction
const (
	ClearingInstruction_PROCESS_NORMALLY                     = "0"
	ClearingInstruction_EXCLUDE_FROM_ALL_NETTING             = "1"
	ClearingInstruction_AUTOMATIC_GIVE_UP_MODE               = "10"
	ClearingInstruction_QUALIFIED_SERVICE_REPRESENTATIVE_QSR = "11"
	ClearingInstruction_CUSTOMER_TRADE                       = "12"
	ClearingInstruction_SELF_CLEARING                        = "13"
	ClearingInstruction_BILATERAL_NETTING_ONLY               = "2"
	ClearingInstruction_EX_CLEARING                          = "3"
	ClearingInstruction_SPECIAL_TRADE                        = "4"
	ClearingInstruction_MULTILATERAL_NETTING                 = "5"
	ClearingInstruction_CLEAR_AGAINST_CENTRAL_COUNTERPARTY   = "6"
	ClearingInstruction_EXCLUDE_FROM_CENTRAL_COUNTERPARTY    = "7"
	ClearingInstruction_MANUAL_MODE                          = "8"
	ClearingInstruction_AUTOMATIC_POSTING_MODE               = "9"
)

//Enum values for CollAction
const (
	CollAction_RETAIN = "0"
	CollAction_ADD    = "1"
	CollAction_REMOVE = "2"
)

//Enum values for CollApplType
const (
	CollApplType_SPECIFIC_DEPOSIT = "0"
	CollApplType_GENERAL          = "1"
)

//Enum values for CollAsgnReason
const (
	CollAsgnReason_INITIAL                   = "0"
	CollAsgnReason_SCHEDULED                 = "1"
	CollAsgnReason_TIME_WARNING              = "2"
	CollAsgnReason_MARGIN_DEFICIENCY         = "3"
	CollAsgnReason_MARGIN_EXCESS             = "4"
	CollAsgnReason_FORWARD_COLLATERAL_DEMAND = "5"
	CollAsgnReason_EVENT_OF_DEFAULT          = "6"
	CollAsgnReason_ADVERSE_TAX_EVENT         = "7"
)

//Enum values for CollAsgnRejectReason
const (
	CollAsgnRejectReason_UNKNOWN_DEAL                  = "0"
	CollAsgnRejectReason_UNKNOWN_OR_INVALID_INSTRUMENT = "1"
	CollAsgnRejectReason_UNAUTHORIZED_TRANSACTION      = "2"
	CollAsgnRejectReason_INSUFFICIENT_COLLATERAL       = "3"
	CollAsgnRejectReason_INVALID_TYPE_OF_COLLATERAL    = "4"
	CollAsgnRejectReason_EXCESSIVE_SUBSTITUTION        = "5"
	CollAsgnRejectReason_OTHER                         = "99"
)

//Enum values for CollAsgnRespType
const (
	CollAsgnRespType_RECEIVED = "0"
	CollAsgnRespType_ACCEPTED = "1"
	CollAsgnRespType_DECLINED = "2"
	CollAsgnRespType_REJECTED = "3"
)

//Enum values for CollAsgnTransType
const (
	CollAsgnTransType_NEW     = "0"
	CollAsgnTransType_REPLACE = "1"
	CollAsgnTransType_CANCEL  = "2"
	CollAsgnTransType_RELEASE = "3"
	CollAsgnTransType_REVERSE = "4"
)

//Enum values for CollInquiryQualifier
const (
	CollInquiryQualifier_TRADE_DATE            = "0"
	CollInquiryQualifier_GC_INSTRUMENT         = "1"
	CollInquiryQualifier_COLLATERAL_INSTRUMENT = "2"
	CollInquiryQualifier_SUBSTITUTION_ELIGIBLE = "3"
	CollInquiryQualifier_NOT_ASSIGNED          = "4"
	CollInquiryQualifier_PARTIALLY_ASSIGNED    = "5"
	CollInquiryQualifier_FULLY_ASSIGNED        = "6"
	CollInquiryQualifier_OUTSTANDING_TRADES    = "7"
)

//Enum values for CollInquiryResult
const (
	CollInquiryResult_SUCCESSFUL                                  = "0"
	CollInquiryResult_INVALID_OR_UNKNOWN_INSTRUMENT               = "1"
	CollInquiryResult_INVALID_OR_UNKNOWN_COLLATERAL_TYPE          = "2"
	CollInquiryResult_INVALID_PARTIES                             = "3"
	CollInquiryResult_INVALID_TRANSPORT_TYPE_REQUESTED            = "4"
	CollInquiryResult_INVALID_DESTINATION_REQUESTED               = "5"
	CollInquiryResult_NO_COLLATERAL_FOUND_FOR_THE_TRADE_SPECIFIED = "6"
	CollInquiryResult_NO_COLLATERAL_FOUND_FOR_THE_ORDER_SPECIFIED = "7"
	CollInquiryResult_COLLATERAL_INQUIRY_TYPE_NOT_SUPPORTED       = "8"
	CollInquiryResult_UNAUTHORIZED_FOR_COLLATERAL_INQUIRY         = "9"
	CollInquiryResult_OTHER                                       = "99"
)

//Enum values for CollInquiryStatus
const (
	CollInquiryStatus_ACCEPTED                = "0"
	CollInquiryStatus_ACCEPTED_WITH_WARNINGS  = "1"
	CollInquiryStatus_COMPLETED               = "2"
	CollInquiryStatus_COMPLETED_WITH_WARNINGS = "3"
	CollInquiryStatus_REJECTED                = "4"
)

//Enum values for CollStatus
const (
	CollStatus_UNASSIGNED          = "0"
	CollStatus_PARTIALLY_ASSIGNED  = "1"
	CollStatus_ASSIGNMENT_PROPOSED = "2"
	CollStatus_ASSIGNED            = "3"
	CollStatus_CHALLENGED          = "4"
)

//Enum values for CommType
const (
	CommType_PER_UNIT                    = "1"
	CommType_PERCENT                     = "2"
	CommType_ABSOLUTE                    = "3"
	CommType_PERCENTAGE_WAIVED_4         = "4"
	CommType_PERCENTAGE_WAIVED_5         = "5"
	CommType_POINTS_PER_BOND_OR_CONTRACT = "6"
)

//Enum values for ComplexEventCondition
const (
	ComplexEventCondition_AND = "1"
	ComplexEventCondition_OR  = "2"
)

//Enum values for ComplexEventPriceBoundaryMethod
const (
	ComplexEventPriceBoundaryMethod_LESS_THAN_COMPLEXEVENTPRICE                = "1"
	ComplexEventPriceBoundaryMethod_LESS_THAN_OR_EQUAL_TO_COMPLEXEVENTPRICE    = "2"
	ComplexEventPriceBoundaryMethod_EQUAL_TO_COMPLEXEVENTPRICE                 = "3"
	ComplexEventPriceBoundaryMethod_GREATER_THAN_OR_EQUAL_TO_COMPLEXEVENTPRICE = "4"
	ComplexEventPriceBoundaryMethod_GREATER_THAN_COMPLEXEVENTPRICE             = "5"
)

//Enum values for ComplexEventPriceTimeType
const (
	ComplexEventPriceTimeType_EXPIRATION          = "1"
	ComplexEventPriceTimeType_IMMEDIATE           = "2"
	ComplexEventPriceTimeType_SPECIFIED_DATE_TIME = "3"
)

//Enum values for ComplexEventType
const (
	ComplexEventType_CAPPED          = "1"
	ComplexEventType_TRIGGER         = "2"
	ComplexEventType_KNOCK_IN_UP     = "3"
	ComplexEventType_KOCK_IN_DOWN    = "4"
	ComplexEventType_KNOCK_OUT_UP    = "5"
	ComplexEventType_KNOCK_OUT_DOWN  = "6"
	ComplexEventType_UNDERLYING      = "7"
	ComplexEventType_RESET_BARRIER   = "8"
	ComplexEventType_ROLLING_BARRIER = "9"
)

//Enum values for ConfirmRejReason
const (
	ConfirmRejReason_MISMATCHED_ACCOUNT              = "1"
	ConfirmRejReason_MISSING_SETTLEMENT_INSTRUCTIONS = "2"
	ConfirmRejReason_OTHER                           = "99"
)

//Enum values for ConfirmStatus
const (
	ConfirmStatus_RECEIVED                        = "1"
	ConfirmStatus_MISMATCHED_ACCOUNT              = "2"
	ConfirmStatus_MISSING_SETTLEMENT_INSTRUCTIONS = "3"
	ConfirmStatus_CONFIRMED                       = "4"
	ConfirmStatus_REQUEST_REJECTED                = "5"
)

//Enum values for ConfirmTransType
const (
	ConfirmTransType_NEW     = "0"
	ConfirmTransType_REPLACE = "1"
	ConfirmTransType_CANCEL  = "2"
)

//Enum values for ConfirmType
const (
	ConfirmType_STATUS                        = "1"
	ConfirmType_CONFIRMATION                  = "2"
	ConfirmType_CONFIRMATION_REQUEST_REJECTED = "3"
)

//Enum values for ContAmtType
const (
	ContAmtType_COMMISSION_AMOUNT                       = "1"
	ContAmtType_EXIT_CHARGE_PERCENT                     = "10"
	ContAmtType_FUND_BASED_RENEWAL_COMMISSION_PERCENT   = "11"
	ContAmtType_PROJECTED_FUND_VALUE                    = "12"
	ContAmtType_FUND_BASED_RENEWAL_COMMISSION_AMOUNT_13 = "13"
	ContAmtType_FUND_BASED_RENEWAL_COMMISSION_AMOUNT_14 = "14"
	ContAmtType_NET_SETTLEMENT_AMOUNT                   = "15"
	ContAmtType_COMMISSION_PERCENT                      = "2"
	ContAmtType_INITIAL_CHARGE_AMOUNT                   = "3"
	ContAmtType_INITIAL_CHARGE_PERCENT                  = "4"
	ContAmtType_DISCOUNT_AMOUNT                         = "5"
	ContAmtType_DISCOUNT_PERCENT                        = "6"
	ContAmtType_DILUTION_LEVY_AMOUNT                    = "7"
	ContAmtType_DILUTION_LEVY_PERCENT                   = "8"
	ContAmtType_EXIT_CHARGE_AMOUNT                      = "9"
)

//Enum values for ContingencyType
const (
	ContingencyType_ONE_CANCELS_THE_OTHER   = "1"
	ContingencyType_ONE_TRIGGERS_THE_OTHER  = "2"
	ContingencyType_ONE_UPDATES_THE_OTHER_3 = "3"
	ContingencyType_ONE_UPDATES_THE_OTHER_4 = "4"
)

//Enum values for ContractMultiplierUnit
const (
	ContractMultiplierUnit_SHARES = "0"
	ContractMultiplierUnit_HOURS  = "1"
	ContractMultiplierUnit_DAYS   = "2"
)

//Enum values for CorporateAction
const (
	CorporateAction_EX_DIVIDEND                  = "A"
	CorporateAction_EX_DISTRIBUTION              = "B"
	CorporateAction_EX_RIGHTS                    = "C"
	CorporateAction_NEW                          = "D"
	CorporateAction_EX_INTEREST                  = "E"
	CorporateAction_CASH_DIVIDEND                = "F"
	CorporateAction_STOCK_DIVIDEND               = "G"
	CorporateAction_NON_INTEGER_STOCK_SPLIT      = "H"
	CorporateAction_REVERSE_STOCK_SPLIT          = "I"
	CorporateAction_STANDARD_INTEGER_STOCK_SPLIT = "J"
	CorporateAction_POSITION_CONSOLIDATION       = "K"
	CorporateAction_LIQUIDATION_REORGANIZATION   = "L"
	CorporateAction_MERGER_REORGANIZATION        = "M"
	CorporateAction_RIGHTS_OFFERING              = "N"
	CorporateAction_SHAREHOLDER_MEETING          = "O"
	CorporateAction_SPINOFF                      = "P"
	CorporateAction_TENDER_OFFER                 = "Q"
	CorporateAction_WARRANT                      = "R"
	CorporateAction_SPECIAL_ACTION               = "S"
	CorporateAction_SYMBOL_CONVERSION            = "T"
	CorporateAction_CUSIP                        = "U"
	CorporateAction_LEAP_ROLLOVER                = "V"
	CorporateAction_SUCCESSION_EVENT             = "W"
)

//Enum values for CoveredOrUncovered
const (
	CoveredOrUncovered_COVERED   = "0"
	CoveredOrUncovered_UNCOVERED = "1"
)

//Enum values for CrossPrioritization
const (
	CrossPrioritization_NONE                     = "0"
	CrossPrioritization_BUY_SIDE_IS_PRIORITIZED  = "1"
	CrossPrioritization_SELL_SIDE_IS_PRIORITIZED = "2"
)

//Enum values for CrossType
const (
	CrossType_CROSS_AON        = "1"
	CrossType_CROSS_IOC        = "2"
	CrossType_CROSS_ONE_SIDE   = "3"
	CrossType_CROSS_SAME_PRICE = "4"
)

//Enum values for CustOrderCapacity
const (
	CustOrderCapacity_MEMBER_TRADING_FOR_THEIR_OWN_ACCOUNT              = "1"
	CustOrderCapacity_CLEARING_FIRM_TRADING_FOR_ITS_PROPRIETARY_ACCOUNT = "2"
	CustOrderCapacity_MEMBER_TRADING_FOR_ANOTHER_MEMBER                 = "3"
	CustOrderCapacity_ALL_OTHER                                         = "4"
)

//Enum values for CustOrderHandlingInst
const (
	CustOrderHandlingInst_ADD_ON_ORDER                      = "ADD"
	CustOrderHandlingInst_ALL_OR_NONE                       = "AON"
	CustOrderHandlingInst_CASH_NOT_HELD                     = "CNH"
	CustOrderHandlingInst_DIRECTED_ORDER                    = "DIR"
	CustOrderHandlingInst_EXCHANGE_FOR_PHYSICAL_TRANSACTION = "E.W"
	CustOrderHandlingInst_FILL_OR_KILL                      = "FOK"
	CustOrderHandlingInst_IMBALANCE_ONLY                    = "IO"
	CustOrderHandlingInst_IMMEDIATE_OR_CANCEL               = "IOC"
	CustOrderHandlingInst_LIMIT_ON_CLOSE                    = "LOC"
	CustOrderHandlingInst_LIMIT_ON_OPEN                     = "LOO"
	CustOrderHandlingInst_MARKET_AT_CLOSE                   = "MAC"
	CustOrderHandlingInst_MARKET_AT_OPEN                    = "MAO"
	CustOrderHandlingInst_MARKET_ON_CLOSE                   = "MOC"
	CustOrderHandlingInst_MARKET_ON_OPEN                    = "MOO"
	CustOrderHandlingInst_MINIMUM_QUANTITY                  = "MQT"
	CustOrderHandlingInst_NOT_HELD                          = "NH"
	CustOrderHandlingInst_OVER_THE_DAY                      = "OVD"
	CustOrderHandlingInst_PEGGED                            = "PEG"
	CustOrderHandlingInst_RESERVE_SIZE_ORDER                = "RSV"
	CustOrderHandlingInst_STOP_STOCK_TRANSACTION            = "S.W"
	CustOrderHandlingInst_SCALE                             = "SCL"
	CustOrderHandlingInst_TIME_ORDER                        = "TMO"
	CustOrderHandlingInst_TRAILING_STOP                     = "TS"
	CustOrderHandlingInst_WORK                              = "WRK"
)

//Enum values for CustomerOrFirm
const (
	CustomerOrFirm_CUSTOMER = "0"
	CustomerOrFirm_FIRM     = "1"
)

//Enum values for CxlRejReason
const (
	CxlRejReason_TOO_LATE_TO_CANCEL                                        = "0"
	CxlRejReason_UNKNOWN_ORDER                                             = "1"
	CxlRejReason_INVALID_PRICE_INCREMENT                                   = "18"
	CxlRejReason_BROKER                                                    = "2"
	CxlRejReason_ORDER_ALREADY_IN_PENDING_CANCEL_OR_PENDING_REPLACE_STATUS = "3"
	CxlRejReason_UNABLE_TO_PROCESS_ORDER_MASS_CANCEL_REQUEST               = "4"
	CxlRejReason_ORIGORDMODTIME                                            = "5"
	CxlRejReason_DUPLICATE_CLORDID                                         = "6"
	CxlRejReason_PRICE_EXCEEDS_CURRENT_PRICE                               = "7"
	CxlRejReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND                          = "8"
	CxlRejReason_OTHER                                                     = "99"
)

//Enum values for CxlRejResponseTo
const (
	CxlRejResponseTo_ORDER_CANCEL_REQUEST         = "1"
	CxlRejResponseTo_ORDER_CANCEL_REPLACE_REQUEST = "2"
)

//Enum values for CxlType
const (
	CxlType_FULL_REMAINING_QUANTITY = "F"
	CxlType_PARTIAL_CANCEL          = "P"
)

//Enum values for DKReason
const (
	DKReason_UNKNOWN_SYMBOL         = "A"
	DKReason_WRONG_SIDE             = "B"
	DKReason_QUANTITY_EXCEEDS_ORDER = "C"
	DKReason_NO_MATCHING_ORDER      = "D"
	DKReason_PRICE_EXCEEDS_LIMIT    = "E"
	DKReason_CALCULATION_DIFFERENCE = "F"
	DKReason_OTHER                  = "Z"
)

//Enum values for DayBookingInst
const (
	DayBookingInst_CAN_TRIGGER_BOOKING_WITHOUT_REFERENCE_TO_THE_ORDER_INITIATOR = "0"
	DayBookingInst_SPEAK_WITH_ORDER_INITIATOR_BEFORE_BOOKING                    = "1"
	DayBookingInst_ACCUMULATE                                                   = "2"
)

//Enum values for DealingCapacity
const (
	DealingCapacity_AGENT              = "A"
	DealingCapacity_PRINCIPAL          = "P"
	DealingCapacity_RISKLESS_PRINCIPAL = "R"
)

//Enum values for DeleteReason
const (
	DeleteReason_CANCELLATION = "0"
	DeleteReason_ERROR        = "1"
)

//Enum values for DeliveryForm
const (
	DeliveryForm_BOOK_ENTRY = "1"
	DeliveryForm_BEARER     = "2"
)

//Enum values for DeliveryType
const (
	DeliveryType_VERSUS_PAYMENT_DELIVER = "0"
	DeliveryType_FREE_DELIVER           = "1"
	DeliveryType_TRI_PARTY              = "2"
	DeliveryType_HOLD_IN_CUSTODY        = "3"
)

//Enum values for DerivativeSecurityListRequestType
const (
	DerivativeSecurityListRequestType_SYMBOL                                    = "0"
	DerivativeSecurityListRequestType_SECURITYTYPE_AND_OR_CFICODE               = "1"
	DerivativeSecurityListRequestType_PRODUCT                                   = "2"
	DerivativeSecurityListRequestType_TRADINGSESSIONID                          = "3"
	DerivativeSecurityListRequestType_ALL_SECURITIES                            = "4"
	DerivativeSecurityListRequestType_UNDELYINGSYMBOL                           = "5"
	DerivativeSecurityListRequestType_UNDERLYING_SECURITYTYPE_AND_OR_CFICODE    = "6"
	DerivativeSecurityListRequestType_UNDERLYING_PRODUCT                        = "7"
	DerivativeSecurityListRequestType_MARKETID_OR_MARKETID_PLUS_MARKETSEGMENTID = "8"
)

//Enum values for DeskOrderHandlingInst
const (
	DeskOrderHandlingInst_ADD_ON_ORDER                      = "ADD"
	DeskOrderHandlingInst_ALL_OR_NONE                       = "AON"
	DeskOrderHandlingInst_CASH_NOT_HELD                     = "CNH"
	DeskOrderHandlingInst_DIRECTED_ORDER                    = "DIR"
	DeskOrderHandlingInst_EXCHANGE_FOR_PHYSICAL_TRANSACTION = "E.W"
	DeskOrderHandlingInst_FILL_OR_KILL                      = "FOK"
	DeskOrderHandlingInst_IMBALANCE_ONLY                    = "IO"
	DeskOrderHandlingInst_IMMEDIATE_OR_CANCEL               = "IOC"
	DeskOrderHandlingInst_LIMIT_ON_CLOSE                    = "LOC"
	DeskOrderHandlingInst_LIMIT_ON_OPEN                     = "LOO"
	DeskOrderHandlingInst_MARKET_AT_CLOSE                   = "MAC"
	DeskOrderHandlingInst_MARKET_AT_OPEN                    = "MAO"
	DeskOrderHandlingInst_MARKET_ON_CLOSE                   = "MOC"
	DeskOrderHandlingInst_MARKET_ON_OPEN                    = "MOO"
	DeskOrderHandlingInst_MINIMUM_QUANTITY                  = "MQT"
	DeskOrderHandlingInst_NOT_HELD                          = "NH"
	DeskOrderHandlingInst_OVER_THE_DAY                      = "OVD"
	DeskOrderHandlingInst_PEGGED                            = "PEG"
	DeskOrderHandlingInst_RESERVE_SIZE_ORDER                = "RSV"
	DeskOrderHandlingInst_STOP_STOCK_TRANSACTION            = "S.W"
	DeskOrderHandlingInst_SCALE                             = "SCL"
	DeskOrderHandlingInst_TIME_ORDER                        = "TMO"
	DeskOrderHandlingInst_TRAILING_STOP                     = "TS"
	DeskOrderHandlingInst_WORK                              = "WRK"
)

//Enum values for DeskType
const (
	DeskType_AGENCY            = "A"
	DeskType_ARBITRAGE         = "AR"
	DeskType_DERIVATIVES       = "D"
	DeskType_INTERNATIONAL     = "IN"
	DeskType_INSTITUTIONAL     = "IS"
	DeskType_OTHER             = "O"
	DeskType_PREFERRED_TRADING = "PF"
	DeskType_PROPRIETARY       = "PR"
	DeskType_PROGRAM_TRADING   = "PT"
	DeskType_SALES             = "S"
	DeskType_TRADING           = "T"
)

//Enum values for DeskTypeSource
const (
	DeskTypeSource_NASD_OATS = "1"
)

//Enum values for DiscretionInst
const (
	DiscretionInst_RELATED_TO_DISPLAYED_PRICE     = "0"
	DiscretionInst_RELATED_TO_MARKET_PRICE        = "1"
	DiscretionInst_RELATED_TO_PRIMARY_PRICE       = "2"
	DiscretionInst_RELATED_TO_LOCAL_PRIMARY_PRICE = "3"
	DiscretionInst_RELATED_TO_MIDPOINT_PRICE      = "4"
	DiscretionInst_RELATED_TO_LAST_TRADE_PRICE    = "5"
	DiscretionInst_RELATED_TO_VWAP                = "6"
	DiscretionInst_AVERAGE_PRICE_GUARANTEE        = "7"
)

//Enum values for DiscretionLimitType
const (
	DiscretionLimitType_OR_BETTER = "0"
	DiscretionLimitType_STRICT    = "1"
	DiscretionLimitType_OR_WORSE  = "2"
)

//Enum values for DiscretionMoveType
const (
	DiscretionMoveType_FLOATING = "0"
	DiscretionMoveType_FIXED    = "1"
)

//Enum values for DiscretionOffsetType
const (
	DiscretionOffsetType_PRICE        = "0"
	DiscretionOffsetType_BASIS_POINTS = "1"
	DiscretionOffsetType_TICKS        = "2"
	DiscretionOffsetType_PRICE_TIER   = "3"
)

//Enum values for DiscretionRoundDirection
const (
	DiscretionRoundDirection_MORE_AGGRESSIVE = "1"
	DiscretionRoundDirection_MORE_PASSIVE    = "2"
)

//Enum values for DiscretionScope
const (
	DiscretionScope_LOCAL                    = "1"
	DiscretionScope_NATIONAL                 = "2"
	DiscretionScope_GLOBAL                   = "3"
	DiscretionScope_NATIONAL_EXCLUDING_LOCAL = "4"
)

//Enum values for DisplayMethod
const (
	DisplayMethod_INITIAL     = "1"
	DisplayMethod_NEW         = "2"
	DisplayMethod_RANDOM      = "3"
	DisplayMethod_UNDISCLOSED = "4"
)

//Enum values for DisplayWhen
const (
	DisplayWhen_IMMEDIATE = "1"
	DisplayWhen_EXHAUST   = "2"
)

//Enum values for DistribPaymentMethod
const (
	DistribPaymentMethod_CREST                            = "1"
	DistribPaymentMethod_BPAY                             = "10"
	DistribPaymentMethod_HIGH_VALUE_CLEARING_SYSTEM_HVACS = "11"
	DistribPaymentMethod_REINVEST_IN_FUND                 = "12"
	DistribPaymentMethod_NSCC                             = "2"
	DistribPaymentMethod_EUROCLEAR                        = "3"
	DistribPaymentMethod_CLEARSTREAM                      = "4"
	DistribPaymentMethod_CHEQUE                           = "5"
	DistribPaymentMethod_TELEGRAPHIC_TRANSFER             = "6"
	DistribPaymentMethod_FED_WIRE                         = "7"
	DistribPaymentMethod_DIRECT_CREDIT                    = "8"
	DistribPaymentMethod_ACH_CREDIT                       = "9"
)

//Enum values for DlvyInstType
const (
	DlvyInstType_CASH       = "C"
	DlvyInstType_SECURITIES = "S"
)

//Enum values for DueToRelated
const (
	DueToRelated_NO  = "N"
	DueToRelated_YES = "Y"
)

//Enum values for EmailType
const (
	EmailType_NEW         = "0"
	EmailType_REPLY       = "1"
	EmailType_ADMIN_REPLY = "2"
)

//Enum values for EncryptMethod
const (
	EncryptMethod_NONE_OTHER  = "0"
	EncryptMethod_PKCS        = "1"
	EncryptMethod_DES         = "2"
	EncryptMethod_PKCS_DES    = "3"
	EncryptMethod_PGP_DES     = "4"
	EncryptMethod_PGP_DES_MD5 = "5"
	EncryptMethod_PEM_DES_MD5 = "6"
)

//Enum values for EventType
const (
	EventType_PUT                        = "1"
	EventType_SWAP_ROLL_DATE             = "10"
	EventType_SWAP_NEXT_START_DATE       = "11"
	EventType_SWAP_NEXT_ROLL_DATE        = "12"
	EventType_FIRST_DELIVERY_DATE        = "13"
	EventType_LAST_DELIVERY_DATE         = "14"
	EventType_INITIAL_INVENTORY_DUE_DATE = "15"
	EventType_FINAL_INVENTORY_DUE_DATE   = "16"
	EventType_FIRST_INTENT_DATE          = "17"
	EventType_LAST_INTENT_DATE           = "18"
	EventType_POSITION_REMOVAL_DATE      = "19"
	EventType_CALL                       = "2"
	EventType_TENDER                     = "3"
	EventType_SINKING_FUND_CALL          = "4"
	EventType_ACTIVATION                 = "5"
	EventType_INACTIVIATION              = "6"
	EventType_LAST_ELIGIBLE_TRADE_DATE   = "7"
	EventType_SWAP_START_DATE            = "8"
	EventType_SWAP_END_DATE              = "9"
	EventType_OTHER                      = "99"
)

//Enum values for ExDestination
const (
	ExDestination_NONE  = "0"
	ExDestination_POSIT = "4"
)

//Enum values for ExDestinationIDSource
const (
	ExDestinationIDSource_BIC                                              = "B"
	ExDestinationIDSource_GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER = "C"
	ExDestinationIDSource_PROPRIETARY                                      = "D"
	ExDestinationIDSource_ISO_COUNTRY_CODE                                 = "E"
	ExDestinationIDSource_MIC                                              = "G"
)

//Enum values for ExchangeForPhysical
const (
	ExchangeForPhysical_NO  = "N"
	ExchangeForPhysical_YES = "Y"
)

//Enum values for ExecAckStatus
const (
	ExecAckStatus_RECEIVED_NOT_YET_PROCESSED = "0"
	ExecAckStatus_ACCEPTED                   = "1"
	ExecAckStatus_DONT_KNOW                  = "2"
)

//Enum values for ExecInst
const (
	ExecInst_STAY_ON_OFFER_SIDE                                    = "0"
	ExecInst_NOT_HELD                                              = "1"
	ExecInst_WORK                                                  = "2"
	ExecInst_GO_ALONG                                              = "3"
	ExecInst_OVER_THE_DAY                                          = "4"
	ExecInst_HELD                                                  = "5"
	ExecInst_PARTICIPANT_DONT_INITIATE                             = "6"
	ExecInst_STRICT_SCALE                                          = "7"
	ExecInst_TRY_TO_SCALE                                          = "8"
	ExecInst_STAY_ON_BID_SIDE                                      = "9"
	ExecInst_NO_CROSS                                              = "A"
	ExecInst_OK_TO_CROSS                                           = "B"
	ExecInst_CALL_FIRST                                            = "C"
	ExecInst_PERCENT_OF_VOLUME                                     = "D"
	ExecInst_DO_NOT_INCREASE                                       = "E"
	ExecInst_DO_NOT_REDUCE                                         = "F"
	ExecInst_ALL_OR_NONE                                           = "G"
	ExecInst_REINSTATE_ON_SYSTEM_FAILURE                           = "H"
	ExecInst_INSTITUTIONS_ONLY                                     = "I"
	ExecInst_REINSTATE_ON_TRADING_HALT                             = "J"
	ExecInst_CANCEL_ON_TRADING_HALT                                = "K"
	ExecInst_LAST_PEG                                              = "L"
	ExecInst_MID_PRICE_PEG                                         = "M"
	ExecInst_NON_NEGOTIABLE                                        = "N"
	ExecInst_OPENING_PEG                                           = "O"
	ExecInst_MARKET_PEG                                            = "P"
	ExecInst_CANCEL_ON_SYSTEM_FAILURE                              = "Q"
	ExecInst_PRIMARY_PEG                                           = "R"
	ExecInst_SUSPEND                                               = "S"
	ExecInst_FIXED_PEG_TO_LOCAL_BEST_BID_OR_OFFER_AT_TIME_OF_ORDER = "T"
	ExecInst_CUSTOMER_DISPLAY_INSTRUCTION                          = "U"
	ExecInst_NETTING                                               = "V"
	ExecInst_PEG_TO_VWAP                                           = "W"
	ExecInst_TRADE_ALONG                                           = "X"
	ExecInst_TRY_TO_STOP                                           = "Y"
	ExecInst_CANCEL_IF_NOT_BEST                                    = "Z"
	ExecInst_TRAILING_STOP_PEG                                     = "a"
	ExecInst_STRICT_LIMIT                                          = "b"
	ExecInst_IGNORE_PRICE_VALIDITY_CHECKS                          = "c"
	ExecInst_PEG_TO_LIMIT_PRICE                                    = "d"
	ExecInst_WORK_TO_TARGET_STRATEGY                               = "e"
	ExecInst_INTERMARKET_SWEEP                                     = "f"
	ExecInst_EXTERNAL_ROUTING_ALLOWED                              = "g"
	ExecInst_EXTERNAL_ROUTING_NOT_ALLOWED                          = "h"
	ExecInst_IMBALANCE_ONLY                                        = "i"
	ExecInst_SINGLE_EXECUTION_REQUESTED_FOR_BLOCK_TRADE            = "j"
	ExecInst_BEST_EXECUTION                                        = "k"
	ExecInst_SUSPEND_ON_SYSTEM_FAILURE                             = "l"
	ExecInst_SUSPEND_ON_TRADING_HALT                               = "m"
	ExecInst_REINSTATE_ON_CONNECTION_LOSS                          = "n"
	ExecInst_CANCEL_ON_CONNECTION_LOSS                             = "o"
	ExecInst_SUSPEND_ON_CONNECTION_LOSS                            = "p"
	ExecInst_RELEASE_FROM_SUSPENSION                               = "q"
	ExecInst_EXECUTE_AS_DELTA_NEUTRAL_USING_VOLATILITY_PROVIDED    = "r"
	ExecInst_EXECUTE_AS_DURATION_NEUTRAL                           = "s"
	ExecInst_EXECUTE_AS_FX_NEUTRAL                                 = "t"
)

//Enum values for ExecPriceType
const (
	ExecPriceType_BID_PRICE                              = "B"
	ExecPriceType_CREATION_PRICE                         = "C"
	ExecPriceType_CREATION_PRICE_PLUS_ADJUSTMENT_PERCENT = "D"
	ExecPriceType_CREATION_PRICE_PLUS_ADJUSTMENT_AMOUNT  = "E"
	ExecPriceType_OFFER_PRICE                            = "O"
	ExecPriceType_OFFER_PRICE_MINUS_ADJUSTMENT_PERCENT   = "P"
	ExecPriceType_OFFER_PRICE_MINUS_ADJUSTMENT_AMOUNT    = "Q"
	ExecPriceType_SINGLE_PRICE                           = "S"
)

//Enum values for ExecRestatementReason
const (
	ExecRestatementReason_GT_CORPORATE_ACTION         = "0"
	ExecRestatementReason_GT_RENEWAL                  = "1"
	ExecRestatementReason_WAREHOUSE_RECAP             = "10"
	ExecRestatementReason_PEG_REFRESH                 = "11"
	ExecRestatementReason_VERBAL_CHANGE               = "2"
	ExecRestatementReason_REPRICING_OF_ORDER          = "3"
	ExecRestatementReason_BROKER_OPTION               = "4"
	ExecRestatementReason_PARTIAL_DECLINE_OF_ORDERQTY = "5"
	ExecRestatementReason_CANCEL_ON_TRADING_HALT      = "6"
	ExecRestatementReason_CANCEL_ON_SYSTEM_FAILURE    = "7"
	ExecRestatementReason_MARKET                      = "8"
	ExecRestatementReason_CANCELED_NOT_BEST           = "9"
	ExecRestatementReason_OTHER                       = "99"
)

//Enum values for ExecTransType
const (
	ExecTransType_NEW     = "0"
	ExecTransType_CANCEL  = "1"
	ExecTransType_CORRECT = "2"
	ExecTransType_STATUS  = "3"
)

//Enum values for ExecType
const (
	ExecType_NEW                                 = "0"
	ExecType_PARTIAL_FILL                        = "1"
	ExecType_FILL                                = "2"
	ExecType_DONE_FOR_DAY                        = "3"
	ExecType_CANCELED                            = "4"
	ExecType_REPLACED                            = "5"
	ExecType_PENDING_CANCEL                      = "6"
	ExecType_STOPPED                             = "7"
	ExecType_REJECTED                            = "8"
	ExecType_SUSPENDED                           = "9"
	ExecType_PENDING_NEW                         = "A"
	ExecType_CALCULATED                          = "B"
	ExecType_EXPIRED                             = "C"
	ExecType_RESTATED                            = "D"
	ExecType_PENDING_REPLACE                     = "E"
	ExecType_TRADE                               = "F"
	ExecType_TRADE_CORRECT                       = "G"
	ExecType_TRADE_CANCEL                        = "H"
	ExecType_ORDER_STATUS                        = "I"
	ExecType_TRADE_IN_A_CLEARING_HOLD            = "J"
	ExecType_TRADE_HAS_BEEN_RELEASED_TO_CLEARING = "K"
	ExecType_TRIGGERED_OR_ACTIVATED_BY_SYSTEM    = "L"
)

//Enum values for ExerciseMethod
const (
	ExerciseMethod_AUTOMATIC = "A"
	ExerciseMethod_MANUAL    = "M"
)

//Enum values for ExerciseStyle
const (
	ExerciseStyle_EUROPEAN = "0"
	ExerciseStyle_AMERICAN = "1"
	ExerciseStyle_BERMUDA  = "2"
)

//Enum values for ExpType
const (
	ExpType_AUTO_EXERCISE           = "1"
	ExpType_NON_AUTO_EXERCISE       = "2"
	ExpType_FINAL_WILL_BE_EXERCISED = "3"
	ExpType_CONTRARY_INTENTION      = "4"
	ExpType_DIFFERENCE              = "5"
)

//Enum values for ExpirationCycle
const (
	ExpirationCycle_EXPIRE_ON_TRADING_SESSION_CLOSE                                                = "0"
	ExpirationCycle_EXPIRE_ON_TRADING_SESSION_OPEN                                                 = "1"
	ExpirationCycle_TRADING_ELIGIBILITY_EXPIRATION_SPECIFIED_IN_THE_DATE_AND_TIME_FIELDS_EVENTDATE = "2"
)

//Enum values for ExpirationQtyType
const (
	ExpirationQtyType_AUTO_EXERCISE           = "1"
	ExpirationQtyType_NON_AUTO_EXERCISE       = "2"
	ExpirationQtyType_FINAL_WILL_BE_EXERCISED = "3"
	ExpirationQtyType_CONTRARY_INTENTION      = "4"
	ExpirationQtyType_DIFFERENCE              = "5"
)

//Enum values for FinancialStatus
const (
	FinancialStatus_BANKRUPT          = "1"
	FinancialStatus_PENDING_DELISTING = "2"
	FinancialStatus_RESTRICTED        = "3"
)

//Enum values for FlowScheduleType
const (
	FlowScheduleType_NERC_EASTERN_OFF_PEAK           = "0"
	FlowScheduleType_NERC_WESTERN_OFF_PEAK           = "1"
	FlowScheduleType_NERC_CALENDAR_ALL_DAYS_IN_MONTH = "2"
	FlowScheduleType_NERC_EASTERN_PEAK               = "3"
	FlowScheduleType_NERC_WESTERN_PEAK               = "4"
)

//Enum values for ForexReq
const (
	ForexReq_NO  = "N"
	ForexReq_YES = "Y"
)

//Enum values for FundRenewWaiv
const (
	FundRenewWaiv_NO  = "N"
	FundRenewWaiv_YES = "Y"
)

//Enum values for FuturesValuationMethod
const (
	FuturesValuationMethod_PREMIUM_STYLE                                  = "EQTY"
	FuturesValuationMethod_FUTURES_STYLE_MARK_TO_MARKET                   = "FUT"
	FuturesValuationMethod_FUTURES_STYLE_WITH_AN_ATTACHED_CASH_ADJUSTMENT = "FUTDA"
)

//Enum values for GTBookingInst
const (
	GTBookingInst_BOOK_OUT_ALL_TRADES_ON_DAY_OF_EXECUTION                 = "0"
	GTBookingInst_ACCUMULATE_EXECTUIONS_UNTIL_FORDER_IS_FILLED_OR_EXPIRES = "1"
	GTBookingInst_ACCUMULATE_UNTIL_VERBALLLY_NOTIFIED_OTHERWISE           = "2"
)

//Enum values for GapFillFlag
const (
	GapFillFlag_NO  = "N"
	GapFillFlag_YES = "Y"
)

//Enum values for HaltReasonChar
const (
	HaltReasonChar_NEWS_DISSEMINATION     = "D"
	HaltReasonChar_ORDER_INFLUX           = "E"
	HaltReasonChar_ORDER_IMBALANCE        = "I"
	HaltReasonChar_ADDITIONAL_INFORMATION = "M"
	HaltReasonChar_NEW_PENDING            = "P"
	HaltReasonChar_EQUIPMENT_CHANGEOVER   = "X"
)

//Enum values for HaltReasonInt
const (
	HaltReasonInt_NEWS_DISSEMINATION     = "0"
	HaltReasonInt_ORDER_INFLUX           = "1"
	HaltReasonInt_ORDER_IMBALANCE        = "2"
	HaltReasonInt_ADDITIONAL_INFORMATION = "3"
	HaltReasonInt_NEWS_PENDING           = "4"
	HaltReasonInt_EQUIPMENT_CHANGEOVER   = "5"
)

//Enum values for HandlInst
const (
	HandlInst_AUTOMATED_EXECUTION_ORDER_PRIVATE_NO_BROKER_INTERVENTION = "1"
	HandlInst_AUTOMATED_EXECUTION_ORDER_PUBLIC_BROKER_INTERVENTION_OK  = "2"
	HandlInst_MANUAL_ORDER_BEST_EXECUTION                              = "3"
)

//Enum values for IDSource
const (
	IDSource_CUSIP                         = "1"
	IDSource_SEDOL                         = "2"
	IDSource_QUIK                          = "3"
	IDSource_ISIN_NUMBER                   = "4"
	IDSource_RIC_CODE                      = "5"
	IDSource_ISO_CURRENCY_CODE             = "6"
	IDSource_ISO_COUNTRY_CODE              = "7"
	IDSource_EXCHANGE_SYMBOL               = "8"
	IDSource_CONSOLIDATED_TAPE_ASSOCIATION = "9"
)

//Enum values for IOINaturalFlag
const (
	IOINaturalFlag_NO  = "N"
	IOINaturalFlag_YES = "Y"
)

//Enum values for IOIOthSvc
const (
	IOIOthSvc_AUTEX  = "A"
	IOIOthSvc_BRIDGE = "B"
)

//Enum values for IOIQltyInd
const (
	IOIQltyInd_HIGH   = "H"
	IOIQltyInd_LOW    = "L"
	IOIQltyInd_MEDIUM = "M"
)

//Enum values for IOIQty
const (
	IOIQty_1000000000           = "0"
	IOIQty_LARGE                = "L"
	IOIQty_MEDIUM               = "M"
	IOIQty_SMALL                = "S"
	IOIQty_UNDISCLOSED_QUANTITY = "U"
)

//Enum values for IOIQualifier
const (
	IOIQualifier_ALL_OR_NONE          = "A"
	IOIQualifier_MARKET_ON_CLOSE      = "B"
	IOIQualifier_AT_THE_CLOSE         = "C"
	IOIQualifier_VWAP                 = "D"
	IOIQualifier_IN_TOUCH_WITH        = "I"
	IOIQualifier_LIMIT                = "L"
	IOIQualifier_MORE_BEHIND          = "M"
	IOIQualifier_AT_THE_OPEN          = "O"
	IOIQualifier_TAKING_A_POSITION    = "P"
	IOIQualifier_AT_THE_MARKET        = "Q"
	IOIQualifier_READY_TO_TRADE       = "R"
	IOIQualifier_PORTFOLIO_SHOWN      = "S"
	IOIQualifier_THROUGH_THE_DAY      = "T"
	IOIQualifier_VERSUS               = "V"
	IOIQualifier_INDICATION           = "W"
	IOIQualifier_CROSSING_OPPORTUNITY = "X"
	IOIQualifier_AT_THE_MIDPOINT      = "Y"
	IOIQualifier_PRE_OPEN             = "Z"
)

//Enum values for IOIShares
const (
	IOIShares_LARGE  = "L"
	IOIShares_MEDIUM = "M"
	IOIShares_SMALL  = "S"
)

//Enum values for IOITransType
const (
	IOITransType_CANCEL  = "C"
	IOITransType_NEW     = "N"
	IOITransType_REPLACE = "R"
)

//Enum values for ImpliedMarketIndicator
const (
	ImpliedMarketIndicator_NOT_IMPLIED                     = "0"
	ImpliedMarketIndicator_IMPLIED_IN                      = "1"
	ImpliedMarketIndicator_IMPLIED_OUT                     = "2"
	ImpliedMarketIndicator_BOTH_IMPLIED_IN_AND_IMPLIED_OUT = "3"
)

//Enum values for InViewOfCommon
const (
	InViewOfCommon_NO  = "N"
	InViewOfCommon_YES = "Y"
)

//Enum values for IncTaxInd
const (
	IncTaxInd_NET   = "1"
	IncTaxInd_GROSS = "2"
)

//Enum values for IndividualAllocType
const (
	IndividualAllocType_SUB_ALLOCATE           = "1"
	IndividualAllocType_THIRD_PARTY_ALLOCATION = "2"
)

//Enum values for InstrAttribType
const (
	InstrAttribType_FLAT                                                                        = "1"
	InstrAttribType_ORIGINAL_ISSUE_DISCOUNT                                                     = "10"
	InstrAttribType_CALLABLE_PUTTABLE                                                           = "11"
	InstrAttribType_ESCROWED_TO_MATURITY                                                        = "12"
	InstrAttribType_ESCROWED_TO_REDEMPTION_DATE                                                 = "13"
	InstrAttribType_PRE_REFUNDED                                                                = "14"
	InstrAttribType_IN_DEFAULT                                                                  = "15"
	InstrAttribType_UNRATED                                                                     = "16"
	InstrAttribType_TAXABLE                                                                     = "17"
	InstrAttribType_INDEXED                                                                     = "18"
	InstrAttribType_SUBJECT_TO_ALTERNATIVE_MINIMUM_TAX                                          = "19"
	InstrAttribType_ZERO_COUPON                                                                 = "2"
	InstrAttribType_ORIGINAL_ISSUE_DISCOUNT_PRICE_SUPPLY_PRICE_IN_THE_INSTRATTRIBVALUE          = "20"
	InstrAttribType_CALLABLE_BELOW_MATURITY_VALUE                                               = "21"
	InstrAttribType_CALLABLE_WITHOUT_NOTICE_BY_MAIL_TO_HOLDER_UNLESS_REGISTERED                 = "22"
	InstrAttribType_PRICE_TICK_RULES_FOR_SECURITY                                               = "23"
	InstrAttribType_TRADE_TYPE_ELIGIBILITY_DETAILS_FOR_SECURITY                                 = "24"
	InstrAttribType_INSTRUMENT_DENOMINATOR                                                      = "25"
	InstrAttribType_INSTRUMENT_NUMERATOR                                                        = "26"
	InstrAttribType_INSTRUMENT_PRICE_PRECISION                                                  = "27"
	InstrAttribType_INSTRUMENT_STRIKE_PRICE                                                     = "28"
	InstrAttribType_TRADEABLE_INDICATOR                                                         = "29"
	InstrAttribType_INTEREST_BEARING                                                            = "3"
	InstrAttribType_NO_PERIODIC_PAYMENTS                                                        = "4"
	InstrAttribType_VARIABLE_RATE                                                               = "5"
	InstrAttribType_LESS_FEE_FOR_PUT                                                            = "6"
	InstrAttribType_STEPPED_COUPON                                                              = "7"
	InstrAttribType_COUPON_PERIOD                                                               = "8"
	InstrAttribType_WHEN_AND_IF_ISSUED                                                          = "9"
	InstrAttribType_TEXT_SUPPLY_THE_TEXT_OF_THE_ATTRIBUTE_OR_DISCLAIMER_IN_THE_INSTRATTRIBVALUE = "99"
)

//Enum values for InstrRegistry
const (
	InstrRegistry_CUSTODIAN = "BIC"
	InstrRegistry_COUNTRY   = "ISO"
	InstrRegistry_PHYSICAL  = "ZZ"
)

//Enum values for LastCapacity
const (
	LastCapacity_AGENT              = "1"
	LastCapacity_CROSS_AS_AGENT     = "2"
	LastCapacity_CROSS_AS_PRINCIPAL = "3"
	LastCapacity_PRINCIPAL          = "4"
)

//Enum values for LastFragment
const (
	LastFragment_NO  = "N"
	LastFragment_YES = "Y"
)

//Enum values for LastLiquidityInd
const (
	LastLiquidityInd_ADDED_LIQUIDITY      = "1"
	LastLiquidityInd_REMOVED_LIQUIDITY    = "2"
	LastLiquidityInd_LIQUIDITY_ROUTED_OUT = "3"
	LastLiquidityInd_AUCTION              = "4"
)

//Enum values for LastRptRequested
const (
	LastRptRequested_NO  = "N"
	LastRptRequested_YES = "Y"
)

//Enum values for LegSwapType
const (
	LegSwapType_PAR_FOR_PAR       = "1"
	LegSwapType_MODIFIED_DURATION = "2"
	LegSwapType_RISK              = "4"
	LegSwapType_PROCEEDS          = "5"
)

//Enum values for LegalConfirm
const (
	LegalConfirm_NO  = "N"
	LegalConfirm_YES = "Y"
)

//Enum values for LiquidityIndType
const (
	LiquidityIndType_5_DAY_MOVING_AVERAGE  = "1"
	LiquidityIndType_20_DAY_MOVING_AVERAGE = "2"
	LiquidityIndType_NORMAL_MARKET_SIZE    = "3"
	LiquidityIndType_OTHER                 = "4"
)

//Enum values for ListExecInstType
const (
	ListExecInstType_IMMEDIATE                   = "1"
	ListExecInstType_WAIT_FOR_EXECUT_INSTRUCTION = "2"
	ListExecInstType_EXCHANGE_SWITCH_CIV_ORDER_3 = "3"
	ListExecInstType_EXCHANGE_SWITCH_CIV_ORDER_4 = "4"
	ListExecInstType_EXCHANGE_SWITCH_CIV_ORDER_5 = "5"
)

//Enum values for ListMethod
const (
	ListMethod_PRE_LISTED_ONLY = "0"
	ListMethod_USER_REQUESTED  = "1"
)

//Enum values for ListOrderStatus
const (
	ListOrderStatus_IN_BIDDING_PROCESS     = "1"
	ListOrderStatus_RECEIVED_FOR_EXECUTION = "2"
	ListOrderStatus_EXECUTING              = "3"
	ListOrderStatus_CANCELLING             = "4"
	ListOrderStatus_ALERT                  = "5"
	ListOrderStatus_ALL_DONE               = "6"
	ListOrderStatus_REJECT                 = "7"
)

//Enum values for ListRejectReason
const (
	ListRejectReason_BROKER                           = "0"
	ListRejectReason_UNSUPPORTED_ORDER_CHARACTERISTIC = "11"
	ListRejectReason_EXCHANGE_CLOSED                  = "2"
	ListRejectReason_TOO_LATE_TO_ENTER                = "4"
	ListRejectReason_UNKNOWN_ORDER                    = "5"
	ListRejectReason_DUPLICATE_ORDER                  = "6"
	ListRejectReason_OTHER                            = "99"
)

//Enum values for ListStatusType
const (
	ListStatusType_ACK          = "1"
	ListStatusType_RESPONSE     = "2"
	ListStatusType_TIMED        = "3"
	ListStatusType_EXEC_STARTED = "4"
	ListStatusType_ALL_DONE     = "5"
	ListStatusType_ALERT        = "6"
)

//Enum values for LocateReqd
const (
	LocateReqd_NO  = "N"
	LocateReqd_YES = "Y"
)

//Enum values for LotType
const (
	LotType_ODD_LOT                            = "1"
	LotType_ROUND_LOT                          = "2"
	LotType_BLOCK_LOT                          = "3"
	LotType_ROUND_LOT_BASED_UPON_UNITOFMEASURE = "4"
)

//Enum values for MDBookType
const (
	MDBookType_TOP_OF_BOOK = "1"
	MDBookType_PRICE_DEPTH = "2"
	MDBookType_ORDER_DEPTH = "3"
)

//Enum values for MDEntryType
const (
	MDEntryType_BID                                             = "0"
	MDEntryType_OFFER                                           = "1"
	MDEntryType_TRADE                                           = "2"
	MDEntryType_INDEX_VALUE                                     = "3"
	MDEntryType_OPENING_PRICE                                   = "4"
	MDEntryType_CLOSING_PRICE                                   = "5"
	MDEntryType_SETTLEMENT_PRICE                                = "6"
	MDEntryType_TRADING_SESSION_HIGH_PRICE                      = "7"
	MDEntryType_TRADING_SESSION_LOW_PRICE                       = "8"
	MDEntryType_TRADING_SESSION_VWAP_PRICE                      = "9"
	MDEntryType_IMBALANCE                                       = "A"
	MDEntryType_TRADE_VOLUME                                    = "B"
	MDEntryType_OPEN_INTEREST                                   = "C"
	MDEntryType_COMPOSITE_UNDERLYING_PRICE                      = "D"
	MDEntryType_SIMULATED_SELL_PRICE                            = "E"
	MDEntryType_SIMULATED_BUY_PRICE                             = "F"
	MDEntryType_MARGIN_RATE                                     = "G"
	MDEntryType_MID_PRICE                                       = "H"
	MDEntryType_EMPTY_BOOK                                      = "J"
	MDEntryType_SETTLE_HIGH_PRICE                               = "K"
	MDEntryType_SETTLE_LOW_PRICE                                = "L"
	MDEntryType_PRIOR_SETTLE_PRICE                              = "M"
	MDEntryType_SESSION_HIGH_BID                                = "N"
	MDEntryType_SESSION_LOW_OFFER                               = "O"
	MDEntryType_EARLY_PRICES                                    = "P"
	MDEntryType_AUCTION_CLEARING_PRICE                          = "Q"
	MDEntryType_DAILY_VALUE_ADJUSTMENT_FOR_LONG_POSITIONS       = "R"
	MDEntryType_SWAP_VALUE_FACTOR                               = "S"
	MDEntryType_CUMULATIVE_VALUE_ADJUSTMENT_FOR_LONG_POSITIONS  = "T"
	MDEntryType_DAILY_VALUE_ADJUSTMENT_FOR_SHORT_POSITIONS      = "U"
	MDEntryType_CUMULATIVE_VALUE_ADJUSTMENT_FOR_SHORT_POSITIONS = "V"
	MDEntryType_FIXING_PRICE                                    = "W"
	MDEntryType_CASH_RATE                                       = "X"
	MDEntryType_RECOVERY_RATE                                   = "Y"
	MDEntryType_RECOVERY_RATE_FOR_LONG                          = "Z"
	MDEntryType_RECOVERY_RATE_FOR_SHORT                         = "a"
)

//Enum values for MDImplicitDelete
const (
	MDImplicitDelete_NO  = "N"
	MDImplicitDelete_YES = "Y"
)

//Enum values for MDOriginType
const (
	MDOriginType_BOOK     = "0"
	MDOriginType_OFF_BOOK = "1"
	MDOriginType_CROSS    = "2"
)

//Enum values for MDQuoteType
const (
	MDQuoteType_INDICATIVE               = "0"
	MDQuoteType_TRADEABLE                = "1"
	MDQuoteType_RESTRICTED_TRADEABLE     = "2"
	MDQuoteType_COUNTER                  = "3"
	MDQuoteType_INDICATIVE_AND_TRADEABLE = "4"
)

//Enum values for MDReqRejReason
const (
	MDReqRejReason_UNKNOWN_SYMBOL                      = "0"
	MDReqRejReason_DUPLICATE_MDREQID                   = "1"
	MDReqRejReason_INSUFFICIENT_BANDWIDTH              = "2"
	MDReqRejReason_INSUFFICIENT_PERMISSIONS            = "3"
	MDReqRejReason_UNSUPPORTED_SUBSCRIPTIONREQUESTTYPE = "4"
	MDReqRejReason_UNSUPPORTED_MARKETDEPTH             = "5"
	MDReqRejReason_UNSUPPORTED_MDUPDATETYPE            = "6"
	MDReqRejReason_UNSUPPORTED_AGGREGATEDBOOK          = "7"
	MDReqRejReason_UNSUPPORTED_MDENTRYTYPE             = "8"
	MDReqRejReason_UNSUPPORTED_TRADINGSESSIONID        = "9"
	MDReqRejReason_UNSUPPORTED_SCOPE                   = "A"
	MDReqRejReason_UNSUPPORTED_OPENCLOSESETTLEFLAG     = "B"
	MDReqRejReason_UNSUPPORTED_MDIMPLICITDELETE        = "C"
	MDReqRejReason_INSUFFICIENT_CREDIT                 = "D"
)

//Enum values for MDSecSizeType
const (
	MDSecSizeType_CUSTOMER = "1"
)

//Enum values for MDUpdateAction
const (
	MDUpdateAction_NEW         = "0"
	MDUpdateAction_CHANGE      = "1"
	MDUpdateAction_DELETE      = "2"
	MDUpdateAction_DELETE_THRU = "3"
	MDUpdateAction_DELETE_FROM = "4"
	MDUpdateAction_OVERLAY     = "5"
)

//Enum values for MDUpdateType
const (
	MDUpdateType_FULL_REFRESH        = "0"
	MDUpdateType_INCREMENTAL_REFRESH = "1"
)

//Enum values for MarketUpdateAction
const (
	MarketUpdateAction_ADD    = "A"
	MarketUpdateAction_DELETE = "D"
	MarketUpdateAction_MODIFY = "M"
)

//Enum values for MassActionRejectReason
const (
	MassActionRejectReason_MASS_ACTION_NOT_SUPPORTED                        = "0"
	MassActionRejectReason_INVALID_OR_UNKNOWN_SECURITY                      = "1"
	MassActionRejectReason_INVALID_OR_UNKNOWN_SECURITY_ISSUER               = "10"
	MassActionRejectReason_INVALID_OR_UNKNOWN_ISSUER_OF_UNDERLYING_SECURITY = "11"
	MassActionRejectReason_INVALID_OR_UNKNOWN_UNDERLYING_SECURITY           = "2"
	MassActionRejectReason_INVALID_OR_UNKNOWN_PRODUCT                       = "3"
	MassActionRejectReason_INVALID_OR_UNKNOWN_CFICODE                       = "4"
	MassActionRejectReason_INVALID_OR_UNKNOWN_SECURITYTYPE                  = "5"
	MassActionRejectReason_INVALID_OR_UNKNOWN_TRADING_SESSION               = "6"
	MassActionRejectReason_INVALID_OR_UNKNOWN_MARKET                        = "7"
	MassActionRejectReason_INVALID_OR_UNKNOWN_MARKET_SEGMENT                = "8"
	MassActionRejectReason_INVALID_OR_UNKNOWN_SECURITY_GROUP                = "9"
	MassActionRejectReason_OTHER                                            = "99"
)

//Enum values for MassActionResponse
const (
	MassActionResponse_REJECTED = "0"
	MassActionResponse_ACCEPTED = "1"
)

//Enum values for MassActionScope
const (
	MassActionScope_ALL_ORDERS_FOR_A_SECURITY                = "1"
	MassActionScope_ALL_ORDERS_FOR_A_SECURITY_GROUP          = "10"
	MassActionScope_CANCEL_FOR_SECURITY_ISSUER               = "11"
	MassActionScope_CANCEL_FOR_ISSUER_OF_UNDERLYING_SECURITY = "12"
	MassActionScope_ALL_ORDERS_FOR_AN_UNDERLYING_SECURITY    = "2"
	MassActionScope_ALL_ORDERS_FOR_A_PRODUCT                 = "3"
	MassActionScope_ALL_ORDERS_FOR_A_CFICODE                 = "4"
	MassActionScope_ALL_ORDERS_FOR_A_SECURITYTYPE            = "5"
	MassActionScope_ALL_ORDERS_FOR_A_TRADING_SESSION         = "6"
	MassActionScope_ALL_ORDERS                               = "7"
	MassActionScope_ALL_ORDERS_FOR_A_MARKET                  = "8"
	MassActionScope_ALL_ORDERS_FOR_A_MARKET_SEGMENT          = "9"
)

//Enum values for MassActionType
const (
	MassActionType_SUSPEND_ORDERS                 = "1"
	MassActionType_RELEASE_ORDERS_FROM_SUSPENSION = "2"
	MassActionType_CANCEL_ORDERS                  = "3"
)

//Enum values for MassCancelRejectReason
const (
	MassCancelRejectReason_MASS_CANCEL_NOT_SUPPORTED                        = "0"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITY                      = "1"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITY_ISSUER               = "10"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_ISSUER_OF_UNDERLYING_SECURITY = "11"
	MassCancelRejectReason_INVALID_OR_UNKOWN_UNDERLYING_SECURITY            = "2"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_PRODUCT                       = "3"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_CFICODE                       = "4"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITYTYPE                  = "5"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_TRADING_SESSION               = "6"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_MARKET                        = "7"
	MassCancelRejectReason_INVALID_OR_UNKOWN_MARKET_SEGMENT                 = "8"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITY_GROUP                = "9"
	MassCancelRejectReason_OTHER                                            = "99"
)

//Enum values for MassCancelRequestType
const (
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_SECURITY             = "1"
	MassCancelRequestType_CANCEL_ORDERS_FOR_AN_UNDERLYING_SECURITY = "2"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_PRODUCT              = "3"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_CFICODE              = "4"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_SECURITYTYPE         = "5"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_TRADING_SESSION      = "6"
	MassCancelRequestType_CANCEL_ALL_ORDERS                        = "7"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_MARKET               = "8"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_MARKET_SEGMENT       = "9"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_SECURITY_GROUP       = "A"
	MassCancelRequestType_CANCEL_FOR_SECURITY_ISSUER               = "B"
	MassCancelRequestType_CANCEL_FOR_ISSUER_OF_UNDERLYING_SECURITY = "C"
)

//Enum values for MassCancelResponse
const (
	MassCancelResponse_CANCEL_REQUEST_REJECTED                         = "0"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITY                    = "1"
	MassCancelResponse_CANCEL_ORDERS_FOR_AN_UNDERLYING_SECURITY        = "2"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_PRODUCT                     = "3"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_CFICODE                     = "4"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITYTYPE                = "5"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_TRADING_SESSION             = "6"
	MassCancelResponse_CANCEL_ALL_ORDERS                               = "7"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_MARKET                      = "8"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_MARKET_SEGMENT              = "9"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITY_GROUP              = "A"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITIES_ISSUER           = "B"
	MassCancelResponse_CANCEL_ORDERS_FOR_ISSUER_OF_UNDERLYING_SECURITY = "C"
)

//Enum values for MassStatusReqType
const (
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_SECURITY             = "1"
	MassStatusReqType_STATUS_FOR_ISSUER_OF_UNDERLYING_SECURITY     = "10"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_AN_UNDERLYING_SECURITY = "2"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_PRODUCT              = "3"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_CFICODE              = "4"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_SECURITYTYPE         = "5"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_TRADING_SESSION      = "6"
	MassStatusReqType_STATUS_FOR_ALL_ORDERS                        = "7"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_PARTYID              = "8"
	MassStatusReqType_STATUS_FOR_SECURITY_ISSUER                   = "9"
)

//Enum values for MatchStatus
const (
	MatchStatus_COMPARED_MATCHED_OR_AFFIRMED       = "0"
	MatchStatus_UNCOMPARED_UNMATCHED_OR_UNAFFIRMED = "1"
	MatchStatus_ADVISORY_OR_ALERT                  = "2"
)

//Enum values for MatchType
const (
	MatchType_ONE_PARTY_TRADE_REPORT                                                                                                           = "1"
	MatchType_TWO_PARTY_TRADE_REPORT                                                                                                           = "2"
	MatchType_CONFIRMED_TRADE_REPORT                                                                                                           = "3"
	MatchType_AUTO_MATCH                                                                                                                       = "4"
	MatchType_CROSS_AUCTION                                                                                                                    = "5"
	MatchType_COUNTER_ORDER_SELECTION                                                                                                          = "6"
	MatchType_ONE_PARTY_PRIVATELY_NEGOTIATED_TRADE_REPORT                                                                                      = "60"
	MatchType_TWO_PARTY_PRIVATELY_NEGOTIATED_TRADE_REPORT                                                                                      = "61"
	MatchType_CONTINUOUS_AUTO_MATCH                                                                                                            = "62"
	MatchType_CROSS_AUCTION_63                                                                                                                 = "63"
	MatchType_COUNTER_ORDER_SELECTION_64                                                                                                       = "64"
	MatchType_CALL_AUCTION_65                                                                                                                  = "65"
	MatchType_CALL_AUCTION                                                                                                                     = "7"
	MatchType_ISSUING_BUY_BACK_AUCTION                                                                                                         = "8"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_FOUR_BADGES_AND_EXECUTION_TIME = "A1"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_FOUR_BADGES                    = "A2"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_TWO_BADGES_AND_EXECUTION_TIME  = "A3"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_TWO_BADGES                     = "A4"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADETYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_EXECUTION_TIME                  = "A5"
	MatchType_NASDAQACTM1MATCH                                                                                                                 = "ACTM1"
	MatchType_NASDAQACTM2MATCH                                                                                                                 = "ACTM2"
	MatchType_NASDAQACTACCEPTEDTRADE                                                                                                           = "ACTM3"
	MatchType_NASDAQACTDEFAULTTRADE                                                                                                            = "ACTM4"
	MatchType_NASDAQACTDEFAULTAFTERM2                                                                                                          = "ACTM5"
	MatchType_NASDAQACTM6MATCH                                                                                                                 = "ACTM6"
	MatchType_NASDAQNONACT                                                                                                                     = "ACTMT"
	MatchType_COMPARED_RECORDS_RESULTING_FROM_STAMPED_ADVISORIES_OR_SPECIALIST_ACCEPTS_PAIR_OFFS                                               = "AQ"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_MINUS_BADGES_AND_TIMES_ACT_M1_MATCH = "M1"
	MatchType_SUMMARIZED_MATCH_MINUS_BADGES_AND_TIMES_ACT_M2_MATCH                                                                             = "M2"
	MatchType_ACT_ACCEPTED_TRADE                                                                                                               = "M3"
	MatchType_ACT_DEFAULT_TRADE                                                                                                                = "M4"
	MatchType_ACT_DEFAULT_AFTER_M2                                                                                                             = "M5"
	MatchType_ACT_M6_MATCH                                                                                                                     = "M6"
	MatchType_OCS_LOCKED_IN_NON_ACT                                                                                                            = "MT"
	MatchType_SUMMARIZED_MATCH_USING_A1_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIED                                                      = "S1"
	MatchType_SUMMARIZED_MATCH_USING_A2_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIZED                                                     = "S2"
	MatchType_SUMMARIZED_MATCH_USING_A3_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIZED                                                     = "S3"
	MatchType_SUMMARIZED_MATCH_USING_A4_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIZED                                                     = "S4"
	MatchType_SUMMARIZED_MATCH_USING_A5_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIZED                                                     = "S5"
)

//Enum values for MaturityMonthYearFormat
const (
	MaturityMonthYearFormat_YEARMONTH_ONLY = "0"
	MaturityMonthYearFormat_YEARMONTHDAY   = "1"
	MaturityMonthYearFormat_YEARMONTHWEEK  = "2"
)

//Enum values for MaturityMonthYearIncrementUnits
const (
	MaturityMonthYearIncrementUnits_MONTHS = "0"
	MaturityMonthYearIncrementUnits_DAYS   = "1"
	MaturityMonthYearIncrementUnits_WEEKS  = "2"
	MaturityMonthYearIncrementUnits_YEARS  = "3"
)

//Enum values for MessageEncoding
const (
	MessageEncoding_EUC_JP      = "EUC-JP"
	MessageEncoding_ISO_2022_JP = "ISO-2022-JP"
	MessageEncoding_SHIFT_JIS   = "SHIFT_JIS"
	MessageEncoding_UTF_8       = "UTF-8"
)

//Enum values for MiscFeeBasis
const (
	MiscFeeBasis_ABSOLUTE   = "0"
	MiscFeeBasis_PER_UNIT   = "1"
	MiscFeeBasis_PERCENTAGE = "2"
)

//Enum values for MiscFeeType
const (
	MiscFeeType_REGULATORY       = "1"
	MiscFeeType_PER_TRANSACTION  = "10"
	MiscFeeType_CONVERSION       = "11"
	MiscFeeType_AGENT            = "12"
	MiscFeeType_TRANSFER_FEE     = "13"
	MiscFeeType_SECURITY_LENDING = "14"
	MiscFeeType_TAX              = "2"
	MiscFeeType_LOCAL_COMMISSION = "3"
	MiscFeeType_EXCHANGE_FEES    = "4"
	MiscFeeType_STAMP            = "5"
	MiscFeeType_LEVY             = "6"
	MiscFeeType_OTHER            = "7"
	MiscFeeType_MARKUP           = "8"
	MiscFeeType_CONSUMPTION_TAX  = "9"
)

//Enum values for ModelType
const (
	ModelType_UTILITY_PROVIDED_STANDARD_MODEL = "0"
	ModelType_PROPRIETARY                     = "1"
)

//Enum values for MoneyLaunderingStatus
const (
	MoneyLaunderingStatus_EXEMPT_1    = "1"
	MoneyLaunderingStatus_EXEMPT_2    = "2"
	MoneyLaunderingStatus_EXEMPT_3    = "3"
	MoneyLaunderingStatus_NOT_CHECKED = "N"
	MoneyLaunderingStatus_PASSED      = "Y"
)

//Enum values for MsgDirection
const (
	MsgDirection_RECEIVE = "R"
	MsgDirection_SEND    = "S"
)

//Enum values for MsgType
const (
	MsgType_HEARTBEAT                          = "0"
	MsgType_TEST_REQUEST                       = "1"
	MsgType_RESEND_REQUEST                     = "2"
	MsgType_REJECT                             = "3"
	MsgType_SEQUENCE_RESET                     = "4"
	MsgType_LOGOUT                             = "5"
	MsgType_INDICATION_OF_INTEREST             = "6"
	MsgType_ADVERTISEMENT                      = "7"
	MsgType_EXECUTION_REPORT                   = "8"
	MsgType_ORDER_CANCEL_REJECT                = "9"
	MsgType_LOGON                              = "A"
	MsgType_DERIVATIVE_SECURITY_LIST           = "AA"
	MsgType_NEW_ORDER_MULTILEG                 = "AB"
	MsgType_MULTILEG_ORDER_CANCEL_REPLACE      = "AC"
	MsgType_TRADE_CAPTURE_REPORT_REQUEST       = "AD"
	MsgType_TRADE_CAPTURE_REPORT               = "AE"
	MsgType_ORDER_MASS_STATUS_REQUEST          = "AF"
	MsgType_QUOTE_REQUEST_REJECT               = "AG"
	MsgType_RFQ_REQUEST                        = "AH"
	MsgType_QUOTE_STATUS_REPORT                = "AI"
	MsgType_QUOTE_RESPONSE                     = "AJ"
	MsgType_CONFIRMATION                       = "AK"
	MsgType_POSITION_MAINTENANCE_REQUEST       = "AL"
	MsgType_POSITION_MAINTENANCE_REPORT        = "AM"
	MsgType_REQUEST_FOR_POSITIONS              = "AN"
	MsgType_REQUEST_FOR_POSITIONS_ACK          = "AO"
	MsgType_POSITION_REPORT                    = "AP"
	MsgType_TRADE_CAPTURE_REPORT_REQUEST_ACK   = "AQ"
	MsgType_TRADE_CAPTURE_REPORT_ACK           = "AR"
	MsgType_ALLOCATION_REPORT                  = "AS"
	MsgType_ALLOCATION_REPORT_ACK              = "AT"
	MsgType_CONFIRMATION_ACK                   = "AU"
	MsgType_SETTLEMENT_INSTRUCTION_REQUEST     = "AV"
	MsgType_ASSIGNMENT_REPORT                  = "AW"
	MsgType_COLLATERAL_REQUEST                 = "AX"
	MsgType_COLLATERAL_ASSIGNMENT              = "AY"
	MsgType_COLLATERAL_RESPONSE                = "AZ"
	MsgType_NEWS                               = "B"
	MsgType_COLLATERAL_REPORT                  = "BA"
	MsgType_COLLATERAL_INQUIRY                 = "BB"
	MsgType_NETWORK_STATUS_REQUEST             = "BC"
	MsgType_NETWORK_STATUS_RESPONSE            = "BD"
	MsgType_USER_REQUEST                       = "BE"
	MsgType_USER_RESPONSE                      = "BF"
	MsgType_COLLATERAL_INQUIRY_ACK             = "BG"
	MsgType_CONFIRMATION_REQUEST               = "BH"
	MsgType_TRADING_SESSION_LIST_REQUEST       = "BI"
	MsgType_TRADING_SESSION_LIST               = "BJ"
	MsgType_SECURITY_LIST_UPDATE_REPORT        = "BK"
	MsgType_ADJUSTED_POSITION_REPORT           = "BL"
	MsgType_ALLOCATION_INSTRUCTION_ALERT       = "BM"
	MsgType_EXECUTION_ACKNOWLEDGEMENT          = "BN"
	MsgType_CONTRARY_INTENTION_REPORT          = "BO"
	MsgType_SECURITY_DEFINITION_UPDATE_REPORT  = "BP"
	MsgType_SETTLEMENTOBLIGATIONREPORT         = "BQ"
	MsgType_DERIVATIVESECURITYLISTUPDATEREPORT = "BR"
	MsgType_TRADINGSESSIONLISTUPDATEREPORT     = "BS"
	MsgType_MARKETDEFINITIONREQUEST            = "BT"
	MsgType_MARKETDEFINITION                   = "BU"
	MsgType_MARKETDEFINITIONUPDATEREPORT       = "BV"
	MsgType_APPLICATIONMESSAGEREQUEST          = "BW"
	MsgType_APPLICATIONMESSAGEREQUESTACK       = "BX"
	MsgType_APPLICATIONMESSAGEREPORT           = "BY"
	MsgType_ORDERMASSACTIONREPORT              = "BZ"
	MsgType_EMAIL                              = "C"
	MsgType_ORDERMASSACTIONREQUEST             = "CA"
	MsgType_USERNOTIFICATION                   = "CB"
	MsgType_STREAMASSIGNMENTREQUEST            = "CC"
	MsgType_STREAMASSIGNMENTREPORT             = "CD"
	MsgType_STREAMASSIGNMENTREPORTACK          = "CE"
	MsgType_PARTYDETAILSLISTREQUEST            = "CF"
	MsgType_PARTYDETAILSLISTREPORT             = "CG"
	MsgType_ORDER_SINGLE                       = "D"
	MsgType_ORDER_LIST                         = "E"
	MsgType_ORDER_CANCEL_REQUEST               = "F"
	MsgType_ORDER_CANCEL_REPLACE_REQUEST       = "G"
	MsgType_ORDER_STATUS_REQUEST               = "H"
	MsgType_ALLOCATION_INSTRUCTION             = "J"
	MsgType_LIST_CANCEL_REQUEST                = "K"
	MsgType_LIST_EXECUTE                       = "L"
	MsgType_LIST_STATUS_REQUEST                = "M"
	MsgType_LIST_STATUS                        = "N"
	MsgType_ALLOCATION_INSTRUCTION_ACK         = "P"
	MsgType_DONT_KNOW_TRADE                    = "Q"
	MsgType_QUOTE_REQUEST                      = "R"
	MsgType_QUOTE                              = "S"
	MsgType_SETTLEMENT_INSTRUCTIONS            = "T"
	MsgType_MARKET_DATA_REQUEST                = "V"
	MsgType_MARKET_DATA_SNAPSHOT_FULL_REFRESH  = "W"
	MsgType_MARKET_DATA_INCREMENTAL_REFRESH    = "X"
	MsgType_MARKET_DATA_REQUEST_REJECT         = "Y"
	MsgType_QUOTE_CANCEL                       = "Z"
	MsgType_QUOTE_STATUS_REQUEST               = "a"
	MsgType_MASS_QUOTE_ACKNOWLEDGEMENT         = "b"
	MsgType_SECURITY_DEFINITION_REQUEST        = "c"
	MsgType_SECURITY_DEFINITION                = "d"
	MsgType_SECURITY_STATUS_REQUEST            = "e"
	MsgType_SECURITY_STATUS                    = "f"
	MsgType_TRADING_SESSION_STATUS_REQUEST     = "g"
	MsgType_TRADING_SESSION_STATUS             = "h"
	MsgType_MASS_QUOTE                         = "i"
	MsgType_BUSINESS_MESSAGE_REJECT            = "j"
	MsgType_BID_REQUEST                        = "k"
	MsgType_BID_RESPONSE                       = "l"
	MsgType_LIST_STRIKE_PRICE                  = "m"
	MsgType_XML_MESSAGE                        = "n"
	MsgType_REGISTRATION_INSTRUCTIONS          = "o"
	MsgType_REGISTRATION_INSTRUCTIONS_RESPONSE = "p"
	MsgType_ORDER_MASS_CANCEL_REQUEST          = "q"
	MsgType_ORDER_MASS_CANCEL_REPORT           = "r"
	MsgType_NEW_ORDER_CROSS                    = "s"
	MsgType_CROSS_ORDER_CANCEL_REPLACE_REQUEST = "t"
	MsgType_CROSS_ORDER_CANCEL_REQUEST         = "u"
	MsgType_SECURITY_TYPE_REQUEST              = "v"
	MsgType_SECURITY_TYPES                     = "w"
	MsgType_SECURITY_LIST_REQUEST              = "x"
	MsgType_SECURITY_LIST                      = "y"
	MsgType_DERIVATIVE_SECURITY_LIST_REQUEST   = "z"
)

//Enum values for MultiLegReportingType
const (
	MultiLegReportingType_SINGLE_SECURITY                        = "1"
	MultiLegReportingType_INDIVIDUAL_LEG_OF_A_MULTI_LEG_SECURITY = "2"
	MultiLegReportingType_MULTI_LEG_SECURITY                     = "3"
)

//Enum values for MultiLegRptTypeReq
const (
	MultiLegRptTypeReq_REPORT_BY_MULITLEG_SECURITY_ONLY                                                      = "0"
	MultiLegRptTypeReq_REPORT_BY_MULTILEG_SECURITY_AND_BY_INSTRUMENT_LEGS_BELONGING_TO_THE_MULTILEG_SECURITY = "1"
	MultiLegRptTypeReq_REPORT_BY_INSTRUMENT_LEGS_BELONGING_TO_THE_MULTILEG_SECURITY_ONLY                     = "2"
)

//Enum values for MultilegModel
const (
	MultilegModel_PREDEFINED_MULTILEG_SECURITY          = "0"
	MultilegModel_USER_DEFINED_MULTLEG_SECURITY         = "1"
	MultilegModel_USER_DEFINED_NON_SECURITIZED_MULTILEG = "2"
)

//Enum values for MultilegPriceMethod
const (
	MultilegPriceMethod_NET_PRICE                       = "0"
	MultilegPriceMethod_REVERSED_NET_PRICE              = "1"
	MultilegPriceMethod_YIELD_DIFFERENCE                = "2"
	MultilegPriceMethod_INDIVIDUAL                      = "3"
	MultilegPriceMethod_CONTRACT_WEIGHTED_AVERAGE_PRICE = "4"
	MultilegPriceMethod_MULTIPLIED_PRICE                = "5"
)

//Enum values for NetGrossInd
const (
	NetGrossInd_NET   = "1"
	NetGrossInd_GROSS = "2"
)

//Enum values for NetworkRequestType
const (
	NetworkRequestType_SNAPSHOT                                        = "1"
	NetworkRequestType_SUBSCRIBE                                       = "2"
	NetworkRequestType_STOP_SUBSCRIBING                                = "4"
	NetworkRequestType_LEVEL_OF_DETAIL_THEN_NOCOMPIDS_BECOMES_REQUIRED = "8"
)

//Enum values for NetworkStatusResponseType
const (
	NetworkStatusResponseType_FULL               = "1"
	NetworkStatusResponseType_INCREMENTAL_UPDATE = "2"
)

//Enum values for NewsCategory
const (
	NewsCategory_COMPANY_NEWS          = "0"
	NewsCategory_MARKETPLACE_NEWS      = "1"
	NewsCategory_FINANCIAL_MARKET_NEWS = "2"
	NewsCategory_TECHNICAL_NEWS        = "3"
	NewsCategory_OTHER_NEWS            = "99"
)

//Enum values for NewsRefType
const (
	NewsRefType_REPLACEMENT    = "0"
	NewsRefType_OTHER_LANGUAGE = "1"
	NewsRefType_COMPLIMENTARY  = "2"
)

//Enum values for NoSides
const (
	NoSides_ONE_SIDE   = "1"
	NoSides_BOTH_SIDES = "2"
)

//Enum values for NotifyBrokerOfCredit
const (
	NotifyBrokerOfCredit_NO  = "N"
	NotifyBrokerOfCredit_YES = "Y"
)

//Enum values for OddLot
const (
	OddLot_NO  = "N"
	OddLot_YES = "Y"
)

//Enum values for OpenClose
const (
	OpenClose_CLOSE = "C"
	OpenClose_OPEN  = "O"
)

//Enum values for OpenCloseSettlFlag
const (
	OpenCloseSettlFlag_DAILY_OPEN                       = "0"
	OpenCloseSettlFlag_SESSION_OPEN                     = "1"
	OpenCloseSettlFlag_DELIVERY_SETTLEMENT_ENTRY        = "2"
	OpenCloseSettlFlag_EXPECTED_ENTRY                   = "3"
	OpenCloseSettlFlag_ENTRY_FROM_PREVIOUS_BUSINESS_DAY = "4"
	OpenCloseSettlFlag_THEORETICAL_PRICE_VALUE          = "5"
)

//Enum values for OpenCloseSettleFlag
const (
	OpenCloseSettleFlag_DAILY_OPEN                       = "0"
	OpenCloseSettleFlag_SESSION_OPEN                     = "1"
	OpenCloseSettleFlag_DELIVERY_SETTLEMENT_PRICE        = "2"
	OpenCloseSettleFlag_EXPECTED_PRICE                   = "3"
	OpenCloseSettleFlag_PRICE_FROM_PREVIOUS_BUSINESS_DAY = "4"
)

//Enum values for OptPayoutType
const (
	OptPayoutType_VANILLA = "1"
	OptPayoutType_CAPPED  = "2"
	OptPayoutType_BINARY  = "3"
)

//Enum values for OrdRejReason
const (
	OrdRejReason_BROKER                                     = "0"
	OrdRejReason_UNKNOWN_SYMBOL                             = "1"
	OrdRejReason_INVALID_INVESTOR_ID                        = "10"
	OrdRejReason_UNSUPPORTED_ORDER_CHARACTERISTIC           = "11"
	OrdRejReason_SURVEILLENCE_OPTION                        = "12"
	OrdRejReason_INCORRECT_QUANTITY                         = "13"
	OrdRejReason_INCORRECT_ALLOCATED_QUANTITY               = "14"
	OrdRejReason_UNKNOWN_ACCOUNT                            = "15"
	OrdRejReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND           = "16"
	OrdRejReason_INVALID_PRICE_INCREMENT                    = "18"
	OrdRejReason_EXCHANGE_CLOSED                            = "2"
	OrdRejReason_ORDER_EXCEEDS_LIMIT                        = "3"
	OrdRejReason_TOO_LATE_TO_ENTER                          = "4"
	OrdRejReason_UNKNOWN_ORDER                              = "5"
	OrdRejReason_DUPLICATE_ORDER                            = "6"
	OrdRejReason_DUPLICATE_OF_A_VERBALLY_COMMUNICATED_ORDER = "7"
	OrdRejReason_STALE_ORDER                                = "8"
	OrdRejReason_TRADE_ALONG_REQUIRED                       = "9"
	OrdRejReason_OTHER                                      = "99"
)

//Enum values for OrdStatus
const (
	OrdStatus_NEW                  = "0"
	OrdStatus_PARTIALLY_FILLED     = "1"
	OrdStatus_FILLED               = "2"
	OrdStatus_DONE_FOR_DAY         = "3"
	OrdStatus_CANCELED             = "4"
	OrdStatus_REPLACED             = "5"
	OrdStatus_PENDING_CANCEL       = "6"
	OrdStatus_STOPPED              = "7"
	OrdStatus_REJECTED             = "8"
	OrdStatus_SUSPENDED            = "9"
	OrdStatus_PENDING_NEW          = "A"
	OrdStatus_CALCULATED           = "B"
	OrdStatus_EXPIRED              = "C"
	OrdStatus_ACCEPTED_FOR_BIDDING = "D"
	OrdStatus_PENDING_REPLACE      = "E"
)

//Enum values for OrdType
const (
	OrdType_MARKET                         = "1"
	OrdType_LIMIT                          = "2"
	OrdType_STOP                           = "3"
	OrdType_STOP_LIMIT                     = "4"
	OrdType_MARKET_ON_CLOSE                = "5"
	OrdType_WITH_OR_WITHOUT                = "6"
	OrdType_LIMIT_OR_BETTER                = "7"
	OrdType_LIMIT_WITH_OR_WITHOUT          = "8"
	OrdType_ON_BASIS                       = "9"
	OrdType_ON_CLOSE                       = "A"
	OrdType_LIMIT_ON_CLOSE                 = "B"
	OrdType_FOREX_MARKET                   = "C"
	OrdType_PREVIOUSLY_QUOTED              = "D"
	OrdType_PREVIOUSLY_INDICATED           = "E"
	OrdType_FOREX_LIMIT                    = "F"
	OrdType_FOREX_SWAP                     = "G"
	OrdType_FOREX_PREVIOUSLY_QUOTED        = "H"
	OrdType_FUNARI                         = "I"
	OrdType_MARKET_IF_TOUCHED              = "J"
	OrdType_MARKET_WITH_LEFT_OVER_AS_LIMIT = "K"
	OrdType_PREVIOUS_FUND_VALUATION_POINT  = "L"
	OrdType_NEXT_FUND_VALUATION_POINT      = "M"
	OrdType_PEGGED                         = "P"
	OrdType_COUNTER_ORDER_SELECTION        = "Q"
)

//Enum values for OrderCapacity
const (
	OrderCapacity_AGENCY                 = "A"
	OrderCapacity_PROPRIETARY            = "G"
	OrderCapacity_INDIVIDUAL             = "I"
	OrderCapacity_PRINCIPAL              = "P"
	OrderCapacity_RISKLESS_PRINCIPAL     = "R"
	OrderCapacity_AGENT_FOR_OTHER_MEMBER = "W"
)

//Enum values for OrderCategory
const (
	OrderCategory_ORDER                      = "1"
	OrderCategory_QUOTE                      = "2"
	OrderCategory_PRIVATELY_NEGOTIATED_TRADE = "3"
	OrderCategory_MULTILEG_ORDER             = "4"
	OrderCategory_LINKED_ORDER               = "5"
	OrderCategory_QUOTE_REQUEST              = "6"
	OrderCategory_IMPLIED_ORDER              = "7"
	OrderCategory_CROSS_ORDER                = "8"
	OrderCategory_STREAMING_PRICE            = "9"
)

//Enum values for OrderDelayUnit
const (
	OrderDelayUnit_SECONDS                = "0"
	OrderDelayUnit_TENTHS_OF_A_SECOND     = "1"
	OrderDelayUnit_MINUTES                = "10"
	OrderDelayUnit_HOURS                  = "11"
	OrderDelayUnit_DAYS                   = "12"
	OrderDelayUnit_WEEKS                  = "13"
	OrderDelayUnit_MONTHS                 = "14"
	OrderDelayUnit_YEARS                  = "15"
	OrderDelayUnit_HUNDREDTHS_OF_A_SECOND = "2"
	OrderDelayUnit_MILLISECONDS           = "3"
	OrderDelayUnit_MICROSECONDS           = "4"
	OrderDelayUnit_NANOSECONDS            = "5"
)

//Enum values for OrderHandlingInstSource
const (
	OrderHandlingInstSource_NASD_OATS = "1"
)

//Enum values for OrderRestrictions
const (
	OrderRestrictions_PROGRAM_TRADE                                                                            = "1"
	OrderRestrictions_INDEX_ARBITRAGE                                                                          = "2"
	OrderRestrictions_NON_INDEX_ARBITRAGE                                                                      = "3"
	OrderRestrictions_COMPETING_MARKET_MAKER                                                                   = "4"
	OrderRestrictions_ACTING_AS_MARKET_MAKER_OR_SPECIALIST_IN_THE_SECURITY                                     = "5"
	OrderRestrictions_ACTING_AS_MARKET_MAKER_OR_SPECIALIST_IN_THE_UNDERLYING_SECURITY_OF_A_DERIVATIVE_SECURITY = "6"
	OrderRestrictions_FOREIGN_ENTITY                                                                           = "7"
	OrderRestrictions_EXTERNAL_MARKET_PARTICIPANT                                                              = "8"
	OrderRestrictions_EXTERNAL_INTER_CONNECTED_MARKET_LINKAGE                                                  = "9"
	OrderRestrictions_RISKLESS_ARBITRAGE                                                                       = "A"
	OrderRestrictions_ISSUER_HOLDING                                                                           = "B"
	OrderRestrictions_ISSUE_PRICE_STABILIZATION                                                                = "C"
	OrderRestrictions_NON_ALGORITHMIC                                                                          = "D"
	OrderRestrictions_ALGORITHMIC                                                                              = "E"
	OrderRestrictions_CROSS                                                                                    = "F"
)

//Enum values for OrigCustOrderCapacity
const (
	OrigCustOrderCapacity_MEMBER_TRADING_FOR_THEIR_OWN_ACCOUNT              = "1"
	OrigCustOrderCapacity_CLEARING_FIRM_TRADING_FOR_ITS_PROPRIETARY_ACCOUNT = "2"
	OrigCustOrderCapacity_MEMBER_TRADING_FOR_ANOTHER_MEMBER                 = "3"
	OrigCustOrderCapacity_ALL_OTHER                                         = "4"
)

//Enum values for OwnerType
const (
	OwnerType_INDIVIDUAL_INVESTOR                 = "1"
	OwnerType_NETWORKING_SUB_ACCOUNT              = "10"
	OwnerType_NON_PROFIT_ORGANIZATION             = "11"
	OwnerType_CORPORATE_BODY                      = "12"
	OwnerType_NOMINEE                             = "13"
	OwnerType_PUBLIC_COMPANY                      = "2"
	OwnerType_PRIVATE_COMPANY                     = "3"
	OwnerType_INDIVIDUAL_TRUSTEE                  = "4"
	OwnerType_COMPANY_TRUSTEE                     = "5"
	OwnerType_PENSION_PLAN                        = "6"
	OwnerType_CUSTODIAN_UNDER_GIFTS_TO_MINORS_ACT = "7"
	OwnerType_TRUSTS                              = "8"
	OwnerType_FIDUCIARIES                         = "9"
)

//Enum values for OwnershipType
const (
	OwnershipType_JOINT_TRUSTEES    = "2"
	OwnershipType_JOINT_INVESTORS   = "J"
	OwnershipType_TENANTS_IN_COMMON = "T"
)

//Enum values for PartyDetailsRequestResult
const (
	PartyDetailsRequestResult_VALID_REQUEST                                                   = "0"
	PartyDetailsRequestResult_INVALID_OR_UNSUPPORTED_REQUEST                                  = "1"
	PartyDetailsRequestResult_NO_PARTIES_OR_PARTY_DETAILS_FOUND_THAT_MATCH_SELECTION_CRITERIA = "2"
	PartyDetailsRequestResult_UNSUPPORTED_PARTYLISTRESPONSETYPE                               = "3"
	PartyDetailsRequestResult_NOT_AUTHORIZED_TO_RETRIEVE_PARTIES_OR_PARTY_DETAILS_DATA        = "4"
	PartyDetailsRequestResult_PARTIES_OR_PARTY_DETAILS_DATA_TEMPORARILY_UNAVAILABLE           = "5"
	PartyDetailsRequestResult_REQUEST_FOR_PARTIES_DATA_NOT_SUPPORTED                          = "6"
	PartyDetailsRequestResult_OTHER                                                           = "99"
)

//Enum values for PartyIDSource
const (
	PartyIDSource_KOREAN_INVESTOR_ID                                                                                = "1"
	PartyIDSource_TAIWANESE_QUALIFIED_FOREIGN_INVESTOR_ID_QFII_FID                                                  = "2"
	PartyIDSource_TAIWANESE_TRADING_ACCT                                                                            = "3"
	PartyIDSource_MALAYSIAN_CENTRAL_DEPOSITORY                                                                      = "4"
	PartyIDSource_CHINESE_INVESTOR_ID                                                                               = "5"
	PartyIDSource_UK_NATIONAL_INSURANCE_OR_PENSION_NUMBER                                                           = "6"
	PartyIDSource_US_SOCIAL_SECURITY_NUMBER                                                                         = "7"
	PartyIDSource_US_EMPLOYER_OR_TAX_ID_NUMBER                                                                      = "8"
	PartyIDSource_AUSTRALIAN_BUSINESS_NUMBER                                                                        = "9"
	PartyIDSource_AUSTRALIAN_TAX_FILE_NUMBER                                                                        = "A"
	PartyIDSource_BIC                                                                                               = "B"
	PartyIDSource_GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER                                                  = "C"
	PartyIDSource_PROPRIETARY                                                                                       = "D"
	PartyIDSource_ISO_COUNTRY_CODE                                                                                  = "E"
	PartyIDSource_SETTLEMENT_ENTITY_LOCATION                                                                        = "F"
	PartyIDSource_MIC                                                                                               = "G"
	PartyIDSource_CSD_PARTICIPANT_MEMBER_CODE                                                                       = "H"
	PartyIDSource_DIRECTED_BROKER_THREE_CHARACTER_ACRONYM_AS_DEFINED_IN_ISITC_ETC_BEST_PRACTICE_GUIDELINES_DOCUMENT = "I"
)

//Enum values for PartyListResponseType
const (
	PartyListResponseType_RETURN_ALL_AVAILABLE_INFORMATION_ON_PARTIES_AND_RELATED_PARTIES = "0"
	PartyListResponseType_RETURN_ONLY_PARTY_INFORMATION                                   = "1"
	PartyListResponseType_INCLUDE_INFORMATION_ON_RELATED_PARTIES                          = "2"
	PartyListResponseType_INCLUDE_RISK_LIMIT_INFORMATION                                  = "3"
)

//Enum values for PartyRelationship
const (
	PartyRelationship_IS_ALSO                       = "0"
	PartyRelationship_CLEARS_FOR                    = "1"
	PartyRelationship_HAS_MEMBERS                   = "10"
	PartyRelationship_PROVIDES_MARKETPLACE_FOR      = "11"
	PartyRelationship_PARTICIPANT_OF_MARKETPLACE    = "12"
	PartyRelationship_CARRIES_POSITIONS_FOR         = "13"
	PartyRelationship_POSTS_TRADES_TO               = "14"
	PartyRelationship_ENTERS_TRADES_FOR             = "15"
	PartyRelationship_ENTERS_TRADES_THROUGH         = "16"
	PartyRelationship_PROVIDES_QUOTES_TO            = "17"
	PartyRelationship_REQUESTS_QUOTES_FROM          = "18"
	PartyRelationship_INVESTS_FOR                   = "19"
	PartyRelationship_CLEARS_THROUGH                = "2"
	PartyRelationship_INVESTS_THROUGH               = "20"
	PartyRelationship_BROKERS_TRADES_FOR            = "21"
	PartyRelationship_BROKERS_TRADES_THROUGH        = "22"
	PartyRelationship_PROVIDES_TRADING_SERVICES_FOR = "23"
	PartyRelationship_USES_TRADING_SERVICES_OF      = "24"
	PartyRelationship_APPROVES_OF                   = "25"
	PartyRelationship_APPROVED_BY                   = "26"
	PartyRelationship_PARENT_FIRM_FOR               = "27"
	PartyRelationship_SUBSIDIARY_OF                 = "28"
	PartyRelationship_REGULATORY_OWNER_OF           = "29"
	PartyRelationship_TRADES_FOR                    = "3"
	PartyRelationship_OWNED_BY_30                   = "30"
	PartyRelationship_CONTROLS                      = "31"
	PartyRelationship_IS_CONTROLLED_BY              = "32"
	PartyRelationship_LEGAL                         = "33"
	PartyRelationship_OWNED_BY_34                   = "34"
	PartyRelationship_BENEFICIAL_OWNER_OF           = "35"
	PartyRelationship_OWNED_BY_36                   = "36"
	PartyRelationship_TRADES_THROUGH                = "4"
	PartyRelationship_SPONSORS                      = "5"
	PartyRelationship_SPONSORED_THROUGH             = "6"
	PartyRelationship_PROVIDES_GUARANTEE_FOR        = "7"
	PartyRelationship_IS_GUARANTEED_BY              = "8"
	PartyRelationship_MEMBER_OF                     = "9"
)

//Enum values for PartyRole
const (
	PartyRole_EXECUTING_FIRM                                                        = "1"
	PartyRole_SETTLEMENT_LOCATION                                                   = "10"
	PartyRole_ORDER_ORIGINATION_TRADER                                              = "11"
	PartyRole_EXECUTING_TRADER                                                      = "12"
	PartyRole_ORDER_ORIGINATION_FIRM                                                = "13"
	PartyRole_GIVEUP_CLEARING_FIRM                                                  = "14"
	PartyRole_CORRESPONDANT_CLEARING_FIRM                                           = "15"
	PartyRole_EXECUTING_SYSTEM                                                      = "16"
	PartyRole_CONTRA_FIRM                                                           = "17"
	PartyRole_CONTRA_CLEARING_FIRM                                                  = "18"
	PartyRole_SPONSORING_FIRM                                                       = "19"
	PartyRole_BROKER_OF_CREDIT                                                      = "2"
	PartyRole_UNDERLYING_CONTRA_FIRM                                                = "20"
	PartyRole_CLEARING_ORGANIZATION                                                 = "21"
	PartyRole_EXCHANGE                                                              = "22"
	PartyRole_CUSTOMER_ACCOUNT                                                      = "24"
	PartyRole_CORRESPONDENT_CLEARING_ORGANIZATION                                   = "25"
	PartyRole_CORRESPONDENT_BROKER                                                  = "26"
	PartyRole_BUYER_SELLER                                                          = "27"
	PartyRole_CUSTODIAN                                                             = "28"
	PartyRole_INTERMEDIARY                                                          = "29"
	PartyRole_CLIENT_ID                                                             = "3"
	PartyRole_AGENT                                                                 = "30"
	PartyRole_SUB_CUSTODIAN                                                         = "31"
	PartyRole_BENEFICIARY                                                           = "32"
	PartyRole_INTERESTED_PARTY                                                      = "33"
	PartyRole_REGULATORY_BODY                                                       = "34"
	PartyRole_LIQUIDITY_PROVIDER                                                    = "35"
	PartyRole_ENTERING_TRADER                                                       = "36"
	PartyRole_CONTRA_TRADER                                                         = "37"
	PartyRole_POSITION_ACCOUNT                                                      = "38"
	PartyRole_CONTRA_INVESTOR_ID                                                    = "39"
	PartyRole_CLEARING_FIRM                                                         = "4"
	PartyRole_TRANSFER_TO_FIRM                                                      = "40"
	PartyRole_CONTRA_POSITION_ACCOUNT                                               = "41"
	PartyRole_CONTRA_EXCHANGE                                                       = "42"
	PartyRole_INTERNAL_CARRY_ACCOUNT                                                = "43"
	PartyRole_ORDER_ENTRY_OPERATOR_ID                                               = "44"
	PartyRole_SECONDARY_ACCOUNT_NUMBER                                              = "45"
	PartyRole_FOREIGN_FIRM                                                          = "46"
	PartyRole_THIRD_PARTY_ALLOCATION_FIRM                                           = "47"
	PartyRole_CLAIMING_ACCOUNT                                                      = "48"
	PartyRole_ASSET_MANAGER                                                         = "49"
	PartyRole_INVESTOR_ID                                                           = "5"
	PartyRole_PLEDGOR_ACCOUNT                                                       = "50"
	PartyRole_PLEDGEE_ACCOUNT                                                       = "51"
	PartyRole_LARGE_TRADER_REPORTABLE_ACCOUNT                                       = "52"
	PartyRole_TRADER_MNEMONIC                                                       = "53"
	PartyRole_SENDER_LOCATION                                                       = "54"
	PartyRole_SESSION_ID                                                            = "55"
	PartyRole_ACCEPTABLE_COUNTERPARTY                                               = "56"
	PartyRole_UNACCEPTABLE_COUNTERPARTY                                             = "57"
	PartyRole_ENTERING_UNIT                                                         = "58"
	PartyRole_EXECUTING_UNIT                                                        = "59"
	PartyRole_INTRODUCING_FIRM                                                      = "6"
	PartyRole_INTRODUCING_BROKER                                                    = "60"
	PartyRole_QUOTE_ORIGINATOR                                                      = "61"
	PartyRole_REPORT_ORIGINATOR                                                     = "62"
	PartyRole_SYSTEMATIC_INTERNALISER                                               = "63"
	PartyRole_MULTILATERAL_TRADING_FACILITY                                         = "64"
	PartyRole_REGULATED_MARKET                                                      = "65"
	PartyRole_MARKET_MAKER                                                          = "66"
	PartyRole_INVESTMENT_FIRM                                                       = "67"
	PartyRole_HOST_COMPETENT_AUTHORITY                                              = "68"
	PartyRole_HOME_COMPETENT_AUTHORITY                                              = "69"
	PartyRole_ENTERING_FIRM                                                         = "7"
	PartyRole_COMPETENT_AUTHORITY_OF_THE_MOST_RELEVANT_MARKET_IN_TERMS_OF_LIQUIDITY = "70"
	PartyRole_COMPETENT_AUTHORITY_OF_THE_TRANSACTION                                = "71"
	PartyRole_REPORTING_INTERMEDIARY                                                = "72"
	PartyRole_EXECUTION_VENUE                                                       = "73"
	PartyRole_MARKET_DATA_ENTRY_ORIGINATOR                                          = "74"
	PartyRole_LOCATION_ID                                                           = "75"
	PartyRole_DESK_ID                                                               = "76"
	PartyRole_MARKET_DATA_MARKET                                                    = "77"
	PartyRole_ALLOCATION_ENTITY                                                     = "78"
	PartyRole_PRIME_BROKER_PROVIDING_GENERAL_TRADE_SERVICES                         = "79"
	PartyRole_LOCATE                                                                = "8"
	PartyRole_STEP_OUT_FIRM                                                         = "80"
	PartyRole_BROKERCLEARINGID                                                      = "81"
	PartyRole_CENTRAL_REGISTRATION_DEPOSITORY                                       = "82"
	PartyRole_CLEARING_ACCOUNT                                                      = "83"
	PartyRole_ACCEPTABLE_SETTLING_COUNTERPARTY                                      = "84"
	PartyRole_UNACCEPTABLE_SETTLING_COUNTERPARTY                                    = "85"
	PartyRole_FUND_MANAGER_CLIENT_ID                                                = "9"
)

//Enum values for PartySubIDType
const (
	PartySubIDType_FIRM                                                          = "1"
	PartySubIDType_SECURITIES_ACCOUNT_NUMBER                                     = "10"
	PartySubIDType_REGISTRATION_NUMBER                                           = "11"
	PartySubIDType_REGISTERED_ADDRESS_12                                         = "12"
	PartySubIDType_REGULATORY_STATUS                                             = "13"
	PartySubIDType_REGISTRATION_NAME                                             = "14"
	PartySubIDType_CASH_ACCOUNT_NUMBER                                           = "15"
	PartySubIDType_BIC                                                           = "16"
	PartySubIDType_CSD_PARTICIPANT_MEMBER_CODE                                   = "17"
	PartySubIDType_REGISTERED_ADDRESS_18                                         = "18"
	PartySubIDType_FUND_ACCOUNT_NAME                                             = "19"
	PartySubIDType_PERSON                                                        = "2"
	PartySubIDType_TELEX_NUMBER                                                  = "20"
	PartySubIDType_FAX_NUMBER                                                    = "21"
	PartySubIDType_SECURITIES_ACCOUNT_NAME                                       = "22"
	PartySubIDType_CASH_ACCOUNT_NAME                                             = "23"
	PartySubIDType_DEPARTMENT                                                    = "24"
	PartySubIDType_LOCATION_DESK                                                 = "25"
	PartySubIDType_POSITION_ACCOUNT_TYPE                                         = "26"
	PartySubIDType_SECURITY_LOCATE_ID                                            = "27"
	PartySubIDType_MARKET_MAKER                                                  = "28"
	PartySubIDType_ELIGIBLE_COUNTERPARTY                                         = "29"
	PartySubIDType_SYSTEM                                                        = "3"
	PartySubIDType_PROFESSIONAL_CLIENT                                           = "30"
	PartySubIDType_LOCATION                                                      = "31"
	PartySubIDType_EXECUTION_VENUE                                               = "32"
	PartySubIDType_CURRENCY_DELIVERY_IDENTIFIER                                  = "33"
	PartySubIDType_APPLICATION                                                   = "4"
	PartySubIDType_RESERVEDANDAVAILABLEFORBILATERALLYAGREEDUPONUSERDEFINEDVALUES = "4000"
	PartySubIDType_FULL_LEGAL_NAME_OF_FIRM                                       = "5"
	PartySubIDType_POSTAL_ADDRESS                                                = "6"
	PartySubIDType_PHONE_NUMBER                                                  = "7"
	PartySubIDType_EMAIL_ADDRESS                                                 = "8"
	PartySubIDType_CONTACT_NAME                                                  = "9"
)

//Enum values for PaymentMethod
const (
	PaymentMethod_CREST                      = "1"
	PaymentMethod_DIRECT_CREDIT              = "10"
	PaymentMethod_CREDIT_CARD                = "11"
	PaymentMethod_ACH_DEBIT                  = "12"
	PaymentMethod_ACH_CREDIT                 = "13"
	PaymentMethod_BPAY                       = "14"
	PaymentMethod_HIGH_VALUE_CLEARING_SYSTEM = "15"
	PaymentMethod_NSCC                       = "2"
	PaymentMethod_EUROCLEAR                  = "3"
	PaymentMethod_CLEARSTREAM                = "4"
	PaymentMethod_CHEQUE                     = "5"
	PaymentMethod_TELEGRAPHIC_TRANSFER       = "6"
	PaymentMethod_FED_WIRE                   = "7"
	PaymentMethod_DEBIT_CARD                 = "8"
	PaymentMethod_DIRECT_DEBIT               = "9"
)

//Enum values for PegLimitType
const (
	PegLimitType_OR_BETTER = "0"
	PegLimitType_STRICT    = "1"
	PegLimitType_OR_WORSE  = "2"
)

//Enum values for PegMoveType
const (
	PegMoveType_FLOATING = "0"
	PegMoveType_FIXED    = "1"
)

//Enum values for PegOffsetType
const (
	PegOffsetType_PRICE        = "0"
	PegOffsetType_BASIS_POINTS = "1"
	PegOffsetType_TICKS        = "2"
	PegOffsetType_PRICE_TIER   = "3"
)

//Enum values for PegPriceType
const (
	PegPriceType_LAST_PEG                                              = "1"
	PegPriceType_MID_PRICE_PEG                                         = "2"
	PegPriceType_OPENING_PEG                                           = "3"
	PegPriceType_MARKET_PEG                                            = "4"
	PegPriceType_PRIMARY_PEG                                           = "5"
	PegPriceType_FIXED_PEG_TO_LOCAL_BEST_BID_OR_OFFER_AT_TIME_OF_ORDER = "6"
	PegPriceType_PEG_TO_VWAP                                           = "7"
	PegPriceType_TRAILING_STOP_PEG                                     = "8"
	PegPriceType_PEG_TO_LIMIT_PRICE                                    = "9"
)

//Enum values for PegRoundDirection
const (
	PegRoundDirection_MORE_AGGRESSIVE = "1"
	PegRoundDirection_MORE_PASSIVE    = "2"
)

//Enum values for PegScope
const (
	PegScope_LOCAL                    = "1"
	PegScope_NATIONAL                 = "2"
	PegScope_GLOBAL                   = "3"
	PegScope_NATIONAL_EXCLUDING_LOCAL = "4"
)

//Enum values for PosAmtType
const (
	PosAmtType_ACCRUED_COUPON_AMOUNT                     = "ACPN"
	PosAmtType_TOTAL_BANKED_AMOUNT                       = "BANK"
	PosAmtType_CASH_AMOUNT                               = "CASH"
	PosAmtType_COLLATERALIZED_MARK_TO_MARKET             = "CMTM"
	PosAmtType_TOTAL_COLLATERALIZED_AMOUNT               = "COLAT"
	PosAmtType_COUPON_AMOUNT                             = "CPN"
	PosAmtType_CASH_RESIDUAL_AMOUNT                      = "CRES"
	PosAmtType_COMPENSATION_AMOUNT                       = "DLV"
	PosAmtType_FINAL_MARK_TO_MARKET_AMOUNT               = "FMTM"
	PosAmtType_INCREMENTAL_ACCRUED_COUPON                = "IACPN"
	PosAmtType_INCREMENTAL_COLLATERALIZED_MARK_TO_MARKET = "ICMTM"
	PosAmtType_INITIAL_TRADE_COUPON_AMOUNT               = "ICPN"
	PosAmtType_INCREMENTAL_MARK_TO_MARKET_AMOUNT         = "IMTM"
	PosAmtType_PREMIUM_AMOUNT                            = "PREM"
	PosAmtType_SETTLEMENT_VALUE                          = "SETL"
	PosAmtType_START_OF_DAY_MARK_TO_MARKET_AMOUNT        = "SMTM"
	PosAmtType_TRADE_VARIATION_AMOUNT                    = "TVAR"
	PosAmtType_VALUE_ADJUSTED_AMOUNT                     = "VADJ"
)

//Enum values for PosMaintAction
const (
	PosMaintAction_NEW     = "1"
	PosMaintAction_REPLACE = "2"
	PosMaintAction_CANCEL  = "3"
	PosMaintAction_REVERSE = "4"
)

//Enum values for PosMaintResult
const (
	PosMaintResult_SUCCESSFUL_COMPLETION = "0"
	PosMaintResult_REJECTED              = "1"
	PosMaintResult_OTHER                 = "99"
)

//Enum values for PosMaintStatus
const (
	PosMaintStatus_ACCEPTED                = "0"
	PosMaintStatus_ACCEPTED_WITH_WARNINGS  = "1"
	PosMaintStatus_REJECTED                = "2"
	PosMaintStatus_COMPLETED               = "3"
	PosMaintStatus_COMPLETED_WITH_WARNINGS = "4"
)

//Enum values for PosQtyStatus
const (
	PosQtyStatus_SUBMITTED = "0"
	PosQtyStatus_ACCEPTED  = "1"
	PosQtyStatus_REJECTED  = "2"
)

//Enum values for PosReqResult
const (
	PosReqResult_VALID_REQUEST                          = "0"
	PosReqResult_INVALID_OR_UNSUPPORTED_REQUEST         = "1"
	PosReqResult_NO_POSITIONS_FOUND_THAT_MATCH_CRITERIA = "2"
	PosReqResult_NOT_AUTHORIZED_TO_REQUEST_POSITIONS    = "3"
	PosReqResult_REQUEST_FOR_POSITION_NOT_SUPPORTED     = "4"
	PosReqResult_OTHER                                  = "99"
)

//Enum values for PosReqStatus
const (
	PosReqStatus_COMPLETED               = "0"
	PosReqStatus_COMPLETED_WITH_WARNINGS = "1"
	PosReqStatus_REJECTED                = "2"
)

//Enum values for PosReqType
const (
	PosReqType_POSITIONS           = "0"
	PosReqType_TRADES              = "1"
	PosReqType_EXERCISES           = "2"
	PosReqType_ASSIGNMENTS         = "3"
	PosReqType_SETTLEMENT_ACTIVITY = "4"
	PosReqType_BACKOUT_MESSAGE     = "5"
	PosReqType_DELTA_POSITIONS     = "6"
)

//Enum values for PosTransType
const (
	PosTransType_EXERCISE                                      = "1"
	PosTransType_DO_NOT_EXERCISE                               = "2"
	PosTransType_POSITION_ADJUSTMENT                           = "3"
	PosTransType_POSITION_CHANGE_SUBMISSION_MARGIN_DISPOSITION = "4"
	PosTransType_PLEDGE                                        = "5"
	PosTransType_LARGE_TRADER_SUBMISSION                       = "6"
)

//Enum values for PosType
const (
	PosType_ALLOCATION_TRADE_QTY           = "ALC"
	PosType_OPTION_ASSIGNMENT              = "AS"
	PosType_AS_OF_TRADE_QTY                = "ASF"
	PosType_CORPORATE_ACTION_ADJUSTMENT    = "CAA"
	PosType_CREDIT_EVENT_ADJUSTMENT        = "CEA"
	PosType_NET_DELTA_QTY                  = "DLT"
	PosType_DELIVERY_QTY                   = "DLV"
	PosType_DELIVERY_NOTICE_QTY            = "DN"
	PosType_EXCHANGE_FOR_PHYSICAL_QTY      = "EP"
	PosType_ELECTRONIC_TRADE_QTY           = "ETR"
	PosType_OPTION_EXERCISE_QTY            = "EX"
	PosType_END_OF_DAY_QTY                 = "FIN"
	PosType_INTRA_SPREAD_QTY               = "IAS"
	PosType_INTER_SPREAD_QTY               = "IES"
	PosType_ADJUSTMENT_QTY                 = "PA"
	PosType_PIT_TRADE_QTY                  = "PIT"
	PosType_PRIVATELY_NEGOTIATED_TRADE_QTY = "PNTN"
	PosType_RECEIVE_QUANTITY               = "RCV"
	PosType_SUCCESSION_EVENT_ADJUSTMENT    = "SEA"
	PosType_START_OF_DAY_QTY               = "SOD"
	PosType_INTEGRAL_SPLIT                 = "SPL"
	PosType_TRANSACTION_FROM_ASSIGNMENT    = "TA"
	PosType_TOTAL_TRANSACTION_QTY          = "TOT"
	PosType_TRANSACTION_QUANTITY           = "TQ"
	PosType_TRANSFER_TRADE_QTY             = "TRF"
	PosType_TRANSACTION_FROM_EXERCISE      = "TX"
	PosType_CROSS_MARGIN_QTY               = "XM"
)

//Enum values for PositionEffect
const (
	PositionEffect_CLOSE                    = "C"
	PositionEffect_DEFAULT                  = "D"
	PositionEffect_FIFO                     = "F"
	PositionEffect_CLOSE_BUT_NOTIFY_ON_OPEN = "N"
	PositionEffect_OPEN                     = "O"
	PositionEffect_ROLLED                   = "R"
)

//Enum values for PossDupFlag
const (
	PossDupFlag_NO  = "N"
	PossDupFlag_YES = "Y"
)

//Enum values for PossResend
const (
	PossResend_NO  = "N"
	PossResend_YES = "Y"
)

//Enum values for PreallocMethod
const (
	PreallocMethod_PRO_RATA        = "0"
	PreallocMethod_DO_NOT_PRO_RATA = "1"
)

//Enum values for PreviouslyReported
const (
	PreviouslyReported_NO  = "N"
	PreviouslyReported_YES = "Y"
)

//Enum values for PriceLimitType
const (
	PriceLimitType_PRICE      = "0"
	PriceLimitType_TICKS      = "1"
	PriceLimitType_PERCENTAGE = "2"
)

//Enum values for PriceProtectionScope
const (
	PriceProtectionScope_NONE     = "0"
	PriceProtectionScope_LOCAL    = "1"
	PriceProtectionScope_NATIONAL = "2"
	PriceProtectionScope_GLOBAL   = "3"
)

//Enum values for PriceQuoteMethod
const (
	PriceQuoteMethod_INTEREST_RATE_INDEX = "INT"
	PriceQuoteMethod_INDEX               = "INX"
	PriceQuoteMethod_PERCENT_OF_PAR      = "PCTPAR"
	PriceQuoteMethod_STANDARD            = "STD"
)

//Enum values for PriceType
const (
	PriceType_PERCENTAGE                         = "1"
	PriceType_FIXED_CABINET_TRADE_PRICE          = "10"
	PriceType_VARIABLE_CABINET_TRADE_PRICE       = "11"
	PriceType_PRODUCT_TICKS_IN_HALFS             = "13"
	PriceType_PRODUCT_TICKS_IN_FOURTHS           = "14"
	PriceType_PRODUCT_TICKS_IN_EIGHTS            = "15"
	PriceType_PRODUCT_TICKS_IN_SIXTEENTHS        = "16"
	PriceType_PRODUCT_TICKS_IN_THIRTY_SECONDS    = "17"
	PriceType_PRODUCT_TICKS_IN_SIXTY_FORTHS      = "18"
	PriceType_PRODUCT_TICKS_IN_ONE_TWENTY_EIGHTS = "19"
	PriceType_PER_UNIT                           = "2"
	PriceType_FIXED_AMOUNT                       = "3"
	PriceType_DISCOUNT                           = "4"
	PriceType_PREMIUM                            = "5"
	PriceType_SPREAD                             = "6"
	PriceType_TED_PRICE                          = "7"
	PriceType_TED_YIELD                          = "8"
	PriceType_YIELD                              = "9"
)

//Enum values for PriorityIndicator
const (
	PriorityIndicator_PRIORITY_UNCHANGED                      = "0"
	PriorityIndicator_LOST_PRIORITY_AS_RESULT_OF_ORDER_CHANGE = "1"
)

//Enum values for ProcessCode
const (
	ProcessCode_REGULAR              = "0"
	ProcessCode_SOFT_DOLLAR          = "1"
	ProcessCode_STEP_IN              = "2"
	ProcessCode_STEP_OUT             = "3"
	ProcessCode_SOFT_DOLLAR_STEP_IN  = "4"
	ProcessCode_SOFT_DOLLAR_STEP_OUT = "5"
	ProcessCode_PLAN_SPONSOR         = "6"
)

//Enum values for Product
const (
	Product_AGENCY      = "1"
	Product_MORTGAGE    = "10"
	Product_MUNICIPAL   = "11"
	Product_OTHER       = "12"
	Product_FINANCING   = "13"
	Product_COMMODITY   = "2"
	Product_CORPORATE   = "3"
	Product_CURRENCY    = "4"
	Product_EQUITY      = "5"
	Product_GOVERNMENT  = "6"
	Product_INDEX       = "7"
	Product_LOAN        = "8"
	Product_MONEYMARKET = "9"
)

//Enum values for ProgRptReqs
const (
	ProgRptReqs_BUY_SIDE_EXPLICITLY_REQUESTS_STATUS_USING_STATUE_REQUEST                                            = "1"
	ProgRptReqs_SELL_SIDE_PERIODICALLY_SENDS_STATUS_USING_LIST_STATUS_PERIOD_OPTIONALLY_SPECIFIED_IN_PROGRESSPERIOD = "2"
	ProgRptReqs_REAL_TIME_EXECUTION_REPORTS                                                                         = "3"
)

//Enum values for PublishTrdIndicator
const (
	PublishTrdIndicator_NO  = "N"
	PublishTrdIndicator_YES = "Y"
)

//Enum values for PutOrCall
const (
	PutOrCall_PUT  = "0"
	PutOrCall_CALL = "1"
)

//Enum values for QtyType
const (
	QtyType_UNITS                          = "0"
	QtyType_CONTRACTS                      = "1"
	QtyType_UNITS_OF_MEASURE_PER_TIME_UNIT = "2"
)

//Enum values for QuantityType
const (
	QuantityType_SHARES       = "1"
	QuantityType_BONDS        = "2"
	QuantityType_CURRENTFACE  = "3"
	QuantityType_ORIGINALFACE = "4"
	QuantityType_CURRENCY     = "5"
	QuantityType_CONTRACTS    = "6"
	QuantityType_OTHER        = "7"
	QuantityType_PAR          = "8"
)

//Enum values for QuoteAckStatus
const (
	QuoteAckStatus_ACCEPTED                   = "0"
	QuoteAckStatus_CANCELED_FOR_SYMBOL        = "1"
	QuoteAckStatus_CANCELED_FOR_SECURITY_TYPE = "2"
	QuoteAckStatus_CANCELED_FOR_UNDERLYING    = "3"
	QuoteAckStatus_CANCELED_ALL               = "4"
	QuoteAckStatus_REJECTED                   = "5"
)

//Enum values for QuoteCancelType
const (
	QuoteCancelType_CANCEL_FOR_ONE_OR_MORE_SECURITIES        = "1"
	QuoteCancelType_CANCEL_FOR_SECURITY_TYPE                 = "2"
	QuoteCancelType_CANCEL_FOR_UNDERLYING_SECURITY           = "3"
	QuoteCancelType_CANCEL_ALL_QUOTES                        = "4"
	QuoteCancelType_CANCEL_QUOTE_SPECIFIED_IN_QUOTEID        = "5"
	QuoteCancelType_CANCEL_BY_QUOTETYPE                      = "6"
	QuoteCancelType_CANCEL_FOR_SECURITY_ISSUER               = "7"
	QuoteCancelType_CANCEL_FOR_ISSUER_OF_UNDERLYING_SECURITY = "8"
)

//Enum values for QuoteCondition
const (
	QuoteCondition_RESERVED_SAM                        = "0"
	QuoteCondition_NO_ACTIVE_SAM                       = "1"
	QuoteCondition_RESTRICTED                          = "2"
	QuoteCondition_REST_OF_BOOK_VWAP                   = "3"
	QuoteCondition_BETTER_PRICES_IN_CONDITIONAL_ORDERS = "4"
	QuoteCondition_MEDIAN_PRICE                        = "5"
	QuoteCondition_FULL_CURVE                          = "6"
	QuoteCondition_FLAT_CURVE                          = "7"
	QuoteCondition_OPEN_ACTIVE                         = "A"
	QuoteCondition_CLOSED_INACTIVE                     = "B"
	QuoteCondition_EXCHANGE_BEST                       = "C"
	QuoteCondition_CONSOLIDATED_BEST                   = "D"
	QuoteCondition_LOCKED                              = "E"
	QuoteCondition_CROSSED                             = "F"
	QuoteCondition_DEPTH                               = "G"
	QuoteCondition_FAST_TRADING                        = "H"
	QuoteCondition_NON_FIRM                            = "I"
	QuoteCondition_OUTRIGHT_PRICE                      = "J"
	QuoteCondition_IMPLIED_PRICE                       = "K"
	QuoteCondition_MANUAL_SLOW_QUOTE                   = "L"
	QuoteCondition_DEPTH_ON_OFFER                      = "M"
	QuoteCondition_DEPTH_ON_BID                        = "N"
	QuoteCondition_CLOSING                             = "O"
	QuoteCondition_NEWS_DISSEMINATION                  = "P"
	QuoteCondition_TRADING_RANGE                       = "Q"
	QuoteCondition_ORDER_INFLUX                        = "R"
	QuoteCondition_DUE_TO_RELATED                      = "S"
	QuoteCondition_NEWS_PENDING                        = "T"
	QuoteCondition_ADDITIONAL_INFO                     = "U"
	QuoteCondition_ADDITIONAL_INFO_DUE_TO_RELATED      = "V"
	QuoteCondition_RESUME                              = "W"
	QuoteCondition_VIEW_OF_COMMON                      = "X"
	QuoteCondition_VOLUME_ALERT                        = "Y"
	QuoteCondition_ORDER_IMBALANCE                     = "Z"
	QuoteCondition_EQUIPMENT_CHANGEOVER                = "a"
	QuoteCondition_NO_OPEN                             = "b"
	QuoteCondition_REGULAR_ETH                         = "c"
	QuoteCondition_AUTOMATIC_EXECUTION                 = "d"
	QuoteCondition_AUTOMATIC_EXECUTION_ETH             = "e"
	QuoteCondition_FAST_MARKET_ETH                     = "f "
	QuoteCondition_INACTIVE_ETH                        = "g"
	QuoteCondition_ROTATION                            = "h"
	QuoteCondition_ROTATION_ETH                        = "i"
	QuoteCondition_HALT                                = "j"
	QuoteCondition_HALT_ETH                            = "k"
	QuoteCondition_DUE_TO_NEWS_DISSEMINATION           = "l"
	QuoteCondition_DUE_TO_NEWS_PENDING                 = "m"
	QuoteCondition_TRADING_RESUME                      = "n"
	QuoteCondition_OUT_OF_SEQUENCE                     = "o"
	QuoteCondition_BID_SPECIALIST                      = "p"
	QuoteCondition_OFFER_SPECIALIST                    = "q"
	QuoteCondition_BID_OFFER_SPECIALIST                = "r"
	QuoteCondition_END_OF_DAY_SAM                      = "s"
	QuoteCondition_FORBIDDEN_SAM                       = "t"
	QuoteCondition_FROZEN_SAM                          = "u"
	QuoteCondition_PREOPENING_SAM                      = "v"
	QuoteCondition_OPENING_SAM                         = "w"
	QuoteCondition_OPEN_SAM                            = "x"
	QuoteCondition_SURVEILLANCE_SAM                    = "y"
	QuoteCondition_SUSPENDED_SAM                       = "z"
)

//Enum values for QuoteEntryRejectReason
const (
	QuoteEntryRejectReason_UNKNOWN_SYMBOL                   = "1"
	QuoteEntryRejectReason_EXHCNAGE                         = "2"
	QuoteEntryRejectReason_QUOTE_EXCEEDS_LIMIT              = "3"
	QuoteEntryRejectReason_TOO_LATE_TO_ENTER                = "4"
	QuoteEntryRejectReason_UNKNOWN_QUOTE                    = "5"
	QuoteEntryRejectReason_DUPLICATE_QUOTE                  = "6"
	QuoteEntryRejectReason_INVALID_BID_ASK_SPREAD           = "7"
	QuoteEntryRejectReason_INVALID_PRICE                    = "8"
	QuoteEntryRejectReason_NOT_AUTHORIZED_TO_QUOTE_SECURITY = "9"
	QuoteEntryRejectReason_OTHER                            = "99"
)

//Enum values for QuoteEntryStatus
const (
	QuoteEntryStatus_ACCEPTED                     = "0"
	QuoteEntryStatus_LOCKED_MARKET_WARNING        = "12"
	QuoteEntryStatus_CROSS_MARKET_WARNING         = "13"
	QuoteEntryStatus_CANCELED_DUE_TO_LOCK_MARKET  = "14"
	QuoteEntryStatus_CANCELED_DUE_TO_CROSS_MARKET = "15"
	QuoteEntryStatus_ACTIVE                       = "16"
	QuoteEntryStatus_REJECTED                     = "5"
	QuoteEntryStatus_REMOVED_FROM_MARKET          = "6"
	QuoteEntryStatus_EXPIRED                      = "7"
)

//Enum values for QuotePriceType
const (
	QuotePriceType_PERCENT      = "1"
	QuotePriceType_YIELD        = "10"
	QuotePriceType_PER_SHARE    = "2"
	QuotePriceType_FIXED_AMOUNT = "3"
	QuotePriceType_DISCOUNT     = "4"
	QuotePriceType_PREMIUM      = "5"
	QuotePriceType_SPREAD       = "6"
	QuotePriceType_TED_PRICE    = "7"
	QuotePriceType_TED_YIELD    = "8"
	QuotePriceType_YIELD_SPREAD = "9"
)

//Enum values for QuoteRejectReason
const (
	QuoteRejectReason_UNKNOWN_SYMBOL                                   = "1"
	QuoteRejectReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND                 = "10"
	QuoteRejectReason_QUOTE_LOCKED                                     = "11"
	QuoteRejectReason_INVALID_OR_UNKNOWN_SECURITY_ISSUER               = "12"
	QuoteRejectReason_INVALID_OR_UNKNOWN_ISSUER_OF_UNDERLYING_SECURITY = "13"
	QuoteRejectReason_EXCHANGE                                         = "2"
	QuoteRejectReason_QUOTE_REQUEST_EXCEEDS_LIMIT                      = "3"
	QuoteRejectReason_TOO_LATE_TO_ENTER                                = "4"
	QuoteRejectReason_UNKNOWN_QUOTE                                    = "5"
	QuoteRejectReason_DUPLICATE_QUOTE                                  = "6"
	QuoteRejectReason_INVALID_BID_ASK_SPREAD                           = "7"
	QuoteRejectReason_INVALID_PRICE                                    = "8"
	QuoteRejectReason_NOT_AUTHORIZED_TO_QUOTE_SECURITY                 = "9"
	QuoteRejectReason_OTHER                                            = "99"
)

//Enum values for QuoteRequestRejectReason
const (
	QuoteRequestRejectReason_UNKNOWN_SYMBOL                  = "1"
	QuoteRequestRejectReason_PASS                            = "10"
	QuoteRequestRejectReason_INSUFFICIENT_CREDIT             = "11"
	QuoteRequestRejectReason_EXCHANGE                        = "2"
	QuoteRequestRejectReason_QUOTE_REQUEST_EXCEEDS_LIMIT     = "3"
	QuoteRequestRejectReason_TOO_LATE_TO_ENTER               = "4"
	QuoteRequestRejectReason_INVALID_PRICE                   = "5"
	QuoteRequestRejectReason_NOT_AUTHORIZED_TO_REQUEST_QUOTE = "6"
	QuoteRequestRejectReason_NO_MATCH_FOR_INQUIRY            = "7"
	QuoteRequestRejectReason_NO_MARKET_FOR_INSTRUMENT        = "8"
	QuoteRequestRejectReason_NO_INVENTORY                    = "9"
	QuoteRequestRejectReason_OTHER                           = "99"
)

//Enum values for QuoteRequestType
const (
	QuoteRequestType_MANUAL    = "1"
	QuoteRequestType_AUTOMATIC = "2"
)

//Enum values for QuoteRespType
const (
	QuoteRespType_HIT_LIFT  = "1"
	QuoteRespType_COUNTER   = "2"
	QuoteRespType_EXPIRED   = "3"
	QuoteRespType_COVER     = "4"
	QuoteRespType_DONE_AWAY = "5"
	QuoteRespType_PASS      = "6"
	QuoteRespType_END_TRADE = "7"
	QuoteRespType_TIMED_OUT = "8"
)

//Enum values for QuoteResponseLevel
const (
	QuoteResponseLevel_NO_ACKNOWLEDGEMENT                            = "0"
	QuoteResponseLevel_ACKNOWLEDGE_ONLY_NEGATIVE_OR_ERRONEOUS_QUOTES = "1"
	QuoteResponseLevel_ACKNOWLEDGE_EACH_QUOTE_MESSAGE                = "2"
	QuoteResponseLevel_SUMMARY_ACKNOWLEDGEMENT                       = "3"
)

//Enum values for QuoteStatus
const (
	QuoteStatus_ACCEPTED                        = "0"
	QuoteStatus_CANCEL_FOR_SYMBOL               = "1"
	QuoteStatus_PENDING                         = "10"
	QuoteStatus_PASS                            = "11"
	QuoteStatus_LOCKED_MARKET_WARNING           = "12"
	QuoteStatus_CROSS_MARKET_WARNING            = "13"
	QuoteStatus_CANCELED_DUE_TO_LOCK_MARKET     = "14"
	QuoteStatus_CANCELED_DUE_TO_CROSS_MARKET    = "15"
	QuoteStatus_ACTIVE                          = "16"
	QuoteStatus_CANCELED                        = "17"
	QuoteStatus_UNSOLICITED_QUOTE_REPLENISHMENT = "18"
	QuoteStatus_PENDING_END_TRADE               = "19"
	QuoteStatus_CANCELED_FOR_SECURITY_TYPE      = "2"
	QuoteStatus_TOO_LATE_TO_END                 = "20"
	QuoteStatus_CANCELED_FOR_UNDERLYING         = "3"
	QuoteStatus_CANCELED_ALL                    = "4"
	QuoteStatus_REJECTED                        = "5"
	QuoteStatus_REMOVED_FROM_MARKET             = "6"
	QuoteStatus_EXPIRED                         = "7"
	QuoteStatus_QUERY                           = "8"
	QuoteStatus_QUOTE_NOT_FOUND                 = "9"
)

//Enum values for QuoteType
const (
	QuoteType_INDICATIVE           = "0"
	QuoteType_TRADEABLE            = "1"
	QuoteType_RESTRICTED_TRADEABLE = "2"
	QuoteType_COUNTER              = "3"
)

//Enum values for RateSource
const (
	RateSource_BLOOMBERG = "0"
	RateSource_REUTERS   = "1"
	RateSource_TELERATE  = "2"
	RateSource_OTHER     = "99"
)

//Enum values for RateSourceType
const (
	RateSourceType_PRIMARY   = "0"
	RateSourceType_SECONDARY = "1"
)

//Enum values for RefOrdIDReason
const (
	RefOrdIDReason_GTC_FROM_PREVIOUS_DAY  = "0"
	RefOrdIDReason_PARTIAL_FILL_REMAINING = "1"
	RefOrdIDReason_ORDER_CHANGED          = "2"
)

//Enum values for RefOrderIDSource
const (
	RefOrderIDSource_SECONDARYORDERID  = "0"
	RefOrderIDSource_ORDERID           = "1"
	RefOrderIDSource_MDENTRYID         = "2"
	RefOrderIDSource_QUOTEENTRYID      = "3"
	RefOrderIDSource_ORIGINAL_ORDER_ID = "4"
)

//Enum values for RegistRejReasonCode
const (
	RegistRejReasonCode_INVALID_UNACCEPTABLE_ACCOUNT_TYPE                  = "1"
	RegistRejReasonCode_INVALID_UNACEEPTABLE_INVESTOR_ID_SOURCE            = "10"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_DATE_OF_BIRTH                 = "11"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_INVESTOR_COUNTRY_OF_RESIDENCE = "12"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_NO_DISTRIB_INSTNS             = "13"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_DISTRIB_PERCENTAGE            = "14"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_DISTRIB_PAYMENT_METHOD        = "15"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_CASH_DISTRIB_AGENT_ACCT_NAME  = "16"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_CASH_DISTRIB_AGENT_CODE       = "17"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_CASH_DISTRIB_AGENT_ACCT_NUM   = "18"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_TAX_EXEMPT_TYPE               = "2"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_OWNERSHIP_TYPE                = "3"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_NO_REG_DETAILS                = "4"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_REG_SEQ_NO                    = "5"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_REG_DETAILS                   = "6"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_MAILING_DETAILS               = "7"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_MAILING_INSTRUCTIONS          = "8"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_INVESTOR_ID                   = "9"
	RegistRejReasonCode_OTHER                                              = "99"
)

//Enum values for RegistStatus
const (
	RegistStatus_ACCEPTED = "A"
	RegistStatus_HELD     = "H"
	RegistStatus_REMINDER = "N"
	RegistStatus_REJECTED = "R"
)

//Enum values for RegistTransType
const (
	RegistTransType_NEW     = "0"
	RegistTransType_REPLACE = "1"
	RegistTransType_CANCEL  = "2"
)

//Enum values for ReportToExch
const (
	ReportToExch_NO  = "N"
	ReportToExch_YES = "Y"
)

//Enum values for ResetSeqNumFlag
const (
	ResetSeqNumFlag_NO  = "N"
	ResetSeqNumFlag_YES = "Y"
)

//Enum values for RespondentType
const (
	RespondentType_ALL_MARKET_PARTICIPANTS       = "1"
	RespondentType_SPECIFIED_MARKET_PARTICIPANTS = "2"
	RespondentType_ALL_MARKET_MAKERS             = "3"
	RespondentType_PRIMARY_MARKET_MAKER          = "4"
)

//Enum values for ResponseTransportType
const (
	ResponseTransportType_INBAND      = "0"
	ResponseTransportType_OUT_OF_BAND = "1"
)

//Enum values for RestructuringType
const (
	RestructuringType_FULL_RESTRUCTURING         = "FR"
	RestructuringType_MODIFIED_MOD_RESTRUCTURING = "MM"
	RestructuringType_MODIFIED_RESTRUCTURING     = "MR"
	RestructuringType_NO_RESTRUCTURING_SPECIFIED = "XR"
)

//Enum values for RiskInstrumentOperator
const (
	RiskInstrumentOperator_INCLUDE = "1"
	RiskInstrumentOperator_EXCLUDE = "2"
)

//Enum values for RiskLimitType
const (
	RiskLimitType_GROSS_LIMIT = "1"
	RiskLimitType_NET_LIMIT   = "2"
	RiskLimitType_EXPOSURE    = "3"
	RiskLimitType_LONG_LIMIT  = "4"
	RiskLimitType_SHORT_LIMIT = "5"
)

//Enum values for RoundingDirection
const (
	RoundingDirection_ROUND_TO_NEAREST = "0"
	RoundingDirection_ROUND_DOWN       = "1"
	RoundingDirection_ROUND_UP         = "2"
)

//Enum values for RoutingType
const (
	RoutingType_TARGET_FIRM = "1"
	RoutingType_TARGET_LIST = "2"
	RoutingType_BLOCK_FIRM  = "3"
	RoutingType_BLOCK_LIST  = "4"
)

//Enum values for Rule80A
const (
	Rule80A_AGENCY_SINGLE_ORDER                                                                                        = "A"
	Rule80A_SHORT_EXEMPT_TRANSACTION_B                                                                                 = "B"
	Rule80A_PROPRIETARY_NON_ALGORITHMIC_PROGRAM_TRADE                                                                  = "C"
	Rule80A_PROGRAM_ORDER_INDEX_ARB_FOR_MEMBER_FIRM_ORG                                                                = "D"
	Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_PRINCIPAL                                                                     = "E"
	Rule80A_SHORT_EXEMPT_TRANSACTION_F                                                                                 = "F"
	Rule80A_SHORT_EXEMPT_TRANSACTION_H                                                                                 = "H"
	Rule80A_INDIVIDUAL_INVESTOR_SINGLE_ORDER                                                                           = "I"
	Rule80A_PROPRIETARY_ALGORITHMIC_PROGRAM_TRADING                                                                    = "J"
	Rule80A_AGENCY_ALGORITHMIC_PROGRAM_TRADING                                                                         = "K"
	Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_MEMBER_COMPETING_MARKET_MAKER_AFFLIATED_WITH_THE_FIRM_CLEARING_THE_TRADE      = "L"
	Rule80A_PROGRAM_ORDER_INDEX_ARB_FOR_OTHER_MEMBER                                                                   = "M"
	Rule80A_AGENT_FOR_OTHER_MEMBER_NON_ALGORITHMIC_PROGRAM_TRADE                                                       = "N"
	Rule80A_PROPRIETARY_TRANSACTIONS_FOR_COMPETING_MARKET_MAKER_THAT_IS_AFFILIATED_WITH_THE_CLEARING_MEMBER            = "O"
	Rule80A_PRINCIPAL                                                                                                  = "P"
	Rule80A_TRANSACTIONS_FOR_THE_ACCOUNT_OF_A_NON_MEMBER_COMPTING_MARKET_MAKER                                         = "R"
	Rule80A_SPECIALIST_TRADES                                                                                          = "S"
	Rule80A_TRANSACTIONS_FOR_THE_ACCOUNT_OF_AN_UNAFFILIATED_MEMBERS_COMPETING_MARKET_MAKER                             = "T"
	Rule80A_AGENCY_INDEX_ARBITRAGE                                                                                     = "U"
	Rule80A_ALL_OTHER_ORDERS_AS_AGENT_FOR_OTHER_MEMBER                                                                 = "W"
	Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_MEMBER_COMPETING_MARKET_MAKER_NOT_AFFILIATED_WITH_THE_FIRM_CLEARING_THE_TRADE = "X"
	Rule80A_AGENCY_NON_ALGORITHMIC_PROGRAM_TRADE                                                                       = "Y"
	Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_NON_MEMBER_COMPETING_MARKET_MAKER                                             = "Z"
)

//Enum values for Scope
const (
	Scope_LOCAL_MARKET = "1"
	Scope_NATIONAL     = "2"
	Scope_GLOBAL       = "3"
)

//Enum values for SecDefStatus
const (
	SecDefStatus_PENDING_APPROVAL           = "0"
	SecDefStatus_APPROVED                   = "1"
	SecDefStatus_REJECTED                   = "2"
	SecDefStatus_UNAUTHORIZED_REQUEST       = "3"
	SecDefStatus_INVALID_DEFINITION_REQUEST = "4"
)

//Enum values for SecurityIDSource
const (
	SecurityIDSource_CUSIP                            = "1"
	SecurityIDSource_SEDOL                            = "2"
	SecurityIDSource_QUIK                             = "3"
	SecurityIDSource_ISIN_NUMBER                      = "4"
	SecurityIDSource_RIC_CODE                         = "5"
	SecurityIDSource_ISO_CURRENCY_CODE                = "6"
	SecurityIDSource_ISO_COUNTRY_CODE                 = "7"
	SecurityIDSource_EXCHANGE_SYMBOL                  = "8"
	SecurityIDSource_CONSOLIDATED_TAPE_ASSOCIATION    = "9"
	SecurityIDSource_BLOOMBERG_SYMBOL                 = "A"
	SecurityIDSource_WERTPAPIER                       = "B"
	SecurityIDSource_DUTCH                            = "C"
	SecurityIDSource_VALOREN                          = "D"
	SecurityIDSource_SICOVAM                          = "E"
	SecurityIDSource_BELGIAN                          = "F"
	SecurityIDSource_COMMON                           = "G"
	SecurityIDSource_CLEARING_HOUSE                   = "H"
	SecurityIDSource_ISDA_FPML_PRODUCT_SPECIFICATION  = "I"
	SecurityIDSource_OPTION_PRICE_REPORTING_AUTHORITY = "J"
	SecurityIDSource_ISDA_FPML_PRODUCT_URL            = "K"
	SecurityIDSource_LETTER_OF_CREDIT                 = "L"
	SecurityIDSource_MARKETPLACE_ASSIGNED_IDENTIFIER  = "M"
)

//Enum values for SecurityListRequestType
const (
	SecurityListRequestType_SYMBOL                                    = "0"
	SecurityListRequestType_SECURITYTYPE_AND_OR_CFICODE               = "1"
	SecurityListRequestType_PRODUCT                                   = "2"
	SecurityListRequestType_TRADINGSESSIONID                          = "3"
	SecurityListRequestType_ALL_SECURITIES                            = "4"
	SecurityListRequestType_MARKETID_OR_MARKETID_PLUS_MARKETSEGMENTID = "5"
)

//Enum values for SecurityListType
const (
	SecurityListType_INDUSTRY_CLASSIFICATION = "1"
	SecurityListType_TRADING_LIST            = "2"
	SecurityListType_MARKET                  = "3"
	SecurityListType_NEWSPAPER_LIST          = "4"
)

//Enum values for SecurityListTypeSource
const (
	SecurityListTypeSource_ICB   = "1"
	SecurityListTypeSource_NAICS = "2"
	SecurityListTypeSource_GICS  = "3"
)

//Enum values for SecurityRequestResult
const (
	SecurityRequestResult_VALID_REQUEST                                      = "0"
	SecurityRequestResult_INVALID_OR_UNSUPPORTED_REQUEST                     = "1"
	SecurityRequestResult_NO_INSTRUMENTS_FOUND_THAT_MATCH_SELECTION_CRITERIA = "2"
	SecurityRequestResult_NOT_AUTHORIZED_TO_RETRIEVE_INSTRUMENT_DATA         = "3"
	SecurityRequestResult_INSTRUMENT_DATA_TEMPORARILY_UNAVAILABLE            = "4"
	SecurityRequestResult_REQUEST_FOR_INSTRUMENT_DATA_NOT_SUPPORTED          = "5"
)

//Enum values for SecurityRequestType
const (
	SecurityRequestType_REQUEST_SECURITY_IDENTITY_AND_SPECIFICATIONS              = "0"
	SecurityRequestType_REQUEST_SECURITY_IDENTITY_FOR_THE_SPECIFICATIONS_PROVIDED = "1"
	SecurityRequestType_REQUEST_LIST_SECURITY_TYPES                               = "2"
	SecurityRequestType_REQUEST_LIST_SECURITIES                                   = "3"
	SecurityRequestType_SYMBOL                                                    = "4"
	SecurityRequestType_SECURITYTYPE_AND_OR_CFICODE                               = "5"
	SecurityRequestType_PRODUCT                                                   = "6"
	SecurityRequestType_TRADINGSESSIONID                                          = "7"
	SecurityRequestType_ALL_SECURITIES                                            = "8"
	SecurityRequestType_MARKETID_OR_MARKETID_PLUS_MARKETSEGMENTID                 = "9"
)

//Enum values for SecurityResponseType
const (
	SecurityResponseType_ACCEPT_SECURITY_PROPOSAL_AS_IS                                      = "1"
	SecurityResponseType_ACCEPT_SECURITY_PROPOSAL_WITH_REVISIONS_AS_INDICATED_IN_THE_MESSAGE = "2"
	SecurityResponseType_LIST_OF_SECURITY_TYPES_RETURNED_PER_REQUEST                         = "3"
	SecurityResponseType_LIST_OF_SECURITIES_RETURNED_PER_REQUEST                             = "4"
	SecurityResponseType_REJECT_SECURITY_PROPOSAL                                            = "5"
	SecurityResponseType_CANNOT_MATCH_SELECTION_CRITERIA                                     = "6"
)

//Enum values for SecurityStatus
const (
	SecurityStatus_ACTIVE   = "1"
	SecurityStatus_INACTIVE = "2"
)

//Enum values for SecurityTradingEvent
const (
	SecurityTradingEvent_ORDER_IMBALANCE_AUCTION_IS_EXTENDED = "1"
	SecurityTradingEvent_TRADING_RESUMES                     = "2"
	SecurityTradingEvent_PRICE_VOLATILITY_INTERRUPTION       = "3"
	SecurityTradingEvent_CHANGE_OF_TRADING_SESSION           = "4"
	SecurityTradingEvent_CHANGE_OF_TRADING_SUBSESSION        = "5"
	SecurityTradingEvent_CHANGE_OF_SECURITY_TRADING_STATUS   = "6"
	SecurityTradingEvent_CHANGE_OF_BOOK_TYPE                 = "7"
	SecurityTradingEvent_CHANGE_OF_MARKET_DEPTH              = "8"
)

//Enum values for SecurityTradingStatus
const (
	SecurityTradingStatus_OPENING_DELAY                  = "1"
	SecurityTradingStatus_MARKET_ON_CLOSE_IMBALANCE_SELL = "10"
	SecurityTradingStatus_11                             = "11"
	SecurityTradingStatus_NO_MARKET_IMBALANCE            = "12"
	SecurityTradingStatus_NO_MARKET_ON_CLOSE_IMBALANCE   = "13"
	SecurityTradingStatus_ITS_PRE_OPENING                = "14"
	SecurityTradingStatus_NEW_PRICE_INDICATION           = "15"
	SecurityTradingStatus_TRADE_DISSEMINATION_TIME       = "16"
	SecurityTradingStatus_READY_TO_TRADE                 = "17"
	SecurityTradingStatus_NOT_AVAILABLE_FOR_TRADING      = "18"
	SecurityTradingStatus_NOT_TRADED_ON_THIS_MARKET      = "19"
	SecurityTradingStatus_TRADING_HALT                   = "2"
	SecurityTradingStatus_UNKNOWN_OR_INVALID             = "20"
	SecurityTradingStatus_PRE_OPEN                       = "21"
	SecurityTradingStatus_OPENING_ROTATION               = "22"
	SecurityTradingStatus_FAST_MARKET                    = "23"
	SecurityTradingStatus_PRE_CROSS                      = "24"
	SecurityTradingStatus_CROSS                          = "25"
	SecurityTradingStatus_POST_CLOSE                     = "26"
	SecurityTradingStatus_RESUME                         = "3"
	SecurityTradingStatus_NO_OPEN                        = "4"
	SecurityTradingStatus_PRICE_INDICATION               = "5"
	SecurityTradingStatus_TRADING_RANGE_INDICATION       = "6"
	SecurityTradingStatus_MARKET_IMBALANCE_BUY           = "7"
	SecurityTradingStatus_MARKET_IMBALANCE_SELL          = "8"
	SecurityTradingStatus_MARKET_ON_CLOSE_IMBALANCE_BUY  = "9"
)

//Enum values for SecurityType
const (
	SecurityType_WILDCARD_ENTRY_FOR_USE_ON_SECURITY_DEFINITION_REQUEST = "?"
	SecurityType_ASSET_BACKED_SECURITIES                               = "ABS"
	SecurityType_AMENDED_RESTATED                                      = "AMENDED"
	SecurityType_OTHER_ANTICIPATION_NOTES                              = "AN"
	SecurityType_BANKERS_ACCEPTANCE                                    = "BA"
	SecurityType_BANK_DEPOSITORY_NOTE                                  = "BDN"
	SecurityType_BANK_NOTES                                            = "BN"
	SecurityType_BILL_OF_EXCHANGES                                     = "BOX"
	SecurityType_BRADY_BOND                                            = "BRADY"
	SecurityType_BRIDGE_LOAN                                           = "BRIDGE"
	SecurityType_BUY_SELLBACK                                          = "BUYSELL"
	SecurityType_CANADIAN_MONEY_MARKETS                                = "CAMM"
	SecurityType_CANADIAN_TREASURY_NOTES                               = "CAN"
	SecurityType_CASH                                                  = "CASH"
	SecurityType_CONVERTIBLE_BOND                                      = "CB"
	SecurityType_CERTIFICATE_OF_DEPOSIT                                = "CD"
	SecurityType_CREDIT_DEFAULT_SWAP                                   = "CDS"
	SecurityType_CALL_LOANS                                            = "CL"
	SecurityType_CANADIAN_MORTGAGE_BONDS                               = "CMB"
	SecurityType_CORP_MORTGAGE_BACKED_SECURITIES                       = "CMBS"
	SecurityType_COLLATERALIZED_MORTGAGE_OBLIGATION                    = "CMO"
	SecurityType_CERTIFICATE_OF_OBLIGATION                             = "COFO"
	SecurityType_CERTIFICATE_OF_PARTICIPATION                          = "COFP"
	SecurityType_CORPORATE_BOND                                        = "CORP"
	SecurityType_COMMERCIAL_PAPER                                      = "CP"
	SecurityType_CORPORATE_PRIVATE_PLACEMENT                           = "CPP"
	SecurityType_COMMON_STOCK                                          = "CS"
	SecurityType_CANADIAN_TREASURY_BILLS                               = "CTB"
	SecurityType_DEFAULTED                                             = "DEFLTED"
	SecurityType_DEBTOR_IN_POSSESSION                                  = "DINP"
	SecurityType_DEPOSIT_NOTES                                         = "DN"
	SecurityType_DUAL_CURRENCY                                         = "DUAL"
	SecurityType_EURO_CERTIFICATE_OF_DEPOSIT                           = "EUCD"
	SecurityType_EURO_CORPORATE_BOND                                   = "EUCORP"
	SecurityType_EURO_COMMERCIAL_PAPER                                 = "EUCP"
	SecurityType_EURO_CORPORATE_FLOATING_RATE_NOTES                    = "EUFRN"
	SecurityType_EURO_SOVEREIGNS                                       = "EUSOV"
	SecurityType_EURO_SUPRANATIONAL_COUPONS                            = "EUSUPRA"
	SecurityType_FEDERAL_AGENCY_COUPON                                 = "FAC"
	SecurityType_FEDERAL_AGENCY_DISCOUNT_NOTE                          = "FADN"
	SecurityType_FEDERAL_HOUSING_AUTHORITY                             = "FHA"
	SecurityType_FEDERAL_HOME_LOAN                                     = "FHL"
	SecurityType_FEDERAL_NATIONAL_MORTGAGE_ASSOCIATION                 = "FN"
	SecurityType_FOREIGN_EXCHANGE_CONTRACT                             = "FOR"
	SecurityType_FORWARD                                               = "FORWARD"
	SecurityType_US_CORPORATE_FLOATING_RATE_NOTES                      = "FRN"
	SecurityType_FUTURE                                                = "FUT"
	SecurityType_FX_FORWARD                                            = "FXFWD"
	SecurityType_NON_DELIVERABLE_FORWARD                               = "FXNDF"
	SecurityType_FX_SPOT                                               = "FXSPOT"
	SecurityType_FX_SWAP                                               = "FXSWAP"
	SecurityType_GOVERNMENT_NATIONAL_MORTGAGE_ASSOCIATION              = "GN"
	SecurityType_GENERAL_OBLIGATION_BONDS                              = "GO"
	SecurityType_TREASURIES_PLUS_AGENCY_DEBENTURE                      = "GOVT"
	SecurityType_IOETTE_MORTGAGE                                       = "IET"
	SecurityType_INTEREST_RATE_SWAP                                    = "IRS"
	SecurityType_LETTER_OF_CREDIT                                      = "LOFC"
	SecurityType_LIQUIDITY_NOTE                                        = "LQN"
	SecurityType_MATURED                                               = "MATURED"
	SecurityType_MORTGAGE_BACKED_SECURITIES                            = "MBS"
	SecurityType_MUTUAL_FUND                                           = "MF"
	SecurityType_MORTGAGE_INTEREST_ONLY                                = "MIO"
	SecurityType_MULTILEG_INSTRUMENT                                   = "MLEG"
	SecurityType_MORTGAGE_PRINCIPAL_ONLY                               = "MPO"
	SecurityType_MORTGAGE_PRIVATE_PLACEMENT                            = "MPP"
	SecurityType_MISCELLANEOUS_PASS_THROUGH                            = "MPT"
	SecurityType_MANDATORY_TENDER                                      = "MT"
	SecurityType_MEDIUM_TERM_NOTES                                     = "MTN"
	SecurityType_MUNICIPAL_BOND                                        = "MUNI"
	SecurityType_NO_SECURITY_TYPE                                      = "NONE"
	SecurityType_OVERNIGHT                                             = "ONITE"
	SecurityType_OPTIONS_ON_COMBO                                      = "OOC"
	SecurityType_OPTIONS_ON_FUTURES                                    = "OOF"
	SecurityType_OPTIONS_ON_PHYSICAL                                   = "OOP"
	SecurityType_OPTION                                                = "OPT"
	SecurityType_PRIVATE_EXPORT_FUNDING                                = "PEF"
	SecurityType_PFANDBRIEFE                                           = "PFAND"
	SecurityType_PROMISSORY_NOTE                                       = "PN"
	SecurityType_AGENCY_POOLS                                          = "POOL"
	SecurityType_CANADIAN_PROVINCIAL_BONDS                             = "PROV"
	SecurityType_PREFERRED_STOCK                                       = "PS"
	SecurityType_PLAZOS_FIJOS                                          = "PZFJ"
	SecurityType_REVENUE_ANTICIPATION_NOTE                             = "RAN"
	SecurityType_REPLACED                                              = "REPLACD"
	SecurityType_REPURCHASE                                            = "REPO"
	SecurityType_RETIRED                                               = "RETIRED"
	SecurityType_REVENUE_BONDS                                         = "REV"
	SecurityType_REPURCHASE_AGREEMENT                                  = "RP"
	SecurityType_REVOLVER_LOAN                                         = "RVLV"
	SecurityType_REVOLVER_TERM_LOAN                                    = "RVLVTRM"
	SecurityType_REVERSE_REPURCHASE_AGREEMENT                          = "RVRP"
	SecurityType_SECURITIES_LOAN                                       = "SECLOAN"
	SecurityType_SECURITIES_PLEDGE                                     = "SECPLEDGE"
	SecurityType_STUDENT_LOAN_MARKETING_ASSOCIATION                    = "SL"
	SecurityType_SECURED_LIQUIDITY_NOTE                                = "SLQN"
	SecurityType_SPECIAL_ASSESSMENT                                    = "SPCLA"
	SecurityType_SPECIAL_OBLIGATION                                    = "SPCLO"
	SecurityType_SPECIAL_TAX                                           = "SPCLT"
	SecurityType_SHORT_TERM_LOAN_NOTE                                  = "STN"
	SecurityType_STRUCTURED_NOTES                                      = "STRUCT"
	SecurityType_USD_SUPRANATIONAL_COUPONS                             = "SUPRA"
	SecurityType_SWING_LINE_FACILITY                                   = "SWING"
	SecurityType_TAX_ANTICIPATION_NOTE                                 = "TAN"
	SecurityType_TAX_ALLOCATION                                        = "TAXA"
	SecurityType_TREASURY_BILL                                         = "TB"
	SecurityType_TO_BE_ANNOUNCED                                       = "TBA"
	SecurityType_US_TREASURY_BILL_TBILL                                = "TBILL"
	SecurityType_US_TREASURY_BOND                                      = "TBOND"
	SecurityType_PRINCIPAL_STRIP_OF_A_CALLABLE_BOND_OR_NOTE            = "TCAL"
	SecurityType_TIME_DEPOSIT                                          = "TD"
	SecurityType_TAX_EXEMPT_COMMERCIAL_PAPER                           = "TECP"
	SecurityType_TERM_LOAN                                             = "TERM"
	SecurityType_INTEREST_STRIP_FROM_ANY_BOND_OR_NOTE                  = "TINT"
	SecurityType_TREASURY_INFLATION_PROTECTED_SECURITIES               = "TIPS"
	SecurityType_TERM_LIQUIDITY_NOTE                                   = "TLQN"
	SecurityType_TAXABLE_MUNICIPAL_CP                                  = "TMCP"
	SecurityType_US_TREASURY_NOTE_TNOTE                                = "TNOTE"
	SecurityType_PRINCIPAL_STRIP_FROM_A_NON_CALLABLE_BOND_OR_NOTE      = "TPRN"
	SecurityType_TAX_REVENUE_ANTICIPATION_NOTE                         = "TRAN"
	SecurityType_US_TREASURY_NOTE_UST                                  = "UST"
	SecurityType_US_TREASURY_BILL_USTB                                 = "USTB"
	SecurityType_VARIABLE_RATE_DEMAND_NOTE                             = "VRDN"
	SecurityType_WARRANT                                               = "WAR"
	SecurityType_WITHDRAWN                                             = "WITHDRN"
	SecurityType_WILDCARD_ENTRY                                        = "WLD"
	SecurityType_EXTENDED_COMM_NOTE                                    = "XCN"
	SecurityType_INDEXED_LINKED                                        = "XLINKD"
	SecurityType_YANKEE_CORPORATE_BOND                                 = "YANK"
	SecurityType_YANKEE_CERTIFICATE_OF_DEPOSIT                         = "YCD"
	SecurityType_CATS_TIGERS_LIONS                                     = "ZOO"
)

//Enum values for SecurityUpdateAction
const (
	SecurityUpdateAction_ADD    = "A"
	SecurityUpdateAction_DELETE = "D"
	SecurityUpdateAction_MODIFY = "M"
)

//Enum values for Seniority
const (
	Seniority_SUBORDINATED   = "SB"
	Seniority_SENIOR_SECURED = "SD"
	Seniority_SENIOR         = "SR"
)

//Enum values for SessionRejectReason
const (
	SessionRejectReason_INVALID_TAG_NUMBER                             = "0"
	SessionRejectReason_REQUIRED_TAG_MISSING                           = "1"
	SessionRejectReason_SENDINGTIME_ACCURACY_PROBLEM                   = "10"
	SessionRejectReason_INVALID_MSGTYPE                                = "11"
	SessionRejectReason_XML_VALIDATION_ERROR                           = "12"
	SessionRejectReason_TAG_APPEARS_MORE_THAN_ONCE                     = "13"
	SessionRejectReason_TAG_SPECIFIED_OUT_OF_REQUIRED_ORDER            = "14"
	SessionRejectReason_REPEATING_GROUP_FIELDS_OUT_OF_ORDER            = "15"
	SessionRejectReason_INCORRECT_NUMINGROUP_COUNT_FOR_REPEATING_GROUP = "16"
	SessionRejectReason_NON_DATA_VALUE_INCLUDES_FIELD_DELIMITER        = "17"
	SessionRejectReason_INVALID_UNSUPPORTED_APPLICATION_VERSION        = "18"
	SessionRejectReason_TAG_NOT_DEFINED_FOR_THIS_MESSAGE_TYPE          = "2"
	SessionRejectReason_UNDEFINED_TAG                                  = "3"
	SessionRejectReason_TAG_SPECIFIED_WITHOUT_A_VALUE                  = "4"
	SessionRejectReason_VALUE_IS_INCORRECT                             = "5"
	SessionRejectReason_INCORRECT_DATA_FORMAT_FOR_VALUE                = "6"
	SessionRejectReason_DECRYPTION_PROBLEM                             = "7"
	SessionRejectReason_SIGNATURE_PROBLEM                              = "8"
	SessionRejectReason_COMPID_PROBLEM                                 = "9"
	SessionRejectReason_OTHER                                          = "99"
)

//Enum values for SessionStatus
const (
	SessionStatus_SESSION_ACTIVE                                   = "0"
	SessionStatus_SESSION_PASSWORD_CHANGED                         = "1"
	SessionStatus_SESSION_PASSWORD_DUE_TO_EXPIRE                   = "2"
	SessionStatus_NEW_SESSION_PASSWORD_DOES_NOT_COMPLY_WITH_POLICY = "3"
	SessionStatus_SESSION_LOGOUT_COMPLETE                          = "4"
	SessionStatus_INVALID_USERNAME_OR_PASSWORD                     = "5"
	SessionStatus_ACCOUNT_LOCKED                                   = "6"
	SessionStatus_LOGONS_ARE_NOT_ALLOWED_AT_THIS_TIME              = "7"
	SessionStatus_PASSWORD_EXPIRED                                 = "8"
)

//Enum values for SettlCurrFxRateCalc
const (
	SettlCurrFxRateCalc_DIVIDE   = "D"
	SettlCurrFxRateCalc_MULTIPLY = "M"
)

//Enum values for SettlDeliveryType
const (
	SettlDeliveryType_VERSUS_PAYMENT_DELIVER = "0"
	SettlDeliveryType_FREE_DELIVER           = "1"
	SettlDeliveryType_TRI_PARTY              = "2"
	SettlDeliveryType_HOLD_IN_CUSTODY        = "3"
)

//Enum values for SettlInstMode
const (
	SettlInstMode_DEFAULT                                = "0"
	SettlInstMode_STANDING_INSTRUCTIONS_PROVIDED         = "1"
	SettlInstMode_SPECIFIC_ALLOCATION_ACCOUNT_OVERRIDING = "2"
	SettlInstMode_SPECIFIC_ALLOCATION_ACCOUNT_STANDING   = "3"
	SettlInstMode_SPECIFIC_ORDER_FOR_A_SINGLE_ACCOUNT    = "4"
	SettlInstMode_REQUEST_REJECT                         = "5"
)

//Enum values for SettlInstReqRejCode
const (
	SettlInstReqRejCode_UNABLE_TO_PROCESS_REQUEST                 = "0"
	SettlInstReqRejCode_UNKNOWN_ACCOUNT                           = "1"
	SettlInstReqRejCode_NO_MATCHING_SETTLEMENT_INSTRUCTIONS_FOUND = "2"
	SettlInstReqRejCode_OTHER                                     = "99"
)

//Enum values for SettlInstSource
const (
	SettlInstSource_BROKERS_INSTRUCTIONS      = "1"
	SettlInstSource_INSTITUTIONS_INSTRUCTIONS = "2"
	SettlInstSource_INVESTOR                  = "3"
)

//Enum values for SettlInstTransType
const (
	SettlInstTransType_CANCEL  = "C"
	SettlInstTransType_NEW     = "N"
	SettlInstTransType_REPLACE = "R"
	SettlInstTransType_RESTATE = "T"
)

//Enum values for SettlLocation
const (
	SettlLocation_CEDEL                        = "CED"
	SettlLocation_DEPOSITORY_TRUST_COMPANY     = "DTC"
	SettlLocation_EURO_CLEAR                   = "EUR"
	SettlLocation_FEDERAL_BOOK_ENTRY           = "FED"
	SettlLocation_LOCAL_MARKET_SETTLE_LOCATION = "ISO_Country_Code"
	SettlLocation_PHYSICAL                     = "PNY"
	SettlLocation_PARTICIPANT_TRUST_COMPANY    = "PTC"
)

//Enum values for SettlMethod
const (
	SettlMethod_CASH_SETTLEMENT_REQUIRED     = "C"
	SettlMethod_PHYSICAL_SETTLEMENT_REQUIRED = "P"
)

//Enum values for SettlObligMode
const (
	SettlObligMode_PRELIMINARY = "1"
	SettlObligMode_FINAL       = "2"
)

//Enum values for SettlObligSource
const (
	SettlObligSource_INSTRUCTIONS_OF_BROKER       = "1"
	SettlObligSource_INSTRUCTIONS_FOR_INSTITUTION = "2"
	SettlObligSource_INVESTOR                     = "3"
)

//Enum values for SettlObligTransType
const (
	SettlObligTransType_CANCEL  = "C"
	SettlObligTransType_NEW     = "N"
	SettlObligTransType_REPLACE = "R"
	SettlObligTransType_RESTATE = "T"
)

//Enum values for SettlPriceType
const (
	SettlPriceType_FINAL       = "1"
	SettlPriceType_THEORETICAL = "2"
)

//Enum values for SettlSessID
const (
	SettlSessID_END_OF_DAY               = "EOD"
	SettlSessID_ELECTRONIC_TRADING_HOURS = "ETH"
	SettlSessID_INTRADAY                 = "ITD"
	SettlSessID_REGULAR_TRADING_HOURS    = "RTH"
)

//Enum values for SettlType
const (
	SettlType_REGULAR                 = "0"
	SettlType_CASH                    = "1"
	SettlType_NEXT_DAY                = "2"
	SettlType_T_PLUS_2                = "3"
	SettlType_T_PLUS_3                = "4"
	SettlType_T_PLUS_4                = "5"
	SettlType_FUTURE                  = "6"
	SettlType_WHEN_AND_IF_ISSUED      = "7"
	SettlType_SELLERS_OPTION          = "8"
	SettlType_T_PLUS_5                = "9"
	SettlType_BROKEN_DATE             = "B"
	SettlType_FX_SPOT_NEXT_SETTLEMENT = "C"
)

//Enum values for SettlmntTyp
const (
	SettlmntTyp_REGULAR            = "0"
	SettlmntTyp_CASH               = "1"
	SettlmntTyp_NEXT_DAY           = "2"
	SettlmntTyp_T_PLUS_2           = "3"
	SettlmntTyp_T_PLUS_3           = "4"
	SettlmntTyp_T_PLUS_4           = "5"
	SettlmntTyp_FUTURE             = "6"
	SettlmntTyp_WHEN_AND_IF_ISSUED = "7"
	SettlmntTyp_SELLERS_OPTION     = "8"
	SettlmntTyp_T_PLUS_5           = "9"
	SettlmntTyp_T_PLUS_1           = "A"
)

//Enum values for ShortSaleReason
const (
	ShortSaleReason_DEALER_SOLD_SHORT                        = "0"
	ShortSaleReason_DEALER_SOLD_SHORT_EXEMPT                 = "1"
	ShortSaleReason_SELLING_CUSTOMER_SOLD_SHORT              = "2"
	ShortSaleReason_SELLING_CUSTOMER_SOLD_SHORT_EXEMPT       = "3"
	ShortSaleReason_QUALIFIED_SERVICE_REPRESENTATIVE         = "4"
	ShortSaleReason_QSR_OR_AGU_CONTRA_SIDE_SOLD_SHORT_EXEMPT = "5"
)

//Enum values for Side
const (
	Side_BUY                = "1"
	Side_SELL               = "2"
	Side_BUY_MINUS          = "3"
	Side_SELL_PLUS          = "4"
	Side_SELL_SHORT         = "5"
	Side_SELL_SHORT_EXEMPT  = "6"
	Side_UNDISCLOSED        = "7"
	Side_CROSS              = "8"
	Side_CROSS_SHORT        = "9"
	Side_CROSS_SHORT_EXEMPT = "A"
	Side_AS_DEFINED         = "B"
	Side_OPPOSITE           = "C"
	Side_SUBSCRIBE          = "D"
	Side_REDEEM             = "E"
	Side_LEND               = "F"
	Side_BORROW             = "G"
)

//Enum values for SideMultiLegReportingType
const (
	SideMultiLegReportingType_SINGLE_SECURITY                       = "1"
	SideMultiLegReportingType_INDIVIDUAL_LEG_OF_A_MULTILEG_SECURITY = "2"
	SideMultiLegReportingType_MULTILEG_SECURITY                     = "3"
)

//Enum values for SideTrdSubTyp
const (
	SideTrdSubTyp_CMTA                                            = "0"
	SideTrdSubTyp_INTERNAL_TRANSFER                               = "1"
	SideTrdSubTyp_TRANSACTION_FROM_ASSIGNMENT                     = "10"
	SideTrdSubTyp_EXTERNAL_TRANSFER                               = "2"
	SideTrdSubTyp_REJECT_FOR_SUBMITTING_TRADE                     = "3"
	SideTrdSubTyp_ADVISORY_FOR_CONTRA_SIDE                        = "4"
	SideTrdSubTyp_OFFSET_DUE_TO_AN_ALLOCATION                     = "5"
	SideTrdSubTyp_ONSET_DUE_TO_AN_ALLOCATION                      = "6"
	SideTrdSubTyp_DIFFERENTIAL_SPREAD                             = "7"
	SideTrdSubTyp_IMPLIED_SPREAD_LEG_EXECUTED_AGAINST_AN_OUTRIGHT = "8"
	SideTrdSubTyp_TRANSACTION_FROM_EXERCISE                       = "9"
)

//Enum values for SideValueInd
const (
	SideValueInd_SIDE_VALUE_1 = "1"
	SideValueInd_SIDE_VALUE_2 = "2"
)

//Enum values for SolicitedFlag
const (
	SolicitedFlag_NO  = "N"
	SolicitedFlag_YES = "Y"
)

//Enum values for StandInstDbType
const (
	StandInstDbType_OTHER              = "0"
	StandInstDbType_DTC_SID            = "1"
	StandInstDbType_THOMSON_ALERT      = "2"
	StandInstDbType_A_GLOBAL_CUSTODIAN = "3"
	StandInstDbType_ACCOUNTNET         = "4"
)

//Enum values for StatsType
const (
	StatsType_EXCHANGE_LAST = "1"
	StatsType_HIGH          = "2"
	StatsType_AVERAGE_PRICE = "3"
	StatsType_TURNOVER      = "4"
)

//Enum values for StatusValue
const (
	StatusValue_CONNECTED       = "1"
	StatusValue_NOT_CONNECTED_2 = "2"
	StatusValue_NOT_CONNECTED_3 = "3"
	StatusValue_IN_PROCESS      = "4"
)

//Enum values for StipulationType
const (
	StipulationType_ABSOLUTE_PREPAYMENT_SPEED                          = "ABS"
	StipulationType_ALTERNATIVE_MINIMUM_TAX                            = "AMT"
	StipulationType_AUTO_REINVESTMENT_AT_RATE_OR_BETTER                = "AUTOREINV"
	StipulationType_AVAILABLE_OFFER_QUANTITY_TO_BE_SHOWN_TO_THE_STREET = "AVAILQTY"
	StipulationType_AVERAGE_FICO_SCORE                                 = "AVFICO"
	StipulationType_AVERAGE_LOAN_SIZE                                  = "AVSIZE"
	StipulationType_BANK_QUALIFIED                                     = "BANKQUAL"
	StipulationType_BARGAIN_CONDITIONS                                 = "BGNCON"
	StipulationType_BROKERS_SALES_CREDIT                               = "BROKERCREDIT"
	StipulationType_COUPON_RANGE                                       = "COUPON"
	StipulationType_CONSTANT_PREPAYMENT_PENALTY                        = "CPP"
	StipulationType_CONSTANT_PREPAYMENT_RATE                           = "CPR"
	StipulationType_CONSTANT_PREPAYMENT_YIELD                          = "CPY"
	StipulationType_ISO_CURRENCY_CODE                                  = "CURRENCY"
	StipulationType_CUSTOM_START_END_DATE                              = "CUSTOMDATE"
	StipulationType_DISCOUNT_RATE                                      = "DISCOUNT"
	StipulationType_GEOGRAPHICS_AND_RANGE                              = "GEOG"
	StipulationType_VALUATION_DISCOUNT                                 = "HAIRCUT"
	StipulationType_FINAL_CPR_OF_HOME_EQUITY_PREPAYMENT_CURVE          = "HEP"
	StipulationType_INSURED                                            = "INSURED"
	StipulationType_OFFER_PRICE_TO_BE_SHOWN_TO_INTERNAL_BROKERS        = "INTERNALPX"
	StipulationType_OFFER_QUANTITY_TO_BE_SHOWN_TO_INTERNAL_BROKERS     = "INTERNALQTY"
	StipulationType_YEAR_OR_YEAR_MONTH_OF_ISSUE                        = "ISSUE"
	StipulationType_ISSUERS_TICKER                                     = "ISSUER"
	StipulationType_ISSUE_SIZE_RANGE                                   = "ISSUESIZE"
	StipulationType_THE_MINIMUM_RESIDUAL_OFFER_QUANTITY                = "LEAVEQTY"
	StipulationType_LOOKBACK_DAYS                                      = "LOOKBACK"
	StipulationType_EXPLICIT_LOT_IDENTIFIER                            = "LOT"
	StipulationType_LOT_VARIANCE                                       = "LOTVAR"
	StipulationType_MATURITY_YEAR_AND_MONTH                            = "MAT"
	StipulationType_MATURITY_RANGE                                     = "MATURITY"
	StipulationType_MAXIMUM_LOAN_BALANCE                               = "MAXBAL"
	StipulationType_MAXIMUMDENOMINATION                                = "MAXDNOM"
	StipulationType_MAXIMUM_ORDER_SIZE                                 = "MAXORDQTY"
	StipulationType_MAXIMUM_SUBSTITUTIONS                              = "MAXSUBS"
	StipulationType_PERCENT_OF_MANUFACTURED_HOUSING_PREPAYMENT_CURVE   = "MHP"
	StipulationType_MINIMUM_DENOMINATION                               = "MINDNOM"
	StipulationType_MINIMUM_INCREMENT                                  = "MININCR"
	StipulationType_MINIMUM_QUANTITY                                   = "MINQTY"
	StipulationType_MONTHLY_PREPAYMENT_RATE                            = "MPR"
	StipulationType_ORDER_QUANTITY_INCREMENT                           = "ORDRINCR"
	StipulationType_PAYMENT_FREQUENCY_CALENDAR                         = "PAYFREQ"
	StipulationType_NUMBER_OF_PIECES                                   = "PIECES"
	StipulationType_POOLS_MAXIMUM                                      = "PMAX"
	StipulationType_POOLSMINIMUM                                       = "PMIN"
	StipulationType_POOL_IDENTIFIER                                    = "POOL"
	StipulationType_PERCENT_OF_PROSPECTUS_PREPAYMENT_CURVE             = "PPC"
	StipulationType_POOLS_PER_LOT                                      = "PPL"
	StipulationType_POOLS_PER_MILLION                                  = "PPM"
	StipulationType_POOLS_PER_TRADE                                    = "PPT"
	StipulationType_PRICE_RANGE                                        = "PRICE"
	StipulationType_PRICING_FREQUENCY                                  = "PRICEFREQ"
	StipulationType_PRIMARY_OR_SECONDARY_MARKET_INDICATOR              = "PRIMARY"
	StipulationType_PRODUCTION_YEAR                                    = "PROD"
	StipulationType_CALL_PROTECTION                                    = "PROTECT"
	StipulationType_PERCENT_OF_BMA_PREPAYMENT_CURVE                    = "PSA"
	StipulationType_PURPOSE                                            = "PURPOSE"
	StipulationType_BENCHMARK_PRICE_SOURCE                             = "PXSOURCE"
	StipulationType_RATING_SOURCE_AND_RANGE                            = "RATING"
	StipulationType_TYPE_OF_REDEMPTION                                 = "REDEMPTION"
	StipulationType_INTEREST_OF_ROLLING_OR_CLOSING_TRADE               = "REFINT"
	StipulationType_PRINCIPAL_OF_ROLLING_OR_CLOSING_TRADE              = "REFPRIN"
	StipulationType_REFERENCE_TO_ROLLING_OR_CLOSING_TRADE              = "REFTRADE"
	StipulationType_RESTRICTED                                         = "RESTRICTED"
	StipulationType_TYPE_OF_ROLL_TRADE                                 = "ROLLTYPE"
	StipulationType_BROKER_SALES_CREDIT_OVERRIDE                       = "SALESCREDITOVR"
	StipulationType_MARKET_SECTOR                                      = "SECTOR"
	StipulationType_SECURITY_TYPE_INCLUDED_OR_EXCLUDED                 = "SECTYPE"
	StipulationType_SINGLE_MONTHLY_MORTALITY                           = "SMM"
	StipulationType_STRUCTURE                                          = "STRUCT"
	StipulationType_SUBSTITUTIONS_FREQUENCY                            = "SUBSFREQ"
	StipulationType_SUBSTITUTIONS_LEFT                                 = "SUBSLEFT"
	StipulationType_FREEFORM_TEXT                                      = "TEXT"
	StipulationType_TRADERS_CREDIT                                     = "TRADERCREDIT"
	StipulationType_TRADE_VARIANCE                                     = "TRDVAR"
	StipulationType_WEIGHTED_AVERAGE_COUPON                            = "WAC"
	StipulationType_WEIGHTED_AVERAGE_LIFE_COUPON                       = "WAL"
	StipulationType_WEIGHTED_AVERAGE_LOAN_AGE                          = "WALA"
	StipulationType_WEIGHTED_AVERAGE_MATURITY                          = "WAM"
	StipulationType_WHOLE_POOL                                         = "WHOLE"
	StipulationType_YIELD_RANGE                                        = "YIELD"
	StipulationType_YIELD_TO_MATURITY                                  = "YTM"
)

//Enum values for StrategyParameterType
const (
	StrategyParameterType_INT                 = "1"
	StrategyParameterType_AMT                 = "10"
	StrategyParameterType_PERCENTAGE          = "11"
	StrategyParameterType_CHAR                = "12"
	StrategyParameterType_BOOLEAN             = "13"
	StrategyParameterType_STRING              = "14"
	StrategyParameterType_MULTIPLECHARVALUE   = "15"
	StrategyParameterType_CURRENCY            = "16"
	StrategyParameterType_EXCHANGE            = "17"
	StrategyParameterType_MONTHYEAR           = "18"
	StrategyParameterType_UTCTIMESTAMP        = "19"
	StrategyParameterType_LENGTH              = "2"
	StrategyParameterType_UTCTIMEONLY         = "20"
	StrategyParameterType_LOCALMKTDATE        = "21"
	StrategyParameterType_UTCDATEONLY         = "22"
	StrategyParameterType_DATA                = "23"
	StrategyParameterType_MULTIPLESTRINGVALUE = "24"
	StrategyParameterType_COUNTRY             = "25"
	StrategyParameterType_LANGUAGE            = "26"
	StrategyParameterType_TZTIMEONLY          = "27"
	StrategyParameterType_TZTIMESTAMP         = "28"
	StrategyParameterType_TENOR               = "29"
	StrategyParameterType_NUMINGROUP          = "3"
	StrategyParameterType_SEQNUM              = "4"
	StrategyParameterType_TAGNUM              = "5"
	StrategyParameterType_FLOAT               = "6"
	StrategyParameterType_QTY                 = "7"
	StrategyParameterType_PRICE               = "8"
	StrategyParameterType_PRICEOFFSET         = "9"
)

//Enum values for StreamAsgnAckType
const (
	StreamAsgnAckType_ASSIGNMENT_ACCEPTED = "0"
	StreamAsgnAckType_ASSIGNMENT_REJECTED = "1"
)

//Enum values for StreamAsgnRejReason
const (
	StreamAsgnRejReason_UNKNOWN_CLIENT                   = "0"
	StreamAsgnRejReason_EXCEEDS_MAXIMUM_SIZE             = "1"
	StreamAsgnRejReason_UNKNOWN_OR_INVALID_CURRENCY_PAIR = "2"
	StreamAsgnRejReason_NO_AVAILABLE_STREAM              = "3"
	StreamAsgnRejReason_OTHER                            = "99"
)

//Enum values for StreamAsgnReqType
const (
	StreamAsgnReqType_STREAM_ASSIGNMENT_FOR_NEW_CUSTOMER      = "1"
	StreamAsgnReqType_STREAM_ASSIGNMENT_FOR_EXISTING_CUSTOMER = "2"
)

//Enum values for StreamAsgnType
const (
	StreamAsgnType_ASSIGNMENT         = "1"
	StreamAsgnType_REJECTED           = "2"
	StreamAsgnType_TERMINATE_UNASSIGN = "3"
)

//Enum values for StrikePriceBoundaryMethod
const (
	StrikePriceBoundaryMethod_LESS_THAN_UNDERLYING_PRICE_IS_IN_THE_MONEY                 = "1"
	StrikePriceBoundaryMethod_LESS_THAN_OR_EQUAL_TO_THE_UNDERLYING_PRICE_IS_IN_THE_MONEY = "2"
	StrikePriceBoundaryMethod_EQUAL_TO_THE_UNDERLYING_PRICE_IS_IN_THE_MONEY              = "3"
	StrikePriceBoundaryMethod_GREATER_THAN_OR_EQUAL_TO_UNDERLYING_PRICE_IS_IN_THE_MONEY  = "4"
	StrikePriceBoundaryMethod_GREATER_THAN_UNDERLYING_IS_IN_THE_MONEY                    = "5"
)

//Enum values for StrikePriceDeterminationMethod
const (
	StrikePriceDeterminationMethod_FIXED_STRIKE                                                                       = "1"
	StrikePriceDeterminationMethod_STRIKE_SET_AT_EXPIRATION_TO_UNDERLYING_OR_OTHER_VALUE                              = "2"
	StrikePriceDeterminationMethod_STRIKE_SET_TO_AVERAGE_OF_UNDERLYING_SETTLEMENT_PRICE_ACROSS_THE_LIFE_OF_THE_OPTION = "3"
	StrikePriceDeterminationMethod_STRIKE_SET_TO_OPTIMAL_VALUE                                                        = "4"
)

//Enum values for SubscriptionRequestType
const (
	SubscriptionRequestType_SNAPSHOT                                      = "0"
	SubscriptionRequestType_SNAPSHOT_PLUS_UPDATES                         = "1"
	SubscriptionRequestType_DISABLE_PREVIOUS_SNAPSHOT_PLUS_UPDATE_REQUEST = "2"
)

//Enum values for SymbolSfx
const (
	SymbolSfx_EUCP_WITH_LUMP_SUM_INTEREST_RATHER_THAN_DISCOUNT_PRICE               = "CD"
	SymbolSfx_WHEN_ISSUED_FOR_A_SECURITY_TO_BE_REISSUED_UNDER_AN_OLD_CUSIP_OR_ISIN = "WI"
)

//Enum values for TargetStrategy
const (
	TargetStrategy_VWAP                                                          = "1"
	TargetStrategy_RESERVEDANDAVAILABLEFORBILATERALLYAGREEDUPONUSERDEFINEDVALUES = "1000"
	TargetStrategy_PARTICIPATE                                                   = "2"
	TargetStrategy_MININIZE_MARKET_IMPACT                                        = "3"
)

//Enum values for TaxAdvantageType
const (
	TaxAdvantageType_NONE_NOT_APPLICABLE              = "0"
	TaxAdvantageType_MAXI_ISA                         = "1"
	TaxAdvantageType_EMPLOYEE_10                      = "10"
	TaxAdvantageType_EMPLOYER_11                      = "11"
	TaxAdvantageType_EMPLOYER_12                      = "12"
	TaxAdvantageType_NON_FUND_PROTOTYPE_IRA           = "13"
	TaxAdvantageType_NON_FUND_QUALIFIED_PLAN          = "14"
	TaxAdvantageType_DEFINED_CONTRIBUTION_PLAN        = "15"
	TaxAdvantageType_INDIVIDUAL_RETIREMENT_ACCOUNT_16 = "16"
	TaxAdvantageType_INDIVIDUAL_RETIREMENT_ACCOUNT_17 = "17"
	TaxAdvantageType_KEOGH                            = "18"
	TaxAdvantageType_PROFIT_SHARING_PLAN              = "19"
	TaxAdvantageType_TESSA                            = "2"
	TaxAdvantageType_401                              = "20"
	TaxAdvantageType_SELF_DIRECTED_IRA                = "21"
	TaxAdvantageType_403                              = "22"
	TaxAdvantageType_457                              = "23"
	TaxAdvantageType_ROTH_IRA_24                      = "24"
	TaxAdvantageType_ROTH_IRA_25                      = "25"
	TaxAdvantageType_ROTH_CONVERSION_IRA_26           = "26"
	TaxAdvantageType_ROTH_CONVERSION_IRA_27           = "27"
	TaxAdvantageType_EDUCATION_IRA_28                 = "28"
	TaxAdvantageType_EDUCATION_IRA_29                 = "29"
	TaxAdvantageType_MINI_CASH_ISA                    = "3"
	TaxAdvantageType_MINI_STOCKS_AND_SHARES_ISA       = "4"
	TaxAdvantageType_MINI_INSURANCE_ISA               = "5"
	TaxAdvantageType_CURRENT_YEAR_PAYMENT             = "6"
	TaxAdvantageType_PRIOR_YEAR_PAYMENT               = "7"
	TaxAdvantageType_ASSET_TRANSFER                   = "8"
	TaxAdvantageType_EMPLOYEE_9                       = "9"
	TaxAdvantageType_OTHER                            = "999"
)

//Enum values for TerminationType
const (
	TerminationType_OVERNIGHT = "1"
	TerminationType_TERM      = "2"
	TerminationType_FLEXIBLE  = "3"
	TerminationType_OPEN      = "4"
)

//Enum values for TestMessageIndicator
const (
	TestMessageIndicator_NO  = "N"
	TestMessageIndicator_YES = "Y"
)

//Enum values for TickDirection
const (
	TickDirection_PLUS_TICK       = "0"
	TickDirection_ZERO_PLUS_TICK  = "1"
	TickDirection_MINUS_TICK      = "2"
	TickDirection_ZERO_MINUS_TICK = "3"
)

//Enum values for TickRuleType
const (
	TickRuleType_REGULAR                 = "0"
	TickRuleType_VARIABLE                = "1"
	TickRuleType_FIXED                   = "2"
	TickRuleType_TRADED_AS_A_SPREAD_LEG  = "3"
	TickRuleType_SETTLED_AS_A_SPREAD_LEG = "4"
)

//Enum values for TimeInForce
const (
	TimeInForce_DAY                   = "0"
	TimeInForce_GOOD_TILL_CANCEL      = "1"
	TimeInForce_AT_THE_OPENING        = "2"
	TimeInForce_IMMEDIATE_OR_CANCEL   = "3"
	TimeInForce_FILL_OR_KILL          = "4"
	TimeInForce_GOOD_TILL_CROSSING    = "5"
	TimeInForce_GOOD_TILL_DATE        = "6"
	TimeInForce_AT_THE_CLOSE          = "7"
	TimeInForce_GOOD_THROUGH_CROSSING = "8"
	TimeInForce_AT_CROSSING           = "9"
)

//Enum values for TimeUnit
const (
	TimeUnit_DAY    = "D"
	TimeUnit_HOUR   = "H"
	TimeUnit_MINUTE = "Min"
	TimeUnit_MONTH  = "Mo"
	TimeUnit_SECOND = "S"
	TimeUnit_WEEK   = "Wk"
	TimeUnit_YEAR   = "Yr"
)

//Enum values for TradSesEvent
const (
	TradSesEvent_TRADING_RESUMES              = "0"
	TradSesEvent_CHANGE_OF_TRADING_SESSION    = "1"
	TradSesEvent_CHANGE_OF_TRADING_SUBSESSION = "2"
	TradSesEvent_CHANGE_OF_TRADING_STATUS     = "3"
)

//Enum values for TradSesMethod
const (
	TradSesMethod_ELECTRONIC  = "1"
	TradSesMethod_OPEN_OUTCRY = "2"
	TradSesMethod_TWO_PARTY   = "3"
)

//Enum values for TradSesMode
const (
	TradSesMode_TESTING    = "1"
	TradSesMode_SIMULATED  = "2"
	TradSesMode_PRODUCTION = "3"
)

//Enum values for TradSesStatus
const (
	TradSesStatus_UNKNOWN          = "0"
	TradSesStatus_HALTED           = "1"
	TradSesStatus_OPEN             = "2"
	TradSesStatus_CLOSED           = "3"
	TradSesStatus_PRE_OPEN         = "4"
	TradSesStatus_PRE_CLOSE        = "5"
	TradSesStatus_REQUEST_REJECTED = "6"
)

//Enum values for TradSesStatusRejReason
const (
	TradSesStatusRejReason_UNKNOWN_OR_INVALID_TRADINGSESSIONID = "1"
	TradSesStatusRejReason_OTHER                               = "99"
)

//Enum values for TradeAllocIndicator
const (
	TradeAllocIndicator_ALLOCATION_NOT_REQUIRED                = "0"
	TradeAllocIndicator_ALLOCATION_REQUIRED                    = "1"
	TradeAllocIndicator_USE_ALLOCATION_PROVIDED_WITH_THE_TRADE = "2"
	TradeAllocIndicator_ALLOCATION_GIVE_UP_EXECUTOR            = "3"
	TradeAllocIndicator_ALLOCATION_FROM_EXECUTOR               = "4"
	TradeAllocIndicator_ALLOCATION_TO_CLAIM_ACCOUNT            = "5"
)

//Enum values for TradeCondition
const (
	TradeCondition_CANCEL                                  = "0"
	TradeCondition_IMPLIED_TRADE                           = "1"
	TradeCondition_MARKETPLACE_ENTERED_TRADE               = "2"
	TradeCondition_MULT_ASSET_CLASS_MULTILEG_TRADE         = "3"
	TradeCondition_MULTILEG_TO_MULTILEG_TRADE              = "4"
	TradeCondition_CASH                                    = "A"
	TradeCondition_SPREAD                                  = "AA"
	TradeCondition_SPREAD_ETH                              = "AB"
	TradeCondition_STRADDLE                                = "AC"
	TradeCondition_STRADDLE_ETH                            = "AD"
	TradeCondition_STOPPED                                 = "AE"
	TradeCondition_STOPPED_ETH                             = "AF"
	TradeCondition_REGULAR_ETH                             = "AG"
	TradeCondition_COMBO                                   = "AH"
	TradeCondition_COMBO_ETH                               = "AI"
	TradeCondition_OFFICIAL_CLOSING_PRICE                  = "AJ"
	TradeCondition_PRIOR_REFERENCE_PRICE                   = "AK"
	TradeCondition_STOPPED_SOLD_LAST                       = "AL"
	TradeCondition_STOPPED_OUT_OF_SEQUENCE                 = "AM"
	TradeCondition_OFFICAL_CLOSING_PRICE                   = "AN"
	TradeCondition_CROSSED_AO                              = "AO"
	TradeCondition_FAST_MARKET                             = "AP"
	TradeCondition_AUTOMATIC_EXECUTION                     = "AQ"
	TradeCondition_FORM_T                                  = "AR"
	TradeCondition_BASKET_INDEX                            = "AS"
	TradeCondition_BURST_BASKET                            = "AT"
	TradeCondition_OUTSIDE_SPREAD                          = "AV"
	TradeCondition_AVERAGE_PRICE_TRADE                     = "B"
	TradeCondition_CASH_TRADE                              = "C"
	TradeCondition_NEXT_DAY                                = "D"
	TradeCondition_OPENING_REOPENING_TRADE_DETAIL          = "E"
	TradeCondition_INTRADAY_TRADE_DETAIL                   = "F"
	TradeCondition_RULE_127_TRADE                          = "G"
	TradeCondition_RULE_155_TRADE                          = "H"
	TradeCondition_SOLD_LAST                               = "I"
	TradeCondition_NEXT_DAY_TRADE                          = "J"
	TradeCondition_OPENED                                  = "K"
	TradeCondition_SELLER                                  = "L"
	TradeCondition_SOLD                                    = "M"
	TradeCondition_STOPPED_STOCK                           = "N"
	TradeCondition_IMBALANCE_MORE_BUYERS                   = "P"
	TradeCondition_IMBALANCE_MORE_SELLERS                  = "Q"
	TradeCondition_OPENING_PRICE                           = "R"
	TradeCondition_BARGAIN_CONDITION                       = "S"
	TradeCondition_CONVERTED_PRICE_INDICATOR               = "T"
	TradeCondition_EXCHANGE_LAST                           = "U"
	TradeCondition_FINAL_PRICE_OF_SESSION                  = "V"
	TradeCondition_EX_PIT                                  = "W"
	TradeCondition_CROSSED_X                               = "X"
	TradeCondition_TRADES_RESULTING_FROM_MANUAL_SLOW_QUOTE = "Y"
	TradeCondition_TRADES_RESULTING_FROM_INTERMARKET_SWEEP = "Z"
	TradeCondition_VOLUME_ONLY                             = "a"
	TradeCondition_DIRECT_PLUS                             = "b"
	TradeCondition_ACQUISITION                             = "c"
	TradeCondition_BUNCHED                                 = "d"
	TradeCondition_DISTRIBUTION                            = "e"
	TradeCondition_BUNCHED_SALE                            = "f"
	TradeCondition_SPLIT_TRADE                             = "g"
	TradeCondition_CANCEL_STOPPED                          = "h"
	TradeCondition_CANCEL_ETH                              = "i"
	TradeCondition_CANCEL_STOPPED_ETH                      = "j"
	TradeCondition_OUT_OF_SEQUENCE_ETH                     = "k"
	TradeCondition_CANCEL_LAST_ETH                         = "l"
	TradeCondition_SOLD_LAST_SALE_ETH                      = "m"
	TradeCondition_CANCEL_LAST                             = "n"
	TradeCondition_SOLD_LAST_SALE                          = "o"
	TradeCondition_CANCEL_OPEN                             = "p"
	TradeCondition_CANCEL_OPEN_ETH                         = "q"
	TradeCondition_OPENED_SALE_ETH                         = "r"
	TradeCondition_CANCEL_ONLY                             = "s"
	TradeCondition_CANCEL_ONLY_ETH                         = "t"
	TradeCondition_LATE_OPEN_ETH                           = "u"
	TradeCondition_AUTO_EXECUTION_ETH                      = "v"
	TradeCondition_REOPEN                                  = "w"
	TradeCondition_REOPEN_ETH                              = "x"
	TradeCondition_ADJUSTED                                = "y"
	TradeCondition_ADJUSTED_ETH                            = "z"
)

//Enum values for TradeHandlingInstr
const (
	TradeHandlingInstr_TRADE_CONFIRMATION                = "0"
	TradeHandlingInstr_TWO_PARTY_REPORT                  = "1"
	TradeHandlingInstr_ONE_PARTY_REPORT_FOR_MATCHING     = "2"
	TradeHandlingInstr_ONE_PARTY_REPORT_FOR_PASS_THROUGH = "3"
	TradeHandlingInstr_AUTOMATED_FLOOR_ORDER_ROUTING     = "4"
	TradeHandlingInstr_TWO_PARTY_REPORT_FOR_CLAIM        = "5"
)

//Enum values for TradePublishIndicator
const (
	TradePublishIndicator_DO_NOT_PUBLISH_TRADE = "0"
	TradePublishIndicator_PUBLISH_TRADE        = "1"
	TradePublishIndicator_DEFERRED_PUBLICATION = "2"
)

//Enum values for TradeReportRejectReason
const (
	TradeReportRejectReason_SUCCESSFUL                    = "0"
	TradeReportRejectReason_INVALID_PARTY_ONFORMATION     = "1"
	TradeReportRejectReason_UNKNOWN_INSTRUMENT            = "2"
	TradeReportRejectReason_UNAUTHORIZED_TO_REPORT_TRADES = "3"
	TradeReportRejectReason_INVALID_TRADE_TYPE            = "4"
	TradeReportRejectReason_OTHER                         = "99"
)

//Enum values for TradeReportTransType
const (
	TradeReportTransType_NEW                             = "0"
	TradeReportTransType_CANCEL                          = "1"
	TradeReportTransType_REPLACE                         = "2"
	TradeReportTransType_RELEASE                         = "3"
	TradeReportTransType_REVERSE                         = "4"
	TradeReportTransType_CANCEL_DUE_TO_BACK_OUT_OF_TRADE = "5"
)

//Enum values for TradeReportType
const (
	TradeReportType_SUBMIT                      = "0"
	TradeReportType_ALLEGED_1                   = "1"
	TradeReportType_PENDED                      = "10"
	TradeReportType_ALLEGED_NEW                 = "11"
	TradeReportType_ALLEGED_ADDENDUM            = "12"
	TradeReportType_ALLEGED_NO_WAS              = "13"
	TradeReportType_ALLEGED_TRADE_REPORT_CANCEL = "14"
	TradeReportType_ALLEGED_15                  = "15"
	TradeReportType_ACCEPT                      = "2"
	TradeReportType_DECLINE                     = "3"
	TradeReportType_ADDENDUM                    = "4"
	TradeReportType_NO_WAS                      = "5"
	TradeReportType_TRADE_REPORT_CANCEL         = "6"
	TradeReportType_7                           = "7"
	TradeReportType_DEFAULTED                   = "8"
	TradeReportType_INVALID_CMTA                = "9"
)

//Enum values for TradeRequestResult
const (
	TradeRequestResult_SUCCESSFUL                       = "0"
	TradeRequestResult_INVALID_OR_UNKNOWN_INSTRUMENT    = "1"
	TradeRequestResult_INVALID_TYPE_OF_TRADE_REQUESTED  = "2"
	TradeRequestResult_INVALID_PARTIES                  = "3"
	TradeRequestResult_INVALID_TRANSPORT_TYPE_REQUESTED = "4"
	TradeRequestResult_INVALID_DESTINATION_REQUESTED    = "5"
	TradeRequestResult_TRADEREQUESTTYPE_NOT_SUPPORTED   = "8"
	TradeRequestResult_NOT_AUTHORIZED                   = "9"
	TradeRequestResult_OTHER                            = "99"
)

//Enum values for TradeRequestStatus
const (
	TradeRequestStatus_ACCEPTED  = "0"
	TradeRequestStatus_COMPLETED = "1"
	TradeRequestStatus_REJECTED  = "2"
)

//Enum values for TradeRequestType
const (
	TradeRequestType_ALL_TRADES                                           = "0"
	TradeRequestType_MATCHED_TRADES_MATCHING_CRITERIA_PROVIDED_ON_REQUEST = "1"
	TradeRequestType_UNMATCHED_TRADES_THAT_MATCH_CRITERIA                 = "2"
	TradeRequestType_UNREPORTED_TRADES_THAT_MATCH_CRITERIA                = "3"
	TradeRequestType_ADVISORIES_THAT_MATCH_CRITERIA                       = "4"
)

//Enum values for TradeType
const (
	TradeType_AGENCY           = "A"
	TradeType_VWAP_GUARANTEE   = "G"
	TradeType_GUARANTEED_CLOSE = "J"
	TradeType_RISK_TRADE       = "R"
)

//Enum values for TradedFlatSwitch
const (
	TradedFlatSwitch_NO  = "N"
	TradedFlatSwitch_YES = "Y"
)

//Enum values for TradingSessionID
const (
	TradingSessionID_DAY         = "1"
	TradingSessionID_HALFDAY     = "2"
	TradingSessionID_MORNING     = "3"
	TradingSessionID_AFTERNOON   = "4"
	TradingSessionID_EVENING     = "5"
	TradingSessionID_AFTER_HOURS = "6"
)

//Enum values for TradingSessionSubID
const (
	TradingSessionSubID_PRE_TRADING                = "1"
	TradingSessionSubID_OPENING_OR_OPENING_AUCTION = "2"
	TradingSessionSubID_3                          = "3"
	TradingSessionSubID_CLOSING_OR_CLOSING_AUCTION = "4"
	TradingSessionSubID_POST_TRADING               = "5"
	TradingSessionSubID_INTRADAY_AUCTION           = "6"
	TradingSessionSubID_QUIESCENT                  = "7"
)

//Enum values for TrdRegTimestampType
const (
	TrdRegTimestampType_EXECUTION_TIME         = "1"
	TrdRegTimestampType_TIME_IN                = "2"
	TrdRegTimestampType_TIME_OUT               = "3"
	TrdRegTimestampType_BROKER_RECEIPT         = "4"
	TrdRegTimestampType_BROKER_EXECUTION       = "5"
	TrdRegTimestampType_DESK_RECEIPT           = "6"
	TrdRegTimestampType_SUBMISSION_TO_CLEARING = "7"
)

//Enum values for TrdRptStatus
const (
	TrdRptStatus_ACCEPTED             = "0"
	TrdRptStatus_REJECTED             = "1"
	TrdRptStatus_ACCEPTED_WITH_ERRORS = "3"
)

//Enum values for TrdSubType
const (
	TrdSubType_CMTA                                            = "0"
	TrdSubType_INTERNAL_TRANSFER_OR_ADJUSTMENT                 = "1"
	TrdSubType_TRANSACTION_FROM_ASSIGNMENT                     = "10"
	TrdSubType_ACATS                                           = "11"
	TrdSubType_AI                                              = "14"
	TrdSubType_B                                               = "15"
	TrdSubType_K                                               = "16"
	TrdSubType_LC                                              = "17"
	TrdSubType_M                                               = "18"
	TrdSubType_N                                               = "19"
	TrdSubType_EXTERNAL_TRANSFER_OR_TRANSFER_OF_ACCOUNT        = "2"
	TrdSubType_NM                                              = "20"
	TrdSubType_NR                                              = "21"
	TrdSubType_P                                               = "22"
	TrdSubType_PA                                              = "23"
	TrdSubType_PC                                              = "24"
	TrdSubType_PN                                              = "25"
	TrdSubType_R                                               = "26"
	TrdSubType_RO                                              = "27"
	TrdSubType_RT                                              = "28"
	TrdSubType_SW                                              = "29"
	TrdSubType_REJECT_FOR_SUBMITTING_SIDE                      = "3"
	TrdSubType_T                                               = "30"
	TrdSubType_WN                                              = "31"
	TrdSubType_WT                                              = "32"
	TrdSubType_OFF_HOURS_TRADE                                 = "33"
	TrdSubType_ON_HOURS_TRADE                                  = "34"
	TrdSubType_OTC_QUOTE                                       = "35"
	TrdSubType_CONVERTED_SWAP                                  = "36"
	TrdSubType_CROSSED_TRADE                                   = "37"
	TrdSubType_INTERIM_PROTECTED_TRADE                         = "38"
	TrdSubType_LARGE_IN_SCALE                                  = "39"
	TrdSubType_ADVISORY_FOR_CONTRA_SIDE                        = "4"
	TrdSubType_OFFSET_DUE_TO_AN_ALLOCATION                     = "5"
	TrdSubType_ONSET_DUE_TO_AN_ALLOCATION                      = "6"
	TrdSubType_DIFFERENTIAL_SPREAD                             = "7"
	TrdSubType_IMPLIED_SPREAD_LEG_EXECUTED_AGAINST_AN_OUTRIGHT = "8"
	TrdSubType_TRANSACTION_FROM_EXERCISE                       = "9"
)

//Enum values for TrdType
const (
	TrdType_REGULAR_TRADE                        = "0"
	TrdType_BLOCK_TRADE_1                        = "1"
	TrdType_AFTER_HOURS_TRADE                    = "10"
	TrdType_EXCHANGE_FOR_RISK                    = "11"
	TrdType_EXCHANGE_FOR_SWAP                    = "12"
	TrdType_EXCHANGE_OF_FUTURES_FOR              = "13"
	TrdType_EXCHANGE_OF_OPTIONS_FOR_OPTIONS      = "14"
	TrdType_TRADING_AT_SETTLEMENT                = "15"
	TrdType_ALL_OR_NONE                          = "16"
	TrdType_FUTURES_LARGE_ORDER_EXECUTION        = "17"
	TrdType_EXCHANGE_OF_FUTURES_FOR_FUTURES      = "18"
	TrdType_OPTION_INTERIM_TRADE                 = "19"
	TrdType_EFP                                  = "2"
	TrdType_OPTION_CABINET_TRADE                 = "20"
	TrdType_PRIVATELY_NEGOTIATED_TRADES          = "22"
	TrdType_SUBSTITUTION_OF_FUTURES_FOR_FORWARDS = "23"
	TrdType_ERROR_TRADE                          = "24"
	TrdType_SPECIAL_CUM_DIVIDEND                 = "25"
	TrdType_SPECIAL_EX_DIVIDEND                  = "26"
	TrdType_SPECIAL_CUM_COUPON                   = "27"
	TrdType_SPECIAL_EX_COUPON                    = "28"
	TrdType_CASH_SETTLEMENT                      = "29"
	TrdType_TRANSFER                             = "3"
	TrdType_SPECIAL_PRICE                        = "30"
	TrdType_GUARANTEED_DELIVERY                  = "31"
	TrdType_SPECIAL_CUM_RIGHTS                   = "32"
	TrdType_SPECIAL_EX_RIGHTS                    = "33"
	TrdType_SPECIAL_CUM_CAPITAL_REPAYMENTS       = "34"
	TrdType_SPECIAL_EX_CAPITAL_REPAYMENTS        = "35"
	TrdType_SPECIAL_CUM_BONUS                    = "36"
	TrdType_SPECIAL_EX_BONUS                     = "37"
	TrdType_BLOCK_TRADE_38                       = "38"
	TrdType_WORKED_PRINCIPAL_TRADE               = "39"
	TrdType_LATE_TRADE                           = "4"
	TrdType_BLOCK_TRADES                         = "40"
	TrdType_NAME_CHANGE                          = "41"
	TrdType_PORTFOLIO_TRANSFER                   = "42"
	TrdType_PROROGATION_BUY                      = "43"
	TrdType_PROROGATION_SELL                     = "44"
	TrdType_OPTION_EXERCISE                      = "45"
	TrdType_DELTA_NEUTRAL_TRANSACTION            = "46"
	TrdType_FINANCING_TRANSACTION                = "47"
	TrdType_NON_STANDARD_SETTLEMENT              = "48"
	TrdType_DERIVATIVE_RELATED_TRANSACTION       = "49"
	TrdType_T_TRADE                              = "5"
	TrdType_PORTFOLIO_TRADE                      = "50"
	TrdType_VOLUME_WEIGHTED_AVERAGE_TRADE        = "51"
	TrdType_EXCHANGE_GRANTED_TRADE               = "52"
	TrdType_REPURCHASE_AGREEMENT                 = "53"
	TrdType_OTC                                  = "54"
	TrdType_EXCHANGE_BASIS_FACILITY              = "55"
	TrdType_WEIGHTED_AVERAGE_PRICE_TRADE         = "6"
	TrdType_BUNCHED_TRADE                        = "7"
	TrdType_LATE_BUNCHED_TRADE                   = "8"
	TrdType_PRIOR_REFERENCE_PRICE_TRADE          = "9"
)

//Enum values for TriggerAction
const (
	TriggerAction_ACTIVATE = "1"
	TriggerAction_MODIFY   = "2"
	TriggerAction_CANCEL   = "3"
)

//Enum values for TriggerOrderType
const (
	TriggerOrderType_MARKET = "1"
	TriggerOrderType_LIMIT  = "2"
)

//Enum values for TriggerPriceDirection
const (
	TriggerPriceDirection_TRIGGER_IF_THE_PRICE_OF_THE_SPECIFIED_TYPE_GOES_DOWN_TO_OR_THROUGH_THE_SPECIFIED_TRIGGER_PRICE = "D"
	TriggerPriceDirection_TRIGGER_IF_THE_PRICE_OF_THE_SPECIFIED_TYPE_GOES_UP_TO_OR_THROUGH_THE_SPECIFIED_TRIGGER_PRICE   = "U"
)

//Enum values for TriggerPriceType
const (
	TriggerPriceType_BEST_OFFER               = "1"
	TriggerPriceType_LAST_TRADE               = "2"
	TriggerPriceType_BEST_BID                 = "3"
	TriggerPriceType_BEST_BID_OR_LAST_TRADE   = "4"
	TriggerPriceType_BEST_OFFER_OR_LAST_TRADE = "5"
	TriggerPriceType_BEST_MID                 = "6"
)

//Enum values for TriggerPriceTypeScope
const (
	TriggerPriceTypeScope_NONE     = "0"
	TriggerPriceTypeScope_LOCAL    = "1"
	TriggerPriceTypeScope_NATIONAL = "2"
	TriggerPriceTypeScope_GLOBAL   = "3"
)

//Enum values for TriggerType
const (
	TriggerType_PARTIAL_EXECUTION         = "1"
	TriggerType_SPECIFIED_TRADING_SESSION = "2"
	TriggerType_NEXT_AUCTION              = "3"
	TriggerType_PRICE_MOVEMENT            = "4"
)

//Enum values for UnderlyingCashType
const (
	UnderlyingCashType_DIFF  = "DIFF"
	UnderlyingCashType_FIXED = "FIXED"
)

//Enum values for UnderlyingFXRateCalc
const (
	UnderlyingFXRateCalc_DIVIDE   = "D"
	UnderlyingFXRateCalc_MULTIPLY = "M"
)

//Enum values for UnderlyingPriceDeterminationMethod
const (
	UnderlyingPriceDeterminationMethod_REGULAR           = "1"
	UnderlyingPriceDeterminationMethod_SPECIAL_REFERENCE = "2"
	UnderlyingPriceDeterminationMethod_OPTIMAL_VALUE     = "3"
	UnderlyingPriceDeterminationMethod_AVERAGE_VALUE     = "4"
)

//Enum values for UnderlyingSettlementType
const (
	UnderlyingSettlementType_T_PLUS_1 = "2"
	UnderlyingSettlementType_T_PLUS_3 = "4"
	UnderlyingSettlementType_T_PLUS_4 = "5"
)

//Enum values for UnitOfMeasure
const (
	UnitOfMeasure_ALLOWANCES         = "Alw"
	UnitOfMeasure_BARRELS            = "Bbl"
	UnitOfMeasure_BILLION_CUBIC_FEET = "Bcf"
	UnitOfMeasure_BUSHELS            = "Bu"
	UnitOfMeasure_GALLONS            = "Gal"
	UnitOfMeasure_ONE_MILLION_BTU    = "MMBtu"
	UnitOfMeasure_MILLION_BARRELS    = "MMbbl"
	UnitOfMeasure_MEGAWATT_HOURS     = "MWh"
	UnitOfMeasure_US_DOLLARS         = "USD"
	UnitOfMeasure_POUNDS             = "lbs"
	UnitOfMeasure_TROY_OUNCES        = "oz_tr"
	UnitOfMeasure_METRIC_TONS        = "t"
	UnitOfMeasure_TONS               = "tn"
)

//Enum values for UnsolicitedIndicator
const (
	UnsolicitedIndicator_NO  = "N"
	UnsolicitedIndicator_YES = "Y"
)

//Enum values for Urgency
const (
	Urgency_NORMAL     = "0"
	Urgency_FLASH      = "1"
	Urgency_BACKGROUND = "2"
)

//Enum values for UserRequestType
const (
	UserRequestType_LOG_ON_USER                    = "1"
	UserRequestType_LOG_OFF_USER                   = "2"
	UserRequestType_CHANGE_PASSWORD_FOR_USER       = "3"
	UserRequestType_REQUEST_INDIVIDUAL_USER_STATUS = "4"
)

//Enum values for UserStatus
const (
	UserStatus_LOGGED_IN                      = "1"
	UserStatus_NOT_LOGGED_IN                  = "2"
	UserStatus_USER_NOT_RECOGNISED            = "3"
	UserStatus_PASSWORD_INCORRECT             = "4"
	UserStatus_PASSWORD_CHANGED               = "5"
	UserStatus_OTHER                          = "6"
	UserStatus_FORCED_USER_LOGOUT_BY_EXCHANGE = "7"
	UserStatus_SESSION_SHUTDOWN_WARNING       = "8"
)

//Enum values for ValuationMethod
const (
	ValuationMethod_CDS_STYLE_COLLATERALIZATION_OF_MARKET_TO_MARKET_AND_COUPON = "CDS"
	ValuationMethod_CDS_IN_DELIVERY                                            = "CDSD"
	ValuationMethod_PREMIUM_STYLE                                              = "EQTY"
	ValuationMethod_FUTURES_STYLE_MARK_TO_MARKET                               = "FUT"
	ValuationMethod_FUTURES_STYLE_WITH_AN_ATTACHED_CASH_ADJUSTMENT             = "FUTDA"
)

//Enum values for VenueType
const (
	VenueType_ELECTRONIC = "E"
	VenueType_PIT        = "P"
	VenueType_EX_PIT     = "X"
)

//Enum values for WorkingIndicator
const (
	WorkingIndicator_NO  = "N"
	WorkingIndicator_YES = "Y"
)

//Enum values for YieldType
const (
	YieldType_AFTER_TAX_YIELD                                                                        = "AFTERTAX"
	YieldType_ANNUAL_YIELD                                                                           = "ANNUAL"
	YieldType_YIELD_AT_ISSUE                                                                         = "ATISSUE"
	YieldType_YIELD_TO_AVERAGE_LIFE_THE_YIELD_ASSUMING_THAT_ALL_SINKS                                = "AVGLIFE"
	YieldType_YIELD_TO_AVG_MATURITY                                                                  = "AVGMATURITY"
	YieldType_BOOK_YIELD                                                                             = "BOOK"
	YieldType_YIELD_TO_NEXT_CALL                                                                     = "CALL"
	YieldType_YIELD_CHANGE_SINCE_CLOSE                                                               = "CHANGE"
	YieldType_CLOSING_YIELD                                                                          = "CLOSE"
	YieldType_COMPOUND_YIELD                                                                         = "COMPOUND"
	YieldType_CURRENT_YIELD                                                                          = "CURRENT"
	YieldType_GVNT_EQUIVALENT_YIELD                                                                  = "GOVTEQUIV"
	YieldType_TRUE_GROSS_YIELD                                                                       = "GROSS"
	YieldType_YIELD_WITH_INFLATION_ASSUMPTION                                                        = "INFLATION"
	YieldType_INVERSE_FLOATER_BOND_YIELD                                                             = "INVERSEFLOATER"
	YieldType_MOST_RECENT_CLOSING_YIELD                                                              = "LASTCLOSE"
	YieldType_CLOSING_YIELD_MOST_RECENT_MONTH                                                        = "LASTMONTH"
	YieldType_CLOSING_YIELD_MOST_RECENT_QUARTER                                                      = "LASTQUARTER"
	YieldType_CLOSING_YIELD_MOST_RECENT_YEAR                                                         = "LASTYEAR"
	YieldType_YIELD_TO_LONGEST_AVERAGE_LIFE                                                          = "LONGAVGLIFE"
	YieldType_YIELD_TO_LONGEST_AVERAGE                                                               = "LONGEST"
	YieldType_MARK_TO_MARKET_YIELD                                                                   = "MARK"
	YieldType_YIELD_TO_MATURITY                                                                      = "MATURITY"
	YieldType_YIELD_TO_NEXT_REFUND                                                                   = "NEXTREFUND"
	YieldType_OPEN_AVERAGE_YIELD                                                                     = "OPENAVG"
	YieldType_PREVIOUS_CLOSE_YIELD                                                                   = "PREVCLOSE"
	YieldType_PROCEEDS_YIELD                                                                         = "PROCEEDS"
	YieldType_YIELD_TO_NEXT_PUT                                                                      = "PUT"
	YieldType_SEMI_ANNUAL_YIELD                                                                      = "SEMIANNUAL"
	YieldType_YIELD_TO_SHORTEST_AVERAGE_LIFE                                                         = "SHORTAVGLIFE"
	YieldType_YIELD_TO_SHORTEST_AVERAGE                                                              = "SHORTEST"
	YieldType_SIMPLE_YIELD                                                                           = "SIMPLE"
	YieldType_TAX_EQUIVALENT_YIELD                                                                   = "TAXEQUIV"
	YieldType_YIELD_TO_TENDER_DATE                                                                   = "TENDER"
	YieldType_TRUE_YIELD                                                                             = "TRUE"
	YieldType_YIELD_VALUE_OF_1_32_THE_AMOUNT_THAT_THE_YIELD_WILL_CHANGE_FOR_A_1_32ND_CHANGE_IN_PRICE = "VALUE1/32"
	YieldType_YIELD_VALUE_OF_1_32                                                                    = "VALUE1_32"
	YieldType_YIELD_TO_WORST                                                                         = "WORST"
)
