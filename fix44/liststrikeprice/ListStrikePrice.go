//Package liststrikeprice msg type = m.
package liststrikeprice

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
)

//NoStrikes is a repeating group in ListStrikePrice
type NoStrikes struct {
	//Instrument Component
	Instrument instrument.Component
}

//NoUnderlyings is a repeating group in ListStrikePrice
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
	//PrevClosePx is a non-required field for NoUnderlyings.
	PrevClosePx *float64 `fix:"140"`
	//ClOrdID is a non-required field for NoUnderlyings.
	ClOrdID *string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NoUnderlyings.
	SecondaryClOrdID *string `fix:"526"`
	//Side is a non-required field for NoUnderlyings.
	Side *string `fix:"54"`
	//Price is a required field for NoUnderlyings.
	Price float64 `fix:"44"`
	//Currency is a non-required field for NoUnderlyings.
	Currency *string `fix:"15"`
	//Text is a non-required field for NoUnderlyings.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoUnderlyings.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoUnderlyings.
	EncodedText *string `fix:"355"`
}

//Message is a ListStrikePrice FIX Message
type Message struct {
	FIXMsgType string `fix:"m"`
	Header     fix44.Header
	//ListID is a required field for ListStrikePrice.
	ListID string `fix:"66"`
	//TotNoStrikes is a required field for ListStrikePrice.
	TotNoStrikes int `fix:"422"`
	//LastFragment is a non-required field for ListStrikePrice.
	LastFragment *bool `fix:"893"`
	//NoStrikes is a required field for ListStrikePrice.
	NoStrikes []NoStrikes `fix:"428"`
	//NoUnderlyings is a non-required field for ListStrikePrice.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	Trailer       fix44.Trailer
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
	return enum.BeginStringFIX44, "m", r
}
