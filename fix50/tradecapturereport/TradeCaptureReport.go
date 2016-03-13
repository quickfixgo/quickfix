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
	fixt11.Header
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
	//Instrument is a required component for TradeCaptureReport.
	instrument.Instrument
	//FinancingDetails is a non-required component for TradeCaptureReport.
	FinancingDetails *financingdetails.FinancingDetails
	//OrderQtyData is a non-required component for TradeCaptureReport.
	OrderQtyData *orderqtydata.OrderQtyData
	//QtyType is a non-required field for TradeCaptureReport.
	QtyType *int `fix:"854"`
	//YieldData is a non-required component for TradeCaptureReport.
	YieldData *yielddata.YieldData
	//UndInstrmtGrp is a non-required component for TradeCaptureReport.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
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
	//SpreadOrBenchmarkCurveData is a non-required component for TradeCaptureReport.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//AvgPxIndicator is a non-required field for TradeCaptureReport.
	AvgPxIndicator *int `fix:"819"`
	//PositionAmountData is a non-required component for TradeCaptureReport.
	PositionAmountData *positionamountdata.PositionAmountData
	//MultiLegReportingType is a non-required field for TradeCaptureReport.
	MultiLegReportingType *string `fix:"442"`
	//TradeLegRefID is a non-required field for TradeCaptureReport.
	TradeLegRefID *string `fix:"824"`
	//TrdInstrmtLegGrp is a non-required component for TradeCaptureReport.
	TrdInstrmtLegGrp *trdinstrmtleggrp.TrdInstrmtLegGrp
	//TransactTime is a non-required field for TradeCaptureReport.
	TransactTime *time.Time `fix:"60"`
	//TrdRegTimestamps is a non-required component for TradeCaptureReport.
	TrdRegTimestamps *trdregtimestamps.TrdRegTimestamps
	//SettlType is a non-required field for TradeCaptureReport.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for TradeCaptureReport.
	SettlDate *string `fix:"64"`
	//MatchStatus is a non-required field for TradeCaptureReport.
	MatchStatus *string `fix:"573"`
	//MatchType is a non-required field for TradeCaptureReport.
	MatchType *string `fix:"574"`
	//TrdCapRptSideGrp is a required component for TradeCaptureReport.
	trdcaprptsidegrp.TrdCapRptSideGrp
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
	//RootParties is a non-required component for TradeCaptureReport.
	RootParties *rootparties.RootParties
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
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetTradeReportID(v string)                               { m.TradeReportID = &v }
func (m *Message) SetTradeReportTransType(v int)                           { m.TradeReportTransType = &v }
func (m *Message) SetTradeReportType(v int)                                { m.TradeReportType = &v }
func (m *Message) SetTradeRequestID(v string)                              { m.TradeRequestID = &v }
func (m *Message) SetTrdType(v int)                                        { m.TrdType = &v }
func (m *Message) SetTrdSubType(v int)                                     { m.TrdSubType = &v }
func (m *Message) SetSecondaryTrdType(v int)                               { m.SecondaryTrdType = &v }
func (m *Message) SetTransferReason(v string)                              { m.TransferReason = &v }
func (m *Message) SetExecType(v string)                                    { m.ExecType = &v }
func (m *Message) SetTotNumTradeReports(v int)                             { m.TotNumTradeReports = &v }
func (m *Message) SetLastRptRequested(v bool)                              { m.LastRptRequested = &v }
func (m *Message) SetUnsolicitedIndicator(v bool)                          { m.UnsolicitedIndicator = &v }
func (m *Message) SetSubscriptionRequestType(v string)                     { m.SubscriptionRequestType = &v }
func (m *Message) SetTradeReportRefID(v string)                            { m.TradeReportRefID = &v }
func (m *Message) SetSecondaryTradeReportRefID(v string)                   { m.SecondaryTradeReportRefID = &v }
func (m *Message) SetSecondaryTradeReportID(v string)                      { m.SecondaryTradeReportID = &v }
func (m *Message) SetTradeLinkID(v string)                                 { m.TradeLinkID = &v }
func (m *Message) SetTrdMatchID(v string)                                  { m.TrdMatchID = &v }
func (m *Message) SetExecID(v string)                                      { m.ExecID = &v }
func (m *Message) SetOrdStatus(v string)                                   { m.OrdStatus = &v }
func (m *Message) SetSecondaryExecID(v string)                             { m.SecondaryExecID = &v }
func (m *Message) SetExecRestatementReason(v int)                          { m.ExecRestatementReason = &v }
func (m *Message) SetPreviouslyReported(v bool)                            { m.PreviouslyReported = &v }
func (m *Message) SetPriceType(v int)                                      { m.PriceType = &v }
func (m *Message) SetInstrument(v instrument.Instrument)                   { m.Instrument = v }
func (m *Message) SetFinancingDetails(v financingdetails.FinancingDetails) { m.FinancingDetails = &v }
func (m *Message) SetOrderQtyData(v orderqtydata.OrderQtyData)             { m.OrderQtyData = &v }
func (m *Message) SetQtyType(v int)                                        { m.QtyType = &v }
func (m *Message) SetYieldData(v yielddata.YieldData)                      { m.YieldData = &v }
func (m *Message) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp)          { m.UndInstrmtGrp = &v }
func (m *Message) SetUnderlyingTradingSessionID(v string)                  { m.UnderlyingTradingSessionID = &v }
func (m *Message) SetUnderlyingTradingSessionSubID(v string)               { m.UnderlyingTradingSessionSubID = &v }
func (m *Message) SetLastQty(v float64)                                    { m.LastQty = v }
func (m *Message) SetLastPx(v float64)                                     { m.LastPx = v }
func (m *Message) SetLastParPx(v float64)                                  { m.LastParPx = &v }
func (m *Message) SetLastSpotRate(v float64)                               { m.LastSpotRate = &v }
func (m *Message) SetLastForwardPoints(v float64)                          { m.LastForwardPoints = &v }
func (m *Message) SetLastMkt(v string)                                     { m.LastMkt = &v }
func (m *Message) SetTradeDate(v string)                                   { m.TradeDate = v }
func (m *Message) SetClearingBusinessDate(v string)                        { m.ClearingBusinessDate = &v }
func (m *Message) SetAvgPx(v float64)                                      { m.AvgPx = &v }
func (m *Message) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *Message) SetAvgPxIndicator(v int) { m.AvgPxIndicator = &v }
func (m *Message) SetPositionAmountData(v positionamountdata.PositionAmountData) {
	m.PositionAmountData = &v
}
func (m *Message) SetMultiLegReportingType(v string)                       { m.MultiLegReportingType = &v }
func (m *Message) SetTradeLegRefID(v string)                               { m.TradeLegRefID = &v }
func (m *Message) SetTrdInstrmtLegGrp(v trdinstrmtleggrp.TrdInstrmtLegGrp) { m.TrdInstrmtLegGrp = &v }
func (m *Message) SetTransactTime(v time.Time)                             { m.TransactTime = &v }
func (m *Message) SetTrdRegTimestamps(v trdregtimestamps.TrdRegTimestamps) { m.TrdRegTimestamps = &v }
func (m *Message) SetSettlType(v string)                                   { m.SettlType = &v }
func (m *Message) SetSettlDate(v string)                                   { m.SettlDate = &v }
func (m *Message) SetMatchStatus(v string)                                 { m.MatchStatus = &v }
func (m *Message) SetMatchType(v string)                                   { m.MatchType = &v }
func (m *Message) SetTrdCapRptSideGrp(v trdcaprptsidegrp.TrdCapRptSideGrp) { m.TrdCapRptSideGrp = v }
func (m *Message) SetCopyMsgIndicator(v bool)                              { m.CopyMsgIndicator = &v }
func (m *Message) SetPublishTrdIndicator(v bool)                           { m.PublishTrdIndicator = &v }
func (m *Message) SetShortSaleReason(v int)                                { m.ShortSaleReason = &v }
func (m *Message) SetTrdRptStatus(v int)                                   { m.TrdRptStatus = &v }
func (m *Message) SetAsOfIndicator(v string)                               { m.AsOfIndicator = &v }
func (m *Message) SetSettlSessID(v string)                                 { m.SettlSessID = &v }
func (m *Message) SetSettlSessSubID(v string)                              { m.SettlSessSubID = &v }
func (m *Message) SetTierCode(v string)                                    { m.TierCode = &v }
func (m *Message) SetMessageEventSource(v string)                          { m.MessageEventSource = &v }
func (m *Message) SetLastUpdateTime(v time.Time)                           { m.LastUpdateTime = &v }
func (m *Message) SetRndPx(v float64)                                      { m.RndPx = &v }
func (m *Message) SetTradeID(v string)                                     { m.TradeID = &v }
func (m *Message) SetSecondaryTradeID(v string)                            { m.SecondaryTradeID = &v }
func (m *Message) SetFirmTradeID(v string)                                 { m.FirmTradeID = &v }
func (m *Message) SetSecondaryFirmTradeID(v string)                        { m.SecondaryFirmTradeID = &v }
func (m *Message) SetCalculatedCcyLastQty(v float64)                       { m.CalculatedCcyLastQty = &v }
func (m *Message) SetLastSwapPoints(v float64)                             { m.LastSwapPoints = &v }
func (m *Message) SetUnderlyingSettlementDate(v string)                    { m.UnderlyingSettlementDate = &v }
func (m *Message) SetGrossTradeAmt(v float64)                              { m.GrossTradeAmt = &v }
func (m *Message) SetRootParties(v rootparties.RootParties)                { m.RootParties = &v }
func (m *Message) SetOrderCategory(v string)                               { m.OrderCategory = &v }
func (m *Message) SetTradeHandlingInstr(v string)                          { m.TradeHandlingInstr = &v }
func (m *Message) SetOrigTradeHandlingInstr(v string)                      { m.OrigTradeHandlingInstr = &v }
func (m *Message) SetOrigTradeDate(v string)                               { m.OrigTradeDate = &v }
func (m *Message) SetOrigTradeID(v string)                                 { m.OrigTradeID = &v }
func (m *Message) SetOrigSecondaryTradeID(v string)                        { m.OrigSecondaryTradeID = &v }
func (m *Message) SetTZTransactTime(v string)                              { m.TZTransactTime = &v }
func (m *Message) SetReportedPxDiff(v bool)                                { m.ReportedPxDiff = &v }

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
