//Package securitystatus msg type = f.
package securitystatus

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a SecurityStatus FIX Message
type Message struct {
	FIXMsgType string `fix:"f"`
	fixt11.Header
	//SecurityStatusReqID is a non-required field for SecurityStatus.
	SecurityStatusReqID *string `fix:"324"`
	//Instrument is a required component for SecurityStatus.
	instrument.Instrument
	//InstrumentExtension is a non-required component for SecurityStatus.
	InstrumentExtension *instrumentextension.InstrumentExtension
	//UndInstrmtGrp is a non-required component for SecurityStatus.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//InstrmtLegGrp is a non-required component for SecurityStatus.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//Currency is a non-required field for SecurityStatus.
	Currency *string `fix:"15"`
	//TradingSessionID is a non-required field for SecurityStatus.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for SecurityStatus.
	TradingSessionSubID *string `fix:"625"`
	//UnsolicitedIndicator is a non-required field for SecurityStatus.
	UnsolicitedIndicator *bool `fix:"325"`
	//SecurityTradingStatus is a non-required field for SecurityStatus.
	SecurityTradingStatus *int `fix:"326"`
	//FinancialStatus is a non-required field for SecurityStatus.
	FinancialStatus *string `fix:"291"`
	//CorporateAction is a non-required field for SecurityStatus.
	CorporateAction *string `fix:"292"`
	//HaltReasonChar is a non-required field for SecurityStatus.
	HaltReasonChar *string `fix:"327"`
	//InViewOfCommon is a non-required field for SecurityStatus.
	InViewOfCommon *bool `fix:"328"`
	//DueToRelated is a non-required field for SecurityStatus.
	DueToRelated *bool `fix:"329"`
	//BuyVolume is a non-required field for SecurityStatus.
	BuyVolume *float64 `fix:"330"`
	//SellVolume is a non-required field for SecurityStatus.
	SellVolume *float64 `fix:"331"`
	//HighPx is a non-required field for SecurityStatus.
	HighPx *float64 `fix:"332"`
	//LowPx is a non-required field for SecurityStatus.
	LowPx *float64 `fix:"333"`
	//LastPx is a non-required field for SecurityStatus.
	LastPx *float64 `fix:"31"`
	//TransactTime is a non-required field for SecurityStatus.
	TransactTime *time.Time `fix:"60"`
	//Adjustment is a non-required field for SecurityStatus.
	Adjustment *int `fix:"334"`
	//Text is a non-required field for SecurityStatus.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SecurityStatus.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SecurityStatus.
	EncodedText *string `fix:"355"`
	//FirstPx is a non-required field for SecurityStatus.
	FirstPx *float64 `fix:"1025"`
	//SecurityTradingEvent is a non-required field for SecurityStatus.
	SecurityTradingEvent *int `fix:"1174"`
	//MDBookType is a non-required field for SecurityStatus.
	MDBookType *int `fix:"1021"`
	//MarketDepth is a non-required field for SecurityStatus.
	MarketDepth *int `fix:"264"`
	//MarketID is a non-required field for SecurityStatus.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for SecurityStatus.
	MarketSegmentID *string `fix:"1300"`
	//ApplicationSequenceControl is a non-required component for SecurityStatus.
	ApplicationSequenceControl *applicationsequencecontrol.ApplicationSequenceControl
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized SecurityStatus instance
func New(instrument instrument.Instrument) *Message {
	var m Message
	m.SetInstrument(instrument)
	return &m
}

func (m *Message) SetSecurityStatusReqID(v string)       { m.SecurityStatusReqID = &v }
func (m *Message) SetInstrument(v instrument.Instrument) { m.Instrument = v }
func (m *Message) SetInstrumentExtension(v instrumentextension.InstrumentExtension) {
	m.InstrumentExtension = &v
}
func (m *Message) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp) { m.UndInstrmtGrp = &v }
func (m *Message) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp) { m.InstrmtLegGrp = &v }
func (m *Message) SetCurrency(v string)                           { m.Currency = &v }
func (m *Message) SetTradingSessionID(v string)                   { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)                { m.TradingSessionSubID = &v }
func (m *Message) SetUnsolicitedIndicator(v bool)                 { m.UnsolicitedIndicator = &v }
func (m *Message) SetSecurityTradingStatus(v int)                 { m.SecurityTradingStatus = &v }
func (m *Message) SetFinancialStatus(v string)                    { m.FinancialStatus = &v }
func (m *Message) SetCorporateAction(v string)                    { m.CorporateAction = &v }
func (m *Message) SetHaltReasonChar(v string)                     { m.HaltReasonChar = &v }
func (m *Message) SetInViewOfCommon(v bool)                       { m.InViewOfCommon = &v }
func (m *Message) SetDueToRelated(v bool)                         { m.DueToRelated = &v }
func (m *Message) SetBuyVolume(v float64)                         { m.BuyVolume = &v }
func (m *Message) SetSellVolume(v float64)                        { m.SellVolume = &v }
func (m *Message) SetHighPx(v float64)                            { m.HighPx = &v }
func (m *Message) SetLowPx(v float64)                             { m.LowPx = &v }
func (m *Message) SetLastPx(v float64)                            { m.LastPx = &v }
func (m *Message) SetTransactTime(v time.Time)                    { m.TransactTime = &v }
func (m *Message) SetAdjustment(v int)                            { m.Adjustment = &v }
func (m *Message) SetText(v string)                               { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                        { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                        { m.EncodedText = &v }
func (m *Message) SetFirstPx(v float64)                           { m.FirstPx = &v }
func (m *Message) SetSecurityTradingEvent(v int)                  { m.SecurityTradingEvent = &v }
func (m *Message) SetMDBookType(v int)                            { m.MDBookType = &v }
func (m *Message) SetMarketDepth(v int)                           { m.MarketDepth = &v }
func (m *Message) SetMarketID(v string)                           { m.MarketID = &v }
func (m *Message) SetMarketSegmentID(v string)                    { m.MarketSegmentID = &v }
func (m *Message) SetApplicationSequenceControl(v applicationsequencecontrol.ApplicationSequenceControl) {
	m.ApplicationSequenceControl = &v
}

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
	return enum.ApplVerID_FIX50SP1, "f", r
}
