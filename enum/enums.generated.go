package enum

//Enum values for AccountType
type AccountType string

const (
	AccountType_ACCOUNT_IS_CARRIED_ON_CUSTOMER_SIDE_OF_THE_BOOKS                       AccountType = "1"
	AccountType_ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS                       AccountType = "2"
	AccountType_HOUSE_TRADER                                                           AccountType = "3"
	AccountType_FLOOR_TRADER                                                           AccountType = "4"
	AccountType_ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS_AND_IS_CROSS_MARGINED AccountType = "6"
	AccountType_ACCOUNT_IS_HOUSE_TRADER_AND_IS_CROSS_MARGINED                          AccountType = "7"
	AccountType_JOINT_BACK_OFFICE_ACCOUNT                                              AccountType = "8"
)

//Enum values for AcctIDSource
type AcctIDSource string

const (
	AcctIDSource_BIC       AcctIDSource = "1"
	AcctIDSource_SID_CODE  AcctIDSource = "2"
	AcctIDSource_TFM       AcctIDSource = "3"
	AcctIDSource_OMGEO     AcctIDSource = "4"
	AcctIDSource_DTCC_CODE AcctIDSource = "5"
	AcctIDSource_OTHER     AcctIDSource = "99"
)

//Enum values for Adjustment
type Adjustment string

const (
	Adjustment_CANCEL     Adjustment = "1"
	Adjustment_ERROR      Adjustment = "2"
	Adjustment_CORRECTION Adjustment = "3"
)

//Enum values for AdjustmentType
type AdjustmentType string

const (
	AdjustmentType_PROCESS_REQUEST_AS_MARGIN_DISPOSITION AdjustmentType = "0"
	AdjustmentType_DELTA_PLUS                            AdjustmentType = "1"
	AdjustmentType_DELTA_MINUS                           AdjustmentType = "2"
	AdjustmentType_FINAL                                 AdjustmentType = "3"
)

//Enum values for AdvSide
type AdvSide string

const (
	AdvSide_BUY   AdvSide = "B"
	AdvSide_SELL  AdvSide = "S"
	AdvSide_TRADE AdvSide = "T"
	AdvSide_CROSS AdvSide = "X"
)

//Enum values for AdvTransType
type AdvTransType string

const (
	AdvTransType_CANCEL  AdvTransType = "C"
	AdvTransType_NEW     AdvTransType = "N"
	AdvTransType_REPLACE AdvTransType = "R"
)

//Enum values for AffirmStatus
type AffirmStatus string

const (
	AffirmStatus_RECEIVED                         AffirmStatus = "1"
	AffirmStatus_CONFIRM_REJECTED_IE_NOT_AFFIRMED AffirmStatus = "2"
	AffirmStatus_AFFIRMED                         AffirmStatus = "3"
)

//Enum values for AggregatedBook
type AggregatedBook string

const (
	AggregatedBook_NO  AggregatedBook = "N"
	AggregatedBook_YES AggregatedBook = "Y"
)

//Enum values for AggressorIndicator
type AggressorIndicator string

const (
	AggressorIndicator_NO  AggressorIndicator = "N"
	AggressorIndicator_YES AggressorIndicator = "Y"
)

//Enum values for AllocAccountType
type AllocAccountType string

const (
	AllocAccountType_ACCOUNT_IS_CARRIED_PN_CUSTOMER_SIDE_OF_BOOKS                           AllocAccountType = "1"
	AllocAccountType_ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS                       AllocAccountType = "2"
	AllocAccountType_HOUSE_TRADER                                                           AllocAccountType = "3"
	AllocAccountType_FLOOR_TRADER                                                           AllocAccountType = "4"
	AllocAccountType_ACCOUNT_IS_CARRIED_ON_NON_CUSTOMER_SIDE_OF_BOOKS_AND_IS_CROSS_MARGINED AllocAccountType = "6"
	AllocAccountType_ACCOUNT_IS_HOUSE_TRADER_AND_IS_CROSS_MARGINED                          AllocAccountType = "7"
	AllocAccountType_JOINT_BACK_OFFICE_ACCOUNT                                              AllocAccountType = "8"
)

//Enum values for AllocCancReplaceReason
type AllocCancReplaceReason string

const (
	AllocCancReplaceReason_ORIGINAL_DETAILS_INCOMPLETE_INCORRECT AllocCancReplaceReason = "1"
	AllocCancReplaceReason_CHANGE_IN_UNDERLYING_ORDER_DETAILS    AllocCancReplaceReason = "2"
	AllocCancReplaceReason_OTHER                                 AllocCancReplaceReason = "99"
)

//Enum values for AllocHandlInst
type AllocHandlInst string

const (
	AllocHandlInst_MATCH             AllocHandlInst = "1"
	AllocHandlInst_FORWARD           AllocHandlInst = "2"
	AllocHandlInst_FORWARD_AND_MATCH AllocHandlInst = "3"
)

//Enum values for AllocIntermedReqType
type AllocIntermedReqType string

const (
	AllocIntermedReqType_PENDING_ACCEPT       AllocIntermedReqType = "1"
	AllocIntermedReqType_PENDING_RELEASE      AllocIntermedReqType = "2"
	AllocIntermedReqType_PENDING_REVERSAL     AllocIntermedReqType = "3"
	AllocIntermedReqType_ACCEPT               AllocIntermedReqType = "4"
	AllocIntermedReqType_BLOCK_LEVEL_REJECT   AllocIntermedReqType = "5"
	AllocIntermedReqType_ACCOUNT_LEVEL_REJECT AllocIntermedReqType = "6"
)

//Enum values for AllocLinkType
type AllocLinkType string

const (
	AllocLinkType_FX_NETTING AllocLinkType = "0"
	AllocLinkType_FX_SWAP    AllocLinkType = "1"
)

//Enum values for AllocMethod
type AllocMethod string

const (
	AllocMethod_AUTOMATIC AllocMethod = "1"
	AllocMethod_GUARANTOR AllocMethod = "2"
	AllocMethod_MANUAL    AllocMethod = "3"
)

//Enum values for AllocNoOrdersType
type AllocNoOrdersType string

const (
	AllocNoOrdersType_NOT_SPECIFIED          AllocNoOrdersType = "0"
	AllocNoOrdersType_EXPLICIT_LIST_PROVIDED AllocNoOrdersType = "1"
)

//Enum values for AllocPositionEffect
type AllocPositionEffect string

const (
	AllocPositionEffect_CLOSE  AllocPositionEffect = "C"
	AllocPositionEffect_FIFO   AllocPositionEffect = "F"
	AllocPositionEffect_OPEN   AllocPositionEffect = "O"
	AllocPositionEffect_ROLLED AllocPositionEffect = "R"
)

//Enum values for AllocRejCode
type AllocRejCode string

const (
	AllocRejCode_UNKNOWN_ACCOUNT                   AllocRejCode = "0"
	AllocRejCode_INCORRECT_QUANTITY                AllocRejCode = "1"
	AllocRejCode_UNKNOWN_OR_STALE_EXECID           AllocRejCode = "10"
	AllocRejCode_MISMATCHED_DATA                   AllocRejCode = "11"
	AllocRejCode_UNKNOWN_CLORDID                   AllocRejCode = "12"
	AllocRejCode_WAREHOUSE_REQUEST_REJECTED        AllocRejCode = "13"
	AllocRejCode_INCORRECT_AVERAGEG_PRICE          AllocRejCode = "2"
	AllocRejCode_UNKNOWN_EXECUTING_BROKER_MNEMONIC AllocRejCode = "3"
	AllocRejCode_COMMISSION_DIFFERENCE             AllocRejCode = "4"
	AllocRejCode_UNKNOWN_ORDERID                   AllocRejCode = "5"
	AllocRejCode_UNKNOWN_LISTID                    AllocRejCode = "6"
	AllocRejCode_OTHER_7                           AllocRejCode = "7"
	AllocRejCode_INCORRECT_ALLOCATED_QUANTITY      AllocRejCode = "8"
	AllocRejCode_CALCULATION_DIFFERENCE            AllocRejCode = "9"
	AllocRejCode_OTHER_99                          AllocRejCode = "99"
)

//Enum values for AllocReportType
type AllocReportType string

const (
	AllocReportType_REJECT                                  AllocReportType = "10"
	AllocReportType_ACCEPT_PENDING                          AllocReportType = "11"
	AllocReportType_COMPLETE                                AllocReportType = "12"
	AllocReportType_REVERSE_PENDING                         AllocReportType = "14"
	AllocReportType_PRELIMINARY_REQUEST_TO_INTERMEDIARY     AllocReportType = "2"
	AllocReportType_SELLSIDE_CALCULATED_USING_PRELIMINARY   AllocReportType = "3"
	AllocReportType_SELLSIDE_CALCULATED_WITHOUT_PRELIMINARY AllocReportType = "4"
	AllocReportType_WAREHOUSE_RECAP                         AllocReportType = "5"
	AllocReportType_REQUEST_TO_INTERMEDIARY                 AllocReportType = "8"
	AllocReportType_ACCEPT                                  AllocReportType = "9"
)

//Enum values for AllocSettlInstType
type AllocSettlInstType string

const (
	AllocSettlInstType_USE_DEFAULT_INSTRUCTIONS        AllocSettlInstType = "0"
	AllocSettlInstType_DERIVE_FROM_PARAMETERS_PROVIDED AllocSettlInstType = "1"
	AllocSettlInstType_FULL_DETAILS_PROVIDED           AllocSettlInstType = "2"
	AllocSettlInstType_SSI_DB_IDS_PROVIDED             AllocSettlInstType = "3"
	AllocSettlInstType_PHONE_FOR_INSTRUCTIONS          AllocSettlInstType = "4"
)

//Enum values for AllocStatus
type AllocStatus string

const (
	AllocStatus_ACCEPTED                 AllocStatus = "0"
	AllocStatus_BLOCK_LEVEL_REJECT       AllocStatus = "1"
	AllocStatus_ACCOUNT_LEVEL_REJECT     AllocStatus = "2"
	AllocStatus_RECEIVED                 AllocStatus = "3"
	AllocStatus_INCOMPLETE               AllocStatus = "4"
	AllocStatus_REJECTED_BY_INTERMEDIARY AllocStatus = "5"
	AllocStatus_ALLOCATION_PENDING       AllocStatus = "6"
	AllocStatus_REVERSED                 AllocStatus = "7"
)

//Enum values for AllocTransType
type AllocTransType string

const (
	AllocTransType_NEW                            AllocTransType = "0"
	AllocTransType_REPLACE                        AllocTransType = "1"
	AllocTransType_CANCEL                         AllocTransType = "2"
	AllocTransType_PRELIMINARY                    AllocTransType = "3"
	AllocTransType_CALCULATED                     AllocTransType = "4"
	AllocTransType_CALCULATED_WITHOUT_PRELIMINARY AllocTransType = "5"
	AllocTransType_REVERSAL                       AllocTransType = "6"
)

//Enum values for AllocType
type AllocType string

const (
	AllocType_CALCULATED                              AllocType = "1"
	AllocType_REJECT                                  AllocType = "10"
	AllocType_ACCEPT_PENDING                          AllocType = "11"
	AllocType_INCOMPLETE_GROUP                        AllocType = "12"
	AllocType_COMPLETE_GROUP                          AllocType = "13"
	AllocType_REVERSAL_PENDING                        AllocType = "14"
	AllocType_PRELIMINARY                             AllocType = "2"
	AllocType_SELLSIDE_CALCULATED_USING_PRELIMINARY   AllocType = "3"
	AllocType_SELLSIDE_CALCULATED_WITHOUT_PRELIMINARY AllocType = "4"
	AllocType_READY_TO_BOOK                           AllocType = "5"
	AllocType_BUYSIDE_READY_TO_BOOK                   AllocType = "6"
	AllocType_WAREHOUSE_INSTRUCTION                   AllocType = "7"
	AllocType_REQUEST_TO_INTERMEDIARY                 AllocType = "8"
	AllocType_ACCEPT                                  AllocType = "9"
)

//Enum values for ApplQueueAction
type ApplQueueAction string

const (
	ApplQueueAction_NO_ACTION_TAKEN ApplQueueAction = "0"
	ApplQueueAction_QUEUE_FLUSHED   ApplQueueAction = "1"
	ApplQueueAction_OVERLAY_LAST    ApplQueueAction = "2"
	ApplQueueAction_END_SESSION     ApplQueueAction = "3"
)

//Enum values for ApplQueueResolution
type ApplQueueResolution string

const (
	ApplQueueResolution_NO_ACTION_TAKEN ApplQueueResolution = "0"
	ApplQueueResolution_QUEUE_FLUSHED   ApplQueueResolution = "1"
	ApplQueueResolution_OVERLAY_LAST    ApplQueueResolution = "2"
	ApplQueueResolution_END_SESSION     ApplQueueResolution = "3"
)

//Enum values for ApplReportType
type ApplReportType string

const (
	ApplReportType_RESET_APPLSEQNUM_TO_NEW_VALUE_SPECIFIED_IN_APPLNEWSEQNUM                               ApplReportType = "0"
	ApplReportType_REPORTS_THAT_THE_LAST_MESSAGE_HAS_BEEN_SENT_FOR_THE_APPLIDS_REFER_TO_REFAPPLLASTSEQNUM ApplReportType = "1"
	ApplReportType_HEARTBEAT_MESSAGE_INDICATING_THAT_APPLICATION_IDENTIFIED_BY_REFAPPLID                  ApplReportType = "2"
	ApplReportType_APPLICATION_MESSAGE_RE_SEND_COMPLETED                                                  ApplReportType = "3"
)

//Enum values for ApplReqType
type ApplReqType string

const (
	ApplReqType_RETRANSMISSION_OF_APPLICATION_MESSAGES_FOR_THE_SPECIFIED_APPLICATIONS        ApplReqType = "0"
	ApplReqType_SUBSCRIPTION_TO_THE_SPECIFIED_APPLICATIONS                                   ApplReqType = "1"
	ApplReqType_REQUEST_FOR_THE_LAST_APPLLASTSEQNUM_PUBLISHED_FOR_THE_SPECIFIED_APPLICATIONS ApplReqType = "2"
	ApplReqType_REQUEST_VALID_SET_OF_APPLICATIONS                                            ApplReqType = "3"
	ApplReqType_UNSUBSCRIBE_TO_THE_SPECIFIED_APPLICATIONS                                    ApplReqType = "4"
	ApplReqType_CANCEL_RETRANSMISSION                                                        ApplReqType = "5"
	ApplReqType_CANCEL_RETRANSMISSION_AND_UNSUBSCRIBE_TO_THE_SPECIFIED_APPLICATIONS          ApplReqType = "6"
)

//Enum values for ApplResponseError
type ApplResponseError string

const (
	ApplResponseError_APPLICATION_DOES_NOT_EXIST           ApplResponseError = "0"
	ApplResponseError_MESSAGES_REQUESTED_ARE_NOT_AVAILABLE ApplResponseError = "1"
	ApplResponseError_USER_NOT_AUTHORIZED_FOR_APPLICATION  ApplResponseError = "2"
)

//Enum values for ApplResponseType
type ApplResponseType string

const (
	ApplResponseType_REQUEST_SUCCESSFULLY_PROCESSED ApplResponseType = "0"
	ApplResponseType_APPLICATION_DOES_NOT_EXIST     ApplResponseType = "1"
	ApplResponseType_MESSAGES_NOT_AVAILABLE         ApplResponseType = "2"
)

//Enum values for ApplVerID
type ApplVerID string

const (
	ApplVerID_FIX27    ApplVerID = "0"
	ApplVerID_FIX30    ApplVerID = "1"
	ApplVerID_FIX40    ApplVerID = "2"
	ApplVerID_FIX41    ApplVerID = "3"
	ApplVerID_FIX42    ApplVerID = "4"
	ApplVerID_FIX43    ApplVerID = "5"
	ApplVerID_FIX44    ApplVerID = "6"
	ApplVerID_FIX50    ApplVerID = "7"
	ApplVerID_FIX50SP1 ApplVerID = "8"
	ApplVerID_FIX50SP2 ApplVerID = "9"
)

//Enum values for AsOfIndicator
type AsOfIndicator string

const (
	AsOfIndicator_FALSE AsOfIndicator = "0"
	AsOfIndicator_TRUE  AsOfIndicator = "1"
)

//Enum values for AssignmentMethod
type AssignmentMethod string

const (
	AssignmentMethod_PRO_RATA AssignmentMethod = "P"
	AssignmentMethod_RANDOM   AssignmentMethod = "R"
)

//Enum values for AvgPxIndicator
type AvgPxIndicator string

const (
	AvgPxIndicator_NO_AVERAGE_PRICING                                                    AvgPxIndicator = "0"
	AvgPxIndicator_TRADE_IS_PART_OF_AN_AVERAGE_PRICE_GROUP_IDENTIFIED_BY_THE_TRADELINKID AvgPxIndicator = "1"
	AvgPxIndicator_LAST_TRADE_IS_THE_AVERAGE_PRICE_GROUP_IDENTIFIED_BY_THE_TRADELINKID   AvgPxIndicator = "2"
)

//Enum values for BasisPxType
type BasisPxType string

const (
	BasisPxType_CLOSING_PRICE_AT_MORNINGN_SESSION             BasisPxType = "2"
	BasisPxType_CLOSING_PRICE                                 BasisPxType = "3"
	BasisPxType_CURRENT_PRICE                                 BasisPxType = "4"
	BasisPxType_SQ                                            BasisPxType = "5"
	BasisPxType_VWAP_THROUGH_A_DAY                            BasisPxType = "6"
	BasisPxType_VWAP_THROUGH_A_MORNING_SESSION                BasisPxType = "7"
	BasisPxType_VWAP_THROUGH_AN_AFTERNOON_SESSION             BasisPxType = "8"
	BasisPxType_VWAP_THROUGH_A_DAY_EXCEPT_YORI                BasisPxType = "9"
	BasisPxType_VWAP_THROUGH_A_MORNING_SESSION_EXCEPT_YORI    BasisPxType = "A"
	BasisPxType_VWAP_THROUGH_AN_AFTERNOON_SESSION_EXCEPT_YORI BasisPxType = "B"
	BasisPxType_STRIKE                                        BasisPxType = "C"
	BasisPxType_OPEN                                          BasisPxType = "D"
	BasisPxType_OTHERS                                        BasisPxType = "Z"
)

//Enum values for Benchmark
type Benchmark string

const (
	Benchmark_CURVE    Benchmark = "1"
	Benchmark_5YR      Benchmark = "2"
	Benchmark_OLD5     Benchmark = "3"
	Benchmark_10YR     Benchmark = "4"
	Benchmark_OLD10    Benchmark = "5"
	Benchmark_30YR     Benchmark = "6"
	Benchmark_OLD30    Benchmark = "7"
	Benchmark_3MOLIBOR Benchmark = "8"
	Benchmark_6MOLIBOR Benchmark = "9"
)

//Enum values for BenchmarkCurveName
type BenchmarkCurveName string

const (
	BenchmarkCurveName_EONIA       BenchmarkCurveName = "EONIA"
	BenchmarkCurveName_EUREPO      BenchmarkCurveName = "EUREPO"
	BenchmarkCurveName_EURIBOR     BenchmarkCurveName = "Euribor"
	BenchmarkCurveName_FUTURESWAP  BenchmarkCurveName = "FutureSWAP"
	BenchmarkCurveName_LIBID       BenchmarkCurveName = "LIBID"
	BenchmarkCurveName_LIBOR       BenchmarkCurveName = "LIBOR"
	BenchmarkCurveName_MUNIAAA     BenchmarkCurveName = "MuniAAA"
	BenchmarkCurveName_OTHER       BenchmarkCurveName = "OTHER"
	BenchmarkCurveName_PFANDBRIEFE BenchmarkCurveName = "Pfandbriefe"
	BenchmarkCurveName_SONIA       BenchmarkCurveName = "SONIA"
	BenchmarkCurveName_SWAP        BenchmarkCurveName = "SWAP"
	BenchmarkCurveName_TREASURY    BenchmarkCurveName = "Treasury"
)

//Enum values for BidDescriptorType
type BidDescriptorType string

const (
	BidDescriptorType_SECTOR  BidDescriptorType = "1"
	BidDescriptorType_COUNTRY BidDescriptorType = "2"
	BidDescriptorType_INDEX   BidDescriptorType = "3"
)

//Enum values for BidRequestTransType
type BidRequestTransType string

const (
	BidRequestTransType_CANCEL BidRequestTransType = "C"
	BidRequestTransType_NO     BidRequestTransType = "N"
)

//Enum values for BidTradeType
type BidTradeType string

const (
	BidTradeType_AGENCY           BidTradeType = "A"
	BidTradeType_VWAP_GUARANTEE   BidTradeType = "G"
	BidTradeType_GUARANTEED_CLOSE BidTradeType = "J"
	BidTradeType_RISK_TRADE       BidTradeType = "R"
)

//Enum values for BidType
type BidType string

const (
	BidType_NON_DISCLOSED_STYLE BidType = "1"
	BidType_DISCLOSED_SYTLE     BidType = "2"
	BidType_NO_BIDDING_PROCESS  BidType = "3"
)

//Enum values for BookingType
type BookingType string

const (
	BookingType_REGULAR_BOOKING   BookingType = "0"
	BookingType_CFD               BookingType = "1"
	BookingType_TOTAL_RETURN_SWAP BookingType = "2"
)

//Enum values for BookingUnit
type BookingUnit string

const (
	BookingUnit_EACH_PARTIAL_EXECUTION_IS_A_BOOKABLE_UNIT                               BookingUnit = "0"
	BookingUnit_AGGREGATE_PARTIAL_EXECUTIONS_ON_THIS_ORDER_AND_BOOK_ONE_TRADE_PER_ORDER BookingUnit = "1"
	BookingUnit_AGGREGATE_EXECUTIONS_FOR_THIS_SYMBOL_SIDE_AND_SETTLEMENT_DATE           BookingUnit = "2"
)

//Enum values for BusinessRejectReason
type BusinessRejectReason string

const (
	BusinessRejectReason_OTHER                                     BusinessRejectReason = "0"
	BusinessRejectReason_UNKNOWN_ID                                BusinessRejectReason = "1"
	BusinessRejectReason_INVALID_PRICE_INCREMENT                   BusinessRejectReason = "18"
	BusinessRejectReason_UNKNOWN_SECURITY                          BusinessRejectReason = "2"
	BusinessRejectReason_UNSUPPORTED_MESSAGE_TYPE                  BusinessRejectReason = "3"
	BusinessRejectReason_APPLICATION_NOT_AVAILABLE                 BusinessRejectReason = "4"
	BusinessRejectReason_CONDITIONALLY_REQUIRED_FIELD_MISSING      BusinessRejectReason = "5"
	BusinessRejectReason_NOT_AUTHORIZED                            BusinessRejectReason = "6"
	BusinessRejectReason_DELIVERTO_FIRM_NOT_AVAILABLE_AT_THIS_TIME BusinessRejectReason = "7"
)

//Enum values for CPProgram
type CPProgram string

const (
	CPProgram_3     CPProgram = "1"
	CPProgram_4     CPProgram = "2"
	CPProgram_OTHER CPProgram = "99"
)

//Enum values for CancellationRights
type CancellationRights string

const (
	CancellationRights_NO_M CancellationRights = "M"
	CancellationRights_NO_N CancellationRights = "N"
	CancellationRights_NO_O CancellationRights = "O"
	CancellationRights_YES  CancellationRights = "Y"
)

//Enum values for CashMargin
type CashMargin string

const (
	CashMargin_CASH         CashMargin = "1"
	CashMargin_MARGIN_OPEN  CashMargin = "2"
	CashMargin_MARGIN_CLOSE CashMargin = "3"
)

//Enum values for ClearingFeeIndicator
type ClearingFeeIndicator string

const (
	ClearingFeeIndicator_1ST_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              ClearingFeeIndicator = "1"
	ClearingFeeIndicator_2ND_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              ClearingFeeIndicator = "2"
	ClearingFeeIndicator_3RD_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              ClearingFeeIndicator = "3"
	ClearingFeeIndicator_4TH_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              ClearingFeeIndicator = "4"
	ClearingFeeIndicator_5TH_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              ClearingFeeIndicator = "5"
	ClearingFeeIndicator_6TH_YEAR_DELEGATE_TRADING_FOR_OWN_ACCOUNT                              ClearingFeeIndicator = "9"
	ClearingFeeIndicator_CBOE_MEMBER                                                            ClearingFeeIndicator = "B"
	ClearingFeeIndicator_NON_MEMBER_AND_CUSTOMER                                                ClearingFeeIndicator = "C"
	ClearingFeeIndicator_EQUITY_MEMBER_AND_CLEARING_MEMBER                                      ClearingFeeIndicator = "E"
	ClearingFeeIndicator_FULL_AND_ASSOCIATE_MEMBER_TRADING_FOR_OWN_ACCOUNT_AND_AS_FLOOR_BROKERS ClearingFeeIndicator = "F"
	ClearingFeeIndicator_106H_AND_106J_FIRMS                                                    ClearingFeeIndicator = "H"
	ClearingFeeIndicator_GIM_IDEM_AND_COM_MEMBERSHIP_INTEREST_HOLDERS                           ClearingFeeIndicator = "I"
	ClearingFeeIndicator_LESSEE_106F_EMPLOYEES                                                  ClearingFeeIndicator = "L"
	ClearingFeeIndicator_ALL_OTHER_OWNERSHIP_TYPES                                              ClearingFeeIndicator = "M"
)

//Enum values for ClearingInstruction
type ClearingInstruction string

const (
	ClearingInstruction_PROCESS_NORMALLY                     ClearingInstruction = "0"
	ClearingInstruction_EXCLUDE_FROM_ALL_NETTING             ClearingInstruction = "1"
	ClearingInstruction_AUTOMATIC_GIVE_UP_MODE               ClearingInstruction = "10"
	ClearingInstruction_QUALIFIED_SERVICE_REPRESENTATIVE_QSR ClearingInstruction = "11"
	ClearingInstruction_CUSTOMER_TRADE                       ClearingInstruction = "12"
	ClearingInstruction_SELF_CLEARING                        ClearingInstruction = "13"
	ClearingInstruction_BILATERAL_NETTING_ONLY               ClearingInstruction = "2"
	ClearingInstruction_EX_CLEARING                          ClearingInstruction = "3"
	ClearingInstruction_SPECIAL_TRADE                        ClearingInstruction = "4"
	ClearingInstruction_MULTILATERAL_NETTING                 ClearingInstruction = "5"
	ClearingInstruction_CLEAR_AGAINST_CENTRAL_COUNTERPARTY   ClearingInstruction = "6"
	ClearingInstruction_EXCLUDE_FROM_CENTRAL_COUNTERPARTY    ClearingInstruction = "7"
	ClearingInstruction_MANUAL_MODE                          ClearingInstruction = "8"
	ClearingInstruction_AUTOMATIC_POSTING_MODE               ClearingInstruction = "9"
)

//Enum values for CollAction
type CollAction string

const (
	CollAction_RETAIN CollAction = "0"
	CollAction_ADD    CollAction = "1"
	CollAction_REMOVE CollAction = "2"
)

//Enum values for CollApplType
type CollApplType string

const (
	CollApplType_SPECIFIC_DEPOSIT CollApplType = "0"
	CollApplType_GENERAL          CollApplType = "1"
)

//Enum values for CollAsgnReason
type CollAsgnReason string

const (
	CollAsgnReason_INITIAL                   CollAsgnReason = "0"
	CollAsgnReason_SCHEDULED                 CollAsgnReason = "1"
	CollAsgnReason_TIME_WARNING              CollAsgnReason = "2"
	CollAsgnReason_MARGIN_DEFICIENCY         CollAsgnReason = "3"
	CollAsgnReason_MARGIN_EXCESS             CollAsgnReason = "4"
	CollAsgnReason_FORWARD_COLLATERAL_DEMAND CollAsgnReason = "5"
	CollAsgnReason_EVENT_OF_DEFAULT          CollAsgnReason = "6"
	CollAsgnReason_ADVERSE_TAX_EVENT         CollAsgnReason = "7"
)

//Enum values for CollAsgnRejectReason
type CollAsgnRejectReason string

const (
	CollAsgnRejectReason_UNKNOWN_DEAL                  CollAsgnRejectReason = "0"
	CollAsgnRejectReason_UNKNOWN_OR_INVALID_INSTRUMENT CollAsgnRejectReason = "1"
	CollAsgnRejectReason_UNAUTHORIZED_TRANSACTION      CollAsgnRejectReason = "2"
	CollAsgnRejectReason_INSUFFICIENT_COLLATERAL       CollAsgnRejectReason = "3"
	CollAsgnRejectReason_INVALID_TYPE_OF_COLLATERAL    CollAsgnRejectReason = "4"
	CollAsgnRejectReason_EXCESSIVE_SUBSTITUTION        CollAsgnRejectReason = "5"
	CollAsgnRejectReason_OTHER                         CollAsgnRejectReason = "99"
)

//Enum values for CollAsgnRespType
type CollAsgnRespType string

const (
	CollAsgnRespType_RECEIVED CollAsgnRespType = "0"
	CollAsgnRespType_ACCEPTED CollAsgnRespType = "1"
	CollAsgnRespType_DECLINED CollAsgnRespType = "2"
	CollAsgnRespType_REJECTED CollAsgnRespType = "3"
)

//Enum values for CollAsgnTransType
type CollAsgnTransType string

const (
	CollAsgnTransType_NEW     CollAsgnTransType = "0"
	CollAsgnTransType_REPLACE CollAsgnTransType = "1"
	CollAsgnTransType_CANCEL  CollAsgnTransType = "2"
	CollAsgnTransType_RELEASE CollAsgnTransType = "3"
	CollAsgnTransType_REVERSE CollAsgnTransType = "4"
)

//Enum values for CollInquiryQualifier
type CollInquiryQualifier string

const (
	CollInquiryQualifier_TRADE_DATE            CollInquiryQualifier = "0"
	CollInquiryQualifier_GC_INSTRUMENT         CollInquiryQualifier = "1"
	CollInquiryQualifier_COLLATERAL_INSTRUMENT CollInquiryQualifier = "2"
	CollInquiryQualifier_SUBSTITUTION_ELIGIBLE CollInquiryQualifier = "3"
	CollInquiryQualifier_NOT_ASSIGNED          CollInquiryQualifier = "4"
	CollInquiryQualifier_PARTIALLY_ASSIGNED    CollInquiryQualifier = "5"
	CollInquiryQualifier_FULLY_ASSIGNED        CollInquiryQualifier = "6"
	CollInquiryQualifier_OUTSTANDING_TRADES    CollInquiryQualifier = "7"
)

//Enum values for CollInquiryResult
type CollInquiryResult string

const (
	CollInquiryResult_SUCCESSFUL                                  CollInquiryResult = "0"
	CollInquiryResult_INVALID_OR_UNKNOWN_INSTRUMENT               CollInquiryResult = "1"
	CollInquiryResult_INVALID_OR_UNKNOWN_COLLATERAL_TYPE          CollInquiryResult = "2"
	CollInquiryResult_INVALID_PARTIES                             CollInquiryResult = "3"
	CollInquiryResult_INVALID_TRANSPORT_TYPE_REQUESTED            CollInquiryResult = "4"
	CollInquiryResult_INVALID_DESTINATION_REQUESTED               CollInquiryResult = "5"
	CollInquiryResult_NO_COLLATERAL_FOUND_FOR_THE_TRADE_SPECIFIED CollInquiryResult = "6"
	CollInquiryResult_NO_COLLATERAL_FOUND_FOR_THE_ORDER_SPECIFIED CollInquiryResult = "7"
	CollInquiryResult_COLLATERAL_INQUIRY_TYPE_NOT_SUPPORTED       CollInquiryResult = "8"
	CollInquiryResult_UNAUTHORIZED_FOR_COLLATERAL_INQUIRY         CollInquiryResult = "9"
	CollInquiryResult_OTHER                                       CollInquiryResult = "99"
)

//Enum values for CollInquiryStatus
type CollInquiryStatus string

const (
	CollInquiryStatus_ACCEPTED                CollInquiryStatus = "0"
	CollInquiryStatus_ACCEPTED_WITH_WARNINGS  CollInquiryStatus = "1"
	CollInquiryStatus_COMPLETED               CollInquiryStatus = "2"
	CollInquiryStatus_COMPLETED_WITH_WARNINGS CollInquiryStatus = "3"
	CollInquiryStatus_REJECTED                CollInquiryStatus = "4"
)

//Enum values for CollStatus
type CollStatus string

const (
	CollStatus_UNASSIGNED          CollStatus = "0"
	CollStatus_PARTIALLY_ASSIGNED  CollStatus = "1"
	CollStatus_ASSIGNMENT_PROPOSED CollStatus = "2"
	CollStatus_ASSIGNED            CollStatus = "3"
	CollStatus_CHALLENGED          CollStatus = "4"
)

//Enum values for CommType
type CommType string

const (
	CommType_PER_UNIT                    CommType = "1"
	CommType_PERCENT                     CommType = "2"
	CommType_ABSOLUTE                    CommType = "3"
	CommType_PERCENTAGE_WAIVED_4         CommType = "4"
	CommType_PERCENTAGE_WAIVED_5         CommType = "5"
	CommType_POINTS_PER_BOND_OR_CONTRACT CommType = "6"
)

//Enum values for ComplexEventCondition
type ComplexEventCondition string

const (
	ComplexEventCondition_AND ComplexEventCondition = "1"
	ComplexEventCondition_OR  ComplexEventCondition = "2"
)

//Enum values for ComplexEventPriceBoundaryMethod
type ComplexEventPriceBoundaryMethod string

const (
	ComplexEventPriceBoundaryMethod_LESS_THAN_COMPLEXEVENTPRICE                ComplexEventPriceBoundaryMethod = "1"
	ComplexEventPriceBoundaryMethod_LESS_THAN_OR_EQUAL_TO_COMPLEXEVENTPRICE    ComplexEventPriceBoundaryMethod = "2"
	ComplexEventPriceBoundaryMethod_EQUAL_TO_COMPLEXEVENTPRICE                 ComplexEventPriceBoundaryMethod = "3"
	ComplexEventPriceBoundaryMethod_GREATER_THAN_OR_EQUAL_TO_COMPLEXEVENTPRICE ComplexEventPriceBoundaryMethod = "4"
	ComplexEventPriceBoundaryMethod_GREATER_THAN_COMPLEXEVENTPRICE             ComplexEventPriceBoundaryMethod = "5"
)

//Enum values for ComplexEventPriceTimeType
type ComplexEventPriceTimeType string

const (
	ComplexEventPriceTimeType_EXPIRATION          ComplexEventPriceTimeType = "1"
	ComplexEventPriceTimeType_IMMEDIATE           ComplexEventPriceTimeType = "2"
	ComplexEventPriceTimeType_SPECIFIED_DATE_TIME ComplexEventPriceTimeType = "3"
)

//Enum values for ComplexEventType
type ComplexEventType string

const (
	ComplexEventType_CAPPED          ComplexEventType = "1"
	ComplexEventType_TRIGGER         ComplexEventType = "2"
	ComplexEventType_KNOCK_IN_UP     ComplexEventType = "3"
	ComplexEventType_KOCK_IN_DOWN    ComplexEventType = "4"
	ComplexEventType_KNOCK_OUT_UP    ComplexEventType = "5"
	ComplexEventType_KNOCK_OUT_DOWN  ComplexEventType = "6"
	ComplexEventType_UNDERLYING      ComplexEventType = "7"
	ComplexEventType_RESET_BARRIER   ComplexEventType = "8"
	ComplexEventType_ROLLING_BARRIER ComplexEventType = "9"
)

//Enum values for ConfirmRejReason
type ConfirmRejReason string

const (
	ConfirmRejReason_MISMATCHED_ACCOUNT              ConfirmRejReason = "1"
	ConfirmRejReason_MISSING_SETTLEMENT_INSTRUCTIONS ConfirmRejReason = "2"
	ConfirmRejReason_OTHER                           ConfirmRejReason = "99"
)

//Enum values for ConfirmStatus
type ConfirmStatus string

const (
	ConfirmStatus_RECEIVED                        ConfirmStatus = "1"
	ConfirmStatus_MISMATCHED_ACCOUNT              ConfirmStatus = "2"
	ConfirmStatus_MISSING_SETTLEMENT_INSTRUCTIONS ConfirmStatus = "3"
	ConfirmStatus_CONFIRMED                       ConfirmStatus = "4"
	ConfirmStatus_REQUEST_REJECTED                ConfirmStatus = "5"
)

//Enum values for ConfirmTransType
type ConfirmTransType string

const (
	ConfirmTransType_NEW     ConfirmTransType = "0"
	ConfirmTransType_REPLACE ConfirmTransType = "1"
	ConfirmTransType_CANCEL  ConfirmTransType = "2"
)

//Enum values for ConfirmType
type ConfirmType string

const (
	ConfirmType_STATUS                        ConfirmType = "1"
	ConfirmType_CONFIRMATION                  ConfirmType = "2"
	ConfirmType_CONFIRMATION_REQUEST_REJECTED ConfirmType = "3"
)

//Enum values for ContAmtType
type ContAmtType string

const (
	ContAmtType_COMMISSION_AMOUNT                       ContAmtType = "1"
	ContAmtType_EXIT_CHARGE_PERCENT                     ContAmtType = "10"
	ContAmtType_FUND_BASED_RENEWAL_COMMISSION_PERCENT   ContAmtType = "11"
	ContAmtType_PROJECTED_FUND_VALUE                    ContAmtType = "12"
	ContAmtType_FUND_BASED_RENEWAL_COMMISSION_AMOUNT_13 ContAmtType = "13"
	ContAmtType_FUND_BASED_RENEWAL_COMMISSION_AMOUNT_14 ContAmtType = "14"
	ContAmtType_NET_SETTLEMENT_AMOUNT                   ContAmtType = "15"
	ContAmtType_COMMISSION_PERCENT                      ContAmtType = "2"
	ContAmtType_INITIAL_CHARGE_AMOUNT                   ContAmtType = "3"
	ContAmtType_INITIAL_CHARGE_PERCENT                  ContAmtType = "4"
	ContAmtType_DISCOUNT_AMOUNT                         ContAmtType = "5"
	ContAmtType_DISCOUNT_PERCENT                        ContAmtType = "6"
	ContAmtType_DILUTION_LEVY_AMOUNT                    ContAmtType = "7"
	ContAmtType_DILUTION_LEVY_PERCENT                   ContAmtType = "8"
	ContAmtType_EXIT_CHARGE_AMOUNT                      ContAmtType = "9"
)

//Enum values for ContingencyType
type ContingencyType string

const (
	ContingencyType_ONE_CANCELS_THE_OTHER   ContingencyType = "1"
	ContingencyType_ONE_TRIGGERS_THE_OTHER  ContingencyType = "2"
	ContingencyType_ONE_UPDATES_THE_OTHER_3 ContingencyType = "3"
	ContingencyType_ONE_UPDATES_THE_OTHER_4 ContingencyType = "4"
)

//Enum values for ContractMultiplierUnit
type ContractMultiplierUnit string

const (
	ContractMultiplierUnit_SHARES ContractMultiplierUnit = "0"
	ContractMultiplierUnit_HOURS  ContractMultiplierUnit = "1"
	ContractMultiplierUnit_DAYS   ContractMultiplierUnit = "2"
)

//Enum values for CorporateAction
type CorporateAction string

const (
	CorporateAction_EX_DIVIDEND                  CorporateAction = "A"
	CorporateAction_EX_DISTRIBUTION              CorporateAction = "B"
	CorporateAction_EX_RIGHTS                    CorporateAction = "C"
	CorporateAction_NEW                          CorporateAction = "D"
	CorporateAction_EX_INTEREST                  CorporateAction = "E"
	CorporateAction_CASH_DIVIDEND                CorporateAction = "F"
	CorporateAction_STOCK_DIVIDEND               CorporateAction = "G"
	CorporateAction_NON_INTEGER_STOCK_SPLIT      CorporateAction = "H"
	CorporateAction_REVERSE_STOCK_SPLIT          CorporateAction = "I"
	CorporateAction_STANDARD_INTEGER_STOCK_SPLIT CorporateAction = "J"
	CorporateAction_POSITION_CONSOLIDATION       CorporateAction = "K"
	CorporateAction_LIQUIDATION_REORGANIZATION   CorporateAction = "L"
	CorporateAction_MERGER_REORGANIZATION        CorporateAction = "M"
	CorporateAction_RIGHTS_OFFERING              CorporateAction = "N"
	CorporateAction_SHAREHOLDER_MEETING          CorporateAction = "O"
	CorporateAction_SPINOFF                      CorporateAction = "P"
	CorporateAction_TENDER_OFFER                 CorporateAction = "Q"
	CorporateAction_WARRANT                      CorporateAction = "R"
	CorporateAction_SPECIAL_ACTION               CorporateAction = "S"
	CorporateAction_SYMBOL_CONVERSION            CorporateAction = "T"
	CorporateAction_CUSIP                        CorporateAction = "U"
	CorporateAction_LEAP_ROLLOVER                CorporateAction = "V"
	CorporateAction_SUCCESSION_EVENT             CorporateAction = "W"
)

//Enum values for CoveredOrUncovered
type CoveredOrUncovered string

const (
	CoveredOrUncovered_COVERED   CoveredOrUncovered = "0"
	CoveredOrUncovered_UNCOVERED CoveredOrUncovered = "1"
)

//Enum values for CrossPrioritization
type CrossPrioritization string

const (
	CrossPrioritization_NONE                     CrossPrioritization = "0"
	CrossPrioritization_BUY_SIDE_IS_PRIORITIZED  CrossPrioritization = "1"
	CrossPrioritization_SELL_SIDE_IS_PRIORITIZED CrossPrioritization = "2"
)

//Enum values for CrossType
type CrossType string

const (
	CrossType_CROSS_AON        CrossType = "1"
	CrossType_CROSS_IOC        CrossType = "2"
	CrossType_CROSS_ONE_SIDE   CrossType = "3"
	CrossType_CROSS_SAME_PRICE CrossType = "4"
)

//Enum values for CustOrderCapacity
type CustOrderCapacity string

const (
	CustOrderCapacity_MEMBER_TRADING_FOR_THEIR_OWN_ACCOUNT              CustOrderCapacity = "1"
	CustOrderCapacity_CLEARING_FIRM_TRADING_FOR_ITS_PROPRIETARY_ACCOUNT CustOrderCapacity = "2"
	CustOrderCapacity_MEMBER_TRADING_FOR_ANOTHER_MEMBER                 CustOrderCapacity = "3"
	CustOrderCapacity_ALL_OTHER                                         CustOrderCapacity = "4"
)

//Enum values for CustOrderHandlingInst
type CustOrderHandlingInst string

const (
	CustOrderHandlingInst_ADD_ON_ORDER                      CustOrderHandlingInst = "ADD"
	CustOrderHandlingInst_ALL_OR_NONE                       CustOrderHandlingInst = "AON"
	CustOrderHandlingInst_CASH_NOT_HELD                     CustOrderHandlingInst = "CNH"
	CustOrderHandlingInst_DIRECTED_ORDER                    CustOrderHandlingInst = "DIR"
	CustOrderHandlingInst_EXCHANGE_FOR_PHYSICAL_TRANSACTION CustOrderHandlingInst = "E.W"
	CustOrderHandlingInst_FILL_OR_KILL                      CustOrderHandlingInst = "FOK"
	CustOrderHandlingInst_IMBALANCE_ONLY                    CustOrderHandlingInst = "IO"
	CustOrderHandlingInst_IMMEDIATE_OR_CANCEL               CustOrderHandlingInst = "IOC"
	CustOrderHandlingInst_LIMIT_ON_CLOSE                    CustOrderHandlingInst = "LOC"
	CustOrderHandlingInst_LIMIT_ON_OPEN                     CustOrderHandlingInst = "LOO"
	CustOrderHandlingInst_MARKET_AT_CLOSE                   CustOrderHandlingInst = "MAC"
	CustOrderHandlingInst_MARKET_AT_OPEN                    CustOrderHandlingInst = "MAO"
	CustOrderHandlingInst_MARKET_ON_CLOSE                   CustOrderHandlingInst = "MOC"
	CustOrderHandlingInst_MARKET_ON_OPEN                    CustOrderHandlingInst = "MOO"
	CustOrderHandlingInst_MINIMUM_QUANTITY                  CustOrderHandlingInst = "MQT"
	CustOrderHandlingInst_NOT_HELD                          CustOrderHandlingInst = "NH"
	CustOrderHandlingInst_OVER_THE_DAY                      CustOrderHandlingInst = "OVD"
	CustOrderHandlingInst_PEGGED                            CustOrderHandlingInst = "PEG"
	CustOrderHandlingInst_RESERVE_SIZE_ORDER                CustOrderHandlingInst = "RSV"
	CustOrderHandlingInst_STOP_STOCK_TRANSACTION            CustOrderHandlingInst = "S.W"
	CustOrderHandlingInst_SCALE                             CustOrderHandlingInst = "SCL"
	CustOrderHandlingInst_TIME_ORDER                        CustOrderHandlingInst = "TMO"
	CustOrderHandlingInst_TRAILING_STOP                     CustOrderHandlingInst = "TS"
	CustOrderHandlingInst_WORK                              CustOrderHandlingInst = "WRK"
)

//Enum values for CustomerOrFirm
type CustomerOrFirm string

const (
	CustomerOrFirm_CUSTOMER CustomerOrFirm = "0"
	CustomerOrFirm_FIRM     CustomerOrFirm = "1"
)

//Enum values for CxlRejReason
type CxlRejReason string

const (
	CxlRejReason_TOO_LATE_TO_CANCEL                                        CxlRejReason = "0"
	CxlRejReason_UNKNOWN_ORDER                                             CxlRejReason = "1"
	CxlRejReason_INVALID_PRICE_INCREMENT                                   CxlRejReason = "18"
	CxlRejReason_BROKER                                                    CxlRejReason = "2"
	CxlRejReason_ORDER_ALREADY_IN_PENDING_CANCEL_OR_PENDING_REPLACE_STATUS CxlRejReason = "3"
	CxlRejReason_UNABLE_TO_PROCESS_ORDER_MASS_CANCEL_REQUEST               CxlRejReason = "4"
	CxlRejReason_ORIGORDMODTIME                                            CxlRejReason = "5"
	CxlRejReason_DUPLICATE_CLORDID                                         CxlRejReason = "6"
	CxlRejReason_PRICE_EXCEEDS_CURRENT_PRICE                               CxlRejReason = "7"
	CxlRejReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND                          CxlRejReason = "8"
	CxlRejReason_OTHER                                                     CxlRejReason = "99"
)

//Enum values for CxlRejResponseTo
type CxlRejResponseTo string

const (
	CxlRejResponseTo_ORDER_CANCEL_REQUEST         CxlRejResponseTo = "1"
	CxlRejResponseTo_ORDER_CANCEL_REPLACE_REQUEST CxlRejResponseTo = "2"
)

//Enum values for CxlType
type CxlType string

const (
	CxlType_FULL_REMAINING_QUANTITY CxlType = "F"
	CxlType_PARTIAL_CANCEL          CxlType = "P"
)

//Enum values for DKReason
type DKReason string

const (
	DKReason_UNKNOWN_SYMBOL         DKReason = "A"
	DKReason_WRONG_SIDE             DKReason = "B"
	DKReason_QUANTITY_EXCEEDS_ORDER DKReason = "C"
	DKReason_NO_MATCHING_ORDER      DKReason = "D"
	DKReason_PRICE_EXCEEDS_LIMIT    DKReason = "E"
	DKReason_CALCULATION_DIFFERENCE DKReason = "F"
	DKReason_OTHER                  DKReason = "Z"
)

//Enum values for DayBookingInst
type DayBookingInst string

const (
	DayBookingInst_CAN_TRIGGER_BOOKING_WITHOUT_REFERENCE_TO_THE_ORDER_INITIATOR DayBookingInst = "0"
	DayBookingInst_SPEAK_WITH_ORDER_INITIATOR_BEFORE_BOOKING                    DayBookingInst = "1"
	DayBookingInst_ACCUMULATE                                                   DayBookingInst = "2"
)

//Enum values for DealingCapacity
type DealingCapacity string

const (
	DealingCapacity_AGENT              DealingCapacity = "A"
	DealingCapacity_PRINCIPAL          DealingCapacity = "P"
	DealingCapacity_RISKLESS_PRINCIPAL DealingCapacity = "R"
)

//Enum values for DeleteReason
type DeleteReason string

const (
	DeleteReason_CANCELLATION DeleteReason = "0"
	DeleteReason_ERROR        DeleteReason = "1"
)

//Enum values for DeliveryForm
type DeliveryForm string

const (
	DeliveryForm_BOOK_ENTRY DeliveryForm = "1"
	DeliveryForm_BEARER     DeliveryForm = "2"
)

//Enum values for DeliveryType
type DeliveryType string

const (
	DeliveryType_VERSUS_PAYMENT_DELIVER DeliveryType = "0"
	DeliveryType_FREE_DELIVER           DeliveryType = "1"
	DeliveryType_TRI_PARTY              DeliveryType = "2"
	DeliveryType_HOLD_IN_CUSTODY        DeliveryType = "3"
)

//Enum values for DerivativeSecurityListRequestType
type DerivativeSecurityListRequestType string

const (
	DerivativeSecurityListRequestType_SYMBOL                                    DerivativeSecurityListRequestType = "0"
	DerivativeSecurityListRequestType_SECURITYTYPE_AND_OR_CFICODE               DerivativeSecurityListRequestType = "1"
	DerivativeSecurityListRequestType_PRODUCT                                   DerivativeSecurityListRequestType = "2"
	DerivativeSecurityListRequestType_TRADINGSESSIONID                          DerivativeSecurityListRequestType = "3"
	DerivativeSecurityListRequestType_ALL_SECURITIES                            DerivativeSecurityListRequestType = "4"
	DerivativeSecurityListRequestType_UNDELYINGSYMBOL                           DerivativeSecurityListRequestType = "5"
	DerivativeSecurityListRequestType_UNDERLYING_SECURITYTYPE_AND_OR_CFICODE    DerivativeSecurityListRequestType = "6"
	DerivativeSecurityListRequestType_UNDERLYING_PRODUCT                        DerivativeSecurityListRequestType = "7"
	DerivativeSecurityListRequestType_MARKETID_OR_MARKETID_PLUS_MARKETSEGMENTID DerivativeSecurityListRequestType = "8"
)

//Enum values for DeskOrderHandlingInst
type DeskOrderHandlingInst string

const (
	DeskOrderHandlingInst_ADD_ON_ORDER                      DeskOrderHandlingInst = "ADD"
	DeskOrderHandlingInst_ALL_OR_NONE                       DeskOrderHandlingInst = "AON"
	DeskOrderHandlingInst_CASH_NOT_HELD                     DeskOrderHandlingInst = "CNH"
	DeskOrderHandlingInst_DIRECTED_ORDER                    DeskOrderHandlingInst = "DIR"
	DeskOrderHandlingInst_EXCHANGE_FOR_PHYSICAL_TRANSACTION DeskOrderHandlingInst = "E.W"
	DeskOrderHandlingInst_FILL_OR_KILL                      DeskOrderHandlingInst = "FOK"
	DeskOrderHandlingInst_IMBALANCE_ONLY                    DeskOrderHandlingInst = "IO"
	DeskOrderHandlingInst_IMMEDIATE_OR_CANCEL               DeskOrderHandlingInst = "IOC"
	DeskOrderHandlingInst_LIMIT_ON_CLOSE                    DeskOrderHandlingInst = "LOC"
	DeskOrderHandlingInst_LIMIT_ON_OPEN                     DeskOrderHandlingInst = "LOO"
	DeskOrderHandlingInst_MARKET_AT_CLOSE                   DeskOrderHandlingInst = "MAC"
	DeskOrderHandlingInst_MARKET_AT_OPEN                    DeskOrderHandlingInst = "MAO"
	DeskOrderHandlingInst_MARKET_ON_CLOSE                   DeskOrderHandlingInst = "MOC"
	DeskOrderHandlingInst_MARKET_ON_OPEN                    DeskOrderHandlingInst = "MOO"
	DeskOrderHandlingInst_MINIMUM_QUANTITY                  DeskOrderHandlingInst = "MQT"
	DeskOrderHandlingInst_NOT_HELD                          DeskOrderHandlingInst = "NH"
	DeskOrderHandlingInst_OVER_THE_DAY                      DeskOrderHandlingInst = "OVD"
	DeskOrderHandlingInst_PEGGED                            DeskOrderHandlingInst = "PEG"
	DeskOrderHandlingInst_RESERVE_SIZE_ORDER                DeskOrderHandlingInst = "RSV"
	DeskOrderHandlingInst_STOP_STOCK_TRANSACTION            DeskOrderHandlingInst = "S.W"
	DeskOrderHandlingInst_SCALE                             DeskOrderHandlingInst = "SCL"
	DeskOrderHandlingInst_TIME_ORDER                        DeskOrderHandlingInst = "TMO"
	DeskOrderHandlingInst_TRAILING_STOP                     DeskOrderHandlingInst = "TS"
	DeskOrderHandlingInst_WORK                              DeskOrderHandlingInst = "WRK"
)

//Enum values for DeskType
type DeskType string

const (
	DeskType_AGENCY            DeskType = "A"
	DeskType_ARBITRAGE         DeskType = "AR"
	DeskType_DERIVATIVES       DeskType = "D"
	DeskType_INTERNATIONAL     DeskType = "IN"
	DeskType_INSTITUTIONAL     DeskType = "IS"
	DeskType_OTHER             DeskType = "O"
	DeskType_PREFERRED_TRADING DeskType = "PF"
	DeskType_PROPRIETARY       DeskType = "PR"
	DeskType_PROGRAM_TRADING   DeskType = "PT"
	DeskType_SALES             DeskType = "S"
	DeskType_TRADING           DeskType = "T"
)

//Enum values for DeskTypeSource
type DeskTypeSource string

const (
	DeskTypeSource_NASD_OATS DeskTypeSource = "1"
)

//Enum values for DiscretionInst
type DiscretionInst string

const (
	DiscretionInst_RELATED_TO_DISPLAYED_PRICE     DiscretionInst = "0"
	DiscretionInst_RELATED_TO_MARKET_PRICE        DiscretionInst = "1"
	DiscretionInst_RELATED_TO_PRIMARY_PRICE       DiscretionInst = "2"
	DiscretionInst_RELATED_TO_LOCAL_PRIMARY_PRICE DiscretionInst = "3"
	DiscretionInst_RELATED_TO_MIDPOINT_PRICE      DiscretionInst = "4"
	DiscretionInst_RELATED_TO_LAST_TRADE_PRICE    DiscretionInst = "5"
	DiscretionInst_RELATED_TO_VWAP                DiscretionInst = "6"
	DiscretionInst_AVERAGE_PRICE_GUARANTEE        DiscretionInst = "7"
)

//Enum values for DiscretionLimitType
type DiscretionLimitType string

const (
	DiscretionLimitType_OR_BETTER DiscretionLimitType = "0"
	DiscretionLimitType_STRICT    DiscretionLimitType = "1"
	DiscretionLimitType_OR_WORSE  DiscretionLimitType = "2"
)

//Enum values for DiscretionMoveType
type DiscretionMoveType string

const (
	DiscretionMoveType_FLOATING DiscretionMoveType = "0"
	DiscretionMoveType_FIXED    DiscretionMoveType = "1"
)

//Enum values for DiscretionOffsetType
type DiscretionOffsetType string

const (
	DiscretionOffsetType_PRICE        DiscretionOffsetType = "0"
	DiscretionOffsetType_BASIS_POINTS DiscretionOffsetType = "1"
	DiscretionOffsetType_TICKS        DiscretionOffsetType = "2"
	DiscretionOffsetType_PRICE_TIER   DiscretionOffsetType = "3"
)

//Enum values for DiscretionRoundDirection
type DiscretionRoundDirection string

const (
	DiscretionRoundDirection_MORE_AGGRESSIVE DiscretionRoundDirection = "1"
	DiscretionRoundDirection_MORE_PASSIVE    DiscretionRoundDirection = "2"
)

//Enum values for DiscretionScope
type DiscretionScope string

const (
	DiscretionScope_LOCAL                    DiscretionScope = "1"
	DiscretionScope_NATIONAL                 DiscretionScope = "2"
	DiscretionScope_GLOBAL                   DiscretionScope = "3"
	DiscretionScope_NATIONAL_EXCLUDING_LOCAL DiscretionScope = "4"
)

//Enum values for DisplayMethod
type DisplayMethod string

const (
	DisplayMethod_INITIAL     DisplayMethod = "1"
	DisplayMethod_NEW         DisplayMethod = "2"
	DisplayMethod_RANDOM      DisplayMethod = "3"
	DisplayMethod_UNDISCLOSED DisplayMethod = "4"
)

//Enum values for DisplayWhen
type DisplayWhen string

const (
	DisplayWhen_IMMEDIATE DisplayWhen = "1"
	DisplayWhen_EXHAUST   DisplayWhen = "2"
)

//Enum values for DistribPaymentMethod
type DistribPaymentMethod string

const (
	DistribPaymentMethod_CREST                            DistribPaymentMethod = "1"
	DistribPaymentMethod_BPAY                             DistribPaymentMethod = "10"
	DistribPaymentMethod_HIGH_VALUE_CLEARING_SYSTEM_HVACS DistribPaymentMethod = "11"
	DistribPaymentMethod_REINVEST_IN_FUND                 DistribPaymentMethod = "12"
	DistribPaymentMethod_NSCC                             DistribPaymentMethod = "2"
	DistribPaymentMethod_EUROCLEAR                        DistribPaymentMethod = "3"
	DistribPaymentMethod_CLEARSTREAM                      DistribPaymentMethod = "4"
	DistribPaymentMethod_CHEQUE                           DistribPaymentMethod = "5"
	DistribPaymentMethod_TELEGRAPHIC_TRANSFER             DistribPaymentMethod = "6"
	DistribPaymentMethod_FED_WIRE                         DistribPaymentMethod = "7"
	DistribPaymentMethod_DIRECT_CREDIT                    DistribPaymentMethod = "8"
	DistribPaymentMethod_ACH_CREDIT                       DistribPaymentMethod = "9"
)

//Enum values for DlvyInstType
type DlvyInstType string

const (
	DlvyInstType_CASH       DlvyInstType = "C"
	DlvyInstType_SECURITIES DlvyInstType = "S"
)

//Enum values for DueToRelated
type DueToRelated string

const (
	DueToRelated_NO  DueToRelated = "N"
	DueToRelated_YES DueToRelated = "Y"
)

//Enum values for EmailType
type EmailType string

const (
	EmailType_NEW         EmailType = "0"
	EmailType_REPLY       EmailType = "1"
	EmailType_ADMIN_REPLY EmailType = "2"
)

//Enum values for EncryptMethod
type EncryptMethod string

const (
	EncryptMethod_NONE_OTHER  EncryptMethod = "0"
	EncryptMethod_PKCS        EncryptMethod = "1"
	EncryptMethod_DES         EncryptMethod = "2"
	EncryptMethod_PKCS_DES    EncryptMethod = "3"
	EncryptMethod_PGP_DES     EncryptMethod = "4"
	EncryptMethod_PGP_DES_MD5 EncryptMethod = "5"
	EncryptMethod_PEM_DES_MD5 EncryptMethod = "6"
)

//Enum values for EventType
type EventType string

const (
	EventType_PUT                        EventType = "1"
	EventType_SWAP_ROLL_DATE             EventType = "10"
	EventType_SWAP_NEXT_START_DATE       EventType = "11"
	EventType_SWAP_NEXT_ROLL_DATE        EventType = "12"
	EventType_FIRST_DELIVERY_DATE        EventType = "13"
	EventType_LAST_DELIVERY_DATE         EventType = "14"
	EventType_INITIAL_INVENTORY_DUE_DATE EventType = "15"
	EventType_FINAL_INVENTORY_DUE_DATE   EventType = "16"
	EventType_FIRST_INTENT_DATE          EventType = "17"
	EventType_LAST_INTENT_DATE           EventType = "18"
	EventType_POSITION_REMOVAL_DATE      EventType = "19"
	EventType_CALL                       EventType = "2"
	EventType_TENDER                     EventType = "3"
	EventType_SINKING_FUND_CALL          EventType = "4"
	EventType_ACTIVATION                 EventType = "5"
	EventType_INACTIVIATION              EventType = "6"
	EventType_LAST_ELIGIBLE_TRADE_DATE   EventType = "7"
	EventType_SWAP_START_DATE            EventType = "8"
	EventType_SWAP_END_DATE              EventType = "9"
	EventType_OTHER                      EventType = "99"
)

//Enum values for ExDestination
type ExDestination string

const (
	ExDestination_NONE  ExDestination = "0"
	ExDestination_POSIT ExDestination = "4"
)

//Enum values for ExDestinationIDSource
type ExDestinationIDSource string

const (
	ExDestinationIDSource_BIC                                              ExDestinationIDSource = "B"
	ExDestinationIDSource_GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER ExDestinationIDSource = "C"
	ExDestinationIDSource_PROPRIETARY                                      ExDestinationIDSource = "D"
	ExDestinationIDSource_ISO_COUNTRY_CODE                                 ExDestinationIDSource = "E"
	ExDestinationIDSource_MIC                                              ExDestinationIDSource = "G"
)

//Enum values for ExchangeForPhysical
type ExchangeForPhysical string

const (
	ExchangeForPhysical_NO  ExchangeForPhysical = "N"
	ExchangeForPhysical_YES ExchangeForPhysical = "Y"
)

//Enum values for ExecAckStatus
type ExecAckStatus string

const (
	ExecAckStatus_RECEIVED_NOT_YET_PROCESSED ExecAckStatus = "0"
	ExecAckStatus_ACCEPTED                   ExecAckStatus = "1"
	ExecAckStatus_DONT_KNOW                  ExecAckStatus = "2"
)

//Enum values for ExecInst
type ExecInst string

const (
	ExecInst_STAY_ON_OFFER_SIDE                                    ExecInst = "0"
	ExecInst_NOT_HELD                                              ExecInst = "1"
	ExecInst_WORK                                                  ExecInst = "2"
	ExecInst_GO_ALONG                                              ExecInst = "3"
	ExecInst_OVER_THE_DAY                                          ExecInst = "4"
	ExecInst_HELD                                                  ExecInst = "5"
	ExecInst_PARTICIPANT_DONT_INITIATE                             ExecInst = "6"
	ExecInst_STRICT_SCALE                                          ExecInst = "7"
	ExecInst_TRY_TO_SCALE                                          ExecInst = "8"
	ExecInst_STAY_ON_BID_SIDE                                      ExecInst = "9"
	ExecInst_NO_CROSS                                              ExecInst = "A"
	ExecInst_OK_TO_CROSS                                           ExecInst = "B"
	ExecInst_CALL_FIRST                                            ExecInst = "C"
	ExecInst_PERCENT_OF_VOLUME                                     ExecInst = "D"
	ExecInst_DO_NOT_INCREASE                                       ExecInst = "E"
	ExecInst_DO_NOT_REDUCE                                         ExecInst = "F"
	ExecInst_ALL_OR_NONE                                           ExecInst = "G"
	ExecInst_REINSTATE_ON_SYSTEM_FAILURE                           ExecInst = "H"
	ExecInst_INSTITUTIONS_ONLY                                     ExecInst = "I"
	ExecInst_REINSTATE_ON_TRADING_HALT                             ExecInst = "J"
	ExecInst_CANCEL_ON_TRADING_HALT                                ExecInst = "K"
	ExecInst_LAST_PEG                                              ExecInst = "L"
	ExecInst_MID_PRICE_PEG                                         ExecInst = "M"
	ExecInst_NON_NEGOTIABLE                                        ExecInst = "N"
	ExecInst_OPENING_PEG                                           ExecInst = "O"
	ExecInst_MARKET_PEG                                            ExecInst = "P"
	ExecInst_CANCEL_ON_SYSTEM_FAILURE                              ExecInst = "Q"
	ExecInst_PRIMARY_PEG                                           ExecInst = "R"
	ExecInst_SUSPEND                                               ExecInst = "S"
	ExecInst_FIXED_PEG_TO_LOCAL_BEST_BID_OR_OFFER_AT_TIME_OF_ORDER ExecInst = "T"
	ExecInst_CUSTOMER_DISPLAY_INSTRUCTION                          ExecInst = "U"
	ExecInst_NETTING                                               ExecInst = "V"
	ExecInst_PEG_TO_VWAP                                           ExecInst = "W"
	ExecInst_TRADE_ALONG                                           ExecInst = "X"
	ExecInst_TRY_TO_STOP                                           ExecInst = "Y"
	ExecInst_CANCEL_IF_NOT_BEST                                    ExecInst = "Z"
	ExecInst_TRAILING_STOP_PEG                                     ExecInst = "a"
	ExecInst_STRICT_LIMIT                                          ExecInst = "b"
	ExecInst_IGNORE_PRICE_VALIDITY_CHECKS                          ExecInst = "c"
	ExecInst_PEG_TO_LIMIT_PRICE                                    ExecInst = "d"
	ExecInst_WORK_TO_TARGET_STRATEGY                               ExecInst = "e"
	ExecInst_INTERMARKET_SWEEP                                     ExecInst = "f"
	ExecInst_EXTERNAL_ROUTING_ALLOWED                              ExecInst = "g"
	ExecInst_EXTERNAL_ROUTING_NOT_ALLOWED                          ExecInst = "h"
	ExecInst_IMBALANCE_ONLY                                        ExecInst = "i"
	ExecInst_SINGLE_EXECUTION_REQUESTED_FOR_BLOCK_TRADE            ExecInst = "j"
	ExecInst_BEST_EXECUTION                                        ExecInst = "k"
	ExecInst_SUSPEND_ON_SYSTEM_FAILURE                             ExecInst = "l"
	ExecInst_SUSPEND_ON_TRADING_HALT                               ExecInst = "m"
	ExecInst_REINSTATE_ON_CONNECTION_LOSS                          ExecInst = "n"
	ExecInst_CANCEL_ON_CONNECTION_LOSS                             ExecInst = "o"
	ExecInst_SUSPEND_ON_CONNECTION_LOSS                            ExecInst = "p"
	ExecInst_RELEASE_FROM_SUSPENSION                               ExecInst = "q"
	ExecInst_EXECUTE_AS_DELTA_NEUTRAL_USING_VOLATILITY_PROVIDED    ExecInst = "r"
	ExecInst_EXECUTE_AS_DURATION_NEUTRAL                           ExecInst = "s"
	ExecInst_EXECUTE_AS_FX_NEUTRAL                                 ExecInst = "t"
)

//Enum values for ExecPriceType
type ExecPriceType string

const (
	ExecPriceType_BID_PRICE                              ExecPriceType = "B"
	ExecPriceType_CREATION_PRICE                         ExecPriceType = "C"
	ExecPriceType_CREATION_PRICE_PLUS_ADJUSTMENT_PERCENT ExecPriceType = "D"
	ExecPriceType_CREATION_PRICE_PLUS_ADJUSTMENT_AMOUNT  ExecPriceType = "E"
	ExecPriceType_OFFER_PRICE                            ExecPriceType = "O"
	ExecPriceType_OFFER_PRICE_MINUS_ADJUSTMENT_PERCENT   ExecPriceType = "P"
	ExecPriceType_OFFER_PRICE_MINUS_ADJUSTMENT_AMOUNT    ExecPriceType = "Q"
	ExecPriceType_SINGLE_PRICE                           ExecPriceType = "S"
)

//Enum values for ExecRestatementReason
type ExecRestatementReason string

const (
	ExecRestatementReason_GT_CORPORATE_ACTION         ExecRestatementReason = "0"
	ExecRestatementReason_GT_RENEWAL                  ExecRestatementReason = "1"
	ExecRestatementReason_WAREHOUSE_RECAP             ExecRestatementReason = "10"
	ExecRestatementReason_PEG_REFRESH                 ExecRestatementReason = "11"
	ExecRestatementReason_VERBAL_CHANGE               ExecRestatementReason = "2"
	ExecRestatementReason_REPRICING_OF_ORDER          ExecRestatementReason = "3"
	ExecRestatementReason_BROKER_OPTION               ExecRestatementReason = "4"
	ExecRestatementReason_PARTIAL_DECLINE_OF_ORDERQTY ExecRestatementReason = "5"
	ExecRestatementReason_CANCEL_ON_TRADING_HALT      ExecRestatementReason = "6"
	ExecRestatementReason_CANCEL_ON_SYSTEM_FAILURE    ExecRestatementReason = "7"
	ExecRestatementReason_MARKET                      ExecRestatementReason = "8"
	ExecRestatementReason_CANCELED_NOT_BEST           ExecRestatementReason = "9"
	ExecRestatementReason_OTHER                       ExecRestatementReason = "99"
)

//Enum values for ExecTransType
type ExecTransType string

const (
	ExecTransType_NEW     ExecTransType = "0"
	ExecTransType_CANCEL  ExecTransType = "1"
	ExecTransType_CORRECT ExecTransType = "2"
	ExecTransType_STATUS  ExecTransType = "3"
)

//Enum values for ExecType
type ExecType string

const (
	ExecType_NEW                                 ExecType = "0"
	ExecType_PARTIAL_FILL                        ExecType = "1"
	ExecType_FILL                                ExecType = "2"
	ExecType_DONE_FOR_DAY                        ExecType = "3"
	ExecType_CANCELED                            ExecType = "4"
	ExecType_REPLACED                            ExecType = "5"
	ExecType_PENDING_CANCEL                      ExecType = "6"
	ExecType_STOPPED                             ExecType = "7"
	ExecType_REJECTED                            ExecType = "8"
	ExecType_SUSPENDED                           ExecType = "9"
	ExecType_PENDING_NEW                         ExecType = "A"
	ExecType_CALCULATED                          ExecType = "B"
	ExecType_EXPIRED                             ExecType = "C"
	ExecType_RESTATED                            ExecType = "D"
	ExecType_PENDING_REPLACE                     ExecType = "E"
	ExecType_TRADE                               ExecType = "F"
	ExecType_TRADE_CORRECT                       ExecType = "G"
	ExecType_TRADE_CANCEL                        ExecType = "H"
	ExecType_ORDER_STATUS                        ExecType = "I"
	ExecType_TRADE_IN_A_CLEARING_HOLD            ExecType = "J"
	ExecType_TRADE_HAS_BEEN_RELEASED_TO_CLEARING ExecType = "K"
	ExecType_TRIGGERED_OR_ACTIVATED_BY_SYSTEM    ExecType = "L"
)

//Enum values for ExerciseMethod
type ExerciseMethod string

const (
	ExerciseMethod_AUTOMATIC ExerciseMethod = "A"
	ExerciseMethod_MANUAL    ExerciseMethod = "M"
)

//Enum values for ExerciseStyle
type ExerciseStyle string

const (
	ExerciseStyle_EUROPEAN ExerciseStyle = "0"
	ExerciseStyle_AMERICAN ExerciseStyle = "1"
	ExerciseStyle_BERMUDA  ExerciseStyle = "2"
)

//Enum values for ExpType
type ExpType string

const (
	ExpType_AUTO_EXERCISE           ExpType = "1"
	ExpType_NON_AUTO_EXERCISE       ExpType = "2"
	ExpType_FINAL_WILL_BE_EXERCISED ExpType = "3"
	ExpType_CONTRARY_INTENTION      ExpType = "4"
	ExpType_DIFFERENCE              ExpType = "5"
)

//Enum values for ExpirationCycle
type ExpirationCycle string

const (
	ExpirationCycle_EXPIRE_ON_TRADING_SESSION_CLOSE                                                ExpirationCycle = "0"
	ExpirationCycle_EXPIRE_ON_TRADING_SESSION_OPEN                                                 ExpirationCycle = "1"
	ExpirationCycle_TRADING_ELIGIBILITY_EXPIRATION_SPECIFIED_IN_THE_DATE_AND_TIME_FIELDS_EVENTDATE ExpirationCycle = "2"
)

//Enum values for ExpirationQtyType
type ExpirationQtyType string

const (
	ExpirationQtyType_AUTO_EXERCISE           ExpirationQtyType = "1"
	ExpirationQtyType_NON_AUTO_EXERCISE       ExpirationQtyType = "2"
	ExpirationQtyType_FINAL_WILL_BE_EXERCISED ExpirationQtyType = "3"
	ExpirationQtyType_CONTRARY_INTENTION      ExpirationQtyType = "4"
	ExpirationQtyType_DIFFERENCE              ExpirationQtyType = "5"
)

//Enum values for FinancialStatus
type FinancialStatus string

const (
	FinancialStatus_BANKRUPT          FinancialStatus = "1"
	FinancialStatus_PENDING_DELISTING FinancialStatus = "2"
	FinancialStatus_RESTRICTED        FinancialStatus = "3"
)

//Enum values for FlowScheduleType
type FlowScheduleType string

const (
	FlowScheduleType_NERC_EASTERN_OFF_PEAK           FlowScheduleType = "0"
	FlowScheduleType_NERC_WESTERN_OFF_PEAK           FlowScheduleType = "1"
	FlowScheduleType_NERC_CALENDAR_ALL_DAYS_IN_MONTH FlowScheduleType = "2"
	FlowScheduleType_NERC_EASTERN_PEAK               FlowScheduleType = "3"
	FlowScheduleType_NERC_WESTERN_PEAK               FlowScheduleType = "4"
)

//Enum values for ForexReq
type ForexReq string

const (
	ForexReq_NO  ForexReq = "N"
	ForexReq_YES ForexReq = "Y"
)

//Enum values for FundRenewWaiv
type FundRenewWaiv string

const (
	FundRenewWaiv_NO  FundRenewWaiv = "N"
	FundRenewWaiv_YES FundRenewWaiv = "Y"
)

//Enum values for FuturesValuationMethod
type FuturesValuationMethod string

const (
	FuturesValuationMethod_PREMIUM_STYLE                                  FuturesValuationMethod = "EQTY"
	FuturesValuationMethod_FUTURES_STYLE_MARK_TO_MARKET                   FuturesValuationMethod = "FUT"
	FuturesValuationMethod_FUTURES_STYLE_WITH_AN_ATTACHED_CASH_ADJUSTMENT FuturesValuationMethod = "FUTDA"
)

//Enum values for GTBookingInst
type GTBookingInst string

const (
	GTBookingInst_BOOK_OUT_ALL_TRADES_ON_DAY_OF_EXECUTION                 GTBookingInst = "0"
	GTBookingInst_ACCUMULATE_EXECTUIONS_UNTIL_FORDER_IS_FILLED_OR_EXPIRES GTBookingInst = "1"
	GTBookingInst_ACCUMULATE_UNTIL_VERBALLLY_NOTIFIED_OTHERWISE           GTBookingInst = "2"
)

//Enum values for GapFillFlag
type GapFillFlag string

const (
	GapFillFlag_NO  GapFillFlag = "N"
	GapFillFlag_YES GapFillFlag = "Y"
)

//Enum values for HaltReasonChar
type HaltReasonChar string

const (
	HaltReasonChar_NEWS_DISSEMINATION     HaltReasonChar = "D"
	HaltReasonChar_ORDER_INFLUX           HaltReasonChar = "E"
	HaltReasonChar_ORDER_IMBALANCE        HaltReasonChar = "I"
	HaltReasonChar_ADDITIONAL_INFORMATION HaltReasonChar = "M"
	HaltReasonChar_NEW_PENDING            HaltReasonChar = "P"
	HaltReasonChar_EQUIPMENT_CHANGEOVER   HaltReasonChar = "X"
)

//Enum values for HaltReasonInt
type HaltReasonInt string

const (
	HaltReasonInt_NEWS_DISSEMINATION     HaltReasonInt = "0"
	HaltReasonInt_ORDER_INFLUX           HaltReasonInt = "1"
	HaltReasonInt_ORDER_IMBALANCE        HaltReasonInt = "2"
	HaltReasonInt_ADDITIONAL_INFORMATION HaltReasonInt = "3"
	HaltReasonInt_NEWS_PENDING           HaltReasonInt = "4"
	HaltReasonInt_EQUIPMENT_CHANGEOVER   HaltReasonInt = "5"
)

//Enum values for HandlInst
type HandlInst string

const (
	HandlInst_AUTOMATED_EXECUTION_ORDER_PRIVATE_NO_BROKER_INTERVENTION HandlInst = "1"
	HandlInst_AUTOMATED_EXECUTION_ORDER_PUBLIC_BROKER_INTERVENTION_OK  HandlInst = "2"
	HandlInst_MANUAL_ORDER_BEST_EXECUTION                              HandlInst = "3"
)

//Enum values for IDSource
type IDSource string

const (
	IDSource_CUSIP                         IDSource = "1"
	IDSource_SEDOL                         IDSource = "2"
	IDSource_QUIK                          IDSource = "3"
	IDSource_ISIN_NUMBER                   IDSource = "4"
	IDSource_RIC_CODE                      IDSource = "5"
	IDSource_ISO_CURRENCY_CODE             IDSource = "6"
	IDSource_ISO_COUNTRY_CODE              IDSource = "7"
	IDSource_EXCHANGE_SYMBOL               IDSource = "8"
	IDSource_CONSOLIDATED_TAPE_ASSOCIATION IDSource = "9"
)

//Enum values for IOINaturalFlag
type IOINaturalFlag string

const (
	IOINaturalFlag_NO  IOINaturalFlag = "N"
	IOINaturalFlag_YES IOINaturalFlag = "Y"
)

//Enum values for IOIOthSvc
type IOIOthSvc string

const (
	IOIOthSvc_AUTEX  IOIOthSvc = "A"
	IOIOthSvc_BRIDGE IOIOthSvc = "B"
)

//Enum values for IOIQltyInd
type IOIQltyInd string

const (
	IOIQltyInd_HIGH   IOIQltyInd = "H"
	IOIQltyInd_LOW    IOIQltyInd = "L"
	IOIQltyInd_MEDIUM IOIQltyInd = "M"
)

//Enum values for IOIQty
type IOIQty string

const (
	IOIQty_1000000000           IOIQty = "0"
	IOIQty_LARGE                IOIQty = "L"
	IOIQty_MEDIUM               IOIQty = "M"
	IOIQty_SMALL                IOIQty = "S"
	IOIQty_UNDISCLOSED_QUANTITY IOIQty = "U"
)

//Enum values for IOIQualifier
type IOIQualifier string

const (
	IOIQualifier_ALL_OR_NONE          IOIQualifier = "A"
	IOIQualifier_MARKET_ON_CLOSE      IOIQualifier = "B"
	IOIQualifier_AT_THE_CLOSE         IOIQualifier = "C"
	IOIQualifier_VWAP                 IOIQualifier = "D"
	IOIQualifier_IN_TOUCH_WITH        IOIQualifier = "I"
	IOIQualifier_LIMIT                IOIQualifier = "L"
	IOIQualifier_MORE_BEHIND          IOIQualifier = "M"
	IOIQualifier_AT_THE_OPEN          IOIQualifier = "O"
	IOIQualifier_TAKING_A_POSITION    IOIQualifier = "P"
	IOIQualifier_AT_THE_MARKET        IOIQualifier = "Q"
	IOIQualifier_READY_TO_TRADE       IOIQualifier = "R"
	IOIQualifier_PORTFOLIO_SHOWN      IOIQualifier = "S"
	IOIQualifier_THROUGH_THE_DAY      IOIQualifier = "T"
	IOIQualifier_VERSUS               IOIQualifier = "V"
	IOIQualifier_INDICATION           IOIQualifier = "W"
	IOIQualifier_CROSSING_OPPORTUNITY IOIQualifier = "X"
	IOIQualifier_AT_THE_MIDPOINT      IOIQualifier = "Y"
	IOIQualifier_PRE_OPEN             IOIQualifier = "Z"
)

//Enum values for IOIShares
type IOIShares string

const (
	IOIShares_LARGE  IOIShares = "L"
	IOIShares_MEDIUM IOIShares = "M"
	IOIShares_SMALL  IOIShares = "S"
)

//Enum values for IOITransType
type IOITransType string

const (
	IOITransType_CANCEL  IOITransType = "C"
	IOITransType_NEW     IOITransType = "N"
	IOITransType_REPLACE IOITransType = "R"
)

//Enum values for ImpliedMarketIndicator
type ImpliedMarketIndicator string

const (
	ImpliedMarketIndicator_NOT_IMPLIED                     ImpliedMarketIndicator = "0"
	ImpliedMarketIndicator_IMPLIED_IN                      ImpliedMarketIndicator = "1"
	ImpliedMarketIndicator_IMPLIED_OUT                     ImpliedMarketIndicator = "2"
	ImpliedMarketIndicator_BOTH_IMPLIED_IN_AND_IMPLIED_OUT ImpliedMarketIndicator = "3"
)

//Enum values for InViewOfCommon
type InViewOfCommon string

const (
	InViewOfCommon_NO  InViewOfCommon = "N"
	InViewOfCommon_YES InViewOfCommon = "Y"
)

//Enum values for IncTaxInd
type IncTaxInd string

const (
	IncTaxInd_NET   IncTaxInd = "1"
	IncTaxInd_GROSS IncTaxInd = "2"
)

//Enum values for IndividualAllocType
type IndividualAllocType string

const (
	IndividualAllocType_SUB_ALLOCATE           IndividualAllocType = "1"
	IndividualAllocType_THIRD_PARTY_ALLOCATION IndividualAllocType = "2"
)

//Enum values for InstrAttribType
type InstrAttribType string

const (
	InstrAttribType_FLAT                                                                        InstrAttribType = "1"
	InstrAttribType_ORIGINAL_ISSUE_DISCOUNT                                                     InstrAttribType = "10"
	InstrAttribType_CALLABLE_PUTTABLE                                                           InstrAttribType = "11"
	InstrAttribType_ESCROWED_TO_MATURITY                                                        InstrAttribType = "12"
	InstrAttribType_ESCROWED_TO_REDEMPTION_DATE                                                 InstrAttribType = "13"
	InstrAttribType_PRE_REFUNDED                                                                InstrAttribType = "14"
	InstrAttribType_IN_DEFAULT                                                                  InstrAttribType = "15"
	InstrAttribType_UNRATED                                                                     InstrAttribType = "16"
	InstrAttribType_TAXABLE                                                                     InstrAttribType = "17"
	InstrAttribType_INDEXED                                                                     InstrAttribType = "18"
	InstrAttribType_SUBJECT_TO_ALTERNATIVE_MINIMUM_TAX                                          InstrAttribType = "19"
	InstrAttribType_ZERO_COUPON                                                                 InstrAttribType = "2"
	InstrAttribType_ORIGINAL_ISSUE_DISCOUNT_PRICE_SUPPLY_PRICE_IN_THE_INSTRATTRIBVALUE          InstrAttribType = "20"
	InstrAttribType_CALLABLE_BELOW_MATURITY_VALUE                                               InstrAttribType = "21"
	InstrAttribType_CALLABLE_WITHOUT_NOTICE_BY_MAIL_TO_HOLDER_UNLESS_REGISTERED                 InstrAttribType = "22"
	InstrAttribType_PRICE_TICK_RULES_FOR_SECURITY                                               InstrAttribType = "23"
	InstrAttribType_TRADE_TYPE_ELIGIBILITY_DETAILS_FOR_SECURITY                                 InstrAttribType = "24"
	InstrAttribType_INSTRUMENT_DENOMINATOR                                                      InstrAttribType = "25"
	InstrAttribType_INSTRUMENT_NUMERATOR                                                        InstrAttribType = "26"
	InstrAttribType_INSTRUMENT_PRICE_PRECISION                                                  InstrAttribType = "27"
	InstrAttribType_INSTRUMENT_STRIKE_PRICE                                                     InstrAttribType = "28"
	InstrAttribType_TRADEABLE_INDICATOR                                                         InstrAttribType = "29"
	InstrAttribType_INTEREST_BEARING                                                            InstrAttribType = "3"
	InstrAttribType_NO_PERIODIC_PAYMENTS                                                        InstrAttribType = "4"
	InstrAttribType_VARIABLE_RATE                                                               InstrAttribType = "5"
	InstrAttribType_LESS_FEE_FOR_PUT                                                            InstrAttribType = "6"
	InstrAttribType_STEPPED_COUPON                                                              InstrAttribType = "7"
	InstrAttribType_COUPON_PERIOD                                                               InstrAttribType = "8"
	InstrAttribType_WHEN_AND_IF_ISSUED                                                          InstrAttribType = "9"
	InstrAttribType_TEXT_SUPPLY_THE_TEXT_OF_THE_ATTRIBUTE_OR_DISCLAIMER_IN_THE_INSTRATTRIBVALUE InstrAttribType = "99"
)

//Enum values for InstrRegistry
type InstrRegistry string

const (
	InstrRegistry_CUSTODIAN InstrRegistry = "BIC"
	InstrRegistry_COUNTRY   InstrRegistry = "ISO"
	InstrRegistry_PHYSICAL  InstrRegistry = "ZZ"
)

//Enum values for LastCapacity
type LastCapacity string

const (
	LastCapacity_AGENT              LastCapacity = "1"
	LastCapacity_CROSS_AS_AGENT     LastCapacity = "2"
	LastCapacity_CROSS_AS_PRINCIPAL LastCapacity = "3"
	LastCapacity_PRINCIPAL          LastCapacity = "4"
)

//Enum values for LastFragment
type LastFragment string

const (
	LastFragment_NO  LastFragment = "N"
	LastFragment_YES LastFragment = "Y"
)

//Enum values for LastLiquidityInd
type LastLiquidityInd string

const (
	LastLiquidityInd_ADDED_LIQUIDITY      LastLiquidityInd = "1"
	LastLiquidityInd_REMOVED_LIQUIDITY    LastLiquidityInd = "2"
	LastLiquidityInd_LIQUIDITY_ROUTED_OUT LastLiquidityInd = "3"
	LastLiquidityInd_AUCTION              LastLiquidityInd = "4"
)

//Enum values for LastRptRequested
type LastRptRequested string

const (
	LastRptRequested_NO  LastRptRequested = "N"
	LastRptRequested_YES LastRptRequested = "Y"
)

//Enum values for LegSwapType
type LegSwapType string

const (
	LegSwapType_PAR_FOR_PAR       LegSwapType = "1"
	LegSwapType_MODIFIED_DURATION LegSwapType = "2"
	LegSwapType_RISK              LegSwapType = "4"
	LegSwapType_PROCEEDS          LegSwapType = "5"
)

//Enum values for LegalConfirm
type LegalConfirm string

const (
	LegalConfirm_NO  LegalConfirm = "N"
	LegalConfirm_YES LegalConfirm = "Y"
)

//Enum values for LiquidityIndType
type LiquidityIndType string

const (
	LiquidityIndType_5_DAY_MOVING_AVERAGE  LiquidityIndType = "1"
	LiquidityIndType_20_DAY_MOVING_AVERAGE LiquidityIndType = "2"
	LiquidityIndType_NORMAL_MARKET_SIZE    LiquidityIndType = "3"
	LiquidityIndType_OTHER                 LiquidityIndType = "4"
)

//Enum values for ListExecInstType
type ListExecInstType string

const (
	ListExecInstType_IMMEDIATE                   ListExecInstType = "1"
	ListExecInstType_WAIT_FOR_EXECUT_INSTRUCTION ListExecInstType = "2"
	ListExecInstType_EXCHANGE_SWITCH_CIV_ORDER_3 ListExecInstType = "3"
	ListExecInstType_EXCHANGE_SWITCH_CIV_ORDER_4 ListExecInstType = "4"
	ListExecInstType_EXCHANGE_SWITCH_CIV_ORDER_5 ListExecInstType = "5"
)

//Enum values for ListMethod
type ListMethod string

const (
	ListMethod_PRE_LISTED_ONLY ListMethod = "0"
	ListMethod_USER_REQUESTED  ListMethod = "1"
)

//Enum values for ListOrderStatus
type ListOrderStatus string

const (
	ListOrderStatus_IN_BIDDING_PROCESS     ListOrderStatus = "1"
	ListOrderStatus_RECEIVED_FOR_EXECUTION ListOrderStatus = "2"
	ListOrderStatus_EXECUTING              ListOrderStatus = "3"
	ListOrderStatus_CANCELLING             ListOrderStatus = "4"
	ListOrderStatus_ALERT                  ListOrderStatus = "5"
	ListOrderStatus_ALL_DONE               ListOrderStatus = "6"
	ListOrderStatus_REJECT                 ListOrderStatus = "7"
)

//Enum values for ListRejectReason
type ListRejectReason string

const (
	ListRejectReason_BROKER                           ListRejectReason = "0"
	ListRejectReason_UNSUPPORTED_ORDER_CHARACTERISTIC ListRejectReason = "11"
	ListRejectReason_EXCHANGE_CLOSED                  ListRejectReason = "2"
	ListRejectReason_TOO_LATE_TO_ENTER                ListRejectReason = "4"
	ListRejectReason_UNKNOWN_ORDER                    ListRejectReason = "5"
	ListRejectReason_DUPLICATE_ORDER                  ListRejectReason = "6"
	ListRejectReason_OTHER                            ListRejectReason = "99"
)

//Enum values for ListStatusType
type ListStatusType string

const (
	ListStatusType_ACK          ListStatusType = "1"
	ListStatusType_RESPONSE     ListStatusType = "2"
	ListStatusType_TIMED        ListStatusType = "3"
	ListStatusType_EXEC_STARTED ListStatusType = "4"
	ListStatusType_ALL_DONE     ListStatusType = "5"
	ListStatusType_ALERT        ListStatusType = "6"
)

//Enum values for LocateReqd
type LocateReqd string

const (
	LocateReqd_NO  LocateReqd = "N"
	LocateReqd_YES LocateReqd = "Y"
)

//Enum values for LotType
type LotType string

const (
	LotType_ODD_LOT                            LotType = "1"
	LotType_ROUND_LOT                          LotType = "2"
	LotType_BLOCK_LOT                          LotType = "3"
	LotType_ROUND_LOT_BASED_UPON_UNITOFMEASURE LotType = "4"
)

//Enum values for MDBookType
type MDBookType string

const (
	MDBookType_TOP_OF_BOOK MDBookType = "1"
	MDBookType_PRICE_DEPTH MDBookType = "2"
	MDBookType_ORDER_DEPTH MDBookType = "3"
)

//Enum values for MDEntryType
type MDEntryType string

const (
	MDEntryType_BID                                             MDEntryType = "0"
	MDEntryType_OFFER                                           MDEntryType = "1"
	MDEntryType_TRADE                                           MDEntryType = "2"
	MDEntryType_INDEX_VALUE                                     MDEntryType = "3"
	MDEntryType_OPENING_PRICE                                   MDEntryType = "4"
	MDEntryType_CLOSING_PRICE                                   MDEntryType = "5"
	MDEntryType_SETTLEMENT_PRICE                                MDEntryType = "6"
	MDEntryType_TRADING_SESSION_HIGH_PRICE                      MDEntryType = "7"
	MDEntryType_TRADING_SESSION_LOW_PRICE                       MDEntryType = "8"
	MDEntryType_TRADING_SESSION_VWAP_PRICE                      MDEntryType = "9"
	MDEntryType_IMBALANCE                                       MDEntryType = "A"
	MDEntryType_TRADE_VOLUME                                    MDEntryType = "B"
	MDEntryType_OPEN_INTEREST                                   MDEntryType = "C"
	MDEntryType_COMPOSITE_UNDERLYING_PRICE                      MDEntryType = "D"
	MDEntryType_SIMULATED_SELL_PRICE                            MDEntryType = "E"
	MDEntryType_SIMULATED_BUY_PRICE                             MDEntryType = "F"
	MDEntryType_MARGIN_RATE                                     MDEntryType = "G"
	MDEntryType_MID_PRICE                                       MDEntryType = "H"
	MDEntryType_EMPTY_BOOK                                      MDEntryType = "J"
	MDEntryType_SETTLE_HIGH_PRICE                               MDEntryType = "K"
	MDEntryType_SETTLE_LOW_PRICE                                MDEntryType = "L"
	MDEntryType_PRIOR_SETTLE_PRICE                              MDEntryType = "M"
	MDEntryType_SESSION_HIGH_BID                                MDEntryType = "N"
	MDEntryType_SESSION_LOW_OFFER                               MDEntryType = "O"
	MDEntryType_EARLY_PRICES                                    MDEntryType = "P"
	MDEntryType_AUCTION_CLEARING_PRICE                          MDEntryType = "Q"
	MDEntryType_DAILY_VALUE_ADJUSTMENT_FOR_LONG_POSITIONS       MDEntryType = "R"
	MDEntryType_SWAP_VALUE_FACTOR                               MDEntryType = "S"
	MDEntryType_CUMULATIVE_VALUE_ADJUSTMENT_FOR_LONG_POSITIONS  MDEntryType = "T"
	MDEntryType_DAILY_VALUE_ADJUSTMENT_FOR_SHORT_POSITIONS      MDEntryType = "U"
	MDEntryType_CUMULATIVE_VALUE_ADJUSTMENT_FOR_SHORT_POSITIONS MDEntryType = "V"
	MDEntryType_FIXING_PRICE                                    MDEntryType = "W"
	MDEntryType_CASH_RATE                                       MDEntryType = "X"
	MDEntryType_RECOVERY_RATE                                   MDEntryType = "Y"
	MDEntryType_RECOVERY_RATE_FOR_LONG                          MDEntryType = "Z"
	MDEntryType_RECOVERY_RATE_FOR_SHORT                         MDEntryType = "a"
)

//Enum values for MDImplicitDelete
type MDImplicitDelete string

const (
	MDImplicitDelete_NO  MDImplicitDelete = "N"
	MDImplicitDelete_YES MDImplicitDelete = "Y"
)

//Enum values for MDOriginType
type MDOriginType string

const (
	MDOriginType_BOOK     MDOriginType = "0"
	MDOriginType_OFF_BOOK MDOriginType = "1"
	MDOriginType_CROSS    MDOriginType = "2"
)

//Enum values for MDQuoteType
type MDQuoteType string

const (
	MDQuoteType_INDICATIVE               MDQuoteType = "0"
	MDQuoteType_TRADEABLE                MDQuoteType = "1"
	MDQuoteType_RESTRICTED_TRADEABLE     MDQuoteType = "2"
	MDQuoteType_COUNTER                  MDQuoteType = "3"
	MDQuoteType_INDICATIVE_AND_TRADEABLE MDQuoteType = "4"
)

//Enum values for MDReqRejReason
type MDReqRejReason string

const (
	MDReqRejReason_UNKNOWN_SYMBOL                      MDReqRejReason = "0"
	MDReqRejReason_DUPLICATE_MDREQID                   MDReqRejReason = "1"
	MDReqRejReason_INSUFFICIENT_BANDWIDTH              MDReqRejReason = "2"
	MDReqRejReason_INSUFFICIENT_PERMISSIONS            MDReqRejReason = "3"
	MDReqRejReason_UNSUPPORTED_SUBSCRIPTIONREQUESTTYPE MDReqRejReason = "4"
	MDReqRejReason_UNSUPPORTED_MARKETDEPTH             MDReqRejReason = "5"
	MDReqRejReason_UNSUPPORTED_MDUPDATETYPE            MDReqRejReason = "6"
	MDReqRejReason_UNSUPPORTED_AGGREGATEDBOOK          MDReqRejReason = "7"
	MDReqRejReason_UNSUPPORTED_MDENTRYTYPE             MDReqRejReason = "8"
	MDReqRejReason_UNSUPPORTED_TRADINGSESSIONID        MDReqRejReason = "9"
	MDReqRejReason_UNSUPPORTED_SCOPE                   MDReqRejReason = "A"
	MDReqRejReason_UNSUPPORTED_OPENCLOSESETTLEFLAG     MDReqRejReason = "B"
	MDReqRejReason_UNSUPPORTED_MDIMPLICITDELETE        MDReqRejReason = "C"
	MDReqRejReason_INSUFFICIENT_CREDIT                 MDReqRejReason = "D"
)

//Enum values for MDSecSizeType
type MDSecSizeType string

const (
	MDSecSizeType_CUSTOMER MDSecSizeType = "1"
)

//Enum values for MDUpdateAction
type MDUpdateAction string

const (
	MDUpdateAction_NEW         MDUpdateAction = "0"
	MDUpdateAction_CHANGE      MDUpdateAction = "1"
	MDUpdateAction_DELETE      MDUpdateAction = "2"
	MDUpdateAction_DELETE_THRU MDUpdateAction = "3"
	MDUpdateAction_DELETE_FROM MDUpdateAction = "4"
	MDUpdateAction_OVERLAY     MDUpdateAction = "5"
)

//Enum values for MDUpdateType
type MDUpdateType string

const (
	MDUpdateType_FULL_REFRESH        MDUpdateType = "0"
	MDUpdateType_INCREMENTAL_REFRESH MDUpdateType = "1"
)

//Enum values for MarketUpdateAction
type MarketUpdateAction string

const (
	MarketUpdateAction_ADD    MarketUpdateAction = "A"
	MarketUpdateAction_DELETE MarketUpdateAction = "D"
	MarketUpdateAction_MODIFY MarketUpdateAction = "M"
)

//Enum values for MassActionRejectReason
type MassActionRejectReason string

const (
	MassActionRejectReason_MASS_ACTION_NOT_SUPPORTED                        MassActionRejectReason = "0"
	MassActionRejectReason_INVALID_OR_UNKNOWN_SECURITY                      MassActionRejectReason = "1"
	MassActionRejectReason_INVALID_OR_UNKNOWN_SECURITY_ISSUER               MassActionRejectReason = "10"
	MassActionRejectReason_INVALID_OR_UNKNOWN_ISSUER_OF_UNDERLYING_SECURITY MassActionRejectReason = "11"
	MassActionRejectReason_INVALID_OR_UNKNOWN_UNDERLYING_SECURITY           MassActionRejectReason = "2"
	MassActionRejectReason_INVALID_OR_UNKNOWN_PRODUCT                       MassActionRejectReason = "3"
	MassActionRejectReason_INVALID_OR_UNKNOWN_CFICODE                       MassActionRejectReason = "4"
	MassActionRejectReason_INVALID_OR_UNKNOWN_SECURITYTYPE                  MassActionRejectReason = "5"
	MassActionRejectReason_INVALID_OR_UNKNOWN_TRADING_SESSION               MassActionRejectReason = "6"
	MassActionRejectReason_INVALID_OR_UNKNOWN_MARKET                        MassActionRejectReason = "7"
	MassActionRejectReason_INVALID_OR_UNKNOWN_MARKET_SEGMENT                MassActionRejectReason = "8"
	MassActionRejectReason_INVALID_OR_UNKNOWN_SECURITY_GROUP                MassActionRejectReason = "9"
	MassActionRejectReason_OTHER                                            MassActionRejectReason = "99"
)

//Enum values for MassActionResponse
type MassActionResponse string

const (
	MassActionResponse_REJECTED MassActionResponse = "0"
	MassActionResponse_ACCEPTED MassActionResponse = "1"
)

//Enum values for MassActionScope
type MassActionScope string

const (
	MassActionScope_ALL_ORDERS_FOR_A_SECURITY                MassActionScope = "1"
	MassActionScope_ALL_ORDERS_FOR_A_SECURITY_GROUP          MassActionScope = "10"
	MassActionScope_CANCEL_FOR_SECURITY_ISSUER               MassActionScope = "11"
	MassActionScope_CANCEL_FOR_ISSUER_OF_UNDERLYING_SECURITY MassActionScope = "12"
	MassActionScope_ALL_ORDERS_FOR_AN_UNDERLYING_SECURITY    MassActionScope = "2"
	MassActionScope_ALL_ORDERS_FOR_A_PRODUCT                 MassActionScope = "3"
	MassActionScope_ALL_ORDERS_FOR_A_CFICODE                 MassActionScope = "4"
	MassActionScope_ALL_ORDERS_FOR_A_SECURITYTYPE            MassActionScope = "5"
	MassActionScope_ALL_ORDERS_FOR_A_TRADING_SESSION         MassActionScope = "6"
	MassActionScope_ALL_ORDERS                               MassActionScope = "7"
	MassActionScope_ALL_ORDERS_FOR_A_MARKET                  MassActionScope = "8"
	MassActionScope_ALL_ORDERS_FOR_A_MARKET_SEGMENT          MassActionScope = "9"
)

//Enum values for MassActionType
type MassActionType string

const (
	MassActionType_SUSPEND_ORDERS                 MassActionType = "1"
	MassActionType_RELEASE_ORDERS_FROM_SUSPENSION MassActionType = "2"
	MassActionType_CANCEL_ORDERS                  MassActionType = "3"
)

//Enum values for MassCancelRejectReason
type MassCancelRejectReason string

const (
	MassCancelRejectReason_MASS_CANCEL_NOT_SUPPORTED                        MassCancelRejectReason = "0"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITY                      MassCancelRejectReason = "1"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITY_ISSUER               MassCancelRejectReason = "10"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_ISSUER_OF_UNDERLYING_SECURITY MassCancelRejectReason = "11"
	MassCancelRejectReason_INVALID_OR_UNKOWN_UNDERLYING_SECURITY            MassCancelRejectReason = "2"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_PRODUCT                       MassCancelRejectReason = "3"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_CFICODE                       MassCancelRejectReason = "4"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITYTYPE                  MassCancelRejectReason = "5"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_TRADING_SESSION               MassCancelRejectReason = "6"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_MARKET                        MassCancelRejectReason = "7"
	MassCancelRejectReason_INVALID_OR_UNKOWN_MARKET_SEGMENT                 MassCancelRejectReason = "8"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITY_GROUP                MassCancelRejectReason = "9"
	MassCancelRejectReason_OTHER                                            MassCancelRejectReason = "99"
)

//Enum values for MassCancelRequestType
type MassCancelRequestType string

const (
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_SECURITY             MassCancelRequestType = "1"
	MassCancelRequestType_CANCEL_ORDERS_FOR_AN_UNDERLYING_SECURITY MassCancelRequestType = "2"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_PRODUCT              MassCancelRequestType = "3"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_CFICODE              MassCancelRequestType = "4"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_SECURITYTYPE         MassCancelRequestType = "5"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_TRADING_SESSION      MassCancelRequestType = "6"
	MassCancelRequestType_CANCEL_ALL_ORDERS                        MassCancelRequestType = "7"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_MARKET               MassCancelRequestType = "8"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_MARKET_SEGMENT       MassCancelRequestType = "9"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_SECURITY_GROUP       MassCancelRequestType = "A"
	MassCancelRequestType_CANCEL_FOR_SECURITY_ISSUER               MassCancelRequestType = "B"
	MassCancelRequestType_CANCEL_FOR_ISSUER_OF_UNDERLYING_SECURITY MassCancelRequestType = "C"
)

//Enum values for MassCancelResponse
type MassCancelResponse string

const (
	MassCancelResponse_CANCEL_REQUEST_REJECTED                         MassCancelResponse = "0"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITY                    MassCancelResponse = "1"
	MassCancelResponse_CANCEL_ORDERS_FOR_AN_UNDERLYING_SECURITY        MassCancelResponse = "2"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_PRODUCT                     MassCancelResponse = "3"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_CFICODE                     MassCancelResponse = "4"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITYTYPE                MassCancelResponse = "5"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_TRADING_SESSION             MassCancelResponse = "6"
	MassCancelResponse_CANCEL_ALL_ORDERS                               MassCancelResponse = "7"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_MARKET                      MassCancelResponse = "8"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_MARKET_SEGMENT              MassCancelResponse = "9"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITY_GROUP              MassCancelResponse = "A"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITIES_ISSUER           MassCancelResponse = "B"
	MassCancelResponse_CANCEL_ORDERS_FOR_ISSUER_OF_UNDERLYING_SECURITY MassCancelResponse = "C"
)

//Enum values for MassStatusReqType
type MassStatusReqType string

const (
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_SECURITY             MassStatusReqType = "1"
	MassStatusReqType_STATUS_FOR_ISSUER_OF_UNDERLYING_SECURITY     MassStatusReqType = "10"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_AN_UNDERLYING_SECURITY MassStatusReqType = "2"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_PRODUCT              MassStatusReqType = "3"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_CFICODE              MassStatusReqType = "4"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_SECURITYTYPE         MassStatusReqType = "5"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_TRADING_SESSION      MassStatusReqType = "6"
	MassStatusReqType_STATUS_FOR_ALL_ORDERS                        MassStatusReqType = "7"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_PARTYID              MassStatusReqType = "8"
	MassStatusReqType_STATUS_FOR_SECURITY_ISSUER                   MassStatusReqType = "9"
)

//Enum values for MatchStatus
type MatchStatus string

const (
	MatchStatus_COMPARED_MATCHED_OR_AFFIRMED       MatchStatus = "0"
	MatchStatus_UNCOMPARED_UNMATCHED_OR_UNAFFIRMED MatchStatus = "1"
	MatchStatus_ADVISORY_OR_ALERT                  MatchStatus = "2"
)

//Enum values for MatchType
type MatchType string

const (
	MatchType_ONE_PARTY_TRADE_REPORT                                                                                                           MatchType = "1"
	MatchType_TWO_PARTY_TRADE_REPORT                                                                                                           MatchType = "2"
	MatchType_CONFIRMED_TRADE_REPORT                                                                                                           MatchType = "3"
	MatchType_AUTO_MATCH                                                                                                                       MatchType = "4"
	MatchType_CROSS_AUCTION                                                                                                                    MatchType = "5"
	MatchType_COUNTER_ORDER_SELECTION                                                                                                          MatchType = "6"
	MatchType_ONE_PARTY_PRIVATELY_NEGOTIATED_TRADE_REPORT                                                                                      MatchType = "60"
	MatchType_TWO_PARTY_PRIVATELY_NEGOTIATED_TRADE_REPORT                                                                                      MatchType = "61"
	MatchType_CONTINUOUS_AUTO_MATCH                                                                                                            MatchType = "62"
	MatchType_CROSS_AUCTION_63                                                                                                                 MatchType = "63"
	MatchType_COUNTER_ORDER_SELECTION_64                                                                                                       MatchType = "64"
	MatchType_CALL_AUCTION_65                                                                                                                  MatchType = "65"
	MatchType_CALL_AUCTION                                                                                                                     MatchType = "7"
	MatchType_ISSUING_BUY_BACK_AUCTION                                                                                                         MatchType = "8"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_FOUR_BADGES_AND_EXECUTION_TIME MatchType = "A1"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_FOUR_BADGES                    MatchType = "A2"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_TWO_BADGES_AND_EXECUTION_TIME  MatchType = "A3"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_TWO_BADGES                     MatchType = "A4"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADETYPE_AND_SPECIAL_TRADE_INDICATOR_PLUS_EXECUTION_TIME                  MatchType = "A5"
	MatchType_NASDAQACTM1MATCH                                                                                                                 MatchType = "ACTM1"
	MatchType_NASDAQACTM2MATCH                                                                                                                 MatchType = "ACTM2"
	MatchType_NASDAQACTACCEPTEDTRADE                                                                                                           MatchType = "ACTM3"
	MatchType_NASDAQACTDEFAULTTRADE                                                                                                            MatchType = "ACTM4"
	MatchType_NASDAQACTDEFAULTAFTERM2                                                                                                          MatchType = "ACTM5"
	MatchType_NASDAQACTM6MATCH                                                                                                                 MatchType = "ACTM6"
	MatchType_NASDAQNONACT                                                                                                                     MatchType = "ACTMT"
	MatchType_COMPARED_RECORDS_RESULTING_FROM_STAMPED_ADVISORIES_OR_SPECIALIST_ACCEPTS_PAIR_OFFS                                               MatchType = "AQ"
	MatchType_EXACT_MATCH_ON_TRADE_DATE_STOCK_SYMBOL_QUANTITY_PRICE_TRADE_TYPE_AND_SPECIAL_TRADE_INDICATOR_MINUS_BADGES_AND_TIMES_ACT_M1_MATCH MatchType = "M1"
	MatchType_SUMMARIZED_MATCH_MINUS_BADGES_AND_TIMES_ACT_M2_MATCH                                                                             MatchType = "M2"
	MatchType_ACT_ACCEPTED_TRADE                                                                                                               MatchType = "M3"
	MatchType_ACT_DEFAULT_TRADE                                                                                                                MatchType = "M4"
	MatchType_ACT_DEFAULT_AFTER_M2                                                                                                             MatchType = "M5"
	MatchType_ACT_M6_MATCH                                                                                                                     MatchType = "M6"
	MatchType_OCS_LOCKED_IN_NON_ACT                                                                                                            MatchType = "MT"
	MatchType_SUMMARIZED_MATCH_USING_A1_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIED                                                      MatchType = "S1"
	MatchType_SUMMARIZED_MATCH_USING_A2_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIZED                                                     MatchType = "S2"
	MatchType_SUMMARIZED_MATCH_USING_A3_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIZED                                                     MatchType = "S3"
	MatchType_SUMMARIZED_MATCH_USING_A4_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIZED                                                     MatchType = "S4"
	MatchType_SUMMARIZED_MATCH_USING_A5_EXACT_MATCH_CRITERIA_EXCEPT_QUANTITY_IS_SUMMARIZED                                                     MatchType = "S5"
)

//Enum values for MaturityMonthYearFormat
type MaturityMonthYearFormat string

const (
	MaturityMonthYearFormat_YEARMONTH_ONLY MaturityMonthYearFormat = "0"
	MaturityMonthYearFormat_YEARMONTHDAY   MaturityMonthYearFormat = "1"
	MaturityMonthYearFormat_YEARMONTHWEEK  MaturityMonthYearFormat = "2"
)

//Enum values for MaturityMonthYearIncrementUnits
type MaturityMonthYearIncrementUnits string

const (
	MaturityMonthYearIncrementUnits_MONTHS MaturityMonthYearIncrementUnits = "0"
	MaturityMonthYearIncrementUnits_DAYS   MaturityMonthYearIncrementUnits = "1"
	MaturityMonthYearIncrementUnits_WEEKS  MaturityMonthYearIncrementUnits = "2"
	MaturityMonthYearIncrementUnits_YEARS  MaturityMonthYearIncrementUnits = "3"
)

//Enum values for MessageEncoding
type MessageEncoding string

const (
	MessageEncoding_EUC_JP      MessageEncoding = "EUC-JP"
	MessageEncoding_ISO_2022_JP MessageEncoding = "ISO-2022-JP"
	MessageEncoding_SHIFT_JIS   MessageEncoding = "SHIFT_JIS"
	MessageEncoding_UTF_8       MessageEncoding = "UTF-8"
)

//Enum values for MiscFeeBasis
type MiscFeeBasis string

const (
	MiscFeeBasis_ABSOLUTE   MiscFeeBasis = "0"
	MiscFeeBasis_PER_UNIT   MiscFeeBasis = "1"
	MiscFeeBasis_PERCENTAGE MiscFeeBasis = "2"
)

//Enum values for MiscFeeType
type MiscFeeType string

const (
	MiscFeeType_REGULATORY       MiscFeeType = "1"
	MiscFeeType_PER_TRANSACTION  MiscFeeType = "10"
	MiscFeeType_CONVERSION       MiscFeeType = "11"
	MiscFeeType_AGENT            MiscFeeType = "12"
	MiscFeeType_TRANSFER_FEE     MiscFeeType = "13"
	MiscFeeType_SECURITY_LENDING MiscFeeType = "14"
	MiscFeeType_TAX              MiscFeeType = "2"
	MiscFeeType_LOCAL_COMMISSION MiscFeeType = "3"
	MiscFeeType_EXCHANGE_FEES    MiscFeeType = "4"
	MiscFeeType_STAMP            MiscFeeType = "5"
	MiscFeeType_LEVY             MiscFeeType = "6"
	MiscFeeType_OTHER            MiscFeeType = "7"
	MiscFeeType_MARKUP           MiscFeeType = "8"
	MiscFeeType_CONSUMPTION_TAX  MiscFeeType = "9"
)

//Enum values for ModelType
type ModelType string

const (
	ModelType_UTILITY_PROVIDED_STANDARD_MODEL ModelType = "0"
	ModelType_PROPRIETARY                     ModelType = "1"
)

//Enum values for MoneyLaunderingStatus
type MoneyLaunderingStatus string

const (
	MoneyLaunderingStatus_EXEMPT_1    MoneyLaunderingStatus = "1"
	MoneyLaunderingStatus_EXEMPT_2    MoneyLaunderingStatus = "2"
	MoneyLaunderingStatus_EXEMPT_3    MoneyLaunderingStatus = "3"
	MoneyLaunderingStatus_NOT_CHECKED MoneyLaunderingStatus = "N"
	MoneyLaunderingStatus_PASSED      MoneyLaunderingStatus = "Y"
)

//Enum values for MsgDirection
type MsgDirection string

const (
	MsgDirection_RECEIVE MsgDirection = "R"
	MsgDirection_SEND    MsgDirection = "S"
)

//Enum values for MsgType
type MsgType string

const (
	MsgType_HEARTBEAT                          MsgType = "0"
	MsgType_TEST_REQUEST                       MsgType = "1"
	MsgType_RESEND_REQUEST                     MsgType = "2"
	MsgType_REJECT                             MsgType = "3"
	MsgType_SEQUENCE_RESET                     MsgType = "4"
	MsgType_LOGOUT                             MsgType = "5"
	MsgType_INDICATION_OF_INTEREST             MsgType = "6"
	MsgType_ADVERTISEMENT                      MsgType = "7"
	MsgType_EXECUTION_REPORT                   MsgType = "8"
	MsgType_ORDER_CANCEL_REJECT                MsgType = "9"
	MsgType_LOGON                              MsgType = "A"
	MsgType_DERIVATIVE_SECURITY_LIST           MsgType = "AA"
	MsgType_NEW_ORDER_MULTILEG                 MsgType = "AB"
	MsgType_MULTILEG_ORDER_CANCEL_REPLACE      MsgType = "AC"
	MsgType_TRADE_CAPTURE_REPORT_REQUEST       MsgType = "AD"
	MsgType_TRADE_CAPTURE_REPORT               MsgType = "AE"
	MsgType_ORDER_MASS_STATUS_REQUEST          MsgType = "AF"
	MsgType_QUOTE_REQUEST_REJECT               MsgType = "AG"
	MsgType_RFQ_REQUEST                        MsgType = "AH"
	MsgType_QUOTE_STATUS_REPORT                MsgType = "AI"
	MsgType_QUOTE_RESPONSE                     MsgType = "AJ"
	MsgType_CONFIRMATION                       MsgType = "AK"
	MsgType_POSITION_MAINTENANCE_REQUEST       MsgType = "AL"
	MsgType_POSITION_MAINTENANCE_REPORT        MsgType = "AM"
	MsgType_REQUEST_FOR_POSITIONS              MsgType = "AN"
	MsgType_REQUEST_FOR_POSITIONS_ACK          MsgType = "AO"
	MsgType_POSITION_REPORT                    MsgType = "AP"
	MsgType_TRADE_CAPTURE_REPORT_REQUEST_ACK   MsgType = "AQ"
	MsgType_TRADE_CAPTURE_REPORT_ACK           MsgType = "AR"
	MsgType_ALLOCATION_REPORT                  MsgType = "AS"
	MsgType_ALLOCATION_REPORT_ACK              MsgType = "AT"
	MsgType_CONFIRMATION_ACK                   MsgType = "AU"
	MsgType_SETTLEMENT_INSTRUCTION_REQUEST     MsgType = "AV"
	MsgType_ASSIGNMENT_REPORT                  MsgType = "AW"
	MsgType_COLLATERAL_REQUEST                 MsgType = "AX"
	MsgType_COLLATERAL_ASSIGNMENT              MsgType = "AY"
	MsgType_COLLATERAL_RESPONSE                MsgType = "AZ"
	MsgType_NEWS                               MsgType = "B"
	MsgType_COLLATERAL_REPORT                  MsgType = "BA"
	MsgType_COLLATERAL_INQUIRY                 MsgType = "BB"
	MsgType_NETWORK_STATUS_REQUEST             MsgType = "BC"
	MsgType_NETWORK_STATUS_RESPONSE            MsgType = "BD"
	MsgType_USER_REQUEST                       MsgType = "BE"
	MsgType_USER_RESPONSE                      MsgType = "BF"
	MsgType_COLLATERAL_INQUIRY_ACK             MsgType = "BG"
	MsgType_CONFIRMATION_REQUEST               MsgType = "BH"
	MsgType_TRADING_SESSION_LIST_REQUEST       MsgType = "BI"
	MsgType_TRADING_SESSION_LIST               MsgType = "BJ"
	MsgType_SECURITY_LIST_UPDATE_REPORT        MsgType = "BK"
	MsgType_ADJUSTED_POSITION_REPORT           MsgType = "BL"
	MsgType_ALLOCATION_INSTRUCTION_ALERT       MsgType = "BM"
	MsgType_EXECUTION_ACKNOWLEDGEMENT          MsgType = "BN"
	MsgType_CONTRARY_INTENTION_REPORT          MsgType = "BO"
	MsgType_SECURITY_DEFINITION_UPDATE_REPORT  MsgType = "BP"
	MsgType_SETTLEMENTOBLIGATIONREPORT         MsgType = "BQ"
	MsgType_DERIVATIVESECURITYLISTUPDATEREPORT MsgType = "BR"
	MsgType_TRADINGSESSIONLISTUPDATEREPORT     MsgType = "BS"
	MsgType_MARKETDEFINITIONREQUEST            MsgType = "BT"
	MsgType_MARKETDEFINITION                   MsgType = "BU"
	MsgType_MARKETDEFINITIONUPDATEREPORT       MsgType = "BV"
	MsgType_APPLICATIONMESSAGEREQUEST          MsgType = "BW"
	MsgType_APPLICATIONMESSAGEREQUESTACK       MsgType = "BX"
	MsgType_APPLICATIONMESSAGEREPORT           MsgType = "BY"
	MsgType_ORDERMASSACTIONREPORT              MsgType = "BZ"
	MsgType_EMAIL                              MsgType = "C"
	MsgType_ORDERMASSACTIONREQUEST             MsgType = "CA"
	MsgType_USERNOTIFICATION                   MsgType = "CB"
	MsgType_STREAMASSIGNMENTREQUEST            MsgType = "CC"
	MsgType_STREAMASSIGNMENTREPORT             MsgType = "CD"
	MsgType_STREAMASSIGNMENTREPORTACK          MsgType = "CE"
	MsgType_PARTYDETAILSLISTREQUEST            MsgType = "CF"
	MsgType_PARTYDETAILSLISTREPORT             MsgType = "CG"
	MsgType_ORDER_SINGLE                       MsgType = "D"
	MsgType_ORDER_LIST                         MsgType = "E"
	MsgType_ORDER_CANCEL_REQUEST               MsgType = "F"
	MsgType_ORDER_CANCEL_REPLACE_REQUEST       MsgType = "G"
	MsgType_ORDER_STATUS_REQUEST               MsgType = "H"
	MsgType_ALLOCATION_INSTRUCTION             MsgType = "J"
	MsgType_LIST_CANCEL_REQUEST                MsgType = "K"
	MsgType_LIST_EXECUTE                       MsgType = "L"
	MsgType_LIST_STATUS_REQUEST                MsgType = "M"
	MsgType_LIST_STATUS                        MsgType = "N"
	MsgType_ALLOCATION_INSTRUCTION_ACK         MsgType = "P"
	MsgType_DONT_KNOW_TRADE                    MsgType = "Q"
	MsgType_QUOTE_REQUEST                      MsgType = "R"
	MsgType_QUOTE                              MsgType = "S"
	MsgType_SETTLEMENT_INSTRUCTIONS            MsgType = "T"
	MsgType_MARKET_DATA_REQUEST                MsgType = "V"
	MsgType_MARKET_DATA_SNAPSHOT_FULL_REFRESH  MsgType = "W"
	MsgType_MARKET_DATA_INCREMENTAL_REFRESH    MsgType = "X"
	MsgType_MARKET_DATA_REQUEST_REJECT         MsgType = "Y"
	MsgType_QUOTE_CANCEL                       MsgType = "Z"
	MsgType_QUOTE_STATUS_REQUEST               MsgType = "a"
	MsgType_MASS_QUOTE_ACKNOWLEDGEMENT         MsgType = "b"
	MsgType_SECURITY_DEFINITION_REQUEST        MsgType = "c"
	MsgType_SECURITY_DEFINITION                MsgType = "d"
	MsgType_SECURITY_STATUS_REQUEST            MsgType = "e"
	MsgType_SECURITY_STATUS                    MsgType = "f"
	MsgType_TRADING_SESSION_STATUS_REQUEST     MsgType = "g"
	MsgType_TRADING_SESSION_STATUS             MsgType = "h"
	MsgType_MASS_QUOTE                         MsgType = "i"
	MsgType_BUSINESS_MESSAGE_REJECT            MsgType = "j"
	MsgType_BID_REQUEST                        MsgType = "k"
	MsgType_BID_RESPONSE                       MsgType = "l"
	MsgType_LIST_STRIKE_PRICE                  MsgType = "m"
	MsgType_XML_MESSAGE                        MsgType = "n"
	MsgType_REGISTRATION_INSTRUCTIONS          MsgType = "o"
	MsgType_REGISTRATION_INSTRUCTIONS_RESPONSE MsgType = "p"
	MsgType_ORDER_MASS_CANCEL_REQUEST          MsgType = "q"
	MsgType_ORDER_MASS_CANCEL_REPORT           MsgType = "r"
	MsgType_NEW_ORDER_CROSS                    MsgType = "s"
	MsgType_CROSS_ORDER_CANCEL_REPLACE_REQUEST MsgType = "t"
	MsgType_CROSS_ORDER_CANCEL_REQUEST         MsgType = "u"
	MsgType_SECURITY_TYPE_REQUEST              MsgType = "v"
	MsgType_SECURITY_TYPES                     MsgType = "w"
	MsgType_SECURITY_LIST_REQUEST              MsgType = "x"
	MsgType_SECURITY_LIST                      MsgType = "y"
	MsgType_DERIVATIVE_SECURITY_LIST_REQUEST   MsgType = "z"
)

//Enum values for MultiLegReportingType
type MultiLegReportingType string

const (
	MultiLegReportingType_SINGLE_SECURITY                        MultiLegReportingType = "1"
	MultiLegReportingType_INDIVIDUAL_LEG_OF_A_MULTI_LEG_SECURITY MultiLegReportingType = "2"
	MultiLegReportingType_MULTI_LEG_SECURITY                     MultiLegReportingType = "3"
)

//Enum values for MultiLegRptTypeReq
type MultiLegRptTypeReq string

const (
	MultiLegRptTypeReq_REPORT_BY_MULITLEG_SECURITY_ONLY                                                      MultiLegRptTypeReq = "0"
	MultiLegRptTypeReq_REPORT_BY_MULTILEG_SECURITY_AND_BY_INSTRUMENT_LEGS_BELONGING_TO_THE_MULTILEG_SECURITY MultiLegRptTypeReq = "1"
	MultiLegRptTypeReq_REPORT_BY_INSTRUMENT_LEGS_BELONGING_TO_THE_MULTILEG_SECURITY_ONLY                     MultiLegRptTypeReq = "2"
)

//Enum values for MultilegModel
type MultilegModel string

const (
	MultilegModel_PREDEFINED_MULTILEG_SECURITY          MultilegModel = "0"
	MultilegModel_USER_DEFINED_MULTLEG_SECURITY         MultilegModel = "1"
	MultilegModel_USER_DEFINED_NON_SECURITIZED_MULTILEG MultilegModel = "2"
)

//Enum values for MultilegPriceMethod
type MultilegPriceMethod string

const (
	MultilegPriceMethod_NET_PRICE                       MultilegPriceMethod = "0"
	MultilegPriceMethod_REVERSED_NET_PRICE              MultilegPriceMethod = "1"
	MultilegPriceMethod_YIELD_DIFFERENCE                MultilegPriceMethod = "2"
	MultilegPriceMethod_INDIVIDUAL                      MultilegPriceMethod = "3"
	MultilegPriceMethod_CONTRACT_WEIGHTED_AVERAGE_PRICE MultilegPriceMethod = "4"
	MultilegPriceMethod_MULTIPLIED_PRICE                MultilegPriceMethod = "5"
)

//Enum values for NetGrossInd
type NetGrossInd string

const (
	NetGrossInd_NET   NetGrossInd = "1"
	NetGrossInd_GROSS NetGrossInd = "2"
)

//Enum values for NetworkRequestType
type NetworkRequestType string

const (
	NetworkRequestType_SNAPSHOT                                        NetworkRequestType = "1"
	NetworkRequestType_SUBSCRIBE                                       NetworkRequestType = "2"
	NetworkRequestType_STOP_SUBSCRIBING                                NetworkRequestType = "4"
	NetworkRequestType_LEVEL_OF_DETAIL_THEN_NOCOMPIDS_BECOMES_REQUIRED NetworkRequestType = "8"
)

//Enum values for NetworkStatusResponseType
type NetworkStatusResponseType string

const (
	NetworkStatusResponseType_FULL               NetworkStatusResponseType = "1"
	NetworkStatusResponseType_INCREMENTAL_UPDATE NetworkStatusResponseType = "2"
)

//Enum values for NewsCategory
type NewsCategory string

const (
	NewsCategory_COMPANY_NEWS          NewsCategory = "0"
	NewsCategory_MARKETPLACE_NEWS      NewsCategory = "1"
	NewsCategory_FINANCIAL_MARKET_NEWS NewsCategory = "2"
	NewsCategory_TECHNICAL_NEWS        NewsCategory = "3"
	NewsCategory_OTHER_NEWS            NewsCategory = "99"
)

//Enum values for NewsRefType
type NewsRefType string

const (
	NewsRefType_REPLACEMENT    NewsRefType = "0"
	NewsRefType_OTHER_LANGUAGE NewsRefType = "1"
	NewsRefType_COMPLIMENTARY  NewsRefType = "2"
)

//Enum values for NoSides
type NoSides string

const (
	NoSides_ONE_SIDE   NoSides = "1"
	NoSides_BOTH_SIDES NoSides = "2"
)

//Enum values for NotifyBrokerOfCredit
type NotifyBrokerOfCredit string

const (
	NotifyBrokerOfCredit_NO  NotifyBrokerOfCredit = "N"
	NotifyBrokerOfCredit_YES NotifyBrokerOfCredit = "Y"
)

//Enum values for OddLot
type OddLot string

const (
	OddLot_NO  OddLot = "N"
	OddLot_YES OddLot = "Y"
)

//Enum values for OpenClose
type OpenClose string

const (
	OpenClose_CLOSE OpenClose = "C"
	OpenClose_OPEN  OpenClose = "O"
)

//Enum values for OpenCloseSettlFlag
type OpenCloseSettlFlag string

const (
	OpenCloseSettlFlag_DAILY_OPEN                       OpenCloseSettlFlag = "0"
	OpenCloseSettlFlag_SESSION_OPEN                     OpenCloseSettlFlag = "1"
	OpenCloseSettlFlag_DELIVERY_SETTLEMENT_ENTRY        OpenCloseSettlFlag = "2"
	OpenCloseSettlFlag_EXPECTED_ENTRY                   OpenCloseSettlFlag = "3"
	OpenCloseSettlFlag_ENTRY_FROM_PREVIOUS_BUSINESS_DAY OpenCloseSettlFlag = "4"
	OpenCloseSettlFlag_THEORETICAL_PRICE_VALUE          OpenCloseSettlFlag = "5"
)

//Enum values for OpenCloseSettleFlag
type OpenCloseSettleFlag string

const (
	OpenCloseSettleFlag_DAILY_OPEN                       OpenCloseSettleFlag = "0"
	OpenCloseSettleFlag_SESSION_OPEN                     OpenCloseSettleFlag = "1"
	OpenCloseSettleFlag_DELIVERY_SETTLEMENT_PRICE        OpenCloseSettleFlag = "2"
	OpenCloseSettleFlag_EXPECTED_PRICE                   OpenCloseSettleFlag = "3"
	OpenCloseSettleFlag_PRICE_FROM_PREVIOUS_BUSINESS_DAY OpenCloseSettleFlag = "4"
)

//Enum values for OptPayoutType
type OptPayoutType string

const (
	OptPayoutType_VANILLA OptPayoutType = "1"
	OptPayoutType_CAPPED  OptPayoutType = "2"
	OptPayoutType_BINARY  OptPayoutType = "3"
)

//Enum values for OrdRejReason
type OrdRejReason string

const (
	OrdRejReason_BROKER                                     OrdRejReason = "0"
	OrdRejReason_UNKNOWN_SYMBOL                             OrdRejReason = "1"
	OrdRejReason_INVALID_INVESTOR_ID                        OrdRejReason = "10"
	OrdRejReason_UNSUPPORTED_ORDER_CHARACTERISTIC           OrdRejReason = "11"
	OrdRejReason_SURVEILLENCE_OPTION                        OrdRejReason = "12"
	OrdRejReason_INCORRECT_QUANTITY                         OrdRejReason = "13"
	OrdRejReason_INCORRECT_ALLOCATED_QUANTITY               OrdRejReason = "14"
	OrdRejReason_UNKNOWN_ACCOUNT                            OrdRejReason = "15"
	OrdRejReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND           OrdRejReason = "16"
	OrdRejReason_INVALID_PRICE_INCREMENT                    OrdRejReason = "18"
	OrdRejReason_EXCHANGE_CLOSED                            OrdRejReason = "2"
	OrdRejReason_ORDER_EXCEEDS_LIMIT                        OrdRejReason = "3"
	OrdRejReason_TOO_LATE_TO_ENTER                          OrdRejReason = "4"
	OrdRejReason_UNKNOWN_ORDER                              OrdRejReason = "5"
	OrdRejReason_DUPLICATE_ORDER                            OrdRejReason = "6"
	OrdRejReason_DUPLICATE_OF_A_VERBALLY_COMMUNICATED_ORDER OrdRejReason = "7"
	OrdRejReason_STALE_ORDER                                OrdRejReason = "8"
	OrdRejReason_TRADE_ALONG_REQUIRED                       OrdRejReason = "9"
	OrdRejReason_OTHER                                      OrdRejReason = "99"
)

//Enum values for OrdStatus
type OrdStatus string

const (
	OrdStatus_NEW                  OrdStatus = "0"
	OrdStatus_PARTIALLY_FILLED     OrdStatus = "1"
	OrdStatus_FILLED               OrdStatus = "2"
	OrdStatus_DONE_FOR_DAY         OrdStatus = "3"
	OrdStatus_CANCELED             OrdStatus = "4"
	OrdStatus_REPLACED             OrdStatus = "5"
	OrdStatus_PENDING_CANCEL       OrdStatus = "6"
	OrdStatus_STOPPED              OrdStatus = "7"
	OrdStatus_REJECTED             OrdStatus = "8"
	OrdStatus_SUSPENDED            OrdStatus = "9"
	OrdStatus_PENDING_NEW          OrdStatus = "A"
	OrdStatus_CALCULATED           OrdStatus = "B"
	OrdStatus_EXPIRED              OrdStatus = "C"
	OrdStatus_ACCEPTED_FOR_BIDDING OrdStatus = "D"
	OrdStatus_PENDING_REPLACE      OrdStatus = "E"
)

//Enum values for OrdType
type OrdType string

const (
	OrdType_MARKET                         OrdType = "1"
	OrdType_LIMIT                          OrdType = "2"
	OrdType_STOP                           OrdType = "3"
	OrdType_STOP_LIMIT                     OrdType = "4"
	OrdType_MARKET_ON_CLOSE                OrdType = "5"
	OrdType_WITH_OR_WITHOUT                OrdType = "6"
	OrdType_LIMIT_OR_BETTER                OrdType = "7"
	OrdType_LIMIT_WITH_OR_WITHOUT          OrdType = "8"
	OrdType_ON_BASIS                       OrdType = "9"
	OrdType_ON_CLOSE                       OrdType = "A"
	OrdType_LIMIT_ON_CLOSE                 OrdType = "B"
	OrdType_FOREX_MARKET                   OrdType = "C"
	OrdType_PREVIOUSLY_QUOTED              OrdType = "D"
	OrdType_PREVIOUSLY_INDICATED           OrdType = "E"
	OrdType_FOREX_LIMIT                    OrdType = "F"
	OrdType_FOREX_SWAP                     OrdType = "G"
	OrdType_FOREX_PREVIOUSLY_QUOTED        OrdType = "H"
	OrdType_FUNARI                         OrdType = "I"
	OrdType_MARKET_IF_TOUCHED              OrdType = "J"
	OrdType_MARKET_WITH_LEFT_OVER_AS_LIMIT OrdType = "K"
	OrdType_PREVIOUS_FUND_VALUATION_POINT  OrdType = "L"
	OrdType_NEXT_FUND_VALUATION_POINT      OrdType = "M"
	OrdType_PEGGED                         OrdType = "P"
	OrdType_COUNTER_ORDER_SELECTION        OrdType = "Q"
)

//Enum values for OrderCapacity
type OrderCapacity string

const (
	OrderCapacity_AGENCY                 OrderCapacity = "A"
	OrderCapacity_PROPRIETARY            OrderCapacity = "G"
	OrderCapacity_INDIVIDUAL             OrderCapacity = "I"
	OrderCapacity_PRINCIPAL              OrderCapacity = "P"
	OrderCapacity_RISKLESS_PRINCIPAL     OrderCapacity = "R"
	OrderCapacity_AGENT_FOR_OTHER_MEMBER OrderCapacity = "W"
)

//Enum values for OrderCategory
type OrderCategory string

const (
	OrderCategory_ORDER                      OrderCategory = "1"
	OrderCategory_QUOTE                      OrderCategory = "2"
	OrderCategory_PRIVATELY_NEGOTIATED_TRADE OrderCategory = "3"
	OrderCategory_MULTILEG_ORDER             OrderCategory = "4"
	OrderCategory_LINKED_ORDER               OrderCategory = "5"
	OrderCategory_QUOTE_REQUEST              OrderCategory = "6"
	OrderCategory_IMPLIED_ORDER              OrderCategory = "7"
	OrderCategory_CROSS_ORDER                OrderCategory = "8"
	OrderCategory_STREAMING_PRICE            OrderCategory = "9"
)

//Enum values for OrderDelayUnit
type OrderDelayUnit string

const (
	OrderDelayUnit_SECONDS                OrderDelayUnit = "0"
	OrderDelayUnit_TENTHS_OF_A_SECOND     OrderDelayUnit = "1"
	OrderDelayUnit_MINUTES                OrderDelayUnit = "10"
	OrderDelayUnit_HOURS                  OrderDelayUnit = "11"
	OrderDelayUnit_DAYS                   OrderDelayUnit = "12"
	OrderDelayUnit_WEEKS                  OrderDelayUnit = "13"
	OrderDelayUnit_MONTHS                 OrderDelayUnit = "14"
	OrderDelayUnit_YEARS                  OrderDelayUnit = "15"
	OrderDelayUnit_HUNDREDTHS_OF_A_SECOND OrderDelayUnit = "2"
	OrderDelayUnit_MILLISECONDS           OrderDelayUnit = "3"
	OrderDelayUnit_MICROSECONDS           OrderDelayUnit = "4"
	OrderDelayUnit_NANOSECONDS            OrderDelayUnit = "5"
)

//Enum values for OrderHandlingInstSource
type OrderHandlingInstSource string

const (
	OrderHandlingInstSource_NASD_OATS OrderHandlingInstSource = "1"
)

//Enum values for OrderRestrictions
type OrderRestrictions string

const (
	OrderRestrictions_PROGRAM_TRADE                                                                            OrderRestrictions = "1"
	OrderRestrictions_INDEX_ARBITRAGE                                                                          OrderRestrictions = "2"
	OrderRestrictions_NON_INDEX_ARBITRAGE                                                                      OrderRestrictions = "3"
	OrderRestrictions_COMPETING_MARKET_MAKER                                                                   OrderRestrictions = "4"
	OrderRestrictions_ACTING_AS_MARKET_MAKER_OR_SPECIALIST_IN_THE_SECURITY                                     OrderRestrictions = "5"
	OrderRestrictions_ACTING_AS_MARKET_MAKER_OR_SPECIALIST_IN_THE_UNDERLYING_SECURITY_OF_A_DERIVATIVE_SECURITY OrderRestrictions = "6"
	OrderRestrictions_FOREIGN_ENTITY                                                                           OrderRestrictions = "7"
	OrderRestrictions_EXTERNAL_MARKET_PARTICIPANT                                                              OrderRestrictions = "8"
	OrderRestrictions_EXTERNAL_INTER_CONNECTED_MARKET_LINKAGE                                                  OrderRestrictions = "9"
	OrderRestrictions_RISKLESS_ARBITRAGE                                                                       OrderRestrictions = "A"
	OrderRestrictions_ISSUER_HOLDING                                                                           OrderRestrictions = "B"
	OrderRestrictions_ISSUE_PRICE_STABILIZATION                                                                OrderRestrictions = "C"
	OrderRestrictions_NON_ALGORITHMIC                                                                          OrderRestrictions = "D"
	OrderRestrictions_ALGORITHMIC                                                                              OrderRestrictions = "E"
	OrderRestrictions_CROSS                                                                                    OrderRestrictions = "F"
)

//Enum values for OrigCustOrderCapacity
type OrigCustOrderCapacity string

const (
	OrigCustOrderCapacity_MEMBER_TRADING_FOR_THEIR_OWN_ACCOUNT              OrigCustOrderCapacity = "1"
	OrigCustOrderCapacity_CLEARING_FIRM_TRADING_FOR_ITS_PROPRIETARY_ACCOUNT OrigCustOrderCapacity = "2"
	OrigCustOrderCapacity_MEMBER_TRADING_FOR_ANOTHER_MEMBER                 OrigCustOrderCapacity = "3"
	OrigCustOrderCapacity_ALL_OTHER                                         OrigCustOrderCapacity = "4"
)

//Enum values for OwnerType
type OwnerType string

const (
	OwnerType_INDIVIDUAL_INVESTOR                 OwnerType = "1"
	OwnerType_NETWORKING_SUB_ACCOUNT              OwnerType = "10"
	OwnerType_NON_PROFIT_ORGANIZATION             OwnerType = "11"
	OwnerType_CORPORATE_BODY                      OwnerType = "12"
	OwnerType_NOMINEE                             OwnerType = "13"
	OwnerType_PUBLIC_COMPANY                      OwnerType = "2"
	OwnerType_PRIVATE_COMPANY                     OwnerType = "3"
	OwnerType_INDIVIDUAL_TRUSTEE                  OwnerType = "4"
	OwnerType_COMPANY_TRUSTEE                     OwnerType = "5"
	OwnerType_PENSION_PLAN                        OwnerType = "6"
	OwnerType_CUSTODIAN_UNDER_GIFTS_TO_MINORS_ACT OwnerType = "7"
	OwnerType_TRUSTS                              OwnerType = "8"
	OwnerType_FIDUCIARIES                         OwnerType = "9"
)

//Enum values for OwnershipType
type OwnershipType string

const (
	OwnershipType_JOINT_TRUSTEES    OwnershipType = "2"
	OwnershipType_JOINT_INVESTORS   OwnershipType = "J"
	OwnershipType_TENANTS_IN_COMMON OwnershipType = "T"
)

//Enum values for PartyDetailsRequestResult
type PartyDetailsRequestResult string

const (
	PartyDetailsRequestResult_VALID_REQUEST                                                   PartyDetailsRequestResult = "0"
	PartyDetailsRequestResult_INVALID_OR_UNSUPPORTED_REQUEST                                  PartyDetailsRequestResult = "1"
	PartyDetailsRequestResult_NO_PARTIES_OR_PARTY_DETAILS_FOUND_THAT_MATCH_SELECTION_CRITERIA PartyDetailsRequestResult = "2"
	PartyDetailsRequestResult_UNSUPPORTED_PARTYLISTRESPONSETYPE                               PartyDetailsRequestResult = "3"
	PartyDetailsRequestResult_NOT_AUTHORIZED_TO_RETRIEVE_PARTIES_OR_PARTY_DETAILS_DATA        PartyDetailsRequestResult = "4"
	PartyDetailsRequestResult_PARTIES_OR_PARTY_DETAILS_DATA_TEMPORARILY_UNAVAILABLE           PartyDetailsRequestResult = "5"
	PartyDetailsRequestResult_REQUEST_FOR_PARTIES_DATA_NOT_SUPPORTED                          PartyDetailsRequestResult = "6"
	PartyDetailsRequestResult_OTHER                                                           PartyDetailsRequestResult = "99"
)

//Enum values for PartyIDSource
type PartyIDSource string

const (
	PartyIDSource_KOREAN_INVESTOR_ID                                                                                PartyIDSource = "1"
	PartyIDSource_TAIWANESE_QUALIFIED_FOREIGN_INVESTOR_ID_QFII_FID                                                  PartyIDSource = "2"
	PartyIDSource_TAIWANESE_TRADING_ACCT                                                                            PartyIDSource = "3"
	PartyIDSource_MALAYSIAN_CENTRAL_DEPOSITORY                                                                      PartyIDSource = "4"
	PartyIDSource_CHINESE_INVESTOR_ID                                                                               PartyIDSource = "5"
	PartyIDSource_UK_NATIONAL_INSURANCE_OR_PENSION_NUMBER                                                           PartyIDSource = "6"
	PartyIDSource_US_SOCIAL_SECURITY_NUMBER                                                                         PartyIDSource = "7"
	PartyIDSource_US_EMPLOYER_OR_TAX_ID_NUMBER                                                                      PartyIDSource = "8"
	PartyIDSource_AUSTRALIAN_BUSINESS_NUMBER                                                                        PartyIDSource = "9"
	PartyIDSource_AUSTRALIAN_TAX_FILE_NUMBER                                                                        PartyIDSource = "A"
	PartyIDSource_BIC                                                                                               PartyIDSource = "B"
	PartyIDSource_GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER                                                  PartyIDSource = "C"
	PartyIDSource_PROPRIETARY                                                                                       PartyIDSource = "D"
	PartyIDSource_ISO_COUNTRY_CODE                                                                                  PartyIDSource = "E"
	PartyIDSource_SETTLEMENT_ENTITY_LOCATION                                                                        PartyIDSource = "F"
	PartyIDSource_MIC                                                                                               PartyIDSource = "G"
	PartyIDSource_CSD_PARTICIPANT_MEMBER_CODE                                                                       PartyIDSource = "H"
	PartyIDSource_DIRECTED_BROKER_THREE_CHARACTER_ACRONYM_AS_DEFINED_IN_ISITC_ETC_BEST_PRACTICE_GUIDELINES_DOCUMENT PartyIDSource = "I"
)

//Enum values for PartyListResponseType
type PartyListResponseType string

const (
	PartyListResponseType_RETURN_ALL_AVAILABLE_INFORMATION_ON_PARTIES_AND_RELATED_PARTIES PartyListResponseType = "0"
	PartyListResponseType_RETURN_ONLY_PARTY_INFORMATION                                   PartyListResponseType = "1"
	PartyListResponseType_INCLUDE_INFORMATION_ON_RELATED_PARTIES                          PartyListResponseType = "2"
	PartyListResponseType_INCLUDE_RISK_LIMIT_INFORMATION                                  PartyListResponseType = "3"
)

//Enum values for PartyRelationship
type PartyRelationship string

const (
	PartyRelationship_IS_ALSO                       PartyRelationship = "0"
	PartyRelationship_CLEARS_FOR                    PartyRelationship = "1"
	PartyRelationship_HAS_MEMBERS                   PartyRelationship = "10"
	PartyRelationship_PROVIDES_MARKETPLACE_FOR      PartyRelationship = "11"
	PartyRelationship_PARTICIPANT_OF_MARKETPLACE    PartyRelationship = "12"
	PartyRelationship_CARRIES_POSITIONS_FOR         PartyRelationship = "13"
	PartyRelationship_POSTS_TRADES_TO               PartyRelationship = "14"
	PartyRelationship_ENTERS_TRADES_FOR             PartyRelationship = "15"
	PartyRelationship_ENTERS_TRADES_THROUGH         PartyRelationship = "16"
	PartyRelationship_PROVIDES_QUOTES_TO            PartyRelationship = "17"
	PartyRelationship_REQUESTS_QUOTES_FROM          PartyRelationship = "18"
	PartyRelationship_INVESTS_FOR                   PartyRelationship = "19"
	PartyRelationship_CLEARS_THROUGH                PartyRelationship = "2"
	PartyRelationship_INVESTS_THROUGH               PartyRelationship = "20"
	PartyRelationship_BROKERS_TRADES_FOR            PartyRelationship = "21"
	PartyRelationship_BROKERS_TRADES_THROUGH        PartyRelationship = "22"
	PartyRelationship_PROVIDES_TRADING_SERVICES_FOR PartyRelationship = "23"
	PartyRelationship_USES_TRADING_SERVICES_OF      PartyRelationship = "24"
	PartyRelationship_APPROVES_OF                   PartyRelationship = "25"
	PartyRelationship_APPROVED_BY                   PartyRelationship = "26"
	PartyRelationship_PARENT_FIRM_FOR               PartyRelationship = "27"
	PartyRelationship_SUBSIDIARY_OF                 PartyRelationship = "28"
	PartyRelationship_REGULATORY_OWNER_OF           PartyRelationship = "29"
	PartyRelationship_TRADES_FOR                    PartyRelationship = "3"
	PartyRelationship_OWNED_BY_30                   PartyRelationship = "30"
	PartyRelationship_CONTROLS                      PartyRelationship = "31"
	PartyRelationship_IS_CONTROLLED_BY              PartyRelationship = "32"
	PartyRelationship_LEGAL                         PartyRelationship = "33"
	PartyRelationship_OWNED_BY_34                   PartyRelationship = "34"
	PartyRelationship_BENEFICIAL_OWNER_OF           PartyRelationship = "35"
	PartyRelationship_OWNED_BY_36                   PartyRelationship = "36"
	PartyRelationship_TRADES_THROUGH                PartyRelationship = "4"
	PartyRelationship_SPONSORS                      PartyRelationship = "5"
	PartyRelationship_SPONSORED_THROUGH             PartyRelationship = "6"
	PartyRelationship_PROVIDES_GUARANTEE_FOR        PartyRelationship = "7"
	PartyRelationship_IS_GUARANTEED_BY              PartyRelationship = "8"
	PartyRelationship_MEMBER_OF                     PartyRelationship = "9"
)

//Enum values for PartyRole
type PartyRole string

const (
	PartyRole_EXECUTING_FIRM                                                        PartyRole = "1"
	PartyRole_SETTLEMENT_LOCATION                                                   PartyRole = "10"
	PartyRole_ORDER_ORIGINATION_TRADER                                              PartyRole = "11"
	PartyRole_EXECUTING_TRADER                                                      PartyRole = "12"
	PartyRole_ORDER_ORIGINATION_FIRM                                                PartyRole = "13"
	PartyRole_GIVEUP_CLEARING_FIRM                                                  PartyRole = "14"
	PartyRole_CORRESPONDANT_CLEARING_FIRM                                           PartyRole = "15"
	PartyRole_EXECUTING_SYSTEM                                                      PartyRole = "16"
	PartyRole_CONTRA_FIRM                                                           PartyRole = "17"
	PartyRole_CONTRA_CLEARING_FIRM                                                  PartyRole = "18"
	PartyRole_SPONSORING_FIRM                                                       PartyRole = "19"
	PartyRole_BROKER_OF_CREDIT                                                      PartyRole = "2"
	PartyRole_UNDERLYING_CONTRA_FIRM                                                PartyRole = "20"
	PartyRole_CLEARING_ORGANIZATION                                                 PartyRole = "21"
	PartyRole_EXCHANGE                                                              PartyRole = "22"
	PartyRole_CUSTOMER_ACCOUNT                                                      PartyRole = "24"
	PartyRole_CORRESPONDENT_CLEARING_ORGANIZATION                                   PartyRole = "25"
	PartyRole_CORRESPONDENT_BROKER                                                  PartyRole = "26"
	PartyRole_BUYER_SELLER                                                          PartyRole = "27"
	PartyRole_CUSTODIAN                                                             PartyRole = "28"
	PartyRole_INTERMEDIARY                                                          PartyRole = "29"
	PartyRole_CLIENT_ID                                                             PartyRole = "3"
	PartyRole_AGENT                                                                 PartyRole = "30"
	PartyRole_SUB_CUSTODIAN                                                         PartyRole = "31"
	PartyRole_BENEFICIARY                                                           PartyRole = "32"
	PartyRole_INTERESTED_PARTY                                                      PartyRole = "33"
	PartyRole_REGULATORY_BODY                                                       PartyRole = "34"
	PartyRole_LIQUIDITY_PROVIDER                                                    PartyRole = "35"
	PartyRole_ENTERING_TRADER                                                       PartyRole = "36"
	PartyRole_CONTRA_TRADER                                                         PartyRole = "37"
	PartyRole_POSITION_ACCOUNT                                                      PartyRole = "38"
	PartyRole_CONTRA_INVESTOR_ID                                                    PartyRole = "39"
	PartyRole_CLEARING_FIRM                                                         PartyRole = "4"
	PartyRole_TRANSFER_TO_FIRM                                                      PartyRole = "40"
	PartyRole_CONTRA_POSITION_ACCOUNT                                               PartyRole = "41"
	PartyRole_CONTRA_EXCHANGE                                                       PartyRole = "42"
	PartyRole_INTERNAL_CARRY_ACCOUNT                                                PartyRole = "43"
	PartyRole_ORDER_ENTRY_OPERATOR_ID                                               PartyRole = "44"
	PartyRole_SECONDARY_ACCOUNT_NUMBER                                              PartyRole = "45"
	PartyRole_FOREIGN_FIRM                                                          PartyRole = "46"
	PartyRole_THIRD_PARTY_ALLOCATION_FIRM                                           PartyRole = "47"
	PartyRole_CLAIMING_ACCOUNT                                                      PartyRole = "48"
	PartyRole_ASSET_MANAGER                                                         PartyRole = "49"
	PartyRole_INVESTOR_ID                                                           PartyRole = "5"
	PartyRole_PLEDGOR_ACCOUNT                                                       PartyRole = "50"
	PartyRole_PLEDGEE_ACCOUNT                                                       PartyRole = "51"
	PartyRole_LARGE_TRADER_REPORTABLE_ACCOUNT                                       PartyRole = "52"
	PartyRole_TRADER_MNEMONIC                                                       PartyRole = "53"
	PartyRole_SENDER_LOCATION                                                       PartyRole = "54"
	PartyRole_SESSION_ID                                                            PartyRole = "55"
	PartyRole_ACCEPTABLE_COUNTERPARTY                                               PartyRole = "56"
	PartyRole_UNACCEPTABLE_COUNTERPARTY                                             PartyRole = "57"
	PartyRole_ENTERING_UNIT                                                         PartyRole = "58"
	PartyRole_EXECUTING_UNIT                                                        PartyRole = "59"
	PartyRole_INTRODUCING_FIRM                                                      PartyRole = "6"
	PartyRole_INTRODUCING_BROKER                                                    PartyRole = "60"
	PartyRole_QUOTE_ORIGINATOR                                                      PartyRole = "61"
	PartyRole_REPORT_ORIGINATOR                                                     PartyRole = "62"
	PartyRole_SYSTEMATIC_INTERNALISER                                               PartyRole = "63"
	PartyRole_MULTILATERAL_TRADING_FACILITY                                         PartyRole = "64"
	PartyRole_REGULATED_MARKET                                                      PartyRole = "65"
	PartyRole_MARKET_MAKER                                                          PartyRole = "66"
	PartyRole_INVESTMENT_FIRM                                                       PartyRole = "67"
	PartyRole_HOST_COMPETENT_AUTHORITY                                              PartyRole = "68"
	PartyRole_HOME_COMPETENT_AUTHORITY                                              PartyRole = "69"
	PartyRole_ENTERING_FIRM                                                         PartyRole = "7"
	PartyRole_COMPETENT_AUTHORITY_OF_THE_MOST_RELEVANT_MARKET_IN_TERMS_OF_LIQUIDITY PartyRole = "70"
	PartyRole_COMPETENT_AUTHORITY_OF_THE_TRANSACTION                                PartyRole = "71"
	PartyRole_REPORTING_INTERMEDIARY                                                PartyRole = "72"
	PartyRole_EXECUTION_VENUE                                                       PartyRole = "73"
	PartyRole_MARKET_DATA_ENTRY_ORIGINATOR                                          PartyRole = "74"
	PartyRole_LOCATION_ID                                                           PartyRole = "75"
	PartyRole_DESK_ID                                                               PartyRole = "76"
	PartyRole_MARKET_DATA_MARKET                                                    PartyRole = "77"
	PartyRole_ALLOCATION_ENTITY                                                     PartyRole = "78"
	PartyRole_PRIME_BROKER_PROVIDING_GENERAL_TRADE_SERVICES                         PartyRole = "79"
	PartyRole_LOCATE                                                                PartyRole = "8"
	PartyRole_STEP_OUT_FIRM                                                         PartyRole = "80"
	PartyRole_BROKERCLEARINGID                                                      PartyRole = "81"
	PartyRole_CENTRAL_REGISTRATION_DEPOSITORY                                       PartyRole = "82"
	PartyRole_CLEARING_ACCOUNT                                                      PartyRole = "83"
	PartyRole_ACCEPTABLE_SETTLING_COUNTERPARTY                                      PartyRole = "84"
	PartyRole_UNACCEPTABLE_SETTLING_COUNTERPARTY                                    PartyRole = "85"
	PartyRole_FUND_MANAGER_CLIENT_ID                                                PartyRole = "9"
)

//Enum values for PartySubIDType
type PartySubIDType string

const (
	PartySubIDType_FIRM                                                          PartySubIDType = "1"
	PartySubIDType_SECURITIES_ACCOUNT_NUMBER                                     PartySubIDType = "10"
	PartySubIDType_REGISTRATION_NUMBER                                           PartySubIDType = "11"
	PartySubIDType_REGISTERED_ADDRESS_12                                         PartySubIDType = "12"
	PartySubIDType_REGULATORY_STATUS                                             PartySubIDType = "13"
	PartySubIDType_REGISTRATION_NAME                                             PartySubIDType = "14"
	PartySubIDType_CASH_ACCOUNT_NUMBER                                           PartySubIDType = "15"
	PartySubIDType_BIC                                                           PartySubIDType = "16"
	PartySubIDType_CSD_PARTICIPANT_MEMBER_CODE                                   PartySubIDType = "17"
	PartySubIDType_REGISTERED_ADDRESS_18                                         PartySubIDType = "18"
	PartySubIDType_FUND_ACCOUNT_NAME                                             PartySubIDType = "19"
	PartySubIDType_PERSON                                                        PartySubIDType = "2"
	PartySubIDType_TELEX_NUMBER                                                  PartySubIDType = "20"
	PartySubIDType_FAX_NUMBER                                                    PartySubIDType = "21"
	PartySubIDType_SECURITIES_ACCOUNT_NAME                                       PartySubIDType = "22"
	PartySubIDType_CASH_ACCOUNT_NAME                                             PartySubIDType = "23"
	PartySubIDType_DEPARTMENT                                                    PartySubIDType = "24"
	PartySubIDType_LOCATION_DESK                                                 PartySubIDType = "25"
	PartySubIDType_POSITION_ACCOUNT_TYPE                                         PartySubIDType = "26"
	PartySubIDType_SECURITY_LOCATE_ID                                            PartySubIDType = "27"
	PartySubIDType_MARKET_MAKER                                                  PartySubIDType = "28"
	PartySubIDType_ELIGIBLE_COUNTERPARTY                                         PartySubIDType = "29"
	PartySubIDType_SYSTEM                                                        PartySubIDType = "3"
	PartySubIDType_PROFESSIONAL_CLIENT                                           PartySubIDType = "30"
	PartySubIDType_LOCATION                                                      PartySubIDType = "31"
	PartySubIDType_EXECUTION_VENUE                                               PartySubIDType = "32"
	PartySubIDType_CURRENCY_DELIVERY_IDENTIFIER                                  PartySubIDType = "33"
	PartySubIDType_APPLICATION                                                   PartySubIDType = "4"
	PartySubIDType_RESERVEDANDAVAILABLEFORBILATERALLYAGREEDUPONUSERDEFINEDVALUES PartySubIDType = "4000"
	PartySubIDType_FULL_LEGAL_NAME_OF_FIRM                                       PartySubIDType = "5"
	PartySubIDType_POSTAL_ADDRESS                                                PartySubIDType = "6"
	PartySubIDType_PHONE_NUMBER                                                  PartySubIDType = "7"
	PartySubIDType_EMAIL_ADDRESS                                                 PartySubIDType = "8"
	PartySubIDType_CONTACT_NAME                                                  PartySubIDType = "9"
)

//Enum values for PaymentMethod
type PaymentMethod string

const (
	PaymentMethod_CREST                      PaymentMethod = "1"
	PaymentMethod_DIRECT_CREDIT              PaymentMethod = "10"
	PaymentMethod_CREDIT_CARD                PaymentMethod = "11"
	PaymentMethod_ACH_DEBIT                  PaymentMethod = "12"
	PaymentMethod_ACH_CREDIT                 PaymentMethod = "13"
	PaymentMethod_BPAY                       PaymentMethod = "14"
	PaymentMethod_HIGH_VALUE_CLEARING_SYSTEM PaymentMethod = "15"
	PaymentMethod_NSCC                       PaymentMethod = "2"
	PaymentMethod_EUROCLEAR                  PaymentMethod = "3"
	PaymentMethod_CLEARSTREAM                PaymentMethod = "4"
	PaymentMethod_CHEQUE                     PaymentMethod = "5"
	PaymentMethod_TELEGRAPHIC_TRANSFER       PaymentMethod = "6"
	PaymentMethod_FED_WIRE                   PaymentMethod = "7"
	PaymentMethod_DEBIT_CARD                 PaymentMethod = "8"
	PaymentMethod_DIRECT_DEBIT               PaymentMethod = "9"
)

//Enum values for PegLimitType
type PegLimitType string

const (
	PegLimitType_OR_BETTER PegLimitType = "0"
	PegLimitType_STRICT    PegLimitType = "1"
	PegLimitType_OR_WORSE  PegLimitType = "2"
)

//Enum values for PegMoveType
type PegMoveType string

const (
	PegMoveType_FLOATING PegMoveType = "0"
	PegMoveType_FIXED    PegMoveType = "1"
)

//Enum values for PegOffsetType
type PegOffsetType string

const (
	PegOffsetType_PRICE        PegOffsetType = "0"
	PegOffsetType_BASIS_POINTS PegOffsetType = "1"
	PegOffsetType_TICKS        PegOffsetType = "2"
	PegOffsetType_PRICE_TIER   PegOffsetType = "3"
)

//Enum values for PegPriceType
type PegPriceType string

const (
	PegPriceType_LAST_PEG                                              PegPriceType = "1"
	PegPriceType_MID_PRICE_PEG                                         PegPriceType = "2"
	PegPriceType_OPENING_PEG                                           PegPriceType = "3"
	PegPriceType_MARKET_PEG                                            PegPriceType = "4"
	PegPriceType_PRIMARY_PEG                                           PegPriceType = "5"
	PegPriceType_FIXED_PEG_TO_LOCAL_BEST_BID_OR_OFFER_AT_TIME_OF_ORDER PegPriceType = "6"
	PegPriceType_PEG_TO_VWAP                                           PegPriceType = "7"
	PegPriceType_TRAILING_STOP_PEG                                     PegPriceType = "8"
	PegPriceType_PEG_TO_LIMIT_PRICE                                    PegPriceType = "9"
)

//Enum values for PegRoundDirection
type PegRoundDirection string

const (
	PegRoundDirection_MORE_AGGRESSIVE PegRoundDirection = "1"
	PegRoundDirection_MORE_PASSIVE    PegRoundDirection = "2"
)

//Enum values for PegScope
type PegScope string

const (
	PegScope_LOCAL                    PegScope = "1"
	PegScope_NATIONAL                 PegScope = "2"
	PegScope_GLOBAL                   PegScope = "3"
	PegScope_NATIONAL_EXCLUDING_LOCAL PegScope = "4"
)

//Enum values for PosAmtType
type PosAmtType string

const (
	PosAmtType_ACCRUED_COUPON_AMOUNT                     PosAmtType = "ACPN"
	PosAmtType_TOTAL_BANKED_AMOUNT                       PosAmtType = "BANK"
	PosAmtType_CASH_AMOUNT                               PosAmtType = "CASH"
	PosAmtType_COLLATERALIZED_MARK_TO_MARKET             PosAmtType = "CMTM"
	PosAmtType_TOTAL_COLLATERALIZED_AMOUNT               PosAmtType = "COLAT"
	PosAmtType_COUPON_AMOUNT                             PosAmtType = "CPN"
	PosAmtType_CASH_RESIDUAL_AMOUNT                      PosAmtType = "CRES"
	PosAmtType_COMPENSATION_AMOUNT                       PosAmtType = "DLV"
	PosAmtType_FINAL_MARK_TO_MARKET_AMOUNT               PosAmtType = "FMTM"
	PosAmtType_INCREMENTAL_ACCRUED_COUPON                PosAmtType = "IACPN"
	PosAmtType_INCREMENTAL_COLLATERALIZED_MARK_TO_MARKET PosAmtType = "ICMTM"
	PosAmtType_INITIAL_TRADE_COUPON_AMOUNT               PosAmtType = "ICPN"
	PosAmtType_INCREMENTAL_MARK_TO_MARKET_AMOUNT         PosAmtType = "IMTM"
	PosAmtType_PREMIUM_AMOUNT                            PosAmtType = "PREM"
	PosAmtType_SETTLEMENT_VALUE                          PosAmtType = "SETL"
	PosAmtType_START_OF_DAY_MARK_TO_MARKET_AMOUNT        PosAmtType = "SMTM"
	PosAmtType_TRADE_VARIATION_AMOUNT                    PosAmtType = "TVAR"
	PosAmtType_VALUE_ADJUSTED_AMOUNT                     PosAmtType = "VADJ"
)

//Enum values for PosMaintAction
type PosMaintAction string

const (
	PosMaintAction_NEW     PosMaintAction = "1"
	PosMaintAction_REPLACE PosMaintAction = "2"
	PosMaintAction_CANCEL  PosMaintAction = "3"
	PosMaintAction_REVERSE PosMaintAction = "4"
)

//Enum values for PosMaintResult
type PosMaintResult string

const (
	PosMaintResult_SUCCESSFUL_COMPLETION PosMaintResult = "0"
	PosMaintResult_REJECTED              PosMaintResult = "1"
	PosMaintResult_OTHER                 PosMaintResult = "99"
)

//Enum values for PosMaintStatus
type PosMaintStatus string

const (
	PosMaintStatus_ACCEPTED                PosMaintStatus = "0"
	PosMaintStatus_ACCEPTED_WITH_WARNINGS  PosMaintStatus = "1"
	PosMaintStatus_REJECTED                PosMaintStatus = "2"
	PosMaintStatus_COMPLETED               PosMaintStatus = "3"
	PosMaintStatus_COMPLETED_WITH_WARNINGS PosMaintStatus = "4"
)

//Enum values for PosQtyStatus
type PosQtyStatus string

const (
	PosQtyStatus_SUBMITTED PosQtyStatus = "0"
	PosQtyStatus_ACCEPTED  PosQtyStatus = "1"
	PosQtyStatus_REJECTED  PosQtyStatus = "2"
)

//Enum values for PosReqResult
type PosReqResult string

const (
	PosReqResult_VALID_REQUEST                          PosReqResult = "0"
	PosReqResult_INVALID_OR_UNSUPPORTED_REQUEST         PosReqResult = "1"
	PosReqResult_NO_POSITIONS_FOUND_THAT_MATCH_CRITERIA PosReqResult = "2"
	PosReqResult_NOT_AUTHORIZED_TO_REQUEST_POSITIONS    PosReqResult = "3"
	PosReqResult_REQUEST_FOR_POSITION_NOT_SUPPORTED     PosReqResult = "4"
	PosReqResult_OTHER                                  PosReqResult = "99"
)

//Enum values for PosReqStatus
type PosReqStatus string

const (
	PosReqStatus_COMPLETED               PosReqStatus = "0"
	PosReqStatus_COMPLETED_WITH_WARNINGS PosReqStatus = "1"
	PosReqStatus_REJECTED                PosReqStatus = "2"
)

//Enum values for PosReqType
type PosReqType string

const (
	PosReqType_POSITIONS           PosReqType = "0"
	PosReqType_TRADES              PosReqType = "1"
	PosReqType_EXERCISES           PosReqType = "2"
	PosReqType_ASSIGNMENTS         PosReqType = "3"
	PosReqType_SETTLEMENT_ACTIVITY PosReqType = "4"
	PosReqType_BACKOUT_MESSAGE     PosReqType = "5"
	PosReqType_DELTA_POSITIONS     PosReqType = "6"
)

//Enum values for PosTransType
type PosTransType string

const (
	PosTransType_EXERCISE                                      PosTransType = "1"
	PosTransType_DO_NOT_EXERCISE                               PosTransType = "2"
	PosTransType_POSITION_ADJUSTMENT                           PosTransType = "3"
	PosTransType_POSITION_CHANGE_SUBMISSION_MARGIN_DISPOSITION PosTransType = "4"
	PosTransType_PLEDGE                                        PosTransType = "5"
	PosTransType_LARGE_TRADER_SUBMISSION                       PosTransType = "6"
)

//Enum values for PosType
type PosType string

const (
	PosType_ALLOCATION_TRADE_QTY           PosType = "ALC"
	PosType_OPTION_ASSIGNMENT              PosType = "AS"
	PosType_AS_OF_TRADE_QTY                PosType = "ASF"
	PosType_CORPORATE_ACTION_ADJUSTMENT    PosType = "CAA"
	PosType_CREDIT_EVENT_ADJUSTMENT        PosType = "CEA"
	PosType_NET_DELTA_QTY                  PosType = "DLT"
	PosType_DELIVERY_QTY                   PosType = "DLV"
	PosType_DELIVERY_NOTICE_QTY            PosType = "DN"
	PosType_EXCHANGE_FOR_PHYSICAL_QTY      PosType = "EP"
	PosType_ELECTRONIC_TRADE_QTY           PosType = "ETR"
	PosType_OPTION_EXERCISE_QTY            PosType = "EX"
	PosType_END_OF_DAY_QTY                 PosType = "FIN"
	PosType_INTRA_SPREAD_QTY               PosType = "IAS"
	PosType_INTER_SPREAD_QTY               PosType = "IES"
	PosType_ADJUSTMENT_QTY                 PosType = "PA"
	PosType_PIT_TRADE_QTY                  PosType = "PIT"
	PosType_PRIVATELY_NEGOTIATED_TRADE_QTY PosType = "PNTN"
	PosType_RECEIVE_QUANTITY               PosType = "RCV"
	PosType_SUCCESSION_EVENT_ADJUSTMENT    PosType = "SEA"
	PosType_START_OF_DAY_QTY               PosType = "SOD"
	PosType_INTEGRAL_SPLIT                 PosType = "SPL"
	PosType_TRANSACTION_FROM_ASSIGNMENT    PosType = "TA"
	PosType_TOTAL_TRANSACTION_QTY          PosType = "TOT"
	PosType_TRANSACTION_QUANTITY           PosType = "TQ"
	PosType_TRANSFER_TRADE_QTY             PosType = "TRF"
	PosType_TRANSACTION_FROM_EXERCISE      PosType = "TX"
	PosType_CROSS_MARGIN_QTY               PosType = "XM"
)

//Enum values for PositionEffect
type PositionEffect string

const (
	PositionEffect_CLOSE                    PositionEffect = "C"
	PositionEffect_DEFAULT                  PositionEffect = "D"
	PositionEffect_FIFO                     PositionEffect = "F"
	PositionEffect_CLOSE_BUT_NOTIFY_ON_OPEN PositionEffect = "N"
	PositionEffect_OPEN                     PositionEffect = "O"
	PositionEffect_ROLLED                   PositionEffect = "R"
)

//Enum values for PossDupFlag
type PossDupFlag string

const (
	PossDupFlag_NO  PossDupFlag = "N"
	PossDupFlag_YES PossDupFlag = "Y"
)

//Enum values for PossResend
type PossResend string

const (
	PossResend_NO  PossResend = "N"
	PossResend_YES PossResend = "Y"
)

//Enum values for PreallocMethod
type PreallocMethod string

const (
	PreallocMethod_PRO_RATA        PreallocMethod = "0"
	PreallocMethod_DO_NOT_PRO_RATA PreallocMethod = "1"
)

//Enum values for PreviouslyReported
type PreviouslyReported string

const (
	PreviouslyReported_NO  PreviouslyReported = "N"
	PreviouslyReported_YES PreviouslyReported = "Y"
)

//Enum values for PriceLimitType
type PriceLimitType string

const (
	PriceLimitType_PRICE      PriceLimitType = "0"
	PriceLimitType_TICKS      PriceLimitType = "1"
	PriceLimitType_PERCENTAGE PriceLimitType = "2"
)

//Enum values for PriceProtectionScope
type PriceProtectionScope string

const (
	PriceProtectionScope_NONE     PriceProtectionScope = "0"
	PriceProtectionScope_LOCAL    PriceProtectionScope = "1"
	PriceProtectionScope_NATIONAL PriceProtectionScope = "2"
	PriceProtectionScope_GLOBAL   PriceProtectionScope = "3"
)

//Enum values for PriceQuoteMethod
type PriceQuoteMethod string

const (
	PriceQuoteMethod_INTEREST_RATE_INDEX PriceQuoteMethod = "INT"
	PriceQuoteMethod_INDEX               PriceQuoteMethod = "INX"
	PriceQuoteMethod_PERCENT_OF_PAR      PriceQuoteMethod = "PCTPAR"
	PriceQuoteMethod_STANDARD            PriceQuoteMethod = "STD"
)

//Enum values for PriceType
type PriceType string

const (
	PriceType_PERCENTAGE                         PriceType = "1"
	PriceType_FIXED_CABINET_TRADE_PRICE          PriceType = "10"
	PriceType_VARIABLE_CABINET_TRADE_PRICE       PriceType = "11"
	PriceType_PRODUCT_TICKS_IN_HALFS             PriceType = "13"
	PriceType_PRODUCT_TICKS_IN_FOURTHS           PriceType = "14"
	PriceType_PRODUCT_TICKS_IN_EIGHTS            PriceType = "15"
	PriceType_PRODUCT_TICKS_IN_SIXTEENTHS        PriceType = "16"
	PriceType_PRODUCT_TICKS_IN_THIRTY_SECONDS    PriceType = "17"
	PriceType_PRODUCT_TICKS_IN_SIXTY_FORTHS      PriceType = "18"
	PriceType_PRODUCT_TICKS_IN_ONE_TWENTY_EIGHTS PriceType = "19"
	PriceType_PER_UNIT                           PriceType = "2"
	PriceType_FIXED_AMOUNT                       PriceType = "3"
	PriceType_DISCOUNT                           PriceType = "4"
	PriceType_PREMIUM                            PriceType = "5"
	PriceType_SPREAD                             PriceType = "6"
	PriceType_TED_PRICE                          PriceType = "7"
	PriceType_TED_YIELD                          PriceType = "8"
	PriceType_YIELD                              PriceType = "9"
)

//Enum values for PriorityIndicator
type PriorityIndicator string

const (
	PriorityIndicator_PRIORITY_UNCHANGED                      PriorityIndicator = "0"
	PriorityIndicator_LOST_PRIORITY_AS_RESULT_OF_ORDER_CHANGE PriorityIndicator = "1"
)

//Enum values for ProcessCode
type ProcessCode string

const (
	ProcessCode_REGULAR              ProcessCode = "0"
	ProcessCode_SOFT_DOLLAR          ProcessCode = "1"
	ProcessCode_STEP_IN              ProcessCode = "2"
	ProcessCode_STEP_OUT             ProcessCode = "3"
	ProcessCode_SOFT_DOLLAR_STEP_IN  ProcessCode = "4"
	ProcessCode_SOFT_DOLLAR_STEP_OUT ProcessCode = "5"
	ProcessCode_PLAN_SPONSOR         ProcessCode = "6"
)

//Enum values for Product
type Product string

const (
	Product_AGENCY      Product = "1"
	Product_MORTGAGE    Product = "10"
	Product_MUNICIPAL   Product = "11"
	Product_OTHER       Product = "12"
	Product_FINANCING   Product = "13"
	Product_COMMODITY   Product = "2"
	Product_CORPORATE   Product = "3"
	Product_CURRENCY    Product = "4"
	Product_EQUITY      Product = "5"
	Product_GOVERNMENT  Product = "6"
	Product_INDEX       Product = "7"
	Product_LOAN        Product = "8"
	Product_MONEYMARKET Product = "9"
)

//Enum values for ProgRptReqs
type ProgRptReqs string

const (
	ProgRptReqs_BUY_SIDE_EXPLICITLY_REQUESTS_STATUS_USING_STATUE_REQUEST                                            ProgRptReqs = "1"
	ProgRptReqs_SELL_SIDE_PERIODICALLY_SENDS_STATUS_USING_LIST_STATUS_PERIOD_OPTIONALLY_SPECIFIED_IN_PROGRESSPERIOD ProgRptReqs = "2"
	ProgRptReqs_REAL_TIME_EXECUTION_REPORTS                                                                         ProgRptReqs = "3"
)

//Enum values for PublishTrdIndicator
type PublishTrdIndicator string

const (
	PublishTrdIndicator_NO  PublishTrdIndicator = "N"
	PublishTrdIndicator_YES PublishTrdIndicator = "Y"
)

//Enum values for PutOrCall
type PutOrCall string

const (
	PutOrCall_PUT  PutOrCall = "0"
	PutOrCall_CALL PutOrCall = "1"
)

//Enum values for QtyType
type QtyType string

const (
	QtyType_UNITS                          QtyType = "0"
	QtyType_CONTRACTS                      QtyType = "1"
	QtyType_UNITS_OF_MEASURE_PER_TIME_UNIT QtyType = "2"
)

//Enum values for QuantityType
type QuantityType string

const (
	QuantityType_SHARES       QuantityType = "1"
	QuantityType_BONDS        QuantityType = "2"
	QuantityType_CURRENTFACE  QuantityType = "3"
	QuantityType_ORIGINALFACE QuantityType = "4"
	QuantityType_CURRENCY     QuantityType = "5"
	QuantityType_CONTRACTS    QuantityType = "6"
	QuantityType_OTHER        QuantityType = "7"
	QuantityType_PAR          QuantityType = "8"
)

//Enum values for QuoteAckStatus
type QuoteAckStatus string

const (
	QuoteAckStatus_ACCEPTED                   QuoteAckStatus = "0"
	QuoteAckStatus_CANCELED_FOR_SYMBOL        QuoteAckStatus = "1"
	QuoteAckStatus_CANCELED_FOR_SECURITY_TYPE QuoteAckStatus = "2"
	QuoteAckStatus_CANCELED_FOR_UNDERLYING    QuoteAckStatus = "3"
	QuoteAckStatus_CANCELED_ALL               QuoteAckStatus = "4"
	QuoteAckStatus_REJECTED                   QuoteAckStatus = "5"
)

//Enum values for QuoteCancelType
type QuoteCancelType string

const (
	QuoteCancelType_CANCEL_FOR_ONE_OR_MORE_SECURITIES        QuoteCancelType = "1"
	QuoteCancelType_CANCEL_FOR_SECURITY_TYPE                 QuoteCancelType = "2"
	QuoteCancelType_CANCEL_FOR_UNDERLYING_SECURITY           QuoteCancelType = "3"
	QuoteCancelType_CANCEL_ALL_QUOTES                        QuoteCancelType = "4"
	QuoteCancelType_CANCEL_QUOTE_SPECIFIED_IN_QUOTEID        QuoteCancelType = "5"
	QuoteCancelType_CANCEL_BY_QUOTETYPE                      QuoteCancelType = "6"
	QuoteCancelType_CANCEL_FOR_SECURITY_ISSUER               QuoteCancelType = "7"
	QuoteCancelType_CANCEL_FOR_ISSUER_OF_UNDERLYING_SECURITY QuoteCancelType = "8"
)

//Enum values for QuoteCondition
type QuoteCondition string

const (
	QuoteCondition_RESERVED_SAM                        QuoteCondition = "0"
	QuoteCondition_NO_ACTIVE_SAM                       QuoteCondition = "1"
	QuoteCondition_RESTRICTED                          QuoteCondition = "2"
	QuoteCondition_REST_OF_BOOK_VWAP                   QuoteCondition = "3"
	QuoteCondition_BETTER_PRICES_IN_CONDITIONAL_ORDERS QuoteCondition = "4"
	QuoteCondition_MEDIAN_PRICE                        QuoteCondition = "5"
	QuoteCondition_FULL_CURVE                          QuoteCondition = "6"
	QuoteCondition_FLAT_CURVE                          QuoteCondition = "7"
	QuoteCondition_OPEN_ACTIVE                         QuoteCondition = "A"
	QuoteCondition_CLOSED_INACTIVE                     QuoteCondition = "B"
	QuoteCondition_EXCHANGE_BEST                       QuoteCondition = "C"
	QuoteCondition_CONSOLIDATED_BEST                   QuoteCondition = "D"
	QuoteCondition_LOCKED                              QuoteCondition = "E"
	QuoteCondition_CROSSED                             QuoteCondition = "F"
	QuoteCondition_DEPTH                               QuoteCondition = "G"
	QuoteCondition_FAST_TRADING                        QuoteCondition = "H"
	QuoteCondition_NON_FIRM                            QuoteCondition = "I"
	QuoteCondition_OUTRIGHT_PRICE                      QuoteCondition = "J"
	QuoteCondition_IMPLIED_PRICE                       QuoteCondition = "K"
	QuoteCondition_MANUAL_SLOW_QUOTE                   QuoteCondition = "L"
	QuoteCondition_DEPTH_ON_OFFER                      QuoteCondition = "M"
	QuoteCondition_DEPTH_ON_BID                        QuoteCondition = "N"
	QuoteCondition_CLOSING                             QuoteCondition = "O"
	QuoteCondition_NEWS_DISSEMINATION                  QuoteCondition = "P"
	QuoteCondition_TRADING_RANGE                       QuoteCondition = "Q"
	QuoteCondition_ORDER_INFLUX                        QuoteCondition = "R"
	QuoteCondition_DUE_TO_RELATED                      QuoteCondition = "S"
	QuoteCondition_NEWS_PENDING                        QuoteCondition = "T"
	QuoteCondition_ADDITIONAL_INFO                     QuoteCondition = "U"
	QuoteCondition_ADDITIONAL_INFO_DUE_TO_RELATED      QuoteCondition = "V"
	QuoteCondition_RESUME                              QuoteCondition = "W"
	QuoteCondition_VIEW_OF_COMMON                      QuoteCondition = "X"
	QuoteCondition_VOLUME_ALERT                        QuoteCondition = "Y"
	QuoteCondition_ORDER_IMBALANCE                     QuoteCondition = "Z"
	QuoteCondition_EQUIPMENT_CHANGEOVER                QuoteCondition = "a"
	QuoteCondition_NO_OPEN                             QuoteCondition = "b"
	QuoteCondition_REGULAR_ETH                         QuoteCondition = "c"
	QuoteCondition_AUTOMATIC_EXECUTION                 QuoteCondition = "d"
	QuoteCondition_AUTOMATIC_EXECUTION_ETH             QuoteCondition = "e"
	QuoteCondition_FAST_MARKET_ETH                     QuoteCondition = "f "
	QuoteCondition_INACTIVE_ETH                        QuoteCondition = "g"
	QuoteCondition_ROTATION                            QuoteCondition = "h"
	QuoteCondition_ROTATION_ETH                        QuoteCondition = "i"
	QuoteCondition_HALT                                QuoteCondition = "j"
	QuoteCondition_HALT_ETH                            QuoteCondition = "k"
	QuoteCondition_DUE_TO_NEWS_DISSEMINATION           QuoteCondition = "l"
	QuoteCondition_DUE_TO_NEWS_PENDING                 QuoteCondition = "m"
	QuoteCondition_TRADING_RESUME                      QuoteCondition = "n"
	QuoteCondition_OUT_OF_SEQUENCE                     QuoteCondition = "o"
	QuoteCondition_BID_SPECIALIST                      QuoteCondition = "p"
	QuoteCondition_OFFER_SPECIALIST                    QuoteCondition = "q"
	QuoteCondition_BID_OFFER_SPECIALIST                QuoteCondition = "r"
	QuoteCondition_END_OF_DAY_SAM                      QuoteCondition = "s"
	QuoteCondition_FORBIDDEN_SAM                       QuoteCondition = "t"
	QuoteCondition_FROZEN_SAM                          QuoteCondition = "u"
	QuoteCondition_PREOPENING_SAM                      QuoteCondition = "v"
	QuoteCondition_OPENING_SAM                         QuoteCondition = "w"
	QuoteCondition_OPEN_SAM                            QuoteCondition = "x"
	QuoteCondition_SURVEILLANCE_SAM                    QuoteCondition = "y"
	QuoteCondition_SUSPENDED_SAM                       QuoteCondition = "z"
)

//Enum values for QuoteEntryRejectReason
type QuoteEntryRejectReason string

const (
	QuoteEntryRejectReason_UNKNOWN_SYMBOL                   QuoteEntryRejectReason = "1"
	QuoteEntryRejectReason_EXHCNAGE                         QuoteEntryRejectReason = "2"
	QuoteEntryRejectReason_QUOTE_EXCEEDS_LIMIT              QuoteEntryRejectReason = "3"
	QuoteEntryRejectReason_TOO_LATE_TO_ENTER                QuoteEntryRejectReason = "4"
	QuoteEntryRejectReason_UNKNOWN_QUOTE                    QuoteEntryRejectReason = "5"
	QuoteEntryRejectReason_DUPLICATE_QUOTE                  QuoteEntryRejectReason = "6"
	QuoteEntryRejectReason_INVALID_BID_ASK_SPREAD           QuoteEntryRejectReason = "7"
	QuoteEntryRejectReason_INVALID_PRICE                    QuoteEntryRejectReason = "8"
	QuoteEntryRejectReason_NOT_AUTHORIZED_TO_QUOTE_SECURITY QuoteEntryRejectReason = "9"
	QuoteEntryRejectReason_OTHER                            QuoteEntryRejectReason = "99"
)

//Enum values for QuoteEntryStatus
type QuoteEntryStatus string

const (
	QuoteEntryStatus_ACCEPTED                     QuoteEntryStatus = "0"
	QuoteEntryStatus_LOCKED_MARKET_WARNING        QuoteEntryStatus = "12"
	QuoteEntryStatus_CROSS_MARKET_WARNING         QuoteEntryStatus = "13"
	QuoteEntryStatus_CANCELED_DUE_TO_LOCK_MARKET  QuoteEntryStatus = "14"
	QuoteEntryStatus_CANCELED_DUE_TO_CROSS_MARKET QuoteEntryStatus = "15"
	QuoteEntryStatus_ACTIVE                       QuoteEntryStatus = "16"
	QuoteEntryStatus_REJECTED                     QuoteEntryStatus = "5"
	QuoteEntryStatus_REMOVED_FROM_MARKET          QuoteEntryStatus = "6"
	QuoteEntryStatus_EXPIRED                      QuoteEntryStatus = "7"
)

//Enum values for QuotePriceType
type QuotePriceType string

const (
	QuotePriceType_PERCENT      QuotePriceType = "1"
	QuotePriceType_YIELD        QuotePriceType = "10"
	QuotePriceType_PER_SHARE    QuotePriceType = "2"
	QuotePriceType_FIXED_AMOUNT QuotePriceType = "3"
	QuotePriceType_DISCOUNT     QuotePriceType = "4"
	QuotePriceType_PREMIUM      QuotePriceType = "5"
	QuotePriceType_SPREAD       QuotePriceType = "6"
	QuotePriceType_TED_PRICE    QuotePriceType = "7"
	QuotePriceType_TED_YIELD    QuotePriceType = "8"
	QuotePriceType_YIELD_SPREAD QuotePriceType = "9"
)

//Enum values for QuoteRejectReason
type QuoteRejectReason string

const (
	QuoteRejectReason_UNKNOWN_SYMBOL                                   QuoteRejectReason = "1"
	QuoteRejectReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND                 QuoteRejectReason = "10"
	QuoteRejectReason_QUOTE_LOCKED                                     QuoteRejectReason = "11"
	QuoteRejectReason_INVALID_OR_UNKNOWN_SECURITY_ISSUER               QuoteRejectReason = "12"
	QuoteRejectReason_INVALID_OR_UNKNOWN_ISSUER_OF_UNDERLYING_SECURITY QuoteRejectReason = "13"
	QuoteRejectReason_EXCHANGE                                         QuoteRejectReason = "2"
	QuoteRejectReason_QUOTE_REQUEST_EXCEEDS_LIMIT                      QuoteRejectReason = "3"
	QuoteRejectReason_TOO_LATE_TO_ENTER                                QuoteRejectReason = "4"
	QuoteRejectReason_UNKNOWN_QUOTE                                    QuoteRejectReason = "5"
	QuoteRejectReason_DUPLICATE_QUOTE                                  QuoteRejectReason = "6"
	QuoteRejectReason_INVALID_BID_ASK_SPREAD                           QuoteRejectReason = "7"
	QuoteRejectReason_INVALID_PRICE                                    QuoteRejectReason = "8"
	QuoteRejectReason_NOT_AUTHORIZED_TO_QUOTE_SECURITY                 QuoteRejectReason = "9"
	QuoteRejectReason_OTHER                                            QuoteRejectReason = "99"
)

//Enum values for QuoteRequestRejectReason
type QuoteRequestRejectReason string

const (
	QuoteRequestRejectReason_UNKNOWN_SYMBOL                  QuoteRequestRejectReason = "1"
	QuoteRequestRejectReason_PASS                            QuoteRequestRejectReason = "10"
	QuoteRequestRejectReason_INSUFFICIENT_CREDIT             QuoteRequestRejectReason = "11"
	QuoteRequestRejectReason_EXCHANGE                        QuoteRequestRejectReason = "2"
	QuoteRequestRejectReason_QUOTE_REQUEST_EXCEEDS_LIMIT     QuoteRequestRejectReason = "3"
	QuoteRequestRejectReason_TOO_LATE_TO_ENTER               QuoteRequestRejectReason = "4"
	QuoteRequestRejectReason_INVALID_PRICE                   QuoteRequestRejectReason = "5"
	QuoteRequestRejectReason_NOT_AUTHORIZED_TO_REQUEST_QUOTE QuoteRequestRejectReason = "6"
	QuoteRequestRejectReason_NO_MATCH_FOR_INQUIRY            QuoteRequestRejectReason = "7"
	QuoteRequestRejectReason_NO_MARKET_FOR_INSTRUMENT        QuoteRequestRejectReason = "8"
	QuoteRequestRejectReason_NO_INVENTORY                    QuoteRequestRejectReason = "9"
	QuoteRequestRejectReason_OTHER                           QuoteRequestRejectReason = "99"
)

//Enum values for QuoteRequestType
type QuoteRequestType string

const (
	QuoteRequestType_MANUAL    QuoteRequestType = "1"
	QuoteRequestType_AUTOMATIC QuoteRequestType = "2"
)

//Enum values for QuoteRespType
type QuoteRespType string

const (
	QuoteRespType_HIT_LIFT  QuoteRespType = "1"
	QuoteRespType_COUNTER   QuoteRespType = "2"
	QuoteRespType_EXPIRED   QuoteRespType = "3"
	QuoteRespType_COVER     QuoteRespType = "4"
	QuoteRespType_DONE_AWAY QuoteRespType = "5"
	QuoteRespType_PASS      QuoteRespType = "6"
	QuoteRespType_END_TRADE QuoteRespType = "7"
	QuoteRespType_TIMED_OUT QuoteRespType = "8"
)

//Enum values for QuoteResponseLevel
type QuoteResponseLevel string

const (
	QuoteResponseLevel_NO_ACKNOWLEDGEMENT                            QuoteResponseLevel = "0"
	QuoteResponseLevel_ACKNOWLEDGE_ONLY_NEGATIVE_OR_ERRONEOUS_QUOTES QuoteResponseLevel = "1"
	QuoteResponseLevel_ACKNOWLEDGE_EACH_QUOTE_MESSAGE                QuoteResponseLevel = "2"
	QuoteResponseLevel_SUMMARY_ACKNOWLEDGEMENT                       QuoteResponseLevel = "3"
)

//Enum values for QuoteStatus
type QuoteStatus string

const (
	QuoteStatus_ACCEPTED                        QuoteStatus = "0"
	QuoteStatus_CANCEL_FOR_SYMBOL               QuoteStatus = "1"
	QuoteStatus_PENDING                         QuoteStatus = "10"
	QuoteStatus_PASS                            QuoteStatus = "11"
	QuoteStatus_LOCKED_MARKET_WARNING           QuoteStatus = "12"
	QuoteStatus_CROSS_MARKET_WARNING            QuoteStatus = "13"
	QuoteStatus_CANCELED_DUE_TO_LOCK_MARKET     QuoteStatus = "14"
	QuoteStatus_CANCELED_DUE_TO_CROSS_MARKET    QuoteStatus = "15"
	QuoteStatus_ACTIVE                          QuoteStatus = "16"
	QuoteStatus_CANCELED                        QuoteStatus = "17"
	QuoteStatus_UNSOLICITED_QUOTE_REPLENISHMENT QuoteStatus = "18"
	QuoteStatus_PENDING_END_TRADE               QuoteStatus = "19"
	QuoteStatus_CANCELED_FOR_SECURITY_TYPE      QuoteStatus = "2"
	QuoteStatus_TOO_LATE_TO_END                 QuoteStatus = "20"
	QuoteStatus_CANCELED_FOR_UNDERLYING         QuoteStatus = "3"
	QuoteStatus_CANCELED_ALL                    QuoteStatus = "4"
	QuoteStatus_REJECTED                        QuoteStatus = "5"
	QuoteStatus_REMOVED_FROM_MARKET             QuoteStatus = "6"
	QuoteStatus_EXPIRED                         QuoteStatus = "7"
	QuoteStatus_QUERY                           QuoteStatus = "8"
	QuoteStatus_QUOTE_NOT_FOUND                 QuoteStatus = "9"
)

//Enum values for QuoteType
type QuoteType string

const (
	QuoteType_INDICATIVE           QuoteType = "0"
	QuoteType_TRADEABLE            QuoteType = "1"
	QuoteType_RESTRICTED_TRADEABLE QuoteType = "2"
	QuoteType_COUNTER              QuoteType = "3"
)

//Enum values for RateSource
type RateSource string

const (
	RateSource_BLOOMBERG RateSource = "0"
	RateSource_REUTERS   RateSource = "1"
	RateSource_TELERATE  RateSource = "2"
	RateSource_OTHER     RateSource = "99"
)

//Enum values for RateSourceType
type RateSourceType string

const (
	RateSourceType_PRIMARY   RateSourceType = "0"
	RateSourceType_SECONDARY RateSourceType = "1"
)

//Enum values for RefOrdIDReason
type RefOrdIDReason string

const (
	RefOrdIDReason_GTC_FROM_PREVIOUS_DAY  RefOrdIDReason = "0"
	RefOrdIDReason_PARTIAL_FILL_REMAINING RefOrdIDReason = "1"
	RefOrdIDReason_ORDER_CHANGED          RefOrdIDReason = "2"
)

//Enum values for RefOrderIDSource
type RefOrderIDSource string

const (
	RefOrderIDSource_SECONDARYORDERID  RefOrderIDSource = "0"
	RefOrderIDSource_ORDERID           RefOrderIDSource = "1"
	RefOrderIDSource_MDENTRYID         RefOrderIDSource = "2"
	RefOrderIDSource_QUOTEENTRYID      RefOrderIDSource = "3"
	RefOrderIDSource_ORIGINAL_ORDER_ID RefOrderIDSource = "4"
)

//Enum values for RegistRejReasonCode
type RegistRejReasonCode string

const (
	RegistRejReasonCode_INVALID_UNACCEPTABLE_ACCOUNT_TYPE                  RegistRejReasonCode = "1"
	RegistRejReasonCode_INVALID_UNACEEPTABLE_INVESTOR_ID_SOURCE            RegistRejReasonCode = "10"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_DATE_OF_BIRTH                 RegistRejReasonCode = "11"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_INVESTOR_COUNTRY_OF_RESIDENCE RegistRejReasonCode = "12"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_NO_DISTRIB_INSTNS             RegistRejReasonCode = "13"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_DISTRIB_PERCENTAGE            RegistRejReasonCode = "14"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_DISTRIB_PAYMENT_METHOD        RegistRejReasonCode = "15"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_CASH_DISTRIB_AGENT_ACCT_NAME  RegistRejReasonCode = "16"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_CASH_DISTRIB_AGENT_CODE       RegistRejReasonCode = "17"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_CASH_DISTRIB_AGENT_ACCT_NUM   RegistRejReasonCode = "18"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_TAX_EXEMPT_TYPE               RegistRejReasonCode = "2"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_OWNERSHIP_TYPE                RegistRejReasonCode = "3"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_NO_REG_DETAILS                RegistRejReasonCode = "4"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_REG_SEQ_NO                    RegistRejReasonCode = "5"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_REG_DETAILS                   RegistRejReasonCode = "6"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_MAILING_DETAILS               RegistRejReasonCode = "7"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_MAILING_INSTRUCTIONS          RegistRejReasonCode = "8"
	RegistRejReasonCode_INVALID_UNACCEPTABLE_INVESTOR_ID                   RegistRejReasonCode = "9"
	RegistRejReasonCode_OTHER                                              RegistRejReasonCode = "99"
)

//Enum values for RegistStatus
type RegistStatus string

const (
	RegistStatus_ACCEPTED RegistStatus = "A"
	RegistStatus_HELD     RegistStatus = "H"
	RegistStatus_REMINDER RegistStatus = "N"
	RegistStatus_REJECTED RegistStatus = "R"
)

//Enum values for RegistTransType
type RegistTransType string

const (
	RegistTransType_NEW     RegistTransType = "0"
	RegistTransType_REPLACE RegistTransType = "1"
	RegistTransType_CANCEL  RegistTransType = "2"
)

//Enum values for ReportToExch
type ReportToExch string

const (
	ReportToExch_NO  ReportToExch = "N"
	ReportToExch_YES ReportToExch = "Y"
)

//Enum values for ResetSeqNumFlag
type ResetSeqNumFlag string

const (
	ResetSeqNumFlag_NO  ResetSeqNumFlag = "N"
	ResetSeqNumFlag_YES ResetSeqNumFlag = "Y"
)

//Enum values for RespondentType
type RespondentType string

const (
	RespondentType_ALL_MARKET_PARTICIPANTS       RespondentType = "1"
	RespondentType_SPECIFIED_MARKET_PARTICIPANTS RespondentType = "2"
	RespondentType_ALL_MARKET_MAKERS             RespondentType = "3"
	RespondentType_PRIMARY_MARKET_MAKER          RespondentType = "4"
)

//Enum values for ResponseTransportType
type ResponseTransportType string

const (
	ResponseTransportType_INBAND      ResponseTransportType = "0"
	ResponseTransportType_OUT_OF_BAND ResponseTransportType = "1"
)

//Enum values for RestructuringType
type RestructuringType string

const (
	RestructuringType_FULL_RESTRUCTURING         RestructuringType = "FR"
	RestructuringType_MODIFIED_MOD_RESTRUCTURING RestructuringType = "MM"
	RestructuringType_MODIFIED_RESTRUCTURING     RestructuringType = "MR"
	RestructuringType_NO_RESTRUCTURING_SPECIFIED RestructuringType = "XR"
)

//Enum values for RiskInstrumentOperator
type RiskInstrumentOperator string

const (
	RiskInstrumentOperator_INCLUDE RiskInstrumentOperator = "1"
	RiskInstrumentOperator_EXCLUDE RiskInstrumentOperator = "2"
)

//Enum values for RiskLimitType
type RiskLimitType string

const (
	RiskLimitType_GROSS_LIMIT RiskLimitType = "1"
	RiskLimitType_NET_LIMIT   RiskLimitType = "2"
	RiskLimitType_EXPOSURE    RiskLimitType = "3"
	RiskLimitType_LONG_LIMIT  RiskLimitType = "4"
	RiskLimitType_SHORT_LIMIT RiskLimitType = "5"
)

//Enum values for RoundingDirection
type RoundingDirection string

const (
	RoundingDirection_ROUND_TO_NEAREST RoundingDirection = "0"
	RoundingDirection_ROUND_DOWN       RoundingDirection = "1"
	RoundingDirection_ROUND_UP         RoundingDirection = "2"
)

//Enum values for RoutingType
type RoutingType string

const (
	RoutingType_TARGET_FIRM RoutingType = "1"
	RoutingType_TARGET_LIST RoutingType = "2"
	RoutingType_BLOCK_FIRM  RoutingType = "3"
	RoutingType_BLOCK_LIST  RoutingType = "4"
)

//Enum values for Rule80A
type Rule80A string

const (
	Rule80A_AGENCY_SINGLE_ORDER                                                                                        Rule80A = "A"
	Rule80A_SHORT_EXEMPT_TRANSACTION_B                                                                                 Rule80A = "B"
	Rule80A_PROPRIETARY_NON_ALGORITHMIC_PROGRAM_TRADE                                                                  Rule80A = "C"
	Rule80A_PROGRAM_ORDER_INDEX_ARB_FOR_MEMBER_FIRM_ORG                                                                Rule80A = "D"
	Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_PRINCIPAL                                                                     Rule80A = "E"
	Rule80A_SHORT_EXEMPT_TRANSACTION_F                                                                                 Rule80A = "F"
	Rule80A_SHORT_EXEMPT_TRANSACTION_H                                                                                 Rule80A = "H"
	Rule80A_INDIVIDUAL_INVESTOR_SINGLE_ORDER                                                                           Rule80A = "I"
	Rule80A_PROPRIETARY_ALGORITHMIC_PROGRAM_TRADING                                                                    Rule80A = "J"
	Rule80A_AGENCY_ALGORITHMIC_PROGRAM_TRADING                                                                         Rule80A = "K"
	Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_MEMBER_COMPETING_MARKET_MAKER_AFFLIATED_WITH_THE_FIRM_CLEARING_THE_TRADE      Rule80A = "L"
	Rule80A_PROGRAM_ORDER_INDEX_ARB_FOR_OTHER_MEMBER                                                                   Rule80A = "M"
	Rule80A_AGENT_FOR_OTHER_MEMBER_NON_ALGORITHMIC_PROGRAM_TRADE                                                       Rule80A = "N"
	Rule80A_PROPRIETARY_TRANSACTIONS_FOR_COMPETING_MARKET_MAKER_THAT_IS_AFFILIATED_WITH_THE_CLEARING_MEMBER            Rule80A = "O"
	Rule80A_PRINCIPAL                                                                                                  Rule80A = "P"
	Rule80A_TRANSACTIONS_FOR_THE_ACCOUNT_OF_A_NON_MEMBER_COMPTING_MARKET_MAKER                                         Rule80A = "R"
	Rule80A_SPECIALIST_TRADES                                                                                          Rule80A = "S"
	Rule80A_TRANSACTIONS_FOR_THE_ACCOUNT_OF_AN_UNAFFILIATED_MEMBERS_COMPETING_MARKET_MAKER                             Rule80A = "T"
	Rule80A_AGENCY_INDEX_ARBITRAGE                                                                                     Rule80A = "U"
	Rule80A_ALL_OTHER_ORDERS_AS_AGENT_FOR_OTHER_MEMBER                                                                 Rule80A = "W"
	Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_MEMBER_COMPETING_MARKET_MAKER_NOT_AFFILIATED_WITH_THE_FIRM_CLEARING_THE_TRADE Rule80A = "X"
	Rule80A_AGENCY_NON_ALGORITHMIC_PROGRAM_TRADE                                                                       Rule80A = "Y"
	Rule80A_SHORT_EXEMPT_TRANSACTION_FOR_NON_MEMBER_COMPETING_MARKET_MAKER                                             Rule80A = "Z"
)

//Enum values for Scope
type Scope string

const (
	Scope_LOCAL_MARKET Scope = "1"
	Scope_NATIONAL     Scope = "2"
	Scope_GLOBAL       Scope = "3"
)

//Enum values for SecDefStatus
type SecDefStatus string

const (
	SecDefStatus_PENDING_APPROVAL           SecDefStatus = "0"
	SecDefStatus_APPROVED                   SecDefStatus = "1"
	SecDefStatus_REJECTED                   SecDefStatus = "2"
	SecDefStatus_UNAUTHORIZED_REQUEST       SecDefStatus = "3"
	SecDefStatus_INVALID_DEFINITION_REQUEST SecDefStatus = "4"
)

//Enum values for SecurityIDSource
type SecurityIDSource string

const (
	SecurityIDSource_CUSIP                            SecurityIDSource = "1"
	SecurityIDSource_SEDOL                            SecurityIDSource = "2"
	SecurityIDSource_QUIK                             SecurityIDSource = "3"
	SecurityIDSource_ISIN_NUMBER                      SecurityIDSource = "4"
	SecurityIDSource_RIC_CODE                         SecurityIDSource = "5"
	SecurityIDSource_ISO_CURRENCY_CODE                SecurityIDSource = "6"
	SecurityIDSource_ISO_COUNTRY_CODE                 SecurityIDSource = "7"
	SecurityIDSource_EXCHANGE_SYMBOL                  SecurityIDSource = "8"
	SecurityIDSource_CONSOLIDATED_TAPE_ASSOCIATION    SecurityIDSource = "9"
	SecurityIDSource_BLOOMBERG_SYMBOL                 SecurityIDSource = "A"
	SecurityIDSource_WERTPAPIER                       SecurityIDSource = "B"
	SecurityIDSource_DUTCH                            SecurityIDSource = "C"
	SecurityIDSource_VALOREN                          SecurityIDSource = "D"
	SecurityIDSource_SICOVAM                          SecurityIDSource = "E"
	SecurityIDSource_BELGIAN                          SecurityIDSource = "F"
	SecurityIDSource_COMMON                           SecurityIDSource = "G"
	SecurityIDSource_CLEARING_HOUSE                   SecurityIDSource = "H"
	SecurityIDSource_ISDA_FPML_PRODUCT_SPECIFICATION  SecurityIDSource = "I"
	SecurityIDSource_OPTION_PRICE_REPORTING_AUTHORITY SecurityIDSource = "J"
	SecurityIDSource_ISDA_FPML_PRODUCT_URL            SecurityIDSource = "K"
	SecurityIDSource_LETTER_OF_CREDIT                 SecurityIDSource = "L"
	SecurityIDSource_MARKETPLACE_ASSIGNED_IDENTIFIER  SecurityIDSource = "M"
)

//Enum values for SecurityListRequestType
type SecurityListRequestType string

const (
	SecurityListRequestType_SYMBOL                                    SecurityListRequestType = "0"
	SecurityListRequestType_SECURITYTYPE_AND_OR_CFICODE               SecurityListRequestType = "1"
	SecurityListRequestType_PRODUCT                                   SecurityListRequestType = "2"
	SecurityListRequestType_TRADINGSESSIONID                          SecurityListRequestType = "3"
	SecurityListRequestType_ALL_SECURITIES                            SecurityListRequestType = "4"
	SecurityListRequestType_MARKETID_OR_MARKETID_PLUS_MARKETSEGMENTID SecurityListRequestType = "5"
)

//Enum values for SecurityListType
type SecurityListType string

const (
	SecurityListType_INDUSTRY_CLASSIFICATION SecurityListType = "1"
	SecurityListType_TRADING_LIST            SecurityListType = "2"
	SecurityListType_MARKET                  SecurityListType = "3"
	SecurityListType_NEWSPAPER_LIST          SecurityListType = "4"
)

//Enum values for SecurityListTypeSource
type SecurityListTypeSource string

const (
	SecurityListTypeSource_ICB   SecurityListTypeSource = "1"
	SecurityListTypeSource_NAICS SecurityListTypeSource = "2"
	SecurityListTypeSource_GICS  SecurityListTypeSource = "3"
)

//Enum values for SecurityRequestResult
type SecurityRequestResult string

const (
	SecurityRequestResult_VALID_REQUEST                                      SecurityRequestResult = "0"
	SecurityRequestResult_INVALID_OR_UNSUPPORTED_REQUEST                     SecurityRequestResult = "1"
	SecurityRequestResult_NO_INSTRUMENTS_FOUND_THAT_MATCH_SELECTION_CRITERIA SecurityRequestResult = "2"
	SecurityRequestResult_NOT_AUTHORIZED_TO_RETRIEVE_INSTRUMENT_DATA         SecurityRequestResult = "3"
	SecurityRequestResult_INSTRUMENT_DATA_TEMPORARILY_UNAVAILABLE            SecurityRequestResult = "4"
	SecurityRequestResult_REQUEST_FOR_INSTRUMENT_DATA_NOT_SUPPORTED          SecurityRequestResult = "5"
)

//Enum values for SecurityRequestType
type SecurityRequestType string

const (
	SecurityRequestType_REQUEST_SECURITY_IDENTITY_AND_SPECIFICATIONS              SecurityRequestType = "0"
	SecurityRequestType_REQUEST_SECURITY_IDENTITY_FOR_THE_SPECIFICATIONS_PROVIDED SecurityRequestType = "1"
	SecurityRequestType_REQUEST_LIST_SECURITY_TYPES                               SecurityRequestType = "2"
	SecurityRequestType_REQUEST_LIST_SECURITIES                                   SecurityRequestType = "3"
	SecurityRequestType_SYMBOL                                                    SecurityRequestType = "4"
	SecurityRequestType_SECURITYTYPE_AND_OR_CFICODE                               SecurityRequestType = "5"
	SecurityRequestType_PRODUCT                                                   SecurityRequestType = "6"
	SecurityRequestType_TRADINGSESSIONID                                          SecurityRequestType = "7"
	SecurityRequestType_ALL_SECURITIES                                            SecurityRequestType = "8"
	SecurityRequestType_MARKETID_OR_MARKETID_PLUS_MARKETSEGMENTID                 SecurityRequestType = "9"
)

//Enum values for SecurityResponseType
type SecurityResponseType string

const (
	SecurityResponseType_ACCEPT_SECURITY_PROPOSAL_AS_IS                                      SecurityResponseType = "1"
	SecurityResponseType_ACCEPT_SECURITY_PROPOSAL_WITH_REVISIONS_AS_INDICATED_IN_THE_MESSAGE SecurityResponseType = "2"
	SecurityResponseType_LIST_OF_SECURITY_TYPES_RETURNED_PER_REQUEST                         SecurityResponseType = "3"
	SecurityResponseType_LIST_OF_SECURITIES_RETURNED_PER_REQUEST                             SecurityResponseType = "4"
	SecurityResponseType_REJECT_SECURITY_PROPOSAL                                            SecurityResponseType = "5"
	SecurityResponseType_CANNOT_MATCH_SELECTION_CRITERIA                                     SecurityResponseType = "6"
)

//Enum values for SecurityStatus
type SecurityStatus string

const (
	SecurityStatus_ACTIVE   SecurityStatus = "1"
	SecurityStatus_INACTIVE SecurityStatus = "2"
)

//Enum values for SecurityTradingEvent
type SecurityTradingEvent string

const (
	SecurityTradingEvent_ORDER_IMBALANCE_AUCTION_IS_EXTENDED SecurityTradingEvent = "1"
	SecurityTradingEvent_TRADING_RESUMES                     SecurityTradingEvent = "2"
	SecurityTradingEvent_PRICE_VOLATILITY_INTERRUPTION       SecurityTradingEvent = "3"
	SecurityTradingEvent_CHANGE_OF_TRADING_SESSION           SecurityTradingEvent = "4"
	SecurityTradingEvent_CHANGE_OF_TRADING_SUBSESSION        SecurityTradingEvent = "5"
	SecurityTradingEvent_CHANGE_OF_SECURITY_TRADING_STATUS   SecurityTradingEvent = "6"
	SecurityTradingEvent_CHANGE_OF_BOOK_TYPE                 SecurityTradingEvent = "7"
	SecurityTradingEvent_CHANGE_OF_MARKET_DEPTH              SecurityTradingEvent = "8"
)

//Enum values for SecurityTradingStatus
type SecurityTradingStatus string

const (
	SecurityTradingStatus_OPENING_DELAY                  SecurityTradingStatus = "1"
	SecurityTradingStatus_MARKET_ON_CLOSE_IMBALANCE_SELL SecurityTradingStatus = "10"
	SecurityTradingStatus_11                             SecurityTradingStatus = "11"
	SecurityTradingStatus_NO_MARKET_IMBALANCE            SecurityTradingStatus = "12"
	SecurityTradingStatus_NO_MARKET_ON_CLOSE_IMBALANCE   SecurityTradingStatus = "13"
	SecurityTradingStatus_ITS_PRE_OPENING                SecurityTradingStatus = "14"
	SecurityTradingStatus_NEW_PRICE_INDICATION           SecurityTradingStatus = "15"
	SecurityTradingStatus_TRADE_DISSEMINATION_TIME       SecurityTradingStatus = "16"
	SecurityTradingStatus_READY_TO_TRADE                 SecurityTradingStatus = "17"
	SecurityTradingStatus_NOT_AVAILABLE_FOR_TRADING      SecurityTradingStatus = "18"
	SecurityTradingStatus_NOT_TRADED_ON_THIS_MARKET      SecurityTradingStatus = "19"
	SecurityTradingStatus_TRADING_HALT                   SecurityTradingStatus = "2"
	SecurityTradingStatus_UNKNOWN_OR_INVALID             SecurityTradingStatus = "20"
	SecurityTradingStatus_PRE_OPEN                       SecurityTradingStatus = "21"
	SecurityTradingStatus_OPENING_ROTATION               SecurityTradingStatus = "22"
	SecurityTradingStatus_FAST_MARKET                    SecurityTradingStatus = "23"
	SecurityTradingStatus_PRE_CROSS                      SecurityTradingStatus = "24"
	SecurityTradingStatus_CROSS                          SecurityTradingStatus = "25"
	SecurityTradingStatus_POST_CLOSE                     SecurityTradingStatus = "26"
	SecurityTradingStatus_RESUME                         SecurityTradingStatus = "3"
	SecurityTradingStatus_NO_OPEN                        SecurityTradingStatus = "4"
	SecurityTradingStatus_PRICE_INDICATION               SecurityTradingStatus = "5"
	SecurityTradingStatus_TRADING_RANGE_INDICATION       SecurityTradingStatus = "6"
	SecurityTradingStatus_MARKET_IMBALANCE_BUY           SecurityTradingStatus = "7"
	SecurityTradingStatus_MARKET_IMBALANCE_SELL          SecurityTradingStatus = "8"
	SecurityTradingStatus_MARKET_ON_CLOSE_IMBALANCE_BUY  SecurityTradingStatus = "9"
)

//Enum values for SecurityType
type SecurityType string

const (
	SecurityType_WILDCARD_ENTRY_FOR_USE_ON_SECURITY_DEFINITION_REQUEST SecurityType = "?"
	SecurityType_ASSET_BACKED_SECURITIES                               SecurityType = "ABS"
	SecurityType_AMENDED_RESTATED                                      SecurityType = "AMENDED"
	SecurityType_OTHER_ANTICIPATION_NOTES                              SecurityType = "AN"
	SecurityType_BANKERS_ACCEPTANCE                                    SecurityType = "BA"
	SecurityType_BANK_DEPOSITORY_NOTE                                  SecurityType = "BDN"
	SecurityType_BANK_NOTES                                            SecurityType = "BN"
	SecurityType_BILL_OF_EXCHANGES                                     SecurityType = "BOX"
	SecurityType_BRADY_BOND                                            SecurityType = "BRADY"
	SecurityType_BRIDGE_LOAN                                           SecurityType = "BRIDGE"
	SecurityType_BUY_SELLBACK                                          SecurityType = "BUYSELL"
	SecurityType_CANADIAN_MONEY_MARKETS                                SecurityType = "CAMM"
	SecurityType_CANADIAN_TREASURY_NOTES                               SecurityType = "CAN"
	SecurityType_CASH                                                  SecurityType = "CASH"
	SecurityType_CONVERTIBLE_BOND                                      SecurityType = "CB"
	SecurityType_CERTIFICATE_OF_DEPOSIT                                SecurityType = "CD"
	SecurityType_CREDIT_DEFAULT_SWAP                                   SecurityType = "CDS"
	SecurityType_CALL_LOANS                                            SecurityType = "CL"
	SecurityType_CANADIAN_MORTGAGE_BONDS                               SecurityType = "CMB"
	SecurityType_CORP_MORTGAGE_BACKED_SECURITIES                       SecurityType = "CMBS"
	SecurityType_COLLATERALIZED_MORTGAGE_OBLIGATION                    SecurityType = "CMO"
	SecurityType_CERTIFICATE_OF_OBLIGATION                             SecurityType = "COFO"
	SecurityType_CERTIFICATE_OF_PARTICIPATION                          SecurityType = "COFP"
	SecurityType_CORPORATE_BOND                                        SecurityType = "CORP"
	SecurityType_COMMERCIAL_PAPER                                      SecurityType = "CP"
	SecurityType_CORPORATE_PRIVATE_PLACEMENT                           SecurityType = "CPP"
	SecurityType_COMMON_STOCK                                          SecurityType = "CS"
	SecurityType_CANADIAN_TREASURY_BILLS                               SecurityType = "CTB"
	SecurityType_DEFAULTED                                             SecurityType = "DEFLTED"
	SecurityType_DEBTOR_IN_POSSESSION                                  SecurityType = "DINP"
	SecurityType_DEPOSIT_NOTES                                         SecurityType = "DN"
	SecurityType_DUAL_CURRENCY                                         SecurityType = "DUAL"
	SecurityType_EURO_CERTIFICATE_OF_DEPOSIT                           SecurityType = "EUCD"
	SecurityType_EURO_CORPORATE_BOND                                   SecurityType = "EUCORP"
	SecurityType_EURO_COMMERCIAL_PAPER                                 SecurityType = "EUCP"
	SecurityType_EURO_CORPORATE_FLOATING_RATE_NOTES                    SecurityType = "EUFRN"
	SecurityType_EURO_SOVEREIGNS                                       SecurityType = "EUSOV"
	SecurityType_EURO_SUPRANATIONAL_COUPONS                            SecurityType = "EUSUPRA"
	SecurityType_FEDERAL_AGENCY_COUPON                                 SecurityType = "FAC"
	SecurityType_FEDERAL_AGENCY_DISCOUNT_NOTE                          SecurityType = "FADN"
	SecurityType_FEDERAL_HOUSING_AUTHORITY                             SecurityType = "FHA"
	SecurityType_FEDERAL_HOME_LOAN                                     SecurityType = "FHL"
	SecurityType_FEDERAL_NATIONAL_MORTGAGE_ASSOCIATION                 SecurityType = "FN"
	SecurityType_FOREIGN_EXCHANGE_CONTRACT                             SecurityType = "FOR"
	SecurityType_FORWARD                                               SecurityType = "FORWARD"
	SecurityType_US_CORPORATE_FLOATING_RATE_NOTES                      SecurityType = "FRN"
	SecurityType_FUTURE                                                SecurityType = "FUT"
	SecurityType_FX_FORWARD                                            SecurityType = "FXFWD"
	SecurityType_NON_DELIVERABLE_FORWARD                               SecurityType = "FXNDF"
	SecurityType_FX_SPOT                                               SecurityType = "FXSPOT"
	SecurityType_FX_SWAP                                               SecurityType = "FXSWAP"
	SecurityType_GOVERNMENT_NATIONAL_MORTGAGE_ASSOCIATION              SecurityType = "GN"
	SecurityType_GENERAL_OBLIGATION_BONDS                              SecurityType = "GO"
	SecurityType_TREASURIES_PLUS_AGENCY_DEBENTURE                      SecurityType = "GOVT"
	SecurityType_IOETTE_MORTGAGE                                       SecurityType = "IET"
	SecurityType_INTEREST_RATE_SWAP                                    SecurityType = "IRS"
	SecurityType_LETTER_OF_CREDIT                                      SecurityType = "LOFC"
	SecurityType_LIQUIDITY_NOTE                                        SecurityType = "LQN"
	SecurityType_MATURED                                               SecurityType = "MATURED"
	SecurityType_MORTGAGE_BACKED_SECURITIES                            SecurityType = "MBS"
	SecurityType_MUTUAL_FUND                                           SecurityType = "MF"
	SecurityType_MORTGAGE_INTEREST_ONLY                                SecurityType = "MIO"
	SecurityType_MULTILEG_INSTRUMENT                                   SecurityType = "MLEG"
	SecurityType_MORTGAGE_PRINCIPAL_ONLY                               SecurityType = "MPO"
	SecurityType_MORTGAGE_PRIVATE_PLACEMENT                            SecurityType = "MPP"
	SecurityType_MISCELLANEOUS_PASS_THROUGH                            SecurityType = "MPT"
	SecurityType_MANDATORY_TENDER                                      SecurityType = "MT"
	SecurityType_MEDIUM_TERM_NOTES                                     SecurityType = "MTN"
	SecurityType_MUNICIPAL_BOND                                        SecurityType = "MUNI"
	SecurityType_NO_SECURITY_TYPE                                      SecurityType = "NONE"
	SecurityType_OVERNIGHT                                             SecurityType = "ONITE"
	SecurityType_OPTIONS_ON_COMBO                                      SecurityType = "OOC"
	SecurityType_OPTIONS_ON_FUTURES                                    SecurityType = "OOF"
	SecurityType_OPTIONS_ON_PHYSICAL                                   SecurityType = "OOP"
	SecurityType_OPTION                                                SecurityType = "OPT"
	SecurityType_PRIVATE_EXPORT_FUNDING                                SecurityType = "PEF"
	SecurityType_PFANDBRIEFE                                           SecurityType = "PFAND"
	SecurityType_PROMISSORY_NOTE                                       SecurityType = "PN"
	SecurityType_AGENCY_POOLS                                          SecurityType = "POOL"
	SecurityType_CANADIAN_PROVINCIAL_BONDS                             SecurityType = "PROV"
	SecurityType_PREFERRED_STOCK                                       SecurityType = "PS"
	SecurityType_PLAZOS_FIJOS                                          SecurityType = "PZFJ"
	SecurityType_REVENUE_ANTICIPATION_NOTE                             SecurityType = "RAN"
	SecurityType_REPLACED                                              SecurityType = "REPLACD"
	SecurityType_REPURCHASE                                            SecurityType = "REPO"
	SecurityType_RETIRED                                               SecurityType = "RETIRED"
	SecurityType_REVENUE_BONDS                                         SecurityType = "REV"
	SecurityType_REPURCHASE_AGREEMENT                                  SecurityType = "RP"
	SecurityType_REVOLVER_LOAN                                         SecurityType = "RVLV"
	SecurityType_REVOLVER_TERM_LOAN                                    SecurityType = "RVLVTRM"
	SecurityType_REVERSE_REPURCHASE_AGREEMENT                          SecurityType = "RVRP"
	SecurityType_SECURITIES_LOAN                                       SecurityType = "SECLOAN"
	SecurityType_SECURITIES_PLEDGE                                     SecurityType = "SECPLEDGE"
	SecurityType_STUDENT_LOAN_MARKETING_ASSOCIATION                    SecurityType = "SL"
	SecurityType_SECURED_LIQUIDITY_NOTE                                SecurityType = "SLQN"
	SecurityType_SPECIAL_ASSESSMENT                                    SecurityType = "SPCLA"
	SecurityType_SPECIAL_OBLIGATION                                    SecurityType = "SPCLO"
	SecurityType_SPECIAL_TAX                                           SecurityType = "SPCLT"
	SecurityType_SHORT_TERM_LOAN_NOTE                                  SecurityType = "STN"
	SecurityType_STRUCTURED_NOTES                                      SecurityType = "STRUCT"
	SecurityType_USD_SUPRANATIONAL_COUPONS                             SecurityType = "SUPRA"
	SecurityType_SWING_LINE_FACILITY                                   SecurityType = "SWING"
	SecurityType_TAX_ANTICIPATION_NOTE                                 SecurityType = "TAN"
	SecurityType_TAX_ALLOCATION                                        SecurityType = "TAXA"
	SecurityType_TREASURY_BILL                                         SecurityType = "TB"
	SecurityType_TO_BE_ANNOUNCED                                       SecurityType = "TBA"
	SecurityType_US_TREASURY_BILL_TBILL                                SecurityType = "TBILL"
	SecurityType_US_TREASURY_BOND                                      SecurityType = "TBOND"
	SecurityType_PRINCIPAL_STRIP_OF_A_CALLABLE_BOND_OR_NOTE            SecurityType = "TCAL"
	SecurityType_TIME_DEPOSIT                                          SecurityType = "TD"
	SecurityType_TAX_EXEMPT_COMMERCIAL_PAPER                           SecurityType = "TECP"
	SecurityType_TERM_LOAN                                             SecurityType = "TERM"
	SecurityType_INTEREST_STRIP_FROM_ANY_BOND_OR_NOTE                  SecurityType = "TINT"
	SecurityType_TREASURY_INFLATION_PROTECTED_SECURITIES               SecurityType = "TIPS"
	SecurityType_TERM_LIQUIDITY_NOTE                                   SecurityType = "TLQN"
	SecurityType_TAXABLE_MUNICIPAL_CP                                  SecurityType = "TMCP"
	SecurityType_US_TREASURY_NOTE_TNOTE                                SecurityType = "TNOTE"
	SecurityType_PRINCIPAL_STRIP_FROM_A_NON_CALLABLE_BOND_OR_NOTE      SecurityType = "TPRN"
	SecurityType_TAX_REVENUE_ANTICIPATION_NOTE                         SecurityType = "TRAN"
	SecurityType_US_TREASURY_NOTE_UST                                  SecurityType = "UST"
	SecurityType_US_TREASURY_BILL_USTB                                 SecurityType = "USTB"
	SecurityType_VARIABLE_RATE_DEMAND_NOTE                             SecurityType = "VRDN"
	SecurityType_WARRANT                                               SecurityType = "WAR"
	SecurityType_WITHDRAWN                                             SecurityType = "WITHDRN"
	SecurityType_WILDCARD_ENTRY                                        SecurityType = "WLD"
	SecurityType_EXTENDED_COMM_NOTE                                    SecurityType = "XCN"
	SecurityType_INDEXED_LINKED                                        SecurityType = "XLINKD"
	SecurityType_YANKEE_CORPORATE_BOND                                 SecurityType = "YANK"
	SecurityType_YANKEE_CERTIFICATE_OF_DEPOSIT                         SecurityType = "YCD"
	SecurityType_CATS_TIGERS_LIONS                                     SecurityType = "ZOO"
)

//Enum values for SecurityUpdateAction
type SecurityUpdateAction string

const (
	SecurityUpdateAction_ADD    SecurityUpdateAction = "A"
	SecurityUpdateAction_DELETE SecurityUpdateAction = "D"
	SecurityUpdateAction_MODIFY SecurityUpdateAction = "M"
)

//Enum values for Seniority
type Seniority string

const (
	Seniority_SUBORDINATED   Seniority = "SB"
	Seniority_SENIOR_SECURED Seniority = "SD"
	Seniority_SENIOR         Seniority = "SR"
)

//Enum values for SessionRejectReason
type SessionRejectReason string

const (
	SessionRejectReason_INVALID_TAG_NUMBER                             SessionRejectReason = "0"
	SessionRejectReason_REQUIRED_TAG_MISSING                           SessionRejectReason = "1"
	SessionRejectReason_SENDINGTIME_ACCURACY_PROBLEM                   SessionRejectReason = "10"
	SessionRejectReason_INVALID_MSGTYPE                                SessionRejectReason = "11"
	SessionRejectReason_XML_VALIDATION_ERROR                           SessionRejectReason = "12"
	SessionRejectReason_TAG_APPEARS_MORE_THAN_ONCE                     SessionRejectReason = "13"
	SessionRejectReason_TAG_SPECIFIED_OUT_OF_REQUIRED_ORDER            SessionRejectReason = "14"
	SessionRejectReason_REPEATING_GROUP_FIELDS_OUT_OF_ORDER            SessionRejectReason = "15"
	SessionRejectReason_INCORRECT_NUMINGROUP_COUNT_FOR_REPEATING_GROUP SessionRejectReason = "16"
	SessionRejectReason_NON_DATA_VALUE_INCLUDES_FIELD_DELIMITER        SessionRejectReason = "17"
	SessionRejectReason_INVALID_UNSUPPORTED_APPLICATION_VERSION        SessionRejectReason = "18"
	SessionRejectReason_TAG_NOT_DEFINED_FOR_THIS_MESSAGE_TYPE          SessionRejectReason = "2"
	SessionRejectReason_UNDEFINED_TAG                                  SessionRejectReason = "3"
	SessionRejectReason_TAG_SPECIFIED_WITHOUT_A_VALUE                  SessionRejectReason = "4"
	SessionRejectReason_VALUE_IS_INCORRECT                             SessionRejectReason = "5"
	SessionRejectReason_INCORRECT_DATA_FORMAT_FOR_VALUE                SessionRejectReason = "6"
	SessionRejectReason_DECRYPTION_PROBLEM                             SessionRejectReason = "7"
	SessionRejectReason_SIGNATURE_PROBLEM                              SessionRejectReason = "8"
	SessionRejectReason_COMPID_PROBLEM                                 SessionRejectReason = "9"
	SessionRejectReason_OTHER                                          SessionRejectReason = "99"
)

//Enum values for SessionStatus
type SessionStatus string

const (
	SessionStatus_SESSION_ACTIVE                                   SessionStatus = "0"
	SessionStatus_SESSION_PASSWORD_CHANGED                         SessionStatus = "1"
	SessionStatus_SESSION_PASSWORD_DUE_TO_EXPIRE                   SessionStatus = "2"
	SessionStatus_NEW_SESSION_PASSWORD_DOES_NOT_COMPLY_WITH_POLICY SessionStatus = "3"
	SessionStatus_SESSION_LOGOUT_COMPLETE                          SessionStatus = "4"
	SessionStatus_INVALID_USERNAME_OR_PASSWORD                     SessionStatus = "5"
	SessionStatus_ACCOUNT_LOCKED                                   SessionStatus = "6"
	SessionStatus_LOGONS_ARE_NOT_ALLOWED_AT_THIS_TIME              SessionStatus = "7"
	SessionStatus_PASSWORD_EXPIRED                                 SessionStatus = "8"
)

//Enum values for SettlCurrFxRateCalc
type SettlCurrFxRateCalc string

const (
	SettlCurrFxRateCalc_DIVIDE   SettlCurrFxRateCalc = "D"
	SettlCurrFxRateCalc_MULTIPLY SettlCurrFxRateCalc = "M"
)

//Enum values for SettlDeliveryType
type SettlDeliveryType string

const (
	SettlDeliveryType_VERSUS_PAYMENT_DELIVER SettlDeliveryType = "0"
	SettlDeliveryType_FREE_DELIVER           SettlDeliveryType = "1"
	SettlDeliveryType_TRI_PARTY              SettlDeliveryType = "2"
	SettlDeliveryType_HOLD_IN_CUSTODY        SettlDeliveryType = "3"
)

//Enum values for SettlInstMode
type SettlInstMode string

const (
	SettlInstMode_DEFAULT                                SettlInstMode = "0"
	SettlInstMode_STANDING_INSTRUCTIONS_PROVIDED         SettlInstMode = "1"
	SettlInstMode_SPECIFIC_ALLOCATION_ACCOUNT_OVERRIDING SettlInstMode = "2"
	SettlInstMode_SPECIFIC_ALLOCATION_ACCOUNT_STANDING   SettlInstMode = "3"
	SettlInstMode_SPECIFIC_ORDER_FOR_A_SINGLE_ACCOUNT    SettlInstMode = "4"
	SettlInstMode_REQUEST_REJECT                         SettlInstMode = "5"
)

//Enum values for SettlInstReqRejCode
type SettlInstReqRejCode string

const (
	SettlInstReqRejCode_UNABLE_TO_PROCESS_REQUEST                 SettlInstReqRejCode = "0"
	SettlInstReqRejCode_UNKNOWN_ACCOUNT                           SettlInstReqRejCode = "1"
	SettlInstReqRejCode_NO_MATCHING_SETTLEMENT_INSTRUCTIONS_FOUND SettlInstReqRejCode = "2"
	SettlInstReqRejCode_OTHER                                     SettlInstReqRejCode = "99"
)

//Enum values for SettlInstSource
type SettlInstSource string

const (
	SettlInstSource_BROKERS_INSTRUCTIONS      SettlInstSource = "1"
	SettlInstSource_INSTITUTIONS_INSTRUCTIONS SettlInstSource = "2"
	SettlInstSource_INVESTOR                  SettlInstSource = "3"
)

//Enum values for SettlInstTransType
type SettlInstTransType string

const (
	SettlInstTransType_CANCEL  SettlInstTransType = "C"
	SettlInstTransType_NEW     SettlInstTransType = "N"
	SettlInstTransType_REPLACE SettlInstTransType = "R"
	SettlInstTransType_RESTATE SettlInstTransType = "T"
)

//Enum values for SettlLocation
type SettlLocation string

const (
	SettlLocation_CEDEL                        SettlLocation = "CED"
	SettlLocation_DEPOSITORY_TRUST_COMPANY     SettlLocation = "DTC"
	SettlLocation_EURO_CLEAR                   SettlLocation = "EUR"
	SettlLocation_FEDERAL_BOOK_ENTRY           SettlLocation = "FED"
	SettlLocation_LOCAL_MARKET_SETTLE_LOCATION SettlLocation = "ISO_Country_Code"
	SettlLocation_PHYSICAL                     SettlLocation = "PNY"
	SettlLocation_PARTICIPANT_TRUST_COMPANY    SettlLocation = "PTC"
)

//Enum values for SettlMethod
type SettlMethod string

const (
	SettlMethod_CASH_SETTLEMENT_REQUIRED     SettlMethod = "C"
	SettlMethod_PHYSICAL_SETTLEMENT_REQUIRED SettlMethod = "P"
)

//Enum values for SettlObligMode
type SettlObligMode string

const (
	SettlObligMode_PRELIMINARY SettlObligMode = "1"
	SettlObligMode_FINAL       SettlObligMode = "2"
)

//Enum values for SettlObligSource
type SettlObligSource string

const (
	SettlObligSource_INSTRUCTIONS_OF_BROKER       SettlObligSource = "1"
	SettlObligSource_INSTRUCTIONS_FOR_INSTITUTION SettlObligSource = "2"
	SettlObligSource_INVESTOR                     SettlObligSource = "3"
)

//Enum values for SettlObligTransType
type SettlObligTransType string

const (
	SettlObligTransType_CANCEL  SettlObligTransType = "C"
	SettlObligTransType_NEW     SettlObligTransType = "N"
	SettlObligTransType_REPLACE SettlObligTransType = "R"
	SettlObligTransType_RESTATE SettlObligTransType = "T"
)

//Enum values for SettlPriceType
type SettlPriceType string

const (
	SettlPriceType_FINAL       SettlPriceType = "1"
	SettlPriceType_THEORETICAL SettlPriceType = "2"
)

//Enum values for SettlSessID
type SettlSessID string

const (
	SettlSessID_END_OF_DAY               SettlSessID = "EOD"
	SettlSessID_ELECTRONIC_TRADING_HOURS SettlSessID = "ETH"
	SettlSessID_INTRADAY                 SettlSessID = "ITD"
	SettlSessID_REGULAR_TRADING_HOURS    SettlSessID = "RTH"
)

//Enum values for SettlType
type SettlType string

const (
	SettlType_REGULAR                 SettlType = "0"
	SettlType_CASH                    SettlType = "1"
	SettlType_NEXT_DAY                SettlType = "2"
	SettlType_T_PLUS_2                SettlType = "3"
	SettlType_T_PLUS_3                SettlType = "4"
	SettlType_T_PLUS_4                SettlType = "5"
	SettlType_FUTURE                  SettlType = "6"
	SettlType_WHEN_AND_IF_ISSUED      SettlType = "7"
	SettlType_SELLERS_OPTION          SettlType = "8"
	SettlType_T_PLUS_5                SettlType = "9"
	SettlType_BROKEN_DATE             SettlType = "B"
	SettlType_FX_SPOT_NEXT_SETTLEMENT SettlType = "C"
)

//Enum values for SettlmntTyp
type SettlmntTyp string

const (
	SettlmntTyp_REGULAR            SettlmntTyp = "0"
	SettlmntTyp_CASH               SettlmntTyp = "1"
	SettlmntTyp_NEXT_DAY           SettlmntTyp = "2"
	SettlmntTyp_T_PLUS_2           SettlmntTyp = "3"
	SettlmntTyp_T_PLUS_3           SettlmntTyp = "4"
	SettlmntTyp_T_PLUS_4           SettlmntTyp = "5"
	SettlmntTyp_FUTURE             SettlmntTyp = "6"
	SettlmntTyp_WHEN_AND_IF_ISSUED SettlmntTyp = "7"
	SettlmntTyp_SELLERS_OPTION     SettlmntTyp = "8"
	SettlmntTyp_T_PLUS_5           SettlmntTyp = "9"
	SettlmntTyp_T_PLUS_1           SettlmntTyp = "A"
)

//Enum values for ShortSaleReason
type ShortSaleReason string

const (
	ShortSaleReason_DEALER_SOLD_SHORT                        ShortSaleReason = "0"
	ShortSaleReason_DEALER_SOLD_SHORT_EXEMPT                 ShortSaleReason = "1"
	ShortSaleReason_SELLING_CUSTOMER_SOLD_SHORT              ShortSaleReason = "2"
	ShortSaleReason_SELLING_CUSTOMER_SOLD_SHORT_EXEMPT       ShortSaleReason = "3"
	ShortSaleReason_QUALIFIED_SERVICE_REPRESENTATIVE         ShortSaleReason = "4"
	ShortSaleReason_QSR_OR_AGU_CONTRA_SIDE_SOLD_SHORT_EXEMPT ShortSaleReason = "5"
)

//Enum values for Side
type Side string

const (
	Side_BUY                Side = "1"
	Side_SELL               Side = "2"
	Side_BUY_MINUS          Side = "3"
	Side_SELL_PLUS          Side = "4"
	Side_SELL_SHORT         Side = "5"
	Side_SELL_SHORT_EXEMPT  Side = "6"
	Side_UNDISCLOSED        Side = "7"
	Side_CROSS              Side = "8"
	Side_CROSS_SHORT        Side = "9"
	Side_CROSS_SHORT_EXEMPT Side = "A"
	Side_AS_DEFINED         Side = "B"
	Side_OPPOSITE           Side = "C"
	Side_SUBSCRIBE          Side = "D"
	Side_REDEEM             Side = "E"
	Side_LEND               Side = "F"
	Side_BORROW             Side = "G"
)

//Enum values for SideMultiLegReportingType
type SideMultiLegReportingType string

const (
	SideMultiLegReportingType_SINGLE_SECURITY                       SideMultiLegReportingType = "1"
	SideMultiLegReportingType_INDIVIDUAL_LEG_OF_A_MULTILEG_SECURITY SideMultiLegReportingType = "2"
	SideMultiLegReportingType_MULTILEG_SECURITY                     SideMultiLegReportingType = "3"
)

//Enum values for SideTrdSubTyp
type SideTrdSubTyp string

const (
	SideTrdSubTyp_CMTA                                            SideTrdSubTyp = "0"
	SideTrdSubTyp_INTERNAL_TRANSFER                               SideTrdSubTyp = "1"
	SideTrdSubTyp_TRANSACTION_FROM_ASSIGNMENT                     SideTrdSubTyp = "10"
	SideTrdSubTyp_EXTERNAL_TRANSFER                               SideTrdSubTyp = "2"
	SideTrdSubTyp_REJECT_FOR_SUBMITTING_TRADE                     SideTrdSubTyp = "3"
	SideTrdSubTyp_ADVISORY_FOR_CONTRA_SIDE                        SideTrdSubTyp = "4"
	SideTrdSubTyp_OFFSET_DUE_TO_AN_ALLOCATION                     SideTrdSubTyp = "5"
	SideTrdSubTyp_ONSET_DUE_TO_AN_ALLOCATION                      SideTrdSubTyp = "6"
	SideTrdSubTyp_DIFFERENTIAL_SPREAD                             SideTrdSubTyp = "7"
	SideTrdSubTyp_IMPLIED_SPREAD_LEG_EXECUTED_AGAINST_AN_OUTRIGHT SideTrdSubTyp = "8"
	SideTrdSubTyp_TRANSACTION_FROM_EXERCISE                       SideTrdSubTyp = "9"
)

//Enum values for SideValueInd
type SideValueInd string

const (
	SideValueInd_SIDE_VALUE_1 SideValueInd = "1"
	SideValueInd_SIDE_VALUE_2 SideValueInd = "2"
)

//Enum values for SolicitedFlag
type SolicitedFlag string

const (
	SolicitedFlag_NO  SolicitedFlag = "N"
	SolicitedFlag_YES SolicitedFlag = "Y"
)

//Enum values for StandInstDbType
type StandInstDbType string

const (
	StandInstDbType_OTHER              StandInstDbType = "0"
	StandInstDbType_DTC_SID            StandInstDbType = "1"
	StandInstDbType_THOMSON_ALERT      StandInstDbType = "2"
	StandInstDbType_A_GLOBAL_CUSTODIAN StandInstDbType = "3"
	StandInstDbType_ACCOUNTNET         StandInstDbType = "4"
)

//Enum values for StatsType
type StatsType string

const (
	StatsType_EXCHANGE_LAST StatsType = "1"
	StatsType_HIGH          StatsType = "2"
	StatsType_AVERAGE_PRICE StatsType = "3"
	StatsType_TURNOVER      StatsType = "4"
)

//Enum values for StatusValue
type StatusValue string

const (
	StatusValue_CONNECTED       StatusValue = "1"
	StatusValue_NOT_CONNECTED_2 StatusValue = "2"
	StatusValue_NOT_CONNECTED_3 StatusValue = "3"
	StatusValue_IN_PROCESS      StatusValue = "4"
)

//Enum values for StipulationType
type StipulationType string

const (
	StipulationType_ABSOLUTE_PREPAYMENT_SPEED                          StipulationType = "ABS"
	StipulationType_ALTERNATIVE_MINIMUM_TAX                            StipulationType = "AMT"
	StipulationType_AUTO_REINVESTMENT_AT_RATE_OR_BETTER                StipulationType = "AUTOREINV"
	StipulationType_AVAILABLE_OFFER_QUANTITY_TO_BE_SHOWN_TO_THE_STREET StipulationType = "AVAILQTY"
	StipulationType_AVERAGE_FICO_SCORE                                 StipulationType = "AVFICO"
	StipulationType_AVERAGE_LOAN_SIZE                                  StipulationType = "AVSIZE"
	StipulationType_BANK_QUALIFIED                                     StipulationType = "BANKQUAL"
	StipulationType_BARGAIN_CONDITIONS                                 StipulationType = "BGNCON"
	StipulationType_BROKERS_SALES_CREDIT                               StipulationType = "BROKERCREDIT"
	StipulationType_COUPON_RANGE                                       StipulationType = "COUPON"
	StipulationType_CONSTANT_PREPAYMENT_PENALTY                        StipulationType = "CPP"
	StipulationType_CONSTANT_PREPAYMENT_RATE                           StipulationType = "CPR"
	StipulationType_CONSTANT_PREPAYMENT_YIELD                          StipulationType = "CPY"
	StipulationType_ISO_CURRENCY_CODE                                  StipulationType = "CURRENCY"
	StipulationType_CUSTOM_START_END_DATE                              StipulationType = "CUSTOMDATE"
	StipulationType_DISCOUNT_RATE                                      StipulationType = "DISCOUNT"
	StipulationType_GEOGRAPHICS_AND_RANGE                              StipulationType = "GEOG"
	StipulationType_VALUATION_DISCOUNT                                 StipulationType = "HAIRCUT"
	StipulationType_FINAL_CPR_OF_HOME_EQUITY_PREPAYMENT_CURVE          StipulationType = "HEP"
	StipulationType_INSURED                                            StipulationType = "INSURED"
	StipulationType_OFFER_PRICE_TO_BE_SHOWN_TO_INTERNAL_BROKERS        StipulationType = "INTERNALPX"
	StipulationType_OFFER_QUANTITY_TO_BE_SHOWN_TO_INTERNAL_BROKERS     StipulationType = "INTERNALQTY"
	StipulationType_YEAR_OR_YEAR_MONTH_OF_ISSUE                        StipulationType = "ISSUE"
	StipulationType_ISSUERS_TICKER                                     StipulationType = "ISSUER"
	StipulationType_ISSUE_SIZE_RANGE                                   StipulationType = "ISSUESIZE"
	StipulationType_THE_MINIMUM_RESIDUAL_OFFER_QUANTITY                StipulationType = "LEAVEQTY"
	StipulationType_LOOKBACK_DAYS                                      StipulationType = "LOOKBACK"
	StipulationType_EXPLICIT_LOT_IDENTIFIER                            StipulationType = "LOT"
	StipulationType_LOT_VARIANCE                                       StipulationType = "LOTVAR"
	StipulationType_MATURITY_YEAR_AND_MONTH                            StipulationType = "MAT"
	StipulationType_MATURITY_RANGE                                     StipulationType = "MATURITY"
	StipulationType_MAXIMUM_LOAN_BALANCE                               StipulationType = "MAXBAL"
	StipulationType_MAXIMUMDENOMINATION                                StipulationType = "MAXDNOM"
	StipulationType_MAXIMUM_ORDER_SIZE                                 StipulationType = "MAXORDQTY"
	StipulationType_MAXIMUM_SUBSTITUTIONS                              StipulationType = "MAXSUBS"
	StipulationType_PERCENT_OF_MANUFACTURED_HOUSING_PREPAYMENT_CURVE   StipulationType = "MHP"
	StipulationType_MINIMUM_DENOMINATION                               StipulationType = "MINDNOM"
	StipulationType_MINIMUM_INCREMENT                                  StipulationType = "MININCR"
	StipulationType_MINIMUM_QUANTITY                                   StipulationType = "MINQTY"
	StipulationType_MONTHLY_PREPAYMENT_RATE                            StipulationType = "MPR"
	StipulationType_ORDER_QUANTITY_INCREMENT                           StipulationType = "ORDRINCR"
	StipulationType_PAYMENT_FREQUENCY_CALENDAR                         StipulationType = "PAYFREQ"
	StipulationType_NUMBER_OF_PIECES                                   StipulationType = "PIECES"
	StipulationType_POOLS_MAXIMUM                                      StipulationType = "PMAX"
	StipulationType_POOLSMINIMUM                                       StipulationType = "PMIN"
	StipulationType_POOL_IDENTIFIER                                    StipulationType = "POOL"
	StipulationType_PERCENT_OF_PROSPECTUS_PREPAYMENT_CURVE             StipulationType = "PPC"
	StipulationType_POOLS_PER_LOT                                      StipulationType = "PPL"
	StipulationType_POOLS_PER_MILLION                                  StipulationType = "PPM"
	StipulationType_POOLS_PER_TRADE                                    StipulationType = "PPT"
	StipulationType_PRICE_RANGE                                        StipulationType = "PRICE"
	StipulationType_PRICING_FREQUENCY                                  StipulationType = "PRICEFREQ"
	StipulationType_PRIMARY_OR_SECONDARY_MARKET_INDICATOR              StipulationType = "PRIMARY"
	StipulationType_PRODUCTION_YEAR                                    StipulationType = "PROD"
	StipulationType_CALL_PROTECTION                                    StipulationType = "PROTECT"
	StipulationType_PERCENT_OF_BMA_PREPAYMENT_CURVE                    StipulationType = "PSA"
	StipulationType_PURPOSE                                            StipulationType = "PURPOSE"
	StipulationType_BENCHMARK_PRICE_SOURCE                             StipulationType = "PXSOURCE"
	StipulationType_RATING_SOURCE_AND_RANGE                            StipulationType = "RATING"
	StipulationType_TYPE_OF_REDEMPTION                                 StipulationType = "REDEMPTION"
	StipulationType_INTEREST_OF_ROLLING_OR_CLOSING_TRADE               StipulationType = "REFINT"
	StipulationType_PRINCIPAL_OF_ROLLING_OR_CLOSING_TRADE              StipulationType = "REFPRIN"
	StipulationType_REFERENCE_TO_ROLLING_OR_CLOSING_TRADE              StipulationType = "REFTRADE"
	StipulationType_RESTRICTED                                         StipulationType = "RESTRICTED"
	StipulationType_TYPE_OF_ROLL_TRADE                                 StipulationType = "ROLLTYPE"
	StipulationType_BROKER_SALES_CREDIT_OVERRIDE                       StipulationType = "SALESCREDITOVR"
	StipulationType_MARKET_SECTOR                                      StipulationType = "SECTOR"
	StipulationType_SECURITY_TYPE_INCLUDED_OR_EXCLUDED                 StipulationType = "SECTYPE"
	StipulationType_SINGLE_MONTHLY_MORTALITY                           StipulationType = "SMM"
	StipulationType_STRUCTURE                                          StipulationType = "STRUCT"
	StipulationType_SUBSTITUTIONS_FREQUENCY                            StipulationType = "SUBSFREQ"
	StipulationType_SUBSTITUTIONS_LEFT                                 StipulationType = "SUBSLEFT"
	StipulationType_FREEFORM_TEXT                                      StipulationType = "TEXT"
	StipulationType_TRADERS_CREDIT                                     StipulationType = "TRADERCREDIT"
	StipulationType_TRADE_VARIANCE                                     StipulationType = "TRDVAR"
	StipulationType_WEIGHTED_AVERAGE_COUPON                            StipulationType = "WAC"
	StipulationType_WEIGHTED_AVERAGE_LIFE_COUPON                       StipulationType = "WAL"
	StipulationType_WEIGHTED_AVERAGE_LOAN_AGE                          StipulationType = "WALA"
	StipulationType_WEIGHTED_AVERAGE_MATURITY                          StipulationType = "WAM"
	StipulationType_WHOLE_POOL                                         StipulationType = "WHOLE"
	StipulationType_YIELD_RANGE                                        StipulationType = "YIELD"
	StipulationType_YIELD_TO_MATURITY                                  StipulationType = "YTM"
)

//Enum values for StrategyParameterType
type StrategyParameterType string

const (
	StrategyParameterType_INT                 StrategyParameterType = "1"
	StrategyParameterType_AMT                 StrategyParameterType = "10"
	StrategyParameterType_PERCENTAGE          StrategyParameterType = "11"
	StrategyParameterType_CHAR                StrategyParameterType = "12"
	StrategyParameterType_BOOLEAN             StrategyParameterType = "13"
	StrategyParameterType_STRING              StrategyParameterType = "14"
	StrategyParameterType_MULTIPLECHARVALUE   StrategyParameterType = "15"
	StrategyParameterType_CURRENCY            StrategyParameterType = "16"
	StrategyParameterType_EXCHANGE            StrategyParameterType = "17"
	StrategyParameterType_MONTHYEAR           StrategyParameterType = "18"
	StrategyParameterType_UTCTIMESTAMP        StrategyParameterType = "19"
	StrategyParameterType_LENGTH              StrategyParameterType = "2"
	StrategyParameterType_UTCTIMEONLY         StrategyParameterType = "20"
	StrategyParameterType_LOCALMKTDATE        StrategyParameterType = "21"
	StrategyParameterType_UTCDATEONLY         StrategyParameterType = "22"
	StrategyParameterType_DATA                StrategyParameterType = "23"
	StrategyParameterType_MULTIPLESTRINGVALUE StrategyParameterType = "24"
	StrategyParameterType_COUNTRY             StrategyParameterType = "25"
	StrategyParameterType_LANGUAGE            StrategyParameterType = "26"
	StrategyParameterType_TZTIMEONLY          StrategyParameterType = "27"
	StrategyParameterType_TZTIMESTAMP         StrategyParameterType = "28"
	StrategyParameterType_TENOR               StrategyParameterType = "29"
	StrategyParameterType_NUMINGROUP          StrategyParameterType = "3"
	StrategyParameterType_SEQNUM              StrategyParameterType = "4"
	StrategyParameterType_TAGNUM              StrategyParameterType = "5"
	StrategyParameterType_FLOAT               StrategyParameterType = "6"
	StrategyParameterType_QTY                 StrategyParameterType = "7"
	StrategyParameterType_PRICE               StrategyParameterType = "8"
	StrategyParameterType_PRICEOFFSET         StrategyParameterType = "9"
)

//Enum values for StreamAsgnAckType
type StreamAsgnAckType string

const (
	StreamAsgnAckType_ASSIGNMENT_ACCEPTED StreamAsgnAckType = "0"
	StreamAsgnAckType_ASSIGNMENT_REJECTED StreamAsgnAckType = "1"
)

//Enum values for StreamAsgnRejReason
type StreamAsgnRejReason string

const (
	StreamAsgnRejReason_UNKNOWN_CLIENT                   StreamAsgnRejReason = "0"
	StreamAsgnRejReason_EXCEEDS_MAXIMUM_SIZE             StreamAsgnRejReason = "1"
	StreamAsgnRejReason_UNKNOWN_OR_INVALID_CURRENCY_PAIR StreamAsgnRejReason = "2"
	StreamAsgnRejReason_NO_AVAILABLE_STREAM              StreamAsgnRejReason = "3"
	StreamAsgnRejReason_OTHER                            StreamAsgnRejReason = "99"
)

//Enum values for StreamAsgnReqType
type StreamAsgnReqType string

const (
	StreamAsgnReqType_STREAM_ASSIGNMENT_FOR_NEW_CUSTOMER      StreamAsgnReqType = "1"
	StreamAsgnReqType_STREAM_ASSIGNMENT_FOR_EXISTING_CUSTOMER StreamAsgnReqType = "2"
)

//Enum values for StreamAsgnType
type StreamAsgnType string

const (
	StreamAsgnType_ASSIGNMENT         StreamAsgnType = "1"
	StreamAsgnType_REJECTED           StreamAsgnType = "2"
	StreamAsgnType_TERMINATE_UNASSIGN StreamAsgnType = "3"
)

//Enum values for StrikePriceBoundaryMethod
type StrikePriceBoundaryMethod string

const (
	StrikePriceBoundaryMethod_LESS_THAN_UNDERLYING_PRICE_IS_IN_THE_MONEY                 StrikePriceBoundaryMethod = "1"
	StrikePriceBoundaryMethod_LESS_THAN_OR_EQUAL_TO_THE_UNDERLYING_PRICE_IS_IN_THE_MONEY StrikePriceBoundaryMethod = "2"
	StrikePriceBoundaryMethod_EQUAL_TO_THE_UNDERLYING_PRICE_IS_IN_THE_MONEY              StrikePriceBoundaryMethod = "3"
	StrikePriceBoundaryMethod_GREATER_THAN_OR_EQUAL_TO_UNDERLYING_PRICE_IS_IN_THE_MONEY  StrikePriceBoundaryMethod = "4"
	StrikePriceBoundaryMethod_GREATER_THAN_UNDERLYING_IS_IN_THE_MONEY                    StrikePriceBoundaryMethod = "5"
)

//Enum values for StrikePriceDeterminationMethod
type StrikePriceDeterminationMethod string

const (
	StrikePriceDeterminationMethod_FIXED_STRIKE                                                                       StrikePriceDeterminationMethod = "1"
	StrikePriceDeterminationMethod_STRIKE_SET_AT_EXPIRATION_TO_UNDERLYING_OR_OTHER_VALUE                              StrikePriceDeterminationMethod = "2"
	StrikePriceDeterminationMethod_STRIKE_SET_TO_AVERAGE_OF_UNDERLYING_SETTLEMENT_PRICE_ACROSS_THE_LIFE_OF_THE_OPTION StrikePriceDeterminationMethod = "3"
	StrikePriceDeterminationMethod_STRIKE_SET_TO_OPTIMAL_VALUE                                                        StrikePriceDeterminationMethod = "4"
)

//Enum values for SubscriptionRequestType
type SubscriptionRequestType string

const (
	SubscriptionRequestType_SNAPSHOT                                      SubscriptionRequestType = "0"
	SubscriptionRequestType_SNAPSHOT_PLUS_UPDATES                         SubscriptionRequestType = "1"
	SubscriptionRequestType_DISABLE_PREVIOUS_SNAPSHOT_PLUS_UPDATE_REQUEST SubscriptionRequestType = "2"
)

//Enum values for SymbolSfx
type SymbolSfx string

const (
	SymbolSfx_EUCP_WITH_LUMP_SUM_INTEREST_RATHER_THAN_DISCOUNT_PRICE               SymbolSfx = "CD"
	SymbolSfx_WHEN_ISSUED_FOR_A_SECURITY_TO_BE_REISSUED_UNDER_AN_OLD_CUSIP_OR_ISIN SymbolSfx = "WI"
)

//Enum values for TargetStrategy
type TargetStrategy string

const (
	TargetStrategy_VWAP                                                          TargetStrategy = "1"
	TargetStrategy_RESERVEDANDAVAILABLEFORBILATERALLYAGREEDUPONUSERDEFINEDVALUES TargetStrategy = "1000"
	TargetStrategy_PARTICIPATE                                                   TargetStrategy = "2"
	TargetStrategy_MININIZE_MARKET_IMPACT                                        TargetStrategy = "3"
)

//Enum values for TaxAdvantageType
type TaxAdvantageType string

const (
	TaxAdvantageType_NONE_NOT_APPLICABLE              TaxAdvantageType = "0"
	TaxAdvantageType_MAXI_ISA                         TaxAdvantageType = "1"
	TaxAdvantageType_EMPLOYEE_10                      TaxAdvantageType = "10"
	TaxAdvantageType_EMPLOYER_11                      TaxAdvantageType = "11"
	TaxAdvantageType_EMPLOYER_12                      TaxAdvantageType = "12"
	TaxAdvantageType_NON_FUND_PROTOTYPE_IRA           TaxAdvantageType = "13"
	TaxAdvantageType_NON_FUND_QUALIFIED_PLAN          TaxAdvantageType = "14"
	TaxAdvantageType_DEFINED_CONTRIBUTION_PLAN        TaxAdvantageType = "15"
	TaxAdvantageType_INDIVIDUAL_RETIREMENT_ACCOUNT_16 TaxAdvantageType = "16"
	TaxAdvantageType_INDIVIDUAL_RETIREMENT_ACCOUNT_17 TaxAdvantageType = "17"
	TaxAdvantageType_KEOGH                            TaxAdvantageType = "18"
	TaxAdvantageType_PROFIT_SHARING_PLAN              TaxAdvantageType = "19"
	TaxAdvantageType_TESSA                            TaxAdvantageType = "2"
	TaxAdvantageType_401                              TaxAdvantageType = "20"
	TaxAdvantageType_SELF_DIRECTED_IRA                TaxAdvantageType = "21"
	TaxAdvantageType_403                              TaxAdvantageType = "22"
	TaxAdvantageType_457                              TaxAdvantageType = "23"
	TaxAdvantageType_ROTH_IRA_24                      TaxAdvantageType = "24"
	TaxAdvantageType_ROTH_IRA_25                      TaxAdvantageType = "25"
	TaxAdvantageType_ROTH_CONVERSION_IRA_26           TaxAdvantageType = "26"
	TaxAdvantageType_ROTH_CONVERSION_IRA_27           TaxAdvantageType = "27"
	TaxAdvantageType_EDUCATION_IRA_28                 TaxAdvantageType = "28"
	TaxAdvantageType_EDUCATION_IRA_29                 TaxAdvantageType = "29"
	TaxAdvantageType_MINI_CASH_ISA                    TaxAdvantageType = "3"
	TaxAdvantageType_MINI_STOCKS_AND_SHARES_ISA       TaxAdvantageType = "4"
	TaxAdvantageType_MINI_INSURANCE_ISA               TaxAdvantageType = "5"
	TaxAdvantageType_CURRENT_YEAR_PAYMENT             TaxAdvantageType = "6"
	TaxAdvantageType_PRIOR_YEAR_PAYMENT               TaxAdvantageType = "7"
	TaxAdvantageType_ASSET_TRANSFER                   TaxAdvantageType = "8"
	TaxAdvantageType_EMPLOYEE_9                       TaxAdvantageType = "9"
	TaxAdvantageType_OTHER                            TaxAdvantageType = "999"
)

//Enum values for TerminationType
type TerminationType string

const (
	TerminationType_OVERNIGHT TerminationType = "1"
	TerminationType_TERM      TerminationType = "2"
	TerminationType_FLEXIBLE  TerminationType = "3"
	TerminationType_OPEN      TerminationType = "4"
)

//Enum values for TestMessageIndicator
type TestMessageIndicator string

const (
	TestMessageIndicator_NO  TestMessageIndicator = "N"
	TestMessageIndicator_YES TestMessageIndicator = "Y"
)

//Enum values for TickDirection
type TickDirection string

const (
	TickDirection_PLUS_TICK       TickDirection = "0"
	TickDirection_ZERO_PLUS_TICK  TickDirection = "1"
	TickDirection_MINUS_TICK      TickDirection = "2"
	TickDirection_ZERO_MINUS_TICK TickDirection = "3"
)

//Enum values for TickRuleType
type TickRuleType string

const (
	TickRuleType_REGULAR                 TickRuleType = "0"
	TickRuleType_VARIABLE                TickRuleType = "1"
	TickRuleType_FIXED                   TickRuleType = "2"
	TickRuleType_TRADED_AS_A_SPREAD_LEG  TickRuleType = "3"
	TickRuleType_SETTLED_AS_A_SPREAD_LEG TickRuleType = "4"
)

//Enum values for TimeInForce
type TimeInForce string

const (
	TimeInForce_DAY                   TimeInForce = "0"
	TimeInForce_GOOD_TILL_CANCEL      TimeInForce = "1"
	TimeInForce_AT_THE_OPENING        TimeInForce = "2"
	TimeInForce_IMMEDIATE_OR_CANCEL   TimeInForce = "3"
	TimeInForce_FILL_OR_KILL          TimeInForce = "4"
	TimeInForce_GOOD_TILL_CROSSING    TimeInForce = "5"
	TimeInForce_GOOD_TILL_DATE        TimeInForce = "6"
	TimeInForce_AT_THE_CLOSE          TimeInForce = "7"
	TimeInForce_GOOD_THROUGH_CROSSING TimeInForce = "8"
	TimeInForce_AT_CROSSING           TimeInForce = "9"
)

//Enum values for TimeUnit
type TimeUnit string

const (
	TimeUnit_DAY    TimeUnit = "D"
	TimeUnit_HOUR   TimeUnit = "H"
	TimeUnit_MINUTE TimeUnit = "Min"
	TimeUnit_MONTH  TimeUnit = "Mo"
	TimeUnit_SECOND TimeUnit = "S"
	TimeUnit_WEEK   TimeUnit = "Wk"
	TimeUnit_YEAR   TimeUnit = "Yr"
)

//Enum values for TradSesEvent
type TradSesEvent string

const (
	TradSesEvent_TRADING_RESUMES              TradSesEvent = "0"
	TradSesEvent_CHANGE_OF_TRADING_SESSION    TradSesEvent = "1"
	TradSesEvent_CHANGE_OF_TRADING_SUBSESSION TradSesEvent = "2"
	TradSesEvent_CHANGE_OF_TRADING_STATUS     TradSesEvent = "3"
)

//Enum values for TradSesMethod
type TradSesMethod string

const (
	TradSesMethod_ELECTRONIC  TradSesMethod = "1"
	TradSesMethod_OPEN_OUTCRY TradSesMethod = "2"
	TradSesMethod_TWO_PARTY   TradSesMethod = "3"
)

//Enum values for TradSesMode
type TradSesMode string

const (
	TradSesMode_TESTING    TradSesMode = "1"
	TradSesMode_SIMULATED  TradSesMode = "2"
	TradSesMode_PRODUCTION TradSesMode = "3"
)

//Enum values for TradSesStatus
type TradSesStatus string

const (
	TradSesStatus_UNKNOWN          TradSesStatus = "0"
	TradSesStatus_HALTED           TradSesStatus = "1"
	TradSesStatus_OPEN             TradSesStatus = "2"
	TradSesStatus_CLOSED           TradSesStatus = "3"
	TradSesStatus_PRE_OPEN         TradSesStatus = "4"
	TradSesStatus_PRE_CLOSE        TradSesStatus = "5"
	TradSesStatus_REQUEST_REJECTED TradSesStatus = "6"
)

//Enum values for TradSesStatusRejReason
type TradSesStatusRejReason string

const (
	TradSesStatusRejReason_UNKNOWN_OR_INVALID_TRADINGSESSIONID TradSesStatusRejReason = "1"
	TradSesStatusRejReason_OTHER                               TradSesStatusRejReason = "99"
)

//Enum values for TradeAllocIndicator
type TradeAllocIndicator string

const (
	TradeAllocIndicator_ALLOCATION_NOT_REQUIRED                TradeAllocIndicator = "0"
	TradeAllocIndicator_ALLOCATION_REQUIRED                    TradeAllocIndicator = "1"
	TradeAllocIndicator_USE_ALLOCATION_PROVIDED_WITH_THE_TRADE TradeAllocIndicator = "2"
	TradeAllocIndicator_ALLOCATION_GIVE_UP_EXECUTOR            TradeAllocIndicator = "3"
	TradeAllocIndicator_ALLOCATION_FROM_EXECUTOR               TradeAllocIndicator = "4"
	TradeAllocIndicator_ALLOCATION_TO_CLAIM_ACCOUNT            TradeAllocIndicator = "5"
)

//Enum values for TradeCondition
type TradeCondition string

const (
	TradeCondition_CANCEL                                  TradeCondition = "0"
	TradeCondition_IMPLIED_TRADE                           TradeCondition = "1"
	TradeCondition_MARKETPLACE_ENTERED_TRADE               TradeCondition = "2"
	TradeCondition_MULT_ASSET_CLASS_MULTILEG_TRADE         TradeCondition = "3"
	TradeCondition_MULTILEG_TO_MULTILEG_TRADE              TradeCondition = "4"
	TradeCondition_CASH                                    TradeCondition = "A"
	TradeCondition_SPREAD                                  TradeCondition = "AA"
	TradeCondition_SPREAD_ETH                              TradeCondition = "AB"
	TradeCondition_STRADDLE                                TradeCondition = "AC"
	TradeCondition_STRADDLE_ETH                            TradeCondition = "AD"
	TradeCondition_STOPPED                                 TradeCondition = "AE"
	TradeCondition_STOPPED_ETH                             TradeCondition = "AF"
	TradeCondition_REGULAR_ETH                             TradeCondition = "AG"
	TradeCondition_COMBO                                   TradeCondition = "AH"
	TradeCondition_COMBO_ETH                               TradeCondition = "AI"
	TradeCondition_OFFICIAL_CLOSING_PRICE                  TradeCondition = "AJ"
	TradeCondition_PRIOR_REFERENCE_PRICE                   TradeCondition = "AK"
	TradeCondition_STOPPED_SOLD_LAST                       TradeCondition = "AL"
	TradeCondition_STOPPED_OUT_OF_SEQUENCE                 TradeCondition = "AM"
	TradeCondition_OFFICAL_CLOSING_PRICE                   TradeCondition = "AN"
	TradeCondition_CROSSED_AO                              TradeCondition = "AO"
	TradeCondition_FAST_MARKET                             TradeCondition = "AP"
	TradeCondition_AUTOMATIC_EXECUTION                     TradeCondition = "AQ"
	TradeCondition_FORM_T                                  TradeCondition = "AR"
	TradeCondition_BASKET_INDEX                            TradeCondition = "AS"
	TradeCondition_BURST_BASKET                            TradeCondition = "AT"
	TradeCondition_OUTSIDE_SPREAD                          TradeCondition = "AV"
	TradeCondition_AVERAGE_PRICE_TRADE                     TradeCondition = "B"
	TradeCondition_CASH_TRADE                              TradeCondition = "C"
	TradeCondition_NEXT_DAY                                TradeCondition = "D"
	TradeCondition_OPENING_REOPENING_TRADE_DETAIL          TradeCondition = "E"
	TradeCondition_INTRADAY_TRADE_DETAIL                   TradeCondition = "F"
	TradeCondition_RULE_127_TRADE                          TradeCondition = "G"
	TradeCondition_RULE_155_TRADE                          TradeCondition = "H"
	TradeCondition_SOLD_LAST                               TradeCondition = "I"
	TradeCondition_NEXT_DAY_TRADE                          TradeCondition = "J"
	TradeCondition_OPENED                                  TradeCondition = "K"
	TradeCondition_SELLER                                  TradeCondition = "L"
	TradeCondition_SOLD                                    TradeCondition = "M"
	TradeCondition_STOPPED_STOCK                           TradeCondition = "N"
	TradeCondition_IMBALANCE_MORE_BUYERS                   TradeCondition = "P"
	TradeCondition_IMBALANCE_MORE_SELLERS                  TradeCondition = "Q"
	TradeCondition_OPENING_PRICE                           TradeCondition = "R"
	TradeCondition_BARGAIN_CONDITION                       TradeCondition = "S"
	TradeCondition_CONVERTED_PRICE_INDICATOR               TradeCondition = "T"
	TradeCondition_EXCHANGE_LAST                           TradeCondition = "U"
	TradeCondition_FINAL_PRICE_OF_SESSION                  TradeCondition = "V"
	TradeCondition_EX_PIT                                  TradeCondition = "W"
	TradeCondition_CROSSED_X                               TradeCondition = "X"
	TradeCondition_TRADES_RESULTING_FROM_MANUAL_SLOW_QUOTE TradeCondition = "Y"
	TradeCondition_TRADES_RESULTING_FROM_INTERMARKET_SWEEP TradeCondition = "Z"
	TradeCondition_VOLUME_ONLY                             TradeCondition = "a"
	TradeCondition_DIRECT_PLUS                             TradeCondition = "b"
	TradeCondition_ACQUISITION                             TradeCondition = "c"
	TradeCondition_BUNCHED                                 TradeCondition = "d"
	TradeCondition_DISTRIBUTION                            TradeCondition = "e"
	TradeCondition_BUNCHED_SALE                            TradeCondition = "f"
	TradeCondition_SPLIT_TRADE                             TradeCondition = "g"
	TradeCondition_CANCEL_STOPPED                          TradeCondition = "h"
	TradeCondition_CANCEL_ETH                              TradeCondition = "i"
	TradeCondition_CANCEL_STOPPED_ETH                      TradeCondition = "j"
	TradeCondition_OUT_OF_SEQUENCE_ETH                     TradeCondition = "k"
	TradeCondition_CANCEL_LAST_ETH                         TradeCondition = "l"
	TradeCondition_SOLD_LAST_SALE_ETH                      TradeCondition = "m"
	TradeCondition_CANCEL_LAST                             TradeCondition = "n"
	TradeCondition_SOLD_LAST_SALE                          TradeCondition = "o"
	TradeCondition_CANCEL_OPEN                             TradeCondition = "p"
	TradeCondition_CANCEL_OPEN_ETH                         TradeCondition = "q"
	TradeCondition_OPENED_SALE_ETH                         TradeCondition = "r"
	TradeCondition_CANCEL_ONLY                             TradeCondition = "s"
	TradeCondition_CANCEL_ONLY_ETH                         TradeCondition = "t"
	TradeCondition_LATE_OPEN_ETH                           TradeCondition = "u"
	TradeCondition_AUTO_EXECUTION_ETH                      TradeCondition = "v"
	TradeCondition_REOPEN                                  TradeCondition = "w"
	TradeCondition_REOPEN_ETH                              TradeCondition = "x"
	TradeCondition_ADJUSTED                                TradeCondition = "y"
	TradeCondition_ADJUSTED_ETH                            TradeCondition = "z"
)

//Enum values for TradeHandlingInstr
type TradeHandlingInstr string

const (
	TradeHandlingInstr_TRADE_CONFIRMATION                TradeHandlingInstr = "0"
	TradeHandlingInstr_TWO_PARTY_REPORT                  TradeHandlingInstr = "1"
	TradeHandlingInstr_ONE_PARTY_REPORT_FOR_MATCHING     TradeHandlingInstr = "2"
	TradeHandlingInstr_ONE_PARTY_REPORT_FOR_PASS_THROUGH TradeHandlingInstr = "3"
	TradeHandlingInstr_AUTOMATED_FLOOR_ORDER_ROUTING     TradeHandlingInstr = "4"
	TradeHandlingInstr_TWO_PARTY_REPORT_FOR_CLAIM        TradeHandlingInstr = "5"
)

//Enum values for TradePublishIndicator
type TradePublishIndicator string

const (
	TradePublishIndicator_DO_NOT_PUBLISH_TRADE TradePublishIndicator = "0"
	TradePublishIndicator_PUBLISH_TRADE        TradePublishIndicator = "1"
	TradePublishIndicator_DEFERRED_PUBLICATION TradePublishIndicator = "2"
)

//Enum values for TradeReportRejectReason
type TradeReportRejectReason string

const (
	TradeReportRejectReason_SUCCESSFUL                    TradeReportRejectReason = "0"
	TradeReportRejectReason_INVALID_PARTY_ONFORMATION     TradeReportRejectReason = "1"
	TradeReportRejectReason_UNKNOWN_INSTRUMENT            TradeReportRejectReason = "2"
	TradeReportRejectReason_UNAUTHORIZED_TO_REPORT_TRADES TradeReportRejectReason = "3"
	TradeReportRejectReason_INVALID_TRADE_TYPE            TradeReportRejectReason = "4"
	TradeReportRejectReason_OTHER                         TradeReportRejectReason = "99"
)

//Enum values for TradeReportTransType
type TradeReportTransType string

const (
	TradeReportTransType_NEW                             TradeReportTransType = "0"
	TradeReportTransType_CANCEL                          TradeReportTransType = "1"
	TradeReportTransType_REPLACE                         TradeReportTransType = "2"
	TradeReportTransType_RELEASE                         TradeReportTransType = "3"
	TradeReportTransType_REVERSE                         TradeReportTransType = "4"
	TradeReportTransType_CANCEL_DUE_TO_BACK_OUT_OF_TRADE TradeReportTransType = "5"
)

//Enum values for TradeReportType
type TradeReportType string

const (
	TradeReportType_SUBMIT                      TradeReportType = "0"
	TradeReportType_ALLEGED_1                   TradeReportType = "1"
	TradeReportType_PENDED                      TradeReportType = "10"
	TradeReportType_ALLEGED_NEW                 TradeReportType = "11"
	TradeReportType_ALLEGED_ADDENDUM            TradeReportType = "12"
	TradeReportType_ALLEGED_NO_WAS              TradeReportType = "13"
	TradeReportType_ALLEGED_TRADE_REPORT_CANCEL TradeReportType = "14"
	TradeReportType_ALLEGED_15                  TradeReportType = "15"
	TradeReportType_ACCEPT                      TradeReportType = "2"
	TradeReportType_DECLINE                     TradeReportType = "3"
	TradeReportType_ADDENDUM                    TradeReportType = "4"
	TradeReportType_NO_WAS                      TradeReportType = "5"
	TradeReportType_TRADE_REPORT_CANCEL         TradeReportType = "6"
	TradeReportType_7                           TradeReportType = "7"
	TradeReportType_DEFAULTED                   TradeReportType = "8"
	TradeReportType_INVALID_CMTA                TradeReportType = "9"
)

//Enum values for TradeRequestResult
type TradeRequestResult string

const (
	TradeRequestResult_SUCCESSFUL                       TradeRequestResult = "0"
	TradeRequestResult_INVALID_OR_UNKNOWN_INSTRUMENT    TradeRequestResult = "1"
	TradeRequestResult_INVALID_TYPE_OF_TRADE_REQUESTED  TradeRequestResult = "2"
	TradeRequestResult_INVALID_PARTIES                  TradeRequestResult = "3"
	TradeRequestResult_INVALID_TRANSPORT_TYPE_REQUESTED TradeRequestResult = "4"
	TradeRequestResult_INVALID_DESTINATION_REQUESTED    TradeRequestResult = "5"
	TradeRequestResult_TRADEREQUESTTYPE_NOT_SUPPORTED   TradeRequestResult = "8"
	TradeRequestResult_NOT_AUTHORIZED                   TradeRequestResult = "9"
	TradeRequestResult_OTHER                            TradeRequestResult = "99"
)

//Enum values for TradeRequestStatus
type TradeRequestStatus string

const (
	TradeRequestStatus_ACCEPTED  TradeRequestStatus = "0"
	TradeRequestStatus_COMPLETED TradeRequestStatus = "1"
	TradeRequestStatus_REJECTED  TradeRequestStatus = "2"
)

//Enum values for TradeRequestType
type TradeRequestType string

const (
	TradeRequestType_ALL_TRADES                                           TradeRequestType = "0"
	TradeRequestType_MATCHED_TRADES_MATCHING_CRITERIA_PROVIDED_ON_REQUEST TradeRequestType = "1"
	TradeRequestType_UNMATCHED_TRADES_THAT_MATCH_CRITERIA                 TradeRequestType = "2"
	TradeRequestType_UNREPORTED_TRADES_THAT_MATCH_CRITERIA                TradeRequestType = "3"
	TradeRequestType_ADVISORIES_THAT_MATCH_CRITERIA                       TradeRequestType = "4"
)

//Enum values for TradeType
type TradeType string

const (
	TradeType_AGENCY           TradeType = "A"
	TradeType_VWAP_GUARANTEE   TradeType = "G"
	TradeType_GUARANTEED_CLOSE TradeType = "J"
	TradeType_RISK_TRADE       TradeType = "R"
)

//Enum values for TradedFlatSwitch
type TradedFlatSwitch string

const (
	TradedFlatSwitch_NO  TradedFlatSwitch = "N"
	TradedFlatSwitch_YES TradedFlatSwitch = "Y"
)

//Enum values for TradingSessionID
type TradingSessionID string

const (
	TradingSessionID_DAY         TradingSessionID = "1"
	TradingSessionID_HALFDAY     TradingSessionID = "2"
	TradingSessionID_MORNING     TradingSessionID = "3"
	TradingSessionID_AFTERNOON   TradingSessionID = "4"
	TradingSessionID_EVENING     TradingSessionID = "5"
	TradingSessionID_AFTER_HOURS TradingSessionID = "6"
)

//Enum values for TradingSessionSubID
type TradingSessionSubID string

const (
	TradingSessionSubID_PRE_TRADING                TradingSessionSubID = "1"
	TradingSessionSubID_OPENING_OR_OPENING_AUCTION TradingSessionSubID = "2"
	TradingSessionSubID_3                          TradingSessionSubID = "3"
	TradingSessionSubID_CLOSING_OR_CLOSING_AUCTION TradingSessionSubID = "4"
	TradingSessionSubID_POST_TRADING               TradingSessionSubID = "5"
	TradingSessionSubID_INTRADAY_AUCTION           TradingSessionSubID = "6"
	TradingSessionSubID_QUIESCENT                  TradingSessionSubID = "7"
)

//Enum values for TrdRegTimestampType
type TrdRegTimestampType string

const (
	TrdRegTimestampType_EXECUTION_TIME         TrdRegTimestampType = "1"
	TrdRegTimestampType_TIME_IN                TrdRegTimestampType = "2"
	TrdRegTimestampType_TIME_OUT               TrdRegTimestampType = "3"
	TrdRegTimestampType_BROKER_RECEIPT         TrdRegTimestampType = "4"
	TrdRegTimestampType_BROKER_EXECUTION       TrdRegTimestampType = "5"
	TrdRegTimestampType_DESK_RECEIPT           TrdRegTimestampType = "6"
	TrdRegTimestampType_SUBMISSION_TO_CLEARING TrdRegTimestampType = "7"
)

//Enum values for TrdRptStatus
type TrdRptStatus string

const (
	TrdRptStatus_ACCEPTED             TrdRptStatus = "0"
	TrdRptStatus_REJECTED             TrdRptStatus = "1"
	TrdRptStatus_ACCEPTED_WITH_ERRORS TrdRptStatus = "3"
)

//Enum values for TrdSubType
type TrdSubType string

const (
	TrdSubType_CMTA                                            TrdSubType = "0"
	TrdSubType_INTERNAL_TRANSFER_OR_ADJUSTMENT                 TrdSubType = "1"
	TrdSubType_TRANSACTION_FROM_ASSIGNMENT                     TrdSubType = "10"
	TrdSubType_ACATS                                           TrdSubType = "11"
	TrdSubType_AI                                              TrdSubType = "14"
	TrdSubType_B                                               TrdSubType = "15"
	TrdSubType_K                                               TrdSubType = "16"
	TrdSubType_LC                                              TrdSubType = "17"
	TrdSubType_M                                               TrdSubType = "18"
	TrdSubType_N                                               TrdSubType = "19"
	TrdSubType_EXTERNAL_TRANSFER_OR_TRANSFER_OF_ACCOUNT        TrdSubType = "2"
	TrdSubType_NM                                              TrdSubType = "20"
	TrdSubType_NR                                              TrdSubType = "21"
	TrdSubType_P                                               TrdSubType = "22"
	TrdSubType_PA                                              TrdSubType = "23"
	TrdSubType_PC                                              TrdSubType = "24"
	TrdSubType_PN                                              TrdSubType = "25"
	TrdSubType_R                                               TrdSubType = "26"
	TrdSubType_RO                                              TrdSubType = "27"
	TrdSubType_RT                                              TrdSubType = "28"
	TrdSubType_SW                                              TrdSubType = "29"
	TrdSubType_REJECT_FOR_SUBMITTING_SIDE                      TrdSubType = "3"
	TrdSubType_T                                               TrdSubType = "30"
	TrdSubType_WN                                              TrdSubType = "31"
	TrdSubType_WT                                              TrdSubType = "32"
	TrdSubType_OFF_HOURS_TRADE                                 TrdSubType = "33"
	TrdSubType_ON_HOURS_TRADE                                  TrdSubType = "34"
	TrdSubType_OTC_QUOTE                                       TrdSubType = "35"
	TrdSubType_CONVERTED_SWAP                                  TrdSubType = "36"
	TrdSubType_CROSSED_TRADE                                   TrdSubType = "37"
	TrdSubType_INTERIM_PROTECTED_TRADE                         TrdSubType = "38"
	TrdSubType_LARGE_IN_SCALE                                  TrdSubType = "39"
	TrdSubType_ADVISORY_FOR_CONTRA_SIDE                        TrdSubType = "4"
	TrdSubType_OFFSET_DUE_TO_AN_ALLOCATION                     TrdSubType = "5"
	TrdSubType_ONSET_DUE_TO_AN_ALLOCATION                      TrdSubType = "6"
	TrdSubType_DIFFERENTIAL_SPREAD                             TrdSubType = "7"
	TrdSubType_IMPLIED_SPREAD_LEG_EXECUTED_AGAINST_AN_OUTRIGHT TrdSubType = "8"
	TrdSubType_TRANSACTION_FROM_EXERCISE                       TrdSubType = "9"
)

//Enum values for TrdType
type TrdType string

const (
	TrdType_REGULAR_TRADE                        TrdType = "0"
	TrdType_BLOCK_TRADE_1                        TrdType = "1"
	TrdType_AFTER_HOURS_TRADE                    TrdType = "10"
	TrdType_EXCHANGE_FOR_RISK                    TrdType = "11"
	TrdType_EXCHANGE_FOR_SWAP                    TrdType = "12"
	TrdType_EXCHANGE_OF_FUTURES_FOR              TrdType = "13"
	TrdType_EXCHANGE_OF_OPTIONS_FOR_OPTIONS      TrdType = "14"
	TrdType_TRADING_AT_SETTLEMENT                TrdType = "15"
	TrdType_ALL_OR_NONE                          TrdType = "16"
	TrdType_FUTURES_LARGE_ORDER_EXECUTION        TrdType = "17"
	TrdType_EXCHANGE_OF_FUTURES_FOR_FUTURES      TrdType = "18"
	TrdType_OPTION_INTERIM_TRADE                 TrdType = "19"
	TrdType_EFP                                  TrdType = "2"
	TrdType_OPTION_CABINET_TRADE                 TrdType = "20"
	TrdType_PRIVATELY_NEGOTIATED_TRADES          TrdType = "22"
	TrdType_SUBSTITUTION_OF_FUTURES_FOR_FORWARDS TrdType = "23"
	TrdType_ERROR_TRADE                          TrdType = "24"
	TrdType_SPECIAL_CUM_DIVIDEND                 TrdType = "25"
	TrdType_SPECIAL_EX_DIVIDEND                  TrdType = "26"
	TrdType_SPECIAL_CUM_COUPON                   TrdType = "27"
	TrdType_SPECIAL_EX_COUPON                    TrdType = "28"
	TrdType_CASH_SETTLEMENT                      TrdType = "29"
	TrdType_TRANSFER                             TrdType = "3"
	TrdType_SPECIAL_PRICE                        TrdType = "30"
	TrdType_GUARANTEED_DELIVERY                  TrdType = "31"
	TrdType_SPECIAL_CUM_RIGHTS                   TrdType = "32"
	TrdType_SPECIAL_EX_RIGHTS                    TrdType = "33"
	TrdType_SPECIAL_CUM_CAPITAL_REPAYMENTS       TrdType = "34"
	TrdType_SPECIAL_EX_CAPITAL_REPAYMENTS        TrdType = "35"
	TrdType_SPECIAL_CUM_BONUS                    TrdType = "36"
	TrdType_SPECIAL_EX_BONUS                     TrdType = "37"
	TrdType_BLOCK_TRADE_38                       TrdType = "38"
	TrdType_WORKED_PRINCIPAL_TRADE               TrdType = "39"
	TrdType_LATE_TRADE                           TrdType = "4"
	TrdType_BLOCK_TRADES                         TrdType = "40"
	TrdType_NAME_CHANGE                          TrdType = "41"
	TrdType_PORTFOLIO_TRANSFER                   TrdType = "42"
	TrdType_PROROGATION_BUY                      TrdType = "43"
	TrdType_PROROGATION_SELL                     TrdType = "44"
	TrdType_OPTION_EXERCISE                      TrdType = "45"
	TrdType_DELTA_NEUTRAL_TRANSACTION            TrdType = "46"
	TrdType_FINANCING_TRANSACTION                TrdType = "47"
	TrdType_NON_STANDARD_SETTLEMENT              TrdType = "48"
	TrdType_DERIVATIVE_RELATED_TRANSACTION       TrdType = "49"
	TrdType_T_TRADE                              TrdType = "5"
	TrdType_PORTFOLIO_TRADE                      TrdType = "50"
	TrdType_VOLUME_WEIGHTED_AVERAGE_TRADE        TrdType = "51"
	TrdType_EXCHANGE_GRANTED_TRADE               TrdType = "52"
	TrdType_REPURCHASE_AGREEMENT                 TrdType = "53"
	TrdType_OTC                                  TrdType = "54"
	TrdType_EXCHANGE_BASIS_FACILITY              TrdType = "55"
	TrdType_WEIGHTED_AVERAGE_PRICE_TRADE         TrdType = "6"
	TrdType_BUNCHED_TRADE                        TrdType = "7"
	TrdType_LATE_BUNCHED_TRADE                   TrdType = "8"
	TrdType_PRIOR_REFERENCE_PRICE_TRADE          TrdType = "9"
)

//Enum values for TriggerAction
type TriggerAction string

const (
	TriggerAction_ACTIVATE TriggerAction = "1"
	TriggerAction_MODIFY   TriggerAction = "2"
	TriggerAction_CANCEL   TriggerAction = "3"
)

//Enum values for TriggerOrderType
type TriggerOrderType string

const (
	TriggerOrderType_MARKET TriggerOrderType = "1"
	TriggerOrderType_LIMIT  TriggerOrderType = "2"
)

//Enum values for TriggerPriceDirection
type TriggerPriceDirection string

const (
	TriggerPriceDirection_TRIGGER_IF_THE_PRICE_OF_THE_SPECIFIED_TYPE_GOES_DOWN_TO_OR_THROUGH_THE_SPECIFIED_TRIGGER_PRICE TriggerPriceDirection = "D"
	TriggerPriceDirection_TRIGGER_IF_THE_PRICE_OF_THE_SPECIFIED_TYPE_GOES_UP_TO_OR_THROUGH_THE_SPECIFIED_TRIGGER_PRICE   TriggerPriceDirection = "U"
)

//Enum values for TriggerPriceType
type TriggerPriceType string

const (
	TriggerPriceType_BEST_OFFER               TriggerPriceType = "1"
	TriggerPriceType_LAST_TRADE               TriggerPriceType = "2"
	TriggerPriceType_BEST_BID                 TriggerPriceType = "3"
	TriggerPriceType_BEST_BID_OR_LAST_TRADE   TriggerPriceType = "4"
	TriggerPriceType_BEST_OFFER_OR_LAST_TRADE TriggerPriceType = "5"
	TriggerPriceType_BEST_MID                 TriggerPriceType = "6"
)

//Enum values for TriggerPriceTypeScope
type TriggerPriceTypeScope string

const (
	TriggerPriceTypeScope_NONE     TriggerPriceTypeScope = "0"
	TriggerPriceTypeScope_LOCAL    TriggerPriceTypeScope = "1"
	TriggerPriceTypeScope_NATIONAL TriggerPriceTypeScope = "2"
	TriggerPriceTypeScope_GLOBAL   TriggerPriceTypeScope = "3"
)

//Enum values for TriggerType
type TriggerType string

const (
	TriggerType_PARTIAL_EXECUTION         TriggerType = "1"
	TriggerType_SPECIFIED_TRADING_SESSION TriggerType = "2"
	TriggerType_NEXT_AUCTION              TriggerType = "3"
	TriggerType_PRICE_MOVEMENT            TriggerType = "4"
)

//Enum values for UnderlyingCashType
type UnderlyingCashType string

const (
	UnderlyingCashType_DIFF  UnderlyingCashType = "DIFF"
	UnderlyingCashType_FIXED UnderlyingCashType = "FIXED"
)

//Enum values for UnderlyingFXRateCalc
type UnderlyingFXRateCalc string

const (
	UnderlyingFXRateCalc_DIVIDE   UnderlyingFXRateCalc = "D"
	UnderlyingFXRateCalc_MULTIPLY UnderlyingFXRateCalc = "M"
)

//Enum values for UnderlyingPriceDeterminationMethod
type UnderlyingPriceDeterminationMethod string

const (
	UnderlyingPriceDeterminationMethod_REGULAR           UnderlyingPriceDeterminationMethod = "1"
	UnderlyingPriceDeterminationMethod_SPECIAL_REFERENCE UnderlyingPriceDeterminationMethod = "2"
	UnderlyingPriceDeterminationMethod_OPTIMAL_VALUE     UnderlyingPriceDeterminationMethod = "3"
	UnderlyingPriceDeterminationMethod_AVERAGE_VALUE     UnderlyingPriceDeterminationMethod = "4"
)

//Enum values for UnderlyingSettlementType
type UnderlyingSettlementType string

const (
	UnderlyingSettlementType_T_PLUS_1 UnderlyingSettlementType = "2"
	UnderlyingSettlementType_T_PLUS_3 UnderlyingSettlementType = "4"
	UnderlyingSettlementType_T_PLUS_4 UnderlyingSettlementType = "5"
)

//Enum values for UnitOfMeasure
type UnitOfMeasure string

const (
	UnitOfMeasure_ALLOWANCES         UnitOfMeasure = "Alw"
	UnitOfMeasure_BARRELS            UnitOfMeasure = "Bbl"
	UnitOfMeasure_BILLION_CUBIC_FEET UnitOfMeasure = "Bcf"
	UnitOfMeasure_BUSHELS            UnitOfMeasure = "Bu"
	UnitOfMeasure_GALLONS            UnitOfMeasure = "Gal"
	UnitOfMeasure_ONE_MILLION_BTU    UnitOfMeasure = "MMBtu"
	UnitOfMeasure_MILLION_BARRELS    UnitOfMeasure = "MMbbl"
	UnitOfMeasure_MEGAWATT_HOURS     UnitOfMeasure = "MWh"
	UnitOfMeasure_US_DOLLARS         UnitOfMeasure = "USD"
	UnitOfMeasure_POUNDS             UnitOfMeasure = "lbs"
	UnitOfMeasure_TROY_OUNCES        UnitOfMeasure = "oz_tr"
	UnitOfMeasure_METRIC_TONS        UnitOfMeasure = "t"
	UnitOfMeasure_TONS               UnitOfMeasure = "tn"
)

//Enum values for UnsolicitedIndicator
type UnsolicitedIndicator string

const (
	UnsolicitedIndicator_NO  UnsolicitedIndicator = "N"
	UnsolicitedIndicator_YES UnsolicitedIndicator = "Y"
)

//Enum values for Urgency
type Urgency string

const (
	Urgency_NORMAL     Urgency = "0"
	Urgency_FLASH      Urgency = "1"
	Urgency_BACKGROUND Urgency = "2"
)

//Enum values for UserRequestType
type UserRequestType string

const (
	UserRequestType_LOG_ON_USER                    UserRequestType = "1"
	UserRequestType_LOG_OFF_USER                   UserRequestType = "2"
	UserRequestType_CHANGE_PASSWORD_FOR_USER       UserRequestType = "3"
	UserRequestType_REQUEST_INDIVIDUAL_USER_STATUS UserRequestType = "4"
)

//Enum values for UserStatus
type UserStatus string

const (
	UserStatus_LOGGED_IN                      UserStatus = "1"
	UserStatus_NOT_LOGGED_IN                  UserStatus = "2"
	UserStatus_USER_NOT_RECOGNISED            UserStatus = "3"
	UserStatus_PASSWORD_INCORRECT             UserStatus = "4"
	UserStatus_PASSWORD_CHANGED               UserStatus = "5"
	UserStatus_OTHER                          UserStatus = "6"
	UserStatus_FORCED_USER_LOGOUT_BY_EXCHANGE UserStatus = "7"
	UserStatus_SESSION_SHUTDOWN_WARNING       UserStatus = "8"
)

//Enum values for ValuationMethod
type ValuationMethod string

const (
	ValuationMethod_CDS_STYLE_COLLATERALIZATION_OF_MARKET_TO_MARKET_AND_COUPON ValuationMethod = "CDS"
	ValuationMethod_CDS_IN_DELIVERY                                            ValuationMethod = "CDSD"
	ValuationMethod_PREMIUM_STYLE                                              ValuationMethod = "EQTY"
	ValuationMethod_FUTURES_STYLE_MARK_TO_MARKET                               ValuationMethod = "FUT"
	ValuationMethod_FUTURES_STYLE_WITH_AN_ATTACHED_CASH_ADJUSTMENT             ValuationMethod = "FUTDA"
)

//Enum values for VenueType
type VenueType string

const (
	VenueType_ELECTRONIC VenueType = "E"
	VenueType_PIT        VenueType = "P"
	VenueType_EX_PIT     VenueType = "X"
)

//Enum values for WorkingIndicator
type WorkingIndicator string

const (
	WorkingIndicator_NO  WorkingIndicator = "N"
	WorkingIndicator_YES WorkingIndicator = "Y"
)

//Enum values for YieldType
type YieldType string

const (
	YieldType_AFTER_TAX_YIELD                                                                        YieldType = "AFTERTAX"
	YieldType_ANNUAL_YIELD                                                                           YieldType = "ANNUAL"
	YieldType_YIELD_AT_ISSUE                                                                         YieldType = "ATISSUE"
	YieldType_YIELD_TO_AVERAGE_LIFE_THE_YIELD_ASSUMING_THAT_ALL_SINKS                                YieldType = "AVGLIFE"
	YieldType_YIELD_TO_AVG_MATURITY                                                                  YieldType = "AVGMATURITY"
	YieldType_BOOK_YIELD                                                                             YieldType = "BOOK"
	YieldType_YIELD_TO_NEXT_CALL                                                                     YieldType = "CALL"
	YieldType_YIELD_CHANGE_SINCE_CLOSE                                                               YieldType = "CHANGE"
	YieldType_CLOSING_YIELD                                                                          YieldType = "CLOSE"
	YieldType_COMPOUND_YIELD                                                                         YieldType = "COMPOUND"
	YieldType_CURRENT_YIELD                                                                          YieldType = "CURRENT"
	YieldType_GVNT_EQUIVALENT_YIELD                                                                  YieldType = "GOVTEQUIV"
	YieldType_TRUE_GROSS_YIELD                                                                       YieldType = "GROSS"
	YieldType_YIELD_WITH_INFLATION_ASSUMPTION                                                        YieldType = "INFLATION"
	YieldType_INVERSE_FLOATER_BOND_YIELD                                                             YieldType = "INVERSEFLOATER"
	YieldType_MOST_RECENT_CLOSING_YIELD                                                              YieldType = "LASTCLOSE"
	YieldType_CLOSING_YIELD_MOST_RECENT_MONTH                                                        YieldType = "LASTMONTH"
	YieldType_CLOSING_YIELD_MOST_RECENT_QUARTER                                                      YieldType = "LASTQUARTER"
	YieldType_CLOSING_YIELD_MOST_RECENT_YEAR                                                         YieldType = "LASTYEAR"
	YieldType_YIELD_TO_LONGEST_AVERAGE_LIFE                                                          YieldType = "LONGAVGLIFE"
	YieldType_YIELD_TO_LONGEST_AVERAGE                                                               YieldType = "LONGEST"
	YieldType_MARK_TO_MARKET_YIELD                                                                   YieldType = "MARK"
	YieldType_YIELD_TO_MATURITY                                                                      YieldType = "MATURITY"
	YieldType_YIELD_TO_NEXT_REFUND                                                                   YieldType = "NEXTREFUND"
	YieldType_OPEN_AVERAGE_YIELD                                                                     YieldType = "OPENAVG"
	YieldType_PREVIOUS_CLOSE_YIELD                                                                   YieldType = "PREVCLOSE"
	YieldType_PROCEEDS_YIELD                                                                         YieldType = "PROCEEDS"
	YieldType_YIELD_TO_NEXT_PUT                                                                      YieldType = "PUT"
	YieldType_SEMI_ANNUAL_YIELD                                                                      YieldType = "SEMIANNUAL"
	YieldType_YIELD_TO_SHORTEST_AVERAGE_LIFE                                                         YieldType = "SHORTAVGLIFE"
	YieldType_YIELD_TO_SHORTEST_AVERAGE                                                              YieldType = "SHORTEST"
	YieldType_SIMPLE_YIELD                                                                           YieldType = "SIMPLE"
	YieldType_TAX_EQUIVALENT_YIELD                                                                   YieldType = "TAXEQUIV"
	YieldType_YIELD_TO_TENDER_DATE                                                                   YieldType = "TENDER"
	YieldType_TRUE_YIELD                                                                             YieldType = "TRUE"
	YieldType_YIELD_VALUE_OF_1_32_THE_AMOUNT_THAT_THE_YIELD_WILL_CHANGE_FOR_A_1_32ND_CHANGE_IN_PRICE YieldType = "VALUE1/32"
	YieldType_YIELD_VALUE_OF_1_32                                                                    YieldType = "VALUE1_32"
	YieldType_YIELD_TO_WORST                                                                         YieldType = "WORST"
)
