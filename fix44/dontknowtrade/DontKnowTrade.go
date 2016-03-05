//Package dontknowtrade msg type = Q.
package dontknowtrade

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/orderqtydata"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
)

//NoUnderlyings is a repeating group in DontKnowTrade
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
}

//NoLegs is a repeating group in DontKnowTrade
type NoLegs struct {
	//InstrumentLeg Component
	instrumentleg.InstrumentLeg
}

//Message is a DontKnowTrade FIX Message
type Message struct {
	FIXMsgType string `fix:"Q"`
	fix44.Header
	//OrderID is a required field for DontKnowTrade.
	OrderID string `fix:"37"`
	//SecondaryOrderID is a non-required field for DontKnowTrade.
	SecondaryOrderID *string `fix:"198"`
	//ExecID is a required field for DontKnowTrade.
	ExecID string `fix:"17"`
	//DKReason is a required field for DontKnowTrade.
	DKReason string `fix:"127"`
	//Instrument Component
	instrument.Instrument
	//NoUnderlyings is a non-required field for DontKnowTrade.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//NoLegs is a non-required field for DontKnowTrade.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//Side is a required field for DontKnowTrade.
	Side string `fix:"54"`
	//OrderQtyData Component
	orderqtydata.OrderQtyData
	//LastQty is a non-required field for DontKnowTrade.
	LastQty *float64 `fix:"32"`
	//LastPx is a non-required field for DontKnowTrade.
	LastPx *float64 `fix:"31"`
	//Text is a non-required field for DontKnowTrade.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for DontKnowTrade.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for DontKnowTrade.
	EncodedText *string `fix:"355"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrderID(v string)                { m.OrderID = v }
func (m *Message) SetSecondaryOrderID(v string)       { m.SecondaryOrderID = &v }
func (m *Message) SetExecID(v string)                 { m.ExecID = v }
func (m *Message) SetDKReason(v string)               { m.DKReason = v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }
func (m *Message) SetNoLegs(v []NoLegs)               { m.NoLegs = v }
func (m *Message) SetSide(v string)                   { m.Side = v }
func (m *Message) SetLastQty(v float64)               { m.LastQty = &v }
func (m *Message) SetLastPx(v float64)                { m.LastPx = &v }
func (m *Message) SetText(v string)                   { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)            { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)            { m.EncodedText = &v }

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
	return enum.BeginStringFIX44, "Q", r
}
