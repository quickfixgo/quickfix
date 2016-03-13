//Package securitydefinition msg type = d.
package securitydefinition

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//ClearingBusinessDate is a repeating group in SecurityDefinition
type ClearingBusinessDate struct {
	//LegOptionRatio is a non-required field for ClearingBusinessDate.
	LegOptionRatio *float64 `fix:"1017"`
	//LegPrice is a non-required field for ClearingBusinessDate.
	LegPrice *float64 `fix:"566"`
}

func (m *ClearingBusinessDate) SetLegOptionRatio(v float64) { m.LegOptionRatio = &v }
func (m *ClearingBusinessDate) SetLegPrice(v float64)       { m.LegPrice = &v }

//Message is a SecurityDefinition FIX Message
type Message struct {
	FIXMsgType string `fix:"d"`
	fixt11.Header
	//SecurityReqID is a non-required field for SecurityDefinition.
	SecurityReqID *string `fix:"320"`
	//SecurityResponseID is a non-required field for SecurityDefinition.
	SecurityResponseID *string `fix:"322"`
	//SecurityResponseType is a non-required field for SecurityDefinition.
	SecurityResponseType *int `fix:"323"`
	//Instrument is a non-required component for SecurityDefinition.
	Instrument *instrument.Instrument
	//InstrumentExtension is a non-required component for SecurityDefinition.
	InstrumentExtension *instrumentextension.InstrumentExtension
	//UndInstrmtGrp is a non-required component for SecurityDefinition.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//Currency is a non-required field for SecurityDefinition.
	Currency *string `fix:"15"`
	//TradingSessionID is a non-required field for SecurityDefinition.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for SecurityDefinition.
	TradingSessionSubID *string `fix:"625"`
	//Text is a non-required field for SecurityDefinition.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SecurityDefinition.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SecurityDefinition.
	EncodedText *string `fix:"355"`
	//InstrmtLegGrp is a non-required component for SecurityDefinition.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//ExpirationCycle is a non-required field for SecurityDefinition.
	ExpirationCycle *int `fix:"827"`
	//RoundLot is a non-required field for SecurityDefinition.
	RoundLot *float64 `fix:"561"`
	//MinTradeVol is a non-required field for SecurityDefinition.
	MinTradeVol *float64 `fix:"562"`
	//SecurityReportID is a non-required field for SecurityDefinition.
	SecurityReportID *int `fix:"964"`
	//ClearingBusinessDate is a non-required field for SecurityDefinition.
	ClearingBusinessDate []ClearingBusinessDate `fix:"715,omitempty"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetSecurityReqID(v string)             { m.SecurityReqID = &v }
func (m *Message) SetSecurityResponseID(v string)        { m.SecurityResponseID = &v }
func (m *Message) SetSecurityResponseType(v int)         { m.SecurityResponseType = &v }
func (m *Message) SetInstrument(v instrument.Instrument) { m.Instrument = &v }
func (m *Message) SetInstrumentExtension(v instrumentextension.InstrumentExtension) {
	m.InstrumentExtension = &v
}
func (m *Message) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp)   { m.UndInstrmtGrp = &v }
func (m *Message) SetCurrency(v string)                             { m.Currency = &v }
func (m *Message) SetTradingSessionID(v string)                     { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)                  { m.TradingSessionSubID = &v }
func (m *Message) SetText(v string)                                 { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                          { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                          { m.EncodedText = &v }
func (m *Message) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp)   { m.InstrmtLegGrp = &v }
func (m *Message) SetExpirationCycle(v int)                         { m.ExpirationCycle = &v }
func (m *Message) SetRoundLot(v float64)                            { m.RoundLot = &v }
func (m *Message) SetMinTradeVol(v float64)                         { m.MinTradeVol = &v }
func (m *Message) SetSecurityReportID(v int)                        { m.SecurityReportID = &v }
func (m *Message) SetClearingBusinessDate(v []ClearingBusinessDate) { m.ClearingBusinessDate = v }

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
	return enum.BeginStringFIX50, "d", r
}
