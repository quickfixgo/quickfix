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
	Header     fixt11.Header
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
	Instrument instrument.Component
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
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
	InstrmtLegGrp instrmtleggrp.Component
	//ExpirationCycle is a non-required field for SecurityDefinitionUpdateReport.
	ExpirationCycle *int `fix:"827"`
	//RoundLot is a non-required field for SecurityDefinitionUpdateReport.
	RoundLot *float64 `fix:"561"`
	//MinTradeVol is a non-required field for SecurityDefinitionUpdateReport.
	MinTradeVol *float64 `fix:"562"`
	Trailer     fixt11.Trailer
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
	return enum.BeginStringFIX50, "BP", r
}
