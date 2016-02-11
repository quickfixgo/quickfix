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

//Message is a SecurityDefinition FIX Message
type Message struct {
	FIXMsgType string `fix:"d"`
	Header     fixt11.Header
	//SecurityReqID is a non-required field for SecurityDefinition.
	SecurityReqID *string `fix:"320"`
	//SecurityResponseID is a non-required field for SecurityDefinition.
	SecurityResponseID *string `fix:"322"`
	//SecurityResponseType is a non-required field for SecurityDefinition.
	SecurityResponseType *int `fix:"323"`
	//Instrument Component
	Instrument instrument.Component
	//InstrumentExtension Component
	InstrumentExtension instrumentextension.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
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
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
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
	Trailer              fixt11.Trailer
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
	return enum.BeginStringFIX50, "d", r
}
