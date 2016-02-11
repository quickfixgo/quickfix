//Package tradecapturereportrequest msg type = AD.
package tradecapturereportrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/financingdetails"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/trdcapdtgrp"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a TradeCaptureReportRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"AD"`
	Header     fixt11.Header
	//TradeRequestID is a required field for TradeCaptureReportRequest.
	TradeRequestID string `fix:"568"`
	//TradeRequestType is a required field for TradeCaptureReportRequest.
	TradeRequestType int `fix:"569"`
	//SubscriptionRequestType is a non-required field for TradeCaptureReportRequest.
	SubscriptionRequestType *string `fix:"263"`
	//TradeReportID is a non-required field for TradeCaptureReportRequest.
	TradeReportID *string `fix:"571"`
	//SecondaryTradeReportID is a non-required field for TradeCaptureReportRequest.
	SecondaryTradeReportID *string `fix:"818"`
	//ExecID is a non-required field for TradeCaptureReportRequest.
	ExecID *string `fix:"17"`
	//ExecType is a non-required field for TradeCaptureReportRequest.
	ExecType *string `fix:"150"`
	//OrderID is a non-required field for TradeCaptureReportRequest.
	OrderID *string `fix:"37"`
	//ClOrdID is a non-required field for TradeCaptureReportRequest.
	ClOrdID *string `fix:"11"`
	//MatchStatus is a non-required field for TradeCaptureReportRequest.
	MatchStatus *string `fix:"573"`
	//TrdType is a non-required field for TradeCaptureReportRequest.
	TrdType *int `fix:"828"`
	//TrdSubType is a non-required field for TradeCaptureReportRequest.
	TrdSubType *int `fix:"829"`
	//TransferReason is a non-required field for TradeCaptureReportRequest.
	TransferReason *string `fix:"830"`
	//SecondaryTrdType is a non-required field for TradeCaptureReportRequest.
	SecondaryTrdType *int `fix:"855"`
	//TradeLinkID is a non-required field for TradeCaptureReportRequest.
	TradeLinkID *string `fix:"820"`
	//TrdMatchID is a non-required field for TradeCaptureReportRequest.
	TrdMatchID *string `fix:"880"`
	//Parties Component
	Parties parties.Component
	//Instrument Component
	Instrument instrument.Component
	//InstrumentExtension Component
	InstrumentExtension instrumentextension.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//TrdCapDtGrp Component
	TrdCapDtGrp trdcapdtgrp.Component
	//ClearingBusinessDate is a non-required field for TradeCaptureReportRequest.
	ClearingBusinessDate *string `fix:"715"`
	//TradingSessionID is a non-required field for TradeCaptureReportRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for TradeCaptureReportRequest.
	TradingSessionSubID *string `fix:"625"`
	//TimeBracket is a non-required field for TradeCaptureReportRequest.
	TimeBracket *string `fix:"943"`
	//Side is a non-required field for TradeCaptureReportRequest.
	Side *string `fix:"54"`
	//MultiLegReportingType is a non-required field for TradeCaptureReportRequest.
	MultiLegReportingType *string `fix:"442"`
	//TradeInputSource is a non-required field for TradeCaptureReportRequest.
	TradeInputSource *string `fix:"578"`
	//TradeInputDevice is a non-required field for TradeCaptureReportRequest.
	TradeInputDevice *string `fix:"579"`
	//ResponseTransportType is a non-required field for TradeCaptureReportRequest.
	ResponseTransportType *int `fix:"725"`
	//ResponseDestination is a non-required field for TradeCaptureReportRequest.
	ResponseDestination *string `fix:"726"`
	//Text is a non-required field for TradeCaptureReportRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for TradeCaptureReportRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for TradeCaptureReportRequest.
	EncodedText *string `fix:"355"`
	//MessageEventSource is a non-required field for TradeCaptureReportRequest.
	MessageEventSource *string `fix:"1011"`
	//TradeID is a non-required field for TradeCaptureReportRequest.
	TradeID *string `fix:"1003"`
	//SecondaryTradeID is a non-required field for TradeCaptureReportRequest.
	SecondaryTradeID *string `fix:"1040"`
	//FirmTradeID is a non-required field for TradeCaptureReportRequest.
	FirmTradeID *string `fix:"1041"`
	//SecondaryFirmTradeID is a non-required field for TradeCaptureReportRequest.
	SecondaryFirmTradeID *string `fix:"1042"`
	//TradeHandlingInstr is a non-required field for TradeCaptureReportRequest.
	TradeHandlingInstr *string `fix:"1123"`
	Trailer            fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

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
	return enum.BeginStringFIX50, "AD", r
}
