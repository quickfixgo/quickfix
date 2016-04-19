//Package tradecapturereportrequest msg type = AD.
package tradecapturereportrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/trdcapdtgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a TradeCaptureReportRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"AD"`
	fixt11.Header
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
	//Parties is a non-required component for TradeCaptureReportRequest.
	Parties *parties.Parties
	//Instrument is a non-required component for TradeCaptureReportRequest.
	Instrument *instrument.Instrument
	//InstrumentExtension is a non-required component for TradeCaptureReportRequest.
	InstrumentExtension *instrumentextension.InstrumentExtension
	//FinancingDetails is a non-required component for TradeCaptureReportRequest.
	FinancingDetails *financingdetails.FinancingDetails
	//UndInstrmtGrp is a non-required component for TradeCaptureReportRequest.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//InstrmtLegGrp is a non-required component for TradeCaptureReportRequest.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//TrdCapDtGrp is a non-required component for TradeCaptureReportRequest.
	TrdCapDtGrp *trdcapdtgrp.TrdCapDtGrp
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
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized TradeCaptureReportRequest instance
func New(traderequestid string, traderequesttype int) *Message {
	var m Message
	m.SetTradeRequestID(traderequestid)
	m.SetTradeRequestType(traderequesttype)
	return &m
}

func (m *Message) SetTradeRequestID(v string)            { m.TradeRequestID = v }
func (m *Message) SetTradeRequestType(v int)             { m.TradeRequestType = v }
func (m *Message) SetSubscriptionRequestType(v string)   { m.SubscriptionRequestType = &v }
func (m *Message) SetTradeReportID(v string)             { m.TradeReportID = &v }
func (m *Message) SetSecondaryTradeReportID(v string)    { m.SecondaryTradeReportID = &v }
func (m *Message) SetExecID(v string)                    { m.ExecID = &v }
func (m *Message) SetExecType(v string)                  { m.ExecType = &v }
func (m *Message) SetOrderID(v string)                   { m.OrderID = &v }
func (m *Message) SetClOrdID(v string)                   { m.ClOrdID = &v }
func (m *Message) SetMatchStatus(v string)               { m.MatchStatus = &v }
func (m *Message) SetTrdType(v int)                      { m.TrdType = &v }
func (m *Message) SetTrdSubType(v int)                   { m.TrdSubType = &v }
func (m *Message) SetTransferReason(v string)            { m.TransferReason = &v }
func (m *Message) SetSecondaryTrdType(v int)             { m.SecondaryTrdType = &v }
func (m *Message) SetTradeLinkID(v string)               { m.TradeLinkID = &v }
func (m *Message) SetTrdMatchID(v string)                { m.TrdMatchID = &v }
func (m *Message) SetParties(v parties.Parties)          { m.Parties = &v }
func (m *Message) SetInstrument(v instrument.Instrument) { m.Instrument = &v }
func (m *Message) SetInstrumentExtension(v instrumentextension.InstrumentExtension) {
	m.InstrumentExtension = &v
}
func (m *Message) SetFinancingDetails(v financingdetails.FinancingDetails) { m.FinancingDetails = &v }
func (m *Message) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp)          { m.UndInstrmtGrp = &v }
func (m *Message) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp)          { m.InstrmtLegGrp = &v }
func (m *Message) SetTrdCapDtGrp(v trdcapdtgrp.TrdCapDtGrp)                { m.TrdCapDtGrp = &v }
func (m *Message) SetClearingBusinessDate(v string)                        { m.ClearingBusinessDate = &v }
func (m *Message) SetTradingSessionID(v string)                            { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)                         { m.TradingSessionSubID = &v }
func (m *Message) SetTimeBracket(v string)                                 { m.TimeBracket = &v }
func (m *Message) SetSide(v string)                                        { m.Side = &v }
func (m *Message) SetMultiLegReportingType(v string)                       { m.MultiLegReportingType = &v }
func (m *Message) SetTradeInputSource(v string)                            { m.TradeInputSource = &v }
func (m *Message) SetTradeInputDevice(v string)                            { m.TradeInputDevice = &v }
func (m *Message) SetResponseTransportType(v int)                          { m.ResponseTransportType = &v }
func (m *Message) SetResponseDestination(v string)                         { m.ResponseDestination = &v }
func (m *Message) SetText(v string)                                        { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                                 { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                                 { m.EncodedText = &v }
func (m *Message) SetMessageEventSource(v string)                          { m.MessageEventSource = &v }
func (m *Message) SetTradeID(v string)                                     { m.TradeID = &v }
func (m *Message) SetSecondaryTradeID(v string)                            { m.SecondaryTradeID = &v }
func (m *Message) SetFirmTradeID(v string)                                 { m.FirmTradeID = &v }
func (m *Message) SetSecondaryFirmTradeID(v string)                        { m.SecondaryFirmTradeID = &v }
func (m *Message) SetTradeHandlingInstr(v string)                          { m.TradeHandlingInstr = &v }

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
	return enum.ApplVerID_FIX50SP2, "AD", r
}
