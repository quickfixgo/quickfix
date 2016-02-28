//Package dontknowtrade msg type = Q.
package dontknowtrade

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix40"
)

//Message is a DontKnowTrade FIX Message
type Message struct {
	FIXMsgType string `fix:"Q"`
	Header     fix40.Header
	//OrderID is a non-required field for DontKnowTrade.
	OrderID *string `fix:"37"`
	//ExecID is a non-required field for DontKnowTrade.
	ExecID *int `fix:"17"`
	//DKReason is a required field for DontKnowTrade.
	DKReason string `fix:"127"`
	//Symbol is a required field for DontKnowTrade.
	Symbol string `fix:"55"`
	//Side is a required field for DontKnowTrade.
	Side string `fix:"54"`
	//OrderQty is a required field for DontKnowTrade.
	OrderQty int `fix:"38"`
	//LastShares is a required field for DontKnowTrade.
	LastShares int `fix:"32"`
	//LastPx is a required field for DontKnowTrade.
	LastPx float64 `fix:"31"`
	//Text is a non-required field for DontKnowTrade.
	Text    *string `fix:"58"`
	Trailer fix40.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrderID(v string)  { m.OrderID = &v }
func (m *Message) SetExecID(v int)      { m.ExecID = &v }
func (m *Message) SetDKReason(v string) { m.DKReason = v }
func (m *Message) SetSymbol(v string)   { m.Symbol = v }
func (m *Message) SetSide(v string)     { m.Side = v }
func (m *Message) SetOrderQty(v int)    { m.OrderQty = v }
func (m *Message) SetLastShares(v int)  { m.LastShares = v }
func (m *Message) SetLastPx(v float64)  { m.LastPx = v }
func (m *Message) SetText(v string)     { m.Text = &v }

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
	return enum.BeginStringFIX40, "Q", r
}
