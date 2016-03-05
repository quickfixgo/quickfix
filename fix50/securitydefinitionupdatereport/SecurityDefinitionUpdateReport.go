//Package securitydefinitionupdatereport msg type = BP.
package securitydefinitionupdatereport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a SecurityDefinitionUpdateReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BP"`
	fixt11.Header
	//SecurityReportID is a non-required field for SecurityDefinitionUpdateReport.
	SecurityReportID *int `fix:"964"`
	//SecurityReqID is a non-required field for SecurityDefinitionUpdateReport.
	SecurityReqID *string `fix:"320"`
	//SecurityResponseID is a non-required field for SecurityDefinitionUpdateReport.
	SecurityResponseID *string `fix:"322"`
	//SecurityResponseType is a non-required field for SecurityDefinitionUpdateReport.
	SecurityResponseType *int `fix:"323"`
	//ClearingBusinessDate is a non-required field for SecurityDefinitionUpdateReport.
	ClearingBusinessDate *string `fix:"715"`
	//SecurityUpdateAction is a non-required field for SecurityDefinitionUpdateReport.
	SecurityUpdateAction *string `fix:"980"`
	//CorporateAction is a non-required field for SecurityDefinitionUpdateReport.
	CorporateAction *string `fix:"292"`
	//Instrument Component
	instrument.Instrument
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
	//Currency is a non-required field for SecurityDefinitionUpdateReport.
	Currency *string `fix:"15"`
	//TradingSessionID is a non-required field for SecurityDefinitionUpdateReport.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for SecurityDefinitionUpdateReport.
	TradingSessionSubID *string `fix:"625"`
	//Text is a non-required field for SecurityDefinitionUpdateReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SecurityDefinitionUpdateReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SecurityDefinitionUpdateReport.
	EncodedText *string `fix:"355"`
	//InstrmtLegGrp Component
	instrmtleggrp.InstrmtLegGrp
	//ExpirationCycle is a non-required field for SecurityDefinitionUpdateReport.
	ExpirationCycle *int `fix:"827"`
	//RoundLot is a non-required field for SecurityDefinitionUpdateReport.
	RoundLot *float64 `fix:"561"`
	//MinTradeVol is a non-required field for SecurityDefinitionUpdateReport.
	MinTradeVol *float64 `fix:"562"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetSecurityReportID(v int)        { m.SecurityReportID = &v }
func (m *Message) SetSecurityReqID(v string)        { m.SecurityReqID = &v }
func (m *Message) SetSecurityResponseID(v string)   { m.SecurityResponseID = &v }
func (m *Message) SetSecurityResponseType(v int)    { m.SecurityResponseType = &v }
func (m *Message) SetClearingBusinessDate(v string) { m.ClearingBusinessDate = &v }
func (m *Message) SetSecurityUpdateAction(v string) { m.SecurityUpdateAction = &v }
func (m *Message) SetCorporateAction(v string)      { m.CorporateAction = &v }
func (m *Message) SetCurrency(v string)             { m.Currency = &v }
func (m *Message) SetTradingSessionID(v string)     { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)  { m.TradingSessionSubID = &v }
func (m *Message) SetText(v string)                 { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)          { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)          { m.EncodedText = &v }
func (m *Message) SetExpirationCycle(v int)         { m.ExpirationCycle = &v }
func (m *Message) SetRoundLot(v float64)            { m.RoundLot = &v }
func (m *Message) SetMinTradeVol(v float64)         { m.MinTradeVol = &v }

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
	return enum.BeginStringFIX50, "BP", r
}
