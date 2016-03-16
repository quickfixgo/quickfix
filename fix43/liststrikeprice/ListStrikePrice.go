//Package liststrikeprice msg type = m.
package liststrikeprice

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
)

//NoStrikes is a repeating group in ListStrikePrice
type NoStrikes struct {
	//Instrument is a required component for NoStrikes.
	instrument.Instrument
	//PrevClosePx is a non-required field for NoStrikes.
	PrevClosePx *float64 `fix:"140"`
	//ClOrdID is a non-required field for NoStrikes.
	ClOrdID *string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NoStrikes.
	SecondaryClOrdID *string `fix:"526"`
	//Side is a non-required field for NoStrikes.
	Side *string `fix:"54"`
	//Price is a required field for NoStrikes.
	Price float64 `fix:"44"`
	//Currency is a non-required field for NoStrikes.
	Currency *string `fix:"15"`
	//Text is a non-required field for NoStrikes.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoStrikes.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoStrikes.
	EncodedText *string `fix:"355"`
}

func (m *NoStrikes) SetInstrument(v instrument.Instrument) { m.Instrument = v }
func (m *NoStrikes) SetPrevClosePx(v float64)              { m.PrevClosePx = &v }
func (m *NoStrikes) SetClOrdID(v string)                   { m.ClOrdID = &v }
func (m *NoStrikes) SetSecondaryClOrdID(v string)          { m.SecondaryClOrdID = &v }
func (m *NoStrikes) SetSide(v string)                      { m.Side = &v }
func (m *NoStrikes) SetPrice(v float64)                    { m.Price = v }
func (m *NoStrikes) SetCurrency(v string)                  { m.Currency = &v }
func (m *NoStrikes) SetText(v string)                      { m.Text = &v }
func (m *NoStrikes) SetEncodedTextLen(v int)               { m.EncodedTextLen = &v }
func (m *NoStrikes) SetEncodedText(v string)               { m.EncodedText = &v }

//Message is a ListStrikePrice FIX Message
type Message struct {
	FIXMsgType string `fix:"m"`
	fix43.Header
	//ListID is a required field for ListStrikePrice.
	ListID string `fix:"66"`
	//TotNoStrikes is a required field for ListStrikePrice.
	TotNoStrikes int `fix:"422"`
	//NoStrikes is a required field for ListStrikePrice.
	NoStrikes []NoStrikes `fix:"428"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetListID(v string)         { m.ListID = v }
func (m *Message) SetTotNoStrikes(v int)      { m.TotNoStrikes = v }
func (m *Message) SetNoStrikes(v []NoStrikes) { m.NoStrikes = v }

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
	return enum.BeginStringFIX43, "m", r
}
