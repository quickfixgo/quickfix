//Package derivativesecuritylist msg type = AA.
package derivativesecuritylist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentextension"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
)

//NoRelatedSym is a repeating group in DerivativeSecurityList
type NoRelatedSym struct {
	//Instrument is a non-required component for NoRelatedSym.
	Instrument *instrument.Instrument
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//ExpirationCycle is a non-required field for NoRelatedSym.
	ExpirationCycle *int `fix:"827"`
	//InstrumentExtension is a non-required component for NoRelatedSym.
	InstrumentExtension *instrumentextension.InstrumentExtension
	//NoLegs is a non-required field for NoRelatedSym.
	NoLegs []NoLegs `fix:"555,omitempty"`
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

func (m *NoRelatedSym) SetInstrument(v instrument.Instrument) { m.Instrument = &v }
func (m *NoRelatedSym) SetCurrency(v string)                  { m.Currency = &v }
func (m *NoRelatedSym) SetExpirationCycle(v int)              { m.ExpirationCycle = &v }
func (m *NoRelatedSym) SetInstrumentExtension(v instrumentextension.InstrumentExtension) {
	m.InstrumentExtension = &v
}
func (m *NoRelatedSym) SetNoLegs(v []NoLegs)            { m.NoLegs = v }
func (m *NoRelatedSym) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *NoRelatedSym) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }
func (m *NoRelatedSym) SetText(v string)                { m.Text = &v }
func (m *NoRelatedSym) SetEncodedTextLen(v int)         { m.EncodedTextLen = &v }
func (m *NoRelatedSym) SetEncodedText(v string)         { m.EncodedText = &v }

//NoLegs is a repeating group in NoRelatedSym
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
}

func (m *NoLegs) SetInstrumentLeg(v instrumentleg.InstrumentLeg) { m.InstrumentLeg = &v }

//Message is a DerivativeSecurityList FIX Message
type Message struct {
	FIXMsgType string `fix:"AA"`
	fix44.Header
	//SecurityReqID is a required field for DerivativeSecurityList.
	SecurityReqID string `fix:"320"`
	//SecurityResponseID is a required field for DerivativeSecurityList.
	SecurityResponseID string `fix:"322"`
	//SecurityRequestResult is a required field for DerivativeSecurityList.
	SecurityRequestResult int `fix:"560"`
	//UnderlyingInstrument is a non-required component for DerivativeSecurityList.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
	//TotNoRelatedSym is a non-required field for DerivativeSecurityList.
	TotNoRelatedSym *int `fix:"393"`
	//LastFragment is a non-required field for DerivativeSecurityList.
	LastFragment *bool `fix:"893"`
	//NoRelatedSym is a non-required field for DerivativeSecurityList.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetSecurityReqID(v string)      { m.SecurityReqID = v }
func (m *Message) SetSecurityResponseID(v string) { m.SecurityResponseID = v }
func (m *Message) SetSecurityRequestResult(v int) { m.SecurityRequestResult = v }
func (m *Message) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
}
func (m *Message) SetTotNoRelatedSym(v int)         { m.TotNoRelatedSym = &v }
func (m *Message) SetLastFragment(v bool)           { m.LastFragment = &v }
func (m *Message) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }

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
	return enum.BeginStringFIX44, "AA", r
}
