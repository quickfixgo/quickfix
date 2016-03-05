//Package tradecapturereportack msg type = AR.
package tradecapturereportack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/positionamountdata"
	"github.com/quickfixgo/quickfix/fix50/rootparties"
	"github.com/quickfixgo/quickfix/fix50/trdcaprptacksidegrp"
	"github.com/quickfixgo/quickfix/fix50/trdinstrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/trdregtimestamps"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a TradeCaptureReportAck FIX Message
type Message struct {
	FIXMsgType string `fix:"AR"`
	fixt11.Header
	//TradeReportID is a non-required field for TradeCaptureReportAck.
	TradeReportID *string `fix:"571"`
	//TradeReportTransType is a non-required field for TradeCaptureReportAck.
	TradeReportTransType *int `fix:"487"`
	//TradeReportType is a non-required field for TradeCaptureReportAck.
	TradeReportType *int `fix:"856"`
	//TrdType is a non-required field for TradeCaptureReportAck.
	TrdType *int `fix:"828"`
	//TrdSubType is a non-required field for TradeCaptureReportAck.
	TrdSubType *int `fix:"829"`
	//SecondaryTrdType is a non-required field for TradeCaptureReportAck.
	SecondaryTrdType *int `fix:"855"`
	//TransferReason is a non-required field for TradeCaptureReportAck.
	TransferReason *string `fix:"830"`
	//ExecType is a non-required field for TradeCaptureReportAck.
	ExecType *string `fix:"150"`
	//TradeReportRefID is a non-required field for TradeCaptureReportAck.
	TradeReportRefID *string `fix:"572"`
	//SecondaryTradeReportRefID is a non-required field for TradeCaptureReportAck.
	SecondaryTradeReportRefID *string `fix:"881"`
	//TrdRptStatus is a non-required field for TradeCaptureReportAck.
	TrdRptStatus *int `fix:"939"`
	//TradeReportRejectReason is a non-required field for TradeCaptureReportAck.
	TradeReportRejectReason *int `fix:"751"`
	//SecondaryTradeReportID is a non-required field for TradeCaptureReportAck.
	SecondaryTradeReportID *string `fix:"818"`
	//SubscriptionRequestType is a non-required field for TradeCaptureReportAck.
	SubscriptionRequestType *string `fix:"263"`
	//TradeLinkID is a non-required field for TradeCaptureReportAck.
	TradeLinkID *string `fix:"820"`
	//TrdMatchID is a non-required field for TradeCaptureReportAck.
	TrdMatchID *string `fix:"880"`
	//ExecID is a non-required field for TradeCaptureReportAck.
	ExecID *string `fix:"17"`
	//SecondaryExecID is a non-required field for TradeCaptureReportAck.
	SecondaryExecID *string `fix:"527"`
	//Instrument Component
	instrument.Instrument
	//TransactTime is a non-required field for TradeCaptureReportAck.
	TransactTime *time.Time `fix:"60"`
	//TrdRegTimestamps Component
	trdregtimestamps.TrdRegTimestamps
	//ResponseTransportType is a non-required field for TradeCaptureReportAck.
	ResponseTransportType *int `fix:"725"`
	//ResponseDestination is a non-required field for TradeCaptureReportAck.
	ResponseDestination *string `fix:"726"`
	//Text is a non-required field for TradeCaptureReportAck.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for TradeCaptureReportAck.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for TradeCaptureReportAck.
	EncodedText *string `fix:"355"`
	//TrdInstrmtLegGrp Component
	trdinstrmtleggrp.TrdInstrmtLegGrp
	//ClearingFeeIndicator is a non-required field for TradeCaptureReportAck.
	ClearingFeeIndicator *string `fix:"635"`
	//OrdStatus is a non-required field for TradeCaptureReportAck.
	OrdStatus *string `fix:"39"`
	//ExecRestatementReason is a non-required field for TradeCaptureReportAck.
	ExecRestatementReason *int `fix:"378"`
	//PreviouslyReported is a non-required field for TradeCaptureReportAck.
	PreviouslyReported *bool `fix:"570"`
	//PriceType is a non-required field for TradeCaptureReportAck.
	PriceType *int `fix:"423"`
	//UnderlyingTradingSessionID is a non-required field for TradeCaptureReportAck.
	UnderlyingTradingSessionID *string `fix:"822"`
	//QtyType is a non-required field for TradeCaptureReportAck.
	QtyType *int `fix:"854"`
	//UnderlyingTradingSessionSubID is a non-required field for TradeCaptureReportAck.
	UnderlyingTradingSessionSubID *string `fix:"823"`
	//LastQty is a non-required field for TradeCaptureReportAck.
	LastQty *float64 `fix:"32"`
	//LastPx is a non-required field for TradeCaptureReportAck.
	LastPx *float64 `fix:"31"`
	//LastParPx is a non-required field for TradeCaptureReportAck.
	LastParPx *float64 `fix:"669"`
	//LastSpotRate is a non-required field for TradeCaptureReportAck.
	LastSpotRate *float64 `fix:"194"`
	//LastForwardPoints is a non-required field for TradeCaptureReportAck.
	LastForwardPoints *float64 `fix:"195"`
	//LastMkt is a non-required field for TradeCaptureReportAck.
	LastMkt *string `fix:"30"`
	//TradeDate is a non-required field for TradeCaptureReportAck.
	TradeDate *string `fix:"75"`
	//ClearingBusinessDate is a non-required field for TradeCaptureReportAck.
	ClearingBusinessDate *string `fix:"715"`
	//AvgPx is a non-required field for TradeCaptureReportAck.
	AvgPx *float64 `fix:"6"`
	//AvgPxIndicator is a non-required field for TradeCaptureReportAck.
	AvgPxIndicator *int `fix:"819"`
	//MultiLegReportingType is a non-required field for TradeCaptureReportAck.
	MultiLegReportingType *string `fix:"442"`
	//TradeLegRefID is a non-required field for TradeCaptureReportAck.
	TradeLegRefID *string `fix:"824"`
	//SettlType is a non-required field for TradeCaptureReportAck.
	SettlType *string `fix:"63"`
	//MatchStatus is a non-required field for TradeCaptureReportAck.
	MatchStatus *string `fix:"573"`
	//MatchType is a non-required field for TradeCaptureReportAck.
	MatchType *string `fix:"574"`
	//CopyMsgIndicator is a non-required field for TradeCaptureReportAck.
	CopyMsgIndicator *bool `fix:"797"`
	//PublishTrdIndicator is a non-required field for TradeCaptureReportAck.
	PublishTrdIndicator *bool `fix:"852"`
	//ShortSaleReason is a non-required field for TradeCaptureReportAck.
	ShortSaleReason *int `fix:"853"`
	//SettlDate is a non-required field for TradeCaptureReportAck.
	SettlDate *string `fix:"64"`
	//SettlSessID is a non-required field for TradeCaptureReportAck.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for TradeCaptureReportAck.
	SettlSessSubID *string `fix:"717"`
	//PositionAmountData Component
	positionamountdata.PositionAmountData
	//TierCode is a non-required field for TradeCaptureReportAck.
	TierCode *string `fix:"994"`
	//MessageEventSource is a non-required field for TradeCaptureReportAck.
	MessageEventSource *string `fix:"1011"`
	//LastUpdateTime is a non-required field for TradeCaptureReportAck.
	LastUpdateTime *time.Time `fix:"779"`
	//RndPx is a non-required field for TradeCaptureReportAck.
	RndPx *float64 `fix:"991"`
	//TrdCapRptAckSideGrp Component
	trdcaprptacksidegrp.TrdCapRptAckSideGrp
	//AsOfIndicator is a non-required field for TradeCaptureReportAck.
	AsOfIndicator *string `fix:"1015"`
	//TradeID is a non-required field for TradeCaptureReportAck.
	TradeID *string `fix:"1003"`
	//SecondaryTradeID is a non-required field for TradeCaptureReportAck.
	SecondaryTradeID *string `fix:"1040"`
	//FirmTradeID is a non-required field for TradeCaptureReportAck.
	FirmTradeID *string `fix:"1041"`
	//SecondaryFirmTradeID is a non-required field for TradeCaptureReportAck.
	SecondaryFirmTradeID *string `fix:"1042"`
	//CalculatedCcyLastQty is a non-required field for TradeCaptureReportAck.
	CalculatedCcyLastQty *float64 `fix:"1056"`
	//LastSwapPoints is a non-required field for TradeCaptureReportAck.
	LastSwapPoints *float64 `fix:"1071"`
	//GrossTradeAmt is a non-required field for TradeCaptureReportAck.
	GrossTradeAmt *float64 `fix:"381"`
	//RootParties Component
	rootparties.RootParties
	//TradeHandlingInstr is a non-required field for TradeCaptureReportAck.
	TradeHandlingInstr *string `fix:"1123"`
	//OrigTradeHandlingInstr is a non-required field for TradeCaptureReportAck.
	OrigTradeHandlingInstr *string `fix:"1124"`
	//OrigTradeDate is a non-required field for TradeCaptureReportAck.
	OrigTradeDate *string `fix:"1125"`
	//OrigTradeID is a non-required field for TradeCaptureReportAck.
	OrigTradeID *string `fix:"1126"`
	//OrigSecondaryTradeID is a non-required field for TradeCaptureReportAck.
	OrigSecondaryTradeID *string `fix:"1127"`
	//UndInstrmtGrp Component
	undinstrmtgrp.UndInstrmtGrp
	//RptSys is a non-required field for TradeCaptureReportAck.
	RptSys *string `fix:"1135"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetTradeReportID(v string)                 { m.TradeReportID = &v }
func (m *Message) SetTradeReportTransType(v int)             { m.TradeReportTransType = &v }
func (m *Message) SetTradeReportType(v int)                  { m.TradeReportType = &v }
func (m *Message) SetTrdType(v int)                          { m.TrdType = &v }
func (m *Message) SetTrdSubType(v int)                       { m.TrdSubType = &v }
func (m *Message) SetSecondaryTrdType(v int)                 { m.SecondaryTrdType = &v }
func (m *Message) SetTransferReason(v string)                { m.TransferReason = &v }
func (m *Message) SetExecType(v string)                      { m.ExecType = &v }
func (m *Message) SetTradeReportRefID(v string)              { m.TradeReportRefID = &v }
func (m *Message) SetSecondaryTradeReportRefID(v string)     { m.SecondaryTradeReportRefID = &v }
func (m *Message) SetTrdRptStatus(v int)                     { m.TrdRptStatus = &v }
func (m *Message) SetTradeReportRejectReason(v int)          { m.TradeReportRejectReason = &v }
func (m *Message) SetSecondaryTradeReportID(v string)        { m.SecondaryTradeReportID = &v }
func (m *Message) SetSubscriptionRequestType(v string)       { m.SubscriptionRequestType = &v }
func (m *Message) SetTradeLinkID(v string)                   { m.TradeLinkID = &v }
func (m *Message) SetTrdMatchID(v string)                    { m.TrdMatchID = &v }
func (m *Message) SetExecID(v string)                        { m.ExecID = &v }
func (m *Message) SetSecondaryExecID(v string)               { m.SecondaryExecID = &v }
func (m *Message) SetTransactTime(v time.Time)               { m.TransactTime = &v }
func (m *Message) SetResponseTransportType(v int)            { m.ResponseTransportType = &v }
func (m *Message) SetResponseDestination(v string)           { m.ResponseDestination = &v }
func (m *Message) SetText(v string)                          { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                   { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                   { m.EncodedText = &v }
func (m *Message) SetClearingFeeIndicator(v string)          { m.ClearingFeeIndicator = &v }
func (m *Message) SetOrdStatus(v string)                     { m.OrdStatus = &v }
func (m *Message) SetExecRestatementReason(v int)            { m.ExecRestatementReason = &v }
func (m *Message) SetPreviouslyReported(v bool)              { m.PreviouslyReported = &v }
func (m *Message) SetPriceType(v int)                        { m.PriceType = &v }
func (m *Message) SetUnderlyingTradingSessionID(v string)    { m.UnderlyingTradingSessionID = &v }
func (m *Message) SetQtyType(v int)                          { m.QtyType = &v }
func (m *Message) SetUnderlyingTradingSessionSubID(v string) { m.UnderlyingTradingSessionSubID = &v }
func (m *Message) SetLastQty(v float64)                      { m.LastQty = &v }
func (m *Message) SetLastPx(v float64)                       { m.LastPx = &v }
func (m *Message) SetLastParPx(v float64)                    { m.LastParPx = &v }
func (m *Message) SetLastSpotRate(v float64)                 { m.LastSpotRate = &v }
func (m *Message) SetLastForwardPoints(v float64)            { m.LastForwardPoints = &v }
func (m *Message) SetLastMkt(v string)                       { m.LastMkt = &v }
func (m *Message) SetTradeDate(v string)                     { m.TradeDate = &v }
func (m *Message) SetClearingBusinessDate(v string)          { m.ClearingBusinessDate = &v }
func (m *Message) SetAvgPx(v float64)                        { m.AvgPx = &v }
func (m *Message) SetAvgPxIndicator(v int)                   { m.AvgPxIndicator = &v }
func (m *Message) SetMultiLegReportingType(v string)         { m.MultiLegReportingType = &v }
func (m *Message) SetTradeLegRefID(v string)                 { m.TradeLegRefID = &v }
func (m *Message) SetSettlType(v string)                     { m.SettlType = &v }
func (m *Message) SetMatchStatus(v string)                   { m.MatchStatus = &v }
func (m *Message) SetMatchType(v string)                     { m.MatchType = &v }
func (m *Message) SetCopyMsgIndicator(v bool)                { m.CopyMsgIndicator = &v }
func (m *Message) SetPublishTrdIndicator(v bool)             { m.PublishTrdIndicator = &v }
func (m *Message) SetShortSaleReason(v int)                  { m.ShortSaleReason = &v }
func (m *Message) SetSettlDate(v string)                     { m.SettlDate = &v }
func (m *Message) SetSettlSessID(v string)                   { m.SettlSessID = &v }
func (m *Message) SetSettlSessSubID(v string)                { m.SettlSessSubID = &v }
func (m *Message) SetTierCode(v string)                      { m.TierCode = &v }
func (m *Message) SetMessageEventSource(v string)            { m.MessageEventSource = &v }
func (m *Message) SetLastUpdateTime(v time.Time)             { m.LastUpdateTime = &v }
func (m *Message) SetRndPx(v float64)                        { m.RndPx = &v }
func (m *Message) SetAsOfIndicator(v string)                 { m.AsOfIndicator = &v }
func (m *Message) SetTradeID(v string)                       { m.TradeID = &v }
func (m *Message) SetSecondaryTradeID(v string)              { m.SecondaryTradeID = &v }
func (m *Message) SetFirmTradeID(v string)                   { m.FirmTradeID = &v }
func (m *Message) SetSecondaryFirmTradeID(v string)          { m.SecondaryFirmTradeID = &v }
func (m *Message) SetCalculatedCcyLastQty(v float64)         { m.CalculatedCcyLastQty = &v }
func (m *Message) SetLastSwapPoints(v float64)               { m.LastSwapPoints = &v }
func (m *Message) SetGrossTradeAmt(v float64)                { m.GrossTradeAmt = &v }
func (m *Message) SetTradeHandlingInstr(v string)            { m.TradeHandlingInstr = &v }
func (m *Message) SetOrigTradeHandlingInstr(v string)        { m.OrigTradeHandlingInstr = &v }
func (m *Message) SetOrigTradeDate(v string)                 { m.OrigTradeDate = &v }
func (m *Message) SetOrigTradeID(v string)                   { m.OrigTradeID = &v }
func (m *Message) SetOrigSecondaryTradeID(v string)          { m.OrigSecondaryTradeID = &v }
func (m *Message) SetRptSys(v string)                        { m.RptSys = &v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX50, "AR", r
}
