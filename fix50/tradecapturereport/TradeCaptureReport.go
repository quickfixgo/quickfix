//Package tradecapturereport msg type = AE.
package tradecapturereport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/financingdetails"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50/positionamountdata"
	"github.com/quickfixgo/quickfix/fix50/rootparties"
	"github.com/quickfixgo/quickfix/fix50/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50/trdcaprptsidegrp"
	"github.com/quickfixgo/quickfix/fix50/trdinstrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/trdregtimestamps"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50/yielddata"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a TradeCaptureReport FIX Message
type Message struct {
	FIXMsgType string `fix:"AE"`
	Header     fixt11.Header
	//TradeReportID is a non-required field for TradeCaptureReport.
	TradeReportID *string `fix:"571"`
	//TradeReportTransType is a non-required field for TradeCaptureReport.
	TradeReportTransType *int `fix:"487"`
	//TradeReportType is a non-required field for TradeCaptureReport.
	TradeReportType *int `fix:"856"`
	//TradeRequestID is a non-required field for TradeCaptureReport.
	TradeRequestID *string `fix:"568"`
	//TrdType is a non-required field for TradeCaptureReport.
	TrdType *int `fix:"828"`
	//TrdSubType is a non-required field for TradeCaptureReport.
	TrdSubType *int `fix:"829"`
	//SecondaryTrdType is a non-required field for TradeCaptureReport.
	SecondaryTrdType *int `fix:"855"`
	//TransferReason is a non-required field for TradeCaptureReport.
	TransferReason *string `fix:"830"`
	//ExecType is a non-required field for TradeCaptureReport.
	ExecType *string `fix:"150"`
	//TotNumTradeReports is a non-required field for TradeCaptureReport.
	TotNumTradeReports *int `fix:"748"`
	//LastRptRequested is a non-required field for TradeCaptureReport.
	LastRptRequested *bool `fix:"912"`
	//UnsolicitedIndicator is a non-required field for TradeCaptureReport.
	UnsolicitedIndicator *bool `fix:"325"`
	//SubscriptionRequestType is a non-required field for TradeCaptureReport.
	SubscriptionRequestType *string `fix:"263"`
	//TradeReportRefID is a non-required field for TradeCaptureReport.
	TradeReportRefID *string `fix:"572"`
	//SecondaryTradeReportRefID is a non-required field for TradeCaptureReport.
	SecondaryTradeReportRefID *string `fix:"881"`
	//SecondaryTradeReportID is a non-required field for TradeCaptureReport.
	SecondaryTradeReportID *string `fix:"818"`
	//TradeLinkID is a non-required field for TradeCaptureReport.
	TradeLinkID *string `fix:"820"`
	//TrdMatchID is a non-required field for TradeCaptureReport.
	TrdMatchID *string `fix:"880"`
	//ExecID is a non-required field for TradeCaptureReport.
	ExecID *string `fix:"17"`
	//OrdStatus is a non-required field for TradeCaptureReport.
	OrdStatus *string `fix:"39"`
	//SecondaryExecID is a non-required field for TradeCaptureReport.
	SecondaryExecID *string `fix:"527"`
	//ExecRestatementReason is a non-required field for TradeCaptureReport.
	ExecRestatementReason *int `fix:"378"`
	//PreviouslyReported is a non-required field for TradeCaptureReport.
	PreviouslyReported *bool `fix:"570"`
	//PriceType is a non-required field for TradeCaptureReport.
	PriceType *int `fix:"423"`
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
	//QtyType is a non-required field for TradeCaptureReport.
	QtyType *int `fix:"854"`
	//YieldData Component
	YieldData yielddata.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//UnderlyingTradingSessionID is a non-required field for TradeCaptureReport.
	UnderlyingTradingSessionID *string `fix:"822"`
	//UnderlyingTradingSessionSubID is a non-required field for TradeCaptureReport.
	UnderlyingTradingSessionSubID *string `fix:"823"`
	//LastQty is a required field for TradeCaptureReport.
	LastQty float64 `fix:"32"`
	//LastPx is a required field for TradeCaptureReport.
	LastPx float64 `fix:"31"`
	//LastParPx is a non-required field for TradeCaptureReport.
	LastParPx *float64 `fix:"669"`
	//LastSpotRate is a non-required field for TradeCaptureReport.
	LastSpotRate *float64 `fix:"194"`
	//LastForwardPoints is a non-required field for TradeCaptureReport.
	LastForwardPoints *float64 `fix:"195"`
	//LastMkt is a non-required field for TradeCaptureReport.
	LastMkt *string `fix:"30"`
	//TradeDate is a required field for TradeCaptureReport.
	TradeDate string `fix:"75"`
	//ClearingBusinessDate is a non-required field for TradeCaptureReport.
	ClearingBusinessDate *string `fix:"715"`
	//AvgPx is a non-required field for TradeCaptureReport.
	AvgPx *float64 `fix:"6"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//AvgPxIndicator is a non-required field for TradeCaptureReport.
	AvgPxIndicator *int `fix:"819"`
	//PositionAmountData Component
	PositionAmountData positionamountdata.Component
	//MultiLegReportingType is a non-required field for TradeCaptureReport.
	MultiLegReportingType *string `fix:"442"`
	//TradeLegRefID is a non-required field for TradeCaptureReport.
	TradeLegRefID *string `fix:"824"`
	//TrdInstrmtLegGrp Component
	TrdInstrmtLegGrp trdinstrmtleggrp.Component
	//TransactTime is a non-required field for TradeCaptureReport.
	TransactTime *time.Time `fix:"60"`
	//TrdRegTimestamps Component
	TrdRegTimestamps trdregtimestamps.Component
	//SettlType is a non-required field for TradeCaptureReport.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for TradeCaptureReport.
	SettlDate *string `fix:"64"`
	//MatchStatus is a non-required field for TradeCaptureReport.
	MatchStatus *string `fix:"573"`
	//MatchType is a non-required field for TradeCaptureReport.
	MatchType *string `fix:"574"`
	//TrdCapRptSideGrp Component
	TrdCapRptSideGrp trdcaprptsidegrp.Component
	//CopyMsgIndicator is a non-required field for TradeCaptureReport.
	CopyMsgIndicator *bool `fix:"797"`
	//PublishTrdIndicator is a non-required field for TradeCaptureReport.
	PublishTrdIndicator *bool `fix:"852"`
	//ShortSaleReason is a non-required field for TradeCaptureReport.
	ShortSaleReason *int `fix:"853"`
	//TrdRptStatus is a non-required field for TradeCaptureReport.
	TrdRptStatus *int `fix:"939"`
	//AsOfIndicator is a non-required field for TradeCaptureReport.
	AsOfIndicator *string `fix:"1015"`
	//SettlSessID is a non-required field for TradeCaptureReport.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for TradeCaptureReport.
	SettlSessSubID *string `fix:"717"`
	//TierCode is a non-required field for TradeCaptureReport.
	TierCode *string `fix:"994"`
	//MessageEventSource is a non-required field for TradeCaptureReport.
	MessageEventSource *string `fix:"1011"`
	//LastUpdateTime is a non-required field for TradeCaptureReport.
	LastUpdateTime *time.Time `fix:"779"`
	//RndPx is a non-required field for TradeCaptureReport.
	RndPx *float64 `fix:"991"`
	//TradeID is a non-required field for TradeCaptureReport.
	TradeID *string `fix:"1003"`
	//SecondaryTradeID is a non-required field for TradeCaptureReport.
	SecondaryTradeID *string `fix:"1040"`
	//FirmTradeID is a non-required field for TradeCaptureReport.
	FirmTradeID *string `fix:"1041"`
	//SecondaryFirmTradeID is a non-required field for TradeCaptureReport.
	SecondaryFirmTradeID *string `fix:"1042"`
	//CalculatedCcyLastQty is a non-required field for TradeCaptureReport.
	CalculatedCcyLastQty *float64 `fix:"1056"`
	//LastSwapPoints is a non-required field for TradeCaptureReport.
	LastSwapPoints *float64 `fix:"1071"`
	//UnderlyingSettlementDate is a non-required field for TradeCaptureReport.
	UnderlyingSettlementDate *string `fix:"987"`
	//GrossTradeAmt is a non-required field for TradeCaptureReport.
	GrossTradeAmt *float64 `fix:"381"`
	//RootParties Component
	RootParties rootparties.Component
	//OrderCategory is a non-required field for TradeCaptureReport.
	OrderCategory *string `fix:"1115"`
	//TradeHandlingInstr is a non-required field for TradeCaptureReport.
	TradeHandlingInstr *string `fix:"1123"`
	//OrigTradeHandlingInstr is a non-required field for TradeCaptureReport.
	OrigTradeHandlingInstr *string `fix:"1124"`
	//OrigTradeDate is a non-required field for TradeCaptureReport.
	OrigTradeDate *string `fix:"1125"`
	//OrigTradeID is a non-required field for TradeCaptureReport.
	OrigTradeID *string `fix:"1126"`
	//OrigSecondaryTradeID is a non-required field for TradeCaptureReport.
	OrigSecondaryTradeID *string `fix:"1127"`
	//TZTransactTime is a non-required field for TradeCaptureReport.
	TZTransactTime *string `fix:"1132"`
	//ReportedPxDiff is a non-required field for TradeCaptureReport.
	ReportedPxDiff *bool `fix:"1134"`
	Trailer        fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX50, "AE", r
}
