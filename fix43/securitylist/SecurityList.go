//Package securitylist msg type = y.
package securitylist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/instrumentleg"
)

//NoRelatedSym is a repeating group in SecurityList
type NoRelatedSym struct {
	//Instrument Component
	Instrument instrument.Component
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//NoLegs is a non-required field for NoRelatedSym.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//RoundLot is a non-required field for NoRelatedSym.
	RoundLot *float64 `fix:"561"`
	//MinTradeVol is a non-required field for NoRelatedSym.
	MinTradeVol *float64 `fix:"562"`
	//TradingSessionID is a non-required field for NoRelatedSym.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoRelatedSym.
	TradingSessionSubID *string `fix:"625"`
	//Text is a non-required field for NoRelatedSym.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoRelatedSym.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoRelatedSym.
	EncodedText *string `fix:"355"`
}

func (m *NoRelatedSym) SetCurrency(v string)            { m.Currency = &v }
func (m *NoRelatedSym) SetNoLegs(v []NoLegs)            { m.NoLegs = v }
func (m *NoRelatedSym) SetRoundLot(v float64)           { m.RoundLot = &v }
func (m *NoRelatedSym) SetMinTradeVol(v float64)        { m.MinTradeVol = &v }
func (m *NoRelatedSym) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *NoRelatedSym) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }
func (m *NoRelatedSym) SetText(v string)                { m.Text = &v }
func (m *NoRelatedSym) SetEncodedTextLen(v int)         { m.EncodedTextLen = &v }
func (m *NoRelatedSym) SetEncodedText(v string)         { m.EncodedText = &v }

//NoLegs is a repeating group in NoRelatedSym
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
	//LegCurrency is a non-required field for NoLegs.
	LegCurrency *string `fix:"556"`
}

func (m *NoLegs) SetLegCurrency(v string) { m.LegCurrency = &v }

//Message is a SecurityList FIX Message
type Message struct {
	FIXMsgType string `fix:"y"`
	Header     fix43.Header
	//SecurityReqID is a required field for SecurityList.
	SecurityReqID string `fix:"320"`
	//SecurityResponseID is a required field for SecurityList.
	SecurityResponseID string `fix:"322"`
	//SecurityRequestResult is a required field for SecurityList.
	SecurityRequestResult int `fix:"560"`
	//TotalNumSecurities is a non-required field for SecurityList.
	TotalNumSecurities *int `fix:"393"`
	//NoRelatedSym is a non-required field for SecurityList.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
	Trailer      fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetSecurityReqID(v string)        { m.SecurityReqID = v }
func (m *Message) SetSecurityResponseID(v string)   { m.SecurityResponseID = v }
func (m *Message) SetSecurityRequestResult(v int)   { m.SecurityRequestResult = v }
func (m *Message) SetTotalNumSecurities(v int)      { m.TotalNumSecurities = &v }
func (m *Message) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }

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
	return enum.BeginStringFIX43, "y", r
}
