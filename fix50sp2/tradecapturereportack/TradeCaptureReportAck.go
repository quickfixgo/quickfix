//Package tradecapturereportack msg type = AR.
package tradecapturereportack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/positionamountdata"
	"github.com/quickfixgo/quickfix/fix50sp2/rootparties"
	"github.com/quickfixgo/quickfix/fix50sp2/trdcaprptacksidegrp"
	"github.com/quickfixgo/quickfix/fix50sp2/trdinstrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/trdregtimestamps"
	"github.com/quickfixgo/quickfix/fix50sp2/trdrepindicatorsgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a TradeCaptureReportAck FIX Message
type Message struct {
	FIXMsgType string `fix:"AR"`
	Header     fixt11.Header
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
	Instrument instrument.Component
	//TransactTime is a non-required field for TradeCaptureReportAck.
	TransactTime *time.Time `fix:"60"`
	//TrdRegTimestamps Component
	TrdRegTimestamps trdregtimestamps.Component
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
	TrdInstrmtLegGrp trdinstrmtleggrp.Component
	//ClearingFeeIndicator is a non-required field for TradeCaptureReportAck.
	ClearingFeeIndicator *string `fix:"635"`
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
	PositionAmountData positionamountdata.Component
	//TierCode is a non-required field for TradeCaptureReportAck.
	TierCode *string `fix:"994"`
	//MessageEventSource is a non-required field for TradeCaptureReportAck.
	MessageEventSource *string `fix:"1011"`
	//LastUpdateTime is a non-required field for TradeCaptureReportAck.
	LastUpdateTime *time.Time `fix:"779"`
	//RndPx is a non-required field for TradeCaptureReportAck.
	RndPx *float64 `fix:"991"`
	//TrdCapRptAckSideGrp Component
	TrdCapRptAckSideGrp trdcaprptacksidegrp.Component
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
	RootParties rootparties.Component
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
	UndInstrmtGrp undinstrmtgrp.Component
	//RptSys is a non-required field for TradeCaptureReportAck.
	RptSys *string `fix:"1135"`
	//Currency is a non-required field for TradeCaptureReportAck.
	Currency *string `fix:"15"`
	//SettlCurrency is a non-required field for TradeCaptureReportAck.
	SettlCurrency *string `fix:"120"`
	//FeeMultiplier is a non-required field for TradeCaptureReportAck.
	FeeMultiplier *float64 `fix:"1329"`
	//TrdRepIndicatorsGrp Component
	TrdRepIndicatorsGrp trdrepindicatorsgrp.Component
	//TradePublishIndicator is a non-required field for TradeCaptureReportAck.
	TradePublishIndicator *int `fix:"1390"`
	//VenueType is a non-required field for TradeCaptureReportAck.
	VenueType *string `fix:"1430"`
	//MarketSegmentID is a non-required field for TradeCaptureReportAck.
	MarketSegmentID *string `fix:"1300"`
	//MarketID is a non-required field for TradeCaptureReportAck.
	MarketID *string `fix:"1301"`
	Trailer  fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP2, "AR", r
}
